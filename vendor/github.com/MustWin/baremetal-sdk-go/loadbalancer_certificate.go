// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// Certificate defines a listener certificate bundle.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Certificate/
// Also https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/requests/CertificateDetails
type Certificate struct {
	OPCRequestIDUnmarshaller
	OPCWorkRequestIDUnmarshaller
	CertificateName   string `header:"-" url:"-" json:"certificateName"`
	PublicCertificate string `header:"-" url:"-" json:"publicCertificate"`
	PrivateKey        string `header:"-" url:"-" json:"privateKey,omitempty"` // Only for create
	// Optional
	CACertificate string `header:"-" url:"-" json:"caCertificate,omitempty"`
	Passphrase    string `header:"-" url:"-" json:"passphrase,omitempty"` // Only for create
}

// ListCertificates contains a list of certificates
//
type ListCertificates struct {
	OPCRequestIDUnmarshaller
	Certificates []Certificate
}

func (l *ListCertificates) GetList() interface{} {
	return &l.Certificates
}

// CreateBackendSet Adds a backend set to a load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/CreateBackendSet
func (c *Client) CreateCertificate(
	loadBalancerID string,
	certificateName string,
	caCertificate string,
	privateKey string,
	passphrase string,
	publicCertificate string,
	opts *LoadBalancerOptions,
) (workRequestID string, e error) {

	required := Certificate{
		CertificateName:   certificateName,
		PublicCertificate: publicCertificate,
		PrivateKey:        privateKey,
		CACertificate:     caCertificate,
		Passphrase:        passphrase,
	}

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceCertificates,
		},
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.postRequest(details); e != nil {
		return
	}

	cert := &Certificate{}
	e = resp.unmarshal(cert)
	if e == nil {
		workRequestID = cert.WorkRequestID
	}
	return
}

// ListCertificates Lists all SSL certificates associated with a given load balancer.
//
// See: https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/Certificate/ListCertificates
func (c *Client) ListCertificates(
	loadBalancerID string,
	opts *ClientRequestOptions,
) (certs *ListCertificates, e error) {
	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceCertificates,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.getRequest(details); e != nil {
		return
	}

	certs = &ListCertificates{}
	e = resp.unmarshal(certs)
	return
}

// Deletes the specified backend set. Note that deleting a backend set removes its backend servers from the load balancer.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/loadbalancer/20170115/BackendSet/DeleteBackendSet
func (c *Client) DeleteCertificate(
	loadBalancerID string,
	certificateName string,
	opts *ClientRequestOptions,
) (workRequestID string, e error) {

	details := &requestDetails{
		name: resourceLoadBalancers,
		ids: urlParts{
			loadBalancerID,
			resourceCertificates,
			certificateName,
		},
		optional: opts,
	}

	var resp *response
	if resp, e = c.loadBalancerApi.request(http.MethodDelete, details); e != nil {
		return
	}

	cert := &Certificate{}
	e = resp.unmarshal(cert)
	if e == nil {
		workRequestID = cert.WorkRequestID
	}
	return
}
