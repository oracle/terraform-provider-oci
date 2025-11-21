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

func BatchBatchJobPoolsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBatchBatchJobPoolsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"batch_context_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"batch_job_pool_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(BatchBatchJobPoolResource()),
						},
					},
				},
			},
		},
	}
}

func readBatchBatchJobPoolsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchJobPoolsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchJobPoolsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.ListBatchJobPoolsResponse
}

func (s *BatchBatchJobPoolsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchJobPoolsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.ListBatchJobPoolsRequest{}

	if batchContextId, ok := s.D.GetOkExists("batch_context_id"); ok {
		tmp := batchContextId.(string)
		request.BatchContextId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_batch.BatchJobPoolLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.ListBatchJobPools(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBatchJobPools(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BatchBatchJobPoolsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BatchBatchJobPoolsDataSource-", BatchBatchJobPoolsDataSource(), s.D))
	resources := []map[string]interface{}{}
	batchJobPool := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BatchJobPoolSummaryToMap(item))
	}
	batchJobPool["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BatchBatchJobPoolsDataSource().Schema["batch_job_pool_collection"].Elem.(*schema.Resource).Schema)
		batchJobPool["items"] = items
	}

	resources = append(resources, batchJobPool)
	if err := s.D.Set("batch_job_pool_collection", resources); err != nil {
		return err
	}

	return nil
}
