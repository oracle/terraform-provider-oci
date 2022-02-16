// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v58/artifacts"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ArtifactsGenericArtifactDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["artifact_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ArtifactsGenericArtifactResource(), fieldMap, readSingularArtifactsGenericArtifact)
}

func readSingularArtifactsGenericArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsGenericArtifactDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsGenericArtifactDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetGenericArtifactResponse
}

func (s *ArtifactsGenericArtifactDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsGenericArtifactDataSourceCrud) Get() error {
	request := oci_artifacts.GetGenericArtifactRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.GetGenericArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsGenericArtifactDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ArtifactPath != nil {
		s.D.Set("artifact_path", *s.Res.ArtifactPath)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RepositoryId != nil {
		s.D.Set("repository_id", *s.Res.RepositoryId)
	}

	if s.Res.Sha256 != nil {
		s.D.Set("sha256", *s.Res.Sha256)
	}

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
