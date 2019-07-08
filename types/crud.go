package types

import(
"net/http"
"encoding/json"
"fmt"
"context"
"time"
"log"
"strconv"
"github.com/go-chi/render"
"github.com/go-chi/chi"
"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo/options"
"go.mongodb.org/mongo-driver/mongo"
)

const DefaultDatabase = "expensedb"


type MongoHandler struct {
	client   *mongo.Client
	database string
}


var expenses []*Expense

//MongoHandler Constructor
func NewHandler(address string) *MongoHandler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))
	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}
	return mh
}


func (mh *MongoHandler) Ctx(next http.Handler) http.Handler{

 return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

 		ID := chi.URLParam(request, "id")
 		id,_:=strconv.Atoi(ID)
 		expense := &Expense{}
 		err := mh.GetOne_DB(expense,bson.M{"id":id})
 		if err !=nil{
          log.Println(err)
        }
        ctx := context.WithValue(request.Context(), "expense", expense )
        next.ServeHTTP(writer, request.WithContext(ctx))

 	})


}

func (mh *MongoHandler) AddOne_DB(expense *Expense)(*mongo.InsertOneResult, error){
    collection := mh.client.Database(mh.database).Collection("expensecoll")
  	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
  	result, err := collection.InsertOne(ctx, expense )
  	return result, err
}

func (mh *MongoHandler) Create (writer http.ResponseWriter , request *http.Request){
var req CreateRequest

	err := render.Bind(request, &req)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

    //TODO implementation
    _,err =mh.AddOne_DB(req.Expense)
    	if err!= nil{
    		log.Println(err)
    	}


		j, _ := json.Marshal(req.Expense)
    	writer.Header().Set("Content-Type", "application/json")
    	writer.WriteHeader(http.StatusCreated)

    	_, _ = fmt.Fprintf(writer, `{"success": true, "data": %v}`, string(j))

}
func (mh *MongoHandler) GetAll_DB(filter interface{}) []*Expense {
    collection := mh.client.Database(mh.database).Collection("expensecoll")
  	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

  	cur, err := collection.Find(ctx, filter)

  	if err != nil {
  		log.Fatal(err)
  	}
  	defer cur.Close(ctx)

  	var result []*Expense
  	for cur.Next(ctx) {
    		expense := &Expense{}
    		er := cur.Decode(expense)
    		if er != nil {
    			log.Fatal(er)
    		}
    		result = append(result, expense)
    	}
    	return result

}

func (mh *MongoHandler) GetAll (writer http.ResponseWriter , request *http.Request){

expenses := mh.GetAll_DB(bson.M{})
err := render.Render(writer, request, ListExpenses(expenses))
 if err != nil{
		render.Render(writer,request,ErrRender(err))
		return
	}

}

func (mh *MongoHandler) Update_DB(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mh.client.Database(mh.database).Collection("expensecoll")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    //TODO implementation
	result, err := collection.UpdateOne(ctx, filter, )
	return result, err
}

func (mh *MongoHandler) Update (writer http.ResponseWriter , request *http.Request){

 expense := request.Context().Value("expense").(Expense)

 var req UpdateRequest
 err:=render.Bind(request,&req)
 if err != nil {
 		render.Render(writer,request,ErrRender(err))
 		return
 	}

 	//TODO implementation
 	//TODO also add the structure sepcific ID attribute for the update_DB to execute
 	_, err = mh.Update_DB()
    	if err!=nil{
    		log.Println(err)
    	}


  _=mh.GetOne_DB()

 if err = render.Render(writer, request, Listexpense(&expense)) ; err != nil{
  			render.Render(writer,request,ErrRender(err))
  			return

  	}

}

func (mh *MongoHandler) GetOne_DB(expense *Expense ,  filter interface{}) error{
    collection := mh.client.Database(mh.database).Collection("expensecoll")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(expense)
	return err
}

func (mh *MongoHandler) GetOne (writer http.ResponseWriter , request *http.Request){
 expense := request.Context().Value("expense").(Expense)

 if err := render.Render(writer, request, Listexpense(&expense)) ; err != nil{
   			render.Render(writer,request,ErrRender(err))
   			return

   	}

}

func (mh *MongoHandler) RemoveOne_DB(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mh.client.Database(mh.database).Collection("expensecoll")
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)

	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}

func (mh *MongoHandler) Delete (writer http.ResponseWriter , request *http.Request){
 expense := request.Context().Value("expense").(Expense)
 //TODO implementation
 _, err := mh.RemoveOne_DB()
 	if err!=nil{
 		log.Println(err)
 		return
 	}



}