// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"bytes"
	"context"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_generic_artifacts_content "github.com/oracle/oci-go-sdk/v46/genericartifactscontent"
)

func init() {
	RegisterResource("oci_generic_artifacts_content_artifact_by_path", GenericArtifactsContentArtifactByPathResource())
}

func GenericArtifactsContentArtifactByPathResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: DefaultTimeout,
		Create:   createGenericArtifactsContentArtifactByPath,
		Read:     readGenericArtifactsContentArtifactByPath,
		Delete:   deleteGenericArtifactsContentArtifactByPath,
		Schema: map[string]*schema.Schema{
			// Required
			"repository_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"artifact_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"content": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Computed
			"artifact_id": {
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
			"sha256": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
				Type:     schema.TypeInt,
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
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func createGenericArtifactsContentArtifactByPath(d *schema.ResourceData, m interface{}) error {
	sync := &GenericArtifactsContentArtifactByPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).genericArtifactsContentClient()

	return CreateResource(d, sync)
}

func readGenericArtifactsContentArtifactByPath(d *schema.ResourceData, m interface{}) error {
	sync := &GenericArtifactsContentArtifactByPathResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).genericArtifactsContentClient()

	return ReadResource(sync)
}

func deleteGenericArtifactsContentArtifactByPath(d *schema.ResourceData, m interface{}) error {
	return nil
}

type GenericArtifactsContentArtifactByPathResourceCrud struct {
	BaseCrud
	Client                 *oci_generic_artifacts_content.GenericArtifactsContentClient
	Res                    *oci_generic_artifacts_content.PutGenericArtifactContentByPathResponse
	Content                *oci_generic_artifacts_content.GetGenericArtifactContentByPathResponse
	DisableNotFoundRetries bool
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generic_artifacts_content.GenericArtifactLifecycleStateAvailable),
	}
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generic_artifacts_content.GenericArtifactLifecycleStateDeleting),
	}
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generic_artifacts_content.GenericArtifactLifecycleStateDeleted),
	}
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) Get() error {
	request := oci_generic_artifacts_content.GetGenericArtifactContentByPathRequest{}

	if artifactPath, ok := s.D.GetOkExists("artifact_path"); ok {
		tmp := artifactPath.(string)
		request.ArtifactPath = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "generic_artifacts_content")

	response, err := s.Client.GetGenericArtifactContentByPath(context.Background(), request)
	if err != nil {
		return err
	}

	s.Content = &response
	return nil
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) Create() error {
	request := oci_generic_artifacts_content.PutGenericArtifactContentByPathRequest{}

	if genericArtifactContentBody, ok := s.D.GetOkExists("content"); ok {
		tmp := []byte(genericArtifactContentBody.(string))
		request.GenericArtifactContentBody = ioutil.NopCloser(bytes.NewBuffer(tmp))
	}

	if artifactPath, ok := s.D.GetOkExists("artifact_path"); ok {
		tmp := artifactPath.(string)
		request.ArtifactPath = &tmp
	}

	if repositoryId, ok := s.D.GetOkExists("repository_id"); ok {
		tmp := repositoryId.(string)
		request.RepositoryId = &tmp
	}

	if version, ok := s.D.GetOkExists("version"); ok {
		tmp := version.(string)
		request.Version = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "generic_artifacts_content")

	response, err := s.Client.PutGenericArtifactContentByPath(context.Background(), request)
	if err != nil {
		return err
	}
	s.Res = &response
	return nil
}

func (s *GenericArtifactsContentArtifactByPathResourceCrud) SetData() error {

	if s.Content != nil {
		contentReader := s.Content.Content
		contentArray, err := ioutil.ReadAll(contentReader)
		if err != nil {
			log.Printf("Unable to read 'content' from response. Error: %q", err)
			return err
		}
		s.D.Set("content", contentArray)
	}

	if s.Res == nil {
		return nil
	}

	if s.Res.ArtifactPath != nil {
		s.D.Set("artifact_path", *s.Res.ArtifactPath)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if s.Res.GenericArtifact.Id != nil {
		s.D.Set("artifact_id", *s.Res.GenericArtifact.Id)
	}

	return nil
}
