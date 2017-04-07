// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

var transportSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"max": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"min": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
}

var icmpSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"code": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	},
}

func SecurityListResource() *schema.Resource {
	return &schema.Resource{
		Create: createSecurityList,
		Read:   readSecurityList,
		Update: updateSecurityList,
		Delete: deleteSecurityList,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"egress_security_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ingress_security_rules": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"icmp_options": icmpSchema,
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tcp_options": transportSchema,
						"udp_options": transportSchema,
						"stateless": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.CreateResource(d, crd)
}

func readSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.ReadResource(crd)
}

func updateSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.UpdateResource(d, crd)
}

func deleteSecurityList(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	crd := &SecurityListResourceCrud{}
	crd.D = d
	crd.Client = client
	return crud.DeleteResource(crd)
}
