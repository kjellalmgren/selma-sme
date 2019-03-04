# Selma-SME

## Docker login
    tetracon/Pdx10p2017!

## Static REST-API

##build container

    $ GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selmasme
    # build with tag selma.sme, -t equals tag
    $ docker build --file Dockerfile.builder -t tetracon/selmasme:0.5.1 .
    # push to hub.docker.com, assumes docker login
    $ docker push tetracon/selmasme:0.5.1
    # create netork
    $ docker network create --driver bridge selmasme-net
    # run container version 0.5.1
    $ docker run -d --name selmasme --network selmasme-net --publish=8000:8000 -t tetracon/selmasme:0.5.1
    # run shell to look into container
    $ docker run -d -t tetracon/selmasme:0.5.1 sh

    # docker service create --name=pingservices --publish=80:9000 tetracon/pingservices:2.19
