package core

import (
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

func InternetGatewayResource() *schema.Resource {
	return &schema.Resource{
		Create: createInternetGateway,
		Read:   readInternetGateway,
		Update: updateInternetGateway,
		Delete: deleteInternetGateway,
		Schema: map[string]*schema.Schema{
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vcn_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
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
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_modified": &schema.Schema{
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

func createInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{
		D:      d,
		Client: m.(client.BareMetalClient)}

	return crud.CreateResource(d, sync)
}

func readInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{
		D:      d,
		Client: m.(client.BareMetalClient)}

	return crud.ReadResource(sync)
}

func updateInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{
		D:      d,
		Client: m.(client.BareMetalClient)}

	return crud.UpdateResource(d, sync)

}

func deleteInternetGateway(d *schema.ResourceData, m interface{}) (e error) {
	sync := &InternetGatewayResourceCrud{D: d, Client: m.(client.BareMetalClient)}
	return crud.DeleteResource(sync)
}
