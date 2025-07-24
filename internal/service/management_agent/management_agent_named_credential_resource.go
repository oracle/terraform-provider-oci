// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package management_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_management_agent "github.com/oracle/oci-go-sdk/v65/managementagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ManagementAgentNamedCredentialResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createManagementAgentNamedCredential,
		Read:     readManagementAgentNamedCredential,
		Update:   updateManagementAgentNamedCredential,
		Delete:   deleteManagementAgentNamedCredential,
		Schema: map[string]*schema.Schema{
			// Required
			"management_agent_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"properties": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value_category": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
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

			// Computed
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

func createManagementAgentNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readManagementAgentNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.ReadResource(sync)
}

func updateManagementAgentNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteManagementAgentNamedCredential(d *schema.ResourceData, m interface{}) error {
	sync := &ManagementAgentNamedCredentialResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ManagementAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type ManagementAgentNamedCredentialResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_management_agent.ManagementAgentClient
	Res                    *oci_management_agent.NamedCredential
	DisableNotFoundRetries bool
}

func (s *ManagementAgentNamedCredentialResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ManagementAgentNamedCredentialResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_management_agent.NamedCredentialLifecycleStateCreating),
	}
}

func (s *ManagementAgentNamedCredentialResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_management_agent.NamedCredentialLifecycleStateActive),
	}
}

func (s *ManagementAgentNamedCredentialResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_management_agent.NamedCredentialLifecycleStateDeleting),
	}
}

func (s *ManagementAgentNamedCredentialResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_management_agent.NamedCredentialLifecycleStateDeleted),
	}
}

func (s *ManagementAgentNamedCredentialResourceCrud) Create() error {
	request := oci_management_agent.CreateNamedCredentialRequest{}

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

	if managementAgentId, ok := s.D.GetOkExists("management_agent_id"); ok {
		tmp := managementAgentId.(string)
		request.ManagementAgentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if properties, ok := s.D.GetOkExists("properties"); ok {
		interfaces := properties.([]interface{})
		tmp := make([]oci_management_agent.NamedCredentialProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToNamedCredentialProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	if type_, ok := s.D.GetOkExists("type"); ok {
		tmp := type_.(string)
		request.Type = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.CreateNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getNamedCredentialFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent"), oci_management_agent.ActionTypesCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ManagementAgentNamedCredentialResourceCrud) getNamedCredentialFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_management_agent.ActionTypesEnum, timeout time.Duration) error {

	// Wait until it finishes
	namedCredentialId, err := namedCredentialWaitForWorkRequest(workId, "namedcredential",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, namedCredentialId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_management_agent.DeleteWorkRequestRequest{
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
	//s.D.SetId(*namedCredentialId)

	return s.Get()
}

func namedCredentialWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "management_agent", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_management_agent.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func namedCredentialWaitForWorkRequest(wId *string, entityType string, action oci_management_agent.ActionTypesEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_management_agent.ManagementAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "management_agent")
	retryPolicy.ShouldRetryOperation = namedCredentialWorkRequestShouldRetryFunc(timeout)

	response := oci_management_agent.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_management_agent.OperationStatusInProgress),
			string(oci_management_agent.OperationStatusAccepted),
			string(oci_management_agent.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_management_agent.OperationStatusSucceeded),
			string(oci_management_agent.OperationStatusFailed),
			string(oci_management_agent.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_management_agent.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_management_agent.OperationStatusFailed || response.Status == oci_management_agent.OperationStatusCanceled {
		return nil, getErrorFromManagementAgentNamedCredentialWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromManagementAgentNamedCredentialWorkRequest(client *oci_management_agent.ManagementAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_management_agent.ActionTypesEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_management_agent.ListWorkRequestErrorsRequest{
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

func (s *ManagementAgentNamedCredentialResourceCrud) Get() error {
	request := oci_management_agent.GetNamedCredentialRequest{}

	tmp := s.D.Id()
	request.NamedCredentialId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.GetNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.NamedCredential
	return nil
}

func (s *ManagementAgentNamedCredentialResourceCrud) Update() error {
	request := oci_management_agent.UpdateNamedCredentialRequest{}

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

	tmp := s.D.Id()
	request.NamedCredentialId = &tmp

	if properties, ok := s.D.GetOkExists("properties"); ok {
		interfaces := properties.([]interface{})
		tmp := make([]oci_management_agent.NamedCredentialProperty, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "properties", stateDataIndex)
			converted, err := s.mapToNamedCredentialProperty(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("properties") {
			request.Properties = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.UpdateNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getNamedCredentialFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent"), oci_management_agent.ActionTypesUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ManagementAgentNamedCredentialResourceCrud) Delete() error {
	request := oci_management_agent.DeleteNamedCredentialRequest{}

	tmp := s.D.Id()
	request.NamedCredentialId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "management_agent")

	response, err := s.Client.DeleteNamedCredential(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := namedCredentialWaitForWorkRequest(workId, "namedcredential",
		oci_management_agent.ActionTypesDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *ManagementAgentNamedCredentialResourceCrud) SetData() error {
	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.ManagementAgentId != nil {
		s.D.Set("management_agent_id", *s.Res.ManagementAgentId)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	properties := []interface{}{}
	for _, item := range s.Res.Properties {
		properties = append(properties, NamedCredentialPropertyToMap(item))
	}
	s.D.Set("properties", properties)

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

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	return nil
}

func (s *ManagementAgentNamedCredentialResourceCrud) mapToNamedCredentialProperty(fieldKeyFormat string) (oci_management_agent.NamedCredentialProperty, error) {
	result := oci_management_agent.NamedCredentialProperty{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if valueCategory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value_category")); ok {
		result.ValueCategory = oci_management_agent.ValueCategoryTypeEnum(valueCategory.(string))
	}

	return result, nil
}

func NamedCredentialPropertyToMap(obj oci_management_agent.NamedCredentialProperty) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	result["value_category"] = string(obj.ValueCategory)

	return result
}

func NamedCredentialSummaryToMap(obj oci_management_agent.NamedCredentialSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ManagementAgentId != nil {
		result["management_agent_id"] = string(*obj.ManagementAgentId)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	properties := []interface{}{}
	for _, item := range obj.Properties {
		properties = append(properties, NamedCredentialPropertyToMap(item))
	}
	result["properties"] = properties

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

	if obj.Type != nil {
		result["type"] = string(*obj.Type)
	}

	return result
}
