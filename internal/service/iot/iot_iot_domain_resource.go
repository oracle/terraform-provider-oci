// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package iot

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_iot "github.com/oracle/oci-go-sdk/v65/iot"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IotIotDomainResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createIotIotDomainWithContext,
		ReadContext:   readIotIotDomainWithContext,
		UpdateContext: updateIotIotDomainWithContext,
		DeleteContext: deleteIotIotDomainWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iot_domain_group_id": {
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
			"data_retention_periods_in_days": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"historized_data": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"raw_command_data": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"raw_data": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rejected_data": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"db_allow_listed_identity_group_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"db_allowed_identity_domain_host": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"device_host": {
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

func createIotIotDomainWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readIotIotDomainWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateIotIotDomainWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteIotIotDomainWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &IotIotDomainResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IotClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type IotIotDomainResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_iot.IotClient
	Res                    *oci_iot.IotDomain
	DisableNotFoundRetries bool
}

func (s *IotIotDomainResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IotIotDomainResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_iot.IotDomainLifecycleStateCreating),
	}
}

func (s *IotIotDomainResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_iot.IotDomainLifecycleStateActive),
	}
}

func (s *IotIotDomainResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_iot.IotDomainLifecycleStateDeleting),
	}
}

func (s *IotIotDomainResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_iot.IotDomainLifecycleStateDeleted),
	}
}

func (s *IotIotDomainResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_iot.CreateIotDomainRequest{}

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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if iotDomainGroupId, ok := s.D.GetOkExists("iot_domain_group_id"); ok {
		tmp := iotDomainGroupId.(string)
		request.IotDomainGroupId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.CreateIotDomain(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getIotDomainFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *IotIotDomainResourceCrud) getIotDomainFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_iot.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	iotDomainId, err := iotDomainWaitForWorkRequest(ctx, workId, "iotdomain",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*iotDomainId)

	return s.GetWithContext(ctx)
}

func iotDomainWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "iot", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_iot.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func iotDomainWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_iot.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_iot.IotClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "iot")
	retryPolicy.ShouldRetryOperation = iotDomainWorkRequestShouldRetryFunc(timeout)

	response := oci_iot.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_iot.OperationStatusInProgress),
			string(oci_iot.OperationStatusAccepted),
			string(oci_iot.OperationStatusWaiting),
		},
		Target: []string{
			string(oci_iot.OperationStatusSucceeded),
			string(oci_iot.OperationStatusFailed),
			string(oci_iot.OperationStatusNeedsAttention),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_iot.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_iot.OperationStatusFailed || response.Status == oci_iot.OperationStatusNeedsAttention {
		return nil, getErrorFromIotIotDomainWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromIotIotDomainWorkRequest(ctx context.Context, client *oci_iot.IotClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_iot.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_iot.ListWorkRequestErrorsRequest{
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

func (s *IotIotDomainResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_iot.GetIotDomainRequest{}

	tmp := s.D.Id()
	request.IotDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.GetIotDomain(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.IotDomain
	return nil
}

func (s *IotIotDomainResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_iot.UpdateIotDomainRequest{}

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
	request.IotDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.UpdateIotDomain(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotDomainFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *IotIotDomainResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_iot.DeleteIotDomainRequest{}

	tmp := s.D.Id()
	request.IotDomainId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.DeleteIotDomain(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := iotDomainWaitForWorkRequest(ctx, workId, "iotdomain",
		oci_iot.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *IotIotDomainResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataRetentionPeriodsInDays != nil {
		s.D.Set("data_retention_periods_in_days", []interface{}{DataRetentionPeriodsInDaysToMap(s.Res.DataRetentionPeriodsInDays)})
	} else {
		s.D.Set("data_retention_periods_in_days", nil)
	}

	s.D.Set("db_allow_listed_identity_group_names", s.Res.DbAllowListedIdentityGroupNames)

	if s.Res.DbAllowedIdentityDomainHost != nil {
		s.D.Set("db_allowed_identity_domain_host", *s.Res.DbAllowedIdentityDomainHost)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DeviceHost != nil {
		s.D.Set("device_host", *s.Res.DeviceHost)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IotDomainGroupId != nil {
		s.D.Set("iot_domain_group_id", *s.Res.IotDomainGroupId)
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

func DataRetentionPeriodsInDaysToMap(obj *oci_iot.DataRetentionPeriodsInDays) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HistorizedData != nil {
		result["historized_data"] = int(*obj.HistorizedData)
	}

	if obj.RawCommandData != nil {
		result["raw_command_data"] = int(*obj.RawCommandData)
	}

	if obj.RawData != nil {
		result["raw_data"] = int(*obj.RawData)
	}

	if obj.RejectedData != nil {
		result["rejected_data"] = int(*obj.RejectedData)
	}

	return result
}

func IotDomainSummaryToMap(obj oci_iot.IotDomainSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IotDomainGroupId != nil {
		result["iot_domain_group_id"] = string(*obj.IotDomainGroupId)
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

func (s *IotIotDomainResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_iot.ChangeIotDomainCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.IotDomainId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot")

	response, err := s.Client.ChangeIotDomainCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getIotDomainFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "iot"), oci_iot.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
