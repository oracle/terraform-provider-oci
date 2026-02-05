// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_software_update

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_software_update "github.com/oracle/oci-go-sdk/v65/fleetsoftwareupdate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetSoftwareUpdateFsuReadinessCheckResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createFleetSoftwareUpdateFsuReadinessCheckWithContext,
		ReadContext:   readFleetSoftwareUpdateFsuReadinessCheckWithContext,
		UpdateContext: updateFleetSoftwareUpdateFsuReadinessCheckWithContext,
		DeleteContext: deleteFleetSoftwareUpdateFsuReadinessCheckWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"TARGET",
				}, true),
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
			"targets": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"entity_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"entity_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"issue_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"issues": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"impacted_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recommended_action": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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
			"time_finished": {
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

func createFleetSoftwareUpdateFsuReadinessCheckWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
}

func readFleetSoftwareUpdateFsuReadinessCheckWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateFleetSoftwareUpdateFsuReadinessCheckWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()

	return tfresource.HandleDiagError(m, tfresource.UpdateResourceWithContext(ctx, d, sync))
}

func deleteFleetSoftwareUpdateFsuReadinessCheckWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &FleetSoftwareUpdateFsuReadinessCheckResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetSoftwareUpdateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type FleetSoftwareUpdateFsuReadinessCheckResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_software_update.FleetSoftwareUpdateClient
	Res                    oci_fleet_software_update.FsuReadinessCheck
	DisableNotFoundRetries bool
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) ID() string {
	return *s.Res.GetId()
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateInProgress),
	}
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateNeedsAttention),
		string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateSucceeded),
	}
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateDeleting),
	}
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_software_update.FsuReadinessCheckLifecycleStateDeleted),
	}
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_fleet_software_update.CreateFsuReadinessCheckRequest{}
	err := s.populateTopLevelPolymorphicCreateFsuReadinessCheckRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.CreateFsuReadinessCheck(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getFsuReadinessCheckFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update"), oci_fleet_software_update.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) getFsuReadinessCheckFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_software_update.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fsuReadinessCheckId, err := fsuReadinessCheckWaitForWorkRequest(ctx, workId, "fsureadinesscheck",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fsuReadinessCheckId)

	return s.GetWithContext(ctx)
}

func fsuReadinessCheckWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_software_update", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_software_update.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fsuReadinessCheckWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_fleet_software_update.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_software_update.FleetSoftwareUpdateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_software_update")
	retryPolicy.ShouldRetryOperation = fsuReadinessCheckWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_software_update.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_fleet_software_update.OperationStatusInProgress),
			string(oci_fleet_software_update.OperationStatusAccepted),
			string(oci_fleet_software_update.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_software_update.OperationStatusSucceeded),
			string(oci_fleet_software_update.OperationStatusFailed),
			string(oci_fleet_software_update.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_fleet_software_update.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_software_update.OperationStatusFailed || response.Status == oci_fleet_software_update.OperationStatusCanceled {
		return nil, getErrorFromFleetSoftwareUpdateFsuReadinessCheckWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetSoftwareUpdateFsuReadinessCheckWorkRequest(ctx context.Context, client *oci_fleet_software_update.FleetSoftwareUpdateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_software_update.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_fleet_software_update.ListWorkRequestErrorsRequest{
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

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_fleet_software_update.GetFsuReadinessCheckRequest{}

	tmp := s.D.Id()
	request.FsuReadinessCheckId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.GetFsuReadinessCheck(ctx, request)
	if err != nil {
		return err
	}

	s.Res = response.FsuReadinessCheck
	return nil
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fleet_software_update.UpdateFsuReadinessCheckRequest{}

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
	request.FsuReadinessCheckId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.UpdateFsuReadinessCheck(ctx, request)
	if err != nil {
		return err
	}

	s.Res = response.FsuReadinessCheck
	return nil
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_fleet_software_update.DeleteFsuReadinessCheckRequest{}

	tmp := s.D.Id()
	request.FsuReadinessCheckId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	response, err := s.Client.DeleteFsuReadinessCheck(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fsuReadinessCheckWaitForWorkRequest(ctx, workId, "fsureadinesscheck",
		oci_fleet_software_update.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	switch v := s.Res.(type) {
	case oci_fleet_software_update.TargetFsuReadinessCheck:
		s.D.Set("type", "TARGET")

		targets := []interface{}{}
		for _, item := range v.Targets {
			targets = append(targets, ReadinessCheckTargetEntryToMap(item))
		}
		s.D.Set("targets", targets)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			// s.D.Set("id", *v.Id)
			s.D.SetId(*v.Id)
		}

		if v.IssueCount != nil {
			s.D.Set("issue_count", *v.IssueCount)
		}

		issues := []interface{}{}
		for _, item := range v.Issues {
			issues = append(issues, PatchingIssueEntryToMap(item))
		}
		s.D.Set("issues", issues)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeFinished != nil {
			s.D.Set("time_finished", v.TimeFinished.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", s.Res)
		return nil
	}
	return nil
}

func FsuReadinessCheckSummaryToMap(obj oci_fleet_software_update.FsuReadinessCheckSummary) map[string]interface{} {
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

	if obj.IssueCount != nil {
		result["issue_count"] = int(*obj.IssueCount)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["type"] = string(obj.Type)

	return result
}

func PatchingIssueEntryToMap(obj oci_fleet_software_update.PatchingIssueEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.ImpactedResourceId != nil {
		result["impacted_resource_id"] = string(*obj.ImpactedResourceId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.RecommendedAction != nil {
		result["recommended_action"] = string(*obj.RecommendedAction)
	}

	return result
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) mapToReadinessCheckTargetEntry(fieldKeyFormat string) (oci_fleet_software_update.ReadinessCheckTargetEntry, error) {
	result := oci_fleet_software_update.ReadinessCheckTargetEntry{}

	if entityId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_id")); ok {
		tmp := entityId.(string)
		result.EntityId = &tmp
	}

	if entityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entity_type")); ok {
		result.EntityType = oci_fleet_software_update.ReadinessCheckTargetEntryEntityTypeEnum(entityType.(string))
	}

	return result, nil
}

func ReadinessCheckTargetEntryToMap(obj oci_fleet_software_update.ReadinessCheckTargetEntry) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EntityId != nil {
		result["entity_id"] = string(*obj.EntityId)
	}

	result["entity_type"] = string(obj.EntityType)

	return result
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) populateTopLevelPolymorphicCreateFsuReadinessCheckRequest(request *oci_fleet_software_update.CreateFsuReadinessCheckRequest) error {
	//discriminator
	typeRaw, ok := s.D.GetOkExists("type")
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TARGET"):
		details := oci_fleet_software_update.CreateTargetFsuReadinessCheckDetails{}
		if targets, ok := s.D.GetOkExists("targets"); ok {
			interfaces := targets.([]interface{})
			tmp := make([]oci_fleet_software_update.ReadinessCheckTargetEntry, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "targets", stateDataIndex)
				converted, err := s.mapToReadinessCheckTargetEntry(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("targets") {
				details.Targets = tmp
			}
		}
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateFsuReadinessCheckDetails = details
	default:
		return fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return nil
}

func (s *FleetSoftwareUpdateFsuReadinessCheckResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_fleet_software_update.ChangeFsuReadinessCheckCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FsuReadinessCheckId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_software_update")

	_, err := s.Client.ChangeFsuReadinessCheckCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
