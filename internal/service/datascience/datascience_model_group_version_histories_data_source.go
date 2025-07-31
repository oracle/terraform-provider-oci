// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupVersionHistoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelGroupVersionHistories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"created_by": {
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
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_group_version_histories": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelGroupVersionHistoryResource()),
			},
		},
	}
}

func readDatascienceModelGroupVersionHistories(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupVersionHistoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelGroupVersionHistoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelGroupVersionHistoriesResponse
}

func (s *DatascienceModelGroupVersionHistoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelGroupVersionHistoriesDataSourceCrud) Get() error {
	request := oci_datascience.ListModelGroupVersionHistoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if createdBy, ok := s.D.GetOkExists("created_by"); ok {
		tmp := createdBy.(string)
		request.CreatedBy = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListModelGroupVersionHistoriesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelGroupVersionHistories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelGroupVersionHistories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelGroupVersionHistoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelGroupVersionHistoriesDataSource-", DatascienceModelGroupVersionHistoriesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelGroupVersionHistory := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			modelGroupVersionHistory["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelGroupVersionHistory["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			modelGroupVersionHistory["display_name"] = *r.DisplayName
		}

		modelGroupVersionHistory["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelGroupVersionHistory["id"] = *r.Id
		}

		if r.LatestModelGroupId != nil {
			modelGroupVersionHistory["latest_model_group_id"] = *r.LatestModelGroupId
		}

		if r.LifecycleDetails != nil {
			modelGroupVersionHistory["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ProjectId != nil {
			modelGroupVersionHistory["project_id"] = *r.ProjectId
		}

		modelGroupVersionHistory["state"] = r.LifecycleState

		if r.SystemTags != nil {
			modelGroupVersionHistory["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			modelGroupVersionHistory["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			modelGroupVersionHistory["time_updated"] = r.TimeUpdated.String()
		}

		resources = append(resources, modelGroupVersionHistory)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelGroupVersionHistoriesDataSource().Schema["model_group_version_histories"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_group_version_histories", resources); err != nil {
		return err
	}

	return nil
}
