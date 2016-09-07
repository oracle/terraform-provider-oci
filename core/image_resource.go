package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ImageResource() *schema.Resource {
	return &schema.Resource{
		Create: createImage,
		Read:   readImage,
		Update: updateImage,
		Delete: deleteImage,
		Schema: map[string]*schema.Schema{
			"base_image_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"create_image_allowed": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"operating_system_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{D: d, Client: client}
	return crud.CreateResource(d, sync)
}

func readImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{D: d, Client: client}
	return crud.ReadResource(sync)
}

func updateImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{D: d, Client: client}
	return crud.UpdateResource(d, sync)
}

func deleteImage(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &ImageResourceCrud{D: d, Client: client}
	return crud.DeleteResource(sync)
}
