// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ArtifactsRepositoryDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["repository_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ArtifactsRepositoryResource(), fieldMap, readSingularArtifactsRepository)
}

func readSingularArtifactsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoryDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsRepositoryDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetRepositoryResponse
}

func (s *ArtifactsRepositoryDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsRepositoryDataSourceCrud) Get() error {
	request := oci_artifacts.GetRepositoryRequest{}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.GetRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsRepositoryDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.GetId())
	switch v := (s.Res.Repository).(type) {
	case oci_artifacts.GenericRepository:
		s.D.Set("repository_type", "GENERIC")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.IsImmutable != nil {
			s.D.Set("is_immutable", *v.IsImmutable)
		}

		s.D.Set("state", v.LifecycleState)
	default:
		log.Printf("[WARN] Received 'repository_type' of unknown type %v", s.Res.Repository)
		return nil
	}

	return nil
}
