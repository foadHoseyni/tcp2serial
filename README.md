# tcp2serial
Converts TCP to Serial in Golang
with this "go run tcp2serial 1234 Com1 9600", the program creates an tcp server and start listening to port #1234 and ready to receive messages and sent it to serial port Com1
with baudrate 9600

with "go run tcpC.go 127.0.0.1:1234", the client is ready to send messages to tcp server.
messages send with Enter.

for the serial side of tcp2serial, you'll need to create a paired virtual serial port, assign one port to tcp2serial, and read and write messages from another port using
some serial monitoring software like https://www.hw-group.com/software/hercules-setup-utility
