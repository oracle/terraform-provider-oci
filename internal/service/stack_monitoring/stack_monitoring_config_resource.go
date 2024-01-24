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
					"LICENSE_AUTO_ASSIGN",
					"LICENSE_ENTERPRISE_EXTENSIBILITY",
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
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
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
	case oci_stack_monitoring.LicenseAutoAssignConfigSummary:
		result["config_type"] = "LICENSE_AUTO_ASSIGN"

		result["license"] = string(v.License)
	case oci_stack_monitoring.LicenseEnterpriseExtensibilityConfigSummary:
		result["config_type"] = "LICENSE_ENTERPRISE_EXTENSIBILITY"

		if v.IsEnabled != nil {
			result["is_enabled"] = bool(*v.IsEnabled)
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", obj)
		return nil
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
