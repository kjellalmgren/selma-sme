# start from hypriot/rpi-alpine-scratch (nginx:alpine)
#
# -------------------------------------------------
# FROM resin/raspbian
# MAINTAINER kjell.almgren[at]tetracon.se
# -------------------------------------------------
#
# FROM resin/rpi-raspbian
FROM alpine

MAINTAINER kjell.almgren@tetracon.se

# make some update to the OS in the container
#RUN apk update && \
#apk upgrade && \
#apk add bash && \
#rm -rf /var/cache/apk/*

#make some changes to the container images (docker dns-bugs)
#COPY docker-compose.yml docker-compose.yaml
#switch to our app directory (/selma-sme)
RUN mkdir -p /selmasme
WORKDIR /selmasme

# COPY all json files
COPY processes.json /selmasme
COPY applicants.json /selmasme
COPY budgets.json /selmasme
COPY collaterals.json /selmasme
COPY companyeconomies.json /selmasme
COPY extloans.json /selmasme
COPY households.json /selmasme
COPY loans.json /selmasme
COPY personaleconomies.json /selmasme
COPY companies.json /selmasme
# COPY executable ./selma-sme
COPY selmasme /selmasme

# copy our self-signed certificate
COPY cert.pem /selmasme
COPY key.pem /selmasme

# tell we are exposing our service on port 8000
EXPOSE 8443

# run it!

ENTRYPOINT ["./selmasme"]
#CMD ["./selmasme"]