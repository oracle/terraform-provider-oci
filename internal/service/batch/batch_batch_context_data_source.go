// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchContextDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["batch_context_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BatchBatchContextResource(), fieldMap, readSingularBatchBatchContextWithContext)
}

func readSingularBatchBatchContextWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchContextDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.GetBatchContextResponse
}

func (s *BatchBatchContextDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchContextDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchContextRequest{}

	if batchContextId, ok := s.D.GetOkExists("batch_context_id"); ok {
		tmp := batchContextId.(string)
		request.BatchContextId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.GetBatchContext(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BatchBatchContextDataSourceCrud) SetData() error {
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

	s.D.Set("entitlements", s.Res.Entitlements)

	fleets := []interface{}{}
	for _, item := range s.Res.Fleets {
		fleets = append(fleets, FleetToMap(item))
	}
	s.D.Set("fleets", fleets)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	jobPriorityConfigurations := []interface{}{}
	for _, item := range s.Res.JobPriorityConfigurations {
		jobPriorityConfigurations = append(jobPriorityConfigurations, JobPriorityConfigurationToMap(item))
	}
	s.D.Set("job_priority_configurations", jobPriorityConfigurations)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LoggingConfiguration != nil {
		loggingConfigurationArray := []interface{}{}
		if loggingConfigurationMap := LoggingConfigurationToMap(&s.Res.LoggingConfiguration); loggingConfigurationMap != nil {
			loggingConfigurationArray = append(loggingConfigurationArray, loggingConfigurationMap)
		}
		s.D.Set("logging_configuration", loggingConfigurationArray)
	} else {
		s.D.Set("logging_configuration", nil)
	}

	if s.Res.Network != nil {
		s.D.Set("network", []interface{}{NetworkToMap(s.Res.Network, true)})
	} else {
		s.D.Set("network", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
