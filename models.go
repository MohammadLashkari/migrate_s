package main

type OldTrip struct {
	// Owner                     string `csv:"owner"`
	// Clients                   string `csv:"clients"`
	// Zones                     string `csv:"zones"`
	// Shares                    string `csv:"shares"`
	// Relations                 string `csv:"relations"`
	// CreatedBy                 string `csv:"created_by"`
	// CreatedAt                 string `csv:"created_at"`
	// CreatedIn                 string `csv:"created_in"`
	Imei                      string `csv:"imei"`
	Tag                       string `csv:"tag"`
	StartAddressOsmType       string `csv:"start_address_osm_type"`
	StartAddressOsmId         string `csv:"start_address_osm_id"`
	StartAddressDisplayName   string `csv:"start_address_display_name"`
	StartAddressName          string `csv:"start_address_name"`
	StartAddressCountry       string `csv:"start_address_country"`
	StartAddressState         string `csv:"start_address_state"`
	StartAddressCity          string `csv:"start_address_city"`
	StartAddressStateDistrict string `csv:"start_address_state_district"`
	StartAddressVillage       string `csv:"start_address_village"`
	EndAddressOsmType         string `csv:"end_address_osm_type"`
	EndAddressOsmId           string `csv:"end_address_osm_id"`
	EndAddressDisplayName     string `csv:"end_address_display_name"`
	EndAddressName            string `csv:"end_address_name"`
	EndAddressCountry         string `csv:"end_address_country"`
	EndAddressState           string `csv:"end_address_state"`
	EndAddressCity            string `csv:"end_address_city"`
	EndAddressStateDistrict   string `csv:"end_address_state_district"`
	EndAddressVillage         string `csv:"end_address_village"`
	Geometries                string `csv:"geometries"`
	Speeds                    string `csv:"speeds"`
	Timestamps                string `csv:"timestamps"`
	Idles                     string `csv:"idles"`
	StartDate                 string `csv:"start_date"`
	EndDate                   string `csv:"end_date"`
	TravelDistance            string `csv:"travel_distance"`
	IsCompletePart            string `csv:"is_complete_part"`
	VehicleId                 string `csv:"vehicle_id"`
	DriverId                  string `csv:"driver_id"`
	VehicleConfigId           string `csv:"vehicle_config_id"`
}

// TripPart 0 isLast 0

type NewTrip struct {
	Imei                 string `csv:"imei"`
	VehicleId            string `csv:"vehicle_id"`
	VehicleConfigId      string `csv:"vehicle_config_id"`
	TrackerId            string `csv:"tracker_id"`
	TripId               string `csv:"trip_id"`
	TripPart             string `csv:"trip_part"`
	IsLast               string `csv:"is_last"`
	Latitude             string `csv:"latitude"`
	Longitude            string `csv:"longitude"`
	Timestamp            string `csv:"timestamp"`
	Speed                string `csv:"speed"`
	Idler                string `csv:"idle"`
	Elapsed              string `csv:"elapsed"`
	PerPointIdle         string `csv:"per_point_idle"`
	TraveledDistance     string `csv:"traveled_distance"`
	LocOutlier           string `csv:"loc_outlier"`
	VehicleVersion       string `csv:"vehicle_version"`
	VehicleConfigVersion string `csv:"vehicle_config_version"`
	TrackerVersion       string `csv:"tracker_version"`
	// UpdateId             string `csv:"update_id"`
}
