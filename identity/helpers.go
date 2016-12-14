package identity

import "github.com/hashicorp/terraform/helper/schema"

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
