// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
)

func HealthChecksPingMonitorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHealthChecksPingMonitors,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"ping_monitors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(HealthChecksPingMonitorResource()),
			},
		},
	}
}

func readHealthChecksPingMonitors(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return ReadResource(sync)
}

type HealthChecksPingMonitorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.ListPingMonitorsResponse
}

func (s *HealthChecksPingMonitorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksPingMonitorsDataSourceCrud) Get() error {
	request := oci_health_checks.ListPingMonitorsRequest{}

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "health_checks")

	response, err := s.Client.ListPingMonitors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPingMonitors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *HealthChecksPingMonitorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		pingMonitor := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			pingMonitor["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			pingMonitor["display_name"] = *r.DisplayName
		}

		pingMonitor["freeform_tags"] = r.FreeformTags

		if r.HomeRegion != nil {
			pingMonitor["home_region"] = *r.HomeRegion
		}

		if r.Id != nil {
			pingMonitor["id"] = *r.Id
		}

		if r.IntervalInSeconds != nil {
			pingMonitor["interval_in_seconds"] = *r.IntervalInSeconds
		}

		if r.IsEnabled != nil {
			pingMonitor["is_enabled"] = *r.IsEnabled
		}

		pingMonitor["protocol"] = r.Protocol

		if r.ResultsUrl != nil {
			pingMonitor["results_url"] = *r.ResultsUrl
		}

		if r.TimeCreated != nil {
			pingMonitor["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, pingMonitor)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, HealthChecksPingMonitorsDataSource().Schema["ping_monitors"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("ping_monitors", resources); err != nil {
		return err
	}

	return nil
}
