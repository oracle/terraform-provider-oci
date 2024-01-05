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

func CloudMigrationsMigrationAssetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createCloudMigrationsMigrationAsset,
		Read:     readCloudMigrationsMigrationAsset,
		Update:   updateCloudMigrationsMigrationAsset,
		Delete:   deleteCloudMigrationsMigrationAsset,
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"inventory_asset_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"migration_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"replication_compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snap_shot_bucket_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"migration_asset_depends_on": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replication_schedule_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"depended_on_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"notifications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parent_snapshot": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshots": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"source_asset_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenancy_id": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createCloudMigrationsMigrationAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.CreateResource(d, sync)
}

func readCloudMigrationsMigrationAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

func updateCloudMigrationsMigrationAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteCloudMigrationsMigrationAsset(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationAssetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type CloudMigrationsMigrationAssetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_cloud_migrations.MigrationClient
	Res                    *oci_cloud_migrations.MigrationAsset
	DisableNotFoundRetries bool
}

func (s *CloudMigrationsMigrationAssetResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *CloudMigrationsMigrationAssetResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationAssetLifecycleStateCreating),
	}
}

func (s *CloudMigrationsMigrationAssetResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationAssetLifecycleStateNeedsAttention),
		string(oci_cloud_migrations.MigrationAssetLifecycleStateActive),
	}
}

func (s *CloudMigrationsMigrationAssetResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_cloud_migrations.MigrationAssetLifecycleStateDeleting),
	}
}

func (s *CloudMigrationsMigrationAssetResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_cloud_migrations.MigrationAssetLifecycleStateDeleted),
	}
}

func (s *CloudMigrationsMigrationAssetResourceCrud) Create() error {
	request := oci_cloud_migrations.CreateMigrationAssetRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if dependsOn, ok := s.D.GetOkExists("migration_asset_depends_on"); ok {
		interfaces := dependsOn.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("migration_asset_depends_on") {
			request.DependsOn = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inventoryAssetId, ok := s.D.GetOkExists("inventory_asset_id"); ok {
		tmp := inventoryAssetId.(string)
		request.InventoryAssetId = &tmp
	}

	if migrationId, ok := s.D.GetOkExists("migration_id"); ok {
		tmp := migrationId.(string)
		request.MigrationId = &tmp
	}

	if replicationCompartmentId, ok := s.D.GetOkExists("replication_compartment_id"); ok {
		tmp := replicationCompartmentId.(string)
		request.ReplicationCompartmentId = &tmp
	}

	if replicationScheduleId, ok := s.D.GetOkExists("replication_schedule_id"); ok {
		tmp := replicationScheduleId.(string)
		request.ReplicationScheduleId = &tmp
	}

	if snapShotBucketName, ok := s.D.GetOkExists("snap_shot_bucket_name"); ok {
		tmp := snapShotBucketName.(string)
		request.SnapShotBucketName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.CreateMigrationAsset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMigrationAssetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations"), oci_cloud_migrations.ActionTypeUpdated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *CloudMigrationsMigrationAssetResourceCrud) getMigrationAssetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_cloud_migrations.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	migrationAssetId, err := migrationAssetWaitForWorkRequest(workId, "migrationAsset",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, migrationAssetId)
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
	s.D.SetId(*migrationAssetId)

	return s.Get()
}

func migrationAssetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func migrationAssetWaitForWorkRequest(wId *string, entityType string, action oci_cloud_migrations.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_cloud_migrations.MigrationClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "cloud_migrations")
	retryPolicy.ShouldRetryOperation = migrationAssetWorkRequestShouldRetryFunc(timeout)

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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_cloud_migrations.OperationStatusFailed || response.Status == oci_cloud_migrations.OperationStatusCanceled {
		return nil, getErrorFromCloudMigrationsMigrationAssetWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromCloudMigrationsMigrationAssetWorkRequest(client *oci_cloud_migrations.MigrationClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_cloud_migrations.ActionTypeEnum) error {
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

func (s *CloudMigrationsMigrationAssetResourceCrud) Get() error {
	request := oci_cloud_migrations.GetMigrationAssetRequest{}

	tmp := s.D.Id()
	request.MigrationAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.GetMigrationAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MigrationAsset
	return nil
}

func (s *CloudMigrationsMigrationAssetResourceCrud) Update() error {
	request := oci_cloud_migrations.UpdateMigrationAssetRequest{}

	if dependsOn, ok := s.D.GetOkExists("migration_asset_depends_on"); ok {
		interfaces := dependsOn.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("migration_asset_depends_on") {
			request.DependsOn = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.MigrationAssetId = &tmp

	if replicationScheduleId, ok := s.D.GetOkExists("replication_schedule_id"); ok {
		tmp := replicationScheduleId.(string)
		request.ReplicationScheduleId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.UpdateMigrationAsset(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MigrationAsset
	return nil
}

func (s *CloudMigrationsMigrationAssetResourceCrud) Delete() error {
	request := oci_cloud_migrations.DeleteMigrationAssetRequest{}

	tmp := s.D.Id()
	request.MigrationAssetId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "cloud_migrations")

	response, err := s.Client.DeleteMigrationAsset(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := migrationAssetWaitForWorkRequest(workId, "migrationAsset",
		oci_cloud_migrations.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *CloudMigrationsMigrationAssetResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("depended_on_by", s.Res.DependedOnBy)
	s.D.Set("depended_on_by", s.Res.DependedOnBy)

	s.D.Set("migration_asset_depends_on", s.Res.DependsOn)
	s.D.Set("migration_asset_depends_on", s.Res.DependsOn)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MigrationId != nil {
		s.D.Set("migration_id", *s.Res.MigrationId)
	}

	s.D.Set("notifications", s.Res.Notifications)
	s.D.Set("notifications", s.Res.Notifications)

	if s.Res.ParentSnapshot != nil {
		s.D.Set("parent_snapshot", *s.Res.ParentSnapshot)
	}

	if s.Res.ReplicationCompartmentId != nil {
		s.D.Set("replication_compartment_id", *s.Res.ReplicationCompartmentId)
	}

	if s.Res.ReplicationScheduleId != nil {
		s.D.Set("replication_schedule_id", *s.Res.ReplicationScheduleId)
	}

	if s.Res.SnapShotBucketName != nil {
		s.D.Set("snap_shot_bucket_name", *s.Res.SnapShotBucketName)
	}

	s.D.Set("snapshots", s.Res.Snapshots)
	s.D.Set("snapshots", s.Res.Snapshots)

	if s.Res.SourceAssetId != nil {
		s.D.Set("source_asset_id", *s.Res.SourceAssetId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TenancyId != nil {
		s.D.Set("tenancy_id", *s.Res.TenancyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func MigrationAssetSummaryToMap(obj oci_cloud_migrations.MigrationAssetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["depended_on_by"] = obj.DependedOnBy
	result["depended_on_by"] = obj.DependedOnBy

	result["migration_asset_depends_on"] = obj.DependsOn
	result["migration_asset_depends_on"] = obj.DependsOn

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MigrationId != nil {
		result["migration_id"] = string(*obj.MigrationId)
	}

	result["notifications"] = obj.Notifications
	result["notifications"] = obj.Notifications

	if obj.ParentSnapshot != nil {
		result["parent_snapshot"] = string(*obj.ParentSnapshot)
	}

	if obj.ReplicationScheduleId != nil {
		result["replication_schedule_id"] = string(*obj.ReplicationScheduleId)
	}

	if obj.SnapshotInfo != nil {
		result["snapshot_info"] = string(*obj.SnapshotInfo)
	}

	result["snapshots"] = obj.Snapshots
	result["snapshots"] = obj.Snapshots

	if obj.SourceAssetId != nil {
		result["source_asset_id"] = string(*obj.SourceAssetId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TenancyId != nil {
		result["tenancy_id"] = string(*obj.TenancyId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
