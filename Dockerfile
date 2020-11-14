FROM golang AS builder
WORKDIR /go/src
COPY . .
RUN go get -u github.com/go-bindata/go-bindata/... && \
  mkdir asset && \
  go-bindata -o asset/bindata.go --pkg asset template/...
RUN CGO_ENABLED=0 go build .

FROM alpine
COPY --from=builder /go/src/command-station .
EXPOSE 8000/tcp
ENTRYPOINT ["./command-station"]
