# ---- build stage ----
    FROM golang:1.23.3-alpine AS build
    ENV GOTOOLCHAIN=auto
    RUN apk add --no-cache build-base ca-certificates
    WORKDIR /src
    
    # Cache modules
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source
    COPY . .
    
    # Build the catalog service (adjust path if needed)
    RUN go build -o /out/app ./catalog/cmd/catalog    
    # ---- runtime stage ----
    FROM alpine:3.20
    WORKDIR /app
    COPY --from=build /out/app /usr/bin/app
    EXPOSE 8080
    CMD ["app"]
    