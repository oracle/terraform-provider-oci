package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ShapeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readShape,
		Schema: map[string]*schema.Schema{
			"shapes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func readShape(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &ShapeDatasourceCrud{
		D:      d,
		Client: client,
	}

	return crud.ReadResource(reader)

}
