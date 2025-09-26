// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotDigitalTwinModelResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotDigitalTwinModelWithContext,
		ReadContext:   readIotDigitalTwinModelWithContext,
		UpdateContext: updateIotDigitalTwinModelWithContext,
		DeleteContext: deleteIotDigitalTwinModelWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"spec": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringIsJSON,
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
			"spec_uri": {
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

func createIotDigitalTwinModelWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotDigitalTwinModelWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotDigitalTwinModelWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotDigitalTwinModelWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinModelResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotDigitalTwinModelResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.DigitalTwinModel
	DisableNotFoundRetries bool
}

func (s *IotDigitalTwinModelResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotDigitalTwinModelResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinModelResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateActive),
	}
}

func (s *IotDigitalTwinModelResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinModelResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateDeleted),
	}
}

func (s *IotDigitalTwinModelResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateDigitalTwinModelRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if spec, ok := s.D.GetOkExists("spec"); ok {
		specJsonStr := spec.(string)
		specMap, err := JsonStringToMap(specJsonStr)
		if err != nil {
			return err
		}
		request.Spec = specMap
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if iotDomainId, ok := s.D.GetOkExists("iot_domain_id"); ok {
		tmp := iotDomainId.(string)
		request.IotDomainId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateDigitalTwinModel(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinModel
	return nil
}

func (s *IotDigitalTwinModelResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinModelRequest{}

	tmp := s.D.Id()
	request.DigitalTwinModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetDigitalTwinModel(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinModel
	return nil
}

func (s *IotDigitalTwinModelResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateDigitalTwinModelRequest{}

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

	tmp := s.D.Id()
	request.DigitalTwinModelId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateDigitalTwinModel(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinModel
	return nil
}

func (s *IotDigitalTwinModelResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteDigitalTwinModelRequest{}

	tmp := s.D.Id()
	request.DigitalTwinModelId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	_, err := s.Client.DeleteDigitalTwinModel(ctx, request)
	return err
}

func (s *IotDigitalTwinModelResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IotDomainId != nil {
		s.D.Set("iot_domain_id", *s.Res.IotDomainId)
	}

	if s.Res.SpecUri != nil {
		s.D.Set("spec_uri", *s.Res.SpecUri)
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

func DigitalTwinModelSummaryToMap(obj oci_iot.DigitalTwinModelSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IotDomainId != nil {
		result["iot_domain_id"] = string(*obj.IotDomainId)
	}

	if obj.SpecUri != nil {
		result["spec_uri"] = string(*obj.SpecUri)
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
