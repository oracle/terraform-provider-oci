// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fusion_apps

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fusion_apps "github.com/oracle/oci-go-sdk/v65/fusionapps"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FusionAppsFusionEnvironmentFamilyResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFusionAppsFusionEnvironmentFamily,
		Read:     readFusionAppsFusionEnvironmentFamily,
		Update:   updateFusionAppsFusionEnvironmentFamily,
		Delete:   deleteFusionAppsFusionEnvironmentFamily,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_ids": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"family_maintenance_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"concurrent_maintenance": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_monthly_patching_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"quarterly_upgrade_begin_times": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"is_subscription_update_needed": {
				Type:     schema.TypeBool,
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
			"system_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFusionAppsFusionEnvironmentFamily(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.CreateResource(d, sync)
}

func readFusionAppsFusionEnvironmentFamily(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.ReadResource(sync)
}

func updateFusionAppsFusionEnvironmentFamily(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFusionAppsFusionEnvironmentFamily(d *schema.ResourceData, m interface{}) error {
	sync := &FusionAppsFusionEnvironmentFamilyResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FusionApplicationsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FusionAppsFusionEnvironmentFamilyResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fusion_apps.FusionApplicationsClient
	Res                    *oci_fusion_apps.FusionEnvironmentFamily
	DisableNotFoundRetries bool
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateCreating),
	}
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateActive),
	}
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateDeleting),
	}
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fusion_apps.FusionEnvironmentFamilyLifecycleStateDeleted),
	}
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) Create() error {
	request := oci_fusion_apps.CreateFusionEnvironmentFamilyRequest{}

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

	if familyMaintenancePolicy, ok := s.D.GetOkExists("family_maintenance_policy"); ok {
		if tmpList := familyMaintenancePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "family_maintenance_policy", 0)
			tmp, err := s.mapToFamilyMaintenancePolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FamilyMaintenancePolicy = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if subscriptionIds, ok := s.D.GetOkExists("subscription_ids"); ok {
		interfaces := subscriptionIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subscription_ids") {
			request.SubscriptionIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.CreateFusionEnvironmentFamily(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_fusion_apps.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_fusion_apps.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "fusionenvironmentfamily") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getFusionEnvironmentFamilyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) getFusionEnvironmentFamilyFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fusion_apps.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	fusionEnvironmentFamilyId, err := fusionEnvironmentFamilyWaitForWorkRequest(workId, "fusionenvironmentfamily",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*fusionEnvironmentFamilyId)

	return s.Get()
}

func fusionEnvironmentFamilyWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fusion_apps", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fusion_apps.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func fusionEnvironmentFamilyWaitForWorkRequest(wId *string, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fusion_apps.FusionApplicationsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fusion_apps")
	retryPolicy.ShouldRetryOperation = fusionEnvironmentFamilyWorkRequestShouldRetryFunc(timeout)

	response := oci_fusion_apps.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fusion_apps.WorkRequestStatusInProgress),
			string(oci_fusion_apps.WorkRequestStatusAccepted),
			string(oci_fusion_apps.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_fusion_apps.WorkRequestStatusSucceeded),
			string(oci_fusion_apps.WorkRequestStatusFailed),
			string(oci_fusion_apps.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fusion_apps.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fusion_apps.WorkRequestStatusFailed || response.Status == oci_fusion_apps.WorkRequestStatusCanceled {
		return nil, getErrorFromFusionAppsFusionEnvironmentFamilyWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFusionAppsFusionEnvironmentFamilyWorkRequest(client *oci_fusion_apps.FusionApplicationsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fusion_apps.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fusion_apps.ListWorkRequestErrorsRequest{
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

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) Get() error {
	request := oci_fusion_apps.GetFusionEnvironmentFamilyRequest{}

	tmp := s.D.Id()
	request.FusionEnvironmentFamilyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.GetFusionEnvironmentFamily(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.FusionEnvironmentFamily
	return nil
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_fusion_apps.UpdateFusionEnvironmentFamilyRequest{}

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

	if familyMaintenancePolicy, ok := s.D.GetOkExists("family_maintenance_policy"); ok {
		if tmpList := familyMaintenancePolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "family_maintenance_policy", 0)
			tmp, err := s.mapToUpdateFamilyMaintenancePolicyDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.FamilyMaintenancePolicy = &tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.FusionEnvironmentFamilyId = &tmp

	if subscriptionIds, ok := s.D.GetOkExists("subscription_ids"); ok {
		interfaces := subscriptionIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("subscription_ids") {
			request.SubscriptionIds = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.UpdateFusionEnvironmentFamily(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFusionEnvironmentFamilyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) Delete() error {
	request := oci_fusion_apps.DeleteFusionEnvironmentFamilyRequest{}

	tmp := s.D.Id()
	request.FusionEnvironmentFamilyId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.DeleteFusionEnvironmentFamily(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := fusionEnvironmentFamilyWaitForWorkRequest(workId, "fusionenvironmentfamily",
		oci_fusion_apps.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FamilyMaintenancePolicy != nil {
		s.D.Set("family_maintenance_policy", []interface{}{FamilyMaintenancePolicyToMap(s.Res.FamilyMaintenancePolicy)})
	} else {
		s.D.Set("family_maintenance_policy", nil)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsSubscriptionUpdateNeeded != nil {
		s.D.Set("is_subscription_update_needed", *s.Res.IsSubscriptionUpdateNeeded)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subscription_ids", s.Res.SubscriptionIds)

	if s.Res.SystemName != nil {
		s.D.Set("system_name", *s.Res.SystemName)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func FamilyMaintenancePolicyToMap(obj *oci_fusion_apps.FamilyMaintenancePolicy) map[string]interface{} {
	result := map[string]interface{}{}

	result["concurrent_maintenance"] = string(obj.ConcurrentMaintenance)

	if obj.IsMonthlyPatchingEnabled != nil {
		result["is_monthly_patching_enabled"] = bool(*obj.IsMonthlyPatchingEnabled)
	}

	if obj.QuarterlyUpgradeBeginTimes != nil {
		result["quarterly_upgrade_begin_times"] = string(*obj.QuarterlyUpgradeBeginTimes)
	}

	return result
}

func FusionEnvironmentFamilySummaryToMap(obj oci_fusion_apps.FusionEnvironmentFamilySummary) map[string]interface{} {
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

	if obj.FamilyMaintenancePolicy != nil {
		result["family_maintenance_policy"] = []interface{}{FamilyMaintenancePolicyToMap(obj.FamilyMaintenancePolicy)}
	}

	result["freeform_tags"] = obj.FreeformTags
	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsSubscriptionUpdateNeeded != nil {
		result["is_subscription_update_needed"] = bool(*obj.IsSubscriptionUpdateNeeded)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	result["subscription_ids"] = obj.SubscriptionIds

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_fusion_apps.ChangeFusionEnvironmentFamilyCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.FusionEnvironmentFamilyId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps")

	response, err := s.Client.ChangeFusionEnvironmentFamilyCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getFusionEnvironmentFamilyFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fusion_apps"), oci_fusion_apps.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) mapToFamilyMaintenancePolicy(fieldKeyFormat string) (oci_fusion_apps.FamilyMaintenancePolicy, error) {
	result := oci_fusion_apps.FamilyMaintenancePolicy{}

	if quarterlyUpgradeBeginTimes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "quarterly_upgrade_begin_times")); ok {
		tmp := quarterlyUpgradeBeginTimes.(string)
		result.QuarterlyUpgradeBeginTimes = &tmp
	}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
	}

	if concurrentMaintenance, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "concurrent_maintenance")); ok {
		result.ConcurrentMaintenance = oci_fusion_apps.FamilyMaintenancePolicyConcurrentMaintenanceEnum(concurrentMaintenance.(string))
	}

	return result, nil
}

func (s *FusionAppsFusionEnvironmentFamilyResourceCrud) mapToUpdateFamilyMaintenancePolicyDetails(fieldKeyFormat string) (oci_fusion_apps.UpdateFamilyMaintenancePolicyDetails, error) {
	result := oci_fusion_apps.UpdateFamilyMaintenancePolicyDetails{}

	if isMonthlyPatchingEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_monthly_patching_enabled")); ok {
		tmp := isMonthlyPatchingEnabled.(bool)
		result.IsMonthlyPatchingEnabled = &tmp
	}

	if concurrentMaintenance, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "concurrent_maintenance")); ok {
		result.ConcurrentMaintenance = oci_fusion_apps.FamilyMaintenancePolicyConcurrentMaintenanceEnum(concurrentMaintenance.(string))
	}

	return result, nil
}
