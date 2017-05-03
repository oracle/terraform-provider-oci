// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
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
				Required: true,
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
