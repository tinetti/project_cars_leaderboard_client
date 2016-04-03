package main

type LapTime struct {
    LapTime        float32 `json:"lapTime"`
    CarName        string  `json:"carName"`
    CarClassName   string  `json:"carClassName"`
    TrackLocation  string  `json:"trackLocation"`
    TrackVariation string  `json:"trackVariation"`
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
