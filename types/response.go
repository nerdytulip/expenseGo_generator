package types

import(
	"net/http"
)

type ListResponse struct{
      *Expense
}


func Listexpense( expense *Expense ) *ListResponse{

 resp := &ListResponse{ Expense : expense }
 return resp

}

func (ListResponse) Render(w http.ResponseWriter, r *http.Request) error {
   return nil
}

type ListResponses struct{
  Expense []*Expense
}

func ListExpenses ( expenses []*Expense ) *ListResponses{
    res := &ListResponses{ Expense : expenses }
     return res
}

func (ListResponses) Render(w http.ResponseWriter, r *http.Request) error {
   return nil
}



