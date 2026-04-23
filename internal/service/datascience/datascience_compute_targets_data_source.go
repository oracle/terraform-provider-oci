// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceComputeTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDatascienceComputeTargetsWithContext,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
			"compute_targets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceComputeTargetResource()),
			},
		},
	}
}

func readDatascienceComputeTargetsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DatascienceComputeTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DatascienceComputeTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListComputeTargetsResponse
}

func (s *DatascienceComputeTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceComputeTargetsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datascience.ListComputeTargetsRequest{}

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
		request.LifecycleState = oci_datascience.ListComputeTargetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListComputeTargets(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListComputeTargets(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceComputeTargetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceComputeTargetsDataSource-", DatascienceComputeTargetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		computeTarget := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			computeTarget["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			computeTarget["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			computeTarget["display_name"] = *r.DisplayName
		}

		computeTarget["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			computeTarget["id"] = *r.Id
		}

		computeTarget["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			computeTarget["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, computeTarget)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceComputeTargetsDataSource().Schema["compute_targets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("compute_targets", resources); err != nil {
		return err
	}

	return nil
}
