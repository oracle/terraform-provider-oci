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

func BatchBatchJobPoolDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["batch_job_pool_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(BatchBatchJobPoolResource(), fieldMap, readSingularBatchBatchJobPoolWithContext)
}

func readSingularBatchBatchJobPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchJobPoolDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.GetBatchJobPoolResponse
}

func (s *BatchBatchJobPoolDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchJobPoolDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchJobPoolRequest{}

	if batchJobPoolId, ok := s.D.GetOkExists("batch_job_pool_id"); ok {
		tmp := batchJobPoolId.(string)
		request.BatchJobPoolId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.GetBatchJobPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BatchBatchJobPoolDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BatchContextId != nil {
		if err := s.D.Set("batch_context_id", *s.Res.BatchContextId); err != nil {
			return err
		}
	}

	if s.Res.CompartmentId != nil {
		if err := s.D.Set("compartment_id", *s.Res.CompartmentId); err != nil {
			return err
		}
	}

	if s.Res.DefinedTags != nil {
		if err := s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags)); err != nil {
			return err
		}
	}

	if s.Res.Description != nil {
		if err := s.D.Set("description", *s.Res.Description); err != nil {
			return err
		}
	}

	if s.Res.DisplayName != nil {
		if err := s.D.Set("display_name", *s.Res.DisplayName); err != nil {
			return err
		}
	}

	if err := s.D.Set("freeform_tags", s.Res.FreeformTags); err != nil {
		return err
	}

	if err := s.D.Set("state", s.Res.LifecycleState); err != nil {
		return err
	}

	if s.Res.SystemTags != nil {
		if err := s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags)); err != nil {
			return err
		}
	}

	if s.Res.TimeCreated != nil {
		if err := s.D.Set("time_created", s.Res.TimeCreated.String()); err != nil {
			return err
		}
	}

	if s.Res.TimeUpdated != nil {
		if err := s.D.Set("time_updated", s.Res.TimeUpdated.String()); err != nil {
			return err
		}
	}

	return nil
}
