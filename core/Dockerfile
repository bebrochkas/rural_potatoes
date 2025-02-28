FROM golang:1.22.1-alpine as builder

WORKDIR /app


COPY core/go.mod .
COPY core/go.sum .
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
COPY core/ .
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o /app/app ./cmd/main.go

# Final stage
FROM golang:1.22.1-alpine

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .
COPY --from=builder /app/tagger/tags_data /app/tags_data
COPY --from=builder /app/core/internal/db/csv_migrations /app/csv_migrations

ENTRYPOINT ["./app", "-env", "./.env"]
