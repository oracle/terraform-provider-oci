// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/v65/healthchecks"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func HealthChecksHttpMonitorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHealthChecksHttpMonitors,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"home_region": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_monitors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(HealthChecksHttpMonitorResource()),
			},
		},
	}
}

func readHealthChecksHttpMonitors(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksHttpMonitorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

type HealthChecksHttpMonitorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.ListHttpMonitorsResponse
}

func (s *HealthChecksHttpMonitorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksHttpMonitorsDataSourceCrud) Get() error {
	request := oci_health_checks.ListHttpMonitorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if homeRegion, ok := s.D.GetOkExists("home_region"); ok {
		tmp := homeRegion.(string)
		request.HomeRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "health_checks")

	response, err := s.Client.ListHttpMonitors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListHttpMonitors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *HealthChecksHttpMonitorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("HealthChecksHttpMonitorsDataSource-", HealthChecksHttpMonitorsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		httpMonitor := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			httpMonitor["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			httpMonitor["display_name"] = *r.DisplayName
		}

		httpMonitor["freeform_tags"] = r.FreeformTags

		if r.HomeRegion != nil {
			httpMonitor["home_region"] = *r.HomeRegion
		}

		if r.Id != nil {
			httpMonitor["id"] = *r.Id
		}

		if r.IntervalInSeconds != nil {
			httpMonitor["interval_in_seconds"] = *r.IntervalInSeconds
		}

		if r.IsEnabled != nil {
			httpMonitor["is_enabled"] = *r.IsEnabled
		}

		httpMonitor["protocol"] = r.Protocol

		if r.ResultsUrl != nil {
			httpMonitor["results_url"] = *r.ResultsUrl
		}

		if r.TimeCreated != nil {
			httpMonitor["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, httpMonitor)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, HealthChecksHttpMonitorsDataSource().Schema["http_monitors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("http_monitors", resources); err != nil {
		return err
	}

	return nil
}
