package objectstorage

import "github.com/hashicorp/terraform/helper/schema"

var bucketSchema = map[string]*schema.Schema{
	"compartment_id": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"namespace": &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
		Computed: false,
	},
	"metadata": &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
	},
}
