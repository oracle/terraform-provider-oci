// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package media_services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_media_services "github.com/oracle/oci-go-sdk/v65/mediaservices"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MediaServicesMediaWorkflowResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createMediaServicesMediaWorkflow,
		Read:     readMediaServicesMediaWorkflow,
		Update:   updateMediaServicesMediaWorkflow,
		Delete:   deleteMediaServicesMediaWorkflow,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
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
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"compartment_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"related_resource_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"time_created": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.TimeDiffSuppressFunction,
						},

						// Computed
					},
				},
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"media_workflow_configuration_ids": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parameters": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				Elem:             schema.TypeString,
				DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
			},
			"tasks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"parameters": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.JsonStringDiffSuppressFunction,
						},
						"type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"version": {
							Type:             schema.TypeString,
							Required:         true,
							ValidateFunc:     tfresource.ValidateInt64TypeString,
							DiffSuppressFunc: tfresource.Int64StringDiffSuppressFunction,
						},

						// Optional
						"enable_parameter_reference": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"enable_when_referenced_parameter_equals": {
							Type:     schema.TypeMap,
							Optional: true,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"prerequisites": {
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

			// Computed
			"lifecyle_details": {
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
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createMediaServicesMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.CreateResource(d, sync)
}

func readMediaServicesMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.ReadResource(sync)
}

func updateMediaServicesMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteMediaServicesMediaWorkflow(d *schema.ResourceData, m interface{}) error {
	sync := &MediaServicesMediaWorkflowResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MediaServicesClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type MediaServicesMediaWorkflowResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_media_services.MediaServicesClient
	Res                    *oci_media_services.MediaWorkflow
	DisableNotFoundRetries bool
}

func (s *MediaServicesMediaWorkflowResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *MediaServicesMediaWorkflowResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *MediaServicesMediaWorkflowResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowLifecycleStateActive),
		string(oci_media_services.MediaWorkflowLifecycleStateNeedsAttention),
	}
}

func (s *MediaServicesMediaWorkflowResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *MediaServicesMediaWorkflowResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_media_services.MediaWorkflowLifecycleStateDeleted),
	}
}

func (s *MediaServicesMediaWorkflowResourceCrud) Create() error {
	request := oci_media_services.CreateMediaWorkflowRequest{}

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

	if locks, ok := s.D.GetOkExists("locks"); ok {
		interfaces := locks.([]interface{})
		tmp := make([]oci_media_services.ResourceLock, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
			converted, err := s.mapToResourceLock(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("locks") {
			request.Locks = tmp
		}
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
			request.MediaWorkflowConfigurationIds = tmp
		}
	}

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		err := json.Unmarshal([]byte(parameters.(string)), &request.Parameters)
		if err != nil {
			return err
		}
	}

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_media_services.MediaWorkflowTask, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToMediaWorkflowTask(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.CreateMediaWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflow
	return nil
}

func (s *MediaServicesMediaWorkflowResourceCrud) Get() error {
	request := oci_media_services.GetMediaWorkflowRequest{}

	tmp := s.D.Id()
	request.MediaWorkflowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.GetMediaWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflow
	return nil
}

func (s *MediaServicesMediaWorkflowResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_media_services.UpdateMediaWorkflowRequest{}

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

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
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
			request.MediaWorkflowConfigurationIds = tmp
		}
	}

	tmp := s.D.Id()
	request.MediaWorkflowId = &tmp

	if parameters, ok := s.D.GetOkExists("parameters"); ok {
		err := json.Unmarshal([]byte(parameters.(string)), &request.Parameters)
		if err != nil {
			return err
		}
	}

	if tasks, ok := s.D.GetOkExists("tasks"); ok {
		interfaces := tasks.([]interface{})
		tmp := make([]oci_media_services.MediaWorkflowTask, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tasks", stateDataIndex)
			converted, err := s.mapToMediaWorkflowTask(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tasks") {
			request.Tasks = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	response, err := s.Client.UpdateMediaWorkflow(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.MediaWorkflow
	return nil
}

func (s *MediaServicesMediaWorkflowResourceCrud) Delete() error {
	request := oci_media_services.DeleteMediaWorkflowRequest{}

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	tmp := s.D.Id()
	request.MediaWorkflowId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.DeleteMediaWorkflow(context.Background(), request)
	return err
}

func (s *MediaServicesMediaWorkflowResourceCrud) SetData() error {
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

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	locks := []interface{}{}
	for _, item := range s.Res.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	s.D.Set("locks", locks)

	s.D.Set("media_workflow_configuration_ids", s.Res.MediaWorkflowConfigurationIds)

	if s.Res.Parameters != nil {
		jsonStr, err := json.Marshal(s.Res.Parameters)
		if err == nil {
			s.D.Set("parameters", string(jsonStr))
		}
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	tasks := []interface{}{}
	for _, item := range s.Res.Tasks {
		tasks = append(tasks, MediaWorkflowTaskToMap(item))
	}
	s.D.Set("tasks", tasks)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", strconv.FormatInt(*s.Res.Version, 10))
	}

	return nil
}

func MediaWorkflowSummaryToMap(obj oci_media_services.MediaWorkflowSummary) map[string]interface{} {
	result := map[string]interface{}{}

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

	locks := []interface{}{}
	for _, item := range obj.Locks {
		locks = append(locks, ResourceLockToMap(item))
	}
	result["locks"] = locks

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

	if obj.Version != nil {
		result["version"] = strconv.FormatInt(*obj.Version, 10)
	}

	return result
}

func (s *MediaServicesMediaWorkflowResourceCrud) mapToMediaWorkflowTask(fieldKeyFormat string) (oci_media_services.MediaWorkflowTask, error) {
	result := oci_media_services.MediaWorkflowTask{}

	if enableParameterReference, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_parameter_reference")); ok {
		tmp := enableParameterReference.(string)
		result.EnableParameterReference = &tmp
	}

	if enableWhenReferencedParameterEquals, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_when_referenced_parameter_equals")); ok {
		result.EnableWhenReferencedParameterEquals = enableWhenReferencedParameterEquals.(map[string]interface{})
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if parameters, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "parameters")); ok {
		err := json.Unmarshal([]byte(parameters.(string)), &result.Parameters)
		if err != nil {
			return oci_media_services.MediaWorkflowTask{}, err
		}
	}

	if prerequisites, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prerequisites")); ok {
		interfaces := prerequisites.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "prerequisites")) {
			result.Prerequisites = tmp
		}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		tmp := type_.(string)
		result.Type = &tmp
	}

	if version, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "version")); ok {
		tmp := version.(string)
		tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			return result, fmt.Errorf("unable to convert version string: %s to an int64 and encountered error: %v", tmp, err)
		}
		result.Version = &tmpInt64
	}

	return result, nil
}

func MediaWorkflowTaskToMap(obj oci_media_services.MediaWorkflowTask) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EnableParameterReference != nil {
		result["enable_parameter_reference"] = string(*obj.EnableParameterReference)
	}

	result["enable_when_referenced_parameter_equals"] = obj.EnableWhenReferencedParameterEquals
	result["enable_when_referenced_parameter_equals"] = obj.EnableWhenReferencedParameterEquals

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	jsonStr, err := json.Marshal(obj.Parameters)
	if err == nil {
		result["parameters"] = string(jsonStr)
	}

	if obj.Prerequisites != nil {
		result["prerequisites"] = obj.Prerequisites
	}

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	if obj.Version != nil {
		result["version"] = strconv.FormatInt(*obj.Version, 10)
	}

	return result
}

func (s *MediaServicesMediaWorkflowResourceCrud) mapToResourceLock(fieldKeyFormat string) (oci_media_services.ResourceLock, error) {
	result := oci_media_services.ResourceLock{}

	if compartmentId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "compartment_id")); ok {
		tmp := compartmentId.(string)
		result.CompartmentId = &tmp
	}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if relatedResourceId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "related_resource_id")); ok {
		tmp := relatedResourceId.(string)
		result.RelatedResourceId = &tmp
	}

	if timeCreated, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "time_created")); ok {
		tmp, err := time.Parse(time.RFC3339, timeCreated.(string))
		if err != nil {
			return result, err
		}
		result.TimeCreated = &oci_common.SDKTime{Time: tmp}
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_media_services.ResourceLockTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *MediaServicesMediaWorkflowResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_media_services.ChangeMediaWorkflowCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	idTmp := s.D.Id()
	changeCompartmentRequest.MediaWorkflowId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "media_services")

	_, err := s.Client.ChangeMediaWorkflowCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
