// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createPsqlConfiguration,
		Read:     readPsqlConfiguration,
		Update:   updatePsqlConfiguration,
		Delete:   deletePsqlConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_configuration_overrides": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"config_key": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"overriden_config_value": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"db_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
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
			"instance_memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"instance_ocpu_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_flexible": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"config_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"allowed_values": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"config_key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"data_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"default_config_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_overridable": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"is_restart_required": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"overriden_config_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
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
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createPsqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.CreateResource(d, sync)
}

func readPsqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

func updatePsqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.UpdateResource(d, sync)
}

func deletePsqlConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type PsqlConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_psql.PostgresqlClient
	Res                    *oci_psql.Configuration
	DisableNotFoundRetries bool
}

func (s *PsqlConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PsqlConfigurationResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *PsqlConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_psql.ConfigurationLifecycleStateActive),
	}
}

func (s *PsqlConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_psql.ConfigurationLifecycleStateDeleting),
	}
}

func (s *PsqlConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_psql.ConfigurationLifecycleStateDeleted),
	}
}

func (s *PsqlConfigurationResourceCrud) Create() error {
	request := oci_psql.CreateConfigurationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbConfigurationOverrides, ok := s.D.GetOkExists("db_configuration_overrides"); ok {
		if tmpList := dbConfigurationOverrides.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "db_configuration_overrides", 0)
			tmp, err := s.mapToDbConfigurationOverrideCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DbConfigurationOverrides = &tmp
		}
	}

	if dbVersion, ok := s.D.GetOkExists("db_version"); ok {
		tmp := dbVersion.(string)
		request.DbVersion = &tmp
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

	if instanceMemorySizeInGBs, ok := s.D.GetOkExists("instance_memory_size_in_gbs"); ok {
		tmp := instanceMemorySizeInGBs.(int)
		request.InstanceMemorySizeInGBs = &tmp
	}

	if instanceOcpuCount, ok := s.D.GetOkExists("instance_ocpu_count"); ok {
		tmp := instanceOcpuCount.(int)
		request.InstanceOcpuCount = &tmp
	}

	if isFlexible, ok := s.D.GetOkExists("is_flexible"); ok {
		tmp := isFlexible.(bool)
		request.IsFlexible = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.Shape = &tmp
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.CreateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *PsqlConfigurationResourceCrud) Get() error {
	request := oci_psql.GetConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.GetConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *PsqlConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_psql.UpdateConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.UpdateConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Configuration
	return nil
}

func (s *PsqlConfigurationResourceCrud) Delete() error {
	request := oci_psql.DeleteConfigurationRequest{}

	tmp := s.D.Id()
	request.ConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	_, err := s.Client.DeleteConfiguration(context.Background(), request)
	return err
}

func (s *PsqlConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("config_type", s.Res.ConfigType)

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", []interface{}{ConfigurationDetailsToMap(s.Res.ConfigurationDetails)})
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.DbVersion != nil {
		s.D.Set("db_version", *s.Res.DbVersion)
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InstanceMemorySizeInGBs != nil {
		s.D.Set("instance_memory_size_in_gbs", *s.Res.InstanceMemorySizeInGBs)
	}

	if s.Res.InstanceOcpuCount != nil {
		s.D.Set("instance_ocpu_count", *s.Res.InstanceOcpuCount)
	}

	if s.Res.IsFlexible != nil {
		s.D.Set("is_flexible", *s.Res.IsFlexible)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", strings.TrimSuffix(*s.Res.Shape, "."+strconv.Itoa(*s.Res.InstanceOcpuCount)+"."+strconv.Itoa(*s.Res.InstanceMemorySizeInGBs)+"GB"))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *PsqlConfigurationResourceCrud) mapToConfigOverrides(fieldKeyFormat string) (oci_psql.ConfigOverrides, error) {
	result := oci_psql.ConfigOverrides{}

	if configKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_key")); ok {
		tmp := configKey.(string)
		result.ConfigKey = &tmp
	}

	if overridenConfigValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "overriden_config_value")); ok {
		tmp := overridenConfigValue.(string)
		result.OverridenConfigValue = &tmp
	}

	return result, nil
}

func ConfigOverridesToMap(obj oci_psql.ConfigOverrides) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.OverridenConfigValue != nil {
		result["overriden_config_value"] = string(*obj.OverridenConfigValue)
	}

	return result
}

func ConfigParamsToMap(obj oci_psql.ConfigParams) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowedValues != nil {
		result["allowed_values"] = string(*obj.AllowedValues)
	}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.DataType != nil {
		result["data_type"] = string(*obj.DataType)
	}

	if obj.DefaultConfigValue != nil {
		result["default_config_value"] = string(*obj.DefaultConfigValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.IsOverridable != nil {
		result["is_overridable"] = bool(*obj.IsOverridable)
	}

	if obj.IsRestartRequired != nil {
		result["is_restart_required"] = bool(*obj.IsRestartRequired)
	}

	if obj.OverridenConfigValue != nil {
		result["overriden_config_value"] = string(*obj.OverridenConfigValue)
	}

	return result
}

func ConfigurationDetailsToMap(obj *oci_psql.ConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ConfigParamsToMap(item))
	}
	result["items"] = items

	return result
}

func ConfigurationSummaryToMap(obj oci_psql.ConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.InstanceMemorySizeInGBs != nil {
		result["instance_memory_size_in_gbs"] = int(*obj.InstanceMemorySizeInGBs)
	}

	if obj.InstanceOcpuCount != nil {
		result["instance_ocpu_count"] = int(*obj.InstanceOcpuCount)
	}

	if obj.IsFlexible != nil {
		result["is_flexible"] = bool(*obj.IsFlexible)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	return result
}

func (s *PsqlConfigurationResourceCrud) mapToDbConfigurationOverrideCollection(fieldKeyFormat string) (oci_psql.DbConfigurationOverrideCollection, error) {
	result := oci_psql.DbConfigurationOverrideCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_psql.ConfigOverrides, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToConfigOverrides(fieldKeyFormatNextLevel)
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

func DbConfigurationOverrideCollectionToMap(obj *oci_psql.DbConfigurationOverrideCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ConfigOverridesToMap(item))
	}
	result["items"] = items

	return result
}

func (s *PsqlConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_psql.ChangeConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	_, err := s.Client.ChangeConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
