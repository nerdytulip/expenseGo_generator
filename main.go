package main

import (
	"flag"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"log"
	"os"
	"text/template"
)

type MetaData struct {
	Struct string
	//varname string
	//sliceofstructure string
}


func main() {
	var struc string

	flag.StringVar(&struc,"structurename","","name of the structure")
	//var_name := strings.ToLower(struc)
	//flag.StringVar(&var_name,"variable name",var_name,"name of the variable for response")
	//slice_Structure := "slice_" + var_name
	//flag.StringVar(&slice_Structure,"variable name",slice_Structure,"name of the variable for response")
	flag.Parse()
	box :=packr.New("temp","./templates")
	t,err := box.FindString("request.gotpl")
	if err != nil {
	  log.Fatal(err)
	}
	tr,err := template.New("request").Parse(t)
	if err != nil {
		log.Fatal(err)
	}


	data:=MetaData{
		Struct:struc,
		//varname:var_name,
		//sliceofstructure:slice_Structure,

	}

	filereq,err :=os.Create("request.go")
	err =tr.Execute(filereq,data)
	if err != nil {
		fmt.Println(err)
	}
}
