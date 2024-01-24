// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsRepositoryCommitsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryCommits,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"author_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"commit_message": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exclude_ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"file_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ref_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timestamp_greater_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timestamp_less_than_or_equal_to": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_commit_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"author_email": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"author_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"commit_id": {
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
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readDevopsRepositoryCommits(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryCommitsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryCommitsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListCommitsResponse
}

func (s *DevopsRepositoryCommitsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryCommitsDataSourceCrud) Get() error {
	request := oci_devops.ListCommitsRequest{}

	if authorName, ok := s.D.GetOkExists("author_name"); ok {
		tmp := authorName.(string)
		request.AuthorName = &tmp
	}

	if commitMessage, ok := s.D.GetOkExists("commit_message"); ok {
		tmp := commitMessage.(string)
		request.CommitMessage = &tmp
	}

	if excludeRefName, ok := s.D.GetOkExists("exclude_ref_name"); ok {
		tmp := excludeRefName.(string)
		request.ExcludeRefName = &tmp
	}

	if filePath, ok := s.D.GetOkExists("file_path"); ok {
		tmp := filePath.(string)
		request.FilePath = &tmp
	}

	if refName, ok := s.D.GetOkExists("ref_name"); ok {
		tmp := refName.(string)
		request.RefName = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if timestampGreaterThanOrEqualTo, ok := s.D.GetOkExists("timestamp_greater_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timestampGreaterThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimestampGreaterThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	if timestampLessThanOrEqualTo, ok := s.D.GetOkExists("timestamp_less_than_or_equal_to"); ok {
		tmp, err := time.Parse(time.RFC3339, timestampLessThanOrEqualTo.(string))
		if err != nil {
			return err
		}
		request.TimestampLessThanOrEqualTo = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListCommits(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCommits(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryCommitsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryCommitsDataSource-", DevopsRepositoryCommitsDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryCommit := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryCommitSummaryToMap(item))
	}
	repositoryCommit["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryCommitsDataSource().Schema["repository_commit_collection"].Elem.(*schema.Resource).Schema)
		repositoryCommit["items"] = items
	}

	resources = append(resources, repositoryCommit)
	if err := s.D.Set("repository_commit_collection", resources); err != nil {
		return err
	}

	return nil
}
