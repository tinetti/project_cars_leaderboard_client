package main

type LapTime struct {
    LapTime        float32 `json:"lap_time"`
    CarName        string  `json:"car_name"`
    CarClassName   string  `json:"car_class_name"`
    TrackLocation  string  `json:"track_location"`
    TrackVariation string  `json:"track_variation"`
}

func NewLapTime(telemetry Telemetry, participants Participants) LapTime {
    lapTime := LapTime{}
    lapTime.LapTime = telemetry.LastLapTime

    lapTime.CarName = participants.GetCarName()
    lapTime.CarClassName = participants.GetCarClassName()
    lapTime.TrackLocation = participants.GetTrackLocation()
    lapTime.TrackVariation = participants.GetTrackVariation()

    return lapTime
}
