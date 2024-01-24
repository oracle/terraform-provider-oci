// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package log_analytics

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_log_analytics "github.com/oracle/oci-go-sdk/v65/loganalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LogAnalyticsNamespaceIngestTimeRuleDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["ingest_time_rule_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["namespace"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(LogAnalyticsNamespaceIngestTimeRuleResource(), fieldMap, readSingularLogAnalyticsNamespaceIngestTimeRule)
}

func readSingularLogAnalyticsNamespaceIngestTimeRule(d *schema.ResourceData, m interface{}) error {
	sync := &LogAnalyticsNamespaceIngestTimeRuleDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LogAnalyticsClient()

	return tfresource.ReadResource(sync)
}

type LogAnalyticsNamespaceIngestTimeRuleDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_log_analytics.LogAnalyticsClient
	Res    *oci_log_analytics.GetIngestTimeRuleResponse
}

func (s *LogAnalyticsNamespaceIngestTimeRuleDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LogAnalyticsNamespaceIngestTimeRuleDataSourceCrud) Get() error {
	request := oci_log_analytics.GetIngestTimeRuleRequest{}

	if ingestTimeRuleId, ok := s.D.GetOkExists("ingest_time_rule_id"); ok {
		tmp := ingestTimeRuleId.(string)
		request.IngestTimeRuleId = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "log_analytics")

	response, err := s.Client.GetIngestTimeRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LogAnalyticsNamespaceIngestTimeRuleDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	actions := []interface{}{}
	for _, item := range s.Res.Actions {
		actions = append(actions, IngestTimeRuleActionToMap(item, true))
	}
	s.D.Set("actions", actions)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Conditions != nil {
		conditionsArray := []interface{}{}
		if conditionsMap := IngestTimeRuleConditionToMap(&s.Res.Conditions, true); conditionsMap != nil {
			conditionsArray = append(conditionsArray, conditionsMap)
		}
		s.D.Set("conditions", conditionsArray)
	} else {
		s.D.Set("conditions", nil)
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

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
