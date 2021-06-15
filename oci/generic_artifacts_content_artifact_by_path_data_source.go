// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_generic_artifacts_content "github.com/oracle/oci-go-sdk/v42/genericartifactscontent"
)

func init() {
	RegisterDatasource("oci_generic_artifacts_content_artifact_by_path", GenericArtifactsContentArtifactByPathDataSource())
}

func GenericArtifactsContentArtifactByPathDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["artifact_path"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["repository_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["version"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(GenericArtifactsContentArtifactByPathResource(), fieldMap, readSingularGenericArtifactsContentArtifactByPath)
}

func readSingularGenericArtifactsContentArtifactByPath(d *schema.ResourceData, m interface{}) error {
	sync := &GenericArtifactsContentArtifactByPathDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).genericArtifactsContentClient()

	return ReadResource(sync)
}

type GenericArtifactsContentArtifactByPathDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generic_artifacts_content.GenericArtifactsContentClient
	Res    *oci_generic_artifacts_content.GetGenericArtifactContentByPathResponse
}

func (s *GenericArtifactsContentArtifactByPathDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenericArtifactsContentArtifactByPathDataSourceCrud) Get() error {
	request := oci_generic_artifacts_content.GetGenericArtifactContentByPathRequest{}

	if artifactPath, ok := s.D.GetOkExists("artifact_path"); ok {
		tmp := artifactPath.(string)
		request.ArtifactPath = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "generic_artifacts_content")

	response, err := s.Client.GetGenericArtifactContentByPath(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenericArtifactsContentArtifactByPathDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("GenericArtifactsContentArtifactByPathDataSource-", GenericArtifactsContentArtifactByPathDataSource(), s.D))

	return nil
}
