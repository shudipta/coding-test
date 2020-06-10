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


CREATE TABLE `order_details` (
  `order_id` smallint(6) NOT NULL,
  `product_id` smallint(6) NOT NULL,
  `unit_price` double NOT NULL,
  `quantity` smallint(6) NOT NULL,
  `discount` double NOT NULL,
  `rating` double DEFAULT NULL,
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `order_details_fk_product_id` (`product_id`),
  CONSTRAINT `order_details_fk_order_id` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_details_fk_product_id` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "order_id": 72}
*/

// OrderDetail struct is a row record of the order_details table in the test database
type OrderDetail struct {
	OrderID   int             `gorm:"column:order_id;type:SMALLINT;primary_key" json:"order_id" protobuf:"int16,0,opt,name=order_id"` //[ 0] order_id                                       smallint             null: false  primary: true   auto: false  col: smallint        len: -1      default: []
	ProductID int             `gorm:"column:product_id;type:SMALLINT;" json:"product_id" protobuf:"int16,1,opt,name=product_id"`      //[ 1] product_id                                     smallint             null: false  primary: false  auto: false  col: smallint        len: -1      default: []
	UnitPrice float64         `gorm:"column:unit_price;type:DOUBLE;" json:"unit_price" protobuf:"float,2,opt,name=unit_price"`        //[ 2] unit_price                                     double               null: false  primary: false  auto: false  col: double          len: -1      default: []
	Quantity  int             `gorm:"column:quantity;type:SMALLINT;" json:"quantity" protobuf:"int16,3,opt,name=quantity"`            //[ 3] quantity                                       smallint             null: false  primary: false  auto: false  col: smallint        len: -1      default: []
	Discount  float64         `gorm:"column:discount;type:DOUBLE;" json:"discount" protobuf:"float,4,opt,name=discount"`              //[ 4] discount                                       double               null: false  primary: false  auto: false  col: double          len: -1      default: []
	Rating    float64		  `gorm:"column:rating;type:DOUBLE;" json:"rating" protobuf:"float,5,opt,name=rating"`                    //[ 5] rating                                         double               null: true   primary: false  auto: false  col: double          len: -1      default: []

}

// TableName sets the insert table name for this struct type
func (o *OrderDetail) TableName() string {
	return "order_details"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OrderDetail) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OrderDetail) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OrderDetail) Validate(action Action) error {
	return nil
}
