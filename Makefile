
run-cli:
	go run nube-cli.go


compile:
	rm -r bin || true
	mkdir "bin"
	echo "Compiling"
	GOOS=linux GOARCH=arm go build -o bin/nube-cli-linux-arm nube-cli.go
	GOOS=linux GOARCH=arm go build -o bin/download-bios-linux-arm download-bios.go

zip:
	cp cli.yml bin/
	zip -r bin/cli-build.zip bin




all: compile zip