package models

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yoshimi-I/Go_RESTAPI/config"
	"log"
)

//主にDB接続を行う
var Db *sql.DB

var err error

//テーブル定義を作成
const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    uid STRING NOT NULL UNIQUE,
    name STRING,
    email STRING,
    paswword STRING,
    created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

}

// uuidとpasswordはユーザーが決めるのではなくGo側で作る必要があるため別途用意する必要がある
func CreateUUID(uuidobj uuid.UUID) {

}
