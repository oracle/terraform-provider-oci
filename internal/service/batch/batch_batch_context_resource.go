// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package batch

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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func BatchBatchContextResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(50 * time.Minute),
			Update: schema.DefaultTimeout(50 * time.Minute),
			Delete: schema.DefaultTimeout(50 * time.Minute),
		},
		CreateContext: createBatchBatchContextWithContext,
		ReadContext:   readBatchBatchContextWithContext,
		UpdateContext: updateBatchBatchContextWithContext,
		DeleteContext: deleteBatchBatchContextWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleets": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"max_concurrent_tasks": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"shape": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"memory_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"ocpus": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"shape_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"SERVICE_MANAGED_FLEET",
							}, true),
						},

						// Optional

						// Computed
						"details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"network": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"subnet_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
						"vnics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"source_ips": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
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
			"entitlements": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"job_priority_configurations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"tag_key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"tag_namespace": {
							Type:     schema.TypeString,
							Required: true,
						},
						"values": {
							Type:     schema.TypeMap,
							Required: true,
							Elem:     schema.TypeString,
						},
						"weight": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"logging_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"log_group_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OCI_LOGGING",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_batch.BatchContextLifecycleStateInactive),
					string(oci_batch.BatchContextLifecycleStateActive),
				}, true),
			},

			// Computed
			"lifecycle_details": {
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

func createBatchBatchContextWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_batch.BatchContextLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_batch.BatchContextLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResourceWithContext(ctx, d, sync); e != nil {
		return tfresource.HandleDiagError(m, e)
	}

	if powerOff {
		if err := sync.StopBatchContext(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchContextLifecycleStateInactive)
	}
	return nil

}

func readBatchBatchContextWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateBatchBatchContextWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_batch.BatchContextLifecycleStateActive == oci_batch.BatchContextLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_batch.BatchContextLifecycleStateInactive == oci_batch.BatchContextLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartBatchContext(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchContextLifecycleStateActive)
	}

	if err := tfresource.UpdateResourceWithContext(ctx, d, sync); err != nil {
		return tfresource.HandleDiagError(m, err)
	}

	if powerOff {
		if err := sync.StopBatchContext(ctx); err != nil {
			return tfresource.HandleDiagError(m, err)
		}
		sync.D.Set("state", oci_batch.BatchContextLifecycleStateInactive)
	}

	return nil
}

func deleteBatchBatchContextWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &BatchBatchContextResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BatchComputingClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type BatchBatchContextResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_batch.BatchComputingClient
	Res                    *oci_batch.BatchContext
	DisableNotFoundRetries bool
}

func (s *BatchBatchContextResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *BatchBatchContextResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_batch.BatchContextLifecycleStateCreating),
	}
}

func (s *BatchBatchContextResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_batch.BatchContextLifecycleStateActive),
		string(oci_batch.BatchContextLifecycleStateNeedsAttention),
	}
}

func (s *BatchBatchContextResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_batch.BatchContextLifecycleStateDeleting),
	}
}

func (s *BatchBatchContextResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_batch.BatchContextLifecycleStateDeleted),
	}
}

func (s *BatchBatchContextResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_batch.CreateBatchContextRequest{}

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

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		convertedEntitlements, err := expandStringIntMap(entitlements.(map[string]interface{}), "entitlements")
		if err != nil {
			return err
		}
		request.Entitlements = convertedEntitlements
	}

	if fleets, ok := s.D.GetOkExists("fleets"); ok {
		interfaces := fleets.([]interface{})
		tmp := make([]oci_batch.CreateFleetDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "fleets", stateDataIndex)
			converted, err := s.mapToCreateFleetDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("fleets") {
			request.Fleets = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobPriorityConfigurations, ok := s.D.GetOkExists("job_priority_configurations"); ok {
		interfaces := jobPriorityConfigurations.([]interface{})
		tmp := make([]oci_batch.JobPriorityConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_priority_configurations", stateDataIndex)
			converted, err := s.mapToJobPriorityConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("job_priority_configurations") {
			request.JobPriorityConfigurations = tmp
		}
	}

	if loggingConfiguration, ok := s.D.GetOkExists("logging_configuration"); ok {
		if tmpList := loggingConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "logging_configuration", 0)
			tmp, err := s.mapToLoggingConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LoggingConfiguration = tmp
		}
	}

	if network, ok := s.D.GetOkExists("network"); ok {
		if tmpList := network.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network", 0)
			tmp, err := s.mapToCreateNetworkDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Network = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.CreateBatchContext(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getBatchContextFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *BatchBatchContextResourceCrud) getBatchContextFromWorkRequest(ctx context.Context, workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_batch.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes, passing resource ID to allow direct state checking
	resourceId := s.D.Id()
	batchContextId, err := batchContextWaitForWorkRequest(ctx, workId, &resourceId, "batchcontext",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	// If work request times out but we have a resource ID, check the resource state directly
	if err != nil && s.D.Id() != "" {
		// Check if resource has reached Active state
		getRequest := oci_batch.GetBatchContextRequest{
			BatchContextId: &resourceId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		}
		getResponse, getErr := s.Client.GetBatchContext(ctx, getRequest)
		if getErr == nil && getResponse.BatchContext.Id != nil {
			state := getResponse.BatchContext.LifecycleState
			if state == oci_batch.BatchContextLifecycleStateActive {
				return s.GetWithContext(ctx)
			}
		}
		// If we can't verify the resource state, return the original error
		return err
	}

	if err != nil {
		return err
	}
	// Only set ID if it's not already set (e.g., if we returned early based on resource state)
	if batchContextId != nil && s.D.Id() == "" {
		s.D.SetId(*batchContextId)
	} else if batchContextId != nil && s.D.Id() != *batchContextId {
		// If ID is different, update it
		s.D.SetId(*batchContextId)
	}

	return s.GetWithContext(ctx)
}

func batchContextWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "batch", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_batch.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func checkBatchContextStateForWorkRequest(ctx context.Context, resourceId *string, action oci_batch.ActionTypeEnum, retryPolicy *oci_common.RetryPolicy, client *oci_batch.BatchComputingClient) (*oci_batch.WorkRequest, string, error) {
	if resourceId == nil || *resourceId == "" {
		return nil, "", nil
	}

	getRequest := oci_batch.GetBatchContextRequest{
		BatchContextId: resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	}
	getResponse, getErr := client.GetBatchContext(ctx, getRequest)
	if getErr == nil && getResponse.BatchContext.Id != nil {
		state := getResponse.BatchContext.LifecycleState
		// Fail fast if resource is in FAILED state
		if state == oci_batch.BatchContextLifecycleStateFailed {
			return nil, "", fmt.Errorf("resource entered FAILED state")
		}
		// For CREATE: If resource is ACTIVE, consider work request succeeded
		if action == oci_batch.ActionTypeCreated && state == oci_batch.BatchContextLifecycleStateActive {
			return &oci_batch.WorkRequest{
				Status: oci_batch.OperationStatusSucceeded,
			}, string(oci_batch.OperationStatusSucceeded), nil
		}
		// For CREATE: If still CREATING, continue waiting
		if action == oci_batch.ActionTypeCreated && state == oci_batch.BatchContextLifecycleStateCreating {
			return &oci_batch.WorkRequest{
				Status: oci_batch.OperationStatusInProgress,
			}, string(oci_batch.OperationStatusInProgress), nil
		}
		// For DELETE: If resource is DELETED, consider work request succeeded
		if action == oci_batch.ActionTypeDeleted && state == oci_batch.BatchContextLifecycleStateDeleted {
			return &oci_batch.WorkRequest{
				Status: oci_batch.OperationStatusSucceeded,
			}, string(oci_batch.OperationStatusSucceeded), nil
		}
		// For DELETE: If still DELETING, continue waiting
		if action == oci_batch.ActionTypeDeleted && state == oci_batch.BatchContextLifecycleStateDeleting {
			return &oci_batch.WorkRequest{
				Status: oci_batch.OperationStatusInProgress,
			}, string(oci_batch.OperationStatusInProgress), nil
		}
	} else if action == oci_batch.ActionTypeDeleted {
		// For DELETE: 404 means resource is deleted
		if failure, isServiceError := oci_common.IsServiceError(getErr); isServiceError && failure.GetHTTPStatusCode() == 404 {
			return &oci_batch.WorkRequest{
				Status: oci_batch.OperationStatusSucceeded,
			}, string(oci_batch.OperationStatusSucceeded), nil
		}
	}
	return nil, "", nil
}

func batchContextWaitForWorkRequest(ctx context.Context, wId *string, resourceId *string, entityType string, action oci_batch.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_batch.BatchComputingClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "batch")
	retryPolicy.ShouldRetryOperation = batchContextWorkRequestShouldRetryFunc(timeout)

	response := oci_batch.GetWorkRequestResponse{}
	var pollCount int
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_batch.OperationStatusInProgress),
			string(oci_batch.OperationStatusAccepted),
			string(oci_batch.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_batch.OperationStatusSucceeded),
			string(oci_batch.OperationStatusFailed),
			string(oci_batch.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			pollCount++
			// First try to get work request status
			var err error
			response, err = client.GetWorkRequest(ctx,
				oci_batch.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})

			// For DELETE operations, check resource state more frequently (every 5 polls = ~2.5 minutes)
			// This helps detect if resource is already deleted even if work request is stuck
			// For CREATE operations, check less frequently (every 20 polls = ~10 minutes)
			checkInterval := 20
			if action == oci_batch.ActionTypeDeleted {
				checkInterval = 5
			}
			if resourceId != nil && *resourceId != "" && pollCount%checkInterval == 0 {
				if wr, status, checkErr := checkBatchContextStateForWorkRequest(ctx, resourceId, action, retryPolicy, client); checkErr != nil {
					return nil, "", checkErr
				} else if status != "" {
					return wr, status, nil
				}
			}

			// If work request succeeds, use it
			if err == nil {
				wr := &response.WorkRequest
				return wr, string(wr.Status), nil
			}

			// If work request fails but we have resource ID, check resource state directly
			// This helps avoid gRPC timeouts by keeping the connection active
			if resourceId != nil && *resourceId != "" {
				if wr, status, checkErr := checkBatchContextStateForWorkRequest(ctx, resourceId, action, retryPolicy, client); checkErr != nil {
					return nil, "", checkErr
				} else if status != "" {
					return wr, status, nil
				}
			}

			// If we can't get work request or resource state, return the error
			return nil, "", err
		},
		Timeout:      timeout,
		MinTimeout:   10 * time.Second,
		Delay:        10 * time.Second,
		PollInterval: 30 * time.Second,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		// If we have a resource ID and it timed out, check if resource is in acceptable state
		if resourceId != nil && *resourceId != "" {
			getRequest := oci_batch.GetBatchContextRequest{
				BatchContextId: resourceId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			}
			getResponse, getErr := client.GetBatchContext(ctx, getRequest)
			if getErr == nil && getResponse.BatchContext.Id != nil {
				state := getResponse.BatchContext.LifecycleState
				// For CREATE: Check if resource is in acceptable state
				if action == oci_batch.ActionTypeCreated &&
					(state == oci_batch.BatchContextLifecycleStateActive ||
						state == oci_batch.BatchContextLifecycleStateNeedsAttention) {
					// Resource is in acceptable state, return the resource ID
					return resourceId, nil
				}
				// For DELETE: Check if resource is deleted
				if action == oci_batch.ActionTypeDeleted &&
					state == oci_batch.BatchContextLifecycleStateDeleted {
					// Resource is deleted, return the resource ID
					return resourceId, nil
				}
			} else if action == oci_batch.ActionTypeDeleted {
				// For DELETE: 404 means resource is deleted
				if failure, isServiceError := oci_common.IsServiceError(getErr); isServiceError && failure.GetHTTPStatusCode() == 404 {
					return resourceId, nil
				}
			}
		}
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	// If response.Resources is empty (e.g., we returned early based on resource state), use resourceId
	if len(response.Resources) > 0 {
		for _, res := range response.Resources {
			if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
				if res.ActionType == action {
					identifier = res.Identifier
					break
				}
			}
		}
	}

	// If identifier not found but we have resource ID, use it
	if identifier == nil && resourceId != nil && *resourceId != "" {
		identifier = resourceId
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	// Only check response.Status if response.WorkRequest is populated
	workRequestStatus := oci_batch.OperationStatusSucceeded
	// Check if we have a valid work request response (not a zero-value struct)
	if response.WorkRequest.Status != "" {
		workRequestStatus = response.WorkRequest.Status
	} else if len(response.Resources) == 0 && resourceId != nil && *resourceId != "" {
		// If we returned early based on resource state (response.Resources is empty), assume success
		// This happens when we detected ACTIVE state directly and returned early
		workRequestStatus = oci_batch.OperationStatusSucceeded
	}
	// If we have an identifier and status is not failed/cancelled, we're good
	if identifier != nil && workRequestStatus != oci_batch.OperationStatusFailed && workRequestStatus != oci_batch.OperationStatusCanceled {
		return identifier, nil
	}
	// Otherwise, check for errors
	if identifier == nil || workRequestStatus == oci_batch.OperationStatusFailed || workRequestStatus == oci_batch.OperationStatusCanceled {
		return nil, getErrorFromBatchBatchContextWorkRequest(ctx, client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromBatchBatchContextWorkRequest(ctx context.Context, client *oci_batch.BatchComputingClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_batch.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(ctx,
		oci_batch.ListWorkRequestErrorsRequest{
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

func (s *BatchBatchContextResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_batch.GetBatchContextRequest{}

	tmp := s.D.Id()
	request.BatchContextId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.GetBatchContext(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.BatchContext
	return nil
}

func (s *BatchBatchContextResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(ctx, compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_batch.UpdateBatchContextRequest{}

	tmp := s.D.Id()
	request.BatchContextId = &tmp

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

	if entitlements, ok := s.D.GetOkExists("entitlements"); ok {
		convertedEntitlements, err := expandStringIntMap(entitlements.(map[string]interface{}), "entitlements")
		if err != nil {
			return err
		}
		request.Entitlements = convertedEntitlements
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobPriorityConfigurations, ok := s.D.GetOkExists("job_priority_configurations"); ok {
		interfaces := jobPriorityConfigurations.([]interface{})
		tmp := make([]oci_batch.JobPriorityConfiguration, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_priority_configurations", stateDataIndex)
			converted, err := s.mapToJobPriorityConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("job_priority_configurations") {
			request.JobPriorityConfigurations = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.UpdateBatchContext(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchContextFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BatchBatchContextResourceCrud) DeleteWithContext(ctx context.Context) error {
	// Check if resource is already deleted before attempting delete
	resourceId := s.D.Id()
	getRequest := oci_batch.GetBatchContextRequest{
		BatchContextId: &resourceId,
		RequestMetadata: oci_common.RequestMetadata{
			RetryPolicy: tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"),
		},
	}
	getResponse, getErr := s.Client.GetBatchContext(ctx, getRequest)
	if getErr != nil {
		// If resource doesn't exist (404), consider delete successful
		if failure, isServiceError := oci_common.IsServiceError(getErr); isServiceError && failure.GetHTTPStatusCode() == 404 {
			log.Printf("[DEBUG] BatchContext already deleted, considering delete successful")
			return nil
		}
		// For other errors, continue with delete attempt
	} else if getResponse.BatchContext.Id != nil && getResponse.BatchContext.LifecycleState == oci_batch.BatchContextLifecycleStateDeleted {
		// Resource is already in DELETED state
		log.Printf("[DEBUG] BatchContext already in DELETED state, considering delete successful")
		return nil
	}

	request := oci_batch.DeleteBatchContextRequest{}
	request.BatchContextId = &resourceId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.DeleteBatchContext(ctx, request)
	if err != nil {
		// If resource is already deleted (404), consider delete successful
		if failure, isServiceError := oci_common.IsServiceError(err); isServiceError && failure.GetHTTPStatusCode() == 404 {
			log.Printf("[DEBUG] BatchContext already deleted, considering delete successful")
			return nil
		}
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := batchContextWaitForWorkRequest(ctx, workId, &resourceId, "batchcontext",
		oci_batch.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *BatchBatchContextResourceCrud) SetData() error {
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

	// Convert entitlements from map[string]int to map[string]string for Terraform schema
	entitlementsMap := make(map[string]string)
	if s.Res.Entitlements != nil {
		for k, v := range s.Res.Entitlements {
			entitlementsMap[k] = strconv.Itoa(v)
		}
	}
	s.D.Set("entitlements", entitlementsMap)

	fleets := []interface{}{}
	for _, item := range s.Res.Fleets {
		fleets = append(fleets, FleetToMap(item))
	}
	s.D.Set("fleets", fleets)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	jobPriorityConfigurations := []interface{}{}
	for _, item := range s.Res.JobPriorityConfigurations {
		jobPriorityConfigurations = append(jobPriorityConfigurations, JobPriorityConfigurationToMap(item))
	}
	s.D.Set("job_priority_configurations", jobPriorityConfigurations)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LoggingConfiguration != nil {
		loggingConfigurationArray := []interface{}{}
		if loggingConfigurationMap := LoggingConfigurationToMap(&s.Res.LoggingConfiguration); loggingConfigurationMap != nil {
			loggingConfigurationArray = append(loggingConfigurationArray, loggingConfigurationMap)
		}
		s.D.Set("logging_configuration", loggingConfigurationArray)
	} else {
		s.D.Set("logging_configuration", nil)
	}

	if s.Res.Network != nil {
		s.D.Set("network", []interface{}{NetworkToMap(s.Res.Network, false)})
	} else {
		s.D.Set("network", nil)
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

func (s *BatchBatchContextResourceCrud) StartBatchContext(ctx context.Context) error {
	request := oci_batch.StartBatchContextRequest{}

	idTmp := s.D.Id()
	request.BatchContextId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.StartBatchContext(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_batch.BatchContextLifecycleStateActive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *BatchBatchContextResourceCrud) StopBatchContext(ctx context.Context) error {
	request := oci_batch.StopBatchContextRequest{}

	idTmp := s.D.Id()
	request.BatchContextId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	_, err := s.Client.StopBatchContext(ctx, request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_batch.BatchContextLifecycleStateInactive }
	return tfresource.WaitForResourceConditionWithContext(ctx, s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func BatchContextSummaryToMap(obj oci_batch.BatchContextSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *BatchBatchContextResourceCrud) mapToCreateFleetDetails(fieldKeyFormat string) (oci_batch.CreateFleetDetails, error) {
	var baseObject oci_batch.CreateFleetDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("SERVICE_MANAGED_FLEET"):
		details := oci_batch.CreateServiceManagedFleetDetails{}
		if maxConcurrentTasks, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_concurrent_tasks")); ok {
			tmp := maxConcurrentTasks.(int)
			details.MaxConcurrentTasks = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
			if tmpList := shape.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape"), 0)
				tmp, err := s.mapToFleetShape(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert shape, encountered error: %v", err)
				}
				details.Shape = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func FleetToMap(obj oci_batch.Fleet) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_batch.CreateServiceManagedFleetDetails:
		// Request-side details (may be seen in some responses)
		result["type"] = "SERVICE_MANAGED_FLEET"

		if v.MaxConcurrentTasks != nil {
			result["max_concurrent_tasks"] = int(*v.MaxConcurrentTasks)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Shape != nil {
			result["shape"] = []interface{}{FleetShapeToMap(v.Shape)}
		}
	case oci_batch.ServiceManagedFleet:
		// Normal response type for BatchContext.Fleets
		result["type"] = "SERVICE_MANAGED_FLEET"

		if v.MaxConcurrentTasks != nil {
			result["max_concurrent_tasks"] = int(*v.MaxConcurrentTasks)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.Shape != nil {
			result["shape"] = []interface{}{FleetShapeToMap(v.Shape)}
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *BatchBatchContextResourceCrud) mapToCreateNetworkDetails(fieldKeyFormat string) (oci_batch.CreateNetworkDetails, error) {
	result := oci_batch.CreateNetworkDetails{}

	if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
		set := nsgIds.(*schema.Set)
		interfaces := set.List()
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
			result.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func NetworkToMap(obj *oci_batch.Network, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	vnics := []interface{}{}
	for _, item := range obj.Vnics {
		vnics = append(vnics, VnicToMap(item))
	}
	result["vnics"] = vnics

	return result
}

func (s *BatchBatchContextResourceCrud) mapToFleetShape(fieldKeyFormat string) (oci_batch.FleetShape, error) {
	result := oci_batch.FleetShape{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(int)
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(int)
		result.Ocpus = &tmp
	}

	if shapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_name")); ok {
		tmp := shapeName.(string)
		result.ShapeName = &tmp
	}

	return result, nil
}

func FleetShapeToMap(obj *oci_batch.FleetShape) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = int(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = int(*obj.Ocpus)
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	return result
}

func (s *BatchBatchContextResourceCrud) mapToJobPriorityConfiguration(fieldKeyFormat string) (oci_batch.JobPriorityConfiguration, error) {
	result := oci_batch.JobPriorityConfiguration{}

	if tagKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_key")); ok {
		tmp := tagKey.(string)
		result.TagKey = &tmp
	}

	if tagNamespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_namespace")); ok {
		tmp := tagNamespace.(string)
		result.TagNamespace = &tmp
	}

	if values, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "values")); ok {
		convertedValues, err := expandStringIntMap(values.(map[string]interface{}), "values")
		if err != nil {
			return result, err
		}
		result.Values = convertedValues
	}

	if weight, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weight")); ok {
		tmp := weight.(int)
		result.Weight = &tmp
	}

	return result, nil
}

func JobPriorityConfigurationToMap(obj oci_batch.JobPriorityConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TagKey != nil {
		result["tag_key"] = string(*obj.TagKey)
	}

	if obj.TagNamespace != nil {
		result["tag_namespace"] = string(*obj.TagNamespace)
	}

	// Convert values from map[string]int to map[string]string for Terraform schema
	valuesMap := make(map[string]string)
	if obj.Values != nil {
		for k, v := range obj.Values {
			valuesMap[k] = strconv.Itoa(v)
		}
	}
	result["values"] = valuesMap

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func expandStringIntMap(input map[string]interface{}, fieldName string) (map[string]int, error) {
	result := make(map[string]int, len(input))
	for key, value := range input {
		switch v := value.(type) {
		case int:
			result[key] = v
		case float64:
			result[key] = int(v)
		case string:
			intValue, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("%s[%s] must be an integer value, got %q", fieldName, key, v)
			}
			result[key] = intValue
		default:
			return nil, fmt.Errorf("%s[%s] must be an integer value, got %T", fieldName, key, value)
		}
	}
	return result, nil
}

func (s *BatchBatchContextResourceCrud) mapToLoggingConfiguration(fieldKeyFormat string) (oci_batch.LoggingConfiguration, error) {
	var baseObject oci_batch.LoggingConfiguration
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("OCI_LOGGING"):
		details := oci_batch.OciLoggingConfiguration{}
		if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
			tmp := logGroupId.(string)
			details.LogGroupId = &tmp
		}
		if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
			tmp := logId.(string)
			details.LogId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func LoggingConfigurationToMap(obj *oci_batch.LoggingConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_batch.OciLoggingConfiguration:
		result["type"] = "OCI_LOGGING"

		if v.LogGroupId != nil {
			result["log_group_id"] = string(*v.LogGroupId)
		}

		if v.LogId != nil {
			result["log_id"] = string(*v.LogId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func VnicToMap(obj oci_batch.Vnic) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["source_ips"] = obj.SourceIps

	return result
}

func (s *BatchBatchContextResourceCrud) updateCompartment(ctx context.Context, compartment interface{}) error {
	changeCompartmentRequest := oci_batch.ChangeBatchContextCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.BatchContextId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch")

	response, err := s.Client.ChangeBatchContextCompartment(ctx, changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBatchContextFromWorkRequest(ctx, workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "batch"), oci_batch.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
