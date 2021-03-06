package orm

import (
	"context"

	"github.com/jinzhu/gorm"
)

type (
	// GORM connection
	GORM struct {
		Type string `config:"type" help:"mysql, postgres, sqlite3, mssql"`
		URI  string `config:"uri" help:"Connection uri"`

		name string
		*gorm.DB
	}
)

var (
	// Default GORM
	Default = New("gorm")
)

// Name config name
func (orm *GORM) Name() string {
	return orm.name
}

// ConfigWillLoad before hook
func (orm *GORM) ConfigWillLoad(context.Context) {

}

// ConfigDidLoad after hook
func (orm *GORM) ConfigDidLoad(context.Context) {
	if orm.URI == "" {
		panic("gorm config is invalid")
	}

	if orm.Type != "mysql" && orm.Type != "postgres" && orm.Type != "sqlite3" && orm.Type != "mssql" {
		panic("gorm connection type invalid.")
	}

	db, err := gorm.Open(orm.Type, orm.URI)
	if err != nil {
		panic(err)
	}

	orm.DB = db
}

// Serve ok
func (orm *GORM) Serve(ctx context.Context) error {
	return nil
}

// Shutdown stop
func (orm *GORM) Shutdown(ctx context.Context) error {
	if orm.DB != nil {
		return orm.DB.Close()
	}

	return nil
}

// New options
func New(name string) *GORM {
	return &GORM{
		name: name,
	}
}
