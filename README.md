# shot-counter-frontend

# Running Go Server

```
gopwd
go run src/github.com/counting-frontend/cmd/counting-frontend/main.go
```

# Making a Request

```
curl -i "localhost:8080/count"
```

# Packaging Go Server into Docker Image and Push it to Docker Hub

```
./buildTagPush.sh
```

# Just building the docker image

```
docker build --no-cache -t shot-counter-frontend -f Dockerfile.build .
```

# Running the Docker Image Locally

```
docker run -it -p 123:8080 shot-counter-frontend
```

# Making a Request Against it

```
curl localhost:123/count
```

# Deploying to K8s

```
cd deploy
./deploy.sh
```

# Mock /sign-in request via curl

NOTE You'll probably need to update the IP

```
curl -i -X POST -d @signin.json 35.227.124.115:8080/sign-in
```

# Mock /count request via curl

NOTE You'll probably need to update the IP

```
curl -i -X POST -d @count.json 35.227.124.115:8080/count
```