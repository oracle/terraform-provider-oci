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

func ArtifactsContainerImageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularArtifactsContainerImage,
		Schema: map[string]*schema.Schema{
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func readSingularArtifactsContainerImage(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsContainerImageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetContainerImageResponse
}

func (s *ArtifactsContainerImageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerImageDataSourceCrud) Get() error {
	request := oci_artifacts.GetContainerImageRequest{}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.GetContainerImage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsContainerImageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.Digest != nil {
		s.D.Set("digest", *s.Res.Digest)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	layers := []interface{}{}
	for _, item := range s.Res.Layers {
		layers = append(layers, ContainerImageLayerToMap(item))
	}
	s.D.Set("layers", layers)

	if s.Res.LayersSizeInBytes != nil {
		s.D.Set("layers_size_in_bytes", strconv.FormatInt(*s.Res.LayersSizeInBytes, 10))
	}

	if s.Res.ManifestSizeInBytes != nil {
		s.D.Set("manifest_size_in_bytes", *s.Res.ManifestSizeInBytes)
	}

	if s.Res.PullCount != nil {
		s.D.Set("pull_count", strconv.FormatInt(*s.Res.PullCount, 10))
	}

	if s.Res.RepositoryId != nil {
		s.D.Set("repository_id", *s.Res.RepositoryId)
	}

	if s.Res.RepositoryName != nil {
		s.D.Set("repository_name", *s.Res.RepositoryName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastPulled != nil {
		s.D.Set("time_last_pulled", s.Res.TimeLastPulled.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	versions := []interface{}{}
	for _, item := range s.Res.Versions {
		versions = append(versions, ContainerVersionToMap(item))
	}
	s.D.Set("versions", versions)

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

	if obj.Digest != nil {
		result["digest"] = string(*obj.Digest)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

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
