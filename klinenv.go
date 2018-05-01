package klinenv

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

type Config interface {
    Get(string) (string, error)
}

type AppConfig struct {
    data map[string]string
}

func (ac AppConfig) Get(key string) (string, error) {
    value, ok := ac.data[key]
    if !ok {
        return "", fmt.Errorf("attempt to read nonexistent config value %s", key)
    }
    return value, nil
}
func NewAppConfig(filename string) AppConfig {
    fin, err := os.Open(filename)
    if err != nil {
        panic("error opening config file: " + err.Error())
    }
    defer func() {
        cerr := fin.Close()
        if cerr != nil {
            panic("error closing config file: " + cerr.Error())
        }
    }()
    scanner := bufio.NewScanner(fin)
    scanner.Split(bufio.ScanLines)
    config := AppConfig{}
    config.data = make(map[string]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
        chunks := strings.Split(line, "=")

        if len(chunks) != 2 {
    //        isfirstquote := string(chunks[1][0])
    //        islastquote := string(line[len(line)-1])
    //        if isfirstquote == "\"" && islastquote == "\"" {
            final := line[len(chunks[0])+1:]
            config.data[chunks[0]] = strings.Trim(final,"\"")
            continue
    //        }
            log.Fatal("config error at ",chunks[0])
        } else {
       //     isfirstquote := string(chunks[1][0])
       //     islastquote := string(line[len(line)-1])
       //     if isfirstquote == "\"" && islastquote == "\"" {
       //         final := line[len(chunks[0])+1:]
       //         config.data[chunks[0]] = final[1:len(final)-1]
       //         continue
       //     }
            config.data[chunks[0]] = strings.Trim(chunks[1],"\"")
        }
    }
    return config
}
//func main() {
//    config := NewAppConfig("configfile.txt")
//    x, err := config.Get("workdir")
//    if err != nil {
//       panic(err)
//   }
//    fmt.Println(x)
//}


