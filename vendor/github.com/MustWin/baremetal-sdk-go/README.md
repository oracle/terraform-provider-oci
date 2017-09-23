# Go BareMetal SDK ![](https://circleci.com/gh/MustWin/baremetal-sdk-go.svg?style=shield&circle-token=fa06ce2af6b594812e3a756f5451a9e101d7b9f5)

Package baremetal provides access to the Oracle BareMetal Cloud APIs.

## Usage

To use the Go BareMetal SDK instantiate a baremetal.Client, supplying
your tenancy OCID, user OCID, RSA public key fingerprint, and RSA private key.
Then call functions as the example below illustrates.

```
import (
  "fmt"
  "github.com/oracle/bmcs-go-sdk"
)

func main() {
  client, err := baremetal.NewClient(
    <user ocid>,
    <tenancy ocid>,
    <public key fingerprint>,
    baremetal.PrivateKeyFilePath(<path to private key>),
  )

  if err != nil {
    fmt.Println("Error setting up bmc client\n", err)
  }

  compartmentList, err := Client.ListCompartments(nil)

  if err != nil {
    fmt.Println("Error listing Compartments\n", err)
  }

  for _, c := range compartments.Compartments {
  	fmt.Println(c.Name)
  }
}
```

For more details, see the [API Docs](https://docs.us-phoenix-1.oraclecloud.com/)

## Unit Testing
Some of the tests rely on GOPATH to build a path where a test private key is located. If
for some reason you have a composite GOPATH i.e /home/foo/go-projects:/usr/stuff
these tests will break.  In that case you export an environment variable with an
explicit path to the test private key.

```
export BAREMETAL_SDK_PEM_DATA_PATH="/home/foo/go-projects/src/github.com/../test/data/private.pem"
```

## Acceptance Tests

Use make to run acceptance tests
```
make acceptance_test
```

In this mode acceptance tests run using fixtures contained in the acceptance-test/fixtures
directory as opposed to calling out to the Oracle Bare Metal API and as such
will not require any special authorization information.

Running the tests in recording mode will make live API calls and record new results in the fixtures directory. This is done by running the `record_acceptance_test` task:
```
make record_acceptance_test
```

To record individual tests, you need to specify the `recording` tag when running your test. For example:
```
go test -v -timeout 120m -tags "recording all" github.com/oracle/bmcs-go-sdk/acceptance-test -run  TestInstanceCRUD
```

You will need to provide credentials to access the Bare Metal API in an .env file
in the acceptance-test directory.  A sample .env file can be found at acceptance-test/sample.env.
Simply copy the sample.env to .env and supply your own credentials. Note that if you
create new tests you'll have to run them in recording mode before they will pass. 

# Vendoring
This project uses the [Go vendor folder](https://blog.gopheracademy.com/advent-2015/vendor-folder/) for dependencies.
If you need to add or update dependency, please review the [go
vendor docs](https://github.com/kardianos/govendor).

# References
[Oracle Bare Metal Iaas API Docs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apiref.htm)
[Oracle Bare Metal Iaas Docs](https://docs.us-phoenix-1.oraclecloud.com/)
