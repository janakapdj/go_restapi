package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"mytest.net/restapi/api/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"mytest.net/restapi/api/db"
	"net/http"
)

func (c *Controller) GetProducts(clts *model.Clients, ctx *gin.Context) {

	var limit int64 = 20
	filter := bson.M{}
	cur, err := db.Find("products", filter, clts.DB, limit)
	if err != nil {
		log.Errorf("error", err)
		return
	}
	var results []*model.Product
	for cur.Next(context.TODO()) {
		var elem model.Product
		err := cur.Decode(&elem)
		if err != nil {
			log.Errorf("error", err)
			return
		}
		results = append(results, &elem)
	}
	ctx.JSON(http.StatusOK, results)
}

func (c *Controller) CreateProduct(clts *model.Clients, ctx *gin.Context) {
	var addProduct model.Product
	if err := ctx.ShouldBindJSON(&addProduct); err != nil {
		log.Errorf("error", err)
		return
	}
	ds := model.Product{
		ID:                 addProduct.ID,
		Name:               addProduct.Name,
		BarCode:            addProduct.BarCode,
		ProductRefNo:       addProduct.ProductRefNo,
		CostPricesArray:    addProduct.CostPricesArray,
		SellingPricesArray: addProduct.SellingPricesArray,
		StoreID:            addProduct.StoreID,
		Weight:             addProduct.Weight,
	}
	res, err := db.InsertOne("products", ds, clts.DB)
	if err != nil {
		log.Errorf("error", err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *Controller) GetProductByID(clts *model.Clients, ctx *gin.Context) {

	sId := ctx.Request.URL.Query().Get("id")
	var sm model.Product
	id, _ := primitive.ObjectIDFromHex(sId)
	d := bson.M{"_id": id}
	err := clts.DB.Collection("products").FindOne(context.TODO(), d).Decode(&sm)
	if err != nil {
		log.Errorf("error on get products by id ", err)
		return
	}
	ctx.JSON(http.StatusOK, sm)
}

func (c *Controller) UpdateProduct(clts *model.Clients, ctx *gin.Context) {
	// function implementation
}

func (c *Controller) DeleteProduct(clts *model.Clients, ctx *gin.Context) {
	// function implementation
}
