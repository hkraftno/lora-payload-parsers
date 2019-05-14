# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12-alpine as builder

RUN apk update
# Add required certificates to be able to call HTTPS endpoints.
RUN apk add --no-cache ca-certificates git
# Add git for installing go dependencies
RUN apk add --no-cache git

# Copy local code to the container image.
WORKDIR /app
COPY . .

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build.
RUN go mod download

# CGO_ENABLED=0 == Don't depend on libc (bigger but more independent binary)
# installsuffix == Cache dir for non cgo build files
RUN CGO_ENABLED=0 go build -installsuffix 'static' -v -o main

FROM scratch
WORKDIR /app

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /app/main .

CMD ["./main"]
