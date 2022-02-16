// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v58/artifacts"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ArtifactsContainerImagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readArtifactsContainerImages,
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
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_versioned": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"repository_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_image_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"remaining_items_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"created_by": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"digest": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"layers": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"digest": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"size_in_bytes": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_created": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"layers_size_in_bytes": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"manifest_size_in_bytes": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"pull_count": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"repository_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"repository_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_last_pulled": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"versions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"created_by": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"time_created": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version": {
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
				},
			},
		},
	}
}

func readArtifactsContainerImages(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsContainerImagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.ListContainerImagesResponse
}

func (s *ArtifactsContainerImagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerImagesDataSourceCrud) Get() error {
	request := oci_artifacts.ListContainerImagesRequest{}

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

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if isVersioned, ok := s.D.GetOkExists("is_versioned"); ok {
		tmp := isVersioned.(bool)
		request.IsVersioned = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if repositoryName, ok := s.D.GetOkExists("repository_name"); ok {
		tmp := repositoryName.(string)
		request.RepositoryName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		tmp := string(oci_artifacts.ContainerImageLifecycleStateEnum(state.(string)))
		request.LifecycleState = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.ListContainerImages(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerImages(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ArtifactsContainerImagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ArtifactsContainerImagesDataSource-", ArtifactsContainerImagesDataSource(), s.D))
	resources := []map[string]interface{}{}
	containerImage := map[string]interface{}{}

	if s.Res.RemainingItemsCount != nil {
		containerImage["remaining_items_count"] = *s.Res.RemainingItemsCount
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerImageSummaryToMap(item))
	}
	containerImage["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ArtifactsContainerImagesDataSource().Schema["container_image_collection"].Elem.(*schema.Resource).Schema)
		containerImage["items"] = items
	}

	resources = append(resources, containerImage)
	if err := s.D.Set("container_image_collection", resources); err != nil {
		return err
	}

	return nil
}
