// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubDynamicSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createOsManagementHubDynamicSetWithContext,
		ReadContext:   readOsManagementHubDynamicSetWithContext,
		UpdateContext: updateOsManagementHubDynamicSetWithContext,
		DeleteContext: deleteOsManagementHubDynamicSetWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"match_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"matching_rule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"architectures": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"display_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"is_reboot_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"locations": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"managed_instance_group_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"managed_instance_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"os_families": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"os_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"statuses": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tags": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DEFINED",
											"FREEFORM",
										}, true),
									},

									// Optional
									"key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value": {
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
			"target_compartments": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"does_include_children": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
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
			"scheduled_job_count": {
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

func createOsManagementHubDynamicSetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readOsManagementHubDynamicSetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateOsManagementHubDynamicSetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteOsManagementHubDynamicSetWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OsManagementHubDynamicSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DynamicSetClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type OsManagementHubDynamicSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.DynamicSetClient
	Res                    *oci_os_management_hub.DynamicSet
	DisableNotFoundRetries bool
}

func (s *OsManagementHubDynamicSetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubDynamicSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.DynamicSetLifecycleStateCreating),
	}
}

func (s *OsManagementHubDynamicSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.DynamicSetLifecycleStateActive),
	}
}

func (s *OsManagementHubDynamicSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.DynamicSetLifecycleStateDeleting),
	}
}

func (s *OsManagementHubDynamicSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.DynamicSetLifecycleStateDeleted),
	}
}

func (s *OsManagementHubDynamicSetResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_os_management_hub.CreateDynamicSetRequest{}

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

	if matchType, ok := s.D.GetOkExists("match_type"); ok {
		request.MatchType = oci_os_management_hub.MatchTypeEnum(matchType.(string))
	}

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		if tmpList := matchingRule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "matching_rule", 0)
			tmp, err := s.mapToMatchingRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MatchingRule = &tmp
		}
	}

	if targetCompartments, ok := s.D.GetOkExists("target_compartments"); ok {
		interfaces := targetCompartments.([]interface{})
		tmp := make([]oci_os_management_hub.TargetCompartmentDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_compartments", stateDataIndex)
			converted, err := s.mapToTargetCompartmentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_compartments") {
			request.TargetCompartments = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateDynamicSet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicSet
	return nil
}

func (s *OsManagementHubDynamicSetResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_os_management_hub.GetDynamicSetRequest{}

	tmp := s.D.Id()
	request.DynamicSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetDynamicSet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicSet
	return nil
}

func (s *OsManagementHubDynamicSetResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateDynamicSetRequest{}

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

	tmp := s.D.Id()
	request.DynamicSetId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if matchType, ok := s.D.GetOkExists("match_type"); ok {
		request.MatchType = oci_os_management_hub.MatchTypeEnum(matchType.(string))
	}

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		if tmpList := matchingRule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "matching_rule", 0)
			tmp, err := s.mapToMatchingRule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MatchingRule = &tmp
		}
	}

	if targetCompartments, ok := s.D.GetOkExists("target_compartments"); ok {
		interfaces := targetCompartments.([]interface{})
		tmp := make([]oci_os_management_hub.TargetCompartmentDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "target_compartments", stateDataIndex)
			converted, err := s.mapToTargetCompartmentDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("target_compartments") {
			request.TargetCompartments = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateDynamicSet(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicSet
	return nil
}

func (s *OsManagementHubDynamicSetResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_os_management_hub.DeleteDynamicSetRequest{}

	tmp := s.D.Id()
	request.DynamicSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteDynamicSet(ctx, request)
	return err
}

func (s *OsManagementHubDynamicSetResourceCrud) SetData() error {
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

	s.D.Set("match_type", s.Res.MatchType)

	if s.Res.MatchingRule != nil {
		s.D.Set("matching_rule", []interface{}{MatchingRuleToMap(s.Res.MatchingRule)})
	} else {
		s.D.Set("matching_rule", nil)
	}

	if s.Res.ScheduledJobCount != nil {
		s.D.Set("scheduled_job_count", *s.Res.ScheduledJobCount)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	targetCompartments := []interface{}{}
	for _, item := range s.Res.TargetCompartments {
		targetCompartments = append(targetCompartments, TargetCompartmentDetailsToMap(item))
	}
	s.D.Set("target_compartments", targetCompartments)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func DynamicSetSummaryToMap(obj oci_os_management_hub.DynamicSetSummary) map[string]interface{} {
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

func (s *OsManagementHubDynamicSetResourceCrud) mapToMatchingRule(fieldKeyFormat string) (oci_os_management_hub.MatchingRule, error) {
	result := oci_os_management_hub.MatchingRule{}

	if architectures, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "architectures")); ok {
		interfaces := architectures.([]interface{})
		tmp := make([]oci_os_management_hub.CpuArchTypeEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.CpuArchTypeEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "architectures")) {
			result.Architectures = tmp
		}
	}

	if displayNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_names")); ok {
		interfaces := displayNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "display_names")) {
			result.DisplayNames = tmp
		}
	}

	if isRebootRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_reboot_required")); ok {
		tmp := isRebootRequired.(bool)
		result.IsRebootRequired = &tmp
	}

	if locations, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "locations")); ok {
		interfaces := locations.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "locations")) {
			result.Locations = tmp
		}
	}

	if managedInstanceGroupIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_group_ids")); ok {
		interfaces := managedInstanceGroupIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "managed_instance_group_ids")) {
			result.ManagedInstanceGroupIds = tmp
		}
	}

	if managedInstanceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")); ok {
		interfaces := managedInstanceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")) {
			result.ManagedInstanceIds = tmp
		}
	}

	if osFamilies, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "os_families")); ok {
		interfaces := osFamilies.([]interface{})
		tmp := make([]oci_os_management_hub.OsFamilyEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.OsFamilyEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "os_families")) {
			result.OsFamilies = tmp
		}
	}

	if osNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "os_names")); ok {
		interfaces := osNames.([]interface{})
		tmp := make([]oci_os_management_hub.OsNameEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.OsNameEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "os_names")) {
			result.OsNames = tmp
		}
	}

	if statuses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "statuses")); ok {
		interfaces := statuses.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceStatusEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceStatusEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "statuses")) {
			result.Statuses = tmp
		}
	}

	if tags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tags")); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_os_management_hub.Tag, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "tags"), stateDataIndex)
			converted, err := s.mapToTag(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "tags")) {
			result.Tags = tmp
		}
	}

	return result, nil
}

func MatchingRuleToMap(obj *oci_os_management_hub.MatchingRule) map[string]interface{} {
	result := map[string]interface{}{}

	result["architectures"] = obj.Architectures

	result["display_names"] = obj.DisplayNames

	if obj.IsRebootRequired != nil {
		result["is_reboot_required"] = bool(*obj.IsRebootRequired)
	}

	result["locations"] = obj.Locations

	result["managed_instance_group_ids"] = obj.ManagedInstanceGroupIds

	result["managed_instance_ids"] = obj.ManagedInstanceIds

	result["os_families"] = obj.OsFamilies

	result["os_names"] = obj.OsNames

	result["statuses"] = obj.Statuses

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, TagToMap(item))
	}
	result["tags"] = tags

	return result
}

func (s *OsManagementHubDynamicSetResourceCrud) mapToTag(fieldKeyFormat string) (oci_os_management_hub.Tag, error) {
	var baseObject oci_os_management_hub.Tag
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFINED"):
		details := oci_os_management_hub.DefinedTag{}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	case strings.ToLower("FREEFORM"):
		details := oci_os_management_hub.FreeFormTag{}
		if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
			tmp := key.(string)
			details.Key = &tmp
		}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func TagToMap(obj oci_os_management_hub.Tag) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_os_management_hub.DefinedTag:
		result["type"] = "DEFINED"

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	case oci_os_management_hub.FreeFormTag:
		result["type"] = "FREEFORM"

		if v.Key != nil {
			result["key"] = string(*v.Key)
		}

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *OsManagementHubDynamicSetResourceCrud) mapToTargetCompartmentDetails(fieldKeyFormat string) (oci_os_management_hub.TargetCompartmentDetails, error) {
	result := oci_os_management_hub.TargetCompartmentDetails{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if doesIncludeChildren, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "does_include_children")); ok {
		tmp := doesIncludeChildren.(bool)
		result.DoesIncludeChildren = &tmp
	}

	return result, nil
}

func TargetCompartmentDetailsToMap(obj oci_os_management_hub.TargetCompartmentDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DoesIncludeChildren != nil {
		result["does_include_children"] = bool(*obj.DoesIncludeChildren)
	}

	return result
}

func (s *OsManagementHubDynamicSetResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeDynamicSetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DynamicSetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeDynamicSetCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
