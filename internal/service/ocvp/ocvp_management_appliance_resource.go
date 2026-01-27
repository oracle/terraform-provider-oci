// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpManagementApplianceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOcvpManagementAppliance,
		Read:     readOcvpManagementAppliance,
		Update:   updateOcvpManagementAppliance,
		Delete:   deleteOcvpManagementAppliance,
		Schema: map[string]*schema.Schema{
			// Required
			"configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"is_log_ingestion_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"is_metrics_collection_enabled": {
							Type:     schema.TypeBool,
							Required: true,
						},

						// Optional
						"metrics": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"support_bundle_bucket_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"connections": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"credentials_secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
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
			"sddc_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"public_ssh_keys": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compute_instance_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heartbeat_connection_states": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_agent_id": {
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
			"time_configuration_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_heartbeat": {
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

func createOcvpManagementAppliance(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpManagementApplianceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementApplianceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readOcvpManagementAppliance(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpManagementApplianceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementApplianceClient()

	return tfresource.ReadResource(sync)
}

func updateOcvpManagementAppliance(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpManagementApplianceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementApplianceClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOcvpManagementAppliance(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpManagementApplianceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementApplianceClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type OcvpManagementApplianceResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.ManagementApplianceClient
	Res                    *oci_ocvp.ManagementAppliance
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpManagementApplianceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpManagementApplianceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.ManagementApplianceLifecycleStateCreating),
	}
}

func (s *OcvpManagementApplianceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.ManagementApplianceLifecycleStateActive),
		string(oci_ocvp.ManagementApplianceLifecycleStateNeedsAttention),
	}
}

func (s *OcvpManagementApplianceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.ManagementApplianceLifecycleStateDeleting),
	}
}

func (s *OcvpManagementApplianceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.ManagementApplianceLifecycleStateDeleted),
	}
}

func (s *OcvpManagementApplianceResourceCrud) Create() error {
	request := oci_ocvp.CreateManagementApplianceRequest{}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", 0)
			tmp, err := s.mapToManagementApplianceConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Configuration = &tmp
		}
	}

	if connections, ok := s.D.GetOkExists("connections"); ok {
		interfaces := connections.([]interface{})
		tmp := make([]oci_ocvp.ManagementApplianceConnection, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connections", stateDataIndex)
			converted, err := s.mapToManagementApplianceConnection(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("connections") {
			request.Connections = tmp
		}
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

	if publicSshKeys, ok := s.D.GetOkExists("public_ssh_keys"); ok {
		tmp := publicSshKeys.(string)
		request.PublicSshKeys = &tmp
	}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateManagementAppliance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "managementappliance") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getManagementApplianceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpManagementApplianceResourceCrud) getManagementApplianceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	managementApplianceId, err := managementApplianceWaitForWorkRequest(workId, "managementappliance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*managementApplianceId)

	return s.Get()
}

func managementApplianceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func managementApplianceWaitForWorkRequest(wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = managementApplianceWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpManagementApplianceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpManagementApplianceWorkRequest(client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpManagementApplianceResourceCrud) Get() error {
	request := oci_ocvp.GetManagementApplianceRequest{}

	tmp := s.D.Id()
	request.ManagementApplianceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetManagementAppliance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ManagementAppliance
	return nil
}

func (s *OcvpManagementApplianceResourceCrud) Update() error {
	request := oci_ocvp.UpdateManagementApplianceRequest{}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", 0)
			tmp, err := s.mapToManagementApplianceConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Configuration = &tmp
		}
	}

	if connections, ok := s.D.GetOkExists("connections"); ok {
		interfaces := connections.([]interface{})
		tmp := make([]oci_ocvp.ManagementApplianceConnection, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "connections", stateDataIndex)
			converted, err := s.mapToManagementApplianceConnection(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("connections") {
			request.Connections = tmp
		}
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

	tmp := s.D.Id()
	request.ManagementApplianceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateManagementAppliance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getManagementApplianceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OcvpManagementApplianceResourceCrud) Delete() error {
	request := oci_ocvp.DeleteManagementApplianceRequest{}

	tmp := s.D.Id()
	request.ManagementApplianceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteManagementAppliance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := managementApplianceWaitForWorkRequest(workId, "managementappliance",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpManagementApplianceResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeInstanceId != nil {
		s.D.Set("compute_instance_id", *s.Res.ComputeInstanceId)
	}

	if s.Res.Configuration != nil {
		s.D.Set("configuration", []interface{}{ManagementApplianceConfigurationToMap(s.Res.Configuration)})
	} else {
		s.D.Set("configuration", nil)
	}

	connections := []interface{}{}
	for _, item := range s.Res.Connections {
		connections = append(connections, ManagementApplianceConnectionToMap(item))
	}
	s.D.Set("connections", connections)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	heartbeatConnectionStates := []interface{}{}
	for _, item := range s.Res.HeartbeatConnectionStates {
		heartbeatConnectionStates = append(heartbeatConnectionStates, ManagementApplianceConnectionStatusToMap(item))
	}
	s.D.Set("heartbeat_connection_states", heartbeatConnectionStates)

	s.D.Set("lifecycle_details", s.Res.LifecycleDetails)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeConfigurationUpdated != nil {
		s.D.Set("time_configuration_updated", s.Res.TimeConfigurationUpdated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastHeartbeat != nil {
		s.D.Set("time_last_heartbeat", s.Res.TimeLastHeartbeat.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *OcvpManagementApplianceResourceCrud) mapToManagementApplianceConfiguration(fieldKeyFormat string) (oci_ocvp.ManagementApplianceConfiguration, error) {
	result := oci_ocvp.ManagementApplianceConfiguration{}

	if isLogIngestionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_log_ingestion_enabled")); ok {
		tmp := isLogIngestionEnabled.(bool)
		result.IsLogIngestionEnabled = &tmp
	}

	if isMetricsCollectionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_metrics_collection_enabled")); ok {
		tmp := isMetricsCollectionEnabled.(bool)
		result.IsMetricsCollectionEnabled = &tmp
	}

	if metrics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metrics")); ok {
		interfaces := metrics.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metrics")) {
			result.Metrics = tmp
		}
	}

	if supportBundleBucketId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "support_bundle_bucket_id")); ok {
		tmp := supportBundleBucketId.(string)
		result.SupportBundleBucketId = &tmp
	}

	return result, nil
}

func ManagementApplianceConfigurationToMap(obj *oci_ocvp.ManagementApplianceConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsLogIngestionEnabled != nil {
		result["is_log_ingestion_enabled"] = bool(*obj.IsLogIngestionEnabled)
	}

	if obj.IsMetricsCollectionEnabled != nil {
		result["is_metrics_collection_enabled"] = bool(*obj.IsMetricsCollectionEnabled)
	}

	result["metrics"] = obj.Metrics

	if obj.SupportBundleBucketId != nil {
		result["support_bundle_bucket_id"] = string(*obj.SupportBundleBucketId)
	}

	return result
}

func (s *OcvpManagementApplianceResourceCrud) mapToManagementApplianceConnection(fieldKeyFormat string) (oci_ocvp.ManagementApplianceConnection, error) {
	result := oci_ocvp.ManagementApplianceConnection{}

	if credentialsSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials_secret_id")); ok {
		tmp := credentialsSecretId.(string)
		result.CredentialsSecretId = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_ocvp.ManagementApplianceConnectionTypeEnum(type_.(string))
	}

	return result, nil
}

func ManagementApplianceConnectionToMap(obj oci_ocvp.ManagementApplianceConnection) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CredentialsSecretId != nil {
		result["credentials_secret_id"] = string(*obj.CredentialsSecretId)
	}

	result["type"] = string(obj.Type)

	return result
}

func ManagementApplianceConnectionStatusToMap(obj oci_ocvp.ManagementApplianceConnectionStatus) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Details != nil {
		result["details"] = string(*obj.Details)
	}

	result["state"] = string(obj.State)

	result["type"] = string(obj.Type)

	return result
}

func ManagementApplianceSummaryToMap(obj oci_ocvp.ManagementApplianceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeInstanceId != nil {
		result["compute_instance_id"] = string(*obj.ComputeInstanceId)
	}

	if obj.Configuration != nil {
		result["configuration"] = []interface{}{ManagementApplianceConfigurationToMap(obj.Configuration)}
	}

	connections := []interface{}{}
	for _, item := range obj.Connections {
		connections = append(connections, ManagementApplianceConnectionToMap(item))
	}
	result["connections"] = connections

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	heartbeatConnectionStates := []interface{}{}
	for _, item := range obj.HeartbeatConnectionStates {
		heartbeatConnectionStates = append(heartbeatConnectionStates, ManagementApplianceConnectionStatusToMap(item))
	}
	result["heartbeat_connection_states"] = heartbeatConnectionStates

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["lifecycle_details"] = string(obj.LifecycleDetails)

	if obj.ManagementAgentId != nil {
		result["management_agent_id"] = string(*obj.ManagementAgentId)
	}

	if obj.SddcId != nil {
		result["sddc_id"] = string(*obj.SddcId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeConfigurationUpdated != nil {
		result["time_configuration_updated"] = obj.TimeConfigurationUpdated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastHeartbeat != nil {
		result["time_last_heartbeat"] = obj.TimeLastHeartbeat.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
