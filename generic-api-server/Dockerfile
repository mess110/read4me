FROM golang:1.22 AS base

ARG WHICH_SERVER=books_api_server

WORKDIR /go/src/github.com/read4me/generic-api-server/

COPY cmd/ cmd/
COPY internal/ internal/
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/$WHICH_SERVER/$WHICH_SERVER.go

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /main .

EXPOSE 3000

CMD ["./main"]
