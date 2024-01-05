// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_stack_monitoring "github.com/oracle/oci-go-sdk/v65/stackmonitoring"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func StackMonitoringMonitoredResourceTaskResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createStackMonitoringMonitoredResourceTask,
		Read:     readStackMonitoringMonitoredResourceTask,
		Update:   updateStackMonitoringMonitoredResourceTask,
		Delete:   deleteStackMonitoringMonitoredResourceTask,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"source": {
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
								"IMPORT_OCI_TELEMETRY_RESOURCES",
							}, true),
						},

						// Optional
						"availability_proxy_metric_collection_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"availability_proxy_metrics": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"tenant_id": {
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
			"work_request_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createStackMonitoringMonitoredResourceTask(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.CreateResource(d, sync)
}

func readStackMonitoringMonitoredResourceTask(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.ReadResource(sync)
}

func updateStackMonitoringMonitoredResourceTask(d *schema.ResourceData, m interface{}) error {
	sync := &StackMonitoringMonitoredResourceTaskResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).StackMonitoringClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteStackMonitoringMonitoredResourceTask(d *schema.ResourceData, m interface{}) error {
	return nil
}

type StackMonitoringMonitoredResourceTaskResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_stack_monitoring.StackMonitoringClient
	Res                    *oci_stack_monitoring.MonitoredResourceTask
	DisableNotFoundRetries bool
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_stack_monitoring.MonitoredResourceTaskLifecycleStateInProgress),
	}
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_stack_monitoring.MonitoredResourceTaskLifecycleStateSucceeded),
		string(oci_stack_monitoring.MonitoredResourceTaskLifecycleStateNeedsAttention),
	}
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) Create() error {
	request := oci_stack_monitoring.CreateMonitoredResourceTaskRequest{}

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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if taskDetails, ok := s.D.GetOkExists("task_details"); ok {
		if tmpList := taskDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "task_details", 0)
			tmp, err := s.mapToMonitoredResourceTaskDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TaskDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.CreateMonitoredResourceTask(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getMonitoredResourceTaskFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring"), oci_stack_monitoring.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) getMonitoredResourceTaskFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_stack_monitoring.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	monitoredResourceTaskId, err := monitoredResourceTaskWaitForWorkRequest(workId, "stackmonitoringresourcetask",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*monitoredResourceTaskId)

	return s.Get()
}

func monitoredResourceTaskWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "stack_monitoring", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_stack_monitoring.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func monitoredResourceTaskWaitForWorkRequest(wId *string, entityType string, action oci_stack_monitoring.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_stack_monitoring.StackMonitoringClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "stack_monitoring")
	retryPolicy.ShouldRetryOperation = monitoredResourceTaskWorkRequestShouldRetryFunc(timeout)

	response := oci_stack_monitoring.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_stack_monitoring.OperationStatusInProgress),
			string(oci_stack_monitoring.OperationStatusAccepted),
			string(oci_stack_monitoring.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_stack_monitoring.OperationStatusSucceeded),
			string(oci_stack_monitoring.OperationStatusFailed),
			string(oci_stack_monitoring.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_stack_monitoring.GetWorkRequestRequest{
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
		log.Printf("res.EntityType: %s, entityType: %s, res.ActionType: %s, action: %s", strings.ToLower(*res.EntityType), entityType, res.ActionType, action)

		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_stack_monitoring.OperationStatusFailed || response.Status == oci_stack_monitoring.OperationStatusCanceled {
		return nil, getErrorFromStackMonitoringMonitoredResourceTaskWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromStackMonitoringMonitoredResourceTaskWorkRequest(client *oci_stack_monitoring.StackMonitoringClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_stack_monitoring.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_stack_monitoring.ListWorkRequestErrorsRequest{
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

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) Get() error {
	request := oci_stack_monitoring.GetMonitoredResourceTaskRequest{}

	tmp := s.D.Id()
	request.MonitoredResourceTaskId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.GetMonitoredResourceTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceTask
	return nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_stack_monitoring.UpdateMonitoredResourceTaskRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.MonitoredResourceTaskId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	response, err := s.Client.UpdateMonitoredResourceTask(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MonitoredResourceTask
	return nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := MonitoredResourceTaskDetailsToMap(&s.Res.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		s.D.Set("task_details", taskDetailsArray)
	} else {
		s.D.Set("task_details", nil)
	}

	if s.Res.TenantId != nil {
		s.D.Set("tenant_id", *s.Res.TenantId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("work_request_ids", s.Res.WorkRequestIds)

	return nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToMonitoredResourceTaskDetails(fieldKeyFormat string) (oci_stack_monitoring.MonitoredResourceTaskDetails, error) {
	var baseObject oci_stack_monitoring.MonitoredResourceTaskDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("IMPORT_OCI_TELEMETRY_RESOURCES"):
		details := oci_stack_monitoring.ImportOciTelemetryResourcesTaskDetails{}
		if availabilityProxyMetricCollectionInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_proxy_metric_collection_interval")); ok {
			tmp := availabilityProxyMetricCollectionInterval.(int)
			details.AvailabilityProxyMetricCollectionInterval = &tmp
		}
		if availabilityProxyMetrics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_proxy_metrics")); ok {
			interfaces := availabilityProxyMetrics.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "availability_proxy_metrics")) {
				details.AvailabilityProxyMetrics = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if resourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_group")); ok {
			tmp := resourceGroup.(string)
			details.ResourceGroup = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			details.Source = oci_stack_monitoring.ImportOciTelemetryResourcesTaskDetailsSourceEnum(source.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func MonitoredResourceTaskDetailsToMap(obj *oci_stack_monitoring.MonitoredResourceTaskDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_stack_monitoring.ImportOciTelemetryResourcesTaskDetails:
		result["type"] = "IMPORT_OCI_TELEMETRY_RESOURCES"

		if v.AvailabilityProxyMetricCollectionInterval != nil {
			result["availability_proxy_metric_collection_interval"] = int(*v.AvailabilityProxyMetricCollectionInterval)
		}

		result["availability_proxy_metrics"] = v.AvailabilityProxyMetrics

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ResourceGroup != nil {
			result["resource_group"] = string(*v.ResourceGroup)
		}

		result["source"] = string(v.Source)
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MonitoredResourceTaskSummaryToMap(obj oci_stack_monitoring.MonitoredResourceTaskSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := MonitoredResourceTaskDetailsToMap(&obj.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		result["task_details"] = taskDetailsArray
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["work_request_ids"] = obj.WorkRequestIds

	return result
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_stack_monitoring.ChangeMonitoredResourceTaskCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MonitoredResourceTaskId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "stack_monitoring")

	_, err := s.Client.ChangeMonitoredResourceTaskCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
