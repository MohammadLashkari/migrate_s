package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {

	oldCSV, err := os.Open("trip_test.csv")
	if err != nil {
		log.Fatal("failed to open new trip", err)
	}
	defer oldCSV.Close()

	oldTrips := []*OldTrip{}
	if err := gocsv.UnmarshalFile(oldCSV, &oldTrips); err != nil {
		log.Fatal("failed unmarshal old trip", err)
	}

	// newCSV, err := os.Open("trip_test_new.csv")
	// if err != nil {
	// 	log.Fatal("failed to open old trip", err)
	// }
	// defer newCSV.Close()
	//
	// newTrips := []*NewTrip{}
	// if err := gocsv.UnmarshalFile(newCSV, &newTrips); err != nil {
	// 	log.Fatal("failed unmarshal new trip", err)
	// }

	newTrips := []*NewTrip{}
	for _, t := range oldTrips {
		newTrip := &NewTrip{}
		newTrip.Imei = t.Imei
		newTrip.VehicleId = t.VehicleId
		newTrip.VehicleConfigId = t.VehicleConfigId
		newTrip.TripPart = "0"
		newTrip.IsLast = "0"

		for _, geo := range t.Geometries {
			newTrip.Latitude = geo[]
		}
		newTrips = append(newTrips, newTrip)
	}
	fmt.Println(len(oldTrips))

	// reader := csv.NewReader(file)
	// for {
	// 	record, err := reader.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		continue
	// 	}
	// 	fmt.Println(record)
	// 	break
	// }

}
