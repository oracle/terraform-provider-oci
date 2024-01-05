// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsBuildRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsBuildRun,
		Read:     readDevopsBuildRun,
		Update:   updateDevopsBuildRun,
		Delete:   deleteDevopsBuildRun,
		Schema: map[string]*schema.Schema{
			// Required
			"build_pipeline_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"build_run_arguments": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"items": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: true,
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
			"commit_info": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"commit_hash": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"repository_branch": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional

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

			// Computed
			"build_outputs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"artifact_override_parameters": {
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
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"delivered_artifacts": {
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
												"artifact_repository_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"artifact_type": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"delivered_artifact_hash": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"delivered_artifact_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"deploy_artifact_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"image_uri": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"output_artifact_name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"path": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"exported_variables": {
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
												"name": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"value": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"vulnerability_audit_summary_collection": {
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
												"build_stage_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"commit_hash": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"vulnerability_audit_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"build_run_progress": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"build_pipeline_stage_run_progress": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
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
			"build_run_source": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"repository_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trigger_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"trigger_info": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"actions": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// Required

												// Optional

												// Computed
												"build_pipeline_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"filter": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															// Required

															// Optional

															// Computed
															"events": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"exclude": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"file_filter": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"file_paths": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															"include": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		// Required

																		// Optional

																		// Computed
																		"base_ref": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"file_filter": {
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					// Required

																					// Optional

																					// Computed
																					"file_paths": {
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Schema{
																							Type: schema.TypeString,
																						},
																					},
																				},
																			},
																		},
																		"head_ref": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"repository_name": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
															"trigger_source": {
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"type": {
													Type:     schema.TypeString,
													Computed: true,
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

func createDevopsBuildRun(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsBuildRun(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsBuildRun(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsBuildRun(d *schema.ResourceData, m interface{}) error {
	return nil
}

type DevopsBuildRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.BuildRun
	DisableNotFoundRetries bool
}

func (s *DevopsBuildRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DevopsBuildRunResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.BuildRunLifecycleStateInProgress),
		string(oci_devops.BuildRunLifecycleStateAccepted),
	}
}

func (s *DevopsBuildRunResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.BuildRunLifecycleStateSucceeded),
	}
}

func (s *DevopsBuildRunResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DevopsBuildRunResourceCrud) DeletedTarget() []string {
	return []string{}
}

func (s *DevopsBuildRunResourceCrud) Create() error {
	request := oci_devops.CreateBuildRunRequest{}

	if buildPipelineId, ok := s.D.GetOkExists("build_pipeline_id"); ok {
		tmp := buildPipelineId.(string)
		request.BuildPipelineId = &tmp
	}

	if buildRunArguments, ok := s.D.GetOkExists("build_run_arguments"); ok {
		if tmpList := buildRunArguments.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_run_arguments", 0)
			tmp, err := s.mapToBuildRunArgumentCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BuildRunArguments = &tmp
		}
	}

	if commitInfo, ok := s.D.GetOkExists("commit_info"); ok {
		if tmpList := commitInfo.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "commit_info", 0)
			tmp, err := s.mapToCommitInfo(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.CommitInfo = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateBuildRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BuildRun
	return nil
}

func (s *DevopsBuildRunResourceCrud) Get() error {
	request := oci_devops.GetBuildRunRequest{}

	tmp := s.D.Id()
	request.BuildRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetBuildRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BuildRun
	return nil
}

func (s *DevopsBuildRunResourceCrud) Update() error {
	request := oci_devops.UpdateBuildRunRequest{}

	tmp := s.D.Id()
	request.BuildRunId = &tmp

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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateBuildRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BuildRun
	return nil
}

func (s *DevopsBuildRunResourceCrud) SetData() error {
	if s.Res.BuildOutputs != nil {
		s.D.Set("build_outputs", []interface{}{BuildOutputsToMap(s.Res.BuildOutputs)})
	} else {
		s.D.Set("build_outputs", nil)
	}

	if s.Res.BuildPipelineId != nil {
		s.D.Set("build_pipeline_id", *s.Res.BuildPipelineId)
	}

	if s.Res.BuildRunArguments != nil {
		s.D.Set("build_run_arguments", []interface{}{BuildRunArgumentCollectionToMap(s.Res.BuildRunArguments)})
	} else {
		s.D.Set("build_run_arguments", nil)
	}

	if s.Res.BuildRunProgress != nil {
		s.D.Set("build_run_progress", []interface{}{BuildRunProgressToMap(s.Res.BuildRunProgress)})
	} else {
		s.D.Set("build_run_progress", nil)
	}

	if s.Res.BuildRunSource != nil {
		buildRunSourceArray := []interface{}{}
		if buildRunSourceMap := BuildRunSourceToMap(&s.Res.BuildRunSource); buildRunSourceMap != nil {
			buildRunSourceArray = append(buildRunSourceArray, buildRunSourceMap)
		}
		s.D.Set("build_run_source", buildRunSourceArray)
	} else {
		s.D.Set("build_run_source", nil)
	}

	if s.Res.CommitInfo != nil {
		s.D.Set("commit_info", []interface{}{CommitInfoToMap(s.Res.CommitInfo)})
	} else {
		s.D.Set("commit_info", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
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

	return nil
}

func (s *DevopsBuildRunResourceCrud) mapToActualBuildRunnerShapeConfig(fieldKeyFormat string) (oci_devops.ActualBuildRunnerShapeConfig, error) {
	result := oci_devops.ActualBuildRunnerShapeConfig{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(float64)
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(float64)
		result.Ocpus = &tmp
	}

	return result, nil
}

func ActualBuildRunnerShapeConfigToMap(obj *oci_devops.ActualBuildRunnerShapeConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float64(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float64(*obj.Ocpus)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToBitbucketCloudFilterAttributes(fieldKeyFormat string) (oci_devops.BitbucketCloudFilterAttributes, error) {
	result := oci_devops.BitbucketCloudFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func BitbucketCloudFilterAttributesToMap(obj *oci_devops.BitbucketCloudFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToBitbucketServerFilterAttributes(fieldKeyFormat string) (oci_devops.BitbucketServerFilterAttributes, error) {
	result := oci_devops.BitbucketServerFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func BitbucketServerFilterAttributesToMap(obj *oci_devops.BitbucketServerFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToBitbucketCloudFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.BitbucketCloudFilterExclusionAttributes, error) {
	result := oci_devops.BitbucketCloudFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func BitbucketCloudFilterExclusionAttributesToMap(obj *oci_devops.BitbucketCloudFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func BuildOutputsToMap(obj *oci_devops.BuildOutputs) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ArtifactOverrideParameters != nil {
		result["artifact_override_parameters"] = []interface{}{DeployArtifactOverrideArgumentCollectionToMap(obj.ArtifactOverrideParameters)}
	}

	if obj.DeliveredArtifacts != nil {
		result["delivered_artifacts"] = []interface{}{DeliveredArtifactCollectionToMap(obj.DeliveredArtifacts)}
	}

	if obj.ExportedVariables != nil {
		result["exported_variables"] = []interface{}{ExportedVariableCollectionToMap(obj.ExportedVariables)}
	}

	if obj.VulnerabilityAuditSummaryCollection != nil {
		result["vulnerability_audit_summary_collection"] = []interface{}{VulnerabilityAuditSummaryCollectionToMap(obj.VulnerabilityAuditSummaryCollection)}
	}

	return result
}

func BuildRunArgumentToMap(obj oci_devops.BuildRunArgument) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func BuildRunArgumentCollectionToMap(obj *oci_devops.BuildRunArgumentCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, BuildRunArgumentToMap(item))
	}
	result["items"] = items

	return result
}

func BuildRunProgressToMap(obj *oci_devops.BuildRunProgress) map[string]interface{} {
	result := map[string]interface{}{}

	result["build_pipeline_stage_run_progress"] = obj.BuildPipelineStageRunProgress

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func BuildRunProgressSummaryToMap(obj *oci_devops.BuildRunProgressSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	return result
}

func BuildRunSourceToMap(obj *oci_devops.BuildRunSource) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.BitbucketCloudBuildRunSource:
		result["source_type"] = "BITBUCKET_CLOUD"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.BitbucketServerBuildRunSource:
		result["source_type"] = "BITBUCKET_SERVER"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.DevopsCodeRepositoryBuildRunSource:
		result["source_type"] = "DEVOPS_CODE_REPOSITORY"

		if v.RepositoryId != nil {
			result["repository_id"] = string(*v.RepositoryId)
		}

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.GithubBuildRunSource:
		result["source_type"] = "GITHUB"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.GitlabBuildRunSource:
		result["source_type"] = "GITLAB"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.GitlabServerBuildRunSource:
		result["source_type"] = "GITLAB_SERVER"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	case oci_devops.ManualBuildRunSource:
		result["source_type"] = "MANUAL"
	case oci_devops.VbsBuildRunSource:
		result["source_type"] = "VBS"

		if v.TriggerId != nil {
			result["trigger_id"] = string(*v.TriggerId)
		}

		if v.TriggerInfo != nil {
			result["trigger_info"] = []interface{}{TriggerInfoToMap(v.TriggerInfo)}
		}
	default:
		log.Printf("[WARN] Received 'source_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func BuildRunSummaryToMap(obj oci_devops.BuildRunSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BuildPipelineId != nil {
		result["build_pipeline_id"] = string(*obj.BuildPipelineId)
	}

	if obj.BuildRunArguments != nil {
		result["build_run_arguments"] = []interface{}{BuildRunArgumentCollectionToMap(obj.BuildRunArguments)}
	}

	if obj.BuildRunProgressSummary != nil {
		result["build_run_progress_summary"] = []interface{}{BuildRunProgressSummaryToMap(obj.BuildRunProgressSummary)}
	}

	if obj.BuildRunSource != nil {
		buildRunSourceArray := []interface{}{}
		if buildRunSourceMap := BuildRunSourceToMap(&obj.BuildRunSource); buildRunSourceMap != nil {
			buildRunSourceArray = append(buildRunSourceArray, buildRunSourceMap)
		}
		result["build_run_source"] = buildRunSourceArray
	}

	if obj.CommitInfo != nil {
		result["commit_info"] = []interface{}{CommitInfoToMap(obj.CommitInfo)}
	}

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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
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

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToBuildSource(fieldKeyFormat string) (oci_devops.BuildSource, error) {
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
	case strings.ToLower("BITBUCKET_CLOUD"):
		details := oci_devops.BitbucketCloudBuildSource{}
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
	case strings.ToLower("BITBUCKET_SERVER"):
		details := oci_devops.BitbucketServerBuildSource{}
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
	case strings.ToLower("GITLAB_SERVER"):
		details := oci_devops.GitlabServerBuildSource{}
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
	case strings.ToLower("VBS"):
		details := oci_devops.VbsBuildSource{}
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

func (s *DevopsBuildRunResourceCrud) mapToBuildSourceCollection(fieldKeyFormat string) (oci_devops.BuildSourceCollection, error) {
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

func (s *DevopsBuildRunResourceCrud) mapToBuildStageRunStep(fieldKeyFormat string) (oci_devops.BuildStageRunStep, error) {
	result := oci_devops.BuildStageRunStep{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if state, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "state")); ok {
		result.State = oci_devops.BuildStageRunStepStateEnum(state.(string))
	}

	if timeFinished, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_finished")); ok {
		tmp, err := time.Parse(time.RFC3339, timeFinished.(string))
		if err != nil {
			return result, err
		}
		result.TimeFinished = &oci_common.SDKTime{Time: tmp}
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

func BuildStageRunStepToMap(obj oci_devops.BuildStageRunStep) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["state"] = string(obj.State)

	if obj.TimeFinished != nil {
		result["time_finished"] = obj.TimeFinished.Format(time.RFC3339Nano)
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.Format(time.RFC3339Nano)
	}

	return result
}

func CommitInfoToMap(obj *oci_devops.CommitInfo) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CommitHash != nil {
		result["commit_hash"] = string(*obj.CommitHash)
	}

	if obj.RepositoryBranch != nil {
		result["repository_branch"] = string(*obj.RepositoryBranch)
	}

	if obj.RepositoryUrl != nil {
		result["repository_url"] = string(*obj.RepositoryUrl)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDeliveredArtifact(fieldKeyFormat string) (oci_devops.DeliveredArtifact, error) {
	var baseObject oci_devops.DeliveredArtifact
	//discriminator
	artifactTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_type"))
	var artifactType string
	if ok {
		artifactType = artifactTypeRaw.(string)
	} else {
		artifactType = "" // default value
	}
	switch strings.ToLower(artifactType) {
	case strings.ToLower("GENERIC_ARTIFACT"):
		details := oci_devops.GenericDeliveredArtifact{}
		if artifactRepositoryId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_repository_id")); ok {
			tmp := artifactRepositoryId.(string)
			details.ArtifactRepositoryId = &tmp
		}
		if deliveredArtifactId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delivered_artifact_id")); ok {
			tmp := deliveredArtifactId.(string)
			details.DeliveredArtifactId = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
			tmp := version.(string)
			details.Version = &tmp
		}
		if deployArtifactId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_id")); ok {
			tmp := deployArtifactId.(string)
			details.DeployArtifactId = &tmp
		}
		if outputArtifactName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_artifact_name")); ok {
			tmp := outputArtifactName.(string)
			details.OutputArtifactName = &tmp
		}
		baseObject = details
	case strings.ToLower("OCIR"):
		details := oci_devops.ContainerRegistryDeliveredArtifact{}
		if deliveredArtifactHash, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "delivered_artifact_hash")); ok {
			tmp := deliveredArtifactHash.(string)
			details.DeliveredArtifactHash = &tmp
		}
		if imageUri, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "image_uri")); ok {
			tmp := imageUri.(string)
			details.ImageUri = &tmp
		}
		if deployArtifactId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deploy_artifact_id")); ok {
			tmp := deployArtifactId.(string)
			details.DeployArtifactId = &tmp
		}
		if outputArtifactName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "output_artifact_name")); ok {
			tmp := outputArtifactName.(string)
			details.OutputArtifactName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown artifact_type '%v' was specified", artifactType)
	}
	return baseObject, nil
}

func DeliveredArtifactToMap(obj oci_devops.DeliveredArtifact) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_devops.GenericDeliveredArtifact:
		result["artifact_type"] = "GENERIC_ARTIFACT"

		if v.ArtifactRepositoryId != nil {
			result["artifact_repository_id"] = string(*v.ArtifactRepositoryId)
		}

		if v.DeliveredArtifactId != nil {
			result["delivered_artifact_id"] = string(*v.DeliveredArtifactId)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Version != nil {
			result["version"] = string(*v.Version)
		}
	case oci_devops.ContainerRegistryDeliveredArtifact:
		result["artifact_type"] = "OCIR"

		if v.DeliveredArtifactHash != nil {
			result["delivered_artifact_hash"] = string(*v.DeliveredArtifactHash)
		}

		if v.ImageUri != nil {
			result["image_uri"] = string(*v.ImageUri)
		}
	default:
		log.Printf("[WARN] Received 'artifact_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDeliveredArtifactCollection(fieldKeyFormat string) (oci_devops.DeliveredArtifactCollection, error) {
	result := oci_devops.DeliveredArtifactCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.DeliveredArtifact, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToDeliveredArtifact(fieldKeyFormatNextLevel)
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

func DeliveredArtifactCollectionToMap(obj *oci_devops.DeliveredArtifactCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeliveredArtifactToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDeployArtifactOverrideArgument(fieldKeyFormat string) (oci_devops.DeployArtifactOverrideArgument, error) {
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

func DeployArtifactOverrideArgumentToMap(obj oci_devops.DeployArtifactOverrideArgument) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DeployArtifactId != nil {
		result["deploy_artifact_id"] = string(*obj.DeployArtifactId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDeployArtifactOverrideArgumentCollection(fieldKeyFormat string) (oci_devops.DeployArtifactOverrideArgumentCollection, error) {
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

func DeployArtifactOverrideArgumentCollectionToMap(obj *oci_devops.DeployArtifactOverrideArgumentCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, DeployArtifactOverrideArgumentToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDevopsCodeRepositoryFilterAttributes(fieldKeyFormat string) (oci_devops.DevopsCodeRepositoryFilterAttributes, error) {
	result := oci_devops.DevopsCodeRepositoryFilterAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func DevopsCodeRepositoryFilterAttributesToMap(obj *oci_devops.DevopsCodeRepositoryFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToDevopsCodeRepositoryFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.DevopsCodeRepositoryFilterExclusionAttributes, error) {
	result := oci_devops.DevopsCodeRepositoryFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func DevopsCodeRepositoryFilterExclusionAttributesToMap(obj *oci_devops.DevopsCodeRepositoryFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToExportedVariable(fieldKeyFormat string) (oci_devops.ExportedVariable, error) {
	result := oci_devops.ExportedVariable{}

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

func ExportedVariableToMap(obj oci_devops.ExportedVariable) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToExportedVariableCollection(fieldKeyFormat string) (oci_devops.ExportedVariableCollection, error) {
	result := oci_devops.ExportedVariableCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.ExportedVariable, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToExportedVariable(fieldKeyFormatNextLevel)
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

func ExportedVariableCollectionToMap(obj *oci_devops.ExportedVariableCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, ExportedVariableToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToFileFilter(fieldKeyFormat string) (oci_devops.FileFilter, error) {
	result := oci_devops.FileFilter{}

	if filePaths, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_paths")); ok {
		interfaces := filePaths.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "file_paths")) {
			result.FilePaths = tmp
		}
	}

	return result, nil
}

func FileFilterToMap(obj *oci_devops.FileFilter) map[string]interface{} {
	result := map[string]interface{}{}

	result["file_paths"] = obj.FilePaths

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToFilter(fieldKeyFormat string) (oci_devops.Filter, error) {
	var baseObject oci_devops.Filter
	//discriminator
	triggerSourceRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_source"))
	var triggerSource string
	if ok {
		triggerSource = triggerSourceRaw.(string)
	} else {
		triggerSource = "" // default value
	}
	switch strings.ToLower(triggerSource) {
	case strings.ToLower("BITBUCKET_CLOUD"):
		details := oci_devops.BitbucketCloudFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.BitbucketCloudFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.BitbucketCloudFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToBitbucketCloudFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToBitbucketCloudFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("BITBUCKET_SERVER"):
		details := oci_devops.BitbucketServerFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.BitbucketServerFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.BitbucketServerFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToBitbucketServerFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("DEVOPS_CODE_REPOSITORY"):
		details := oci_devops.DevopsCodeRepositoryFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.DevopsCodeRepositoryFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(oci_devops.DevopsCodeRepositoryFilterEventsEnum)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToDevopsCodeRepositoryFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToDevopsCodeRepositoryFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITHUB"):
		details := oci_devops.GithubFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GithubFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(oci_devops.GithubFilterEventsEnum)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToGithubFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGithubFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITLAB"):
		details := oci_devops.GitlabFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GitlabFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(oci_devops.GitlabFilterEventsEnum)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToGitlabFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGitlabFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("GITLAB_SERVER"):
		details := oci_devops.GitlabServerFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.GitlabServerFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.GitlabServerFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToGitlabServerFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("VBS"):
		details := oci_devops.VbsFilter{}
		if events, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "events")); ok {
			interfaces := events.([]interface{})
			tmp := make([]oci_devops.VbsFilterEventsEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = oci_devops.VbsFilterEventsEnum(interfaces[i].(string))
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "events")) {
				details.Events = tmp
			}
		}
		if exclude, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "exclude")); ok {
			if tmpList := exclude.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "exclude"), 0)
				tmp, err := s.mapToVbsFilterExclusionAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert exclude, encountered error: %v", err)
				}
				details.Exclude = &tmp
			}
		}
		if include, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include")); ok {
			if tmpList := include.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "include"), 0)
				tmp, err := s.mapToVbsFilterAttributes(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert include, encountered error: %v", err)
				}
				details.Include = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown trigger_source '%v' was specified", triggerSource)
	}
	return baseObject, nil
}

func FilterToMap(obj *oci_devops.Filter) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_devops.BitbucketCloudFilter:
		result["trigger_source"] = "BITBUCKET_CLOUD"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{BitbucketCloudFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{BitbucketCloudFilterAttributesToMap(v.Include)}
		}
	case oci_devops.BitbucketServerFilter:
		result["trigger_source"] = "BITBUCKET_SERVER"

		result["events"] = v.Events

		if v.Include != nil {
			result["include"] = []interface{}{BitbucketServerFilterAttributesToMap(v.Include)}
		}
	case oci_devops.DevopsCodeRepositoryFilter:
		result["trigger_source"] = "DEVOPS_CODE_REPOSITORY"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{DevopsCodeRepositoryFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{DevopsCodeRepositoryFilterAttributesToMap(v.Include)}
		}
	case oci_devops.GithubFilter:
		result["trigger_source"] = "GITHUB"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{GithubFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{GithubFilterAttributesToMap(v.Include)}
		}
	case oci_devops.GitlabFilter:
		result["trigger_source"] = "GITLAB"

		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{GitlabFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{GitlabFilterAttributesToMap(v.Include)}
		}
	case oci_devops.GitlabServerFilter:
		result["trigger_source"] = "GITLAB_SERVER"

		result["events"] = v.Events

		if v.Include != nil {
			result["include"] = []interface{}{GitlabServerFilterAttributesToMap(v.Include)}
		}
	case oci_devops.VbsFilter:
		result["trigger_source"] = "VBS"

		result["events"] = v.Events
		result["events"] = v.Events

		if v.Exclude != nil {
			result["exclude"] = []interface{}{VbsFilterExclusionAttributesToMap(v.Exclude)}
		}

		if v.Include != nil {
			result["include"] = []interface{}{VbsFilterAttributesToMap(v.Include)}
		}
	default:
		log.Printf("[WARN] Received 'trigger_source' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToGithubFilterAttributes(fieldKeyFormat string) (oci_devops.GithubFilterAttributes, error) {
	result := oci_devops.GithubFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func GithubFilterAttributesToMap(obj *oci_devops.GithubFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToGithubFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.GithubFilterExclusionAttributes, error) {
	result := oci_devops.GithubFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func GithubFilterExclusionAttributesToMap(obj *oci_devops.GithubFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToGitlabFilterAttributes(fieldKeyFormat string) (oci_devops.GitlabFilterAttributes, error) {
	result := oci_devops.GitlabFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func GitlabFilterAttributesToMap(obj *oci_devops.GitlabFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToGitlabFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.GitlabFilterExclusionAttributes, error) {
	result := oci_devops.GitlabFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func GitlabFilterExclusionAttributesToMap(obj *oci_devops.GitlabFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToGitlabServerFilterAttributes(fieldKeyFormat string) (oci_devops.GitlabServerFilterAttributes, error) {
	result := oci_devops.GitlabServerFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	return result, nil
}

func GitlabServerFilterAttributesToMap(obj *oci_devops.GitlabServerFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToNetworkChannel(fieldKeyFormat string) (oci_devops.NetworkChannel, error) {
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
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
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
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
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

func (s *DevopsBuildRunResourceCrud) mapToTriggerAction(fieldKeyFormat string) (oci_devops.TriggerAction, error) {
	var baseObject oci_devops.TriggerAction
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("TRIGGER_BUILD_PIPELINE"):
		details := oci_devops.TriggerBuildPipelineAction{}
		if buildPipelineId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "build_pipeline_id")); ok {
			tmp := buildPipelineId.(string)
			details.BuildPipelineId = &tmp
		}
		if filter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filter")); ok {
			if tmpList := filter.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "filter"), 0)
				tmp, err := s.mapToFilter(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert filter, encountered error: %v", err)
				}
				details.Filter = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func (s *DevopsBuildRunResourceCrud) mapToTriggerInfo(fieldKeyFormat string) (oci_devops.TriggerInfo, error) {
	result := oci_devops.TriggerInfo{}

	if actions, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "actions")); ok {
		interfaces := actions.([]interface{})
		tmp := make([]oci_devops.TriggerAction, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "actions"), stateDataIndex)
			converted, err := s.mapToTriggerAction(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "actions")) {
			result.Actions = tmp
		}
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	return result, nil
}

func TriggerInfoToMap(obj *oci_devops.TriggerInfo) map[string]interface{} {
	result := map[string]interface{}{}

	actions := []interface{}{}
	for _, item := range obj.Actions {
		actions = append(actions, TriggerActionToMap(item))
	}
	result["actions"] = actions

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToVbsFilterAttributes(fieldKeyFormat string) (oci_devops.VbsFilterAttributes, error) {
	result := oci_devops.VbsFilterAttributes{}

	if baseRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "base_ref")); ok {
		tmp := baseRef.(string)
		result.BaseRef = &tmp
	}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	if headRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "head_ref")); ok {
		tmp := headRef.(string)
		result.HeadRef = &tmp
	}

	if repositoryName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_name")); ok {
		tmp := repositoryName.(string)
		result.RepositoryName = &tmp
	}

	return result, nil
}

func VbsFilterAttributesToMap(obj *oci_devops.VbsFilterAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BaseRef != nil {
		result["base_ref"] = string(*obj.BaseRef)
	}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	if obj.HeadRef != nil {
		result["head_ref"] = string(*obj.HeadRef)
	}

	if obj.RepositoryName != nil {
		result["repository_name"] = string(*obj.RepositoryName)
	}

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToVbsFilterExclusionAttributes(fieldKeyFormat string) (oci_devops.VbsFilterExclusionAttributes, error) {
	result := oci_devops.VbsFilterExclusionAttributes{}

	if fileFilter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "file_filter")); ok {
		if tmpList := fileFilter.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "file_filter"), 0)
			tmp, err := s.mapToFileFilter(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert file_filter, encountered error: %v", err)
			}
			result.FileFilter = &tmp
		}
	}

	return result, nil
}

func VbsFilterExclusionAttributesToMap(obj *oci_devops.VbsFilterExclusionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FileFilter != nil {
		result["file_filter"] = []interface{}{FileFilterToMap(obj.FileFilter)}
	}

	return result
}

func VulnerabilityAuditSummaryToMap(obj oci_devops.VulnerabilityAuditSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BuildStageId != nil {
		result["build_stage_id"] = string(*obj.BuildStageId)
	}

	if obj.CommitHash != nil {
		result["commit_hash"] = string(*obj.CommitHash)
	}

	if obj.VulnerabilityAuditId != nil {
		result["vulnerability_audit_id"] = string(*obj.VulnerabilityAuditId)
	}

	return result
}

func VulnerabilityAuditSummaryCollectionToMap(obj *oci_devops.VulnerabilityAuditSummaryCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, VulnerabilityAuditSummaryToMap(item))
	}
	result["items"] = items

	return result
}

func (s *DevopsBuildRunResourceCrud) mapToBuildRunArgumentCollection(fieldKeyFormat string) (oci_devops.BuildRunArgumentCollection, error) {
	result := oci_devops.BuildRunArgumentCollection{}
	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.BuildRunArgument, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToBuildRunArgument(fieldKeyFormatNextLevel)
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

func (s *DevopsBuildRunResourceCrud) mapToBuildRunArgument(fieldKeyFormat string) (oci_devops.BuildRunArgument, error) {
	result := oci_devops.BuildRunArgument{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}
	if defaultValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := defaultValue.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *DevopsBuildRunResourceCrud) mapToCommitInfo(fieldKeyFormat string) (oci_devops.CommitInfo, error) {
	result := oci_devops.CommitInfo{}

	if commitHash, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "commit_hash")); ok {
		tmp := commitHash.(string)
		result.CommitHash = &tmp
	}

	if repositoryBranch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_branch")); ok {
		tmp := repositoryBranch.(string)
		result.RepositoryBranch = &tmp
	}

	if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
		tmp := repositoryUrl.(string)
		result.RepositoryUrl = &tmp
	}
	return result, nil
}
