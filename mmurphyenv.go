package mmurphyenv

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Config interface {
    Get(string) string
}

type AppConfig struct {
    data map[string]string
}

func (ac AppConfig) Get(key string) string {
    return ac.data[key]
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
        if len(chunks) == 2 {
            config.data[chunks[0]] = chunks[1]
        }
    }
    return config
}

func main() {
    config := NewAppConfig("configfile.txt")
    fmt.Println(config.Get("keyname"))
    fmt.Println(config.Get("otherkey"))
}
