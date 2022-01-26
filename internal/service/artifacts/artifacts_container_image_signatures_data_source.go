// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func ArtifactsContainerImageSignaturesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readArtifactsContainerImageSignatures,
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
			"image_digest": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
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
			"signing_algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"container_image_signature_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(ArtifactsContainerImageSignatureResource()),
						},

						"remaining_items_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readArtifactsContainerImageSignatures(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageSignaturesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsContainerImageSignaturesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.ListContainerImageSignaturesResponse
}

func (s *ArtifactsContainerImageSignaturesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerImageSignaturesDataSourceCrud) Get() error {
	request := oci_artifacts.ListContainerImageSignaturesRequest{}

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

	if imageDigest, ok := s.D.GetOkExists("image_digest"); ok {
		tmp := imageDigest.(string)
		request.ImageDigest = &tmp
	}

	if imageId, ok := s.D.GetOkExists("image_id"); ok {
		tmp := imageId.(string)
		request.ImageId = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if kmsKeyVersionId, ok := s.D.GetOkExists("kms_key_version_id"); ok {
		tmp := kmsKeyVersionId.(string)
		request.KmsKeyVersionId = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if repositoryName, ok := s.D.GetOkExists("repository_name"); ok {
		tmp := repositoryName.(string)
		request.RepositoryName = &tmp
	}

	if signingAlgorithm, ok := s.D.GetOkExists("signing_algorithm"); ok {
		request.SigningAlgorithm = oci_artifacts.ListContainerImageSignaturesSigningAlgorithmEnum(signingAlgorithm.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.ListContainerImageSignatures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListContainerImageSignatures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *ArtifactsContainerImageSignaturesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("ArtifactsContainerImageSignaturesDataSource-", ArtifactsContainerImageSignaturesDataSource(), s.D))
	resources := []map[string]interface{}{}
	containerImageSignature := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, ContainerImageSignatureSummaryToMap(item))
	}
	containerImageSignature["items"] = items

	if s.Res.RemainingItemsCount != nil {
		containerImageSignature["remaining_items_count"] = *s.Res.RemainingItemsCount
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, ArtifactsContainerImageSignaturesDataSource().Schema["container_image_signature_collection"].Elem.(*schema.Resource).Schema)
		containerImageSignature["items"] = items
	}

	resources = append(resources, containerImageSignature)
	if err := s.D.Set("container_image_signature_collection", resources); err != nil {
		return err
	}

	return nil
}
