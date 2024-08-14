package main

import (
	"database/sql"
)

type SqlStatStore struct {
	db *sql.DB
}

func NewSqlStatStore(db *sql.DB) *SqlStatStore {
	return &SqlStatStore{db: db}
}

func (s *SqlStatStore) GetTeams() ([]Team, error) {
	rows, err := s.db.Query("SELECT DISTINCT team, year FROM batting")
	if err != nil {
		return nil, err
	}

	data := []Team{}
	for rows.Next() {
		var team Team
		if err := rows.Scan(&team.Name, &team.Year); err != nil {
			return nil, err
		}
		data = append(data, team)
	}

	return data, nil
}

func (s *SqlStatStore) GetBattingStat(team string, year int) ([]map[string]any, error) {
	rows, err := s.db.Query("SELECT * FROM batting WHERE team = $1 AND year = $2", team, year)
	if err != nil {
		return nil, err
	}

	return s.rowsToMap(rows)
}

func (s *SqlStatStore) rowsToMap(rows *sql.Rows) ([]map[string]any, error) {
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
