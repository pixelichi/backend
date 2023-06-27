VERSION 0.7

builder:
  FROM --platform='linux/amd64' golang:1.19
  ENV USER=appuser
  ENV UID=10001

  # See https://stackoverflow.com/a/55757473/12429735RUN
  RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"

  WORKDIR /app
  COPY . .

  # Fetch dependencies.
  RUN go mod download
  RUN go mod verify

  # Build the binary
  RUN go build -o /usr/local/bin/backend

  SAVE ARTIFACT /etc/passwd
  SAVE ARTIFACT /etc/group
  SAVE ARTIFACT /usr/local/bin/backend

final-image:
  ARG ENV="prod"
  FROM --platform='linux/amd64' gcr.io/distroless/base

  # Use an unprivileged user.
  COPY +builder/passwd /etc/passwd
  COPY +builder/group /etc/group
  USER appuser:appuser

  EXPOSE 1337

  WORKDIR /app

  # Copy our static executable
  COPY +builder/backend .
  COPY env/$ENV.env .env

  # Run the binary.
  CMD ["/app/backend"]
  SAVE IMAGE backend:latest
