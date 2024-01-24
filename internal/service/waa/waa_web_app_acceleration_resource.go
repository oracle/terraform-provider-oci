// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package waa

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_waa "github.com/oracle/oci-go-sdk/v65/waa"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func WaaWebAppAccelerationResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createWaaWebAppAcceleration,
		Read:     readWaaWebAppAcceleration,
		Update:   updateWaaWebAppAcceleration,
		Delete:   deleteWaaWebAppAcceleration,
		Schema: map[string]*schema.Schema{
			// Required
			"backend_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"LOAD_BALANCER",
				}, true),
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"web_app_acceleration_policy_id": {
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
			"system_tags": {
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

func createWaaWebAppAcceleration(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.CreateResource(d, sync)
}

func readWaaWebAppAcceleration(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()

	return tfresource.ReadResource(sync)
}

func updateWaaWebAppAcceleration(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteWaaWebAppAcceleration(d *schema.ResourceData, m interface{}) error {
	sync := &WaaWebAppAccelerationResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).WaaClient()
	sync.DisableNotFoundRetries = true
	sync.WorkRequestClient = m.(*client.OracleClients).WaaWorkRequestClient()

	return tfresource.DeleteResource(d, sync)
}

type WaaWebAppAccelerationResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_waa.WaaClient
	Res                    *oci_waa.WebAppAcceleration
	DisableNotFoundRetries bool
	WorkRequestClient      *oci_waa.WorkRequestClient
}

func (s *WaaWebAppAccelerationResourceCrud) ID() string {
	webAppAcceleration := *s.Res
	return *webAppAcceleration.GetId()
}

func (s *WaaWebAppAccelerationResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_waa.WebAppAccelerationLifecycleStateCreating),
	}
}

func (s *WaaWebAppAccelerationResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_waa.WebAppAccelerationLifecycleStateActive),
	}
}

func (s *WaaWebAppAccelerationResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_waa.WebAppAccelerationLifecycleStateDeleting),
	}
}

func (s *WaaWebAppAccelerationResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_waa.WebAppAccelerationLifecycleStateDeleted),
	}
}

func (s *WaaWebAppAccelerationResourceCrud) Create() error {
	request := oci_waa.CreateWebAppAccelerationRequest{}
	err := s.populateTopLevelPolymorphicCreateWebAppAccelerationRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.CreateWebAppAcceleration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getWebAppAccelerationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *WaaWebAppAccelerationResourceCrud) getWebAppAccelerationFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_waa.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	webAppAccelerationId, err := webAppAccelerationWaitForWorkRequest(workId, "webappacceleration",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestClient)

	if err != nil {
		return err
	}
	s.D.SetId(*webAppAccelerationId)

	return s.Get()
}

func webAppAccelerationWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "waa", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_waa.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func webAppAccelerationWaitForWorkRequest(wId *string, entityType string, action oci_waa.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_waa.WorkRequestClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "waa")
	retryPolicy.ShouldRetryOperation = webAppAccelerationWorkRequestShouldRetryFunc(timeout)

	response := oci_waa.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_waa.WorkRequestStatusInProgress),
			string(oci_waa.WorkRequestStatusAccepted),
			string(oci_waa.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_waa.WorkRequestStatusSucceeded),
			string(oci_waa.WorkRequestStatusFailed),
			string(oci_waa.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_waa.GetWorkRequestRequest{
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
	if identifier == nil || response.Status == oci_waa.WorkRequestStatusFailed || response.Status == oci_waa.WorkRequestStatusCanceled {
		return nil, getErrorFromWaaWebAppAccelerationWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromWaaWebAppAccelerationWorkRequest(client *oci_waa.WorkRequestClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_waa.WorkRequestResourceActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_waa.ListWorkRequestErrorsRequest{
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

func (s *WaaWebAppAccelerationResourceCrud) Get() error {
	request := oci_waa.GetWebAppAccelerationRequest{}

	tmp := s.D.Id()
	request.WebAppAccelerationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.GetWebAppAcceleration(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.WebAppAcceleration
	return nil
}

func (s *WaaWebAppAccelerationResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_waa.UpdateWebAppAccelerationRequest{}

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

	if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
		convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.SystemTags = convertedSystemTags
	}

	tmp := s.D.Id()
	request.WebAppAccelerationId = &tmp

	if webAppAccelerationPolicyId, ok := s.D.GetOkExists("web_app_acceleration_policy_id"); ok {
		tmp := webAppAccelerationPolicyId.(string)
		request.WebAppAccelerationPolicyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.UpdateWebAppAcceleration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppAccelerationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *WaaWebAppAccelerationResourceCrud) Delete() error {
	request := oci_waa.DeleteWebAppAccelerationRequest{}

	tmp := s.D.Id()
	request.WebAppAccelerationId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.DeleteWebAppAcceleration(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := webAppAccelerationWaitForWorkRequest(workId, "webappacceleration",
		oci_waa.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestClient)
	return delWorkRequestErr
}

func (s *WaaWebAppAccelerationResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_waa.WebAppAccelerationLoadBalancer:
		s.D.Set("backend_type", "LOAD_BALANCER")

		if v.LoadBalancerId != nil {
			s.D.Set("load_balancer_id", *v.LoadBalancerId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		s.D.Set("state", v.LifecycleState)

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.WebAppAccelerationPolicyId != nil {
			s.D.Set("web_app_acceleration_policy_id", *v.WebAppAccelerationPolicyId)
		}
	default:
		log.Printf("[WARN] Received 'backend_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func WebAppAccelerationSummaryToMap(obj oci_waa.WebAppAccelerationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (obj).(type) {
	case oci_waa.WebAppAccelerationLoadBalancerSummary:
		result["backend_type"] = "LOAD_BALANCER"

		if v.LoadBalancerId != nil {
			result["load_balancer_id"] = string(*v.LoadBalancerId)
		}
	default:
		log.Printf("[WARN] Received 'backend_type' of unknown type %v", obj)
		return nil
	}

	if obj.GetCompartmentId() != nil {
		result["compartment_id"] = string(*obj.GetCompartmentId())
	}

	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}

	if obj.GetDisplayName() != nil {
		result["display_name"] = string(*obj.GetDisplayName())
	}

	result["freeform_tags"] = obj.GetFreeformTags()

	if obj.GetId() != nil {
		result["id"] = string(*obj.GetId())
	}

	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = string(*obj.GetLifecycleDetails())
	}

	result["state"] = string(obj.GetLifecycleState())

	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}

	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}

	if obj.GetWebAppAccelerationPolicyId() != nil {
		result["web_app_acceleration_policy_id"] = obj.GetWebAppAccelerationPolicyId()
	}

	return result
}

func (s *WaaWebAppAccelerationResourceCrud) populateTopLevelPolymorphicCreateWebAppAccelerationRequest(request *oci_waa.CreateWebAppAccelerationRequest) error {
	//discriminator
	backendTypeRaw, ok := s.D.GetOkExists("backend_type")
	var backendType string
	if ok {
		backendType = backendTypeRaw.(string)
	} else {
		backendType = "" // default value
	}
	switch strings.ToLower(backendType) {
	case strings.ToLower("LOAD_BALANCER"):
		details := oci_waa.CreateWebAppAccelerationLoadBalancerDetails{}
		if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
			tmp := loadBalancerId.(string)
			details.LoadBalancerId = &tmp
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
		if systemTags, ok := s.D.GetOkExists("system_tags"); ok {
			convertedSystemTags, err := tfresource.MapToSystemTags(systemTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.SystemTags = convertedSystemTags
		}
		if webAppAccelerationPolicyId, ok := s.D.GetOkExists("web_app_acceleration_policy_id"); ok {
			tmp := webAppAccelerationPolicyId.(string)
			details.WebAppAccelerationPolicyId = &tmp
		}
		request.CreateWebAppAccelerationDetails = details
	default:
		return fmt.Errorf("unknown backend_type '%v' was specified", backendType)
	}
	return nil
}

func (s *WaaWebAppAccelerationResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_waa.ChangeWebAppAccelerationCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.WebAppAccelerationId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa")

	response, err := s.Client.ChangeWebAppAccelerationCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getWebAppAccelerationFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "waa"), oci_waa.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
