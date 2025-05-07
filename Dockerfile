FROM golang:1.24 AS build
WORKDIR /src
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/gotth-stack ./main.go

FROM scratch
COPY --from=build /bin/gotth-stack /gotth-stack
COPY --from=build /src/static /static
EXPOSE 3000
CMD ["/gotth-stack"]