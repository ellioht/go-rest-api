FROM docker.io/library/golang:1.21.8-alpine3.19 AS base
RUN apk --no-cache add build-base linux-headers git bash ca-certificates libstdc++
RUN mkdir /api
WORKDIR /api
ADD go.mod .
ADD go.sum .
RUN go mod download


FROM base as builder
WORKDIR /api
RUN ls
ADD . .
RUN ls
RUN go build -o apiexecutable ./cmd/api


FROM docker.io/library/alpine:3.19
COPY --from=builder /api/apiexecutable /apiexecutable

EXPOSE 8080
CMD [ "/apiexecutable", "serve" ]