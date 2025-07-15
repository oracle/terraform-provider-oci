// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package jms

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_jms "github.com/oracle/oci-go-sdk/v65/jms"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func JmsTaskScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createJmsTaskSchedule,
		Read:     readJmsTaskSchedule,
		Update:   updateJmsTaskSchedule,
		Delete:   deleteJmsTaskSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"execution_recurrences": {
				Type:     schema.TypeString,
				Required: true,
			},
			"fleet_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"task_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ADD_INSTALLATION_SITE",
								"CRYPTO",
								"DEPLOYED_APPLICATION_MIGRATION",
								"JAVA_MIGRATION",
								"JFR",
								"PERFORMANCE_TUNING",
								"REMOVE_INSTALLATION_SITE",
								"SCAN_JAVA_SERVER",
								"SCAN_LIBRARY",
							}, true),
						},

						// Optional
						"add_installation_site_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"installation_sites": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"artifact_content_type": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"force_install": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"headless_mode": {
													Type:     schema.TypeBool,
													Optional: true,
													Computed: true,
												},
												"installation_path": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"release_version": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"post_installation_actions": {
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
						"crypto_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"recording_duration_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"application_installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"application_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"container_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"jre_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"waiting_period_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"deployed_application_migration_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"deployed_application_installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"exclude_package_prefixes": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"include_package_prefixes": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"source_jdk_version": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"target_jdk_version": {
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
						"java_migration_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"application_installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"exclude_package_prefixes": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"include_package_prefixes": {
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"source_jdk_version": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"target_jdk_version": {
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
						"jfr_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"jfc_profile_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"jfc_v1": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"jfc_v2": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"recording_duration_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"recording_size_in_mb": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"application_installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"application_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"container_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"jre_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"waiting_period_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"performance_tuning_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"recording_duration_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"targets": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"application_installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"application_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"container_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"jre_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"managed_instance_id": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"waiting_period_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"remove_installation_site_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"installation_sites": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"installation_key": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"managed_instance_id": {
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
						"scan_java_server_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"managed_instance_ids": {
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
						"scan_library_task_request": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"dynamic_scan_duration_in_minutes": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"is_dynamic_scan": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"managed_instance_ids": {
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

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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
			"time_last_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_next_run": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createJmsTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.CreateResource(d, sync)
}

func readJmsTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.ReadResource(sync)
}

func updateJmsTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteJmsTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &JmsTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).JavaManagementServiceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type JmsTaskScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_jms.JavaManagementServiceClient
	Res                    *oci_jms.TaskSchedule
	DisableNotFoundRetries bool
}

func (s *JmsTaskScheduleResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *JmsTaskScheduleResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *JmsTaskScheduleResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_jms.TaskScheduleLifecycleStateActive),
	}
}

func (s *JmsTaskScheduleResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *JmsTaskScheduleResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *JmsTaskScheduleResourceCrud) Create() error {
	request := oci_jms.CreateTaskScheduleRequest{}

	if executionRecurrences, ok := s.D.GetOkExists("execution_recurrences"); ok {
		tmp := executionRecurrences.(string)
		request.ExecutionRecurrences = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if taskDetails, ok := s.D.GetOkExists("task_details"); ok {
		if tmpList := taskDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "task_details", 0)
			tmp, err := s.mapToTaskDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TaskDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.CreateTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *JmsTaskScheduleResourceCrud) Get() error {
	request := oci_jms.GetTaskScheduleRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.TaskScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.GetTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *JmsTaskScheduleResourceCrud) Update() error {
	request := oci_jms.UpdateTaskScheduleRequest{}

	if executionRecurrences, ok := s.D.GetOkExists("execution_recurrences"); ok {
		tmp := executionRecurrences.(string)
		request.ExecutionRecurrences = &tmp
	}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	if taskDetails, ok := s.D.GetOkExists("task_details"); ok {
		if tmpList := taskDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "task_details", 0)
			tmp, err := s.mapToTaskDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TaskDetails = tmp
		}
	}

	tmp := s.D.Id()
	request.TaskScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	response, err := s.Client.UpdateTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *JmsTaskScheduleResourceCrud) Delete() error {
	request := oci_jms.DeleteTaskScheduleRequest{}

	if fleetId, ok := s.D.GetOkExists("fleet_id"); ok {
		tmp := fleetId.(string)
		request.FleetId = &tmp
	}

	tmp := s.D.Id()
	request.TaskScheduleId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "jms")

	_, err := s.Client.DeleteTaskSchedule(context.Background(), request)
	return err
}

func (s *JmsTaskScheduleResourceCrud) SetData() error {
	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.ExecutionRecurrences != nil {
		s.D.Set("execution_recurrences", *s.Res.ExecutionRecurrences)
	}

	if s.Res.FleetId != nil {
		s.D.Set("fleet_id", *s.Res.FleetId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := TaskDetailsToMap(&s.Res.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		s.D.Set("task_details", taskDetailsArray)
	} else {
		s.D.Set("task_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRun != nil {
		s.D.Set("time_last_run", s.Res.TimeLastRun.String())
	}

	if s.Res.TimeLastUpdated != nil {
		s.D.Set("time_last_updated", s.Res.TimeLastUpdated.String())
	}

	if s.Res.TimeNextRun != nil {
		s.D.Set("time_next_run", s.Res.TimeNextRun.String())
	}

	return nil
}

func (s *JmsTaskScheduleResourceCrud) mapToAddFleetInstallationSitesDetails(fieldKeyFormat string) (oci_jms.AddFleetInstallationSitesDetails, error) {
	result := oci_jms.AddFleetInstallationSitesDetails{}

	if installationSites, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "installation_sites")); ok {
		interfaces := installationSites.([]interface{})
		tmp := make([]oci_jms.NewInstallationSite, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "installation_sites"), stateDataIndex)
			converted, err := s.mapToNewInstallationSite(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "installation_sites")) {
			result.InstallationSites = tmp
		}
	}

	if postInstallationActions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "post_installation_actions")); ok {
		interfaces := postInstallationActions.([]interface{})
		tmp := make([]oci_jms.PostInstallationActionsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_jms.PostInstallationActionsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "post_installation_actions")) {
			result.PostInstallationActions = tmp
		}
	}

	return result, nil
}

func AddFleetInstallationSitesDetailsToMap(obj *oci_jms.AddFleetInstallationSitesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	installationSites := []interface{}{}
	for _, item := range obj.InstallationSites {
		installationSites = append(installationSites, NewInstallationSiteToMap(item))
	}
	result["installation_sites"] = installationSites

	result["post_installation_actions"] = obj.PostInstallationActions

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToDeployedApplicationMigrationAnalysesTarget(fieldKeyFormat string) (oci_jms.DeployedApplicationMigrationAnalysesTarget, error) {
	result := oci_jms.DeployedApplicationMigrationAnalysesTarget{}

	if deployedApplicationInstallationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployed_application_installation_key")); ok {
		tmp := deployedApplicationInstallationKey.(string)
		result.DeployedApplicationInstallationKey = &tmp
	}

	if excludePackagePrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_package_prefixes")); ok {
		interfaces := excludePackagePrefixes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_package_prefixes")) {
			result.ExcludePackagePrefixes = tmp
		}
	}

	if includePackagePrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_package_prefixes")); ok {
		interfaces := includePackagePrefixes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "include_package_prefixes")) {
			result.IncludePackagePrefixes = tmp
		}
	}

	if managedInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_id")); ok {
		tmp := managedInstanceId.(string)
		result.ManagedInstanceId = &tmp
	}

	if sourceJdkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_jdk_version")); ok {
		tmp := sourceJdkVersion.(string)
		result.SourceJdkVersion = &tmp
	}

	if targetJdkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_jdk_version")); ok {
		tmp := targetJdkVersion.(string)
		result.TargetJdkVersion = &tmp
	}

	return result, nil
}

func DeployedApplicationMigrationAnalysesTargetToMap(obj oci_jms.DeployedApplicationMigrationAnalysesTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeployedApplicationInstallationKey != nil {
		result["deployed_application_installation_key"] = string(*obj.DeployedApplicationInstallationKey)
	}

	result["exclude_package_prefixes"] = obj.ExcludePackagePrefixes

	result["include_package_prefixes"] = obj.IncludePackagePrefixes

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.SourceJdkVersion != nil {
		result["source_jdk_version"] = string(*obj.SourceJdkVersion)
	}

	if obj.TargetJdkVersion != nil {
		result["target_jdk_version"] = string(*obj.TargetJdkVersion)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToExistingInstallationSiteId(fieldKeyFormat string) (oci_jms.ExistingInstallationSiteId, error) {
	result := oci_jms.ExistingInstallationSiteId{}

	if installationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "installation_key")); ok {
		tmp := installationKey.(string)
		result.InstallationKey = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_id")); ok {
		tmp := managedInstanceId.(string)
		result.ManagedInstanceId = &tmp
	}

	return result, nil
}

func ExistingInstallationSiteIdToMap(obj oci_jms.ExistingInstallationSiteId) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.InstallationKey != nil {
		result["installation_key"] = string(*obj.InstallationKey)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToJavaMigrationAnalysisTarget(fieldKeyFormat string) (oci_jms.JavaMigrationAnalysisTarget, error) {
	result := oci_jms.JavaMigrationAnalysisTarget{}

	if applicationInstallationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_installation_key")); ok {
		tmp := applicationInstallationKey.(string)
		result.ApplicationInstallationKey = &tmp
	}

	if excludePackagePrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude_package_prefixes")); ok {
		interfaces := excludePackagePrefixes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "exclude_package_prefixes")) {
			result.ExcludePackagePrefixes = tmp
		}
	}

	if includePackagePrefixes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_package_prefixes")); ok {
		interfaces := includePackagePrefixes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "include_package_prefixes")) {
			result.IncludePackagePrefixes = tmp
		}
	}

	if managedInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_id")); ok {
		tmp := managedInstanceId.(string)
		result.ManagedInstanceId = &tmp
	}

	if sourceJdkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "source_jdk_version")); ok {
		tmp := sourceJdkVersion.(string)
		result.SourceJdkVersion = &tmp
	}

	if targetJdkVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_jdk_version")); ok {
		tmp := targetJdkVersion.(string)
		result.TargetJdkVersion = &tmp
	}

	return result, nil
}

func JavaMigrationAnalysisTargetToMap(obj oci_jms.JavaMigrationAnalysisTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationInstallationKey != nil {
		result["application_installation_key"] = string(*obj.ApplicationInstallationKey)
	}

	result["exclude_package_prefixes"] = obj.ExcludePackagePrefixes

	result["include_package_prefixes"] = obj.IncludePackagePrefixes

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.SourceJdkVersion != nil {
		result["source_jdk_version"] = string(*obj.SourceJdkVersion)
	}

	if obj.TargetJdkVersion != nil {
		result["target_jdk_version"] = string(*obj.TargetJdkVersion)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToJfrAttachmentTarget(fieldKeyFormat string) (oci_jms.JfrAttachmentTarget, error) {
	result := oci_jms.JfrAttachmentTarget{}

	if applicationInstallationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_installation_key")); ok {
		tmp := applicationInstallationKey.(string)
		result.ApplicationInstallationKey = &tmp
	}

	if applicationKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "application_key")); ok {
		tmp := applicationKey.(string)
		result.ApplicationKey = &tmp
	}

	if containerKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_key")); ok {
		tmp := containerKey.(string)
		result.ContainerKey = &tmp
	}

	if jreKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jre_key")); ok {
		tmp := jreKey.(string)
		result.JreKey = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_id")); ok {
		tmp := managedInstanceId.(string)
		result.ManagedInstanceId = &tmp
	}

	return result, nil
}

func JfrAttachmentTargetToMap(obj oci_jms.JfrAttachmentTarget) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ApplicationInstallationKey != nil {
		result["application_installation_key"] = string(*obj.ApplicationInstallationKey)
	}

	if obj.ApplicationKey != nil {
		result["application_key"] = string(*obj.ApplicationKey)
	}

	if obj.ContainerKey != nil {
		result["container_key"] = string(*obj.ContainerKey)
	}

	if obj.JreKey != nil {
		result["jre_key"] = string(*obj.JreKey)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToNewInstallationSite(fieldKeyFormat string) (oci_jms.NewInstallationSite, error) {
	result := oci_jms.NewInstallationSite{}

	if artifactContentType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_content_type")); ok {
		result.ArtifactContentType = oci_jms.ArtifactContentTypeEnum(artifactContentType.(string))
	}

	if forceInstall, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "force_install")); ok {
		tmp := forceInstall.(bool)
		result.ForceInstall = &tmp
	}

	if headlessMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "headless_mode")); ok {
		tmp := headlessMode.(bool)
		result.HeadlessMode = &tmp
	}

	if installationPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "installation_path")); ok {
		tmp := installationPath.(string)
		result.InstallationPath = &tmp
	}

	if managedInstanceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_id")); ok {
		tmp := managedInstanceId.(string)
		result.ManagedInstanceId = &tmp
	}

	if releaseVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "release_version")); ok {
		tmp := releaseVersion.(string)
		result.ReleaseVersion = &tmp
	}

	return result, nil
}

func NewInstallationSiteToMap(obj oci_jms.NewInstallationSite) map[string]interface{} {
	result := map[string]interface{}{}

	result["artifact_content_type"] = string(obj.ArtifactContentType)

	if obj.ForceInstall != nil {
		result["force_install"] = bool(*obj.ForceInstall)
	}

	if obj.HeadlessMode != nil {
		result["headless_mode"] = bool(*obj.HeadlessMode)
	}

	if obj.InstallationPath != nil {
		result["installation_path"] = string(*obj.InstallationPath)
	}

	if obj.ManagedInstanceId != nil {
		result["managed_instance_id"] = string(*obj.ManagedInstanceId)
	}

	if obj.ReleaseVersion != nil {
		result["release_version"] = string(*obj.ReleaseVersion)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRemoveFleetInstallationSitesDetails(fieldKeyFormat string) (oci_jms.RemoveFleetInstallationSitesDetails, error) {
	result := oci_jms.RemoveFleetInstallationSitesDetails{}

	if installationSites, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "installation_sites")); ok {
		interfaces := installationSites.([]interface{})
		tmp := make([]oci_jms.ExistingInstallationSiteId, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "installation_sites"), stateDataIndex)
			converted, err := s.mapToExistingInstallationSiteId(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "installation_sites")) {
			result.InstallationSites = tmp
		}
	}

	return result, nil
}

func RemoveFleetInstallationSitesDetailsToMap(obj *oci_jms.RemoveFleetInstallationSitesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	installationSites := []interface{}{}
	for _, item := range obj.InstallationSites {
		installationSites = append(installationSites, ExistingInstallationSiteIdToMap(item))
	}
	result["installation_sites"] = installationSites

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRequestCryptoAnalysesDetails(fieldKeyFormat string) (oci_jms.RequestCryptoAnalysesDetails, error) {
	result := oci_jms.RequestCryptoAnalysesDetails{}

	if recordingDurationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recording_duration_in_minutes")); ok {
		tmp := recordingDurationInMinutes.(int)
		result.RecordingDurationInMinutes = &tmp
	}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_jms.JfrAttachmentTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToJfrAttachmentTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	if waitingPeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "waiting_period_in_minutes")); ok {
		tmp := waitingPeriodInMinutes.(int)
		result.WaitingPeriodInMinutes = &tmp
	}

	return result, nil
}

func RequestCryptoAnalysesDetailsToMap(obj *oci_jms.RequestCryptoAnalysesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RecordingDurationInMinutes != nil {
		result["recording_duration_in_minutes"] = int(*obj.RecordingDurationInMinutes)
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, JfrAttachmentTargetToMap(item))
	}
	result["targets"] = targets

	if obj.WaitingPeriodInMinutes != nil {
		result["waiting_period_in_minutes"] = int(*obj.WaitingPeriodInMinutes)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRequestDeployedApplicationMigrationAnalysesDetails(fieldKeyFormat string) (oci_jms.RequestDeployedApplicationMigrationAnalysesDetails, error) {
	result := oci_jms.RequestDeployedApplicationMigrationAnalysesDetails{}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_jms.DeployedApplicationMigrationAnalysesTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToDeployedApplicationMigrationAnalysesTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	return result, nil
}

func RequestDeployedApplicationMigrationAnalysesDetailsToMap(obj *oci_jms.RequestDeployedApplicationMigrationAnalysesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, DeployedApplicationMigrationAnalysesTargetToMap(item))
	}
	result["targets"] = targets

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRequestJavaMigrationAnalysesDetails(fieldKeyFormat string) (oci_jms.RequestJavaMigrationAnalysesDetails, error) {
	result := oci_jms.RequestJavaMigrationAnalysesDetails{}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_jms.JavaMigrationAnalysisTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToJavaMigrationAnalysisTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	return result, nil
}

func RequestJavaMigrationAnalysesDetailsToMap(obj *oci_jms.RequestJavaMigrationAnalysesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, JavaMigrationAnalysisTargetToMap(item))
	}
	result["targets"] = targets

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRequestJfrRecordingsDetails(fieldKeyFormat string) (oci_jms.RequestJfrRecordingsDetails, error) {
	result := oci_jms.RequestJfrRecordingsDetails{}

	if jfcProfileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jfc_profile_name")); ok {
		tmp := jfcProfileName.(string)
		result.JfcProfileName = &tmp
	}

	if jfcV1, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jfc_v1")); ok {
		tmp := jfcV1.(string)
		result.JfcV1 = &tmp
	}

	if jfcV2, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jfc_v2")); ok {
		tmp := jfcV2.(string)
		result.JfcV2 = &tmp
	}

	if recordingDurationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recording_duration_in_minutes")); ok {
		tmp := recordingDurationInMinutes.(int)
		result.RecordingDurationInMinutes = &tmp
	}

	if recordingSizeInMb, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recording_size_in_mb")); ok {
		tmp := recordingSizeInMb.(int)
		result.RecordingSizeInMb = &tmp
	}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_jms.JfrAttachmentTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToJfrAttachmentTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	if waitingPeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "waiting_period_in_minutes")); ok {
		tmp := waitingPeriodInMinutes.(int)
		result.WaitingPeriodInMinutes = &tmp
	}

	return result, nil
}

func RequestJfrRecordingsDetailsToMap(obj *oci_jms.RequestJfrRecordingsDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.JfcProfileName != nil {
		result["jfc_profile_name"] = string(*obj.JfcProfileName)
	}

	if obj.JfcV1 != nil {
		result["jfc_v1"] = string(*obj.JfcV1)
	}

	if obj.JfcV2 != nil {
		result["jfc_v2"] = string(*obj.JfcV2)
	}

	if obj.RecordingDurationInMinutes != nil {
		result["recording_duration_in_minutes"] = int(*obj.RecordingDurationInMinutes)
	}

	if obj.RecordingSizeInMb != nil {
		result["recording_size_in_mb"] = int(*obj.RecordingSizeInMb)
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, JfrAttachmentTargetToMap(item))
	}
	result["targets"] = targets

	if obj.WaitingPeriodInMinutes != nil {
		result["waiting_period_in_minutes"] = int(*obj.WaitingPeriodInMinutes)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToRequestPerformanceTuningAnalysesDetails(fieldKeyFormat string) (oci_jms.RequestPerformanceTuningAnalysesDetails, error) {
	result := oci_jms.RequestPerformanceTuningAnalysesDetails{}

	if recordingDurationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "recording_duration_in_minutes")); ok {
		tmp := recordingDurationInMinutes.(int)
		result.RecordingDurationInMinutes = &tmp
	}

	if targets, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "targets")); ok {
		interfaces := targets.([]interface{})
		tmp := make([]oci_jms.JfrAttachmentTarget, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "targets"), stateDataIndex)
			converted, err := s.mapToJfrAttachmentTarget(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "targets")) {
			result.Targets = tmp
		}
	}

	if waitingPeriodInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "waiting_period_in_minutes")); ok {
		tmp := waitingPeriodInMinutes.(int)
		result.WaitingPeriodInMinutes = &tmp
	}

	return result, nil
}

func RequestPerformanceTuningAnalysesDetailsToMap(obj *oci_jms.RequestPerformanceTuningAnalysesDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.RecordingDurationInMinutes != nil {
		result["recording_duration_in_minutes"] = int(*obj.RecordingDurationInMinutes)
	}

	targets := []interface{}{}
	for _, item := range obj.Targets {
		targets = append(targets, JfrAttachmentTargetToMap(item))
	}
	result["targets"] = targets

	if obj.WaitingPeriodInMinutes != nil {
		result["waiting_period_in_minutes"] = int(*obj.WaitingPeriodInMinutes)
	}

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToScanJavaServerUsageDetails(fieldKeyFormat string) (oci_jms.ScanJavaServerUsageDetails, error) {
	result := oci_jms.ScanJavaServerUsageDetails{}

	if managedInstanceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")); ok {
		interfaces := managedInstanceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")) {
			result.ManagedInstanceIds = tmp
		}
	}

	return result, nil
}

func ScanJavaServerUsageDetailsToMap(obj *oci_jms.ScanJavaServerUsageDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["managed_instance_ids"] = obj.ManagedInstanceIds

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToScanLibraryUsageDetails(fieldKeyFormat string) (oci_jms.ScanLibraryUsageDetails, error) {
	result := oci_jms.ScanLibraryUsageDetails{}

	if dynamicScanDurationInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dynamic_scan_duration_in_minutes")); ok {
		tmp := dynamicScanDurationInMinutes.(int)
		result.DynamicScanDurationInMinutes = &tmp
	}

	if isDynamicScan, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_dynamic_scan")); ok {
		tmp := isDynamicScan.(bool)
		result.IsDynamicScan = &tmp
	}

	if managedInstanceIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")); ok {
		interfaces := managedInstanceIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "managed_instance_ids")) {
			result.ManagedInstanceIds = tmp
		}
	}

	return result, nil
}

func ScanLibraryUsageDetailsToMap(obj *oci_jms.ScanLibraryUsageDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DynamicScanDurationInMinutes != nil {
		result["dynamic_scan_duration_in_minutes"] = int(*obj.DynamicScanDurationInMinutes)
	}

	if obj.IsDynamicScan != nil {
		result["is_dynamic_scan"] = bool(*obj.IsDynamicScan)
	}

	result["managed_instance_ids"] = obj.ManagedInstanceIds

	return result
}

func (s *JmsTaskScheduleResourceCrud) mapToTaskDetails(fieldKeyFormat string) (oci_jms.TaskDetails, error) {
	var baseObject oci_jms.TaskDetails
	//discriminator
	taskTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "task_type"))
	var taskType string
	if ok {
		taskType = taskTypeRaw.(string)
	} else {
		taskType = "" // default value
	}
	switch strings.ToLower(taskType) {
	case strings.ToLower("ADD_INSTALLATION_SITE"):
		details := oci_jms.AddInstallationSiteTaskDetails{}
		if addInstallationSiteTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "add_installation_site_task_request")); ok {
			if tmpList := addInstallationSiteTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "add_installation_site_task_request"), 0)
				tmp, err := s.mapToAddFleetInstallationSitesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert add_installation_site_task_request, encountered error: %v", err)
				}
				details.AddInstallationSiteTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("CRYPTO"):
		details := oci_jms.CryptoTaskDetails{}
		if cryptoTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "crypto_task_request")); ok {
			if tmpList := cryptoTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "crypto_task_request"), 0)
				tmp, err := s.mapToRequestCryptoAnalysesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert crypto_task_request, encountered error: %v", err)
				}
				details.CryptoTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DEPLOYED_APPLICATION_MIGRATION"):
		details := oci_jms.DeployedApplicationMigrationTaskDetails{}
		if deployedApplicationMigrationTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deployed_application_migration_task_request")); ok {
			if tmpList := deployedApplicationMigrationTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "deployed_application_migration_task_request"), 0)
				tmp, err := s.mapToRequestDeployedApplicationMigrationAnalysesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert deployed_application_migration_task_request, encountered error: %v", err)
				}
				details.DeployedApplicationMigrationTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("JAVA_MIGRATION"):
		details := oci_jms.JavaMigrationTaskDetails{}
		if javaMigrationTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "java_migration_task_request")); ok {
			if tmpList := javaMigrationTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "java_migration_task_request"), 0)
				tmp, err := s.mapToRequestJavaMigrationAnalysesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert java_migration_task_request, encountered error: %v", err)
				}
				details.JavaMigrationTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("JFR"):
		details := oci_jms.JfrTaskDetails{}
		if jfrTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "jfr_task_request")); ok {
			if tmpList := jfrTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "jfr_task_request"), 0)
				tmp, err := s.mapToRequestJfrRecordingsDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert jfr_task_request, encountered error: %v", err)
				}
				details.JfrTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PERFORMANCE_TUNING"):
		details := oci_jms.PerformanceTuningTaskDetails{}
		if performanceTuningTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "performance_tuning_task_request")); ok {
			if tmpList := performanceTuningTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "performance_tuning_task_request"), 0)
				tmp, err := s.mapToRequestPerformanceTuningAnalysesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert performance_tuning_task_request, encountered error: %v", err)
				}
				details.PerformanceTuningTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("REMOVE_INSTALLATION_SITE"):
		details := oci_jms.RemoveInstallationSiteTaskDetails{}
		if removeInstallationSiteTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "remove_installation_site_task_request")); ok {
			if tmpList := removeInstallationSiteTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "remove_installation_site_task_request"), 0)
				tmp, err := s.mapToRemoveFleetInstallationSitesDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert remove_installation_site_task_request, encountered error: %v", err)
				}
				details.RemoveInstallationSiteTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SCAN_JAVA_SERVER"):
		details := oci_jms.ScanJavaServerTaskDetails{}
		if scanJavaServerTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_java_server_task_request")); ok {
			if tmpList := scanJavaServerTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scan_java_server_task_request"), 0)
				tmp, err := s.mapToScanJavaServerUsageDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scan_java_server_task_request, encountered error: %v", err)
				}
				details.ScanJavaServerTaskRequest = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("SCAN_LIBRARY"):
		details := oci_jms.ScanLibraryTaskDetails{}
		if scanLibraryTaskRequest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "scan_library_task_request")); ok {
			if tmpList := scanLibraryTaskRequest.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "scan_library_task_request"), 0)
				tmp, err := s.mapToScanLibraryUsageDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert scan_library_task_request, encountered error: %v", err)
				}
				details.ScanLibraryTaskRequest = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown task_type '%v' was specified", taskType)
	}
	return baseObject, nil
}

func TaskDetailsToMap(obj *oci_jms.TaskDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_jms.AddInstallationSiteTaskDetails:
		result["task_type"] = "ADD_INSTALLATION_SITE"

		if v.AddInstallationSiteTaskRequest != nil {
			result["add_installation_site_task_request"] = []interface{}{AddFleetInstallationSitesDetailsToMap(v.AddInstallationSiteTaskRequest)}
		}
	case oci_jms.CryptoTaskDetails:
		result["task_type"] = "CRYPTO"

		if v.CryptoTaskRequest != nil {
			result["crypto_task_request"] = []interface{}{RequestCryptoAnalysesDetailsToMap(v.CryptoTaskRequest)}
		}
	case oci_jms.DeployedApplicationMigrationTaskDetails:
		result["task_type"] = "DEPLOYED_APPLICATION_MIGRATION"

		if v.DeployedApplicationMigrationTaskRequest != nil {
			result["deployed_application_migration_task_request"] = []interface{}{RequestDeployedApplicationMigrationAnalysesDetailsToMap(v.DeployedApplicationMigrationTaskRequest)}
		}
	case oci_jms.JavaMigrationTaskDetails:
		result["task_type"] = "JAVA_MIGRATION"

		if v.JavaMigrationTaskRequest != nil {
			result["java_migration_task_request"] = []interface{}{RequestJavaMigrationAnalysesDetailsToMap(v.JavaMigrationTaskRequest)}
		}
	case oci_jms.JfrTaskDetails:
		result["task_type"] = "JFR"

		if v.JfrTaskRequest != nil {
			result["jfr_task_request"] = []interface{}{RequestJfrRecordingsDetailsToMap(v.JfrTaskRequest)}
		}
	case oci_jms.PerformanceTuningTaskDetails:
		result["task_type"] = "PERFORMANCE_TUNING"

		if v.PerformanceTuningTaskRequest != nil {
			result["performance_tuning_task_request"] = []interface{}{RequestPerformanceTuningAnalysesDetailsToMap(v.PerformanceTuningTaskRequest)}
		}
	case oci_jms.RemoveInstallationSiteTaskDetails:
		result["task_type"] = "REMOVE_INSTALLATION_SITE"

		if v.RemoveInstallationSiteTaskRequest != nil {
			result["remove_installation_site_task_request"] = []interface{}{RemoveFleetInstallationSitesDetailsToMap(v.RemoveInstallationSiteTaskRequest)}
		}
	case oci_jms.ScanJavaServerTaskDetails:
		result["task_type"] = "SCAN_JAVA_SERVER"

		if v.ScanJavaServerTaskRequest != nil {
			result["scan_java_server_task_request"] = []interface{}{ScanJavaServerUsageDetailsToMap(v.ScanJavaServerTaskRequest)}
		}
	case oci_jms.ScanLibraryTaskDetails:
		result["task_type"] = "SCAN_LIBRARY"

		if v.ScanLibraryTaskRequest != nil {
			result["scan_library_task_request"] = []interface{}{ScanLibraryUsageDetailsToMap(v.ScanLibraryTaskRequest)}
		}
	default:
		log.Printf("[WARN] Received 'task_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func TaskScheduleSummaryToMap(obj oci_jms.TaskScheduleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.ExecutionRecurrences != nil {
		result["execution_recurrences"] = string(*obj.ExecutionRecurrences)
	}

	if obj.FleetId != nil {
		result["fleet_id"] = string(*obj.FleetId)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TaskDetails != nil {
		taskDetailsArray := []interface{}{}
		if taskDetailsMap := TaskDetailsToMap(&obj.TaskDetails); taskDetailsMap != nil {
			taskDetailsArray = append(taskDetailsArray, taskDetailsMap)
		}
		result["task_details"] = taskDetailsArray
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeLastRun != nil {
		result["time_last_run"] = obj.TimeLastRun.String()
	}

	if obj.TimeLastUpdated != nil {
		result["time_last_updated"] = obj.TimeLastUpdated.String()
	}

	if obj.TimeNextRun != nil {
		result["time_next_run"] = obj.TimeNextRun.String()
	}

	return result
}
