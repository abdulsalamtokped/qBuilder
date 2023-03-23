package db

import "fmt"

var q *QueryBuilder

type QueryBuilder struct {
	condition []string
	query     map[string]interface{}
}

func GetInstance() *QueryBuilder {
	if q == nil {
		q = &QueryBuilder{
			query: make(map[string]interface{}),
		}
	}

	return q
}

func (q *QueryBuilder) Select(column string) *QueryBuilder {
	q.query["column"] = column
	return q
}
func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.query["table"] = table
	return q
}
func (q *QueryBuilder) Limit(limit int) *QueryBuilder {
	q.query["limit"] = limit
	return q
}
func (q *QueryBuilder) Offset(offset int) *QueryBuilder {
	q.query["offset"] = offset
	return q
}
func (q *QueryBuilder) Where(clause, operator string, value interface{}) *QueryBuilder {
	query := fmt.Sprintf("%s %s %v", clause, operator, value)
	switch operator {
	case "IN":
		query = fmt.Sprintf("%s %s(%v)", clause, operator, value)
	default:
	}

	q.condition = append(q.condition, query)
	return q
}
func (q *QueryBuilder) Build() string {
	var qWhere string
	var qLimit, qOffset int = -1, -1

	if len(q.condition) > 0 {
		qWhere = fmt.Sprintf("WHERE %s", q.condition[0])
		for i := 1; i < len(q.condition); i++ {
			qWhere += fmt.Sprintf(" AND %s", q.condition[i])
		}
	}
	if _, ok := q.query["limit"]; ok {
		qLimit = q.query["limit"].(int)
	}
	if _, ok := q.query["offset"]; ok {
		qOffset = q.query["offset"].(int)
	}
	query := fmt.Sprintf(
		`SELECT %s FROM %s %s LIMIT %d OFFSET %d`,
		q.query["column"],
		q.query["table"],
		qWhere,
		qLimit,
		qOffset,
	)

	return query
}
