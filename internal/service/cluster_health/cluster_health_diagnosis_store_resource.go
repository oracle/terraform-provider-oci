// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cluster_health

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cluster_health "github.com/oracle/oci-go-sdk/v65/clusterhealth"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ClusterHealthDiagnosisStoreResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createClusterHealthDiagnosisStoreWithContext,
		ReadContext:   readClusterHealthDiagnosisStoreWithContext,
		UpdateContext: updateClusterHealthDiagnosisStoreWithContext,
		DeleteContext: deleteClusterHealthDiagnosisStoreWithContext,
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
			"object_store_bucket": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_store_namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Computed
			"last_accepted_request_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_completed_request_token": {
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

func createClusterHealthDiagnosisStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readClusterHealthDiagnosisStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateClusterHealthDiagnosisStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteClusterHealthDiagnosisStoreWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &ClusterHealthDiagnosisStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DiagnosisClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type ClusterHealthDiagnosisStoreResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cluster_health.DiagnosisClient
	Res                    *oci_cluster_health.DiagnosisStore
	DisableNotFoundRetries bool
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cluster_health.DiagnosisStoreLifecycleStateCreating),
	}
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cluster_health.DiagnosisStoreLifecycleStateActive),
	}
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cluster_health.DiagnosisStoreLifecycleStateDeleting),
	}
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cluster_health.DiagnosisStoreLifecycleStateDeleted),
	}
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_cluster_health.CreateDiagnosisStoreRequest{}

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

	if objectStoreBucket, ok := s.D.GetOkExists("object_store_bucket"); ok {
		tmp := objectStoreBucket.(string)
		request.ObjectStoreBucket = &tmp
	}

	if objectStoreNamespace, ok := s.D.GetOkExists("object_store_namespace"); ok {
		tmp := objectStoreNamespace.(string)
		request.ObjectStoreNamespace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health")

	response, err := s.Client.CreateDiagnosisStore(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDiagnosisStoreFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health"), oci_cluster_health.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) getDiagnosisStoreFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cluster_health.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	diagnosisStoreId, err := diagnosisStoreWaitForWorkRequest(ctx, workId, "diagnosisstore",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, diagnosisStoreId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_cluster_health.CancelWorkRequestRequest{
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
	s.D.SetId(*diagnosisStoreId)

	return s.GetWithContext(ctx)
}

func diagnosisStoreWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cluster_health", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cluster_health.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func diagnosisStoreWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_cluster_health.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cluster_health.DiagnosisClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cluster_health")
	retryPolicy.ShouldRetryOperation = diagnosisStoreWorkRequestShouldRetryFunc(timeout)

	response := oci_cluster_health.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_cluster_health.OperationStatusInProgress),
			string(oci_cluster_health.OperationStatusAccepted),
			string(oci_cluster_health.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cluster_health.OperationStatusSucceeded),
			string(oci_cluster_health.OperationStatusFailed),
			string(oci_cluster_health.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_cluster_health.GetWorkRequestRequest{
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
	if _, e := stateConf.WaitForStateContext(ctx); e != nil {
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
	if identifier == nil || response.Status == oci_cluster_health.OperationStatusFailed || response.Status == oci_cluster_health.OperationStatusCanceled {
		return nil, getErrorFromClusterHealthDiagnosisStoreWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromClusterHealthDiagnosisStoreWorkRequest(ctx context.Context, client *oci_cluster_health.DiagnosisClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cluster_health.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_cluster_health.ListWorkRequestErrorsRequest{
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

func (s *ClusterHealthDiagnosisStoreResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_cluster_health.GetDiagnosisStoreRequest{}

	tmp := s.D.Id()
	request.DiagnosisStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health")

	response, err := s.Client.GetDiagnosisStore(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.DiagnosisStore
	return nil
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cluster_health.UpdateDiagnosisStoreRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	tmp := s.D.Id()
	request.DiagnosisStoreId = &tmp

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if objectStoreBucket, ok := s.D.GetOkExists("object_store_bucket"); ok {
		tmp := objectStoreBucket.(string)
		request.ObjectStoreBucket = &tmp
	}

	if objectStoreNamespace, ok := s.D.GetOkExists("object_store_namespace"); ok {
		tmp := objectStoreNamespace.(string)
		request.ObjectStoreNamespace = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health")

	response, err := s.Client.UpdateDiagnosisStore(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDiagnosisStoreFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health"), oci_cluster_health.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_cluster_health.DeleteDiagnosisStoreRequest{}

	tmp := s.D.Id()
	request.DiagnosisStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health")

	response, err := s.Client.DeleteDiagnosisStore(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := diagnosisStoreWaitForWorkRequest(ctx, workId, "diagnosisstore",
		oci_cluster_health.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ClusterHealthDiagnosisStoreResourceCrud) SetData() error {
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

	if s.Res.LastAcceptedRequestToken != nil {
		s.D.Set("last_accepted_request_token", *s.Res.LastAcceptedRequestToken)
	}

	if s.Res.LastCompletedRequestToken != nil {
		s.D.Set("last_completed_request_token", *s.Res.LastCompletedRequestToken)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ObjectStoreBucket != nil {
		s.D.Set("object_store_bucket", *s.Res.ObjectStoreBucket)
	}

	if s.Res.ObjectStoreNamespace != nil {
		s.D.Set("object_store_namespace", *s.Res.ObjectStoreNamespace)
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

func DiagnosisStoreSummaryToMap(obj oci_cluster_health.DiagnosisStoreSummary) map[string]interface{} {
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

	if obj.ObjectStoreBucket != nil {
		result["object_store_bucket"] = string(*obj.ObjectStoreBucket)
	}

	if obj.ObjectStoreNamespace != nil {
		result["object_store_namespace"] = string(*obj.ObjectStoreNamespace)
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

func (s *ClusterHealthDiagnosisStoreResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_cluster_health.ChangeDiagnosisStoreCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.DiagnosisStoreId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health")

	response, err := s.Client.ChangeDiagnosisStoreCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDiagnosisStoreFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cluster_health"), oci_cluster_health.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
