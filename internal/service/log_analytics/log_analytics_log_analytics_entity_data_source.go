// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v56/loganalytics"
)

func LogAnalyticsLogAnalyticsEntityDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["log_analytics_entity_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LogAnalyticsLogAnalyticsEntityResource(), fieldMap, readSingularLogAnalyticsLogAnalyticsEntity)
}

func readSingularLogAnalyticsLogAnalyticsEntity(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsLogAnalyticsEntityDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsLogAnalyticsEntityDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetLogAnalyticsEntityResponse
}

func (s *LogAnalyticsLogAnalyticsEntityDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsLogAnalyticsEntityDataSourceCrud) Get() error {
	request := oci_log_analytics.GetLogAnalyticsEntityRequest{}

	if logAnalyticsEntityId, ok := s.D.GetOkExists("log_analytics_entity_id"); ok {
		tmp := logAnalyticsEntityId.(string)
		request.LogAnalyticsEntityId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetLogAnalyticsEntity(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsLogAnalyticsEntityDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AreLogsCollected != nil {
		s.D.Set("are_logs_collected", *s.Res.AreLogsCollected)
	}

	if s.Res.CloudResourceId != nil {
		s.D.Set("cloud_resource_id", *s.Res.CloudResourceId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.EntityTypeInternalName != nil {
		s.D.Set("entity_type_internal_name", *s.Res.EntityTypeInternalName)
	}

	if s.Res.EntityTypeName != nil {
		s.D.Set("entity_type_name", *s.Res.EntityTypeName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ManagementAgentCompartmentId != nil {
		s.D.Set("management_agent_compartment_id", *s.Res.ManagementAgentCompartmentId)
	}

	if s.Res.ManagementAgentDisplayName != nil {
		s.D.Set("management_agent_display_name", *s.Res.ManagementAgentDisplayName)
	}

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("properties", s.Res.Properties)

	if s.Res.SourceId != nil {
		s.D.Set("source_id", *s.Res.SourceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimezoneRegion != nil {
		s.D.Set("timezone_region", *s.Res.TimezoneRegion)
	}

	return nil
}
