// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import "github.com/hashicorp/terraform/helper/schema"

var HealthCheckerSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"interval_ms": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  30000,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"response_body_regex": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	},
}

var SSLConfigSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	MaxItems: 1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			"certificate_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"verify_depth": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5,
			},
			"verify_peer_certificate": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	},
}
