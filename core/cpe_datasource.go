package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func CpeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCpeList,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CpeResource(),
			},
		},
	}
}

func readCpeList(d *schema.ResourceData, m interface{}) (e error) {
	reader := &CPEDatasourceCrud{
		D:      d,
		Client: m.(client.BareMetalClient),
	}
	return crud.ReadResource(reader)

}
