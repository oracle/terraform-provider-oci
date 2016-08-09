package main

import (
	"time"

	"github.com/MustWin/baremtlclient"
	"github.com/hashicorp/terraform/helper/schema"
)

type ShapeReader struct {
	resourceData      *schema.ResourceData
	client            BareMetalClient
	shapeListResponse *baremtlsdk.ShapeList
}

func ResourceCoreShape() *schema.Resource {
	return &schema.Resource{
		Read: readShape,
		Schema: map[string]*schema.Schema{
			"shapes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
	client := m.(BareMetalClient)
	reader := &ShapeReader{
		resourceData: d,
		client:       client,
	}

	return readResource(reader)

}

func (r *ShapeReader) getOptions() (opts []baremtlsdk.CoreOptions) {
	keys := []string{
		"availability_domain",
		"image_id",
	}

	opts = []baremtlsdk.CoreOptions{}

	for _, key := range keys {
		if val, ok := r.resourceData.GetOk(key); ok {
			if len(opts) == 0 {
				opts = append(opts, baremtlsdk.CoreOptions{})
			}

			switch key {
			case "availability_domain":
				opts[0].AvailabilityDomain = val.(string)
			case "image_id":
				opts[0].ImageID = val.(string)

			}
		}
	}

	return

}

func (r *ShapeReader) Get() (e error) {
	compartmentID := r.resourceData.Get("compartment_id").(string)
	opts := r.getOptions()

	if r.shapeListResponse, e = r.client.ListShapes(compartmentID, opts...); e != nil {
		return
	}

	return
}

func (r *ShapeReader) SetData() {
	if r.shapeListResponse != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		r.resourceData.SetId(time.Now().UTC().String())
		r.resourceData.Set("shapes", r.shapeListResponse.Shapes)
	}
	return
}
