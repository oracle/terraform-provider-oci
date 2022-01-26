// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryCommitDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryCommit,
		Schema: map[string]*schema.Schema{
			"commit_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"author_email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"author_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"committer_email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"committer_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_commit_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tree_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDevopsRepositoryCommit(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryCommitDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryCommitDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.GetCommitResponse
}

func (s *DevopsRepositoryCommitDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryCommitDataSourceCrud) Get() error {
	request := oci_devops.GetCommitRequest{}

	if commitId, ok := s.D.GetOkExists("commit_id"); ok {
		tmp := commitId.(string)
		request.CommitId = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.GetCommit(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryCommitDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryCommitDataSource-", DevopsRepositoryCommitDataSource(), s.D))

	if s.Res.AuthorEmail != nil {
		s.D.Set("author_email", *s.Res.AuthorEmail)
	}

	if s.Res.AuthorName != nil {
		s.D.Set("author_name", *s.Res.AuthorName)
	}

	if s.Res.CommitMessage != nil {
		s.D.Set("commit_message", *s.Res.CommitMessage)
	}

	if s.Res.CommitterEmail != nil {
		s.D.Set("committer_email", *s.Res.CommitterEmail)
	}

	if s.Res.CommitterName != nil {
		s.D.Set("committer_name", *s.Res.CommitterName)
	}

	s.D.Set("parent_commit_ids", s.Res.ParentCommitIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TreeId != nil {
		s.D.Set("tree_id", *s.Res.TreeId)
	}

	return nil
}

func RepositoryCommitSummaryToMap(obj oci_devops.RepositoryCommitSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthorEmail != nil {
		result["author_email"] = string(*obj.AuthorEmail)
	}

	if obj.AuthorName != nil {
		result["author_name"] = string(*obj.AuthorName)
	}

	if obj.CommitId != nil {
		result["commit_id"] = string(*obj.CommitId)
	}

	if obj.CommitMessage != nil {
		result["commit_message"] = string(*obj.CommitMessage)
	}

	if obj.CommitterEmail != nil {
		result["committer_email"] = string(*obj.CommitterEmail)
	}

	if obj.CommitterName != nil {
		result["committer_name"] = string(*obj.CommitterName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["parent_commit_ids"] = obj.ParentCommitIds

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TreeId != nil {
		result["tree_id"] = string(*obj.TreeId)
	}

	return result
}
