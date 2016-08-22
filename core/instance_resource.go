package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func ResourceCoreInstance() *schema.Resource {
	return &schema.Resource{
		Create: createInstance,
		Read:   readInstance,
		Update: updateInstance,
		Delete: deleteInstance,
		Schema: map[string]*schema.Schema{
			"availability_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem:     schema.TypeString,
				ForceNew: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"shape": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"time_created": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.CreateResource(d, sync)
}

func readInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.ReadResource(sync)
}

func updateInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InstanceSync{D: d, Client: m.(client.BareMetalClient)}
	return crud.UpdateResource(d, sync)
}

func deleteInstance(d *schema.ResourceData, m interface{}) (e error) {
	sync := InstanceSync{D: d, Client: m.(client.BareMetalClient)}
	return sync.Delete()
}
