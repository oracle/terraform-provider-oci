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

func ArtifactsContainerImageSignatureResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createArtifactsContainerImageSignature,
		Read:     readArtifactsContainerImageSignature,
		Delete:   deleteArtifactsContainerImageSignature,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kms_key_version_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"message": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"signing_algorithm": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createArtifactsContainerImageSignature(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageSignatureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.CreateResource(d, sync)
}

func readArtifactsContainerImageSignature(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageSignatureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

func deleteArtifactsContainerImageSignature(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageSignatureResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ArtifactsContainerImageSignatureResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_artifacts.ArtifactsClient
	Res                    *oci_artifacts.ContainerImageSignature
	DisableNotFoundRetries bool
}

func (s *ArtifactsContainerImageSignatureResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ArtifactsContainerImageSignatureResourceCrud) Create() error {
	request := oci_artifacts.CreateContainerImageSignatureRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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

	if message, ok := s.D.GetOkExists("message"); ok {
		tmp := message.(string)
		request.Message = &tmp
	}

	if signature, ok := s.D.GetOkExists("signature"); ok {
		tmp := signature.(string)
		request.Signature = &tmp
	}

	if signingAlgorithm, ok := s.D.GetOkExists("signing_algorithm"); ok {
		request.SigningAlgorithm = oci_artifacts.CreateContainerImageSignatureDetailsSigningAlgorithmEnum(signingAlgorithm.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.CreateContainerImageSignature(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerImageSignature
	return nil
}

func (s *ArtifactsContainerImageSignatureResourceCrud) Get() error {
	request := oci_artifacts.GetContainerImageSignatureRequest{}

	imageSignatureId := s.D.Id()
	request.ImageSignatureId = &imageSignatureId

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.GetContainerImageSignature(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerImageSignature
	return nil
}

func (s *ArtifactsContainerImageSignatureResourceCrud) Delete() error {
	request := oci_artifacts.DeleteContainerImageSignatureRequest{}

	tmp := s.D.Id()
	request.ImageSignatureId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.DeleteContainerImageSignature(context.Background(), request)
	return err
}

func (s *ArtifactsContainerImageSignatureResourceCrud) SetData() error {

	if s.Res.Id != nil {
		s.D.Set("image_signature_id", *s.Res.Id)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ImageId != nil {
		s.D.Set("image_id", *s.Res.ImageId)
	}

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.KmsKeyVersionId != nil {
		s.D.Set("kms_key_version_id", *s.Res.KmsKeyVersionId)
	}

	if s.Res.Message != nil {
		s.D.Set("message", *s.Res.Message)
	}

	if s.Res.Signature != nil {
		s.D.Set("signature", *s.Res.Signature)
	}

	s.D.Set("signing_algorithm", s.Res.SigningAlgorithm)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func ContainerImageSignatureSummaryToMap(obj oci_artifacts.ContainerImageSignatureSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ImageId != nil {
		result["image_id"] = string(*obj.ImageId)
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.KmsKeyVersionId != nil {
		result["kms_key_version_id"] = string(*obj.KmsKeyVersionId)
	}

	if obj.Message != nil {
		result["message"] = string(*obj.Message)
	}

	if obj.Signature != nil {
		result["signature"] = string(*obj.Signature)
	}

	result["signing_algorithm"] = string(obj.SigningAlgorithm)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}
