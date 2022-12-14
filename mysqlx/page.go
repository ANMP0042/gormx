/**
 * @Author: YMBoom
 * @Description:
 * @File:  page
 * @Version: 1.0.0
 * @Date: 2022/12/14 11:33
 */
package mysqlx

type Page struct {
	Page      int
	Size      int
	Order     string
	OrderType string
}

func NewPage(page, size int, order, orderType string) *Page {
	p := new(Page)

	if page == 0 {
		p.Page = 1
	}

	if size == 0 {
		p.Size = 10
	}

	if order == "" {
		p.Order = "id"
	}

	if orderType == "" {
		p.OrderType = "asc"
	}
	return p
}

func newPage() *Page {
	return &Page{
		Page:      1,
		Size:      10,
		Order:     "id",
		OrderType: "asc",
	}
}
