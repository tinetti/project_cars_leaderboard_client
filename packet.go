package main

import (
    "encoding/binary"
    "bytes"
    "io/ioutil"
)

type Header struct {
    BuildVersionNumber uint16
    RawSequenceNumber  uint8
}

type Packet struct {
    Header          Header
    Telemetry       Telemetry
    Participants    Participants
    ParticipantsExt ParticipantsExt
}

type Telemetry struct {
    GameSessionState           uint8
    ViewedParticipantIndex     int8
    NumParticipants            int8
    // Unfiltered input
    UnfilteredThrottle         uint8
    UnfilteredBrake            uint8
    UnfilteredSteering         int8
    UnfilteredClutch           uint8
    // ?
    RaceStateFlags             RaceState
    LapsInEvent                uint8
    // Timing info
    BestLapTime                float32
    LastLapTime                float32
    CurrentTime                float32
    SplitTimeAhead             float32
    SplitTimeBehind            float32
    SplitTime                  float32
    EventTimeRemaining         float32
    PersonalFastestLapTime     float32
    WorldFastestLapTime        float32
    CurrentSector1Time         float32
    CurrentSector2Time         float32
    CurrentSector3Time         float32
    FastestSector1Time         float32
    FastestSector2Time         float32
    FastestSector3Time         float32
    PersonalFastestSector1Time float32
    PersonalFastestSector2Time float32
    PersonalFastestSector3Time float32
    WorldFastestSector1Time    float32
    WorldFastestSector2Time    float32
    WorldFastestSector3Time    float32
    // Joypad state?
    JoyPad                     uint16
    // Flags
    HighestFlag                uint8
    // Pit schedule
    PitModeSchedule            uint8
    // Car state
    OilTempCelsius             int16
    OilPressureKPa             uint16
    WaterTempCelsius           int16
    WaterPressureKPa           uint16
    FuelPressureKPa            uint16
    CarFlags                   uint8
    FuelCapacity               uint8
    Brake                      uint8
    Throttle                   uint8
    Clutch                     uint8
    Steering                   int8
    FuelLevel                  float32
    Speed                      float32
    Rpm                        uint16
    MaxRpm                     uint16
    GearNumGears               uint8
    BoostAmount                uint8
    EnforcedPitStopLap         int8
    CrashState                 uint8
    // Motion and device
    OdometerKM                 float32
    OrientationX               float32
    OrientationY               float32
    OrientationZ               float32
    LocalVelocityX             float32
    LocalVelocityY             float32
    LocalVelocityZ             float32
    WorldVelocityX             float32
    WorldVelocityY             float32
    WorldVelocityZ             float32
    AngularVelocityX           float32
    AngularVelocityY           float32
    AngularVelocityZ           float32
    LocalAccelerationX         float32
    LocalAccelerationY         float32
    LocalAccelerationZ         float32
    WorldAccelerationX         float32
    WorldAccelerationY         float32
    WorldAccelerationZ         float32
    ExtentsCentreX             float32
    ExtentsCentreY             float32
    ExtentsCentreZ             float32
}

type Participants struct {
    RawCarName         [64]byte
    RawCarClassName    [64]byte
    RawTrackLocation   [64]byte
    RawTrackVariation  [64]byte
    RawNames           [16][64]byte
    RawFastestLapTimes [16]float32
}

type ParticipantsExt struct {

}

type PacketType uint8

const (
    PacketType_TELEMETRY PacketType = iota
    PacketType_PARTICIPANT PacketType = iota
    PacketType_PARTICIPANT_ADDITIONAL PacketType = iota
)

type GameState uint8

const (
    GameState_EXITED GameState = iota
    GameState_FRONT_END GameState = iota
    GameState_INGAME_PLAYING GameState = iota
    GameState_INGAME_PAUSED GameState = iota
)

type SessionState uint8

const (
    SessionState_INVALID SessionState = iota
    SessionState_PRACTICE SessionState = iota
    SessionState_TEST SessionState = iota
    SessionState_QUALIFY SessionState = iota
    SessionState_FORMATION_LAP SessionState = iota
    SessionState_RACE SessionState = iota
    SessionState_TIME_ATTACK SessionState = iota
)

type RaceState uint8

const (
    RaceState_INVALID RaceState = iota
    RaceState_NOT_STARTED RaceState = iota
    RaceState_RACING RaceState = iota
    RaceState_FINISHED RaceState = iota
    RaceState_DISQUALIFIED RaceState = iota
    RaceState_RETIRED RaceState = iota
    RaceState_DNF RaceState = iota
)

func (packet *Packet) Marshal() []byte {
    buffer := &bytes.Buffer{}
    binary.Write(buffer, binary.LittleEndian, packet)
    return buffer.Bytes()
}

func ReadFile(filename string) (*Packet, error) {
    contents, err := ioutil.ReadFile(filename)
    if err != nil {
        return &Packet{}, err
    }

    return Unmarshal(contents)
}

func Unmarshal(raw_message []byte) (*Packet, error) {
    packet := &Packet{}
    buf := bytes.NewReader(raw_message)
    err := binary.Read(buf, binary.LittleEndian, &packet.Header)
    if LogError(err, "unmarshaling packet") {
        return packet, err
    }

    switch (packet.Header.GetPacketType()) {
    case PacketType_TELEMETRY:
        telemetryData := Telemetry{}
        err = binary.Read(buf, binary.LittleEndian, &telemetryData)
        err = WrapError(err, "reading telemetry packet")
        packet.Telemetry = telemetryData

    case PacketType_PARTICIPANT:
        participantInfo := Participants{}
        err = binary.Read(buf, binary.LittleEndian, &participantInfo)
        err = WrapError(err, "reading participant packet")
        packet.Participants = participantInfo

    case PacketType_PARTICIPANT_ADDITIONAL:
        participantInfoAdditional := ParticipantsExt{}
        err = binary.Read(buf, binary.LittleEndian, &participantInfoAdditional)
        err = WrapError(err, "reading ex participant packet")
        packet.ParticipantsExt = participantInfoAdditional
    }

    return packet, err
}

func (header Header) GetSequenceNumber() uint8 {
    return uint8((header.RawSequenceNumber & 0xFC) >> 2)
}

func (header Header) GetPacketType() PacketType {
    return PacketType(header.RawSequenceNumber & 0x3)
}

func toStrings(buffer [16][64]byte) []string {
    strings := []string{}
    for i := 0; i < len(buffer); i++ {
        strings = append(strings, toString(buffer[i]))
    }
    return strings
}

func toString(buffer [64]byte) string {
    n := 0
    for ; buffer[n] != 0; n++ {}
    return string(buffer[:n])
}

func (participantInfo Participants) GetCarName() string {
    return toString(participantInfo.RawCarName)
}

func (participantInfo Participants) GetCarClassName() string {
    return toString(participantInfo.RawCarClassName)
}

func (participantInfo Participants) GetTrackLocation() string {
    return toString(participantInfo.RawTrackLocation)
}

func (participantInfo Participants) GetTrackVariation() string {
    return toString(participantInfo.RawTrackVariation)
}

func (participantInfo Participants) GetNames() []string {
    return toStrings(participantInfo.RawNames)
}

func (participantInfo Participants) GetFastestLapTimes() []float32 {
    return participantInfo.RawFastestLapTimes[:len(participantInfo.RawFastestLapTimes)]
}
