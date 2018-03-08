# Oracle Cloud Infrastructure Golang SDK
This is the official go sdk for oracle cloud infrastructure

## Installing
Simply clone this repo into your go sdk you can use the following command

```
git clone git@github.com:oracle/oci-go-sdk.git  $GOPATH/src/github.com/oracle
```

In addition you'll need to download a testing dependency before you can build or test the sdk
```
go get github.com/stretchr/testify
```

## Configuring the SDK
You can configure the sdk with your credentials by creating a settings file in:
 $HOME/.oci/config
 ```
 [DEFAULT]
 user=[user ocid]
 fingerprint=[fingerprint]
 key_file=[path to pem file containing private key]
 tenancy=[tenancy ocid]
 region=[region for your tenancy]
 ```

## Making a request
Here is a quick example showing how to make a request to the identity service
```
	c := identity.NewIdentityClientForRegion(getRegion())
	request := identity.ListAvailabilityDomainsRequest{}
	response, err := c.ListAvailabilityDomains(context.Background(), request)
	fmt.Println(response.Items)
```

## Organization of the SDK

## Building and testing
Building is provided by the make file at the root of the project. To build the project execute

```
make build
```

To execute the tests:
```
make test
```

