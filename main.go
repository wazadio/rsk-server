package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var LocationCoordinate string = ""

func main() {
	http.HandleFunc("/", GetLocation)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://nominatim.openstreetmap.org/search.php?q=-0.9127291074293501,100.45824322784344&polygon_geojson=1&format=json")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	w.Write([]byte("hello world"))
}
