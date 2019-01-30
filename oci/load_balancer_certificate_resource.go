// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func CertificateResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createCertificate,
		Read:     readCertificate,
		Delete:   deleteCertificate,
		Schema: map[string]*schema.Schema{
			// Required
			"certificate_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passphrase": {
				Type:      schema.TypeString,
				Optional:  true,
				Default:   "",
				ForceNew:  true,
				Sensitive: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				ForceNew:  true,
				Sensitive: true,
			},
			"public_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return CreateResource(d, sync)
}

func readCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return ReadResource(sync)
}

func deleteCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &CertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type CertificateResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Certificate
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *CertificateResourceCrud) ID() string {
	id, workSuccess := LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("certificate_name").(string)
	}
	return ""
}

func (s *CertificateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *CertificateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *CertificateResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *CertificateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *CertificateResourceCrud) Create() error {
	request := oci_load_balancer.CreateCertificateRequest{}

	if caCertificate, ok := s.D.GetOkExists("ca_certificate"); ok {
		tmp := caCertificate.(string)
		request.CaCertificate = &tmp
	}

	if certificateName, ok := s.D.GetOkExists("certificate_name"); ok {
		tmp := certificateName.(string)
		request.CertificateName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if passphrase, ok := s.D.GetOkExists("passphrase"); ok {
		tmp := passphrase.(string)
		request.Passphrase = &tmp
	}

	if privateKey, ok := s.D.GetOkExists("private_key"); ok {
		tmp := privateKey.(string)
		request.PrivateKey = &tmp
	}

	if publicCertificate, ok := s.D.GetOkExists("public_certificate"); ok {
		tmp := publicCertificate.(string)
		request.PublicCertificate = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateCertificate(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *CertificateResourceCrud) Get() error {
	_, stillWorking, err := LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.ListCertificatesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	certificateName := s.D.Get("certificate_name").(string)
	for _, item := range response.Items {
		if *item.CertificateName == certificateName {
			s.Res = &item
			return nil
		}
	}
	return errors.New("Certificate with expected identifier not found")

}

func (s *CertificateResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteCertificateRequest{}

	if certificateName, ok := s.D.GetOkExists("certificate_name"); ok {
		tmp := certificateName.(string)
		request.CertificateName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteCertificate(context.Background(), request)

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *CertificateResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}
	if s.Res.CaCertificate != nil {
		s.D.Set("ca_certificate", *s.Res.CaCertificate)
	}

	if s.Res.CertificateName != nil {
		s.D.Set("certificate_name", *s.Res.CertificateName)
	}

	if s.Res.PublicCertificate != nil {
		s.D.Set("public_certificate", *s.Res.PublicCertificate)
	}

	return nil
}
