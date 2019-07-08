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
	DBname string
	Collname string
}


func main() {
	var struc string
	var varname string
	var slice_Structure string
	var database_name string
	var collection_name string

	flag.StringVar(&struc,"structurename","","name of the structure")
	//varname := strings.ToLower(struc)
	flag.StringVar(&varname,"variablename", "" ,"name of the types variable")
	//slice_Structure := "slice_" + var_name
	flag.StringVar(&slice_Structure,"slicevar","","name of the slice variable")
	flag.StringVar(&database_name,"dbname","","name od database")
	flag.StringVar(&collection_name,"collname","","name of collection")
	flag.Parse()


	box :=packr.New("temp","./templates")
	t,err := box.FindString("crud.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	t1,err := box.FindString("request.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	t2,err := box.FindString("response.gotpl")
	if err != nil {
		log.Fatal(err)
	}
	t3,err := box.FindString("routes.gotpl")
	if err != nil {
		log.Fatal(err)
	}

	t4,err := box.FindString("errors.gotpl")
	if err != nil {
		log.Fatal(err)
	}

	tc,err := template.New("crud").Parse(t)
	if err != nil {
		log.Fatal(err)
	}
	tr,err := template.New("routes").Parse(t3)
	if err != nil {
		log.Fatal(err)
	}
	tres,err := template.New("response").Parse(t2)
	if err != nil {
		log.Fatal(err)
	}
	treq,err := template.New("request").Parse(t1)
	if err != nil {
		log.Fatal(err)
	}
	terr,err := template.New("errors").Parse(t4)
	if err != nil {
		log.Fatal(err)
	}


	data:=MetaData{
		Struct:struc,
		Varname:varname,
		Sliceofstructure:slice_Structure,
		DBname:database_name,
		Collname:collection_name,
	}
	filecrud,err :=os.Create("crud.go")
	fileroutes,err :=os.Create("routes.go")
	filereq,err :=os.Create("request.go")
	fileres,err :=os.Create("response.go")
	fileerr,err :=os.Create("errors.go")
	err =tc.Execute(filecrud,data)
	if err != nil {
		fmt.Println(err)
	}
	err =tr.Execute(fileroutes,data)
	if err != nil {
		fmt.Println(err)
	}
	err =tres.Execute(fileres,data)
	if err != nil {
		fmt.Println(err)
	}
	err =treq.Execute(filereq,data)
	if err != nil {
		fmt.Println(err)
	}
	err = terr.Execute(fileerr,data)
	if err != nil {
		fmt.Println(err)
	}
}

