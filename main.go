package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/jftuga/geodist"
)

func main() {

	input := flag.String("in", "./trip_test.csv", "input csv file path")

	baseName := filepath.Base(*input)
	ext := filepath.Ext(baseName)
	nameWithoutExt := baseName[:len(baseName)-len(ext)]
	outputDefault := nameWithoutExt + "_result" + ext

	output := flag.String("out", outputDefault, "output csv file path")
	flag.Parse()

	file, err := os.Open(*input)
	if err != nil {
		log.Fatalf("failed to open %s : %v", *input, err)
	}
	defer file.Close()

	oldTrips := []*OldTrip{}
	if err := gocsv.UnmarshalFile(file, &oldTrips); err != nil {
		log.Fatalf("failed unmarshal old trip: %v", err)
	}

	newTrips := []*NewTrip{}
	for _, t := range oldTrips {
		for i, geo := range t.Geometries {
			nt := &NewTrip{
				Imei:                 t.Imei,
				VehicleId:            t.VehicleId,
				VehicleConfigId:      t.VehicleConfigId,
				TripId:               t.Tag,
				DriverId:             t.DriverId,
				TripPart:             "0",
				IsLast:               "0",
				TrackerVersion:       "0",
				VehicleVersion:       "0",
				VehicleConfigVersion: "0",
				LocOutlier:           "0",
				Longitude:            geo[0],
				Latitude:             geo[1],
				Speed:                getSafeValue(t.Speeds, i, 0),
				Idle:                 getSafeValue(t.Idles, i, 0),
				Timestamp:            time.Unix(getSafeValue(t.Timestamps, i, 0), 0).Format(time.RFC3339),
				// ASK
				// TrackerId: ,
			}
			if i == 0 {
				nt.Speed = 0
				nt.Idle = 0
				nt.TraveledDistance = 0
				nt.Elapsed = 0
			} else {
				nt.Elapsed = t.Timestamps[i] - t.Timestamps[i-1]
				current := geodist.Coord{Lon: geo[0], Lat: geo[1]}
				before := geodist.Coord{Lon: t.Geometries[i-1][0], Lat: t.Geometries[i-1][1]}
				_, nt.TraveledDistance, _ = geodist.VincentyDistance(current, before)

			}
			newTrips = append(newTrips, nt)
		}
	}

	outFile, err := os.Create(*output)
	if err != nil {
		log.Fatalf("failed to create CSV result: %v", err)
	}
	defer outFile.Close()

	if err := gocsv.MarshalFile(&newTrips, outFile); err != nil {
		log.Fatalf("failed to  write to CSV result: %v", err)
	}

	log.Println("CSV result written successfully")
}

func getSafeValue[T any](slice []T, index int, defaultValue T) T {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}
