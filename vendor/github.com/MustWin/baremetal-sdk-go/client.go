// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

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
	"fmt"
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

type NewClientOptions struct {
	Transport   http.RoundTripper
	UserAgent   string
	keyPassword *string
	keyPath     *string
	keyBytes    []byte
}

type NewClientOptionsFunc func(o *NewClientOptions)

// PrivateKeyPassword password to decrypt an encrypted private key
func PrivateKeyPassword(pwd string) func(o *NewClientOptions) {
	return func(o *NewClientOptions) {
		o.keyPassword = &pwd
	}
}

// PrivateKeyFilePath provide the path a file containing a private key.
func PrivateKeyFilePath(path string) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.keyPath = &path
	}
}

// PrivateKeyBytes supply bytes for private key
func PrivateKeyBytes(buff []byte) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.keyBytes = buff
	}
}

// CustomTransport can be used to assign a custom http.RoundTripper
// to the Bare Metal API connection. For example, you could wrap the default
// http transport in your won transport the logs interactions for diagnotic
// purposes.
func CustomTransport(tr http.RoundTripper) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.Transport = tr
	}
}

// UserAgent assigns a custom user agent for API connection
func UserAgent(userAgent string) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.UserAgent = userAgent
	}
}

// NewClient creates and authenticates a BareMetal API client
func NewClient(userOCID, tenancyOCID, keyFingerprint string, opts ...NewClientOptionsFunc) (*Client, error) {
	var err error
	auth := &authenticationInfo{
		tenancyOCID:    tenancyOCID,
		userOCID:       userOCID,
		keyFingerPrint: keyFingerprint,
	}
	nco := &NewClientOptions{
		Transport: &http.Transport{},
	}
	for _, opt := range opts {
		opt(nco)
	}

	if nco.keyPath != nil {
		auth.privateRSAKey, err = PrivateKeyFromFile(*nco.keyPath, nco.keyPassword)
	} else {
		auth.privateRSAKey, err = PrivateKeyFromBytes(nco.keyBytes, nco.keyPassword)
	}
	if err != nil {
		return nil, err
	}
	return &Client{
		userAgent:        nco.UserAgent,
		authInfo:         auth,
		identityApi:      newIdentityAPIRequestor(auth, nco),
		coreApi:          newCoreAPIRequestor(auth, nco),
		objectStorageApi: newObjectStorageAPIRequestor(auth, nco),
		databaseApi:      newDatabaseAPIRequestor(auth, nco),
	}, nil
}

// New creates a new client to access Oracle BareMetal services.
// userOCI, tenancyOCID and fingerprint arguments are accessed from the BareMetal identity
// console. privateKey is an RSA key associated with the user accessing the API.
func New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey, opts ...func(o *NewClientOptions)) (c *Client) {
	auth := &authenticationInfo{
		privateRSAKey:  privateKey,
		tenancyOCID:    tenancyOCID,
		userOCID:       userOCID,
		keyFingerPrint: keyFingerPrint,
	}

	nco := &NewClientOptions{
		Transport: &http.Transport{},
	}

	for _, opt := range opts {
		opt(nco)
	}

	return &Client{
		userAgent:        nco.UserAgent,
		authInfo:         auth,
		identityApi:      newIdentityAPIRequestor(auth, nco),
		coreApi:          newCoreAPIRequestor(auth, nco),
		objectStorageApi: newObjectStorageAPIRequestor(auth, nco),
		databaseApi:      newDatabaseAPIRequestor(auth, nco),
	}
}

// PrivateKeyFromBytes is a helper function that will produce a RSA private
// key from bytes.
func PrivateKeyFromBytes(pemData []byte, password *string) (key *rsa.PrivateKey, e error) {
	if pemBlock, _ := pem.Decode(pemData); pemBlock != nil {

		decrypted := pemBlock.Bytes

		if x509.IsEncryptedPEMBlock(pemBlock) {
			if password == nil {
				e = fmt.Errorf("private_key_password is required for encrypted private keys")
				return
			}
			if decrypted, e = x509.DecryptPEMBlock(pemBlock, []byte(*password)); e != nil {
				return
			}
		}

		key, e = x509.ParsePKCS1PrivateKey(decrypted)

	} else {
		e = errors.New("PEM data was not found in buffer")
		return
	}

	return
}

func PrivateKeyFromFile(pemFilePath string, password *string) (key *rsa.PrivateKey, e error) {
	var fileData []byte
	if fileData, e = ioutil.ReadFile(pemFilePath); e != nil {
		return
	}

	key, e = PrivateKeyFromBytes(fileData, password)

	return

}
func PrivateKeyFromUnencryptedBytes(pemBytes []byte) (*rsa.PrivateKey, error) {
	pemBlock, _ := pem.Decode(pemBytes)
	return x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
}

func PrivateKeyFromUnencryptedFile(path string) (*rsa.PrivateKey, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return PrivateKeyFromUnencryptedBytes(buff)
}
