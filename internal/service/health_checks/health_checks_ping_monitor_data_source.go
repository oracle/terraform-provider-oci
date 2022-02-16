// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package health_checks

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/v58/healthchecks"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func HealthChecksPingMonitorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitor_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(HealthChecksPingMonitorResource(), fieldMap, readSingularHealthChecksPingMonitor)
}

func readSingularHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

type HealthChecksPingMonitorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_health_checks.HealthChecksClient
	Res    *oci_health_checks.GetPingMonitorResponse
}

func (s *HealthChecksPingMonitorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HealthChecksPingMonitorDataSourceCrud) Get() error {
	request := oci_health_checks.GetPingMonitorRequest{}

	if monitorId, ok := s.D.GetOkExists("monitor_id"); ok {
		tmp := monitorId.(string)
		request.MonitorId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "health_checks")

	response, err := s.Client.GetPingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *HealthChecksPingMonitorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HomeRegion != nil {
		s.D.Set("home_region", *s.Res.HomeRegion)
	}

	if s.Res.IntervalInSeconds != nil {
		s.D.Set("interval_in_seconds", *s.Res.IntervalInSeconds)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	s.D.Set("protocol", s.Res.Protocol)

	if s.Res.ResultsUrl != nil {
		s.D.Set("results_url", *s.Res.ResultsUrl)
	}

	s.D.Set("targets", s.Res.Targets)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	s.D.Set("vantage_point_names", s.Res.VantagePointNames)

	return nil
}
