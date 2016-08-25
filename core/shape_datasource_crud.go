package core

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/hashicorp/terraform/helper/schema"
)

type ShapeDatasourceCrud struct {
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.ShapeList
}

func (r *ShapeDatasourceCrud) Get() (e error) {
	compartmentID := r.D.Get("compartment_id").(string)
	opts := getCoreOptionsFromResourceData(r.D, "availability_doresource", "image_id")

	if r.Res, e = r.Client.ListShapes(compartmentID, opts...); e != nil {
		return
	}

	return
}

func (r *ShapeDatasourceCrud) SetData() {
	if r.Res != nil {
		// Important, if you don't have an ID, make one up for your datasource
		// or things will end in tears
		r.D.SetId(time.Now().UTC().String())
		shapes := []map[string]string{}
		for _, v := range r.Res.Shapes {
			shape := map[string]string{
				"name": v.Name,
			}
			shapes = append(shapes, shape)
		}
		r.D.Set("shapes", shapes)
	}
	return
}
