// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_bridge

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudBridgeAgentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeAgent,
		Read:     readCloudBridgeAgent,
		Update:   updateCloudBridgeAgent,
		Delete:   deleteCloudBridgeAgent,
		Schema: map[string]*schema.Schema{
			// Required
			"agent_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"agent_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"environment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"os_version": {
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

			// Computed
			"agent_pub_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"heart_beat_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plugin_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"agent_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"plugin_version": {
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
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_expire_agent_key_in_ms": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_sync_received": {
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

func createCloudBridgeAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeAgent(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeAgentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.OcbAgentSvcClient
	WorkRequestClient      *oci_cloud_bridge.CommonClient
	Res                    *oci_cloud_bridge.Agent
	DisableNotFoundRetries bool
}

func (s *CloudBridgeAgentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudBridgeAgentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_bridge.AgentLifecycleStateCreating),
	}
}

func (s *CloudBridgeAgentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AgentLifecycleStateActive),
	}
}

func (s *CloudBridgeAgentResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *CloudBridgeAgentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AgentLifecycleStateDeleted),
	}
}

func (s *CloudBridgeAgentResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateAgentRequest{}

	if agentType, ok := s.D.GetOkExists("agent_type"); ok {
		request.AgentType = oci_cloud_bridge.AgentAgentTypeEnum(agentType.(string))
	}

	if agentVersion, ok := s.D.GetOkExists("agent_version"); ok {
		tmp := agentVersion.(string)
		request.AgentVersion = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if environmentId, ok := s.D.GetOkExists("environment_id"); ok {
		tmp := environmentId.(string)
		request.EnvironmentId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if osVersion, ok := s.D.GetOkExists("os_version"); ok {
		tmp := osVersion.(string)
		request.OsVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *CloudBridgeAgentResourceCrud) getAgentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_bridge.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	agentId, err := agentWaitForWorkRequest(workId, "ocbagent",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, agentId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_cloud_bridge.CancelWorkRequestRequest{
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
	s.D.SetId(*agentId)

	return s.Get()
}

func agentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cloud_bridge", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cloud_bridge.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func agentWaitForWorkRequest(wId *string, entityType string, action oci_cloud_bridge.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_bridge.CommonClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_bridge")
	retryPolicy.ShouldRetryOperation = agentWorkRequestShouldRetryFunc(timeout)

	response := oci_cloud_bridge.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cloud_bridge.OperationStatusInProgress),
			string(oci_cloud_bridge.OperationStatusAccepted),
			string(oci_cloud_bridge.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cloud_bridge.OperationStatusSucceeded),
			string(oci_cloud_bridge.OperationStatusFailed),
			string(oci_cloud_bridge.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cloud_bridge.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_cloud_bridge.OperationStatusFailed || response.Status == oci_cloud_bridge.OperationStatusCanceled {
		return nil, getErrorFromCloudBridgeAgentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudBridgeAgentWorkRequest(client *oci_cloud_bridge.CommonClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_bridge.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cloud_bridge.ListWorkRequestErrorsRequest{
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

func (s *CloudBridgeAgentResourceCrud) Get() error {
	request := oci_cloud_bridge.GetAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *CloudBridgeAgentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdateAgent(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Agent
	return nil
}

func (s *CloudBridgeAgentResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteAgentRequest{}

	tmp := s.D.Id()
	request.AgentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.DeleteAgent(context.Background(), request)
	return err
}

func (s *CloudBridgeAgentResourceCrud) SetData() error {
	if s.Res.AgentPubKey != nil {
		s.D.Set("agent_pub_key", *s.Res.AgentPubKey)
	}

	s.D.Set("agent_type", s.Res.AgentType)

	if s.Res.AgentVersion != nil {
		s.D.Set("agent_version", *s.Res.AgentVersion)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EnvironmentId != nil {
		s.D.Set("environment_id", *s.Res.EnvironmentId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("heart_beat_status", s.Res.HeartBeatStatus)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OsVersion != nil {
		s.D.Set("os_version", *s.Res.OsVersion)
	}

	pluginList := []interface{}{}
	for _, item := range s.Res.PluginList {
		pluginList = append(pluginList, PluginSummaryToMap(item))
	}
	s.D.Set("plugin_list", pluginList)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeExpireAgentKeyInMs != nil {
		s.D.Set("time_expire_agent_key_in_ms", s.Res.TimeExpireAgentKeyInMs.String())
	}

	if s.Res.TimeLastSyncReceived != nil {
		s.D.Set("time_last_sync_received", s.Res.TimeLastSyncReceived.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func AgentSummaryToMap(obj oci_cloud_bridge.AgentSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["agent_type"] = string(obj.AgentType)

	if obj.AgentVersion != nil {
		result["agent_version"] = string(*obj.AgentVersion)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.EnvironmentId != nil {
		result["environment_id"] = string(*obj.EnvironmentId)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["heart_beat_status"] = string(obj.HeartBeatStatus)

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OsVersion != nil {
		result["os_version"] = string(*obj.OsVersion)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastSyncReceived != nil {
		result["time_last_sync_received"] = obj.TimeLastSyncReceived.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func PluginSummaryToMap(obj oci_cloud_bridge.PluginSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AgentId != nil {
		result["agent_id"] = string(*obj.AgentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.PluginVersion != nil {
		result["plugin_version"] = string(*obj.PluginVersion)
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

func (s *CloudBridgeAgentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeAgentCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AgentId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.ChangeAgentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
