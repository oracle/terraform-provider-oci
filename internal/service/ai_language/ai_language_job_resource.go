// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ai_language

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_ai_language "github.com/oracle/oci-go-sdk/v65/ailanguage"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func AiLanguageJobResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(120 * time.Minute),
			Update: schema.DefaultTimeout(120 * time.Minute),
			Delete: schema.DefaultTimeout(120 * time.Minute),
		},
		Create: createAiLanguageJob,
		Read:   readAiLanguageJob,
		Update: updateAiLanguageJob,
		Delete: deleteAiLanguageJob,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"input_location": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"location_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"OBJECT_STORAGE_FILE_LIST",
								"OBJECT_STORAGE_PREFIX",
							}, true),
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"object_names": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							//	return true
							//},
						},
						//"prefix": {
						//	Type:     schema.TypeString,
						//	Optional: true,
						//	Computed: true,
						//	ForceNew: true,
						//	//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
						//	//	return true
						//	//},
						//},

						// Computed
					},
				},
			},
			"model_metadata_details": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"configuration": {
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
							// Elem must be a simple type (string) for Terraform SDK, cannot be Resource
							Elem: &schema.Schema{
								Type: schema.TypeMap, // nested map for CSV -> config
							},
						},
						"endpoint_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"language_code": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"model_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"model_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"output_location": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"bucket": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"prefix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
							//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
							//	// Ignore OCI-generated prefixes
							//	return strings.HasPrefix(new, "ocid1.")
							//},
						},

						// Computed
					},
				},
			},

			// Optional
			"description": {
				Type: schema.TypeString,
				//DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Optional: true,
				Computed: true,
			},
			"display_name": {
				Type: schema.TypeString,
				//DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				Optional: true,
				Computed: true,
			},
			"input_configuration": {
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
						//"configuration": {
						//	Type:     schema.TypeMap,
						//	Optional: true,
						//	ForceNew: true,
						//	// Elem must be a simple type (string) for Terraform SDK, cannot be Resource
						//	Elem: &schema.Schema{
						//		Type: schema.TypeMap, // nested map for CSV -> config
						//	},
						//},
						//"configuration": {
						//	Type:     schema.TypeList,
						//	Optional: true,
						//	ForceNew: true,
						//	MaxItems: 1,
						//	Elem: &schema.Resource{
						//		Schema: map[string]*schema.Schema{
						//			"csv": {
						//				Type:     schema.TypeList,
						//				Optional: true,
						//				MaxItems: 1,
						//				Elem: &schema.Resource{
						//					Schema: map[string]*schema.Schema{
						//						"config": {
						//							Type:     schema.TypeMap,
						//							Optional: true,
						//							Elem: &schema.Schema{
						//								Type: schema.TypeString,
						//							},
						//						},
						//					},
						//				},
						//			},
						//		},
						//	},
						//},
						"configuration": {
							Type:     schema.TypeMap,
							Optional: true,
							ForceNew: true,
							// Elem must be a simple type (string) for Terraform SDK, cannot be Resource
							Elem: &schema.Schema{
								Type: schema.TypeMap,
								// nested map for CSV -> config
							},
						},
						"document_types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Computed
					},
				},
				//DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				//	// ignore differences since service doesn’t return same config
				//	return true
				//},
			},

			// Computed
			"completed_documents": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failed_documents": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pending_documents": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"percent_complete": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_accepted": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_started": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_documents": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ttl_in_days": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"warnings_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func createAiLanguageJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.CreateResource(d, sync)
}

func readAiLanguageJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.ReadResource(sync)
}

func updateAiLanguageJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteAiLanguageJob(d *schema.ResourceData, m interface{}) error {
	sync := &AiLanguageJobResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AiServiceLanguageClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AiLanguageJobResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_ai_language.AIServiceLanguageClient
	Res                    *oci_ai_language.Job
	DisableNotFoundRetries bool
}

func (s *AiLanguageJobResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AiLanguageJobResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_ai_language.JobLifecycleStateInProgress),
	}
}

func (s *AiLanguageJobResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_ai_language.JobLifecycleStateSucceeded),
	}
}

func (s *AiLanguageJobResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_ai_language.JobLifecycleStateDeleting),
	}
}

func (s *AiLanguageJobResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_ai_language.JobLifecycleStateDeleted),
	}
}

func (s *AiLanguageJobResourceCrud) Create() error {
	ctx := context.Background()
	request := oci_ai_language.CreateJobRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if inputConfiguration, ok := s.D.GetOkExists("input_configuration"); ok {
		if tmpList := inputConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_configuration", 0)
			tmp, err := s.mapToInputConfiguration(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InputConfiguration = &tmp
		}
	}

	if inputLocation, ok := s.D.GetOkExists("input_location"); ok {
		if tmpList := inputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "input_location", 0)
			tmp, err := s.mapToInputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.InputLocation = tmp
		}
	}

	if modelMetadataDetails, ok := s.D.GetOkExists("model_metadata_details"); ok {
		interfaces := modelMetadataDetails.([]interface{})
		tmp := make([]oci_ai_language.ModelMetadataDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "model_metadata_details", stateDataIndex)
			converted, err := s.mapToModelMetadataDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("model_metadata_details") {
			request.ModelMetadataDetails = tmp
		}
	}

	if outputLocation, ok := s.D.GetOkExists("output_location"); ok {
		if tmpList := outputLocation.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "output_location", 0)
			tmp, err := s.mapToObjectPrefixOutputLocation(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.OutputLocation = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	if request.InputConfiguration != nil {
		data, _ := json.MarshalIndent(request.InputConfiguration, "", "  ")
		fmt.Println("DEBUG InputConfiguration =>", string(data))
	} else {
		fmt.Println("DEBUG InputConfiguration => nil")
	}
	response, err := s.Client.CreateJob(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	// ---- Wait until job finishes ----
	utils.Logf("[Info] Dump createJob response id: %s,     workReqId: %s ", identifier, workId)
	s.waitForJobCompletion(context.Background(), identifier, s.D.Timeout(schema.TimeoutCreate))
	utils.Logf("[Info] Dump createJob response id after wait: %s,     workReqId: %s ", identifier, workId)

	s.D.SetId(*identifier)
	return s.Get()
	//return s.getJobFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language"), oci_ai_language.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AiLanguageJobResourceCrud) waitForJobCompletion(ctx context.Context, jobId *string, timeout time.Duration) {
	if jobId == nil {
		utils.Logf("[Error] jobid is nil, cannot wait for job completion")
		return
	}

	finishedStates := map[oci_ai_language.JobLifecycleStateEnum]bool{
		oci_ai_language.JobLifecycleStateSucceeded: true,
		oci_ai_language.JobLifecycleStateFailed:    true,
		oci_ai_language.JobLifecycleStateCanceled:  true,
	}

	start := time.Now()
	for {
		resp, err := s.Client.GetJob(ctx, oci_ai_language.GetJobRequest{JobId: jobId})
		if err != nil {
			utils.Logf("[Error] error fetching work request: %v", err)
			return
		}

		state := resp.LifecycleState
		utils.Logf("[DEBUG] Job %s current state: %s", *jobId, state)

		if finishedStates[state] {
			if state == oci_ai_language.JobLifecycleStateFailed {
				utils.Logf("[Error] job %s failed: %v", *jobId, resp.LifecycleDetails)
			}
			utils.Logf("[INFO] Job %s completed successfully", *jobId)
			return
		}

		if time.Since(start) > timeout {
			utils.Logf("[Error] timeout waiting for job %s to complete", *jobId)
			return
		}

		time.Sleep(30 * time.Second) // polling interval
	}
}

func (s *AiLanguageJobResourceCrud) getJobFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_ai_language.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	jobId, err := jobWaitForWorkRequest(workId, "job",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*jobId)

	return s.Get()
}

func jobWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "ai_language", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_ai_language.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func jobWaitForWorkRequest(wId *string, entityType string, action oci_ai_language.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_ai_language.AIServiceLanguageClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "ai_language")
	retryPolicy.ShouldRetryOperation = jobWorkRequestShouldRetryFunc(timeout)

	response := oci_ai_language.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_ai_language.ActionTypeInProgress),
			string(oci_ai_language.JobLifecycleStateAccepted),
		},
		Target: []string{
			string("ACTIVE"),
			string("SUCCEEDED"),
			string("FAILED"),
			string(oci_ai_language.JobLifecycleStateCanceled),
			string(oci_ai_language.ActionTypeDeleted),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_ai_language.GetWorkRequestRequest{
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
	//if identifier == nil || response.Status == oci_ai_language.OperationStatusFailed || response.Status == oci_ai_language.OperationStatusCanceled {
	//	return nil, getErrorFromAiLanguageJobWorkRequest(client, wId, retryPolicy, entityType, action)
	//}

	return identifier, nil
}

func getErrorFromAiLanguageJobWorkRequest(client *oci_ai_language.AIServiceLanguageClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_ai_language.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_ai_language.ListWorkRequestErrorsRequest{
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

func (s *AiLanguageJobResourceCrud) Get() error {
	request := oci_ai_language.GetJobRequest{}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.GetJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *AiLanguageJobResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_ai_language.UpdateJobRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.JobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	response, err := s.Client.UpdateJob(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Job
	return nil
}

func (s *AiLanguageJobResourceCrud) Delete() error {
	request := oci_ai_language.DeleteJobRequest{}
	getRequest := oci_ai_language.GetJobRequest{}

	finishedLifecycleStates := map[string]bool{
		string(oci_ai_language.JobLifecycleStateSucceeded): true,
		string(oci_ai_language.JobLifecycleStateFailed):    true,
	}

	tmp := s.D.Id()
	request.JobId = &tmp
	getRequest.JobId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")
	//request.RequestMetadata.RetryPolicy.MaximumNumberAttempts = 2
	for {
		// Re-fetch job state each iteration
		getResponse, err := s.Client.GetJob(context.Background(), getRequest)
		if err != nil {
			if se, ok := oci_common.IsServiceError(err); ok {
				if se.GetHTTPStatusCode() == 404 {
					// Job not found → consider deleted
					return nil
				}
			}
			// handle error (e.g., log or break)
			fmt.Println("Error fetching job:", err)
			break
		}

		// Check if job is in finished states
		if _, ok := finishedLifecycleStates[string(getResponse.LifecycleState)]; ok {
			break
		}

		// Sleep to avoid busy loop
		time.Sleep(5 * time.Second)
	}

	response, err := s.Client.DeleteJob(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	if workId == nil {
		fmt.Println("[Warn] WorkRequestID is nil; retrying GetJob to fetch latest work request ID")

		// try to re-fetch job details to get WorkRequestId (some services set it after delete call)
		getResponse, err := s.Client.GetJob(context.Background(), getRequest)
		if err == nil && getResponse.OpcRequestId != nil {
			workId = getResponse.OpcRequestId
			fmt.Println("[Info] Retrieved WorkRequestID from GetJob:", *workId)
		}
	}
	// Fallback: if still nil, handle deletion manually
	if workId == nil {
		fmt.Println("[Warn] WorkRequestID still nil; waiting until job is deleted manually")

		timeout := time.After(s.D.Timeout(schema.TimeoutDelete))
		for {
			select {
			case <-timeout:
				return fmt.Errorf("timeout waiting for job deletion (no WorkRequestID available)")
			default:
				_, err := s.Client.GetJob(context.Background(), getRequest)
				if err != nil {
					if se, ok := oci_common.IsServiceError(err); ok {
						if se.GetHTTPStatusCode() == 404 {
							// Job not found → consider deleted
							fmt.Println("[Info] Job deleted successfully (no work request)")
							return nil
						}
					}
					fmt.Println("Error polling job:", err)
				}
				time.Sleep(10 * time.Second)
			}
		}
	}
	fmt.Println("WorkRequestID error while deleting:", workId)
	// Wait until it finishes
	//_, delWorkRequestErr := jobWaitForWorkRequest(workId, "job",
	//	oci_ai_language.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	//return delWorkRequestErr
	return nil
}

func (s *AiLanguageJobResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompletedDocuments != nil {
		s.D.Set("completed_documents", *s.Res.CompletedDocuments)
	}

	if s.Res.CreatedBy != nil {
		s.D.Set("created_by", *s.Res.CreatedBy)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.FailedDocuments != nil {
		s.D.Set("failed_documents", *s.Res.FailedDocuments)
	}

	if s.Res.InputConfiguration != nil {
		s.D.Set("input_configuration", []interface{}{InputConfigurationToMap(s.Res.InputConfiguration)})
	} else {
		s.D.Set("input_configuration", nil)
	}

	if s.Res.InputLocation != nil {
		inputLocationArray := []interface{}{}
		if inputLocationMap := InputLocationToMap(&s.Res.InputLocation); inputLocationMap != nil {
			inputLocationArray = append(inputLocationArray, inputLocationMap)
		}
		s.D.Set("input_location", inputLocationArray)
	} else {
		s.D.Set("input_location", nil)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	modelMetadataDetails := []interface{}{}
	for _, item := range s.Res.ModelMetadataDetails {
		modelMetadataDetails = append(modelMetadataDetails, ModelMetadataDetailsToMap(item))
	}
	s.D.Set("model_metadata_details", modelMetadataDetails)

	if s.Res.OutputLocation != nil {
		s.D.Set("output_location", []interface{}{ObjectPrefixOutputLocationToMap(s.Res.OutputLocation)})
	} else {
		s.D.Set("output_location", nil)
	}

	if s.Res.PendingDocuments != nil {
		s.D.Set("pending_documents", *s.Res.PendingDocuments)
	}

	if s.Res.PercentComplete != nil {
		s.D.Set("percent_complete", *s.Res.PercentComplete)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeAccepted != nil {
		s.D.Set("time_accepted", s.Res.TimeAccepted.String())
	}

	if s.Res.TimeCompleted != nil {
		s.D.Set("time_completed", s.Res.TimeCompleted.String())
	}

	if s.Res.TimeStarted != nil {
		s.D.Set("time_started", s.Res.TimeStarted.String())
	}

	if s.Res.TotalDocuments != nil {
		s.D.Set("total_documents", *s.Res.TotalDocuments)
	}

	if s.Res.TtlInDays != nil {
		s.D.Set("ttl_in_days", *s.Res.TtlInDays)
	}

	if s.Res.WarningsCount != nil {
		s.D.Set("warnings_count", *s.Res.WarningsCount)
	}

	return nil
}

func (s *AiLanguageJobResourceCrud) mapToInputConfiguration(fieldKeyFormat string) (oci_ai_language.InputConfiguration, error) {
	result := oci_ai_language.InputConfiguration{}

	// Handle configuration block
	if configList, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
		if len(configList.([]interface{})) > 0 {
			configMap := configList.([]interface{})[0].(map[string]interface{})
			docsConfigMap := make(map[string]oci_ai_language.DocumentsConfiguration)

			// Handle "csv" block
			if csvList, ok := configMap["csv"].([]interface{}); ok && len(csvList) > 0 {
				csvMap := csvList[0].(map[string]interface{})
				if cfg, ok := csvMap["config"].(map[string]interface{}); ok {
					docCfg := oci_ai_language.DocumentsConfiguration{
						Config: map[string]string{},
					}
					for k, v := range cfg {
						if v != nil {
							docCfg.Config[k] = fmt.Sprintf("%v", v)
						}
					}
					docsConfigMap["CSV"] = docCfg
				}
			}
			result.Configuration = docsConfigMap
		}
	}

	// Handle document_types
	if documentTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "document_types")); ok {
		interfaces := documentTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) > 0 {
			result.DocumentTypes = tmp
		}
	}

	return result, nil
}

//func (s *AiLanguageJobResourceCrud) mapToInputConfiguration(fieldKeyFormat string) (oci_ai_language.InputConfiguration, error) {
//	result := oci_ai_language.InputConfiguration{}
//
//	if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
//		//result.Configuration = tfresource.ObjectMapToStringMap(configuration.(map[string]interface{}))
//		stringMap := tfresource.ObjectMapToStringMap(configuration.(map[string]interface{}))
//
//		docsConfigMap := make(map[string]oci_ai_language.DocumentsConfiguration)
//
//		for k, v := range stringMap {
//			docsConfigMap[k] = oci_ai_language.DocumentsConfiguration{
//				Config: map[string]string{
//					"value": v,
//				},
//			}
//		}
//
//		result.Configuration = docsConfigMap
//	}
//
//	if documentTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "document_types")); ok {
//		interfaces := documentTypes.([]interface{})
//		tmp := make([]string, len(interfaces))
//		for i := range interfaces {
//			if interfaces[i] != nil {
//				tmp[i] = interfaces[i].(string)
//			}
//		}
//		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "document_types")) {
//			result.DocumentTypes = tmp
//		}
//	}
//
//	return result, nil
//}

func InputConfigurationToMap(obj *oci_ai_language.InputConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	result["configuration"] = obj.Configuration

	result["document_types"] = obj.DocumentTypes

	return result
}

func (s *AiLanguageJobResourceCrud) mapToInputLocation(fieldKeyFormat string) (oci_ai_language.InputLocation, error) {
	var baseObject oci_ai_language.InputLocation
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OBJECT_STORAGE_FILE_LIST"):
		details := oci_ai_language.ObjectStorageFileNameLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		if objectNames, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_names")); ok {
			interfaces := objectNames.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_names")) {
				details.ObjectNames = tmp
			}
		}
		baseObject = details
	case strings.ToLower("OBJECT_STORAGE_PREFIX"):
		details := oci_ai_language.ObjectStoragePrefixLocation{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.BucketName = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.NamespaceName = &tmp
		}
		//if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
		//	tmp := prefix.(string)
		//	details.Prefix = &tmp
		//}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	fmt.Println("DEBUG InputConfiguration =>", baseObject)
	return baseObject, nil
}

func InputLocationToMap(obj *oci_ai_language.InputLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_ai_language.ObjectStorageFileNameLocation:
		result["location_type"] = "OBJECT_STORAGE_FILE_LIST"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		result["object_names"] = v.ObjectNames
	case oci_ai_language.ObjectStoragePrefixLocation:
		result["location_type"] = "OBJECT_STORAGE_PREFIX"

		if v.BucketName != nil {
			result["bucket"] = string(*v.BucketName)
		}

		if v.NamespaceName != nil {
			result["namespace"] = string(*v.NamespaceName)
		}

		//if v.Prefix != nil {
		//	result["prefix"] = string(*v.Prefix)
		//}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func JobSummaryToMap(obj oci_ai_language.JobSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.CompletedDocuments != nil {
		result["completed_documents"] = int(*obj.CompletedDocuments)
	}

	if obj.CreatedBy != nil {
		result["created_by"] = string(*obj.CreatedBy)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.FailedDocuments != nil {
		result["failed_documents"] = int(*obj.FailedDocuments)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	if obj.PendingDocuments != nil {
		result["pending_documents"] = int(*obj.PendingDocuments)
	}

	if obj.PercentComplete != nil {
		result["percent_complete"] = int(*obj.PercentComplete)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.TimeAccepted != nil {
		result["time_accepted"] = obj.TimeAccepted.String()
	}

	if obj.TimeCompleted != nil {
		result["time_completed"] = obj.TimeCompleted.String()
	}

	if obj.TimeStarted != nil {
		result["time_started"] = obj.TimeStarted.String()
	}

	if obj.TotalDocuments != nil {
		result["total_documents"] = int(*obj.TotalDocuments)
	}

	if obj.WarningsCount != nil {
		result["warnings_count"] = int(*obj.WarningsCount)
	}

	return result
}

func StringMapToConfigDetails(in map[string]string) map[string]oci_ai_language.ConfigurationDetails {
	out := make(map[string]oci_ai_language.ConfigurationDetails, len(in))
	for k, v := range in {
		out[k] = oci_ai_language.ConfigurationDetails{
			ConfigurationMap: map[string]string{
				"value": v,
			},
		}
	}
	return out
}

func (s *AiLanguageJobResourceCrud) mapToModelMetadataDetails(fieldKeyFormat string) (oci_ai_language.ModelMetadataDetails, error) {
	result := oci_ai_language.ModelMetadataDetails{}

	if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
		//result.Configuration = tfresource.ObjectMapToStringMap(configuration.(map[string]interface{}))
		stringMap := tfresource.ObjectMapToStringMap(configuration.(map[string]interface{}))
		result.Configuration = StringMapToConfigDetails(stringMap)
	}

	if endpointId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_id")); ok {
		tmp := endpointId.(string)
		result.EndpointId = &tmp
	}

	if languageCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "language_code")); ok {
		tmp := languageCode.(string)
		result.LanguageCode = &tmp
	}

	if modelId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_id")); ok {
		tmp := modelId.(string)
		result.ModelId = &tmp
	}

	if modelType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "model_type")); ok {
		tmp := modelType.(string)
		result.ModelType = &tmp
	}

	return result, nil
}

func ModelMetadataDetailsToMap(obj oci_ai_language.ModelMetadataDetails) map[string]interface{} {
	result := map[string]interface{}{}

	result["configuration"] = obj.Configuration

	if obj.EndpointId != nil {
		result["endpoint_id"] = string(*obj.EndpointId)
	}

	if obj.LanguageCode != nil {
		result["language_code"] = string(*obj.LanguageCode)
	}

	if obj.ModelId != nil {
		result["model_id"] = string(*obj.ModelId)
	}

	if obj.ModelType != nil {
		result["model_type"] = string(*obj.ModelType)
	}

	return result
}

func (s *AiLanguageJobResourceCrud) mapToObjectPrefixOutputLocation(fieldKeyFormat string) (oci_ai_language.ObjectPrefixOutputLocation, error) {
	result := oci_ai_language.ObjectPrefixOutputLocation{}

	if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
		tmp := bucket.(string)
		result.BucketName = &tmp
	}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.NamespaceName = &tmp
	}

	if prefix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "prefix")); ok {
		tmp := prefix.(string)
		result.Prefix = &tmp
	}

	return result, nil
}

func ObjectPrefixOutputLocationToMap(obj *oci_ai_language.ObjectPrefixOutputLocation) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.BucketName != nil {
		result["bucket"] = string(*obj.BucketName)
	}

	if obj.NamespaceName != nil {
		result["namespace"] = string(*obj.NamespaceName)
	}

	if obj.Prefix != nil {
		result["prefix"] = string(*obj.Prefix)
	}

	return result
}

func (s *AiLanguageJobResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_ai_language.ChangeJobCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.JobId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "ai_language")

	_, err := s.Client.ChangeJobCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
