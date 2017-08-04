// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"errors"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerCertificateResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancerCertificate,
		Read:   readLoadBalancerCertificate,
		Delete: deleteLoadBalancerCertificate,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ca_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"certificate_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"private_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_certificate": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerCertificate(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readLoadBalancerCertificate(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func deleteLoadBalancerCertificate(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerCertificateResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

type LoadBalancerCertificateResourceCrud struct {
	crud.BaseCrud
	WorkRequest *baremetal.WorkRequest
	Resource    *baremetal.Certificate
}

func (s *LoadBalancerCertificateResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Resource, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("certificate_name").(string)
	}
	return ""
}

func (s *LoadBalancerCertificateResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerCertificateResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
		baremetal.ResourceFailed,
	}
}

func (s *LoadBalancerCertificateResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerCertificateResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
	}
}

func (s *LoadBalancerCertificateResourceCrud) Create() (e error) {
	opts := &baremetal.LoadBalancerOptions{}

	var workReqID string
	workReqID, e = s.Client.CreateCertificate(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("certificate_name").(string),
		s.D.Get("ca_certificate").(string),
		s.D.Get("private_key").(string),
		s.D.Get("passphrase").(string),
		s.D.Get("public_certificate").(string),
		opts,
	)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerCertificateResourceCrud) Get() (e error) {

	_, stillWorking, err := crud.LoadBalancerResourceGet(s.BaseCrud, s.WorkRequest)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}

	var list *baremetal.ListCertificates
	list, e = s.Client.ListCertificates(s.D.Get("load_balancer_id").(string), nil)
	if e != nil {
		return
	}
	for _, cert := range list.Certificates {
		if cert.CertificateName == s.D.Get("certificate_name").(string) {
			s.Resource = &cert
			return
		}
	}
	e = errors.New("Certificate does not exist")
	return
}

func (s *LoadBalancerCertificateResourceCrud) SetData() {
	// Noop for this resource
}

func (s *LoadBalancerCertificateResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteCertificate(s.D.Get("load_balancer_id").(string), s.D.Get("certificate_name").(string), nil)
	if e != nil {
		return
	}
	s.D.SetId(workReqID)
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
