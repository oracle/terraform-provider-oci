package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func SecurityListDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readSecurityLists,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"security_lists": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     SecurityListResource(),
			},
		},
	}
}

func readSecurityLists(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &SecurityListDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
