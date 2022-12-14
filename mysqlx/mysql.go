/**
 * @Author: YMBoom
 * @Description:
 * @File:  mysql
 * @Version: 1.0.0
 * @Date: 2022/12/14 9:56
 */
package mysqlx

import (
	"context"
	"fmt"
	"github.com/ANMP0042/gormx/common/errorx"
	"github.com/ANMP0042/gormx/configx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
)

type (
	Mysql struct {
		Ctx context.Context
		DB  *gorm.DB
	}

	MysqlOption func(mysql *Mysql)

	Where  map[string]interface{}
	Update map[string]interface{}
)

// New 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
// 中文文档地址：https://learnku.com/docs/gorm/v2
func New(config *configx.Config, opts ...MysqlOption) (*Mysql, error) {
	dsn, err := ParseDSN(config)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	msl := &Mysql{DB: db}

	for _, opt := range opts {
		opt(msl)
	}
	return msl, nil
}

func ParseDSN(cfg *configx.Config) (dsn string, err error) {
	if cfg.Dbname == "" {
		return "", &errorx.Errorx{Text: "ParseDSN:Dbname cant be null"}
	}

	if cfg.User == "" {
		cfg.User = "root"
	}

	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}

	if cfg.Port == 0 {
		cfg.Port = 3306
	}

	if cfg.Dsn == "" {
		cfg.Dsn = "charset=utf8mb4&parseTime=True&loc=Local"
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", cfg.User, cfg.Passwd, cfg.Host, cfg.Port, cfg.Dbname, cfg.Dsn)
	return
}

// Create  插入数据 v指针类型
// 要有效地插入大量记录，请将一个 slice 传递给 Create 方法
func (m *Mysql) Create(v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &errorx.InvalidNonPointerError{Text: "insert"}
	}

	return m.create(v)
}

func (m *Mysql) create(v any) error {
	if m.DB == nil {
		return &errorx.InvalidNonDBError{Text: "insert"}
	}
	return m.DB.Create(v).Error
}

// FirstById 通过id查询单条数据
func (m *Mysql) FirstById(id int64, q string, v any) error {
	if id == 0 {
		return &errorx.NotFoundError{}
	}

	w := map[string]interface{}{
		"id": id,
	}
	return m.first(newQuery(w, q, v))
}

// First 简单的查询单条数据
func (m *Mysql) First(w Where, q string, v any) error {
	return m.first(newQuery(w, q, v))
}

// FirstInTbName 简单的查询单条数据 指定表名
func (m *Mysql) FirstInTbName(w Where, tbName, q string, v any) error {
	return m.first(newQuery(w, q, v, WithTbName(tbName)))
}

// FirstInWhereIn whereIn
func (m *Mysql) FirstInWhereIn(w Where, q string, column string, value, v any) error {
	in := newWhereIn(column, value)
	if in == nil {
		return &errorx.InvalidParamError{Text: "FirstInWhereIn column or value cant be null"}
	}

	return m.first(newQuery(w, q, v, WithWhereIn(in)))
}

// FirstInBetween Between
func (m *Mysql) FirstInBetween(w Where, q string, column string, begin, end, v any) error {
	between := newWhereBetween(column, begin, end)
	if between == nil {
		return &errorx.InvalidParamError{Text: "FirstInBetween:column or value cant be null"}
	}

	return m.first(newQuery(w, q, v, WithWhereBetween(between)))
}

// FirstInQuery todo 直接使用query
func (m *Mysql) FirstInQuery(q *query) {

}

func (m *Mysql) first(fq *query) error {
	rv := reflect.ValueOf(fq.v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &errorx.InvalidNonPointerError{Text: "first"}
	}

	db := m.DB.Where(fq.w)
	if fq.tbName != "" {
		db = db.Table(fq.tbName)
	}

	if fq.in != nil {
		db = db.Where("? in (?) ", fq.in.column, fq.in.value)
	}

	if fq.between != nil {
		db = db.Where("? between ? and ? ", fq.between.column, fq.between.begin, fq.between.end)
	}

	return db.Select(fq.q).First(fq.v).Error
}

// Find 简单的查询多条数据
func (m *Mysql) Find(w Where, q string, v any) error {
	qy := newQuery(w, q, v)
	return m.find(newFindQuery(qy))
}

func (m *Mysql) FindInTbName(w Where, tbName, q string, v any) error {
	qy := newQuery(w, q, v, WithTbName(tbName))
	return m.find(newFindQuery(qy))
}

// FindInPage 简单的分页多条数据
func (m *Mysql) FindInPage(w Where, q string, v any, p *Page) error {
	qy := newQuery(w, q, v)
	return m.find(newFindQuery(qy, WithPage(p)))
}

// FindInPageInTbName 简单的分页多条数据
func (m *Mysql) FindInPageInTbName(w Where, tbName, q string, v any, p *Page) error {
	qy := newQuery(w, q, v, WithTbName(tbName))
	return m.find(newFindQuery(qy, WithPage(p)))
}

// FindInWhereIn whereIn
func (m *Mysql) FindInWhereIn(w Where, q string, column string, value, v any) error {
	in := newWhereIn(column, value)
	if in == nil {
		return &errorx.InvalidParamError{Text: "FindInWhereIn column or value cant be null"}
	}

	qy := newQuery(w, q, v, WithWhereIn(in))
	return m.find(newFindQuery(qy))
}

// FindInBetween Between
func (m *Mysql) FindInBetween(w Where, q string, column string, begin, end, v any) error {
	between := newWhereBetween(column, begin, end)
	if between == nil {
		return &errorx.InvalidParamError{Text: "FirstInBetween:column or value cant be null"}
	}
	qy := newQuery(w, q, v, WithWhereBetween(between))
	return m.find(newFindQuery(qy))
}

func (m *Mysql) find(fq *findQuery) error {
	rv := reflect.ValueOf(fq.query.v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &errorx.InvalidNonPointerError{Text: "find"}
	}

	db := m.DB.Where(fq.query.w)
	if fq.query.tbName != "" {
		db = db.Table(fq.query.tbName)
	}

	if fq.query.in != nil {
		db = db.Where("? in (?) ", fq.query.in.column, fq.query.in.value)
	}

	if fq.query.between != nil {
		db = db.Where("? between ? and ? ", fq.query.between.column, fq.query.between.begin, fq.query.between.end)
	}

	if fq.page != nil {
		db = db.Limit(fq.page.Size).Offset((fq.page.Page - 1) * fq.page.Size).Order(fq.page.Order + " " + fq.page.OrderType)
	}

	return db.Select(fq.query.q).Find(fq.query.v).Error
}

// RecordIsExist 记录是否存在 查询发生错误并且错误不是notfound 返回true
func (m *Mysql) RecordIsExist(w Where, tbName string, model any) bool {
	return m.recordIsExist(w, tbName, model)
}

func (m *Mysql) recordIsExist(w Where, tbName string, model any) bool {
	err := m.DB.Table(tbName).Where(w).First(model).Error

	if err == nil {
		return true
	} else {
		if err == gorm.ErrRecordNotFound {
			return false
		} else {
			return true
		}
	}
}

func (m *Mysql) UpdateById(id int64, uv *UpdateValue) error {
	if id == 0 {
		return &errorx.NotFoundError{}
	}

	w := map[string]interface{}{
		"id": id,
	}

	u := map[string]interface{}{
		uv.column: uv.value,
	}
	return m.update(w, u, uv.tbName)
}

func (m *Mysql) Update(w Where, uv *UpdateValue) error {
	u := map[string]interface{}{
		uv.column: uv.value,
	}
	return m.update(w, u, uv.tbName)
}

func (m *Mysql) UpdatesById(id int64, u Update, tbName string) error {
	if id == 0 || tbName == "" {
		return &errorx.NotFoundError{}
	}

	w := map[string]interface{}{
		"id": id,
	}
	return m.update(w, u, tbName)
}

func (m *Mysql) Updates(w Where, u Update, tbName string) error {
	return m.update(w, u, tbName)
}

func (m *Mysql) update(w Where, u Update, tbName string) error {
	db := m.DB.Where(w)

	if tbName != "" {
		db = db.Table(tbName)
	}

	return db.Updates(u).Error
}

func (m *Mysql) DeleteById(id int64, v any) error {
	if id == 0 {
		return &errorx.NotFoundError{}
	}

	w := map[string]interface{}{
		"id": id,
	}
	return m.delete(w, v)
}

func (m *Mysql) Delete(w Where, v any) error {
	if w == nil {
		return &errorx.NotFoundError{}
	}
	return m.delete(w, v)
}

func (m *Mysql) delete(w Where, v any) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &errorx.InvalidNonPointerError{Text: "delete"}
	}

	return m.DB.Where(w).Delete(v).Error
}
