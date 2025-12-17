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

func BatchBatchTaskProfilesDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readBatchBatchTaskProfilesWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"batch_task_profile_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(BatchBatchTaskProfileResource()),
						},
					},
				},
			},
		},
	}
}

func readBatchBatchTaskProfilesWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchTaskProfilesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type BatchBatchTaskProfilesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_batch.BatchComputingClient
	Res    *oci_batch.ListBatchTaskProfilesResponse
}

func (s *BatchBatchTaskProfilesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BatchBatchTaskProfilesDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.ListBatchTaskProfilesRequest{}

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
		request.LifecycleState = oci_batch.BatchTaskProfileLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "batch")

	response, err := s.Client.ListBatchTaskProfiles(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListBatchTaskProfiles(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *BatchBatchTaskProfilesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("BatchBatchTaskProfilesDataSource-", BatchBatchTaskProfilesDataSource(), s.D))
	resources := []map[string]interface{}{}
	batchTaskProfile := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, BatchTaskProfileSummaryToMap(item))
	}
	batchTaskProfile["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, BatchBatchTaskProfilesDataSource().Schema["batch_task_profile_collection"].Elem.(*schema.Resource).Schema)
		batchTaskProfile["items"] = items
	}

	resources = append(resources, batchTaskProfile)
	if err := s.D.Set("batch_task_profile_collection", resources); err != nil {
		return err
	}

	return nil
}
