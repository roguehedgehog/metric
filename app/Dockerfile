FROM golang:1.18 as build
WORKDIR /app
COPY go.* .
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/metric ./metric
CMD ["./metric", "serve", "-p", "8080"]
EXPOSE 8080
