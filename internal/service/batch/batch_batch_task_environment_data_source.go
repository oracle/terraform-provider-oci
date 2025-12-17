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

func BatchBatchTaskEnvironmentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["batch_task_environment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BatchBatchTaskEnvironmentResource(), fieldMap, readSingularBatchBatchTaskEnvironmentWithContext)
}

func readSingularBatchBatchTaskEnvironmentWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskEnvironmentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchTaskEnvironmentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.GetBatchTaskEnvironmentResponse
}

func (s *BatchBatchTaskEnvironmentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchTaskEnvironmentDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchTaskEnvironmentRequest{}

	if batchTaskEnvironmentId, ok := s.D.GetOkExists("batch_task_environment_id"); ok {
		tmp := batchTaskEnvironmentId.(string)
		request.BatchTaskEnvironmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.GetBatchTaskEnvironment(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BatchBatchTaskEnvironmentDataSourceCrud) SetData() error {
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

	if s.Res.ImageUrl != nil {
		s.D.Set("image_url", *s.Res.ImageUrl)
	}

	if s.Res.SecurityContext != nil {
		s.D.Set("security_context", []interface{}{SecurityContextToMap(s.Res.SecurityContext)})
	} else {
		s.D.Set("security_context", nil)
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

	volumes := []interface{}{}
	for _, item := range s.Res.Volumes {
		volumes = append(volumes, BatchTaskEnvironmentVolumeToMap(item))
	}
	s.D.Set("volumes", volumes)

	if s.Res.WorkingDirectory != nil {
		s.D.Set("working_directory", *s.Res.WorkingDirectory)
	}

	return nil
}
