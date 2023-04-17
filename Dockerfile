#=============================================================================
# STEP 1 build executable binary
#=============================================================================

# Targeting linux amd64, the docker image base needs to match otherwise you will need
# to install a cross compiler when golang ultimately compiles C code.
FROM --platform='linux/amd64' golang:1.19 as builder

# # Install git + SSL ca certificates.
# # Git is required for fetching the dependencies.
# # Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser
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

# WORKDIR /app
COPY . .

# Fetch dependencies.
RUN go mod download
RUN go mod verify

# Build the binary
RUN go build -o /usr/local/bin/backend


# #=============================================================================
# # STEP 2 build a small image
# #=============================================================================

FROM --platform='linux/amd64' gcr.io/distroless/base

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Use an unprivileged user.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER appuser:appuser

EXPOSE 1337

WORKDIR /app

# Copy our static executable
COPY --from=builder /usr/local/bin/backend .

# Run the binary.
CMD ["./backend"]
