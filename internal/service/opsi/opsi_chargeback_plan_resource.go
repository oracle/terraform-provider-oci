// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opsi

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpsiChargebackPlanResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createOpsiChargebackPlanWithContext,
		ReadContext:   readOpsiChargebackPlanWithContext,
		UpdateContext: updateOpsiChargebackPlanWithContext,
		DeleteContext: deleteOpsiChargebackPlanWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_source": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"plan_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"plan_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice(
					[]string{
						"EQUAL_ALLOCATION",
						"WEIGHTED_ALLOCATION",
						"UNUSED_ALLOCATION",
					},
					true,
				),
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
			"plan_custom_items": {
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"is_customizable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Computed
					},
				},
			},
			"plan_description": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Computed
			"is_customizable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"plan_category": {
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

func createOpsiChargebackPlanWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readOpsiChargebackPlanWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateOpsiChargebackPlanWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteOpsiChargebackPlanWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &OpsiChargebackPlanResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OperationsInsightsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type OpsiChargebackPlanResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_opsi.OperationsInsightsClient
	Res                    *oci_opsi.ChargebackPlan
	DisableNotFoundRetries bool
}

func (s *OpsiChargebackPlanResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OpsiChargebackPlanResourceCrud) resolveChargebackPlanID() (*string, error) {
	if chargebackplanId, ok := s.D.GetOkExists("chargebackplan_id"); ok {
		if id := strings.TrimSpace(chargebackplanId.(string)); id != "" {
			return &id, nil
		}
	}

	if id := strings.TrimSpace(s.D.Id()); id != "" {
		return &id, nil
	}

	return nil, fmt.Errorf("chargeback plan identifier is not set")
}

func (s *OpsiChargebackPlanResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateCreating),
	}
}

func (s *OpsiChargebackPlanResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateActive),
		string(oci_opsi.LifecycleStateNeedsAttention),
	}
}

func (s *OpsiChargebackPlanResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleting),
	}
}

func (s *OpsiChargebackPlanResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_opsi.LifecycleStateDeleted),
	}
}

func (s *OpsiChargebackPlanResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_opsi.CreateChargebackPlanRequest{}
	err := s.populateTopLevelPolymorphicCreateChargebackPlanRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.CreateChargebackPlan(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getChargebackPlanFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *OpsiChargebackPlanResourceCrud) getChargebackPlanFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_opsi.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	chargebackPlanId, err := chargebackPlanWaitForWorkRequest(ctx, workId, "chargebackplan",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*chargebackPlanId)

	return s.GetWithContext(ctx)
}

func chargebackPlanWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func chargebackPlanWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_opsi.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_opsi.OperationsInsightsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "opsi")
	retryPolicy.ShouldRetryOperation = chargebackPlanWorkRequestShouldRetryFunc(timeout)

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
			response, err = client.GetWorkRequest(ctx,
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
		return nil, getErrorFromOpsiChargebackPlanWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromOpsiChargebackPlanWorkRequest(ctx context.Context, client *oci_opsi.OperationsInsightsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_opsi.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
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

func (s *OpsiChargebackPlanResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_opsi.GetChargebackPlanRequest{}

	if chargebackplanId, err := s.resolveChargebackPlanID(); err == nil {
		request.ChargebackplanId = chargebackplanId
	} else {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.GetChargebackPlan(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.ChargebackPlan
	return nil
}

func (s *OpsiChargebackPlanResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_opsi.UpdateChargebackPlanRequest{}

	if chargebackplanId, err := s.resolveChargebackPlanID(); err == nil {
		request.ChargebackplanId = chargebackplanId
	} else {
		return err
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if planCustomItems, ok := s.D.GetOkExists("plan_custom_items"); ok {
		interfaces := planCustomItems.([]interface{})
		tmp := make([]oci_opsi.CreatePlanCustomItemDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "plan_custom_items", stateDataIndex)
			converted, err := s.mapToCreatePlanCustomItemDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("plan_custom_items") {
			request.PlanCustomItems = tmp
		}
	}

	if planDescription, ok := s.D.GetOkExists("plan_description"); ok {
		tmp := planDescription.(string)
		request.PlanDescription = &tmp
	}

	if planName, ok := s.D.GetOkExists("plan_name"); ok {
		tmp := planName.(string)
		request.PlanName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.UpdateChargebackPlan(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getChargebackPlanFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *OpsiChargebackPlanResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_opsi.DeleteChargebackPlanRequest{}

	if chargebackplanId, err := s.resolveChargebackPlanID(); err == nil {
		request.ChargebackplanId = chargebackplanId
	} else {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.DeleteChargebackPlan(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := chargebackPlanWaitForWorkRequest(ctx, workId, "chargebackplan",
		oci_opsi.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *OpsiChargebackPlanResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("entity_source", s.Res.EntitySource)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCustomizable != nil {
		s.D.Set("is_customizable", *s.Res.IsCustomizable)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("plan_category", s.Res.PlanCategory)

	planCustomItems := []interface{}{}
	for _, item := range s.Res.PlanCustomItems {
		planCustomItems = append(planCustomItems, CreatePlanCustomItemDetailsToMap(item))
	}
	s.D.Set("plan_custom_items", planCustomItems)

	if s.Res.PlanDescription != nil {
		s.D.Set("plan_description", *s.Res.PlanDescription)
	}

	if s.Res.PlanName != nil {
		s.D.Set("plan_name", *s.Res.PlanName)
	}

	if s.Res.PlanType != nil {
		s.D.Set("plan_type", *s.Res.PlanType)
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

func ChargebackPlanSummaryToMap(obj oci_opsi.ChargebackPlanSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["entity_source"] = string(obj.EntitySource)

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCustomizable != nil {
		result["is_customizable"] = bool(*obj.IsCustomizable)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["plan_category"] = string(obj.PlanCategory)

	planCustomItems := []interface{}{}
	for _, item := range obj.PlanCustomItems {
		planCustomItems = append(planCustomItems, CreatePlanCustomItemDetailsToMap(item))
	}
	result["plan_custom_items"] = planCustomItems

	if obj.PlanDescription != nil {
		result["plan_description"] = string(*obj.PlanDescription)
	}

	if obj.PlanName != nil {
		result["plan_name"] = string(*obj.PlanName)
	}

	if obj.PlanType != nil {
		result["plan_type"] = string(*obj.PlanType)
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

func (s *OpsiChargebackPlanResourceCrud) mapToCreatePlanCustomItemDetails(fieldKeyFormat string) (oci_opsi.CreatePlanCustomItemDetails, error) {
	result := oci_opsi.CreatePlanCustomItemDetails{}

	if isCustomizable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_customizable")); ok {
		tmp := isCustomizable.(bool)
		result.IsCustomizable = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func CreatePlanCustomItemDetailsToMap(obj oci_opsi.CreatePlanCustomItemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsCustomizable != nil {
		result["is_customizable"] = bool(*obj.IsCustomizable)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *OpsiChargebackPlanResourceCrud) populateTopLevelPolymorphicCreateChargebackPlanRequest(request *oci_opsi.CreateChargebackPlanRequest) error {
	//discriminator
	entitySourceRaw, ok := s.D.GetOkExists("entity_source")
	var entitySource string
	if ok {
		entitySource = entitySourceRaw.(string)
	} else {
		entitySource = "" // default value
	}
	switch strings.ToLower(entitySource) {
	case strings.ToLower("CHARGEBACK_EXADATA"):
		details := oci_opsi.CreateChargebackPlanExadataDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if planCustomItems, ok := s.D.GetOkExists("plan_custom_items"); ok {
			interfaces := planCustomItems.([]interface{})
			tmp := make([]oci_opsi.CreatePlanCustomItemDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "plan_custom_items", stateDataIndex)
				converted, err := s.mapToCreatePlanCustomItemDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("plan_custom_items") {
				details.PlanCustomItems = tmp
			}
		}
		if planDescription, ok := s.D.GetOkExists("plan_description"); ok {
			tmp := planDescription.(string)
			details.PlanDescription = &tmp
		}
		if planName, ok := s.D.GetOkExists("plan_name"); ok {
			tmp := planName.(string)
			details.PlanName = &tmp
		}
		if planType, ok := s.D.GetOkExists("plan_type"); ok {
			tmp := planType.(string)
			details.PlanType = &tmp
		}
		request.CreateChargebackPlanDetails = details
	default:
		return fmt.Errorf("unknown entity_source '%v' was specified", entitySource)
	}
	return nil
}

func (s *OpsiChargebackPlanResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_opsi.ChangeChargebackPlanCompartmentRequest{}

	if chargebackplanId, err := s.resolveChargebackPlanID(); err == nil {
		changeCompartmentRequest.ChargebackplanId = chargebackplanId
	} else {
		return err
	}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi")

	response, err := s.Client.ChangeChargebackPlanCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getChargebackPlanFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "opsi"), oci_opsi.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
