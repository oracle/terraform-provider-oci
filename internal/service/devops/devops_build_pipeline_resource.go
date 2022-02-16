// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
	oci_devops "github.com/oracle/oci-go-sdk/v58/devops"
)

func DevopsBuildPipelineResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsBuildPipeline,
		Read:     readDevopsBuildPipeline,
		Update:   updateDevopsBuildPipeline,
		Delete:   deleteDevopsBuildPipeline,
		Schema: map[string]*schema.Schema{
			// Required
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"build_pipeline_parameters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
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
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"default_value": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
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

			// Computed
			"compartment_id": {
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

func createDevopsBuildPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsBuildPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsBuildPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsBuildPipeline(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsBuildPipelineResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsBuildPipelineResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.BuildPipeline
	DisableNotFoundRetries bool
}

func (s *DevopsBuildPipelineResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DevopsBuildPipelineResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.BuildPipelineLifecycleStateCreating),
	}
}

func (s *DevopsBuildPipelineResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.BuildPipelineLifecycleStateActive),
	}
}

func (s *DevopsBuildPipelineResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_devops.BuildPipelineLifecycleStateDeleting),
	}
}

func (s *DevopsBuildPipelineResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.BuildPipelineLifecycleStateDeleted),
	}
}

func (s *DevopsBuildPipelineResourceCrud) Create() error {
	request := oci_devops.CreateBuildPipelineRequest{}

	if buildPipelineParameters, ok := s.D.GetOkExists("build_pipeline_parameters"); ok {
		if tmpList := buildPipelineParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_parameters", 0)
			tmp, err := s.mapToBuildPipelineParameterCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BuildPipelineParameters = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateBuildPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBuildPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsBuildPipelineResourceCrud) getBuildPipelineFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	buildPipelineId, err := buildPipelineWaitForWorkRequest(workId, "buildPipeline",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*buildPipelineId)

	return s.Get()
}

func buildPipelineWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func buildPipelineWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = buildPipelineWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromDevopsBuildPipelineWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsBuildPipelineWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsBuildPipelineResourceCrud) Get() error {
	request := oci_devops.GetBuildPipelineRequest{}

	tmp := s.D.Id()
	request.BuildPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetBuildPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BuildPipeline
	return nil
}

func (s *DevopsBuildPipelineResourceCrud) Update() error {
	request := oci_devops.UpdateBuildPipelineRequest{}

	tmp := s.D.Id()
	request.BuildPipelineId = &tmp

	if buildPipelineParameters, ok := s.D.GetOkExists("build_pipeline_parameters"); ok {
		if tmpList := buildPipelineParameters.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "build_pipeline_parameters", 0)
			tmp, err := s.mapToBuildPipelineParameterCollection(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BuildPipelineParameters = &tmp
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateBuildPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBuildPipelineFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsBuildPipelineResourceCrud) Delete() error {
	request := oci_devops.DeleteBuildPipelineRequest{}

	tmp := s.D.Id()
	request.BuildPipelineId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteBuildPipeline(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := buildPipelineWaitForWorkRequest(workId, "buildPipeline",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsBuildPipelineResourceCrud) SetData() error {
	if s.Res.BuildPipelineParameters != nil {
		s.D.Set("build_pipeline_parameters", []interface{}{BuildPipelineParameterCollectionToMap(s.Res.BuildPipelineParameters)})
	} else {
		s.D.Set("build_pipeline_parameters", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
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

func BuildPipelineParameterToMap(obj oci_devops.BuildPipelineParameter) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultValue != nil {
		result["default_value"] = string(*obj.DefaultValue)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func BuildPipelineParameterCollectionToMap(obj *oci_devops.BuildPipelineParameterCollection) map[string]interface{} {
	result := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range obj.Items {
		items = append(items, BuildPipelineParameterToMap(item))
	}
	result["items"] = items

	return result
}

func BuildPipelineSummaryToMap(obj oci_devops.BuildPipelineSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BuildPipelineParameters != nil {
		result["build_pipeline_parameters"] = []interface{}{BuildPipelineParameterCollectionToMap(obj.BuildPipelineParameters)}
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
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

func (s *DevopsBuildPipelineResourceCrud) mapToBuildPipelineParameterCollection(fieldKeyFormat string) (oci_devops.BuildPipelineParameterCollection, error) {
	result := oci_devops.BuildPipelineParameterCollection{}

	if items, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "items")); ok {
		interfaces := items.([]interface{})
		tmp := make([]oci_devops.BuildPipelineParameter, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "items"), stateDataIndex)
			converted, err := s.mapToBuildPipelineParameter(fieldKeyFormatNextLevel)
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

func (s *DevopsBuildPipelineResourceCrud) mapToBuildPipelineParameter(fieldKeyFormat string) (oci_devops.BuildPipelineParameter, error) {
	result := oci_devops.BuildPipelineParameter{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}
	if defaultValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_value")); ok {
		tmp := defaultValue.(string)
		result.DefaultValue = &tmp
	}
	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	return result, nil
}
