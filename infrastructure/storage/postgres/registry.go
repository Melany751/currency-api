package postgres

import (
	"challenge/domain/model"
	"database/sql"
	"reflect"
)

var (
	_psqlInsertRegistry = `INSERT INTO domain.registries ("transaction", "date", "duration") VALUES ($1, $2, $3)`
)

type Registry struct {
	db *sql.DB
}

func NewRegistryStorage(db *sql.DB) Registry {
	return Registry{db}
}

func (r Registry) CreateRegistryStorage(registry model.Registry) error {
	args := r.readModelRegistry(registry)

	stmt, err := r.db.Prepare(_psqlInsertRegistry)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	return nil
}

func (r Registry) readModelRegistry(registryEntry model.Registry) []any {
	v := reflect.ValueOf(registryEntry)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}
