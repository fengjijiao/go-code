package main

import(
    "flag"
    "github.com/jinzhu/configor"
    "os"
    "fmt"
    hoststat "github.com/likexian/host-stat-go"
    "net/http"
    "encoding/json"
)

var ConfigPATH string

var Config struct {
    ServerPort string `default:":8090"`
    TrafficDevice string `default:"eth0"`
}

type deviceTraffic struct {
    TXByteString   string      `json:"txBytes"`
    RXByteString   string      `json:"rxBytes"`
    TXPacketString  string       `json:"rxPackets"`
    RXPacketString  string       `json:"txPackets"`
}

func init() {
    flag.StringVar(&ConfigPATH, "c", "config.yml", "config file path")
    flag.Parse()
    if !IsFile(ConfigPATH) || !Exists(ConfigPATH) {
        fmt.Printf("Failed to find configuration %s\n", ConfigPATH)
        os.Exit(3)
        return
    }
    configor.Load(&Config, ConfigPATH)
}

func main() {
    http.HandleFunc("/traffic-api", trafficHttpHandler)
    http.HandleFunc("/", defaultHttpHandler)
    http.ListenAndServe(Config.ServerPort, nil)
    
}

func trafficHttpHandler(w http.ResponseWriter, req *http.Request) {
    netStat,err := hoststat.GetNetStat()
    if err != nil {
        fmt.Println("err")
        return
    }
    for _, ns := range netStat {
        //fmt.Printf("%d：%v\n", i, ns)
        if ns.Device == Config.TrafficDevice {
            dt := &deviceTraffic{
                TXByteString : elegantByteValue(ns.TXBytes),
                RXByteString : elegantByteValue(ns.RXBytes),
                TXPacketString : elegantPacketValue(ns.TXPackets),
                RXPacketString : elegantPacketValue(ns.RXPackets)}
            resJson, _ := json.Marshal(dt)
            fmt.Fprintf(w, string(resJson))
            return
        }
    }
    fmt.Fprintf(w, "{}")
}

func defaultHttpHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello, world!")
}
func elegantByteValue(tr uint64) string {
    if tr >= 1024*1024*1024*1024 {
        return fmt.Sprintf("%.2fTib",float64(tr)/(1024*1024*1024*1024))
    }else if tr >= 1024*1024*1024 {
        return fmt.Sprintf("%.2fGib",float64(tr)/(1024*1024*1024))
    }else if tr >= 1024*1024 {
        return fmt.Sprintf("%.2fMib",float64(tr)/(1024*1024))
    }else if tr >= 1024 {
        return fmt.Sprintf("%.2fKib",float64(tr)/(1024))
    }else {
        return fmt.Sprintf("%.2fByte",float64(tr))
    }
}

func elegantPacketValue(tr uint64) string {
    if tr >= 1000*1000*1000*1000 {
        return fmt.Sprintf("%.2fT",float64(tr)/(1000*1000*1000*1000))
    }else if tr >= 1000*1000*1000 {
        return fmt.Sprintf("%.2fB",float64(tr)/(1000*1000*1000))
    }else if tr >= 1000*1000 {
        return fmt.Sprintf("%.2fM",float64(tr)/(1000*1000))
    }else if tr >= 1000 {
        return fmt.Sprintf("%.2fK",float64(tr)/(1000))
    }else {
        return fmt.Sprintf("%d",tr)
    }
}

//utility

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
    _, err := os.Stat(path)    //os.Stat获取文件信息
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
    s, err := os.Stat(path)
    if err != nil {
        return false
    }
    return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
    return !IsDir(path)
}
