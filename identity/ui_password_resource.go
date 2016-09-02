package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

// Id is exposed to allow resetting an existing user's password.
func UIPasswordResource() *schema.Resource {
	return &schema.Resource{
		Create: createUIPassword,
		Read:   readUIPassword,
		Delete: deleteUIPassword,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
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
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &UIPasswordResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	return nil
}

func deleteUIPassword(d *schema.ResourceData, m interface{}) (e error) {
	return nil
}
