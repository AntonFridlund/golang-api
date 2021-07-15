##
## Build
##
## docker build -t golang-api:tags-go-here -f Dockerfile .

FROM golang:1.16.6

WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o golang-api

##
## Deploy
##

FROM alpine:3.14

WORKDIR /
COPY --from=0 golang-api /app/golang-api
EXPOSE 8080
ENTRYPOINT ["/app/golang-api"]

## docker run -d -p 8080:8080 --name golang-api golang-api:tags-go-here
