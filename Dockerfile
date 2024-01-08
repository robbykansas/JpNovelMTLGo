# ImageBuilder
FROM golang:1.19.4-buster as builder

LABEL maintainer="robby"

# CREATE A WORKING DIRECTORY
RUN mkdir /app
WORKDIR /app

#COPY SOURCE CODE
COPY . .

#BUILD BINARY FILE
RUN go build -ldflags '-linkmode=external' -o jpnovelmtl cmd/main/main.go

# Distribution Image Debian
FROM debian:buster-slim

LABEL maintainer="robby"


# SET TIMEZONE
ENV TZ="Asia/Jakarta"
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone && dpkg-reconfigure -f noninteractive tzdata

#CREATE WORKDIR
RUN mkdir /app
WORKDIR /app

RUN chgrp -R 0 /app && \
    chmod -R g=u /app

#COPY BINARY FILE FROM BUILDER
COPY --from=builder /app/jpnovelmtl /app
COPY --from=builder /app/.env /app

EXPOSE 3000
CMD /app/jpnovelmtl