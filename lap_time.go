package main

type LapTime struct {
    CarName        string
    CarClassName   string
    TrackLocation  string
    TrackVariation string
    Names          []string
}

func CreateLapTime(packet Packet) LapTime {
    lapTime := LapTime{}
    lapTime.CarName = packet.Participants.GetCarName()
    lapTime.CarClassName = packet.Participants.GetCarClassName()
    lapTime.TrackLocation = packet.Participants.GetTrackLocation()
    lapTime.TrackVariation = packet.Participants.GetTrackVariation()
    lapTime.Names = packet.Participants.GetNames()

    return lapTime
}
