// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// /Licensed under the Mozilla Public License v2.0

package fleet_apps_management

import (
	"context"
	"math/rand"

	// "crypto/rand"
	"fmt"
	"log"
	"strings"
	"time"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"
	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func FleetAppsManagementSchedulerDefinitionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createFleetAppsManagementSchedulerDefinition,
		Read:     readFleetAppsManagementSchedulerDefinition,
		Update:   updateFleetAppsManagementSchedulerDefinition,
		Delete:   deleteFleetAppsManagementSchedulerDefinition,
		Schema: map[string]*schema.Schema{
			// Required
			"action_groups": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"fleet_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"kind": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FLEET_USING_RUNBOOK",
							}, true),
						},
						"runbook_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"runbook_version_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sequence": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"execution_startdate": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CUSTOM",
								"MAINTENANCE_WINDOW",
							}, true),
						},

						// Optional
						"duration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"maintenance_window_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"recurrences": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"run_books": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"runbook_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"runbook_version_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"input_parameters": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"step_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"arguments": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"kind": {
													Type:             schema.TypeString,
													Required:         true,
													DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
													ValidateFunc: validation.StringInSlice([]string{
														"FILE",
														"STRING",
													}, true),
												},
												"name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
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
												"value": {
													Type:     schema.TypeString,
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

						// Computed
					},
				},
			},

			// Computed
			"count_of_affected_action_groups": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"count_of_affected_resources": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"count_of_affected_targets": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"products": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"time_of_next_run": {
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

func createFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()

	return tfresource.ReadResource(sync)
}

func updateFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).FleetAppsManagementFleetAppsManagementWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type FleetAppsManagementSchedulerDefinitionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_fleet_apps_management.FleetAppsManagementOperationsClient
	Res                    *oci_fleet_apps_management.SchedulerDefinition
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateCreating),
	}
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateActive),
	}
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleting),
	}
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted),
	}
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Create() error {
	request := oci_fleet_apps_management.CreateSchedulerDefinitionRequest{}

	if actionGroups, ok := s.D.GetOkExists("action_groups"); ok {
		interfaces := actionGroups.([]interface{})
		tmp := make([]oci_fleet_apps_management.ActionGroup, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_groups", stateDataIndex)
			converted, err := s.mapToActionGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("action_groups") {
			request.ActionGroups = tmp
		}
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

	// Added to randomize the OpcRetryToken
	rand.Seed(time.Now().UnixNano())
	request.OpcRetryToken = oci_common.String(string(randSeq(20)))

	// var DisplayNameRandomString string = string(randSeq(20))
	// request.DisplayName = oci_common.String(DisplayNameRandomString)
	// fmt.Printf("Requested display name is: %s \n", *request.DisplayName)

	if runBooks, ok := s.D.GetOkExists("run_books"); ok {
		interfaces := runBooks.([]interface{})
		tmp := make([]oci_fleet_apps_management.OperationRunbook, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "run_books", stateDataIndex)
			converted, err := s.mapToOperationRunbook(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("run_books") {
			request.RunBooks = tmp
		}
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
			tmp, err := s.mapToSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schedule = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.CreateSchedulerDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSchedulerDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) getSchedulerDefinitionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	schedulerDefinitionId, err := schedulerDefinitionWaitForWorkRequest(workId, "schedulerdefinition",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*schedulerDefinitionId)

	return s.Get()
}

func schedulerDefinitionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func schedulerDefinitionWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
	retryPolicy.ShouldRetryOperation = schedulerDefinitionWorkRequestShouldRetryFunc(timeout)

	response := oci_fleet_apps_management.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
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
		// fmt.Printf("============================")
		// fmt.Printf("Value of action: %s\n", action)
		// fmt.Printf("Value of res.ActionType: %s\n", res.ActionType)
		// fmt.Printf("Value of entityType: %s\n", entityType)
		// fmt.Printf("Value of *res.EntityType: %s\n", *res.EntityType)
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				// fmt.Printf("============breaking================")
				break
			}
		}
	}

	if identifier != nil {
		fmt.Printf("Value of identifier: %s\n", *identifier)
	} else {
		fmt.Printf("Value of identifier: <nil>\n")
	}

	fmt.Printf("Value of response status: %s\n", response.Status)

	// fmt.Printf("Full response: %+v\n", response)

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
		return nil, getErrorFromFleetAppsManagementSchedulerDefinitionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromFleetAppsManagementSchedulerDefinitionWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementWorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
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

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Get() error {
	request := oci_fleet_apps_management.GetSchedulerDefinitionRequest{}

	tmp := s.D.Id()
	request.SchedulerDefinitionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.GetSchedulerDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SchedulerDefinition
	return nil
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Update() error {
	request := oci_fleet_apps_management.UpdateSchedulerDefinitionRequest{}

	if actionGroups, ok := s.D.GetOkExists("action_groups"); ok {
		interfaces := actionGroups.([]interface{})
		tmp := make([]oci_fleet_apps_management.ActionGroup, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_groups", stateDataIndex)
			converted, err := s.mapToActionGroup(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("action_groups") {
			request.ActionGroups = tmp
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

	if runBooks, ok := s.D.GetOkExists("run_books"); ok {
		interfaces := runBooks.([]interface{})
		tmp := make([]oci_fleet_apps_management.OperationRunbook, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "run_books", stateDataIndex)
			converted, err := s.mapToOperationRunbook(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("run_books") {
			request.RunBooks = tmp
		}
	}

	if schedule, ok := s.D.GetOkExists("schedule"); ok {
		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
			tmp, err := s.mapToSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Schedule = tmp
		}
	}

	tmp := s.D.Id()
	request.SchedulerDefinitionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	response, err := s.Client.UpdateSchedulerDefinition(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSchedulerDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Delete() error {
	request := oci_fleet_apps_management.DeleteSchedulerDefinitionRequest{}

	tmp := s.D.Id()
	request.SchedulerDefinitionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")

	_, err := s.Client.DeleteSchedulerDefinition(context.Background(), request)
	return err
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) SetData() error {
	actionGroups := []interface{}{}
	for _, item := range s.Res.ActionGroups {
		actionGroups = append(actionGroups, ActionGroupToMap(item))
	}
	s.D.Set("action_groups", actionGroups)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	} else {
		s.D.Set("compartment_id", nil)
	}

	if s.Res.CountOfAffectedActionGroups != nil {
		s.D.Set("count_of_affected_action_groups", *s.Res.CountOfAffectedActionGroups)
	}

	if s.Res.CountOfAffectedResources != nil {
		s.D.Set("count_of_affected_resources", *s.Res.CountOfAffectedResources)
	}

	if s.Res.CountOfAffectedTargets != nil {
		s.D.Set("count_of_affected_targets", *s.Res.CountOfAffectedTargets)
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

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("lifecycle_operations", s.Res.LifecycleOperations)

	s.D.Set("products", s.Res.Products)

	if s.Res.ResourceRegion != nil {
		s.D.Set("resource_region", *s.Res.ResourceRegion)
	}

	runBooks := []interface{}{}
	for _, item := range s.Res.RunBooks {
		runBooks = append(runBooks, OperationRunbookToMap(item))
	}
	s.D.Set("run_books", runBooks)

	if s.Res.Schedule != nil {
		scheduleArray := []interface{}{}
		if scheduleMap := ScheduleToMap(&s.Res.Schedule); scheduleMap != nil {
			scheduleArray = append(scheduleArray, scheduleMap)
		}
		s.D.Set("schedule", scheduleArray)
	} else {
		s.D.Set("schedule", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfNextRun != nil {
		s.D.Set("time_of_next_run", s.Res.TimeOfNextRun.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToActionGroup(fieldKeyFormat string) (oci_fleet_apps_management.ActionGroup, error) {
	var baseObject oci_fleet_apps_management.ActionGroup
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("FLEET_USING_RUNBOOK"):
		details := oci_fleet_apps_management.FleetBasedActionGroup{}
		if fleetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fleet_id")); ok {
			tmp := fleetId.(string)
			details.FleetId = &tmp
		}
		if runbookId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runbook_id")); ok {
			tmp := runbookId.(string)
			details.RunbookId = &tmp
		}
		if runbookVersionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runbook_version_name")); ok {
			tmp := runbookVersionName.(string)
			details.RunbookVersionName = &tmp
		}
		if sequence, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sequence")); ok {
			tmp := sequence.(int)
			details.Sequence = &tmp
		}
		if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func ActionGroupToMap(obj oci_fleet_apps_management.ActionGroup) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_apps_management.FleetBasedActionGroup:
		result["kind"] = "FLEET_USING_RUNBOOK"

		if v.FleetId != nil {
			result["fleet_id"] = string(*v.FleetId)
		}

		if v.RunbookId != nil {
			result["runbook_id"] = string(*v.RunbookId)
		}

		if v.RunbookVersionName != nil {
			result["runbook_version_name"] = string(*v.RunbookVersionName)
		}

		if v.Sequence != nil {
			result["sequence"] = int(*v.Sequence)
		}

		if v.DisplayName != nil {
			result["display_name"] = string(*v.DisplayName)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToInputFileContentDetails(fieldKeyFormat string) (oci_fleet_apps_management.InputFileContentDetails, error) {
	var baseObject oci_fleet_apps_management.InputFileContentDetails
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
		details := oci_fleet_apps_management.InputFileObjectStorageBucketContentDetails{}
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

func InputFileContentDetailsToMap(obj *oci_fleet_apps_management.InputFileContentDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.InputFileObjectStorageBucketContentDetails:
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

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToInputParameter(fieldKeyFormat string) (oci_fleet_apps_management.InputParameter, error) {
	result := oci_fleet_apps_management.InputParameter{}

	if arguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arguments")); ok {
		interfaces := arguments.([]interface{})
		tmp := make([]oci_fleet_apps_management.TaskArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "arguments"), stateDataIndex)
			converted, err := s.mapToTaskArgument(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "arguments")) {
			result.Arguments = tmp
		}
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		result.StepName = &tmp
	}

	return result, nil
}

func InputParameterToMap(obj oci_fleet_apps_management.InputParameter) map[string]interface{} {
	result := map[string]interface{}{}

	arguments := []interface{}{}
	for _, item := range obj.Arguments {
		arguments = append(arguments, TaskArgumentToMap(item))
	}
	result["arguments"] = arguments

	if obj.StepName != nil {
		result["step_name"] = string(*obj.StepName)
	}

	return result
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToOperationRunbook(fieldKeyFormat string) (oci_fleet_apps_management.OperationRunbook, error) {
	result := oci_fleet_apps_management.OperationRunbook{}

	if inputParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_parameters")); ok {
		interfaces := inputParameters.([]interface{})
		tmp := make([]oci_fleet_apps_management.InputParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "input_parameters"), stateDataIndex)
			converted, err := s.mapToInputParameter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "input_parameters")) {
			result.InputParameters = tmp
		}
	}

	if runbookId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runbook_id")); ok {
		tmp := runbookId.(string)
		result.RunbookId = &tmp
	}

	if runbookVersionName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "runbook_version_name")); ok {
		tmp := runbookVersionName.(string)
		result.RunbookVersionName = &tmp
	}

	return result, nil
}

func OperationRunbookToMap(obj oci_fleet_apps_management.OperationRunbook) map[string]interface{} {
	result := map[string]interface{}{}

	inputParameters := []interface{}{}
	for _, item := range obj.InputParameters {
		inputParameters = append(inputParameters, InputParameterToMap(item))
	}
	result["input_parameters"] = inputParameters

	if obj.RunbookId != nil {
		result["runbook_id"] = string(*obj.RunbookId)
	}

	if obj.RunbookVersionName != nil {
		result["runbook_version_name"] = string(*obj.RunbookVersionName)
	}

	return result
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToSchedule(fieldKeyFormat string) (oci_fleet_apps_management.Schedule, error) {
	var baseObject oci_fleet_apps_management.Schedule
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("CUSTOM"):
		details := oci_fleet_apps_management.CustomSchedule{}
		if duration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duration")); ok {
			tmp := duration.(string)
			details.Duration = &tmp
		}
		if recurrences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recurrences")); ok {
			tmp := recurrences.(string)
			details.Recurrences = &tmp
		}
		if executionStartdate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_startdate")); ok {
			tmp, err := time.Parse(time.RFC3339, executionStartdate.(string))
			if err != nil {
				return details, err
			}
			details.ExecutionStartdate = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	case strings.ToLower("MAINTENANCE_WINDOW"):
		details := oci_fleet_apps_management.MaintenanceWindowSchedule{}
		if maintenanceWindowId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maintenance_window_id")); ok {
			tmp := maintenanceWindowId.(string)
			details.MaintenanceWindowId = &tmp
		}
		if executionStartdate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_startdate")); ok {
			tmp, err := time.Parse(time.RFC3339, executionStartdate.(string))
			if err != nil {
				return details, err
			}
			details.ExecutionStartdate = &oci_common.SDKTime{Time: tmp}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ScheduleToMap(obj *oci_fleet_apps_management.Schedule) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_fleet_apps_management.CustomSchedule:
		result["type"] = "CUSTOM"

		if v.Duration != nil {
			result["duration"] = string(*v.Duration)
		}

		if v.Recurrences != nil {
			result["recurrences"] = string(*v.Recurrences)
		}

		if v.ExecutionStartdate != nil {
			result["execution_startdate"] = v.ExecutionStartdate.Format(time.RFC3339Nano)
		}
	case oci_fleet_apps_management.MaintenanceWindowSchedule:
		result["type"] = "MAINTENANCE_WINDOW"

		if v.MaintenanceWindowId != nil {
			result["maintenance_window_id"] = string(*v.MaintenanceWindowId)
		}

		if v.ExecutionStartdate != nil {
			result["execution_startdate"] = v.ExecutionStartdate.Format(time.RFC3339Nano)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func SchedulerDefinitionSummaryToMap(obj oci_fleet_apps_management.SchedulerDefinitionSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CountOfAffectedActionGroups != nil {
		result["count_of_affected_action_groups"] = int(*obj.CountOfAffectedActionGroups)
	}

	if obj.CountOfAffectedResources != nil {
		result["count_of_affected_resources"] = int(*obj.CountOfAffectedResources)
	}

	if obj.CountOfAffectedTargets != nil {
		result["count_of_affected_targets"] = int(*obj.CountOfAffectedTargets)
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

	result["lifecycle_operations"] = obj.LifecycleOperations

	result["products"] = obj.Products

	if obj.ResourceRegion != nil {
		result["resource_region"] = string(*obj.ResourceRegion)
	}

	if obj.Schedule != nil {
		scheduleArray := []interface{}{}
		if scheduleMap := ScheduleToMap(&obj.Schedule); scheduleMap != nil {
			scheduleArray = append(scheduleArray, scheduleMap)
		}
		result["schedule"] = scheduleArray
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeOfNextRun != nil {
		result["time_of_next_run"] = obj.TimeOfNextRun.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToTaskArgument(fieldKeyFormat string) (oci_fleet_apps_management.TaskArgument, error) {
	var baseObject oci_fleet_apps_management.TaskArgument
	//discriminator
	kindRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "kind"))
	var kind string
	if ok {
		kind = kindRaw.(string)
	} else {
		kind = "" // default value
	}
	switch strings.ToLower(kind) {
	case strings.ToLower("FILE"):
		details := oci_fleet_apps_management.FileTaskArgument{}
		if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
			if tmpList := content.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "content"), 0)
				tmp, err := s.mapToInputFileContentDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert content, encountered error: %v", err)
				}
				details.Content = tmp
			}
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	case strings.ToLower("STRING"):
		details := oci_fleet_apps_management.StringTaskArgument{}
		if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
			tmp := value.(string)
			details.Value = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown kind '%v' was specified", kind)
	}
	return baseObject, nil
}

func TaskArgumentToMap(obj oci_fleet_apps_management.TaskArgument) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_fleet_apps_management.FileTaskArgument:
		result["kind"] = "FILE"

		if v.Content != nil {
			contentArray := []interface{}{}
			if contentMap := InputFileContentDetailsToMap(&v.Content); contentMap != nil {
				contentArray = append(contentArray, contentMap)
			}
			result["content"] = contentArray
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	case oci_fleet_apps_management.StringTaskArgument:
		result["kind"] = "STRING"

		if v.Value != nil {
			result["value"] = string(*v.Value)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}
	default:
		log.Printf("[WARN] Received 'kind' of unknown type %v", obj)
		return nil
	}

	return result
}

//=======
//
//func createFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
//	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
//	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
//
//	return tfresource.CreateResource(d, sync)
//}
//
//func readFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
//	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
//	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
//
//	return tfresource.ReadResource(sync)
//}
//
//func updateFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
//	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
//	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
//
//	return tfresource.UpdateResource(d, sync)
//}
//
//func deleteFleetAppsManagementSchedulerDefinition(d *schema.ResourceData, m interface{}) error {
//	sync := &FleetAppsManagementSchedulerDefinitionResourceCrud{}
//	sync.D = d
//	sync.Client = m.(*client.OracleClients).FleetAppsManagementOperationsClient()
//	sync.FleetClient = m.(*client.OracleClients).FleetAppsManagementClient()
//	sync.DisableNotFoundRetries = true
//
//	return tfresource.DeleteResource(d, sync)
//}
//
//type FleetAppsManagementSchedulerDefinitionResourceCrud struct {
//	tfresource.BaseCrud
//	Client                 *oci_fleet_apps_management.FleetAppsManagementOperationsClient
//	FleetClient            *oci_fleet_apps_management.FleetAppsManagementClient
//	Res                    *oci_fleet_apps_management.SchedulerDefinition
//	DisableNotFoundRetries bool
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) ID() string {
//	return *s.Res.Id
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) CreatedPending() []string {
//	return []string{
//		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateCreating),
//	}
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) CreatedTarget() []string {
//	return []string{
//		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateActive),
//	}
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) DeletedPending() []string {
//	return []string{
//		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleting),
//	}
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) DeletedTarget() []string {
//	return []string{
//		string(oci_fleet_apps_management.SchedulerDefinitionLifecycleStateDeleted),
//	}
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Create() error {
//	request := oci_fleet_apps_management.CreateSchedulerDefinitionRequest{}
//
//	if actionGroups, ok := s.D.GetOkExists("action_groups"); ok {
//		interfaces := actionGroups.([]interface{})
//		tmp := make([]oci_fleet_apps_management.ActionGroup, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_groups", stateDataIndex)
//			converted, err := s.mapToActionGroup(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange("action_groups") {
//			request.ActionGroups = tmp
//		}
//	}
//
//	//if activityInitiationCutOff, ok := s.D.GetOkExists("activity_initiation_cut_off"); ok {
//	//	tmp := activityInitiationCutOff.(int)
//	//	request.ActivityInitiationCutOff = &tmp
//	//}
//
//	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
//		tmp := compartmentId.(string)
//		request.CompartmentId = &tmp
//	}
//
//	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
//		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
//		if err != nil {
//			return err
//		}
//		request.DefinedTags = convertedDefinedTags
//	}
//
//	if description, ok := s.D.GetOkExists("description"); ok {
//		tmp := description.(string)
//		request.Description = &tmp
//	}
//
//	if displayName, ok := s.D.GetOkExists("display_name"); ok {
//		tmp := displayName.(string)
//		request.DisplayName = &tmp
//	}
//
//	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
//		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
//	}
//
//	if runBooks, ok := s.D.GetOkExists("run_books"); ok {
//		interfaces := runBooks.([]interface{})
//		tmp := make([]oci_fleet_apps_management.OperationRunbook, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "run_books", stateDataIndex)
//			converted, err := s.mapToOperationRunbook(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange("run_books") {
//			request.RunBooks = tmp
//		}
//	}
//
//	if schedule, ok := s.D.GetOkExists("schedule"); ok {
//		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
//			tmp, err := s.mapToSchedule(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			request.Schedule = &tmp
//		}
//	}
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")
//
//	response, err := s.Client.CreateSchedulerDefinition(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	// CreateSchedulerDefinition returns an incorrect workRequest Id.
//	//workId := response.OpcWorkRequestId
//	var identifier *string
//	identifier = response.Id
//	if identifier != nil {
//		s.D.SetId(*identifier)
//	}
//	return s.Get()
//	//return s.getSchedulerDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) getSchedulerDefinitionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
//	actionTypeEnum oci_fleet_apps_management.ActionTypeEnum, timeout time.Duration) error {
//
//	// Wait until it finishes
//	schedulerDefinitionId, err := schedulerDefinitionWaitForWorkRequest(workId, "schedulerdefinition",
//		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.FleetClient)
//
//	if err != nil {
//		return err
//	}
//	s.D.SetId(*schedulerDefinitionId)
//
//	return s.Get()
//}
//
//func schedulerDefinitionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
//	startTime := time.Now()
//	stopTime := startTime.Add(timeout)
//	return func(response oci_common.OCIOperationResponse) bool {
//
//		// Stop after timeout has elapsed
//		if time.Now().After(stopTime) {
//			return false
//		}
//
//		// Make sure we stop on default rules
//		if tfresource.ShouldRetry(response, false, "fleet_apps_management", startTime) {
//			return true
//		}
//
//		// Only stop if the time Finished is set
//		if workRequestResponse, ok := response.Response.(oci_fleet_apps_management.GetWorkRequestResponse); ok {
//			return workRequestResponse.TimeFinished == nil
//		}
//		return false
//	}
//}
//
//func schedulerDefinitionWaitForWorkRequest(wId *string, entityType string, action oci_fleet_apps_management.ActionTypeEnum,
//	timeout time.Duration, disableFoundRetries bool, client *oci_fleet_apps_management.FleetAppsManagementClient) (*string, error) {
//	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "fleet_apps_management")
//	retryPolicy.ShouldRetryOperation = schedulerDefinitionWorkRequestShouldRetryFunc(timeout)
//
//	response := oci_fleet_apps_management.GetWorkRequestResponse{}
//	stateConf := &retry.StateChangeConf{
//		Pending: []string{
//			string(oci_fleet_apps_management.OperationStatusInProgress),
//			string(oci_fleet_apps_management.OperationStatusAccepted),
//			string(oci_fleet_apps_management.OperationStatusCanceling),
//		},
//		Target: []string{
//			string(oci_fleet_apps_management.OperationStatusSucceeded),
//			string(oci_fleet_apps_management.OperationStatusFailed),
//			string(oci_fleet_apps_management.OperationStatusCanceled),
//		},
//		Refresh: func() (interface{}, string, error) {
//			var err error
//			response, err = client.GetWorkRequest(context.Background(),
//				oci_fleet_apps_management.GetWorkRequestRequest{
//					WorkRequestId: wId,
//					RequestMetadata: oci_common.RequestMetadata{
//						RetryPolicy: retryPolicy,
//					},
//				})
//			wr := &response.WorkRequest
//			return wr, string(wr.Status), err
//		},
//		Timeout: timeout,
//	}
//	if _, e := stateConf.WaitForState(); e != nil {
//		return nil, e
//	}
//
//	var identifier *string
//	// The work request response contains an array of objects that finished the operation
//	for _, res := range response.Resources {
//		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
//			if res.ActionType == action {
//				identifier = res.Identifier
//				break
//			}
//		}
//	}
//
//	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
//	if identifier == nil || response.Status == oci_fleet_apps_management.OperationStatusFailed || response.Status == oci_fleet_apps_management.OperationStatusCanceled {
//		return nil, getErrorFromFleetAppsManagementSchedulerDefinitionWorkRequest(client, wId, retryPolicy, entityType, action)
//	}
//
//	return identifier, nil
//}
//
//func getErrorFromFleetAppsManagementSchedulerDefinitionWorkRequest(client *oci_fleet_apps_management.FleetAppsManagementClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_fleet_apps_management.ActionTypeEnum) error {
//	response, err := client.ListWorkRequestErrors(context.Background(),
//		oci_fleet_apps_management.ListWorkRequestErrorsRequest{
//			WorkRequestId: workId,
//			RequestMetadata: oci_common.RequestMetadata{
//				RetryPolicy: retryPolicy,
//			},
//		})
//	if err != nil {
//		return err
//	}
//
//	allErrs := make([]string, 0)
//	for _, wrkErr := range response.Items {
//		allErrs = append(allErrs, *wrkErr.Message)
//	}
//	errorMessage := strings.Join(allErrs, "\n")
//
//	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)
//
//	return workRequestErr
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Get() error {
//	request := oci_fleet_apps_management.GetSchedulerDefinitionRequest{}
//
//	tmp := s.D.Id()
//	request.SchedulerDefinitionId = &tmp
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")
//
//	response, err := s.Client.GetSchedulerDefinition(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	s.Res = &response.SchedulerDefinition
//	return nil
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Update() error {
//	request := oci_fleet_apps_management.UpdateSchedulerDefinitionRequest{}
//
//	if actionGroups, ok := s.D.GetOkExists("action_groups"); ok {
//		interfaces := actionGroups.([]interface{})
//		tmp := make([]oci_fleet_apps_management.ActionGroup, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "action_groups", stateDataIndex)
//			converted, err := s.mapToActionGroup(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange("action_groups") {
//			request.ActionGroups = tmp
//		}
//	}
//
//	if activityInitiationCutOff, ok := s.D.GetOkExists("activity_initiation_cut_off"); ok {
//		tmp := activityInitiationCutOff.(int)
//		request.ActivityInitiationCutOff = &tmp
//	}
//
//	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
//		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
//		if err != nil {
//			return err
//		}
//		request.DefinedTags = convertedDefinedTags
//	}
//
//	if description, ok := s.D.GetOkExists("description"); ok {
//		tmp := description.(string)
//		request.Description = &tmp
//	}
//
//	if displayName, ok := s.D.GetOkExists("display_name"); ok {
//		tmp := displayName.(string)
//		request.DisplayName = &tmp
//	}
//
//	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
//		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
//	}
//
//	if runBooks, ok := s.D.GetOkExists("run_books"); ok {
//		interfaces := runBooks.([]interface{})
//		tmp := make([]oci_fleet_apps_management.OperationRunbook, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "run_books", stateDataIndex)
//			converted, err := s.mapToOperationRunbook(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange("run_books") {
//			request.RunBooks = tmp
//		}
//	}
//
//	if schedule, ok := s.D.GetOkExists("schedule"); ok {
//		if tmpList := schedule.([]interface{}); len(tmpList) > 0 {
//			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule", 0)
//			tmp, err := s.mapToSchedule(fieldKeyFormat)
//			if err != nil {
//				return err
//			}
//			request.Schedule = &tmp
//		}
//	}
//
//	tmp := s.D.Id()
//	request.SchedulerDefinitionId = &tmp
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")
//
//	_, err := s.Client.UpdateSchedulerDefinition(context.Background(), request)
//	if err != nil {
//		return err
//	}
//
//	// Update returns incorrect workRequestId header
//	//workId := response.OpcWorkRequestId
//	return s.Get()
//	//return s.getSchedulerDefinitionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management"), oci_fleet_apps_management.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
//
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) Delete() error {
//	request := oci_fleet_apps_management.DeleteSchedulerDefinitionRequest{}
//
//	tmp := s.D.Id()
//	request.SchedulerDefinitionId = &tmp
//
//	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "fleet_apps_management")
//
//	_, err := s.Client.DeleteSchedulerDefinition(context.Background(), request)
//	return err
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) SetData() error {
//	s.D.Set("action_group_types", s.Res.ActionGroupTypes)
//
//	actionGroups := []interface{}{}
//	for _, item := range s.Res.ActionGroups {
//		actionGroups = append(actionGroups, ActionGroupToMap(item))
//	}
//	s.D.Set("action_groups", actionGroups)
//
//	if s.Res.ActivityInitiationCutOff != nil {
//		s.D.Set("activity_initiation_cut_off", *s.Res.ActivityInitiationCutOff)
//	}
//
//	s.D.Set("application_types", s.Res.ApplicationTypes)
//
//	if s.Res.CompartmentId != nil {
//		s.D.Set("compartment_id", *s.Res.CompartmentId)
//	} else {
//		s.D.Set("compartment_id", nil)
//	}
//
//	if s.Res.CountOfAffectedActionGroups != nil {
//		s.D.Set("count_of_affected_action_groups", *s.Res.CountOfAffectedActionGroups)
//	}
//
//	if s.Res.CountOfAffectedResources != nil {
//		s.D.Set("count_of_affected_resources", *s.Res.CountOfAffectedResources)
//	}
//
//	if s.Res.CountOfAffectedTargets != nil {
//		s.D.Set("count_of_affected_targets", *s.Res.CountOfAffectedTargets)
//	}
//
//	if s.Res.DefinedTags != nil {
//		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
//	}
//
//	if s.Res.Description != nil {
//		s.D.Set("description", *s.Res.Description)
//	}
//
//	if s.Res.DisplayName != nil {
//		s.D.Set("display_name", *s.Res.DisplayName)
//	}
//
//	s.D.Set("freeform_tags", s.Res.FreeformTags)
//
//	if s.Res.LifecycleDetails != nil {
//		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
//	}
//
//	s.D.Set("lifecycle_operations", s.Res.LifecycleOperations)
//
//	s.D.Set("products", s.Res.Products)
//
//	if s.Res.ResourceRegion != nil {
//		s.D.Set("resource_region", *s.Res.ResourceRegion)
//	}
//
//	runBooks := []interface{}{}
//	for _, item := range s.Res.RunBooks {
//		runBooks = append(runBooks, OperationRunbookToMap(item))
//	}
//	s.D.Set("run_books", runBooks)
//
//	if s.Res.Schedule != nil {
//		s.D.Set("schedule", []interface{}{ScheduleToMap(s.Res.Schedule)})
//	} else {
//		s.D.Set("schedule", nil)
//	}
//
//	s.D.Set("state", s.Res.LifecycleState)
//
//	if s.Res.SystemTags != nil {
//		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
//	}
//
//	if s.Res.TimeCreated != nil {
//		s.D.Set("time_created", s.Res.TimeCreated.String())
//	}
//
//	if s.Res.TimeOfNextRun != nil {
//		s.D.Set("time_of_next_run", s.Res.TimeOfNextRun.String())
//	}
//
//	if s.Res.TimeUpdated != nil {
//		s.D.Set("time_updated", s.Res.TimeUpdated.String())
//	}
//
//	return nil
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToActionGroup(fieldKeyFormat string) (oci_fleet_apps_management.ActionGroup, error) {
//	result := oci_fleet_apps_management.ActionGroup{}
//
//	if applicationType, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "application_type")); ok {
//		tmp := applicationType.(string)
//		result.ApplicationType = &tmp
//	}
//
//	if lifecycleOperation, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "lifecycle_operation")); ok {
//		tmp := lifecycleOperation.(string)
//		result.LifecycleOperation = &tmp
//	}
//
//	if product, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "product")); ok {
//		tmp := product.(string)
//		result.Product = &tmp
//	}
//
//	if resourceId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "resource_id")); ok {
//		tmp := resourceId.(string)
//		result.ResourceId = &tmp
//	}
//
//	if runbookId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "runbook_id")); ok {
//		tmp := runbookId.(string)
//		result.RunbookId = &tmp
//	}
//
//	if subjects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subjects")); ok {
//		interfaces := subjects.([]interface{})
//		tmp := make([]string, len(interfaces))
//		for i := range interfaces {
//			if interfaces[i] != nil {
//				tmp[i] = interfaces[i].(string)
//			}
//		}
//		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "subjects")) {
//			result.Subjects = tmp
//		}
//	}
//
//	if targetId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "target_id")); ok {
//		tmp := targetId.(string)
//		result.TargetId = &tmp
//	}
//
//	if type_, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "type")); ok {
//		result.Type = oci_fleet_apps_management.LifeCycleActionGroupTypeEnum(type_.(string))
//	}
//	log.Printf("[DEBUG] ActionGroup is %s", result)
//	return result, nil
//}
//
//func ActionGroupToMap(obj oci_fleet_apps_management.ActionGroup) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.ApplicationType != nil {
//		result["application_type"] = string(*obj.ApplicationType)
//	}
//
//	if obj.LifecycleOperation != nil {
//		result["lifecycle_operation"] = string(*obj.LifecycleOperation)
//	}
//
//	if obj.Product != nil {
//		result["product"] = string(*obj.Product)
//	}
//
//	if obj.ResourceId != nil {
//		result["resource_id"] = string(*obj.ResourceId)
//	}
//
//	if obj.RunbookId != nil {
//		result["runbook_id"] = string(*obj.RunbookId)
//	}
//
//	result["subjects"] = obj.Subjects
//
//	if obj.TargetId != nil {
//		result["target_id"] = string(*obj.TargetId)
//	}
//
//	result["type"] = string(obj.Type)
//
//	return result
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToInputParameter(fieldKeyFormat string) (oci_fleet_apps_management.InputParameter, error) {
//	result := oci_fleet_apps_management.InputParameter{}
//
//	if arguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "arguments")); ok {
//		interfaces := arguments.([]interface{})
//		tmp := make([]oci_fleet_apps_management.TaskArgument, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "arguments"), stateDataIndex)
//			converted, err := s.mapToTaskArgument(fieldKeyFormatNextLevel)
//			if err != nil {
//				return result, err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "arguments")) {
//			result.Arguments = tmp
//		}
//	}
//
//	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
//		tmp := stepName.(string)
//		result.StepName = &tmp
//	}
//
//	return result, nil
//}
//
//func InputParameterToMap(obj oci_fleet_apps_management.InputParameter) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	arguments := []interface{}{}
//	for _, item := range obj.Arguments {
//		arguments = append(arguments, TaskArgumentToMap(item))
//	}
//	result["arguments"] = arguments
//
//	if obj.StepName != nil {
//		result["step_name"] = string(*obj.StepName)
//	}
//
//	return result
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToOperationRunbook(fieldKeyFormat string) (oci_fleet_apps_management.OperationRunbook, error) {
//	result := oci_fleet_apps_management.OperationRunbook{}
//
//	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
//		tmp := id.(string)
//		result.Id = &tmp
//	}
//
//	if inputParameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "input_parameters")); ok {
//		interfaces := inputParameters.([]interface{})
//		tmp := make([]oci_fleet_apps_management.InputParameter, len(interfaces))
//		for i := range interfaces {
//			stateDataIndex := i
//			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "input_parameters"), stateDataIndex)
//			converted, err := s.mapToInputParameter(fieldKeyFormatNextLevel)
//			if err != nil {
//				return result, err
//			}
//			tmp[i] = converted
//		}
//		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "input_parameters")) {
//			result.InputParameters = tmp
//		}
//	}
//
//	return result, nil
//}
//
//func OperationRunbookToMap(obj oci_fleet_apps_management.OperationRunbook) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.Id != nil {
//		result["id"] = string(*obj.Id)
//	}
//
//	inputParameters := []interface{}{}
//	for _, item := range obj.InputParameters {
//		inputParameters = append(inputParameters, InputParameterToMap(item))
//	}
//	result["input_parameters"] = inputParameters
//
//	return result
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToSchedule(fieldKeyFormat string) (oci_fleet_apps_management.Schedule, error) {
//	result := oci_fleet_apps_management.Schedule{}
//
//	if duration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "duration")); ok {
//		tmp := duration.(string)
//		result.Duration = &tmp
//	}
//
//	if executionStartdate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "execution_startdate")); ok {
//		tmp, err := time.Parse(time.RFC3339, executionStartdate.(string))
//		if err != nil {
//			return result, err
//		}
//		result.ExecutionStartdate = &oci_common.SDKTime{Time: tmp}
//	}
//
//	if maintenanceWindowId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maintenance_window_id")); ok {
//		tmp := maintenanceWindowId.(string)
//		result.MaintenanceWindowId = &tmp
//	}
//
//	if recurrences, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recurrences")); ok {
//		tmp := recurrences.(string)
//		result.Recurrences = &tmp
//	}
//
//	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
//		result.Type = oci_fleet_apps_management.ScheduleTypeEnum(type_.(string))
//	}
//
//	return result, nil
//}
//
//func ScheduleToMap(obj *oci_fleet_apps_management.Schedule) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.Duration != nil {
//		result["duration"] = string(*obj.Duration)
//	}
//
//	if obj.ExecutionStartdate != nil {
//		result["execution_startdate"] = obj.ExecutionStartdate.Format(time.RFC3339Nano)
//	}
//
//	if obj.MaintenanceWindowId != nil {
//		result["maintenance_window_id"] = string(*obj.MaintenanceWindowId)
//	}
//
//	if obj.Recurrences != nil {
//		result["recurrences"] = string(*obj.Recurrences)
//	}
//
//	result["type"] = string(obj.Type)
//
//	return result
//}
//
//func SchedulerDefinitionSummaryToMap(obj oci_fleet_apps_management.SchedulerDefinitionSummary) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	result["action_group_types"] = obj.ActionGroupTypes
//
//	result["application_types"] = obj.ApplicationTypes
//
//	if obj.CompartmentId != nil {
//		result["compartment_id"] = string(*obj.CompartmentId)
//	}
//
//	if obj.CountOfAffectedActionGroups != nil {
//		result["count_of_affected_action_groups"] = int(*obj.CountOfAffectedActionGroups)
//	}
//
//	if obj.CountOfAffectedResources != nil {
//		result["count_of_affected_resources"] = int(*obj.CountOfAffectedResources)
//	}
//
//	if obj.CountOfAffectedTargets != nil {
//		result["count_of_affected_targets"] = int(*obj.CountOfAffectedTargets)
//	}
//
//	if obj.DefinedTags != nil {
//		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
//	}
//
//	if obj.Description != nil {
//		result["description"] = string(*obj.Description)
//	}
//
//	if obj.DisplayName != nil {
//		result["display_name"] = string(*obj.DisplayName)
//	}
//
//	result["freeform_tags"] = obj.FreeformTags
//
//	if obj.Id != nil {
//		result["id"] = string(*obj.Id)
//	}
//
//	if obj.LifecycleDetails != nil {
//		result["lifecycle_details"] = string(*obj.LifecycleDetails)
//	}
//
//	result["lifecycle_operations"] = obj.LifecycleOperations
//
//	result["products"] = obj.Products
//
//	if obj.ResourceRegion != nil {
//		result["resource_region"] = string(*obj.ResourceRegion)
//	}
//
//	if obj.Schedule != nil {
//		result["schedule"] = []interface{}{ScheduleToMap(obj.Schedule)}
//	}
//
//	result["state"] = string(obj.LifecycleState)
//
//	if obj.SystemTags != nil {
//		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
//	}
//
//	if obj.TimeCreated != nil {
//		result["time_created"] = obj.TimeCreated.String()
//	}
//
//	if obj.TimeOfNextRun != nil {
//		result["time_of_next_run"] = obj.TimeOfNextRun.String()
//	}
//
//	if obj.TimeUpdated != nil {
//		result["time_updated"] = obj.TimeUpdated.String()
//	}
//
//	return result
//}
//
//func (s *FleetAppsManagementSchedulerDefinitionResourceCrud) mapToTaskArgument(fieldKeyFormat string) (oci_fleet_apps_management.TaskArgument, error) {
//	result := oci_fleet_apps_management.TaskArgument{}
//
//	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
//		tmp := name.(string)
//		result.Name = &tmp
//	}
//
//	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
//		tmp := value.(string)
//		result.Value = &tmp
//	}
//
//	return result, nil
//}
//
//func TaskArgumentToMap(obj oci_fleet_apps_management.TaskArgument) map[string]interface{} {
//	result := map[string]interface{}{}
//
//	if obj.Name != nil {
//		result["name"] = string(*obj.Name)
//	}
//
//	if obj.Value != nil {
//		result["value"] = string(*obj.Value)
//	}
//
//	return result
//}
//>>>>>>> theirs
