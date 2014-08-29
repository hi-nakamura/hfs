package main

import (
    "fmt"
    "net"
    "net/http"
    "bufio"
    "os"
    "github.com/BurntSushi/toml"
)

type Config struct {
    Port string
    Path string
}

func wait() {
    reader := bufio.NewReader(os.Stdin)
    reader.ReadString('\n')
}

func output_error(err error) {
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}

func start_server(path, port string) (err error) {
    fs := http.FileServer(http.Dir(path))
    err = http.ListenAndServe(port, fs)
    return
}

func main() {
    var config Config
    var err error

    // load setting file
    _, err = toml.DecodeFile("setting.tml", &config)
    output_error(err)

    // information
    fmt.Println("Http File Server")
    ifaces, err := net.Interfaces()
    output_error(err)
    for _, iface := range ifaces {
        addrs, err := iface.Addrs()
        output_error(err)
        for _, addr := range addrs {
            switch addr.(type) {
            case *net.IPAddr:
                fmt.Println("IP address:", addr)
            }
        }
    }
    fmt.Println("Port:", config.Port)
    fmt.Println("Start.")

    // start server
    err = start_server(config.Path, config.Port)
    output_error(err)

    wait()
}
