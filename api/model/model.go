package model

import (
	"go.mongodb.org/mongo-driver/mongo"
	"mytest.net/restapi/config"
	"time"
)

type Clients struct {
	DB     *mongo.Database
	Config config.Config
}

type Conf struct {
	Config config.Config
}

type Product struct {
	ID                 string        `json:"id,omitempty"`
	Name               string        `json:"name,omitempty" binding:"required"`
	BarCode            string        `json:"barCode" binding:"required"`
	ProductRefNo       string        `json:"productRefNo" binding:"required"`
	CostPricesArray    []interface{} `json:"costPricesArray" binding:"required"`
	SellingPricesArray []interface{} `json:"sellingPricesArray,omitempty" binding:"required"`
	StoreID            string        `json:"storeID,omitempty" binding:"required"`
	Weight             string        `json:"weight,omitempty" binding:"required"`
}

type Sales struct {
	ID               string        `json:"id,omitempty"`
	InvoiceNo        string        `json:"invoiceNo,omitempty" binding:"required"`
	SalesDate        time.Time     `json:"salesDate,omitempty"`
	CreatedDate      time.Time     `json:"createdDate,omitempty"`
	ModifiedDate     time.Time     `json:"modifiedDate,omitempty"`
	CustomerName     string        `json:"customerName,omitempty" binding:"required"`
	TotalValue       float32       `json:"totalValue,omitempty" binding:"required"`
	Currency         string        `json:"currency,omitempty" binding:"required"`
	SalesOrderDetail []interface{} `json:"salesOrderDetail,omitempty" binding:"required"`
}

type SalesOrderDetails struct {
	ID               string                 `json:"id,omitempty"`
	InvoiceNo        string                 `json:"invoiceNo,omitempty"`
	SalesDate        time.Time              `json:"salesDate,omitempty"`
	CreatedDate      time.Time              `json:"createdDate,omitempty"`
	ModifiedDate     time.Time              `json:"modifiedDate,omitempty"`
	CustomerName     string                 `json:"customerName,omitempty"`
	TotalValue       float32                `json:"totalValue,omitempty"`
	Currency         string                 `json:"currency,omitempty"`
	SalesOrderDetail map[string]interface{} `json:"salesOrderDetail,omitempty"`
	Products         []interface{}          `json:"products,omitempty"`
}
