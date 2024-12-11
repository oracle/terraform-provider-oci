// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package fleet_apps_management

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
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FleetAppsManagementTaskRecordResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementTaskRecord,
		Read:     readFleetAppsManagementTaskRecord,
		Update:   updateFleetAppsManagementTaskRecord,
		Delete:   deleteFleetAppsManagementTaskRecord,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"execution_details": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"execution_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"API",
											"SCRIPT",
										}, true),
									},

									// Optional
									"command": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"bucket": {
													Type:     schema.TypeString,
													Required: true,
												},
												"checksum": {
													Type:     schema.TypeString,
													Required: true,
												},
												"namespace": {
													Type:     schema.TypeString,
													Required: true,
												},
												"object": {
													Type:     schema.TypeString,
													Required: true,
												},
												"source_type": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"OBJECT_STORAGE_BUCKET",
													}, true),
												},

												// Optional

												// Computed
											},
										},
									},
									"credentials": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"display_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"endpoint": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"variables": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"input_variables": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"description": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"name": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},
												"output_variables": {
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
						"os_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"scope": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"is_apply_subject_task": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_discovery_output_task": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"platform": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"num_retries": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"timeout_in_seconds": {
										Type:     schema.TypeInt,
										Required: true,
									},

									// Optional

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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_region": {
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createFleetAppsManagementTaskRecord(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementTaskRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementTaskRecord(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementTaskRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementTaskRecord(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementTaskRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementTaskRecord(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementTaskRecordResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementRunbooksClient()
	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementTaskRecordResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementRunbooksClient
	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
	Res                    *oci_fleet_apps_management.TaskRecord
	DisableNotFoundRetries bool
}

func (s *FleetAppsManagementTaskRecordResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementTaskRecordResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *FleetAppsManagementTaskRecordResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.TaskRecordLifecycleStateActive),
	}
}

func (s *FleetAppsManagementTaskRecordResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.TaskRecordLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementTaskRecordResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.TaskRecordLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementTaskRecordResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateTaskRecordRequest{}

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

	if details, ok := s.D.GetOkExists("details"); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "details", 0)
			tmp, err := s.mapToDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Details = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateTaskRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskRecord
	return nil
}

func (s *FleetAppsManagementTaskRecordResourceCrud) getTaskRecordFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	taskRecordId, err := taskRecordWaitForWorkRequest(workId, "task",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)

	if err != nil {
		return err
	}
	s.D.SetId(*taskRecordId)

	return s.Get()
}

func taskRecordWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func taskRecordWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = taskRecordWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_fleet_apps_management.OperationStatusInProgress),
			string(oci_fleet_apps_management.OperationStatusAccepted),
			string(oci_fleet_apps_management.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_fleet_apps_management.OperationStatusSucceeded),
			string(oci_fleet_apps_management.OperationStatusFailed),
			string(oci_fleet_apps_management.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_fleet_apps_management.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementTaskRecordWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementTaskRecordWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
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

func (s *FleetAppsManagementTaskRecordResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetTaskRecordRequest{}

	tmp := s.D.Id()
	request.TaskRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetTaskRecord(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskRecord
	return nil
}

func (s *FleetAppsManagementTaskRecordResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateTaskRecordRequest{}

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

	if details, ok := s.D.GetOkExists("details"); ok {
		if tmpList := details.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "details", 0)
			tmp, err := s.mapToDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Details = &tmp
		}
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.TaskRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateTaskRecord(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getTaskRecordFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementTaskRecordResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteTaskRecordRequest{}

	tmp := s.D.Id()
	request.TaskRecordId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.DeleteTaskRecord(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := taskRecordWaitForWorkRequest(workId, "task",
		oci_fleet_apps_management.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.FleetClient)
	return delWorkRequestErr
}

func (s *FleetAppsManagementTaskRecordResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Details != nil {
		s.D.Set("details", []interface{}{DetailsToMap(s.Res.Details)})
	} else {
		s.D.Set("details", nil)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
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

	s.D.Set("type", s.Res.Type)

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	return nil
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToConfigAssociationDetails(fieldKeyFormat string) (oci_fleet_apps_management.ConfigAssociationDetails, error) {
	result := oci_fleet_apps_management.ConfigAssociationDetails{}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToContentDetails(fieldKeyFormat string) (oci_fleet_apps_management.ContentDetails, error) {
	var baseObject oci_fleet_apps_management.ContentDetails
	//discriminator
	sourceTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_type"))
	var sourceType string
	if ok {
		sourceType = sourceTypeRaw.(string)
	} else {
		sourceType = "" // default value
	}
	switch strings.ToLower(sourceType) {
	case strings.ToLower("OBJECT_STORAGE_BUCKET"):
		details := oci_fleet_apps_management.ObjectStorageBucketContentDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if checksum, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "checksum")); ok {
			tmp := checksum.(string)
			details.Checksum = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if object, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object")); ok {
			tmp := object.(string)
			details.ObjectName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown source_type '%v' was specified", sourceType)
	}
	return baseObject, nil
}

func ContentDetailsToMap(obj *oci_fleet_apps_management.ContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.ObjectStorageBucketContentDetails:
		result["source_type"] = "OBJECT_STORAGE_BUCKET"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.Checksum != nil {
			result["checksum"] = string(*v.Checksum)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		if v.ObjectName != nil {
			result["object"] = string(*v.ObjectName)
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToDetails(fieldKeyFormat string) (oci_fleet_apps_management.Details, error) {
	result := oci_fleet_apps_management.Details{}

	if executionDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_details")); ok {
		if tmpList := executionDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "execution_details"), 0)
			tmp, err := s.mapToExecutionDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert execution_details, encountered error: %v", err)
			}
			result.ExecutionDetails = tmp
		}
	}

	if isApplySubjectTask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_apply_subject_task")); ok {
		tmp := isApplySubjectTask.(bool)
		result.IsApplySubjectTask = &tmp
	}

	if isDiscoveryOutputTask, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_discovery_output_task")); ok {
		tmp := isDiscoveryOutputTask.(bool)
		result.IsDiscoveryOutputTask = &tmp
	}

	if osType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "os_type")); ok {
		result.OsType = oci_fleet_apps_management.OsTypeEnum(osType.(string))
	}

	if platform, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "platform")); ok {
		tmp := platform.(string)
		result.Platform = &tmp
	}

	if properties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
		if tmpList := properties.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "properties"), 0)
			tmp, err := s.mapToProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert properties, encountered error: %v", err)
			}
			result.Properties = &tmp
		}
	}

	if scope, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scope")); ok {
		result.Scope = oci_fleet_apps_management.TaskScopeEnum(scope.(string))
	}

	return result, nil
}

func DetailsToMap(obj *oci_fleet_apps_management.Details) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ExecutionDetails != nil {
		executionDetailsArray := []interface{}{}
		if executionDetailsMap := ExecutionDetailsToMap(&obj.ExecutionDetails); executionDetailsMap != nil {
			executionDetailsArray = append(executionDetailsArray, executionDetailsMap)
		}
		result["execution_details"] = executionDetailsArray
	}

	if obj.IsApplySubjectTask != nil {
		result["is_apply_subject_task"] = bool(*obj.IsApplySubjectTask)
	}

	if obj.IsDiscoveryOutputTask != nil {
		result["is_discovery_output_task"] = bool(*obj.IsDiscoveryOutputTask)
	}

	result["os_type"] = string(obj.OsType)

	if obj.Platform != nil {
		result["platform"] = string(*obj.Platform)
	}

	if obj.Properties != nil {
		result["properties"] = []interface{}{PropertiesToMap(obj.Properties)}
	}

	result["scope"] = string(obj.Scope)

	return result
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToExecutionDetails(fieldKeyFormat string) (oci_fleet_apps_management.ExecutionDetails, error) {
	var baseObject oci_fleet_apps_management.ExecutionDetails
	//discriminator
	executionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_type"))
	var executionType string
	if ok {
		executionType = executionTypeRaw.(string)
	} else {
		executionType = "" // default value
	}
	switch strings.ToLower(executionType) {
	case strings.ToLower("API"):
		details := oci_fleet_apps_management.ApiBasedExecutionDetails{}
		if endpoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("SCRIPT"):
		details := oci_fleet_apps_management.ScriptBasedExecutionDetails{}
		if command, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command")); ok {
			tmp := command.(string)
			details.Command = &tmp
		}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			if tmpList := content.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
				tmp, err := s.mapToContentDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert content, encountered error: %v", err)
				}
				details.Content = tmp
			}
		}
		if credentials, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "credentials")); ok {
			interfaces := credentials.([]interface{})
			tmp := make([]oci_fleet_apps_management.ConfigAssociationDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "credentials"), stateDataIndex)
				converted, err := s.mapToConfigAssociationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "credentials")) {
				details.Credentials = tmp
			}
		}
		if variables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "variables")); ok {
			if tmpList := variables.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "variables"), 0)
				tmp, err := s.mapToTaskVariable(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert variables, encountered error: %v", err)
				}
				details.Variables = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown execution_type '%v' was specified", executionType)
	}
	return baseObject, nil
}

func ExecutionDetailsToMap(obj *oci_fleet_apps_management.ExecutionDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.ApiBasedExecutionDetails:
		result["execution_type"] = "API"

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_fleet_apps_management.ScriptBasedExecutionDetails:
		result["execution_type"] = "SCRIPT"

		if v.Command != nil {
			result["command"] = string(*v.Command)
		}

		if v.Content != nil {
			contentArray := []interface{}{}
			if contentMap := ContentDetailsToMap(&v.Content); contentMap != nil {
				contentArray = append(contentArray, contentMap)
			}
			result["content"] = contentArray
		}

		credentials := []interface{}{}
		for _, item := range v.Credentials {
			credentials = append(credentials, ConfigAssociationDetailsToMap(item))
		}
		result["credentials"] = credentials

		if v.Variables != nil {
			result["variables"] = []interface{}{TaskVariableToMap(v.Variables)}
		}
	default:
		log.Printf("[WARN] Received 'execution_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToInputArgument(fieldKeyFormat string) (oci_fleet_apps_management.InputArgument, error) {
	var baseObject oci_fleet_apps_management.InputArgument
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	nameRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name"))
	var name_ string
	if ok {
		name_ = nameRaw.(string)
	} else {
		name_ = ""
	}
	descriptionRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description"))
	var description string
	if ok {
		description = descriptionRaw.(string)
	} else {
		description = ""
	}

	switch strings.ToLower(type_) {
	case strings.ToLower("OUTPUT_VARIABLE"):
		details := oci_fleet_apps_management.OutputVariableInputArgument{Name: &name_, Description: &description}
		baseObject = details
	case strings.ToLower("STRING"):
		details := oci_fleet_apps_management.StringInputArgument{Name: &name_, Description: &description}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func InputArgumentToMap(obj oci_fleet_apps_management.InputArgument) map[string]interface{} {
	result := map[string]interface{}{}
	switch (obj).(type) {
	case oci_fleet_apps_management.OutputVariableInputArgument:
		result["type"] = "OUTPUT_VARIABLE"
	case oci_fleet_apps_management.StringInputArgument:
		result["type"] = "STRING"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", obj)
		return nil
	}
	result["name"] = obj.GetName()
	result["description"] = obj.GetDescription()

	return result
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToProperties(fieldKeyFormat string) (oci_fleet_apps_management.Properties, error) {
	result := oci_fleet_apps_management.Properties{}

	if numRetries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_retries")); ok {
		tmp := numRetries.(int)
		result.NumRetries = &tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_seconds")); ok {
		tmp := timeoutInSeconds.(int)
		result.TimeoutInSeconds = &tmp
	}

	return result, nil
}

func PropertiesToMap(obj *oci_fleet_apps_management.Properties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NumRetries != nil {
		result["num_retries"] = int(*obj.NumRetries)
	}

	if obj.TimeoutInSeconds != nil {
		result["timeout_in_seconds"] = int(*obj.TimeoutInSeconds)
	}

	return result
}

func TaskRecordSummaryToMap(obj oci_fleet_apps_management.TaskRecordSummary) map[string]interface{} {
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

	if obj.Details != nil {
		result["details"] = []interface{}{DetailsToMap(obj.Details)}
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

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
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

	result["type"] = string(obj.Type)

	if obj.Version != nil {
		result["version"] = string(*obj.Version)
	}

	return result
}

func (s *FleetAppsManagementTaskRecordResourceCrud) mapToTaskVariable(fieldKeyFormat string) (oci_fleet_apps_management.TaskVariable, error) {
	result := oci_fleet_apps_management.TaskVariable{}

	if inputVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_variables")); ok {
		interfaces := inputVariables.([]interface{})
		tmp := make([]oci_fleet_apps_management.InputArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "input_variables"), stateDataIndex)
			converted, err := s.mapToInputArgument(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "input_variables")) {
			result.InputVariables = tmp
		}
	}

	if outputVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_variables")); ok {
		interfaces := outputVariables.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "output_variables")) {
			result.OutputVariables = tmp
		}
	}

	return result, nil
}

func TaskVariableToMap(obj *oci_fleet_apps_management.TaskVariable) map[string]interface{} {
	result := map[string]interface{}{}

	inputVariables := []interface{}{}
	for _, item := range obj.InputVariables {
		inputVariables = append(inputVariables, InputArgumentToMap(item))
	}
	result["input_variables"] = inputVariables

	result["output_variables"] = obj.OutputVariables

	return result
}
