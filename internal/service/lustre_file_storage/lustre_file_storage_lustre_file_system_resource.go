// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package lustre_file_storage

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_lustre_file_storage "github.com/oracle/oci-go-sdk/v65/lustrefilestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func LustreFileStorageLustreFileSystemResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		CreateContext: createLustreFileStorageLustreFileSystemWithContext,
		ReadContext:   readLustreFileStorageLustreFileSystemWithContext,
		UpdateContext: updateLustreFileStorageLustreFileSystemWithContext,
		DeleteContext: deleteLustreFileStorageLustreFileSystemWithContext,
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("2h0m"),
			Update: tfresource.GetTimeoutDuration("2h0m"),
			Delete: tfresource.GetTimeoutDuration("2h0m"),
		},
		Schema: map[string]*schema.Schema{
			// Required
			"availability_domain": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
			},
			"capacity_in_gbs": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"file_system_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"performance_tier": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"root_squash_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"client_exceptions": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"identity_squash": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"squash_gid": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"squash_uid": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"cluster_placement_group_id": {
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
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_system_description": {
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
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maintenance_window": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"day_of_week": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"time_start": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"override_maintenance_trigger": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"date_time_details": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:     schema.TypeString,
							Required: true, // YYYY-MM-DD
						},
						"time": {
							Type:     schema.TypeString,
							Required: true, // HH:MM
						},
					},
				},
			},
			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lnet": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_window_metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"active_or_next_planned_maintenance": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"finished_maintenance": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"date": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_maintenance_in_progress": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"major_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_service_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_billing_cycle_end": {
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

func createLustreFileStorageLustreFileSystemWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageLustreFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	if err := tfresource.CreateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if _, ok := sync.D.GetOkExists("override_maintenance_trigger"); ok {
		if err := sync.OverrideMaintenance(); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
	}
	return tfresource.HandleDiagError(m, nil)
}

func readLustreFileStorageLustreFileSystemWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageLustreFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateLustreFileStorageLustreFileSystemWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageLustreFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()

	if _, ok := sync.D.GetOkExists("override_maintenance_trigger"); ok && sync.D.HasChange("override_maintenance_trigger") {
		oldRaw, newRaw := sync.D.GetChange("override_maintenance_trigger")
		oldValue := oldRaw.(int)
		newValue := newRaw.(int)

		if oldValue < newValue {
			if err := sync.OverrideMaintenance(); err != nil {
				return tfresource.HandleDiagError(m, err)
			}
		} else {
			_ = sync.D.Set("override_maintenance_trigger", oldRaw)
			return tfresource.HandleDiagError(m, fmt.Errorf("new value of trigger should be greater than the old value"))
		}
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	return tfresource.HandleDiagError(m, nil)
}

func deleteLustreFileStorageLustreFileSystemWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &LustreFileStorageLustreFileSystemResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LustreFileStorageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type LustreFileStorageLustreFileSystemResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_lustre_file_storage.LustreFileStorageClient
	Res                    *oci_lustre_file_storage.LustreFileSystem
	DisableNotFoundRetries bool
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_lustre_file_storage.LustreFileSystemLifecycleStateCreating),
	}
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_lustre_file_storage.LustreFileSystemLifecycleStateActive),
	}
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_lustre_file_storage.LustreFileSystemLifecycleStateDeleting),
	}
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_lustre_file_storage.LustreFileSystemLifecycleStateDeleted),
	}
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.CreateLustreFileSystemRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if capacityInGBs, ok := s.D.GetOkExists("capacity_in_gbs"); ok {
		tmp := capacityInGBs.(int)
		request.CapacityInGBs = &tmp
	}

	if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
		tmp := clusterPlacementGroupId.(string)
		request.ClusterPlacementGroupId = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if fileSystemDescription, ok := s.D.GetOkExists("file_system_description"); ok {
		tmp := fileSystemDescription.(string)
		request.FileSystemDescription = &tmp
	}

	if fileSystemName, ok := s.D.GetOkExists("file_system_name"); ok {
		tmp := fileSystemName.(string)
		request.FileSystemName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if performanceTier, ok := s.D.GetOkExists("performance_tier"); ok {
		request.PerformanceTier = oci_lustre_file_storage.CreateLustreFileSystemDetailsPerformanceTierEnum(performanceTier.(string))
	}

	if rootSquashConfiguration, ok := s.D.GetOkExists("root_squash_configuration"); ok {
		if tmpList := rootSquashConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "root_squash_configuration", 0)
			tmp, err := s.mapToRootSquashConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RootSquashConfiguration = &tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.CreateLustreFileSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getLustreFileSystemFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage"), oci_lustre_file_storage.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) getLustreFileSystemFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_lustre_file_storage.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	lustreFileSystemId, err := lustreFileSystemWaitForWorkRequest(ctx, workId, "lustrefilesystem",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, lustreFileSystemId)
		_, cancelErr := s.Client.CancelWorkRequest(ctx,
			oci_lustre_file_storage.CancelWorkRequestRequest{
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
	s.D.SetId(*lustreFileSystemId)

	return s.GetWithContext(ctx)
}

func lustreFileSystemWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "lustre_file_storage", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_lustre_file_storage.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func lustreFileSystemWaitForWorkRequest(ctx context.Context, wId *string, entityType string, action oci_lustre_file_storage.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_lustre_file_storage.LustreFileStorageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "lustre_file_storage")
	retryPolicy.ShouldRetryOperation = lustreFileSystemWorkRequestShouldRetryFunc(timeout)

	response := oci_lustre_file_storage.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_lustre_file_storage.OperationStatusInProgress),
			string(oci_lustre_file_storage.OperationStatusAccepted),
			string(oci_lustre_file_storage.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_lustre_file_storage.OperationStatusSucceeded),
			string(oci_lustre_file_storage.OperationStatusFailed),
			string(oci_lustre_file_storage.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_lustre_file_storage.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_lustre_file_storage.OperationStatusFailed || response.Status == oci_lustre_file_storage.OperationStatusCanceled {
		return nil, getErrorFromLustreFileStorageLustreFileSystemWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLustreFileStorageLustreFileSystemWorkRequest(ctx context.Context, client *oci_lustre_file_storage.LustreFileStorageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_lustre_file_storage.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_lustre_file_storage.ListWorkRequestErrorsRequest{
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

func (s *LustreFileStorageLustreFileSystemResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.GetLustreFileSystemRequest{}

	tmp := s.D.Id()
	request.LustreFileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.GetLustreFileSystem(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.LustreFileSystem
	return nil
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_lustre_file_storage.UpdateLustreFileSystemRequest{}

	if capacityInGBs, ok := s.D.GetOkExists("capacity_in_gbs"); ok {
		if s.D.HasChange("capacity_in_gbs") {
			if capacityInGBs == nil {
				request.CapacityInGBs = nil
			} else {
				tmp := capacityInGBs.(int)
				request.CapacityInGBs = &tmp
			}
		} else {
			request.CapacityInGBs = nil
		}
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

	if fileSystemDescription, ok := s.D.GetOkExists("file_system_description"); ok {
		tmp := fileSystemDescription.(string)
		request.FileSystemDescription = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	tmp := s.D.Id()
	request.LustreFileSystemId = &tmp

	if maintenanceWindow, ok := s.D.GetOkExists("maintenance_window"); ok {
		if tmpList := maintenanceWindow.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window", 0)
			tmp, err := s.mapToMaintenanceWindow(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindow = &tmp
		}
	}

	if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
			request.NsgIds = tmp
		}
	}

	if rootSquashConfiguration, ok := s.D.GetOkExists("root_squash_configuration"); ok {
		if tmpList := rootSquashConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "root_squash_configuration", 0)
			tmp, err := s.mapToRootSquashConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RootSquashConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.UpdateLustreFileSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLustreFileSystemFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage"), oci_lustre_file_storage.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_lustre_file_storage.DeleteLustreFileSystemRequest{}

	tmp := s.D.Id()
	request.LustreFileSystemId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.DeleteLustreFileSystem(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := lustreFileSystemWaitForWorkRequest(ctx, workId, "lustrefilesystem",
		oci_lustre_file_storage.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) SetData() error {
	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CapacityInGBs != nil {
		s.D.Set("capacity_in_gbs", *s.Res.CapacityInGBs)
	}

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FileSystemDescription != nil {
		s.D.Set("file_system_description", *s.Res.FileSystemDescription)
	}

	if s.Res.FileSystemName != nil {
		s.D.Set("file_system_name", *s.Res.FileSystemName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.Lnet != nil {
		s.D.Set("lnet", *s.Res.Lnet)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.MaintenanceWindowMetadata != nil {
		s.D.Set("maintenance_window_metadata", []interface{}{MaintenanceWindowMetadataDetailsToMap(s.Res.MaintenanceWindowMetadata)})
	} else {
		s.D.Set("maintenance_window_metadata", nil)
	}

	if s.Res.MajorVersion != nil {
		s.D.Set("major_version", *s.Res.MajorVersion)
	}

	if s.Res.ManagementServiceAddress != nil {
		s.D.Set("management_service_address", *s.Res.ManagementServiceAddress)
	}

	nsgIds := []interface{}{}
	for _, item := range s.Res.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

	s.D.Set("performance_tier", s.Res.PerformanceTier)

	if s.Res.RootSquashConfiguration != nil {
		s.D.Set("root_squash_configuration", []interface{}{RootSquashConfigurationToMap(s.Res.RootSquashConfiguration)})
	} else {
		s.D.Set("root_squash_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeBillingCycleEnd != nil {
		s.D.Set("time_billing_cycle_end", s.Res.TimeBillingCycleEnd.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) OverrideMaintenance() error {
	request := oci_lustre_file_storage.OverrideMaintenanceRequest{}

	if dateTimeDetails, ok := s.D.GetOkExists("date_time_details"); ok {
		if tmpList := dateTimeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "date_time_details", 0)
			tmp, err := s.mapToDateAndTime(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.DateTimeDetails = &tmp
		}
	}

	idTmp := s.D.Id()
	request.LustreFileSystemId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	_, err := s.Client.OverrideMaintenance(context.Background(), request)
	if err != nil {
		return err
	}

	val := s.D.Get("override_maintenance_trigger")
	s.D.Set("override_maintenance_trigger", val)

	return nil
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) mapToDateAndTime(fieldKeyFormat string) (oci_lustre_file_storage.DateAndTime, error) {
	result := oci_lustre_file_storage.DateAndTime{}

	if dateVal, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "date")); ok {
		tmp := dateVal.(string)
		result.Date = &tmp
	}

	if timeVal, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
		tmp := timeVal.(string)
		result.Time = &tmp
	}

	return result, nil
}

func DateAndTimeToMap(obj *oci_lustre_file_storage.DateAndTime) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Date != nil {
		result["date"] = string(*obj.Date)
	}

	if obj.Time != nil {
		result["time"] = string(*obj.Time)
	}

	return result
}

func LustreFileSystemSummaryToMap(obj oci_lustre_file_storage.LustreFileSystemSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.CapacityInGBs != nil {
		result["capacity_in_gbs"] = int(*obj.CapacityInGBs)
	}

	if obj.ClusterPlacementGroupId != nil {
		result["cluster_placement_group_id"] = string(*obj.ClusterPlacementGroupId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FileSystemDescription != nil {
		result["file_system_description"] = string(*obj.FileSystemDescription)
	}

	if obj.FileSystemName != nil {
		result["file_system_name"] = string(*obj.FileSystemName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.KmsKeyId != nil {
		result["kms_key_id"] = string(*obj.KmsKeyId)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.Lnet != nil {
		result["lnet"] = string(*obj.Lnet)
	}

	if obj.MajorVersion != nil {
		result["major_version"] = string(*obj.MajorVersion)
	}

	if obj.ManagementServiceAddress != nil {
		result["management_service_address"] = string(*obj.ManagementServiceAddress)
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	result["performance_tier"] = string(obj.PerformanceTier)

	if obj.RootSquashConfiguration != nil {
		result["root_squash_configuration"] = []interface{}{RootSquashConfigurationToMap(obj.RootSquashConfiguration)}
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeBillingCycleEnd != nil {
		result["time_billing_cycle_end"] = obj.TimeBillingCycleEnd.String()
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) mapToMaintenanceWindow(fieldKeyFormat string) (oci_lustre_file_storage.MaintenanceWindow, error) {
	result := oci_lustre_file_storage.MaintenanceWindow{}

	if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
		result.DayOfWeek = oci_lustre_file_storage.MaintenanceWindowDayOfWeekEnum(dayOfWeek.(string))
	}

	if timeStart, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_start")); ok {
		tmp := timeStart.(string)
		result.TimeStart = &tmp
	}

	return result, nil
}

func MaintenanceWindowToMap(obj *oci_lustre_file_storage.MaintenanceWindow) map[string]interface{} {
	result := map[string]interface{}{}

	result["day_of_week"] = string(obj.DayOfWeek)

	if obj.TimeStart != nil {
		result["time_start"] = string(*obj.TimeStart)
	}

	return result
}

func MaintenanceWindowMetadataDetailsToMap(obj *oci_lustre_file_storage.MaintenanceWindowMetadataDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveOrNextPlannedMaintenance != nil {
		result["active_or_next_planned_maintenance"] = []interface{}{DateAndTimeToMap(obj.ActiveOrNextPlannedMaintenance)}
	}

	if obj.FinishedMaintenance != nil {
		result["finished_maintenance"] = []interface{}{DateAndTimeToMap(obj.FinishedMaintenance)}
	}

	if obj.IsMaintenanceInProgress != nil {
		result["is_maintenance_in_progress"] = bool(*obj.IsMaintenanceInProgress)
	}

	return result
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) mapToRootSquashConfiguration(fieldKeyFormat string) (oci_lustre_file_storage.RootSquashConfiguration, error) {
	result := oci_lustre_file_storage.RootSquashConfiguration{}

	if clientExceptions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_exceptions")); ok {
		interfaces := clientExceptions.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "client_exceptions")) {
			result.ClientExceptions = tmp
		}
	}

	if identitySquash, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_squash")); ok {
		result.IdentitySquash = oci_lustre_file_storage.RootSquashConfigurationIdentitySquashEnum(identitySquash.(string))
	}

	if squashGid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "squash_gid")); ok {
		tmp := squashGid.(string)
		if tmp != "null" && tmp != "" {
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return result, fmt.Errorf("unable to convert squashGid string: %s to an int64 and encountered error: %v", tmp, err)
			}
			result.SquashGid = &tmpInt64
		}
	}

	if squashUid, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "squash_uid")); ok {
		tmp := squashUid.(string)
		if tmp != "null" && tmp != "" {
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return result, fmt.Errorf("unable to convert squashUid string: %s to an int64 and encountered error: %v", tmp, err)
			}
			result.SquashUid = &tmpInt64
		}
	}

	return result, nil
}

func RootSquashConfigurationToMap(obj *oci_lustre_file_storage.RootSquashConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["client_exceptions"] = obj.ClientExceptions

	result["identity_squash"] = string(obj.IdentitySquash)

	if obj.SquashGid != nil {
		result["squash_gid"] = strconv.FormatInt(*obj.SquashGid, 10)
	}

	if obj.SquashUid != nil {
		result["squash_uid"] = strconv.FormatInt(*obj.SquashUid, 10)
	}

	return result
}

func (s *LustreFileStorageLustreFileSystemResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_lustre_file_storage.ChangeLustreFileSystemCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.LustreFileSystemId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage")

	response, err := s.Client.ChangeLustreFileSystemCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getLustreFileSystemFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "lustre_file_storage"), oci_lustre_file_storage.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
