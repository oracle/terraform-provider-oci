// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/terraform-providers/terraform-provider-oci/crud"
)

func ServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readServices,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"cidr_block": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
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

func readServices(d *schema.ResourceData, m interface{}) error {
	sync := &ServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.ReadResource(sync)
}

type ServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.VirtualNetworkClient
	Res    *oci_core.ListServicesResponse
}

func (s *ServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ServicesDataSourceCrud) Get() error {
	request := oci_core.ListServicesRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListServices(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListServices(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ServicesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		service := map[string]interface{}{}

		if r.CidrBlock != nil {
			service["cidr_block"] = *r.CidrBlock
		}

		if r.Description != nil {
			service["description"] = *r.Description
		}

		if r.Id != nil {
			service["id"] = *r.Id
		}

		if r.Name != nil {
			service["name"] = *r.Name
		}

		resources = append(resources, service)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources)
	}

	if err := s.D.Set("services", resources); err != nil {
		panic(err)
	}

	return
}
