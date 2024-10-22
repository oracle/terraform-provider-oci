// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatascienceModelsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModels,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"model_version_set_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"model_version_set_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version_label": {
				Type:     schema.TypeString,
				Optional: true,
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
			"models": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelResource()),
			},
		},
	}
}

func readDatascienceModels(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelsResponse
}

func (s *DatascienceModelsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelsDataSourceCrud) Get() error {
	request := oci_datascience.ListModelsRequest{}

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

	if model_version_set_name, ok := s.D.GetOkExists("model_version_set_name"); ok {
		tmp := model_version_set_name.(string)
		request.ModelVersionSetName = &tmp
	}

	if version_label, ok := s.D.GetOkExists("version_label"); ok {
		tmp := version_label.(string)
		request.VersionLabel = &tmp
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
		request.LifecycleState = oci_datascience.ListModelsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModels(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModels(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelsDataSource-", DatascienceModelsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		model := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CreatedBy != nil {
			model["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			model["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			model["display_name"] = *r.DisplayName
		}

		model["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			model["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			model["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ProjectId != nil {
			model["project_id"] = *r.ProjectId
		}

		if r.ModelVersionSetName != nil {
			model["model_version_set_name"] = *r.ModelVersionSetName
		}

		if r.ModelVersionSetId != nil {
			model["model_version_set_id"] = *r.ModelVersionSetId
		}

		if r.VersionLabel != nil {
			model["version_label"] = *r.VersionLabel
		}

		model["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			model["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, model)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelsDataSource().Schema["models"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("models", resources); err != nil {
		return err
	}

	return nil
}
