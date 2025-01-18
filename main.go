package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/jftuga/geodist"
)

func main() {

	path := flag.String("file", "./trip_test.csv", "csv file path")
	flag.Parse()

	oldCSV, err := os.Open(*path)
	if err != nil {
		log.Fatalf("failed to open new trip: %v", err)
	}
	defer oldCSV.Close()

	oldTrips := []*OldTrip{}
	if err := gocsv.UnmarshalFile(oldCSV, &oldTrips); err != nil {
		log.Fatalf("failed unmarshal old trip: %v", err)
	}

	newTrips := []*NewTrip{}
	for _, t := range oldTrips {
		nt := &NewTrip{
			Imei:                 t.Imei,
			VehicleId:            t.VehicleId,
			VehicleConfigId:      t.VehicleConfigId,
			TripId:               t.Tag,
			TripPart:             "0",
			IsLast:               "0",
			TrackerVersion:       "0",
			VehicleVersion:       "0",
			VehicleConfigVersion: "0",
			DriverId:             t.DriverId,
			// ASK
			// TrackerId: ,
		}

		// Handle the first geometry record separately
		if len(t.Geometries) > 0 {
			nt.Longitude = t.Geometries[0][0]
			nt.Latitude = t.Geometries[0][1]
			nt.Speed = 0
			nt.Idle = 0
			nt.TraveledDistance = 0
			nt.Elapsed = 0
		}

		for i := 1; i < len(t.Geometries); i++ {
			nt.Longitude = t.Geometries[i][0]
			nt.Latitude = t.Geometries[i][1]
			if i < len(t.Speeds) {
				nt.Speed = t.Speeds[i]
			} else {
				nt.Speed = 0
			}
			if i < len(t.Idles) {
				nt.Idle = t.Idles[i]
			} else {
				nt.Idle = 0
			}
			if i < len(t.Timestamps) {
				nt.Timestamp = t.Timestamps[i]
			} else {
				nt.Timestamp = 0
			}
			current := geodist.Coord{Lon: t.Geometries[i][0], Lat: t.Geometries[i][1]}
			before := geodist.Coord{Lon: t.Geometries[i-1][0], Lat: t.Geometries[i-1][1]}
			_, nt.TraveledDistance, _ = geodist.VincentyDistance(current, before)
			nt.Elapsed = t.Timestamps[i] - t.Timestamps[i-1]

		}
		newTrips = append(newTrips, nt)
	}
	fmt.Println(*newTrips[0])
	fmt.Println(*newTrips[1])

	resFile, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("failed to create CSV file: %v", err)
	}
	defer resFile.Close()

	if err := gocsv.MarshalFile(&newTrips, resFile); err != nil {
		log.Fatalf("failed to  write to CSV: %v", err)
	}

	log.Println("CSV file written successfully")
}
