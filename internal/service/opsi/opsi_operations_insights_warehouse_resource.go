// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func OpsiOperationsInsightsWarehouseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOpsiOperationsInsightsWarehouse,
		Read:     readOpsiOperationsInsightsWarehouse,
		Update:   updateOpsiOperationsInsightsWarehouse,
		Delete:   deleteOpsiOperationsInsightsWarehouse,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu_allocated": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"compute_model": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"storage_allocated_in_gbs": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},

			// Computed
			"cpu_used": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"dynamic_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operations_insights_tenancy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_used_in_gbs": {
				Type:     schema.TypeFloat,
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
			"time_last_wallet_rotated": {
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

func createOpsiOperationsInsightsWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.CreateResource(d, sync)
}

func readOpsiOperationsInsightsWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.ReadResource(sync)
}

func updateOpsiOperationsInsightsWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOpsiOperationsInsightsWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &OpsiOperationsInsightsWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OpsiOperationsInsightsWarehouseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.OperationsInsightsWarehouse
	DisableNotFoundRetries bool
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateCreating),
	}
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateActive),
	}
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateDeleting),
	}
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.OperationsInsightsWarehouseLifecycleStateDeleted),
	}
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) Create() error {
	request := oci_opsi.CreateOperationsInsightsWarehouseRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
		request.ComputeModel = oci_opsi.OperationsInsightsWarehouseComputeModelEnum(computeModel.(string))
	}

	if cpuAllocated, ok := s.D.GetOkExists("cpu_allocated"); ok {
		tmp := cpuAllocated.(float64)
		request.CpuAllocated = &tmp
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

	if storageAllocatedInGBs, ok := s.D.GetOkExists("storage_allocated_in_gbs"); ok {
		tmp := storageAllocatedInGBs.(float64)
		request.StorageAllocatedInGBs = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateOperationsInsightsWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOperationsInsightsWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) getOperationsInsightsWarehouseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	operationsInsightsWarehouseId, err := operationsInsightsWarehouseWaitForWorkRequest(workId, "opsi",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*operationsInsightsWarehouseId)

	return s.Get()
}

func operationsInsightsWarehouseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func operationsInsightsWarehouseWaitForWorkRequest(wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = operationsInsightsWarehouseWorkRequestShouldRetryFunc(timeout)

	response := oci_opsi.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		return nil, getErrorFromOpsiOperationsInsightsWarehouseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiOperationsInsightsWarehouseWorkRequest(client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
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

func (s *OpsiOperationsInsightsWarehouseResourceCrud) Get() error {
	request := oci_opsi.GetOperationsInsightsWarehouseRequest{}

	tmp := s.D.Id()
	request.OperationsInsightsWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetOperationsInsightsWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.OperationsInsightsWarehouse
	return nil
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateOperationsInsightsWarehouseRequest{}

	if computeModel, ok := s.D.GetOkExists("compute_model"); ok {
		request.ComputeModel = oci_opsi.OperationsInsightsWarehouseComputeModelEnum(computeModel.(string))
	}

	if cpuAllocated, ok := s.D.GetOkExists("cpu_allocated"); ok {
		tmp := cpuAllocated.(float64)
		request.CpuAllocated = &tmp
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
	request.OperationsInsightsWarehouseId = &tmp

	if storageAllocatedInGBs, ok := s.D.GetOkExists("storage_allocated_in_gbs"); ok {
		tmp := storageAllocatedInGBs.(float64)
		request.StorageAllocatedInGBs = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateOperationsInsightsWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOperationsInsightsWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeRelated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) Delete() error {
	request := oci_opsi.DeleteOperationsInsightsWarehouseRequest{}

	tmp := s.D.Id()
	request.OperationsInsightsWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteOperationsInsightsWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := operationsInsightsWarehouseWaitForWorkRequest(workId, "opsi",
		oci_opsi.ActionTypeRelated, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("compute_model", s.Res.ComputeModel)

	if s.Res.CpuAllocated != nil {
		s.D.Set("cpu_allocated", *s.Res.CpuAllocated)
	}

	if s.Res.CpuUsed != nil {
		s.D.Set("cpu_used", *s.Res.CpuUsed)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DynamicGroupId != nil {
		s.D.Set("dynamic_group_id", *s.Res.DynamicGroupId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OperationsInsightsTenancyId != nil {
		s.D.Set("operations_insights_tenancy_id", *s.Res.OperationsInsightsTenancyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageAllocatedInGBs != nil {
		s.D.Set("storage_allocated_in_gbs", *s.Res.StorageAllocatedInGBs)
	}

	if s.Res.StorageUsedInGBs != nil {
		s.D.Set("storage_used_in_gbs", *s.Res.StorageUsedInGBs)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastWalletRotated != nil {
		s.D.Set("time_last_wallet_rotated", s.Res.TimeLastWalletRotated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func OperationsInsightsWarehouseSummaryToMap(obj oci_opsi.OperationsInsightsWarehouseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["compute_model"] = string(obj.ComputeModel)

	if obj.CpuAllocated != nil {
		result["cpu_allocated"] = float64(*obj.CpuAllocated)
	}

	if obj.CpuUsed != nil {
		result["cpu_used"] = float64(*obj.CpuUsed)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DynamicGroupId != nil {
		result["dynamic_group_id"] = string(*obj.DynamicGroupId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OperationsInsightsTenancyId != nil {
		result["operations_insights_tenancy_id"] = string(*obj.OperationsInsightsTenancyId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.StorageAllocatedInGBs != nil {
		result["storage_allocated_in_gbs"] = float64(*obj.StorageAllocatedInGBs)
	}

	if obj.StorageUsedInGBs != nil {
		result["storage_used_in_gbs"] = float64(*obj.StorageUsedInGBs)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastWalletRotated != nil {
		result["time_last_wallet_rotated"] = obj.TimeLastWalletRotated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *OpsiOperationsInsightsWarehouseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeOperationsInsightsWarehouseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OperationsInsightsWarehouseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeOperationsInsightsWarehouseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getOperationsInsightsWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
