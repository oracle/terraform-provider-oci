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

func ArtifactsGenericArtifactResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createArtifactsGenericArtifact,
		Read:     readArtifactsGenericArtifact,
		Update:   updateArtifactsGenericArtifact,
		Delete:   deleteArtifactsGenericArtifact,
		Schema: map[string]*schema.Schema{
			// Required
			"artifact_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"artifact_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"repository_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sha256": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createArtifactsGenericArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsGenericArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.CreateResource(d, sync)
}

func readArtifactsGenericArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsGenericArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

func updateArtifactsGenericArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsGenericArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteArtifactsGenericArtifact(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsGenericArtifactResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ArtifactsGenericArtifactResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_artifacts.ArtifactsClient
	Res                    *oci_artifacts.GenericArtifact
	DisableNotFoundRetries bool
}

func (s *ArtifactsGenericArtifactResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ArtifactsGenericArtifactResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ArtifactsGenericArtifactResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_artifacts.GenericArtifactLifecycleStateAvailable),
	}
}

func (s *ArtifactsGenericArtifactResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_artifacts.GenericArtifactLifecycleStateDeleting),
	}
}

func (s *ArtifactsGenericArtifactResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_artifacts.GenericArtifactLifecycleStateDeleted),
	}
}

func (s *ArtifactsGenericArtifactResourceCrud) Create() error {
	request := oci_artifacts.UpdateGenericArtifactRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateGenericArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GenericArtifact
	return nil
}

func (s *ArtifactsGenericArtifactResourceCrud) Get() error {
	request := oci_artifacts.GetGenericArtifactRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	} else {
		tmp := s.D.Id()
		request.ArtifactId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.GetGenericArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GenericArtifact
	return nil
}

func (s *ArtifactsGenericArtifactResourceCrud) Update() error {
	request := oci_artifacts.UpdateGenericArtifactRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateGenericArtifact(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.GenericArtifact
	return nil
}

func (s *ArtifactsGenericArtifactResourceCrud) Delete() error {
	request := oci_artifacts.DeleteGenericArtifactRequest{}

	if artifactId, ok := s.D.GetOkExists("artifact_id"); ok {
		tmp := artifactId.(string)
		request.ArtifactId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.DeleteGenericArtifact(context.Background(), request)
	return err
}

func (s *ArtifactsGenericArtifactResourceCrud) SetData() error {
	if s.Res.Id != nil {
		s.D.Set("artifact_id", *s.Res.Id)
	}

	if s.Res.ArtifactPath != nil {
		s.D.Set("artifact_path", *s.Res.ArtifactPath)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.RepositoryId != nil {
		s.D.Set("repository_id", *s.Res.RepositoryId)
	}

	if s.Res.Sha256 != nil {
		s.D.Set("sha256", *s.Res.Sha256)
	}

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}
func GenericArtifactSummaryToMap(obj oci_artifacts.GenericArtifactSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArtifactPath != nil {
		result["artifact_path"] = string(*obj.ArtifactPath)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
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

	if obj.Sha256 != nil {
		result["sha256"] = string(*obj.Sha256)
	}

	if obj.SizeInBytes != nil {
		result["size_in_bytes"] = strconv.FormatInt(*obj.SizeInBytes, 10)
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
