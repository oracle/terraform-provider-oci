// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v38/loadbalancer"
)

func init() {
	RegisterResource("oci_load_balancer_certificate", LoadBalancerCertificateResource())
}

func LoadBalancerCertificateResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createLoadBalancerCertificate,
		Read:     readLoadBalancerCertificate,
		Delete:   deleteLoadBalancerCertificate,
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

func createLoadBalancerCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return CreateResource(d, sync)
}

func readLoadBalancerCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

func deleteLoadBalancerCertificate(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type LoadBalancerCertificateResourceCrud struct {
	BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Certificate
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerCertificateResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			return getCertificateCompositeId(s.D.Get("certificate_name").(string), s.D.Get("load_balancer_id").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerCertificateResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerCertificateResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerCertificateResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerCertificateResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerCertificateResourceCrud) Create() error {
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

func (s *LoadBalancerCertificateResourceCrud) Get() error {
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

	certificateName := s.D.Get("certificate_name").(string)

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		certNameFromId, loadBalancerId, err := parseCertificateCompositeId(s.D.Id())
		if err == nil {
			certificateName = certNameFromId
			request.LoadBalancerId = &loadBalancerId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.ListCertificates(context.Background(), request)
	if err != nil {
		return err
	}

	for _, item := range response.Items {
		if *item.CertificateName == certificateName {
			s.Res = &item
			return nil
		}
	}
	return errors.New("Certificate with expected identifier not found")

}

func (s *LoadBalancerCertificateResourceCrud) Delete() error {
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

func (s *LoadBalancerCertificateResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	certificateName, loadBalancerId, err := parseCertificateCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("certificate_name", &certificateName)
		s.D.Set("load_balancer_id", &loadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
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

func getCertificateCompositeId(certificateName string, loadBalancerId string) string {
	certificateName = url.PathEscape(certificateName)
	loadBalancerId = url.PathEscape(loadBalancerId)
	compositeId := "loadBalancers/" + loadBalancerId + "/certificates/" + certificateName
	return compositeId
}

func parseCertificateCompositeId(compositeId string) (certificateName string, loadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/certificates/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	certificateName, _ = url.PathUnescape(parts[3])

	return
}
