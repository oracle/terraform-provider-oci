// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package psql

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_psql "github.com/oracle/oci-go-sdk/v65/psql"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func PsqlBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createPsqlBackup,
		Read:     readPsqlBackup,
		Update:   updatePsqlBackup,
		Delete:   deletePsqlBackup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_system_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ForceNew:      true,
				Computed:      true,
				ConflictsWith: []string{"source_backup_details"},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"retention_period": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"backup_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"copy_status": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"backup_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"db_system_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"config_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"system_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_accepted_request_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_completed_request_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_backup_details": {
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"db_system_id"},
				MaxItems:      1,
				MinItems:      1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source_backup_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source_region": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"source_type": {
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
			"time_created_precise": {
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

func createPsqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.CreateResource(d, sync)
}

func readPsqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.ReadResource(sync)
}

func updatePsqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()

	return tfresource.UpdateResource(d, sync)
}

func deletePsqlBackup(d *schema.ResourceData, m interface{}) error {
	sync := &PsqlBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).PostgresqlClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type PsqlBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_psql.PostgresqlClient
	SourceRegionClient     *oci_psql.PostgresqlClient
	Res                    *oci_psql.Backup
	DisableNotFoundRetries bool
}

func (s *PsqlBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *PsqlBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_psql.BackupLifecycleStateCreating),
	}
}

func (s *PsqlBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_psql.BackupLifecycleStateActive),
	}
}

func (s *PsqlBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_psql.BackupLifecycleStateDeleting),
	}
}

func (s *PsqlBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_psql.BackupLifecycleStateDeleted),
	}
}

func (s *PsqlBackupResourceCrud) Create() error {
	if s.isCopyCreate() {
		return s.createBackupCopyPsql()
	}
	return s.createBackupPsql()
}

func (s *PsqlBackupResourceCrud) isCopyCreate() bool {
	if sourceBackupDetails, ok := s.D.GetOkExists("source_backup_details"); ok {
		if tmpList := sourceBackupDetails.([]interface{}); len(tmpList) > 0 {
			return true
		}
	}
	return false
}
func (s *PsqlBackupResourceCrud) createBackupCopyPsql() error {
	copyPsqlBackupRequest := oci_psql.BackupCopyRequest{}

	configProvider := *s.Client.ConfigurationProvider()
	if configProvider == nil {
		return fmt.Errorf("cannot access ConfigurationProvider")
	}
	currentRegion, error := configProvider.Region()
	if error != nil {
		return fmt.Errorf("cannot access Region for the current ConfigurationProvider")
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		copyPsqlBackupRequest.CompartmentId = &tmp
	}

	if sourceBackupDetails, ok := s.D.GetOkExists("source_backup_details"); ok && sourceBackupDetails != nil {
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "source_backup_details", 0)

		if sourceBackupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_backup_id")); ok && sourceBackupId != nil {
			tmp := sourceBackupId.(string)
			copyPsqlBackupRequest.BackupId = &tmp
		}
		result := oci_psql.BackupCopyDetails{}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			result.CompartmentId = &tmp
		}
		if retentionPeriod, ok := s.D.GetOkExists("retention_period"); ok {
			tmp := retentionPeriod.(int)
			result.RetentionPeriod = &tmp
		}

		tmp := make([]string, 1)
		tmp[0] = currentRegion
		result.Regions = tmp

		copyPsqlBackupRequest.BackupCopyDetails = result

		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_region")); ok {
			srcRegion := region.(string)
			err := s.createPsqlSourceRegionClient(srcRegion)
			if err != nil {
				return err
			}
		}
	}

	response, err := s.SourceRegionClient.BackupCopy(context.Background(), copyPsqlBackupRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string

	res, err := s.SourceRegionClient.GetWorkRequest(context.Background(),
		oci_psql.GetWorkRequestRequest{
			WorkRequestId: workId,
		})
	if err == nil {
		identifier = res.Id
	}

	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getBackupFromCopyWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql"), oci_psql.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *PsqlBackupResourceCrud) createBackupPsql() error {
	request := oci_psql.CreateBackupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
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

	if retentionPeriod, ok := s.D.GetOkExists("retention_period"); ok {
		tmp := retentionPeriod.(int)
		request.RetentionPeriod = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.CreateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string

	res, err := s.Client.GetWorkRequest(context.Background(),
		oci_psql.GetWorkRequestRequest{
			WorkRequestId: workId,
		})
	if err == nil {
		identifier = res.Id
	}

	if identifier != nil {
		s.D.SetId(*identifier)
	}

	return s.getBackupFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql"), oci_psql.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *PsqlBackupResourceCrud) getBackupFromCopyWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_psql.ActionTypeEnum, timeout time.Duration) error {

	configProvider := *s.Client.ConfigurationProvider()
	if configProvider == nil {
		return fmt.Errorf("cannot access ConfigurationProvider")
	}
	currentRegion, error := configProvider.Region()
	if error != nil {
		return fmt.Errorf("cannot access Region for the current ConfigurationProvider")
	}

	// Wait until it finishes
	// changes required here for the backup copy request
	backupId, err := backupWaitForWorkRequest(workId, "backup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.SourceRegionClient)

	if err != nil {
		return err
	}
	s.D.SetId(*backupId)
	err = s.GetSource()
	if err != nil {
		return err
	}

	for _, item := range s.Res.CopyStatus {
		if item.Region != nil {
			if item.BackupId != nil {
				if *item.Region == currentRegion {
					s.D.SetId(*item.BackupId)
					err = tfresource.WaitForResourceCondition(s, func() bool { return s.Res.LifecycleState == oci_psql.BackupLifecycleStateActive }, s.D.Timeout(schema.TimeoutCreate))
					if err != nil {
						return err
					}
				}
			}
		}
	}

	err = s.Update()
	if err != nil {
		return err
	}
	return s.Get()
}

func (s *PsqlBackupResourceCrud) getBackupFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_psql.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	backupId, err := backupWaitForWorkRequest(workId, "backup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*backupId)

	return s.Get()
}

func backupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "psql", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_psql.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func backupWaitForWorkRequest(wId *string, entityType string, action oci_psql.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_psql.PostgresqlClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "psql")
	retryPolicy.ShouldRetryOperation = backupWorkRequestShouldRetryFunc(timeout)

	response := oci_psql.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_psql.OperationStatusInProgress),
			string(oci_psql.OperationStatusAccepted),
			string(oci_psql.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_psql.OperationStatusSucceeded),
			string(oci_psql.OperationStatusFailed),
			string(oci_psql.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_psql.GetWorkRequestRequest{
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
			if response.Status == oci_psql.OperationStatusSucceeded {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_psql.OperationStatusFailed || response.Status == oci_psql.OperationStatusCanceled {
		return nil, getErrorFromPsqlBackupWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromPsqlBackupWorkRequest(client *oci_psql.PostgresqlClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_psql.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_psql.ListWorkRequestErrorsRequest{
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

func (s *PsqlBackupResourceCrud) Get() error {
	request := oci_psql.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *PsqlBackupResourceCrud) GetSource() error {
	request := oci_psql.GetBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.SourceRegionClient.GetBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *PsqlBackupResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_psql.UpdateBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

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

	if retentionPeriod, ok := s.D.GetOkExists("retention_period"); ok {
		tmp := retentionPeriod.(int)
		request.RetentionPeriod = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.UpdateBackup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backup
	return nil
}

func (s *PsqlBackupResourceCrud) Delete() error {
	request := oci_psql.DeleteBackupRequest{}

	tmp := s.D.Id()
	request.BackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	response, err := s.Client.DeleteBackup(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := backupWaitForWorkRequest(workId, "backup",
		oci_psql.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *PsqlBackupResourceCrud) SetData() error {
	if s.Res.BackupSize != nil {
		s.D.Set("backup_size", *s.Res.BackupSize)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	copyStatus := []interface{}{}
	for _, item := range s.Res.CopyStatus {
		copyStatus = append(copyStatus, BackupCopyStatusDetailsToMap(item))
	}
	s.D.Set("copy_status", copyStatus)

	if s.Res.DbSystemDetails != nil {
		s.D.Set("db_system_details", []interface{}{DbSystemDetailsToMap(s.Res.DbSystemDetails)})
	} else {
		s.D.Set("db_system_details", nil)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
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
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LastAcceptedRequestToken != nil {
		s.D.Set("last_accepted_request_token", *s.Res.LastAcceptedRequestToken)
	}

	if s.Res.LastCompletedRequestToken != nil {
		s.D.Set("last_completed_request_token", *s.Res.LastCompletedRequestToken)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.RetentionPeriod != nil {
		s.D.Set("retention_period", *s.Res.RetentionPeriod)
	}

	if s.Res.SourceBackupDetails != nil {
		s.D.Set("source_backup_details", []interface{}{SourceBackupDetailsToMap(s.Res.SourceBackupDetails)})
	} else {
		s.D.Set("source_backup_details", nil)
	}

	s.D.Set("source_type", s.Res.SourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeCreatedPrecise != nil {
		s.D.Set("time_created_precise", s.Res.TimeCreatedPrecise.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func BackupCopyStatusDetailsToMap(obj oci_psql.BackupCopyStatusDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupId != nil {
		result["backup_id"] = string(*obj.BackupId)
	}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	result["state"] = string(obj.State)

	if obj.StateDetails != nil {
		result["state_details"] = string(*obj.StateDetails)
	}

	return result
}

func BackupSummaryToMap(obj oci_psql.BackupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupSize != nil {
		result["backup_size"] = int(*obj.BackupSize)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	copyStatus := []interface{}{}
	for _, item := range obj.CopyStatus {
		copyStatus = append(copyStatus, BackupCopyStatusDetailsToMap(item))
	}
	result["copy_status"] = copyStatus

	if obj.DbSystemId != nil {
		result["db_system_id"] = string(*obj.DbSystemId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.RetentionPeriod != nil {
		result["retention_period"] = int(*obj.RetentionPeriod)
	}

	result["source_type"] = string(obj.SourceType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeCreatedPrecise != nil {
		result["time_created_precise"] = obj.TimeCreatedPrecise.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func DbSystemDetailsToMap(obj *oci_psql.DbSystemDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConfigId != nil {
		result["config_id"] = string(*obj.ConfigId)
	}

	if obj.DbVersion != nil {
		result["db_version"] = string(*obj.DbVersion)
	}

	result["system_type"] = string(obj.SystemType)

	return result
}

func SourceBackupDetailsToMap(obj *oci_psql.SourceBackupDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.SourceBackupId != nil {
		result["source_backup_id"] = string(*obj.SourceBackupId)
	}

	if obj.SourceRegion != nil {
		result["source_region"] = string(*obj.SourceRegion)
	}

	return result
}

func (s *PsqlBackupResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_psql.ChangeBackupCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BackupId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "psql")

	_, err := s.Client.ChangeBackupCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	return nil
}
