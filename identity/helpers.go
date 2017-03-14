// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import "github.com/hashicorp/terraform/helper/schema"

var baseIdentitySchemaWithID = map[string]*schema.Schema{
	"id": {
		Type: schema.TypeString,
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

// Just has a computed compartment_id
var baseIdentitySchema = map[string]*schema.Schema{
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
	"time_created": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"time_modified": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

var identitySchema = map[string]*schema.Schema{
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
		Required: true,
	},
	"state": {
		Type:     schema.TypeString,
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

var identitySchemaWithID = map[string]*schema.Schema{
	"id": {
		Type: schema.TypeString,
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
		Required: true,
	},
	"state": {
		Type:     schema.TypeString,
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

