// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import "github.com/hashicorp/terraform/helper/schema"

// User and group happen to have the same schema and share this
var baseIdentitySchema = map[string]*schema.Schema{
	"id": {
		Type:     schema.TypeString,
		Computed: true,
		ForceNew: true,
	},
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"description": {
		Type:     schema.TypeString,
		Required: true,
	},
	"compartment_id": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"state": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"inactive_state": {
		Type:     schema.TypeInt,
		Computed: true,
	},
	"time_created": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"time_modified": {
		Type:     schema.TypeString,
		Computed: true,
	},
}
