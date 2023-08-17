alias b := build
alias i := install_to_home

binary_name := "gops"
home_dir := env_var('HOME')
arch := `uname -m`
os := `uname -s | tr "[:upper:]" "[:lower:]"`

## Builds the binary for linux and darwin on amd64/arm64
build: clean
	GOARCH=amd64 GOOS=linux go build -o bin/{{binary_name}}-linux-amd64 main.go
	GOARCH=arm64 GOOS=linux go build -o bin/{{binary_name}}-linux-arm64 main.go
	GOARCH=amd64 GOOS=darwin go build -o bin/{{binary_name}}-darwin-amd64 main.go
	GOARCH=arm64 GOOS=darwin go build -o bin/{{binary_name}}-darwin-arm64 main.go

# Build the binary for the current OS
_build_current:
	go build -o ./bin/{{binary_name}}-{{os}}-{{arch}} main.go

# Run the binary with the given url
run url:
	go run main.go {{url}}

# Clean up the bin directory
clean:
	rm -f ./bin/{{binary_name}}-*

# Install gops to $HOME/.bin, or to a custom directory if specified
install_to_home bin_dir=(home_dir + '/.bin'): _build_current
	cp bin/{{binary_name}}-{{os}}-{{arch}} {{bin_dir}}/{{binary_name}}
	chmod +x {{bin_dir}}/{{binary_name}}
	just clean