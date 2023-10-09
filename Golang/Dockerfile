# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build
# Set the Current Working Directory inside the container
WORKDIR /app
RUN go get -d -v ./...

# Build the Go application
RUN go install -v ./...
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -ldflags='-s' .

## Deploy ##
FROM gcr.io/distroless/static-debian11
WORKDIR /app
COPY --from=build /app /app
USER nonroot:nonroot

# ENTRYPOINT ["sh", "-c", "source env.sh && \"$@\"", "-s"]
EXPOSE 5001 8080

CMD [ "./recordsApp" ]
