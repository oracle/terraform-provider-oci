// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package em_warehouse

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_em_warehouse "github.com/oracle/oci-go-sdk/v65/emwarehouse"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func EmWarehouseEmWarehouseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createEmWarehouseEmWarehouse,
		Read:     readEmWarehouseEmWarehouse,
		Update:   updateEmWarehouseEmWarehouse,
		Delete:   deleteEmWarehouseEmWarehouse,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"em_bridge_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operations_insights_warehouse_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"em_warehouse_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_etl_run_message": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_etl_run_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latest_etl_run_time": {
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

func createEmWarehouseEmWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.CreateResource(d, sync)
}

func readEmWarehouseEmWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.ReadResource(sync)
}

func updateEmWarehouseEmWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteEmWarehouseEmWarehouse(d *schema.ResourceData, m interface{}) error {
	sync := &EmWarehouseEmWarehouseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).EmWarehouseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type EmWarehouseEmWarehouseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_em_warehouse.EmWarehouseClient
	Res                    *oci_em_warehouse.EmWarehouse
	DisableNotFoundRetries bool
}

func (s *EmWarehouseEmWarehouseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *EmWarehouseEmWarehouseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_em_warehouse.EmWarehouseLifecycleStateCreating),
	}
}

func (s *EmWarehouseEmWarehouseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_em_warehouse.EmWarehouseLifecycleStateActive),
	}
}

func (s *EmWarehouseEmWarehouseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_em_warehouse.EmWarehouseLifecycleStateDeleting),
	}
}

func (s *EmWarehouseEmWarehouseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_em_warehouse.EmWarehouseLifecycleStateDeleted),
	}
}

func (s *EmWarehouseEmWarehouseResourceCrud) Create() error {
	request := oci_em_warehouse.CreateEmWarehouseRequest{}

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

	if emBridgeId, ok := s.D.GetOkExists("em_bridge_id"); ok {
		tmp := emBridgeId.(string)
		request.EmBridgeId = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if operationsInsightsWarehouseId, ok := s.D.GetOkExists("operations_insights_warehouse_id"); ok {
		tmp := operationsInsightsWarehouseId.(string)
		request.OperationsInsightsWarehouseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse")

	response, err := s.Client.CreateEmWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_em_warehouse.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_em_warehouse.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "emwarehouse") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getEmWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse"), oci_em_warehouse.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *EmWarehouseEmWarehouseResourceCrud) getEmWarehouseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_em_warehouse.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	emWarehouseId, err := emWarehouseWaitForWorkRequest(workId, "emwarehouse",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, emWarehouseId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_em_warehouse.CancelWorkRequestRequest{
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
	s.D.SetId(*emWarehouseId)

	return s.Get()
}

func emWarehouseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "emwarehouse", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_em_warehouse.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func emWarehouseWaitForWorkRequest(wId *string, entityType string, action oci_em_warehouse.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_em_warehouse.EmWarehouseClient) (*string, error) {

	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "emwarehouse")
	retryPolicy.ShouldRetryOperation = emWarehouseWorkRequestShouldRetryFunc(timeout)

	response := oci_em_warehouse.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_em_warehouse.OperationStatusInProgress),
			string(oci_em_warehouse.OperationStatusAccepted),
			string(oci_em_warehouse.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_em_warehouse.OperationStatusSucceeded),
			string(oci_em_warehouse.OperationStatusFailed),
			string(oci_em_warehouse.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_em_warehouse.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_em_warehouse.OperationStatusFailed || response.Status == oci_em_warehouse.OperationStatusCanceled {
		return nil, getErrorFromEmWarehouseEmWarehouseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromEmWarehouseEmWarehouseWorkRequest(client *oci_em_warehouse.EmWarehouseClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_em_warehouse.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_em_warehouse.ListWorkRequestErrorsRequest{
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

func (s *EmWarehouseEmWarehouseResourceCrud) Get() error {
	request := oci_em_warehouse.GetEmWarehouseRequest{}

	tmp := s.D.Id()
	request.EmWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse")

	response, err := s.Client.GetEmWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.EmWarehouse
	return nil
}

func (s *EmWarehouseEmWarehouseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_em_warehouse.UpdateEmWarehouseRequest{}

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

	if emBridgeId, ok := s.D.GetOkExists("em_bridge_id"); ok {
		tmp := emBridgeId.(string)
		request.EmBridgeId = &tmp
	}

	tmp := s.D.Id()
	request.EmWarehouseId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse")

	response, err := s.Client.UpdateEmWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse"), oci_em_warehouse.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *EmWarehouseEmWarehouseResourceCrud) Delete() error {
	request := oci_em_warehouse.DeleteEmWarehouseRequest{}

	tmp := s.D.Id()
	request.EmWarehouseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse")

	response, err := s.Client.DeleteEmWarehouse(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := emWarehouseWaitForWorkRequest(workId, "emwarehouse",
		oci_em_warehouse.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *EmWarehouseEmWarehouseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EmBridgeId != nil {
		s.D.Set("em_bridge_id", *s.Res.EmBridgeId)
	}

	if s.Res.EmWarehouseType != nil {
		s.D.Set("em_warehouse_type", *s.Res.EmWarehouseType)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LatestEtlRunMessage != nil {
		s.D.Set("latest_etl_run_message", *s.Res.LatestEtlRunMessage)
	}

	if s.Res.LatestEtlRunStatus != nil {
		s.D.Set("latest_etl_run_status", *s.Res.LatestEtlRunStatus)
	}

	if s.Res.LatestEtlRunTime != nil {
		s.D.Set("latest_etl_run_time", *s.Res.LatestEtlRunTime)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OperationsInsightsWarehouseId != nil {
		s.D.Set("operations_insights_warehouse_id", *s.Res.OperationsInsightsWarehouseId)
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

func EmWarehouseSummaryToMap(obj oci_em_warehouse.EmWarehouseSummary) map[string]interface{} {
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

	if obj.EmBridgeId != nil {
		result["em_bridge_id"] = string(*obj.EmBridgeId)
	}

	if obj.EmWarehouseType != nil {
		result["em_warehouse_type"] = string(*obj.EmWarehouseType)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LatestEtlRunMessage != nil {
		result["latest_etl_run_message"] = string(*obj.LatestEtlRunMessage)
	}

	if obj.LatestEtlRunStatus != nil {
		result["latest_etl_run_status"] = string(*obj.LatestEtlRunStatus)
	}

	if obj.LatestEtlRunTime != nil {
		result["latest_etl_run_time"] = string(*obj.LatestEtlRunTime)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OperationsInsightsWarehouseId != nil {
		result["operations_insights_warehouse_id"] = string(*obj.OperationsInsightsWarehouseId)
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

func (s *EmWarehouseEmWarehouseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_em_warehouse.ChangeEmWarehouseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.EmWarehouseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse")

	response, err := s.Client.ChangeEmWarehouseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getEmWarehouseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "emwarehouse"), oci_em_warehouse.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
