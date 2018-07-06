// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func CrossConnectLocationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCrossConnectLocations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cross_connect_locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCrossConnectLocations(d *schema.ResourceData, m interface{}) error {
	sync := &CrossConnectLocationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type CrossConnectLocationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListCrossConnectLocationsResponse
}

func (s *CrossConnectLocationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CrossConnectLocationsDataSourceCrud) Get() error {
	request := oci_core.ListCrossConnectLocationsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListCrossConnectLocations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCrossConnectLocations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CrossConnectLocationsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		crossConnectLocation := map[string]interface{}{}

		if r.Description != nil {
			crossConnectLocation["description"] = *r.Description
		}

		if r.Name != nil {
			crossConnectLocation["name"] = *r.Name
		}

		resources = append(resources, crossConnectLocation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, CrossConnectLocationsDataSource().Schema["cross_connect_locations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cross_connect_locations", resources); err != nil {
		panic(err)
	}

	return
}
