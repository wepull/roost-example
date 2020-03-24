# Deploy a simple Node.js app on Kubernetes (GKE)
 It will be executed on localhost:4000 in browser.
 User need to give a secret along with a message in browser and the message will be encrypted and cipher text will be genereated.
 On another tab in the browser the cipher text along with secret should be given as input. Then given message will be displayed on the browser as the output

## Deploy on Kubernetes (GKE)


[Read here](https://medium.com/@onufrienkos/deploying-a-node-js-app-to-the-google-kubernetes-engine-gke-d6af1f3a954c)


## Run

```shell
git clone https://github.com/sonufrienko/gke-simple-app
cd gke-simple-app/app
npm i
npm run start
```


## App

The app built with Node.js and allow AES encryption/decryption using HTTP request.

#### Browser

Encrypt a "message" with "secret"

```http://localhost:4000/encrypt?secret=8650&message=i-love-you```

Decrypt a "message" with "secret"

```http://localhost:4000/decrypt?secret=8650&message=12840030619419b8d8ec4fe61e275d99```

#### CURL

Encrypt a "message" with "secret"

```shell
curl -G 'http://localhost:4000/encrypt' \
-d secret=8650 \
-d message=i-love-you
```

Decrypt a "message" with "secret"

```shell
curl -G 'http://localhost:4000/decrypt' \
-d secret=8650 \
-d message=12840030619419b8d8ec4fe61e275d99
```
