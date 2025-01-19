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

	tripPath := flag.String("trip", "./trip_test.csv", "input trips csv file path")
	trackerPath := flag.String("tracker", "./tracker_test.csv", "input trackers csv file path")

	baseName := filepath.Base(*tripPath)
	ext := filepath.Ext(baseName)
	nameWithoutExt := baseName[:len(baseName)-len(ext)]
	defaultOutput := nameWithoutExt + "_result" + ext

	output := flag.String("out", defaultOutput, "output csv file path")
	flag.Parse()

	tripFile, err := os.Open(*tripPath)
	if err != nil {
		log.Fatalf("failed to open trip file : %v", err)
	}
	defer tripFile.Close()

	trackerFile, err := os.Open(*trackerPath)
	if err != nil {
		log.Fatalf("failed to open tracker file : %v", err)
	}
	defer trackerFile.Close()

	trackers := []*Tracker{}
	if err := gocsv.UnmarshalFile(trackerFile, &trackers); err != nil {
		log.Fatalf("failed unmarshal trackers: %v", err)
	}

	imeiToId := make(map[string]string)

	for _, tracker := range trackers {
		imeiToId[tracker.Imei] = tracker.Id
	}

	oldTrips := []*OldTrip{}
	if err := gocsv.UnmarshalFile(tripFile, &oldTrips); err != nil {
		log.Fatalf("failed unmarshal old trip: %v", err)
	}

	outFile, err := os.Create(*output)
	if err != nil {
		log.Fatalf("failed to create CSV result: %v", err)
	}
	defer outFile.Close()

	// newTrips := []*NewTrip{}
	if err := gocsv.MarshalFile(&[]*NewTrip{}, outFile); err != nil {
		log.Fatalf("failed to  write to CSV result: %v", err)
	}
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
				TrackerId:            imeiToId[t.Imei],
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
			if err := gocsv.MarshalWithoutHeaders(&[]*NewTrip{nt}, outFile); err != nil {
				log.Fatalf("failed to  write to CSV result: %v", err)
			}
			// newTrips = append(newTrips, nt)
		}
	}

	// if err := gocsv.MarshalFile(&newTrips, outFile); err != nil {
	// 	log.Fatalf("failed to  write to CSV result: %v", err)
	// }

	log.Println("CSV result written successfully")
}

func getSafeValue[T any](slice []T, index int, defaultValue T) T {
	if index < len(slice) {
		return slice[index]
	}
	return defaultValue
}
