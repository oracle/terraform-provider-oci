// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package redis

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_redis "github.com/oracle/oci-go-sdk/v65/redis"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func RedisOciCacheBackupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createRedisOciCacheBackupWithContext,
		ReadContext:   readRedisOciCacheBackupWithContext,
		UpdateContext: updateRedisOciCacheBackupWithContext,
		DeleteContext: deleteRedisOciCacheBackupWithContext,
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
			"source_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"backup_source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"retention_period_in_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"export_to_object_storage_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},

			// Computed
			"backup_size_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_memory_in_gbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"cluster_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shard_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"software_version": {
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

func createRedisOciCacheBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if _, ok := sync.D.GetOkExists("export_to_object_storage_trigger"); ok {
		err := sync.ExportOciCacheBackupToObjectStorage(ctx)
		if err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return nil

}

func readRedisOciCacheBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateRedisOciCacheBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()

	if _, ok := sync.D.GetOkExists("export_to_object_storage_trigger"); ok && sync.D.HasChange("export_to_object_storage_trigger") {
		oldRaw, newRaw := sync.D.GetChange("export_to_object_storage_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)
		if oldValue < newValue {
			err := sync.ExportOciCacheBackupToObjectStorage(ctx)

			if err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			sync.D.Set("export_to_object_storage_trigger", oldRaw)
			err := fmt.Errorf("new value of trigger should be greater than the old value")
			return tfresource.HandleDiagError(m, err)
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return nil
}

func deleteRedisOciCacheBackupWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &RedisOciCacheBackupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OciCacheBackupClient()
	sync.RedisClusterClient = m.(*client.OracleClients).RedisClusterClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type RedisOciCacheBackupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_redis.OciCacheBackupClient
	Res                    *oci_redis.OciCacheBackup
	RedisClusterClient     *oci_redis.RedisClusterClient
	DisableNotFoundRetries bool
}

func (s *RedisOciCacheBackupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *RedisOciCacheBackupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_redis.OciCacheBackupLifecycleStateCreating),
	}
}

func (s *RedisOciCacheBackupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_redis.OciCacheBackupLifecycleStateActive),
	}
}

func (s *RedisOciCacheBackupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_redis.OciCacheBackupLifecycleStateDeleting),
	}
}

func (s *RedisOciCacheBackupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_redis.OciCacheBackupLifecycleStateDeleted),
	}
}

func (s *RedisOciCacheBackupResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_redis.CreateOciCacheBackupRequest{}

	if backupSource, ok := s.D.GetOkExists("backup_source"); ok {
		request.BackupSource = oci_redis.OciCacheBackupBackupSourceEnum(backupSource.(string))
	}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	if sourceClusterId, ok := s.D.GetOkExists("source_cluster_id"); ok {
		tmp := sourceClusterId.(string)
		request.SourceClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.CreateOciCacheBackup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getOciCacheBackupFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *RedisOciCacheBackupResourceCrud) getOciCacheBackupFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_redis.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	ociCacheBackupId, err := ociCacheBackupWaitForWorkRequest(ctx, workId, "backup",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client, s.RedisClusterClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, ociCacheBackupId)
		_, cancelErr := s.RedisClusterClient.CancelWorkRequest(ctx,
			oci_redis.CancelWorkRequestRequest{
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
	s.D.SetId(*ociCacheBackupId)

	return s.GetWithContext(ctx)
}

func ociCacheBackupWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "redis", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_redis.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func ociCacheBackupWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_redis.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_redis.OciCacheBackupClient, redisClusterClient *oci_redis.RedisClusterClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "redis")
	retryPolicy.ShouldRetryOperation = ociCacheBackupWorkRequestShouldRetryFunc(timeout)

	response := oci_redis.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_redis.OperationStatusInProgress),
			string(oci_redis.OperationStatusAccepted),
			string(oci_redis.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_redis.OperationStatusSucceeded),
			string(oci_redis.OperationStatusFailed),
			string(oci_redis.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = redisClusterClient.GetWorkRequest(ctx,
				oci_redis.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_redis.OperationStatusFailed || response.Status == oci_redis.OperationStatusCanceled {
		return nil, getErrorFromRedisOciCacheBackupWorkRequest(ctx, client, redisClusterClient, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromRedisOciCacheBackupWorkRequest(ctx context.Context, client *oci_redis.OciCacheBackupClient, redisClusterClient *oci_redis.RedisClusterClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_redis.ActionTypeEnum) error {
	response, err := redisClusterClient.ListWorkRequestErrors(ctx,
		oci_redis.ListWorkRequestErrorsRequest{
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

func (s *RedisOciCacheBackupResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_redis.GetOciCacheBackupRequest{}

	tmp := s.D.Id()
	request.OciCacheBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.GetOciCacheBackup(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.OciCacheBackup
	return nil
}

func (s *RedisOciCacheBackupResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_redis.UpdateOciCacheBackupRequest{}

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

	tmp := s.D.Id()
	request.OciCacheBackupId = &tmp

	if retentionPeriodInDays, ok := s.D.GetOkExists("retention_period_in_days"); ok {
		tmp := retentionPeriodInDays.(int)
		request.RetentionPeriodInDays = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.UpdateOciCacheBackup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId == nil || *workId == "" {
		// Synchronous update: just refresh
		return s.GetWithContext(ctx)
	}
	return s.getOciCacheBackupFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *RedisOciCacheBackupResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_redis.DeleteOciCacheBackupRequest{}

	tmp := s.D.Id()
	request.OciCacheBackupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.DeleteOciCacheBackup(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := ociCacheBackupWaitForWorkRequest(ctx, workId, "backup",
		oci_redis.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client, s.RedisClusterClient)
	return delWorkRequestErr
}

func (s *RedisOciCacheBackupResourceCrud) SetData() error {
	if s.Res == nil {
		// Nothing to set; avoids nil dereference in edge cases.
		return nil
	}

	if s.Res.BackupSizeInGBs != nil {
		s.D.Set("backup_size_in_gbs", *s.Res.BackupSizeInGBs)
	}

	s.D.Set("backup_source", s.Res.BackupSource)

	s.D.Set("backup_type", s.Res.BackupType)

	if s.Res.ClusterMemoryInGBs != nil {
		s.D.Set("cluster_memory_in_gbs", *s.Res.ClusterMemoryInGBs)
	}

	s.D.Set("cluster_mode", s.Res.ClusterMode)

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

	if s.Res.RetentionPeriodInDays != nil {
		s.D.Set("retention_period_in_days", *s.Res.RetentionPeriodInDays)
	}

	if s.Res.ShardCount != nil {
		s.D.Set("shard_count", *s.Res.ShardCount)
	}

	s.D.Set("software_version", s.Res.SoftwareVersion)

	if s.Res.SourceClusterId != nil {
		s.D.Set("source_cluster_id", *s.Res.SourceClusterId)
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

func (s *RedisOciCacheBackupResourceCrud) ExportOciCacheBackupToObjectStorage(ctx context.Context) error {
	request := oci_redis.ExportOciCacheBackupToObjectStorageRequest{}

	if bucket, ok := s.D.GetOkExists("bucket"); ok {
		tmp := bucket.(string)
		request.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists("namespace"); ok {
		tmp := namespace.(string)
		request.NamespaceName = &tmp
	}

	idTmp := s.D.Id()
	request.OciCacheBackupId = &idTmp

	if prefix, ok := s.D.GetOkExists("prefix"); ok {
		tmp := prefix.(string)
		request.Prefix = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	_, exportErr := s.Client.ExportOciCacheBackupToObjectStorage(ctx, request)
	if exportErr != nil {
		return exportErr
	}
	if waitErr := tfresource.WaitForUpdatedStateWithContext(ctx, s.D, s); waitErr != nil {
		return waitErr
	}

	val := s.D.Get("export_to_object_storage_trigger")
	s.D.Set("export_to_object_storage_trigger", val)

	return nil
}

func OciCacheBackupSummaryToMap(obj oci_redis.OciCacheBackupSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackupSizeInGBs != nil {
		result["backup_size_in_gbs"] = float32(*obj.BackupSizeInGBs)
	}

	result["backup_source"] = string(obj.BackupSource)

	result["backup_type"] = string(obj.BackupType)

	result["cluster_mode"] = string(obj.ClusterMode)

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

	if obj.RetentionPeriodInDays != nil {
		result["retention_period_in_days"] = int(*obj.RetentionPeriodInDays)
	}

	if obj.SourceClusterId != nil {
		result["source_cluster_id"] = string(*obj.SourceClusterId)
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

func (s *RedisOciCacheBackupResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_redis.ChangeOciCacheBackupCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.OciCacheBackupId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis")

	response, err := s.Client.ChangeOciCacheBackupCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId == nil || *workId == "" {
		return fmt.Errorf("ChangeOciCacheBackupCompartment returned nil opc-work-request-id")
	}
	return s.getOciCacheBackupFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "redis"), oci_redis.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
