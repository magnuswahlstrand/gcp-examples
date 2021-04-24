
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

### Deploy

```
NAME=$(basename "$PWD")
TAG=gcr.io/bytesgcptemplate/$NAME
docker build -t $TAG .  
docker push $TAG

gcloud run deploy $NAME --image $TAG --platform=managed --region=europe-west1
```

or shorted

```
NAME=$(basename "$PWD")
TAG=gcr.io/bytesgcptemplate/$NAME
gcloud builds submit --tag $TAG
gcloud run deploy $NAME --image $TAG --platform=managed --region=europe-west1
```

### Set defaults

```
gcloud config set run/platform managed
gcloud config set run/region europe-west1
```


#### Things to explore
* [ ] Use secret manager for environment vars
* [ ] Service to service authentication
  *   https://cloud.google.com/run/docs/authenticating/service-to-service#gcloud
* [ ] Setup continuous deployment
* [ ] Setup staging enviroment