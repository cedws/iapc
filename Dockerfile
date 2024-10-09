FROM golang:1.23 AS build
ENV CGO_ENABLED 0
WORKDIR /workdir
COPY . .
RUN go build -trimpath

FROM scratch
COPY --from=build /workdir/iapc /usr/bin/iapc
ENTRYPOINT [ "/usr/bin/iapc" ]
