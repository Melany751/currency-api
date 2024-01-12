package postgres

import (
	"challenge/domain/model"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4"
	"reflect"
)

var (
	_psqlGetCurrency    = `SELECT * FROM domain.currencies WHERE create_date >= $1 and create_date <= $2`
	_psqlInsertCurrency = `INSERT INTO domain.currencies (id, "code", "value", "create_date") VALUES ($1, $2, $3, $4, $5)`
)

type Currency struct {
	db *sql.DB
}

func NewCurrencyStorage(db *sql.DB) Currency {
	return Currency{db}
}

func (c Currency) GetCurrenciesStorage(currencyParam, fInit, fEnd string) (model.Currencies, error) {
	args := []any{fInit, fEnd}

	querySql := _psqlGetCurrency

	if currencyParam != "all" {
		fmt.Println(_psqlGetCurrency)
		querySql += " and code = $3"
		args = append(args, currencyParam)
	}

	stmt, err := c.db.Prepare(querySql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var currency model.Currency
	var currencies model.Currencies

	for rows.Next() {
		currency, err = c.scanRowCurrency(rows)
		if err != nil {
			break
		}
		currencies = append(currencies, currency)
	}
	return currencies, nil
}

func (c Currency) CreateCurrencyStorage(currency model.Currency) error {
	args := c.readModelCurrency(currency)

	stmt, err := c.db.Prepare(_psqlInsertCurrency)
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

func (c Currency) scanRowCurrency(s pgx.Row) (model.Currency, error) {
	m := model.Currency{}

	err := s.Scan(
		&m.ID,
		&m.Code,
		&m.Value,
		&m.CreateDate,
	)
	if err != nil {
		return model.Currency{}, err
	}

	return m, nil
}

func (c Currency) readModelCurrency(currency model.Currency) []any {
	v := reflect.ValueOf(currency)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	return values
}

func (c Currency) GetCurrenciesValues(currencies model.Currencies) (bool, error) {

	return false, nil
}
