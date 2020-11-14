# command-station

A command station for the common nerd.

Create buttons for shell commands/tasks and trigger them from a no-frills web UI.

![Demo](doc/demo.png)

## Usage

### Standalone

```bash
git clone https://github.com/matthewjwhite/command-station && cd command-station
go get -u github.com/go-bindata/go-bindata/... && \
  mkdir asset && \
  go-bindata -o asset/bindata.go --pkg asset template/...
go build .
./command-station sample.yml # Button created for each command.
```

Navigate to `http://127.0.0.1:8000`.

### Docker

```bash
git clone https://github.com/matthewjwhite/command-station && cd command-station
docker build -t command-station .
docker run -it -v "$PWD/sample.yml:/sample.yml" -p 8000:8000 command-station /sample.yml
```

Navigate to `http://127.0.0.1:8000`.
