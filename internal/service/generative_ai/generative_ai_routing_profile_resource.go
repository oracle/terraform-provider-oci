// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_generative_ai "github.com/oracle/oci-go-sdk/v65/generativeai"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiRoutingProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGenerativeAiRoutingProfileWithContext,
		ReadContext:   readGenerativeAiRoutingProfileWithContext,
		UpdateContext: updateGenerativeAiRoutingProfileWithContext,
		DeleteContext: deleteGenerativeAiRoutingProfileWithContext,
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
			"description": {
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
			"model_routing_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allowed_models": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},
			"region_routing_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allowed_regions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_state": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_routing_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_models": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"region_routing_policy": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_regions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
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
				},
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

func createGenerativeAiRoutingProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiRoutingProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readGenerativeAiRoutingProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiRoutingProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGenerativeAiRoutingProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiRoutingProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteGenerativeAiRoutingProfileWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GenerativeAiRoutingProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GenerativeAiRoutingProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai.GenerativeAiClient
	Res                    *oci_generative_ai.RoutingProfile
	DisableNotFoundRetries bool
}

func (s *GenerativeAiRoutingProfileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiRoutingProfileResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai.RoutingProfileLifecycleStateCreating),
	}
}

func (s *GenerativeAiRoutingProfileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai.RoutingProfileLifecycleStateActive),
	}
}

func (s *GenerativeAiRoutingProfileResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai.RoutingProfileLifecycleStateDeleting),
	}
}

func (s *GenerativeAiRoutingProfileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai.RoutingProfileLifecycleStateDeleted),
	}
}

func (s *GenerativeAiRoutingProfileResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_generative_ai.CreateRoutingProfileRequest{}

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

	if modelRoutingPolicy, ok := s.D.GetOkExists("model_routing_policy"); ok {
		if tmpList := modelRoutingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_routing_policy", 0)
			tmp, err := s.mapToModelRoutingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelRoutingPolicy = &tmp
		}
	}

	if regionRoutingPolicy, ok := s.D.GetOkExists("region_routing_policy"); ok {
		if tmpList := regionRoutingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "region_routing_policy", 0)
			tmp, err := s.mapToRegionRoutingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegionRoutingPolicy = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.CreateRoutingProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.RoutingProfile
	return nil
}

func (s *GenerativeAiRoutingProfileResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_generative_ai.GetRoutingProfileRequest{}

	tmp := s.D.Id()
	request.RoutingProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.GetRoutingProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.RoutingProfile
	return nil
}

func (s *GenerativeAiRoutingProfileResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_generative_ai.UpdateRoutingProfileRequest{}

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

	if modelRoutingPolicy, ok := s.D.GetOkExists("model_routing_policy"); ok {
		if tmpList := modelRoutingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_routing_policy", 0)
			tmp, err := s.mapToModelRoutingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ModelRoutingPolicy = &tmp
		}
	}

	if regionRoutingPolicy, ok := s.D.GetOkExists("region_routing_policy"); ok {
		if tmpList := regionRoutingPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "region_routing_policy", 0)
			tmp, err := s.mapToRegionRoutingPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegionRoutingPolicy = &tmp
		}
	}

	tmp := s.D.Id()
	request.RoutingProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	response, err := s.Client.UpdateRoutingProfile(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.RoutingProfile
	return nil
}

func (s *GenerativeAiRoutingProfileResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_generative_ai.DeleteRoutingProfileRequest{}

	tmp := s.D.Id()
	request.RoutingProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.DeleteRoutingProfile(ctx, request)
	return err
}

func (s *GenerativeAiRoutingProfileResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ModelRoutingPolicy != nil {
		s.D.Set("model_routing_policy", []interface{}{ModelRoutingPolicyToMap(s.Res.ModelRoutingPolicy)})
	} else {
		s.D.Set("model_routing_policy", nil)
	}

	if s.Res.PreviousState != nil {
		s.D.Set("previous_state", []interface{}{PreviousRoutingProfileStateToMap(s.Res.PreviousState)})
	} else {
		s.D.Set("previous_state", nil)
	}

	if s.Res.RegionRoutingPolicy != nil {
		s.D.Set("region_routing_policy", []interface{}{RegionRoutingPolicyToMap(s.Res.RegionRoutingPolicy)})
	} else {
		s.D.Set("region_routing_policy", nil)
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

func ModelRoutingPolicyToMap(obj *oci_generative_ai.ModelRoutingPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_models"] = obj.AllowedModels

	return result
}

func PreviousRoutingProfileStateToMap(obj *oci_generative_ai.PreviousRoutingProfileState) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ModelRoutingPolicy != nil {
		result["model_routing_policy"] = []interface{}{ModelRoutingPolicyToMap(obj.ModelRoutingPolicy)}
	}

	if obj.RegionRoutingPolicy != nil {
		result["region_routing_policy"] = []interface{}{RegionRoutingPolicyToMap(obj.RegionRoutingPolicy)}
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

func (s *GenerativeAiRoutingProfileResourceCrud) mapToRegionRoutingPolicy(fieldKeyFormat string) (oci_generative_ai.RegionRoutingPolicy, error) {
	result := oci_generative_ai.RegionRoutingPolicy{}

	if allowedRegions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_regions")); ok {
		interfaces := allowedRegions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_regions")) {
			result.AllowedRegions = tmp
		}
	}

	return result, nil
}

func RegionRoutingPolicyToMap(obj *oci_generative_ai.RegionRoutingPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["allowed_regions"] = obj.AllowedRegions

	return result
}

func RoutingProfileSummaryToMap(obj oci_generative_ai.RoutingProfileSummary) map[string]interface{} {
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ModelRoutingPolicy != nil {
		result["model_routing_policy"] = []interface{}{ModelRoutingPolicyToMap(obj.ModelRoutingPolicy)}
	}

	if obj.RegionRoutingPolicy != nil {
		result["region_routing_policy"] = []interface{}{RegionRoutingPolicyToMap(obj.RegionRoutingPolicy)}
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

func (s *GenerativeAiRoutingProfileResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_generative_ai.ChangeRoutingProfileCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.RoutingProfileId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai")

	_, err := s.Client.ChangeRoutingProfileCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *GenerativeAiRoutingProfileResourceCrud) mapToModelRoutingPolicy(fieldKeyFormat string) (oci_generative_ai.ModelRoutingPolicy, error) {
	result := oci_generative_ai.ModelRoutingPolicy{}

	if allowedModels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allowed_models")); ok {
		interfaces := allowedModels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "allowed_models")) {
			result.AllowedModels = tmp
		}
	}

	return result, nil
}
