package repository

import (
	"github.com/shudipta/coding-test/model"
)

// GetAllCategories is a function to get a slice of record(s) from categories table in the test database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllTopCategories() (categories []*model.Category, err error) {
	categories = []*model.Category{}

	res := DB.Raw(`
SELECT prod.category_id id
FROM order_details details

INNER JOIN products prod
ON details.product_id = prod.id AND details.rating IS NOT NULL

GROUP BY prod.category_id
ORDER BY AVG(details.rating) DESC;
`).Find(&categories)

	if res.RecordNotFound() {
		return nil, nil
	}
	if res.Error != nil {
		return nil, err
	}

	return categories, nil
}

/*

   "CommandLine"          [string] "conv-db --sqltype=mysql --connstr root:root@tcp(localhost:3306)/test --database test --json --gorm --rest --out ./ --module github.com/shudipta/coding-test --mod --server --makefile --json-fmt=snake --generate-proj --overwrite --db --generate-dao --generate-proj --dao=repository"
   "DatabaseName"         [string] "test"
   "ShortStructName"      [string] "c"
   "StructName"           [string] "Category"
   "SwaggerInfo"          [*main.swaggerInfoDetails] &main.swaggerInfoDetails{Version:"1.0", Host:"localhost:8080", BasePath:"/", Title:"Sample CRUD api for test db", Description:"Sample CRUD api for test db", TOS:"", ContactName:"Me", ContactURL:"http://me.com/terms.html", ContactEmail:"me@me.com"}
   "TableInfo"            [*dbmeta.ModelInfo] &dbmeta.ModelInfo{Index:0, IndexPlus1:1, PackageName:"model", StructName:"Category", ShortStructName:"c", TableName:"categories", Fields:[]string{"ID int `gorm:\"AUTO_INCREMENT;column:id;type:SMALLINT;primary_key\" json:\"id\" protobuf:\"int16,0,opt,name=id\"` //[ 0] id                                             smallint             null: false  primary: true   auto: true   col: smallint        len: -1      default: []", "CategoryName string `gorm:\"column:category_name;type:VARCHAR;size:15;\" json:\"category_name\" protobuf:\"string,1,opt,name=category_name\"` //[ 1] category_name                                  varchar(15)          null: false  primary: false  auto: false  col: varchar         len: 15      default: []", "Description sql.NullString `gorm:\"column:description;type:TEXT;size:65535;\" json:\"description\" protobuf:\"string,2,opt,name=description\"` //[ 2] description                                    text(65535)          null: true   primary: false  auto: false  col: text            len: 65535   default: []"}, DBMeta:(*dbmeta.dbTableMeta)(0xc0003e4120), Instance:(*struct { ID int "json:\"id\"" })(0xc000355fb8), CodeFields:[]*dbmeta.FieldInfo{(*dbmeta.FieldInfo)(0xc0001169c0), (*dbmeta.FieldInfo)(0xc000116a80), (*dbmeta.FieldInfo)(0xc000116b40)}, PrimaryKeyField:0, PrimaryKeyGoType:"int", PrimaryKeyFieldParser:"parseInt"}
   "TableName"            [string] "categories"
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

table test.categories {

}
*/
