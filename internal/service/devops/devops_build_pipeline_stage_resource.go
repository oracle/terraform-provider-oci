// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

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

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsBuildPipelineStageResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsBuildPipelineStage,
		Read:     readDevopsBuildPipelineStage,
		Update:   updateDevopsBuildPipelineStage,
		Delete:   deleteDevopsBuildPipelineStage,
		Schema: map[string]*schema.Schema{
			// Required
			"build_pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"build_pipeline_stage_predecessor_collection": {
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
			"build_pipeline_stage_type": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"BUILD",
					"DELIVER_ARTIFACT",
					"TRIGGER_DEPLOYMENT_PIPELINE",
					"WAIT",
				}, true),
			},

			// Optional
			"build_source_collection": {
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"connection_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"DEVOPS_CODE_REPOSITORY",
											"GITHUB",
											"GITLAB",
										}, true),
									},

									// Optional
									"branch": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"connection_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"repository_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"repository_url": {
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
			"build_spec_file": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"deliver_artifact_collection": {
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"artifact_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"artifact_name": {
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
			"deploy_pipeline_id": {
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
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"image": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_pass_all_parameters_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"primary_build_source": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stage_execution_timeout_in_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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

func createDevopsBuildPipelineStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsBuildPipelineStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsBuildPipelineStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsBuildPipelineStage(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineStageResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsBuildPipelineStageResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.BuildPipelineStage
	DisableNotFoundRetries bool
}

func (s *DevopsBuildPipelineStageResourceCrud) ID() string {
	buildPipelineStage := *s.Res
	return *buildPipelineStage.GetId()
}

func (s *DevopsBuildPipelineStageResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.BuildPipelineStageLifecycleStateCreating),
	}
}

func (s *DevopsBuildPipelineStageResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.BuildPipelineStageLifecycleStateActive),
	}
}

func (s *DevopsBuildPipelineStageResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_devops.BuildPipelineStageLifecycleStateDeleting),
	}
}

func (s *DevopsBuildPipelineStageResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.BuildPipelineStageLifecycleStateDeleted),
	}
}

func (s *DevopsBuildPipelineStageResourceCrud) Create() error {
	request := oci_devops.CreateBuildPipelineStageRequest{}
	err := s.populateTopLevelPolymorphicCreateBuildPipelineStageRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateBuildPipelineStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBuildPipelineStageFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsBuildPipelineStageResourceCrud) getBuildPipelineStageFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	buildPipelineStageId, err := buildPipelineStageWaitForWorkRequest(workId, "buildPipelineStage",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*buildPipelineStageId)

	return s.Get()
}

func buildPipelineStageWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func buildPipelineStageWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = buildPipelineStageWorkRequestShouldRetryFunc(timeout)

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
		if strings.Contains(strings.ToLower(*res.EntityType), strings.ToLower(entityType)) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_devops.OperationStatusFailed || response.Status == oci_devops.OperationStatusCanceled {
		return nil, getErrorFromDevopsBuildPipelineStageWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsBuildPipelineStageWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsBuildPipelineStageResourceCrud) Get() error {
	request := oci_devops.GetBuildPipelineStageRequest{}

	tmp := s.D.Id()
	request.BuildPipelineStageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetBuildPipelineStage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BuildPipelineStage
	return nil
}

func (s *DevopsBuildPipelineStageResourceCrud) Update() error {
	request := oci_devops.UpdateBuildPipelineStageRequest{}
	err := s.populateTopLevelPolymorphicUpdateBuildPipelineStageRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateBuildPipelineStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBuildPipelineStageFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsBuildPipelineStageResourceCrud) Delete() error {
	request := oci_devops.DeleteBuildPipelineStageRequest{}

	tmp := s.D.Id()
	request.BuildPipelineStageId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteBuildPipelineStage(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := buildPipelineStageWaitForWorkRequest(workId, "buildPipelineStage",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsBuildPipelineStageResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_devops.BuildStage:
		s.D.Set("build_pipeline_stage_type", "BUILD")

		if v.BuildSourceCollection != nil {
			s.D.Set("build_source_collection", []interface{}{BuildSourceCollectionToMap(v.BuildSourceCollection)})
		} else {
			s.D.Set("build_source_collection", nil)
		}

		if v.BuildSpecFile != nil {
			s.D.Set("build_spec_file", *v.BuildSpecFile)
		}

		s.D.Set("image", v.Image)

		if v.PrimaryBuildSource != nil {
			s.D.Set("primary_build_source", *v.PrimaryBuildSource)
		}

		if v.StageExecutionTimeoutInSeconds != nil {
			s.D.Set("stage_execution_timeout_in_seconds", *v.StageExecutionTimeoutInSeconds)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

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
	case oci_devops.DeliverArtifactStage:
		s.D.Set("build_pipeline_stage_type", "DELIVER_ARTIFACT")

		if v.DeliverArtifactCollection != nil {
			s.D.Set("deliver_artifact_collection", []interface{}{DeliverArtifactCollectionToMap(v.DeliverArtifactCollection)})
		} else {
			s.D.Set("deliver_artifact_collection", nil)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

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
	case oci_devops.TriggerDeploymentStage:
		s.D.Set("build_pipeline_stage_type", "TRIGGER_DEPLOYMENT_PIPELINE")

		if v.DeployPipelineId != nil {
			s.D.Set("deploy_pipeline_id", *v.DeployPipelineId)
		}

		if v.IsPassAllParametersEnabled != nil {
			s.D.Set("is_pass_all_parameters_enabled", *v.IsPassAllParametersEnabled)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

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
	case oci_devops.WaitStage:
		s.D.Set("build_pipeline_stage_type", "WAIT")

		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			s.D.Set("wait_criteria", waitCriteriaArray)
		} else {
			s.D.Set("wait_criteria", nil)
		}

		if v.BuildPipelineId != nil {
			s.D.Set("build_pipeline_id", *v.BuildPipelineId)
		}

		if v.BuildPipelineStagePredecessorCollection != nil {
			s.D.Set("build_pipeline_stage_predecessor_collection", []interface{}{BuildPipelineStagePredecessorCollectionToMap(v.BuildPipelineStagePredecessorCollection)})
		} else {
			s.D.Set("build_pipeline_stage_predecessor_collection", nil)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.Set("id", *v.Id)
		}

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
	default:
		log.Printf("[WARN] Received 'build_pipeline_stage_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func BuildPipelineStagePredecessorToMap(obj oci_devops.BuildPipelineStagePredecessor) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func BuildPipelineStagePredecessorCollectionToMap(obj *oci_devops.BuildPipelineStagePredecessorCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, BuildPipelineStagePredecessorToMap(item))
	}
	result["items"] = items

	return result
}

func BuildPipelineStageSummaryToMap(obj oci_devops.BuildPipelineStageSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetId() != nil {
		result["id"] = string(*obj.GetId())
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = string(*obj.GetCompartmentId())
	}

	if obj.GetProjectId() != nil {
		result["project_id"] = string(*obj.GetProjectId())
	}

	if obj.GetProjectId() != nil {
		result["build_pipeline_id"] = string(*obj.GetBuildPipelineId())
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = string(*obj.GetDisplayName())
	}

	if obj.GetDescription() != nil {
		result["description"] = string(*obj.GetDescription())
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	result["state"] = obj.GetLifecycleState()

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = string(*obj.GetLifecycleDetails())
	}

	if obj.GetBuildPipelineStagePredecessorCollection() != nil {
		result["build_pipeline_stage_predecessor_collection"] = []interface{}{BuildPipelineStagePredecessorCollectionToMap(obj.GetBuildPipelineStagePredecessorCollection())}
	}

	if obj.GetFreeformTags() != nil {
		result["freeform_tags"] = obj.GetFreeformTags()
	}

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	switch v := (obj).(type) {
	case oci_devops.BuildStageSummary:
		result["build_pipeline_stage_type"] = "BUILD"

		if v.BuildSourceCollection != nil {
			result["build_source_collection"] = []interface{}{BuildSourceCollectionToMap(v.BuildSourceCollection)}
		}

		if v.BuildSpecFile != nil {
			result["build_spec_file"] = string(*v.BuildSpecFile)
		}

		result["image"] = string(v.Image)

		if v.PrimaryBuildSource != nil {
			result["primary_build_source"] = string(*v.PrimaryBuildSource)
		}

		if v.StageExecutionTimeoutInSeconds != nil {
			result["stage_execution_timeout_in_seconds"] = int(*v.StageExecutionTimeoutInSeconds)
		}
	case oci_devops.DeliverArtifactStageSummary:
		result["build_pipeline_stage_type"] = "DELIVER_ARTIFACT"

		if v.DeliverArtifactCollection != nil {
			result["deliver_artifact_collection"] = []interface{}{DeliverArtifactCollectionToMap(v.DeliverArtifactCollection)}
		}
	case oci_devops.TriggerDeploymentStageSummary:
		result["build_pipeline_stage_type"] = "TRIGGER_DEPLOYMENT_PIPELINE"

		if v.DeployPipelineId != nil {
			result["deploy_pipeline_id"] = string(*v.DeployPipelineId)
		}

		if v.IsPassAllParametersEnabled != nil {
			result["is_pass_all_parameters_enabled"] = bool(*v.IsPassAllParametersEnabled)
		}
	case oci_devops.WaitStageSummary:
		result["build_pipeline_stage_type"] = "WAIT"

		if v.WaitCriteria != nil {
			waitCriteriaArray := []interface{}{}
			if waitCriteriaMap := WaitCriteriaToMap(&v.WaitCriteria); waitCriteriaMap != nil {
				waitCriteriaArray = append(waitCriteriaArray, waitCriteriaMap)
			}
			result["wait_criteria"] = waitCriteriaArray
		}
	default:
		log.Printf("[WARN] Received 'build_pipeline_stage_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToBuildSource(fieldKeyFormat string) (oci_devops.BuildSource, error) {
	var baseObject oci_devops.BuildSource
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_type"))
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("DEVOPS_CODE_REPOSITORY"):
		details := oci_devops.DevopsCodeRepositoryBuildSource{}
		if repositoryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_id")); ok {
			tmp := repositoryId.(string)
			details.RepositoryId = &tmp
		}
		if branch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("GITHUB"):
		details := oci_devops.GithubBuildSource{}
		if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if branch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("GITLAB"):
		details := oci_devops.GitlabBuildSource{}
		if connectionId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_id")); ok {
			tmp := connectionId.(string)
			details.ConnectionId = &tmp
		}
		if branch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
			tmp := repositoryUrl.(string)
			details.RepositoryUrl = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return baseObject, nil
}

func BuildSourceToMap(obj oci_devops.BuildSource) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.GetBranch() != nil {
		result["branch"] = string(*obj.GetBranch())
	}

	if obj.GetName() != nil {
		result["name"] = string(*obj.GetName())
	}

	if obj.GetRepositoryUrl() != nil {
		result["repository_url"] = string(*obj.GetRepositoryUrl())
	}

	switch v := (obj).(type) {
	case oci_devops.DevopsCodeRepositoryBuildSource:
		result["connection_type"] = "DEVOPS_CODE_REPOSITORY"

		if v.RepositoryId != nil {
			result["repository_id"] = string(*v.RepositoryId)
		}
	case oci_devops.GithubBuildSource:
		result["connection_type"] = "GITHUB"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	case oci_devops.GitlabBuildSource:
		result["connection_type"] = "GITLAB"

		if v.ConnectionId != nil {
			result["connection_id"] = string(*v.ConnectionId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToBuildSourceCollection(fieldKeyFormat string) (oci_devops.BuildSourceCollection, error) {
	result := oci_devops.BuildSourceCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.BuildSource, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToBuildSource(fieldKeyFormatNextLevel)
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

func BuildSourceCollectionToMap(obj *oci_devops.BuildSourceCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, BuildSourceToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToCreateWaitCriteriaDetails(fieldKeyFormat string) (oci_devops.CreateWaitCriteriaDetails, error) {
	var baseObject oci_devops.CreateWaitCriteriaDetails
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
		details := oci_devops.CreateAbsoluteWaitCriteriaDetails{}
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

func CreateWaitCriteriaDetailsToMap(obj *oci_devops.CreateWaitCriteriaDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.CreateAbsoluteWaitCriteriaDetails:
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

func (s *DevopsBuildPipelineStageResourceCrud) mapToDeliverArtifact(fieldKeyFormat string) (oci_devops.DeliverArtifact, error) {
	result := oci_devops.DeliverArtifact{}

	if artifactId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_id")); ok {
		tmp := artifactId.(string)
		result.ArtifactId = &tmp
	}

	if artifactName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_name")); ok {
		tmp := artifactName.(string)
		result.ArtifactName = &tmp
	}

	return result, nil
}

func DeliverArtifactToMap(obj oci_devops.DeliverArtifact) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArtifactId != nil {
		result["artifact_id"] = string(*obj.ArtifactId)
	}

	if obj.ArtifactName != nil {
		result["artifact_name"] = string(*obj.ArtifactName)
	}

	return result
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToDeliverArtifactCollection(fieldKeyFormat string) (oci_devops.DeliverArtifactCollection, error) {
	result := oci_devops.DeliverArtifactCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeliverArtifact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeliverArtifact(fieldKeyFormatNextLevel)
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

func DeliverArtifactCollectionToMap(obj *oci_devops.DeliverArtifactCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeliverArtifactToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToUpdateWaitCriteriaDetails(fieldKeyFormat string) (oci_devops.UpdateWaitCriteriaDetails, error) {
	var baseObject oci_devops.UpdateWaitCriteriaDetails
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
		details := oci_devops.UpdateAbsoluteWaitCriteriaDetails{}
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

func UpdateWaitCriteriaDetailsToMap(obj *oci_devops.UpdateWaitCriteriaDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.UpdateAbsoluteWaitCriteriaDetails:
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

func (s *DevopsBuildPipelineStageResourceCrud) mapToWaitCriteria(fieldKeyFormat string) (oci_devops.WaitCriteria, error) {
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

func WaitCriteriaToMap(obj *oci_devops.WaitCriteria) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.AbsoluteWaitCriteria:
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

func (s *DevopsBuildPipelineStageResourceCrud) populateTopLevelPolymorphicCreateBuildPipelineStageRequest(request *oci_devops.CreateBuildPipelineStageRequest) error {
	//discriminator
	buildPipelineStageTypeRaw, ok := s.D.GetOkExists("build_pipeline_stage_type")
	var buildPipelineStageType string
	if ok {
		buildPipelineStageType = buildPipelineStageTypeRaw.(string)
	} else {
		buildPipelineStageType = "" // default value
	}
	switch strings.ToLower(buildPipelineStageType) {
	case strings.ToLower("BUILD"):
		details := oci_devops.CreateBuildStageDetails{}
		if buildSourceCollection, ok := s.D.GetOkExists("build_source_collection"); ok {
			if tmpList := buildSourceCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_source_collection", 0)
				tmp, err := s.mapToBuildSourceCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildSourceCollection = &tmp
			}
		}
		if buildSpecFile, ok := s.D.GetOkExists("build_spec_file"); ok {
			tmp := buildSpecFile.(string)
			details.BuildSpecFile = &tmp
		}
		if image, ok := s.D.GetOkExists("image"); ok {
			details.Image = oci_devops.BuildStageImageEnum(image.(string))
		}
		if primaryBuildSource, ok := s.D.GetOkExists("primary_build_source"); ok {
			tmp := primaryBuildSource.(string)
			details.PrimaryBuildSource = &tmp
		}
		if stageExecutionTimeoutInSeconds, ok := s.D.GetOkExists("stage_execution_timeout_in_seconds"); ok {
			tmp := stageExecutionTimeoutInSeconds.(int)
			details.StageExecutionTimeoutInSeconds = &tmp
		}
		if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBuildPipelineStageDetails = details
	case strings.ToLower("DELIVER_ARTIFACT"):
		details := oci_devops.CreateDeliverArtifactStageDetails{}
		if deliverArtifactCollection, ok := s.D.GetOkExists("deliver_artifact_collection"); ok {
			if tmpList := deliverArtifactCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deliver_artifact_collection", 0)
				tmp, err := s.mapToDeliverArtifactCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeliverArtifactCollection = &tmp
			}
		}
		if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBuildPipelineStageDetails = details
	case strings.ToLower("TRIGGER_DEPLOYMENT_PIPELINE"):
		details := oci_devops.CreateTriggerDeploymentStageDetails{}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if isPassAllParametersEnabled, ok := s.D.GetOkExists("is_pass_all_parameters_enabled"); ok {
			tmp := isPassAllParametersEnabled.(bool)
			details.IsPassAllParametersEnabled = &tmp
		}
		if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBuildPipelineStageDetails = details
	case strings.ToLower("WAIT"):
		details := oci_devops.CreateWaitStageDetails{}
		if waitCriteria, ok := s.D.GetOkExists("wait_criteria"); ok {
			if tmpList := waitCriteria.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "wait_criteria", 0)
				tmp, err := s.mapToCreateWaitCriteriaDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.WaitCriteria = tmp
			}
		}
		if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.CreateBuildPipelineStageDetails = details
	default:
		return fmt.Errorf("unknown build_pipeline_stage_type '%v' was specified", buildPipelineStageType)
	}
	return nil
}

func (s *DevopsBuildPipelineStageResourceCrud) populateTopLevelPolymorphicUpdateBuildPipelineStageRequest(request *oci_devops.UpdateBuildPipelineStageRequest) error {
	//discriminator
	buildPipelineStageTypeRaw, ok := s.D.GetOkExists("build_pipeline_stage_type")
	var buildPipelineStageType string
	if ok {
		buildPipelineStageType = buildPipelineStageTypeRaw.(string)
	} else {
		buildPipelineStageType = "" // default value
	}
	switch strings.ToLower(buildPipelineStageType) {
	case strings.ToLower("BUILD"):
		details := oci_devops.UpdateBuildStageDetails{}
		if buildSourceCollection, ok := s.D.GetOkExists("build_source_collection"); ok {
			if tmpList := buildSourceCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_source_collection", 0)
				tmp, err := s.mapToBuildSourceCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildSourceCollection = &tmp
			}
		}
		if buildSpecFile, ok := s.D.GetOkExists("build_spec_file"); ok {
			tmp := buildSpecFile.(string)
			details.BuildSpecFile = &tmp
		}
		if image, ok := s.D.GetOkExists("image"); ok {
			details.Image = oci_devops.BuildStageImageEnum(image.(string))
		}
		if primaryBuildSource, ok := s.D.GetOkExists("primary_build_source"); ok {
			tmp := primaryBuildSource.(string)
			details.PrimaryBuildSource = &tmp
		}
		if stageExecutionTimeoutInSeconds, ok := s.D.GetOkExists("stage_execution_timeout_in_seconds"); ok {
			tmp := stageExecutionTimeoutInSeconds.(int)
			details.StageExecutionTimeoutInSeconds = &tmp
		}
		tmp := s.D.Id()
		request.BuildPipelineStageId = &tmp
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateBuildPipelineStageDetails = details
	case strings.ToLower("DELIVER_ARTIFACT"):
		details := oci_devops.UpdateDeliverArtifactStageDetails{}
		if deliverArtifactCollection, ok := s.D.GetOkExists("deliver_artifact_collection"); ok {
			if tmpList := deliverArtifactCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "deliver_artifact_collection", 0)
				tmp, err := s.mapToDeliverArtifactCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.DeliverArtifactCollection = &tmp
			}
		}
		tmp := s.D.Id()
		request.BuildPipelineStageId = &tmp
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateBuildPipelineStageDetails = details
	case strings.ToLower("TRIGGER_DEPLOYMENT_PIPELINE"):
		details := oci_devops.UpdateTriggerDeploymentStageDetails{}
		if deployPipelineId, ok := s.D.GetOkExists("deploy_pipeline_id"); ok {
			tmp := deployPipelineId.(string)
			details.DeployPipelineId = &tmp
		}
		if isPassAllParametersEnabled, ok := s.D.GetOkExists("is_pass_all_parameters_enabled"); ok {
			tmp := isPassAllParametersEnabled.(bool)
			details.IsPassAllParametersEnabled = &tmp
		}
		tmp := s.D.Id()
		request.BuildPipelineStageId = &tmp
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateBuildPipelineStageDetails = details
	case strings.ToLower("WAIT"):
		details := oci_devops.UpdateWaitStageDetails{}
		if waitCriteria, ok := s.D.GetOkExists("wait_criteria"); ok {
			if tmpList := waitCriteria.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "wait_criteria", 0)
				tmp, err := s.mapToUpdateWaitCriteriaDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.WaitCriteria = tmp
			}
		}
		tmp := s.D.Id()
		request.BuildPipelineStageId = &tmp
		if buildPipelineStagePredecessorCollection, ok := s.D.GetOkExists("build_pipeline_stage_predecessor_collection"); ok {
			if tmpList := buildPipelineStagePredecessorCollection.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_stage_predecessor_collection", 0)
				tmp, err := s.mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.BuildPipelineStagePredecessorCollection = &tmp
			}
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
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
			details.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		request.UpdateBuildPipelineStageDetails = details
	default:
		return fmt.Errorf("unknown build_pipeline_stage_type '%v' was specified", buildPipelineStageType)
	}
	return nil
}

func (s *DevopsBuildPipelineStageResourceCrud) mapToBuildPipelineStagePredecessorCollection(fieldKeyFormat string) (oci_devops.BuildPipelineStagePredecessorCollection, error) {
	result := oci_devops.BuildPipelineStagePredecessorCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.BuildPipelineStagePredecessor, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToBuildPipelineStagePredecessor(fieldKeyFormatNextLevel)
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

func (s *DevopsBuildPipelineStageResourceCrud) mapToBuildPipelineStagePredecessor(fieldKeyFormat string) (oci_devops.BuildPipelineStagePredecessor, error) {
	result := oci_devops.BuildPipelineStagePredecessor{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	return result, nil
}
