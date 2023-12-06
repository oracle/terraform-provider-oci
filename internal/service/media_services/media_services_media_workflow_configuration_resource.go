// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesMediaWorkflowConfiguration,
		Read:     readMediaServicesMediaWorkflowConfiguration,
		Update:   updateMediaServicesMediaWorkflowConfiguration,
		Delete:   deleteMediaServicesMediaWorkflowConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parameters": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
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
			"lifecyle_details": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMediaServicesMediaWorkflowConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesMediaWorkflowConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesMediaWorkflowConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesMediaWorkflowConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesMediaWorkflowConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.MediaWorkflowConfiguration
	DisableNotFoundRetries bool
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowConfigurationLifecycleStateActive),
		string(oci_media_services.MediaWorkflowLifecycleStateNeedsAttention),
	}
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowConfigurationLifecycleStateDeleted),
	}
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) Create() error {
	request := oci_media_services.CreateMediaWorkflowConfigurationRequest{}

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

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		err := json.Unmarshal([]byte(parameters.(string)), &request.Parameters)
		if err != nil {
			return err
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateMediaWorkflowConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowConfiguration
	return nil
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) Get() error {
	request := oci_media_services.GetMediaWorkflowConfigurationRequest{}

	tmp := s.D.Id()
	request.MediaWorkflowConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetMediaWorkflowConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowConfiguration
	return nil
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_media_services.UpdateMediaWorkflowConfigurationRequest{}

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

	tmp := s.D.Id()
	request.MediaWorkflowConfigurationId = &tmp

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		err := json.Unmarshal([]byte(parameters.(string)), &request.Parameters)
		if err != nil {
			return err
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateMediaWorkflowConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowConfiguration
	return nil
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) Delete() error {
	request := oci_media_services.DeleteMediaWorkflowConfigurationRequest{}

	tmp := s.D.Id()
	request.MediaWorkflowConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteMediaWorkflowConfiguration(context.Background(), request)
	return err
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) SetData() error {
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

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.Parameters != nil {
		jsonStr, err := json.Marshal(s.Res.Parameters)
		if err == nil {
			s.D.Set("parameters", string(jsonStr))
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func MediaWorkflowConfigurationSummaryToMap(obj oci_media_services.MediaWorkflowConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *MediaServicesMediaWorkflowConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_media_services.ChangeMediaWorkflowConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MediaWorkflowConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.ChangeMediaWorkflowConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
