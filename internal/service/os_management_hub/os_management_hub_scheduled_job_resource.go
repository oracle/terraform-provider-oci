// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package os_management_hub

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OsManagementHubScheduledJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createOsManagementHubScheduledJob,
		Read:     readOsManagementHubScheduledJob,
		Update:   updateOsManagementHubScheduledJob,
		Delete:   deleteOsManagementHubScheduledJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"operations": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"operation_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"manage_module_streams_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"disable": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"module_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"stream_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"software_source_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"enable": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"module_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"stream_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"software_source_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"install": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"module_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"profile_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"stream_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"software_source_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"remove": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"module_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"profile_name": {
													Type:     schema.TypeString,
													Required: true,
												},
												"stream_name": {
													Type:     schema.TypeString,
													Required: true,
												},

												// Optional
												"software_source_id": {
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
						"package_names": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"software_source_ids": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"switch_module_streams_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"module_name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"stream_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"software_source_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"windows_update_names": {
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
			"schedule_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"time_next_execution": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
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
			"is_managed_by_autonomous_linux": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_subcompartment_included": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"lifecycle_stage_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"locations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_compartment_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_instance_group_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"managed_instance_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"recurring_rule": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retry_intervals": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			// Computed
			"is_restricted": {
				Type:     schema.TypeBool,
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
			"time_last_execution": {
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

func createOsManagementHubScheduledJob(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()

	return tfresource.CreateResource(d, sync)
}

func readOsManagementHubScheduledJob(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()

	return tfresource.ReadResource(sync)
}

func updateOsManagementHubScheduledJob(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteOsManagementHubScheduledJob(d *schema.ResourceData, m interface{}) error {
	sync := &OsManagementHubScheduledJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ScheduledJobClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type OsManagementHubScheduledJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_os_management_hub.ScheduledJobClient
	Res                    *oci_os_management_hub.ScheduledJob
	DisableNotFoundRetries bool
}

func (s *OsManagementHubScheduledJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *OsManagementHubScheduledJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_os_management_hub.ScheduledJobLifecycleStateCreating),
	}
}

func (s *OsManagementHubScheduledJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_os_management_hub.ScheduledJobLifecycleStateActive),
	}
}

func (s *OsManagementHubScheduledJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_os_management_hub.ScheduledJobLifecycleStateDeleting),
	}
}

func (s *OsManagementHubScheduledJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_os_management_hub.ScheduledJobLifecycleStateDeleted),
	}
}

func (s *OsManagementHubScheduledJobResourceCrud) Create() error {
	request := oci_os_management_hub.CreateScheduledJobRequest{}

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

	if isManagedByAutonomousLinux, ok := s.D.GetOkExists("is_managed_by_autonomous_linux"); ok {
		tmp := isManagedByAutonomousLinux.(bool)
		request.IsManagedByAutonomousLinux = &tmp
	}

	if isSubcompartmentIncluded, ok := s.D.GetOkExists("is_subcompartment_included"); ok {
		tmp := isSubcompartmentIncluded.(bool)
		request.IsSubcompartmentIncluded = &tmp
	}

	if lifecycleStageIds, ok := s.D.GetOkExists("lifecycle_stage_ids"); ok {
		interfaces := lifecycleStageIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("lifecycle_stage_ids") {
			request.LifecycleStageIds = tmp
		}
	}

	if locations, ok := s.D.GetOkExists("locations"); ok {
		interfaces := locations.([]interface{})
		tmp := make([]oci_os_management_hub.ManagedInstanceLocationEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_os_management_hub.ManagedInstanceLocationEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("locations") {
			request.Locations = tmp
		}
	}

	if managedCompartmentIds, ok := s.D.GetOkExists("managed_compartment_ids"); ok {
		interfaces := managedCompartmentIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_compartment_ids") {
			request.ManagedCompartmentIds = tmp
		}
	}

	if managedInstanceGroupIds, ok := s.D.GetOkExists("managed_instance_group_ids"); ok {
		interfaces := managedInstanceGroupIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_instance_group_ids") {
			request.ManagedInstanceGroupIds = tmp
		}
	}

	if managedInstanceIds, ok := s.D.GetOkExists("managed_instance_ids"); ok {
		interfaces := managedInstanceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("managed_instance_ids") {
			request.ManagedInstanceIds = tmp
		}
	}

	if operations, ok := s.D.GetOkExists("operations"); ok {
		interfaces := operations.([]interface{})
		tmp := make([]oci_os_management_hub.ScheduledJobOperation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "operations", stateDataIndex)
			converted, err := s.mapToScheduledJobOperation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("operations") {
			request.Operations = tmp
		}
	}

	if recurringRule, ok := s.D.GetOkExists("recurring_rule"); ok {
		tmp := recurringRule.(string)
		request.RecurringRule = &tmp
	}

	if retryIntervals, ok := s.D.GetOkExists("retry_intervals"); ok {
		interfaces := retryIntervals.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("retry_intervals") {
			request.RetryIntervals = tmp
		}
	}

	if scheduleType, ok := s.D.GetOkExists("schedule_type"); ok {
		request.ScheduleType = oci_os_management_hub.ScheduleTypesEnum(scheduleType.(string))
	}

	if timeNextExecution, ok := s.D.GetOkExists("time_next_execution"); ok {
		tmp, err := time.Parse(time.RFC3339, timeNextExecution.(string))
		if err != nil {
			return err
		}
		request.TimeNextExecution = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.CreateScheduledJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledJob
	return nil
}

func (s *OsManagementHubScheduledJobResourceCrud) Get() error {
	request := oci_os_management_hub.GetScheduledJobRequest{}

	tmp := s.D.Id()
	request.ScheduledJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.GetScheduledJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledJob
	return nil
}

func (s *OsManagementHubScheduledJobResourceCrud) Update() error {

	if _, ok := s.D.GetOkExists("compartmentId"); ok && s.D.HasChange("compartmentId") {
		err := s.ChangeScheduledJobCompartment()
		if err != nil {
			return err
		}
	}
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_os_management_hub.UpdateScheduledJobRequest{}

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

	if operations, ok := s.D.GetOkExists("operations"); ok {
		interfaces := operations.([]interface{})
		tmp := make([]oci_os_management_hub.ScheduledJobOperation, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "operations", stateDataIndex)
			converted, err := s.mapToScheduledJobOperation(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("operations") {
			request.Operations = tmp
		}
	}

	if recurringRule, ok := s.D.GetOkExists("recurring_rule"); ok {
		tmp := recurringRule.(string)
		request.RecurringRule = &tmp
	}

	if retryIntervals, ok := s.D.GetOkExists("retry_intervals"); ok {
		interfaces := retryIntervals.([]interface{})
		tmp := make([]int, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(int)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("retry_intervals") {
			request.RetryIntervals = tmp
		}
	}

	if scheduleType, ok := s.D.GetOkExists("schedule_type"); ok {
		request.ScheduleType = oci_os_management_hub.ScheduleTypesEnum(scheduleType.(string))
	}

	tmp := s.D.Id()
	request.ScheduledJobId = &tmp

	if timeNextExecution, ok := s.D.GetOkExists("time_next_execution"); ok {
		tmp, err := time.Parse(time.RFC3339, timeNextExecution.(string))
		if err != nil {
			return err
		}
		request.TimeNextExecution = &oci_common.SDKTime{Time: tmp}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	response, err := s.Client.UpdateScheduledJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.ScheduledJob
	return nil
}

func (s *OsManagementHubScheduledJobResourceCrud) Delete() error {
	request := oci_os_management_hub.DeleteScheduledJobRequest{}

	tmp := s.D.Id()
	request.ScheduledJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.DeleteScheduledJob(context.Background(), request)
	return err
}

func (s *OsManagementHubScheduledJobResourceCrud) SetData() error {
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

	if s.Res.IsManagedByAutonomousLinux != nil {
		s.D.Set("is_managed_by_autonomous_linux", *s.Res.IsManagedByAutonomousLinux)
	}

	if s.Res.IsRestricted != nil {
		s.D.Set("is_restricted", *s.Res.IsRestricted)
	}

	if s.Res.IsSubcompartmentIncluded != nil {
		s.D.Set("is_subcompartment_included", *s.Res.IsSubcompartmentIncluded)
	}

	s.D.Set("lifecycle_stage_ids", s.Res.LifecycleStageIds)

	s.D.Set("locations", s.Res.Locations)

	s.D.Set("managed_compartment_ids", s.Res.ManagedCompartmentIds)

	s.D.Set("managed_instance_group_ids", s.Res.ManagedInstanceGroupIds)

	s.D.Set("managed_instance_ids", s.Res.ManagedInstanceIds)

	operations := []interface{}{}
	for _, item := range s.Res.Operations {
		operations = append(operations, ScheduledJobOperationToMap(item))
	}
	s.D.Set("operations", operations)

	if s.Res.RecurringRule != nil {
		s.D.Set("recurring_rule", *s.Res.RecurringRule)
	}

	s.D.Set("retry_intervals", s.Res.RetryIntervals)

	s.D.Set("schedule_type", s.Res.ScheduleType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastExecution != nil {
		s.D.Set("time_last_execution", s.Res.TimeLastExecution.String())
	}

	if s.Res.TimeNextExecution != nil {
		s.D.Set("time_next_execution", s.Res.TimeNextExecution.Format(time.RFC3339Nano))
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("work_request_ids", s.Res.WorkRequestIds)

	return nil
}

func (s *OsManagementHubScheduledJobResourceCrud) ChangeScheduledJobCompartment() error {
	request := oci_os_management_hub.ChangeScheduledJobCompartmentRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	idTmp := s.D.Id()
	request.ScheduledJobId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeScheduledJobCompartment(context.Background(), request)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *OsManagementHubScheduledJobResourceCrud) mapToManageModuleStreamsInScheduledJobDetails(fieldKeyFormat string) (oci_os_management_hub.ManageModuleStreamsInScheduledJobDetails, error) {
	result := oci_os_management_hub.ManageModuleStreamsInScheduledJobDetails{}

	if disable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "disable")); ok {
		interfaces := disable.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "disable"), stateDataIndex)
			converted, err := s.mapToModuleStreamDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "disable")) {
			result.Disable = tmp
		}
	}

	if enable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable")); ok {
		interfaces := enable.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "enable"), stateDataIndex)
			converted, err := s.mapToModuleStreamDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "enable")) {
			result.Enable = tmp
		}
	}

	if install, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "install")); ok {
		interfaces := install.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamProfileDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "install"), stateDataIndex)
			converted, err := s.mapToModuleStreamProfileDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "install")) {
			result.Install = tmp
		}
	}

	if remove, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remove")); ok {
		interfaces := remove.([]interface{})
		tmp := make([]oci_os_management_hub.ModuleStreamProfileDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "remove"), stateDataIndex)
			converted, err := s.mapToModuleStreamProfileDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "remove")) {
			result.Remove = tmp
		}
	}

	return result, nil
}

func ManageModuleStreamsInScheduledJobDetailsToMap(obj *oci_os_management_hub.ManageModuleStreamsInScheduledJobDetails) map[string]interface{} {
	result := map[string]interface{}{}

	disable := []interface{}{}
	for _, item := range obj.Disable {
		disable = append(disable, ModuleStreamDetailsToMap(item))
	}
	result["disable"] = disable

	enable := []interface{}{}
	for _, item := range obj.Enable {
		enable = append(enable, ModuleStreamDetailsToMap(item))
	}
	result["enable"] = enable

	install := []interface{}{}
	for _, item := range obj.Install {
		install = append(install, ModuleStreamProfileDetailsToMap(item))
	}
	result["install"] = install

	remove := []interface{}{}
	for _, item := range obj.Remove {
		remove = append(remove, ModuleStreamProfileDetailsToMap(item))
	}
	result["remove"] = remove

	return result
}

func (s *OsManagementHubScheduledJobResourceCrud) mapToModuleStreamDetails(fieldKeyFormat string) (oci_os_management_hub.ModuleStreamDetails, error) {
	result := oci_os_management_hub.ModuleStreamDetails{}

	if moduleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_name")); ok {
		tmp := moduleName.(string)
		result.ModuleName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_id")); ok {
		tmp := softwareSourceId.(string)
		result.SoftwareSourceId = &tmp
	}

	if streamName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_name")); ok {
		tmp := streamName.(string)
		result.StreamName = &tmp
	}

	return result, nil
}

func (s *OsManagementHubScheduledJobResourceCrud) mapToModuleStreamProfileDetails(fieldKeyFormat string) (oci_os_management_hub.ModuleStreamProfileDetails, error) {
	result := oci_os_management_hub.ModuleStreamProfileDetails{}

	if moduleName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "module_name")); ok {
		tmp := moduleName.(string)
		result.ModuleName = &tmp
	}

	if profileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "profile_name")); ok {
		tmp := profileName.(string)
		result.ProfileName = &tmp
	}

	if softwareSourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_id")); ok {
		tmp := softwareSourceId.(string)
		result.SoftwareSourceId = &tmp
	}

	if streamName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "stream_name")); ok {
		tmp := streamName.(string)
		result.StreamName = &tmp
	}

	return result, nil
}

func (s *OsManagementHubScheduledJobResourceCrud) mapToScheduledJobOperation(fieldKeyFormat string) (oci_os_management_hub.ScheduledJobOperation, error) {
	result := oci_os_management_hub.ScheduledJobOperation{}

	if manageModuleStreamsDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "manage_module_streams_details")); ok {
		if tmpList := manageModuleStreamsDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "manage_module_streams_details"), 0)
			tmp, err := s.mapToManageModuleStreamsInScheduledJobDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert manage_module_streams_details, encountered error: %v", err)
			}
			result.ManageModuleStreamsDetails = &tmp
			if result.ManageModuleStreamsDetails.Disable == nil && result.ManageModuleStreamsDetails.Enable == nil && result.ManageModuleStreamsDetails.Install == nil && result.ManageModuleStreamsDetails.Remove == nil {
				result.ManageModuleStreamsDetails = nil
			}
		}
	}

	if operationType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "operation_type")); ok {
		result.OperationType = oci_os_management_hub.OperationTypesEnum(operationType.(string))
	}

	if packageNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "package_names")); ok {
		interfaces := packageNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "package_names")) {
			result.PackageNames = tmp
		}
	}

	if softwareSourceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "software_source_ids")); ok {
		interfaces := softwareSourceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "software_source_ids")) {
			result.SoftwareSourceIds = tmp
		}
	}

	if switchModuleStreamsDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "switch_module_streams_details")); ok {
		if tmpList := switchModuleStreamsDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "switch_module_streams_details"), 0)
			tmp, err := s.mapToModuleStreamDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert switch_module_streams_details, encountered error: %v", err)
			}
			result.SwitchModuleStreamsDetails = &tmp
			if result.SwitchModuleStreamsDetails.ModuleName == nil && result.SwitchModuleStreamsDetails.StreamName == nil {
				result.SwitchModuleStreamsDetails = nil
			}
		}
	}

	if windowsUpdateNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "windows_update_names")); ok {
		interfaces := windowsUpdateNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "windows_update_names")) {
			result.WindowsUpdateNames = tmp
		}
	}

	return result, nil
}

func ScheduledJobOperationToMap(obj oci_os_management_hub.ScheduledJobOperation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ManageModuleStreamsDetails != nil {
		result["manage_module_streams_details"] = []interface{}{ManageModuleStreamsInScheduledJobDetailsToMap(obj.ManageModuleStreamsDetails)}
	}

	result["operation_type"] = string(obj.OperationType)

	result["package_names"] = obj.PackageNames

	result["software_source_ids"] = obj.SoftwareSourceIds

	if obj.SwitchModuleStreamsDetails != nil {
		result["switch_module_streams_details"] = []interface{}{ModuleStreamDetailsToMap(*obj.SwitchModuleStreamsDetails)}
	}

	result["windows_update_names"] = obj.WindowsUpdateNames

	return result
}

func ScheduledJobSummaryToMap(obj oci_os_management_hub.ScheduledJobSummary) map[string]interface{} {
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

	if obj.IsManagedByAutonomousLinux != nil {
		result["is_managed_by_autonomous_linux"] = bool(*obj.IsManagedByAutonomousLinux)
	}

	if obj.IsRestricted != nil {
		result["is_restricted"] = bool(*obj.IsRestricted)
	}

	result["lifecycle_stage_ids"] = obj.LifecycleStageIds

	result["locations"] = obj.Locations

	result["managed_compartment_ids"] = obj.ManagedCompartmentIds

	result["managed_instance_group_ids"] = obj.ManagedInstanceGroupIds

	result["managed_instance_ids"] = obj.ManagedInstanceIds

	operations := []interface{}{}
	for _, item := range obj.Operations {
		operations = append(operations, ScheduledJobOperationToMap(item))
	}
	result["operations"] = operations

	result["retry_intervals"] = obj.RetryIntervals

	result["schedule_type"] = string(obj.ScheduleType)

	result["state"] = string(obj.LifecycleState)

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastExecution != nil {
		result["time_last_execution"] = obj.TimeLastExecution.String()
	}

	if obj.TimeNextExecution != nil {
		result["time_next_execution"] = obj.TimeNextExecution.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}

func (s *OsManagementHubScheduledJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_os_management_hub.ChangeScheduledJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ScheduledJobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "os_management_hub")

	_, err := s.Client.ChangeScheduledJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
