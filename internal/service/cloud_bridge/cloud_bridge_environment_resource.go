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

func CloudBridgeEnvironmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudBridgeEnvironment,
		Read:     readCloudBridgeEnvironment,
		Update:   updateCloudBridgeEnvironment,
		Delete:   deleteCloudBridgeEnvironment,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
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

func createCloudBridgeEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudBridgeEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.ReadResource(sync)
}

func updateCloudBridgeEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudBridgeEnvironment(d *schema.ResourceData, m interface{}) error {
	sync := &CloudBridgeEnvironmentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OcbAgentSvcClient()
	sync.WorkRequestClient = m.(*client.OracleClients).CommonClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudBridgeEnvironmentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_bridge.OcbAgentSvcClient
	WorkRequestClient      *oci_cloud_bridge.CommonClient
	Res                    *oci_cloud_bridge.Environment
	DisableNotFoundRetries bool
}

func (s *CloudBridgeEnvironmentResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudBridgeEnvironmentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_bridge.EnvironmentLifecycleStateCreating),
	}
}

func (s *CloudBridgeEnvironmentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_bridge.EnvironmentLifecycleStateActive),
	}
}

func (s *CloudBridgeEnvironmentResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_bridge.EnvironmentLifecycleStateDeleting),
	}
}

func (s *CloudBridgeEnvironmentResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_bridge.EnvironmentLifecycleStateDeleted),
	}
}

func (s *CloudBridgeEnvironmentResourceCrud) Create() error {
	request := oci_cloud_bridge.CreateEnvironmentRequest{}

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.CreateEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Environment
	return nil
}

func (s *CloudBridgeEnvironmentResourceCrud) getEnvironmentFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_bridge.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	environmentId, err := environmentWaitForWorkRequest(workId, "environment",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, environmentId)
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
	s.D.SetId(*environmentId)

	return s.Get()
}

func environmentWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func environmentWaitForWorkRequest(wId *string, entityType string, action oci_cloud_bridge.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_bridge.CommonClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_bridge")
	retryPolicy.ShouldRetryOperation = environmentWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromCloudBridgeEnvironmentWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudBridgeEnvironmentWorkRequest(client *oci_cloud_bridge.CommonClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_bridge.ActionTypeEnum) error {
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

func (s *CloudBridgeEnvironmentResourceCrud) Get() error {
	request := oci_cloud_bridge.GetEnvironmentRequest{}

	tmp := s.D.Id()
	request.EnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.GetEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Environment
	return nil
}

func (s *CloudBridgeEnvironmentResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_bridge.UpdateEnvironmentRequest{}

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

	tmp := s.D.Id()
	request.EnvironmentId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.UpdateEnvironment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Environment
	return nil
}

func (s *CloudBridgeEnvironmentResourceCrud) Delete() error {
	request := oci_cloud_bridge.DeleteEnvironmentRequest{}

	tmp := s.D.Id()
	request.EnvironmentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	_, err := s.Client.DeleteEnvironment(context.Background(), request)
	return err
}

func (s *CloudBridgeEnvironmentResourceCrud) SetData() error {
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

func EnvironmentSummaryToMap(obj oci_cloud_bridge.EnvironmentSummary) map[string]interface{} {
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

func (s *CloudBridgeEnvironmentResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_bridge.ChangeEnvironmentCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EnvironmentId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge")

	response, err := s.Client.ChangeEnvironmentCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEnvironmentFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_bridge"), oci_cloud_bridge.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
