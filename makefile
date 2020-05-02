server:
	go build -o ./bin/server ./main.go
windowsserver:
	env GOOS=windows GOARCH=amd64 go build -o ./bin/server.exe ./main.go