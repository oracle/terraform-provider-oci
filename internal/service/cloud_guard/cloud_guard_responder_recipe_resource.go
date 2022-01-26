// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func CloudGuardResponderRecipeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardResponderRecipe,
		Read:     readCloudGuardResponderRecipe,
		Update:   updateCloudGuardResponderRecipe,
		Delete:   deleteCloudGuardResponderRecipe,
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
			"source_responder_recipe_id": {
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
			"responder_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"is_enabled": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional

									// Computed
									"condition": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configurations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"config_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"responder_rule_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_modes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Computed
			"effective_responder_rules": {
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
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"details": {
							Type:     schema.TypeList,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"condition": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"configurations": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"config_key": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"mode": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"policies": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"responder_rule_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"supported_modes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
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
			"owner": {
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

func createCloudGuardResponderRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardResponderRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardResponderRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardResponderRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardResponderRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardResponderRecipeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.ResponderRecipe
	DisableNotFoundRetries bool
}

func (s *CloudGuardResponderRecipeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardResponderRecipeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardResponderRecipeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardResponderRecipeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardResponderRecipeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardResponderRecipeResourceCrud) Create() error {
	request := oci_cloud_guard.CreateResponderRecipeRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if responderRules, ok := s.D.GetOkExists("responder_rules"); ok {
		interfaces := responderRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateResponderRecipeResponderRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "responder_rules", stateDataIndex)
			converted, err := s.mapToUpdateResponderRecipeResponderRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("responder_rules") {
			request.ResponderRules = tmp
		}
	}

	if sourceResponderRecipeId, ok := s.D.GetOkExists("source_responder_recipe_id"); ok {
		tmp := sourceResponderRecipeId.(string)
		request.SourceResponderRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateResponderRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResponderRecipe
	return nil
}

func (s *CloudGuardResponderRecipeResourceCrud) Get() error {
	request := oci_cloud_guard.GetResponderRecipeRequest{}

	tmp := s.D.Id()
	request.ResponderRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetResponderRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResponderRecipe
	return nil
}

func (s *CloudGuardResponderRecipeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_guard.UpdateResponderRecipeRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.ResponderRecipeId = &tmp

	if responderRules, ok := s.D.GetOkExists("responder_rules"); ok {
		interfaces := responderRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateResponderRecipeResponderRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "responder_rules", stateDataIndex)
			converted, err := s.mapToUpdateResponderRecipeResponderRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("responder_rules") {
			request.ResponderRules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateResponderRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ResponderRecipe
	return nil
}

func (s *CloudGuardResponderRecipeResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteResponderRecipeRequest{}

	tmp := s.D.Id()
	request.ResponderRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteResponderRecipe(context.Background(), request)
	return err
}

func (s *CloudGuardResponderRecipeResourceCrud) SetData() error {
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

	effectiveResponderRules := []interface{}{}
	for _, item := range s.Res.EffectiveResponderRules {
		effectiveResponderRules = append(effectiveResponderRules, ResponderRecipeResponderRuleToMap(item))
	}
	s.D.Set("effective_responder_rules", effectiveResponderRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("owner", s.Res.Owner)

	responderRules := []interface{}{}
	for _, item := range s.Res.ResponderRules {
		responderRules = append(responderRules, ResponderRecipeResponderRuleToMap(item))
	}
	s.D.Set("responder_rules", responderRules)

	if s.Res.SourceResponderRecipeId != nil {
		s.D.Set("source_responder_recipe_id", *s.Res.SourceResponderRecipeId)
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

func ResponderConfigurationToMap(obj oci_cloud_guard.ResponderConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func ResponderRecipeResponderRuleToMap(obj oci_cloud_guard.ResponderRecipeResponderRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{ResponderRuleDetailsToMap(obj.Details)}
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["policies"] = obj.Policies

	if obj.ResponderRuleId != nil {
		result["responder_rule_id"] = string(*obj.ResponderRuleId)
	}

	result["state"] = string(obj.LifecycleState)

	result["supported_modes"] = obj.SupportedModes

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func ResponderRecipeSummaryToMap(obj oci_cloud_guard.ResponderRecipeSummary) map[string]interface{} {
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

	result["owner"] = string(obj.Owner)

	responderRules := []interface{}{}
	for _, item := range obj.ResponderRules {
		responderRules = append(responderRules, ResponderRecipeResponderRuleToMap(item))
	}
	result["responder_rules"] = responderRules

	if obj.SourceResponderRecipeId != nil {
		result["source_responder_recipe_id"] = string(*obj.SourceResponderRecipeId)
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

func ResponderRuleDetailsToMap(obj *oci_cloud_guard.ResponderRuleDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Condition != nil {
		condition, err := tfresource.ConvertObjectToJsonString(obj.Condition)
		if err == nil {
			result["condition"] = condition
		}
	}

	configurations := []interface{}{}
	for _, item := range obj.Configurations {
		configurations = append(configurations, ResponderConfigurationToMap(item))
	}
	result["configurations"] = configurations

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["mode"] = string(obj.Mode)

	return result
}

func (s *CloudGuardResponderRecipeResourceCrud) mapToUpdateResponderRecipeResponderRule(fieldKeyFormat string) (oci_cloud_guard.UpdateResponderRecipeResponderRule, error) {
	result := oci_cloud_guard.UpdateResponderRecipeResponderRule{}

	if details, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "details")); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "details"), 0)
			tmp, err := s.mapToUpdateResponderRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert details, encountered error: %v", err)
			}
			result.Details = &tmp
		}
	}

	if responderRuleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "responder_rule_id")); ok {
		tmp := responderRuleId.(string)
		result.ResponderRuleId = &tmp
	}

	return result, nil
}

func (s *CloudGuardResponderRecipeResourceCrud) mapToUpdateResponderRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateResponderRuleDetails, error) {
	result := oci_cloud_guard.UpdateResponderRuleDetails{}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	return result, nil
}

func (s *CloudGuardResponderRecipeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_guard.ChangeResponderRecipeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ResponderRecipeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.ChangeResponderRecipeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
