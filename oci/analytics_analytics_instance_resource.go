// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"

	oci_analytics "github.com/oracle/oci-go-sdk/analytics"
	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterResource("oci_analytics_analytics_instance", AnalyticsAnalyticsInstanceResource())
}

func AnalyticsAnalyticsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: getTimeoutDuration("1h"),
			Update: getTimeoutDuration("1h"),
			Delete: getTimeoutDuration("1h"),
		},
		Create: createAnalyticsAnalyticsInstance,
		Read:   readAnalyticsAnalyticsInstance,
		Update: updateAnalyticsAnalyticsInstance,
		Delete: deleteAnalyticsAnalyticsInstance,
		Schema: map[string]*schema.Schema{
			// Required
			"capacity": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"capacity_type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"capacity_value": {
							Type:     schema.TypeInt,
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
			"feature_set": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_notification": {
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
			"idcs_access_token": {
				Type:      schema.TypeString,
				Optional:  true,
				StateFunc: getMd5Hash,
				Sensitive: true,
			},
			"state": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					string(oci_analytics.AnalyticsInstanceLifecycleStateInactive),
					string(oci_analytics.AnalyticsInstanceLifecycleStateActive),
				}, true),
			},

			// Computed
			"service_url": {
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

func createAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).analyticsClient()

	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_analytics.AnalyticsInstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_analytics.AnalyticsInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopOacInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateInactive)
	}
	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) StartOacInstance() error {
	request := oci_analytics.StartAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.StartAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultStarted, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) StopOacInstance() error {
	request := oci_analytics.StopAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.StopAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultStopped, s.D.Timeout(schema.TimeoutUpdate))
}

func readAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).analyticsClient()

	return ReadResource(sync)
}

func updateAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).analyticsClient()

	// switch to power on
	powerOn, powerOff := false, false

	if sync.D.HasChange("state") {
		wantedState := strings.ToUpper(sync.D.Get("state").(string))
		if oci_analytics.AnalyticsInstanceLifecycleStateActive == oci_analytics.AnalyticsInstanceLifecycleStateEnum(wantedState) {
			powerOn = true
		} else if oci_analytics.AnalyticsInstanceLifecycleStateInactive == oci_analytics.AnalyticsInstanceLifecycleStateEnum(wantedState) {
			powerOff = true
		}
	}

	if powerOn {
		if err := sync.StartOacInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateActive)
	}

	if err := UpdateResource(d, sync); err != nil {
		return err
	}

	// switch to power off
	if powerOff {
		if err := sync.StopOacInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateInactive)
	}
	return nil
}

func deleteAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).analyticsClient()
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type AnalyticsAnalyticsInstanceResourceCrud struct {
	BaseCrud
	Client                 *oci_analytics.AnalyticsClient
	Res                    *oci_analytics.AnalyticsInstance
	DisableNotFoundRetries bool
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateCreating),
	}
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateActive),
	}
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateDeleting),
	}
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_analytics.AnalyticsInstanceLifecycleStateDeleted),
	}
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) Create() error {
	request := oci_analytics.CreateAnalyticsInstanceRequest{}

	if capacity, ok := s.D.GetOkExists("capacity"); ok {
		if tmpList := capacity.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "capacity", 0)
			tmp, err := s.mapToCapacity(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.Capacity = &tmp
		}
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if emailNotification, ok := s.D.GetOkExists("email_notification"); ok {
		tmp := emailNotification.(string)
		request.EmailNotification = &tmp
	}

	if featureSet, ok := s.D.GetOkExists("feature_set"); ok {
		request.FeatureSet = oci_analytics.FeatureSetEnum(featureSet.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_analytics.LicenseTypeEnum(licenseType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.CreateAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) getAnalyticsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_analytics.WorkRequestActionResultEnum, timeout time.Duration) error {

	// Wait until it finishes
	analyticsInstanceId, err := analyticsInstanceWaitForWorkRequest(workId, "analytics",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed: %v for identifier: %v\n", workId, analyticsInstanceId)
		return err
	}

	if analyticsInstanceId == nil {
		return fmt.Errorf("operation failed: %v for identifier: %v\n", workId, analyticsInstanceId)
	}

	s.D.SetId(*analyticsInstanceId)

	return s.Get()
}

func analyticsInstanceWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "analytics", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_analytics.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func analyticsInstanceWaitForWorkRequest(wId *string, entityType string, action oci_analytics.WorkRequestActionResultEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_analytics.AnalyticsClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "analytics")
	retryPolicy.ShouldRetryOperation = analyticsInstanceWorkRequestShouldRetryFunc(timeout)

	response := oci_analytics.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_analytics.WorkRequestStatusInProgress),
			string(oci_analytics.WorkRequestStatusAccepted),
			string(oci_analytics.WorkRequestStatusCanceling),
		},
		Target: []string{
			string(oci_analytics.WorkRequestStatusSucceeded),
			string(oci_analytics.WorkRequestStatusFailed),
			string(oci_analytics.WorkRequestStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_analytics.GetWorkRequestRequest{
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
	for _, res := range response.WorkRequest.Resources {
		if res.ResourceType == "ANALYTICS_INSTANCE" {
			if res.ActionResult == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The OAC workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_analytics.WorkRequestStatusFailed || response.Status == oci_analytics.WorkRequestStatusCanceled {
		return nil, getErrorFromAnalyticsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAnalyticsInstanceWorkRequest(client *oci_analytics.AnalyticsClient, wId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_analytics.WorkRequestActionResultEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_analytics.ListWorkRequestErrorsRequest{
			WorkRequestId: wId,
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

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *wId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) Get() error {
	request := oci_analytics.GetAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.GetAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyticsInstance
	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_analytics.UpdateAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if emailNotification, ok := s.D.GetOkExists("email_notification"); ok {
		tmp := emailNotification.(string)
		request.EmailNotification = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_analytics.LicenseTypeEnum(licenseType.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.UpdateAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.AnalyticsInstance

	// capacity change if any, perform upscaling or downscaling
	if capacity, ok := s.D.GetOkExists("capacity"); ok && s.D.HasChange("capacity") {
		scaleRequest := oci_analytics.ScaleAnalyticsInstanceRequest{}
		if tmpList := capacity.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "capacity", 0)
			tmp, err := s.mapToCapacity(fieldKeyFormat)
			if err != nil {
				return err
			}

			// instance id
			id := s.D.Id()
			scaleRequest.AnalyticsInstanceId = &id
			scaleRequest.Capacity = &tmp

			scaleRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")
			scaleResponse, err := s.Client.ScaleAnalyticsInstance(context.Background(), scaleRequest)

			if err != nil {
				return err
			}

			workId := scaleResponse.OpcWorkRequestId
			return s.getAnalyticsInstanceFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultScaled, s.D.Timeout(schema.TimeoutUpdate))
		}
	}

	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) Delete() error {
	request := oci_analytics.DeleteAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.DeleteAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := analyticsInstanceWaitForWorkRequest(workId, "analytics",
		oci_analytics.WorkRequestActionResultDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) SetData() error {
	if s.Res.Capacity != nil {
		s.D.Set("capacity", []interface{}{AnalyticsCapacityToMap(s.Res.Capacity)})
	} else {
		s.D.Set("capacity", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.EmailNotification != nil {
		s.D.Set("email_notification", *s.Res.EmailNotification)
	}

	s.D.Set("feature_set", s.Res.FeatureSet)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("license_type", s.Res.LicenseType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.ServiceUrl != nil {
		s.D.Set("service_url", *s.Res.ServiceUrl)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}
	s.D.Set("tim", "")
	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) mapToCapacity(fieldKeyFormat string) (oci_analytics.Capacity, error) {
	result := oci_analytics.Capacity{}

	if capacityType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_type")); ok {
		result.CapacityType = oci_analytics.CapacityTypeEnum(capacityType.(string))
	}

	if capacityValue, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capacity_value")); ok {
		tmp := capacityValue.(int)
		result.CapacityValue = &tmp
	}

	return result, nil
}

func AnalyticsCapacityToMap(obj *oci_analytics.Capacity) map[string]interface{} {
	result := map[string]interface{}{}

	result["capacity_type"] = string(obj.CapacityType)

	if obj.CapacityValue != nil {
		result["capacity_value"] = int(*obj.CapacityValue)
	}

	return result
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_analytics.ChangeAnalyticsInstanceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AnalyticsInstanceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "analytics")

	_, err := s.Client.ChangeAnalyticsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}
	return nil
}
