# tempcontroller_go
Rpi Temperature Controller

# To compile for RPI B+
GOOS=linux GOARCH=arm GOARM=6 go build  -a -ldflags '-extldflags "-static"' .

