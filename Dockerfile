# syntax=docker/dockerfile:1

FROM golang:1.19 as build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY . .
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM scratch

# Copy ca-certs for app web access
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /app/app /app

# Expose port
EXPOSE 8080

# Set the entry point
ENTRYPOINT ["/app"]
