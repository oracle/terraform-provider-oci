// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsDeploymentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		Create:        createDevopsDeployment,
		Read:          readDevopsDeployment,
		Update:        updateDevopsDeployment,
		Delete:        deleteDevopsDeployment,
		CustomizeDiff: resourceOkeClusterHelmReleaseDiff,
		Schema: map[string]*schema.Schema{
			// Required
			"deploy_pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"deployment_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"PIPELINE_DEPLOYMENT",
					"PIPELINE_REDEPLOYMENT",
					"SINGLE_STAGE_DEPLOYMENT",
					"SINGLE_STAGE_REDEPLOYMENT",
				}, true),
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"deploy_artifact_override_arguments": {
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
						"items": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"deploy_artifact_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
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

						// Computed
					},
				},
			},
			"deploy_stage_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"deploy_stage_override_arguments": {
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
						"items": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"deploy_stage_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},
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

						// Computed
					},
				},
			},
			"deployment_arguments": {
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
						"items": {
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

						// Computed
					},
				},
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
			"previous_deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"trigger_new_devops_deployment": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deploy_pipeline_artifacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"deploy_artifact_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_pipeline_stages": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"deploy_stage_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"deploy_pipeline_environments": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"deploy_environment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_pipeline_stages": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"items": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"deploy_stage_id": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"display_name": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"deployment_execution_progress": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"deploy_stage_execution_progress": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"deploy_stage_display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_stage_execution_progress_details": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"target_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"target_group": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"steps": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"state": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"time_started": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"time_finished": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"rollback_steps": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"state": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"time_started": {
																Type:     schema.TypeString,
																Computed: true,
															},
															"time_finished": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"deploy_stage_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_stage_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"deploy_stage_predecessors": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"deploy_stage_predecessor": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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

func createDevopsDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	checkForHydrationWorkRequest(d, m)
	return tfresource.CreateResource(d, sync)
}

func readDevopsDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsDeployment(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsDeploymentResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsDeployment(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DevopsDeploymentResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.Deployment
	DisableNotFoundRetries bool
}

func (s *DevopsDeploymentResourceCrud) ID() string {
	deployment := *s.Res
	return *deployment.GetId()
}

func (s *DevopsDeploymentResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.DeploymentLifecycleStateInProgress),
		string(oci_devops.DeploymentLifecycleStateAccepted),
	}
}

func (s *DevopsDeploymentResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.DeploymentLifecycleStateSucceeded),
	}
}

func (s *DevopsDeploymentResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DevopsDeploymentResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DevopsDeploymentResourceCrud) Create() error {
	request := oci_devops.CreateDeploymentRequest{}
	err := s.populateTopLevelPolymorphicCreateDeploymentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Deployment
	return nil
}

func (s *DevopsDeploymentResourceCrud) Get() error {
	request := oci_devops.GetDeploymentRequest{}

	tmp := s.D.Id()
	request.DeploymentId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Deployment
	return nil
}

func (s *DevopsDeploymentResourceCrud) Update() error {
	request := oci_devops.UpdateDeploymentRequest{}
	err := s.populateTopLevelPolymorphicUpdateDeploymentRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateDeployment(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Deployment
	return nil
}

func (s *DevopsDeploymentResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.DeployPipelineDeployment:
		s.D.Set("deployment_type", "PIPELINE_DEPLOYMENT")

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStageOverrideArguments != nil {
			s.D.Set("deploy_stage_override_arguments", []interface{}{DeployStageOverrideArgumentCollectionToMap(v.DeployStageOverrideArguments)})
		} else {
			s.D.Set("deploy_stage_override_arguments", nil)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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

		s.D.Set("trigger_new_devops_deployment", false)
	case oci_devops.DeployPipelineRedeployment:
		s.D.Set("deployment_type", "PIPELINE_REDEPLOYMENT")

		if v.PreviousDeploymentId != nil {
			s.D.Set("previous_deployment_id", *v.PreviousDeploymentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStageOverrideArguments != nil {
			s.D.Set("deploy_stage_override_arguments", []interface{}{DeployStageOverrideArgumentCollectionToMap(v.DeployStageOverrideArguments)})
		} else {
			s.D.Set("deploy_stage_override_arguments", nil)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}
		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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

		s.D.Set("trigger_new_devops_deployment", false)
	case oci_devops.SingleDeployStageDeployment:
		s.D.Set("deployment_type", "SINGLE_STAGE_DEPLOYMENT")

		if v.DeployStageId != nil {
			s.D.Set("deploy_stage_id", *v.DeployStageId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStageOverrideArguments != nil {
			s.D.Set("deploy_stage_override_arguments", []interface{}{DeployStageOverrideArgumentCollectionToMap(v.DeployStageOverrideArguments)})
		} else {
			s.D.Set("deploy_stage_override_arguments", nil)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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

		s.D.Set("trigger_new_devops_deployment", false)
	case oci_devops.SingleDeployStageRedeployment:
		s.D.Set("deployment_type", "SINGLE_STAGE_REDEPLOYMENT")

		if v.DeployStageId != nil {
			s.D.Set("deploy_stage_id", *v.DeployStageId)
		}

		if v.PreviousDeploymentId != nil {
			s.D.Set("previous_deployment_id", *v.PreviousDeploymentId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DeployArtifactOverrideArguments != nil {
			s.D.Set("deploy_artifact_override_arguments", []interface{}{DeployArtifactOverrideArgumentCollectionToMap(v.DeployArtifactOverrideArguments)})
		} else {
			s.D.Set("deploy_artifact_override_arguments", nil)
		}

		if v.DeployPipelineArtifacts != nil {
			s.D.Set("deploy_pipeline_artifacts", []interface{}{DeployPipelineArtifactCollectionToMap(v.DeployPipelineArtifacts)})
		} else {
			s.D.Set("deploy_pipeline_artifacts", nil)
		}

		if v.DeployPipelineEnvironments != nil {
			s.D.Set("deploy_pipeline_environments", []interface{}{DeployPipelineEnvironmentCollectionToMap(v.DeployPipelineEnvironments)})
		} else {
			s.D.Set("deploy_pipeline_environments", nil)
		}

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.DeployStageOverrideArguments != nil {
			s.D.Set("deploy_stage_override_arguments", []interface{}{DeployStageOverrideArgumentCollectionToMap(v.DeployStageOverrideArguments)})
		} else {
			s.D.Set("deploy_stage_override_arguments", nil)
		}

		if v.DeploymentArguments != nil {
			s.D.Set("deployment_arguments", []interface{}{DeploymentArgumentCollectionToMap(v.DeploymentArguments)})
		} else {
			s.D.Set("deployment_arguments", nil)
		}

		if v.DeploymentExecutionProgress != nil {
			s.D.Set("deployment_execution_progress", []interface{}{DeploymentExecutionProgressToMap(v.DeploymentExecutionProgress)})
		} else {
			s.D.Set("deployment_execution_progress", nil)
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

		s.D.Set("trigger_new_devops_deployment", false)
	default:
		log.Printf("[WARN] Received 'deployment_type' of unknown type %v", *s.Res)
		return nil
	}

	return nil
}

func (s *DevopsDeploymentResourceCrud) mapToApprovalAction(fieldKeyFormat string) (oci_devops.ApprovalAction, error) {
	result := oci_devops.ApprovalAction{}

	if action, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "action")); ok {
		result.Action = oci_devops.ApprovalActionActionEnum(action.(string))
	}

	if reason, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reason")); ok {
		tmp := reason.(string)
		result.Reason = &tmp
	}

	if subjectId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subject_id")); ok {
		tmp := subjectId.(string)
		result.SubjectId = &tmp
	}

	return result, nil
}

func ApprovalActionToMap(obj oci_devops.ApprovalAction) map[string]interface{} {
	result := map[string]interface{}{}

	result["action"] = string(obj.Action)

	if obj.Reason != nil {
		result["reason"] = string(*obj.Reason)
	}

	if obj.SubjectId != nil {
		result["subject_id"] = string(*obj.SubjectId)
	}

	return result
}

func (s *DevopsDeploymentResourceCrud) mapToDeployArtifactOverrideArgument(fieldKeyFormat string) (oci_devops.DeployArtifactOverrideArgument, error) {
	result := oci_devops.DeployArtifactOverrideArgument{}

	if deployArtifactId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_id")); ok {
		tmp := deployArtifactId.(string)
		result.DeployArtifactId = &tmp
	}

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

func (s *DevopsDeploymentResourceCrud) mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat string) (oci_devops.DeployArtifactOverrideArgumentCollection, error) {
	result := oci_devops.DeployArtifactOverrideArgumentCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeployArtifactOverrideArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeployArtifactOverrideArgument(fieldKeyFormatNextLevel)
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

func DeployStageExecutionProgressDetailsToMap(obj oci_devops.DeployStageExecutionProgressDetails) map[string]interface{} {
	result := map[string]interface{}{}

	rollbackSteps := []interface{}{}
	for _, item := range obj.RollbackSteps {
		rollbackSteps = append(rollbackSteps, DeployStageExecutionStepToMap(item))
	}
	result["rollback_steps"] = rollbackSteps

	steps := []interface{}{}
	for _, item := range obj.Steps {
		steps = append(steps, DeployStageExecutionStepToMap(item))
	}
	result["steps"] = steps

	if obj.TargetGroup != nil {
		result["target_group"] = string(*obj.TargetGroup)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}

func DeployStageExecutionStepToMap(obj oci_devops.DeployStageExecutionStep) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.State)

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func DeploymentExecutionProgressToMap(obj *oci_devops.DeploymentExecutionProgress) map[string]interface{} {
	result := map[string]interface{}{}
	deployStageExecutionProgressInfo := []interface{}{}

	for _, item := range obj.DeployStageExecutionProgress {
		deployStageExecutionProgressInfo = append(deployStageExecutionProgressInfo, DeploymentStageExecutionProgressToMap(item))
	}

	result["deploy_stage_execution_progress"] = deployStageExecutionProgressInfo

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func DeploymentStageExecutionProgressToMap(obj oci_devops.DeployStageExecutionProgress) map[string]interface{} {
	result := map[string]interface{}{}

	deployStageExecutionProgressDetails := []interface{}{}

	for _, item := range obj.GetDeployStageExecutionProgressDetails() {
		deployStageExecutionProgressDetails = append(deployStageExecutionProgressDetails, DeployStageExecutionProgressDetailsToMap(item))
	}

	result["deploy_stage_execution_progress_details"] = deployStageExecutionProgressDetails

	if obj.GetDeployStageDisplayName() != nil {
		result["deploy_stage_display_name"] = string(*obj.GetDeployStageDisplayName())
	}

	if obj.GetDeployStageId() != nil {
		result["deploy_stage_id"] = string(*obj.GetDeployStageId())
	}

	if obj.GetTimeFinished() != nil {
		result["time_finished"] = obj.GetTimeFinished().String()
	}

	if obj.GetTimeStarted() != nil {
		result["time_started"] = obj.GetTimeStarted().String()
	}

	if obj.GetDeployStagePredecessors() != nil {
		result["deploy_stage_predecessors"] = obj.GetDeployStagePredecessors
	}

	result["status"] = obj.GetStatus()

	return result
}

func (s *DevopsDeploymentResourceCrud) mapToDeployStageOverrideArgument(fieldKeyFormat string) (oci_devops.DeployStageOverrideArgument, error) {
	result := oci_devops.DeployStageOverrideArgument{}

	if deployStageId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_stage_id")); ok {
		tmp := deployStageId.(string)
		result.DeployStageId = &tmp
	}

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

func DeployStageOverrideArgumentToMap(obj oci_devops.DeployStageOverrideArgument) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeployStageId != nil {
		result["deploy_stage_id"] = string(*obj.DeployStageId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DevopsDeploymentResourceCrud) mapToDeployStageOverrideArgumentCollection(fieldKeyFormat string) (oci_devops.DeployStageOverrideArgumentCollection, error) {
	result := oci_devops.DeployStageOverrideArgumentCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeployStageOverrideArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeployStageOverrideArgument(fieldKeyFormatNextLevel)
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

func DeployStageOverrideArgumentCollectionToMap(obj *oci_devops.DeployStageOverrideArgumentCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeployStageOverrideArgumentToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsDeploymentResourceCrud) mapToDeploymentArgument(fieldKeyFormat string) (oci_devops.DeploymentArgument, error) {
	result := oci_devops.DeploymentArgument{}

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

func DeploymentArgumentToMap(obj oci_devops.DeploymentArgument) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DevopsDeploymentResourceCrud) mapToDeploymentArgumentCollection(fieldKeyFormat string) (oci_devops.DeploymentArgumentCollection, error) {
	result := oci_devops.DeploymentArgumentCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeploymentArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeploymentArgument(fieldKeyFormatNextLevel)
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

func DeploymentArgumentCollectionToMap(obj *oci_devops.DeploymentArgumentCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeploymentArgumentToMap(item))
	}
	result["items"] = items

	return result
}

func DevopsDeploymentSummaryToMap(obj oci_devops.DeploymentSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.DeployPipelineDeploymentSummary:
		result["deployment_type"] = "PIPELINE_DEPLOYMENT"
	case oci_devops.DeployPipelineRedeploymentSummary:
		result["deployment_type"] = "PIPELINE_REDEPLOYMENT"

		if v.PreviousDeploymentId != nil {
			result["previous_deployment_id"] = string(*v.PreviousDeploymentId)
		}
	case oci_devops.SingleDeployStageDeploymentSummary:
		result["deployment_type"] = "SINGLE_STAGE_DEPLOYMENT"

		if v.DeployStageId != nil {
			result["deploy_stage_id"] = string(*v.DeployStageId)
		}
	case oci_devops.SingleDeployStageRedeploymentSummary:
		result["deployment_type"] = "SINGLE_STAGE_REDEPLOYMENT"

		if v.DeployStageId != nil {
			result["deploy_stage_id"] = string(*v.DeployStageId)
		}

		if v.PreviousDeploymentId != nil {
			result["previous_deployment_id"] = string(*v.PreviousDeploymentId)
		}
	default:
		log.Printf("[WARN] Received 'deployment_type' of unknown type %v", obj)
		return nil
	}

	if obj.GetId() != nil {
		result["id"] = obj.GetId()
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = obj.GetDisplayName()
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = obj.GetCompartmentId()
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

	if obj.GetDeployArtifactOverrideArguments() != nil {
		result["deploy_artifact_override_arguments"] = DeployArtifactOverrideArgumentCollectionToMap(obj.GetDeployArtifactOverrideArguments())
	} else {
		result["deploy_artifact_override_arguments"] = nil
	}

	//if obj.GetDeploymentArguments() != nil {
	//	result["deployment_arguments"] = DeploymentArgumentCollectionToMap(obj.GetDeploymentArguments())
	//} else {
	//	result["deployment_arguments"] = nil
	//}

	if obj.GetDeployPipelineId() != nil {
		result["deploy_pipeline_id"] = obj.GetDeployPipelineId()
	}

	return result
}

func (s *DevopsDeploymentResourceCrud) populateTopLevelPolymorphicCreateDeploymentRequest(request *oci_devops.CreateDeploymentRequest) error {
	//discriminator
	deploymentTypeRaw, ok := s.D.GetOkExists("deployment_type")
	var deploymentType string
	if ok {
		deploymentType = deploymentTypeRaw.(string)
	} else {
		deploymentType = "" // default value
	}
	switch strings.ToLower(deploymentType) {
	case strings.ToLower("PIPELINE_DEPLOYMENT"):
		details := oci_devops.CreateDeployPipelineDeploymentDetails{}
		if deployArtifactOverrideArguments, ok := s.D.GetOkExists("deploy_artifact_override_arguments"); ok {
			if tmpList := deployArtifactOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_override_arguments", 0)
				tmp, err := s.mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployArtifactOverrideArguments = &tmp
			}
		}
		if deploymentArguments, ok := s.D.GetOkExists("deployment_arguments"); ok {
			if tmpList := deploymentArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deployment_arguments", 0)
				tmp, err := s.mapToDeploymentArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeploymentArguments = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployArtifactOverrideArguments, ok := s.D.GetOkExists("deploy_artifact_override_arguments"); ok {
			if tmpList := deployArtifactOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_override_arguments", 0)
				tmp, err := s.mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployArtifactOverrideArguments = &tmp
			}
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStageOverrideArguments, ok := s.D.GetOkExists("deploy_stage_override_arguments"); ok {
			if tmpList := deployStageOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_override_arguments", 0)
				tmp, err := s.mapToDeployStageOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStageOverrideArguments = &tmp
			}
		}
		if deploymentArguments, ok := s.D.GetOkExists("deployment_arguments"); ok {
			if tmpList := deploymentArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deployment_arguments", 0)
				tmp, err := s.mapToDeploymentArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeploymentArguments = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeploymentDetails = details
	case strings.ToLower("PIPELINE_REDEPLOYMENT"):
		details := oci_devops.CreateDeployPipelineRedeploymentDetails{}
		if previousDeploymentId, ok := s.D.GetOkExists("previous_deployment_id"); ok {
			tmp := previousDeploymentId.(string)
			details.PreviousDeploymentId = &tmp
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeploymentDetails = details
	case strings.ToLower("SINGLE_STAGE_DEPLOYMENT"):
		details := oci_devops.CreateSingleDeployStageDeploymentDetails{}
		if deployArtifactOverrideArguments, ok := s.D.GetOkExists("deploy_artifact_override_arguments"); ok {
			if tmpList := deployArtifactOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_override_arguments", 0)
				tmp, err := s.mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployArtifactOverrideArguments = &tmp
			}
		}
		if deployStageId, ok := s.D.GetOkExists("deploy_stage_id"); ok {
			tmp := deployStageId.(string)
			details.DeployStageId = &tmp
		}
		if deploymentArguments, ok := s.D.GetOkExists("deployment_arguments"); ok {
			if tmpList := deploymentArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deployment_arguments", 0)
				tmp, err := s.mapToDeploymentArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeploymentArguments = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if deployArtifactOverrideArguments, ok := s.D.GetOkExists("deploy_artifact_override_arguments"); ok {
			if tmpList := deployArtifactOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_artifact_override_arguments", 0)
				tmp, err := s.mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployArtifactOverrideArguments = &tmp
			}
		}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if deployStageOverrideArguments, ok := s.D.GetOkExists("deploy_stage_override_arguments"); ok {
			if tmpList := deployStageOverrideArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deploy_stage_override_arguments", 0)
				tmp, err := s.mapToDeployStageOverrideArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeployStageOverrideArguments = &tmp
			}
		}
		if deploymentArguments, ok := s.D.GetOkExists("deployment_arguments"); ok {
			if tmpList := deploymentArguments.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deployment_arguments", 0)
				tmp, err := s.mapToDeploymentArgumentCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeploymentArguments = &tmp
			}
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeploymentDetails = details
	case strings.ToLower("SINGLE_STAGE_REDEPLOYMENT"):
		details := oci_devops.CreateSingleDeployStageRedeploymentDetails{}
		if deployStageId, ok := s.D.GetOkExists("deploy_stage_id"); ok {
			tmp := deployStageId.(string)
			details.DeployStageId = &tmp
		}
		if previousDeploymentId, ok := s.D.GetOkExists("previous_deployment_id"); ok {
			tmp := previousDeploymentId.(string)
			details.PreviousDeploymentId = &tmp
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
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateDeploymentDetails = details
	default:
		return fmt.Errorf("unknown deployment_type '%v' was specified", deploymentType)
	}
	return nil
}

func (s *DevopsDeploymentResourceCrud) populateTopLevelPolymorphicUpdateDeploymentRequest(request *oci_devops.UpdateDeploymentRequest) error {
	//discriminator
	deploymentTypeRaw, ok := s.D.GetOkExists("deployment_type")
	var deploymentType string
	if ok {
		deploymentType = deploymentTypeRaw.(string)
	} else {
		deploymentType = "" // default value
	}
	switch strings.ToLower(deploymentType) {
	case strings.ToLower("PIPELINE_DEPLOYMENT"):
		details := oci_devops.UpdateDeployPipelineDeploymentDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeploymentId = &tmp
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeploymentDetails = details
	case strings.ToLower("PIPELINE_REDEPLOYMENT"):
		details := oci_devops.UpdateDeployPipelineRedeploymentDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeploymentId = &tmp
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeploymentDetails = details
	case strings.ToLower("SINGLE_STAGE_DEPLOYMENT"):
		details := oci_devops.UpdateSingleDeployStageDeploymentDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeploymentId = &tmp
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeploymentDetails = details
	case strings.ToLower("SINGLE_STAGE_REDEPLOYMENT"):
		details := oci_devops.UpdateSingleDeployStageRedeploymentDetails{}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		tmp := s.D.Id()
		request.DeploymentId = &tmp
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateDeploymentDetails = details
	default:
		return fmt.Errorf("unknown deployment_type '%v' was specified", deploymentType)
	}
	return nil
}
