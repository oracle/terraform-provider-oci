// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package stack_monitoring

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
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
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"IMPORT_OCI_TELEMETRY_RESOURCES",
								"UPDATE_AGENT_RECEIVER",
								"UPDATE_RESOURCE_TYPE_CONFIGS",
							}, true),
						},

						// Optional
						"agent_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
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
						"console_path_prefix": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"external_id_mapping": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"handler_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"is_enable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"lifecycle_status_mappings_for_up_status": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"receiver_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"listener_port": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"resource_group": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"resource_name_filter": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_name_mapping": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_type_filter": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_type_mapping": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"resource_types_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"availability_metrics_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"collection_interval_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"metrics": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												// Computed
											},
										},
									},
									"handler_config": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"collectd_resource_name_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"exclude_properties": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"include_properties": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"suffix": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"collector_types": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"handler_properties": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"value": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"metric_mappings": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"collector_metric_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_skip_upload": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"metric_upload_interval_in_seconds": {
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"telemetry_metric_name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"metric_name_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"exclude_pattern_on_prefix": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},
															"is_prefix_with_collector_type": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"metric_upload_interval_in_seconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"telegraf_resource_name_config": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													ForceNew: true,
													MaxItems: 1,
													MinItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"exclude_tags": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"include_tags": {
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																ForceNew: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"is_use_tags_only": {
																Type:     schema.TypeBool,
																Optional: true,
																Computed: true,
																ForceNew: true,
															},

															// Computed
														},
													},
												},
												"telemetry_resource_group": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"resource_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"service_base_url": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"should_use_metrics_flow_for_status": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"source": {
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
			"type": {
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
	stateConf := &retry.StateChangeConf{
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

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	s.D.Set("work_request_ids", s.Res.WorkRequestIds)

	return nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToAgentExtensionHandlerConfiguration(fieldKeyFormat string) (oci_stack_monitoring.AgentExtensionHandlerConfiguration, error) {
	result := oci_stack_monitoring.AgentExtensionHandlerConfiguration{}

	if collectdResourceNameConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collectd_resource_name_config")); ok {
		if tmpList := collectdResourceNameConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "collectd_resource_name_config"), 0)
			tmp, err := s.mapToCollectdResourceNameConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert collectd_resource_name_config, encountered error: %v", err)
			}
			result.CollectdResourceNameConfig = &tmp
		}
	}

	if collectorTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collector_types")); ok {
		interfaces := collectorTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "collector_types")) {
			result.CollectorTypes = tmp
		}
	}

	if handlerProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler_properties")); ok {
		interfaces := handlerProperties.([]interface{})
		tmp := make([]oci_stack_monitoring.AgentExtensionHandlerConfigurationProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "handler_properties"), stateDataIndex)
			converted, err := s.mapToAgentExtensionHandlerConfigurationProperty(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "handler_properties")) {
			result.HandlerProperties = tmp
		}
	}

	if metricMappings, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_mappings")); ok {
		interfaces := metricMappings.([]interface{})
		tmp := make([]oci_stack_monitoring.AgentExtensionHandlerMetricMappingDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metric_mappings"), stateDataIndex)
			converted, err := s.mapToAgentExtensionHandlerMetricMappingDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metric_mappings")) {
			result.MetricMappings = tmp
		}
	}

	if metricNameConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_name_config")); ok {
		if tmpList := metricNameConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metric_name_config"), 0)
			tmp, err := s.mapToMetricNameConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metric_name_config, encountered error: %v", err)
			}
			result.MetricNameConfig = &tmp
		}
	}

	if metricUploadIntervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_upload_interval_in_seconds")); ok {
		tmp := metricUploadIntervalInSeconds.(int)
		result.MetricUploadIntervalInSeconds = &tmp
	}

	if telegrafResourceNameConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "telegraf_resource_name_config")); ok {
		if tmpList := telegrafResourceNameConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "telegraf_resource_name_config"), 0)
			tmp, err := s.mapToTelegrafResourceNameConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert telegraf_resource_name_config, encountered error: %v", err)
			}
			result.TelegrafResourceNameConfig = &tmp
		}
	}

	if telemetryResourceGroup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "telemetry_resource_group")); ok {
		tmp := telemetryResourceGroup.(string)
		result.TelemetryResourceGroup = &tmp
	}

	return result, nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToAgentExtensionHandlerConfigurationProperty(fieldKeyFormat string) (oci_stack_monitoring.AgentExtensionHandlerConfigurationProperty, error) {
	result := oci_stack_monitoring.AgentExtensionHandlerConfigurationProperty{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToAgentExtensionHandlerMetricMappingDetails(fieldKeyFormat string) (oci_stack_monitoring.AgentExtensionHandlerMetricMappingDetails, error) {
	result := oci_stack_monitoring.AgentExtensionHandlerMetricMappingDetails{}

	if collectorMetricName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collector_metric_name")); ok {
		tmp := collectorMetricName.(string)
		result.CollectorMetricName = &tmp
	}

	if isSkipUpload, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_skip_upload")); ok {
		tmp := isSkipUpload.(bool)
		result.IsSkipUpload = &tmp
	}

	if metricUploadIntervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metric_upload_interval_in_seconds")); ok {
		tmp := metricUploadIntervalInSeconds.(int)
		result.MetricUploadIntervalInSeconds = &tmp
	}

	if telemetryMetricName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "telemetry_metric_name")); ok {
		tmp := telemetryMetricName.(string)
		result.TelemetryMetricName = &tmp
	}

	return result, nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToAgentReceiverProperties(fieldKeyFormat string) (oci_stack_monitoring.AgentReceiverProperties, error) {
	result := oci_stack_monitoring.AgentReceiverProperties{}

	if listenerPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_port")); ok {
		tmp := listenerPort.(int)
		result.ListenerPort = &tmp
	}

	return result, nil
}

func AgentReceiverPropertiesToMap(obj *oci_stack_monitoring.AgentReceiverProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ListenerPort != nil {
		result["listener_port"] = int(*obj.ListenerPort)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToAvailabilityMetricsDetails(fieldKeyFormat string) (oci_stack_monitoring.AvailabilityMetricsDetails, error) {
	result := oci_stack_monitoring.AvailabilityMetricsDetails{}

	if collectionIntervalInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "collection_interval_in_seconds")); ok {
		tmp := collectionIntervalInSeconds.(int)
		result.CollectionIntervalInSeconds = &tmp
	}

	if metrics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metrics")); ok {
		interfaces := metrics.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "metrics")) {
			result.Metrics = tmp
		}
	}

	return result, nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToCollectdResourceNameConfigurationDetails(fieldKeyFormat string) (oci_stack_monitoring.CollectdResourceNameConfigurationDetails, error) {
	result := oci_stack_monitoring.CollectdResourceNameConfigurationDetails{}

	if excludeProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_properties")); ok {
		interfaces := excludeProperties.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_properties")) {
			result.ExcludeProperties = tmp
		}
	}

	if includeProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_properties")); ok {
		interfaces := includeProperties.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "include_properties")) {
			result.IncludeProperties = tmp
		}
	}

	if suffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "suffix")); ok {
		tmp := suffix.(string)
		result.Suffix = &tmp
	}

	return result, nil
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToMetricNameConfigurationDetails(fieldKeyFormat string) (oci_stack_monitoring.MetricNameConfigurationDetails, error) {
	result := oci_stack_monitoring.MetricNameConfigurationDetails{}

	if excludePatternOnPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_pattern_on_prefix")); ok {
		tmp := excludePatternOnPrefix.(string)
		result.ExcludePatternOnPrefix = &tmp
	}

	if isPrefixWithCollectorType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_prefix_with_collector_type")); ok {
		tmp := isPrefixWithCollectorType.(bool)
		result.IsPrefixWithCollectorType = &tmp
	}

	return result, nil
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
		if consolePathPrefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "console_path_prefix")); ok {
			tmp := consolePathPrefix.(string)
			details.ConsolePathPrefix = &tmp
		}
		if externalIdMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "external_id_mapping")); ok {
			tmp := externalIdMapping.(string)
			details.ExternalIdMapping = &tmp
		}
		if lifecycleStatusMappingsForUpStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "lifecycle_status_mappings_for_up_status")); ok {
			interfaces := lifecycleStatusMappingsForUpStatus.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "lifecycle_status_mappings_for_up_status")) {
				details.LifecycleStatusMappingsForUpStatus = tmp
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
		if resourceNameFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name_filter")); ok {
			tmp := resourceNameFilter.(string)
			details.ResourceNameFilter = &tmp
		}
		if resourceNameMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_name_mapping")); ok {
			tmp := resourceNameMapping.(string)
			details.ResourceNameMapping = &tmp
		}
		if resourceTypeFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type_filter")); ok {
			tmp := resourceTypeFilter.(string)
			details.ResourceTypeFilter = &tmp
		}
		if resourceTypeMapping, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type_mapping")); ok {
			tmp := resourceTypeMapping.(string)
			details.ResourceTypeMapping = &tmp
		}
		if serviceBaseUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_base_url")); ok {
			tmp := serviceBaseUrl.(string)
			details.ServiceBaseUrl = &tmp
		}
		if shouldUseMetricsFlowForStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_use_metrics_flow_for_status")); ok {
			tmp := shouldUseMetricsFlowForStatus.(bool)
			details.ShouldUseMetricsFlowForStatus = &tmp
		}
		if source, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source")); ok {
			details.Source = oci_stack_monitoring.ImportOciTelemetryResourcesTaskDetailsSourceEnum(source.(string))
		}
		baseObject = details
	case strings.ToLower("UPDATE_AGENT_RECEIVER"):
		details := oci_stack_monitoring.UpdateAgentReceiverTaskDetails{}
		if agentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "agent_id")); ok {
			tmp := agentId.(string)
			details.AgentId = &tmp
		}
		if handlerType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler_type")); ok {
			details.HandlerType = oci_stack_monitoring.HandlerTypeEnum(handlerType.(string))
		}
		if isEnable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_enable")); ok {
			tmp := isEnable.(bool)
			details.IsEnable = &tmp
		}
		if receiverProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "receiver_properties")); ok {
			if tmpList := receiverProperties.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "receiver_properties"), 0)
				tmp, err := s.mapToAgentReceiverProperties(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert receiver_properties, encountered error: %v", err)
				}
				details.ReceiverProperties = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("UPDATE_RESOURCE_TYPE_CONFIGS"):
		details := oci_stack_monitoring.UpdateResourceTypeConfigTaskDetails{}
		if handlerType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler_type")); ok {
			details.HandlerType = oci_stack_monitoring.HandlerTypeEnum(handlerType.(string))
		}
		if resourceTypesConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_types_configuration")); ok {
			interfaces := resourceTypesConfiguration.([]interface{})
			tmp := make([]oci_stack_monitoring.ResourceTypeConfigDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "resource_types_configuration"), stateDataIndex)
				converted, err := s.mapToResourceTypeConfigDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "resource_types_configuration")) {
				details.ResourceTypesConfiguration = tmp
			}
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

		if v.ConsolePathPrefix != nil {
			result["console_path_prefix"] = string(*v.ConsolePathPrefix)
		}

		if v.ExternalIdMapping != nil {
			result["external_id_mapping"] = string(*v.ExternalIdMapping)
		}

		result["lifecycle_status_mappings_for_up_status"] = v.LifecycleStatusMappingsForUpStatus

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.ResourceGroup != nil {
			result["resource_group"] = string(*v.ResourceGroup)
		}

		if v.ResourceNameFilter != nil {
			result["resource_name_filter"] = string(*v.ResourceNameFilter)
		}

		if v.ResourceNameMapping != nil {
			result["resource_name_mapping"] = string(*v.ResourceNameMapping)
		}

		if v.ResourceTypeFilter != nil {
			result["resource_type_filter"] = string(*v.ResourceTypeFilter)
		}

		if v.ResourceTypeMapping != nil {
			result["resource_type_mapping"] = string(*v.ResourceTypeMapping)
		}

		if v.ServiceBaseUrl != nil {
			result["service_base_url"] = string(*v.ServiceBaseUrl)
		}

		if v.ShouldUseMetricsFlowForStatus != nil {
			result["should_use_metrics_flow_for_status"] = bool(*v.ShouldUseMetricsFlowForStatus)
		}

		result["source"] = string(v.Source)
	case oci_stack_monitoring.UpdateAgentReceiverTaskDetails:
		result["type"] = "UPDATE_AGENT_RECEIVER"

		if v.AgentId != nil {
			result["agent_id"] = string(*v.AgentId)
		}

		result["handler_type"] = string(v.HandlerType)

		if v.IsEnable != nil {
			result["is_enable"] = bool(*v.IsEnable)
		}

		if v.ReceiverProperties != nil {
			result["receiver_properties"] = []interface{}{AgentReceiverPropertiesToMap(v.ReceiverProperties)}
		}
	case oci_stack_monitoring.UpdateResourceTypeConfigTaskDetails:
		result["type"] = "UPDATE_RESOURCE_TYPE_CONFIGS"

		result["handler_type"] = string(v.HandlerType)

		resourceTypesConfiguration := []interface{}{}
		for _, item := range v.ResourceTypesConfiguration {
			resourceTypesConfiguration = append(resourceTypesConfiguration, ResourceTypeConfigDetailsToMap(item))
		}
		result["resource_types_configuration"] = resourceTypesConfiguration
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

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	result["work_request_ids"] = obj.WorkRequestIds

	return result
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToResourceTypeConfigDetails(fieldKeyFormat string) (oci_stack_monitoring.ResourceTypeConfigDetails, error) {
	result := oci_stack_monitoring.ResourceTypeConfigDetails{}

	if availabilityMetricsConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_metrics_config")); ok {
		if tmpList := availabilityMetricsConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "availability_metrics_config"), 0)
			tmp, err := s.mapToAvailabilityMetricsDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert availability_metrics_config, encountered error: %v", err)
			}
			result.AvailabilityMetricsConfig = &tmp
		}
	}

	if handlerConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "handler_config")); ok {
		if tmpList := handlerConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "handler_config"), 0)
			tmp, err := s.mapToAgentExtensionHandlerConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert handler_config, encountered error: %v", err)
			}
			result.HandlerConfig = &tmp
		}
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	return result, nil
}

func ResourceTypeConfigDetailsToMap(obj oci_stack_monitoring.ResourceTypeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityMetricsConfig != nil {
		result["availability_metrics_config"] = []interface{}{AvailabilityMetricsDetailsToMap(obj.AvailabilityMetricsConfig)}
	}

	if obj.HandlerConfig != nil {
		result["handler_config"] = []interface{}{AgentExtensionHandlerConfigurationToMap(obj.HandlerConfig)}
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	return result
}

func (s *StackMonitoringMonitoredResourceTaskResourceCrud) mapToTelegrafResourceNameConfigurationDetails(fieldKeyFormat string) (oci_stack_monitoring.TelegrafResourceNameConfigurationDetails, error) {
	result := oci_stack_monitoring.TelegrafResourceNameConfigurationDetails{}

	if excludeTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_tags")); ok {
		interfaces := excludeTags.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_tags")) {
			result.ExcludeTags = tmp
		}
	}

	if includeTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_tags")); ok {
		interfaces := includeTags.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "include_tags")) {
			result.IncludeTags = tmp
		}
	}

	if isUseTagsOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_use_tags_only")); ok {
		tmp := isUseTagsOnly.(bool)
		result.IsUseTagsOnly = &tmp
	}

	return result, nil
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
