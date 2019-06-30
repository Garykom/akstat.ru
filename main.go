package main

import (
	"fmt"
	//"github.com/icza/gowut/gwu"
	//"io/ioutil"
	"net/http"
	"strings"
	//"path"
	//"bytes"
	//"net/smtp"
	//"log"
	//"os"
	//"os/exec"
	//"time"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	categories := dbGetCategories()

	json.NewEncoder(w).Encode(categories)
}

func getPoints(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	points := dbGetPoints()

	featureCollection := FeatureCollection{
		Type: "FeatureCollection"}

	for _, point := range points {
		fmt.Println("id:", point.ID)
		feature := Feature{
			Type: "Feature",
			Id:   point.ID}

		feature.Geometry = Geometry{
			Type: "Point"}
		feature.Geometry.Coordinates = append(feature.Geometry.Coordinates, point.Latitude)
		feature.Geometry.Coordinates = append(feature.Geometry.Coordinates, point.Longitude)

		feature.Properties = Properties{
			BalloonContent: "Содержимое балуна",
			ClusterCaption: "Метка с iconContent",
			HintContent:    "Текст подсказки",
			IconContent:    "1"}

		feature.Options = Options{
			IconLayout:    "default#image",
			IconImageHref: "http://akstat.ru/assets/bus.png"}
		feature.Options.IconImageSize = append(feature.Options.IconImageSize, 20)
		feature.Options.IconImageSize = append(feature.Options.IconImageSize, 20)

		featureCollection.Features = append(featureCollection.Features, feature)
	}

	userVar2, _ := json.Marshal(featureCollection)
	fmt.Println(string(userVar2))

	json.NewEncoder(w).Encode(featureCollection)
}

func createPoint(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rstr := formatRequest(r)
	fmt.Println(rstr)

	var point Point
	_ = json.NewDecoder(r.Body).Decode(&point)

	dbCreatePoint(point)

	json.NewEncoder(w).Encode(point)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/points", getPoints).Methods("GET")
	r.HandleFunc("/points", createPoint).Methods("POST")
	r.HandleFunc("/categories", getCategories).Methods("GET")

	port := "2222"
	// sudo kill $(sudo lsof -t -i:2222)
	// you may need to change the address
	fmt.Println("Open http://1cvpn.ru:" + port + " in your browser to see the result")

	handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(":"+port, handler); nil != err {
		fmt.Println(err)
	}
}
