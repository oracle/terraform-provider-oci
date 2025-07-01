// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_traces

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_apm_traces "github.com/oracle/oci-go-sdk/v65/apmtraces"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmTracesScheduledQueryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["apm_domain_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["scheduled_query_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ApmTracesScheduledQueryResource(), fieldMap, readSingularApmTracesScheduledQuery)
}

func readSingularApmTracesScheduledQuery(d *schema.ResourceData, m interface{}) error {
	sync := &ApmTracesScheduledQueryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledQueryClient()

	return tfresource.ReadResource(sync)
}

type ApmTracesScheduledQueryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_apm_traces.ScheduledQueryClient
	Res    *oci_apm_traces.GetScheduledQueryResponse
}

func (s *ApmTracesScheduledQueryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ApmTracesScheduledQueryDataSourceCrud) Get() error {
	request := oci_apm_traces.GetScheduledQueryRequest{}

	if scheduledQueryId, ok := s.D.GetOkExists("scheduled_query_id"); ok {
		tmp := scheduledQueryId.(string)
		apmDomainId, scheduledQueryId, err := parseScheduledQueryCompositeId(tmp)
		if err == nil {
			request.ScheduledQueryId = &scheduledQueryId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", tmp)
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "apm_traces")

	response, err := s.Client.GetScheduledQuery(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ApmTracesScheduledQueryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ScheduledQueryDescription != nil {
		s.D.Set("scheduled_query_description", *s.Res.ScheduledQueryDescription)
	}

	if s.Res.ScheduledQueryInstances != nil {
		s.D.Set("scheduled_query_instances", *s.Res.ScheduledQueryInstances)
	}

	if s.Res.ScheduledQueryMaximumRuntimeInSeconds != nil {
		s.D.Set("scheduled_query_maximum_runtime_in_seconds", strconv.FormatInt(*s.Res.ScheduledQueryMaximumRuntimeInSeconds, 10))
	}

	if s.Res.ScheduledQueryName != nil {
		s.D.Set("scheduled_query_name", *s.Res.ScheduledQueryName)
	}

	if s.Res.ScheduledQueryNextRunInMs != nil {
		s.D.Set("scheduled_query_next_run_in_ms", strconv.FormatInt(*s.Res.ScheduledQueryNextRunInMs, 10))
	}

	if s.Res.ScheduledQueryProcessingConfiguration != nil {
		s.D.Set("scheduled_query_processing_configuration", []interface{}{ScheduledQueryProcessingConfigToMap(s.Res.ScheduledQueryProcessingConfiguration)})
	} else {
		s.D.Set("scheduled_query_processing_configuration", nil)
	}

	s.D.Set("scheduled_query_processing_sub_type", s.Res.ScheduledQueryProcessingSubType)

	s.D.Set("scheduled_query_processing_type", s.Res.ScheduledQueryProcessingType)

	s.D.Set("scheduled_query_retention_criteria", s.Res.ScheduledQueryRetentionCriteria)

	if s.Res.ScheduledQueryRetentionPeriodInMs != nil {
		s.D.Set("scheduled_query_retention_period_in_ms", strconv.FormatInt(*s.Res.ScheduledQueryRetentionPeriodInMs, 10))
	}

	if s.Res.ScheduledQuerySchedule != nil {
		s.D.Set("scheduled_query_schedule", *s.Res.ScheduledQuerySchedule)
	}

	if s.Res.ScheduledQueryText != nil {
		s.D.Set("scheduled_query_text", *s.Res.ScheduledQueryText)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	return nil
}
