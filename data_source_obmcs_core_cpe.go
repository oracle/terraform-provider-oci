// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
)

func CpeDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readCpeList,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"page": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cpes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     CpeResource(),
			},
		},
	}
}

func readCpeList(d *schema.ResourceData, m interface{}) (e error) {
	reader := &CPEDatasourceCrud{}
	reader.D = d
	reader.Client = m.(*baremetal.Client)
	return crud.ReadResource(reader)

}

type CPEDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.ListCpes
}

func (s *CPEDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Resource = &baremetal.ListCpes{Cpes: []baremetal.Cpe{}}

	for {
		var list *baremetal.ListCpes
		if list, e = s.Client.ListCpes(compartmentID, opts); e != nil {
			break
		}

		s.Resource.Cpes = append(s.Resource.Cpes, list.Cpes...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s *CPEDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())

		cpes := []map[string]interface{}{}

		for _, v := range s.Resource.Cpes {
			cpe := map[string]interface{}{
				"id":             v.ID,
				"compartment_id": v.CompartmentID,
				"display_name":   v.DisplayName,
				"ip_address":     v.IPAddress,
				"time_created":   v.TimeCreated.String(),
			}

			cpes = append(cpes, cpe)
		}

		s.D.Set("cpes", cpes)

	}
}
