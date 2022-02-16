// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v58/datascience"
)

func DatascienceModelDeploymentsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatascienceModelDeployments,
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
			"model_deployments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatascienceModelDeploymentResource()),
			},
		},
	}
}

func readDatascienceModelDeployments(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

type DatascienceModelDeploymentsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.ListModelDeploymentsResponse
}

func (s *DatascienceModelDeploymentsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelDeploymentsDataSourceCrud) Get() error {
	request := oci_datascience.ListModelDeploymentsRequest{}

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
		request.LifecycleState = oci_datascience.ListModelDeploymentsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datascience")

	response, err := s.Client.ListModelDeployments(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListModelDeployments(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatascienceModelDeploymentsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatascienceModelDeploymentsDataSource-", DatascienceModelDeploymentsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		modelDeployment := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CategoryLogDetails != nil {
			modelDeployment["category_log_details"] = []interface{}{CategoryLogDetailsToMap(r.CategoryLogDetails)}
		} else {
			modelDeployment["category_log_details"] = nil
		}

		if r.CreatedBy != nil {
			modelDeployment["created_by"] = *r.CreatedBy
		}

		if r.DefinedTags != nil {
			modelDeployment["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			modelDeployment["description"] = *r.Description
		}

		if r.DisplayName != nil {
			modelDeployment["display_name"] = *r.DisplayName
		}

		modelDeployment["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			modelDeployment["id"] = *r.Id
		}

		if r.ModelDeploymentConfigurationDetails != nil {
			modelDeploymentConfigurationDetailsArray := []interface{}{}
			if modelDeploymentConfigurationDetailsMap := ModelDeploymentConfigurationDetailsToMap(&r.ModelDeploymentConfigurationDetails); modelDeploymentConfigurationDetailsMap != nil {
				modelDeploymentConfigurationDetailsArray = append(modelDeploymentConfigurationDetailsArray, modelDeploymentConfigurationDetailsMap)
			}
			modelDeployment["model_deployment_configuration_details"] = modelDeploymentConfigurationDetailsArray
		} else {
			modelDeployment["model_deployment_configuration_details"] = nil
		}

		if r.ModelDeploymentUrl != nil {
			modelDeployment["model_deployment_url"] = *r.ModelDeploymentUrl
		}

		if r.ProjectId != nil {
			modelDeployment["project_id"] = *r.ProjectId
		}

		modelDeployment["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			modelDeployment["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, modelDeployment)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatascienceModelDeploymentsDataSource().Schema["model_deployments"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("model_deployments", resources); err != nil {
		return err
	}

	return nil
}
