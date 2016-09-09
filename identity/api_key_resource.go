package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func APIKeyResource() *schema.Resource {
	return &schema.Resource{
		Create: createAPIKey,
		Delete: deleteAPIKey,
		Schema: map[string]*schema.Schema{
			"key_value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func deleteAPIKey(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyResourceCrud{D: d, Client: client}
	return crud.DeleteResource(sync)
}
