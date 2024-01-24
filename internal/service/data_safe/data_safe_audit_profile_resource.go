// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"
)

func DataSafeAuditProfileResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataSafeAuditProfile,
		Read:     readDataSafeAuditProfile,
		Update:   updateDataSafeAuditProfile,
		Delete:   deleteDataSafeAuditProfile,
		Schema: map[string]*schema.Schema{
			// Required
			"audit_profile_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"compartment_id": {
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
			"is_paid_usage_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"change_retention_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"audit_collected_volume": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"audit_trails": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"audit_collection_start_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"audit_profile_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"freeform_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_auto_purge_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_target_database_key": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"purge_job_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_job_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"purge_job_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_collected": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trail_location": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trail_source": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"work_request_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_override_global_retention_setting": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"offline_months": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"online_months": {
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
			"target_id": {
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

func createDataSafeAuditProfile(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()
	compartment, ok := sync.D.GetOkExists("compartment_id")

	err := tfresource.CreateResource(d, sync)
	if err != nil {
		return err
	}

	if _, ok := sync.D.GetOkExists("change_retention_trigger"); ok {
		err := sync.ChangeRetention()
		if err != nil {
			return err
		}
	}

	if ok && compartment != *sync.Res.CompartmentId {
		err = sync.updateCompartment(compartment)
		if err != nil {
			return err
		}
		tmp := compartment.(string)
		sync.Res.CompartmentId = &tmp
		err := sync.Get()
		if err != nil {
			log.Printf("error doing a Get() after compartment update: %v", err)
		}
		err = sync.SetData()
		if err != nil {
			log.Printf("error doing a SetData() after compartment update: %v", err)
		}
	}
	return nil
}

func readDataSafeAuditProfile(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

func updateDataSafeAuditProfile(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeAuditProfileResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	if _, ok := sync.D.GetOkExists("change_retention_trigger"); ok && sync.D.HasChange("change_retention_trigger") {
		oldRaw, newRaw := sync.D.GetChange("change_retention_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ChangeRetention()

			if err != nil {
				return err
			}
		} else {
			sync.D.Set("change_retention_trigger", oldRaw)
			return fmt.Errorf("new value of trigger should be greater than the old value")
		}
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	return nil
}

func deleteDataSafeAuditProfile(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DataSafeAuditProfileResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_data_safe.DataSafeClient
	Res                    *oci_data_safe.AuditProfile
	DisableNotFoundRetries bool
}

func (s *DataSafeAuditProfileResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataSafeAuditProfileResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateCreating),
	}
}

func (s *DataSafeAuditProfileResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateActive),
		string(oci_data_safe.AuditProfileLifecycleStateNeedsAttention),
	}
}

func (s *DataSafeAuditProfileResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateDeleting),
	}
}

func (s *DataSafeAuditProfileResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_data_safe.AuditProfileLifecycleStateDeleted),
	}
}

func (s *DataSafeAuditProfileResourceCrud) Create() error {
	request := oci_data_safe.UpdateAuditProfileRequest{}

	if auditProfileId, ok := s.D.GetOkExists("audit_profile_id"); ok {
		tmp := auditProfileId.(string)
		request.AuditProfileId = &tmp
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

	if isPaidUsageEnabled, ok := s.D.GetOkExists("is_paid_usage_enabled"); ok {
		tmp := isPaidUsageEnabled.(bool)
		request.IsPaidUsageEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	workRequestResponse := oci_data_safe.GetWorkRequestResponse{}
	workRequestResponse, err = s.Client.GetWorkRequest(context.Background(),
		oci_data_safe.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"),
			},
		})
	if err == nil {
		// The work request response contains an array of objects
		for _, res := range workRequestResponse.Resources {
			if res.EntityType != nil && strings.Contains(strings.ToLower(*res.EntityType), "auditprofile") && res.Identifier != nil {
				s.D.SetId(*res.Identifier)
				break
			}
		}
	}
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DataSafeAuditProfileResourceCrud) getAuditProfileFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_data_safe.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	auditProfileId, err := auditProfileWaitForWorkRequest(workId, "auditprofile",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*auditProfileId)

	return s.Get()
}

func auditProfileWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "data_safe", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_data_safe.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func auditProfileWaitForWorkRequest(wId *string, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_data_safe.DataSafeClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "data_safe")
	retryPolicy.ShouldRetryOperation = auditProfileWorkRequestShouldRetryFunc(timeout)

	response := oci_data_safe.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_data_safe.WorkRequestStatusInProgress),
			string(oci_data_safe.WorkRequestStatusAccepted),
			string(oci_data_safe.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_data_safe.WorkRequestStatusSucceeded),
			string(oci_data_safe.WorkRequestStatusFailed),
			string(oci_data_safe.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_data_safe.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_data_safe.WorkRequestStatusFailed || response.Status == oci_data_safe.WorkRequestStatusCanceled {
		return nil, getErrorFromDataSafeAuditProfileWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataSafeAuditProfileWorkRequest(client *oci_data_safe.DataSafeClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_data_safe.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_data_safe.ListWorkRequestErrorsRequest{
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

func (s *DataSafeAuditProfileResourceCrud) Get() error {
	request := oci_data_safe.GetAuditProfileRequest{}

	tmp := s.D.Id()
	request.AuditProfileId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.GetAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AuditProfile
	return nil
}

func (s *DataSafeAuditProfileResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_data_safe.UpdateAuditProfileRequest{}

	tmp := s.D.Id()
	request.AuditProfileId = &tmp

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

	if isPaidUsageEnabled, ok := s.D.GetOkExists("is_paid_usage_enabled"); ok {
		tmp := isPaidUsageEnabled.(bool)
		request.IsPaidUsageEnabled = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.UpdateAuditProfile(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataSafeAuditProfileResourceCrud) SetData() error {
	if s.Res.AuditCollectedVolume != nil {
		s.D.Set("audit_collected_volume", strconv.FormatInt(*s.Res.AuditCollectedVolume, 10))
	}

	auditTrails := []interface{}{}
	for _, item := range s.Res.AuditTrails {
		auditTrails = append(auditTrails, AuditTrailToMap(item))
	}
	s.D.Set("audit_trails", auditTrails)

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

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsOverrideGlobalRetentionSetting != nil {
		s.D.Set("is_override_global_retention_setting", *s.Res.IsOverrideGlobalRetentionSetting)
	}

	if s.Res.IsPaidUsageEnabled != nil {
		s.D.Set("is_paid_usage_enabled", *s.Res.IsPaidUsageEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OfflineMonths != nil {
		s.D.Set("offline_months", *s.Res.OfflineMonths)
	}

	if s.Res.OnlineMonths != nil {
		s.D.Set("online_months", *s.Res.OnlineMonths)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataSafeAuditProfileResourceCrud) ChangeRetention() error {
	request := oci_data_safe.ChangeRetentionRequest{}

	idTmp := s.D.Id()
	request.AuditProfileId = &idTmp

	if isOverrideGlobalRetentionSetting, ok := s.D.GetOkExists("is_override_global_retention_setting"); ok {
		tmp := isOverrideGlobalRetentionSetting.(bool)
		request.IsOverrideGlobalRetentionSetting = &tmp
	}

	if offlineMonths, ok := s.D.GetOkExists("offline_months"); ok {
		tmp := offlineMonths.(int)
		request.OfflineMonths = &tmp
	}

	if onlineMonths, ok := s.D.GetOkExists("online_months"); ok {
		tmp := onlineMonths.(int)
		request.OnlineMonths = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeRetention(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("change_retention_trigger")
	s.D.Set("change_retention_trigger", val)

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))

}

func AuditProfileSummaryToMap(obj oci_data_safe.AuditProfileSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditCollectedVolume != nil {
		result["audit_collected_volume"] = strconv.FormatInt(*obj.AuditCollectedVolume, 10)
	}

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

	if obj.IsOverrideGlobalRetentionSetting != nil {
		result["is_override_global_retention_setting"] = bool(*obj.IsOverrideGlobalRetentionSetting)
	}

	if obj.IsPaidUsageEnabled != nil {
		result["is_paid_usage_enabled"] = bool(*obj.IsPaidUsageEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.OfflineMonths != nil {
		result["offline_months"] = int(*obj.OfflineMonths)
	}

	if obj.OnlineMonths != nil {
		result["online_months"] = int(*obj.OnlineMonths)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func AuditTrailToMap(obj oci_data_safe.AuditTrail) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AuditCollectionStartTime != nil {
		result["audit_collection_start_time"] = obj.AuditCollectionStartTime.String()
	}

	if obj.AuditProfileId != nil {
		result["audit_profile_id"] = string(*obj.AuditProfileId)
	}

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

	if obj.IsAutoPurgeEnabled != nil {
		result["is_auto_purge_enabled"] = bool(*obj.IsAutoPurgeEnabled)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PeerTargetDatabaseKey != nil {
		result["peer_target_database_key"] = int(*obj.PeerTargetDatabaseKey)
	}

	if obj.PurgeJobDetails != nil {
		result["purge_job_details"] = string(*obj.PurgeJobDetails)
	}

	result["purge_job_status"] = string(obj.PurgeJobStatus)

	if obj.PurgeJobTime != nil {
		result["purge_job_time"] = obj.PurgeJobTime.String()
	}

	result["state"] = string(obj.LifecycleState)

	result["status"] = string(obj.Status)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastCollected != nil {
		result["time_last_collected"] = obj.TimeLastCollected.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TrailLocation != nil {
		result["trail_location"] = string(*obj.TrailLocation)
	}

	result["trail_source"] = string(obj.TrailSource)

	if obj.WorkRequestId != nil {
		result["work_request_id"] = string(*obj.WorkRequestId)
	}

	return result
}

func (s *DataSafeAuditProfileResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_data_safe.ChangeAuditProfileCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AuditProfileId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe")

	response, err := s.Client.ChangeAuditProfileCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAuditProfileFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "data_safe"), oci_data_safe.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
