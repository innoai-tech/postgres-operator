package db

type pgDatabase struct {
	DatabaseName string `db:"datname"`
	IsTemplate   bool   `db:"datistemplate"`

	CharacterType    string `db:"datctype"`
	Collation        string `db:"datcollate"`
	CollationVersion string `db:"datcollversion"`
}

func (pgDatabase) TableName() string {
	return "pg_catalog.pg_database"
}
