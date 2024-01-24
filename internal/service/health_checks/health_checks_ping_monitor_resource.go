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

func HealthChecksPingMonitorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createHealthChecksPingMonitor,
		Read:     readHealthChecksPingMonitor,
		Update:   updateHealthChecksPingMonitor,
		Delete:   deleteHealthChecksPingMonitor,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"interval_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"targets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vantage_point_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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

func createHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.CreateResource(d, sync)
}

func readHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.ReadResource(sync)
}

func updateHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteHealthChecksPingMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &HealthChecksPingMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).HealthChecksClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type HealthChecksPingMonitorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_health_checks.HealthChecksClient
	Res                    *oci_health_checks.PingMonitor
	DisableNotFoundRetries bool
}

func (s *HealthChecksPingMonitorResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *HealthChecksPingMonitorResourceCrud) Create() error {
	request := oci_health_checks.CreatePingMonitorRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
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

	response, err := s.Client.CreatePingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Get() error {
	request := oci_health_checks.GetPingMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	response, err := s.Client.GetPingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_health_checks.UpdatePingMonitorRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if intervalInSeconds, ok := s.D.GetOkExists("interval_in_seconds"); ok {
		tmp := intervalInSeconds.(int)
		request.IntervalInSeconds = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	tmp := s.D.Id()
	request.MonitorId = &tmp

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

	response, err := s.Client.UpdatePingMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PingMonitor
	return nil
}

func (s *HealthChecksPingMonitorResourceCrud) Delete() error {
	request := oci_health_checks.DeletePingMonitorRequest{}

	tmp := s.D.Id()
	request.MonitorId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	_, err := s.Client.DeletePingMonitor(context.Background(), request)
	return err
}

func (s *HealthChecksPingMonitorResourceCrud) SetData() error {
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

func (s *HealthChecksPingMonitorResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_health_checks.ChangePingMonitorCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MonitorId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "health_checks")

	_, err := s.Client.ChangePingMonitorCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
