// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
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
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"digest": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func ContainerImageLayerToMap(obj oci_artifacts.ContainerImageLayer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Digest != nil {
		result["digest"] = string(*obj.Digest)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = strconv.FormatInt(*obj.SizeInBytes, 10)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func ContainerImageSummaryToMap(obj oci_artifacts.ContainerImageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Digest != nil {
		result["digest"] = string(*obj.Digest)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.RepositoryId != nil {
		result["repository_id"] = string(*obj.RepositoryId)
	}

	if obj.RepositoryName != nil {
		result["repository_name"] = string(*obj.RepositoryName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func ContainerVersionToMap(obj oci_artifacts.ContainerVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}
