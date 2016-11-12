FROM scratch
COPY ./wisselgeld /wisselgeld
ENTRYPOINT ["/wisselgeld"]
