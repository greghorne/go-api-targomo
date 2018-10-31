package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// _ "github.com/lib/pq"
	"os"
	// "database/sql"
	"encoding/json"
	// "fmt"
)

// ============================================================
// connet string to pg server
var const_connect_string =" host=" + os.Getenv("GO_HOST") + 
	" database=" + os.Getenv("GO_DATABASE") + 
	" user=" + os.Getenv("GO_USER") + 
	" password=" + os.Getenv("GO_PASSWORD") + 
	" port=" + os.Getenv("GO_PORT") + 
	" sslmode=require"

var const_db_type = "postgres"
// ============================================================


// ============================================================
func main() {
	
	router := mux.NewRouter()
	router.HandleFunc("/v1/targomo-isochrone/{lng}/{lat}/{time}", v1TargomoIsochrone).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
// ============================================================


// ============================================================
func v1TargomoIsochrone (w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	bIntersects := v1DoTargomoIsochrone(params["lng"], params["lat"], params["time"], params["key"])
	// jsonResult  := map[string]bool{"intersects": bIntersects}

	json.NewEncoder(w).Encode("x")

}
// ============================================================


func getRegion(xLng string, yLat string) (region string) {

	lng := strconv.ParseFloat(xLng, 64)
	lat := strconv.ParseFloat(yLat, 64)

	if xLng <= -100 && lat <= 36 {
		region := "southwest"
	} else if lng <= -100 && lat >= 36 {
		region := "northwest"
	} else if lng >= -100 && lat >= 36 {
		region := "northeast"
	} else if lng >= -100 && lat <= 36 {
		region := "southeast"
	} else {
		region := "none"
	}

}


// ============================================================
func v1DoTargomoIsochrone(sxLng string, syLat string, sTime string, sKey string) (polygons string) {

	if region := getRegion(sxLng, syLat); region != "none" { 

		r360_key := "GYSWOYA0HD8JM1LAJYDX"

    	r360_url_string := "https://service.route360.net/na_" +
			region + "/v1/polygon?cfg={'sources':[{'lat':" + 
			xyLat + ",'lng':" + sxLng + 
			",'id':'Mappy','tm':{'car':{}}}],'polygon':" +
			"{'serializer':'geojson','srid':'4326'," +
			"'values':[" + sTime + "]}}&key=" + r360_key
		polygons = "good"
	} else {
		polgons = "none"
	}
				  
	// // connect to db
	// db, err := sql.Open(const_db_type, const_connect_string)
	
	// defer db.Close()
	// if err != nil { log.Fatal(err) }
	
	// // call pg function select z_tl_2016_us_state(lng, lat)
	// var strQuery = "select z_tl_2016_us_state($1, $2);"
	// rows, err   := db.Query(strQuery, xLng, yLat)
	
	// defer rows.Close()
	// if err != nil { log.Fatal(err) }

	// bIntersects = false
	// for rows.Next() {
	// 	bIntersects = true
	// } 

	return

}
// ============================================================


