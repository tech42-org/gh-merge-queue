FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/pets .

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/pets /pets
EXPOSE 8080
ENTRYPOINT ["/pets"]
