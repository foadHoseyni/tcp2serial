package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tarm/serial"
)

func main() {
        arguments := os.Args
        if len(arguments) == 1 {
                fmt.Println("Please provide port number")
                return
        }

        PORT := ":" + arguments[1]
        serial_port := arguments[2]
        baud_rate, _ := strconv.Atoi(arguments[3])

        l, err := net.Listen("tcp", PORT)
        if err != nil {
                fmt.Println(err)
                return
        }
        defer l.Close()

        c, err := l.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }
        config := &serial.Config{
                Name: serial_port,
                Baud: baud_rate,
                ReadTimeout: time.Second * 5,
                Size: 8,
        }
        stream, err := serial.OpenPort(config)
        if err != nil {
            log.Fatal(err)
        }
        for {
                netData, err := bufio.NewReader(c).ReadString('\n')
                if err != nil {
                        fmt.Println(err)
                        return
                }
                if strings.TrimSpace(string(netData)) == "STOP" {
                        fmt.Println("Exiting TCP server!")
                        return
                }

                stream.Write([]byte(netData))
                
                serial_data, err := bufio.NewReader(stream).ReadBytes('\n')
                if err != nil {
                        log.Fatal(err)
                        return
                }
                c.Write([]byte(serial_data))                
        }
}
