// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func LoadBalancerBackendResource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backendset_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"drain": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"offline": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
