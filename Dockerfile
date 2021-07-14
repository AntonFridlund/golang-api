##
## Build
##
## docker build -t golang-api:multistage -f Dockerfile .

FROM golang:1.16.6 AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /bin/app .

##
## Deploy
##

FROM alpine:3.14

WORKDIR /
COPY --from=build /bin/app /bin/app
EXPOSE 8080

ENTRYPOINT ["/bin/app"]

##
## Run
##
## docker run -d -p 8080:8080 --name golang-api golang-api:multistage