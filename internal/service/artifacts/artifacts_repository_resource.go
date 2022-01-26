// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_artifacts "github.com/oracle/oci-go-sdk/v56/artifacts"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func ArtifactsRepositoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createArtifactsRepository,
		Read:     readArtifactsRepository,
		Update:   updateArtifactsRepository,
		Delete:   deleteArtifactsRepository,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_immutable": {
				Type:     schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
			"repository_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"state": {
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

func createArtifactsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.CreateResource(d, sync)
}

func readArtifactsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

func updateArtifactsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteArtifactsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ArtifactsRepositoryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_artifacts.ArtifactsClient
	Res                    *oci_artifacts.Repository
	DisableNotFoundRetries bool
}

func (s *ArtifactsRepositoryResourceCrud) ID() string {
	repository := *s.Res
	return *repository.GetId()
}

func (s *ArtifactsRepositoryResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ArtifactsRepositoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_artifacts.RepositoryLifecycleStateAvailable),
	}
}

func (s *ArtifactsRepositoryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_artifacts.RepositoryLifecycleStateDeleting),
	}
}

func (s *ArtifactsRepositoryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_artifacts.RepositoryLifecycleStateDeleted),
	}
}

func (s *ArtifactsRepositoryResourceCrud) Create() error {
	request := oci_artifacts.CreateRepositoryRequest{}
	err := s.populateTopLevelPolymorphicCreateRepositoryRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.CreateRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Repository
	return nil
}

func (s *ArtifactsRepositoryResourceCrud) Get() error {
	request := oci_artifacts.GetRepositoryRequest{}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.GetRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Repository
	return nil
}

func (s *ArtifactsRepositoryResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_artifacts.UpdateRepositoryRequest{}
	err := s.populateTopLevelPolymorphicUpdateRepositoryRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Repository
	return nil
}

func (s *ArtifactsRepositoryResourceCrud) Delete() error {
	request := oci_artifacts.DeleteRepositoryRequest{}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.DeleteRepository(context.Background(), request)
	return err
}

func (s *ArtifactsRepositoryResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_artifacts.GenericRepository:
		s.D.Set("repository_type", "GENERIC")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

		if v.IsImmutable != nil {
			s.D.Set("is_immutable", *v.IsImmutable)
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.GetTimeCreated().String())
		}

		s.D.Set("state", v.LifecycleState)
	default:
		log.Printf("[WARN] Received 'repository_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func RepositorySummaryToMap(obj oci_artifacts.RepositorySummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_artifacts.GenericRepositorySummary:
		result["repository_type"] = "GENERIC"
		if obj.GetId() != nil {
			result["id"] = string(*obj.GetId())
		}

		if obj.GetCompartmentId() != nil {
			result["compartment_id"] = string(*obj.GetCompartmentId())
		}

		if obj.GetDefinedTags() != nil {
			result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
		}

		if obj.GetDisplayName() != nil {
			result["display_name"] = string(*obj.GetDisplayName())
		}

		if obj.GetDescription() != nil {
			result["description"] = string(*obj.GetDescription())
		}

		if obj.GetIsImmutable() != nil {
			result["is_immutable"] = bool(*obj.GetIsImmutable())
		}

		result["state"] = string(obj.GetLifecycleState())

		if obj.GetTimeCreated() != nil {
			result["time_created"] = obj.GetTimeCreated().String()
		}

		result["freeform_tags"] = obj.GetFreeformTags()
	default:
		log.Printf("[WARN] Received 'repository_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *ArtifactsRepositoryResourceCrud) populateTopLevelPolymorphicCreateRepositoryRequest(request *oci_artifacts.CreateRepositoryRequest) error {
	//discriminator
	repositoryTypeRaw, ok := s.D.GetOkExists("repository_type")
	var repositoryType string
	if ok {
		repositoryType = repositoryTypeRaw.(string)
	} else {
		repositoryType = "" // default value
	}
	switch strings.ToLower(repositoryType) {
	case strings.ToLower("GENERIC"):
		details := oci_artifacts.CreateGenericRepositoryDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isImmutable, ok := s.D.GetOkExists("is_immutable"); ok {
			tmp := isImmutable.(bool)
			details.IsImmutable = &tmp
		}
		request.CreateRepositoryDetails = details
	default:
		return fmt.Errorf("unknown repository_type '%v' was specified", repositoryType)
	}
	return nil
}

func (s *ArtifactsRepositoryResourceCrud) populateTopLevelPolymorphicUpdateRepositoryRequest(request *oci_artifacts.UpdateRepositoryRequest) error {
	//discriminator
	repositoryTypeRaw, ok := s.D.GetOkExists("repository_type")
	var repositoryType string
	if ok {
		repositoryType = repositoryTypeRaw.(string)
	} else {
		repositoryType = "" // default value
	}
	switch strings.ToLower(repositoryType) {
	case strings.ToLower("GENERIC"):
		details := oci_artifacts.UpdateGenericRepositoryDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		tmp := s.D.Id()
		request.RepositoryId = &tmp
		request.UpdateRepositoryDetails = details
	default:
		return fmt.Errorf("unknown repository_type '%v' was specified", repositoryType)
	}
	return nil
}

func (s *ArtifactsRepositoryResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_artifacts.ChangeRepositoryCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RepositoryId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.ChangeRepositoryCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
