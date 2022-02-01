package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"mytest.net/restapi/api/db"
	"mytest.net/restapi/api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

func (c *Controller) GetSales(clts *model.Clients, ctx *gin.Context) {

	var limit int64 = 20
	filter := bson.M{}
	cur, err := db.Find("sales", filter, clts.DB, limit)
	if err != nil {
		log.Errorf("error", err)
		return
	}
	var results []*model.Sales
	for cur.Next(context.TODO()) {
		var elem model.Sales
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("error", err)
			return
		}
		results = append(results, &elem)
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *Controller) CreateSale(clts *model.Clients, ctx *gin.Context) {
	var addSales model.Sales
	if err := ctx.ShouldBindJSON(&addSales); err != nil {
		log.Errorf("error", err)
		return
	}
	ds := model.Sales{
		ID:               addSales.ID,
		InvoiceNo:        addSales.InvoiceNo,
		SalesDate:        addSales.SalesDate,
		CreatedDate:      addSales.CreatedDate,
		ModifiedDate:     addSales.ModifiedDate,
		CustomerName:     addSales.CustomerName,
		TotalValue:       addSales.TotalValue,
		Currency:         addSales.Currency,
		SalesOrderDetail: addSales.SalesOrderDetail,
	}
	res, err := db.InsertOne("sales", ds, clts.DB)
	if err != nil {
		log.Errorf("error", err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetSaleByID(clts *model.Clients, ctx *gin.Context) {

	sId := ctx.Request.URL.Query().Get("id")
	var sm model.Sales
	id, _ := primitive.ObjectIDFromHex(sId)
	d := bson.M{"_id": id}
	err := clts.DB.Collection("sales").FindOne(context.TODO(), d).Decode(&sm)
	if err != nil {
		log.Errorf("error on get sales by id ", err)
		return
	}
	ctx.JSON(http.StatusOK, sm)
}

func (c *Controller) UpdateSale(clts *model.Clients, ctx *gin.Context) {
	// function implementation
}

func (c *Controller) SalesOrderDetails(clts *model.Clients, ctx *gin.Context) {

	var limit int64 = 20
	startTime := ctx.Request.URL.Query().Get("startTime")
	endTime := ctx.Request.URL.Query().Get("endTime")
	const (
        layoutISO = "2006-01-02T15:04:05.000Z"
    )
    st, _ := time.Parse(layoutISO, startTime) //converted to ISODate format
    et, _ := time.Parse(layoutISO, endTime)   //converted to ISODate format
	pipeline := bson.A{
		bson.M{"$match": bson.M{"createddate":bson.M{"$gte": st,"$lt": et }}},
		bson.M{"$unwind": bson.M{"path": "$salesorderdetail"}},
		bson.M{"$lookup": bson.M{"from": "products", "localField": "salesorderdetail.CloudProductID", "foreignField": "id", "as": "products"}},
		bson.M{"$limit": limit},
	}
	cur, err := db.Aggregate("sales", pipeline, clts.DB, limit)
	if err != nil {
		log.Errorf("error", err)
		return
	}
	var results []*model.SalesOrderDetails
	for cur.Next(context.TODO()) {
		var elem model.SalesOrderDetails
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("error", err)
			return
		}
		results = append(results, &elem)
	}
	ctx.JSON(http.StatusOK, results)
}
