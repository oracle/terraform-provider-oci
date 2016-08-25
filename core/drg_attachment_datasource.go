package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func DrgAttachmentDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readDrgAttachments,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"drg_attachments": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     DrgAttachmentResource(),
			},
		},
	}
}

func readDrgAttachments(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &DrgAttachmentDatasourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}
