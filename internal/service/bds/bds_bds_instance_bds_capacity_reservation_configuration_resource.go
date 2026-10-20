// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bds "github.com/oracle/oci-go-sdk/v65/bds"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BdsBdsInstanceBdsCapacityReservationConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBdsBdsInstanceBdsCapacityReservationConfigurationWithContext,
		ReadContext:   readBdsBdsInstanceBdsCapacityReservationConfigurationWithContext,
		UpdateContext: updateBdsBdsInstanceBdsCapacityReservationConfigurationWithContext,
		DeleteContext: deleteBdsBdsInstanceBdsCapacityReservationConfigurationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"bds_capacity_reservation_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bds_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"activate_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"deactivate_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
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
	}
}

func createBdsBdsInstanceBdsCapacityReservationConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if _, ok := sync.D.GetOkExists("activate_trigger"); ok {
		err := sync.ActivateBdsCapacityReservationConfiguration(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("deactivate_trigger"); ok {
		err := sync.DeactivateBdsCapacityReservationConfiguration(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readBdsBdsInstanceBdsCapacityReservationConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBdsBdsInstanceBdsCapacityReservationConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	if _, ok := sync.D.GetOkExists("activate_trigger"); ok && sync.D.HasChange("activate_trigger") {
		oldRaw, newRaw := sync.D.GetChange("activate_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ActivateBdsCapacityReservationConfiguration(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("activate_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if _, ok := sync.D.GetOkExists("deactivate_trigger"); ok && sync.D.HasChange("deactivate_trigger") {
		oldRaw, newRaw := sync.D.GetChange("deactivate_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.DeactivateBdsCapacityReservationConfiguration(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("deactivate_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteBdsBdsInstanceBdsCapacityReservationConfigurationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsCapacityReservationConfiguration
	DisableNotFoundRetries bool
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) ID() string {
	return GetBdsInstanceBdsCapacityReservationConfigurationCompositeId(*s.Res.Id, s.D.Get("bds_instance_id").(string))
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationConfigurationLifecycleStateCreating),
	}
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationConfigurationLifecycleStateActive),
	}
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationConfigurationLifecycleStateDeleting),
	}
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationConfigurationLifecycleStateDeleted),
	}
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_bds.CreateBdsCapacityReservationConfigurationRequest{}

	if bdsCapacityReservationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationId.(string)
		request.BdsCapacityReservationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsCapacityReservationConfiguration(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.Res = &response.BdsCapacityReservationConfiguration
	if s.Res.Id != nil {
		s.D.SetId(GetBdsInstanceBdsCapacityReservationConfigurationCompositeId(*s.Res.Id, s.D.Get("bds_instance_id").(string)))
	}
	if workId == nil {
		return s.GetWithContext(ctx)
	}
	return s.getBdsInstanceBdsCapacityReservationConfigurationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) getBdsInstanceBdsCapacityReservationConfigurationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsInstanceBdsCapacityReservationConfigurationId, err := bdsInstanceBdsCapacityReservationConfigurationWaitForWorkRequest(ctx, workId, "bdscapacityreservationconfiguration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsInstanceBdsCapacityReservationConfigurationId)

	return s.GetWithContext(ctx)
}

func bdsInstanceBdsCapacityReservationConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "bds", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_bds.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func bdsInstanceBdsCapacityReservationConfigurationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsInstanceBdsCapacityReservationConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_bds.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_bds.OperationStatusAccepted),
			string(oci_bds.OperationStatusInProgress),
			string(oci_bds.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_bds.OperationStatusFailed),
			string(oci_bds.OperationStatusSucceeded),
			string(oci_bds.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_bds.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_bds.OperationStatusFailed || response.Status == oci_bds.OperationStatusCanceled {
		return nil, getErrorFromBdsBdsInstanceBdsCapacityReservationConfigurationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsInstanceBdsCapacityReservationConfigurationWorkRequest(ctx context.Context, client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_bds.ListWorkRequestErrorsRequest{
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

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.GetBdsCapacityReservationConfigurationRequest{}

	if bdsCapacityReservationConfigurationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationConfigurationId.(string)
		request.BdsCapacityReservationConfigurationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCapacityReservationConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCapacityReservationConfigurationId = &bdsCapacityReservationConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsCapacityReservationConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCapacityReservationConfiguration
	return nil
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) UpdateWithContext(ctx context.Context) error {
	request := oci_bds.UpdateBdsCapacityReservationConfigurationRequest{}

	if bdsCapacityReservationConfigurationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationConfigurationId.(string)
		request.BdsCapacityReservationConfigurationId = &tmp
	}

	if bdsCapacityReservationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationId.(string)
		request.BdsCapacityReservationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	bdsCapacityReservationConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCapacityReservationConfigurationId = &bdsCapacityReservationConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateBdsCapacityReservationConfiguration(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCapacityReservationConfiguration
	return nil
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_bds.DeleteBdsCapacityReservationConfigurationRequest{}

	if bdsCapacityReservationConfigurationId, ok := s.D.GetOkExists("bds_capacity_reservation_id"); ok {
		tmp := bdsCapacityReservationConfigurationId.(string)
		request.BdsCapacityReservationConfigurationId = &tmp
	}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCapacityReservationConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCapacityReservationConfigurationId = &bdsCapacityReservationConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err = s.Client.DeleteBdsCapacityReservationConfiguration(ctx, request)
	return err
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) SetData() error {

	_, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("bds_instance_id", &bdsInstanceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.BdsCapacityReservationId != nil {
		s.D.Set("bds_capacity_reservation_id", *s.Res.BdsCapacityReservationId)
	}

	if s.Res.BdsInstanceId != nil {
		s.D.Set("bds_instance_id", *s.Res.BdsInstanceId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GetBdsInstanceBdsCapacityReservationConfigurationCompositeId(bdsCapacityReservationConfigurationId string, bdsInstanceId string) string {
	bdsCapacityReservationConfigurationId = url.PathEscape(bdsCapacityReservationConfigurationId)
	bdsInstanceId = url.PathEscape(bdsInstanceId)
	compositeId := "bdsInstances/" + bdsInstanceId + "/bdsCapacityReservationConfigurations/" + bdsCapacityReservationConfigurationId
	return compositeId
}

func parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(compositeId string) (bdsCapacityReservationConfigurationId string, bdsInstanceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("bdsInstances/.*/bdsCapacityReservationConfigurations/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	bdsInstanceId, _ = url.PathUnescape(parts[1])
	bdsCapacityReservationConfigurationId, _ = url.PathUnescape(parts[3])

	return
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) ActivateBdsCapacityReservationConfiguration(ctx context.Context) error {
	request := oci_bds.ActivateBdsCapacityReservationConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCapacityReservationConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCapacityReservationConfigurationId = &bdsCapacityReservationConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.ActivateBdsCapacityReservationConfiguration(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = bdsInstanceBdsCapacityReservationConfigurationWaitForWorkRequest(ctx, workRequestId, "bdscapacityreservationconfiguration", oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("activate_trigger")
	s.D.Set("activate_trigger", val)

	return nil
}

func (s *BdsBdsInstanceBdsCapacityReservationConfigurationResourceCrud) DeactivateBdsCapacityReservationConfiguration(ctx context.Context) error {
	request := oci_bds.DeactivateBdsCapacityReservationConfigurationRequest{}

	if bdsInstanceId, ok := s.D.GetOkExists("bds_instance_id"); ok {
		tmp := bdsInstanceId.(string)
		request.BdsInstanceId = &tmp
	}

	bdsCapacityReservationConfigurationId, bdsInstanceId, err := parseBdsInstanceBdsCapacityReservationConfigurationCompositeId(s.D.Id())
	if err == nil {
		request.BdsCapacityReservationConfigurationId = &bdsCapacityReservationConfigurationId
		request.BdsInstanceId = &bdsInstanceId
	}
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.DeactivateBdsCapacityReservationConfiguration(ctx, request)
	if err != nil {
		return err
	}
	var workRequestId *string
	if response.RawResponse != nil {
		if wr := response.RawResponse.Header.Get("opc-work-request-id"); wr != "" {
			workRequestId = &wr
		}
	}
	if workRequestId != nil {
		_, err = bdsInstanceBdsCapacityReservationConfigurationWaitForWorkRequest(ctx, workRequestId, "bdscapacityreservationconfiguration", oci_bds.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
		if err != nil {
			return err
		}
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("deactivate_trigger")
	s.D.Set("deactivate_trigger", val)

	return nil
}

func BdsCapacityReservationConfigurationSummaryToMap(obj oci_bds.BdsCapacityReservationConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BdsCapacityReservationId != nil {
		result["bds_capacity_reservation_id"] = string(*obj.BdsCapacityReservationId)
	}

	if obj.BdsInstanceId != nil {
		result["bds_instance_id"] = string(*obj.BdsInstanceId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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
