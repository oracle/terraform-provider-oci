// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_guard

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_cloud_guard "github.com/oracle/oci-go-sdk/v58/cloudguard"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

func CloudGuardDataMaskRuleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudGuardDataMaskRule,
		Read:     readCloudGuardDataMaskRule,
		Update:   updateCloudGuardDataMaskRule,
		Delete:   deleteCloudGuardDataMaskRule,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_mask_categories": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iam_group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"target_selected": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ALL",
								"TARGETIDS",
								"TARGETTYPES",
							}, true),
						},

						// Optional
						"values": {
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
			"data_mask_rule_status": {
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
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecyle_details": {
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

func createCloudGuardDataMaskRule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudGuardDataMaskRule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.ReadResource(sync)
}

func updateCloudGuardDataMaskRule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudGuardDataMaskRule(d *schema.ResourceData, m interface{}) error {
	sync := &CloudGuardDataMaskRuleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).CloudGuardClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudGuardDataMaskRuleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_guard.CloudGuardClient
	Res                    *oci_cloud_guard.DataMaskRule
	DisableNotFoundRetries bool
}

func (s *CloudGuardDataMaskRuleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudGuardDataMaskRuleResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateCreating),
	}
}

func (s *CloudGuardDataMaskRuleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateActive),
	}
}

func (s *CloudGuardDataMaskRuleResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleting),
	}
}

func (s *CloudGuardDataMaskRuleResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_guard.LifecycleStateDeleted),
	}
}

func (s *CloudGuardDataMaskRuleResourceCrud) Create() error {
	request := oci_cloud_guard.CreateDataMaskRuleRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataMaskCategories, ok := s.D.GetOkExists("data_mask_categories"); ok {
		interfaces := dataMaskCategories.([]interface{})
		tmp := make([]oci_cloud_guard.DataMaskCategoryEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_cloud_guard.DataMaskCategoryEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("data_mask_categories") {
			request.DataMaskCategories = tmp
		}
	}

	if dataMaskRuleStatus, ok := s.D.GetOkExists("data_mask_rule_status"); ok {
		request.DataMaskRuleStatus = oci_cloud_guard.DataMaskRuleStatusEnum(dataMaskRuleStatus.(string))
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

	if iamGroupId, ok := s.D.GetOkExists("iam_group_id"); ok {
		tmp := iamGroupId.(string)
		request.IamGroupId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_cloud_guard.LifecycleStateEnum(state.(string))
	}

	if targetSelected, ok := s.D.GetOkExists("target_selected"); ok {
		if tmpList := targetSelected.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_selected", 0)
			tmp, err := s.mapToTargetSelected(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetSelected = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.CreateDataMaskRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataMaskRule
	return nil
}

func (s *CloudGuardDataMaskRuleResourceCrud) Get() error {
	request := oci_cloud_guard.GetDataMaskRuleRequest{}

	tmp := s.D.Id()
	request.DataMaskRuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.GetDataMaskRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataMaskRule
	return nil
}

func (s *CloudGuardDataMaskRuleResourceCrud) Update() error {
	request := oci_cloud_guard.UpdateDataMaskRuleRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataMaskCategories, ok := s.D.GetOkExists("data_mask_categories"); ok {
		interfaces := dataMaskCategories.([]interface{})
		tmp := make([]oci_cloud_guard.DataMaskCategoryEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_cloud_guard.DataMaskCategoryEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("data_mask_categories") {
			request.DataMaskCategories = tmp
		}
	}

	tmp := s.D.Id()
	request.DataMaskRuleId = &tmp

	if dataMaskRuleStatus, ok := s.D.GetOkExists("data_mask_rule_status"); ok {
		request.DataMaskRuleStatus = oci_cloud_guard.DataMaskRuleStatusEnum(dataMaskRuleStatus.(string))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if iamGroupId, ok := s.D.GetOkExists("iam_group_id"); ok {
		tmp := iamGroupId.(string)
		request.IamGroupId = &tmp
	}

	if targetSelected, ok := s.D.GetOkExists("target_selected"); ok {
		if tmpList := targetSelected.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_selected", 0)
			tmp, err := s.mapToTargetSelected(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TargetSelected = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	response, err := s.Client.UpdateDataMaskRule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataMaskRule
	return nil
}

func (s *CloudGuardDataMaskRuleResourceCrud) Delete() error {
	request := oci_cloud_guard.DeleteDataMaskRuleRequest{}

	tmp := s.D.Id()
	request.DataMaskRuleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_guard")

	_, err := s.Client.DeleteDataMaskRule(context.Background(), request)
	return err
}

func (s *CloudGuardDataMaskRuleResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("data_mask_categories", s.Res.DataMaskCategories)

	s.D.Set("data_mask_rule_status", s.Res.DataMaskRuleStatus)

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

	if s.Res.IamGroupId != nil {
		s.D.Set("iam_group_id", *s.Res.IamGroupId)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetSelected != nil {
		targetSelectedArray := []interface{}{}
		if targetSelectedMap := TargetSelectedToMap(&s.Res.TargetSelected); targetSelectedMap != nil {
			targetSelectedArray = append(targetSelectedArray, targetSelectedMap)
		}
		s.D.Set("target_selected", targetSelectedArray)
	} else {
		s.D.Set("target_selected", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func DataMaskRuleSummaryToMap(obj oci_cloud_guard.DataMaskRuleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["data_mask_categories"] = obj.DataMaskCategories

	result["data_mask_rule_status"] = string(obj.DataMaskRuleStatus)

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

	if obj.IamGroupId != nil {
		result["iam_group_id"] = string(*obj.IamGroupId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecyleDetails != nil {
		result["lifecyle_details"] = string(*obj.LifecyleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetSelected != nil {
		targetSelectedArray := []interface{}{}
		if targetSelectedMap := TargetSelectedToMap(&obj.TargetSelected); targetSelectedMap != nil {
			targetSelectedArray = append(targetSelectedArray, targetSelectedMap)
		}
		result["target_selected"] = targetSelectedArray
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *CloudGuardDataMaskRuleResourceCrud) mapToTargetSelected(fieldKeyFormat string) (oci_cloud_guard.TargetSelected, error) {
	var baseObject oci_cloud_guard.TargetSelected
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("ALL"):
		details := oci_cloud_guard.AllTargetsSelected{}
		baseObject = details
	case strings.ToLower("TARGETIDS"):
		details := oci_cloud_guard.TargetIdsSelected{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = tmp
			}
		}
		baseObject = details
	case strings.ToLower("TARGETTYPES"):
		details := oci_cloud_guard.TargetResourceTypesSelected{}
		if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
			interfaces := values.([]interface{})
			tmp := make([]oci_cloud_guard.TargetResourceTypeEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_cloud_guard.TargetResourceTypeEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "values")) {
				details.Values = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func TargetSelectedToMap(obj *oci_cloud_guard.TargetSelected) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_cloud_guard.AllTargetsSelected:
		result["kind"] = "ALL"
	case oci_cloud_guard.TargetIdsSelected:
		result["kind"] = "TARGETIDS"

		result["values"] = v.Values
	case oci_cloud_guard.TargetResourceTypesSelected:
		result["kind"] = "TARGETTYPES"

		result["values"] = v.Values
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", *obj)
		return nil
	}

	return result
}
