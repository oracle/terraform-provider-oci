// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"log"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v56/apmsynthetics"
)

func ApmSyntheticsMonitorDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["monitor_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmSyntheticsMonitorResource(), fieldMap, readSingularApmSyntheticsMonitor)
}

func readSingularApmSyntheticsMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

type ApmSyntheticsMonitorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_synthetics.ApmSyntheticClient
	Res    *oci_apm_synthetics.GetMonitorResponse
}

func (s *ApmSyntheticsMonitorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmSyntheticsMonitorDataSourceCrud) Get() error {
	request := oci_apm_synthetics.GetMonitorRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if monitorId, ok := s.D.GetOkExists("monitor_id"); ok {
		tmp := monitorId.(string)

		monitorId, apmDomainId, err := parseMonitorCompositeId(tmp)
		if err == nil {
			request.MonitorId = &monitorId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_synthetics")

	response, err := s.Client.GetMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmSyntheticsMonitorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GetMonitorCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string)))

	if s.Res.Configuration != nil {
		configurationArray := []interface{}{}
		if configurationMap := MonitorConfigurationToMap(&s.Res.Configuration); configurationMap != nil {
			configurationArray = append(configurationArray, configurationMap)
		}
		s.D.Set("configuration", configurationArray)
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRunOnce != nil {
		s.D.Set("is_run_once", *s.Res.IsRunOnce)
	}

	s.D.Set("monitor_type", s.Res.MonitorType)

	if s.Res.RepeatIntervalInSeconds != nil {
		s.D.Set("repeat_interval_in_seconds", *s.Res.RepeatIntervalInSeconds)
	}

	if s.Res.ScriptId != nil {
		s.D.Set("script_id", GetMonitorCompositeId(*s.Res.ScriptId, s.D.Get("apm_domain_id").(string)))
	}

	if s.Res.ScriptName != nil {
		s.D.Set("script_name", *s.Res.ScriptName)
	}

	scriptParameters := []interface{}{}
	for _, item := range s.Res.ScriptParameters {
		scriptParameters = append(scriptParameters, MonitorScriptParameterInfoToMap(item))
	}
	s.D.Set("script_parameters", scriptParameters)

	s.D.Set("status", s.Res.Status)

	if s.Res.Target != nil {
		s.D.Set("target", *s.Res.Target)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	if s.Res.VantagePointCount != nil {
		s.D.Set("vantage_point_count", *s.Res.VantagePointCount)
	}

	vantagePoints := []interface{}{}
	for _, item := range s.Res.VantagePoints {
		vantagePoints = append(vantagePoints, item.Name)
	}
	s.D.Set("vantage_points", vantagePoints)

	return nil
}
