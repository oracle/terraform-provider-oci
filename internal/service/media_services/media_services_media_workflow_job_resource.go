// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesMediaWorkflowJob,
		Read:     readMediaServicesMediaWorkflowJob,
		Update:   updateMediaServicesMediaWorkflowJob,
		Delete:   deleteMediaServicesMediaWorkflowJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"workflow_identifier_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"ID",
					"NAME",
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
			"media_workflow_configuration_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"media_workflow_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"media_workflow_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"outputs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"asset_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"object": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"runnable": {
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
			"task_lifecycle_state": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"key": {
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
					},
				},
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_ended": {
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

func createMediaServicesMediaWorkflowJob(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesMediaWorkflowJob(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesMediaWorkflowJob(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesMediaWorkflowJob(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesMediaWorkflowJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.MediaWorkflowJob
	DisableNotFoundRetries bool
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_media_services.MediaWorkflowJobLifecycleStateInProgress),
	}
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowJobLifecycleStateAccepted),
		string(oci_media_services.MediaWorkflowJobLifecycleStateFailed),
		string(oci_media_services.MediaWorkflowJobLifecycleStateSucceeded),
	}
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_media_services.MediaWorkflowJobLifecycleStateCanceling),
	}
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowJobLifecycleStateSucceeded),
		string(oci_media_services.MediaWorkflowJobLifecycleStateCanceled),
		string(oci_media_services.MediaWorkflowJobLifecycleStateFailed),
	}
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) Create() error {
	request := oci_media_services.CreateMediaWorkflowJobRequest{}
	err := s.populateTopLevelPolymorphicCreateMediaWorkflowJobRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateMediaWorkflowJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowJob
	return nil
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) Get() error {
	request := oci_media_services.GetMediaWorkflowJobRequest{}

	tmp := s.D.Id()
	request.MediaWorkflowJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetMediaWorkflowJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowJob
	return nil
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_media_services.UpdateMediaWorkflowJobRequest{}

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
	request.MediaWorkflowJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateMediaWorkflowJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflowJob
	return nil
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) Delete() error {
	request := oci_media_services.DeleteMediaWorkflowJobRequest{}

	tmp := s.D.Id()
	request.MediaWorkflowJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")
	request.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 1

	response, err := s.Client.DeleteMediaWorkflowJob(context.Background(), request)
	if response.RawResponse.StatusCode == 409 {
		err = nil
	}
	return err
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) SetData() error {
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

	s.D.Set("media_workflow_configuration_ids", s.Res.MediaWorkflowConfigurationIds)

	if s.Res.MediaWorkflowId != nil {
		s.D.Set("media_workflow_id", *s.Res.MediaWorkflowId)
	}

	outputs := []interface{}{}
	for _, item := range s.Res.Outputs {
		outputs = append(outputs, JobOutputToMap(item))
	}
	s.D.Set("outputs", outputs)

	if s.Res.Parameters != nil {
		jsonStr, err := json.Marshal(s.Res.Parameters)
		if err == nil {
			s.D.Set("parameters", string(jsonStr))
		}
	}

	if s.Res.Runnable != nil {
		jsonStr, err := json.Marshal(s.Res.Runnable)
		if err == nil {
			s.D.Set("runnable", string(jsonStr))
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	taskLifecycleState := []interface{}{}
	for _, item := range s.Res.TaskLifecycleState {
		taskLifecycleState = append(taskLifecycleState, MediaWorkflowTaskStateToMap(item))
	}
	s.D.Set("task_lifecycle_state", taskLifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeEnded != nil {
		s.D.Set("time_ended", s.Res.TimeEnded.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func JobOutputToMap(obj oci_media_services.JobOutput) map[string]interface{} {
	result := map[string]interface{}{}

	result["asset_type"] = string(obj.AssetType)

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.ObjectName != nil {
		result["object"] = string(*obj.ObjectName)
	}

	return result
}

func MediaWorkflowJobSummaryToMap(obj oci_media_services.MediaWorkflowJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	if obj.MediaWorkflowId != nil {
		result["media_workflow_id"] = string(*obj.MediaWorkflowId)
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

func MediaWorkflowTaskStateToMap(obj oci_media_services.MediaWorkflowTaskState) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	return result
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) populateTopLevelPolymorphicCreateMediaWorkflowJobRequest(request *oci_media_services.CreateMediaWorkflowJobRequest) error {
	//discriminator
	workflowIdentifierTypeRaw, ok := s.D.GetOkExists("workflow_identifier_type")
	var workflowIdentifierType string
	if ok {
		workflowIdentifierType = workflowIdentifierTypeRaw.(string)
	} else {
		workflowIdentifierType = "" // default value
	}
	switch strings.ToLower(workflowIdentifierType) {
	case strings.ToLower("ID"):
		details := oci_media_services.CreateMediaWorkflowJobByIdDetails{}
		if mediaWorkflowId, ok := s.D.GetOkExists("media_workflow_id"); ok {
			tmp := mediaWorkflowId.(string)
			details.MediaWorkflowId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if mediaWorkflowConfigurationIds, ok := s.D.GetOkExists("media_workflow_configuration_ids"); ok {
			interfaces := mediaWorkflowConfigurationIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("media_workflow_configuration_ids") {
				details.MediaWorkflowConfigurationIds = tmp
			}
		}
		if mediaWorkflowId, ok := s.D.GetOkExists("media_workflow_id"); ok {
			tmp := mediaWorkflowId.(string)
			details.MediaWorkflowId = &tmp
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			err := json.Unmarshal([]byte(parameters.(string)), &details.Parameters)
			if err != nil {
				return err
			}
		}
		request.CreateMediaWorkflowJobDetails = details
	case strings.ToLower("NAME"):
		details := oci_media_services.CreateMediaWorkflowJobByNameDetails{}
		if mediaWorkflowName, ok := s.D.GetOkExists("media_workflow_name"); ok {
			tmp := mediaWorkflowName.(string)
			details.MediaWorkflowName = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if mediaWorkflowConfigurationIds, ok := s.D.GetOkExists("media_workflow_configuration_ids"); ok {
			interfaces := mediaWorkflowConfigurationIds.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("media_workflow_configuration_ids") {
				details.MediaWorkflowConfigurationIds = tmp
			}
		}
		if parameters, ok := s.D.GetOkExists("parameters"); ok {
			err := json.Unmarshal([]byte(parameters.(string)), &details.Parameters)
			if err != nil {
				return err
			}
		}
		request.CreateMediaWorkflowJobDetails = details
	default:
		return fmt.Errorf("unknown workflow_identifier_type '%v' was specified", workflowIdentifierType)
	}
	return nil
}

func (s *MediaServicesMediaWorkflowJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_media_services.ChangeMediaWorkflowJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.MediaWorkflowJobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.ChangeMediaWorkflowJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
