package repository

import (
	"context"

	"github.com/shudipta/coding-test/model"
)

// GetAllOrderDetails is a function to get a slice of record(s) from order_details table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllOrderDetails(ctx context.Context, page, pagesize int64, order string) (orderdetails []*model.OrderDetail, totalRows int, err error) {

	orderdetails = []*model.OrderDetail{}

	orderdetailsOrm := DB.Model(&model.OrderDetail{})
	orderdetailsOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		orderdetailsOrm = orderdetailsOrm.Offset(offset).Limit(pagesize)
	} else {
		orderdetailsOrm = orderdetailsOrm.Limit(pagesize)
	}

	if order != "" {
		orderdetailsOrm = orderdetailsOrm.Order(order)
	}

	if err = orderdetailsOrm.Find(&orderdetails).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return orderdetails, totalRows, nil
}

// GetOrderDetail is a function to get a single record to order_details table in the test database
// error - ErrNotFound, db Find error
func GetOrderDetail(ctx context.Context, id int) (record model.OrderDetail, err error) {
	if err = DB.First(&record, id).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddOrderDetail is a function to add a single record to order_details table in the test database
// error - ErrInsertFailed, db save call failed
func AddOrderDetail(ctx context.Context, orderdetail *model.OrderDetail) (result *model.OrderDetail, RowsAffected int64, err error) {
	db := DB.Save(orderdetail)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return orderdetail, db.RowsAffected, nil
}

// UpdateOrderDetail is a function to update a single record from order_details table in the test database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateOrderDetail(updated *model.OrderDetail) (err error) {
	res := DB.Model(&model.OrderDetail{}).Where("product_id = ? AND order_id = ?", updated.ProductID, updated.OrderID).
		Update("rating", updated.Rating)
	if res.RecordNotFound() {
		return ErrNotFound
	}
	if res.Error != nil {
		return ErrUpdateFailed
	}

	return nil
}

// DeleteOrderDetail is a function to delete a single record from order_details table in the test database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteOrderDetail(ctx context.Context, id int) (rowsAffected int64, err error) {

	orderdetail := &model.OrderDetail{}
	db := DB.First(orderdetail, id)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(orderdetail)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}

/*

   "CommandLine"          [string] "conv-db --sqltype=mysql --connstr root:root@tcp(localhost:3306)/test --database test --json --gorm --rest --out ./ --module github.com/shudipta/coding-test --mod --server --makefile --json-fmt=snake --generate-proj --overwrite --db --generate-dao --generate-proj --dao=repository"
   "DatabaseName"         [string] "test"
   "ShortStructName"      [string] "o"
   "StructName"           [string] "OrderDetail"
   "SwaggerInfo"          [*main.swaggerInfoDetails] &main.swaggerInfoDetails{Version:"1.0", Host:"localhost:8080", BasePath:"/", Title:"Sample CRUD api for test db", Description:"Sample CRUD api for test db", TOS:"", ContactName:"Me", ContactURL:"http://me.com/terms.html", ContactEmail:"me@me.com"}
   "TableInfo"            [*dbmeta.ModelInfo] &dbmeta.ModelInfo{Index:2, IndexPlus1:3, PackageName:"model", StructName:"OrderDetail", ShortStructName:"o", TableName:"order_details", Fields:[]string{"OrderID int `gorm:\"column:order_id;type:SMALLINT;primary_key\" json:\"order_id\" protobuf:\"int16,0,opt,name=order_id\"` //[ 0] order_id                                       smallint             null: false  primary: true   auto: false  col: smallint        len: -1      default: []", "ProductID int `gorm:\"column:product_id;type:SMALLINT;\" json:\"product_id\" protobuf:\"int16,1,opt,name=product_id\"` //[ 1] product_id                                     smallint             null: false  primary: false  auto: false  col: smallint        len: -1      default: []", "UnitPrice float64 `gorm:\"column:unit_price;type:DOUBLE;\" json:\"unit_price\" protobuf:\"float,2,opt,name=unit_price\"` //[ 2] unit_price                                     double               null: false  primary: false  auto: false  col: double          len: -1      default: []", "Quantity int `gorm:\"column:quantity;type:SMALLINT;\" json:\"quantity\" protobuf:\"int16,3,opt,name=quantity\"` //[ 3] quantity                                       smallint             null: false  primary: false  auto: false  col: smallint        len: -1      default: []", "Discount float64 `gorm:\"column:discount;type:DOUBLE;\" json:\"discount\" protobuf:\"float,4,opt,name=discount\"` //[ 4] discount                                       double               null: false  primary: false  auto: false  col: double          len: -1      default: []", "Rating sql.NullFloat64 `gorm:\"column:rating;type:DOUBLE;\" json:\"rating\" protobuf:\"float,5,opt,name=rating\"` //[ 5] rating                                         double               null: true   primary: false  auto: false  col: double          len: -1      default: []"}, DBMeta:(*dbmeta.dbTableMeta)(0xc0003e0a80), Instance:(*struct { OrderID int "json:\"order_id\"" })(0xc0004126d0), CodeFields:[]*dbmeta.FieldInfo{(*dbmeta.FieldInfo)(0xc00035c000), (*dbmeta.FieldInfo)(0xc00035c0c0), (*dbmeta.FieldInfo)(0xc00035c180), (*dbmeta.FieldInfo)(0xc00035c240), (*dbmeta.FieldInfo)(0xc00035c300), (*dbmeta.FieldInfo)(0xc00035c3c0)}, PrimaryKeyField:0, PrimaryKeyGoType:"int", PrimaryKeyFieldParser:"parseInt"}
   "TableName"            [string] "order_details"
   "apiFQPN"              [string] "github.com/shudipta/coding-test/api"
   "apiPackageName"       [string] "api"
   "daoFQPN"              [string] "github.com/shudipta/coding-test/repository"
   "daoPackageName"       [string] "repository"
   "modelFQPN"            [string] "github.com/shudipta/coding-test/model"
   "modelPackageName"     [string] "model"
   "module"               [string] "github.com/shudipta/coding-test"
   "outDir"               [string] "./"
   "serverHost"           [string] "localhost"
   "serverPort"           [int] 8080
   "sqlConnStr"           [string] "root:root@tcp(localhost:3306)/test"
   "sqlType"              [string] "mysql"
   "structs"              [[]string] []string{"Category", "Customer", "OrderDetail", "Order", "Product"}
   "tableInfos"           [map[string]*dbmeta.ModelInfo] map[string]*dbmeta.ModelInfo{"categories":(*dbmeta.ModelInfo)(0xc000373110), "customers":(*dbmeta.ModelInfo)(0xc00009a1a0), "order_details":(*dbmeta.ModelInfo)(0xc000124270), "orders":(*dbmeta.ModelInfo)(0xc00009a340), "products":(*dbmeta.ModelInfo)(0xc00009a4e0)}
   "tables"               [[]string] []string{"categories", "customers", "order_details", "orders", "products"}

table test.order_details {

}
*/
