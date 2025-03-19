Go WebDev
=========

Implementing my way through https://gowebexamples.com/ and also adding in
Docker and other parts along the way.


Docker
------

I've implemented a multi-stage build Dockerfile. This builds the smallest
container image I can in Go. Stage one is the build tools. Then stage two
copies in the binary for the production image. I'm amazed how small this is
compareed to previous python based images. This is understandable given the
basic server is standalone. It also doesn't need libc or other dynamic linked
libraries::

    # Build and product production image
    docker build --target build-env .
    docker build --target runtime --tag basic .

    # Running the production image
    docker run --rm -p 8080:8080 -it basic-latest

    $ docker images
    REPOSITORY      TAG       IMAGE ID       CREATED          SIZE
    basic-alpine    latest    3d80ec994a88   18 minutes ago   25MB
    basic-scratch   latest    264abaac06ff   39 minutes ago   12.1MB
    :