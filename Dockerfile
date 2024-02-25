FROM  --platform= --platform=arm64 golang:1.21 AS build
WORKDIR /src
ADD go.mod main.go ./
RUN go build -o app

FROM gcr.io/distroless/static AS final
COPY --from=build --chown=nonroot:nonroot /go-app-1 /app
ENTRYPOINT ["/app"]
