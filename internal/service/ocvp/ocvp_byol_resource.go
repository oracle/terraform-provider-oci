// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpByolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createOcvpByolWithContext,
		ReadContext:   readOcvpByolWithContext,
		UpdateContext: updateOcvpByolWithContext,
		DeleteContext: deleteOcvpByolWithContext,
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
			"entitlement_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"software_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_term_end": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"time_term_start": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"total_units": {
				Type:     schema.TypeInt,
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
			"available_units": {
				Type:     schema.TypeInt,
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

func createOcvpByolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readOcvpByolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateOcvpByolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolClient()
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteOcvpByolWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OcvpByolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ByolClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).OcvpWorkRequestClient()

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type OcvpByolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ocvp.ByolClient
	Res                    *oci_ocvp.Byol
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_ocvp.WorkRequestClient
}

func (s *OcvpByolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OcvpByolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ocvp.ByolLifecycleStateCreating),
	}
}

func (s *OcvpByolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ocvp.ByolLifecycleStateActive),
	}
}

func (s *OcvpByolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ocvp.ByolLifecycleStateDeleting),
	}
}

func (s *OcvpByolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ocvp.ByolLifecycleStateDeleted),
	}
}

func (s *OcvpByolResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_ocvp.CreateByolRequest{}

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

	if entitlementKey, ok := s.D.GetOkExists("entitlement_key"); ok {
		tmp := entitlementKey.(string)
		request.EntitlementKey = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if softwareType, ok := s.D.GetOkExists("software_type"); ok {
		request.SoftwareType = oci_ocvp.ByolSoftwareTypeEnum(softwareType.(string))
	}

	if timeTermEnd, ok := s.D.GetOkExists("time_term_end"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeTermEnd.(string))
		if err != nil {
			return err
		}
		request.TimeTermEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeTermStart, ok := s.D.GetOkExists("time_term_start"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeTermStart.(string))
		if err != nil {
			return err
		}
		request.TimeTermStart = &oci_common.SDKTime{Time: tmp}
	}

	if totalUnits, ok := s.D.GetOkExists("total_units"); ok {
		tmp := totalUnits.(int)
		request.TotalUnits = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.CreateByol(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_ocvp.GetWorkRequestResponse{}
	workRequestResponse, err = s.WorkRequestClient.GetWorkRequest(ctx,
		oci_ocvp.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "byol") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getByolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OcvpByolResourceCrud) getByolFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ocvp.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	byolId, err := byolWaitForWorkRequest(ctx, workId, "byol",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*byolId)

	return s.GetWithContext(ctx)
}

func byolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ocvp", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ocvp.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func byolWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_ocvp.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ocvp.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ocvp")
	retryPolicy.ShouldRetryOperation = byolWorkRequestShouldRetryFunc(timeout)

	response := oci_ocvp.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ocvp.OperationStatusInProgress),
			string(oci_ocvp.OperationStatusAccepted),
			string(oci_ocvp.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_ocvp.OperationStatusSucceeded),
			string(oci_ocvp.OperationStatusFailed),
			string(oci_ocvp.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_ocvp.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_ocvp.OperationStatusFailed || response.Status == oci_ocvp.OperationStatusCanceled {
		return nil, getErrorFromOcvpByolWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOcvpByolWorkRequest(ctx context.Context, client *oci_ocvp.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ocvp.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_ocvp.ListWorkRequestErrorsRequest{
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

func (s *OcvpByolResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_ocvp.GetByolRequest{}

	tmp := s.D.Id()
	request.ByolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.GetByol(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Byol
	return nil
}

func (s *OcvpByolResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ocvp.UpdateByolRequest{}

	tmp := s.D.Id()
	request.ByolId = &tmp

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

	if entitlementKey, ok := s.D.GetOkExists("entitlement_key"); ok {
		tmp := entitlementKey.(string)
		request.EntitlementKey = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if softwareType, ok := s.D.GetOkExists("software_type"); ok {
		request.SoftwareType = oci_ocvp.ByolSoftwareTypeEnum(softwareType.(string))
	}

	if timeTermEnd, ok := s.D.GetOkExists("time_term_end"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeTermEnd.(string))
		if err != nil {
			return err
		}
		request.TimeTermEnd = &oci_common.SDKTime{Time: tmp}
	}

	if timeTermStart, ok := s.D.GetOkExists("time_term_start"); ok {
		tmp, err := time.Parse(time.RFC3339Nano, timeTermStart.(string))
		if err != nil {
			return err
		}
		request.TimeTermStart = &oci_common.SDKTime{Time: tmp}
	}

	if totalUnits, ok := s.D.GetOkExists("total_units"); ok {
		tmp := totalUnits.(int)
		request.TotalUnits = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.UpdateByol(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getByolFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp"), oci_ocvp.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OcvpByolResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_ocvp.DeleteByolRequest{}

	tmp := s.D.Id()
	request.ByolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	response, err := s.Client.DeleteByol(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := byolWaitForWorkRequest(ctx, workId, "byol",
		oci_ocvp.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *OcvpByolResourceCrud) SetData() error {
	if s.Res.AvailableUnits != nil {
		s.D.Set("available_units", *s.Res.AvailableUnits)
	}

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

	if s.Res.EntitlementKey != nil {
		s.D.Set("entitlement_key", *s.Res.EntitlementKey)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("software_type", s.Res.SoftwareType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeTermEnd != nil {
		s.D.Set("time_term_end", s.Res.TimeTermEnd.Format(time.RFC3339Nano))
	}

	if s.Res.TimeTermStart != nil {
		s.D.Set("time_term_start", s.Res.TimeTermStart.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalUnits != nil {
		s.D.Set("total_units", *s.Res.TotalUnits)
	}

	return nil
}

func ByolSummaryToMap(obj oci_ocvp.ByolSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailableUnits != nil {
		result["available_units"] = int(*obj.AvailableUnits)
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

	result["software_type"] = string(obj.SoftwareType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeTermEnd != nil {
		result["time_term_end"] = obj.TimeTermEnd.String()
	}

	if obj.TimeTermStart != nil {
		result["time_term_start"] = obj.TimeTermStart.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TotalUnits != nil {
		result["total_units"] = int(*obj.TotalUnits)
	}

	return result
}

func (s *OcvpByolResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_ocvp.ChangeByolCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.ByolId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ocvp")

	_, err := s.Client.ChangeByolCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
