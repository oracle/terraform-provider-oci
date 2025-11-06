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

func DbmulticloudOracleDbAwsKeyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createDbmulticloudOracleDbAwsKeyWithContext,
		ReadContext:   readDbmulticloudOracleDbAwsKeyWithContext,
		UpdateContext: updateDbmulticloudOracleDbAwsKeyWithContext,
		DeleteContext: deleteDbmulticloudOracleDbAwsKeyWithContext,
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
			"oracle_db_connector_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"aws_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"aws_key_arn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
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
			"is_aws_key_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"properties": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"lifecycle_state_details": {
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

func createDbmulticloudOracleDbAwsKeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readDbmulticloudOracleDbAwsKeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateDbmulticloudOracleDbAwsKeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteDbmulticloudOracleDbAwsKeyWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsKeyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).DbmulticloudWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type DbmulticloudOracleDbAwsKeyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res                    *oci_dbmulticloud.OracleDbAwsKey
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_dbmulticloud.WorkRequestClient
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateCreating),
	}
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateActive),
	}
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateDeleting),
	}
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dbmulticloud.OracleDbAwsKeyLifecycleStateDeleted),
	}
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.CreateOracleDbAwsKeyRequest{}

	if awsAccountId, ok := s.D.GetOkExists("aws_account_id"); ok {
		tmp := awsAccountId.(string)
		request.AwsAccountId = &tmp
	}

	if awsKeyArn, ok := s.D.GetOkExists("aws_key_arn"); ok {
		tmp := awsKeyArn.(string)
		request.AwsKeyArn = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isAwsKeyEnabled, ok := s.D.GetOkExists("is_aws_key_enabled"); ok {
		tmp := isAwsKeyEnabled.(bool)
		request.IsAwsKeyEnabled = &tmp
	}

	if location, ok := s.D.GetOkExists("location"); ok {
		tmp := location.(string)
		request.Location = &tmp
	}

	if oracleDbConnectorId, ok := s.D.GetOkExists("oracle_db_connector_id"); ok {
		tmp := oracleDbConnectorId.(string)
		request.OracleDbConnectorId = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		request.Properties = tfresource.ObjectMapToStringMap(properties.(map[string]interface{}))
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.CreateOracleDbAwsKey(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOracleDbAwsKeyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) getOracleDbAwsKeyFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dbmulticloud.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	oracleDbAwsKeyId, err := oracleDbAwsKeyWaitForWorkRequest(ctx, workId, "oracledbawskey",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, oracleDbAwsKeyId)
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
	s.D.SetId(*oracleDbAwsKeyId)

	return s.GetWithContext(ctx)
}

func oracleDbAwsKeyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func oracleDbAwsKeyWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_dbmulticloud.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dbmulticloud.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dbmulticloud")
	retryPolicy.ShouldRetryOperation = oracleDbAwsKeyWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_dbmulticloud.OperationStatusFailed || response.Status == oci_dbmulticloud.OperationStatusCanceled {
		return nil, getErrorFromDbmulticloudOracleDbAwsKeyWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDbmulticloudOracleDbAwsKeyWorkRequest(ctx context.Context, client *oci_dbmulticloud.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dbmulticloud.ActionTypeEnum) error {
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

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.GetOracleDbAwsKeyRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.GetOracleDbAwsKey(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OracleDbAwsKey
	return nil
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dbmulticloud.UpdateOracleDbAwsKeyRequest{}

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
	request.OracleDbAwsKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.UpdateOracleDbAwsKey(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsKeyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.DeleteOracleDbAwsKeyRequest{}

	tmp := s.D.Id()
	request.OracleDbAwsKeyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.DeleteOracleDbAwsKey(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := oracleDbAwsKeyWaitForWorkRequest(ctx, workId, "oracledbawskey",
		oci_dbmulticloud.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) SetData() error {
	if s.Res.AwsAccountId != nil {
		s.D.Set("aws_account_id", *s.Res.AwsAccountId)
	}

	if s.Res.AwsKeyArn != nil {
		s.D.Set("aws_key_arn", *s.Res.AwsKeyArn)
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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAwsKeyEnabled != nil {
		s.D.Set("is_aws_key_enabled", *s.Res.IsAwsKeyEnabled)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.Location != nil {
		s.D.Set("location", *s.Res.Location)
	}

	if s.Res.OracleDbConnectorId != nil {
		s.D.Set("oracle_db_connector_id", *s.Res.OracleDbConnectorId)
	}

	s.D.Set("properties", s.Res.Properties)

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

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func OracleDbAwsKeySummaryToMap(obj oci_dbmulticloud.OracleDbAwsKeySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AwsAccountId != nil {
		result["aws_account_id"] = string(*obj.AwsAccountId)
	}

	if obj.AwsKeyArn != nil {
		result["aws_key_arn"] = string(*obj.AwsKeyArn)
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

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsAwsKeyEnabled != nil {
		result["is_aws_key_enabled"] = bool(*obj.IsAwsKeyEnabled)
	}

	if obj.LifecycleStateDetails != nil {
		result["lifecycle_state_details"] = string(*obj.LifecycleStateDetails)
	}

	if obj.Location != nil {
		result["location"] = string(*obj.Location)
	}

	if obj.OracleDbConnectorId != nil {
		result["oracle_db_connector_id"] = string(*obj.OracleDbConnectorId)
	}

	result["properties"] = obj.Properties

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

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DbmulticloudOracleDbAwsKeyResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_dbmulticloud.ChangeOracleDbAwsKeyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OracleDbAwsKeyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud")

	response, err := s.Client.ChangeOracleDbAwsKeyCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOracleDbAwsKeyFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dbmulticloud"), oci_dbmulticloud.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
