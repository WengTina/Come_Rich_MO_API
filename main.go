package main

// Project Author: Shane, shane871112@hotmail.com
// GCC require!! https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe
import (
	"net/http"

	"eirc.app/internal/pkg/dao/gorm"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/v1/router"
	routeAccount "eirc.app/internal/v1/router/account"
	routerCustomer "eirc.app/internal/v1/router/customer"
	routerFile "eirc.app/internal/v1/router/file"
	routerLogin "eirc.app/internal/v1/router/login"
	routerManuOrder "eirc.app/internal/v1/router/manu_order"
	routerSalesInfo "eirc.app/internal/v1/router/sales_info"

	AccountModel "eirc.app/internal/v1/structure/accounts"
	fileModel "eirc.app/internal/v1/structure/file"
	ManuOrderModel "eirc.app/internal/v1/structure/manu_order"
	RawMaterialModel "eirc.app/internal/v1/structure/raw_material"
)

// @version 0.1
// @author Shane
// @description COME RICH 製令平台

func main() {
	dbLY, err := gorm.New()
	if err != nil {
		log.Error(err)
		return
	}

	db, err := gorm.NewSQLite()
	if err != nil {

		log.Error(err)
		return
	}
	db.AutoMigrate(&fileModel.Table{})
	db.AutoMigrate(&AccountModel.Table{})
	db.AutoMigrate(&ManuOrderModel.Table{})
	db.AutoMigrate(&RawMaterialModel.Table{})
	route := router.Default()
	route = routerCustomer.GetRoute(route, dbLY)      //客戶路由
	route = routerSalesInfo.GetRoute(route, dbLY, db) //銷貨單路由
	route = routerFile.GetRoute(route, db, dbLY)      //檔案上傳路由
	route = routerLogin.GetRoute(route, db)
	route = routeAccount.GetRoute(route, db)
	route = routerManuOrder.GetRoute(route, db)
	log.Fatal(http.ListenAndServe("192.168.50.254:8090", route))
}
