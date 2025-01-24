// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func RecoveryLongTermBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("24h"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("24h"),
		},
		Create: createRecoveryLongTermBackup,
		Read:   readRecoveryLongTermBackup,
		Update: updateRecoveryLongTermBackup,
		Delete: deleteRecoveryLongTermBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"protected_database_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"retention_period": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      retentionPeriodHashCodeForSets,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"retention_count": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"retention_period_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
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
			"retention_point_in_time": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
			},
			"retention_scn": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_substate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retention_until_date_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rman_tag": {
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
			"time_backup_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_backup_initiated": {
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

func createRecoveryLongTermBackup(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryLongTermBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readRecoveryLongTermBackup(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryLongTermBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateRecoveryLongTermBackup(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryLongTermBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRecoveryLongTermBackup(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryLongTermBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RecoveryLongTermBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_recovery.DatabaseRecoveryClient
	Res                    *oci_recovery.LongTermBackup
	DisableNotFoundRetries bool
}

func (s *RecoveryLongTermBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RecoveryLongTermBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_recovery.LongTermBackupLifecycleStateCreating),
	}
}

func (s *RecoveryLongTermBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_recovery.LongTermBackupLifecycleStateActive),
	}
}

func (s *RecoveryLongTermBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_recovery.LongTermBackupLifecycleStateDeleting),
	}
}

func (s *RecoveryLongTermBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_recovery.LongTermBackupLifecycleStateDeleted),
	}
}

func (s *RecoveryLongTermBackupResourceCrud) Create() error {
	request := oci_recovery.CreateLongTermBackupRequest{}

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

	if protectedDatabaseId, ok := s.D.GetOkExists("protected_database_id"); ok {
		tmp := protectedDatabaseId.(string)
		request.ProtectedDatabaseId = &tmp
	}

	if retentionPeriod, ok := s.D.GetOkExists("retention_period"); ok {
		set := retentionPeriod.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_recovery.RetentionPeriodValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := retentionPeriodHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "retention_period", stateDataIndex)
			converted, err := s.mapToRetentionPeriodValue(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("retention_period") {
			request.RetentionPeriod = tmp
		}
	}

	if retentionPointInTime, ok := s.D.GetOkExists("retention_point_in_time"); ok {
		tmp, err := time.Parse(time.RFC3339, retentionPointInTime.(string))
		if err != nil {
			return err
		}
		request.RetentionPointInTime = &oci_common.SDKTime{Time: tmp}
	}

	if retentionScn, ok := s.D.GetOkExists("retention_scn"); ok {
		tmp := retentionScn.(int)
		request.RetentionScn = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.CreateLongTermBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getLongTermBackupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RecoveryLongTermBackupResourceCrud) getLongTermBackupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	longTermBackupId, err := longTermBackupWaitForWorkRequest(workId, "longtermbackup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*longTermBackupId)

	return s.Get()
}

func longTermBackupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "recovery", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_recovery.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func longTermBackupWaitForWorkRequest(wId *string, entityType string, action oci_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_recovery.DatabaseRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "recovery")
	retryPolicy.ShouldRetryOperation = longTermBackupWorkRequestShouldRetryFunc(timeout)

	response := oci_recovery.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_recovery.OperationStatusInProgress),
			string(oci_recovery.OperationStatusAccepted),
			string(oci_recovery.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_recovery.OperationStatusSucceeded),
			string(oci_recovery.OperationStatusFailed),
			string(oci_recovery.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_recovery.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_recovery.OperationStatusFailed || response.Status == oci_recovery.OperationStatusCanceled {
		return nil, getErrorFromRecoveryLongTermBackupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRecoveryLongTermBackupWorkRequest(client *oci_recovery.DatabaseRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_recovery.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_recovery.ListWorkRequestErrorsRequest{
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

func (s *RecoveryLongTermBackupResourceCrud) Get() error {
	request := oci_recovery.GetLongTermBackupRequest{}

	tmp := s.D.Id()
	request.LongTermBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.GetLongTermBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LongTermBackup
	return nil
}

func (s *RecoveryLongTermBackupResourceCrud) Update() error {
	request := oci_recovery.UpdateLongTermBackupRequest{}

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
	request.LongTermBackupId = &tmp

	if retentionPeriod, ok := s.D.GetOkExists("retention_period"); ok {
		set := retentionPeriod.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_recovery.RetentionPeriodValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := retentionPeriodHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "retention_period", stateDataIndex)
			converted, err := s.mapToRetentionPeriodValue(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("retention_period") {
			request.RetentionPeriod = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.UpdateLongTermBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLongTermBackupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RecoveryLongTermBackupResourceCrud) Delete() error {
	request := oci_recovery.DeleteLongTermBackupRequest{}

	tmp := s.D.Id()
	request.LongTermBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.DeleteLongTermBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := longTermBackupWaitForWorkRequest(workId, "longtermbackup",
		oci_recovery.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *RecoveryLongTermBackupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseIdentifier != nil {
		s.D.Set("database_identifier", *s.Res.DatabaseIdentifier)
	}

	if s.Res.DatabaseSizeInGBs != nil {
		s.D.Set("database_size_in_gbs", *s.Res.DatabaseSizeInGBs)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_substate", s.Res.LifecycleSubstate)

	if s.Res.ProtectedDatabaseId != nil {
		s.D.Set("protected_database_id", *s.Res.ProtectedDatabaseId)
	}

	retentionPeriod := []interface{}{}
	for _, item := range s.Res.RetentionPeriod {
		retentionPeriod = append(retentionPeriod, RetentionPeriodValueToMap(item))
	}
	s.D.Set("retention_period", schema.NewSet(retentionPeriodHashCodeForSets, retentionPeriod))

	if s.Res.RetentionPointInTime != nil {
		s.D.Set("retention_point_in_time", s.Res.RetentionPointInTime.Format(time.RFC3339Nano))
	}

	if s.Res.RetentionScn != nil {
		s.D.Set("retention_scn", *s.Res.RetentionScn)
	}

	if s.Res.RetentionUntilDateTime != nil {
		s.D.Set("retention_until_date_time", s.Res.RetentionUntilDateTime.String())
	}

	if s.Res.RmanTag != nil {
		s.D.Set("rman_tag", *s.Res.RmanTag)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBackupCompleted != nil {
		s.D.Set("time_backup_completed", s.Res.TimeBackupCompleted.String())
	}

	if s.Res.TimeBackupInitiated != nil {
		s.D.Set("time_backup_initiated", s.Res.TimeBackupInitiated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func LongTermBackupSummaryToMap(obj oci_recovery.LongTermBackupSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseIdentifier != nil {
		result["database_identifier"] = string(*obj.DatabaseIdentifier)
	}

	if obj.DatabaseSizeInGBs != nil {
		result["database_size_in_gbs"] = int(*obj.DatabaseSizeInGBs)
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["lifecycle_substate"] = string(obj.LifecycleSubstate)

	if obj.ProtectedDatabaseId != nil {
		result["protected_database_id"] = string(*obj.ProtectedDatabaseId)
	}

	retentionPeriod := []interface{}{}
	for _, item := range obj.RetentionPeriod {
		retentionPeriod = append(retentionPeriod, RetentionPeriodValueToMap(item))
	}
	if datasource {
		result["retention_period"] = retentionPeriod
	} else {
		result["retention_period"] = schema.NewSet(retentionPeriodHashCodeForSets, retentionPeriod)
	}

	if obj.RetentionPointInTime != nil {
		result["retention_point_in_time"] = obj.RetentionPointInTime.String()
	}

	if obj.RetentionScn != nil {
		result["retention_scn"] = int(*obj.RetentionScn)
	}

	if obj.RetentionUntilDateTime != nil {
		result["retention_until_date_time"] = obj.RetentionUntilDateTime.String()
	}

	if obj.RmanTag != nil {
		result["rman_tag"] = string(*obj.RmanTag)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeBackupCompleted != nil {
		result["time_backup_completed"] = obj.TimeBackupCompleted.String()
	}

	if obj.TimeBackupInitiated != nil {
		result["time_backup_initiated"] = obj.TimeBackupInitiated.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *RecoveryLongTermBackupResourceCrud) mapToRetentionPeriodValue(fieldKeyFormat string) (oci_recovery.RetentionPeriodValue, error) {
	result := oci_recovery.RetentionPeriodValue{}

	if retentionCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_count")); ok {
		tmp := retentionCount.(int)
		result.RetentionCount = &tmp
	}

	if retentionPeriodType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retention_period_type")); ok {
		result.RetentionPeriodType = oci_recovery.RetentionPeriodValueRetentionPeriodTypeEnum(retentionPeriodType.(string))
	}

	return result, nil
}

func RetentionPeriodValueToMap(obj oci_recovery.RetentionPeriodValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RetentionCount != nil {
		result["retention_count"] = int(*obj.RetentionCount)
	}

	result["retention_period_type"] = string(obj.RetentionPeriodType)

	return result
}

func retentionPeriodHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if retentionCount, ok := m["retention_count"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", retentionCount))
	}
	if retentionPeriodType, ok := m["retention_period_type"]; ok && retentionPeriodType != "" {
		buf.WriteString(fmt.Sprintf("%v-", retentionPeriodType))
	}
	return utils.GetStringHashcode(buf.String())
}
