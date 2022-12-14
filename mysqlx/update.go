/**
 * @Author: YMBoom
 * @Description:
 * @File:  update
 * @Version: 1.0.0
 * @Date: 2022/12/14 16:43
 */
package mysqlx

type (
	UpdateValue struct {
		column string
		value  any
		tbName string
	}

	UpdateOption func(uv *UpdateValue)
)

func NewUpdateValue(column string, value any, opts ...UpdateOption) *UpdateValue {
	uv := &UpdateValue{}
	if column == "" || value == nil {
		return nil
	}

	uv.column = column
	uv.value = value

	for _, opt := range opts {
		opt(uv)
	}

	return uv
}

func WithUpdateTbName(tbName string) UpdateOption {
	return func(uv *UpdateValue) {
		uv.tbName = tbName
	}
}
