FROM golang:1.18-alpine AS build

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 go build .

FROM busybox

COPY --from=build /app/ipg /ipg

ENTRYPOINT [ "/ipg" ]
