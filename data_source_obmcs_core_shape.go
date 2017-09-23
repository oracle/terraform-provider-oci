// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/oracle/bmcs-go-sdk"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InstanceShapeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInstanceShape,
		Schema: map[string]*schema.Schema{
			"shapes": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func readInstanceShape(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &InstanceShapeDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)

}

type InstanceShapeDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListShapes
}

func (r *InstanceShapeDatasourceCrud) Get() (e error) {
	compartmentID := r.D.Get("compartment_id").(string)

	opts := &baremetal.ListShapesOptions{}
	options.SetListOptions(r.D, &opts.ListOptions)
	if val, ok := r.D.GetOk("availability_domain"); ok {
		opts.AvailabilityDomain = val.(string)
	}
	if val, ok := r.D.GetOk("image_id"); ok {
		opts.ImageID = val.(string)
	}

	r.Res = &baremetal.ListShapes{Shapes: []baremetal.Shape{}}

	for {
		var list *baremetal.ListShapes
		if list, e = r.Client.ListShapes(compartmentID, opts); e != nil {
			break
		}

		r.Res.Shapes = append(r.Res.Shapes, list.Shapes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (r *InstanceShapeDatasourceCrud) SetData() {
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
