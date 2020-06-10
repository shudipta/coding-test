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


CREATE TABLE `products` (
  `id` smallint(6) NOT NULL AUTO_INCREMENT,
  `product_name` varchar(40) NOT NULL,
  `supplier_id` smallint(6) DEFAULT NULL,
  `category_id` smallint(6) DEFAULT NULL,
  `quantity_per_unit` varchar(20) DEFAULT NULL,
  `unit_price` double DEFAULT NULL,
  `units_in_stock` smallint(6) DEFAULT NULL,
  `units_on_order` smallint(6) DEFAULT NULL,
  `reorder_level` smallint(6) DEFAULT NULL,
  `discontinued` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `products_fk_category_id` (`category_id`),
  CONSTRAINT `products_fk_category_id` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=78 DEFAULT CHARSET=latin1

JSON Sample
-------------------------------------
{    "id": 46}
*/

// Product struct is a row record of the products table in the test database
type Product struct {
	ID              int             `gorm:"AUTO_INCREMENT;column:id;type:SMALLINT;primary_key" json:"id" protobuf:"int16,0,opt,name=id"`                            //[ 0] id                                             smallint             null: false  primary: true   auto: true   col: smallint        len: -1      default: []
	ProductName     string          `gorm:"column:product_name;type:VARCHAR;size:40;" json:"product_name" protobuf:"string,1,opt,name=product_name"`                //[ 1] product_name                                   varchar(40)          null: false  primary: false  auto: false  col: varchar         len: 40      default: []
	SupplierID      sql.NullInt64   `gorm:"column:supplier_id;type:SMALLINT;" json:"supplier_id" protobuf:"int16,2,opt,name=supplier_id"`                           //[ 2] supplier_id                                    smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	CategoryID      sql.NullInt64   `gorm:"column:category_id;type:SMALLINT;" json:"category_id" protobuf:"int16,3,opt,name=category_id"`                           //[ 3] category_id                                    smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	QuantityPerUnit sql.NullString  `gorm:"column:quantity_per_unit;type:VARCHAR;size:20;" json:"quantity_per_unit" protobuf:"string,4,opt,name=quantity_per_unit"` //[ 4] quantity_per_unit                              varchar(20)          null: true   primary: false  auto: false  col: varchar         len: 20      default: []
	UnitPrice       sql.NullFloat64 `gorm:"column:unit_price;type:DOUBLE;" json:"unit_price" protobuf:"float,5,opt,name=unit_price"`                                //[ 5] unit_price                                     double               null: true   primary: false  auto: false  col: double          len: -1      default: []
	UnitsInStock    sql.NullInt64   `gorm:"column:units_in_stock;type:SMALLINT;" json:"units_in_stock" protobuf:"int16,6,opt,name=units_in_stock"`                  //[ 6] units_in_stock                                 smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	UnitsOnOrder    sql.NullInt64   `gorm:"column:units_on_order;type:SMALLINT;" json:"units_on_order" protobuf:"int16,7,opt,name=units_on_order"`                  //[ 7] units_on_order                                 smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	ReorderLevel    sql.NullInt64   `gorm:"column:reorder_level;type:SMALLINT;" json:"reorder_level" protobuf:"int16,8,opt,name=reorder_level"`                     //[ 8] reorder_level                                  smallint             null: true   primary: false  auto: false  col: smallint        len: -1      default: []
	Discontinued    int             `gorm:"column:discontinued;type:INT;" json:"discontinued" protobuf:"int32,9,opt,name=discontinued"`                             //[ 9] discontinued                                   int                  null: false  primary: false  auto: false  col: int             len: -1      default: []

}

// TableName sets the insert table name for this struct type
func (p *Product) TableName() string {
	return "products"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (p *Product) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (p *Product) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (p *Product) Validate(action Action) error {
	return nil
}
