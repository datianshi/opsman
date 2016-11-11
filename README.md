# opsman

## Why create this tool?

* Pivotal Ops Manager has API endpoints
* People want a CLI Tool to interact with OpsManager (I like curl though)

## What does this tool do?

```
COMMANDS:
     token    retieve token
     pivnet   pivnet
     upload   upload
     help, h  Shows a list of commands or help for one command
```

* Retrieve Ops Manager UAA Token

```
./opsman-cli token --opsmanurl OPS_MGR_URL -u admin -p password --skipssl
```
* Show the latest release

```
./opsman-cli pivnet latest-release --productname pcf-metrics
{  
   "Id":2381,
   "Version":"1.1.3",
   "AcceptUrl":"https://network.pivotal.io/api/v2/products/pcf-metrics/releases/2381/eula_acceptance",
   "Files":[  
      {  
         "Name":"PCF Metrics",
         "DownloadUrl":"https://network.pivotal.io/api/v2/products/pcf-metrics/releases/2381/product_files/7569/download"
      },
      {  
         "Name":"PCF Metrics v1.1 OSL",
         "DownloadUrl":"https://network.pivotal.io/api/v2/products/pcf-metrics/releases/2381/product_files/5186/download"
      }
   ]
}
```

* Accept one release

```
./opsman-cli pivnet accept-eula --eulaurl EULA_URL --token XXX
```
* Download a tile from Pivnet

```
 ./opsman-cli pivnet download --producturl https://network.pivotal.io/api/v2/products/apigee-edge-for-pcf-service-broker/releases/1773/product_files/4698/download --token XXXX --dest /tmp/apigee.pivotal
```

* Upload a product tile to OpsManager

```
./opsman-cli upload product --opsmanurl OPS_MGR_URL -u admin -p password --skipssl --from /tmp/apigee.pivotal
```

* Upload a stemcell to OpsManager

```
./opsman-cli upload stemcell --opsmanurl OPS_MGR_URL -u admin -p password --skipssl --from /tmp/stemcell.zip
```
