// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v40/artifacts"
)

func init() {
	RegisterDatasource("oci_artifacts_container_configuration", ArtifactsContainerConfigurationDataSource())
}

func ArtifactsContainerConfigurationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(ArtifactsContainerConfigurationResource(), fieldMap, readSingularArtifactsContainerConfiguration)
}

func readSingularArtifactsContainerConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerConfigurationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).artifactsClient()

	return ReadResource(sync)
}

type ArtifactsContainerConfigurationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetContainerConfigurationResponse
}

func (s *ArtifactsContainerConfigurationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerConfigurationDataSourceCrud) Get() error {
	request := oci_artifacts.GetContainerConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "artifacts")

	response, err := s.Client.GetContainerConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsContainerConfigurationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(s.D.Get("compartment_id").(string))

	if s.Res.IsRepositoryCreatedOnFirstPush != nil {
		s.D.Set("is_repository_created_on_first_push", *s.Res.IsRepositoryCreatedOnFirstPush)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	return nil
}
