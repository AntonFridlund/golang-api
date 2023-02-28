##
## Build
##
## docker build -t golang-api:tags-go-here -f Dockerfile .

FROM golang:1.20.1-bullseye AS build

WORKDIR /
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o golang-api

##
## Deploy
##

FROM scratch

WORKDIR /app
COPY --from=build /golang-api /app/golang-api

# Set the PORT environment variable
ENV PORT=8080

EXPOSE 8080
ENTRYPOINT ["/app/golang-api"]

## docker run -d -p 8080:8080 --name golang-api golang-api:tags-go-here