// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringConfig,
		Read:     readStackMonitoringConfig,
		Update:   updateStackMonitoringConfig,
		Delete:   deleteStackMonitoringConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"AUTO_PROMOTE",
					"COMPUTE_AUTO_ACTIVATE_PLUGIN",
					"LICENSE_AUTO_ASSIGN",
					"LICENSE_ENTERPRISE_EXTENSIBILITY",
					"ONBOARD",
				}, true),
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"license": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Optional
			"additional_configurations": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"properties_map": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     schema.TypeString,
						},

						// Computed
					},
				},
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stack_monitoring_assignment": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			"is_manually_onboarded": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"policy_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"user_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"stack_monitoring_role": {
							Type:     schema.TypeString,
							Optional: true,
						},

						// Computed
					},
				},
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
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

func createStackMonitoringConfig(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringConfig(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringConfig(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringConfig(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type StackMonitoringConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.Config
	DisableNotFoundRetries bool
}

func (s *StackMonitoringConfigResourceCrud) ID() string {
	config := *s.Res
	return *config.GetId()
}

func (s *StackMonitoringConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.ConfigLifecycleStateCreating),
	}
}

func (s *StackMonitoringConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ConfigLifecycleStateActive),
	}
}

func (s *StackMonitoringConfigResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_stack_monitoring.ConfigLifecycleStateDeleting),
	}
}

func (s *StackMonitoringConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_stack_monitoring.ConfigLifecycleStateDeleted),
	}
}

func (s *StackMonitoringConfigResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateConfigRequest{}
	err := s.populateTopLevelPolymorphicCreateConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *StackMonitoringConfigResourceCrud) Get() error {
	request := oci_stack_monitoring.GetConfigRequest{}

	tmp := s.D.Id()
	request.ConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *StackMonitoringConfigResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_stack_monitoring.UpdateConfigRequest{}
	err := s.populateTopLevelPolymorphicUpdateConfigRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Config
	return nil
}

func (s *StackMonitoringConfigResourceCrud) Delete() error {
	request := oci_stack_monitoring.DeleteConfigRequest{}

	tmp := s.D.Id()
	request.ConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.DeleteConfig(context.Background(), request)
	return err
}

func (s *StackMonitoringConfigResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_stack_monitoring.AutoPromoteConfigDetails:
		s.D.Set("config_type", "AUTO_PROMOTE")

		if v.IsEnabled != nil {
			s.D.Set("is_enabled", *v.IsEnabled)
		}

		s.D.Set("resource_type", v.ResourceType)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_stack_monitoring.ComputeAutoActivatePluginConfigDetails:
		s.D.Set("config_type", "COMPUTE_AUTO_ACTIVATE_PLUGIN")

		if v.IsEnabled != nil {
			s.D.Set("is_enabled", *v.IsEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_stack_monitoring.LicenseAutoAssignConfigDetails:
		s.D.Set("config_type", "LICENSE_AUTO_ASSIGN")

		s.D.Set("license", v.License)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_stack_monitoring.LicenseEnterpriseExtensibilityConfigDetails:
		s.D.Set("config_type", "LICENSE_ENTERPRISE_EXTENSIBILITY")

		if v.IsEnabled != nil {
			s.D.Set("is_enabled", *v.IsEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_stack_monitoring.OnboardConfigDetails:
		s.D.Set("config_type", "ONBOARD")

		if v.AdditionalConfigurations != nil {
			s.D.Set("additional_configurations", []interface{}{AdditionalConfigurationDetailsToMap(v.AdditionalConfigurations)})
		} else {
			s.D.Set("additional_configurations", nil)
		}

		dynamicGroups := []interface{}{}
		for _, item := range v.DynamicGroups {
			dynamicGroups = append(dynamicGroups, DynamicGroupDetailsToMap(item))
		}
		s.D.Set("dynamic_groups", dynamicGroups)

		if v.IsManuallyOnboarded != nil {
			s.D.Set("is_manually_onboarded", *v.IsManuallyOnboarded)
		}

		s.D.Set("policy_names", v.PolicyNames)

		userGroups := []interface{}{}
		for _, item := range v.UserGroups {
			userGroups = append(userGroups, GroupDetailsToMap(item))
		}
		s.D.Set("user_groups", userGroups)

		if v.Version != nil {
			s.D.Set("version", *v.Version)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *StackMonitoringConfigResourceCrud) mapToAdditionalConfigurationDetails(fieldKeyFormat string) (oci_stack_monitoring.AdditionalConfigurationDetails, error) {
	result := oci_stack_monitoring.AdditionalConfigurationDetails{}

	if propertiesMap, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties_map")); ok {
		result.PropertiesMap = tfresource.ObjectMapToStringMap(propertiesMap.(map[string]interface{}))
	}

	return result, nil
}

func AdditionalConfigurationDetailsToMap(obj *oci_stack_monitoring.AdditionalConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["properties_map"] = obj.PropertiesMap

	return result
}

func ConfigSummaryToMap(obj oci_stack_monitoring.ConfigSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = string(*obj.GetCompartmentId())
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = string(*obj.GetDisplayName())
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetId() != nil {
		result["id"] = string(*obj.GetId())
	}

	result["state"] = obj.GetLifecycleState()

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	switch v := (obj).(type) {
	case oci_stack_monitoring.AutoPromoteConfigSummary:
		result["config_type"] = "AUTO_PROMOTE"

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}

		result["resource_type"] = string(v.ResourceType)
	case oci_stack_monitoring.ComputeAutoActivatePluginConfigSummary:
		result["config_type"] = "COMPUTE_AUTO_ACTIVATE_PLUGIN"

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}
	case oci_stack_monitoring.LicenseAutoAssignConfigSummary:
		result["config_type"] = "LICENSE_AUTO_ASSIGN"

		result["license"] = string(v.License)
	case oci_stack_monitoring.LicenseEnterpriseExtensibilityConfigSummary:
		result["config_type"] = "LICENSE_ENTERPRISE_EXTENSIBILITY"

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}
	case oci_stack_monitoring.OnboardConfigSummary:
		result["config_type"] = "ONBOARD"

		if v.IsManuallyOnboarded != nil {
			result["is_manually_onboarded"] = bool(*v.IsManuallyOnboarded)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *StackMonitoringConfigResourceCrud) mapToDynamicGroupDetails(fieldKeyFormat string) (oci_stack_monitoring.DynamicGroupDetails, error) {
	result := oci_stack_monitoring.DynamicGroupDetails{}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		result.Domain = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if stackMonitoringAssignment, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stack_monitoring_assignment")); ok {
		result.StackMonitoringAssignment = oci_stack_monitoring.DynamicGroupDetailsStackMonitoringAssignmentEnum(stackMonitoringAssignment.(string))
	}

	return result, nil
}

func DynamicGroupDetailsToMap(obj oci_stack_monitoring.DynamicGroupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Domain != nil {
		result["domain"] = string(*obj.Domain)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["stack_monitoring_assignment"] = string(obj.StackMonitoringAssignment)

	return result
}

func (s *StackMonitoringConfigResourceCrud) mapToGroupDetails(fieldKeyFormat string) (oci_stack_monitoring.GroupDetails, error) {
	result := oci_stack_monitoring.GroupDetails{}

	if domain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain")); ok {
		tmp := domain.(string)
		result.Domain = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if stackMonitoringRole, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stack_monitoring_role")); ok {
		tmp := stackMonitoringRole.(string)
		result.StackMonitoringRole = &tmp
	}

	return result, nil
}

func GroupDetailsToMap(obj oci_stack_monitoring.GroupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Domain != nil {
		result["domain"] = string(*obj.Domain)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.StackMonitoringRole != nil {
		result["stack_monitoring_role"] = string(*obj.StackMonitoringRole)
	}

	return result
}

func (s *StackMonitoringConfigResourceCrud) populateTopLevelPolymorphicCreateConfigRequest(request *oci_stack_monitoring.CreateConfigRequest) error {
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists("config_type")
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("AUTO_PROMOTE"):
		details := oci_stack_monitoring.CreateAutoPromoteConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		if resourceType, ok := s.D.GetOkExists("resource_type"); ok {
			details.ResourceType = oci_stack_monitoring.CreateAutoPromoteConfigDetailsResourceTypeEnum(resourceType.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateConfigDetails = details
	case strings.ToLower("COMPUTE_AUTO_ACTIVATE_PLUGIN"):
		details := oci_stack_monitoring.CreateComputeAutoActivatePluginConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateConfigDetails = details
	case strings.ToLower("LICENSE_AUTO_ASSIGN"):
		details := oci_stack_monitoring.CreateLicenseAutoAssignConfigDetails{}
		if license, ok := s.D.GetOkExists("license"); ok {
			details.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateConfigDetails = details
	case strings.ToLower("LICENSE_ENTERPRISE_EXTENSIBILITY"):
		details := oci_stack_monitoring.CreateLicenseEnterpriseExtensibilityConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateConfigDetails = details
	case strings.ToLower("ONBOARD"):
		details := oci_stack_monitoring.CreateOnboardConfigDetails{}
		if additionalConfigurations, ok := s.D.GetOkExists("additional_configurations"); ok {
			if tmpList := additionalConfigurations.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_configurations", 0)
				tmp, err := s.mapToAdditionalConfigurationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdditionalConfigurations = &tmp
			}
		}
		if dynamicGroups, ok := s.D.GetOkExists("dynamic_groups"); ok {
			interfaces := dynamicGroups.([]interface{})
			tmp := make([]oci_stack_monitoring.DynamicGroupDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dynamic_groups", stateDataIndex)
				converted, err := s.mapToDynamicGroupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("dynamic_groups") {
				details.DynamicGroups = tmp
			}
		}
		if isManuallyOnboarded, ok := s.D.GetOkExists("is_manually_onboarded"); ok {
			tmp := isManuallyOnboarded.(bool)
			details.IsManuallyOnboarded = &tmp
		}
		if policyNames, ok := s.D.GetOkExists("policy_names"); ok {
			interfaces := policyNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("policy_names") {
				details.PolicyNames = tmp
			}
		}
		if userGroups, ok := s.D.GetOkExists("user_groups"); ok {
			interfaces := userGroups.([]interface{})
			tmp := make([]oci_stack_monitoring.GroupDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_groups", stateDataIndex)
				converted, err := s.mapToGroupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("user_groups") {
				details.UserGroups = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateConfigDetails = details
	default:
		return fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return nil
}

func (s *StackMonitoringConfigResourceCrud) populateTopLevelPolymorphicUpdateConfigRequest(request *oci_stack_monitoring.UpdateConfigRequest) error {
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists("config_type")
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("AUTO_PROMOTE"):
		details := oci_stack_monitoring.UpdateAutoPromoteConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		tmp := s.D.Id()
		request.ConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("COMPUTE_AUTO_ACTIVATE_PLUGIN"):
		details := oci_stack_monitoring.UpdateComputeAutoActivatePluginConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		tmp := s.D.Id()
		request.ConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("LICENSE_AUTO_ASSIGN"):
		details := oci_stack_monitoring.UpdateLicenseAutoAssignConfigDetails{}
		if license, ok := s.D.GetOkExists("license"); ok {
			details.License = oci_stack_monitoring.LicenseTypeEnum(license.(string))
		}
		tmp := s.D.Id()
		request.ConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("LICENSE_ENTERPRISE_EXTENSIBILITY"):
		details := oci_stack_monitoring.UpdateLicenseEnterpriseExtensibilityConfigDetails{}
		if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
			tmp := isEnabled.(bool)
			details.IsEnabled = &tmp
		}
		tmp := s.D.Id()
		request.ConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateConfigDetails = details
	case strings.ToLower("ONBOARD"):
		details := oci_stack_monitoring.UpdateOnboardConfigDetails{}
		if additionalConfigurations, ok := s.D.GetOkExists("additional_configurations"); ok {
			if tmpList := additionalConfigurations.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_configurations", 0)
				tmp, err := s.mapToAdditionalConfigurationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AdditionalConfigurations = &tmp
			}
		}
		if dynamicGroups, ok := s.D.GetOkExists("dynamic_groups"); ok {
			interfaces := dynamicGroups.([]interface{})
			tmp := make([]oci_stack_monitoring.DynamicGroupDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "dynamic_groups", stateDataIndex)
				converted, err := s.mapToDynamicGroupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("dynamic_groups") {
				details.DynamicGroups = tmp
			}
		}
		if isManuallyOnboarded, ok := s.D.GetOkExists("is_manually_onboarded"); ok {
			tmp := isManuallyOnboarded.(bool)
			details.IsManuallyOnboarded = &tmp
		}
		if policyNames, ok := s.D.GetOkExists("policy_names"); ok {
			interfaces := policyNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("policy_names") {
				details.PolicyNames = tmp
			}
		}
		if userGroups, ok := s.D.GetOkExists("user_groups"); ok {
			interfaces := userGroups.([]interface{})
			tmp := make([]oci_stack_monitoring.GroupDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "user_groups", stateDataIndex)
				converted, err := s.mapToGroupDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("user_groups") {
				details.UserGroups = tmp
			}
		}
		if version, ok := s.D.GetOkExists("version"); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		tmp := s.D.Id()
		request.ConfigId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateConfigDetails = details
	default:
		return fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return nil
}

func (s *StackMonitoringConfigResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_stack_monitoring.ChangeConfigCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConfigId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.ChangeConfigCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
