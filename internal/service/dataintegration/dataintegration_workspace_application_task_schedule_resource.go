// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dataintegration

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_dataintegration "github.com/oracle/oci-go-sdk/v65/dataintegration"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataintegrationWorkspaceApplicationTaskScheduleResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDataintegrationWorkspaceApplicationTaskSchedule,
		Read:     readDataintegrationWorkspaceApplicationTaskSchedule,
		Update:   updateDataintegrationWorkspaceApplicationTaskSchedule,
		Delete:   deleteDataintegrationWorkspaceApplicationTaskSchedule,
		Schema: map[string]*schema.Schema{
			// Required
			"application_key": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"identifier": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workspace_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"auth_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_provider_delegate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_time_millis": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"expected_duration": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"expected_duration_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_backfill_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_concurrent_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"model_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"number_of_retries": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"next_run_time_millis": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"object_version": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"parent_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"parent": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"root_doc_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"registry_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"aggregator_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"retry_delay": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"retry_delay_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_ref": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"frequency_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"model_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CUSTOM",
											"DAILY",
											"HOURLY",
											"MONTHLY",
											"MONTHLY_RULE",
											"WEEKLY",
										}, true),
									},

									// Optional
									"custom_expression": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"day_of_week": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"days": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"frequency": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interval": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"time": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"hour": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"minute": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"second": {
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"week_of_month": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"identifier": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_daylight_adjustment_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"aggregator": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"description": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"identifier": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": {
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
									"aggregator_key": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"count_statistics": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required
												"object_type_count_list": {
													Type:     schema.TypeList,
													Required: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional
															"object_count": {
																Type:             schema.TypeString,
																Optional:         true,
																Computed:         true,
																ValidateFunc:     tfresource.ValidateInt64TypeString,
																DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
															},
															"object_type": {
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},

															// Computed
														},
													},
												},

												// Optional

												// Computed
											},
										},
									},
									"created_by": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"created_by_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"identifier_path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"info_fields": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"is_favorite": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"labels": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"registry_version": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"time_created": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"time_updated": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
									},
									"updated_by": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"updated_by_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"model_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"model_version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"parent": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"timezone": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"start_time_millis": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},

			// Computed
			"last_run_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_run_time_millis": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"model_version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object_status": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"object_version": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"parent_ref": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"parent": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"root_doc_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"metadata": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"aggregator": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"identifier": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"key": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"aggregator_key": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"count_statistics": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"object_type_count_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"object_count": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"object_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"created_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"identifier_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"info_fields": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"is_favorite": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"labels": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"registry_version": {
							Type:     schema.TypeInt,
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
						"updated_by": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"updated_by_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"retry_attempts": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createDataintegrationWorkspaceApplicationTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.CreateResource(d, sync)
}

func readDataintegrationWorkspaceApplicationTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.ReadResource(sync)
}

func updateDataintegrationWorkspaceApplicationTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDataintegrationWorkspaceApplicationTaskSchedule(d *schema.ResourceData, m interface{}) error {
	sync := &DataintegrationWorkspaceApplicationTaskScheduleResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataIntegrationClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DataintegrationWorkspaceApplicationTaskScheduleResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_dataintegration.DataIntegrationClient
	Res                    *oci_dataintegration.TaskSchedule
	DisableNotFoundRetries bool
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) ID() string {
	return GetWorkspaceApplicationTaskScheduleCompositeId(s.D.Get("application_key").(string), *s.Res.Key, s.D.Get("workspace_id").(string))
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) Create() error {
	request := oci_dataintegration.CreateTaskScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if authMode, ok := s.D.GetOkExists("auth_mode"); ok {
		request.AuthMode = oci_dataintegration.CreateTaskScheduleDetailsAuthModeEnum(authMode.(string))
	}

	if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
		tmp := configProviderDelegate.(string)
		var configProviderDelegateObj oci_dataintegration.ConfigProvider
		err := json.Unmarshal([]byte(tmp), &configProviderDelegateObj)
		if err != nil {
			return err
		}
		request.ConfigProviderDelegate = &configProviderDelegateObj
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if endTimeMillis, ok := s.D.GetOkExists("end_time_millis"); ok {
		tmp := endTimeMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert endTimeMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.EndTimeMillis = &tmpInt64
	}

	if expectedDuration, ok := s.D.GetOkExists("expected_duration"); ok {
		tmp := expectedDuration.(float64)
		request.ExpectedDuration = &tmp
	}

	if expectedDurationUnit, ok := s.D.GetOkExists("expected_duration_unit"); ok {
		request.ExpectedDurationUnit = oci_dataintegration.CreateTaskScheduleDetailsExpectedDurationUnitEnum(expectedDurationUnit.(string))
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isBackfillEnabled, ok := s.D.GetOkExists("is_backfill_enabled"); ok {
		tmp := isBackfillEnabled.(bool)
		request.IsBackfillEnabled = &tmp
	}

	if isConcurrentAllowed, ok := s.D.GetOkExists("is_concurrent_allowed"); ok {
		tmp := isConcurrentAllowed.(bool)
		request.IsConcurrentAllowed = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numberOfRetries, ok := s.D.GetOkExists("number_of_retries"); ok {
		tmp := numberOfRetries.(int)
		request.NumberOfRetries = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if retryDelay, ok := s.D.GetOkExists("retry_delay"); ok {
		tmp := retryDelay.(float64)
		request.RetryDelay = &tmp
	}

	if retryDelayUnit, ok := s.D.GetOkExists("retry_delay_unit"); ok {
		request.RetryDelayUnit = oci_dataintegration.CreateTaskScheduleDetailsRetryDelayUnitEnum(retryDelayUnit.(string))
	}

	if scheduleRef, ok := s.D.GetOkExists("schedule_ref"); ok {
		if tmpList := scheduleRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule_ref", 0)
			tmp, err := s.mapToSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScheduleRef = &tmp
		}
	}

	if startTimeMillis, ok := s.D.GetOkExists("start_time_millis"); ok {
		tmp := startTimeMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert startTimeMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.StartTimeMillis = &tmpInt64
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.CreateTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) Get() error {
	request := oci_dataintegration.GetTaskScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if taskScheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskScheduleKey.(string)
		request.TaskScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	applicationKey, taskScheduleKey, workspaceId, err := parseWorkspaceApplicationTaskScheduleCompositeId(s.D.Id())
	if err == nil {
		request.ApplicationKey = &applicationKey
		request.TaskScheduleKey = &taskScheduleKey
		request.WorkspaceId = &workspaceId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.GetTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) Update() error {
	request := oci_dataintegration.UpdateTaskScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if authMode, ok := s.D.GetOkExists("auth_mode"); ok {
		request.AuthMode = oci_dataintegration.UpdateTaskScheduleDetailsAuthModeEnum(authMode.(string))
	}

	if configProviderDelegate, ok := s.D.GetOkExists("config_provider_delegate"); ok {
		tmp := configProviderDelegate.(string)
		var configProviderDelegateObj oci_dataintegration.ConfigProvider
		err := json.Unmarshal([]byte(tmp), &configProviderDelegateObj)
		if err != nil {
			return err
		}
		request.ConfigProviderDelegate = &configProviderDelegateObj
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if endTimeMillis, ok := s.D.GetOkExists("end_time_millis"); ok {
		tmp := endTimeMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert endTimeMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.EndTimeMillis = &tmpInt64
	}

	if expectedDuration, ok := s.D.GetOkExists("expected_duration"); ok {
		tmp := expectedDuration.(float64)
		request.ExpectedDuration = &tmp
	}

	if expectedDurationUnit, ok := s.D.GetOkExists("expected_duration_unit"); ok {
		request.ExpectedDurationUnit = oci_dataintegration.UpdateTaskScheduleDetailsExpectedDurationUnitEnum(expectedDurationUnit.(string))
	}

	if identifier, ok := s.D.GetOkExists("identifier"); ok {
		tmp := identifier.(string)
		request.Identifier = &tmp
	}

	if isBackfillEnabled, ok := s.D.GetOkExists("is_backfill_enabled"); ok {
		tmp := isBackfillEnabled.(bool)
		request.IsBackfillEnabled = &tmp
	}

	if isConcurrentAllowed, ok := s.D.GetOkExists("is_concurrent_allowed"); ok {
		tmp := isConcurrentAllowed.(bool)
		request.IsConcurrentAllowed = &tmp
	}

	if isEnabled, ok := s.D.GetOkExists("is_enabled"); ok {
		tmp := isEnabled.(bool)
		request.IsEnabled = &tmp
	}

	if key, ok := s.D.GetOkExists("key"); ok {
		tmp := key.(string)
		request.Key = &tmp
	}

	if modelType, ok := s.D.GetOkExists("model_type"); ok {
		tmp := modelType.(string)
		request.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists("model_version"); ok {
		tmp := modelVersion.(string)
		request.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if numberOfRetries, ok := s.D.GetOkExists("number_of_retries"); ok {
		tmp := numberOfRetries.(int)
		request.NumberOfRetries = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists("object_status"); ok {
		tmp := objectStatus.(int)
		request.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists("object_version"); ok {
		tmp := objectVersion.(int)
		request.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists("parent_ref"); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "parent_ref", 0)
			tmp, err := s.mapToParentReference(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ParentRef = &tmp
		}
	}

	if registryMetadata, ok := s.D.GetOkExists("registry_metadata"); ok {
		if tmpList := registryMetadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "registry_metadata", 0)
			tmp, err := s.mapToRegistryMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RegistryMetadata = &tmp
		}
	}

	if retryDelay, ok := s.D.GetOkExists("retry_delay"); ok {
		tmp := retryDelay.(float64)
		request.RetryDelay = &tmp
	}

	if retryDelayUnit, ok := s.D.GetOkExists("retry_delay_unit"); ok {
		request.RetryDelayUnit = oci_dataintegration.UpdateTaskScheduleDetailsRetryDelayUnitEnum(retryDelayUnit.(string))
	}

	if scheduleRef, ok := s.D.GetOkExists("schedule_ref"); ok {
		if tmpList := scheduleRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "schedule_ref", 0)
			tmp, err := s.mapToSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ScheduleRef = &tmp
		}
	}

	if startTimeMillis, ok := s.D.GetOkExists("start_time_millis"); ok {
		tmp := startTimeMillis.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert startTimeMillis string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.StartTimeMillis = &tmpInt64
	}

	if taskScheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskScheduleKey.(string)
		request.TaskScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	response, err := s.Client.UpdateTaskSchedule(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.TaskSchedule
	return nil
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) Delete() error {
	request := oci_dataintegration.DeleteTaskScheduleRequest{}

	if applicationKey, ok := s.D.GetOkExists("application_key"); ok {
		tmp := applicationKey.(string)
		request.ApplicationKey = &tmp
	}

	if taskScheduleKey, ok := s.D.GetOkExists("key"); ok {
		tmp := taskScheduleKey.(string)
		request.TaskScheduleKey = &tmp
	}

	if workspaceId, ok := s.D.GetOkExists("workspace_id"); ok {
		tmp := workspaceId.(string)
		request.WorkspaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "dataintegration")

	_, err := s.Client.DeleteTaskSchedule(context.Background(), request)
	return err
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) SetData() error {

	applicationKey, taskScheduleKey, workspaceId, err := parseWorkspaceApplicationTaskScheduleCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("application_key", &applicationKey)
		s.D.Set("key", &taskScheduleKey)
		s.D.Set("workspace_id", &workspaceId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	s.D.Set("auth_mode", s.Res.AuthMode)

	if s.Res.ConfigProviderDelegate != nil {
		tempDelegate, err := json.Marshal(*s.Res.ConfigProviderDelegate)
		if err == nil {
			// Convert the byte slice to a string
			jsonString := string(tempDelegate)

			// Set the JSON string in the data structure
			s.D.Set("config_provider_delegate", jsonString)
		} else {
			log.Printf("error in parsing delegate object: %v", err)
		}
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EndTimeMillis != nil {
		s.D.Set("end_time_millis", strconv.FormatInt(*s.Res.EndTimeMillis, 10))
	}

	if s.Res.ExpectedDuration != nil {
		s.D.Set("expected_duration", *s.Res.ExpectedDuration)
	}

	s.D.Set("expected_duration_unit", s.Res.ExpectedDurationUnit)

	if s.Res.Identifier != nil {
		s.D.Set("identifier", *s.Res.Identifier)
	}

	if s.Res.IsBackfillEnabled != nil {
		s.D.Set("is_backfill_enabled", *s.Res.IsBackfillEnabled)
	}

	if s.Res.IsConcurrentAllowed != nil {
		s.D.Set("is_concurrent_allowed", *s.Res.IsConcurrentAllowed)
	}

	if s.Res.IsEnabled != nil {
		s.D.Set("is_enabled", *s.Res.IsEnabled)
	}

	if s.Res.Key != nil {
		s.D.Set("key", *s.Res.Key)
	}

	if s.Res.LastRunDetails != nil {
		s.D.Set("last_run_details", []interface{}{LastRunDetailsToMap(s.Res.LastRunDetails)})
	} else {
		s.D.Set("last_run_details", nil)
	}

	if s.Res.Metadata != nil {
		s.D.Set("metadata", []interface{}{ObjectMetadataToMapForTaskSchedule(s.Res.Metadata)})
	} else {
		s.D.Set("metadata", nil)
	}

	if s.Res.ModelType != nil {
		s.D.Set("model_type", *s.Res.ModelType)
	}

	if s.Res.ModelVersion != nil {
		s.D.Set("model_version", *s.Res.ModelVersion)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ObjectStatus != nil {
		s.D.Set("object_status", *s.Res.ObjectStatus)
	}

	if s.Res.ObjectVersion != nil {
		s.D.Set("object_version", *s.Res.ObjectVersion)
	}

	if s.Res.ParentRef != nil {
		s.D.Set("parent_ref", []interface{}{ParentReferenceToMapTaskSchedule(s.Res.ParentRef)})
	} else {
		s.D.Set("parent_ref", nil)
	}

	if s.Res.RetryAttempts != nil {
		s.D.Set("retry_attempts", *s.Res.RetryAttempts)
	}

	if s.Res.RetryDelay != nil {
		s.D.Set("retry_delay", *s.Res.RetryDelay)
	}

	s.D.Set("retry_delay_unit", s.Res.RetryDelayUnit)

	if s.Res.ScheduleRef != nil {
		s.D.Set("schedule_ref", []interface{}{ScheduleToMap(s.Res.ScheduleRef)})
	} else {
		s.D.Set("schedule_ref", nil)
	}

	if s.Res.StartTimeMillis != nil {
		s.D.Set("start_time_millis", strconv.FormatInt(*s.Res.StartTimeMillis, 10))
	}

	return nil
}

func GetWorkspaceApplicationTaskScheduleCompositeId(applicationKey string, taskScheduleKey string, workspaceId string) string {
	applicationKey = url.PathEscape(applicationKey)
	taskScheduleKey = url.PathEscape(taskScheduleKey)
	workspaceId = url.PathEscape(workspaceId)
	compositeId := "workspaces/" + workspaceId + "/applications/" + applicationKey + "/taskSchedules/" + taskScheduleKey
	return compositeId
}

func parseWorkspaceApplicationTaskScheduleCompositeId(compositeId string) (applicationKey string, taskScheduleKey string, workspaceId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("workspaces/.*/applications/.*/taskSchedules/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	workspaceId, _ = url.PathUnescape(parts[1])
	applicationKey, _ = url.PathUnescape(parts[3])
	taskScheduleKey, _ = url.PathUnescape(parts[5])

	return
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToAbstractFrequencyDetails(fieldKeyFormat string) (oci_dataintegration.AbstractFrequencyDetails, error) {
	var baseObject oci_dataintegration.AbstractFrequencyDetails
	//discriminator
	modelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type"))
	var modelType string
	if ok {
		modelType = modelTypeRaw.(string)
	} else {
		modelType = "" // default value
	}
	switch strings.ToLower(modelType) {
	case strings.ToLower("CUSTOM"):
		details := oci_dataintegration.CustomFrequencyDetails{}
		if customExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_expression")); ok {
			tmp := customExpression.(string)
			details.CustomExpression = &tmp
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("DAILY"):
		details := oci_dataintegration.DailyFrequencyDetails{}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("HOURLY"):
		details := oci_dataintegration.HourlyFrequencyDetails{}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("MONTHLY"):
		details := oci_dataintegration.MonthlyFrequencyDetails{}
		if days, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days")); ok {
			interfaces := days.([]interface{})
			tmp := make([]int, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(int)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days")) {
				details.Days = tmp
			}
		}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("MONTHLY_RULE"):
		details := oci_dataintegration.MonthlyRuleFrequencyDetails{}
		if dayOfWeek, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "day_of_week")); ok {
			details.DayOfWeek = oci_dataintegration.MonthlyRuleFrequencyDetailsDayOfWeekEnum(dayOfWeek.(string))
		}
		if interval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval")); ok {
			tmp := interval.(int)
			details.Interval = &tmp
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if weekOfMonth, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "week_of_month")); ok {
			details.WeekOfMonth = oci_dataintegration.MonthlyRuleFrequencyDetailsWeekOfMonthEnum(weekOfMonth.(string))
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	case strings.ToLower("WEEKLY"):
		details := oci_dataintegration.WeeklyFrequencyDetails{}
		if days, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "days")); ok {
			interfaces := days.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "days")) {
				details.Days = make([]oci_dataintegration.WeeklyFrequencyDetailsDaysEnum, len(tmp))
				for i := range tmp {
					details.Days[i] = oci_dataintegration.WeeklyFrequencyDetailsDaysEnum(tmp[i])
				}
			}
		}
		if time, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time")); ok {
			if tmpList := time.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "time"), 0)
				tmp, err := s.mapToTime(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert time, encountered error: %v", err)
				}
				details.Time = &tmp
			}
		}
		if frequency, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency")); ok {
			details.Frequency = oci_dataintegration.AbstractFrequencyDetailsFrequencyEnum(frequency.(string))
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown model_type '%v' was specified", modelType)
	}
	return baseObject, nil
}

func AbstractFrequencyDetailsToMap(obj *oci_dataintegration.AbstractFrequencyDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_dataintegration.CustomFrequencyDetails:
		result["model_type"] = "CUSTOM"

		if v.CustomExpression != nil {
			result["custom_expression"] = string(*v.CustomExpression)
		}

		result["frequency"] = string(v.Frequency)
	case oci_dataintegration.DailyFrequencyDetails:
		result["model_type"] = "DAILY"

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMap(v.Time)}
		}

		result["frequency"] = string(v.Frequency)
	case oci_dataintegration.HourlyFrequencyDetails:
		result["model_type"] = "HOURLY"

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMap(v.Time)}
		}

		result["frequency"] = string(v.Frequency)
	case oci_dataintegration.MonthlyFrequencyDetails:
		result["model_type"] = "MONTHLY"

		result["days"] = v.Days

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMap(v.Time)}
		}

		result["frequency"] = string(v.Frequency)
	case oci_dataintegration.MonthlyRuleFrequencyDetails:
		result["model_type"] = "MONTHLY_RULE"

		result["day_of_week"] = string(v.DayOfWeek)

		if v.Interval != nil {
			result["interval"] = int(*v.Interval)
		}

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMap(v.Time)}
		}

		result["week_of_month"] = string(v.WeekOfMonth)

		result["frequency"] = string(v.Frequency)
	case oci_dataintegration.WeeklyFrequencyDetails:
		result["model_type"] = "WEEKLY"

		result["days"] = v.Days

		if v.Time != nil {
			result["time"] = []interface{}{TimeToMap(v.Time)}
		}

		result["frequency"] = string(v.Frequency)
	default:
		log.Printf("[WARN] Received 'model_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToAggregatorSummary(fieldKeyFormat string) (oci_dataintegration.AggregatorSummary, error) {
	result := oci_dataintegration.AggregatorSummary{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	return result, nil
}

func AggregatorSummaryToMapForTaskSchedule(obj *oci_dataintegration.AggregatorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToCountStatistic(fieldKeyFormat string) (oci_dataintegration.CountStatistic, error) {
	result := oci_dataintegration.CountStatistic{}

	if objectTypeCountList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")); ok {
		interfaces := objectTypeCountList.([]interface{})
		tmp := make([]oci_dataintegration.CountStatisticSummary, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_type_count_list"), stateDataIndex)
			converted, err := s.mapToCountStatisticSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_type_count_list")) {
			result.ObjectTypeCountList = tmp
		}
	}

	return result, nil
}

func CountStatisticToMapForTaskSchedule(obj *oci_dataintegration.CountStatistic) map[string]interface{} {
	result := map[string]interface{}{}

	objectTypeCountList := []interface{}{}
	for _, item := range obj.ObjectTypeCountList {
		objectTypeCountList = append(objectTypeCountList, CountStatisticSummaryToMapForTaskSchedule(item))
	}
	result["object_type_count_list"] = objectTypeCountList

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToCountStatisticSummary(fieldKeyFormat string) (oci_dataintegration.CountStatisticSummary, error) {
	result := oci_dataintegration.CountStatisticSummary{}

	if objectCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_count")); ok {
		tmp := objectCount.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert objectCount string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.ObjectCount = &tmpInt64
	}

	if objectType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_type")); ok {
		result.ObjectType = oci_dataintegration.CountStatisticSummaryObjectTypeEnum(objectType.(string))
	}

	return result, nil
}

func CountStatisticSummaryToMapForTaskSchedule(obj oci_dataintegration.CountStatisticSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ObjectCount != nil {
		result["object_count"] = strconv.FormatInt(*obj.ObjectCount, 10)
	}

	result["object_type"] = string(obj.ObjectType)

	return result
}

func LastRunDetailsToMap(obj *oci_dataintegration.LastRunDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LastRunTimeMillis != nil {
		result["last_run_time_millis"] = strconv.FormatInt(*obj.LastRunTimeMillis, 10)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMapTaskSchedule(obj.ParentRef)}
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToObjectMetadata(fieldKeyFormat string) (oci_dataintegration.ObjectMetadata, error) {
	result := oci_dataintegration.ObjectMetadata{}

	if aggregator, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator")); ok {
		if tmpList := aggregator.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "aggregator"), 0)
			tmp, err := s.mapToAggregatorSummary(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert aggregator, encountered error: %v", err)
			}
			result.Aggregator = &tmp
		}
	}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if countStatistics, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "count_statistics")); ok {
		if tmpList := countStatistics.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "count_statistics"), 0)
			tmp, err := s.mapToCountStatistic(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert count_statistics, encountered error: %v", err)
			}
			result.CountStatistics = &tmp
		}
	}

	if createdBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by")); ok {
		tmp := createdBy.(string)
		result.CreatedBy = &tmp
	}

	if createdByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "created_by_name")); ok {
		tmp := createdByName.(string)
		result.CreatedByName = &tmp
	}

	if identifierPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier_path")); ok {
		tmp := identifierPath.(string)
		result.IdentifierPath = &tmp
	}

	if infoFields, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "info_fields")); ok {
		result.InfoFields = tfresource.ObjectMapToStringMap(infoFields.(map[string]interface{}))
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if timeUpdated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_updated")); ok {
		tmp, err := time.Parse(time.RFC3339, timeUpdated.(string))
		if err != nil {
			return result, err
		}
		result.TimeUpdated = &oci_common.SDKTime{Time: tmp}
	}

	if updatedBy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by")); ok {
		tmp := updatedBy.(string)
		result.UpdatedBy = &tmp
	}

	if updatedByName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "updated_by_name")); ok {
		tmp := updatedByName.(string)
		result.UpdatedByName = &tmp
	}

	return result, nil
}

func ObjectMetadataToMapForTaskSchedule(obj *oci_dataintegration.ObjectMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Aggregator != nil {
		result["aggregator"] = []interface{}{AggregatorSummaryToMapForTaskSchedule(obj.Aggregator)}
	}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.CountStatistics != nil {
		result["count_statistics"] = []interface{}{CountStatisticToMapForTaskSchedule(obj.CountStatistics)}
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.CreatedByName != nil {
		result["created_by_name"] = string(*obj.CreatedByName)
	}

	if obj.IdentifierPath != nil {
		result["identifier_path"] = string(*obj.IdentifierPath)
	}

	result["info_fields"] = obj.InfoFields

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.Format(time.RFC3339Nano)
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.Format(time.RFC3339Nano)
	}

	if obj.UpdatedBy != nil {
		result["updated_by"] = string(*obj.UpdatedBy)
	}

	if obj.UpdatedByName != nil {
		result["updated_by_name"] = string(*obj.UpdatedByName)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToParentReference(fieldKeyFormat string) (oci_dataintegration.ParentReference, error) {
	result := oci_dataintegration.ParentReference{}

	if parent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent")); ok {
		tmp := parent.(string)
		result.Parent = &tmp
	}

	if rootDocId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "root_doc_id")); ok {
		tmp := rootDocId.(string)
		result.RootDocId = &tmp
	}

	return result, nil
}

func ParentReferenceToMapTaskSchedule(obj *oci_dataintegration.ParentReference) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Parent != nil {
		result["parent"] = string(*obj.Parent)
	}

	if obj.RootDocId != nil {
		result["root_doc_id"] = string(*obj.RootDocId)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToRegistryMetadata(fieldKeyFormat string) (oci_dataintegration.RegistryMetadata, error) {
	result := oci_dataintegration.RegistryMetadata{}

	if aggregatorKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "aggregator_key")); ok {
		tmp := aggregatorKey.(string)
		result.AggregatorKey = &tmp
	}

	if isFavorite, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_favorite")); ok {
		tmp := isFavorite.(bool)
		result.IsFavorite = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if labels, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "labels")); ok {
		interfaces := labels.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "labels")) {
			result.Labels = tmp
		}
	}

	if registryVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "registry_version")); ok {
		tmp := registryVersion.(int)
		result.RegistryVersion = &tmp
	}

	return result, nil
}

func RegistryMetadataToMapForTaskSchedule(obj *oci_dataintegration.RegistryMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AggregatorKey != nil {
		result["aggregator_key"] = string(*obj.AggregatorKey)
	}

	if obj.IsFavorite != nil {
		result["is_favorite"] = bool(*obj.IsFavorite)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	result["labels"] = obj.Labels

	if obj.RegistryVersion != nil {
		result["registry_version"] = int(*obj.RegistryVersion)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToSchedule(fieldKeyFormat string) (oci_dataintegration.Schedule, error) {
	result := oci_dataintegration.Schedule{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if frequencyDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "frequency_details")); ok {
		if tmpList := frequencyDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "frequency_details"), 0)
			tmp, err := s.mapToAbstractFrequencyDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert frequency_details, encountered error: %v", err)
			}
			result.FrequencyDetails = tmp
		}
	}

	if identifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identifier")); ok {
		tmp := identifier.(string)
		result.Identifier = &tmp
	}

	if isDaylightAdjustmentEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_daylight_adjustment_enabled")); ok {
		tmp := isDaylightAdjustmentEnabled.(bool)
		result.IsDaylightAdjustmentEnabled = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if tmpList := metadata.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "metadata"), 0)
			tmp, err := s.mapToObjectMetadata(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert metadata, encountered error: %v", err)
			}
			result.Metadata = &tmp
		}
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	if modelVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_version")); ok {
		tmp := modelVersion.(string)
		result.ModelVersion = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if objectStatus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_status")); ok {
		tmp := objectStatus.(int)
		result.ObjectStatus = &tmp
	}

	if objectVersion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_version")); ok {
		tmp := objectVersion.(int)
		result.ObjectVersion = &tmp
	}

	if parentRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parent_ref")); ok {
		if tmpList := parentRef.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "parent_ref"), 0)
			tmp, err := s.mapToParentReference(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert parent_ref, encountered error: %v", err)
			}
			result.ParentRef = &tmp
		}
	}

	if timezone, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timezone")); ok {
		tmp := timezone.(string)
		result.Timezone = &tmp
	}

	return result, nil
}

func ScheduleToMap(obj *oci_dataintegration.Schedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.FrequencyDetails != nil {
		frequencyDetailsArray := []interface{}{}
		if frequencyDetailsMap := AbstractFrequencyDetailsToMap(&obj.FrequencyDetails); frequencyDetailsMap != nil {
			frequencyDetailsArray = append(frequencyDetailsArray, frequencyDetailsMap)
		}
		result["frequency_details"] = frequencyDetailsArray
	}

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.IsDaylightAdjustmentEnabled != nil {
		result["is_daylight_adjustment_enabled"] = bool(*obj.IsDaylightAdjustmentEnabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMapForTaskSchedule(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMapTaskSchedule(obj.ParentRef)}
	}

	if obj.Timezone != nil {
		result["timezone"] = string(*obj.Timezone)
	}

	return result
}

func TaskScheduleSummaryToMap(obj oci_dataintegration.TaskScheduleSummary) map[string]interface{} {
	result := map[string]interface{}{}

	result["auth_mode"] = string(obj.AuthMode)

	if obj.ConfigProviderDelegate != nil {
		tmp, _ := json.Marshal(obj.ConfigProviderDelegate)
		result["config_provider_delegate"] = string(tmp)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.EndTimeMillis != nil {
		result["end_time_millis"] = strconv.FormatInt(*obj.EndTimeMillis, 10)
	}

	if obj.ExpectedDuration != nil {
		result["expected_duration"] = float64(*obj.ExpectedDuration)
	}

	result["expected_duration_unit"] = string(obj.ExpectedDurationUnit)

	if obj.Identifier != nil {
		result["identifier"] = string(*obj.Identifier)
	}

	if obj.IsBackfillEnabled != nil {
		result["is_backfill_enabled"] = bool(*obj.IsBackfillEnabled)
	}

	if obj.IsConcurrentAllowed != nil {
		result["is_concurrent_allowed"] = bool(*obj.IsConcurrentAllowed)
	}

	if obj.IsEnabled != nil {
		result["is_enabled"] = bool(*obj.IsEnabled)
	}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LastRunDetails != nil {
		result["last_run_details"] = []interface{}{LastRunDetailsToMap(obj.LastRunDetails)}
	}

	if obj.Metadata != nil {
		result["metadata"] = []interface{}{ObjectMetadataToMapForTaskSchedule(obj.Metadata)}
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	if obj.ModelVersion != nil {
		result["model_version"] = string(*obj.ModelVersion)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.NextRunTimeMillis != nil {
		result["next_run_time_millis"] = strconv.FormatInt(*obj.NextRunTimeMillis, 10)
	}

	if obj.NumberOfRetries != nil {
		result["number_of_retries"] = int(*obj.NumberOfRetries)
	}

	if obj.ObjectStatus != nil {
		result["object_status"] = int(*obj.ObjectStatus)
	}

	if obj.ObjectVersion != nil {
		result["object_version"] = int(*obj.ObjectVersion)
	}

	if obj.ParentRef != nil {
		result["parent_ref"] = []interface{}{ParentReferenceToMapTaskSchedule(obj.ParentRef)}
	}

	if obj.RetryDelay != nil {
		result["retry_delay"] = float64(*obj.RetryDelay)
	}

	result["retry_delay_unit"] = string(obj.RetryDelayUnit)

	if obj.ScheduleRef != nil {
		result["schedule_ref"] = []interface{}{ScheduleToMap(obj.ScheduleRef)}
	}

	if obj.StartTimeMillis != nil {
		result["start_time_millis"] = strconv.FormatInt(*obj.StartTimeMillis, 10)
	}

	return result
}

func (s *DataintegrationWorkspaceApplicationTaskScheduleResourceCrud) mapToTime(fieldKeyFormat string) (oci_dataintegration.Time, error) {
	result := oci_dataintegration.Time{}

	if hour, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hour")); ok {
		tmp := hour.(int)
		result.Hour = &tmp
	}

	if minute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "minute")); ok {
		tmp := minute.(int)
		result.Minute = &tmp
	}

	if second, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "second")); ok {
		tmp := second.(int)
		result.Second = &tmp
	}

	return result, nil
}

func TimeToMap(obj *oci_dataintegration.Time) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Hour != nil {
		result["hour"] = int(*obj.Hour)
	}

	if obj.Minute != nil {
		result["minute"] = int(*obj.Minute)
	}

	if obj.Second != nil {
		result["second"] = int(*obj.Second)
	}

	return result
}
