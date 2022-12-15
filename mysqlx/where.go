/**
 * @Author: YMBoom
 * @Description:
 * @File:  where
 * @Version: 1.0.0
 * @Date: 2022/12/15 14:47
 */
package mysqlx

import (
	"fmt"
	"strings"
)

//type Where map[string]any
//
//func MakeWhere() Where {
//	return make(Where)
//}
//
//func (w Where) Set(k string, v any) {
//	w[k] = v
//}

type (
	Operator string

	Wherex struct {
		Sql []Sql
	}

	Sql struct {
		Column   string
		Operator Operator
		Value    any
	}

	WherexOption func(w *Wherex)
)

const (
	EQ      Operator = "="
	NE      Operator = "!="
	GT      Operator = ">"
	LT      Operator = "<"
	GE      Operator = ">="
	LE      Operator = "<="
	IN      Operator = "IN"
	NIN     Operator = "NOT IN"
	BETWEEN Operator = "BETWEEN"
)

func NewWherex(opts ...WherexOption) *Wherex {
	wx := &Wherex{}

	for _, opt := range opts {
		opt(wx)
	}
	return wx
}

func (w *Wherex) SAdd(column string, operator Operator, value any) *Wherex {
	if column == "" || value == nil {
		return w
	}

	w.Add(Sql{
		Column:   column,
		Operator: operator,
		Value:    value,
	})
	return w
}

func (w *Wherex) Add(s Sql) *Wherex {
	if s.Operator == "" {
		s.Operator = EQ
	}

	w.Sql = append(w.Sql, s)
	return w
}

func (w *Wherex) Set(sql []Sql) *Wherex {
	if len(sql) == 0 {
		return w
	}

	for i, s := range sql {
		if s.Operator == "" {
			sql[i].Operator = EQ
		}
	}

	w.Sql = sql
	return w
}

// notice whereBetween 值能是any string int三种类型 其他类型不会加入到where中
// 如果需要其他类型 可以使用FirstInBetween 详见mysql_test.go
func (w *Wherex) toSql() (sql string, args []any) {
	if len(w.Sql) == 0 {
		return
	}

	var s []string
	for _, wSql := range w.Sql {
		tmp := ""
		switch wSql.Operator {
		case NIN, IN:
			tmp = fmt.Sprintf(" %s %s (?) ", wSql.Column, wSql.Operator)
			args = append(args, wSql.Value)
		case BETWEEN:
			switch wSql.Value.(type) {
			case []any:
				v := wSql.Value.([]any)
				args = append(args, v[0])
				args = append(args, v[1])
			case []string:
				v := wSql.Value.([]string)
				args = append(args, v[0])
				args = append(args, v[1])
			case []int:
				v := wSql.Value.([]int)
				args = append(args, v[0])
				args = append(args, v[1])
			default:
				continue
			}

			tmp = fmt.Sprintf(" %s %s ? and ? ", wSql.Column, wSql.Operator)
		default:
			tmp = fmt.Sprintf("%s %s ?", wSql.Column, wSql.Operator)
			args = append(args, wSql.Value)
		}

		fmt.Println("13132123132131")
		s = append(s, tmp)
	}
	return strings.Join(s, " AND "), args
}
