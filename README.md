# Selma-SME

## Docker login
    tetracon/Pdx10p2017!

## Static REST-API

##build container

    $ GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' -a -installsuffix cgo -o selmasme
    # build with tag selma.sme, -t equals tag
    $ docker build --file Dockerfile.builder -t tetracon/selmasme:0.4.8 .
    # push to hub.docker.com, assumes docker login
    $ docker push tetracon/selma-sme:0.4.8
    # run container version 0.4.8
    $ docker run -it --name selmasme --publish=8000:8000 -t tetracon/selmasme:0.4.8
    # run shell to look into cantainer
    $ docker run -it -t tetracon/selmasme:0.4.8 sh

    # docker service create --name=pingservices --publish=80:9000 tetracon/pingservices:2.19
