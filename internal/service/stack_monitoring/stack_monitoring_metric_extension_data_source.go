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

func StackMonitoringMetricExtensionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["metric_extension_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(StackMonitoringMetricExtensionResource(), fieldMap, readSingularStackMonitoringMetricExtension)
}

func readSingularStackMonitoringMetricExtension(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMetricExtensionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

type StackMonitoringMetricExtensionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_stack_monitoring.StackMonitoringClient
	Res    *oci_stack_monitoring.GetMetricExtensionResponse
}

func (s *StackMonitoringMetricExtensionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *StackMonitoringMetricExtensionDataSourceCrud) Get() error {
	request := oci_stack_monitoring.GetMetricExtensionRequest{}

	if metricExtensionId, ok := s.D.GetOkExists("metric_extension_id"); ok {
		tmp := metricExtensionId.(string)
		request.MetricExtensionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "stack_monitoring")

	response, err := s.Client.GetMetricExtension(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *StackMonitoringMetricExtensionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CollectionMethod != nil {
		s.D.Set("collection_method", *s.Res.CollectionMethod)
	}

	if s.Res.CollectionRecurrences != nil {
		s.D.Set("collection_recurrences", *s.Res.CollectionRecurrences)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	enabledOnResources := []interface{}{}
	for _, item := range s.Res.EnabledOnResources {
		enabledOnResources = append(enabledOnResources, EnabledResourceDetailsToMap(item))
	}
	s.D.Set("enabled_on_resources", enabledOnResources)

	if s.Res.EnabledOnResourcesCount != nil {
		s.D.Set("enabled_on_resources_count", *s.Res.EnabledOnResourcesCount)
	}

	if s.Res.LastUpdatedBy != nil {
		s.D.Set("last_updated_by", *s.Res.LastUpdatedBy)
	}

	metricList := []interface{}{}
	for _, item := range s.Res.MetricList {
		metricList = append(metricList, MetricToMap(item))
	}
	s.D.Set("metric_list", metricList)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.QueryProperties != nil {
		queryPropertiesArray := []interface{}{}
		if queryPropertiesMap := MetricExtensionQueryPropertiesToMap(&s.Res.QueryProperties); queryPropertiesMap != nil {
			queryPropertiesArray = append(queryPropertiesArray, queryPropertiesMap)
		}
		s.D.Set("query_properties", queryPropertiesArray)
	} else {
		s.D.Set("query_properties", nil)
	}

	if s.Res.ResourceType != nil {
		s.D.Set("resource_type", *s.Res.ResourceType)
	}

	if s.Res.ResourceUri != nil {
		s.D.Set("resource_uri", *s.Res.ResourceUri)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("status", s.Res.Status)

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
