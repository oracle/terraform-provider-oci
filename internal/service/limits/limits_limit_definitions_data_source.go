// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v56/limits"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LimitsLimitDefinitionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLimitsLimitDefinitions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit_definitions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"are_quotas_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_deprecated": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_dynamic": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_eligible_for_limit_increase": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_resource_availability_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"scope_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLimitsLimitDefinitions(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsLimitDefinitionsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LimitsClient()

	return tfresource.ReadResource(sync)
}

type LimitsLimitDefinitionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.LimitsClient
	Res    *oci_limits.ListLimitDefinitionsResponse
}

func (s *LimitsLimitDefinitionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsLimitDefinitionsDataSourceCrud) Get() error {
	request := oci_limits.ListLimitDefinitionsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "limits")

	response, err := s.Client.ListLimitDefinitions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLimitDefinitions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LimitsLimitDefinitionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LimitsLimitDefinitionsDataSource-", LimitsLimitDefinitionsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		limitDefinition := map[string]interface{}{}

		if r.AreQuotasSupported != nil {
			limitDefinition["are_quotas_supported"] = *r.AreQuotasSupported
		}

		if r.Description != nil {
			limitDefinition["description"] = *r.Description
		}

		if r.IsDeprecated != nil {
			limitDefinition["is_deprecated"] = *r.IsDeprecated
		}

		if r.IsDynamic != nil {
			limitDefinition["is_dynamic"] = *r.IsDynamic
		}

		if r.IsEligibleForLimitIncrease != nil {
			limitDefinition["is_eligible_for_limit_increase"] = *r.IsEligibleForLimitIncrease
		}

		if r.IsResourceAvailabilitySupported != nil {
			limitDefinition["is_resource_availability_supported"] = *r.IsResourceAvailabilitySupported
		}

		if r.Name != nil {
			limitDefinition["name"] = *r.Name
		}

		limitDefinition["scope_type"] = r.ScopeType

		if r.ServiceName != nil {
			limitDefinition["service_name"] = *r.ServiceName
		}

		resources = append(resources, limitDefinition)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LimitsLimitDefinitionsDataSource().Schema["limit_definitions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("limit_definitions", resources); err != nil {
		return err
	}

	return nil
}
