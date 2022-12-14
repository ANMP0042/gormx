/**
 * @Author: YMBoom
 * @Description:
 * @File:  query
 * @Version: 1.0.0
 * @Date: 2022/12/14 11:59
 */
package mysqlx

type (
	query struct {
		tbName string
		q      string
		v      any

		w       Where
		in      *whereIn
		between *whereBetween
	}

	whereIn struct {
		column string
		value  any
	}

	whereBetween struct {
		column     string
		begin, end any
	}

	QueryOption     func(q *query)
	FindQueryOption func(fq *findQuery)
)

func newQuery(w Where, q string, v any, opts ...QueryOption) *query {
	qy := &query{w: w, q: q, v: v}

	for _, opt := range opts {
		opt(qy)
	}

	return qy
}

func WithTbName(tbName string) QueryOption {
	return func(q *query) {
		q.tbName = tbName
	}
}

func newWhereIn(column string, value any) *whereIn {
	if column == "" {
		return nil
	}

	if value == nil {
		return nil
	}

	return &whereIn{
		column: column,
		value:  value,
	}
}

func WithWhereIn(in *whereIn) QueryOption {
	return func(q *query) {
		q.in = in
	}
}

func newWhereBetween(column string, begin, end any) *whereBetween {
	if column == "" {
		return nil
	}

	if begin == nil || end == nil {
		return nil
	}

	return &whereBetween{
		column: column,
		begin:  begin,
		end:    end,
	}
}

func WithWhereBetween(between *whereBetween) QueryOption {
	return func(q *query) {
		q.between = between
	}
}

type findQuery struct {
	query *query
	page  *Page
}

func newFindQuery(q *query, opts ...FindQueryOption) *findQuery {
	fq := &findQuery{query: q}

	for _, opt := range opts {
		opt(fq)
	}
	return fq
}

func WithPage(p *Page) FindQueryOption {
	return func(fq *findQuery) {
		if p == nil {
			p = newPage()
		}
		fq.page = p
	}
}
