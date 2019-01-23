// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_health_checks "github.com/oracle/oci-go-sdk/healthchecks"
)

func HealthChecksPingMonitorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularHealthChecksPingMonitor,
		Schema: map[string]*schema.Schema{
			"monitor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"interval_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"results_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vantage_point_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).healthChecksClient

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "health_checks")

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
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	s.D.Set("vantage_point_names", s.Res.VantagePointNames)

	return nil
}
