FROM golang:1.24-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd/server && go build -o /bin/favorites

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=build /bin/favorites /bin/favorites
EXPOSE 8080
CMD ["/bin/favorites"]
