// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apm_synthetics

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_apm_synthetics "github.com/oracle/oci-go-sdk/v65/apmsynthetics"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApmSyntheticsMonitorResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApmSyntheticsMonitor,
		Read:     readApmSyntheticsMonitor,
		Update:   updateApmSyntheticsMonitor,
		Delete:   deleteApmSyntheticsMonitor,
		Schema: map[string]*schema.Schema{
			// Required
			"apm_domain_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"monitor_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repeat_interval_in_seconds": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"vantage_points": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 255,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Computed
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
				},
			},

			// Optional
			"availability_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"max_allowed_failures_per_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_allowed_runs_per_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"batch_interval_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"client_certificate_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"client_certificate": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"content": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"file_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"private_key": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"content": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"file_name": {
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
						"config_type": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"BROWSER_CONFIG",
								"DNSSEC_CONFIG",
								"DNS_SERVER_CONFIG",
								"DNS_TRACE_CONFIG",
								"NETWORK_CONFIG",
								"REST_CONFIG",
								"SCRIPTED_BROWSER_CONFIG",
								"SCRIPTED_REST_CONFIG",
							}, true),
						},
						"dns_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"is_override_dns": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"override_dns_ip": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"is_certificate_validation_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_default_snapshot_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_failure_retried": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_query_recursive": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_redirection_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name_server": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"network_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"number_of_hops": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"probe_mode": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"probe_per_hop": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"transmission_rate": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"record_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"req_authentication_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"auth_headers": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional
												"header_name": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"header_value": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},

												// Computed
											},
										},
									},
									"auth_request_method": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_request_post_body": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_token": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_url": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_user_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_user_password": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										Sensitive: true,
									},
									"oauth_scheme": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"req_authentication_scheme": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"request_headers": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"header_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"header_value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"request_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"request_post_body": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"request_query_params": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"param_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"param_value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"verify_response_codes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 3,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"verify_response_content": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"verify_texts": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"text": {
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
			"is_run_now": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_run_once": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"maintenance_window_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"time_ended": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},
						"time_started": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"scheduling_policy": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"param_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"param_value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"is_overwritten": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_secret": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"monitor_script_parameter": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"param_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"param_value": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"script_name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"vantage_point_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createApmSyntheticsMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.CreateResource(d, sync)
}

func readApmSyntheticsMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.ReadResource(sync)
}

func updateApmSyntheticsMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApmSyntheticsMonitor(d *schema.ResourceData, m interface{}) error {
	sync := &ApmSyntheticsMonitorResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ApmSyntheticClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ApmSyntheticsMonitorResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apm_synthetics.ApmSyntheticClient
	Res                    *oci_apm_synthetics.Monitor
	DisableNotFoundRetries bool
}

func (s *ApmSyntheticsMonitorResourceCrud) ID() string {
	return GetMonitorCompositeId(*s.Res.Id, s.D.Get("apm_domain_id").(string))
}

func (s *ApmSyntheticsMonitorResourceCrud) Create() error {
	request := oci_apm_synthetics.CreateMonitorRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if availabilityConfiguration, ok := s.D.GetOkExists("availability_configuration"); ok {
		if tmpList := availabilityConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_configuration", 0)
			tmp, err := s.mapToAvailabilityConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityConfiguration = &tmp
		}
	}

	if batchIntervalInSeconds, ok := s.D.GetOkExists("batch_interval_in_seconds"); ok {
		tmp := batchIntervalInSeconds.(int)
		request.BatchIntervalInSeconds = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", 0)
			tmp, err := s.mapToMonitorConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Configuration = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRunNow, ok := s.D.GetOkExists("is_run_now"); ok {
		tmp := isRunNow.(bool)
		request.IsRunNow = &tmp
	}

	if isRunOnce, ok := s.D.GetOkExists("is_run_once"); ok {
		tmp := isRunOnce.(bool)
		request.IsRunOnce = &tmp
	}

	if maintenanceWindowSchedule, ok := s.D.GetOkExists("maintenance_window_schedule"); ok {
		if tmpList := maintenanceWindowSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_schedule", 0)
			tmp, err := s.mapToMaintenanceWindowSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowSchedule = &tmp
		}
	}

	if monitorType, ok := s.D.GetOkExists("monitor_type"); ok {
		request.MonitorType = oci_apm_synthetics.MonitorTypesEnum(monitorType.(string))
	}

	if repeatIntervalInSeconds, ok := s.D.GetOkExists("repeat_interval_in_seconds"); ok {
		tmp := repeatIntervalInSeconds.(int)
		request.RepeatIntervalInSeconds = &tmp
	}

	if schedulingPolicy, ok := s.D.GetOkExists("scheduling_policy"); ok {
		request.SchedulingPolicy = oci_apm_synthetics.SchedulingPolicyEnum(schedulingPolicy.(string))
	}

	if compositeId, ok := s.D.GetOkExists("script_id"); ok {
		tmp := compositeId.(string)
		scriptId, apmDomainId, err := parseScriptCompositeId(tmp)
		if err == nil {
			request.ScriptId = &scriptId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	if scriptParameters, ok := s.D.GetOkExists("script_parameters"); ok {
		interfaces := scriptParameters.([]interface{})
		tmp := make([]oci_apm_synthetics.MonitorScriptParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "script_parameters", stateDataIndex)
			converted, err := s.mapToMonitorScriptParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("script_parameters") {
			request.ScriptParameters = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.MonitorStatusEnum(status.(string))
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		tmp := target.(string)
		request.Target = &tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	if vantagePoints, ok := s.D.GetOkExists("vantage_points"); ok {
		interfaces := vantagePoints.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vantage_points", stateDataIndex)
			converted, err := s.mapToVantagePointInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			name := converted.Name
			tmp[i] = *name
		}
		if len(tmp) != 0 || s.D.HasChange("vantage_points") {
			request.VantagePoints = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.CreateMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Monitor
	return nil
}

func (s *ApmSyntheticsMonitorResourceCrud) Get() error {
	request := oci_apm_synthetics.GetMonitorRequest{}

	monitorId, apmDomainId, err := parseMonitorCompositeId(s.D.Id())
	if err == nil {
		request.MonitorId = &monitorId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.GetMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Monitor
	return nil
}

func (s *ApmSyntheticsMonitorResourceCrud) Update() error {
	request := oci_apm_synthetics.UpdateMonitorRequest{}

	if apmDomainId, ok := s.D.GetOkExists("apm_domain_id"); ok {
		tmp := apmDomainId.(string)
		request.ApmDomainId = &tmp
	}

	if availabilityConfiguration, ok := s.D.GetOkExists("availability_configuration"); ok {
		if tmpList := availabilityConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "availability_configuration", 0)
			tmp, err := s.mapToAvailabilityConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AvailabilityConfiguration = &tmp
		}
	}

	if batchIntervalInSeconds, ok := s.D.GetOkExists("batch_interval_in_seconds"); ok {
		tmp := batchIntervalInSeconds.(int)
		request.BatchIntervalInSeconds = &tmp
	}

	if configuration, ok := s.D.GetOkExists("configuration"); ok {
		if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration", 0)
			tmp, err := s.mapToMonitorConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Configuration = tmp
		}
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if isRunNow, ok := s.D.GetOkExists("is_run_now"); ok {
		tmp := isRunNow.(bool)
		request.IsRunNow = &tmp
	}

	if isRunOnce, ok := s.D.GetOkExists("is_run_once"); ok {
		tmp := isRunOnce.(bool)
		request.IsRunOnce = &tmp
	}

	if maintenanceWindowSchedule, ok := s.D.GetOkExists("maintenance_window_schedule"); ok {
		if tmpList := maintenanceWindowSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "maintenance_window_schedule", 0)
			tmp, err := s.mapToMaintenanceWindowSchedule(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MaintenanceWindowSchedule = &tmp
		}
	}

	monitorId, apmDomainId, err := parseMonitorCompositeId(s.D.Id())
	if err == nil {
		request.MonitorId = &monitorId
		request.ApmDomainId = &apmDomainId
	} else {
		log.Printf("[WARN] Update() unable to parse current ID: %s", s.D.Id())
	}

	if repeatIntervalInSeconds, ok := s.D.GetOkExists("repeat_interval_in_seconds"); ok {
		tmp := repeatIntervalInSeconds.(int)
		request.RepeatIntervalInSeconds = &tmp
	}

	if schedulingPolicy, ok := s.D.GetOkExists("scheduling_policy"); ok {
		request.SchedulingPolicy = oci_apm_synthetics.SchedulingPolicyEnum(schedulingPolicy.(string))
	}

	if compositeId, ok := s.D.GetOkExists("script_id"); ok {
		tmp := compositeId.(string)
		scriptId, apmDomainId, err := parseScriptCompositeId(tmp)
		if err == nil {
			request.ScriptId = &scriptId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Update() unable to parse current ID: %s", s.D.Id())
		}
	}

	if scriptParameters, ok := s.D.GetOkExists("script_parameters"); ok {
		interfaces := scriptParameters.([]interface{})
		tmp := make([]oci_apm_synthetics.MonitorScriptParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "script_parameters", stateDataIndex)
			converted, err := s.mapToMonitorScriptParameter(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("script_parameters") {
			request.ScriptParameters = tmp
		}
	}

	if status, ok := s.D.GetOkExists("status"); ok {
		request.Status = oci_apm_synthetics.MonitorStatusEnum(status.(string))
	}

	if target, ok := s.D.GetOkExists("target"); ok {
		tmp := target.(string)
		request.Target = &tmp
	}

	if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
		tmp := timeoutInSeconds.(int)
		request.TimeoutInSeconds = &tmp
	}

	if vantagePoints, ok := s.D.GetOkExists("vantage_points"); ok {
		interfaces := vantagePoints.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "vantage_points", stateDataIndex)
			converted, err := s.mapToVantagePointInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			name := converted.Name
			tmp[i] = *name
		}
		if len(tmp) != 0 || s.D.HasChange("vantage_points") {
			request.VantagePoints = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	response, err := s.Client.UpdateMonitor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Monitor
	return nil
}

func (s *ApmSyntheticsMonitorResourceCrud) Delete() error {
	request := oci_apm_synthetics.DeleteMonitorRequest{}

	tmp := s.D.Id()

	if tmp != "" {
		monitorId, apmDomainId, err := parseMonitorCompositeId(s.D.Id())
		if err == nil {
			request.MonitorId = &monitorId
			request.ApmDomainId = &apmDomainId
		} else {
			log.Printf("[WARN] Delete() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apm_synthetics")

	_, err := s.Client.DeleteMonitor(context.Background(), request)
	return err
}

func (s *ApmSyntheticsMonitorResourceCrud) SetData() error {

	_, apmDomainId, err := parseMonitorCompositeId(s.D.Id())
	if err == nil {
		if apmDomainId != "" {

		}
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AvailabilityConfiguration != nil {
		s.D.Set("availability_configuration", []interface{}{AvailabilityConfigurationToMap(s.Res.AvailabilityConfiguration)})
	} else {
		s.D.Set("availability_configuration", nil)
	}

	if s.Res.BatchIntervalInSeconds != nil {
		s.D.Set("batch_interval_in_seconds", *s.Res.BatchIntervalInSeconds)
	}

	if s.Res.Configuration != nil {
		configurationArray := []interface{}{}
		if configurationMap := MonitorConfigurationToMap(&s.Res.Configuration); configurationMap != nil {
			configurationArray = append(configurationArray, configurationMap)
		}
		s.D.Set("configuration", configurationArray)
	} else {
		s.D.Set("configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsRunNow != nil {
		s.D.Set("is_run_now", *s.Res.IsRunNow)
	}

	if s.Res.IsRunOnce != nil {
		s.D.Set("is_run_once", *s.Res.IsRunOnce)
	}

	if s.Res.MaintenanceWindowSchedule != nil {
		s.D.Set("maintenance_window_schedule", []interface{}{MaintenanceWindowScheduleToMap(s.Res.MaintenanceWindowSchedule)})
	} else {
		s.D.Set("maintenance_window_schedule", nil)
	}

	s.D.Set("monitor_type", s.Res.MonitorType)

	if s.Res.RepeatIntervalInSeconds != nil {
		s.D.Set("repeat_interval_in_seconds", *s.Res.RepeatIntervalInSeconds)
	}

	s.D.Set("scheduling_policy", s.Res.SchedulingPolicy)

	if s.Res.ScriptId != nil {
		s.D.Set("script_id", GetScriptCompositeId(*s.Res.ScriptId, apmDomainId))
	}

	if s.Res.ScriptName != nil {
		s.D.Set("script_name", *s.Res.ScriptName)
	}

	scriptParameters := []interface{}{}
	for _, item := range s.Res.ScriptParameters {
		scriptParameters = append(scriptParameters, MonitorScriptParameterInfoToMap(item))
	}
	s.D.Set("script_parameters", scriptParameters)

	s.D.Set("status", s.Res.Status)

	if s.Res.Target != nil {
		s.D.Set("target", *s.Res.Target)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeoutInSeconds != nil {
		s.D.Set("timeout_in_seconds", *s.Res.TimeoutInSeconds)
	}

	if s.Res.VantagePointCount != nil {
		s.D.Set("vantage_point_count", *s.Res.VantagePointCount)
	}

	vantagePoints := []interface{}{}
	for _, item := range s.Res.VantagePoints {
		vantagePoints = append(vantagePoints, VantagePointInfoToMap(item))
	}
	s.D.Set("vantage_points", vantagePoints)

	return nil
}

func GetMonitorCompositeId(monitorId string, apmDomainId string) string {
	apmDomainId = url.PathEscape(apmDomainId)
	monitorId = url.PathEscape(monitorId)
	compositeId := "monitors/" + monitorId + "/apmDomainId/" + apmDomainId
	return compositeId
}

func parseMonitorCompositeId(compositeId string) (monitorId string, apmDomainId string, err error) {
	parts := strings.Split(compositeId, "/")

	match, _ := regexp.MatchString("monitors/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	monitorId, _ = url.PathUnescape(parts[1])
	apmDomainId, _ = url.PathUnescape(parts[3])

	return
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToAvailabilityConfiguration(fieldKeyFormat string) (oci_apm_synthetics.AvailabilityConfiguration, error) {
	result := oci_apm_synthetics.AvailabilityConfiguration{}

	if maxAllowedFailuresPerInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_allowed_failures_per_interval")); ok {
		tmp := maxAllowedFailuresPerInterval.(int)
		result.MaxAllowedFailuresPerInterval = &tmp
	}

	if minAllowedRunsPerInterval, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_allowed_runs_per_interval")); ok {
		tmp := minAllowedRunsPerInterval.(int)
		result.MinAllowedRunsPerInterval = &tmp
	}

	return result, nil
}

func AvailabilityConfigurationToMap(obj *oci_apm_synthetics.AvailabilityConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxAllowedFailuresPerInterval != nil {
		result["max_allowed_failures_per_interval"] = int(*obj.MaxAllowedFailuresPerInterval)
	}

	if obj.MinAllowedRunsPerInterval != nil {
		result["min_allowed_runs_per_interval"] = int(*obj.MinAllowedRunsPerInterval)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToClientCertificate(fieldKeyFormat string) (oci_apm_synthetics.ClientCertificate, error) {
	result := oci_apm_synthetics.ClientCertificate{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		tmp := content.(string)
		result.Content = &tmp
	}

	if fileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_name")); ok {
		tmp := fileName.(string)
		result.FileName = &tmp
	}

	return result, nil
}

func ClientCertificateToMap(obj *oci_apm_synthetics.ClientCertificate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		result["content"] = string(*obj.Content)
	}

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToClientCertificateDetails(fieldKeyFormat string) (oci_apm_synthetics.ClientCertificateDetails, error) {
	result := oci_apm_synthetics.ClientCertificateDetails{}

	if clientCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_certificate")); ok {
		if tmpList := clientCertificate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "client_certificate"), 0)
			tmp, err := s.mapToClientCertificate(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert client_certificate, encountered error: %v", err)
			}
			result.ClientCertificate = &tmp
		}
	}

	if privateKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key")); ok {
		if tmpList := privateKey.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "private_key"), 0)
			tmp, err := s.mapToPrivateKey(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert private_key, encountered error: %v", err)
			}
			result.PrivateKey = &tmp
		}
	}

	return result, nil
}

func ClientCertificateDetailsToMap(obj *oci_apm_synthetics.ClientCertificateDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ClientCertificate != nil {
		result["client_certificate"] = []interface{}{ClientCertificateToMap(obj.ClientCertificate)}
	}

	if obj.PrivateKey != nil {
		result["private_key"] = []interface{}{PrivateKeyToMap(obj.PrivateKey)}
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToDnsConfiguration(fieldKeyFormat string) (oci_apm_synthetics.DnsConfiguration, error) {
	result := oci_apm_synthetics.DnsConfiguration{}

	if isOverrideDns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_override_dns")); ok {
		tmp := isOverrideDns.(bool)
		result.IsOverrideDns = &tmp
	}

	if overrideDnsIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "override_dns_ip")); ok {
		tmp := overrideDnsIp.(string)
		result.OverrideDnsIp = &tmp
	}

	return result, nil
}

func DnsConfigurationToMap(obj *oci_apm_synthetics.DnsConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IsOverrideDns != nil {
		result["is_override_dns"] = bool(*obj.IsOverrideDns)
	}

	if obj.OverrideDnsIp != nil {
		result["override_dns_ip"] = string(*obj.OverrideDnsIp)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToHeader(fieldKeyFormat string) (oci_apm_synthetics.Header, error) {
	result := oci_apm_synthetics.Header{}

	if headerName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_name")); ok {
		tmp := headerName.(string)
		result.HeaderName = &tmp
	}

	if headerValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "header_value")); ok {
		tmp := headerValue.(string)
		result.HeaderValue = &tmp
	}

	return result, nil
}

func SyntheticHeaderToMap(obj oci_apm_synthetics.Header) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.HeaderName != nil {
		result["header_name"] = string(*obj.HeaderName)
	}

	if obj.HeaderValue != nil {
		result["header_value"] = string(*obj.HeaderValue)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToMaintenanceWindowSchedule(fieldKeyFormat string) (oci_apm_synthetics.MaintenanceWindowSchedule, error) {
	result := oci_apm_synthetics.MaintenanceWindowSchedule{}

	if timeEnded, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_ended")); ok {
		tmp, err := time.Parse(time.RFC3339, timeEnded.(string))
		if err != nil {
			return result, err
		}
		result.TimeEnded = &oci_common.SDKTime{Time: tmp}
	}

	if timeStarted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_started")); ok {
		tmp, err := time.Parse(time.RFC3339, timeStarted.(string))
		if err != nil {
			return result, err
		}
		result.TimeStarted = &oci_common.SDKTime{Time: tmp}
	}

	return result, nil
}

func MaintenanceWindowScheduleToMap(obj *oci_apm_synthetics.MaintenanceWindowSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeEnded != nil {
		result["time_ended"] = obj.TimeEnded.Format(time.RFC3339Nano)
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.Format(time.RFC3339Nano)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToMonitorConfiguration(fieldKeyFormat string) (oci_apm_synthetics.MonitorConfiguration, error) {
	var baseObject oci_apm_synthetics.MonitorConfiguration
	//discriminator
	configTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "config_type"))
	var configType string
	if ok {
		configType = configTypeRaw.(string)
	} else {
		configType = "" // default value
	}
	switch strings.ToLower(configType) {
	case strings.ToLower("BROWSER_CONFIG"):
		details := oci_apm_synthetics.BrowserMonitorConfiguration{}
		if isCertificateValidationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_certificate_validation_enabled")); ok {
			tmp := isCertificateValidationEnabled.(bool)
			details.IsCertificateValidationEnabled = &tmp
		}
		if isDefaultSnapshotEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_default_snapshot_enabled")); ok {
			tmp := isDefaultSnapshotEnabled.(bool)
			details.IsDefaultSnapshotEnabled = &tmp
		}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if verifyResponseCodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")); ok {
			interfaces := verifyResponseCodes.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")) {
				details.VerifyResponseCodes = tmp
			}
		}
		if verifyTexts, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_texts")); ok {
			interfaces := verifyTexts.([]interface{})
			tmp := make([]oci_apm_synthetics.VerifyText, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "verify_texts"), stateDataIndex)
				converted, err := s.mapToVerifyText(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_texts")) {
				details.VerifyTexts = tmp
			}
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("DNSSEC_CONFIG"):
		details := oci_apm_synthetics.DnsSecMonitorConfiguration{}
		if recordType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_type")); ok {
			details.RecordType = oci_apm_synthetics.DnsRecordTypeEnum(recordType.(string))
		}
		if verifyResponseContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_content")); ok {
			tmp := verifyResponseContent.(string)
			details.VerifyResponseContent = &tmp
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("DNS_SERVER_CONFIG"):
		details := oci_apm_synthetics.DnsServerMonitorConfiguration{}
		if isQueryRecursive, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_query_recursive")); ok {
			tmp := isQueryRecursive.(bool)
			details.IsQueryRecursive = &tmp
		}
		if nameServer, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name_server")); ok {
			tmp := nameServer.(string)
			details.NameServer = &tmp
		}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
			details.Protocol = oci_apm_synthetics.DnsTransportProtocolEnum(protocol.(string))
		}
		if recordType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_type")); ok {
			details.RecordType = oci_apm_synthetics.DnsRecordTypeEnum(recordType.(string))
		}
		if verifyResponseContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_content")); ok {
			tmp := verifyResponseContent.(string)
			details.VerifyResponseContent = &tmp
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("DNS_TRACE_CONFIG"):
		details := oci_apm_synthetics.DnsTraceMonitorConfiguration{}
		if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
			details.Protocol = oci_apm_synthetics.DnsTransportProtocolEnum(protocol.(string))
		}
		if recordType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "record_type")); ok {
			details.RecordType = oci_apm_synthetics.DnsRecordTypeEnum(recordType.(string))
		}
		if verifyResponseContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_content")); ok {
			tmp := verifyResponseContent.(string)
			details.VerifyResponseContent = &tmp
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("NETWORK_CONFIG"):
		details := oci_apm_synthetics.NetworkMonitorConfiguration{}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("REST_CONFIG"):
		details := oci_apm_synthetics.RestMonitorConfiguration{}
		if clientCertificateDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_certificate_details")); ok {
			if tmpList := clientCertificateDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "client_certificate_details"), 0)
				tmp, err := s.mapToClientCertificateDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert client_certificate_details, encountered error: %v", err)
				}
				details.ClientCertificateDetails = &tmp
			}
		}
		if isCertificateValidationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_certificate_validation_enabled")); ok {
			tmp := isCertificateValidationEnabled.(bool)
			details.IsCertificateValidationEnabled = &tmp
		}
		if isRedirectionEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_redirection_enabled")); ok {
			tmp := isRedirectionEnabled.(bool)
			details.IsRedirectionEnabled = &tmp
		}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if reqAuthenticationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "req_authentication_details")); ok {
			if tmpList := reqAuthenticationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "req_authentication_details"), 0)
				tmp, err := s.mapToRequestAuthenticationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert req_authentication_details, encountered error: %v", err)
				}
				details.ReqAuthenticationDetails = &tmp
			}
		}
		if reqAuthenticationScheme, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "req_authentication_scheme")); ok {
			details.ReqAuthenticationScheme = oci_apm_synthetics.RequestAuthenticationSchemesEnum(reqAuthenticationScheme.(string))
		}
		if requestHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_headers")); ok {
			interfaces := requestHeaders.([]interface{})
			tmp := make([]oci_apm_synthetics.Header, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "request_headers"), stateDataIndex)
				converted, err := s.mapToHeader(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "request_headers")) {
				details.RequestHeaders = tmp
			}
		}
		if requestMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_method")); ok {
			details.RequestMethod = oci_apm_synthetics.RequestMethodsEnum(requestMethod.(string))
		}
		if requestPostBody, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_post_body")); ok {
			tmp := requestPostBody.(string)
			details.RequestPostBody = &tmp
		}
		if requestQueryParams, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_query_params")); ok {
			interfaces := requestQueryParams.([]interface{})
			tmp := make([]oci_apm_synthetics.RequestQueryParam, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "request_query_params"), stateDataIndex)
				converted, err := s.mapToRequestQueryParam(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "request_query_params")) {
				details.RequestQueryParams = tmp
			}
		}
		if verifyResponseCodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")); ok {
			interfaces := verifyResponseCodes.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")) {
				details.VerifyResponseCodes = tmp
			}
		}
		if verifyResponseContent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_content")); ok {
			tmp := verifyResponseContent.(string)
			details.VerifyResponseContent = &tmp
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("SCRIPTED_BROWSER_CONFIG"):
		details := oci_apm_synthetics.ScriptedBrowserMonitorConfiguration{}
		if isCertificateValidationEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_certificate_validation_enabled")); ok {
			tmp := isCertificateValidationEnabled.(bool)
			details.IsCertificateValidationEnabled = &tmp
		}
		if isDefaultSnapshotEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_default_snapshot_enabled")); ok {
			tmp := isDefaultSnapshotEnabled.(bool)
			details.IsDefaultSnapshotEnabled = &tmp
		}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	case strings.ToLower("SCRIPTED_REST_CONFIG"):
		details := oci_apm_synthetics.ScriptedRestMonitorConfiguration{}
		if networkConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_configuration")); ok {
			if tmpList := networkConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_configuration"), 0)
				tmp, err := s.mapToNetworkConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_configuration, encountered error: %v", err)
				}
				details.NetworkConfiguration = &tmp
			}
		}
		if reqAuthenticationScheme, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "req_authentication_scheme")); ok {
			details.ReqAuthenticationScheme = oci_apm_synthetics.RequestAuthenticationSchemesForScriptedRestEnum(reqAuthenticationScheme.(string))
		}
		if verifyResponseCodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")); ok {
			interfaces := verifyResponseCodes.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "verify_response_codes")) {
				details.VerifyResponseCodes = tmp
			}
		}
		if dnsConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns_configuration")); ok {
			if tmpList := dnsConfiguration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns_configuration"), 0)
				tmp, err := s.mapToDnsConfiguration(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert dns_configuration, encountered error: %v", err)
				}
				details.DnsConfiguration = &tmp
			}
		}
		if isFailureRetried, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_failure_retried")); ok {
			tmp := isFailureRetried.(bool)
			details.IsFailureRetried = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown config_type '%v' was specified", configType)
	}
	return baseObject, nil
}

func MonitorConfigurationToMap(obj *oci_apm_synthetics.MonitorConfiguration) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apm_synthetics.BrowserMonitorConfiguration:
		result["config_type"] = "BROWSER_CONFIG"

		if v.IsCertificateValidationEnabled != nil {
			result["is_certificate_validation_enabled"] = bool(*v.IsCertificateValidationEnabled)
		}

		if v.IsDefaultSnapshotEnabled != nil {
			result["is_default_snapshot_enabled"] = bool(*v.IsDefaultSnapshotEnabled)
		}

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		result["verify_response_codes"] = v.VerifyResponseCodes
		result["verify_response_codes"] = v.VerifyResponseCodes

		verifyTexts := []interface{}{}
		for _, item := range v.VerifyTexts {
			verifyTexts = append(verifyTexts, VerifyTextToMap(item))
		}
		result["verify_texts"] = verifyTexts

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.DnsSecMonitorConfiguration:
		result["config_type"] = "DNSSEC_CONFIG"

		result["record_type"] = string(v.RecordType)

		if v.VerifyResponseContent != nil {
			result["verify_response_content"] = string(*v.VerifyResponseContent)
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.DnsServerMonitorConfiguration:
		result["config_type"] = "DNS_SERVER_CONFIG"

		if v.IsQueryRecursive != nil {
			result["is_query_recursive"] = bool(*v.IsQueryRecursive)
		}

		if v.NameServer != nil {
			result["name_server"] = string(*v.NameServer)
		}

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		result["protocol"] = string(v.Protocol)

		result["record_type"] = string(v.RecordType)

		if v.VerifyResponseContent != nil {
			result["verify_response_content"] = string(*v.VerifyResponseContent)
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.DnsTraceMonitorConfiguration:
		result["config_type"] = "DNS_TRACE_CONFIG"

		result["protocol"] = string(v.Protocol)

		result["record_type"] = string(v.RecordType)

		if v.VerifyResponseContent != nil {
			result["verify_response_content"] = string(*v.VerifyResponseContent)
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.NetworkMonitorConfiguration:
		result["config_type"] = "NETWORK_CONFIG"

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.RestMonitorConfiguration:
		result["config_type"] = "REST_CONFIG"

		if v.ClientCertificateDetails != nil {
			result["client_certificate_details"] = []interface{}{ClientCertificateDetailsToMap(v.ClientCertificateDetails)}
		}

		if v.IsCertificateValidationEnabled != nil {
			result["is_certificate_validation_enabled"] = bool(*v.IsCertificateValidationEnabled)
		}

		if v.IsRedirectionEnabled != nil {
			result["is_redirection_enabled"] = bool(*v.IsRedirectionEnabled)
		}

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		if v.ReqAuthenticationDetails != nil {
			result["req_authentication_details"] = []interface{}{RequestAuthenticationDetailsToMap(v.ReqAuthenticationDetails)}
		}

		result["req_authentication_scheme"] = string(v.ReqAuthenticationScheme)

		requestHeaders := []interface{}{}
		for _, item := range v.RequestHeaders {
			requestHeaders = append(requestHeaders, SyntheticHeaderToMap(item))
		}
		result["request_headers"] = requestHeaders

		result["request_method"] = string(v.RequestMethod)

		if v.RequestPostBody != nil {
			result["request_post_body"] = string(*v.RequestPostBody)
		}

		requestQueryParams := []interface{}{}
		for _, item := range v.RequestQueryParams {
			requestQueryParams = append(requestQueryParams, RequestQueryParamToMap(item))
		}
		result["request_query_params"] = requestQueryParams

		result["verify_response_codes"] = v.VerifyResponseCodes

		if v.VerifyResponseContent != nil {
			result["verify_response_content"] = string(*v.VerifyResponseContent)
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.ScriptedBrowserMonitorConfiguration:
		result["config_type"] = "SCRIPTED_BROWSER_CONFIG"

		if v.IsCertificateValidationEnabled != nil {
			result["is_certificate_validation_enabled"] = bool(*v.IsCertificateValidationEnabled)
		}

		if v.IsDefaultSnapshotEnabled != nil {
			result["is_default_snapshot_enabled"] = bool(*v.IsDefaultSnapshotEnabled)
		}

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	case oci_apm_synthetics.ScriptedRestMonitorConfiguration:
		result["config_type"] = "SCRIPTED_REST_CONFIG"

		if v.NetworkConfiguration != nil {
			result["network_configuration"] = []interface{}{NetworkConfigurationToMap(v.NetworkConfiguration)}
		}

		result["req_authentication_scheme"] = string(v.ReqAuthenticationScheme)

		result["verify_response_codes"] = v.VerifyResponseCodes
		result["verify_response_codes"] = v.VerifyResponseCodes

		if v.DnsConfiguration != nil {
			result["dns_configuration"] = []interface{}{DnsConfigurationToMap(v.DnsConfiguration)}
		}

		if v.IsFailureRetried != nil {
			result["is_failure_retried"] = bool(*v.IsFailureRetried)
		}
	default:
		log.Printf("[WARN] Received 'config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func MonitorScriptParameterToMap(obj *oci_apm_synthetics.MonitorScriptParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ParamName != nil {
		result["param_name"] = string(*obj.ParamName)
	}

	if obj.ParamValue != nil {
		result["param_value"] = string(*obj.ParamValue)
	}

	return result
}

func MonitorSummaryToMap(obj oci_apm_synthetics.MonitorSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BatchIntervalInSeconds != nil {
		result["batch_interval_in_seconds"] = int(*obj.BatchIntervalInSeconds)
	}

	if obj.Configuration != nil {
		configurationArray := []interface{}{}
		if configurationMap := MonitorConfigurationToMap(&obj.Configuration); configurationMap != nil {
			configurationArray = append(configurationArray, configurationMap)
		}
		result["configuration"] = configurationArray
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

	if obj.IsRunNow != nil {
		result["is_run_now"] = bool(*obj.IsRunNow)
	}

	if obj.IsRunOnce != nil {
		result["is_run_once"] = bool(*obj.IsRunOnce)
	}

	if obj.MaintenanceWindowSchedule != nil {
		result["maintenance_window_schedule"] = []interface{}{MaintenanceWindowScheduleToMap(obj.MaintenanceWindowSchedule)}
	}

	result["monitor_type"] = string(obj.MonitorType)

	if obj.RepeatIntervalInSeconds != nil {
		result["repeat_interval_in_seconds"] = int(*obj.RepeatIntervalInSeconds)
	}

	result["scheduling_policy"] = string(obj.SchedulingPolicy)

	if obj.ScriptId != nil {
		result["script_id"] = string(*obj.ScriptId)
	}

	if obj.ScriptName != nil {
		result["script_name"] = string(*obj.ScriptName)
	}

	result["status"] = string(obj.Status)

	if obj.Target != nil {
		result["target"] = string(*obj.Target)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.TimeoutInSeconds != nil {
		result["timeout_in_seconds"] = int(*obj.TimeoutInSeconds)
	}

	if obj.VantagePointCount != nil {
		result["vantage_point_count"] = int(*obj.VantagePointCount)
	}

	vantagePoints := []interface{}{}
	for _, item := range obj.VantagePoints {
		vantagePoints = append(vantagePoints, VantagePointInfoToMap(item))
	}
	result["vantage_points"] = vantagePoints

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToNetworkConfiguration(fieldKeyFormat string) (oci_apm_synthetics.NetworkConfiguration, error) {
	result := oci_apm_synthetics.NetworkConfiguration{}

	if numberOfHops, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "number_of_hops")); ok {
		tmp := numberOfHops.(int)
		result.NumberOfHops = &tmp
	}

	if probeMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "probe_mode")); ok {
		result.ProbeMode = oci_apm_synthetics.ProbeModeEnum(probeMode.(string))
	}

	if probePerHop, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "probe_per_hop")); ok {
		tmp := probePerHop.(int)
		result.ProbePerHop = &tmp
	}

	if protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol")); ok {
		result.Protocol = oci_apm_synthetics.ProtocolEnum(protocol.(string))
	}

	if transmissionRate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transmission_rate")); ok {
		tmp := transmissionRate.(int)
		result.TransmissionRate = &tmp
	}

	return result, nil
}

func NetworkConfigurationToMap(obj *oci_apm_synthetics.NetworkConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.NumberOfHops != nil {
		result["number_of_hops"] = int(*obj.NumberOfHops)
	}

	result["probe_mode"] = string(obj.ProbeMode)

	if obj.ProbePerHop != nil {
		result["probe_per_hop"] = int(*obj.ProbePerHop)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.TransmissionRate != nil {
		result["transmission_rate"] = int(*obj.TransmissionRate)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToPrivateKey(fieldKeyFormat string) (oci_apm_synthetics.PrivateKey, error) {
	result := oci_apm_synthetics.PrivateKey{}

	if content, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "content")); ok {
		tmp := content.(string)
		result.Content = &tmp
	}

	if fileName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_name")); ok {
		tmp := fileName.(string)
		result.FileName = &tmp
	}

	return result, nil
}

func PrivateKeyToMap(obj *oci_apm_synthetics.PrivateKey) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Content != nil {
		result["content"] = string(*obj.Content)
	}

	if obj.FileName != nil {
		result["file_name"] = string(*obj.FileName)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToRequestAuthenticationDetails(fieldKeyFormat string) (oci_apm_synthetics.RequestAuthenticationDetails, error) {
	result := oci_apm_synthetics.RequestAuthenticationDetails{}

	if authHeaders, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_headers")); ok {
		interfaces := authHeaders.([]interface{})
		tmp := make([]oci_apm_synthetics.Header, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "auth_headers"), stateDataIndex)
			converted, err := s.mapToHeader(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "auth_headers")) {
			result.AuthHeaders = tmp
		}
	}

	if authRequestMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_request_method")); ok {
		result.AuthRequestMethod = oci_apm_synthetics.RequestMethodsEnum(authRequestMethod.(string))
	}

	if authRequestPostBody, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_request_post_body")); ok {
		tmp := authRequestPostBody.(string)
		result.AuthRequestPostBody = &tmp
	}

	if authToken, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_token")); ok {
		tmp := authToken.(string)
		result.AuthToken = &tmp
	}

	if authUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_url")); ok {
		tmp := authUrl.(string)
		result.AuthUrl = &tmp
	}

	if authUserName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_user_name")); ok {
		tmp := authUserName.(string)
		result.AuthUserName = &tmp
	}

	if authUserPassword, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_user_password")); ok {
		tmp := authUserPassword.(string)
		result.AuthUserPassword = &tmp
	}

	if oauthScheme, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "oauth_scheme")); ok {
		result.OauthScheme = oci_apm_synthetics.OAuthSchemesEnum(oauthScheme.(string))
	}

	return result, nil
}

func RequestAuthenticationDetailsToMap(obj *oci_apm_synthetics.RequestAuthenticationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	authHeaders := []interface{}{}
	for _, item := range obj.AuthHeaders {
		authHeaders = append(authHeaders, SyntheticHeaderToMap(item))
	}
	result["auth_headers"] = authHeaders

	result["auth_request_method"] = string(obj.AuthRequestMethod)

	if obj.AuthRequestPostBody != nil {
		result["auth_request_post_body"] = string(*obj.AuthRequestPostBody)
	}

	if obj.AuthToken != nil {
		result["auth_token"] = string(*obj.AuthToken)
	}

	if obj.AuthUrl != nil {
		result["auth_url"] = string(*obj.AuthUrl)
	}

	if obj.AuthUserName != nil {
		result["auth_user_name"] = string(*obj.AuthUserName)
	}

	if obj.AuthUserPassword != nil {
		result["auth_user_password"] = string(*obj.AuthUserPassword)
	}

	result["oauth_scheme"] = string(obj.OauthScheme)

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToRequestQueryParam(fieldKeyFormat string) (oci_apm_synthetics.RequestQueryParam, error) {
	result := oci_apm_synthetics.RequestQueryParam{}

	if paramName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_name")); ok {
		tmp := paramName.(string)
		result.ParamName = &tmp
	}

	if paramValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_value")); ok {
		tmp := paramValue.(string)
		result.ParamValue = &tmp
	}

	return result, nil
}

func RequestQueryParamToMap(obj oci_apm_synthetics.RequestQueryParam) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ParamName != nil {
		result["param_name"] = string(*obj.ParamName)
	}

	if obj.ParamValue != nil {
		result["param_value"] = string(*obj.ParamValue)
	}

	return result
}

func VantagePointInfoToMap(obj oci_apm_synthetics.VantagePointInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToVerifyText(fieldKeyFormat string) (oci_apm_synthetics.VerifyText, error) {
	result := oci_apm_synthetics.VerifyText{}

	if text, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "text")); ok {
		tmp := text.(string)
		result.Text = &tmp
	}

	return result, nil
}

func VerifyTextToMap(obj oci_apm_synthetics.VerifyText) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Text != nil {
		result["text"] = string(*obj.Text)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToMonitorScriptParameter(fieldKeyFormat string) (oci_apm_synthetics.MonitorScriptParameter, error) {
	result := oci_apm_synthetics.MonitorScriptParameter{}

	if paramName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_name")); ok {
		tmp := paramName.(string)
		result.ParamName = &tmp
	}

	if paramValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "param_value")); ok {
		tmp := paramValue.(string)
		result.ParamValue = &tmp
	}

	return result, nil
}

func MonitorScriptParameterInfoToMap(obj oci_apm_synthetics.MonitorScriptParameterInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MonitorScriptParameter != nil {
		MonitorScriptParameters := []interface{}{}
		MonitorScriptParameters = append(MonitorScriptParameters, MonitorScriptParameterToMap(obj.MonitorScriptParameter))

		result["monitor_script_parameter"] = MonitorScriptParameters
		result["param_name"] = string(*obj.MonitorScriptParameter.ParamName)
		result["param_value"] = string(*obj.MonitorScriptParameter.ParamValue)
	}

	if obj.IsSecret != nil {
		result["is_secret"] = bool(*obj.IsSecret)
	}

	if obj.IsOverwritten != nil {
		result["is_overwritten"] = bool(*obj.IsOverwritten)
	}

	return result
}

func (s *ApmSyntheticsMonitorResourceCrud) mapToVantagePointInfo(fieldKeyFormat string) (oci_apm_synthetics.VantagePointInfo, error) {
	result := oci_apm_synthetics.VantagePointInfo{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}
