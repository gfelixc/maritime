FROM golang:1.20-alpine3.17 as build

WORKDIR /src

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /app /src/cmd/port-domain/main.go

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]
