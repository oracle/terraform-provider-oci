// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/options"

	"github.com/oracle/terraform-provider-oci/crud"
)

func InternetGatewayDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readInternetGateways,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcn_id": {
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
			"gateways": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     InternetGatewayResource(),
			},
		},
	}
}

func readInternetGateways(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*OracleClients)
	reader := &InternetGatewayDatasourceCrud{}
	reader.D = d
	reader.Client = client.client

	return crud.ReadResource(reader)
}

type InternetGatewayDatasourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.ListInternetGateways
}

func (s *InternetGatewayDatasourceCrud) Get() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	vcnID := s.D.Get("vcn_id").(string)

	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	s.Resource = &baremetal.ListInternetGateways{
		Gateways: []baremetal.InternetGateway{},
	}

	for {
		var list *baremetal.ListInternetGateways
		if list, e = s.Client.ListInternetGateways(compartmentID, vcnID, opts); e != nil {
			break
		}

		s.Resource.Gateways = append(s.Resource.Gateways, list.Gateways...)

		if hasNextPage := options.SetNextPageOption(list.NextPage, &opts.PageListOptions); !hasNextPage {
			break
		}
	}

	return
}

func (s InternetGatewayDatasourceCrud) SetData() {
	if s.Resource == nil {
		return
	}
	s.D.SetId(time.Now().UTC().String())
	resources := []map[string]interface{}{}

	for _, v := range s.Resource.Gateways {

		resource := map[string]interface{}{
			"compartment_id": v.CompartmentID,
			"display_name":   v.DisplayName,
			"id":             v.ID,
			"enabled":        v.IsEnabled,
			"state":          v.State,
			"time_modified":  v.ModifiedTime.String(),
			"time_created":   v.TimeCreated.String(),
		}

		resources = append(resources, resource)
	}

	if f, fOk := s.D.GetOk("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("gateways", resources); err != nil {
		panic(err)
	}

	return
}
