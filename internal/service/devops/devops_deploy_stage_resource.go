// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsDeployStageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsDeployStage,
		Read:     readDevopsDeployStage,
		Update:   updateDevopsDeployStage,
		Delete:   deleteDevopsDeployStage,
		Schema: map[string]*schema.Schema{
			// Required
			"deploy_pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deploy_stage_predecessor_collection": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"deploy_stage_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT",
					"COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT",
					"COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL",
					"COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT",
					"COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT",
					"COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT",
					"DEPLOY_FUNCTION",
					"INVOKE_FUNCTION",
					"LOAD_BALANCER_TRAFFIC_SHIFT",
					"SHELL",
					"MANUAL_APPROVAL",
					"OKE_BLUE_GREEN_DEPLOYMENT",
					"OKE_BLUE_GREEN_TRAFFIC_SHIFT",
					"OKE_CANARY_APPROVAL",
					"OKE_CANARY_DEPLOYMENT",
					"OKE_CANARY_TRAFFIC_SHIFT",
					"OKE_DEPLOYMENT",
					"OKE_HELM_CHART_DEPLOYMENT",
					"WAIT",
				}, true),
			},

			// Optional
			"approval_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"approval_policy_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COUNT_BASED_APPROVAL",
							}, true),
						},
						"number_of_approvals_required": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"are_hooks_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"blue_backend_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"items": {
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
			"blue_green_strategy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ingress_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace_a": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace_b": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"strategy_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NGINX_BLUE_GREEN_STRATEGY",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"canary_strategy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"ingress_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"strategy_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"NGINX_CANARY_STRATEGY",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},
			"command_spec_deploy_artifact_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"compute_instance_group_blue_green_deployment_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compute_instance_group_canary_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compute_instance_group_canary_traffic_shift_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"compute_instance_group_deploy_environment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem:     schema.TypeString,
			},
			"container_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"container_config_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CONTAINER_INSTANCE_CONFIG",
							}, true),
						},
						"network_channel": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"network_channel_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"PRIVATE_ENDPOINT_CHANNEL",
											"SERVICE_VNIC_CHANNEL",
										}, true),
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"nsg_ids": {
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
						"shape_config": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"ocpus": {
										Type:     schema.TypeFloat,
										Required: true,
									},

									// Optional
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"shape_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"availability_domain": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
						},
						"compartment_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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
			"deploy_artifact_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"deploy_artifact_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"deploy_environment_id_a": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"deploy_environment_id_b": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"deployment_spec_deploy_artifact_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"docker_image_deploy_artifact_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"policy_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT",
								"COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE",
							}, true),
						},

						// Optional
						"failure_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"failure_percentage": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"function_deploy_environment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"function_timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"green_backend_ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"items": {
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
			"helm_chart_deploy_artifact_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"helm_command_artifact_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"is_async": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_debug_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_force_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_uninstall_on_stage_delete": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_validation_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"kubernetes_manifest_deploy_artifact_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"load_balancer_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"backend_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"listener_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"max_history": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"max_memory_in_mbs": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     tfresource.ValidateInt64TypeString,
				DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
			},
			"namespace": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oke_blue_green_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"oke_canary_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"oke_canary_traffic_shift_deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"oke_cluster_deploy_environment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"production_load_balancer_config": {
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
						"backend_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"listener_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"purpose": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"release_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rollback_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"policy_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"rollout_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"policy_type": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT",
								"COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE",
							}, true),
						},

						// Optional
						"batch_count": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"batch_delay_in_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"batch_percentage": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ramp_limit_percent": {
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"set_string": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"items": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
			"set_values": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 0,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"items": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
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
			"should_cleanup_on_fail": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_not_wait": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_reset_values": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_reuse_values": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_skip_crds": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_skip_render_subchart_notes": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"test_load_balancer_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"backend_port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"listener_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"load_balancer_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						// internal for work request access
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"traffic_shift_target": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"values_artifact_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: false,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"wait_criteria": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"wait_duration": {
							Type:     schema.TypeString,
							Required: true,
						},
						"wait_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ABSOLUTE_WAIT",
							}, true),
						},

						// Optional

						// Computed
					},
				},
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_id": {
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
		},
	}
}

func createDevopsDeployStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsDeployStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsDeployStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsDeployStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeployStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsDeployStageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.DeployStage
	DisableNotFoundRetries bool
}

func (s *DevopsDeployStageResourceCrud) ID() string {
	deployStage := *s.Res
	return *deployStage.GetId()
}

func (s *DevopsDeployStageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.DeployStageLifecycleStateCreating),
	}
}

func (s *DevopsDeployStageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.DeployStageLifecycleStateActive),
	}
}

func (s *DevopsDeployStageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_devops.DeployStageLifecycleStateDeleting),
	}
}

func (s *DevopsDeployStageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.DeployStageLifecycleStateDeleted),
	}
}

func (s *DevopsDeployStageResourceCrud) Create() error {
	request := oci_devops.CreateDeployStageRequest{}
	err := s.populateTopLevelPolymorphicCreateDeployStageRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateDeployStage(context.Background(), request)
	if err != nil {
		return err
	}
	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	log.Printf("##########################################")
	log.Printf("##########################################")
	fmt.Println(response)
	log.Printf("##########################################")
	log.Printf("##########################################")
	return s.getDeployStageFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsDeployStageResourceCrud) getDeployStageFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	deployStageId, err := deployStageWaitForWorkRequest(workId, "stage",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*deployStageId)

	return s.Get()
}

func deployStageWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "devops", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_devops.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func deployStageWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = deployStageWorkRequestShouldRetryFunc(timeout)

	response := oci_devops.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_devops.OperationStatusInProgress),
			string(oci_devops.OperationStatusAccepted),
			string(oci_devops.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_devops.OperationStatusSucceeded),
			string(oci_devops.OperationStatusFailed),
			string(oci_devops.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_devops.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_devops.OperationStatusFailed {
		return nil, getErrorFromDevopsDeployStageWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsDeployStageWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_devops.ListWorkRequestErrorsRequest{
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

func (s *DevopsDeployStageResourceCrud) Get() error {
	request := oci_devops.GetDeployStageRequest{}

	tmp := s.D.Id()
	request.DeployStageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetDeployStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DeployStage
	return nil
}

func (s *DevopsDeployStageResourceCrud) Update() error {
	request := oci_devops.UpdateDeployStageRequest{}
	err := s.populateTopLevelPolymorphicUpdateDeployStageRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateDeployStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getDeployStageFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsDeployStageResourceCrud) Delete() error {
	request := oci_devops.DeleteDeployStageRequest{}

	tmp := s.D.Id()
	request.DeployStageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteDeployStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := deployStageWaitForWorkRequest(workId, "stage",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsDeployStageResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.ComputeInstanceGroupBlueGreenDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT")

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeployEnvironmentIdA != nil {
			s.D.Set("deploy_environment_id_a", *v.DeployEnvironmentIdA)
		}

		if v.DeployEnvironmentIdB != nil {
			s.D.Set("deploy_environment_id_b", *v.DeployEnvironmentIdB)
		}

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			s.D.Set("failure_policy", failurePolicyArray)
		} else {
			s.D.Set("failure_policy", nil)
		}

		if v.ProductionLoadBalancerConfig != nil {
			s.D.Set("production_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)})
		} else {
			s.D.Set("production_load_balancer_config", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.TestLoadBalancerConfig != nil {
			s.D.Set("test_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)})
		} else {
			s.D.Set("test_load_balancer_config", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupBlueGreenTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT")

		if v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId != nil {
			s.D.Set("compute_instance_group_blue_green_deployment_deploy_stage_id", *v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryApprovalDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId != nil {
			s.D.Set("compute_instance_group_canary_traffic_shift_deploy_stage_id", *v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT")

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			s.D.Set("compute_instance_group_deploy_environment_id", *v.ComputeInstanceGroupDeployEnvironmentId)
		}

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.ProductionLoadBalancerConfig != nil {
			s.D.Set("production_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)})
		} else {
			s.D.Set("production_load_balancer_config", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.TestLoadBalancerConfig != nil {
			s.D.Set("test_load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)})
		} else {
			s.D.Set("test_load_balancer_config", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupCanaryTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT")

		if v.ComputeInstanceGroupCanaryDeployStageId != nil {
			s.D.Set("compute_instance_group_canary_deploy_stage_id", *v.ComputeInstanceGroupCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ComputeInstanceGroupDeployStage:
		s.D.Set("deploy_stage_type", "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT")

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			s.D.Set("compute_instance_group_deploy_environment_id", *v.ComputeInstanceGroupDeployEnvironmentId)
		}

		s.D.Set("deploy_artifact_ids", v.DeployArtifactIds)

		if v.DeploymentSpecDeployArtifactId != nil {
			s.D.Set("deployment_spec_deploy_artifact_id", *v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			s.D.Set("failure_policy", failurePolicyArray)
		} else {
			s.D.Set("failure_policy", nil)
		}

		if v.LoadBalancerConfig != nil {
			s.D.Set("load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)})
		} else {
			s.D.Set("load_balancer_config", nil)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			s.D.Set("rollout_policy", rolloutPolicyArray)
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.FunctionDeployStage:
		s.D.Set("deploy_stage_type", "DEPLOY_FUNCTION")

		s.D.Set("config", v.Config)

		if v.DockerImageDeployArtifactId != nil {
			s.D.Set("docker_image_deploy_artifact_id", *v.DockerImageDeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			s.D.Set("function_deploy_environment_id", *v.FunctionDeployEnvironmentId)
		}

		if v.FunctionTimeoutInSeconds != nil {
			s.D.Set("function_timeout_in_seconds", *v.FunctionTimeoutInSeconds)
		}

		if v.MaxMemoryInMBs != nil {
			s.D.Set("max_memory_in_mbs", strconv.FormatInt(*v.MaxMemoryInMBs, 10))
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.InvokeFunctionDeployStage:
		s.D.Set("deploy_stage_type", "INVOKE_FUNCTION")

		if v.DeployArtifactId != nil {
			s.D.Set("deploy_artifact_id", *v.DeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			s.D.Set("function_deploy_environment_id", *v.FunctionDeployEnvironmentId)
		}

		if v.IsAsync != nil {
			s.D.Set("is_async", *v.IsAsync)
		}

		if v.IsValidationEnabled != nil {
			s.D.Set("is_validation_enabled", *v.IsValidationEnabled)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.LoadBalancerTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "LOAD_BALANCER_TRAFFIC_SHIFT")

		if v.BlueBackendIps != nil {
			s.D.Set("blue_backend_ips", []interface{}{BackendSetIpCollectionToMap(v.BlueBackendIps)})
		} else {
			s.D.Set("blue_backend_ips", nil)
		}

		if v.GreenBackendIps != nil {
			s.D.Set("green_backend_ips", []interface{}{BackendSetIpCollectionToMap(v.GreenBackendIps)})
		} else {
			s.D.Set("green_backend_ips", nil)
		}

		if v.LoadBalancerConfig != nil {
			s.D.Set("load_balancer_config", []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)})
		} else {
			s.D.Set("load_balancer_config", nil)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		s.D.Set("traffic_shift_target", v.TrafficShiftTarget)

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ShellDeployStage:
		s.D.Set("deploy_stage_type", "SHELL")

		if v.CommandSpecDeployArtifactId != nil {
			s.D.Set("command_spec_deploy_artifact_id", *v.CommandSpecDeployArtifactId)
		}

		if v.ContainerConfig != nil {
			containerConfigArray := []interface{}{}
			if containerConfigMap := ContainerConfigToMap(&v.ContainerConfig); containerConfigMap != nil {
				containerConfigArray = append(containerConfigArray, containerConfigMap)
			}
			s.D.Set("container_config", containerConfigArray)
		} else {
			s.D.Set("container_config", nil)
		}

		if v.TimeoutInSeconds != nil {
			s.D.Set("timeout_in_seconds", *v.TimeoutInSeconds)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.ManualApprovalDeployStage:
		s.D.Set("deploy_stage_type", "MANUAL_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeBlueGreenDeployStage:
		s.D.Set("deploy_stage_type", "OKE_BLUE_GREEN_DEPLOYMENT")

		if v.BlueGreenStrategy != nil {
			blueGreenStrategyArray := []interface{}{}
			if blueGreenStrategyMap := OkeBlueGreenStrategyToMap(&v.BlueGreenStrategy); blueGreenStrategyMap != nil {
				blueGreenStrategyArray = append(blueGreenStrategyArray, blueGreenStrategyMap)
			}
			s.D.Set("blue_green_strategy", blueGreenStrategyArray)
		} else {
			s.D.Set("blue_green_strategy", nil)
		}

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeBlueGreenTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "OKE_BLUE_GREEN_TRAFFIC_SHIFT")

		if v.OkeBlueGreenDeployStageId != nil {
			s.D.Set("oke_blue_green_deploy_stage_id", *v.OkeBlueGreenDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryApprovalDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_APPROVAL")

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			s.D.Set("approval_policy", approvalPolicyArray)
		} else {
			s.D.Set("approval_policy", nil)
		}

		if v.OkeCanaryTrafficShiftDeployStageId != nil {
			s.D.Set("oke_canary_traffic_shift_deploy_stage_id", *v.OkeCanaryTrafficShiftDeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_DEPLOYMENT")

		if v.CanaryStrategy != nil {
			canaryStrategyArray := []interface{}{}
			if canaryStrategyMap := OkeCanaryStrategyToMap(&v.CanaryStrategy); canaryStrategyMap != nil {
				canaryStrategyArray = append(canaryStrategyArray, canaryStrategyMap)
			}
			s.D.Set("canary_strategy", canaryStrategyArray)
		} else {
			s.D.Set("canary_strategy", nil)
		}

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeCanaryTrafficShiftDeployStage:
		s.D.Set("deploy_stage_type", "OKE_CANARY_TRAFFIC_SHIFT")

		if v.OkeCanaryDeployStageId != nil {
			s.D.Set("oke_canary_deploy_stage_id", *v.OkeCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			s.D.Set("rollout_policy", []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)})
		} else {
			s.D.Set("rollout_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeDeployStage:
		s.D.Set("deploy_stage_type", "OKE_DEPLOYMENT")

		s.D.Set("kubernetes_manifest_deploy_artifact_ids", v.KubernetesManifestDeployArtifactIds)

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.WaitDeployStage:
		s.D.Set("deploy_stage_type", "WAIT")
		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			s.D.Set("wait_criteria", waitCriteriaArray)
		} else {
			s.D.Set("wait_criteria", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	case oci_devops.OkeHelmChartDeployStage:
		s.D.Set("deploy_stage_type", "OKE_HELM_CHART_DEPLOYMENT")

		if v.AreHooksEnabled != nil {
			s.D.Set("are_hooks_enabled", *v.AreHooksEnabled)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		s.D.Set("state", v.LifecycleState)

		if v.HelmChartDeployArtifactId != nil {
			s.D.Set("helm_chart_deploy_artifact_id", *v.HelmChartDeployArtifactId)
		}

		if v.HelmCommandArtifactIds != nil {
			s.D.Set("helm_command_artifact_ids", v.HelmCommandArtifactIds)
		}

		if v.IsDebugEnabled != nil {
			s.D.Set("is_debug_enabled", *v.IsDebugEnabled)
		}

		if v.IsForceEnabled != nil {
			s.D.Set("is_force_enabled", *v.IsForceEnabled)
		}

		if v.IsUninstallOnStageDelete != nil {
			s.D.Set("is_uninstall_on_stage_delete", *v.IsUninstallOnStageDelete)
		}

		if v.MaxHistory != nil {
			s.D.Set("max_history", *v.MaxHistory)
		}

		if v.Namespace != nil {
			s.D.Set("namespace", *v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			s.D.Set("oke_cluster_deploy_environment_id", *v.OkeClusterDeployEnvironmentId)
		}

		s.D.Set("purpose", v.Purpose)

		if v.ReleaseName != nil {
			s.D.Set("release_name", *v.ReleaseName)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			s.D.Set("rollback_policy", rollbackPolicyArray)
		} else {
			s.D.Set("rollback_policy", nil)
		}

		if v.SetValues != nil {
			setValues := []interface{}{HelmSetValueCollectionToMap(v.SetValues)}
			common.Debugf("OkeHelmChartDeployStage: setValues= %v\n", setValues)
			if len(setValues) > 0 {
				common.Debugf("OkeHelmChartDeployStage: setValues= %v\n", setValues[0])
				setValuesMap := setValues[0].(map[string]interface{})
				items := setValuesMap["items"].([]interface{})
				if len(items) > 0 {
					s.D.Set("set_values", setValues)
				}
			}
		}

		if v.SetString != nil {
			setString := []interface{}{HelmSetValueCollectionToMap(v.SetString)}
			common.Debugf("OkeHelmChartDeployStage: SetString= %v\n", setString)
			if len(setString) > 0 {
				common.Debugf("OkeHelmChartDeployStage: SetString= %v\n", setString[0])
				setStringMap := setString[0].(map[string]interface{})
				items := setStringMap["items"].([]interface{})
				if len(items) > 0 {
					s.D.Set("set_string", setString)
				}
			}
		}

		if v.ShouldCleanupOnFail != nil {
			s.D.Set("should_cleanup_on_fail", *v.ShouldCleanupOnFail)
		}

		if v.ShouldNotWait != nil {
			s.D.Set("should_not_wait", *v.ShouldNotWait)
		}

		if v.ShouldResetValues != nil {
			s.D.Set("should_reset_values", *v.ShouldResetValues)
		}

		if v.ShouldReuseValues != nil {
			s.D.Set("should_reuse_values", *v.ShouldReuseValues)
		}

		if v.ShouldSkipCrds != nil {
			s.D.Set("should_skip_crds", *v.ShouldSkipCrds)
		}

		if v.ShouldSkipRenderSubchartNotes != nil {
			s.D.Set("should_skip_render_subchart_notes", *v.ShouldSkipRenderSubchartNotes)
		}

		if v.TimeoutInSeconds != nil {
			s.D.Set("timeout_in_seconds", *v.TimeoutInSeconds)
		}

		if v.ValuesArtifactIds != nil {
			s.D.Set("values_artifact_ids", v.ValuesArtifactIds)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStagePredecessorCollection != nil {
			s.D.Set("deploy_stage_predecessor_collection", []interface{}{DeployStagePredecessorCollectionToMap(v.DeployStagePredecessorCollection)})
		} else {
			s.D.Set("deploy_stage_predecessor_collection", nil)
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		if v.ProjectId != nil {
			s.D.Set("project_id", *v.ProjectId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}
	default:
		log.Printf("[WARN] Received 'deploy_stage_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func (s *DevopsDeployStageResourceCrud) mapToApprovalPolicy(fieldKeyFormat string) (oci_devops.ApprovalPolicy, error) {
	var baseObject oci_devops.ApprovalPolicy
	//discriminator
	approvalPolicyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "approval_policy_type"))
	var approvalPolicyType string
	if ok {
		approvalPolicyType = approvalPolicyTypeRaw.(string)
	} else {
		approvalPolicyType = "" // default value
	}
	switch strings.ToLower(approvalPolicyType) {
	case strings.ToLower("COUNT_BASED_APPROVAL"):
		details := oci_devops.CountBasedApprovalPolicy{}
		if numberOfApprovalsRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "number_of_approvals_required")); ok {
			tmp := numberOfApprovalsRequired.(int)
			details.NumberOfApprovalsRequired = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown approval_policy_type '%v' was specified", approvalPolicyType)
	}
	return baseObject, nil
}

func ApprovalPolicyToMap(obj *oci_devops.ApprovalPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.CountBasedApprovalPolicy:
		result["approval_policy_type"] = "COUNT_BASED_APPROVAL"

		if v.NumberOfApprovalsRequired != nil {
			result["number_of_approvals_required"] = int(*v.NumberOfApprovalsRequired)
		}
	default:
		log.Printf("[WARN] Received 'approval_policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToBackendSetIpCollection(fieldKeyFormat string) (oci_devops.BackendSetIpCollection, error) {
	result := oci_devops.BackendSetIpCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func BackendSetIpCollectionToMap(obj *oci_devops.BackendSetIpCollection) map[string]interface{} {
	result := map[string]interface{}{}

	result["items"] = obj.Items

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToComputeInstanceGroupFailurePolicy(fieldKeyFormat string) (oci_devops.ComputeInstanceGroupFailurePolicy, error) {
	var baseObject oci_devops.ComputeInstanceGroupFailurePolicy
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"):
		details := oci_devops.ComputeInstanceGroupFailurePolicyByCount{}
		if failureCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_count")); ok {
			tmp := failureCount.(int)
			details.FailureCount = &tmp
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE"):
		details := oci_devops.ComputeInstanceGroupFailurePolicyByPercentage{}
		if failurePercentage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "failure_percentage")); ok {
			tmp := failurePercentage.(int)
			details.FailurePercentage = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func ComputeInstanceGroupFailurePolicyToMap(obj *oci_devops.ComputeInstanceGroupFailurePolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.ComputeInstanceGroupFailurePolicyByCount:
		result["policy_type"] = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_COUNT"

		if v.FailureCount != nil {
			result["failure_count"] = int(*v.FailureCount)
		}
	case oci_devops.ComputeInstanceGroupFailurePolicyByPercentage:
		result["policy_type"] = "COMPUTE_INSTANCE_GROUP_FAILURE_POLICY_BY_PERCENTAGE"

		if v.FailurePercentage != nil {
			result["failure_percentage"] = int(*v.FailurePercentage)
		}
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat string) (oci_devops.ComputeInstanceGroupRolloutPolicy, error) {
	var baseObject oci_devops.ComputeInstanceGroupRolloutPolicy
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"):
		details := oci_devops.ComputeInstanceGroupLinearRolloutPolicyByCount{}
		if batchCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_count")); ok {
			tmp := batchCount.(int)
			details.BatchCount = &tmp
		}
		if batchDelayInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_delay_in_seconds")); ok {
			tmp := batchDelayInSeconds.(int)
			details.BatchDelayInSeconds = &tmp
		}
		baseObject = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE"):
		details := oci_devops.ComputeInstanceGroupLinearRolloutPolicyByPercentage{}
		if batchPercentage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_percentage")); ok {
			tmp := batchPercentage.(int)
			details.BatchPercentage = &tmp
		}
		if batchDelayInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_delay_in_seconds")); ok {
			tmp := batchDelayInSeconds.(int)
			details.BatchDelayInSeconds = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func ComputeInstanceGroupRolloutPolicyToMap(obj *oci_devops.ComputeInstanceGroupRolloutPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.ComputeInstanceGroupLinearRolloutPolicyByCount:
		result["policy_type"] = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_COUNT"

		if v.BatchCount != nil {
			result["batch_count"] = int(*v.BatchCount)
		}
		if v.BatchDelayInSeconds != nil {
			result["batch_delay_in_seconds"] = int(*v.BatchDelayInSeconds)
		}
	case oci_devops.ComputeInstanceGroupLinearRolloutPolicyByPercentage:
		result["policy_type"] = "COMPUTE_INSTANCE_GROUP_LINEAR_ROLLOUT_POLICY_BY_PERCENTAGE"

		if v.BatchPercentage != nil {
			result["batch_percentage"] = int(*v.BatchPercentage)
		}
		if v.BatchDelayInSeconds != nil {
			result["batch_delay_in_seconds"] = int(*v.BatchDelayInSeconds)
		}
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToContainerConfig(fieldKeyFormat string) (oci_devops.ContainerConfig, error) {
	var baseObject oci_devops.ContainerConfig
	//discriminator
	containerConfigTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "container_config_type"))
	var containerConfigType string
	if ok {
		containerConfigType = containerConfigTypeRaw.(string)
	} else {
		containerConfigType = "" // default value
	}
	switch strings.ToLower(containerConfigType) {
	case strings.ToLower("CONTAINER_INSTANCE_CONFIG"):
		details := oci_devops.ContainerInstanceConfig{}
		if availabilityDomain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "availability_domain")); ok {
			tmp := availabilityDomain.(string)
			details.AvailabilityDomain = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if networkChannel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_channel")); ok {
			if tmpList := networkChannel.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_channel"), 0)
				tmp, err := s.mapToNetworkChannel(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert network_channel, encountered error: %v", err)
				}
				details.NetworkChannel = tmp
			}
		}
		if shapeConfig, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_config")); ok {
			if tmpList := shapeConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "shape_config"), 0)
				tmp, err := s.mapToShapeConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert shape_config, encountered error: %v", err)
				}
				details.ShapeConfig = &tmp
			}
		}
		if shapeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "shape_name")); ok {
			tmp := shapeName.(string)
			details.ShapeName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown container_config_type '%v' was specified", containerConfigType)
	}
	return baseObject, nil
}

func ContainerConfigToMap(obj *oci_devops.ContainerConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.ContainerInstanceConfig:
		result["container_config_type"] = "CONTAINER_INSTANCE_CONFIG"

		if v.AvailabilityDomain != nil {
			result["availability_domain"] = string(*v.AvailabilityDomain)
		}

		if v.CompartmentId != nil {
			result["compartment_id"] = string(*v.CompartmentId)
		}

		if v.NetworkChannel != nil {
			networkChannelArray := []interface{}{}
			if networkChannelMap := NetworkChannelToMapForShellStage(&v.NetworkChannel, false); networkChannelMap != nil {
				networkChannelArray = append(networkChannelArray, networkChannelMap)
			}
			result["network_channel"] = networkChannelArray
		}

		if v.ShapeConfig != nil {
			result["shape_config"] = []interface{}{ShapeConfigToMap(v.ShapeConfig)}
		}

		if v.ShapeName != nil {
			result["shape_name"] = string(*v.ShapeName)
		}
	default:
		log.Printf("[WARN] Received 'container_config_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DeployStagePredecessorToMap(obj oci_devops.DeployStagePredecessor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func DeployStagePredecessorCollectionToMap(obj *oci_devops.DeployStagePredecessorCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeployStagePredecessorToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToDeployStagePredecessorCollection(fieldKeyFormat string) (oci_devops.DeployStagePredecessorCollection, error) {
	result := oci_devops.DeployStagePredecessorCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeployStagePredecessor, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeployStagePredecessor(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func (s *DevopsDeployStageResourceCrud) mapToDeployStagePredecessor(fieldKeyFormat string) (oci_devops.DeployStagePredecessor, error) {
	result := oci_devops.DeployStagePredecessor{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}

func (s *DevopsDeployStageResourceCrud) mapToDeployStageRollbackPolicy(fieldKeyFormat string) (oci_devops.DeployStageRollbackPolicy, error) {
	var baseObject oci_devops.DeployStageRollbackPolicy
	//discriminator
	policyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "policy_type"))
	var policyType string
	if ok {
		policyType = policyTypeRaw.(string)
	} else {
		policyType = "" // default value
	}
	switch strings.ToLower(policyType) {
	case strings.ToLower("AUTOMATED_STAGE_ROLLBACK_POLICY"):
		details := oci_devops.AutomatedDeployStageRollbackPolicy{}
		baseObject = details
	case strings.ToLower("NO_STAGE_ROLLBACK_POLICY"):
		details := oci_devops.NoDeployStageRollbackPolicy{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown policy_type '%v' was specified", policyType)
	}
	return baseObject, nil
}

func DeployStageRollbackPolicyToMap(obj *oci_devops.DeployStageRollbackPolicy) map[string]interface{} {
	result := map[string]interface{}{}
	switch (*obj).(type) {
	case oci_devops.AutomatedDeployStageRollbackPolicy:
		result["policy_type"] = "AUTOMATED_STAGE_ROLLBACK_POLICY"
	case oci_devops.NoDeployStageRollbackPolicy:
		result["policy_type"] = "NO_STAGE_ROLLBACK_POLICY"
	default:
		log.Printf("[WARN] Received 'policy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func DeployStageSummaryToMap(obj oci_devops.DeployStageSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.ComputeInstanceGroupBlueGreenDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT"

		result["deploy_artifact_ids"] = v.DeployArtifactIds

		if v.DeployEnvironmentIdA != nil {
			result["deploy_environment_id_a"] = string(*v.DeployEnvironmentIdA)
		}

		if v.DeployEnvironmentIdB != nil {
			result["deploy_environment_id_b"] = string(*v.DeployEnvironmentIdB)
		}

		if v.DeploymentSpecDeployArtifactId != nil {
			result["deployment_spec_deploy_artifact_id"] = string(*v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			result["failure_policy"] = failurePolicyArray
		}

		if v.ProductionLoadBalancerConfig != nil {
			result["production_load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)}
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			result["rollout_policy"] = rolloutPolicyArray
		}

		if v.TestLoadBalancerConfig != nil {
			result["test_load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)}
		}
	case oci_devops.ComputeInstanceGroupBlueGreenTrafficShiftDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT"

		if v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId != nil {
			result["compute_instance_group_blue_green_deployment_deploy_stage_id"] = string(*v.ComputeInstanceGroupBlueGreenDeploymentDeployStageId)
		}
	case oci_devops.ComputeInstanceGroupCanaryApprovalDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			result["approval_policy"] = approvalPolicyArray
		}

		if v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId != nil {
			result["compute_instance_group_canary_traffic_shift_deploy_stage_id"] = string(*v.ComputeInstanceGroupCanaryTrafficShiftDeployStageId)
		}
	case oci_devops.ComputeInstanceGroupCanaryDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT"

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			result["compute_instance_group_deploy_environment_id"] = string(*v.ComputeInstanceGroupDeployEnvironmentId)
		}

		result["deploy_artifact_ids"] = v.DeployArtifactIds

		if v.DeploymentSpecDeployArtifactId != nil {
			result["deployment_spec_deploy_artifact_id"] = string(*v.DeploymentSpecDeployArtifactId)
		}

		if v.ProductionLoadBalancerConfig != nil {
			result["production_load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.ProductionLoadBalancerConfig)}
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			result["rollout_policy"] = rolloutPolicyArray
		}

		if v.TestLoadBalancerConfig != nil {
			result["test_load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.TestLoadBalancerConfig)}
		}
	case oci_devops.ComputeInstanceGroupCanaryTrafficShiftDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT"

		if v.ComputeInstanceGroupCanaryDeployStageId != nil {
			result["compute_instance_group_canary_deploy_stage_id"] = string(*v.ComputeInstanceGroupCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			result["rollout_policy"] = []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)}
		}
	case oci_devops.ComputeInstanceGroupDeployStageSummary:
		result["deploy_stage_type"] = "COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"

		if v.ComputeInstanceGroupDeployEnvironmentId != nil {
			result["compute_instance_group_deploy_environment_id"] = string(*v.ComputeInstanceGroupDeployEnvironmentId)
		}

		result["deploy_artifact_ids"] = v.DeployArtifactIds

		if v.DeploymentSpecDeployArtifactId != nil {
			result["deployment_spec_deploy_artifact_id"] = string(*v.DeploymentSpecDeployArtifactId)
		}

		if v.FailurePolicy != nil {
			failurePolicyArray := []interface{}{}
			if failurePolicyMap := ComputeInstanceGroupFailurePolicyToMap(&v.FailurePolicy); failurePolicyMap != nil {
				failurePolicyArray = append(failurePolicyArray, failurePolicyMap)
			}
			result["failure_policy"] = failurePolicyArray
		}

		if v.LoadBalancerConfig != nil {
			result["load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)}
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			result["rollback_policy"] = rollbackPolicyArray
		}

		if v.RolloutPolicy != nil {
			rolloutPolicyArray := []interface{}{}
			if rolloutPolicyMap := ComputeInstanceGroupRolloutPolicyToMap(&v.RolloutPolicy); rolloutPolicyMap != nil {
				rolloutPolicyArray = append(rolloutPolicyArray, rolloutPolicyMap)
			}
			result["rollout_policy"] = rolloutPolicyArray
		}
	case oci_devops.FunctionDeployStageSummary:
		result["deploy_stage_type"] = "DEPLOY_FUNCTION"

		result["config"] = v.Config

		if v.DockerImageDeployArtifactId != nil {
			result["docker_image_deploy_artifact_id"] = string(*v.DockerImageDeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			result["function_deploy_environment_id"] = string(*v.FunctionDeployEnvironmentId)
		}

		if v.FunctionTimeoutInSeconds != nil {
			result["function_timeout_in_seconds"] = int(*v.FunctionTimeoutInSeconds)
		}

		if v.MaxMemoryInMBs != nil {
			result["max_memory_in_mbs"] = strconv.FormatInt(*v.MaxMemoryInMBs, 10)
		}
	case oci_devops.InvokeFunctionDeployStageSummary:
		result["deploy_stage_type"] = "INVOKE_FUNCTION"

		if v.DeployArtifactId != nil {
			result["deploy_artifact_id"] = string(*v.DeployArtifactId)
		}

		if v.FunctionDeployEnvironmentId != nil {
			result["function_deploy_environment_id"] = string(*v.FunctionDeployEnvironmentId)
		}

		if v.IsAsync != nil {
			result["is_async"] = bool(*v.IsAsync)
		}

		if v.IsValidationEnabled != nil {
			result["is_validation_enabled"] = bool(*v.IsValidationEnabled)
		}
	case oci_devops.LoadBalancerTrafficShiftDeployStageSummary:
		result["deploy_stage_type"] = "LOAD_BALANCER_TRAFFIC_SHIFT"

		if v.BlueBackendIps != nil {
			result["blue_backend_ips"] = []interface{}{BackendSetIpCollectionToMap(v.BlueBackendIps)}
		}

		if v.GreenBackendIps != nil {
			result["green_backend_ips"] = []interface{}{BackendSetIpCollectionToMap(v.GreenBackendIps)}
		}

		if v.LoadBalancerConfig != nil {
			result["load_balancer_config"] = []interface{}{LoadBalancerConfigToMap(v.LoadBalancerConfig)}
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			result["rollback_policy"] = rollbackPolicyArray
		}

		if v.RolloutPolicy != nil {
			result["rollout_policy"] = []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)}
		}

		result["traffic_shift_target"] = string(v.TrafficShiftTarget)
	case oci_devops.ShellDeployStageSummary:
		result["deploy_stage_type"] = "SHELL"

		if v.CommandSpecDeployArtifactId != nil {
			result["command_spec_deploy_artifact_id"] = string(*v.CommandSpecDeployArtifactId)
		}

		if v.ContainerConfig != nil {
			containerConfigArray := []interface{}{}
			if containerConfigMap := ContainerConfigToMap(&v.ContainerConfig); containerConfigMap != nil {
				containerConfigArray = append(containerConfigArray, containerConfigMap)
			}
			result["container_config"] = containerConfigArray
		}

		if v.TimeoutInSeconds != nil {
			result["timeout_in_seconds"] = int(*v.TimeoutInSeconds)
		}
	case oci_devops.ManualApprovalDeployStageSummary:
		result["deploy_stage_type"] = "MANUAL_APPROVAL"

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			result["approval_policy"] = approvalPolicyArray
		}
	case oci_devops.OkeBlueGreenDeployStageSummary:
		result["deploy_stage_type"] = "OKE_BLUE_GREEN_DEPLOYMENT"

		if v.BlueGreenStrategy != nil {
			blueGreenStrategyArray := []interface{}{}
			if blueGreenStrategyMap := OkeBlueGreenStrategyToMap(&v.BlueGreenStrategy); blueGreenStrategyMap != nil {
				blueGreenStrategyArray = append(blueGreenStrategyArray, blueGreenStrategyMap)
			}
			result["blue_green_strategy"] = blueGreenStrategyArray
		}

		result["kubernetes_manifest_deploy_artifact_ids"] = v.KubernetesManifestDeployArtifactIds

		if v.OkeClusterDeployEnvironmentId != nil {
			result["oke_cluster_deploy_environment_id"] = string(*v.OkeClusterDeployEnvironmentId)
		}
	case oci_devops.OkeBlueGreenTrafficShiftDeployStageSummary:
		result["deploy_stage_type"] = "OKE_BLUE_GREEN_TRAFFIC_SHIFT"

		if v.OkeBlueGreenDeployStageId != nil {
			result["oke_blue_green_deploy_stage_id"] = string(*v.OkeBlueGreenDeployStageId)
		}
	case oci_devops.OkeCanaryApprovalDeployStageSummary:
		result["deploy_stage_type"] = "OKE_CANARY_APPROVAL"

		if v.ApprovalPolicy != nil {
			approvalPolicyArray := []interface{}{}
			if approvalPolicyMap := ApprovalPolicyToMap(&v.ApprovalPolicy); approvalPolicyMap != nil {
				approvalPolicyArray = append(approvalPolicyArray, approvalPolicyMap)
			}
			result["approval_policy"] = approvalPolicyArray
		}

		if v.OkeCanaryTrafficShiftDeployStageId != nil {
			result["oke_canary_traffic_shift_deploy_stage_id"] = string(*v.OkeCanaryTrafficShiftDeployStageId)
		}
	case oci_devops.OkeCanaryDeployStageSummary:
		result["deploy_stage_type"] = "OKE_CANARY_DEPLOYMENT"

		if v.CanaryStrategy != nil {
			canaryStrategyArray := []interface{}{}
			if canaryStrategyMap := OkeCanaryStrategyToMap(&v.CanaryStrategy); canaryStrategyMap != nil {
				canaryStrategyArray = append(canaryStrategyArray, canaryStrategyMap)
			}
			result["canary_strategy"] = canaryStrategyArray
		}

		result["kubernetes_manifest_deploy_artifact_ids"] = v.KubernetesManifestDeployArtifactIds

		if v.OkeClusterDeployEnvironmentId != nil {
			result["oke_cluster_deploy_environment_id"] = string(*v.OkeClusterDeployEnvironmentId)
		}
	case oci_devops.OkeCanaryTrafficShiftDeployStageSummary:
		result["deploy_stage_type"] = "OKE_CANARY_TRAFFIC_SHIFT"

		if v.OkeCanaryDeployStageId != nil {
			result["oke_canary_deploy_stage_id"] = string(*v.OkeCanaryDeployStageId)
		}

		if v.RolloutPolicy != nil {
			result["rollout_policy"] = []interface{}{LoadBalancerTrafficShiftRolloutPolicyToMap(v.RolloutPolicy)}
		}
	case oci_devops.OkeDeployStageSummary:
		result["deploy_stage_type"] = "OKE_DEPLOYMENT"

		result["kubernetes_manifest_deploy_artifact_ids"] = v.KubernetesManifestDeployArtifactIds

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			result["oke_cluster_deploy_environment_id"] = string(*v.OkeClusterDeployEnvironmentId)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			result["rollback_policy"] = rollbackPolicyArray
		}
	case oci_devops.WaitDeployStageSummary:
		result["deploy_stage_type"] = "WAIT"

		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaSummaryToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			result["wait_criteria"] = waitCriteriaArray
		}
	case oci_devops.OkeHelmChartDeployStageSummary:
		result["deploy_stage_type"] = "OKE_HELM_CHART_DEPLOYMENT"

		if v.AreHooksEnabled != nil {
			result["are_hooks_enabled"] = bool(*v.AreHooksEnabled)
		}

		if v.HelmChartDeployArtifactId != nil {
			result["helm_chart_deploy_artifact_id"] = string(*v.HelmChartDeployArtifactId)
		}

		if v.HelmCommandArtifactIds != nil {
			result["helm_command_artifact_ids"] = v.HelmCommandArtifactIds
		}

		if v.IsDebugEnabled != nil {
			result["is_debug_enabled"] = bool(*v.IsDebugEnabled)
		}

		if v.IsForceEnabled != nil {
			result["is_force_enabled"] = bool(*v.IsForceEnabled)
		}

		if v.IsUninstallOnStageDelete != nil {
			result["is_uninstall_on_stage_delete"] = bool(*v.IsUninstallOnStageDelete)
		}

		if v.MaxHistory != nil {
			result["max_history"] = int(*v.MaxHistory)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.OkeClusterDeployEnvironmentId != nil {
			result["oke_cluster_deploy_environment_id"] = string(*v.OkeClusterDeployEnvironmentId)
		}

		if v.Purpose != "" {
			result["purpose"] = string(v.Purpose)
		} else {
			v.Purpose = oci_devops.OkeHelmChartDeployStageSummaryPurposeEnum("EXECUTE_HELM_UPGRADE")
			result["purpose"] = v.Purpose
		}

		if v.ReleaseName != nil {
			result["release_name"] = string(*v.ReleaseName)
		}

		if v.RollbackPolicy != nil {
			rollbackPolicyArray := []interface{}{}
			if rollbackPolicyMap := DeployStageRollbackPolicyToMap(&v.RollbackPolicy); rollbackPolicyMap != nil {
				rollbackPolicyArray = append(rollbackPolicyArray, rollbackPolicyMap)
			}
			result["rollback_policy"] = rollbackPolicyArray
		}

		if v.SetString != nil {
			setString := []interface{}{HelmSetValueCollectionToMap(v.SetString)}
			if len(setString) > 0 {
				setStringMap := setString[0].(map[string]interface{})
				items := setStringMap["items"].([]interface{})
				if len(items) > 0 {
					result["set_string"] = setString
				}
			}
		}

		if v.SetValues != nil {
			setValues := []interface{}{HelmSetValueCollectionToMap(v.SetValues)}
			if len(setValues) > 0 {
				setValuesMap := setValues[0].(map[string]interface{})
				items := setValuesMap["items"].([]interface{})
				if len(items) > 0 {
					result["set_values"] = setValues
				}
			}
		}

		if v.ShouldCleanupOnFail != nil {
			result["should_cleanup_on_fail"] = bool(*v.ShouldCleanupOnFail)
		}

		if v.ShouldNotWait != nil {
			result["should_not_wait"] = bool(*v.ShouldNotWait)
		}

		if v.ShouldResetValues != nil {
			result["should_reset_values"] = bool(*v.ShouldResetValues)
		}

		if v.ShouldReuseValues != nil {
			result["should_reuse_values"] = bool(*v.ShouldReuseValues)
		}

		if v.ShouldSkipCrds != nil {
			result["should_skip_crds"] = bool(*v.ShouldSkipCrds)
		}

		if v.ShouldSkipRenderSubchartNotes != nil {
			result["should_skip_render_subchart_notes"] = bool(*v.ShouldSkipRenderSubchartNotes)
		}

		if v.TimeoutInSeconds != nil {
			result["timeout_in_seconds"] = int(*v.TimeoutInSeconds)
		}

		result["values_artifact_ids"] = v.ValuesArtifactIds
	default:
		log.Printf("[WARN] Received 'deploy_stage_type' of unknown type %v", obj)
		return nil
	}

	if obj.GetId() != nil {
		result["id"] = obj.GetId()
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = obj.GetCompartmentId()
	}

	if obj.GetDescription() != nil {
		result["description"] = obj.GetDescription()
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	result["state"] = obj.GetLifecycleState()

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = obj.GetLifecycleDetails()
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetProjectId() != nil {
		result["project_id"] = obj.GetProjectId()
	}

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToHelmSetValue(fieldKeyFormat string) (oci_devops.HelmSetValue, error) {
	result := oci_devops.HelmSetValue{}

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

func HelmSetValueToMap(obj oci_devops.HelmSetValue) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToHelmSetValueCollection(fieldKeyFormat string) (oci_devops.HelmSetValueCollection, error) {
	result := oci_devops.HelmSetValueCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.HelmSetValue, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToHelmSetValue(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "items")) {
			result.Items = tmp
		}
	}

	return result, nil
}

func HelmSetValueCollectionToMap(obj *oci_devops.HelmSetValueCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, HelmSetValueToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToLoadBalancerConfig(fieldKeyFormat string) (oci_devops.LoadBalancerConfig, error) {
	result := oci_devops.LoadBalancerConfig{}

	if backendPort, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "backend_port")); ok {
		tmp := backendPort.(int)
		result.BackendPort = &tmp
	}

	if listenerName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "listener_name")); ok {
		tmp := listenerName.(string)
		result.ListenerName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "load_balancer_id")); ok {
		tmp := loadBalancerId.(string)
		result.LoadBalancerId = &tmp
	}

	return result, nil
}

func LoadBalancerConfigToMap(obj *oci_devops.LoadBalancerConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BackendPort != nil {
		result["backend_port"] = int(*obj.BackendPort)
	}

	if obj.ListenerName != nil {
		result["listener_name"] = string(*obj.ListenerName)
	}

	if obj.LoadBalancerId != nil {
		result["load_balancer_id"] = string(*obj.LoadBalancerId)
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat string) (oci_devops.LoadBalancerTrafficShiftRolloutPolicy, error) {
	result := oci_devops.LoadBalancerTrafficShiftRolloutPolicy{}

	if batchCount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_count")); ok {
		tmp := batchCount.(int)
		result.BatchCount = &tmp
	}

	if batchDelayInSeconds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "batch_delay_in_seconds")); ok {
		tmp := batchDelayInSeconds.(int)
		result.BatchDelayInSeconds = &tmp
	}

	if rampLimitPercent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ramp_limit_percent")); ok {
		tmp := float32(rampLimitPercent.(float64))
		result.RampLimitPercent = &tmp
	}

	return result, nil
}

func LoadBalancerTrafficShiftRolloutPolicyToMap(obj *oci_devops.LoadBalancerTrafficShiftRolloutPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BatchCount != nil {
		result["batch_count"] = int(*obj.BatchCount)
	}

	if obj.BatchDelayInSeconds != nil {
		result["batch_delay_in_seconds"] = int(*obj.BatchDelayInSeconds)
	}

	if obj.RampLimitPercent != nil {
		result["ramp_limit_percent"] = float32(*obj.RampLimitPercent)
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToNetworkChannel(fieldKeyFormat string) (oci_devops.NetworkChannel, error) {
	var baseObject oci_devops.NetworkChannel
	//discriminator
	networkChannelTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_channel_type"))
	var networkChannelType string
	if ok {
		networkChannelType = networkChannelTypeRaw.(string)
	} else {
		networkChannelType = "" // default value
	}
	switch strings.ToLower(networkChannelType) {
	case strings.ToLower("PRIVATE_ENDPOINT_CHANNEL"):
		details := oci_devops.PrivateEndpointChannel{}
		if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
			interfaces := nsgIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
				details.NsgIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	case strings.ToLower("SERVICE_VNIC_CHANNEL"):
		details := oci_devops.ServiceVnicChannel{}
		if nsgIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "nsg_ids")); ok {
			interfaces := nsgIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "nsg_ids")) {
				details.NsgIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown network_channel_type '%v' was specified", networkChannelType)
	}
	return baseObject, nil
}

func NetworkChannelToMapForShellStage(obj *oci_devops.NetworkChannel, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.PrivateEndpointChannel:
		result["network_channel_type"] = "PRIVATE_ENDPOINT_CHANNEL"

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		result["nsg_ids"] = nsgIds

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	case oci_devops.ServiceVnicChannel:
		result["network_channel_type"] = "SERVICE_VNIC_CHANNEL"

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		result["nsg_ids"] = nsgIds

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}
	default:
		log.Printf("[WARN] Received 'network_channel_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToOkeBlueGreenStrategy(fieldKeyFormat string) (oci_devops.OkeBlueGreenStrategy, error) {
	var baseObject oci_devops.OkeBlueGreenStrategy
	//discriminator
	strategyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy_type"))
	var strategyType string
	if ok {
		strategyType = strategyTypeRaw.(string)
	} else {
		strategyType = "" // default value
	}
	switch strings.ToLower(strategyType) {
	case strings.ToLower("NGINX_BLUE_GREEN_STRATEGY"):
		details := oci_devops.NginxBlueGreenStrategy{}
		if ingressName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ingress_name")); ok {
			tmp := ingressName.(string)
			details.IngressName = &tmp
		}
		if namespaceA, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace_a")); ok {
			tmp := namespaceA.(string)
			details.NamespaceA = &tmp
		}
		if namespaceB, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace_b")); ok {
			tmp := namespaceB.(string)
			details.NamespaceB = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy_type '%v' was specified", strategyType)
	}
	return baseObject, nil
}

func OkeBlueGreenStrategyToMap(obj *oci_devops.OkeBlueGreenStrategy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.NginxBlueGreenStrategy:
		result["strategy_type"] = "NGINX_BLUE_GREEN_STRATEGY"

		if v.IngressName != nil {
			result["ingress_name"] = string(*v.IngressName)
		}

		if v.NamespaceA != nil {
			result["namespace_a"] = string(*v.NamespaceA)
		}

		if v.NamespaceB != nil {
			result["namespace_b"] = string(*v.NamespaceB)
		}
	default:
		log.Printf("[WARN] Received 'strategy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToOkeCanaryStrategy(fieldKeyFormat string) (oci_devops.OkeCanaryStrategy, error) {
	var baseObject oci_devops.OkeCanaryStrategy
	//discriminator
	strategyTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "strategy_type"))
	var strategyType string
	if ok {
		strategyType = strategyTypeRaw.(string)
	} else {
		strategyType = "" // default value
	}
	switch strings.ToLower(strategyType) {
	case strings.ToLower("NGINX_CANARY_STRATEGY"):
		details := oci_devops.NginxCanaryStrategy{}
		if ingressName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ingress_name")); ok {
			tmp := ingressName.(string)
			details.IngressName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown strategy_type '%v' was specified", strategyType)
	}
	return baseObject, nil
}

func OkeCanaryStrategyToMap(obj *oci_devops.OkeCanaryStrategy) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.NginxCanaryStrategy:
		result["strategy_type"] = "NGINX_CANARY_STRATEGY"

		if v.IngressName != nil {
			result["ingress_name"] = string(*v.IngressName)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}
	default:
		log.Printf("[WARN] Received 'strategy_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToShapeConfig(fieldKeyFormat string) (oci_devops.ShapeConfig, error) {
	result := oci_devops.ShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := float32(memoryInGBs.(float64))
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := float32(ocpus.(float64))
		result.Ocpus = &tmp
	}

	return result, nil
}

func ShapeConfigToMap(obj *oci_devops.ShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) mapToWaitCriteria(fieldKeyFormat string) (oci_devops.WaitCriteria, error) {
	var baseObject oci_devops.WaitCriteria
	//discriminator
	waitTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_type"))
	var waitType string
	if ok {
		waitType = waitTypeRaw.(string)
	} else {
		waitType = "" // default value
	}
	switch strings.ToLower(waitType) {
	case strings.ToLower("ABSOLUTE_WAIT"):
		details := oci_devops.AbsoluteWaitCriteria{}
		if waitDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_duration")); ok {
			tmp := waitDuration.(string)
			details.WaitDuration = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown wait_type '%v' was specified", waitType)
	}
	return baseObject, nil
}

func (s *DevopsDeployStageResourceCrud) mapToWaitCriteriaSummary(fieldKeyFormat string) (oci_devops.WaitCriteriaSummary, error) {
	var baseObject oci_devops.WaitCriteriaSummary
	//discriminator
	waitTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_type"))
	var waitType string
	if ok {
		waitType = waitTypeRaw.(string)
	} else {
		waitType = "" // default value
	}
	switch strings.ToLower(waitType) {
	case strings.ToLower("ABSOLUTE_WAIT"):
		details := oci_devops.AbsoluteWaitCriteriaSummary{}
		if waitDuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "wait_duration")); ok {
			tmp := waitDuration.(string)
			details.WaitDuration = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown wait_type '%v' was specified", waitType)
	}
	return baseObject, nil
}

func WaitCriteriaSummaryToMap(obj *oci_devops.WaitCriteriaSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.AbsoluteWaitCriteriaSummary:
		result["wait_type"] = "ABSOLUTE_WAIT"

		if v.WaitDuration != nil {
			result["wait_duration"] = string(*v.WaitDuration)
		}
	default:
		log.Printf("[WARN] Received 'wait_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsDeployStageResourceCrud) populateTopLevelPolymorphicCreateDeployStageRequest(request *oci_devops.CreateDeployStageRequest) error {
	//discriminator
	deployStageTypeRaw, ok := s.D.GetOkExists("deploy_stage_type")
	var deployStageType string
	if ok {
		deployStageType = deployStageTypeRaw.(string)
	} else {
		deployStageType = "" // default value
	}
	switch strings.ToLower(deployStageType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT"):
		details := oci_devops.CreateComputeInstanceGroupBlueGreenDeployStageDetails{}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deployEnvironmentIdA, ok := s.D.GetOkExists("deploy_environment_id_a"); ok {
			tmp := deployEnvironmentIdA.(string)
			details.DeployEnvironmentIdA = &tmp
		}
		if deployEnvironmentIdB, ok := s.D.GetOkExists("deploy_environment_id_b"); ok {
			tmp := deployEnvironmentIdB.(string)
			details.DeployEnvironmentIdB = &tmp
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if failurePolicy, ok := s.D.GetOkExists("failure_policy"); ok {
			if tmpList := failurePolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupFailurePolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FailurePolicy = tmp
			}
		}
		if productionLoadBalancerConfig, ok := s.D.GetOkExists("production_load_balancer_config"); ok {
			if tmpList := productionLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "production_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProductionLoadBalancerConfig = &tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if testLoadBalancerConfig, ok := s.D.GetOkExists("test_load_balancer_config"); ok {
			if tmpList := testLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "test_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TestLoadBalancerConfig = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT"):
		details := oci_devops.CreateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails{}
		if computeInstanceGroupBlueGreenDeploymentDeployStageId, ok := s.D.GetOkExists("compute_instance_group_blue_green_deployment_deploy_stage_id"); ok {
			tmp := computeInstanceGroupBlueGreenDeploymentDeployStageId.(string)
			details.ComputeInstanceGroupBlueGreenDeploymentDeployStageId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"):
		details := oci_devops.CreateComputeInstanceGroupCanaryApprovalDeployStageDetails{}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if computeInstanceGroupCanaryTrafficShiftDeployStageId, ok := s.D.GetOkExists("compute_instance_group_canary_traffic_shift_deploy_stage_id"); ok {
			tmp := computeInstanceGroupCanaryTrafficShiftDeployStageId.(string)
			details.ComputeInstanceGroupCanaryTrafficShiftDeployStageId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT"):
		details := oci_devops.CreateComputeInstanceGroupCanaryDeployStageDetails{}
		if computeInstanceGroupDeployEnvironmentId, ok := s.D.GetOkExists("compute_instance_group_deploy_environment_id"); ok {
			tmp := computeInstanceGroupDeployEnvironmentId.(string)
			details.ComputeInstanceGroupDeployEnvironmentId = &tmp
		}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if productionLoadBalancerConfig, ok := s.D.GetOkExists("production_load_balancer_config"); ok {
			if tmpList := productionLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "production_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ProductionLoadBalancerConfig = &tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if testLoadBalancerConfig, ok := s.D.GetOkExists("test_load_balancer_config"); ok {
			if tmpList := testLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "test_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TestLoadBalancerConfig = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT"):
		details := oci_devops.CreateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails{}
		if computeInstanceGroupCanaryDeployStageId, ok := s.D.GetOkExists("compute_instance_group_canary_deploy_stage_id"); ok {
			tmp := computeInstanceGroupCanaryDeployStageId.(string)
			details.ComputeInstanceGroupCanaryDeployStageId = &tmp
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"):
		details := oci_devops.CreateComputeInstanceGroupDeployStageDetails{}
		if computeInstanceGroupDeployEnvironmentId, ok := s.D.GetOkExists("compute_instance_group_deploy_environment_id"); ok {
			tmp := computeInstanceGroupDeployEnvironmentId.(string)
			details.ComputeInstanceGroupDeployEnvironmentId = &tmp
		}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if failurePolicy, ok := s.D.GetOkExists("failure_policy"); ok {
			if tmpList := failurePolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupFailurePolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FailurePolicy = tmp
			}
		}
		if loadBalancerConfig, ok := s.D.GetOkExists("load_balancer_config"); ok {
			if tmpList := loadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.LoadBalancerConfig = &tmp
			}
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("DEPLOY_FUNCTION"):
		details := oci_devops.CreateFunctionDeployStageDetails{}
		if config, ok := s.D.GetOkExists("config"); ok {
			details.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
		}
		if dockerImageDeployArtifactId, ok := s.D.GetOkExists("docker_image_deploy_artifact_id"); ok {
			tmp := dockerImageDeployArtifactId.(string)
			details.DockerImageDeployArtifactId = &tmp
		}
		if functionDeployEnvironmentId, ok := s.D.GetOkExists("function_deploy_environment_id"); ok {
			tmp := functionDeployEnvironmentId.(string)
			details.FunctionDeployEnvironmentId = &tmp
		}
		if functionTimeoutInSeconds, ok := s.D.GetOkExists("function_timeout_in_seconds"); ok {
			tmp := functionTimeoutInSeconds.(int)
			details.FunctionTimeoutInSeconds = &tmp
		}
		if maxMemoryInMBs, ok := s.D.GetOkExists("max_memory_in_mbs"); ok {
			tmp := maxMemoryInMBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return fmt.Errorf("unable to convert maxMemoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaxMemoryInMBs = &tmpInt64
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("INVOKE_FUNCTION"):
		details := oci_devops.CreateInvokeFunctionDeployStageDetails{}
		if deployArtifactId, ok := s.D.GetOkExists("deploy_artifact_id"); ok {
			tmp := deployArtifactId.(string)
			details.DeployArtifactId = &tmp
		}
		if functionDeployEnvironmentId, ok := s.D.GetOkExists("function_deploy_environment_id"); ok {
			tmp := functionDeployEnvironmentId.(string)
			details.FunctionDeployEnvironmentId = &tmp
		}
		if isAsync, ok := s.D.GetOkExists("is_async"); ok {
			tmp := isAsync.(bool)
			details.IsAsync = &tmp
		}
		if isValidationEnabled, ok := s.D.GetOkExists("is_validation_enabled"); ok {
			tmp := isValidationEnabled.(bool)
			details.IsValidationEnabled = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("LOAD_BALANCER_TRAFFIC_SHIFT"):
		details := oci_devops.CreateLoadBalancerTrafficShiftDeployStageDetails{}
		if blueBackendIps, ok := s.D.GetOkExists("blue_backend_ips"); ok {
			if tmpList := blueBackendIps.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "blue_backend_ips", 0)
				tmp, err := s.mapToBackendSetIpCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BlueBackendIps = &tmp
			}
		}
		if greenBackendIps, ok := s.D.GetOkExists("green_backend_ips"); ok {
			if tmpList := greenBackendIps.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "green_backend_ips", 0)
				tmp, err := s.mapToBackendSetIpCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GreenBackendIps = &tmp
			}
		}
		if loadBalancerConfig, ok := s.D.GetOkExists("load_balancer_config"); ok {
			if tmpList := loadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.LoadBalancerConfig = &tmp
			}
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if trafficShiftTarget, ok := s.D.GetOkExists("traffic_shift_target"); ok {
			details.TrafficShiftTarget = oci_devops.LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum(trafficShiftTarget.(string))
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("SHELL"):
		details := oci_devops.CreateShellDeployStageDetails{}
		if commandSpecDeployArtifactId, ok := s.D.GetOkExists("command_spec_deploy_artifact_id"); ok {
			tmp := commandSpecDeployArtifactId.(string)
			details.CommandSpecDeployArtifactId = &tmp
		}
		if containerConfig, ok := s.D.GetOkExists("container_config"); ok {
			if tmpList := containerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "container_config", 0)
				tmp, err := s.mapToContainerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ContainerConfig = tmp
			}
		}
		if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("MANUAL_APPROVAL"):
		details := oci_devops.CreateManualApprovalDeployStageDetails{}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_DEPLOYMENT"):
		details := oci_devops.CreateOkeCanaryDeployStageDetails{}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if canaryStrategy, ok := s.D.GetOkExists("canary_strategy"); ok {
			if tmpList := canaryStrategy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "canary_strategy", 0)
				tmp, err := s.mapToOkeCanaryStrategy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.CanaryStrategy = tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_TRAFFIC_SHIFT"):
		details := oci_devops.CreateOkeCanaryTrafficShiftDeployStageDetails{}
		if okeCanaryDeployStageId, ok := s.D.GetOkExists("oke_canary_deploy_stage_id"); ok {
			tmp := okeCanaryDeployStageId.(string)
			details.OkeCanaryDeployStageId = &tmp
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_APPROVAL"):
		details := oci_devops.CreateOkeCanaryApprovalDeployStageDetails{}
		if okeCanaryTrafficShiftDeployStageId, ok := s.D.GetOkExists("oke_canary_traffic_shift_deploy_stage_id"); ok {
			tmp := okeCanaryTrafficShiftDeployStageId.(string)
			details.OkeCanaryTrafficShiftDeployStageId = &tmp
		}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_DEPLOYMENT"):
		details := oci_devops.CreateOkeDeployStageDetails{}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_BLUE_GREEN_DEPLOYMENT"):
		details := oci_devops.CreateOkeBlueGreenDeployStageDetails{}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if blueGreenStrategy, ok := s.D.GetOkExists("blue_green_strategy"); ok {
			if tmpList := blueGreenStrategy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "blue_green_strategy", 0)
				tmp, err := s.mapToOkeBlueGreenStrategy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BlueGreenStrategy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_BLUE_GREEN_TRAFFIC_SHIFT"):
		details := oci_devops.CreateOkeBlueGreenTrafficShiftDeployStageDetails{}
		if okeBlueGreenDeployStageId, ok := s.D.GetOkExists("oke_blue_green_deploy_stage_id"); ok {
			tmp := okeBlueGreenDeployStageId.(string)
			details.OkeBlueGreenDeployStageId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("WAIT"):
		details := oci_devops.CreateWaitDeployStageDetails{}
		if waitCriteria, ok := s.D.GetOkExists("wait_criteria"); ok {
			if tmpList := waitCriteria.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "wait_criteria", 0)
				tmp, err := s.mapToWaitCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.WaitCriteria = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	case strings.ToLower("OKE_HELM_CHART_DEPLOYMENT"):
		details := oci_devops.CreateOkeHelmChartDeployStageDetails{}
		if areHooksEnabled, ok := s.D.GetOkExists("are_hooks_enabled"); ok {
			tmp := areHooksEnabled.(bool)
			details.AreHooksEnabled = &tmp
		}
		if helmChartDeployArtifactId, ok := s.D.GetOkExists("helm_chart_deploy_artifact_id"); ok {
			tmp := helmChartDeployArtifactId.(string)
			details.HelmChartDeployArtifactId = &tmp
		}

		if helmCommandArtifactIds, ok := s.D.GetOkExists("helm_command_artifact_ids"); ok {
			interfaces := helmCommandArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("helm_command_artifact_ids") {
				details.HelmCommandArtifactIds = tmp
			}
		}

		if isDebugEnabled, ok := s.D.GetOkExists("is_debug_enabled"); ok {
			tmp := isDebugEnabled.(bool)
			details.IsDebugEnabled = &tmp
		}
		if isForceEnabled, ok := s.D.GetOkExists("is_force_enabled"); ok {
			tmp := isForceEnabled.(bool)
			details.IsForceEnabled = &tmp
		}
		if isUninstallOnStageDelete, ok := s.D.GetOkExists("is_uninstall_on_stage_delete"); ok {
			tmp := isUninstallOnStageDelete.(bool)
			details.IsUninstallOnStageDelete = &tmp
		}
		if maxHistory, ok := s.D.GetOkExists("max_history"); ok {
			tmp := maxHistory.(int)
			details.MaxHistory = &tmp
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if purpose, ok := s.D.GetOkExists("purpose"); ok {
			details.Purpose = oci_devops.CreateOkeHelmChartDeployStageDetailsPurposeEnum(purpose.(string))
		}
		if releaseName, ok := s.D.GetOkExists("release_name"); ok {
			tmp := releaseName.(string)
			details.ReleaseName = &tmp
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if setString, ok := s.D.GetOkExists("set_string"); ok {
			if tmpList := setString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "set_string", 0)
				tmp, err := s.mapToHelmSetValueCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				if tmp.Items != nil && len(tmp.Items) > 0 {
					details.SetString = &tmp
				}
			}
		}
		if setValues, ok := s.D.GetOkExists("set_values"); ok {
			if tmpList := setValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "set_values", 0)
				tmp, err := s.mapToHelmSetValueCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				if tmp.Items != nil && len(tmp.Items) > 0 {
					details.SetValues = &tmp
				}
			}
		}
		if shouldCleanupOnFail, ok := s.D.GetOkExists("should_cleanup_on_fail"); ok {
			tmp := shouldCleanupOnFail.(bool)
			details.ShouldCleanupOnFail = &tmp
		}
		if shouldNotWait, ok := s.D.GetOkExists("should_not_wait"); ok {
			tmp := shouldNotWait.(bool)
			details.ShouldNotWait = &tmp
		}
		if shouldResetValues, ok := s.D.GetOkExists("should_reset_values"); ok {
			tmp := shouldResetValues.(bool)
			details.ShouldResetValues = &tmp
		}
		if shouldReuseValues, ok := s.D.GetOkExists("should_reuse_values"); ok {
			tmp := shouldReuseValues.(bool)
			details.ShouldReuseValues = &tmp
		}
		if shouldSkipCrds, ok := s.D.GetOkExists("should_skip_crds"); ok {
			tmp := shouldSkipCrds.(bool)
			details.ShouldSkipCrds = &tmp
		}
		if shouldSkipRenderSubchartNotes, ok := s.D.GetOkExists("should_skip_render_subchart_notes"); ok {
			tmp := shouldSkipRenderSubchartNotes.(bool)
			details.ShouldSkipRenderSubchartNotes = &tmp
		}
		if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		if valuesArtifactIds, ok := s.D.GetOkExists("values_artifact_ids"); ok {
			interfaces := valuesArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("values_artifact_ids") {
				details.ValuesArtifactIds = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeployStageDetails = details
	default:
		return fmt.Errorf("unknown deploy_stage_type '%v' was specified", deployStageType)
	}
	return nil
}

func (s *DevopsDeployStageResourceCrud) populateTopLevelPolymorphicUpdateDeployStageRequest(request *oci_devops.UpdateDeployStageRequest) error {
	//discriminator
	deployStageTypeRaw, ok := s.D.GetOkExists("deploy_stage_type")
	var deployStageType string
	if ok {
		deployStageType = deployStageTypeRaw.(string)
	} else {
		deployStageType = "" // default value
	}
	switch strings.ToLower(deployStageType) {
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_BLUE_GREEN_DEPLOYMENT"):
		details := oci_devops.UpdateComputeInstanceGroupBlueGreenDeployStageDetails{}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if failurePolicy, ok := s.D.GetOkExists("failure_policy"); ok {
			if tmpList := failurePolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupFailurePolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FailurePolicy = tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if testLoadBalancerConfig, ok := s.D.GetOkExists("test_load_balancer_config"); ok {
			if tmpList := testLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "test_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TestLoadBalancerConfig = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_BLUE_GREEN_TRAFFIC_SHIFT"):
		details := oci_devops.UpdateComputeInstanceGroupBlueGreenTrafficShiftDeployStageDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_APPROVAL"):
		details := oci_devops.UpdateComputeInstanceGroupCanaryApprovalDeployStageDetails{}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_DEPLOYMENT"):
		details := oci_devops.UpdateComputeInstanceGroupCanaryDeployStageDetails{}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if testLoadBalancerConfig, ok := s.D.GetOkExists("test_load_balancer_config"); ok {
			if tmpList := testLoadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "test_load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.TestLoadBalancerConfig = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_CANARY_TRAFFIC_SHIFT"):
		details := oci_devops.UpdateComputeInstanceGroupCanaryTrafficShiftDeployStageDetails{}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("COMPUTE_INSTANCE_GROUP_ROLLING_DEPLOYMENT"):
		details := oci_devops.UpdateComputeInstanceGroupDeployStageDetails{}
		if computeInstanceGroupDeployEnvironmentId, ok := s.D.GetOkExists("compute_instance_group_deploy_environment_id"); ok {
			tmp := computeInstanceGroupDeployEnvironmentId.(string)
			details.ComputeInstanceGroupDeployEnvironmentId = &tmp
		}
		if deployArtifactIds, ok := s.D.GetOkExists("deploy_artifact_ids"); ok {
			interfaces := deployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("deploy_artifact_ids") {
				details.DeployArtifactIds = tmp
			}
		}
		if deploymentSpecDeployArtifactId, ok := s.D.GetOkExists("deployment_spec_deploy_artifact_id"); ok {
			tmp := deploymentSpecDeployArtifactId.(string)
			details.DeploymentSpecDeployArtifactId = &tmp
		}
		if failurePolicy, ok := s.D.GetOkExists("failure_policy"); ok {
			if tmpList := failurePolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "failure_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupFailurePolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.FailurePolicy = tmp
			}
		}
		if loadBalancerConfig, ok := s.D.GetOkExists("load_balancer_config"); ok {
			if tmpList := loadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.LoadBalancerConfig = &tmp
			}
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToComputeInstanceGroupRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("DEPLOY_FUNCTION"):
		details := oci_devops.UpdateFunctionDeployStageDetails{}
		if config, ok := s.D.GetOkExists("config"); ok {
			details.Config = tfresource.ObjectMapToStringMap(config.(map[string]interface{}))
		}
		if dockerImageDeployArtifactId, ok := s.D.GetOkExists("docker_image_deploy_artifact_id"); ok {
			tmp := dockerImageDeployArtifactId.(string)
			details.DockerImageDeployArtifactId = &tmp
		}
		if functionDeployEnvironmentId, ok := s.D.GetOkExists("function_deploy_environment_id"); ok {
			tmp := functionDeployEnvironmentId.(string)
			details.FunctionDeployEnvironmentId = &tmp
		}
		if functionTimeoutInSeconds, ok := s.D.GetOkExists("function_timeout_in_seconds"); ok {
			tmp := functionTimeoutInSeconds.(int)
			details.FunctionTimeoutInSeconds = &tmp
		}
		if maxMemoryInMBs, ok := s.D.GetOkExists("max_memory_in_mbs"); ok {
			tmp := maxMemoryInMBs.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return fmt.Errorf("unable to convert maxMemoryInMBs string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.MaxMemoryInMBs = &tmpInt64
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("INVOKE_FUNCTION"):
		details := oci_devops.UpdateInvokeFunctionDeployStageDetails{}
		if deployArtifactId, ok := s.D.GetOkExists("deploy_artifact_id"); ok {
			tmp := deployArtifactId.(string)
			details.DeployArtifactId = &tmp
		}
		if functionDeployEnvironmentId, ok := s.D.GetOkExists("function_deploy_environment_id"); ok {
			tmp := functionDeployEnvironmentId.(string)
			details.FunctionDeployEnvironmentId = &tmp
		}
		if isAsync, ok := s.D.GetOkExists("is_async"); ok {
			tmp := isAsync.(bool)
			details.IsAsync = &tmp
		}
		if isValidationEnabled, ok := s.D.GetOkExists("is_validation_enabled"); ok {
			tmp := isValidationEnabled.(bool)
			details.IsValidationEnabled = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("LOAD_BALANCER_TRAFFIC_SHIFT"):
		details := oci_devops.UpdateLoadBalancerTrafficShiftDeployStageDetails{}
		if blueBackendIps, ok := s.D.GetOkExists("blue_backend_ips"); ok {
			if tmpList := blueBackendIps.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "blue_backend_ips", 0)
				tmp, err := s.mapToBackendSetIpCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BlueBackendIps = &tmp
			}
		}
		if greenBackendIps, ok := s.D.GetOkExists("green_backend_ips"); ok {
			if tmpList := greenBackendIps.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "green_backend_ips", 0)
				tmp, err := s.mapToBackendSetIpCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.GreenBackendIps = &tmp
			}
		}
		if loadBalancerConfig, ok := s.D.GetOkExists("load_balancer_config"); ok {
			if tmpList := loadBalancerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "load_balancer_config", 0)
				tmp, err := s.mapToLoadBalancerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.LoadBalancerConfig = &tmp
			}
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if trafficShiftTarget, ok := s.D.GetOkExists("traffic_shift_target"); ok {
			details.TrafficShiftTarget = oci_devops.LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum(trafficShiftTarget.(string))
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("SHELL"):
		details := oci_devops.UpdateShellDeployStageDetails{}
		if commandSpecDeployArtifactId, ok := s.D.GetOkExists("command_spec_deploy_artifact_id"); ok {
			tmp := commandSpecDeployArtifactId.(string)
			details.CommandSpecDeployArtifactId = &tmp
		}
		if containerConfig, ok := s.D.GetOkExists("container_config"); ok {
			if tmpList := containerConfig.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "container_config", 0)
				tmp, err := s.mapToContainerConfig(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ContainerConfig = tmp
			}
		}
		if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("MANUAL_APPROVAL"):
		details := oci_devops.UpdateManualApprovalDeployStageDetails{}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_BLUE_GREEN_DEPLOYMENT"):
		details := oci_devops.UpdateOkeBlueGreenDeployStageDetails{}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_BLUE_GREEN_TRAFFIC_SHIFT"):
		details := oci_devops.UpdateOkeBlueGreenTrafficShiftDeployStageDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_APPROVAL"):
		details := oci_devops.UpdateOkeCanaryApprovalDeployStageDetails{}
		if approvalPolicy, ok := s.D.GetOkExists("approval_policy"); ok {
			if tmpList := approvalPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "approval_policy", 0)
				tmp, err := s.mapToApprovalPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ApprovalPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_DEPLOYMENT"):
		details := oci_devops.UpdateOkeCanaryDeployStageDetails{}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_CANARY_TRAFFIC_SHIFT"):
		details := oci_devops.UpdateOkeCanaryTrafficShiftDeployStageDetails{}
		if rolloutPolicy, ok := s.D.GetOkExists("rollout_policy"); ok {
			if tmpList := rolloutPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollout_policy", 0)
				tmp, err := s.mapToLoadBalancerTrafficShiftRolloutPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RolloutPolicy = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_DEPLOYMENT"):
		details := oci_devops.UpdateOkeDeployStageDetails{}
		if kubernetesManifestDeployArtifactIds, ok := s.D.GetOkExists("kubernetes_manifest_deploy_artifact_ids"); ok {
			interfaces := kubernetesManifestDeployArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("kubernetes_manifest_deploy_artifact_ids") {
				details.KubernetesManifestDeployArtifactIds = tmp
			}
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("WAIT"):
		details := oci_devops.UpdateWaitDeployStageDetails{}
		if waitCriteria, ok := s.D.GetOkExists("wait_criteria"); ok {
			if tmpList := waitCriteria.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "wait_criteria", 0)
				tmp, err := s.mapToWaitCriteria(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.WaitCriteria = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	case strings.ToLower("OKE_HELM_CHART_DEPLOYMENT"):
		details := oci_devops.UpdateOkeHelmChartDeployStageDetails{}
		if areHooksEnabled, ok := s.D.GetOkExists("are_hooks_enabled"); ok {
			tmp := areHooksEnabled.(bool)
			details.AreHooksEnabled = &tmp
		}
		if helmChartDeployArtifactId, ok := s.D.GetOkExists("helm_chart_deploy_artifact_id"); ok {
			tmp := helmChartDeployArtifactId.(string)
			details.HelmChartDeployArtifactId = &tmp
		}
		if helmCommandArtifactIds, ok := s.D.GetOkExists("helm_command_artifact_ids"); ok {
			interfaces := helmCommandArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("helm_command_artifact_ids") {
				details.HelmCommandArtifactIds = tmp
			}
		}
		if isDebugEnabled, ok := s.D.GetOkExists("is_debug_enabled"); ok {
			tmp := isDebugEnabled.(bool)
			details.IsDebugEnabled = &tmp
		}
		if isForceEnabled, ok := s.D.GetOkExists("is_force_enabled"); ok {
			tmp := isForceEnabled.(bool)
			details.IsForceEnabled = &tmp
		}
		if isUninstallOnStageDelete, ok := s.D.GetOkExists("is_uninstall_on_stage_delete"); ok {
			tmp := isUninstallOnStageDelete.(bool)
			details.IsUninstallOnStageDelete = &tmp
		}
		if maxHistory, ok := s.D.GetOkExists("max_history"); ok {
			tmp := maxHistory.(int)
			details.MaxHistory = &tmp
		}
		if namespace, ok := s.D.GetOkExists("namespace"); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if okeClusterDeployEnvironmentId, ok := s.D.GetOkExists("oke_cluster_deploy_environment_id"); ok {
			tmp := okeClusterDeployEnvironmentId.(string)
			details.OkeClusterDeployEnvironmentId = &tmp
		}
		if purpose, ok := s.D.GetOkExists("purpose"); ok {
			details.Purpose = oci_devops.UpdateOkeHelmChartDeployStageDetailsPurposeEnum(purpose.(string))
		}
		if releaseName, ok := s.D.GetOkExists("release_name"); ok {
			tmp := releaseName.(string)
			details.ReleaseName = &tmp
		}
		if rollbackPolicy, ok := s.D.GetOkExists("rollback_policy"); ok {
			if tmpList := rollbackPolicy.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "rollback_policy", 0)
				tmp, err := s.mapToDeployStageRollbackPolicy(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.RollbackPolicy = tmp
			}
		}
		if setString, ok := s.D.GetOkExists("set_string"); ok {
			if tmpList := setString.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "set_string", 0)
				tmp, err := s.mapToHelmSetValueCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				if tmp.Items != nil && len(tmp.Items) > 0 {
					details.SetString = &tmp
				}
			}
		}
		if setValues, ok := s.D.GetOkExists("set_values"); ok {
			if tmpList := setValues.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "set_values", 0)
				tmp, err := s.mapToHelmSetValueCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				if tmp.Items != nil && len(tmp.Items) > 0 {
					details.SetValues = &tmp
				}
			}
		}
		if shouldCleanupOnFail, ok := s.D.GetOkExists("should_cleanup_on_fail"); ok {
			tmp := shouldCleanupOnFail.(bool)
			details.ShouldCleanupOnFail = &tmp
		}
		if shouldNotWait, ok := s.D.GetOkExists("should_not_wait"); ok {
			tmp := shouldNotWait.(bool)
			details.ShouldNotWait = &tmp
		}
		if shouldResetValues, ok := s.D.GetOkExists("should_reset_values"); ok {
			tmp := shouldResetValues.(bool)
			details.ShouldResetValues = &tmp
		}
		if shouldReuseValues, ok := s.D.GetOkExists("should_reuse_values"); ok {
			tmp := shouldReuseValues.(bool)
			details.ShouldReuseValues = &tmp
		}
		if shouldSkipCrds, ok := s.D.GetOkExists("should_skip_crds"); ok {
			tmp := shouldSkipCrds.(bool)
			details.ShouldSkipCrds = &tmp
		}
		if shouldSkipRenderSubchartNotes, ok := s.D.GetOkExists("should_skip_render_subchart_notes"); ok {
			tmp := shouldSkipRenderSubchartNotes.(bool)
			details.ShouldSkipRenderSubchartNotes = &tmp
		}
		if timeoutInSeconds, ok := s.D.GetOkExists("timeout_in_seconds"); ok {
			tmp := timeoutInSeconds.(int)
			details.TimeoutInSeconds = &tmp
		}
		if valuesArtifactIds, ok := s.D.GetOkExists("values_artifact_ids"); ok {
			interfaces := valuesArtifactIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("values_artifact_ids") {
				details.ValuesArtifactIds = tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeployStageId = &tmp
		if deployStagePredecessorCollection, ok := s.D.GetOkExists("deploy_stage_predecessor_collection"); ok {
			if tmpList := deployStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_predecessor_collection", 0)
				tmp, err := s.mapToDeployStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStagePredecessorCollection = &tmp
			}
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeployStageDetails = details
	default:
		return fmt.Errorf("unknown deploy_stage_type '%v' was specified", deployStageType)
	}
	return nil
}
