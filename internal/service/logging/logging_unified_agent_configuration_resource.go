// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package logging

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_logging "github.com/oracle/oci-go-sdk/v56/logging"
)

func LoggingUnifiedAgentConfigurationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoggingUnifiedAgentConfiguration,
		Read:     readLoggingUnifiedAgentConfiguration,
		Update:   updateLoggingUnifiedAgentConfiguration,
		Delete:   deleteLoggingUnifiedAgentConfiguration,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"service_configuration": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"configuration_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"LOGGING",
							}, true),
						},

						// Optional
						"destination": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"log_object_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"sources": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"source_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"LOG_TAIL",
											"WINDOWS_EVENT_LOG",
										}, true),
									},

									// Optional
									"channels": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"parser": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"parser_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"APACHE2",
														"APACHE_ERROR",
														"AUDITD",
														"CSV",
														"GROK",
														"JSON",
														"MSGPACK",
														"MULTILINE",
														"MULTILINE_GROK",
														"NONE",
														"REGEXP",
														"SYSLOG",
														"TSV",
													}, true),
												},

												// Optional
												"delimiter": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"expression": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"field_time_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"format": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"format_firstline": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_failure_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"grok_name_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"is_estimate_current_event": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_keep_time_key": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_null_empty_string": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_support_colonless_ident": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"is_with_priority": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"keys": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"message_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"multi_line_start_regexp": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"null_value_pattern": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"patterns": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"field_time_format": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_key": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"field_time_zone": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"pattern": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"rfc5424time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"syslog_parser_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_format": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"time_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"timeout_in_milliseconds": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"types": {
													Type:     schema.TypeMap,
													Optional: true,
													Computed: true,
													Elem:     schema.TypeString,
												},

												// Computed
											},
										},
									},
									"paths": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"group_association": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"group_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
			},

			// Computed
			"configuration_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.ReadResource(sync)
}

func updateLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoggingUnifiedAgentConfiguration(d *schema.ResourceData, m interface{}) error {
	sync := &LoggingUnifiedAgentConfigurationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoggingManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoggingUnifiedAgentConfigurationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_logging.LoggingManagementClient
	Res                    *oci_logging.UnifiedAgentConfiguration
	DisableNotFoundRetries bool
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateCreating),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_logging.LogLifecycleStateActive),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_logging.LogLifecycleStateDeleting),
	}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Create() error {
	request := oci_logging.CreateUnifiedAgentConfigurationRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if groupAssociation, ok := s.D.GetOkExists("group_association"); ok {
		if tmpList := groupAssociation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "group_association", 0)
			tmp, err := s.mapToGroupAssociationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GroupAssociation = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if serviceConfiguration, ok := s.D.GetOkExists("service_configuration"); ok {
		if tmpList := serviceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_configuration", 0)
			tmp, err := s.mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceConfiguration = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.CreateUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) getUnifiedAgentConfigurationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_logging.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	unifiedAgentConfigurationId, err := unifiedAgentConfigurationWaitForWorkRequest(workId, "unifiedagentconfiguration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*unifiedAgentConfigurationId)

	return s.Get()
}

func unifiedAgentConfigurationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "logging", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_logging.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func unifiedAgentConfigurationWaitForWorkRequest(wId *string, entityType string, action oci_logging.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_logging.LoggingManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "logging")
	retryPolicy.ShouldRetryOperation = unifiedAgentConfigurationWorkRequestShouldRetryFunc(timeout)

	response := oci_logging.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_logging.OperationStatusInProgress),
			string(oci_logging.OperationStatusAccepted),
			string(oci_logging.OperationStatusCancelling),
		},
		Target: []string{
			string(oci_logging.OperationStatusSucceeded),
			string(oci_logging.OperationStatusFailed),
			string(oci_logging.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_logging.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_logging.OperationStatusFailed || response.Status == oci_logging.OperationStatusCanceled {
		return nil, getErrorFromLoggingUnifiedAgentConfigurationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromLoggingUnifiedAgentConfigurationWorkRequest(client *oci_logging.LoggingManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_logging.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_logging.ListWorkRequestErrorsRequest{
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

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Get() error {
	request := oci_logging.GetUnifiedAgentConfigurationRequest{}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.GetUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UnifiedAgentConfiguration
	return nil
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_logging.UpdateUnifiedAgentConfigurationRequest{}

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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if groupAssociation, ok := s.D.GetOkExists("group_association"); ok {
		if tmpList := groupAssociation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "group_association", 0)
			tmp, err := s.mapToGroupAssociationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.GroupAssociation = &tmp
		}
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if serviceConfiguration, ok := s.D.GetOkExists("service_configuration"); ok {
		if tmpList := serviceConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_configuration", 0)
			tmp, err := s.mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ServiceConfiguration = tmp
		}
	}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.UpdateUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) Delete() error {
	request := oci_logging.DeleteUnifiedAgentConfigurationRequest{}

	tmp := s.D.Id()
	request.UnifiedAgentConfigurationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.DeleteUnifiedAgentConfiguration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId

	// Wait until it finishes
	_, delWorkRequestErr := unifiedAgentConfigurationWaitForWorkRequest(workId, "unifiedagentconfiguration",
		oci_logging.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("configuration_state", s.Res.ConfigurationState)

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

	if s.Res.GroupAssociation != nil {
		s.D.Set("group_association", []interface{}{GroupAssociationDetailsToMap(s.Res.GroupAssociation)})
	} else {
		s.D.Set("group_association", nil)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.ServiceConfiguration != nil {
		serviceConfigurationArray := []interface{}{}
		if serviceConfigurationMap := UnifiedAgentServiceConfigurationDetailsToMap(&s.Res.ServiceConfiguration); serviceConfigurationMap != nil {
			serviceConfigurationArray = append(serviceConfigurationArray, serviceConfigurationMap)
		}
		s.D.Set("service_configuration", serviceConfigurationArray)
	} else {
		s.D.Set("service_configuration", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastModified != nil {
		s.D.Set("time_last_modified", s.Res.TimeLastModified.String())
	}

	return nil
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGrokPattern(fieldKeyFormat string) (oci_logging.GrokPattern, error) {
	result := oci_logging.GrokPattern{}

	if fieldTimeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_format")); ok {
		tmp := fieldTimeFormat.(string)
		result.FieldTimeFormat = &tmp
	}

	if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
		tmp := fieldTimeKey.(string)
		result.FieldTimeKey = &tmp
	}

	if fieldTimeZone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_zone")); ok {
		tmp := fieldTimeZone.(string)
		result.FieldTimeZone = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if pattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pattern")); ok {
		tmp := pattern.(string)
		result.Pattern = &tmp
	}

	return result, nil
}

func GrokPatternToMap(obj oci_logging.GrokPattern) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FieldTimeFormat != nil {
		result["field_time_format"] = string(*obj.FieldTimeFormat)
	}

	if obj.FieldTimeKey != nil {
		result["field_time_key"] = string(*obj.FieldTimeKey)
	}

	if obj.FieldTimeZone != nil {
		result["field_time_zone"] = string(*obj.FieldTimeZone)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Pattern != nil {
		result["pattern"] = string(*obj.Pattern)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToGroupAssociationDetails(fieldKeyFormat string) (oci_logging.GroupAssociationDetails, error) {
	result := oci_logging.GroupAssociationDetails{}

	if groupList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_list")); ok {
		interfaces := groupList.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_list")) {
			result.GroupList = tmp
		}
	}

	return result, nil
}

func GroupAssociationDetailsToMap(obj *oci_logging.GroupAssociationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["group_list"] = obj.GroupList

	return result
}

func UnifiedAgentConfigurationSummaryToMap(obj oci_logging.UnifiedAgentConfigurationSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	result["configuration_state"] = string(obj.ConfigurationState)

	result["configuration_type"] = string(obj.ConfigurationType)

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

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastModified != nil {
		result["time_last_modified"] = obj.TimeLastModified.String()
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentLoggingDestination(fieldKeyFormat string) (oci_logging.UnifiedAgentLoggingDestination, error) {
	result := oci_logging.UnifiedAgentLoggingDestination{}

	if logObjectId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_object_id")); ok {
		tmp := logObjectId.(string)
		result.LogObjectId = &tmp
	}

	return result, nil
}

func UnifiedAgentLoggingDestinationToMap(obj *oci_logging.UnifiedAgentLoggingDestination) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogObjectId != nil {
		result["log_object_id"] = string(*obj.LogObjectId)
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentLoggingSource(fieldKeyFormat string) (oci_logging.UnifiedAgentLoggingSource, error) {
	var baseObject oci_logging.UnifiedAgentLoggingSource
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("LOG_TAIL"):
		details := oci_logging.UnifiedAgentTailLogSource{}
		if parser, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser")); ok {
			if tmpList := parser.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parser"), 0)
				tmp, err := s.mapToUnifiedAgentParser(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert parser, encountered error: %v", err)
				}
				details.Parser = tmp
			}
		}
		if paths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "paths")); ok {
			interfaces := paths.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "paths")) {
				details.Paths = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("WINDOWS_EVENT_LOG"):
		details := oci_logging.UnifiedAgentWindowsEventSource{}
		if channels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "channels")); ok {
			interfaces := channels.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "channels")) {
				details.Channels = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func UnifiedAgentLoggingSourceToMap(obj oci_logging.UnifiedAgentLoggingSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_logging.UnifiedAgentTailLogSource:
		result["source_type"] = "LOG_TAIL"

		if v.Parser != nil {
			parserArray := []interface{}{}
			if parserMap := UnifiedAgentParserToMap(&v.Parser); parserMap != nil {
				parserArray = append(parserArray, parserMap)
			}
			result["parser"] = parserArray
		}
		result["name"] = v.Name
		result["paths"] = v.Paths
	case oci_logging.UnifiedAgentWindowsEventSource:
		result["source_type"] = "WINDOWS_EVENT_LOG"
		result["name"] = v.Name
		result["channels"] = v.Channels
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentParser(fieldKeyFormat string) (oci_logging.UnifiedAgentParser, error) {
	var baseObject oci_logging.UnifiedAgentParser
	//discriminator
	parserTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parser_type"))
	var parserType string
	if ok {
		parserType = parserTypeRaw.(string)
	} else {
		parserType = "" // default value
	}
	switch strings.ToLower(parserType) {
	case strings.ToLower("APACHE2"):
		details := oci_logging.UnifiedAgentApache2Parser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("APACHE_ERROR"):
		details := oci_logging.UnifiedAgentApacheErrorParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("AUDITD"):
		details := oci_logging.UnifiedAgentAuditdParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("CSV"):
		details := oci_logging.UnifiedAgentCsvParser{}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if keys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keys")); ok {
			interfaces := keys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keys")) {
				details.Keys = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("GROK"):
		details := oci_logging.UnifiedAgentGrokParser{}
		if grokFailureKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_failure_key")); ok {
			tmp := grokFailureKey.(string)
			details.GrokFailureKey = &tmp
		}
		if grokNameKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_name_key")); ok {
			tmp := grokNameKey.(string)
			details.GrokNameKey = &tmp
		}
		if patterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patterns")); ok {
			interfaces := patterns.([]interface{})
			tmp := make([]oci_logging.GrokPattern, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patterns"), stateDataIndex)
				converted, err := s.mapToGrokPattern(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patterns")) {
				details.Patterns = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("JSON"):
		details := oci_logging.UnifiedJsonParser{}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if timeType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_type")); ok {
			details.TimeType = oci_logging.UnifiedJsonParserTimeTypeEnum(timeType.(string))
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MSGPACK"):
		details := oci_logging.UnifiedAgentMsgpackParser{}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MULTILINE"):
		details := oci_logging.UnifiedAgentMultilineParser{}
		if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
			interfaces := format.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "format")) {
				details.Format = tmp
			}
		}
		if formatFirstline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format_firstline")); ok {
			tmp := formatFirstline.(string)
			details.FormatFirstline = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("MULTILINE_GROK"):
		details := oci_logging.UnifiedAgentMultilineGrokParser{}
		if grokFailureKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_failure_key")); ok {
			tmp := grokFailureKey.(string)
			details.GrokFailureKey = &tmp
		}
		if grokNameKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "grok_name_key")); ok {
			tmp := grokNameKey.(string)
			details.GrokNameKey = &tmp
		}
		if multiLineStartRegexp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "multi_line_start_regexp")); ok {
			tmp := multiLineStartRegexp.(string)
			details.MultiLineStartRegexp = &tmp
		}
		if patterns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "patterns")); ok {
			interfaces := patterns.([]interface{})
			tmp := make([]oci_logging.GrokPattern, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "patterns"), stateDataIndex)
				converted, err := s.mapToGrokPattern(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "patterns")) {
				details.Patterns = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_logging.UnifiedAgentNoneParser{}
		if messageKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message_key")); ok {
			tmp := messageKey.(string)
			details.MessageKey = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("REGEXP"):
		details := oci_logging.UnifiedAgentRegexParser{}
		if expression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "expression")); ok {
			tmp := expression.(string)
			details.Expression = &tmp
		}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("SYSLOG"):
		details := oci_logging.UnifiedAgentSyslogParser{}
		if isSupportColonlessIdent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_support_colonless_ident")); ok {
			tmp := isSupportColonlessIdent.(bool)
			details.IsSupportColonlessIdent = &tmp
		}
		if isWithPriority, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_with_priority")); ok {
			tmp := isWithPriority.(bool)
			details.IsWithPriority = &tmp
		}
		if messageFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message_format")); ok {
			details.MessageFormat = oci_logging.UnifiedAgentSyslogParserMessageFormatEnum(messageFormat.(string))
		}
		if rfc5424TimeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rfc5424time_format")); ok {
			tmp := rfc5424TimeFormat.(string)
			details.Rfc5424TimeFormat = &tmp
		}
		if syslogParserType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "syslog_parser_type")); ok {
			details.SyslogParserType = oci_logging.UnifiedAgentSyslogParserSyslogParserTypeEnum(syslogParserType.(string))
		}
		if timeFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_format")); ok {
			tmp := timeFormat.(string)
			details.TimeFormat = &tmp
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	case strings.ToLower("TSV"):
		details := oci_logging.UnifiedAgentTsvParser{}
		if delimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delimiter")); ok {
			tmp := delimiter.(string)
			details.Delimiter = &tmp
		}
		if keys, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "keys")); ok {
			interfaces := keys.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "keys")) {
				details.Keys = tmp
			}
		}
		if fieldTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "field_time_key")); ok {
			tmp := fieldTimeKey.(string)
			details.FieldTimeKey = &tmp
		}
		if isEstimateCurrentEvent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_estimate_current_event")); ok {
			tmp := isEstimateCurrentEvent.(bool)
			details.IsEstimateCurrentEvent = &tmp
		}
		if isKeepTimeKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_keep_time_key")); ok {
			tmp := isKeepTimeKey.(bool)
			details.IsKeepTimeKey = &tmp
		}
		if isNullEmptyString, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_null_empty_string")); ok {
			tmp := isNullEmptyString.(bool)
			details.IsNullEmptyString = &tmp
		}
		if nullValuePattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "null_value_pattern")); ok {
			tmp := nullValuePattern.(string)
			details.NullValuePattern = &tmp
		}
		if timeoutInMilliseconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_milliseconds")); ok {
			tmp := timeoutInMilliseconds.(int)
			details.TimeoutInMilliseconds = &tmp
		}
		if types, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "types")); ok {
			details.Types = utils.ObjectMapToStringMap(types.(map[string]interface{}))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown parser_type '%v' was specified", parserType)
	}
	return baseObject, nil
}

func UnifiedAgentParserToMap(obj *oci_logging.UnifiedAgentParser) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_logging.UnifiedAgentApache2Parser:
		result["parser_type"] = "APACHE2"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentApacheErrorParser:
		result["parser_type"] = "APACHE_ERROR"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentAuditdParser:
		result["parser_type"] = "AUDITD"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentCsvParser:
		result["parser_type"] = "CSV"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		if v.Delimiter != nil {
			result["delimiter"] = string(*v.Delimiter)
		}

		result["keys"] = v.Keys
	case oci_logging.UnifiedAgentGrokParser:
		result["parser_type"] = "GROK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}

		if v.GrokFailureKey != nil {
			result["grok_failure_key"] = string(*v.GrokFailureKey)
		}

		if v.GrokNameKey != nil {
			result["grok_name_key"] = string(*v.GrokNameKey)
		}

		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}

		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}

		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}

		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}

		patterns := []interface{}{}
		for _, item := range v.Patterns {
			patterns = append(patterns, GrokPatternToMap(item))
		}
		result["patterns"] = patterns
	case oci_logging.UnifiedJsonParser:
		result["parser_type"] = "JSON"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}

		result["time_type"] = string(v.TimeType)
	case oci_logging.UnifiedAgentMsgpackParser:
		result["parser_type"] = "MSGPACK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
	case oci_logging.UnifiedAgentMultilineParser:
		result["parser_type"] = "MULTILINE"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		result["format"] = v.Format

		if v.FormatFirstline != nil {
			result["format_firstline"] = string(*v.FormatFirstline)
		}
	case oci_logging.UnifiedAgentMultilineGrokParser:
		result["parser_type"] = "MULTILINE_GROK"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		if v.GrokFailureKey != nil {
			result["grok_failure_key"] = string(*v.GrokFailureKey)
		}

		if v.GrokNameKey != nil {
			result["grok_name_key"] = string(*v.GrokNameKey)
		}

		if v.MultiLineStartRegexp != nil {
			result["multi_line_start_regexp"] = string(*v.MultiLineStartRegexp)
		}

		patterns := []interface{}{}
		for _, item := range v.Patterns {
			patterns = append(patterns, GrokPatternToMap(item))
		}
		result["patterns"] = patterns
	case oci_logging.UnifiedAgentNoneParser:
		result["parser_type"] = "NONE"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		if v.MessageKey != nil {
			result["message_key"] = string(*v.MessageKey)
		}
	case oci_logging.UnifiedAgentRegexParser:
		result["parser_type"] = "REGEXP"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}

		if v.Expression != nil {
			result["expression"] = string(*v.Expression)
		}

		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}
	case oci_logging.UnifiedAgentSyslogParser:
		result["parser_type"] = "SYSLOG"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}

		if v.IsSupportColonlessIdent != nil {
			result["is_support_colonless_ident"] = bool(*v.IsSupportColonlessIdent)
		}

		if v.IsWithPriority != nil {
			result["is_with_priority"] = bool(*v.IsWithPriority)
		}

		result["message_format"] = string(v.MessageFormat)

		if v.Rfc5424TimeFormat != nil {
			result["rfc5424time_format"] = string(*v.Rfc5424TimeFormat)
		}

		result["syslog_parser_type"] = string(v.SyslogParserType)

		if v.TimeFormat != nil {
			result["time_format"] = string(*v.TimeFormat)
		}
	case oci_logging.UnifiedAgentTsvParser:
		result["parser_type"] = "TSV"
		if v.FieldTimeKey != nil {
			result["field_time_key"] = v.FieldTimeKey
		}
		if v.IsEstimateCurrentEvent != nil {
			result["is_estimate_current_event"] = v.IsEstimateCurrentEvent
		}
		if v.IsKeepTimeKey != nil {
			result["is_keep_time_key"] = v.IsKeepTimeKey
		}
		if v.IsNullEmptyString != nil {
			result["is_null_empty_string"] = v.IsNullEmptyString
		}
		if v.NullValuePattern != nil {
			result["null_value_pattern"] = v.NullValuePattern
		}
		if v.TimeoutInMilliseconds != nil {
			result["timeout_in_milliseconds"] = v.TimeoutInMilliseconds
		}
		if v.Types != nil {
			result["types"] = utils.StringMapToObjectMap(v.Types)
		}
		if v.Delimiter != nil {
			result["delimiter"] = string(*v.Delimiter)
		}

		result["keys"] = v.Keys
	default:
		log.Printf("[WARN] Received 'parser_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) mapToUnifiedAgentServiceConfigurationDetails(fieldKeyFormat string) (oci_logging.UnifiedAgentServiceConfigurationDetails, error) {
	var baseObject oci_logging.UnifiedAgentServiceConfigurationDetails
	//discriminator
	configurationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration_type"))
	var configurationType string
	if ok {
		configurationType = configurationTypeRaw.(string)
	} else {
		configurationType = "" // default value
	}
	switch strings.ToLower(configurationType) {
	case strings.ToLower("LOGGING"):
		details := oci_logging.UnifiedAgentLoggingConfiguration{}
		if destination, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination")); ok {
			if tmpList := destination.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "destination"), 0)
				tmp, err := s.mapToUnifiedAgentLoggingDestination(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert destination, encountered error: %v", err)
				}
				details.Destination = &tmp
			}
		}
		if sources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sources")); ok {
			interfaces := sources.([]interface{})
			tmp := make([]oci_logging.UnifiedAgentLoggingSource, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "sources"), stateDataIndex)
				converted, err := s.mapToUnifiedAgentLoggingSource(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "sources")) {
				details.Sources = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown configuration_type '%v' was specified", configurationType)
	}
	return baseObject, nil
}

func UnifiedAgentServiceConfigurationDetailsToMap(obj *oci_logging.UnifiedAgentServiceConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_logging.UnifiedAgentLoggingConfiguration:
		result["configuration_type"] = "LOGGING"

		if v.Destination != nil {
			result["destination"] = []interface{}{UnifiedAgentLoggingDestinationToMap(v.Destination)}
		}

		sources := []interface{}{}
		for _, item := range v.Sources {
			sources = append(sources, UnifiedAgentLoggingSourceToMap(item))
		}
		result["sources"] = sources
	default:
		log.Printf("[WARN] Received 'configuration_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *LoggingUnifiedAgentConfigurationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_logging.ChangeUnifiedAgentConfigurationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.UnifiedAgentConfigurationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging")

	response, err := s.Client.ChangeUnifiedAgentConfigurationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getUnifiedAgentConfigurationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "logging"), oci_logging.ActionTypesRelated, s.D.Timeout(schema.TimeoutUpdate))
}
