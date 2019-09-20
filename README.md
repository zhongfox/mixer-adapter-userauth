# Istio Mixer Check Adapter Demo

![image-20190920165024485](https://zhongfox-blogimage-1256048497.cos.ap-guangzhou.myqcloud.com/2019-09-20-085224.png)

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

6. 按照`https://venilnoronha.io/seamless-cloud-native-apps-with-grpc-web-and-istio` 搭建 grpc-web demo

7. install handler, instance and rule

    ```
    kubectl -n istio-system apply -f testdata/operator-cfg.yaml
    ```

## test



1. Login as fox (cookie: user=fox)

   ![image-20190920165141136](https://zhongfox-blogimage-1256048497.cos.ap-guangzhou.myqcloud.com/2019-09-20-085148.png)

2. Login as other user:

   ![image-20190920165215683](https://zhongfox-blogimage-1256048497.cos.ap-guangzhou.myqcloud.com/2019-09-20-085218.png)

