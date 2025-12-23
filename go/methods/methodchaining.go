package main

import (
	"fmt"
	"strconv"
	"strings"
)

type QueryBuilder struct {
	table      string
	conditions []string
	parameters []interface{}
	orderBy    string
	limit      int
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		table:      table,
		conditions: []string{},
		parameters: []interface{}{},
		limit:      -1,
	}
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.conditions = append(qb.conditions, condition)
	qb.parameters = append(qb.parameters, args...)
	return qb
}

func (qb *QueryBuilder) OrderBy(column, direction string) *QueryBuilder {
	qb.orderBy = column + " " + direction
	return qb
}

func (qb *QueryBuilder) Limit(n int) *QueryBuilder {
	qb.limit = n
	return qb
}

func (qb *QueryBuilder) Build() string {
	query := "SELECT * FROM " + qb.table

	if len(qb.conditions) > 0 {
		query += " WHERE " + strings.Join(qb.conditions, " AND ")
	}
	if qb.orderBy != "" {
		query += " ORDER BY " + qb.orderBy
	}
	if qb.limit > 0 {
		query += " LIMIT " + strconv.Itoa(qb.limit)
	}
	return query
}

func main() {
	query := NewQueryBuilder("users").
		Where("age > ?", 18).
		Where("status = ?", "active").
		OrderBy("name", "ASC").
		Limit(10).
		Build()

	fmt.Println(query)
}
