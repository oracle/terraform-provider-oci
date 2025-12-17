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

func BatchBatchTaskProfileDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["batch_task_profile_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BatchBatchTaskProfileResource(), fieldMap, readSingularBatchBatchTaskProfileWithContext)
}

func readSingularBatchBatchTaskProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfileDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchTaskProfileDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.GetBatchTaskProfileResponse
}

func (s *BatchBatchTaskProfileDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchTaskProfileDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchTaskProfileRequest{}

	if batchTaskProfileId, ok := s.D.GetOkExists("batch_task_profile_id"); ok {
		tmp := batchTaskProfileId.(string)
		request.BatchTaskProfileId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.GetBatchTaskProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BatchBatchTaskProfileDataSourceCrud) SetData() error {
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

	if s.Res.MinMemoryInGBs != nil {
		s.D.Set("min_memory_in_gbs", *s.Res.MinMemoryInGBs)
	}

	if s.Res.MinOcpus != nil {
		s.D.Set("min_ocpus", *s.Res.MinOcpus)
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
