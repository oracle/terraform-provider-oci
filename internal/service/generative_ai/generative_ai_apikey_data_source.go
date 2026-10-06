// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiApikeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["api_key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(generativeAiApikeyDataSourceItemSchema(), fieldMap, readSingularGenerativeAiApikeyWithContext)
}

// generativeAiApikeyDataSourceItemSchema defines the read-only API key fields
// shared by the singular and collection data sources. API keys intentionally do
// not have a corresponding Terraform managed resource because create and renew
// return generated credential values that cannot be managed safely in state.
func generativeAiApikeyDataSourceItemSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_mask": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_activated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_deactivated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_expiry": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_used": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_revoked": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
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

func readSingularGenerativeAiApikeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiApikeyDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type GenerativeAiApikeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_generative_ai.GenerativeAiClient
	Res    *oci_generative_ai.GetApiKeyResponse
}

func (s *GenerativeAiApikeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *GenerativeAiApikeyDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetApiKeyRequest{}

	if apiKeyId, ok := s.D.GetOkExists("api_key_id"); ok {
		tmp := apiKeyId.(string)
		request.ApiKeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "generative_ai")

	response, err := s.Client.GetApiKey(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *GenerativeAiApikeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	keys := []interface{}{}
	for _, item := range s.Res.Keys {
		keys = append(keys, ApiKeyItemToMap(item))
	}
	s.D.Set("keys", keys)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

func ApiKeyItemToMap(obj oci_generative_ai.ApiKeyItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.KeyMask != nil {
		result["key_mask"] = string(*obj.KeyMask)
	}

	if obj.KeyName != nil {
		result["key_name"] = string(*obj.KeyName)
	}

	result["state"] = string(obj.State)

	if obj.TimeActivated != nil {
		result["time_activated"] = obj.TimeActivated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeDeactivated != nil {
		result["time_deactivated"] = obj.TimeDeactivated.String()
	}

	if obj.TimeExpiry != nil {
		result["time_expiry"] = obj.TimeExpiry.String()
	}

	if obj.TimeLastUsed != nil {
		result["time_last_used"] = obj.TimeLastUsed.String()
	}

	if obj.TimeRevoked != nil {
		result["time_revoked"] = obj.TimeRevoked.String()
	}

	return result
}

func ApiKeySummaryToMap(obj oci_generative_ai.ApiKeySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
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

	keys := []interface{}{}
	for _, item := range obj.Keys {
		keys = append(keys, ApiKeyItemToMap(item))
	}
	result["keys"] = keys

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
