// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesStreamDistributionChannelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesStreamDistributionChannel,
		Read:     readMediaServicesStreamDistributionChannel,
		Update:   updateMediaServicesStreamDistributionChannel,
		Delete:   deleteMediaServicesStreamDistributionChannel,
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
			"domain_name": {
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

func createMediaServicesStreamDistributionChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamDistributionChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesStreamDistributionChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamDistributionChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesStreamDistributionChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamDistributionChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesStreamDistributionChannel(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesStreamDistributionChannelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesStreamDistributionChannelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.StreamDistributionChannel
	DisableNotFoundRetries bool
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.StreamDistributionChannelLifecycleStateActive),
		string(oci_media_services.StreamDistributionChannelLifecycleStateNeedsAttention),
	}
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.StreamDistributionChannelLifecycleStateDeleted),
	}
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) Create() error {
	request := oci_media_services.CreateStreamDistributionChannelRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateStreamDistributionChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamDistributionChannel
	return nil
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) Get() error {
	request := oci_media_services.GetStreamDistributionChannelRequest{}

	tmp := s.D.Id()
	request.StreamDistributionChannelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetStreamDistributionChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamDistributionChannel
	return nil
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_media_services.UpdateStreamDistributionChannelRequest{}

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
	request.StreamDistributionChannelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateStreamDistributionChannel(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.StreamDistributionChannel
	return nil
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) Delete() error {
	request := oci_media_services.DeleteStreamDistributionChannelRequest{}

	tmp := s.D.Id()
	request.StreamDistributionChannelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteStreamDistributionChannel(context.Background(), request)
	return err
}

func (s *MediaServicesStreamDistributionChannelResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainName != nil {
		s.D.Set("domain_name", *s.Res.DomainName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

func StreamDistributionChannelSummaryToMap(obj oci_media_services.StreamDistributionChannelSummary) map[string]interface{} {
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

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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

func (s *MediaServicesStreamDistributionChannelResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_media_services.ChangeStreamDistributionChannelCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.StreamDistributionChannelId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.ChangeStreamDistributionChannelCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
