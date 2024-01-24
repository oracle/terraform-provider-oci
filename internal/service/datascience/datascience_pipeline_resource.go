// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatasciencePipelineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatasciencePipeline,
		Read:     readDatasciencePipeline,
		Update:   updateDatasciencePipeline,
		Delete:   deleteDatasciencePipeline,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"step_details": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"step_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"step_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"CUSTOM_SCRIPT",
								"ML_JOB",
							}, true),
						},

						// Optional
						"depends_on": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_artifact_uploaded": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"job_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"step_configuration_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
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
									},
									"environment_variables": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"maximum_runtime_in_minutes": {
										Type:             schema.TypeString,
										Optional:         true,
										Computed:         true,
										ValidateFunc:     tfresource.ValidateInt64TypeString,
										DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
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

									// Optional
									"block_storage_size_in_gbs": {
										Type:     schema.TypeInt,
										Required: true,
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
												"ocpus": {
													Type:     schema.TypeFloat,
													Optional: true,
													Computed: true,
													ForceNew: true,
												},

												// Computed
											},
										},
									},
									"shape_name": {
										Type:     schema.TypeString,
										Required: true,
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

			// Optional
			"configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
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
						},
						"environment_variables": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"maximum_runtime_in_minutes": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
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
			"delete_related_pipeline_runs": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			// Optional
			"step_artifact": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"pipeline_step_artifact": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"artifact_content_length": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},
						"artifact_content_disposition": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						// Computed
						"artifact_content_md5": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"artifact_last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"step_name": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"infrastructure_configuration_details": {
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
									"ocpus": {
										Type:     schema.TypeFloat,
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
			"log_configuration_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
						},
						"enable_logging": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"log_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"created_by": {
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

func createDatasciencePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	// create pipeline before step artifact because the pipeline id is required in the case of custom
	// script
	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if StepArtifact, ok := d.GetOkExists("step_artifact"); ok {
		if tmpList := StepArtifact.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "step_artifact", 0)
			err := sync.CreateArtifact(fieldKeyFormat)
			if err != nil {
				return err
			}
		}
	}

	return tfresource.ReadResource(sync)
}

func (s *DatasciencePipelineResourceCrud) CreateArtifact(fieldKeyFormat string) error {
	request := oci_datascience.CreateStepArtifactRequest{}

	if contentDisposition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_content_disposition")); ok {
		tmp := contentDisposition.(string)
		request.ContentDisposition = &tmp
	}

	if contentLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "artifact_content_length")); ok {
		tmp := contentLength.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return fmt.Errorf("unable to convert Content-Length string: %s to an int64 and encountered error: %v", tmp, err)
		}
		request.ContentLength = &tmpInt64
	}

	if pipelineStepArtifact, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "pipeline_step_artifact")); ok {
		tmp := pipelineStepArtifact.(string)
		var artifactReader io.Reader
		artifactReader, err := os.Open(tmp)
		if err != nil {
			return fmt.Errorf("the specified pipeline step artifact is not available: %q", err)
		}
		request.StepArtifact = ioutil.NopCloser(artifactReader)
	}

	if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
		tmp := stepName.(string)
		request.StepName = &tmp
	}

	request.PipelineId = s.Res.Id

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.CreateStepArtifact(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func readDatasciencePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatasciencePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatasciencePipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DatasciencePipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type HeadPipelineArtifact struct {
	ContentLength      *int64
	ContentDisposition *string
	ContentMd5         *string
	LastModified       *common.SDKTime
}

type StepArtifact struct {
	StepName     *string
	StepArtifact *HeadPipelineArtifact
}

type DatasciencePipelineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.Pipeline
	StepArtifactRes        []StepArtifact
	DisableNotFoundRetries bool
}

func (s *DatasciencePipelineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatasciencePipelineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_datascience.PipelineLifecycleStateCreating),
	}
}

func (s *DatasciencePipelineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_datascience.PipelineLifecycleStateCreating),
		string(oci_datascience.PipelineLifecycleStateActive),
	}
}

func (s *DatasciencePipelineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_datascience.PipelineLifecycleStateDeleting),
	}
}

func (s *DatasciencePipelineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.PipelineLifecycleStateDeleted),
	}
}

func (s *DatasciencePipelineResourceCrud) mapToPipelineStepDetails(fieldKeyFormat string) (oci_datascience.PipelineStepDetails, error) {
	var baseObject oci_datascience.PipelineStepDetails
	//discriminator
	stepTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_type"))
	var stepType string
	if ok {
		stepType = stepTypeRaw.(string)
	} else {
		stepType = "" // default value
	}
	switch strings.ToLower(stepType) {
	case strings.ToLower("CUSTOM_SCRIPT"):
		details := oci_datascience.PipelineCustomScriptStepDetails{}
		if dependsOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "depends_on")); ok {
			interfaces := dependsOn.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "depends_on")) {
				details.DependsOn = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if stepInfrastructureConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_infrastructure_configuration_details")); ok {
			if tmpList := stepInfrastructureConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_infrastructure_configuration_details"), 0)
				tmp, err := s.mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormat)
				if err != nil {
					return details, fmt.Errorf("unable to convert step_infrastructure_configuration_details, encountered error: %v", err)
				}
				details.StepInfrastructureConfigurationDetails = &tmp
			}
		}
		if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
			if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
				tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
				}
				details.StepConfigurationDetails = &tmp
			}
		}
		if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
			tmp := stepName.(string)
			details.StepName = &tmp
		}
		baseObject = details
	case strings.ToLower("ML_JOB"):
		details := oci_datascience.PipelineMlJobStepDetails{}
		if dependsOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "depends_on")); ok {
			interfaces := dependsOn.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "depends_on")) {
				details.DependsOn = tmp
			}
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if jobId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_id")); ok {
			tmp := jobId.(string)
			details.JobId = &tmp
		}
		if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
			if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
				tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
				}
				details.StepConfigurationDetails = &tmp
			}
		}
		if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
			tmp := stepName.(string)
			details.StepName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown step_type '%v' was specified", stepType)
	}
	return baseObject, nil
}

func (s *DatasciencePipelineResourceCrud) mapToPipelineStepUpdateDetails(fieldKeyFormat string) (oci_datascience.PipelineStepUpdateDetails, error) {
	var baseObject oci_datascience.PipelineStepUpdateDetails
	//discriminator
	stepTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_type"))
	var stepType string
	if ok {
		stepType = stepTypeRaw.(string)
	} else {
		stepType = "" // default value
	}
	switch strings.ToLower(stepType) {
	case strings.ToLower("CUSTOM_SCRIPT"):
		details := oci_datascience.PipelineCustomScriptStepUpdateDetails{}
		if dependsOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "depends_on")); ok {
			interfaces := dependsOn.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			// if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "depends_on")) {
			// 	details.DependsOn = tmp
			// }
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
			if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
				tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
				}
				details.StepConfigurationDetails = &tmp
			}
		}
		if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
			tmp := stepName.(string)
			details.StepName = &tmp
		}
		baseObject = details
	case strings.ToLower("ML_JOB"):
		details := oci_datascience.PipelineMlJobStepUpdateDetails{}
		if dependsOn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "depends_on")); ok {
			interfaces := dependsOn.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			// if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "depends_on")) {
			// 	details.DependsOn = tmp
			// }
		}
		if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if stepConfigurationDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_configuration_details")); ok {
			if tmpList := stepConfigurationDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "step_configuration_details"), 0)
				tmp, err := s.mapToPipelineStepConfigurationDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert step_configuration_details, encountered error: %v", err)
				}
				details.StepConfigurationDetails = &tmp
			}
		}
		if stepName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "step_name")); ok {
			tmp := stepName.(string)
			details.StepName = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown step_type '%v' was specified", stepType)
	}
	return baseObject, nil
}

func (s *DatasciencePipelineResourceCrud) Create() error {
	request := oci_datascience.CreatePipelineRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if configurationDetails, ok := s.D.GetOkExists("configuration_details"); ok {
		if tmpList := configurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_details", 0)
			tmp, err := s.mapToPipelineConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationDetails = tmp
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

	if infrastructureConfigurationDetails, ok := s.D.GetOkExists("infrastructure_configuration_details"); ok {
		if tmpList := infrastructureConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "infrastructure_configuration_details", 0)
			tmp, err := s.mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InfrastructureConfigurationDetails = &tmp
		}
	}

	if logConfigurationDetails, ok := s.D.GetOkExists("log_configuration_details"); ok {
		if tmpList := logConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_configuration_details", 0)
			tmp, err := s.mapToPipelineLogConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogConfigurationDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if stepDetails, ok := s.D.GetOkExists("step_details"); ok {
		interfaces := stepDetails.([]interface{})
		tmp := make([]oci_datascience.PipelineStepDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "step_details", stateDataIndex)
			converted, err := s.mapToPipelineStepDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("step_details") {
			request.StepDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreatePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Pipeline
	return nil
}

func (s *DatasciencePipelineResourceCrud) getPipelineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_datascience.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	pipelineId, err := pipelineWaitForWorkRequest(workId, "datascience",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, pipelineId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_datascience.CancelWorkRequestRequest{
				WorkRequestId: workId,
				RequestMetadata: oci_common.RequestMetadata{
					RetryPolicy: retryPolicy,
				},
			})
		if cancelErr != nil {
			log.Printf("[DEBUG] cleanup cancelWorkRequest failed with the error: %v\n", cancelErr)
		}
		return err
	}
	s.D.SetId(*pipelineId)

	return s.Get()
}

func pipelineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "datascience", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_datascience.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func pipelineWaitForWorkRequest(wId *string, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_datascience.DataScienceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "datascience")
	retryPolicy.ShouldRetryOperation = pipelineWorkRequestShouldRetryFunc(timeout)

	response := oci_datascience.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_datascience.WorkRequestStatusInProgress),
			string(oci_datascience.WorkRequestStatusAccepted),
			string(oci_datascience.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_datascience.WorkRequestStatusSucceeded),
			string(oci_datascience.WorkRequestStatusFailed),
			string(oci_datascience.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_datascience.GetWorkRequestRequest{
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
	if response.Status == oci_datascience.WorkRequestStatusFailed || response.Status == oci_datascience.WorkRequestStatusCanceled {
		return nil, getErrorFromDatasciencePipelineWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDatasciencePipelineWorkRequest(client *oci_datascience.DataScienceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_datascience.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_datascience.ListWorkRequestErrorsRequest{
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

func (s *DatasciencePipelineResourceCrud) Get() error {
	request := oci_datascience.GetPipelineRequest{}

	tmp := s.D.Id()
	request.PipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Pipeline
	log.Printf("[DEBUG] the response from the pipeline get call is %v", s.Res)
	// after we get the pipeline response, GET the head artifact for each of the custom_script steps
	for _, item := range s.Res.StepDetails {
		err := s.ExtractPipelineStepDetailsAndCallHeadArtifact(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DatasciencePipelineResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdatePipelineRequest{}

	if configurationDetails, ok := s.D.GetOkExists("configuration_details"); ok {
		if tmpList := configurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "configuration_details", 0)
			tmp, err := s.mapToPipelineConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ConfigurationDetails = tmp
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

	if logConfigurationDetails, ok := s.D.GetOkExists("log_configuration_details"); ok {
		if tmpList := logConfigurationDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "log_configuration_details", 0)
			tmp, err := s.mapToPipelineLogConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.LogConfigurationDetails = &tmp
		}
	}

	tmp := s.D.Id()
	request.PipelineId = &tmp

	if stepDetails, ok := s.D.GetOkExists("step_details"); ok {
		interfaces := stepDetails.([]interface{})
		tmp := make([]oci_datascience.PipelineStepUpdateDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "step_details", stateDataIndex)
			converted, err := s.mapToPipelineStepUpdateDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("step_details") {
			request.StepDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdatePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Pipeline
	return nil
}

func (s *DatasciencePipelineResourceCrud) Delete() error {
	request := oci_datascience.DeletePipelineRequest{}

	if deleteRelatedJobRuns, ok := s.D.GetOkExists("delete_related_job_runs"); ok {
		tmp := deleteRelatedJobRuns.(bool)
		request.DeleteRelatedJobRuns = &tmp
	}

	if deleteRelatedPipelineRuns, ok := s.D.GetOkExists("delete_related_pipeline_runs"); ok {
		tmp := deleteRelatedPipelineRuns.(bool)
		request.DeleteRelatedPipelineRuns = &tmp
	}

	tmp := s.D.Id()
	request.PipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.DeletePipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := pipelineWaitForWorkRequest(workId, "datascience",
		oci_datascience.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DatasciencePipelineResourceCrud) SetData() error {
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

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
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

	if s.Res.InfrastructureConfigurationDetails != nil {
		s.D.Set("infrastructure_configuration_details", []interface{}{PipelineInfrastructureConfigurationDetailsToMap(s.Res.InfrastructureConfigurationDetails)})
	} else {
		s.D.Set("infrastructure_configuration_details", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogConfigurationDetails != nil {
		s.D.Set("log_configuration_details", []interface{}{PipelineLogConfigurationDetailsToMap(s.Res.LogConfigurationDetails)})
	} else {
		s.D.Set("log_configuration_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	stepDetails := []interface{}{}
	for _, item := range s.Res.StepDetails {
		stepDetails = append(stepDetails, PipelineStepDetailsToMap(item))
	}
	s.D.Set("step_details", stepDetails)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}
	// add head artifact [stepName:{artifact_content}]
	stepArtifacts := []interface{}{}
	for _, item := range s.StepArtifactRes {
		stepArtifacts = append(stepArtifacts, PipelineStepArtifactToMap(item))
	}
	if s.Res.LifecycleDetails != nil && *s.Res.LifecycleDetails == "ACTIVE" {
		log.Printf("[DEBUG] Step artifact content is :%v", stepArtifacts)
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	return nil
}

func (s *DatasciencePipelineResourceCrud) mapToPipelineConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineConfigurationDetails, error) {
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

// func PipelineConfigurationDetailsToMap(obj *oci_datascience.PipelineConfigurationDetails) map[string]interface{} {
// 	result := map[string]interface{}{}
// 	switch v := (*obj).(type) {
// 	case oci_datascience.PipelineDefaultConfigurationDetails:
// 		result["type"] = "DEFAULT"

// 		if v.CommandLineArguments != nil {
// 			result["command_line_arguments"] = string(*v.CommandLineArguments)
// 		}

// 		result["environment_variables"] = v.EnvironmentVariables

// 		if v.MaximumRuntimeInMinutes != nil {
// 			result["maximum_runtime_in_minutes"] = strconv.FormatInt(*v.MaximumRuntimeInMinutes, 10)
// 		}
// 	default:
// 		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
// 		return nil
// 	}

// 	return result
// }

func (s *DatasciencePipelineResourceCrud) mapToPipelineInfrastructureConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineInfrastructureConfigurationDetails, error) {
	result := oci_datascience.PipelineInfrastructureConfigurationDetails{}

	if blockStorageSizeInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "block_storage_size_in_gbs")); ok {
		tmp := blockStorageSizeInGBs.(int)
		result.BlockStorageSizeInGBs = &tmp
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

	return result, nil
}

func PipelineInfrastructureConfigurationDetailsToMap(obj *oci_datascience.PipelineInfrastructureConfigurationDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BlockStorageSizeInGBs != nil {
		result["block_storage_size_in_gbs"] = int(*obj.BlockStorageSizeInGBs)
	}

	if obj.ShapeConfigDetails != nil {
		result["shape_config_details"] = []interface{}{PipelineShapeConfigDetailsToMap(obj.ShapeConfigDetails)}
	}

	if obj.ShapeName != nil {
		result["shape_name"] = string(*obj.ShapeName)
	}

	return result
}

func (s *DatasciencePipelineResourceCrud) mapToPipelineLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineLogConfigurationDetails, error) {
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

// func PipelineLogConfigurationDetailsToMap(obj *oci_datascience.PipelineLogConfigurationDetails) map[string]interface{} {
// 	result := map[string]interface{}{}

// 	if obj.EnableAutoLogCreation != nil {
// 		result["enable_auto_log_creation"] = bool(*obj.EnableAutoLogCreation)
// 	}

// 	if obj.EnableLogging != nil {
// 		result["enable_logging"] = bool(*obj.EnableLogging)
// 	}

// 	if obj.LogGroupId != nil {
// 		result["log_group_id"] = string(*obj.LogGroupId)
// 	}

// 	if obj.LogId != nil {
// 		result["log_id"] = string(*obj.LogId)
// 	}

// 	return result
// }

func (s *DatasciencePipelineResourceCrud) mapToPipelineShapeConfigDetails(fieldKeyFormat string) (oci_datascience.PipelineShapeConfigDetails, error) {
	result := oci_datascience.PipelineShapeConfigDetails{}

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

func PipelineShapeConfigDetailsToMap(obj *oci_datascience.PipelineShapeConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	return result
}

func (s *DatasciencePipelineResourceCrud) mapToPipelineStepConfigurationDetails(fieldKeyFormat string) (oci_datascience.PipelineStepConfigurationDetails, error) {
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

// func PipelineStepConfigurationDetailsToMap(obj *oci_datascience.PipelineStepConfigurationDetails) map[string]interface{} {
// 	result := map[string]interface{}{}

// 	if obj.CommandLineArguments != nil {
// 		result["command_line_arguments"] = string(*obj.CommandLineArguments)
// 	}

// 	result["environment_variables"] = obj.EnvironmentVariables

// 	if obj.MaximumRuntimeInMinutes != nil {
// 		result["maximum_runtime_in_minutes"] = strconv.FormatInt(*obj.MaximumRuntimeInMinutes, 10)
// 	}

// 	return result
// }

func (s *DatasciencePipelineResourceCrud) ExtractPipelineStepDetailsAndCallHeadArtifact(obj oci_datascience.PipelineStepDetails) error {
	switch v := (obj).(type) {
	case oci_datascience.PipelineCustomScriptStepDetails:
		if v.StepName != nil && v.IsArtifactUploaded != nil {
			if stepArtifact, err := s.GetArtifactHead(*v.StepName); stepArtifact != nil {
				if err != nil {
					return err
				} // if err
				s.StepArtifactRes = append(s.StepArtifactRes, stepArtifact.(StepArtifact))
			}
		} // if the step name and artifact uploaded items are not nil
	}
	return nil
}

func PipelineStepArtifactToMap(v StepArtifact) map[string]interface{} {
	result := map[string]interface{}{}
	result["step_name"] = v.StepName
	if v.StepArtifact != nil {
		if v.StepArtifact.ContentLength != nil {
			result["artifact_content_length"] = v.StepArtifact.ContentLength
		}
		if v.StepArtifact.ContentMd5 != nil {
			result["artifact_content_md5"] = v.StepArtifact.ContentMd5
		}
		if v.StepArtifact.ContentDisposition != nil {
			result["artifact_content_disposition"] = v.StepArtifact.ContentDisposition
		}
		if v.StepArtifact.LastModified != nil {
			result["artifact_last_modified"] = v.StepArtifact.LastModified.String()
		}
	}
	return result
}

func PipelineStepDetailsToMap(obj oci_datascience.PipelineStepDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_datascience.PipelineCustomScriptStepDetails:
		result["step_type"] = "CUSTOM_SCRIPT"

		result["depends_on"] = v.DependsOn

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.StepConfigurationDetails != nil {
			result["step_configuration_details"] = []interface{}{PipelineStepConfigurationDetailsToMap(v.StepConfigurationDetails)}
		}

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}

		result["is_artifact_uploaded"] = v.IsArtifactUploaded

		if v.StepInfrastructureConfigurationDetails != nil {
			result["step_infrastructure_configuration_details"] = []interface{}{PipelineInfrastructureConfigurationDetailsToMap(v.StepInfrastructureConfigurationDetails)}
		}
	case oci_datascience.PipelineMlJobStepDetails:
		result["step_type"] = "ML_JOB"

		result["depends_on"] = v.DependsOn

		if v.Description != nil {
			result["description"] = string(*v.Description)
		}

		if v.JobId != nil {
			result["job_id"] = string(*v.JobId)
		}

		if v.StepConfigurationDetails != nil {
			result["step_configuration_details"] = []interface{}{PipelineStepConfigurationDetailsToMap(v.StepConfigurationDetails)}
		}

		if v.StepName != nil {
			result["step_name"] = string(*v.StepName)
		}
	default:
		log.Printf("[WARN] Received 'step_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *DatasciencePipelineResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangePipelineCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.PipelineId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangePipelineCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}

func (s *DatasciencePipelineResourceCrud) GetArtifactHead(stepName string) (interface{}, error) {
	request := oci_datascience.HeadStepArtifactRequest{}
	result := StepArtifact{}

	tmp := s.D.Id()
	request.PipelineId = &tmp
	request.StepName = &stepName

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "datascience")

	response, err := s.Client.HeadStepArtifact(context.Background(), request)
	if err != nil {
		return nil, err
	}

	result.StepName = &stepName

	result.StepArtifact = &HeadPipelineArtifact{
		ContentLength:      response.ContentLength,
		ContentDisposition: response.ContentDisposition,
		ContentMd5:         response.ContentMd5,
		LastModified:       response.LastModified,
	}
	return result, nil
}

// func (s *DatasciencePipelineResourceCrud) SetArtifactDataForStep() map[string]interface{} {
// 	result := map[string]interface{}{}
// 	if s.ArtifactHeadRes == nil {
// 		return nil
// 	}

// 	if s.ArtifactHeadRes.ContentDisposition != nil {
// 		result["artifact_content_disposition"] = *s.ArtifactHeadRes.ContentDisposition
// 	}

// 	if s.ArtifactHeadRes.ContentLength != nil {
// 		result["artifact_content_length"] = *s.ArtifactHeadRes.ContentLength
// 	}

// 	if s.ArtifactHeadRes.ContentMd5 != nil {
// 		result["artifact_content_md5"] = *s.ArtifactHeadRes.ContentMd5
// 	}

// 	if s.ArtifactHeadRes.LastModified != nil {
// 		result["artifact_last_modified"] = s.ArtifactHeadRes.LastModified.String()
// 	}

// 	return result
// }
