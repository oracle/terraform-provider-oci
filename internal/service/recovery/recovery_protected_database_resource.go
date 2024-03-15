// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package recovery

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_recovery "github.com/oracle/oci-go-sdk/v65/recovery"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RecoveryProtectedDatabaseResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createRecoveryProtectedDatabase,
		Read:     readRecoveryProtectedDatabase,
		Update:   updateRecoveryProtectedDatabase,
		Delete:   deleteRecoveryProtectedDatabase,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_unique_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"protection_policy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"recovery_service_subnets": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"recovery_service_subnet_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			// Optional
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"database_size": {
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
			"deletion_schedule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_recovery.DeleteProtectedDatabaseDeletionSchedule72Hours),
					string(oci_recovery.DeleteProtectedDatabaseDeletionScheduleRetentionPeriod),
				}, true),
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"is_redo_logs_shipped": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"health": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_read_only_resource": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metrics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_space_estimate_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"backup_space_used_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"current_retention_period_in_seconds": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"db_size_in_gbs": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"is_redo_logs_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"minimum_recovery_needed_in_days": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"retention_period_in_days": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"unprotected_window_in_seconds": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"policy_locked_date_time": {
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
			"vpc_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createRecoveryProtectedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.CreateResource(d, sync)
}

func readRecoveryProtectedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.ReadResource(sync)
}

func updateRecoveryProtectedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteRecoveryProtectedDatabase(d *schema.ResourceData, m interface{}) error {
	sync := &RecoveryProtectedDatabaseResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseRecoveryClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type RecoveryProtectedDatabaseResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_recovery.DatabaseRecoveryClient
	Res                    *oci_recovery.ProtectedDatabase
	DisableNotFoundRetries bool
}

func (s *RecoveryProtectedDatabaseResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RecoveryProtectedDatabaseResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateCreating),
	}
}

func (s *RecoveryProtectedDatabaseResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateActive),
	}
}

func (s *RecoveryProtectedDatabaseResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleting),
	}
}

func (s *RecoveryProtectedDatabaseResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_recovery.LifecycleStateDeleteScheduled),
		string(oci_recovery.LifecycleStateDeleted),
	}
}

func (s *RecoveryProtectedDatabaseResourceCrud) Create() error {
	request := oci_recovery.CreateProtectedDatabaseRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	if databaseSize, ok := s.D.GetOkExists("database_size"); ok {
		request.DatabaseSize = oci_recovery.DatabaseSizesEnum(databaseSize.(string))
	}

	if dbUniqueName, ok := s.D.GetOkExists("db_unique_name"); ok {
		tmp := dbUniqueName.(string)
		request.DbUniqueName = &tmp
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

	if isRedoLogsShipped, ok := s.D.GetOkExists("is_redo_logs_shipped"); ok {
		tmp := isRedoLogsShipped.(bool)
		request.IsRedoLogsShipped = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok {
		tmp := password.(string)
		request.Password = &tmp
	}

	if protectionPolicyId, ok := s.D.GetOkExists("protection_policy_id"); ok {
		tmp := protectionPolicyId.(string)
		request.ProtectionPolicyId = &tmp
	}

	if recoveryServiceSubnets, ok := s.D.GetOkExists("recovery_service_subnets"); ok {
		interfaces := recoveryServiceSubnets.([]interface{})
		tmp := make([]oci_recovery.RecoveryServiceSubnetInput, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "recovery_service_subnets", stateDataIndex)
			converted, err := s.mapToRecoveryServiceSubnetInput(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("recovery_service_subnets") {
			request.RecoveryServiceSubnets = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.CreateProtectedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getProtectedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RecoveryProtectedDatabaseResourceCrud) getProtectedDatabaseFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_recovery.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	protectedDatabaseId, err := protectedDatabaseWaitForWorkRequest(workId, "protecteddatabase",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*protectedDatabaseId)

	return s.Get()
}

func protectedDatabaseWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func protectedDatabaseWaitForWorkRequest(wId *string, entityType string, action oci_recovery.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_recovery.DatabaseRecoveryClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "recovery")
	retryPolicy.ShouldRetryOperation = protectedDatabaseWorkRequestShouldRetryFunc(timeout)

	response := oci_recovery.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
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
		return nil, getErrorFromRecoveryProtectedDatabaseWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRecoveryProtectedDatabaseWorkRequest(client *oci_recovery.DatabaseRecoveryClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_recovery.ActionTypeEnum) error {
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

func (s *RecoveryProtectedDatabaseResourceCrud) Get() error {
	request := oci_recovery.GetProtectedDatabaseRequest{}

	tmp := s.D.Id()
	request.ProtectedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.GetProtectedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ProtectedDatabase
	return nil
}

func (s *RecoveryProtectedDatabaseResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_recovery.UpdateProtectedDatabaseRequest{}

	if databaseSize, ok := s.D.GetOkExists("database_size"); ok {
		request.DatabaseSize = oci_recovery.DatabaseSizesEnum(databaseSize.(string))
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

	if isRedoLogsShipped, ok := s.D.GetOkExists("is_redo_logs_shipped"); ok {
		tmp := isRedoLogsShipped.(bool)
		request.IsRedoLogsShipped = &tmp
	}

	if password, ok := s.D.GetOkExists("password"); ok && s.D.HasChange("password") {
		tmp := password.(string)
		request.Password = &tmp
	}

	tmp := s.D.Id()
	request.ProtectedDatabaseId = &tmp

	if protectionPolicyId, ok := s.D.GetOkExists("protection_policy_id"); ok {
		tmp := protectionPolicyId.(string)
		request.ProtectionPolicyId = &tmp
	}

	if recoveryServiceSubnets, ok := s.D.GetOkExists("recovery_service_subnets"); ok {
		interfaces := recoveryServiceSubnets.([]interface{})
		tmp := make([]oci_recovery.RecoveryServiceSubnetInput, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "recovery_service_subnets", stateDataIndex)
			converted, err := s.mapToRecoveryServiceSubnetInput(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("recovery_service_subnets") {
			request.RecoveryServiceSubnets = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.UpdateProtectedDatabase(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProtectedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RecoveryProtectedDatabaseResourceCrud) Delete() error {
	request := oci_recovery.ScheduleProtectedDatabaseDeletionRequest{}

	if deletionSchedule, ok := s.D.GetOkExists("deletion_schedule"); ok {
		request.DeletionSchedule = oci_recovery.DeletionScheduleEnum(deletionSchedule.(string))
	}

	tmp := s.D.Id()
	request.ProtectedDatabaseId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	_, err := s.Client.ScheduleProtectedDatabaseDeletion(context.Background(), request)
	return err
}

func (s *RecoveryProtectedDatabaseResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DatabaseId != nil {
		s.D.Set("database_id", *s.Res.DatabaseId)
	}

	s.D.Set("database_size", s.Res.DatabaseSize)

	if s.Res.DbUniqueName != nil {
		s.D.Set("db_unique_name", *s.Res.DbUniqueName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("health", s.Res.Health)

	if s.Res.HealthDetails != nil {
		s.D.Set("health_details", *s.Res.HealthDetails)
	}

	if s.Res.IsReadOnlyResource != nil {
		s.D.Set("is_read_only_resource", *s.Res.IsReadOnlyResource)
	}

	if s.Res.IsRedoLogsShipped != nil {
		s.D.Set("is_redo_logs_shipped", *s.Res.IsRedoLogsShipped)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Metrics != nil {
		s.D.Set("metrics", []interface{}{MetricsToMap(s.Res.Metrics)})
	} else {
		s.D.Set("metrics", nil)
	}

	if s.Res.PolicyLockedDateTime != nil {
		s.D.Set("policy_locked_date_time", *s.Res.PolicyLockedDateTime)
	}

	if s.Res.ProtectionPolicyId != nil {
		s.D.Set("protection_policy_id", *s.Res.ProtectionPolicyId)
	}

	recoveryServiceSubnets := []interface{}{}
	for _, item := range s.Res.RecoveryServiceSubnets {
		recoveryServiceSubnets = append(recoveryServiceSubnets, RecoveryServiceSubnetDetailsToMap(item))
	}
	s.D.Set("recovery_service_subnets", recoveryServiceSubnets)

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

	if s.Res.VpcUserName != nil {
		s.D.Set("vpc_user_name", *s.Res.VpcUserName)
	}

	return nil
}

func MetricsToMap(obj *oci_recovery.Metrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupSpaceEstimateInGBs != nil {
		result["backup_space_estimate_in_gbs"] = float32(*obj.BackupSpaceEstimateInGBs)
	}

	if obj.BackupSpaceUsedInGBs != nil {
		result["backup_space_used_in_gbs"] = float32(*obj.BackupSpaceUsedInGBs)
	}

	if obj.CurrentRetentionPeriodInSeconds != nil {
		result["current_retention_period_in_seconds"] = float32(*obj.CurrentRetentionPeriodInSeconds)
	}

	if obj.DbSizeInGBs != nil {
		result["db_size_in_gbs"] = float32(*obj.DbSizeInGBs)
	}

	if obj.IsRedoLogsEnabled != nil {
		result["is_redo_logs_enabled"] = bool(*obj.IsRedoLogsEnabled)
	}

	if obj.MinimumRecoveryNeededInDays != nil {
		result["minimum_recovery_needed_in_days"] = float32(*obj.MinimumRecoveryNeededInDays)
	}

	if obj.RetentionPeriodInDays != nil {
		result["retention_period_in_days"] = float32(*obj.RetentionPeriodInDays)
	}

	if obj.UnprotectedWindowInSeconds != nil {
		result["unprotected_window_in_seconds"] = float32(*obj.UnprotectedWindowInSeconds)
	}

	return result
}

func MetricsSummaryToMap(obj *oci_recovery.MetricsSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupSpaceEstimateInGBs != nil {
		result["backup_space_estimate_in_gbs"] = float32(*obj.BackupSpaceEstimateInGBs)
	}

	if obj.BackupSpaceUsedInGBs != nil {
		result["backup_space_used_in_gbs"] = float32(*obj.BackupSpaceUsedInGBs)
	}

	if obj.CurrentRetentionPeriodInSeconds != nil {
		result["current_retention_period_in_seconds"] = float32(*obj.CurrentRetentionPeriodInSeconds)
	}

	if obj.DbSizeInGBs != nil {
		result["db_size_in_gbs"] = float32(*obj.DbSizeInGBs)
	}

	if obj.IsRedoLogsEnabled != nil {
		result["is_redo_logs_enabled"] = bool(*obj.IsRedoLogsEnabled)
	}

	if obj.MinimumRecoveryNeededInDays != nil {
		result["minimum_recovery_needed_in_days"] = float32(*obj.MinimumRecoveryNeededInDays)
	}

	if obj.RetentionPeriodInDays != nil {
		result["retention_period_in_days"] = float32(*obj.RetentionPeriodInDays)
	}

	if obj.UnprotectedWindowInSeconds != nil {
		result["unprotected_window_in_seconds"] = float32(*obj.UnprotectedWindowInSeconds)
	}

	return result
}

func ProtectedDatabaseSummaryToMap(obj oci_recovery.ProtectedDatabaseSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DatabaseId != nil {
		result["database_id"] = string(*obj.DatabaseId)
	}

	result["database_size"] = string(obj.DatabaseSize)

	if obj.DbUniqueName != nil {
		result["db_unique_name"] = string(*obj.DbUniqueName)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	result["health"] = string(obj.Health)

	if obj.HealthDetails != nil {
		result["health_details"] = string(*obj.HealthDetails)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsReadOnlyResource != nil {
		result["is_read_only_resource"] = bool(*obj.IsReadOnlyResource)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Metrics != nil {
		result["metrics"] = []interface{}{MetricsSummaryToMap(obj.Metrics)}
	}

	if obj.PolicyLockedDateTime != nil {
		result["policy_locked_date_time"] = string(*obj.PolicyLockedDateTime)
	}

	if obj.ProtectionPolicyId != nil {
		result["protection_policy_id"] = string(*obj.ProtectionPolicyId)
	}

	recoveryServiceSubnets := []interface{}{}
	for _, item := range obj.RecoveryServiceSubnets {
		recoveryServiceSubnets = append(recoveryServiceSubnets, RecoveryServiceSubnetDetailsToMap(item))
	}
	result["recovery_service_subnets"] = recoveryServiceSubnets

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

	if obj.VpcUserName != nil {
		result["vpc_user_name"] = string(*obj.VpcUserName)
	}

	return result
}

func RecoveryServiceSubnetDetailsToMap(obj oci_recovery.RecoveryServiceSubnetDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RecoveryServiceSubnetId != nil {
		result["recovery_service_subnet_id"] = string(*obj.RecoveryServiceSubnetId)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func (s *RecoveryProtectedDatabaseResourceCrud) mapToRecoveryServiceSubnetInput(fieldKeyFormat string) (oci_recovery.RecoveryServiceSubnetInput, error) {
	result := oci_recovery.RecoveryServiceSubnetInput{}

	if recoveryServiceSubnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recovery_service_subnet_id")); ok {
		tmp := recoveryServiceSubnetId.(string)
		result.RecoveryServiceSubnetId = &tmp
	}

	return result, nil
}

func (s *RecoveryProtectedDatabaseResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_recovery.ChangeProtectedDatabaseCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ProtectedDatabaseId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery")

	response, err := s.Client.ChangeProtectedDatabaseCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getProtectedDatabaseFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "recovery"), oci_recovery.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
