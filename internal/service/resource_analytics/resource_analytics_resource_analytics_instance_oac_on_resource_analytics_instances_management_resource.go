// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package resource_analytics

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_resource_analytics "github.com/oracle/oci-go-sdk/v65/resourceanalytics"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ResourceAnalyticsResourceAnalyticsInstanceOacManagementResource() *schema.Resource {
	return &schema.Resource{
		Timeouts: &schema.ResourceTimeout{
			Create: &tfresource.TwoHours,
			Update: &tfresource.TwoHours,
			Delete: &tfresource.TwoHours,
			Read:   &tfresource.TwoHours,
		},
		Create: createResourceAnalyticsResourceAnalyticsInstanceOacManagement,
		Read:   readResourceAnalyticsResourceAnalyticsInstanceOacManagement,
		Update: updateResourceAnalyticsResourceAnalyticsInstanceOacManagement,
		Delete: deleteResourceAnalyticsResourceAnalyticsInstanceOacManagement,
		Schema: map[string]*schema.Schema{
			// Required
			"resource_analytics_instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enable_oac": {
				Type:     schema.TypeBool,
				Required: true,
			},

			// Optional
			"attachment_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						//"enable_oac": {
						//	Type:     schema.TypeBool,
						//	Required: true,
						//},

						// Optional
						"idcs_domain_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"license_model": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
						"network_details": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							ForceNew: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									//"enable_oac": {
									//	Type:     schema.TypeBool,
									//	Required: true,
									//},

									// Optional
									"nsg_ids": {
										Type:     schema.TypeSet,
										Optional: true,
										Computed: true,
										ForceNew: true,
										Set:      tfresource.LiteralTypeHashCodeForSets,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										ForceNew: true,
									},

									// Computed
								},
							},
						},
						"nsg_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							ForceNew: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
					},
				},
			},
			"attachment_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createResourceAnalyticsResourceAnalyticsInstanceOacManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.ResourceAnalyticsInstanceClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.Res = &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse{}
	sync.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusSucceeded
	return tfresource.CreateResource(d, sync)
}

func readResourceAnalyticsResourceAnalyticsInstanceOacManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.ResourceAnalyticsInstanceClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.Res = &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse{}

	return tfresource.ReadResource(sync)
}

func updateResourceAnalyticsResourceAnalyticsInstanceOacManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.ResourceAnalyticsInstanceClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.Res = &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse{}

	return tfresource.UpdateResource(d, sync)
}

func deleteResourceAnalyticsResourceAnalyticsInstanceOacManagement(d *schema.ResourceData, m interface{}) error {
	sync := &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud{}
	sync.D = d
	clients := m.(*client.OracleClients)
	sync.Client = clients.ResourceAnalyticsInstanceClient()
	sync.WorkReqClient = clients.ResourceAnalyticsInstanceClient()
	sync.Res = &ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse{}
	sync.DisableNotFoundRetries = true

	if sync.Res.LifecycleState == "" {
		sync.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusSucceeded
	}

	return tfresource.DeleteResource(d, sync)
}

type ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse struct {
	enableResponse  *oci_resource_analytics.ResourceAnalyticsInstanceEnableOacResponse
	disableResponse *oci_resource_analytics.ResourceAnalyticsInstanceDisableOacResponse
	LifecycleState  oci_resource_analytics.ListWorkRequestsStatusEnum `mandatory:"true" json:"lifecycleState"`
}

type ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_resource_analytics.ResourceAnalyticsInstanceClient
	WorkReqClient          *oci_resource_analytics.ResourceAnalyticsInstanceClient
	Res                    *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResponse
	DisableNotFoundRetries bool
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) ID() string {
	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		return resourceAnalyticsInstanceId.(string)
	}
	return s.D.Id()
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_resource_analytics.OperationStatusInProgress),
		string(oci_resource_analytics.OperationStatusAccepted),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_resource_analytics.OperationStatusSucceeded),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_resource_analytics.OperationStatusInProgress),
		string(oci_resource_analytics.OperationStatusAccepted),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_resource_analytics.OperationStatusSucceeded),
	}
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) State() string {
	log.Printf("[INFO] State() Before fetch = %v\n", s.Res.LifecycleState)
	var err error
	retryPolicy := tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	if s.Res.disableResponse != nil {
		log.Printf("[INFO] Fetching state from disableResponse: %v\n", s.Res.disableResponse)
		s.setResToFetchedWorkRequest(err, s.Res.disableResponse.OpcWorkRequestId, retryPolicy)
	} else if s.Res.enableResponse != nil {
		log.Printf("[INFO] Fetching state from enableResponse: %v\n", s.Res.enableResponse)
		s.setResToFetchedWorkRequest(err, s.Res.enableResponse.OpcWorkRequestId, retryPolicy)
	}
	log.Printf("[INFO] State() After fetch = %v\n", s.Res.LifecycleState)
	return string(s.Res.LifecycleState)
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) Get() error {
	// For management resources, we don't need to fetch actual data
	// The resource state is managed through the enable/disable operations
	// This method is required by the interface but can be a no-op
	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) Create() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_oac"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		request := oci_resource_analytics.ResourceAnalyticsInstanceEnableOacRequest{}

		if attachmentDetails, ok := s.D.GetOkExists("attachment_details"); ok {
			if tmpList := attachmentDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attachment_details", 0)
				tmp, err := s.mapToResourceAnalyticsInstanceOacAttachmentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.AttachmentDetails = &tmp
			}
		}

		if attachmentType, ok := s.D.GetOkExists("attachment_type"); ok {
			request.AttachmentType = oci_resource_analytics.ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum(attachmentType.(string))
		}

		if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
			tmp := resourceAnalyticsInstanceId.(string)
			request.ResourceAnalyticsInstanceId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

		response, err := s.executeEnableOacRequest(request)
		if err != nil {
			return err
		}

		s.setResToFetchedWorkRequest(err, response.OpcWorkRequestId, request.RequestMetadata.RetryPolicy)

		err = s.getResourceAnalyticsInstanceOacManagementFromWorkRequest(response.OpcWorkRequestId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
		return nil
	}

	// Disable OAC req
	request := oci_resource_analytics.ResourceAnalyticsInstanceDisableOacRequest{}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.executeDisableOacRequest(request)
	if err != nil {
		return err
	}

	s.setResToFetchedWorkRequest(err, response.OpcWorkRequestId, request.RequestMetadata.RetryPolicy)

	workId := response.OpcWorkRequestId
	err = s.getResourceAnalyticsInstanceOacManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	//s.Res.disableResponse = &response
	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) setResToFetchedWorkRequest(err error, workId *string, retryPolicy *oci_common.RetryPolicy) *oci_resource_analytics.WorkRequest {
	log.Printf("setResToFetchedWorkRequest called with workId: %s, err: %v\n", *workId, err)
	wrResponse, err := s.WorkReqClient.GetWorkRequest(context.Background(),
		oci_resource_analytics.GetWorkRequestRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	log.Printf("setResToFetchedWorkRequest: wrResponse: %v\n", wrResponse)
	wr := &wrResponse.WorkRequest

	s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusEnum(wr.Status)
	log.Printf("setResToFetchedWorkRequest: setting state to : %v\n", s.Res.LifecycleState)
	return wr
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) getResourceAnalyticsInstanceOacManagementFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_resource_analytics.ActionTypeEnum, timeout time.Duration) error {

	log.Printf("getResourceAnalyticsInstanceOacManagementFromWorkRequest called with workId: %s, actionTypeEnum: %v, timeout: %v\n", *workId, actionTypeEnum, timeout)

	// Wait until it finishes
	resourceAnalyticsInstanceOacManagementId, err := resourceAnalyticsInstanceOacManagementWaitForWorkRequest(s, workId, "resourceanalyticsinstance",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	log.Printf("getResourceAnalyticsInstanceOacManagementFromWorkRequest. Finished!\n")

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] getResourceAnalyticsInstanceOacManagementFromWorkRequest: creation failed, attempting to cancel the workrequest: %s for identifier: %v\n", *workId, resourceAnalyticsInstanceOacManagementId)
		_, cancelErr := s.WorkReqClient.CancelWorkRequest(context.Background(),
			oci_resource_analytics.CancelWorkRequestRequest{
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

	return nil
}

func resourceAnalyticsInstanceOacManagementWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "resource_analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_resource_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func resourceAnalyticsInstanceOacManagementWaitForWorkRequest(s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud, wId *string, entityType string, action oci_resource_analytics.ActionTypeEnum, timeout time.Duration, disableFoundRetries bool, client *oci_resource_analytics.ResourceAnalyticsInstanceClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "resource_analytics")
	retryPolicy.ShouldRetryOperation = resourceAnalyticsInstanceOacManagementWorkRequestShouldRetryFunc(timeout)

	response := oci_resource_analytics.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_resource_analytics.OperationStatusInProgress),
			string(oci_resource_analytics.OperationStatusAccepted),
			string(oci_resource_analytics.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_resource_analytics.OperationStatusSucceeded),
			string(oci_resource_analytics.OperationStatusFailed),
			string(oci_resource_analytics.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			log.Printf("[DEBUG] REFRESH() resourceAnalyticsInstanceOacManagementWaitForWorkRequest. Getting WR...")
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_resource_analytics.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			log.Printf("[DEBUG] resourceAnalyticsInstanceOacManagementWaitForWorkRequest | Got Work Request.. : wr.Status: %v\nSetting s.Res.LifecycleState!\n", wr.Status)
			s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusEnum(wr.Status)
			log.Printf("[DEBUG] Res.LifecycleState = %s\n", s.Res.LifecycleState)
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	log.Printf("[DEBUG] ABOUT TO WAIT FOR STATE... resourceAnalyticsInstanceOacManagementWaitForWorkRequest: stateConf: %v\n", stateConf)
	if _, e := stateConf.WaitForState(); e != nil {
		log.Printf("[DEBUG] ABOUT TO WAIT FOR STATE... DONE! Bad :( resourceAnalyticsInstanceOacManagementWaitForWorkRequest: stateConf: %v\n", stateConf)
		return nil, e
	}
	log.Printf("[DEBUG] ABOUT TO WAIT FOR STATE... DONE! GOOD?! resourceAnalyticsInstanceOacManagementWaitForWorkRequest: stateConf: %v\n", stateConf)

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		log.Printf("[DEBUG] resourceAnalyticsInstanceOacManagementWaitForWorkRequest: res.entityType: %v = entityType: %v, action: %v = res.ActionType: %v, res: %v\n", strings.ToLower(*res.EntityType), entityType, action, res.ActionType, res)
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				log.Printf("[DEBUG] resourceAnalyticsInstanceOacManagementWaitForWorkRequest: MATCH!")
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_resource_analytics.OperationStatusFailed || response.Status == oci_resource_analytics.OperationStatusCanceled {
		log.Printf("[DEBUG] resourceAnalyticsInstanceOacManagementWaitForWorkRequest: DIDNT FIND!!!!")
		return nil, getErrorFromResourceAnalyticsResourceAnalyticsInstanceOacManagementWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromResourceAnalyticsResourceAnalyticsInstanceOacManagementWorkRequest(client *oci_resource_analytics.ResourceAnalyticsInstanceClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_resource_analytics.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_resource_analytics.ListWorkRequestErrorsRequest{
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

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) Update() error {
	log.Printf("[DEBUG] Update called for resourceAnalyticsInstanceOacManagementResourceCrud")
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_oac"); ok {
		operation = enableOperation.(bool)
	}

	if operation {
		log.Printf("[DEBUG] ENABLE Update called for resourceAnalyticsInstanceOacManagementResourceCrud")
		request := oci_resource_analytics.ResourceAnalyticsInstanceEnableOacRequest{}

		if attachmentDetails, ok := s.D.GetOkExists("attachment_details"); ok {
			if tmpList := attachmentDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attachment_details", 0)
				tmp, err := s.mapToResourceAnalyticsInstanceOacAttachmentDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				request.AttachmentDetails = &tmp
			}
		}

		if attachmentType, ok := s.D.GetOkExists("attachment_type"); ok {
			request.AttachmentType = oci_resource_analytics.ResourceAnalyticsInstanceEnableOacDetailsAttachmentTypeEnum(attachmentType.(string))
		}

		if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
			tmp := resourceAnalyticsInstanceId.(string)
			request.ResourceAnalyticsInstanceId = &tmp
		}

		request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

		response, err := s.executeEnableOacRequest(request)
		if err != nil {
			return err
		}

		workId := response.OpcWorkRequestId
		err = s.getResourceAnalyticsInstanceOacManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}

		s.setResToFetchedWorkRequest(err, response.OpcWorkRequestId, request.RequestMetadata.RetryPolicy)

		return nil
	}

	log.Printf("[DEBUG] DISABLE Update called for resourceAnalyticsInstanceOacManagementResourceCrud")
	request := oci_resource_analytics.ResourceAnalyticsInstanceDisableOacRequest{}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	// Set a small JSON payload to ensure Content-Length header is properly set
	// requestBody := "{}"
	// request.RequestBody = &requestBody

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.executeDisableOacRequest(request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getResourceAnalyticsInstanceOacManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}
	s.setResToFetchedWorkRequest(err, response.OpcWorkRequestId, request.RequestMetadata.RetryPolicy)

	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) executeEnableOacRequest(request oci_resource_analytics.ResourceAnalyticsInstanceEnableOacRequest) (oci_resource_analytics.ResourceAnalyticsInstanceEnableOacResponse, error) {

	s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusInProgress
	log.Printf("[DEBUG] executeEnableOacRequest: pre - SETTING LifeCycleState to: %v\n", s.Res.LifecycleState)
	response, err := s.Client.ResourceAnalyticsInstanceEnableOac(context.Background(), request)
	s.Res.enableResponse = &response
	s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusSucceeded
	log.Printf("[DEBUG] executeEnableOacRequest: post - SETTING LifeCycleState to: %v\n", s.Res.LifecycleState)
	return response, err
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) executeDisableOacRequest(request oci_resource_analytics.ResourceAnalyticsInstanceDisableOacRequest) (oci_resource_analytics.ResourceAnalyticsInstanceDisableOacResponse, error) {
	s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusInProgress
	log.Printf("[DEBUG] executeDisableOacRequest: pre - SETTING LifeCycleState to: %v\n", s.Res.LifecycleState)
	response, err := s.Client.ResourceAnalyticsInstanceDisableOac(context.Background(), request)
	s.Res.disableResponse = &response
	s.Res.LifecycleState = oci_resource_analytics.ListWorkRequestsStatusSucceeded
	log.Printf("[DEBUG] executeDisableOacRequest: post - SETTING LifeCycleState to: %v\n", s.Res.LifecycleState)
	return response, err
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) Delete() error {
	var operation bool
	if enableOperation, ok := s.D.GetOkExists("enable_oac"); ok {
		operation = enableOperation.(bool)
	}

	if !operation {
		return nil
	}

	request := oci_resource_analytics.ResourceAnalyticsInstanceDisableOacRequest{}

	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		tmp := resourceAnalyticsInstanceId.(string)
		request.ResourceAnalyticsInstanceId = &tmp
	}

	// Set a small JSON payload to ensure Content-Length header is properly set
	// requestBody := "{}"
	// request.RequestBody = &requestBody

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics")

	response, err := s.Client.ResourceAnalyticsInstanceDisableOac(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	err = s.getResourceAnalyticsInstanceOacManagementFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "resource_analytics"), oci_resource_analytics.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
	if err != nil {
		return err
	}

	s.setResToFetchedWorkRequest(err, response.OpcWorkRequestId, request.RequestMetadata.RetryPolicy)
	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) SetData() error {
	// Check if we have any response data to work with
	if s.Res == nil {
		return fmt.Errorf("resource state is nil")
	}

	// Check if we have either enable or disable response
	//if s.Res.enableResponse == nil && s.Res.disableResponse == nil {
	//	return fmt.Errorf("no valid response data available for state synchronization")
	//}

	// Set the resource analytics instance ID if available
	if resourceAnalyticsInstanceId, ok := s.D.GetOkExists("resource_analytics_instance_id"); ok {
		s.D.Set("resource_analytics_instance_id", resourceAnalyticsInstanceId)
	}

	// Set the enable operation status
	if enableOperation, ok := s.D.GetOkExists("enable_oac"); ok {
		s.D.Set("enable_oac", enableOperation)
		//s.D.Set("state", s.Res.LifecycleState)
		//s.D.Set("state", s.Res.LifecycleState)
		//if s.Res.enableResponse != nil {
		//	s.D.Set("state", oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateActive)
		//} else {
		//	s.D.Set("state", oci_resource_analytics.ResourceAnalyticsInstanceLifecycleStateDeleted)
		//}
	}

	// Set attachment details if they exist
	if attachmentDetails, ok := s.D.GetOkExists("attachment_details"); ok {
		s.D.Set("attachment_details", attachmentDetails)
	}

	// Set attachment type if it exists
	if attachmentType, ok := s.D.GetOkExists("attachment_type"); ok {
		s.D.Set("attachment_type", attachmentType)
	}

	return nil
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) mapToResourceAnalyticsInstanceOacAttachmentDetails(fieldKeyFormat string) (oci_resource_analytics.ResourceAnalyticsInstanceOacAttachmentDetails, error) {
	result := oci_resource_analytics.ResourceAnalyticsInstanceOacAttachmentDetails{}

	if idcsDomainId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "idcs_domain_id")); ok {
		tmp := idcsDomainId.(string)
		result.IdcsDomainId = &tmp
	}

	if licenseModel, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "license_model")); ok {
		result.LicenseModel = oci_resource_analytics.ResourceAnalyticsInstanceOacAttachmentDetailsLicenseModelEnum(licenseModel.(string))
	}

	if networkDetails, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_details")); ok {
		if tmpList := networkDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "network_details"), 0)
			tmp, err := s.mapToResourceAnalyticsInstanceOacNetworkDetails(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert network_details, encountered error: %v", err)
			}
			result.NetworkDetails = &tmp
		}
	}

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
			result.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func ResourceAnalyticsInstanceOacAttachmentDetailsToMap(obj *oci_resource_analytics.ResourceAnalyticsInstanceOacAttachmentDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdcsDomainId != nil {
		result["idcs_domain_id"] = string(*obj.IdcsDomainId)
	}

	result["license_model"] = string(obj.LicenseModel)

	if obj.NetworkDetails != nil {
		result["network_details"] = []interface{}{ResourceAnalyticsInstanceOacNetworkDetailsToMap(obj.NetworkDetails, false)}
	}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}

func (s *ResourceAnalyticsResourceAnalyticsInstanceOacManagementResourceCrud) mapToResourceAnalyticsInstanceOacNetworkDetails(fieldKeyFormat string) (oci_resource_analytics.ResourceAnalyticsInstanceOacNetworkDetails, error) {
	result := oci_resource_analytics.ResourceAnalyticsInstanceOacNetworkDetails{}

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
			result.NsgIds = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
		tmp := subnetId.(string)
		result.SubnetId = &tmp
	}

	return result, nil
}

func ResourceAnalyticsInstanceOacNetworkDetailsToMap(obj *oci_resource_analytics.ResourceAnalyticsInstanceOacNetworkDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	nsgIds := []interface{}{}
	for _, item := range obj.NsgIds {
		nsgIds = append(nsgIds, item)
	}
	if datasource {
		result["nsg_ids"] = nsgIds
	} else {
		result["nsg_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	return result
}
