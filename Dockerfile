FROM golang:1.22-alpine3.19 AS builder

COPY ${PWD} /app
WORKDIR /app

# Toggle CGO based on your app requirement. CGO_ENABLED=1 for enabling CGO
RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/appbin *.go

FROM alpine:3.19
LABEL MAINTAINER Author <guvian14@gmail.com>

# Add new user 'appuser'
RUN adduser -D appuser
USER appuser

COPY --from=builder /app /home/appuser/app

WORKDIR /home/appuser/app

CMD ["./appbin"]