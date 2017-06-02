// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceObjectStorageMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

var bucketSchema = map[string]*schema.Schema{
	"compartment_id": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"namespace": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"access_type": {
		Type:     schema.TypeString,
		Computed: false,
		Optional: true,
	},
	"metadata": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}

var objectSchema = map[string]*schema.Schema{
	"namespace": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"bucket": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"object": {
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"content": {
		Type:     schema.TypeString,
		Optional: true,
	},
	"metadata": {
		Type:     schema.TypeMap,
		Optional: true,
	},
}
