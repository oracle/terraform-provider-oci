package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

type ShapeReader struct {
	resourceData      *schema.ResourceData
	client            client.BareMetalClient
	shapeListResponse *baremetal.ShapeList
}

func ResourceCoreShape() *schema.Resource {
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
	reader := &ShapeReader{
		resourceData: d,
		client:       client,
	}

	return crud.ReadResource(reader)

}

func (r *ShapeReader) Get() (e error) {
	compartmentID := r.resourceData.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(r.resourceData, "availability_domain", "image_id")

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
		shapes := []map[string]string{}
		for _, v := range r.shapeListResponse.Shapes {
			shape := map[string]string{
				"name": v.Name,
			}
			shapes = append(shapes, shape)
		}
		r.resourceData.Set("shapes", shapes)
	}
	return
}
