// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IPSecConnectionsDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readIPSecConnections,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpe_id": {
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
			"connections": {
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
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"drg_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpe_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_routes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readIPSecConnections(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	reader := &IPSecConnectionsDatasourceCrud{}
	reader.D = d
	reader.Client = client

	return crud.ReadResource(reader)
}

type IPSecConnectionsDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.ListIPSecConnections
}

func (s *IPSecConnectionsDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)

	opts := &baremetal.ListIPSecConnsOptions{}
	options.SetListOptions(s.D, &opts.ListOptions)
	if val, ok := s.D.GetOk("cpe_id"); ok {
		opts.CpeID = val.(string)
	}
	if val, ok := s.D.GetOk("drg_id"); ok {
		opts.DrgID = val.(string)
	}

	s.Resource = &baremetal.ListIPSecConnections{
		Connections: []baremetal.IPSecConnection{},
	}

	for {
		var list *baremetal.ListIPSecConnections
		if list, e = s.Client.ListIPSecConnections(compartmentID, opts); e != nil {
			break
		}

		s.Resource.Connections = append(s.Resource.Connections, list.Connections...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.ListOptions.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s IPSecConnectionsDatasourceCrud) SetData() {
	if s.Resource != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}

		for _, v := range s.Resource.Connections {

			resource := map[string]interface{}{
				"compartment_id": v.CompartmentID,
				"drg_id":         v.DrgID,
				"cpe_id":         v.CpeID,
				"display_name":   v.DisplayName,
				"id":             v.ID,
				"state":          v.State,
				"static_routes":  v.StaticRoutes,
				"time_created":   v.TimeCreated.String(),
			}

			resources = append(resources, resource)
		}

		s.D.Set("connections", resources)

	}

	return
}
