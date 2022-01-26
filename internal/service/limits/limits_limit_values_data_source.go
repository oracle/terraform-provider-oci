// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package limits

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_limits "github.com/oracle/oci-go-sdk/v56/limits"

	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LimitsLimitValuesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLimitsLimitValues,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"scope_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"availability_domain": {
							Type:     schema.TypeString,
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
						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readLimitsLimitValues(d *schema.ResourceData, m interface{}) error {
	sync := &LimitsLimitValuesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LimitsClient()

	return tfresource.ReadResource(sync)
}

type LimitsLimitValuesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_limits.LimitsClient
	Res    *oci_limits.ListLimitValuesResponse
}

func (s *LimitsLimitValuesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LimitsLimitValuesDataSourceCrud) Get() error {
	request := oci_limits.ListLimitValuesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if scopeType, ok := s.D.GetOkExists("scope_type"); ok {
		request.ScopeType = oci_limits.ListLimitValuesScopeTypeEnum(scopeType.(string))
	}

	if serviceName, ok := s.D.GetOkExists("service_name"); ok {
		tmp := serviceName.(string)
		request.ServiceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "limits")

	response, err := s.Client.ListLimitValues(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListLimitValues(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *LimitsLimitValuesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LimitsLimitValuesDataSource-", LimitsLimitValuesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		limitValue := map[string]interface{}{}

		if r.AvailabilityDomain != nil {
			limitValue["availability_domain"] = *r.AvailabilityDomain
		}

		if r.Name != nil {
			limitValue["name"] = *r.Name
		}

		limitValue["scope_type"] = r.ScopeType

		if r.Value != nil {
			limitValue["value"] = strconv.FormatInt(*r.Value, 10)
		}

		resources = append(resources, limitValue)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LimitsLimitValuesDataSource().Schema["limit_values"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("limit_values", resources); err != nil {
		return err
	}

	return nil
}
