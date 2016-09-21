# opsman

## Why create this tool?

* Pivotal Ops Manager has API endpoints
* People want a CLI Tool to interact with OpsManager (I like curl though)

## What does this tool doing?

* Retrieve Ops Manager UAA Token

```
./opsman-cli token --opsmanurl OPS_MGR_URL -u admin -p password --skipssl
```

* Download a tile from Pivnet

```
 ./opsman-cli download --producturl https://network.pivotal.io/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download --token XXXX --dest /tmp/apigee.pivotal
```

* Upload a tile to OpsManager

```
./opsman-cli upload --opsmanurl OPS_MGR_URL -u admin -p password --skipssl --from /tmp/apigee.pivotal
```
