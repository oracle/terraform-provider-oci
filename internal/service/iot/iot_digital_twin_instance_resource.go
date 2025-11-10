// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotDigitalTwinInstanceWithContext,
		ReadContext:   readIotDigitalTwinInstanceWithContext,
		UpdateContext: updateIotDigitalTwinInstanceWithContext,
		DeleteContext: deleteIotDigitalTwinInstanceWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"auth_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
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
			"digital_twin_adapter_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digital_twin_model_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digital_twin_model_spec_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_key": {
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

func createIotDigitalTwinInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotDigitalTwinInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotDigitalTwinInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotDigitalTwinInstanceWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotDigitalTwinInstanceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.DigitalTwinInstance
	DisableNotFoundRetries bool
}

func (s *IotDigitalTwinInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotDigitalTwinInstanceResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateActive),
	}
}

func (s *IotDigitalTwinInstanceResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateDeleted),
	}
}

func (s *IotDigitalTwinInstanceResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateDigitalTwinInstanceRequest{}

	if authId, ok := s.D.GetOkExists("auth_id"); ok {
		tmp := authId.(string)
		request.AuthId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if digitalTwinAdapterId, ok := s.D.GetOkExists("digital_twin_adapter_id"); ok {
		tmp := digitalTwinAdapterId.(string)
		request.DigitalTwinAdapterId = &tmp
	}

	if digitalTwinModelId, ok := s.D.GetOkExists("digital_twin_model_id"); ok {
		tmp := digitalTwinModelId.(string)
		request.DigitalTwinModelId = &tmp
	}

	if digitalTwinModelSpecUri, ok := s.D.GetOkExists("digital_twin_model_spec_uri"); ok {
		tmp := digitalTwinModelSpecUri.(string)
		request.DigitalTwinModelSpecUri = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalKey, ok := s.D.GetOkExists("external_key"); ok {
		tmp := externalKey.(string)
		request.ExternalKey = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateDigitalTwinInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinInstance
	return nil
}

func (s *IotDigitalTwinInstanceResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinInstanceRequest{}

	tmp := s.D.Id()
	request.DigitalTwinInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetDigitalTwinInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinInstance
	return nil
}

func (s *IotDigitalTwinInstanceResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateDigitalTwinInstanceRequest{}

	if authId, ok := s.D.GetOkExists("auth_id"); ok {
		tmp := authId.(string)
		request.AuthId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if digitalTwinAdapterId, ok := s.D.GetOkExists("digital_twin_adapter_id"); ok {
		tmp := digitalTwinAdapterId.(string)
		request.DigitalTwinAdapterId = &tmp
	}

	tmp := s.D.Id()
	request.DigitalTwinInstanceId = &tmp

	if digitalTwinModelId, ok := s.D.GetOkExists("digital_twin_model_id"); ok {
		tmp := digitalTwinModelId.(string)
		request.DigitalTwinModelId = &tmp
	}

	if digitalTwinModelSpecUri, ok := s.D.GetOkExists("digital_twin_model_spec_uri"); ok {
		tmp := digitalTwinModelSpecUri.(string)
		request.DigitalTwinModelSpecUri = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if externalKey, ok := s.D.GetOkExists("external_key"); ok {
		tmp := externalKey.(string)
		request.ExternalKey = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateDigitalTwinInstance(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinInstance
	return nil
}

func (s *IotDigitalTwinInstanceResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteDigitalTwinInstanceRequest{}

	tmp := s.D.Id()
	request.DigitalTwinInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	_, err := s.Client.DeleteDigitalTwinInstance(ctx, request)
	return err
}

func (s *IotDigitalTwinInstanceResourceCrud) SetData() error {
	if s.Res.AuthId != nil {
		s.D.Set("auth_id", *s.Res.AuthId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DigitalTwinAdapterId != nil {
		s.D.Set("digital_twin_adapter_id", *s.Res.DigitalTwinAdapterId)
	}

	if s.Res.DigitalTwinModelId != nil {
		s.D.Set("digital_twin_model_id", *s.Res.DigitalTwinModelId)
	}

	if s.Res.DigitalTwinModelSpecUri != nil {
		s.D.Set("digital_twin_model_spec_uri", *s.Res.DigitalTwinModelSpecUri)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExternalKey != nil {
		s.D.Set("external_key", *s.Res.ExternalKey)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
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

func DigitalTwinInstanceSummaryToMap(obj oci_iot.DigitalTwinInstanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuthId != nil {
		result["auth_id"] = string(*obj.AuthId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DigitalTwinAdapterId != nil {
		result["digital_twin_adapter_id"] = string(*obj.DigitalTwinAdapterId)
	}

	if obj.DigitalTwinModelId != nil {
		result["digital_twin_model_id"] = string(*obj.DigitalTwinModelId)
	}

	if obj.DigitalTwinModelSpecUri != nil {
		result["digital_twin_model_spec_uri"] = string(*obj.DigitalTwinModelSpecUri)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ExternalKey != nil {
		result["external_key"] = string(*obj.ExternalKey)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IotDomainId != nil {
		result["iot_domain_id"] = string(*obj.IotDomainId)
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
