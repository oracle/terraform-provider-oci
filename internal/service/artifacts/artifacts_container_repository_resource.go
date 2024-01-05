// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package artifacts

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_artifacts "github.com/oracle/oci-go-sdk/v65/artifacts"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ArtifactsContainerRepositoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createArtifactsContainerRepository,
		Read:     readArtifactsContainerRepository,
		Update:   updateArtifactsContainerRepository,
		Delete:   deleteArtifactsContainerRepository,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
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
			"is_immutable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_public": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"readme": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"content": {
							Type:     schema.TypeString,
							Required: true,
						},
						"format": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"billable_size_in_gbs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"layer_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"layers_size_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
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
			"time_last_pushed": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createArtifactsContainerRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.CreateResource(d, sync)
}

func readArtifactsContainerRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.ReadResource(sync)
}

func updateArtifactsContainerRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteArtifactsContainerRepository(d *schema.ResourceData, m interface{}) error {
	sync := &ArtifactsContainerRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ArtifactsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ArtifactsContainerRepositoryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_artifacts.ArtifactsClient
	Res                    *oci_artifacts.ContainerRepository
	DisableNotFoundRetries bool
}

func (s *ArtifactsContainerRepositoryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ArtifactsContainerRepositoryResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *ArtifactsContainerRepositoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_artifacts.ContainerRepositoryLifecycleStateAvailable),
	}
}

func (s *ArtifactsContainerRepositoryResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_artifacts.ContainerRepositoryLifecycleStateDeleting),
	}
}

func (s *ArtifactsContainerRepositoryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_artifacts.ContainerRepositoryLifecycleStateDeleted),
	}
}

func (s *ArtifactsContainerRepositoryResourceCrud) Create() error {
	request := oci_artifacts.CreateContainerRepositoryRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isImmutable, ok := s.D.GetOkExists("is_immutable"); ok {
		tmp := isImmutable.(bool)
		request.IsImmutable = &tmp
	}

	if isPublic, ok := s.D.GetOkExists("is_public"); ok {
		tmp := isPublic.(bool)
		request.IsPublic = &tmp
	}

	if readme, ok := s.D.GetOkExists("readme"); ok {
		if tmpList := readme.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "readme", 0)
			tmp, err := s.mapToContainerRepositoryReadme(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Readme = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.CreateContainerRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerRepository
	return nil
}

func (s *ArtifactsContainerRepositoryResourceCrud) Get() error {
	request := oci_artifacts.GetContainerRepositoryRequest{}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.GetContainerRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerRepository
	return nil
}

func (s *ArtifactsContainerRepositoryResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_artifacts.UpdateContainerRepositoryRequest{}

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

	if isImmutable, ok := s.D.GetOkExists("is_immutable"); ok {
		tmp := isImmutable.(bool)
		request.IsImmutable = &tmp
	}

	if isPublic, ok := s.D.GetOkExists("is_public"); ok {
		tmp := isPublic.(bool)
		request.IsPublic = &tmp
	}

	if readme, ok := s.D.GetOkExists("readme"); ok {
		if tmpList := readme.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "readme", 0)
			tmp, err := s.mapToContainerRepositoryReadme(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Readme = &tmp
		}
	}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	response, err := s.Client.UpdateContainerRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ContainerRepository
	return nil
}

func (s *ArtifactsContainerRepositoryResourceCrud) Delete() error {
	request := oci_artifacts.DeleteContainerRepositoryRequest{}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.DeleteContainerRepository(context.Background(), request)
	return err
}

func (s *ArtifactsContainerRepositoryResourceCrud) SetData() error {

	if s.Res.BillableSizeInGBs != nil {
		s.D.Set("billable_size_in_gbs", strconv.FormatInt(*s.Res.BillableSizeInGBs, 10))
	}

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

	if s.Res.ImageCount != nil {
		s.D.Set("image_count", *s.Res.ImageCount)
	}

	if s.Res.IsImmutable != nil {
		s.D.Set("is_immutable", *s.Res.IsImmutable)
	}

	if s.Res.IsPublic != nil {
		s.D.Set("is_public", *s.Res.IsPublic)
	}

	if s.Res.LayerCount != nil {
		s.D.Set("layer_count", *s.Res.LayerCount)
	}

	if s.Res.LayersSizeInBytes != nil {
		s.D.Set("layers_size_in_bytes", strconv.FormatInt(*s.Res.LayersSizeInBytes, 10))
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.Readme != nil {
		s.D.Set("readme", []interface{}{ContainerRepositoryReadmeToMap(s.Res.Readme)})
	} else {
		s.D.Set("readme", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastPushed != nil {
		s.D.Set("time_last_pushed", s.Res.TimeLastPushed.String())
	}

	return nil
}

func (s *ArtifactsContainerRepositoryResourceCrud) mapToContainerRepositoryReadme(fieldKeyFormat string) (oci_artifacts.ContainerRepositoryReadme, error) {
	result := oci_artifacts.ContainerRepositoryReadme{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		tmp := content.(string)
		result.Content = &tmp
	}

	if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
		result.Format = oci_artifacts.ContainerRepositoryReadmeFormatEnum(format.(string))
	}

	return result, nil
}

func ContainerRepositoryReadmeToMap(obj *oci_artifacts.ContainerRepositoryReadme) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		result["content"] = string(*obj.Content)
	}

	result["format"] = string(obj.Format)

	return result
}

func ContainerRepositorySummaryToMap(obj oci_artifacts.ContainerRepositorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BillableSizeInGBs != nil {
		result["billable_size_in_gbs"] = strconv.FormatInt(*obj.BillableSizeInGBs, 10)
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

	if obj.ImageCount != nil {
		result["image_count"] = int(*obj.ImageCount)
	}

	if obj.IsPublic != nil {
		result["is_public"] = bool(*obj.IsPublic)
	}

	if obj.LayerCount != nil {
		result["layer_count"] = int(*obj.LayerCount)
	}

	if obj.LayersSizeInBytes != nil {
		result["layers_size_in_bytes"] = strconv.FormatInt(*obj.LayersSizeInBytes, 10)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *ArtifactsContainerRepositoryResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_artifacts.ChangeContainerRepositoryCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	tmp := s.D.Id()
	changeCompartmentRequest.RepositoryId = &tmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "artifacts")

	_, err := s.Client.ChangeContainerRepositoryCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
