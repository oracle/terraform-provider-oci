// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bds

import (
	"context"
	"fmt"
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

func BdsBdsCapacityReservationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createBdsBdsCapacityReservationWithContext,
		ReadContext:   readBdsBdsCapacityReservationWithContext,
		UpdateContext: updateBdsBdsCapacityReservationWithContext,
		DeleteContext: deleteBdsBdsCapacityReservationWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compute_capacity_reservations": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain1reservation_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain2reservation_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain3reservation_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
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

			// Computed
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

func createBdsBdsCapacityReservationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readBdsBdsCapacityReservationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBdsBdsCapacityReservationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteBdsBdsCapacityReservationWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BdsBdsCapacityReservationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BdsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BdsBdsCapacityReservationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_bds.BdsClient
	Res                    *oci_bds.BdsCapacityReservation
	DisableNotFoundRetries bool
}

func (s *BdsBdsCapacityReservationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BdsBdsCapacityReservationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationLifecycleStateCreating),
	}
}

func (s *BdsBdsCapacityReservationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationLifecycleStateActive),
	}
}

func (s *BdsBdsCapacityReservationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationLifecycleStateDeleting),
	}
}

func (s *BdsBdsCapacityReservationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_bds.BdsCapacityReservationLifecycleStateDeleted),
	}
}

func (s *BdsBdsCapacityReservationResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_bds.CreateBdsCapacityReservationRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if computeCapacityReservations, ok := s.D.GetOkExists("compute_capacity_reservations"); ok {
		if tmpList := computeCapacityReservations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_capacity_reservations", 0)
			tmp, err := s.mapToComputeCapacityReservations(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ComputeCapacityReservations = &tmp
		}
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.CreateBdsCapacityReservation(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	if workId == nil {
		return s.GetWithContext(ctx)
	}
	return s.getBdsCapacityReservationFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds"), oci_bds.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BdsBdsCapacityReservationResourceCrud) getBdsCapacityReservationFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_bds.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	bdsCapacityReservationId, err := bdsCapacityReservationWaitForWorkRequest(ctx, workId, "bdscapacityreservation",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*bdsCapacityReservationId)

	return s.GetWithContext(ctx)
}

func bdsCapacityReservationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func bdsCapacityReservationWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_bds.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_bds.BdsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "bds")
	retryPolicy.ShouldRetryOperation = bdsCapacityReservationWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromBdsBdsCapacityReservationWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBdsBdsCapacityReservationWorkRequest(ctx context.Context, client *oci_bds.BdsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_bds.ActionTypesEnum) error {
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

func (s *BdsBdsCapacityReservationResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_bds.GetBdsCapacityReservationRequest{}

	tmp := s.D.Id()
	request.BdsCapacityReservationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.GetBdsCapacityReservation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCapacityReservation
	return nil
}

func (s *BdsBdsCapacityReservationResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_bds.UpdateBdsCapacityReservationRequest{}

	tmp := s.D.Id()
	request.BdsCapacityReservationId = &tmp

	if computeCapacityReservations, ok := s.D.GetOkExists("compute_capacity_reservations"); ok {
		if tmpList := computeCapacityReservations.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "compute_capacity_reservations", 0)
			tmp, err := s.mapToComputeCapacityReservations(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ComputeCapacityReservations = &tmp
		}
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	response, err := s.Client.UpdateBdsCapacityReservation(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BdsCapacityReservation
	return nil
}

func (s *BdsBdsCapacityReservationResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_bds.DeleteBdsCapacityReservationRequest{}

	tmp := s.D.Id()
	request.BdsCapacityReservationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err := s.Client.DeleteBdsCapacityReservation(ctx, request)
	return err
}

func (s *BdsBdsCapacityReservationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCapacityReservations != nil {
		s.D.Set("compute_capacity_reservations", []interface{}{ComputeCapacityReservationsToMap(s.Res.ComputeCapacityReservations)})
	} else {
		s.D.Set("compute_capacity_reservations", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

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

func BdsCapacityReservationSummaryToMap(obj oci_bds.BdsCapacityReservationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.ComputeCapacityReservations != nil {
		result["compute_capacity_reservations"] = []interface{}{ComputeCapacityReservationsToMap(obj.ComputeCapacityReservations)}
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

func (s *BdsBdsCapacityReservationResourceCrud) mapToComputeCapacityReservations(fieldKeyFormat string) (oci_bds.ComputeCapacityReservations, error) {
	result := oci_bds.ComputeCapacityReservations{}

	if domain1ReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain1reservation_id")); ok {
		tmp := domain1ReservationId.(string)
		result.Domain1ReservationId = &tmp
	}

	if domain2ReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain2reservation_id")); ok {
		tmp := domain2ReservationId.(string)
		result.Domain2ReservationId = &tmp
	}

	if domain3ReservationId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain3reservation_id")); ok {
		tmp := domain3ReservationId.(string)
		result.Domain3ReservationId = &tmp
	}

	return result, nil
}

func ComputeCapacityReservationsToMap(obj *oci_bds.ComputeCapacityReservations) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Domain1ReservationId != nil {
		result["domain1reservation_id"] = string(*obj.Domain1ReservationId)
	}

	if obj.Domain2ReservationId != nil {
		result["domain2reservation_id"] = string(*obj.Domain2ReservationId)
	}

	if obj.Domain3ReservationId != nil {
		result["domain3reservation_id"] = string(*obj.Domain3ReservationId)
	}

	return result
}

func (s *BdsBdsCapacityReservationResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_bds.ChangeBdsCapacityReservationCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BdsCapacityReservationId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "bds")

	_, err := s.Client.ChangeBdsCapacityReservationCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
