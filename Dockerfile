FROM golang AS builder
WORKDIR /go/src
COPY . .
RUN CGO_ENABLED=0 go build .

FROM alpine
COPY --from=builder /go/src/command-station .
EXPOSE 8000/tcp
ENTRYPOINT ["./command-station"]
