#############
# BUILD APP #
#############
FROM registry-nexus.gid.team/golang:1.22.0-bullseye AS builder
ARG GITLAB_LOGIN="its-not-used-in-auth-just-dummy-value"
ARG GITLAB_TOKEN

ENV GOPRIVATE="gitlab.gid.team/*"
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /gid
COPY . .

RUN echo "machine gitlab.gid.team\nlogin $GITLAB_LOGIN\npassword $GITLAB_TOKEN" > /root/.netrc
RUN go clean --modcache
RUN go mod download
RUN go mod verify
RUN go build -o /bin/app ./cmd/grpc


###########
# RUN APP #
###########

FROM registry-nexus.gid.team/golang:1.22.0-bullseye AS app
WORKDIR /gid

COPY app.toml /gid
COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["/bin/app"]
