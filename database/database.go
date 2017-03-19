package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var (
	// DbConn is connection to database
	DbConn *pg.DB
)

// TmpUser struct
type TmpUser struct {
	ID        uint
	Email     string
	Name      string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (u TmpUser) String() string {
	return fmt.Sprintf("User<ID:%d, Email:%s, Name:%s, UpdatedAt:%s, CreatedAt:%s>", u.ID, u.Email, u.Name, u.UpdatedAt, u.CreatedAt)
}

// User struct
type User struct {
	ID    uint
	Email string
	Name  string
	Age   uint8
}

func init() {
	if DbConn == nil {
		DbConn = pg.Connect(&pg.Options{
			User:     "postgres",
			Password: "password",
			Database: "test",
		})
	}
}

// Init desc
func Init() {
	err := DbConn.CreateTable(&TmpUser{}, &orm.CreateTableOptions{
		Temp: true,
	})
	if err != nil {
		panic(err)
	}
	sob := TmpUser{ID: 1, Email: "admin@test.oleg", Name: "Oleg", UpdatedAt: time.Now(), CreatedAt: time.Now()}
	if err := DbConn.Insert(&sob); err != nil {
		panic(err)
	}
	tmpU := TmpUser{}
	DbConn.Model(&tmpU).Where("tmp_user.name=?", "Oleg").Select()
	fmt.Println(tmpU)

	users := []User{}
	DbConn.Model(&users).OrderExpr("id").Select(&users)
	fmt.Println(users)
}

func Migrate() {
	_, err := DbConn.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id serial PRIMARY KEY,
      email varchar(255) NOT NULL,
      name varchar(255) NOT NULL,
      age integer
    )
  `)
	if err != nil {
		fmt.Println(err)
	}
}
