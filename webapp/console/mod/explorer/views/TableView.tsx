import { component } from "@nodepkg/runtime";
import type { DatabaseV1Table } from "@webapp/console/client/postgresOperator.ts";
import {
  Icon,
  mdiKey,
  mdiKeyOutline,
  SectionRow,
  styled,
} from "@nodepkg/dashboard";

export const SummaryLabel = styled("span")({
  px: 4,
});

export const TableView = component<{ table: DatabaseV1Table }>((props) => {
  return () => {
    return (
      <TableViewContainer>
        <summary>
          <SummaryLabel>{props.table.code}</SummaryLabel>
        </summary>
        <TableColumnListView>
          {props.table.spec?.columns?.map((c) => {
            return (
              <>
                <dt>{c.code}</dt>
                <dd>{c.spec.type}</dd>
              </>
            );
          })}
        </TableColumnListView>

        {props.table.spec?.constraints && (
          <TableConstraintListView>
            {props.table.spec?.constraints?.map((constraint) => {
              return (
                <TableConstraintItemView
                  data-primary={constraint.spec.primary}
                  data-unique={constraint.spec.unique}
                >
                  <Icon
                    path={constraint.spec.unique ? mdiKeyOutline : mdiKey}
                  />
                  <SectionRow>
                    <span>{constraint.code}</span>
                    <span>
                      {[
                        `(${constraint.spec.columns
                          .map((constraint) => {
                            return [
                              constraint.code,
                              ...(constraint.options ?? []),
                            ].join(" ");
                          })
                          .join(", ")})`,
                      ]}
                    </span>
                    <span>{constraint.spec.method}</span>
                  </SectionRow>
                </TableConstraintItemView>
              );
            })}
          </TableConstraintListView>
        )}
      </TableViewContainer>
    );
  };
});

const TableViewContainer = styled("details")({
  summary: {
    textStyle: "sys.label-small",
  },

  "&[open] > summary": {
    color: "sys.primary",
    fontWeight: "bold",
  },
});

const TableColumnListView = styled("dl")({
  textStyle: "sys.label-small",
  pl: "1em",
  dt: {
    fontWeight: "bold",
  },
  dd: {
    opacity: 0.7,
  },
});

const TableConstraintListView = styled("div")({
  display: "flex",
  flexDirection: "column",
  gap: 4,
  textStyle: "sys.label-small",
  pl: "1em",
});

const TableConstraintItemView = styled("div")({
  display: "flex",
  alignItems: "center",
  gap: "1em",
  textStyle: "sys.label-small",

  _primary: {
    [`${Icon}`]: {
      color: "sys.primary",
    },
  },

  _unique: {
    [`${Icon}`]: {
      position: "relative",
      "&::before": {
        content: JSON.stringify("1"),
        position: "absolute",
        right: -6,
        textStyle: "sys.label-small",
        fontSize: 8,
      },
    },
  },
});
