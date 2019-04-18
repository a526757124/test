GOOS=linux go build main.go

GOOS=linux go build -ldflags "-w" -o main main.go
