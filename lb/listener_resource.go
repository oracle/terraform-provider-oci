// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerListenerResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancerListener,
		Read:   readLoadBalancerListener,
		Update: updateLoadBalancerListener,
		Delete: deleteLoadBalancerListener,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"default_backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_configuration": SSLConfigSchema,
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}
