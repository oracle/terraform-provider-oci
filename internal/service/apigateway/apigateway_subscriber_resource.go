// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package apigateway

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func ApigatewaySubscriberResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createApigatewaySubscriber,
		Read:     readApigatewaySubscriber,
		Update:   updateApigatewaySubscriber,
		Delete:   deleteApigatewaySubscriber,
		Schema: map[string]*schema.Schema{
			// Required
			"clients": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"token": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"usage_plans": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

			// Computed
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
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

func createApigatewaySubscriber(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewaySubscriberResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.CreateResource(d, sync)
}

func readApigatewaySubscriber(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewaySubscriberResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()

	return tfresource.ReadResource(sync)
}

func updateApigatewaySubscriber(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewaySubscriberResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteApigatewaySubscriber(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewaySubscriberResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SubscribersClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).ApigatewayWorkRequestsClient()

	return tfresource.DeleteResource(d, sync)
}

type ApigatewaySubscriberResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_apigateway.SubscribersClient
	Res                    *oci_apigateway.Subscriber
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_apigateway.WorkRequestsClient
}

func (s *ApigatewaySubscriberResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApigatewaySubscriberResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apigateway.SubscriberLifecycleStateCreating),
	}
}

func (s *ApigatewaySubscriberResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apigateway.SubscriberLifecycleStateActive),
	}
}

func (s *ApigatewaySubscriberResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apigateway.SubscriberLifecycleStateDeleting),
	}
}

func (s *ApigatewaySubscriberResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apigateway.SubscriberLifecycleStateDeleted),
	}
}

func (s *ApigatewaySubscriberResourceCrud) Create() error {
	request := oci_apigateway.CreateSubscriberRequest{}

	if clients, ok := s.D.GetOkExists("clients"); ok {
		interfaces := clients.([]interface{})
		tmp := make([]oci_apigateway.Client, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "clients", stateDataIndex)
			converted, err := s.mapToClient(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("clients") {
			request.Clients = tmp
		}
	}

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

	if usagePlans, ok := s.D.GetOkExists("usage_plans"); ok {
		interfaces := usagePlans.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("usage_plans") {
			request.UsagePlans = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateSubscriber(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getSubscriberFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApigatewaySubscriberResourceCrud) getSubscriberFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	subscriberId, err := subscriberWaitForWorkRequest(workId, "subscriber",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, subscriberId)
		_, cancelErr := s.WorkRequestClient.CancelWorkRequest(context.Background(),
			oci_apigateway.CancelWorkRequestRequest{
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
	s.D.SetId(*subscriberId)

	return s.Get()
}

func subscriberWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "apigateway", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apigateway.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func subscriberWaitForWorkRequest(wId *string, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.WorkRequestsClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = subscriberWorkRequestShouldRetryFunc(timeout)

	response := oci_apigateway.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_apigateway.WorkRequestStatusInProgress),
			string(oci_apigateway.WorkRequestStatusAccepted),
			string(oci_apigateway.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_apigateway.WorkRequestStatusSucceeded),
			string(oci_apigateway.WorkRequestStatusFailed),
			string(oci_apigateway.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_apigateway.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_apigateway.WorkRequestStatusFailed || response.Status == oci_apigateway.WorkRequestStatusCanceled {
		return nil, getErrorFromApigatewaySubscriberWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApigatewaySubscriberWorkRequest(client *oci_apigateway.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_apigateway.ListWorkRequestErrorsRequest{
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

func (s *ApigatewaySubscriberResourceCrud) Get() error {
	request := oci_apigateway.GetSubscriberRequest{}

	tmp := s.D.Id()
	request.SubscriberId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.GetSubscriber(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Subscriber
	return nil
}

func (s *ApigatewaySubscriberResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apigateway.UpdateSubscriberRequest{}

	if clients, ok := s.D.GetOkExists("clients"); ok {
		interfaces := clients.([]interface{})
		tmp := make([]oci_apigateway.Client, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "clients", stateDataIndex)
			converted, err := s.mapToClient(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("clients") {
			request.Clients = tmp
		}
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

	tmp := s.D.Id()
	request.SubscriberId = &tmp

	if usagePlans, ok := s.D.GetOkExists("usage_plans"); ok {
		interfaces := usagePlans.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("usage_plans") {
			request.UsagePlans = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateSubscriber(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubscriberFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApigatewaySubscriberResourceCrud) Delete() error {
	request := oci_apigateway.DeleteSubscriberRequest{}

	tmp := s.D.Id()
	request.SubscriberId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteSubscriber(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := subscriberWaitForWorkRequest(workId, "subscriber",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *ApigatewaySubscriberResourceCrud) SetData() error {
	clients := []interface{}{}
	for _, item := range s.Res.Clients {
		clients = append(clients, ClientToMap(item))
	}
	s.D.Set("clients", clients)

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

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	s.D.Set("usage_plans", s.Res.UsagePlans)

	return nil
}

func (s *ApigatewaySubscriberResourceCrud) mapToClient(fieldKeyFormat string) (oci_apigateway.Client, error) {
	result := oci_apigateway.Client{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if token, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "token")); ok {
		tmp := token.(string)
		result.Token = &tmp
	}

	return result, nil
}

func ClientToMap(obj oci_apigateway.Client) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Token != nil {
		result["token"] = string(*obj.Token)
	}

	return result
}

func ClientSummaryToMap(obj oci_apigateway.ClientSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func SubscriberSummaryToMap(obj oci_apigateway.SubscriberSummary) map[string]interface{} {
	result := map[string]interface{}{}

	clients := []interface{}{}
	for _, item := range obj.Clients {
		clients = append(clients, ClientSummaryToMap(item))
	}
	result["clients"] = clients

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

	result["state"] = string(obj.LifecycleState)

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	result["usage_plans"] = obj.UsagePlans

	return result
}

func (s *ApigatewaySubscriberResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apigateway.ChangeSubscriberCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.SubscriberId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.ChangeSubscriberCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getSubscriberFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
