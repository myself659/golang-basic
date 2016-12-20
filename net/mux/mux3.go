package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	Title  string `json:"title"`
	Rating string `json:"rating"`
	Year   string `json:"year"`
}

var movies = map[string]*Movie{
	"tt0076759": {Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
	"tt0082971": {Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/movies", handleMovies).Methods("GET")
	// 后面的内容认为是变量，一个简单的模式
	router.HandleFunc("/movie/{imdbKey}", handleMovie).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func handleMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	// 获取变量
	vars := mux.Vars(req)
	imdbKey := vars["imdbKey"]

	log.Println("Request for:", imdbKey)

	if movie, ok := movies[imdbKey]; ok {
		outgoingJSON, error := json.Marshal(movie)

		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res, string(outgoingJSON))
	} else {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res, string("Movie not found"))
	}
}

func handleMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	outgoingJSON, error := json.Marshal(movies)

	if error != nil {
		log.Println(error.Error())
		http.Error(res, error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(res, string(outgoingJSON))
}

/*
$ curl -X GET http://127.0.0.1:8080/movie/tt0082971 -v
Results:

* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /movie/tt0082971 HTTP/1.1
> User-Agent: curl/7.37.1
> Host: 127.0.0.1:8080
> Accept:
*/
/*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Mon, 06 Jul 2015 23:11:15 GMT
< Content-Length: 79
<
* Connection #0 to host 127.0.0.1 left intact
{"title":"Indiana Jones: Raiders of the Lost Ark","rating":"8.6","year":"1981"}


*/
