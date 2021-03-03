FROM golang:1.16-alpine as build

WORKDIR /ws/betting
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o /out/betting .

FROM alpine:latest
WORKDIR /ws/betting
COPY --from=build /out/betting /ws/betting/betting
COPY --from=build /ws/betting/migrations /ws/betting/migrations
EXPOSE 8080
ENTRYPOINT ["/ws/betting/betting"]