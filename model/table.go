package model

import (
	"database/sql"
	"time"
)

//sql语句:"show table status"对应的结构体
type Table struct {
	Name            string         `db:"Name"`
	Engine          string         `db:"Engine"`
	Version         string         `db:"Version"`
	Row_format      string         `db:"Row_format"`
	Rows            string         `db:"Rows"`
	Avg_row_length  string         `db:"Avg_row_length"`
	Data_length     string         `db:"Data_length"`
	Max_data_length string         `db:"Max_data_length"`
	Index_length    string         `db:"Index_length"`
	Data_free       string         `db:"Data_free"`
	Auto_increment  sql.NullString `db:"Auto_increment"`
	CreateTime      time.Time      `db:"Create_time"`
	Update_time     sql.NullTime   `db:"Update_time"`
	Check_time      sql.NullTime   `db:"Check_time"`
	Collation       string         `db:"Collation"` //使用的编码
	Checksum        sql.NullTime   `db:"Checksum"`
	Create_options  string         `db:"Create_options"`
	Comment         string         `db:"Comment"`
	TableCreateStr  string
}

type CreateTable struct {
	TableName string `db:"Table"`
	CreateStr string `db:"Create Table"`
}

type Column struct {
}
