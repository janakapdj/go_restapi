# Developing REST API with MongoDB and Go

### clone the repo
$ git clone https://github.com/janakapdj/go_restapi.git <project_name>

### Connect to Monogodb
update api/db/db.go file<br/>
ApplyURI("mongodb+srv://<db_user>:<db_password>@cluster0.c2bfq.mongodb.net/<data_base>?retryWrites=true&w=majority")

### Next step

### start local server
go run main.go<br/>

Run API end points on a postman<br/>
Products
<br/>
GET: http://localhost:4000/api/v1/products - get all products<br/>
POST: http://localhost:4000/api/v1/products/create - create new product<br/>
POST: http://localhost:4000/api/v1/products/byId?id=<product_id> - get product by id<br/>
<br/>
Sales

GET: http://localhost:4000/api/v1/sales - get all sales<br/>
POST: http://localhost:4000/api/v1/sales/create - create new sale<br/>
POST: http://localhost:4000/api/v1/sales/byId?id=<sale_id> - get sale by id<br/>
POST: http://localhost:4000/api/v1/sales/salesOrderDetails?startTime=<start_time>&endTime=<end_time> - get sales order details by