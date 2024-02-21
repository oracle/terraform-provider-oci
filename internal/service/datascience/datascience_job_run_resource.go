// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datascience

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_datascience "github.com/oracle/oci-go-sdk/v65/datascience"
)

func DatascienceJobRunResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatascienceJobRun,
		Read:     readDatascienceJobRun,
		Update:   updateDatascienceJobRun,
		Delete:   deleteDatascienceJobRun,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"job_id": {
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
			"asynchronous": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
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
			"job_configuration_override_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"job_type": {
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
			"job_environment_configuration_override_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"image": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"job_environment_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OCIR_CONTAINER",
							}, true),
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
			"job_log_configuration_override_details": {
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

			// Computed
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"job_infrastructure_configuration_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"block_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"job_infrastructure_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"job_shape_config_details": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"shape_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"job_storage_mount_configuration_details_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"bucket": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_directory_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"destination_path": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"export_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"mount_target_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"storage_type": {
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
		},
	}
}

func createDatascienceJobRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.CreateResource(d, sync)
}

func readDatascienceJobRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.ReadResource(sync)
}

func updateDatascienceJobRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatascienceJobRun(d *schema.ResourceData, m interface{}) error {
	sync := &DatascienceJobRunResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataScienceClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatascienceJobRunResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_datascience.DataScienceClient
	Res                    *oci_datascience.JobRun
	DisableNotFoundRetries bool
}

func (s *DatascienceJobRunResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatascienceJobRunResourceCrud) CreatedPending() []string {
	invokeAsynchronously := true
	if async, ok := s.D.GetOkExists("asynchronous"); ok {
		invokeAsynchronously = async.(bool)
	}

	if invokeAsynchronously {
		return []string{}
	}

	return []string{
		string(oci_datascience.JobRunLifecycleStateAccepted),
		string(oci_datascience.JobRunLifecycleStateInProgress),
	}
}

func (s *DatascienceJobRunResourceCrud) CreatedTarget() []string {
	invokeAsynchronously := true
	if async, ok := s.D.GetOkExists("asynchronous"); ok {
		invokeAsynchronously = async.(bool)
	}

	if invokeAsynchronously {
		return []string{
			string(oci_datascience.JobRunLifecycleStateAccepted),
		}
	}

	return []string{
		string(oci_datascience.JobRunLifecycleStateSucceeded),
		string(oci_datascience.JobRunLifecycleStateNeedsAttention),
		string(oci_datascience.JobRunLifecycleStateFailed),
		string(oci_datascience.JobRunLifecycleStateCanceled),
	}
}

func (s *DatascienceJobRunResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatascienceJobRunResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_datascience.JobRunLifecycleStateDeleted),
	}
}

func (s *DatascienceJobRunResourceCrud) Create() error {
	request := oci_datascience.CreateJobRunRequest{}

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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if jobConfigurationOverrideDetails, ok := s.D.GetOkExists("job_configuration_override_details"); ok {
		if tmpList := jobConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_configuration_override_details", 0)
			tmp, err := s.mapToJobConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobConfigurationOverrideDetails = tmp
		}
	}

	if jobEnvironmentConfigurationOverrideDetails, ok := s.D.GetOkExists("job_environment_configuration_override_details"); ok {
		if tmpList := jobEnvironmentConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_environment_configuration_override_details", 0)
			tmp, err := s.mapToJobEnvironmentConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobEnvironmentConfigurationOverrideDetails = tmp
		}
	}

	if jobId, ok := s.D.GetOkExists("job_id"); ok {
		tmp := jobId.(string)
		request.JobId = &tmp
	}

	if jobLogConfigurationOverrideDetails, ok := s.D.GetOkExists("job_log_configuration_override_details"); ok {
		if tmpList := jobLogConfigurationOverrideDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "job_log_configuration_override_details", 0)
			tmp, err := s.mapToJobLogConfigurationDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.JobLogConfigurationOverrideDetails = &tmp
		}
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.CreateJobRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JobRun
	return nil
}

func (s *DatascienceJobRunResourceCrud) Get() error {
	request := oci_datascience.GetJobRunRequest{}

	tmp := s.D.Id()
	request.JobRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.GetJobRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JobRun
	return nil
}

func (s *DatascienceJobRunResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_datascience.UpdateJobRunRequest{}

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
	request.JobRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	response, err := s.Client.UpdateJobRun(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.JobRun
	return nil
}

func (s *DatascienceJobRunResourceCrud) Delete() error {
	request := oci_datascience.DeleteJobRunRequest{}

	tmp := s.D.Id()
	request.JobRunId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.DeleteJobRun(context.Background(), request)
	return err
}

func (s *DatascienceJobRunResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

	if s.Res.JobConfigurationOverrideDetails != nil {
		jobConfigurationOverrideDetailsArray := []interface{}{}
		if jobConfigurationOverrideDetailsMap := JobConfigurationDetailsToMap(&s.Res.JobConfigurationOverrideDetails); jobConfigurationOverrideDetailsMap != nil {
			jobConfigurationOverrideDetailsArray = append(jobConfigurationOverrideDetailsArray, jobConfigurationOverrideDetailsMap)
		}
		s.D.Set("job_configuration_override_details", jobConfigurationOverrideDetailsArray)
	} else {
		s.D.Set("job_configuration_override_details", nil)
	}

	if s.Res.JobEnvironmentConfigurationOverrideDetails != nil {
		jobEnvironmentConfigurationOverrideDetailsArray := []interface{}{}
		if jobEnvironmentConfigurationOverrideDetailsMap := JobEnvironmentConfigurationDetailsToMap(&s.Res.JobEnvironmentConfigurationOverrideDetails); jobEnvironmentConfigurationOverrideDetailsMap != nil {
			jobEnvironmentConfigurationOverrideDetailsArray = append(jobEnvironmentConfigurationOverrideDetailsArray, jobEnvironmentConfigurationOverrideDetailsMap)
		}
		s.D.Set("job_environment_configuration_override_details", jobEnvironmentConfigurationOverrideDetailsArray)
	} else {
		s.D.Set("job_environment_configuration_override_details", nil)
	}

	if s.Res.JobId != nil {
		s.D.Set("job_id", *s.Res.JobId)
	}

	if s.Res.JobInfrastructureConfigurationDetails != nil {
		jobInfrastructureConfigurationDetailsArray := []interface{}{}
		if jobInfrastructureConfigurationDetailsMap := JobInfrastructureConfigurationDetailsToMap(&s.Res.JobInfrastructureConfigurationDetails); jobInfrastructureConfigurationDetailsMap != nil {
			jobInfrastructureConfigurationDetailsArray = append(jobInfrastructureConfigurationDetailsArray, jobInfrastructureConfigurationDetailsMap)
		}
		s.D.Set("job_infrastructure_configuration_details", jobInfrastructureConfigurationDetailsArray)
	} else {
		s.D.Set("job_infrastructure_configuration_details", nil)
	}

	if s.Res.JobLogConfigurationOverrideDetails != nil {
		s.D.Set("job_log_configuration_override_details", []interface{}{JobLogConfigurationDetailsToMap(s.Res.JobLogConfigurationOverrideDetails)})
	} else {
		s.D.Set("job_log_configuration_override_details", nil)
	}

	jobStorageMountConfigurationDetailsList := []interface{}{}
	for _, item := range s.Res.JobStorageMountConfigurationDetailsList {
		jobStorageMountConfigurationDetailsList = append(jobStorageMountConfigurationDetailsList, StorageMountConfigurationDetailsToMap(item))
	}
	s.D.Set("job_storage_mount_configuration_details_list", jobStorageMountConfigurationDetailsList)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LogDetails != nil {
		s.D.Set("log_details", []interface{}{JobRunLogDetailsToMap(s.Res.LogDetails)})
	} else {
		s.D.Set("log_details", nil)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeFinished != nil {
		s.D.Set("time_finished", s.Res.TimeFinished.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	return nil
}

func (s *DatascienceJobRunResourceCrud) mapToJobConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobConfigurationDetails, error) {
	var baseObject oci_datascience.JobConfigurationDetails
	//discriminator
	jobTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_type"))
	var jobType string
	if ok {
		jobType = jobTypeRaw.(string)
	} else {
		jobType = "" // default value
	}
	switch strings.ToLower(jobType) {
	case strings.ToLower("DEFAULT"):
		details := oci_datascience.DefaultJobConfigurationDetails{}
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
		return nil, fmt.Errorf("unknown job_type '%v' was specified", jobType)
	}
	return baseObject, nil
}

func (s *DatascienceJobRunResourceCrud) mapToJobEnvironmentConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobEnvironmentConfigurationDetails, error) {
	var baseObject oci_datascience.JobEnvironmentConfigurationDetails
	//discriminator
	jobEnvironmentTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "job_environment_type"))
	var jobEnvironmentType string
	if ok {
		jobEnvironmentType = jobEnvironmentTypeRaw.(string)
	} else {
		jobEnvironmentType = "" // default value
	}
	switch strings.ToLower(jobEnvironmentType) {
	case strings.ToLower("OCIR_CONTAINER"):
		details := oci_datascience.OcirContainerJobEnvironmentConfigurationDetails{}
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
		return nil, fmt.Errorf("unknown job_environment_type '%v' was specified", jobEnvironmentType)
	}
	return baseObject, nil
}

func (s *DatascienceJobRunResourceCrud) mapToJobLogConfigurationDetails(fieldKeyFormat string) (oci_datascience.JobLogConfigurationDetails, error) {
	result := oci_datascience.JobLogConfigurationDetails{}

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

func JobRunLogDetailsToMap(obj *oci_datascience.JobRunLogDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LogGroupId != nil {
		result["log_group_id"] = string(*obj.LogGroupId)
	}

	if obj.LogId != nil {
		result["log_id"] = string(*obj.LogId)
	}

	return result
}

func (s *DatascienceJobRunResourceCrud) mapToJobShapeConfigDetails(fieldKeyFormat string) (oci_datascience.JobShapeConfigDetails, error) {
	result := oci_datascience.JobShapeConfigDetails{}

	if memoryInGBs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "memory_in_gbs")); ok {
		tmp := memoryInGBs.(float32)
		result.MemoryInGBs = &tmp
	}

	if ocpus, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ocpus")); ok {
		tmp := ocpus.(float32)
		result.Ocpus = &tmp
	}

	return result, nil
}

func (s *DatascienceJobRunResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_datascience.ChangeJobRunCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.JobRunId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "datascience")

	_, err := s.Client.ChangeJobRunCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
