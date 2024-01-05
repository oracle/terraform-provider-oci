// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ArtifactsContainerImageSignatureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["image_signature_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(ArtifactsContainerImageSignatureResource(), fieldMap, readSingularArtifactsContainerImageSignature)
}

func readSingularArtifactsContainerImageSignature(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerImageSignatureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

type ArtifactsContainerImageSignatureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_artifacts.ArtifactsClient
	Res    *oci_artifacts.GetContainerImageSignatureResponse
}

func (s *ArtifactsContainerImageSignatureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *ArtifactsContainerImageSignatureDataSourceCrud) Get() error {
	request := oci_artifacts.GetContainerImageSignatureRequest{}

	if imageSignatureId, ok := s.D.GetOkExists("image_signature_id"); ok {
		tmp := imageSignatureId.(string)
		request.ImageSignatureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "artifacts")

	response, err := s.Client.GetContainerImageSignature(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *ArtifactsContainerImageSignatureDataSourceCrud) SetData() error {
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

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
