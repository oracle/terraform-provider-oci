// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_datascience "github.com/oracle/oci-go-sdk/v39/datascience"
)

func init() {
	RegisterDatasource("oci_datascience_model_deployment", DatascienceModelDeploymentDataSource())
}

func DatascienceModelDeploymentDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["model_deployment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatascienceModelDeploymentResource(), fieldMap, readSingularDatascienceModelDeployment)
}

func readSingularDatascienceModelDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceModelDeploymentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).dataScienceClient()

	return ReadResource(sync)
}

type DatascienceModelDeploymentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datascience.DataScienceClient
	Res    *oci_datascience.GetModelDeploymentResponse
}

func (s *DatascienceModelDeploymentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatascienceModelDeploymentDataSourceCrud) Get() error {
	request := oci_datascience.GetModelDeploymentRequest{}

	if modelDeploymentId, ok := s.D.GetOkExists("model_deployment_id"); ok {
		tmp := modelDeploymentId.(string)
		request.ModelDeploymentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "datascience")

	response, err := s.Client.GetModelDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatascienceModelDeploymentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CategoryLogDetails != nil {
		s.D.Set("category_log_details", []interface{}{CategoryLogDetailsToMap(s.Res.CategoryLogDetails)})
	} else {
		s.D.Set("category_log_details", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelDeploymentConfigurationDetails != nil {
		modelDeploymentConfigurationDetailsArray := []interface{}{}
		if modelDeploymentConfigurationDetailsMap := ModelDeploymentConfigurationDetailsToMap(&s.Res.ModelDeploymentConfigurationDetails); modelDeploymentConfigurationDetailsMap != nil {
			modelDeploymentConfigurationDetailsArray = append(modelDeploymentConfigurationDetailsArray, modelDeploymentConfigurationDetailsMap)
		}
		s.D.Set("model_deployment_configuration_details", modelDeploymentConfigurationDetailsArray)
	} else {
		s.D.Set("model_deployment_configuration_details", nil)
	}

	if s.Res.ModelDeploymentUrl != nil {
		s.D.Set("model_deployment_url", *s.Res.ModelDeploymentUrl)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
