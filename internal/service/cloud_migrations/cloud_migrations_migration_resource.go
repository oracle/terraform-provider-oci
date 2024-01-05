// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsMigrationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudMigrationsMigration,
		Read:     readCloudMigrationsMigration,
		Update:   updateCloudMigrationsMigration,
		Delete:   deleteCloudMigrationsMigration,
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
			"is_completed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"replication_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCloudMigrationsMigration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudMigrationsMigration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

func updateCloudMigrationsMigration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudMigrationsMigration(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudMigrationsMigrationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_migrations.MigrationClient
	Res                    *oci_cloud_migrations.Migration
	DisableNotFoundRetries bool
}

func (s *CloudMigrationsMigrationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudMigrationsMigrationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationLifecycleStateCreating),
	}
}

func (s *CloudMigrationsMigrationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationLifecycleStateActive),
	}
}

func (s *CloudMigrationsMigrationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationLifecycleStateDeleting),
	}
}

func (s *CloudMigrationsMigrationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationLifecycleStateDeleted),
	}
}

func (s *CloudMigrationsMigrationResourceCrud) Create() error {
	request := oci_cloud_migrations.CreateMigrationRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isCompleted, ok := s.D.GetOkExists("is_completed"); ok {
		tmp := isCompleted.(bool)
		request.IsCompleted = &tmp
	}

	if replicationScheduleId, ok := s.D.GetOkExists("replication_schedule_id"); ok {
		tmp := replicationScheduleId.(string)
		request.ReplicationScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.CreateMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Migration
	return nil
}

func (s *CloudMigrationsMigrationResourceCrud) getMigrationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_migrations.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	migrationId, err := migrationWaitForWorkRequest(workId, "migration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, migrationId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_cloud_migrations.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*migrationId)

	return s.Get()
}

func migrationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "cloud_migrations", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_cloud_migrations.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func migrationWaitForWorkRequest(wId *string, entityType string, action oci_cloud_migrations.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_migrations.MigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_migrations")
	retryPolicy.ShouldRetryOperation = migrationWorkRequestShouldRetryFunc(timeout)

	response := oci_cloud_migrations.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_cloud_migrations.OperationStatusInProgress),
			string(oci_cloud_migrations.OperationStatusAccepted),
			string(oci_cloud_migrations.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_cloud_migrations.OperationStatusSucceeded),
			string(oci_cloud_migrations.OperationStatusFailed),
			string(oci_cloud_migrations.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_cloud_migrations.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_cloud_migrations.OperationStatusFailed || response.Status == oci_cloud_migrations.OperationStatusCanceled {
		return nil, getErrorFromCloudMigrationsMigrationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudMigrationsMigrationWorkRequest(client *oci_cloud_migrations.MigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_migrations.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_cloud_migrations.ListWorkRequestErrorsRequest{
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

func (s *CloudMigrationsMigrationResourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationRequest{}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.GetMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Migration
	return nil
}

func (s *CloudMigrationsMigrationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_cloud_migrations.UpdateMigrationRequest{}

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

	if isCompleted, ok := s.D.GetOkExists("is_completed"); ok {
		tmp := isCompleted.(bool)
		request.IsCompleted = &tmp
	}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	if replicationScheduleId, ok := s.D.GetOkExists("replication_schedule_id"); ok {
		tmp := replicationScheduleId.(string)
		request.ReplicationScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.UpdateMigration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Migration
	return nil
}

func (s *CloudMigrationsMigrationResourceCrud) Delete() error {
	request := oci_cloud_migrations.DeleteMigrationRequest{}

	tmp := s.D.Id()
	request.MigrationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.DeleteMigration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := migrationWaitForWorkRequest(workId, "migration",
		oci_cloud_migrations.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)

	return delWorkRequestErr
}

func (s *CloudMigrationsMigrationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsCompleted != nil {
		s.D.Set("is_completed", *s.Res.IsCompleted)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ReplicationScheduleId != nil {
		s.D.Set("replication_schedule_id", *s.Res.ReplicationScheduleId)
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

func MigrationSummaryToMap(obj oci_cloud_migrations.MigrationSummary) map[string]interface{} {
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
	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsCompleted != nil {
		result["is_completed"] = bool(*obj.IsCompleted)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ReplicationScheduleId != nil {
		result["replication_schedule_id"] = string(*obj.ReplicationScheduleId)
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

func (s *CloudMigrationsMigrationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_cloud_migrations.ChangeMigrationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MigrationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.ChangeMigrationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getMigrationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
