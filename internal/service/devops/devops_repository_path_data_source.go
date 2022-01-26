// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_devops "github.com/oracle/oci-go-sdk/v56/devops"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DevopsRepositoryPathDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDevopsRepositoryPath,
		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"folder_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"paths_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ref": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"submodule_git_url": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readSingularDevopsRepositoryPath(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryPathDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryPathDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListPathsResponse
}

func (s *DevopsRepositoryPathDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryPathDataSourceCrud) Get() error {
	request := oci_devops.ListPathsRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if folderPath, ok := s.D.GetOkExists("folder_path"); ok {
		tmp := folderPath.(string)
		request.FolderPath = &tmp
	}

	if pathsInSubtree, ok := s.D.GetOkExists("paths_in_subtree"); ok {
		tmp := pathsInSubtree.(bool)
		request.PathsInSubtree = &tmp
	}

	if ref, ok := s.D.GetOkExists("ref"); ok {
		tmp := ref.(string)
		request.Ref = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "devops")

	response, err := s.Client.ListPaths(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DevopsRepositoryPathDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryPathDataSource-", DevopsRepositoryPathDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryPathSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}

func RepositoryPathSummaryToMap(obj oci_devops.RepositoryPathSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Path != nil {
		result["path"] = string(*obj.Path)
	}

	if obj.Sha != nil {
		result["sha"] = string(*obj.Sha)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = strconv.FormatInt(*obj.SizeInBytes, 10)
	}

	if obj.SubmoduleGitUrl != nil {
		result["submodule_git_url"] = string(*obj.SubmoduleGitUrl)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
