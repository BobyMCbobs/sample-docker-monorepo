FROM cgr.dev/chainguard/go:1.20 AS build
WORKDIR /app
COPY go.* /app
RUN go mod download
COPY main.go /app
RUN CGO_ENABLED=0 go build -mod=readonly -a -installsuffix cgo -ldflags "-extldflags '-static' -s -w" -o bin/app main.go

FROM cgr.dev/chainguard/static:latest AS product
WORKDIR /app
COPY --from=build /app/bin/app /app/bin/app
USER nonroot
ENTRYPOINT ["/app/bin/app"]
