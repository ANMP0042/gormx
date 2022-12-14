/**
 * @Author: YMBoom
 * @Description:
 * @File:  error
 * @Version: 1.0.0
 * @Date: 2022/12/14 10:09
 */
package errorx

type Errorx struct {
	Text string
}

func (ex *Errorx) Error() string {
	return ex.Text
}

type NotFoundError struct{}

func (e *NotFoundError) Error() string {
	return "gormx: record not found"
}

type InvalidParamError struct {
	Text string
}

func (e *InvalidParamError) Error() string {
	return "gormx: (param " + e.Text + ")"
}

type InvalidNonDBError struct {
	Text string
}

func (e *InvalidNonDBError) Error() string {
	return "gormx: (non-db " + e.Text + ")"
}

// InvalidNonPointerError 必须是指针类型
type InvalidNonPointerError struct {
	Text string
}

func (e *InvalidNonPointerError) Error() string {
	return "gormx: (non-pointer " + e.Text + ")"
}
