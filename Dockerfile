FROM alpine
COPY migrations /migrations
ADD whatsforlunch /
ENTRYPOINT ["./whatsforlunch"]
EXPOSE 3000