# Get Go image from DockerHub.
FROM golang:1.19 AS api

# Set working directory.
WORKDIR /compiler

# Copy dependency locks so we can cache.
COPY go.mod go.sum ./

# Get all of our dependencies.
RUN go mod download

# Copy all of our remaining application.
COPY . .

# Build our application.
RUN go build -o cake-store-api ./main.go

# Use 'scratch' image for super-mini build.
FROM scratch AS dev

# Set working directory for this stage.
WORKDIR /dev

# Copy our compiled executable and allfolders from the last stage.
COPY --from=api /compiler/ .

# Run application and expose port 8000.
EXPOSE 8000
CMD ["./cake-store-api"]