FROM golang:1.19.5 AS develop
WORKDIR /app
COPY ./src/ /app
CMD ["go", "run", "main.go"]
FROM develop AS builder
RUN go build -o tracker.app

FROM builder AS test
WORKDIR /app
RUN apt-get update && apt-get -y install cron
COPY --from=develop /app/tracker_output.json /app/tracker-cron /app/start.sh ./
RUN cp ./tracker-cron /etc/cron.d/tracker-cron && chmod +x /app/start.sh && chmod 0644 /etc/cron.d/tracker-cron && crontab /etc/cron.d/tracker-cron && touch ./cron.log
COPY --from=builder /app/tracker.app .
CMD ["cron", "-f"]

FROM ubuntu:23.04 AS production
WORKDIR /app
RUN apt-get update && apt-get -y install cron && apt-get -y install wget
COPY --from=builder /app/tracker_output.json /app/tracker-cron /app/start.sh ./
RUN cp ./tracker-cron /etc/cron.d/tracker-cron && chmod +x /app/start.sh && chmod 0644 /etc/cron.d/tracker-cron && crontab /etc/cron.d/tracker-cron && touch ./cron.log
COPY --from=builder /app/tracker.app .
CMD ["cron", "-f"]