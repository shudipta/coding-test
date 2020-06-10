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


CREATE TABLE `orders` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `customer_id` smallint(6) DEFAULT NULL,
  `order_date` date DEFAULT NULL,
  `required_date` date DEFAULT NULL,
  `shipped_date` date DEFAULT NULL,
  `freight` double DEFAULT NULL,
  `ship_name` varchar(40) DEFAULT NULL,
  `ship_address` varchar(60) DEFAULT NULL,
  `ship_city` varchar(15) DEFAULT NULL,
  `ship_region` varchar(15) DEFAULT NULL,
  `ship_postal_code` varchar(10) DEFAULT NULL,
  `ship_country` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_fk_customer_id` (`customer_id`),
  CONSTRAINT `orders_fk_customer_id` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 89}
*/

// Order struct is a row record of the orders table in the test database
type Order struct {
	ID             int             `gorm:"AUTO_INCREMENT;column:id;type:SMALLINT;primary_key" json:"id" protobuf:"int16,0,opt,name=id"`                          //[ 0] id                                             smallint             null: false  primary: true   auto: true   col: smallint        len: -1      default: []
	CustomerID     sql.NullInt64   `gorm:"column:customer_id;type:SMALLINT;" json:"customer_id" protobuf:"int16,1,opt,name=customer_id"`                         //[ 1] customer_id                                    smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	OrderDate      time.Time       `gorm:"column:order_date;type:DATE;" json:"order_date" protobuf:"uint64,2,opt,name=order_date"`                               //[ 2] order_date                                     date                 null: true   primary: false  auto: false  col: date            len: -1      default: []
	RequiredDate   time.Time       `gorm:"column:required_date;type:DATE;" json:"required_date" protobuf:"uint64,3,opt,name=required_date"`                      //[ 3] required_date                                  date                 null: true   primary: false  auto: false  col: date            len: -1      default: []
	ShippedDate    time.Time       `gorm:"column:shipped_date;type:DATE;" json:"shipped_date" protobuf:"uint64,4,opt,name=shipped_date"`                         //[ 4] shipped_date                                   date                 null: true   primary: false  auto: false  col: date            len: -1      default: []
	Freight        sql.NullFloat64 `gorm:"column:freight;type:DOUBLE;" json:"freight" protobuf:"float,5,opt,name=freight"`                                       //[ 5] freight                                        double               null: true   primary: false  auto: false  col: double          len: -1      default: []
	ShipName       sql.NullString  `gorm:"column:ship_name;type:VARCHAR;size:40;" json:"ship_name" protobuf:"string,6,opt,name=ship_name"`                       //[ 6] ship_name                                      varchar(40)          null: true   primary: false  auto: false  col: varchar         len: 40      default: []
	ShipAddress    sql.NullString  `gorm:"column:ship_address;type:VARCHAR;size:60;" json:"ship_address" protobuf:"string,7,opt,name=ship_address"`              //[ 7] ship_address                                   varchar(60)          null: true   primary: false  auto: false  col: varchar         len: 60      default: []
	ShipCity       sql.NullString  `gorm:"column:ship_city;type:VARCHAR;size:15;" json:"ship_city" protobuf:"string,8,opt,name=ship_city"`                       //[ 8] ship_city                                      varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []
	ShipRegion     sql.NullString  `gorm:"column:ship_region;type:VARCHAR;size:15;" json:"ship_region" protobuf:"string,9,opt,name=ship_region"`                 //[ 9] ship_region                                    varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []
	ShipPostalCode sql.NullString  `gorm:"column:ship_postal_code;type:VARCHAR;size:10;" json:"ship_postal_code" protobuf:"string,10,opt,name=ship_postal_code"` //[10] ship_postal_code                               varchar(10)          null: true   primary: false  auto: false  col: varchar         len: 10      default: []
	ShipCountry    sql.NullString  `gorm:"column:ship_country;type:VARCHAR;size:15;" json:"ship_country" protobuf:"string,11,opt,name=ship_country"`             //[11] ship_country                                   varchar(15)          null: true   primary: false  auto: false  col: varchar         len: 15      default: []

}

// TableName sets the insert table name for this struct type
func (o *Order) TableName() string {
	return "orders"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *Order) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *Order) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *Order) Validate(action Action) error {
	return nil
}
