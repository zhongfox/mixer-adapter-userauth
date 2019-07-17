
## install

1. install istio

2. make sure istio mixer check enabled, see: <https://istio.io/docs/tasks/policy-enforcement/enabling-policy/>

3. install authorization template

   ```
   kubectl -n istio-system apply -f config/authorization-template.yaml
   ```

4. install mixer adapter: userauth

    ```
    kubectl -n istio-system apply -f config/userauth-adapter.yaml
    ```

5. install userauth deployment and service

    ```
    kubectl -n istio-system apply -f testdata/userauth-service.yaml
    ```

6. install handler, instance and rule

    ```
    kubectl -n istio-system apply -f testdata/operator-cfg.yaml
    ```

## test

1. create a namespace for test, and enable sidecar auto injection

    ```
    kubectl create namespace foo
    kubectl label namespace foo istio-injection=enabled
    ```
2. install demo apps: sleep and httpbin

    ```
    kubectl -nfoo apply -f testdata/demo-apps.yaml
    ```
3. login to sleep pod, try to curl httpbin service

    login:

    ```
    export SLEEP_POD=$(kubectl -nfoo get pod -l app=sleep -o jsonpath={.items..metadata.name})

    kubectl -nfoo exec $SLEEP_POD -c sleep -- sh
    ```

    try curl without token:

    ```
    / # curl -i httpbin:8000/headers
    HTTP/1.1 403 Forbidden
    content-length: 63
    content-type: text/plain
    date: Wed, 17 Jul 2019 16:45:54 GMT
    server: envoy
    x-envoy-upstream-service-time: 2

    PERMISSION_DENIED
    ```

    try curl with valid token:
    ```
    / # curl -i -H "x-custom-token:admintoken" httpbin:8000/headers
    HTTP/1.1 200 OK
    server: envoy
    date: Wed, 17 Jul 2019 16:45:14 GMT
    content-type: application/json
    content-length: 340
    access-control-allow-origin: *
    access-control-allow-credentials: true
    x-envoy-upstream-service-time: 4

    {
      "headers": {
        "Accept": "*/*",
        "Content-Length": "0",
        "Host": "httpbin:8000",
        "User-Agent": "curl/7.60.0",
        "X-B3-Parentspanid": "28c9c15be81624a9",
        "X-B3-Sampled": "1",
        "X-B3-Spanid": "5b8a6e948d4b1a49",
        "X-B3-Traceid": "e66d3ece8d65398f28c9c15be81624a9",
        "X-Custom-Token": "admintoken"
      }
    }
    ```

    try curl with invalid token:

    ```
    / # curl -i -H "x-custom-token:hackertoken" httpbin:8000/headers
    HTTP/1.1 403 Forbidden
    content-length: 63
    content-type: text/plain
    date: Wed, 17 Jul 2019 16:51:02 GMT
    server: envoy
    x-envoy-upstream-service-time: 2

    PERMISSION_DENIED
    ```
