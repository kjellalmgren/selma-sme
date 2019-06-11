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
#
# COPY all json files
RUN mkdir -p /selmasme/json
COPY json/processes.json /selmasme/json
COPY json/applicants.json /selmasme/json
COPY json/budgets.json /selmasme/json
COPY json/collaterals.json /selmasme/json
COPY json/companyeconomies.json /selmasme/json
COPY json/extloans.json /selmasme/json
COPY json/households.json /selmasme/json
COPY json/loans.json /selmasme/json
COPY json/personaleconomies.json /selmasme/json
COPY json/companies.json /selmasme/json
COPY json/kycinformations.json /selmasme/json
COPY json/eusupports.json /selmasme/json
COPY json/guarantors.json /selmasme/json
COPY json/maintenancecosts.json /selma/json
COPY json/mainpurposes.json /selma/json
COPY json/selma-en-v0.6.9.yaml /selmasme/json
# COPY executable ./selma-sme
COPY selmasme /selmasme
#
# copy our self-signed certificate
COPY cert.pem /selmasme
COPY key.pem /selmasme
#
# tell we are exposing our service on port 8443 for https & 8400 for http
EXPOSE 8443
# run it!
ENTRYPOINT ["./selmasme"]
#CMD ["./selmasme"]