// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package generative_ai_agent

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_generative_ai_agent "github.com/oracle/oci-go-sdk/v65/generativeaiagent"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GenerativeAiAgentDataIngestionJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("40m"),
			Update: tfresource.GetTimeoutDuration("20m"),
			Delete: tfresource.GetTimeoutDuration("20m"),
		},
		Create: createGenerativeAiAgentDataIngestionJob,
		Read:   readGenerativeAiAgentDataIngestionJob,
		Delete: deleteGenerativeAiAgentDataIngestionJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"data_source_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"data_ingestion_job_statistics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"duration_in_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_failed_files": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_of_ingested_files": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
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

func createGenerativeAiAgentDataIngestionJob(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.CreateResource(d, sync)
}

func readGenerativeAiAgentDataIngestionJob(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()

	return tfresource.ReadResource(sync)
}

func deleteGenerativeAiAgentDataIngestionJob(d *schema.ResourceData, m interface{}) error {
	sync := &GenerativeAiAgentDataIngestionJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GenerativeAiAgentClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type GenerativeAiAgentDataIngestionJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_generative_ai_agent.GenerativeAiAgentClient
	Res                    *oci_generative_ai_agent.DataIngestionJob
	DisableNotFoundRetries bool
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_generative_ai_agent.DataIngestionJobLifecycleStateInProgress),
	}
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.DataIngestionJobLifecycleStateSucceeded),
	}
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_generative_ai_agent.DataIngestionJobLifecycleStateDeleting),
	}
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_generative_ai_agent.DataIngestionJobLifecycleStateDeleted),
	}
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) Create() error {
	request := oci_generative_ai_agent.CreateDataIngestionJobRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dataSourceId, ok := s.D.GetOkExists("data_source_id"); ok {
		tmp := dataSourceId.(string)
		request.DataSourceId = &tmp
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

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.CreateDataIngestionJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getDataIngestionJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent"), oci_generative_ai_agent.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) getDataIngestionJobFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_generative_ai_agent.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	dataIngestionJobId, err := dataIngestionJobWaitForWorkRequest(workId, "dataingestionjob",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, dataIngestionJobId)
		_, cancelErr := s.Client.CancelWorkRequest(context.Background(),
			oci_generative_ai_agent.CancelWorkRequestRequest{
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
	s.D.SetId(*dataIngestionJobId)

	return s.Get()
}

func dataIngestionJobWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "generative_ai_agent", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_generative_ai_agent.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func dataIngestionJobWaitForWorkRequest(wId *string, entityType string, action oci_generative_ai_agent.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_generative_ai_agent.GenerativeAiAgentClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "generative_ai_agent")
	retryPolicy.ShouldRetryOperation = dataIngestionJobWorkRequestShouldRetryFunc(timeout)

	response := oci_generative_ai_agent.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_generative_ai_agent.OperationStatusInProgress),
			string(oci_generative_ai_agent.OperationStatusAccepted),
			string(oci_generative_ai_agent.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_generative_ai_agent.OperationStatusSucceeded),
			string(oci_generative_ai_agent.OperationStatusFailed),
			string(oci_generative_ai_agent.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_generative_ai_agent.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_generative_ai_agent.OperationStatusFailed || response.Status == oci_generative_ai_agent.OperationStatusCanceled {
		return nil, getErrorFromGenerativeAiAgentDataIngestionJobWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGenerativeAiAgentDataIngestionJobWorkRequest(client *oci_generative_ai_agent.GenerativeAiAgentClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_generative_ai_agent.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_generative_ai_agent.ListWorkRequestErrorsRequest{
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

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) Get() error {
	request := oci_generative_ai_agent.GetDataIngestionJobRequest{}

	tmp := s.D.Id()
	request.DataIngestionJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.GetDataIngestionJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DataIngestionJob
	return nil
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) Delete() error {
	request := oci_generative_ai_agent.DeleteDataIngestionJobRequest{}

	tmp := s.D.Id()
	request.DataIngestionJobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "generative_ai_agent")

	response, err := s.Client.DeleteDataIngestionJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := dataIngestionJobWaitForWorkRequest(workId, "dataingestionjob",
		oci_generative_ai_agent.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GenerativeAiAgentDataIngestionJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataIngestionJobStatistics != nil {
		s.D.Set("data_ingestion_job_statistics", []interface{}{DataIngestionJobStatisticsToMap(s.Res.DataIngestionJobStatistics)})
	} else {
		s.D.Set("data_ingestion_job_statistics", nil)
	}

	if s.Res.DataSourceId != nil {
		s.D.Set("data_source_id", *s.Res.DataSourceId)
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

func DataIngestionJobStatisticsToMap(obj *oci_generative_ai_agent.DataIngestionJobStatistics) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DurationInSeconds != nil {
		result["duration_in_seconds"] = int(*obj.DurationInSeconds)
	}

	if obj.NumberOfFailedFiles != nil {
		result["number_of_failed_files"] = int(*obj.NumberOfFailedFiles)
	}

	if obj.NumberOfIngestedFiles != nil {
		result["number_of_ingested_files"] = int(*obj.NumberOfIngestedFiles)
	}

	return result
}

func DataIngestionJobSummaryToMap(obj oci_generative_ai_agent.DataIngestionJobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DataSourceId != nil {
		result["data_source_id"] = string(*obj.DataSourceId)
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

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
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
