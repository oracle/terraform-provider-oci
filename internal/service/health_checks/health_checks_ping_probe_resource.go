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

func HealthChecksPingProbeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createHealthChecksPingProbe,
		Read:     readHealthChecksPingProbe,
		Delete:   deleteHealthChecksPingProbe,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vantage_point_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Computed
			"home_region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"results_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createHealthChecksPingProbe(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingProbeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.CreateResource(d, sync)
}

func readHealthChecksPingProbe(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteHealthChecksPingProbe(d *schema.ResourceData, m interface{}) error {
	return nil
}

type HealthChecksPingProbeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_health_checks.HealthChecksClient
	Res                    *oci_health_checks.PingProbe
	DisableNotFoundRetries bool
}

func (s *HealthChecksPingProbeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *HealthChecksPingProbeResourceCrud) Create() error {
	request := oci_health_checks.CreateOnDemandPingProbeRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_health_checks.PingProbeProtocolEnum(protocol.(string))
	}

	if targets, ok := s.D.GetOkExists("targets"); ok {
		interfaces := targets.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("targets") {
			request.Targets = tmp
		}
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	if vantagePointNames, ok := s.D.GetOkExists("vantage_point_names"); ok {
		interfaces := vantagePointNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("vantage_point_names") {
			request.VantagePointNames = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.CreateOnDemandPingProbe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingProbe
	return nil
}

func (s *HealthChecksPingProbeResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.HomeRegion != nil {
		s.D.Set("home_region", *s.Res.HomeRegion)
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
