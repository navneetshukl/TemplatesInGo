package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Student struct {
	Name string
	Roll int
	Hobby string
}

func renderTemplate(w http.ResponseWriter,tmpl string,td *Student){
	files,err:=template.ParseFiles("./templates/"+tmpl)
	if err!=nil{
		fmt.Println("Not Running Part")
		fmt.Println(err)
	}
	err=files.Execute(w,td)
	if err!=nil{
		fmt.Println(err)
		return
	}
}

func main() {
	
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		name1 := "Navneet Shukla"
		roll1 := 34
		renderTemplate(w,"home.page.tmpl",&Student{
			Name: name1,
			Roll: roll1,
		})
		
	})
	http.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request) {
		hobby:="I Like to Play Video Game"
		renderTemplate(w,"about.page.tmpl",&Student{
			Hobby: hobby,
		})

	})
	http.ListenAndServe(":8000",nil)
}
