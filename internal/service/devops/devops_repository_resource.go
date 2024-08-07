// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package devops

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_devops "github.com/oracle/oci-go-sdk/v65/devops"
)

func DevopsRepositoryResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDevopsRepository,
		Read:     readDevopsRepository,
		Update:   updateDevopsRepository,
		Delete:   deleteDevopsRepository,
		Schema: map[string]*schema.Schema{
			// Required
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"repository_type": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"default_branch": {
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
			"description": {
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
			"mirror_repository_config": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"connector_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"repository_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trigger_schedule": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"schedule_type": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"custom_schedule": {
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
			"parent_repository_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"branch_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"commit_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecyle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size_in_bytes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_url": {
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
			"trigger_build_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createDevopsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.CreateResource(d, sync)
}

func readDevopsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.ReadResource(sync)
}

func updateDevopsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDevopsRepository(d *schema.ResourceData, m interface{}) error {
	sync := &DevopsRepositoryResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DevopsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DevopsRepositoryResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_devops.DevopsClient
	Res                    *oci_devops.Repository
	DisableNotFoundRetries bool
}

func (s *DevopsRepositoryResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DevopsRepositoryResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_devops.RepositoryLifecycleStateCreating),
	}
}

func (s *DevopsRepositoryResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_devops.RepositoryLifecycleStateActive),
	}
}

func (s *DevopsRepositoryResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DevopsRepositoryResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_devops.RepositoryLifecycleStateDeleted),
	}
}

func (s *DevopsRepositoryResourceCrud) Create() error {
	request := oci_devops.CreateRepositoryRequest{}

	if defaultBranch, ok := s.D.GetOkExists("default_branch"); ok {
		tmp := defaultBranch.(string)
		request.DefaultBranch = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mirrorRepositoryConfig, ok := s.D.GetOkExists("mirror_repository_config"); ok {
		if tmpList := mirrorRepositoryConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mirror_repository_config", 0)
			tmp, err := s.mapToMirrorRepositoryConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MirrorRepositoryConfig = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if parentRepositoryId, ok := s.D.GetOkExists("parent_repository_id"); ok {
		tmp := parentRepositoryId.(string)
		request.ParentRepositoryId = &tmp
	}

	if projectId, ok := s.D.GetOkExists("project_id"); ok {
		tmp := projectId.(string)
		request.ProjectId = &tmp
	}

	if repositoryType, ok := s.D.GetOkExists("repository_type"); ok {
		request.RepositoryType = oci_devops.RepositoryRepositoryTypeEnum(repositoryType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.CreateRepository(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getRepositoryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *DevopsRepositoryResourceCrud) getRepositoryFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_devops.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	repositoryId, err := repositoryWaitForWorkRequest(workId, "repository",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*repositoryId)

	return s.Get()
}

func repositoryWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func repositoryWaitForWorkRequest(wId *string, entityType string, action oci_devops.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_devops.DevopsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "devops")
	retryPolicy.ShouldRetryOperation = repositoryWorkRequestShouldRetryFunc(timeout)

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
	if identifier == nil || response.Status == oci_devops.OperationStatusFailed || response.Status == oci_devops.OperationStatusCanceled {
		return nil, getErrorFromDevopsRepositoryWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromDevopsRepositoryWorkRequest(client *oci_devops.DevopsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_devops.ActionTypeEnum) error {
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

func (s *DevopsRepositoryResourceCrud) Get() error {
	request := oci_devops.GetRepositoryRequest{}

	if fields, ok := s.D.GetOkExists("fields"); ok {
		interfaces := fields.([]interface{})
		tmp := make([]oci_devops.GetRepositoryFieldsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(oci_devops.GetRepositoryFieldsEnum)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("fields") {
			request.Fields = tmp
		}
	}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.GetRepository(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Repository
	return nil
}

func (s *DevopsRepositoryResourceCrud) Update() error {
	request := oci_devops.UpdateRepositoryRequest{}

	if defaultBranch, ok := s.D.GetOkExists("default_branch"); ok {
		tmp := defaultBranch.(string)
		request.DefaultBranch = &tmp
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

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if mirrorRepositoryConfig, ok := s.D.GetOkExists("mirror_repository_config"); ok {
		if tmpList := mirrorRepositoryConfig.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "mirror_repository_config", 0)
			tmp, err := s.mapToMirrorRepositoryConfig(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.MirrorRepositoryConfig = &tmp
		}
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	if repositoryType, ok := s.D.GetOkExists("repository_type"); ok {
		request.RepositoryType = oci_devops.RepositoryRepositoryTypeEnum(repositoryType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.UpdateRepository(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getRepositoryFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops"), oci_devops.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *DevopsRepositoryResourceCrud) Delete() error {
	request := oci_devops.DeleteRepositoryRequest{}

	tmp := s.D.Id()
	request.RepositoryId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "devops")

	response, err := s.Client.DeleteRepository(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := repositoryWaitForWorkRequest(workId, "repository",
		oci_devops.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *DevopsRepositoryResourceCrud) SetData() error {
	if s.Res.BranchCount != nil {
		s.D.Set("branch_count", *s.Res.BranchCount)
	}

	if s.Res.CommitCount != nil {
		s.D.Set("commit_count", *s.Res.CommitCount)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefaultBranch != nil {
		s.D.Set("default_branch", *s.Res.DefaultBranch)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HttpUrl != nil {
		s.D.Set("http_url", *s.Res.HttpUrl)
	}

	if s.Res.LifecyleDetails != nil {
		s.D.Set("lifecyle_details", *s.Res.LifecyleDetails)
	}

	if s.Res.MirrorRepositoryConfig != nil {
		s.D.Set("mirror_repository_config", []interface{}{MirrorRepositoryConfigToMap(s.Res.MirrorRepositoryConfig)})
	} else {
		s.D.Set("mirror_repository_config", nil)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Namespace != nil {
		s.D.Set("namespace", *s.Res.Namespace)
	}

	if s.Res.ParentRepositoryId != nil {
		s.D.Set("parent_repository_id", *s.Res.ParentRepositoryId)
	}

	if s.Res.ProjectId != nil {
		s.D.Set("project_id", *s.Res.ProjectId)
	}

	if s.Res.ProjectName != nil {
		s.D.Set("project_name", *s.Res.ProjectName)
	}

	s.D.Set("repository_type", s.Res.RepositoryType)

	if s.Res.SizeInBytes != nil {
		s.D.Set("size_in_bytes", strconv.FormatInt(*s.Res.SizeInBytes, 10))
	}

	if s.Res.SshUrl != nil {
		s.D.Set("ssh_url", *s.Res.SshUrl)
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

	s.D.Set("trigger_build_events", s.Res.TriggerBuildEvents)

	return nil
}

func (s *DevopsRepositoryResourceCrud) mapToMirrorRepositoryConfig(fieldKeyFormat string) (oci_devops.MirrorRepositoryConfig, error) {
	result := oci_devops.MirrorRepositoryConfig{}

	if connectorId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_id")); ok {
		tmp := connectorId.(string)
		result.ConnectorId = &tmp
	}

	if repositoryUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "repository_url")); ok {
		tmp := repositoryUrl.(string)
		result.RepositoryUrl = &tmp
	}

	if triggerSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "trigger_schedule")); ok {
		if tmpList := triggerSchedule.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "trigger_schedule"), 0)
			tmp, err := s.mapToTriggerSchedule(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert trigger_schedule, encountered error: %v", err)
			}
			result.TriggerSchedule = &tmp
		}
	}

	return result, nil
}

func MirrorRepositoryConfigToMap(obj *oci_devops.MirrorRepositoryConfig) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ConnectorId != nil {
		result["connector_id"] = string(*obj.ConnectorId)
	}

	if obj.RepositoryUrl != nil {
		result["repository_url"] = string(*obj.RepositoryUrl)
	}

	if obj.TriggerSchedule != nil {
		result["trigger_schedule"] = []interface{}{TriggerScheduleToMap(obj.TriggerSchedule)}
	}

	return result
}

func DevopsRepositorySummaryToMap(obj oci_devops.RepositorySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DefaultBranch != nil {
		result["default_branch"] = string(*obj.DefaultBranch)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.HttpUrl != nil {
		result["http_url"] = string(*obj.HttpUrl)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.MirrorRepositoryConfig != nil {
		result["mirror_repository_config"] = []interface{}{MirrorRepositoryConfigToMap(obj.MirrorRepositoryConfig)}
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Namespace != nil {
		result["namespace"] = string(*obj.Namespace)
	}

	if obj.ParentRepositoryId != nil {
		result["parent_repository_id"] = string(*obj.ParentRepositoryId)
	}

	if obj.ProjectId != nil {
		result["project_id"] = string(*obj.ProjectId)
	}

	if obj.ProjectName != nil {
		result["project_name"] = string(*obj.ProjectName)
	}

	result["repository_type"] = string(obj.RepositoryType)

	if obj.SshUrl != nil {
		result["ssh_url"] = string(*obj.SshUrl)
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

func (s *DevopsRepositoryResourceCrud) mapToTriggerSchedule(fieldKeyFormat string) (oci_devops.TriggerSchedule, error) {
	result := oci_devops.TriggerSchedule{}

	if customSchedule, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "custom_schedule")); ok {
		tmp := customSchedule.(string)
		result.CustomSchedule = &tmp
	}

	if scheduleType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "schedule_type")); ok {
		result.ScheduleType = oci_devops.TriggerScheduleScheduleTypeEnum(scheduleType.(string))
	}

	return result, nil
}

func TriggerScheduleToMap(obj *oci_devops.TriggerSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomSchedule != nil {
		result["custom_schedule"] = string(*obj.CustomSchedule)
	}

	result["schedule_type"] = string(obj.ScheduleType)

	return result
}
