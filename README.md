Running on docker locally:
```
docker build -t test
docker run -t test
```
Docker local logs:
```
2020/07/24 21:39:12 Proxy listening on port: 8080
[2020-07-24 21:39:12 +0000] [7] [INFO] Starting gunicorn 20.0.4
[2020-07-24 21:39:12 +0000] [7] [INFO] Listening at: http://0.0.0.0:54321 (7)
[2020-07-24 21:39:12 +0000] [7] [INFO] Using worker: threads
[2020-07-24 21:39:12 +0000] [16] [INFO] Booting worker with pid: 16
```

Deploying to Cloud Run (fully managed):
```
docker build -t gcr.io/<some gcr bucket>
docker push gcr.io/<some gcr bucket>
gcloud run deploy --image gcr.io/<some gcr bucket> --platform managed --region us-central1 --allow-unauthenticated mockscf --set-env-vars PROXY_PORT=8080 --port 8080
```

Cloud Run (fully managed) logs
```
2020-07-24 14:27:14.612 PDT2020/07/24 21:27:14 Proxy listening on port: 8080
```