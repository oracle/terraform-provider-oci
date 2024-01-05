// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v65/limits"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LimitsServicesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLimitsServices,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"services": {
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

func readLimitsServices(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsServicesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LimitsClient()

	return tfresource.ReadResource(sync)
}

type LimitsServicesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.LimitsClient
	Res    *oci_limits.ListServicesResponse
}

func (s *LimitsServicesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsServicesDataSourceCrud) Get() error {
	request := oci_limits.ListServicesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "limits")

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

func (s *LimitsServicesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LimitsServicesDataSource-", LimitsServicesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		service := map[string]interface{}{}

		if r.Description != nil {
			service["description"] = *r.Description
		}

		if r.Name != nil {
			service["name"] = *r.Name
		}

		resources = append(resources, service)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LimitsServicesDataSource().Schema["services"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("services", resources); err != nil {
		return err
	}

	return nil
}
