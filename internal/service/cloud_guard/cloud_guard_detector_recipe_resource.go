// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v56/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func CloudGuardDetectorRecipeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardDetectorRecipe,
		Read:     readCloudGuardDetectorRecipe,
		Update:   updateCloudGuardDetectorRecipe,
		Delete:   deleteCloudGuardDetectorRecipe,
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
			"source_detector_recipe_id": {
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
			"detector_rules": {
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
									"risk_level": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: utils.JsonStringDiffSuppressFunction,
									},
									"configurations": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"config_key": {
													Type:     schema.TypeString,
													Required: true,
												},
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"data_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"values": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required
															"list_type": {
																Type:     schema.TypeString,
																Required: true,
															},
															"managed_list_type": {
																Type:     schema.TypeString,
																Required: true,
															},
															"value": {
																Type:     schema.TypeString,
																Required: true,
															},

															// Optional

															// Computed
														},
													},
												},

												// Computed
											},
										},
									},
									"labels": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
									"is_configuration_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"detector_rule_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"candidate_responder_rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_preferred": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"detector": {
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
						"managed_list_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"recommendation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"detector": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"effective_detector_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"candidate_responder_rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_preferred": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
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
												"data_type": {
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
												"values": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"list_type": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"managed_list_type": {
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
											},
										},
									},
									"is_configuration_allowed": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"labels": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"risk_level": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"detector": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"detector_rule_id": {
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
						"managed_list_types": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"recommendation": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"service_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
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

func createCloudGuardDetectorRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardDetectorRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardDetectorRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardDetectorRecipe(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDetectorRecipeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardDetectorRecipeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.DetectorRecipe
	DisableNotFoundRetries bool
}

func (s *CloudGuardDetectorRecipeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardDetectorRecipeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardDetectorRecipeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardDetectorRecipeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardDetectorRecipeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardDetectorRecipeResourceCrud) Create() error {
	request := oci_cloud_guard.CreateDetectorRecipeRequest{}

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

	if detectorRules, ok := s.D.GetOkExists("detector_rules"); ok {
		interfaces := detectorRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateDetectorRecipeDetectorRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "detector_rules", stateDataIndex)
			converted, err := s.mapToUpdateDetectorRecipeDetectorRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("detector_rules") {
			request.DetectorRules = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if sourceDetectorRecipeId, ok := s.D.GetOkExists("source_detector_recipe_id"); ok {
		tmp := sourceDetectorRecipeId.(string)
		request.SourceDetectorRecipeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateDetectorRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectorRecipe
	return nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) Get() error {
	request := oci_cloud_guard.GetDetectorRecipeRequest{}

	tmp := s.D.Id()
	request.DetectorRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetDetectorRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectorRecipe
	return nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_guard.UpdateDetectorRecipeRequest{}

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
	request.DetectorRecipeId = &tmp

	if detectorRules, ok := s.D.GetOkExists("detector_rules"); ok {
		interfaces := detectorRules.([]interface{})
		tmp := make([]oci_cloud_guard.UpdateDetectorRecipeDetectorRule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "detector_rules", stateDataIndex)
			converted, err := s.mapToUpdateDetectorRecipeDetectorRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("detector_rules") {
			request.DetectorRules = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateDetectorRecipe(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DetectorRecipe
	return nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteDetectorRecipeRequest{}

	tmp := s.D.Id()
	request.DetectorRecipeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteDetectorRecipe(context.Background(), request)
	return err
}

func (s *CloudGuardDetectorRecipeResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("detector", s.Res.Detector)

	detectorRules := []interface{}{}
	for _, item := range s.Res.DetectorRules {
		detectorRules = append(detectorRules, DetectorRecipeDetectorRuleToMap(item))
	}
	s.D.Set("detector_rules", detectorRules)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	effectiveDetectorRules := []interface{}{}
	for _, item := range s.Res.EffectiveDetectorRules {
		effectiveDetectorRules = append(effectiveDetectorRules, DetectorRecipeDetectorRuleToMap(item))
	}
	s.D.Set("effective_detector_rules", effectiveDetectorRules)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("owner", s.Res.Owner)

	if s.Res.SourceDetectorRecipeId != nil {
		s.D.Set("source_detector_recipe_id", *s.Res.SourceDetectorRecipeId)
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

func CandidateResponderRuleToMap(obj oci_cloud_guard.CandidateResponderRule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsPreferred != nil {
		result["is_preferred"] = bool(*obj.IsPreferred)
	}

	return result
}

func ConfigValueToMap(obj oci_cloud_guard.ConfigValue) map[string]interface{} {
	result := map[string]interface{}{}

	result["list_type"] = string(obj.ListType)

	if obj.ManagedListType != nil {
		result["managed_list_type"] = string(*obj.ManagedListType)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func DetectorConfigurationToMap(obj oci_cloud_guard.DetectorConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	values := []interface{}{}
	for _, item := range obj.Values {
		values = append(values, ConfigValueToMap(item))
	}
	result["values"] = values

	return result
}

func DetectorDetailsToMap(obj *oci_cloud_guard.DetectorDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Condition != nil {
		condition, err := tfresource.ConvertObjectToJsonString(obj.Condition)
		if err == nil {
			result["condition"] = condition
		}
	}

	configurations := []interface{}{}
	for _, item := range obj.Configurations {
		configurations = append(configurations, DetectorConfigurationToMap(item))
	}
	result["configurations"] = configurations

	if obj.IsConfigurationAllowed != nil {
		result["is_configuration_allowed"] = bool(*obj.IsConfigurationAllowed)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["labels"] = obj.Labels

	result["risk_level"] = string(obj.RiskLevel)

	return result
}

func DetectorRecipeDetectorRuleToMap(obj oci_cloud_guard.DetectorRecipeDetectorRule) map[string]interface{} {
	result := map[string]interface{}{}

	candidateResponderRules := []interface{}{}
	for _, item := range obj.CandidateResponderRules {
		candidateResponderRules = append(candidateResponderRules, CandidateResponderRuleToMap(item))
	}
	result["candidate_responder_rules"] = candidateResponderRules

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Details != nil {
		result["details"] = []interface{}{DetectorDetailsToMap(obj.Details)}
	}

	result["detector"] = string(obj.Detector)

	if obj.DetectorRuleId != nil {
		result["detector_rule_id"] = string(*obj.DetectorRuleId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["managed_list_types"] = obj.ManagedListTypes

	if obj.Recommendation != nil {
		result["recommendation"] = string(*obj.Recommendation)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	if obj.ServiceType != nil {
		result["service_type"] = string(*obj.ServiceType)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DetectorRecipeSummaryToMap(obj oci_cloud_guard.DetectorRecipeSummary) map[string]interface{} {
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

	result["detector"] = string(obj.Detector)

	detectorRules := []interface{}{}
	for _, item := range obj.DetectorRules {
		detectorRules = append(detectorRules, DetectorRecipeDetectorRuleToMap(item))
	}
	result["detector_rules"] = detectorRules

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["owner"] = string(obj.Owner)

	if obj.SourceDetectorRecipeId != nil {
		result["source_detector_recipe_id"] = string(*obj.SourceDetectorRecipeId)
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

func (s *CloudGuardDetectorRecipeResourceCrud) mapToUpdateDetectorRecipeDetectorRule(fieldKeyFormat string) (oci_cloud_guard.UpdateDetectorRecipeDetectorRule, error) {
	result := oci_cloud_guard.UpdateDetectorRecipeDetectorRule{}

	if details, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "details")); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "details"), 0)
			tmp, err := s.mapToUpdateDetectorRuleDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert details, encountered error: %v", err)
			}
			result.Details = &tmp
		}
	}

	if detectorRuleId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "detector_rule_id")); ok {
		tmp := detectorRuleId.(string)
		result.DetectorRuleId = &tmp
	}

	return result, nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) mapToUpdateDetectorRuleDetails(fieldKeyFormat string) (oci_cloud_guard.UpdateDetectorRuleDetails, error) {
	result := oci_cloud_guard.UpdateDetectorRuleDetails{}

	//Condition Modelling
	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		tmp := condition.(string)
		if len(tmp) > 0 {
			var err error
			result.Condition, err = jsonToCondition(tmp)
			if err != nil {
				return result, err
			}
		}
	}

	if configurations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configurations")); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_cloud_guard.DetectorConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configurations"), stateDataIndex)
			converted, err := s.mapToDetectorConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "configurations")) {
			result.Configurations = tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enabled")); ok {
		tmp := isEnabled.(bool)
		result.IsEnabled = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if riskLevel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "risk_level")); ok {
		result.RiskLevel = oci_cloud_guard.RiskLevelEnum(riskLevel.(string))
	}

	return result, nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_guard.ChangeDetectorRecipeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DetectorRecipeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.ChangeDetectorRecipeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

type cloudGuardCondition struct {
	JsonData []byte
	Kind     string `json:"kind"`
}

func jsonToCondition(data string) (oci_cloud_guard.Condition, error) {
	var val cloudGuardCondition
	if err := json.Unmarshal([]byte(data), &val); err == nil {
		if schemaData, err := UnmarshalPolymorphicConditionJSON(val.Kind, data); err == nil {
			return schemaData, nil
		} else {
			return nil, err
		}
	}
	return nil, nil
}

func UnmarshalPolymorphicConditionJSON(kind string, data string) (oci_cloud_guard.Condition, error) {
	var err error
	switch kind {
	case "SIMPLE":
		mm := oci_cloud_guard.SimpleCondition{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	case "COMPOSITE":
		mm := oci_cloud_guard.CompositeCondition{}
		err = json.Unmarshal([]byte(data), &mm)
		return mm, err
	default:
		return nil, nil
	}
}

func (s *CloudGuardDetectorRecipeResourceCrud) mapToDetectorConfiguration(fieldKeyFormat string) (oci_cloud_guard.DetectorConfiguration, error) {
	result := oci_cloud_guard.DetectorConfiguration{}

	if configKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_key")); ok {
		tmp := configKey.(string)
		result.ConfigKey = &tmp
	}

	if dataType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "data_type")); ok {
		tmp := dataType.(string)
		result.DataType = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		interfaces := values.([]interface{})
		tmp := make([]oci_cloud_guard.ConfigValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "values"), stateDataIndex)
			converted, err := s.mapToConfigValue(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
			result.Values = tmp
		}
	}
	return result, nil
}

func (s *CloudGuardDetectorRecipeResourceCrud) mapToConfigValue(fieldKeyFormat string) (oci_cloud_guard.ConfigValue, error) {
	result := oci_cloud_guard.ConfigValue{}

	if listType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "list_type")); ok {
		result.ListType = oci_cloud_guard.ConfigurationListItemTypeEnum(listType.(string))
	}

	if managedListType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_list_type")); ok {
		tmp := managedListType.(string)
		result.ManagedListType = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
