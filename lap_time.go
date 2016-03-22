package main

type LapTime struct {
    LapTime        float32
    CarName        string
    CarClassName   string
    TrackLocation  string
    TrackVariation string
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
