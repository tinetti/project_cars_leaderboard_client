package main

import (
    "bytes"
    "encoding/json"
    "net/http"
    "fmt"
    "io/ioutil"
)

type ServerWriterHandler struct {
    URL          string
    Username     string
    Password     string

    userId       string
    authToken    string

    LastLapTime  float32
    Participants Participants
}

func (handler *ServerWriterHandler) HandlePacket(packet *Packet) {
    if len(handler.Username) > 0 && len(handler.userId) == 0 {
        err := handler.Login()
        if LogError(err) {
            return
        }
    }

    switch (packet.Header.GetPacketType()) {
    case PacketType_TELEMETRY:
        if packet.Telemetry.LastLapTime != handler.LastLapTime {
            lapTime := NewLapTime(packet.Telemetry, handler.Participants)
            _, _, err := handler.PostLapTime(lapTime)
            LogError(err, "posting lap time")
        }
        handler.LastLapTime = packet.Telemetry.LastLapTime

    case PacketType_PARTICIPANT:
        //fmt.Printf("participant packet -> car_name:%v, car_class:%v, track_location:%v, track_variation:%v\n", packet.Participants.GetCarName(), packet.Participants.GetCarClassName(), packet.Participants.GetTrackLocation(), packet.Participants.GetTrackVariation())
        handler.Participants = packet.Participants
        break

    case PacketType_PARTICIPANT_ADDITIONAL:
    default:
        break
    }
}

func (handler *ServerWriterHandler) PostLapTime(lapTime LapTime) (int, string, error) {
    var err error
    status := -1
    body := ""

    jsonBytes, err := json.Marshal(lapTime)
    if err != nil {
        return status, body, err
    }

    url := handler.URL + "/api/v1/laps"
    fmt.Printf("posting lap time to [ %v ]: %v\n", url, string(jsonBytes))

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
    req.Header.Set("Content-Type", "application/json")
    if len(handler.userId) > 0 {
        req.Header.Set("X-User-Id", handler.userId)
    }
    if len(handler.authToken) > 0 {
        req.Header.Set("X-Auth-Token", handler.authToken)
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()

    status = resp.StatusCode
    if err != nil {
        return status, body, WrapError(err, "posting lap")
    }

    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return status, body, WrapError(err, "reading lap response body")
    }

    body = string(contents)

    if status < 200 || status > 299 {
        err = fmt.Errorf("error posting lap [status:%v, body:%v]", status, body)
    }

    if err == nil {
        fmt.Printf("successfully posted lap (%v): %v\n", status, body)
    }

    return status, body, err
}

func (handler *ServerWriterHandler) Login() error {
    url := handler.URL + "/api/v1/login/"
    reqBodyText := fmt.Sprintf("username=%v&password=%v", handler.Username, handler.Password)
    reqReader := bytes.NewBuffer([]byte(reqBodyText))
    req, err := http.NewRequest("POST", url, reqReader)
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    fmt.Printf("logging in to [ %v ]\n", url)

    client := &http.Client{}
    resp, err := client.Do(req)
    defer resp.Body.Close()

    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return WrapError(err, "reading login response body")
    }

    loginResponse := LoginResponse{}
    err = json.Unmarshal(contents, &loginResponse)
    if err != nil {
        return WrapError(err, "unmarshaling json response")
    }

    if len(loginResponse.Data.UserId) > 0  && len(loginResponse.Data.AuthToken) > 0 {
        handler.userId = loginResponse.Data.UserId
        handler.authToken = loginResponse.Data.AuthToken
        fmt.Println("successfully logged in")
    } else {
        err = fmt.Errorf("error logging in [status=%v, body=%v]", resp.StatusCode, string(contents))
    }

    return err
}

type LoginResponse struct {
    Status string
    Data   LoginResponseData
}

type LoginResponseData struct {
    AuthToken string
    UserId    string
}
