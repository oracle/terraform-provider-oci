// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package optimizer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_optimizer "github.com/oracle/oci-go-sdk/v65/optimizer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OptimizerProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOptimizerProfile,
		Read:     readOptimizerProfile,
		Update:   updateOptimizerProfile,
		Delete:   deleteOptimizerProfile,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"levels_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"items": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"level": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"recommendation_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"aggregation_interval_in_days": {
				Type:     schema.TypeInt,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"target_compartments": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"target_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"tag_definition_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"tag_namespace_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"tag_value_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"tag_values": {
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

						// Optional

						// Computed
					},
				},
			},

			// Computed
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
	}
}

func createOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.CreateResource(d, sync)
}

func readOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.ReadResource(sync)
}

func updateOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OptimizerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OptimizerProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_optimizer.OptimizerClient
	Res                    *oci_optimizer.Profile
	DisableNotFoundRetries bool
}

func (s *OptimizerProfileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OptimizerProfileResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateAttaching),
		string(oci_optimizer.LifecycleStateCreating),
	}
}

func (s *OptimizerProfileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateActive),
	}
}

func (s *OptimizerProfileResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDetaching),
		string(oci_optimizer.LifecycleStateDeleting),
	}
}

func (s *OptimizerProfileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_optimizer.LifecycleStateDeleted),
	}
}

func (s *OptimizerProfileResourceCrud) Create() error {
	request := oci_optimizer.CreateProfileRequest{}

	if aggregationIntervalInDays, ok := s.D.GetOkExists("aggregation_interval_in_days"); ok {
		tmp := aggregationIntervalInDays.(int)
		request.AggregationIntervalInDays = &tmp
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if levelsConfiguration, ok := s.D.GetOkExists("levels_configuration"); ok {
		if tmpList := levelsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "levels_configuration", 0)
			tmp, err := s.mapToLevelsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LevelsConfiguration = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if targetCompartments, ok := s.D.GetOkExists("target_compartments"); ok {
		if tmpList := targetCompartments.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_compartments", 0)
			tmp, err := s.mapToTargetCompartments(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetCompartments = &tmp
		}
	}

	if targetTags, ok := s.D.GetOkExists("target_tags"); ok {
		if tmpList := targetTags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_tags", 0)
			tmp, err := s.mapToTargetTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetTags = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.CreateProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OptimizerProfileResourceCrud) Get() error {
	request := oci_optimizer.GetProfileRequest{}

	tmp := s.D.Id()
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.GetProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OptimizerProfileResourceCrud) Update() error {
	request := oci_optimizer.UpdateProfileRequest{}

	if aggregationIntervalInDays, ok := s.D.GetOkExists("aggregation_interval_in_days"); ok {
		tmp := aggregationIntervalInDays.(int)
		request.AggregationIntervalInDays = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if levelsConfiguration, ok := s.D.GetOkExists("levels_configuration"); ok {
		if tmpList := levelsConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "levels_configuration", 0)
			tmp, err := s.mapToLevelsConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LevelsConfiguration = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	tmp := s.D.Id()
	request.ProfileId = &tmp

	if targetCompartments, ok := s.D.GetOkExists("target_compartments"); ok {
		if tmpList := targetCompartments.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_compartments", 0)
			tmp, err := s.mapToTargetCompartments(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetCompartments = &tmp
		}
	}

	if targetTags, ok := s.D.GetOkExists("target_tags"); ok {
		if tmpList := targetTags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_tags", 0)
			tmp, err := s.mapToTargetTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetTags = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.UpdateProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OptimizerProfileResourceCrud) Delete() error {
	request := oci_optimizer.DeleteProfileRequest{}

	tmp := s.D.Id()
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	_, err := s.Client.DeleteProfile(context.Background(), request)
	return err
}

func (s *OptimizerProfileResourceCrud) SetData() error {
	if s.Res.AggregationIntervalInDays != nil {
		s.D.Set("aggregation_interval_in_days", *s.Res.AggregationIntervalInDays)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LevelsConfiguration != nil {
		s.D.Set("levels_configuration", []interface{}{LevelsConfigurationToMap(s.Res.LevelsConfiguration)})
	} else {
		s.D.Set("levels_configuration", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetCompartments != nil {
		s.D.Set("target_compartments", []interface{}{TargetCompartmentsToMap(s.Res.TargetCompartments)})
	} else {
		s.D.Set("target_compartments", nil)
	}

	if s.Res.TargetTags != nil {
		s.D.Set("target_tags", []interface{}{TargetTagsToMap(s.Res.TargetTags)})
	} else {
		s.D.Set("target_tags", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *OptimizerProfileResourceCrud) mapToLevelConfiguration(fieldKeyFormat string) (oci_optimizer.LevelConfiguration, error) {
	result := oci_optimizer.LevelConfiguration{}

	if level, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "level")); ok {
		tmp := level.(string)
		result.Level = &tmp
	}

	if recommendationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recommendation_id")); ok {
		tmp := recommendationId.(string)
		result.RecommendationId = &tmp
	}

	return result, nil
}

func LevelConfigurationToMap(obj oci_optimizer.LevelConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Level != nil {
		result["level"] = string(*obj.Level)
	}

	if obj.RecommendationId != nil {
		result["recommendation_id"] = string(*obj.RecommendationId)
	}

	return result
}

func (s *OptimizerProfileResourceCrud) mapToLevelsConfiguration(fieldKeyFormat string) (oci_optimizer.LevelsConfiguration, error) {
	result := oci_optimizer.LevelsConfiguration{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_optimizer.LevelConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToLevelConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func LevelsConfigurationToMap(obj *oci_optimizer.LevelsConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, LevelConfigurationToMap(item))
	}
	result["items"] = items

	return result
}

func ProfileSummaryToMap(obj oci_optimizer.ProfileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregationIntervalInDays != nil {
		result["aggregation_interval_in_days"] = int(*obj.AggregationIntervalInDays)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LevelsConfiguration != nil {
		result["levels_configuration"] = []interface{}{LevelsConfigurationToMap(obj.LevelsConfiguration)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetCompartments != nil {
		result["target_compartments"] = []interface{}{TargetCompartmentsToMap(obj.TargetCompartments)}
	}

	if obj.TargetTags != nil {
		result["target_tags"] = []interface{}{TargetTagsToMap(obj.TargetTags)}
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *OptimizerProfileResourceCrud) mapToTargetCompartments(fieldKeyFormat string) (oci_optimizer.TargetCompartments, error) {
	result := oci_optimizer.TargetCompartments{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func TargetCompartmentsToMap(obj *oci_optimizer.TargetCompartments) map[string]interface{} {
	result := map[string]interface{}{}

	result["items"] = obj.Items

	return result
}

func (s *OptimizerProfileResourceCrud) mapToTargetTag(fieldKeyFormat string) (oci_optimizer.TargetTag, error) {
	result := oci_optimizer.TargetTag{}

	if tagDefinitionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_definition_name")); ok {
		tmp := tagDefinitionName.(string)
		result.TagDefinitionName = &tmp
	}

	if tagNamespaceName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_namespace_name")); ok {
		tmp := tagNamespaceName.(string)
		result.TagNamespaceName = &tmp
	}

	if tagValueType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_value_type")); ok {
		result.TagValueType = oci_optimizer.TagValueTypeEnum(tagValueType.(string))
	}

	if tagValues, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_values")); ok {
		interfaces := tagValues.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tag_values")) {
			result.TagValues = tmp
		}
	}

	return result, nil
}

func TargetTagToMap(obj oci_optimizer.TargetTag) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TagDefinitionName != nil {
		result["tag_definition_name"] = string(*obj.TagDefinitionName)
	}

	if obj.TagNamespaceName != nil {
		result["tag_namespace_name"] = string(*obj.TagNamespaceName)
	}

	result["tag_value_type"] = string(obj.TagValueType)

	result["tag_values"] = obj.TagValues

	return result
}

func (s *OptimizerProfileResourceCrud) mapToTargetTags(fieldKeyFormat string) (oci_optimizer.TargetTags, error) {
	result := oci_optimizer.TargetTags{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_optimizer.TargetTag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToTargetTag(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func TargetTagsToMap(obj *oci_optimizer.TargetTags) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, TargetTagToMap(item))
	}
	result["items"] = items

	return result
}
