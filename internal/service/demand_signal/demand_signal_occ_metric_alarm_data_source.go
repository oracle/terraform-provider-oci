// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package demand_signal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_demand_signal "github.com/oracle/oci-go-sdk/v65/demandsignal"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DemandSignalOccMetricAlarmDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["occ_metric_alarm_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DemandSignalOccMetricAlarmResource(), fieldMap, readSingularDemandSignalOccMetricAlarmWithContext)
}

func readSingularDemandSignalOccMetricAlarmWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DemandSignalOccMetricAlarmDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OccMetricAlarmClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DemandSignalOccMetricAlarmDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_demand_signal.OccMetricAlarmClient
	Res    *oci_demand_signal.GetOccMetricAlarmResponse
}

func (s *DemandSignalOccMetricAlarmDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DemandSignalOccMetricAlarmDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_demand_signal.GetOccMetricAlarmRequest{}

	if occMetricAlarmId, ok := s.D.GetOkExists("occ_metric_alarm_id"); ok {
		tmp := occMetricAlarmId.(string)
		request.OccMetricAlarmId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "demand_signal")

	response, err := s.Client.GetOccMetricAlarm(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DemandSignalOccMetricAlarmDataSourceCrud) SetData() error {
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

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("frequency", s.Res.Frequency)

	if s.Res.IsActive != nil {
		s.D.Set("is_active", *s.Res.IsActive)
	}

	if s.Res.ResourceConfiguration != nil {
		resourceConfigurationArray := []interface{}{}
		if resourceConfigurationMap := BaseResourceConfigurationToMap(&s.Res.ResourceConfiguration); resourceConfigurationMap != nil {
			resourceConfigurationArray = append(resourceConfigurationArray, resourceConfigurationMap)
		}
		s.D.Set("resource_configuration", resourceConfigurationArray)
	} else {
		s.D.Set("resource_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscribers", s.Res.Subscribers)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.Threshold != nil {
		s.D.Set("threshold", *s.Res.Threshold)
	}

	s.D.Set("threshold_type", s.Res.ThresholdType)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
