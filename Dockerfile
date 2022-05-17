#
# build application
#

FROM golang:1.18-bullseye as build

WORKDIR /build

# download required Go packages
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# copy application files
COPY . .

# build Go application
RUN go build -v -o /build/arena

#
# deploy application
#

FROM debian:bullseye-slim
ENV GIN_MODE=release

# copy freshly built application binary
COPY --from=build /build/arena /usr/local/bin/arena

# create application user
RUN adduser \
    --disabled-password \
    --uid 1001 \
    --gecos "" \
    --shell /bin/sh \
    arena
USER arena

# run production-ready application :)
EXPOSE 8080
ENTRYPOINT ["arena"]
