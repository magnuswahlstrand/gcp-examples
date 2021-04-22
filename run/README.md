
#### Tag an image

```bash
docker tag gcp-run-hello gcr.io/bytesgcptemplate/hello:v1;  docker push gcr.io/bytesgcptemplate/hello:v1
```

#### Deploy an image
```bash
gcloud run deploy hello --image gcr.io/bytesgcptemplate/hello:v1
```

#### Set environment variable
```bash
gcloud run services update --platform=managed --region=europe-west1 --update-env-vars NAME=Magnus
```


#### Things to explore
* [ ] Terraform or similar for setup 
* [ ] Use secret manager for environment vars

### Deploy

```
NAME=$(basename "$PWD")
TAG=gcr.io/bytesgcptemplate/$NAME
docker build -t $TAG .  
docker push $TAG

gcloud run deploy $NAME --image $TAG --platform=managed --region=europe-west1
```