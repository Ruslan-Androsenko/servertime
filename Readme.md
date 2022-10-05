# Building and running in your container
```
docker build -t servertime .
docker run --publish 8080:8080 --name servertime --rm servertime
```
