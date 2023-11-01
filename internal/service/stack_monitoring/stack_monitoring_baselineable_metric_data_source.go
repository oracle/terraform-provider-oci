// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringBaselineableMetricDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["baselineable_metric_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringBaselineableMetricResource(), fieldMap, readSingularStackMonitoringBaselineableMetric)
}

func readSingularStackMonitoringBaselineableMetric(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringBaselineableMetricDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringBaselineableMetricDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetBaselineableMetricResponse
}

func (s *StackMonitoringBaselineableMetricDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringBaselineableMetricDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetBaselineableMetricRequest{}

	if baselineableMetricId, ok := s.D.GetOkExists("baselineable_metric_id"); ok {
		tmp := baselineableMetricId.(string)
		request.BaselineableMetricId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetBaselineableMetric(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringBaselineableMetricDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Column != nil {
		s.D.Set("column", *s.Res.Column)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOutOfBox != nil {
		s.D.Set("is_out_of_box", *s.Res.IsOutOfBox)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", *s.Res.LastUpdatedBy)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ResourceGroup != nil {
		s.D.Set("resource_group", *s.Res.ResourceGroup)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	return nil
}
