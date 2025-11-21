// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatascienceModelGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelGroups,
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
			"model_group_version_history_id": {
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
			"model_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelGroupResource()),
			},
		},
	}
}

func readDatascienceModelGroups(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelGroupsResponse
}

func (s *DatascienceModelGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelGroupsDataSourceCrud) Get() error {
	request := oci_datascience.ListModelGroupsRequest{}

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

	if modelGroupVersionHistoryId, ok := s.D.GetOkExists("model_group_version_history_id"); ok {
		tmp := modelGroupVersionHistoryId.(string)
		request.ModelGroupVersionHistoryId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_datascience.ListModelGroupsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelGroupsDataSource-", DatascienceModelGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelGroup := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			modelGroup["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelGroup["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			modelGroup["display_name"] = *r.DisplayName
		}

		modelGroup["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelGroup["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			modelGroup["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ModelGroupDetails != nil {
			modelGroupDetailsArray := []interface{}{}
			if modelGroupDetailsMap := ModelGroupDetailsToMap(&r.ModelGroupDetails); modelGroupDetailsMap != nil {
				modelGroupDetailsArray = append(modelGroupDetailsArray, modelGroupDetailsMap)
			}
			modelGroup["model_group_details"] = modelGroupDetailsArray
		} else {
			modelGroup["model_group_details"] = nil
		}

		if r.ModelGroupVersionHistoryId != nil {
			modelGroup["model_group_version_history_id"] = *r.ModelGroupVersionHistoryId
		}

		if r.ModelGroupVersionHistoryName != nil {
			modelGroup["model_group_version_history_name"] = *r.ModelGroupVersionHistoryName
		}

		if r.ProjectId != nil {
			modelGroup["project_id"] = *r.ProjectId
		}

		modelGroup["state"] = r.LifecycleState

		if r.SystemTags != nil {
			modelGroup["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			modelGroup["time_created"] = r.TimeCreated.String()
		}

		if r.TimeUpdated != nil {
			modelGroup["time_updated"] = r.TimeUpdated.String()
		}

		if r.VersionId != nil {
			modelGroup["version_id"] = strconv.FormatInt(*r.VersionId, 10)
		}

		if r.VersionLabel != nil {
			modelGroup["version_label"] = *r.VersionLabel
		}

		resources = append(resources, modelGroup)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelGroupsDataSource().Schema["model_groups"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_groups", resources); err != nil {
		return err
	}

	return nil
}
