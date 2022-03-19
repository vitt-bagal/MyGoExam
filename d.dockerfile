# Base image
FROM golang:1.17-alpine

# Add Maintainer Info
LABEL maintainer="Vitthal Bagal <bagalvitthal94@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# COPY code to workspace 
COPY . .

# Build service
RUN go mod download && go build -o /csvparser-service

# Expose port to service
EXPOSE 8080

# Default entrypoint to service
CMD [ "/csvparser-service" ]
# End of Dockerfile