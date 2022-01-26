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

func DevopsRepositoryPathsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDevopsRepositoryPaths,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"repository_path_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
				},
			},
		},
	}
}

func readDevopsRepositoryPaths(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryPathsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

type DevopsRepositoryPathsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_devops.DevopsClient
	Res    *oci_devops.ListPathsResponse
}

func (s *DevopsRepositoryPathsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DevopsRepositoryPathsDataSourceCrud) Get() error {
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
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPaths(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DevopsRepositoryPathsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DevopsRepositoryPathsDataSource-", DevopsRepositoryPathsDataSource(), s.D))
	resources := []map[string]interface{}{}
	repositoryPath := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, RepositoryPathSummaryToMap(item))
	}
	repositoryPath["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DevopsRepositoryPathsDataSource().Schema["repository_path_collection"].Elem.(*schema.Resource).Schema)
		repositoryPath["items"] = items
	}

	resources = append(resources, repositoryPath)
	if err := s.D.Set("repository_path_collection", resources); err != nil {
		return err
	}

	return nil
}
