package app

import (
	"fmt"
	"net/http"
	"text/template"
)



func init() {
	http.HandleFunc("/", handlePata)
}

type File struct{
     First string
}

type File2 struct{
     First string
     Second string
     Mix string
}

func handlePata(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t,_:=template.ParseFiles("edit.html")
	p:=File{"Input 2 word"}
	t.Execute(w,p)

	first:=r.FormValue("first")
	second:=r.FormValue("second")
	firstr:=first
	secondr:=second

	if len(first)!=len(second) {
//	fmt.Fprintf(w,"different length!")
	if (len([]rune(first))<len([]rune(second))){
	 //fmt.Fprintf(w,"\nfirst")
	counter:=	len([]rune(second))-len([]rune(first))

 	for i:=0;i<counter;i++{
	firstr=firstr+"!"
	}

	}else
	{//fmt.Fprintf(w,"\nsecond")
	counter:=	len([]rune(first))-len([]rune(second))
	for i:=0;i<counter;i++{
	secondr=secondr+"!"
	}
	fmt.Fprintf(w,"%s %s",firstr,secondr)
	}
	} 
	if(len(firstr)==0){
	firstr = "  "
	secondr = "  "
	}
	mix:="  "

	for i:=0;i<len([]rune(firstr));i++{
	    firstRune:=string([]rune(firstr)[i:i+1])
	    secondRune:=string([]rune(secondr)[i:i+1])
	    mix = mix[:]+firstRune+secondRune
	}
	t2,_:=template.ParseFiles("edit2.html")
	answer:=File2{first,second,mix}

	t2.Execute(w,answer)
		
	


}

