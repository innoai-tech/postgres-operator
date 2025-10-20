package db

import (
	"cmp"
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/go-json-experiment/json"
	databasev1 "github.com/innoai-tech/postgres-operator/pkg/apis/database/v1"
	"github.com/jackc/pgx/v5"
	metav1 "github.com/octohelm/objectkind/pkg/apis/meta/v1"
	"github.com/octohelm/objectkind/pkg/runtime"
	"github.com/octohelm/storage/pkg/session"
	"github.com/octohelm/storage/pkg/sqlbuilder"
	"github.com/octohelm/storage/pkg/sqlfrag"
	"github.com/octohelm/x/ptr"

	_ "github.com/octohelm/storage/pkg/session/db"
)

func New(dsn *url.URL) *Controller {
	return &Controller{
		dsn:    dsn,
		dbCode: databasev1.DatabaseCode(strings.TrimSuffix(dsn.Path, "/")),
	}
}

type Controller struct {
	dbCode databasev1.DatabaseCode
	dsn    *url.URL
}

func (c Controller) WithName(databaseCode databasev1.DatabaseCode) *Controller {
	c.dbCode = databaseCode
	c.dsn = ptr.Ptr(*c.dsn)
	c.dsn.Path = "/" + string(c.dbCode)
	return &c
}

func (c *Controller) QueryResult(ctx context.Context, sql string) (*databasev1.Result, error) {
	a, err := c.Open(ctx)
	if err != nil {
		return nil, err
	}
	defer a.Close()

	rows, err := a.Query(ctx, sqlfrag.Pair(sql))
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	ret := &databasev1.Result{}

	columns, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	ret.Columns = make([]*databasev1.ResultColumn, 0, len(columns))
	for _, column := range columns {
		ret.Columns = append(ret.Columns, &databasev1.ResultColumn{
			Code: databasev1.ColumnCode(column.Name()),
			Type: column.DatabaseTypeName(),
		})
	}

	count := 0

	for rows.Next() {
		data := make([]any, len(columns))
		for i := range data {
			data[i] = any(&recv{})
		}

		if err := rows.Scan(data...); err != nil {
			return nil, err
		}

		ret.Data = append(ret.Data, data)

		count++
		if count > 500 {
			break
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	return ret, nil
}

type recv struct {
	v any
}

var _ = sql.Scanner(&recv{})

func (r *recv) Scan(x any) error {
	r.v = x
	return nil
}

func (r recv) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.v)
}

func (c *Controller) ListDatabase(ctx context.Context) (*metav1.List[databasev1.Database], error) {
	a, err := c.WithName("postgres").Open(ctx)
	if err != nil {
		return nil, err
	}
	defer a.Close()

	databaseSchema := sqlbuilder.TableFromModel(&pgDatabase{})

	stmt := sqlbuilder.Select(sqlbuilder.ColumnCollect(databaseSchema.Cols())).From(databaseSchema,
		sqlbuilder.Where(
			sqlbuilder.And(
				sqlbuilder.TypedColOf[bool](databaseSchema, "datistemplate").V(sqlbuilder.Eq(false)),
			),
		),
	)

	rows, err := a.Query(ctx, stmt)
	if err != nil {
		return nil, err
	}

	list := &metav1.List[databasev1.Database]{}

	if err := session.Scan(ctx, rows, session.Recv(func(d *pgDatabase) error {
		list.Add(runtime.Build(func(o *databasev1.Database) {
			o.Code = databasev1.DatabaseCode(d.DatabaseName)
			o.Spec.CharacterType = d.CharacterType
			o.Spec.Collation = d.Collation
			o.Spec.CollationVersion = d.CollationVersion
		}))
		return nil
	})); err != nil {
		return nil, err
	}

	return list, nil
}

func (c *Controller) ListTableOfDatabase(ctx context.Context, databaseName databasev1.DatabaseCode) (*metav1.List[databasev1.Table], error) {
	a, err := c.WithName(databaseName).Open(ctx)
	if err != nil {
		return nil, err
	}

	cat, err := a.Catalog(ctx)
	if err != nil {
		return nil, err
	}

	tables := slices.SortedFunc(cat.Tables(), func(t1 sqlbuilder.Table, t2 sqlbuilder.Table) int {
		return cmp.Compare(t1.TableName(), t2.TableName())
	})

	list := &metav1.List[databasev1.Table]{}

	for _, t := range tables {
		list.Add(runtime.Build(func(o *databasev1.Table) {
			o.Code = databasev1.TableCode(t.TableName())

			cols := slices.SortedFunc(t.Cols(), func(c1 sqlbuilder.Column, c2 sqlbuilder.Column) int {
				return cmp.Compare(c1.Name(), c1.Name())
			})

			for _, col := range cols {
				tpe, _ := sqlfrag.Collect(ctx, a.Dialect().DataType(sqlbuilder.GetColumnDef(col)))

				o.Spec.Columns = append(o.Spec.Columns, runtime.Build(func(o *databasev1.Column) {
					o.Code = databasev1.ColumnCode(col.Name())
					o.Spec.Type = tpe
				}))
			}

			keys := slices.SortedFunc(t.Keys(), func(k1 sqlbuilder.Key, k2 sqlbuilder.Key) int {
				return cmp.Compare(k1.Name(), k2.Name())
			})

			for _, key := range keys {
				if key.Name() == "primary" {
					continue
				}

				o.Spec.Constraints = append(o.Spec.Constraints, runtime.Build(func(o *databasev1.Constraint) {
					o.Code = databasev1.ConstraintCode(key.Name())
					o.Spec.Unique = key.IsUnique()
					o.Spec.Primary = key.IsPrimary()

					keyDef := sqlbuilder.GetKeyDef(key)
					if keyDef != nil {
						o.Spec.Method = keyDef.Method()
					}

					for col := range key.Cols() {
						cn := runtime.Build(func(o *databasev1.ConstraintColumn) {
							o.Code = databasev1.ColumnCode(col.Name())

							for _, opt := range keyDef.FieldNameAndOptions() {
								if name := opt.Name(); name == col.Name() || name == col.FieldName() {
									o.Options = opt.Options()
								}
							}
						})

						o.Spec.Columns = append(o.Spec.Columns, cn)
					}
				}))
			}
		}))
	}

	return list, nil
}

func (c *Controller) Open(ctx context.Context) (session.Adapter, error) {
	return session.Open(ctx, c.dsn.String())
}

func (c *Controller) IsReady(ctx context.Context) error {
	pgConn, err := pgx.Connect(ctx, c.dsn.String())
	if err != nil {
		if isErrorUnknownDatabase(err) {
			return c.createDatabase(ctx)
		}
		return err
	}
	defer func() {
		_ = pgConn.Close(ctx)
	}()

	if err := pgConn.Ping(ctx); err != nil {
		if isErrorUnknownDatabase(err) {
			return c.createDatabase(ctx)
		}
		return err
	}
	return nil
}

func (c *Controller) createDatabase(ctx context.Context) error {
	dsn := *c.dsn
	dsn.Path = "/postgres"

	pgConn, err := pgx.Connect(ctx, dsn.String())
	if err != nil {
		return fmt.Errorf("connect to database failed: %w", err)
	}
	defer func() {
		_ = pgConn.Close(ctx)
	}()

	q, args := sqlfrag.Collect(ctx, sqlfrag.Pair("CREATE DATABASE ?;", sqlfrag.Const(c.dbCode)))
	if _, err := pgConn.Exec(ctx, q, args...); err != nil {
		return fmt.Errorf("database creation failed: %w", err)
	}

	return nil
}
