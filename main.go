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
	Varname string
	Sliceofstructure string
}


func main() {
	var struc string
	var varname string
	var slice_Structure string

	flag.StringVar(&struc,"structurename","","name of the structure")
	//varname := strings.ToLower(struc)
	flag.StringVar(&varname,"variablename", "" ,"name of the types variable")
	//slice_Structure := "slice_" + var_name
	flag.StringVar(&slice_Structure,"slicevar","","name of the slice variable")
	flag.Parse()
	box :=packr.New("temp","./templates")
	t,err := box.FindString("response.gotpl")
	if err != nil {
	  log.Fatal(err)
	}
	tr,err := template.New("response").Parse(t)
	if err != nil {
		log.Fatal(err)
	}


	data:=MetaData{
		Struct:struc,
		Varname:varname,
		Sliceofstructure:slice_Structure,
	}

	fileres,err :=os.Create("response.go")
	err =tr.Execute(fileres,data)
	if err != nil {
		fmt.Println(err)
	}
}
