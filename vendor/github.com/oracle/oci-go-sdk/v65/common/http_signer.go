// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

package common

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// HTTPRequestSigner the interface to sign a request
type HTTPRequestSigner interface {
	Sign(r *http.Request) error
}

// KeyProvider interface that wraps information about the key's account owner
type KeyProvider interface {
	PrivateRSAKey() (*rsa.PrivateKey, error)
	KeyID() (string, error)
}

const signerVersion = "1"

// SignerBodyHashPredicate a function that allows to disable/enable body hashing
// of requests and headers associated with body content
type SignerBodyHashPredicate func(r *http.Request) bool

// ociRequestSigner implements the http-signatures-draft spec
// as described in https://tools.ietf.org/html/draft-cavage-http-signatures-08
type ociRequestSigner struct {
	KeyProvider    KeyProvider
	GenericHeaders []string
	BodyHeaders    []string
	ShouldHashBody SignerBodyHashPredicate
	SigningMethod  SigningMethod
}

// SigningMethod wraps the information about signature scheme and options
type SigningMethod struct {
	SignatureScheme RSASignatureScheme
	RSAPSSOptions   rsa.PSSOptions
}

// RSASignatureScheme type of the RSA Signature Scheme
type RSASignatureScheme string

// Set of constants representing the allowable values for RSASignatureScheme
const (
	PKCS1v15 RSASignatureScheme = "PKCS1v15"
	PSS      RSASignatureScheme = "PSS"
)

var (
	defaultGenericHeaders    = []string{"date", "(request-target)", "host"}
	defaultBodyHeaders       = []string{"content-length", "content-type", "x-content-sha256"}
	defaultBodyHashPredicate = func(r *http.Request) bool {
		return r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch
	}
	defaultSigningMethod = SigningMethod{
		SignatureScheme: RSASignatureScheme(PKCS1v15)}
)

// DefaultGenericHeaders list of default generic headers that is used in signing
func DefaultGenericHeaders() []string {
	return makeACopy(defaultGenericHeaders)
}

// DefaultBodyHeaders list of default body headers that is used in signing
func DefaultBodyHeaders() []string {
	return makeACopy(defaultBodyHeaders)
}

// DefaultRequestSigner creates a signer with default parameters.
func DefaultRequestSigner(provider KeyProvider) HTTPRequestSigner {
	return RequestSigner(provider, defaultGenericHeaders, defaultBodyHeaders)
}

// RequestSignerExcludeBody creates a signer without hash the body.
func RequestSignerExcludeBody(provider KeyProvider) HTTPRequestSigner {
	bodyHashPredicate := func(r *http.Request) bool {
		// week request signer will not hash the body
		return false
	}
	return RequestSignerWithBodyHashingPredicate(provider, defaultGenericHeaders, defaultBodyHeaders, bodyHashPredicate)
}

// NewSignerFromOCIRequestSigner creates a copy of the request signer and attaches the new SignerBodyHashPredicate
// returns an error if the passed signer is not of type ociRequestSigner
func NewSignerFromOCIRequestSigner(oldSigner HTTPRequestSigner, predicate SignerBodyHashPredicate) (HTTPRequestSigner, error) {
	if oldS, ok := oldSigner.(ociRequestSigner); ok {
		s := ociRequestSigner{
			KeyProvider:    oldS.KeyProvider,
			GenericHeaders: oldS.GenericHeaders,
			BodyHeaders:    oldS.BodyHeaders,
			ShouldHashBody: predicate,
			SigningMethod:  oldS.SigningMethod,
		}
		return s, nil

	}
	return nil, fmt.Errorf("can not create a signer, input signer needs to be of type ociRequestSigner")
}

// RequestSigner creates a signer that utilizes the specified headers for signing
// and the default predicate for using the body of the request as part of the signature
func RequestSigner(provider KeyProvider, genericHeaders, bodyHeaders []string) HTTPRequestSigner {
	return ociRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: genericHeaders,
		BodyHeaders:    bodyHeaders,
		ShouldHashBody: defaultBodyHashPredicate,
		SigningMethod:  defaultSigningMethod}
}

// RequestSignerWithBodyHashingPredicate creates a signer that utilizes the specified headers for signing, as well as a predicate for using
// the body of the request and bodyHeaders parameter as part of the signature
func RequestSignerWithBodyHashingPredicate(provider KeyProvider, genericHeaders, bodyHeaders []string, shouldHashBody SignerBodyHashPredicate) HTTPRequestSigner {
	return ociRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: genericHeaders,
		BodyHeaders:    bodyHeaders,
		ShouldHashBody: shouldHashBody,
		SigningMethod:  defaultSigningMethod}
}

// RequestSignerWithSigningMethodAndBodyHashingPredicate creates a signer with the specified signature scheme and
// utilizes the specified headers for signing, as well as a predicate for using the body of the request and
// bodyHeaders parameter as part of the signature
func RequestSignerWithSigningMethodAndBodyHashingPredicate(provider KeyProvider, genericHeaders, bodyHeaders []string, shouldHashBody SignerBodyHashPredicate, signingMethod SigningMethod) HTTPRequestSigner {
	return ociRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: genericHeaders,
		BodyHeaders:    bodyHeaders,
		ShouldHashBody: shouldHashBody,
		SigningMethod:  signingMethod}
}

// RequestSignerWithSigningMethod creates a signer with the specified signature scheme
func RequestSignerWithSigningMethod(provider KeyProvider, genericHeaders, bodyHeaders []string, signingMethod SigningMethod) HTTPRequestSigner {
	return ociRequestSigner{
		KeyProvider:    provider,
		GenericHeaders: genericHeaders,
		BodyHeaders:    bodyHeaders,
		ShouldHashBody: defaultBodyHashPredicate,
		SigningMethod:  signingMethod}
}

func (signer ociRequestSigner) getSigningHeaders(r *http.Request) []string {
	var result []string
	result = append(result, signer.GenericHeaders...)

	if signer.ShouldHashBody(r) {
		result = append(result, signer.BodyHeaders...)
	}

	return result
}

func (signer ociRequestSigner) getSigningStringAndHeaders(request *http.Request) (string, []string) {
	headersToSign := signer.getSigningHeaders(request)
	signedHeaderNames := make([]string, len(headersToSign))
	signedHeaders := make([]string, len(headersToSign))
	signedHeaderCount := 0
	for _, headerName := range headersToSign {
		headerName = strings.ToLower(headerName)
		var value string
		switch headerName {
		case "(request-target)":
			value = getRequestTarget(request)
		case "host":
			value = request.URL.Host
			if len(value) == 0 {
				value = request.Host
			}
		default:
			value = request.Header.Get(headerName)
		}
		if value != "" {
			signedHeaders[signedHeaderCount] = fmt.Sprintf("%s: %s", headerName, value)
			signedHeaderNames[signedHeaderCount] = headerName
			signedHeaderCount++
		}
	}

	signingString := strings.Join(signedHeaders[0:signedHeaderCount], "\n")
	return signingString, signedHeaderNames[0:signedHeaderCount]
}

func getRequestTarget(request *http.Request) string {
	lowercaseMethod := strings.ToLower(request.Method)
	return fmt.Sprintf("%s %s", lowercaseMethod, request.URL.RequestURI())
}

func calculateHashOfBody(request *http.Request) (err error) {
	var hash string
	hash, err = GetBodyHash(request)
	if err != nil {
		return
	}
	request.Header.Set(requestHeaderXContentSHA256, hash)
	return
}

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
//
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func hashAndEncode(data []byte) string {
	hashedContent := sha256.Sum256(data)
	hash := base64.StdEncoding.EncodeToString(hashedContent[:])
	return hash
}

// GetBodyHash creates a base64 string from the hash of body the request
func GetBodyHash(request *http.Request) (hashString string, err error) {
	if request.Body == nil {
		request.ContentLength = 0
		request.Header.Set(requestHeaderContentLength, fmt.Sprintf("%v", request.ContentLength))
		return hashAndEncode([]byte("")), nil
	}

	var data []byte
	var bReader io.Reader
	bReader, request.Body, err = drainBody(request.Body)
	if err != nil {
		return "", fmt.Errorf("can not read body of request while calculating body hash: %s", err.Error())
	}

	data, err = ioutil.ReadAll(bReader)
	if err != nil {
		return "", fmt.Errorf("can not read body of request while calculating body hash: %s", err.Error())
	}

	// Since the request can be coming from a binary body. Make an attempt to set the body length
	request.ContentLength = int64(len(data))
	request.Header.Set(requestHeaderContentLength, fmt.Sprintf("%v", request.ContentLength))

	hashString = hashAndEncode(data)

	return
}

func (signer ociRequestSigner) computeSignature(request *http.Request) (signature string, signingHeaders []string, err error) {
	signingString, signingHeaders := signer.getSigningStringAndHeaders(request)
	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	hashed := hasher.Sum(nil)

	privateKey, err := signer.KeyProvider.PrivateRSAKey()
	if err != nil {
		return
	}

	var unencodedSig []byte
	var e error
	switch signatureScheme := signer.SigningMethod.SignatureScheme; signatureScheme {
	case PSS:
		unencodedSig, e = rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, hashed, &signer.SigningMethod.RSAPSSOptions)
	default:
		unencodedSig, e = rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	}

	if e != nil {
		err = fmt.Errorf("can not compute signature while signing the request %s: ", e.Error())
		return
	}

	signature = base64.StdEncoding.EncodeToString(unencodedSig)
	return
}

// Sign signs the http request, by inspecting the necessary headers. Once signed
// the request will have the proper 'Authorization' header set, otherwise
// and error is returned
func (signer ociRequestSigner) Sign(request *http.Request) (err error) {
	if signer.ShouldHashBody(request) {
		err = calculateHashOfBody(request)
		if err != nil {
			return
		}
	}

	var signature string
	var signingHeaders []string
	if signature, signingHeaders, err = signer.computeSignature(request); err != nil {
		return
	}

	var keyID string
	if keyID, err = signer.KeyProvider.KeyID(); err != nil {
		return
	}

	authValue := fmt.Sprintf("Signature version=\"%s\",headers=\"%s\",keyId=\"%s\",algorithm=\"rsa-sha256\",signature=\"%s\"",
		signerVersion, strings.Join(signingHeaders, " "), keyID, signature)

	request.Header.Set(requestHeaderAuthorization, authValue)

	return
}
