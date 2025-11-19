# ---- builder ----
FROM golang:1.20-alpine AS builder
WORKDIR /src

# copy module files first to leverage Docker layer cache
COPY go.mod go.sum ./

# git is needed for some modules that fetch via VCS
RUN apk add --no-cache git

# download modules (cached unless go.mod/go.sum changes)
RUN go mod download

# copy source and build static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# ---- final image ----
FROM alpine:3.18
# create non-root user
RUN addgroup -S app && adduser -S -G app app

# copy binary from builder
COPY --from=builder /app /app

USER app
EXPOSE 8080
ENTRYPOINT ["/app"]

