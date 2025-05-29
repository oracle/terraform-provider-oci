// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package managed_kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_managed_kafka "github.com/oracle/oci-go-sdk/v65/managedkafka"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagedKafkaKafkaClusterConfigResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagedKafkaKafkaClusterConfig,
		Read:     readManagedKafkaKafkaClusterConfig,
		Update:   updateManagedKafkaKafkaClusterConfig,
		Delete:   deleteManagedKafkaKafkaClusterConfig,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"latest_config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"properties": {
							Type:     schema.TypeMap,
							Required: true,
							Elem:     schema.TypeString,
						},

						// Optional
						"config_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version_number": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
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
			"lifecycle_details": {
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

func createManagedKafkaKafkaClusterConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.CreateResource(d, sync)
}

func readManagedKafkaKafkaClusterConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.ReadResource(sync)
}

func updateManagedKafkaKafkaClusterConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagedKafkaKafkaClusterConfig(d *schema.ResourceData, m interface{}) error {
	sync := &ManagedKafkaKafkaClusterConfigResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KafkaClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagedKafkaKafkaClusterConfigResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_managed_kafka.KafkaClusterClient
	Res                    *oci_managed_kafka.KafkaClusterConfig
	DisableNotFoundRetries bool
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterConfigLifecycleStateCreating),
	}
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterConfigLifecycleStateActive),
	}
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_managed_kafka.KafkaClusterConfigLifecycleStateDeleted),
	}
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) Create() error {
	request := oci_managed_kafka.CreateKafkaClusterConfigRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if latestConfig, ok := s.D.GetOkExists("latest_config"); ok {
		if tmpList := latestConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "latest_config", 0)
			tmp, err := s.mapToKafkaClusterConfigVersion(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LatestConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.CreateKafkaClusterConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KafkaClusterConfig
	return nil
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) Get() error {
	request := oci_managed_kafka.GetKafkaClusterConfigRequest{}

	tmp := s.D.Id()
	request.KafkaClusterConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.GetKafkaClusterConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KafkaClusterConfig
	return nil
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_managed_kafka.UpdateKafkaClusterConfigRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.KafkaClusterConfigId = &tmp

	if latestConfig, ok := s.D.GetOkExists("latest_config"); ok {
		if tmpList := latestConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "latest_config", 0)
			tmp, err := s.mapToKafkaClusterConfigVersion(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LatestConfig = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	response, err := s.Client.UpdateKafkaClusterConfig(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KafkaClusterConfig
	return nil
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) Delete() error {
	request := oci_managed_kafka.DeleteKafkaClusterConfigRequest{}

	tmp := s.D.Id()
	request.KafkaClusterConfigId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	_, err := s.Client.DeleteKafkaClusterConfig(context.Background(), request)
	return err
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LatestConfig != nil {
		s.D.Set("latest_config", []interface{}{KafkaClusterConfigVersionToMap(s.Res.LatestConfig)})
	} else {
		s.D.Set("latest_config", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
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

func KafkaClusterConfigSummaryToMap(obj oci_managed_kafka.KafkaClusterConfigSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
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

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) mapToKafkaClusterConfigVersion(fieldKeyFormat string) (oci_managed_kafka.KafkaClusterConfigVersion, error) {
	result := oci_managed_kafka.KafkaClusterConfigVersion{}

	if configId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_id")); ok {
		tmp := configId.(string)
		result.ConfigId = &tmp
	}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		result.Properties = tfresource.ObjectMapToStringMap(properties.(map[string]interface{}))
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if versionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version_number")); ok {
		tmp := versionNumber.(int)
		result.VersionNumber = &tmp
	}

	return result, nil
}

func KafkaClusterConfigVersionToMap(obj *oci_managed_kafka.KafkaClusterConfigVersion) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigId != nil {
		result["config_id"] = string(*obj.ConfigId)
	}

	result["properties"] = obj.Properties

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.VersionNumber != nil {
		result["version_number"] = int(*obj.VersionNumber)
	}

	return result
}

func (s *ManagedKafkaKafkaClusterConfigResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_managed_kafka.ChangeKafkaClusterConfigCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KafkaClusterConfigId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "managed_kafka")

	_, err := s.Client.ChangeKafkaClusterConfigCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
