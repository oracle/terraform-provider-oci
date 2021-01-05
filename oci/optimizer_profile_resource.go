// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_optimizer "github.com/oracle/oci-go-sdk/v31/optimizer"
)

func init() {
	RegisterResource("oci_optimizer_profile", OptimizerProfileResource())
}

func OptimizerProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
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
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
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
	sync.Client = m.(*OracleClients).optimizerClient()

	return CreateResource(d, sync)
}

func readOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).optimizerClient()

	return ReadResource(sync)
}

func updateOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).optimizerClient()

	return UpdateResource(d, sync)
}

func deleteOptimizerProfile(d *schema.ResourceData, m interface{}) error {
	sync := &OptimizerProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).optimizerClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type OptimizerProfileResourceCrud struct {
	BaseCrud
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

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "optimizer")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	response, err := s.Client.GetProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Profile
	return nil
}

func (s *OptimizerProfileResourceCrud) Update() error {
	request := oci_optimizer.UpdateProfileRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
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

	tmp := s.D.Id()
	request.ProfileId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "optimizer")

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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "optimizer")

	_, err := s.Client.DeleteProfile(context.Background(), request)
	return err
}

func (s *OptimizerProfileResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
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

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
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

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
