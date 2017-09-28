// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

/*
  Package baremetal provides access to the Oracle Bare Metal Cloud API's. See Readme for usage example.
*/

package baremetal

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Client is used to access Oracle BareMetal Services
type Client struct {
	userAgent        string
	authInfo         *authenticationInfo
	identityApi      requestor
	coreApi          requestor
	databaseApi      requestor
	objectStorageApi requestor
	loadBalancerApi  requestor
	identityEndPoint string
}

type NewClientOptions struct {
	Region                 string
	Transport              http.RoundTripper
	UrlTemplate            string
	UserAgent              string
	keyPassword            *string
	keyPath                *string
	keyBytes               []byte
	ShortRetryTime         time.Duration
	LongRetryTime          time.Duration
	RandGen                *rand.Rand
	DisableAutoRetries     bool
	DisableNotFoundRetries bool
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
// http transport in your own transport and log interactions for diagnostic
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

// Region assigns a region override for API connections
func Region(region string) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.Region = region
	}
}

// UrlTemplate lets you override the production url scheme of https://%s.%s.oraclecloud.com
func UrlTemplate(urlTemplate string) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.UrlTemplate = urlTemplate
	}
}

func ShortRetryTime(retryTime time.Duration) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.ShortRetryTime = retryTime
	}
}

func LongRetryTime(retryTime time.Duration) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.LongRetryTime = retryTime
	}
}

func DisableAutoRetries(disableAutoRetries bool) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.DisableAutoRetries = disableAutoRetries
	}
}

func DisableNotFoundRetries(disableNotFoundRetries bool) NewClientOptionsFunc {
	return func(o *NewClientOptions) {
		o.DisableNotFoundRetries = disableNotFoundRetries
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

	//create random number generator for creating Retry Tokens
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	nco := &NewClientOptions{
		Transport:      &http.Transport{},
		Region:         us_phoenix_1,
		UrlTemplate:    baseUrlTemplate,
		ShortRetryTime: shortRetryTime,
		LongRetryTime:  longRetryTime,
		RandGen:        randGen,
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
		loadBalancerApi:  newLoadBalancerAPIRequestor(auth, nco),
	}, nil
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
