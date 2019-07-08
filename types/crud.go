package types

import(
"net/http"
"encoding/json"
"fmt"
"strconv"
"github.com/go-chi/render"
"github.com/go-chi/chi"
)


var expenses []*Expense




func Ctx(http.Handler) http.Handler{

 return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

 		ID := chi.URLParam(request, "id")
 		id,_:=strconv.Atoi(ID)
 		for _,Expense := range expenses{
 		//TODO implementation
 		}

 	})


}

func Create (writer http.ResponseWriter , request *http.Request){
var req CreateRequest

	err := render.Bind(request, &req)
	if err != nil {
		render.Render(writer, request, ErrInvalidRequest(err))
		return
	}

	//TODO implementation

		j, _ := json.Marshal(req.Expense)
    	writer.Header().Set("Content-Type", "application/json")
    	writer.WriteHeader(http.StatusCreated)

    	_, _ = fmt.Fprintf(writer, `{"success": true, "data": %v}`, string(j))

}
func GetAll (writer http.ResponseWriter , request *http.Request){

err := render.Render(writer, request, ListExpenses(expenses))
 if err != nil{
		render.Render(writer,request,ErrRender(err))
		return
	}

}
func Update (writer http.ResponseWriter , request *http.Request){

 expense := request.Context().Value("expense").(Expense)

 var req UpdateRequest
 err:=render.Bind(request,&req)
 if err != nil {
 		render.Render(writer,request,ErrRender(err))
 		return
 	}

 	//TODO implementation

 if err = render.Render(writer, request, Listexpense(&expense)) ; err != nil{
  			render.Render(writer,request,ErrRender(err))
  			return

  	}

}
func GetOne (writer http.ResponseWriter , request *http.Request){
 expense := request.Context().Value("expense").(Expense)

 if err := render.Render(writer, request, Listexpense(&expense)) ; err != nil{
   			render.Render(writer,request,ErrRender(err))
   			return

   	}

}

func Delete (writer http.ResponseWriter , request *http.Request){
 expense := request.Context().Value("expense").(Expense)
 //TODO implementation
 err := render.Render(writer, request, ListExpenses(expenses))
 if err != nil{
 		render.Render(writer,request,ErrRender(err))
 		return
 	}


}