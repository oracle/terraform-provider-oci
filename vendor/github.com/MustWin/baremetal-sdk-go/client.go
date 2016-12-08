/*
	Package baremetal provides access to the Oracle Bare Metal Cloud API's

	Usage:

	To use the Go BareMetal SDK instantiate a baremetal.Client, supplying
	your tenancy OCID, user OCID, RSA public key fingerprint, and RSA private key.
	Then call functions as the example below illustrates.  Note that error
	handling has been omitted to add clarity.

		import (
		  "fmt"
		  "crypto/rsa"
		  "github.com/MustWin/baremetal-sdk-go"
		)

		func main() {
		  privateKey, _ := baremetal.PrivateKeyFromFile("/path/to/key.pem", "keyPassword")

		  client := baremetal.New(
		    "ocid1.tenancy.oc1..aaaaaaaaq3hulfjvrouw3e6qx2ncxtp256aq7etiabqqtzunnhxjslzkfyxq",
		    "ocid1.user.oc1..aaaaaaaaflxvsdpjs5ztahmsf7vjxy5kdqnuzyqpvwnncbkfhavexwd4w5ra",
		    "b4:8a:7d:54:e6:81:04:b2:99:8e:b3:ed:10:e2:12:2b",
		    privateKey,
		  )

		  availabilityDomains, _ := client.ListAvailablityDomains()

		  for _, ad := range availabilityDomains {
		    fmt.Println(ad.Name)
		  }

		}

	For more details, see the API docs located here:
	https://docs.us-az-phoenix-1.oracleiaas.com/#API/Concepts/usingapi.htm
*/
package baremetal

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"net/http"
)

// Client is used to access Oracle BareMetal Services
type Client struct {
	userAgent        string
	authInfo         *authenticationInfo
	identityApi      requestor
	coreApi          requestor
	databaseApi      requestor
	objectStorageApi requestor
	identityEndPoint string
}

// New creates a new client to access Oracle BareMetal services.
// userOCI, tenancyOCID and fingerprint arguments are accessed from the BareMetal identity
// console. privateKey is an RSA key associated with the user accessing the API.
func New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey, userAgent string) (c *Client) {
	auth := &authenticationInfo{
		privateRSAKey:  privateKey,
		tenancyOCID:    tenancyOCID,
		userOCID:       userOCID,
		keyFingerPrint: keyFingerPrint,
	}

	// TODO: set configuration for real https client
	tr := &http.Transport{}

	return &Client{
		userAgent:        userAgent,
		authInfo:         auth,
		identityApi:      newIdentityAPIRequestor(auth, tr),
		coreApi:          newCoreAPIRequestor(auth, tr),
		objectStorageApi: newObjectStorageAPIRequestor(auth, tr),
		databaseApi:      newDatabaseAPIRequestor(auth, tr),
	}
}

// NewFromKeyPath creates a client reading an RSA private key from a file. The
// userOCID and tenancyOCID are obtained from the BareMetal console.
// The fingerprint can be obtained from the BareMetal console or running
//     openssl rsa -pubout -outform DER -in private.pem | openssl md5 -c
func NewFromKeyPath(userOCID, tenancyOCID, keyFingerPrint, privateKeyPath, keyPassword string) (c *Client, e error) {
	var key *rsa.PrivateKey

	if key, e = PrivateKeyFromFile(privateKeyPath, keyPassword); e != nil {
		return
	}

	c = New(userOCID, tenancyOCID, keyFingerPrint, key, "")

	return
}

// PrivateKeyFromBytes is a helper function that will produce a RSA private
// key from bytes.
func PrivateKeyFromBytes(pemData []byte, password string) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {

		var decrypted []byte

		if decrypted, e = x509.DecryptPEMBlock(pemBlock, []byte(password)); e != nil {
			return
		}

		key, e = x509.ParsePKCS1PrivateKey(decrypted)

	} else {
		e = errors.New("PEM data was not found in buffer")
		return
	}

	return
}

// PrivateKeyFromFile is a helper function that will produce an RSA private
// key from a PEM file.  The PEM file MUST be created with a password which
// is supplied as an argument.
func PrivateKeyFromFile(pemFilePath, password string) (key *rsa.PrivateKey, e error) {
	var fileData []byte
	if fileData, e = ioutil.ReadFile(pemFilePath); e != nil {
		return
	}

	key, e = PrivateKeyFromBytes(fileData, password)

	return

}
