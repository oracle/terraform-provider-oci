package identity

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func APIKeyDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readAPIKeys,
		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"api_keys": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     APIKeyResource(),
			},
		},
	}
}

func readAPIKeys(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &APIKeyDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
