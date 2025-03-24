Go WebDev
=========

Implementing my way through https://gowebexamples.com/ and also adding in
Docker, Docker Compose and other bits and things I expect along the way.


Docker
------

I've implemented a multi-stage build Dockerfile. This builds the smallest
container image I can in Go. Stage one is the build tools. Then stage two
copies in the binary for the production image. I'm amazed how small this is
compareed to previous python based images. This is understandable given the
basic server is standalone. It also doesn't need libc or other dynamic linked
libraries::

    # Build and product production image
    docker build --target production --tag basic .

    # Running the production image
    docker run --rm -p 8080:8080 -it basic

    $ docker images
    REPOSITORY      TAG       IMAGE ID       CREATED          SIZE
    basic           latest    3d80ec994a88   18 minutes ago   25MB
    basic-scratch   latest    264abaac06ff   39 minutes ago   12.1MB
    :


Development
-----------

I'm adding in docker compose to manage the build and database dependancies.

::
    # build and run all parts
    docker compose up

    # or just the database
    docker compose up -d db

    # building and running the web server outside of docker compose
    ./build.sh && ./web-server --database-dsn postgres://postgres:password@127.0.0.1:7432/webdev

    $ PGPASSWORD=password psql -h localhost -p 7432 -U postgres -d webdev
    psql (17.4, server 17.2 (Debian 17.2-1.pgdg120+1))
    Type "help" for help.

    webdev=# select * from users;
    id | username | password | created_at
    ----+----------+----------+------------
    (0 rows)

    webdev=#

    # Run all tests which will test against the running postgreql. It will
    # run the tests in their own isolated DB instances.
    go test ./...

I'm using docker compose and postgres as the database. There original tutorial
uses MySQL, however I prefer to work with Postgres instead. For test I'm using
https://github.com/peterldowns/pgtestdb
