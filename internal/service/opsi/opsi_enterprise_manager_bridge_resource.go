// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v58/opsi"
)

func OpsiEnterpriseManagerBridgeResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiEnterpriseManagerBridge,
		Read:     readOpsiEnterpriseManagerBridge,
		Update:   updateOpsiEnterpriseManagerBridge,
		Delete:   deleteOpsiEnterpriseManagerBridge,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"object_storage_bucket_name": {
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
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_storage_bucket_status_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_storage_namespace_name": {
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

func createOpsiEnterpriseManagerBridge(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiEnterpriseManagerBridge(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiEnterpriseManagerBridge(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiEnterpriseManagerBridge(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiEnterpriseManagerBridgeResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiEnterpriseManagerBridgeResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.EnterpriseManagerBridge
	DisableNotFoundRetries bool
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	}
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) Create() error {
	request := oci_opsi.CreateEnterpriseManagerBridgeRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if objectStorageBucketName, ok := s.D.GetOkExists("object_storage_bucket_name"); ok {
		tmp := objectStorageBucketName.(string)
		request.ObjectStorageBucketName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateEnterpriseManagerBridge(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEnterpriseManagerBridgeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) getEnterpriseManagerBridgeFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	enterpriseManagerBridgeId, err := enterpriseManagerBridgeWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*enterpriseManagerBridgeId)

	return s.Get()
}

func enterpriseManagerBridgeWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "opsi", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_opsi.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func enterpriseManagerBridgeWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = enterpriseManagerBridgeWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_opsi.OperationStatusInProgress),
			string(oci_opsi.OperationStatusAccepted),
			string(oci_opsi.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_opsi.OperationStatusSucceeded),
			string(oci_opsi.OperationStatusFailed),
			string(oci_opsi.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_opsi.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_opsi.OperationStatusFailed || response.Status == oci_opsi.OperationStatusCanceled {
		return nil, getErrorFromOpsiEnterpriseManagerBridgeWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiEnterpriseManagerBridgeWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_opsi.ListWorkRequestErrorsRequest{
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

func (s *OpsiEnterpriseManagerBridgeResourceCrud) Get() error {
	request := oci_opsi.GetEnterpriseManagerBridgeRequest{}

	tmp := s.D.Id()
	request.EnterpriseManagerBridgeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetEnterpriseManagerBridge(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EnterpriseManagerBridge
	return nil
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateEnterpriseManagerBridgeRequest{}

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
	request.EnterpriseManagerBridgeId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateEnterpriseManagerBridge(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEnterpriseManagerBridgeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) Delete() error {
	request := oci_opsi.DeleteEnterpriseManagerBridgeRequest{}

	tmp := s.D.Id()
	request.EnterpriseManagerBridgeId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteEnterpriseManagerBridge(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := enterpriseManagerBridgeWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiEnterpriseManagerBridgeResourceCrud) SetData() error {
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectStorageBucketName != nil {
		s.D.Set("object_storage_bucket_name", *s.Res.ObjectStorageBucketName)
	}

	if s.Res.ObjectStorageBucketStatusDetails != nil {
		s.D.Set("object_storage_bucket_status_details", *s.Res.ObjectStorageBucketStatusDetails)
	}

	if s.Res.ObjectStorageNamespaceName != nil {
		s.D.Set("object_storage_namespace_name", *s.Res.ObjectStorageNamespaceName)
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

func EnterpriseManagerBridgeSummaryToMap(obj oci_opsi.EnterpriseManagerBridgeSummary) map[string]interface{} {
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

	if obj.ObjectStorageBucketName != nil {
		result["object_storage_bucket_name"] = string(*obj.ObjectStorageBucketName)
	}

	if obj.ObjectStorageBucketStatusDetails != nil {
		result["object_storage_bucket_status_details"] = string(*obj.ObjectStorageBucketStatusDetails)
	}

	if obj.ObjectStorageNamespaceName != nil {
		result["object_storage_namespace_name"] = string(*obj.ObjectStorageNamespaceName)
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

func (s *OpsiEnterpriseManagerBridgeResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeEnterpriseManagerBridgeCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EnterpriseManagerBridgeId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeEnterpriseManagerBridgeCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEnterpriseManagerBridgeFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
