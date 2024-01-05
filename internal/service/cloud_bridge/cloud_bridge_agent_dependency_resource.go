// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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

func CloudBridgeAgentDependencyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeAgentDependency,
		Read:     readCloudBridgeAgentDependency,
		Update:   updateCloudBridgeAgentDependency,
		Delete:   deleteCloudBridgeAgentDependency,
		Schema: map[string]*schema.Schema{
			// Required
			"bucket": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dependency_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"dependency_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"e_tag": {
				Type:     schema.TypeString,
				Computed: true,
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

func createCloudBridgeAgentDependency(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDependencyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeAgentDependency(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDependencyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeAgentDependency(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDependencyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeAgentDependency(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeAgentDependencyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeAgentDependencyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.OcbAgentSvcClient
	WorkRequestClient      *oci_cloud_bridge.CommonClient
	Res                    *oci_cloud_bridge.AgentDependency
	DisableNotFoundRetries bool
}

func (s *CloudBridgeAgentDependencyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudBridgeAgentDependencyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_bridge.AgentDependencyLifecycleStateCreating),
	}
}

func (s *CloudBridgeAgentDependencyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AgentDependencyLifecycleStateActive),
	}
}

func (s *CloudBridgeAgentDependencyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_bridge.AgentDependencyLifecycleStateDeleting),
	}
}

func (s *CloudBridgeAgentDependencyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.AgentDependencyLifecycleStateDeleted),
	}
}

func (s *CloudBridgeAgentDependencyResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateAgentDependencyRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.Bucket = &tmp
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

	if dependencyName, ok := s.D.GetOkExists("dependency_name"); ok {
		tmp := dependencyName.(string)
		request.DependencyName = &tmp
	}

	if dependencyVersion, ok := s.D.GetOkExists("dependency_version"); ok {
		tmp := dependencyVersion.(string)
		request.DependencyVersion = &tmp
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

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateAgentDependency(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_cloud_bridge.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(context.Background(),
		oci_cloud_bridge.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "agentdependency") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAgentDependencyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudBridgeAgentDependencyResourceCrud) getAgentDependencyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_bridge.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	agentDependencyId, err := agentDependencyWaitForWorkRequest(workId, "agentdependency",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, agentDependencyId)
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
	s.D.SetId(*agentDependencyId)

	return s.Get()
}

func agentDependencyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func agentDependencyWaitForWorkRequest(wId *string, entityType string, action oci_cloud_bridge.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_bridge.CommonClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_bridge")
	retryPolicy.ShouldRetryOperation = agentDependencyWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromCloudBridgeAgentDependencyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudBridgeAgentDependencyWorkRequest(client *oci_cloud_bridge.CommonClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_bridge.ActionTypeEnum) error {
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

func (s *CloudBridgeAgentDependencyResourceCrud) Get() error {
	request := oci_cloud_bridge.GetAgentDependencyRequest{}

	tmp := s.D.Id()
	request.AgentDependencyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetAgentDependency(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AgentDependency
	return nil
}

func (s *CloudBridgeAgentDependencyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateAgentDependencyRequest{}

	tmp := s.D.Id()
	request.AgentDependencyId = &tmp

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.Bucket = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if dependencyName, ok := s.D.GetOkExists("dependency_name"); ok {
		tmp := dependencyName.(string)
		request.DependencyName = &tmp
	}

	if dependencyVersion, ok := s.D.GetOkExists("dependency_version"); ok {
		tmp := dependencyVersion.(string)
		request.DependencyVersion = &tmp
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

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.Namespace = &tmp
	}

	if object, ok := s.D.GetOkExists("object"); ok {
		tmp := object.(string)
		request.ObjectName = &tmp
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdateAgentDependency(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAgentDependencyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *CloudBridgeAgentDependencyResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteAgentDependencyRequest{}

	tmp := s.D.Id()
	request.AgentDependencyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.DeleteAgentDependency(context.Background(), request)
	return err
}

func (s *CloudBridgeAgentDependencyResourceCrud) SetData() error {
	if s.Res.Bucket != nil {
		s.D.Set("bucket", *s.Res.Bucket)
	}

	if s.Res.Checksum != nil {
		s.D.Set("checksum", *s.Res.Checksum)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DependencyName != nil {
		s.D.Set("dependency_name", *s.Res.DependencyName)
	}

	if s.Res.DependencyVersion != nil {
		s.D.Set("dependency_version", *s.Res.DependencyVersion)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ETag != nil {
		s.D.Set("e_tag", *s.Res.ETag)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ObjectName != nil {
		s.D.Set("object", *s.Res.ObjectName)
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

func AgentDependencySummaryToMap(obj oci_cloud_bridge.AgentDependencySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Bucket != nil {
		result["bucket"] = string(*obj.Bucket)
	}

	if obj.Checksum != nil {
		result["checksum"] = string(*obj.Checksum)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DependencyName != nil {
		result["dependency_name"] = string(*obj.DependencyName)
	}

	if obj.DependencyVersion != nil {
		result["dependency_version"] = string(*obj.DependencyVersion)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.ETag != nil {
		result["e_tag"] = string(*obj.ETag)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
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

func (s *CloudBridgeAgentDependencyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeAgentDependencyCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AgentDependencyId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.ChangeAgentDependencyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
