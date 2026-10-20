// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAwsKeyPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDbmulticloudOracleDbAwsKeyPoolWithContext,
		ReadContext:   readDbmulticloudOracleDbAwsKeyPoolWithContext,
		UpdateContext: updateDbmulticloudOracleDbAwsKeyPoolWithContext,
		DeleteContext: deleteDbmulticloudOracleDbAwsKeyPoolWithContext,
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
			"is_multi_region_key_pool": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_reuse_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_state_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"replication_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"replication_state": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_region": {
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDbmulticloudOracleDbAwsKeyPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDbmulticloudOracleDbAwsKeyPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDbmulticloudOracleDbAwsKeyPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDbmulticloudOracleDbAwsKeyPoolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DbmulticloudOracleDbAwsKeyPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res                    *oci_dbmulticloud.OracleDbAwsKeyPool
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyPoolLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.CreateOracleDbAwsKeyPoolRequest{}

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

	if isMultiRegionKeyPool, ok := s.D.GetOkExists("is_multi_region_key_pool"); ok {
		tmp := isMultiRegionKeyPool.(bool)
		request.IsMultiRegionKeyPool = &tmp
	}

	if isReuseAllowed, ok := s.D.GetOkExists("is_reuse_allowed"); ok {
		tmp := isReuseAllowed.(bool)
		request.IsReuseAllowed = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbAwsKeyPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOracleDbAwsKeyPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) getOracleDbAwsKeyPoolFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbAwsKeyPoolId, err := oracleDbAwsKeyPoolWaitForWorkRequest(ctx, workId, "oracledbawskeypool",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbAwsKeyPoolId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(ctx,
			oci_dbmulticloud.CancelWorkRequestRequest{
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
	s.D.SetId(*oracleDbAwsKeyPoolId)

	return s.GetWithContext(ctx)
}

func oracleDbAwsKeyPoolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dbmulticloud", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dbmulticloud.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func oracleDbAwsKeyPoolWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbAwsKeyPoolWorkRequestShouldRetryFunc(timeout)

	response := oci_dbmulticloud.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_dbmulticloud.OperationStatusInProgress),
			string(oci_dbmulticloud.OperationStatusAccepted),
			string(oci_dbmulticloud.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_dbmulticloud.OperationStatusSucceeded),
			string(oci_dbmulticloud.OperationStatusFailed),
			string(oci_dbmulticloud.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_dbmulticloud.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_dbmulticloud.OperationStatusFailed || response.Status == oci_dbmulticloud.OperationStatusCanceled {
		return nil, getErrorFromDbmulticloudOracleDbAwsKeyPoolWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbAwsKeyPoolWorkRequest(ctx context.Context, client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_dbmulticloud.ListWorkRequestErrorsRequest{
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

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.GetOracleDbAwsKeyPoolRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsKeyPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbAwsKeyPool(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbAwsKeyPool
	return nil
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) UpdateWithContext(ctx context.Context) error {

	if _, ok := s.D.GetOkExists("action"); ok && s.D.HasChange("action") {
		err := s.ReplicateOracleDbAwsKeyPool(ctx)
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbAwsKeyPoolRequest{}

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

	if isMultiRegionKeyPool, ok := s.D.GetOkExists("is_multi_region_key_pool"); ok {
		tmp := isMultiRegionKeyPool.(bool)
		request.IsMultiRegionKeyPool = &tmp
	}

	if isReuseAllowed, ok := s.D.GetOkExists("is_reuse_allowed"); ok {
		tmp := isReuseAllowed.(bool)
		request.IsReuseAllowed = &tmp
	}

	tmp := s.D.Id()
	request.OracleDbAwsKeyPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbAwsKeyPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsKeyPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.DeleteOracleDbAwsKeyPoolRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsKeyPoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbAwsKeyPool(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbAwsKeyPoolWaitForWorkRequest(ctx, workId, "oracledbawskeypool",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) SetData() error {
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

	if s.Res.IsMultiRegionKeyPool != nil {
		s.D.Set("is_multi_region_key_pool", *s.Res.IsMultiRegionKeyPool)
	}

	if s.Res.IsReuseAllowed != nil {
		s.D.Set("is_reuse_allowed", *s.Res.IsReuseAllowed)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.ReplicationMetadata != nil {
		s.D.Set("replication_metadata", []interface{}{ReplicationMetadataToMap(s.Res.ReplicationMetadata)})
	} else {
		s.D.Set("replication_metadata", nil)
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

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) ReplicateOracleDbAwsKeyPool(ctx context.Context) error {
	request := oci_dbmulticloud.ReplicateOracleDbAwsKeyPoolRequest{}

	if action, ok := s.D.GetOkExists("action"); ok {
		request.Action = oci_dbmulticloud.ReplicationActionsEnum(action.(string))
	}

	idTmp := s.D.Id()
	request.OracleDbAwsKeyPoolId = &idTmp

	if targetRegion, ok := s.D.GetOkExists("target_region"); ok {
		tmp := targetRegion.(string)
		request.TargetRegion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	_, err := s.Client.ReplicateOracleDbAwsKeyPool(ctx, request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func OracleDbAwsKeyPoolSummaryToMap(obj oci_dbmulticloud.OracleDbAwsKeyPoolSummary) map[string]interface{} {
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

	if obj.IsMultiRegionKeyPool != nil {
		result["is_multi_region_key_pool"] = bool(*obj.IsMultiRegionKeyPool)
	}

	if obj.IsReuseAllowed != nil {
		result["is_reuse_allowed"] = bool(*obj.IsReuseAllowed)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.ReplicationMetadata != nil {
		result["replication_metadata"] = []interface{}{ReplicationMetadataToMap(obj.ReplicationMetadata)}
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

func ReplicationDetailsToMap(obj oci_dbmulticloud.ReplicationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["replication_state"] = string(obj.ReplicationState)

	if obj.TargetRegion != nil {
		result["target_region"] = string(*obj.TargetRegion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func ReplicationMetadataToMap(obj *oci_dbmulticloud.ReplicationMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	replicationDetails := []interface{}{}
	for _, item := range obj.ReplicationDetails {
		replicationDetails = append(replicationDetails, ReplicationDetailsToMap(item))
	}
	result["replication_details"] = replicationDetails

	return result
}

func (s *DbmulticloudOracleDbAwsKeyPoolResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbAwsKeyPoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbAwsKeyPoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbAwsKeyPoolCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsKeyPoolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
