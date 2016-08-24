package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DatasourceCoreIPSecConnections() *schema.Resource {
	return &schema.Resource{
		Read: readIPSecConnections,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpe_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"connections": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     datasourceIPSecConnections(),
			},
		},
	}
}

func datasourceIPSecConnections() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"drg_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpe_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_routes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readIPSecConnections(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	reader := &IPSecDatasourceCrud{
		D:      d,
		Client: client,
	}

	return crud.ReadResource(reader)
}
