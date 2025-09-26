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

func IotDigitalTwinRelationshipResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotDigitalTwinRelationshipWithContext,
		ReadContext:   readIotDigitalTwinRelationshipWithContext,
		UpdateContext: updateIotDigitalTwinRelationshipWithContext,
		DeleteContext: deleteIotDigitalTwinRelationshipWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"content_path": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"iot_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_digital_twin_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target_digital_twin_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"content": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsJSON,
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

func createIotDigitalTwinRelationshipWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotDigitalTwinRelationshipWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotDigitalTwinRelationshipWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotDigitalTwinRelationshipWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotDigitalTwinRelationshipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotDigitalTwinRelationshipResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.DigitalTwinRelationship
	DisableNotFoundRetries bool
}

func (s *IotDigitalTwinRelationshipResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotDigitalTwinRelationshipResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinRelationshipResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateActive),
	}
}

func (s *IotDigitalTwinRelationshipResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *IotDigitalTwinRelationshipResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.LifecycleStateDeleted),
	}
}

func (s *IotDigitalTwinRelationshipResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateDigitalTwinRelationshipRequest{}

	if content, ok := s.D.GetOkExists("content"); ok {
		contentStr := content.(string)
		contentMap, err := JsonStringToMap(contentStr)
		if err != nil {
			return err
		}
		request.Content = contentMap
	}

	if contentPath, ok := s.D.GetOkExists("content_path"); ok {
		tmp := contentPath.(string)
		request.ContentPath = &tmp
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

	if sourceDigitalTwinInstanceId, ok := s.D.GetOkExists("source_digital_twin_instance_id"); ok {
		tmp := sourceDigitalTwinInstanceId.(string)
		request.SourceDigitalTwinInstanceId = &tmp
	}

	if targetDigitalTwinInstanceId, ok := s.D.GetOkExists("target_digital_twin_instance_id"); ok {
		tmp := targetDigitalTwinInstanceId.(string)
		request.TargetDigitalTwinInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateDigitalTwinRelationship(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinRelationship
	return nil
}

func (s *IotDigitalTwinRelationshipResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetDigitalTwinRelationshipRequest{}

	tmp := s.D.Id()
	request.DigitalTwinRelationshipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetDigitalTwinRelationship(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinRelationship
	return nil
}

func (s *IotDigitalTwinRelationshipResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_iot.UpdateDigitalTwinRelationshipRequest{}

	if content, ok := s.D.GetOkExists("content"); ok {
		contentStr := content.(string)
		contentMap, err := JsonStringToMap(contentStr)
		if err != nil {
			return err
		}
		request.Content = contentMap
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

	tmp := s.D.Id()
	request.DigitalTwinRelationshipId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateDigitalTwinRelationship(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DigitalTwinRelationship
	return nil
}

func (s *IotDigitalTwinRelationshipResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteDigitalTwinRelationshipRequest{}

	tmp := s.D.Id()
	request.DigitalTwinRelationshipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	_, err := s.Client.DeleteDigitalTwinRelationship(ctx, request)
	return err
}

func (s *IotDigitalTwinRelationshipResourceCrud) SetData() error {

	if s.Res.Content != nil {
		contentStr, err := MapToJsonString(s.Res.Content)
		if err != nil {
			return err
		}
		s.D.Set("content", contentStr)
	}

	if s.Res.ContentPath != nil {
		s.D.Set("content_path", *s.Res.ContentPath)
	}

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

	if s.Res.SourceDigitalTwinInstanceId != nil {
		s.D.Set("source_digital_twin_instance_id", *s.Res.SourceDigitalTwinInstanceId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetDigitalTwinInstanceId != nil {
		s.D.Set("target_digital_twin_instance_id", *s.Res.TargetDigitalTwinInstanceId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func DigitalTwinRelationshipSummaryToMap(obj oci_iot.DigitalTwinRelationshipSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ContentPath != nil {
		result["content_path"] = string(*obj.ContentPath)
	}

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

	if obj.SourceDigitalTwinInstanceId != nil {
		result["source_digital_twin_instance_id"] = string(*obj.SourceDigitalTwinInstanceId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetDigitalTwinInstanceId != nil {
		result["target_digital_twin_instance_id"] = string(*obj.TargetDigitalTwinInstanceId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
