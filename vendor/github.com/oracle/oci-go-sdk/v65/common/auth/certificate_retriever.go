// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package auth

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/fs"
	"os"
	"sync"
	"time"

	"github.com/oracle/oci-go-sdk/v65/common"
)

// Returns a copy of the input bytes data
func copyOfBytes(in []byte) []byte {
	if in == nil {
		return nil
	} // avoid allocations
	c := make([]byte, len(in))
	copy(c, in)
	return c
}

// x509CertificateRetriever provides an X509 certificate with the RSA private key
type x509CertificateRetriever interface {
	Refresh() error
	CertificatePemRaw() []byte
	Certificate() *x509.Certificate
	PrivateKeyPemRaw() []byte
	PrivateKey() *rsa.PrivateKey
}

// urlBasedX509CertificateRetriever retrieves PEM-encoded X509 certificates from the given URLs.
type urlBasedX509CertificateRetriever struct {
	certURL           string
	privateKeyURL     string
	passphrase        string
	certificatePemRaw []byte
	certificate       *x509.Certificate
	privateKeyPemRaw  []byte
	privateKey        *rsa.PrivateKey
	mux               sync.Mutex
	dispatcher        common.HTTPRequestDispatcher
}

func newURLBasedX509CertificateRetriever(dispatcher common.HTTPRequestDispatcher, certURL, privateKeyURL, passphrase string) x509CertificateRetriever {
	return &urlBasedX509CertificateRetriever{
		certURL:       certURL,
		privateKeyURL: privateKeyURL,
		passphrase:    passphrase,
		mux:           sync.Mutex{},
		dispatcher:    dispatcher,
	}
}

// Refresh() is failure atomic, i.e., CertificatePemRaw(), Certificate(), PrivateKeyPemRaw(), and PrivateKey() would
// return their previous values if Refresh() fails.
func (r *urlBasedX509CertificateRetriever) Refresh() error {
	common.Debugln("Refreshing certificate")

	r.mux.Lock()
	defer r.mux.Unlock()

	var err error

	var certificatePemRaw []byte
	var certificate *x509.Certificate
	if certificatePemRaw, certificate, err = r.renewCertificate(r.certURL); err != nil {
		return fmt.Errorf("failed to renew certificate: %w", err)
	}

	var privateKeyPemRaw []byte
	var privateKey *rsa.PrivateKey
	if r.privateKeyURL != "" {
		if privateKeyPemRaw, privateKey, err = r.renewPrivateKey(r.privateKeyURL, r.passphrase); err != nil {
			return fmt.Errorf("failed to renew private key: %w", err)
		}
	}

	r.certificatePemRaw = certificatePemRaw
	r.certificate = certificate
	r.privateKeyPemRaw = privateKeyPemRaw
	r.privateKey = privateKey
	return nil
}

func (r *urlBasedX509CertificateRetriever) renewCertificate(url string) (certificatePemRaw []byte, certificate *x509.Certificate, err error) {
	var body bytes.Buffer
	if body, _, err = httpGet(r.dispatcher, url); err != nil {
		return nil, nil, fmt.Errorf("failed to get certificate from %s: %w", url, err)
	}

	certificatePemRaw = body.Bytes()
	var block *pem.Block
	block, _ = pem.Decode(certificatePemRaw)
	if block == nil {
		return nil, nil, fmt.Errorf("failed to parse the new certificate, not valid pem data")
	}

	if certificate, err = x509.ParseCertificate(block.Bytes); err != nil {
		return nil, nil, fmt.Errorf("failed to parse the new certificate: %w", err)
	}

	return certificatePemRaw, certificate, nil
}

func (r *urlBasedX509CertificateRetriever) renewPrivateKey(url, passphrase string) (privateKeyPemRaw []byte, privateKey *rsa.PrivateKey, err error) {
	var body bytes.Buffer
	if body, _, err = httpGet(r.dispatcher, url); err != nil {
		return nil, nil, fmt.Errorf("failed to get private key from %s: %w", url, err)
	}

	privateKeyPemRaw = body.Bytes()
	if privateKey, err = common.PrivateKeyFromBytes(privateKeyPemRaw, &passphrase); err != nil {
		return nil, nil, fmt.Errorf("failed to parse the new private key: %w", err)
	}

	return privateKeyPemRaw, privateKey, nil
}

func (r *urlBasedX509CertificateRetriever) CertificatePemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.certificatePemRaw)
}

func (r *urlBasedX509CertificateRetriever) Certificate() *x509.Certificate {
	r.mux.Lock()
	defer r.mux.Unlock()

	if r.certificate == nil {
		return nil
	}

	c := *r.certificate
	return &c
}

func (r *urlBasedX509CertificateRetriever) PrivateKeyPemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.privateKeyPemRaw)
}

func (r *urlBasedX509CertificateRetriever) PrivateKey() *rsa.PrivateKey {
	r.mux.Lock()
	defer r.mux.Unlock()

	//Nil Private keys are supported as part of a certificate
	if r.privateKey == nil {
		return nil
	}

	c := *r.privateKey
	return &c
}

// newStaticX509CertificateRetriever creates a static memory based retriever.
func newStaticX509CertificateRetriever(certificatePemRaw, privateKeyPemRaw []byte, passphrase []byte) x509CertificateRetriever {
	return &staticCertificateRetriever{
		CertificatePem: certificatePemRaw,
		PrivateKeyPem:  privateKeyPemRaw,
		Passphrase:     passphrase,
	}
}

// staticCertificateRetriever serves certificates from static data
type staticCertificateRetriever struct {
	Passphrase     []byte
	CertificatePem []byte
	PrivateKeyPem  []byte
	certificate    *x509.Certificate
	privateKey     *rsa.PrivateKey
	mux            sync.Mutex
}

// Refresh proccess the inputs into appropiate keys and certificates
func (r *staticCertificateRetriever) Refresh() error {
	r.mux.Lock()
	defer r.mux.Unlock()

	certifcate, err := r.readCertificate()
	if err != nil {
		r.certificate = nil
		return err
	}
	r.certificate = certifcate

	key, err := r.readPrivateKey()
	if err != nil {
		r.privateKey = nil
		return err
	}
	r.privateKey = key

	return nil
}

func (r *staticCertificateRetriever) Certificate() *x509.Certificate {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.certificate
}

func (r *staticCertificateRetriever) PrivateKey() *rsa.PrivateKey {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.privateKey
}

func (r *staticCertificateRetriever) CertificatePemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.CertificatePem)
}

func (r *staticCertificateRetriever) PrivateKeyPemRaw() []byte {
	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.PrivateKeyPem)
}

func (r *staticCertificateRetriever) readCertificate() (certificate *x509.Certificate, err error) {
	block, _ := pem.Decode(r.CertificatePem)
	if block == nil {
		return nil, fmt.Errorf("failed to parse the new certificate, not valid pem data")
	}

	if certificate, err = x509.ParseCertificate(block.Bytes); err != nil {
		return nil, fmt.Errorf("failed to parse the new certificate: %w", err)
	}
	return certificate, nil
}

func (r *staticCertificateRetriever) readPrivateKey() (*rsa.PrivateKey, error) {
	if r.PrivateKeyPem == nil {
		return nil, nil
	}

	var pass *string
	if r.Passphrase == nil {
		pass = nil
	} else {
		ss := string(r.Passphrase)
		pass = &ss
	}
	return common.PrivateKeyFromBytes(r.PrivateKeyPem, pass)
}

type fileBasedCertificateRetriever struct {
	Passphrase         []byte
	CertificatePemPath string
	PrivateKeyPemPath  string
	RefreshRate        time.Duration

	mux                 sync.Mutex
	certificatePem      []byte
	privateKeyPem       []byte
	certificate         *x509.Certificate
	privateKey          *rsa.PrivateKey
	lastRefreshedAt     time.Time
	certificateFileStat os.FileInfo
	privateKeyFileStat  os.FileInfo
}

// Check if refresh is needed
func (r *fileBasedCertificateRetriever) refreshIfNeeded() {
	if r.shouldRefresh() {
		if r.certFileChanged() || r.privateKeyFileChanged() {
			// Refresh only if files changed
			r.Refresh()
		} else {
			// Update the last refresh time interval to skip this refresh till next interval
			r.mux.Lock()
			defer r.mux.Unlock()
			r.lastRefreshedAt = time.Now()
		}
	}
}

// Check if the time since the last refresh was greater than the refresh rate
func (r *fileBasedCertificateRetriever) shouldRefresh() bool {
	return r.RefreshRate > 0 && time.Since(r.lastRefreshedAt) > r.RefreshRate
}

// Checks if the cert file has changed
func (r *fileBasedCertificateRetriever) certFileChanged() bool {
	currentCertificateStat, err := os.Stat(r.CertificatePemPath)
	if err != nil {
		return false
	}
	if r.certificateFileStat == nil {
		return false
	}
	if r.certificateFileStat.Size() != currentCertificateStat.Size() {
		return true
	}
	if r.certificateFileStat.ModTime() != currentCertificateStat.ModTime() {
		return true
	}
	return false
}

// Checks if the private key file has changed
func (r *fileBasedCertificateRetriever) privateKeyFileChanged() bool {
	if r.PrivateKeyPemPath == "" {
		return false
	}
	currentPrivateKeyStat, err := os.Stat(r.PrivateKeyPemPath)
	if err != nil {
		return false
	}
	if r.privateKeyFileStat == nil {
		return false
	}
	if r.privateKeyFileStat.Size() != currentPrivateKeyStat.Size() {
		return true
	}
	if r.privateKeyFileStat.ModTime() != currentPrivateKeyStat.ModTime() {
		return true
	}
	return false
}

// Refresh process the inputs into appropriate keys and certificates
func (r *fileBasedCertificateRetriever) Refresh() error {
	certificate, certPem, certStat, certificateError := r.readCertificate()
	if certificateError != nil {
		return certificateError
	}
	key, keyPem, keyStat, keyError := r.readPrivateKey()
	if keyError != nil {
		return keyError
	}
	r.mux.Lock()
	defer r.mux.Unlock()
	r.certificate = certificate
	r.certificatePem = certPem
	r.certificateFileStat = certStat
	r.privateKey = key
	r.privateKeyPem = keyPem
	r.privateKeyFileStat = keyStat
	r.lastRefreshedAt = time.Now()
	return nil
}

func (r *fileBasedCertificateRetriever) readCertificate() (*x509.Certificate, []byte, fs.FileInfo, error) {
	leafCert, err := os.ReadFile(r.CertificatePemPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error reading leafCertificate at path:%s due to error:%w", r.CertificatePemPath, err)
	}
	block, _ := pem.Decode(leafCert)
	if block == nil {
		return nil, nil, nil, fmt.Errorf("failed to parse the new certificate, not valid pem data")
	}
	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to parse the new certificate: %w", err)
	}
	certificateStat, err := os.Stat(r.CertificatePemPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to stat certificate file at path:%s due to:%w", r.CertificatePemPath, err)
	}
	return certificate, leafCert, certificateStat, nil
}

func (r *fileBasedCertificateRetriever) readPrivateKey() (*rsa.PrivateKey, []byte, fs.FileInfo, error) {
	if r.PrivateKeyPemPath == "" {
		return nil, nil, nil, nil
	}
	var pass *string
	if r.Passphrase == nil {
		pass = nil
	} else {
		ss := string(r.Passphrase)
		pass = &ss
	}
	leafKey, err := os.ReadFile(r.PrivateKeyPemPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("reading leafPrivateKey at path: %s due to: %w", r.PrivateKeyPemPath, err)
	}
	privateKey, err := common.PrivateKeyFromBytes(leafKey, pass)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("getting private key due to :%w", err)
	}
	privateKeyFileStat, err := os.Stat(r.PrivateKeyPemPath)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to stat PrivateKey file at path:%s due to :%w", r.PrivateKeyPemPath, err)
	}
	return privateKey, leafKey, privateKeyFileStat, nil
}

func (r *fileBasedCertificateRetriever) Certificate() *x509.Certificate {
	// Check for refresh
	r.refreshIfNeeded()

	r.mux.Lock()
	defer r.mux.Unlock()
	return r.certificate
}

func (r *fileBasedCertificateRetriever) PrivateKey() *rsa.PrivateKey {
	// Check for refresh
	r.refreshIfNeeded()

	r.mux.Lock()
	defer r.mux.Unlock()
	return r.privateKey
}

func (r *fileBasedCertificateRetriever) CertificatePemRaw() []byte {
	// Check for refresh
	r.refreshIfNeeded()

	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.certificatePem)
}

func (r *fileBasedCertificateRetriever) PrivateKeyPemRaw() []byte {
	// Check for refresh
	r.refreshIfNeeded()

	r.mux.Lock()
	defer r.mux.Unlock()
	return copyOfBytes(r.privateKeyPem)
}
