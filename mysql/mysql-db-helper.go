package mysql

import (
	"container/list"
	"database/sql"
)

//delegate function for sql data read
type delegate func(rows *sql.Rows, collection *list.List)

//ExecuteReader - executes @query and returns results in a list
func ExecuteReader(db *sql.DB, query string, fn delegate, params ...interface{}) (data *list.List, err error) {
	data = list.New()
	if rows, err := db.Query(query, params...); err == nil {
		defer rows.Close()
		for rows.Next() {
			fn(rows, data)
		}
		err = rows.Err()
	}
	return
}

//ExecuteInsert - performs a insert operation and returns latest inserted id
func ExecuteInsert(db *sql.DB, query string, params ...interface{}) (lastInsertedID int64, err error) {
	lastInsertedID = -1
	if stmt, err := db.Prepare(query); err == nil {
		defer stmt.Close()
		if result, err := stmt.Exec(params...); err == nil {
			lastInsertedID, err = result.LastInsertId()
		}
	}
	return
}

//ExecuteUpdateDelete - performs update and returns affected row count
func ExecuteUpdateDelete(db *sql.DB, query string, params ...interface{}) (updatedRowCount int64, err error) {
	updatedRowCount = 0
	if stmt, err := db.Prepare(query); err == nil {
		defer stmt.Close()
		if result, err := stmt.Exec(params...); err == nil {
			updatedRowCount, err = result.RowsAffected()
		}
	}
	return
}
