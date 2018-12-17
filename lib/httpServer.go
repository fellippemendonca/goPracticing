package lib;

import (
	"fmt";
	"log";
	"net/http";
	"encoding/json";
)


type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
};

type Articles []Article;


func HttpServer() {
	fmt.Println("\n\n ## The following content is exploring Go Http Server ## \n");
	handleRequests();
}


func handleRequests() {
	myMux := http.NewServeMux();
	myMux.HandleFunc("/", homePage);
	myMux.HandleFunc("/articles", allArticles);
	log.Fatal(http.ListenAndServe(":8080", myMux));
}


func homePage(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "homePage Endpoint Hit");
}


func allArticles(responseWriter http.ResponseWriter, request *http.Request) {
	
	articles := Articles{
		Article{ Title: "Test Title", Desc: "Test Desctiption", Content: "Test Content"},
	};

	json.NewEncoder(responseWriter).Encode(articles);
	
	fmt.Println("allArticles Endpoint Hit");
}


