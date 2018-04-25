package lib;

import (
	"fmt";
	"net/http";
)


func HttpServer() {

	fmt.Println("\n\n ## The following content is exploring Go Http Server ## \n");
	
	http.HandleFunc("/", handler);

	http.HandleFunc("/earth", handler2);

	http.ListenAndServe(":8080", nil);

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new access to Hello World");
	fmt.Fprintf(w, "Hello World\n");
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("new access to Hello Earth");
	fmt.Fprintf(w, "Hello Earth\n");
}
