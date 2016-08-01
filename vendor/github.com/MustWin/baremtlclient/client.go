package baremtlsdk

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
	authInfo         *authenticationInfo
	identityAPI      requestor
	identityEndPoint string
}

// New creates a new client to access Oracle BareMetal services
// OCID and fingerprint arguments are accessed from BareMetal identity
// console, private key is RSA key associated with user accessing api, and must
// be associated with user.
func New(userOCID, tenancyOCID, keyFingerPrint string, privateKey *rsa.PrivateKey) (c *Client) {
	auth := &authenticationInfo{
		privateRSAKey:  privateKey,
		tenancyOCID:    tenancyOCID,
		userOCID:       userOCID,
		keyFingerPrint: keyFingerPrint,
	}

	// TODO: set configuration for real https client
	tr := &http.Transport{}

	return &Client{
		authInfo:    auth,
		identityAPI: newAPIRequestor(auth, tr),
	}
}

// NewFromKeyPath creates a client reading an RSA private key from a file. The
// userOCID and tenancyOCID are obtained from the BareMetal console.
// The fingerprint can be obtained from the BareMetal console or
// openssl rsa -pubout -outform DER -in private.pem | openssl md5 -c
func NewFromKeyPath(userOCID, tenancyOCID, keyFingerPrint, privateKeyPath, keyPassword string) (c *Client, e error) {
	var key *rsa.PrivateKey

	if key, e = PrivateKeyFromFile(privateKeyPath, keyPassword); e != nil {
		return
	}

	c = New(userOCID, tenancyOCID, keyFingerPrint, key)

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
