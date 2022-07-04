# Go Training

## Install Go on ubuntu

```wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz```

```sudo tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz```

- This runs only in the current session

```export PATH=$PATH:/usr/local/go/bin```

- open .profile

```vim .profile```

- add path env variable to the end of the .profile file

```export PATH=$PATH:/usr/local/go/bin```

```source .profile```

```go version```

## Go Tools

### Run

- ```go run main.go```
- ```go run --work main.go```

### Compile

- ```go tool compile main.go```
- ```go tool compile -S main.go```
  
### Build

- ```go build main.go```
- ```go build -o app main.go```

#### Cross compilation and build

- list of all go cross compilation os and archs
- ```go tool dist list```
- ```GOARCH=amd64 GOOS=windows go build main.go```

### Install

- ```go install main.go```
- ```go install github.com/cweill/gotests/gotests@latest```

### Mod

- ```go mod init demo```
- ```go mod tidy```
