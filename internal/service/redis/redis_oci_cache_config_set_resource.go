// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisOciCacheConfigSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRedisOciCacheConfigSet,
		Read:     readRedisOciCacheConfigSet,
		Update:   updateRedisOciCacheConfigSet,
		Delete:   deleteRedisOciCacheConfigSet,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configuration_details": {
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
									"config_value": {
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
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_version": {
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

			// Computed
			"default_config_set_id": {
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

func createRedisOciCacheConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()
	return tfresource.CreateResource(d, sync)
}

func readRedisOciCacheConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()
	return tfresource.ReadResource(sync)
}

func updateRedisOciCacheConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()
	return tfresource.UpdateResource(d, sync)
}

func deleteRedisOciCacheConfigSet(d *schema.ResourceData, m interface{}) error {
	sync := &RedisOciCacheConfigSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheConfigSetClient()
	sync.DisableNotFoundRetries = true
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()
	return tfresource.DeleteResource(d, sync)
}

type RedisOciCacheConfigSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.OciCacheConfigSetClient
	Res                    *oci_redis.OciCacheConfigSet
	RedisClusterClient     *oci_redis.RedisClusterClient
	DisableNotFoundRetries bool
}

func (s *RedisOciCacheConfigSetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RedisOciCacheConfigSetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_redis.OciCacheConfigSetLifecycleStateCreating),
	}
}

func (s *RedisOciCacheConfigSetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_redis.OciCacheConfigSetLifecycleStateActive),
	}
}

func (s *RedisOciCacheConfigSetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_redis.OciCacheConfigSetLifecycleStateDeleting),
	}
}

func (s *RedisOciCacheConfigSetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_redis.OciCacheConfigSetLifecycleStateDeleted),
	}
}

func (s *RedisOciCacheConfigSetResourceCrud) Create() error {
	request := oci_redis.CreateOciCacheConfigSetRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationDetails, ok := s.D.GetOkExists("configuration_details"); ok {
		if tmpList := configurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_details", 0)
			tmp, err := s.mapToConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationDetails = &tmp
		}
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

	if softwareVersion, ok := s.D.GetOkExists("software_version"); ok {
		request.SoftwareVersion = oci_redis.OciCacheConfigSetSoftwareVersionEnum(softwareVersion.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.CreateOciCacheConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_redis.GetWorkRequestResponse{}
	workRequestResponse, err = s.RedisClusterClient.GetWorkRequest(context.Background(),
		oci_redis.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "ocicacheconfigset") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getOciCacheConfigSetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RedisOciCacheConfigSetResourceCrud) getOciCacheConfigSetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_redis.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	ociCacheConfigSetId, err := ociCacheConfigSetWaitForWorkRequest(workId, "ocicacheconfigset",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.RedisClusterClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, ociCacheConfigSetId)
		_, cancelErr := s.RedisClusterClient.CancelWorkRequest(context.Background(),
			oci_redis.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*ociCacheConfigSetId)

	return s.Get()
}

func ociCacheConfigSetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "redis", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_redis.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func ociCacheConfigSetWaitForWorkRequest(wId *string, entityType string, action oci_redis.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_redis.RedisClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "redis")
	retryPolicy.ShouldRetryOperation = ociCacheConfigSetWorkRequestShouldRetryFunc(timeout)

	response := oci_redis.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_redis.OperationStatusInProgress),
			string(oci_redis.OperationStatusAccepted),
			string(oci_redis.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_redis.OperationStatusSucceeded),
			string(oci_redis.OperationStatusFailed),
			string(oci_redis.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_redis.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_redis.OperationStatusFailed || response.Status == oci_redis.OperationStatusCanceled {
		return nil, getErrorFromRedisOciCacheConfigSetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRedisOciCacheConfigSetWorkRequest(client *oci_redis.RedisClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_redis.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_redis.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *RedisOciCacheConfigSetResourceCrud) Get() error {
	request := oci_redis.GetOciCacheConfigSetRequest{}

	tmp := s.D.Id()
	request.OciCacheConfigSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.GetOciCacheConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OciCacheConfigSet
	return nil
}

func (s *RedisOciCacheConfigSetResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_redis.UpdateOciCacheConfigSetRequest{}

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

	tmp := s.D.Id()
	request.OciCacheConfigSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.UpdateOciCacheConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOciCacheConfigSetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RedisOciCacheConfigSetResourceCrud) Delete() error {
	request := oci_redis.DeleteOciCacheConfigSetRequest{}

	tmp := s.D.Id()
	request.OciCacheConfigSetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.DeleteOciCacheConfigSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := ociCacheConfigSetWaitForWorkRequest(workId, "ocicacheconfigset",
		oci_redis.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.RedisClusterClient)
	return delWorkRequestErr
}

func (s *RedisOciCacheConfigSetResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationDetails != nil {
		s.D.Set("configuration_details", []interface{}{ConfigurationDetailsToMap(s.Res.ConfigurationDetails)})
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.DefaultConfigSetId != nil {
		s.D.Set("default_config_set_id", *s.Res.DefaultConfigSetId)
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

	s.D.Set("software_version", s.Res.SoftwareVersion)

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

func (s *RedisOciCacheConfigSetResourceCrud) mapToConfigurationDetails(fieldKeyFormat string) (oci_redis.ConfigurationDetails, error) {
	result := oci_redis.ConfigurationDetails{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_redis.ConfigurationInfo, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToConfigurationInfo(fieldKeyFormatNextLevel)
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

func ConfigurationDetailsToMap(obj *oci_redis.ConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ConfigurationInfoToMap(item))
	}
	result["items"] = items

	return result
}

func (s *RedisOciCacheConfigSetResourceCrud) mapToConfigurationInfo(fieldKeyFormat string) (oci_redis.ConfigurationInfo, error) {
	result := oci_redis.ConfigurationInfo{}

	if configKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_key")); ok {
		tmp := configKey.(string)
		result.ConfigKey = &tmp
	}

	if configValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_value")); ok {
		tmp := configValue.(string)
		result.ConfigValue = &tmp
	}

	return result, nil
}

func ConfigurationInfoToMap(obj oci_redis.ConfigurationInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigKey != nil {
		result["config_key"] = string(*obj.ConfigKey)
	}

	if obj.ConfigValue != nil {
		result["config_value"] = string(*obj.ConfigValue)
	}

	return result
}

func OciCacheConfigSetSummaryToMap(obj oci_redis.OciCacheConfigSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefaultConfigSetId != nil {
		result["default_config_set_id"] = string(*obj.DefaultConfigSetId)
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

	result["software_version"] = string(obj.SoftwareVersion)

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

func (s *RedisOciCacheConfigSetResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_redis.ChangeOciCacheConfigSetCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OciCacheConfigSetId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ChangeOciCacheConfigSetCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOciCacheConfigSetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
