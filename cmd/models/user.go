package models

import "time"

//テーブルと同じ型を作成する
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	CreatedAt time.Time
}

func (u User) CreateUser() (err error) {
	cmd := `insert into users(
       			  uuid,
                  name,
                  email,
                  password,
                  created_at) values (?,?,?,?,?)`
	_, err := Db.Exec(cmd) //Dbばmodesパッケージに入っているから使える
}
