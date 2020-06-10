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


CREATE TABLE `customers` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `company_name` varchar(40) NOT NULL,
  `contact_name` varchar(30) DEFAULT NULL,
  `contact_title` varchar(30) DEFAULT NULL,
  `address` varchar(60) DEFAULT NULL,
  `city` varchar(15) DEFAULT NULL,
  `region` varchar(15) DEFAULT NULL,
  `postal_code` varchar(10) DEFAULT NULL,
  `country` varchar(15) DEFAULT NULL,
  `phone` varchar(24) DEFAULT NULL,
  `fax` varchar(24) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 97}
*/

// Customer struct is a row record of the customers table in the test database
type Customer struct {
	ID           int            `gorm:"AUTO_INCREMENT;column:id;type:SMALLINT;primary_key" json:"id" protobuf:"int16,0,opt,name=id"`                //[ 0] id                                             smallint             null: false  primary: true   auto: true   col: smallint        len: -1      default: []
	CompanyName  string         `gorm:"column:company_name;type:VARCHAR;size:40;" json:"company_name" protobuf:"string,1,opt,name=company_name"`    //[ 1] company_name                                   varchar(40)          null: false  primary: false  auto: false  col: varchar         len: 40      default: []
	ContactName  sql.NullString `gorm:"column:contact_name;type:VARCHAR;size:30;" json:"contact_name" protobuf:"string,2,opt,name=contact_name"`    //[ 2] contact_name                                   varchar(30)          null: true   primary: false  auto: false  col: varchar         len: 30      default: []
	ContactTitle sql.NullString `gorm:"column:contact_title;type:VARCHAR;size:30;" json:"contact_title" protobuf:"string,3,opt,name=contact_title"` //[ 3] contact_title                                  varchar(30)          null: true   primary: false  auto: false  col: varchar         len: 30      default: []
	Address      sql.NullString `gorm:"column:address;type:VARCHAR;size:60;" json:"address" protobuf:"string,4,opt,name=address"`                   //[ 4] address                                        varchar(60)          null: true   primary: false  auto: false  col: varchar         len: 60      default: []
	City         sql.NullString `gorm:"column:city;type:VARCHAR;size:15;" json:"city" protobuf:"string,5,opt,name=city"`                            //[ 5] city                                           varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []
	Region       sql.NullString `gorm:"column:region;type:VARCHAR;size:15;" json:"region" protobuf:"string,6,opt,name=region"`                      //[ 6] region                                         varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []
	PostalCode   sql.NullString `gorm:"column:postal_code;type:VARCHAR;size:10;" json:"postal_code" protobuf:"string,7,opt,name=postal_code"`       //[ 7] postal_code                                    varchar(10)          null: true   primary: false  auto: false  col: varchar         len: 10      default: []
	Country      sql.NullString `gorm:"column:country;type:VARCHAR;size:15;" json:"country" protobuf:"string,8,opt,name=country"`                   //[ 8] country                                        varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []
	Phone        sql.NullString `gorm:"column:phone;type:VARCHAR;size:24;" json:"phone" protobuf:"string,9,opt,name=phone"`                         //[ 9] phone                                          varchar(24)          null: true   primary: false  auto: false  col: varchar         len: 24      default: []
	Fax          sql.NullString `gorm:"column:fax;type:VARCHAR;size:24;" json:"fax" protobuf:"string,10,opt,name=fax"`                              //[10] fax                                            varchar(24)          null: true   primary: false  auto: false  col: varchar         len: 24      default: []

}

// TableName sets the insert table name for this struct type
func (c *Customer) TableName() string {
	return "customers"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Customer) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Customer) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Customer) Validate(action Action) error {
	return nil
}
