// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ArtifactsContainerRepositoriesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readArtifactsContainerRepositories,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_public": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_repository_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"image_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ArtifactsContainerRepositoryResource()),
						},

						"layer_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"layers_size_in_bytes": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"remaining_items_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"repository_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readArtifactsContainerRepositories(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoriesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsContainerRepositoriesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.ListContainerRepositoriesResponse
}

func (s *ArtifactsContainerRepositoriesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerRepositoriesDataSourceCrud) Get() error {
	request := oci_artifacts.ListContainerRepositoriesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isPublic, ok := s.D.GetOkExists("is_public"); ok {
		tmp := isPublic.(bool)
		request.IsPublic = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := state.(string)
		request.LifecycleState = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.ListContainerRepositories(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerRepositories(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ArtifactsContainerRepositoriesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ArtifactsContainerRepositoriesDataSource-", ArtifactsContainerRepositoriesDataSource(), s.D))
	resources := []map[string]interface{}{}
	containerRepository := map[string]interface{}{}

	if s.Res.ImageCount != nil {
		containerRepository["image_count"] = *s.Res.ImageCount
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerRepositorySummaryToMap(item))
	}
	containerRepository["items"] = items

	if s.Res.LayerCount != nil {
		containerRepository["layer_count"] = *s.Res.LayerCount
	}

	if s.Res.LayersSizeInBytes != nil {
		containerRepository["layers_size_in_bytes"] = strconv.FormatInt(*s.Res.LayersSizeInBytes, 10)
	}

	if s.Res.RemainingItemsCount != nil {
		containerRepository["remaining_items_count"] = *s.Res.RemainingItemsCount
	}

	if s.Res.RepositoryCount != nil {
		containerRepository["repository_count"] = *s.Res.RepositoryCount
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ArtifactsContainerRepositoriesDataSource().Schema["container_repository_collection"].Elem.(*schema.Resource).Schema)
		containerRepository["items"] = items
	}

	resources = append(resources, containerRepository)
	if err := s.D.Set("container_repository_collection", resources); err != nil {
		return err
	}

	return nil
}
