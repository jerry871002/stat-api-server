package main

import "database/sql"

type SqliteStatStore struct {
	db *sql.DB
}

func NewSqliteStatStore(db *sql.DB) *SqliteStatStore {
	return &SqliteStatStore{db: db}
}

func (s *SqliteStatStore) GetPitching() ([]map[string]any, error) {
	rows, err := s.db.Query("SELECT * FROM pitching")
	if err != nil {
		return nil, err
	}

	return s.rowsToMap(rows)
}

func (s *SqliteStatStore) GetBatting() ([]map[string]any, error) {
	rows, err := s.db.Query("SELECT * FROM batting")
	if err != nil {
		return nil, err
	}

	return s.rowsToMap(rows)
}

func (s *SqliteStatStore) GetFielding() ([]map[string]any, error) {
	rows, err := s.db.Query("SELECT * FROM fielding")
	if err != nil {
		return nil, err
	}

	return s.rowsToMap(rows)
}

func (s *SqliteStatStore) rowsToMap(rows *sql.Rows) ([]map[string]any, error) {
	var data []map[string]any
	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]any, len(cols))
		columnPointers := make([]any, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]any)
		for i, colName := range cols {
			val := columnPointers[i].(*any)
			m[colName] = *val
		}
		data = append(data, m)
	}

	return data, nil
}
