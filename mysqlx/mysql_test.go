/**
 * @Author: YMBoom
 * @Description:
 * @File:  mysql_test
 * @Version: 1.0.0
 * @Date: 2022/12/15 13:46
 */
package mysqlx

import (
	"fmt"
	"github.com/ANMP0042/gormx/common/logger"
	"github.com/ANMP0042/gormx/configx"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	cfg := &configx.Config{
		Dbname:        "gormx_test",
		Passwd:        "123456",
		SingularTable: true,
	}
	m, err := New(cfg)
	if err != nil {
		fmt.Println("new mysql err", err)
		return
	}

	fmt.Println(m)
}

type User struct {
	Id        int
	Name      string
	Phone     string
	CreatedAt *time.Time
}

func TestMysql_FirstById(t *testing.T) {
	cfg := &configx.Config{
		Dbname:        "gormx_test",
		Passwd:        "123456",
		SingularTable: true,
	}
	m, err := New(cfg)
	if err != nil {
		fmt.Println("new mysql err", err)
		return
	}

	u := User{}
	err = m.FirstById(1, "name,phone", &u)
	if err != nil {
		fmt.Println("firstById err", err)
		return
	}

	fmt.Println("FirstById 查询结果：", u)

}

func TestMysql_FirstInWhereIn(t *testing.T) {
	cfg := &configx.Config{
		Dbname:        "gormx_test",
		Passwd:        "123456",
		SingularTable: true,
		LogLevel:      logger.Info,
	}
	m, err := New(cfg)
	if err != nil {
		fmt.Println("new mysql err", err)
		return
	}

	inUser := User{}
	w := NewWherex()
	w.SAdd("created_at", BETWEEN, []int64{1, 2})
	err = m.First(w, "id", &inUser)
	fmt.Println(err)
	fmt.Println("inUser", inUser)
}

func TestMysql_FirstInBetween(t *testing.T) {
	cfg := &configx.Config{
		Dbname:        "gormx_test",
		Passwd:        "123456",
		SingularTable: true,
		LogLevel:      logger.Info,
	}
	m, err := New(cfg)
	if err != nil {
		fmt.Println("new mysql err", err)
		return
	}

	betweenUser := User{}
	err = m.FirstInBetween(nil, "id,name,created_at", "created_at", "2022-12-15 00:00:00", "2022-12-15 11:00:00",
		&betweenUser)
	if err != nil {
		fmt.Println("FirstInBetween err", err)
		return
	}

	fmt.Println("FirstInBetween 查询结果：", betweenUser)
}
