// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePipelineRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatasciencePipelineRun,
		Read:     readDatasciencePipelineRun,
		Update:   updateDatasciencePipelineRun,
		Delete:   deleteDatasciencePipelineRun,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"configuration_override_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
								"DEFAULT",
							}, true),
						},

						// Optional
						"command_line_arguments": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"environment_variables": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem:     schema.TypeString,
						},
						"maximum_runtime_in_minutes": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
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
			"infrastructure_configuration_override_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Required: true,
							ForceNew: true,
						},
						"shape_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"block_storage_size_in_gbs_parameterized": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"shape_config_details": {
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
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"memory_in_gbs_parameterized": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"ocpus_parameterized": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"cpu_baseline": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"delete_related_job_runs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"log_configuration_override_details": {
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
						"enable_auto_log_creation": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"enable_logging": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"log_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"opc_parent_rpt_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parameters_override": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},
			"step_override_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"step_configuration_details": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"command_line_arguments": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"environment_variables": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem:     schema.TypeString,
									},
									"maximum_runtime_in_minutes": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ForceNew:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
									},

									// Computed
								},
							},
						},
						"step_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"step_container_configuration_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"container_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"OCIR_CONTAINER",
										}, true),
									},
									"image": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"cmd": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"entrypoint": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"image_digest": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"image_signature_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"step_dataflow_configuration_details": {
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
									"configuration": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										Elem:             schema.TypeString,
										DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
									},
									"driver_shape": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"driver_shape_config_details": {
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
												"memory_in_gbs": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"memory_in_gbs_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"executor_shape": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"executor_shape_config_details": {
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
												"memory_in_gbs": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"memory_in_gbs_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"logs_bucket_uri": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"num_executors": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"warehouse_bucket_uri": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"step_infrastructure_configuration_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"block_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: true,
									},
									"shape_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},

									// Optional
									"block_storage_size_in_gbs_parameterized": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"shape_config_details": {
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
												"memory_in_gbs": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"memory_in_gbs_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"ocpus_parameterized": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},
												"cpu_baseline": {
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"step_storage_mount_configuration_details_list": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"destination_directory_name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"storage_type": {
										Type:             schema.TypeString,
										Required:         true,
										ForceNew:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"FILE_STORAGE",
											"OBJECT_STORAGE",
										}, true),
									},

									// Optional
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"destination_path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"export_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"mount_target_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
									"prefix": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},

						// Computed
					},
				},
			},
			"storage_mount_configuration_override_details_list": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"destination_directory_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"storage_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"FILE_STORAGE",
								"OBJECT_STORAGE",
							}, true),
						},

						// Optional
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"destination_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"export_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"mount_target_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"configuration_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"command_line_arguments": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"environment_variables": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"maximum_runtime_in_minutes": {
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
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"log_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"step_runs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"dataflow_run_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_run_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_run_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_finished": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_finished": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
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

func createDatasciencePipelineRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatasciencePipelineRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatasciencePipelineRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatasciencePipelineRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatasciencePipelineRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.PipelineRun
	DisableNotFoundRetries bool
}

func (s *DatasciencePipelineRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatasciencePipelineRunResourceCrud) CreatedPending() []string {
	// invokeAsynchronously := true
	// if async, ok := s.D.GetOkExists("asynchronous"); ok {
	// 	invokeAsynchronously = async.(bool)
	// }

	// if invokeAsynchronously {
	// 	return []string{} // nothing
	// }

	return []string{
		string(oci_datascience.PipelineRunLifecycleStateAccepted),
		string(oci_datascience.PipelineRunLifecycleStateInProgress),
	}
}

func (s *DatasciencePipelineRunResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.PipelineRunLifecycleStateSucceeded),
		string(oci_datascience.PipelineRunLifecycleStateFailed),
		string(oci_datascience.PipelineRunLifecycleStateCanceled),
	}
}

func (s *DatasciencePipelineRunResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.PipelineRunLifecycleStateDeleting),
	}
}

func (s *DatasciencePipelineRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.PipelineRunLifecycleStateDeleted),
	}
}

func (s *DatasciencePipelineRunResourceCrud) Create() error {
	request := oci_datascience.CreatePipelineRunRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationOverrideDetails, ok := s.D.GetOkExists("configuration_override_details"); ok {
		if tmpList := configurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_override_details", 0)
			tmp, err := s.mapToPipelineConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationOverrideDetails = tmp
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

	if infrastructureConfigurationOverrideDetails, ok := s.D.GetOkExists("infrastructure_configuration_override_details"); ok {
		if tmpList := infrastructureConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "infrastructure_configuration_override_details", 0)
			tmp, err := s.mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InfrastructureConfigurationOverrideDetails = &tmp
		}
	}

	if logConfigurationOverrideDetails, ok := s.D.GetOkExists("log_configuration_override_details"); ok {
		if tmpList := logConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_configuration_override_details", 0)
			tmp, err := s.mapToPipelineLogConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogConfigurationOverrideDetails = &tmp
		}
	}

	if opcParentRptUrl, ok := s.D.GetOkExists("opc_parent_rpt_url"); ok {
		tmp := opcParentRptUrl.(string)
		log.Printf("[INFO] Received opc_parent_rpt_url with value:%s", opcParentRptUrl)
		request.OpcParentRptUrl = &tmp
	}

	if parametersOverride, ok := s.D.GetOkExists("parameters_override"); ok {
		request.ParametersOverride = tfresource.ObjectMapToStringMap(parametersOverride.(map[string]interface{}))
	}

	if pipelineId, ok := s.D.GetOkExists("pipeline_id"); ok {
		tmp := pipelineId.(string)
		request.PipelineId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if stepOverrideDetails, ok := s.D.GetOkExists("step_override_details"); ok {
		interfaces := stepOverrideDetails.([]interface{})
		tmp := make([]oci_datascience.PipelineStepOverrideDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "step_override_details", stateDataIndex)
			converted, err := s.mapToPipelineStepOverrideDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("step_override_details") {
			request.StepOverrideDetails = tmp
		}
	}

	if storageMountConfigurationOverrideDetailsList, ok := s.D.GetOkExists("storage_mount_configuration_override_details_list"); ok {
		interfaces := storageMountConfigurationOverrideDetailsList.([]interface{})
		tmp := make([]oci_datascience.StorageMountConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage_mount_configuration_override_details_list", stateDataIndex)
			converted, err := s.mapToStorageMountConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("storage_mount_configuration_override_details_list") {
			request.StorageMountConfigurationOverrideDetailsList = tmp
		}
	}

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreatePipelineRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PipelineRun
	return nil
}

func (s *DatasciencePipelineRunResourceCrud) Get() error {
	request := oci_datascience.GetPipelineRunRequest{}

	tmp := s.D.Id()
	request.PipelineRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetPipelineRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PipelineRun
	return nil
}

func (s *DatasciencePipelineRunResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdatePipelineRunRequest{}

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

	tmp := s.D.Id()
	request.PipelineRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdatePipelineRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.PipelineRun
	return nil
}

func (s *DatasciencePipelineRunResourceCrud) Delete() error {
	request := oci_datascience.DeletePipelineRunRequest{}

	if deleteRelatedJobRuns, ok := s.D.GetOkExists("delete_related_job_runs"); ok {
		tmp := deleteRelatedJobRuns.(bool)
		request.DeleteRelatedJobRuns = &tmp
	}

	tmp := s.D.Id()
	request.PipelineRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeletePipelineRun(context.Background(), request)
	return err
}

func (s *DatasciencePipelineRunResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConfigurationDetails != nil {
		configurationDetailsArray := []interface{}{}
		if configurationDetailsMap := PipelineConfigurationDetailsToMap(&s.Res.ConfigurationDetails); configurationDetailsMap != nil {
			configurationDetailsArray = append(configurationDetailsArray, configurationDetailsMap)
		}
		s.D.Set("configuration_details", configurationDetailsArray)
	} else {
		s.D.Set("configuration_details", nil)
	}

	if s.Res.ConfigurationOverrideDetails != nil {
		configurationOverrideDetailsArray := []interface{}{}
		if configurationOverrideDetailsMap := PipelineConfigurationDetailsToMap(&s.Res.ConfigurationOverrideDetails); configurationOverrideDetailsMap != nil {
			configurationOverrideDetailsArray = append(configurationOverrideDetailsArray, configurationOverrideDetailsMap)
		}
		s.D.Set("configuration_override_details", configurationOverrideDetailsArray)
	} else {
		s.D.Set("configuration_override_details", nil)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InfrastructureConfigurationOverrideDetails != nil {
		s.D.Set("infrastructure_configuration_override_details", []interface{}{PipelineInfrastructureConfigurationDetailsToMap(s.Res.InfrastructureConfigurationOverrideDetails)})
	} else {
		s.D.Set("infrastructure_configuration_override_details", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogConfigurationOverrideDetails != nil {
		s.D.Set("log_configuration_override_details", []interface{}{PipelineLogConfigurationDetailsToMap(s.Res.LogConfigurationOverrideDetails)})
	} else {
		s.D.Set("log_configuration_override_details", nil)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{PipelineRunLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
	}

	s.D.Set("parameters_override", s.Res.ParametersOverride)

	if s.Res.PipelineId != nil {
		s.D.Set("pipeline_id", *s.Res.PipelineId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	stepOverrideDetails := []interface{}{}
	for _, item := range s.Res.StepOverrideDetails {
		stepOverrideDetails = append(stepOverrideDetails, PipelineStepOverrideDetailsToMap(item))
	}
	s.D.Set("step_override_details", stepOverrideDetails)

	stepRuns := []interface{}{}
	for _, item := range s.Res.StepRuns {
		stepRuns = append(stepRuns, PipelineStepRunToMap(item))
	}
	s.D.Set("step_runs", stepRuns)

	storageMountConfigurationOverrideDetailsList := []interface{}{}
	for _, item := range s.Res.StorageMountConfigurationOverrideDetailsList {
		storageMountConfigurationOverrideDetailsList = append(storageMountConfigurationOverrideDetailsList, StorageMountConfigurationDetailsToMap(item))
	}
	s.D.Set("storage_mount_configuration_override_details_list", storageMountConfigurationOverrideDetailsList)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineConfigurationDetails, error) {
	var baseObject oci_datascience.PipelineConfigurationDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.PipelineDefaultConfigurationDetails{}
		if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
			tmp := commandLineArguments.(string)
			details.CommandLineArguments = &tmp
		}
		if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
			details.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
		}
		if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
			tmp := maximumRuntimeInMinutes.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaximumRuntimeInMinutes = &tmpInt64
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func PipelineConfigurationDetailsToMap(obj *oci_datascience.PipelineConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_datascience.PipelineDefaultConfigurationDetails:
		result["type"] = "DEFAULT"

		if v.CommandLineArguments != nil {
			result["command_line_arguments"] = string(*v.CommandLineArguments)
		}

		result["environment_variables"] = v.EnvironmentVariables
		result["environment_variables"] = v.EnvironmentVariables

		if v.MaximumRuntimeInMinutes != nil {
			result["maximum_runtime_in_minutes"] = strconv.FormatInt(*v.MaximumRuntimeInMinutes, 10)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineStepContainerConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineContainerConfigurationDetails, error) {
	var baseObject oci_datascience.PipelineContainerConfigurationDetails
	//discriminator
	containerTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_type"))
	var containerType string
	if ok {
		containerType = containerTypeRaw.(string)
	} else {
		containerType = "" // default value
	}
	switch strings.ToLower(containerType) {
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.PipelineOcirContainerConfigurationDetails{}
		if cmd, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cmd")); ok {
			interfaces := cmd.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "cmd")) {
				details.Cmd = tmp
			}
		}
		if entrypoint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "entrypoint")); ok {
			interfaces := entrypoint.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "entrypoint")) {
				details.Entrypoint = tmp
			}
		}
		if image, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image")); ok {
			tmp := image.(string)
			details.Image = &tmp
		}
		if imageDigest, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_digest")); ok {
			tmp := imageDigest.(string)
			details.ImageDigest = &tmp
		}
		if imageSignatureId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_signature_id")); ok {
			tmp := imageSignatureId.(string)
			details.ImageSignatureId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown container_type '%v' was specified", containerType)
	}
	return baseObject, nil
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineDataflowConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineDataflowConfigurationDetails, error) {
	result := oci_datascience.PipelineDataflowConfigurationDetails{}

	if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
		tmp := configuration.(map[string]interface{})
		pointerToMap := &tmp
		// Wrap the map pointer in an interface{}
		var i interface{} = pointerToMap

		// Take the address of the interface{}
		result.Configuration = &i
	}

	if driverShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape")); ok {
		tmp := driverShape.(string)
		result.DriverShape = &tmp
	}

	if driverShapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "driver_shape_config_details")); ok {
		if tmpList := driverShapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "driver_shape_config_details"), 0)
			tmp, err := s.mapToPipelineShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert driver_shape_config_details, encountered error: %v", err)
			}
			result.DriverShapeConfigDetails = &tmp
		}
	}

	if executorShape, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape")); ok {
		tmp := executorShape.(string)
		result.ExecutorShape = &tmp
	}

	if executorShapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "executor_shape_config_details")); ok {
		if tmpList := executorShapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "executor_shape_config_details"), 0)
			tmp, err := s.mapToPipelineShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert executor_shape_config_details, encountered error: %v", err)
			}
			result.ExecutorShapeConfigDetails = &tmp
		}
	}

	if logsBucketUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logs_bucket_uri")); ok {
		tmp := logsBucketUri.(string)
		result.LogsBucketUri = &tmp
	}

	if numExecutors, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "num_executors")); ok {
		tmp := numExecutors.(int)
		result.NumExecutors = &tmp
	}

	if warehouseBucketUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "warehouse_bucket_uri")); ok {
		tmp := warehouseBucketUri.(string)
		result.WarehouseBucketUri = &tmp
	}

	return result, nil
}

// func PipelineDataflowConfigurationDetailsToMap(obj *oci_datascience.PipelineDataflowConfigurationDetails) map[string]interface{} {
// 	result := map[string]interface{}{}

// 	if obj.Configuration != nil {
// 		result["configuration"] = []interface{}{objectToMap(obj.Configuration)}
// 	}

// 	if obj.DriverShape != nil {
// 		result["driver_shape"] = string(*obj.DriverShape)
// 	}

// 	if obj.DriverShapeConfigDetails != nil {
// 		result["driver_shape_config_details"] = []interface{}{PipelineShapeConfigDetailsToMap(obj.DriverShapeConfigDetails)}
// 	}

// 	if obj.ExecutorShape != nil {
// 		result["executor_shape"] = string(*obj.ExecutorShape)
// 	}

// 	if obj.ExecutorShapeConfigDetails != nil {
// 		result["executor_shape_config_details"] = []interface{}{PipelineShapeConfigDetailsToMap(obj.ExecutorShapeConfigDetails)}
// 	}

// 	if obj.LogsBucketUri != nil {
// 		result["logs_bucket_uri"] = string(*obj.LogsBucketUri)
// 	}

// 	if obj.NumExecutors != nil {
// 		result["num_executors"] = int(*obj.NumExecutors)
// 	}

// 	if obj.WarehouseBucketUri != nil {
// 		result["warehouse_bucket_uri"] = string(*obj.WarehouseBucketUri)
// 	}

// 	return result
// }

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineInfrastructureConfigurationDetails, error) {
	result := oci_datascience.PipelineInfrastructureConfigurationDetails{}

	if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
		tmp := blockStorageSizeInGBs.(int)
		result.BlockStorageSizeInGBs = &tmp
	}

	if blockStorageSizeInGBsParameterized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs_parameterized")); ok {
		tmp := blockStorageSizeInGBsParameterized.(string)
		result.BlockStorageSizeInGBsParameterized = &tmp
	}

	if shapeConfigDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config_details")); ok {
		if tmpList := shapeConfigDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config_details"), 0)
			tmp, err := s.mapToPipelineShapeConfigDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert shape_config_details, encountered error: %v", err)
			}
			result.ShapeConfigDetails = &tmp
		}
	}

	if shapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_name")); ok {
		tmp := shapeName.(string)
		result.ShapeName = &tmp
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineLogConfigurationDetails, error) {
	result := oci_datascience.PipelineLogConfigurationDetails{}

	if enableAutoLogCreation, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_auto_log_creation")); ok {
		tmp := enableAutoLogCreation.(bool)
		result.EnableAutoLogCreation = &tmp
	}

	if enableLogging, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_logging")); ok {
		tmp := enableLogging.(bool)
		result.EnableLogging = &tmp
	}

	if logGroupId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_group_id")); ok {
		tmp := logGroupId.(string)
		result.LogGroupId = &tmp
	}

	if logId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "log_id")); ok {
		tmp := logId.(string)
		result.LogId = &tmp
	}

	return result, nil
}

func PipelineLogConfigurationDetailsToMap(obj *oci_datascience.PipelineLogConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnableAutoLogCreation != nil {
		result["enable_auto_log_creation"] = bool(*obj.EnableAutoLogCreation)
	}

	if obj.EnableLogging != nil {
		result["enable_logging"] = bool(*obj.EnableLogging)
	}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func PipelineRunLogDetailsToMap(obj *oci_datascience.PipelineRunLogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineShapeConfigDetails(fieldKeyFormat string) (oci_datascience.PipelineShapeConfigDetails, error) {
	result := oci_datascience.PipelineShapeConfigDetails{}

	if cpuBaseline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cpu_baseline")); ok {
		result.CpuBaseline = oci_datascience.PipelineShapeConfigDetailsCpuBaselineEnum(cpuBaseline.(string))
	}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if memoryInGBsParameterized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs_parameterized")); ok {
		tmp := memoryInGBsParameterized.(string)
		result.MemoryInGBsParameterized = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	if ocpusParameterized, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus_parameterized")); ok {
		tmp := ocpusParameterized.(string)
		result.OcpusParameterized = &tmp
	}

	return result, nil
}

func PipelineShapeConfigDetailsToMap(obj *oci_datascience.PipelineShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["cpu_baseline"] = string(obj.CpuBaseline)

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.MemoryInGBsParameterized != nil {
		result["memory_in_gbs_parameterized"] = string(*obj.MemoryInGBsParameterized)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	if obj.OcpusParameterized != nil {
		result["ocpus_parameterized"] = string(*obj.OcpusParameterized)
	}

	return result
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineStepConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineStepConfigurationDetails, error) {
	result := oci_datascience.PipelineStepConfigurationDetails{}

	if commandLineArguments, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "command_line_arguments")); ok {
		tmp := commandLineArguments.(string)
		result.CommandLineArguments = &tmp
	}

	if environmentVariables, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "environment_variables")); ok {
		result.EnvironmentVariables = tfresource.ObjectMapToStringMap(environmentVariables.(map[string]interface{}))
	}

	if maximumRuntimeInMinutes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "maximum_runtime_in_minutes")); ok {
		tmp := maximumRuntimeInMinutes.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert maximumRuntimeInMinutes string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.MaximumRuntimeInMinutes = &tmpInt64
	}

	return result, nil
}

func PipelineStepConfigurationDetailsToMap(obj *oci_datascience.PipelineStepConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommandLineArguments != nil {
		result["command_line_arguments"] = string(*obj.CommandLineArguments)
	}

	result["environment_variables"] = obj.EnvironmentVariables

	if obj.MaximumRuntimeInMinutes != nil {
		result["maximum_runtime_in_minutes"] = strconv.FormatInt(*obj.MaximumRuntimeInMinutes, 10)
	}

	return result
}

func (s *DatasciencePipelineRunResourceCrud) mapToPipelineStepOverrideDetails(fieldKeyFormat string) (oci_datascience.PipelineStepOverrideDetails, error) {
	result := oci_datascience.PipelineStepOverrideDetails{}

	if stepContainerConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_container_configuration_details")); ok {
		if tmpList := stepContainerConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_container_configuration_details"), 0)
			tmp, err := s.mapToPipelineStepContainerConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_container_configuration_details, encountered error: %v", err)
			}
			result.StepContainerConfigurationDetails = tmp
		}
	}

	if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
		if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
			tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
			}
			result.StepConfigurationDetails = &tmp
		}
	}

	if stepDataflowConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_dataflow_configuration_details")); ok {
		if tmpList := stepDataflowConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_dataflow_configuration_details"), 0)
			tmp, err := s.mapToPipelineDataflowConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_dataflow_configuration_details, encountered error: %v", err)
			}
			result.StepDataflowConfigurationDetails = &tmp
		}
	}

	if stepInfrastructureConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_infrastructure_configuration_details")); ok {
		if tmpList := stepInfrastructureConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_infrastructure_configuration_details"), 0)
			tmp, err := s.mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert step_infrastructure_configuration_details, encountered error: %v", err)
			}
			result.StepInfrastructureConfigurationDetails = &tmp
		}
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		result.StepName = &tmp
	}

	if stepStorageMountConfigurationDetailsList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_storage_mount_configuration_details_list")); ok {
		interfaces := stepStorageMountConfigurationDetailsList.([]interface{})
		tmp := make([]oci_datascience.StorageMountConfigurationDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_storage_mount_configuration_details_list"), stateDataIndex)
			converted, err := s.mapToStorageMountConfigurationDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "step_storage_mount_configuration_details_list")) {
			result.StepStorageMountConfigurationDetailsList = tmp
		}
	}

	return result, nil
}

func PipelineStepOverrideDetailsToMap(obj oci_datascience.PipelineStepOverrideDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.StepContainerConfigurationDetails != nil {
		result["step_container_configuration_details"] = []interface{}{PipelineContainerConfigurationDetailsToMap(obj.StepContainerConfigurationDetails)}
	}

	if obj.StepConfigurationDetails != nil {
		result["step_configuration_details"] = []interface{}{PipelineStepConfigurationDetailsToMap(obj.StepConfigurationDetails)}
	}

	if obj.StepDataflowConfigurationDetails != nil {
		result["step_dataflow_configuration_details"] = []interface{}{PipelineDataflowConfigurationDetailsToMap(obj.StepDataflowConfigurationDetails)}
	}

	if obj.StepInfrastructureConfigurationDetails != nil {
		result["step_infrastructure_configuration_details"] = []interface{}{PipelineInfrastructureConfigurationDetailsToMap(obj.StepInfrastructureConfigurationDetails)}
	}

	if obj.StepName != nil {
		result["step_name"] = string(*obj.StepName)
	}

	stepStorageMountConfigurationDetailsList := []interface{}{}
	for _, item := range obj.StepStorageMountConfigurationDetailsList {
		stepStorageMountConfigurationDetailsList = append(stepStorageMountConfigurationDetailsList, StorageMountConfigurationDetailsToMap(item))
	}
	result["step_storage_mount_configuration_details_list"] = stepStorageMountConfigurationDetailsList
	return result
}

func PipelineStepRunToMap(obj oci_datascience.PipelineStepRun) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.PipelineContainerStepRun:
		result["step_type"] = "CONTAINER"

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}

		if v.TimeFinished != nil {
			result["time_finished"] = v.TimeFinished.String()
		}

		if v.TimeStarted != nil {
			result["time_started"] = v.TimeStarted.String()
		}
	case oci_datascience.PipelineCustomScriptStepRun:
		result["step_type"] = "CUSTOM_SCRIPT"

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}

		if v.TimeFinished != nil {
			result["time_finished"] = v.TimeFinished.String()
		}

		if v.TimeStarted != nil {
			result["time_started"] = v.TimeStarted.String()
		}
	case oci_datascience.PipelineDataflowStepRun:
		result["step_type"] = "DATAFLOW"

		if v.DataflowRunId != nil {
			result["dataflow_run_id"] = string(*v.DataflowRunId)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}

		if v.TimeFinished != nil {
			result["time_finished"] = v.TimeFinished.String()
		}

		if v.TimeStarted != nil {
			result["time_started"] = v.TimeStarted.String()
		}
	case oci_datascience.PipelineMlJobStepRun:
		result["step_type"] = "ML_JOB"

		if v.JobRunId != nil {
			result["job_run_id"] = string(*v.JobRunId)
		}

		if v.StepRunName != nil {
			result["step_run_name"] = string(*v.StepRunName)
		}

		if v.LifecycleDetails != nil {
			result["lifecycle_details"] = string(*v.LifecycleDetails)
		}

		result["state"] = string(v.LifecycleState)

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}

		if v.TimeFinished != nil {
			result["time_finished"] = v.TimeFinished.String()
		}

		if v.TimeStarted != nil {
			result["time_started"] = v.TimeStarted.String()
		}
	default:
		log.Printf("[WARN] Received 'step_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatasciencePipelineRunResourceCrud) mapToStorageMountConfigurationDetails(fieldKeyFormat string) (oci_datascience.StorageMountConfigurationDetails, error) {
	var baseObject oci_datascience.StorageMountConfigurationDetails
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("FILE_STORAGE"):
		details := oci_datascience.FileStorageMountConfigurationDetails{}
		if exportId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "export_id")); ok {
			tmp := exportId.(string)
			details.ExportId = &tmp
		}
		if mountTargetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "mount_target_id")); ok {
			tmp := mountTargetId.(string)
			details.MountTargetId = &tmp
		}
		if destinationDirectoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_directory_name")); ok {
			tmp := destinationDirectoryName.(string)
			details.DestinationDirectoryName = &tmp
		}
		if destinationPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_path")); ok {
			tmp := destinationPath.(string)
			details.DestinationPath = &tmp
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE"):
		details := oci_datascience.ObjectStorageMountConfigurationDetails{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
			tmp := prefix.(string)
			details.Prefix = &tmp
		}
		if destinationDirectoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_directory_name")); ok {
			tmp := destinationDirectoryName.(string)
			details.DestinationDirectoryName = &tmp
		}
		if destinationPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "destination_path")); ok {
			tmp := destinationPath.(string)
			details.DestinationPath = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func (s *DatasciencePipelineRunResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangePipelineRunCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PipelineRunId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangePipelineRunCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
