// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataflow

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dataflow "github.com/oracle/oci-go-sdk/v65/dataflow"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataflowPoolResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataflowPool,
		Read:     readDataflowPool,
		Update:   updateDataflowPool,
		Delete:   deleteDataflowPool,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configurations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"max": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"shape": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"shape_config": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
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
			"idle_timeout_in_minutes": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"schedules": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"day_of_week": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"start_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"stop_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"state": {
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_dataflow.PoolLifecycleStateAccepted),
					string(oci_dataflow.PoolLifecycleStateScheduled),
					string(oci_dataflow.PoolLifecycleStateCreating),
					string(oci_dataflow.PoolLifecycleStateActive),
					string(oci_dataflow.PoolLifecycleStateStopping),
					string(oci_dataflow.PoolLifecycleStateStopped),
					string(oci_dataflow.PoolLifecycleStateUpdating),
					string(oci_dataflow.PoolLifecycleStateDeleting),
					string(oci_dataflow.PoolLifecycleStateDeleted),
					string(oci_dataflow.PoolLifecycleStateFailed),
				}, true),
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_principal_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_metrics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"active_runs_count": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"actively_used_node_count": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"logical_shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"pool_count": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"time_last_metrics_updated": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_stopped": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_used": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
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

func createDataflowPool(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_dataflow.PoolLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_dataflow.PoolLifecycleStateDeleted {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopPool(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.PoolLifecycleStateDeleted)
	}
	return nil

}

func readDataflowPool(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	return tfresource.ReadResource(sync)
}

func updateDataflowPool(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()

	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_dataflow.PoolLifecycleStateActive == oci_dataflow.PoolLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_dataflow.PoolLifecycleStateDeleted == oci_dataflow.PoolLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartPool(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.PoolLifecycleStateActive)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopPool(); err != nil {
			return err
		}
		sync.D.Set("state", oci_dataflow.PoolLifecycleStateDeleted)
	}

	return nil
}

func deleteDataflowPool(d *schema.ResourceData, m interface{}) error {
	sync := &DataflowPoolResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataFlowClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataflowPoolResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataflow.DataFlowClient
	Res                    *oci_dataflow.Pool
	DisableNotFoundRetries bool
}

func (s *DataflowPoolResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DataflowPoolResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_dataflow.PoolLifecycleStateCreating),
	}
}

func (s *DataflowPoolResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_dataflow.PoolLifecycleStateActive),
		string(oci_dataflow.PoolLifecycleStateScheduled),
		string(oci_dataflow.PoolLifecycleStateAccepted),
	}
}

func (s *DataflowPoolResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_dataflow.PoolLifecycleStateDeleting),
		string(oci_dataflow.PoolLifecycleStateStopping),
	}
}

func (s *DataflowPoolResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_dataflow.PoolLifecycleStateDeleted),
	}
}

func (s *DataflowPoolResourceCrud) Create() error {
	request := oci_dataflow.CreatePoolRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_dataflow.PoolConfig, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToPoolConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
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

	if idleTimeoutInMinutes, ok := s.D.GetOkExists("idle_timeout_in_minutes"); ok {
		tmp := idleTimeoutInMinutes.(int)
		request.IdleTimeoutInMinutes = &tmp
	}

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		interfaces := schedules.([]interface{})
		tmp := make([]oci_dataflow.PoolSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToPoolSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.CreatePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		fmt.Println("returning from create fn here is workId: ", workId)
		return s.getPoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow"), oci_dataflow.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
	} else {
		fmt.Println("returning from create fn here is workId is nill")
		s.Res = &response.Pool
		return nil
	}

}

func (s *DataflowPoolResourceCrud) getPoolFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_dataflow.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	fmt.Println("about to call poolWaitForWorkRequest")
	// Wait until it finishes
	poolId, err := poolWaitForWorkRequest(workId, "dataflow",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*poolId)

	return s.Get()
}

func poolWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "dataflow", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_dataflow.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func poolWaitForWorkRequest(wId *string, entityType string, action oci_dataflow.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_dataflow.DataFlowClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "dataflow")
	retryPolicy.ShouldRetryOperation = poolWorkRequestShouldRetryFunc(timeout)

	response := oci_dataflow.GetWorkRequestResponse{}
	fmt.Println("Here is the response: ", response)
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_dataflow.WorkRequestStatusInprogress),
			string(oci_dataflow.WorkRequestStatusAccepted),
			string(oci_dataflow.WorkRequestStatusCancelling),
		},
		Target: []string{
			string(oci_dataflow.WorkRequestStatusSucceeded),
			string(oci_dataflow.WorkRequestStatusFailed),
			string(oci_dataflow.WorkRequestStatusCancelled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_dataflow.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			fmt.Println("inner response: ", response)
			wr := &response.WorkRequest
			fmt.Println("inner response work request: ", wr)
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	fmt.Println("Got out of the waititng loop")
	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		fmt.Println("looping through response resources", res)
		if strings.Contains(strings.ToLower(*res.ResourceType), entityType) {
			if res.ActionType == action {
				identifier = res.ResourceId
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_dataflow.WorkRequestStatusFailed || response.Status == oci_dataflow.WorkRequestStatusCancelled {
		return nil, getErrorFromDataflowPoolWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDataflowPoolWorkRequest(client *oci_dataflow.DataFlowClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_dataflow.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_dataflow.ListWorkRequestErrorsRequest{
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

func (s *DataflowPoolResourceCrud) Get() error {
	request := oci_dataflow.GetPoolRequest{}

	tmp := s.D.Id()
	request.PoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.GetPool(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Pool
	return nil
}

func (s *DataflowPoolResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_dataflow.UpdatePoolRequest{}

	if configurations, ok := s.D.GetOkExists("configurations"); ok {
		interfaces := configurations.([]interface{})
		tmp := make([]oci_dataflow.PoolConfig, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configurations", stateDataIndex)
			converted, err := s.mapToPoolConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("configurations") {
			request.Configurations = tmp
		}
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

	if idleTimeoutInMinutes, ok := s.D.GetOkExists("idle_timeout_in_minutes"); ok {
		tmp := idleTimeoutInMinutes.(int)
		request.IdleTimeoutInMinutes = &tmp
	}

	tmp := s.D.Id()
	request.PoolId = &tmp

	if schedules, ok := s.D.GetOkExists("schedules"); ok {
		interfaces := schedules.([]interface{})
		tmp := make([]oci_dataflow.PoolSchedule, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedules", stateDataIndex)
			converted, err := s.mapToPoolSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("schedules") {
			request.Schedules = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	response, err := s.Client.UpdatePool(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId != nil {
		fmt.Println("returning from update fn here is workId: ", workId)
		return s.getPoolFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow"), oci_dataflow.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	} else {
		fmt.Println("returning from create fn here is workId is nill")
		err := s.Get()
		if err != nil {
			return err
		}
		return nil
	}
}

func (s *DataflowPoolResourceCrud) Delete() error {
	request := oci_dataflow.DeletePoolRequest{}

	tmp := s.D.Id()
	request.PoolId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.DeletePool(context.Background(), request)
	return err
}

func (s *DataflowPoolResourceCrud) SetData() error {
	fmt.Println("in setData, ", s.Res)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	configurations := []interface{}{}
	for _, item := range s.Res.Configurations {
		configurations = append(configurations, PoolConfigToMap(item))
	}
	s.D.Set("configurations", configurations)

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

	if s.Res.IdleTimeoutInMinutes != nil {
		s.D.Set("idle_timeout_in_minutes", *s.Res.IdleTimeoutInMinutes)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.OwnerPrincipalId != nil {
		s.D.Set("owner_principal_id", *s.Res.OwnerPrincipalId)
	}

	if s.Res.OwnerUserName != nil {
		s.D.Set("owner_user_name", *s.Res.OwnerUserName)
	}

	if s.Res.PoolMetrics != nil {
		s.D.Set("pool_metrics", []interface{}{PoolMetricsToMap(s.Res.PoolMetrics)})
	} else {
		s.D.Set("pool_metrics", nil)
	}

	schedules := []interface{}{}
	for _, item := range s.Res.Schedules {
		schedules = append(schedules, PoolScheduleToMap(item))
	}
	s.D.Set("schedules", schedules)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DataflowPoolResourceCrud) StartPool() error {
	request := oci_dataflow.StartPoolRequest{}

	idTmp := s.D.Id()
	request.PoolId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.StartPool(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_dataflow.PoolLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DataflowPoolResourceCrud) StopPool() error {
	fmt.Println("ML: in stop pool")
	request := oci_dataflow.StopPoolRequest{}

	idTmp := s.D.Id()
	request.PoolId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.StopPool(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_dataflow.PoolLifecycleStateStopped }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func NodeCountToMap(obj oci_dataflow.NodeCount) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogicalShape != nil {
		result["logical_shape"] = string(*obj.LogicalShape)
	}

	if obj.Count != nil {
		result["pool_count"] = int(*obj.Count)
	}

	return result
}

func (s *DataflowPoolResourceCrud) mapToPoolConfig(fieldKeyFormat string) (oci_dataflow.PoolConfig, error) {
	result := oci_dataflow.PoolConfig{}

	if max, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max")); ok {
		tmp := max.(int)
		result.Max = &tmp
	}

	if min, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min")); ok {
		tmp := min.(int)
		result.Min = &tmp
	}

	if shape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape")); ok {
		tmp := shape.(string)
		result.Shape = &tmp
	}

	if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
		if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
			tmp, err := s.mapToShapeConfig(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
			}
			result.ShapeConfig = &tmp
		}
	}

	return result, nil
}

func PoolConfigToMap(obj oci_dataflow.PoolConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Max != nil {
		result["max"] = int(*obj.Max)
	}

	if obj.Min != nil {
		result["min"] = int(*obj.Min)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.ShapeConfig != nil {
		result["shape_config"] = []interface{}{PoolShapeConfigToMap(obj.ShapeConfig)}
	}

	return result
}

func PoolMetricsToMap(obj *oci_dataflow.PoolMetrics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ActiveRunsCount != nil {
		result["active_runs_count"] = strconv.FormatInt(*obj.ActiveRunsCount, 10)
	}

	activelyUsedNodeCount := []interface{}{}
	for _, item := range obj.ActivelyUsedNodeCount {
		activelyUsedNodeCount = append(activelyUsedNodeCount, NodeCountToMap(item))
	}
	result["actively_used_node_count"] = activelyUsedNodeCount

	if obj.TimeLastMetricsUpdated != nil {
		result["time_last_metrics_updated"] = obj.TimeLastMetricsUpdated.String()
	}

	if obj.TimeLastStarted != nil {
		result["time_last_started"] = obj.TimeLastStarted.String()
	}

	if obj.TimeLastStopped != nil {
		result["time_last_stopped"] = obj.TimeLastStopped.String()
	}

	if obj.TimeLastUsed != nil {
		result["time_last_used"] = obj.TimeLastUsed.String()
	}

	return result
}

func (s *DataflowPoolResourceCrud) mapToPoolSchedule(fieldKeyFormat string) (oci_dataflow.PoolSchedule, error) {
	result := oci_dataflow.PoolSchedule{}

	if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
		result.DayOfWeek = oci_dataflow.DayOfWeekEnum(dayOfWeek.(string))
	}

	if startTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "start_time")); ok {
		tmp := startTime.(int)
		result.StartTime = &tmp
	}

	if stopTime, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stop_time")); ok {
		tmp := stopTime.(int)
		result.StopTime = &tmp
	}

	return result, nil
}

func PoolScheduleToMap(obj oci_dataflow.PoolSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	result["day_of_week"] = string(obj.DayOfWeek)

	if obj.StartTime != nil {
		result["start_time"] = int(*obj.StartTime)
	}

	if obj.StopTime != nil {
		result["stop_time"] = int(*obj.StopTime)
	}

	return result
}

func PoolSummaryToMap(obj oci_dataflow.PoolSummary) map[string]interface{} {
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

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.OwnerPrincipalId != nil {
		result["owner_principal_id"] = string(*obj.OwnerPrincipalId)
	}

	if obj.OwnerUserName != nil {
		result["owner_user_name"] = string(*obj.OwnerUserName)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *DataflowPoolResourceCrud) mapToShapeConfig(fieldKeyFormat string) (oci_dataflow.ShapeConfig, error) {
	result := oci_dataflow.ShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func PoolShapeConfigToMap(obj *oci_dataflow.ShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *DataflowPoolResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_dataflow.ChangePoolCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PoolId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataflow")

	_, err := s.Client.ChangePoolCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
