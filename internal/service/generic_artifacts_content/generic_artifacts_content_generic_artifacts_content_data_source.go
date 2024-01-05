// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generic_artifacts_content

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generic_artifacts_content "github.com/oracle/oci-go-sdk/v65/genericartifactscontent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenericArtifactsContentGenericArtifactsContentDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularGenericArtifactsContentGenericArtifactsContent,
		Schema: map[string]*schema.Schema{
			"artifact_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
		},
	}
}

func readSingularGenericArtifactsContentGenericArtifactsContent(d *schema.ResourceData, m interface{}) error {
	sync := &GenericArtifactsContentGenericArtifactsContentDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenericArtifactsContentClient()

	return tfresource.ReadResource(sync)
}

type GenericArtifactsContentGenericArtifactsContentDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generic_artifacts_content.GenericArtifactsContentClient
	Res    *oci_generic_artifacts_content.GetGenericArtifactContentResponse
}

func (s *GenericArtifactsContentGenericArtifactsContentDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenericArtifactsContentGenericArtifactsContentDataSourceCrud) Get() error {
	request := oci_generic_artifacts_content.GetGenericArtifactContentRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generic_artifacts_content")

	response, err := s.Client.GetGenericArtifactContent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenericArtifactsContentGenericArtifactsContentDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("GenericArtifactsContentGenericArtifactsContentDataSource-", GenericArtifactsContentGenericArtifactsContentDataSource(), s.D))

	return nil
}
