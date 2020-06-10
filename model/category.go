package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `categories` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `category_name` varchar(15) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 10}
*/

// Category struct is a row record of the categories table in the test database
type Category struct {
	ID           int            `gorm:"AUTO_INCREMENT;column:id;type:SMALLINT;primary_key" json:"id" protobuf:"int16,0,opt,name=id"`                //[ 0] id                                             smallint             null: false  primary: true   auto: true   col: smallint        len: -1      default: []
	CategoryName string         `gorm:"column:category_name;type:VARCHAR;size:15;" json:"category_name" protobuf:"string,1,opt,name=category_name"` //[ 1] category_name                                  varchar(15)          null: false  primary: false  auto: false  col: varchar         len: 15      default: []
	Description  sql.NullString `gorm:"column:description;type:TEXT;size:65535;" json:"description" protobuf:"string,2,opt,name=description"`       //[ 2] description                                    text(65535)          null: true   primary: false  auto: false  col: text            len: 65535   default: []
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "categories"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Category) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Category) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Category) Validate(action Action) error {
	return nil
}
