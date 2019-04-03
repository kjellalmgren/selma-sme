# Selma-SME

## Docker login
    tetracon/Pdx10p2017!

## Static REST-API

##build container

    $ GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selmasme
    # Use this for alpine container
    $ GOOS=linux GOARCH=arm64 go build -v
    # build with tag selma.sme, -t equals tag
    $ docker build --file Dockerfile.builder -t tetracon/selmasme:0.5.8 .
    # push to hub.docker.com, assumes docker login
    $ docker push tetracon/selmasme:0.5.8
    # create netork
    $ docker network create --driver bridge selmasme-net
    # run container version 0.5.8
    $ docker run -d --name selmasme --network selmasme-net --publish=8443:8443 -t tetracon/selmasme:0.5.8
    # run shell to look into container
    $ docker run -d -t tetracon/selmasme:0.5.8 sh
    #
    # Remove container
    $ docker stop [CONTAINERID]
    $ docker rm [CONTAINERID]
    # Remove network
    $ docker network rm selmasme-net
    # look into the logs
    $ docker logs [CONTAINERID]
