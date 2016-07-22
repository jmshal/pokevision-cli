package main

import (
    "fmt"
    "time"
)

const (
    UPDATE_MAP_DATA_URL = "https://pokevision.com/map/scan/%v/%v"
    CHECK_UPDATE_MAP_DATA_URL = "https://pokevision.com/map/data/%v/%v/%v"
)

func UpdatePokemonInRegion(lat, lon float64) (error) {
    data, err := RequestAPI(fmt.Sprintf(UPDATE_MAP_DATA_URL, lat, lon))
    if err != nil {
        return err
    }
    jobId := data["jobId"].(string)
    for {
        check, err := RequestAPI(fmt.Sprintf(CHECK_UPDATE_MAP_DATA_URL, lat, lon, jobId))
        if err != nil {
            return err
        }
        if check["jobStatus"] == "in_progress" {
            // Try to check the job status in another few seconds
            time.Sleep(time.Second * 3)
        } else {
            // Job has complete
            break
        }
    }
    return nil
}
