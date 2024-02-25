FROM golang:1.21 AS build
WORKDIR /src
ADD go.mod main.go ./
RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static AS final
COPY --from=build --chown=nonroot:nonroot /src/app /app
ENTRYPOINT ["/app"]
