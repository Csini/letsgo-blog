# Go API Server for letsgo-blog

Mini Application providing something to something like a wannabe Blog.

### Running the server

1) build executable run
```
mvn clean install -Passembly
```

executable is under letsgo-bog-server/bin

2) navigate to letsgo-bog-server/bin and start there

3) sqlite db with data is copied there from letsgo-blog/server/res

4) user/pw is userid/userid1 to every user

5) server url is for example

http://localhost:8085/letsgo/v1/api/statistics/200

api documentation is under letsgo-blog-api


### Running the server with docker TODO (not working yet) 


To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```
