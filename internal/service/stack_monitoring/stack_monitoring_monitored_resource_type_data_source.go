// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceTypeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["monitored_resource_type_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMonitoredResourceTypeResource(), fieldMap, readSingularStackMonitoringMonitoredResourceType)
}

func readSingularStackMonitoringMonitoredResourceType(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTypeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMonitoredResourceTypeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMonitoredResourceTypeResponse
}

func (s *StackMonitoringMonitoredResourceTypeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMonitoredResourceTypeDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceTypeRequest{}

	if monitoredResourceTypeId, ok := s.D.GetOkExists("monitored_resource_type_id"); ok {
		tmp := monitoredResourceTypeId.(string)
		request.MonitoredResourceTypeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMonitoredResourceType(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMonitoredResourceTypeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("additional_namespace_map", s.Res.AdditionalNamespaceMap)

	if s.Res.AvailabilityMetricsConfig != nil {
		s.D.Set("availability_metrics_config", []interface{}{AvailabilityMetricsDetailsToMap(s.Res.AvailabilityMetricsConfig)})
	} else {
		s.D.Set("availability_metrics_config", nil)
	}

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

	if s.Res.HandlerConfig != nil {
		s.D.Set("handler_config", []interface{}{AgentExtensionHandlerConfigurationToMap(s.Res.HandlerConfig)})
	} else {
		s.D.Set("handler_config", nil)
	}

	if s.Res.IsSystemDefined != nil {
		s.D.Set("is_system_defined", *s.Res.IsSystemDefined)
	}

	if s.Res.Metadata != nil {
		metadataArray := []interface{}{}
		if metadataMap := ResourceTypeMetadataDetailsToMap(&s.Res.Metadata); metadataMap != nil {
			metadataArray = append(metadataArray, metadataMap)
		}
		s.D.Set("metadata", metadataArray)
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.MetricNamespace != nil {
		s.D.Set("metric_namespace", *s.Res.MetricNamespace)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("resource_category", s.Res.ResourceCategory)

	s.D.Set("source_type", s.Res.SourceType)

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

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
