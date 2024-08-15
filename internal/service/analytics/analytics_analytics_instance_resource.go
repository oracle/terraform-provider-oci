// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package analytics

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_analytics "github.com/oracle/oci-go-sdk/v65/analytics"
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func AnalyticsAnalyticsInstanceResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: tfresource.GetTimeoutDuration("1h"),
			Update: tfresource.GetTimeoutDuration("1h"),
			Delete: tfresource.GetTimeoutDuration("1h"),
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
			"idcs_access_token": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				StateFunc: tfresource.GetMd5Hash,
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
			"admin_user": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"email_notification": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"feature_bundle": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"kms_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"network_endpoint_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"network_endpoint_type": {
							Type:             schema.TypeString,
							Required:         true,
							ForceNew:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"PRIVATE",
								"PUBLIC",
							}, true),
						},

						// Optional
						"network_security_group_ids": {
							Type:     schema.TypeSet,
							Optional: true,
							Computed: true,
							Set:      tfresource.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vcn_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"whitelisted_ips": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"whitelisted_services": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"whitelisted_vcns": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"whitelisted_ips": {
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
					},
				},
			},
			"state": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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

func createAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()
	var powerOff = false
	if powerState, ok := sync.D.GetOkExists("state"); ok {
		wantedPowerState := oci_analytics.AnalyticsInstanceLifecycleStateEnum(strings.ToUpper(powerState.(string)))
		if wantedPowerState == oci_analytics.AnalyticsInstanceLifecycleStateInactive {
			powerOff = true
		}
	}

	if e := tfresource.CreateResource(d, sync); e != nil {
		return e
	}

	if powerOff {
		if err := sync.StopAnalyticsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateInactive)
	}
	return nil

}

func (s *AnalyticsAnalyticsInstanceResourceCrud) SetKmsKey(kmsKeyId *string) error {
	request := oci_analytics.SetKmsKeyRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp
	request.KmsKeyId = kmsKeyId
	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.SetKmsKey(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	_, err = analyticsInstanceWaitForWorkRequest(workId, "analytics",
		oci_analytics.WorkRequestActionResultCreated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	return err
}

func readAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

	return tfresource.ReadResource(sync)
}

func updateAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()

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
		if err := sync.StartAnalyticsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateActive)
	}
	if sync.D.HasChange("kms_key_id") {
		wantedKmsKeyId := sync.D.Get("kms_key_id").(string)
		if err := sync.SetKmsKey(&wantedKmsKeyId); err != nil {
			// Re-read the instance to update the state file with correct values after failure
			err = tfresource.ReadResource(sync)
			return err
		}

		sync.D.Set("kms_key_id", wantedKmsKeyId)
	}

	if err := tfresource.UpdateResource(d, sync); err != nil {
		return err
	}

	if powerOff {
		if err := sync.StopAnalyticsInstance(); err != nil {
			return err
		}
		sync.D.Set("state", oci_analytics.AnalyticsInstanceLifecycleStateInactive)
	}

	return nil
}

func deleteAnalyticsAnalyticsInstance(d *schema.ResourceData, m interface{}) error {
	sync := &AnalyticsAnalyticsInstanceResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).AnalyticsClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type AnalyticsAnalyticsInstanceResourceCrud struct {
	tfresource.BaseCrud
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

	if adminUser, ok := s.D.GetOkExists("admin_user"); ok {
		tmp := adminUser.(string)
		request.AdminUser = &tmp
	}

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

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	if emailNotification, ok := s.D.GetOkExists("email_notification"); ok {
		tmp := emailNotification.(string)
		request.EmailNotification = &tmp
	}

	if featureBundle, ok := s.D.GetOkExists("feature_bundle"); ok {
		request.FeatureBundle = oci_analytics.FeatureBundleEnum(featureBundle.(string))
	}

	if featureSet, ok := s.D.GetOkExists("feature_set"); ok {
		request.FeatureSet = oci_analytics.FeatureSetEnum(featureSet.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if idcsAccessToken, ok := s.D.GetOkExists("idcs_access_token"); ok {
		tmp := idcsAccessToken.(string)
		request.IdcsAccessToken = &tmp
	}

	if kmsKeyId, ok := s.D.GetOkExists("kms_key_id"); ok {
		tmp := kmsKeyId.(string)
		request.KmsKeyId = &tmp
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_analytics.LicenseTypeEnum(licenseType.(string))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkEndpointDetails, ok := s.D.GetOkExists("network_endpoint_details"); ok {
		if tmpList := networkEndpointDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_endpoint_details", 0)
			tmp, err := s.mapToNetworkEndpointDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.NetworkEndpointDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.CreateAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.Id
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) getAnalyticsInstanceFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_analytics.WorkRequestActionResultEnum, timeout time.Duration) error {

	// Wait until it finishes
	analyticsInstanceId, err := analyticsInstanceWaitForWorkRequest(workId, "analytics",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, analyticsInstanceId)
		_, cancelErr := s.Client.DeleteWorkRequest(context.Background(),
			oci_analytics.DeleteWorkRequestRequest{
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
		if tfresource.ShouldRetry(response, false, "analytics", startTime) {
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
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "analytics")
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
	for _, res := range response.Resources {
		if res.ResourceType == "ANALYTICS_INSTANCE" {
			if res.ActionResult == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_analytics.WorkRequestStatusFailed || response.Status == oci_analytics.WorkRequestStatusCanceled {
		return nil, getErrorFromAnalyticsAnalyticsInstanceWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromAnalyticsAnalyticsInstanceWorkRequest(client *oci_analytics.AnalyticsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_analytics.WorkRequestActionResultEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_analytics.ListWorkRequestErrorsRequest{
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

func (s *AnalyticsAnalyticsInstanceResourceCrud) Get() error {
	request := oci_analytics.GetAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

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

	if networkEndpointDetails, ok := s.D.GetOkExists("network_endpoint_details"); ok && s.D.HasChange("network_endpoint_details") {
		oldNetworkEndpoint, newNetworkEndpoint := s.D.GetChange("network_endpoint_details")
		if oldNetworkEndpoint != nil && newNetworkEndpoint != nil {

			if tmpList := networkEndpointDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "network_endpoint_details", 0)
				//check if the new network endpoint details and old are different
				hasNpChanged, err := s.hasNEPChanged(fieldKeyFormat)
				if err != nil {
					return err
				}
				if hasNpChanged {
					tmp, err := s.mapToNetworkEndpointDetails(fieldKeyFormat)
					if err != nil {
						return err
					}
					err = s.updateNetworkEndpoint(tmp)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	if emailNotification, ok := s.D.GetOkExists("email_notification"); ok {
		tmp := emailNotification.(string)
		request.EmailNotification = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if licenseType, ok := s.D.GetOkExists("license_type"); ok {
		request.LicenseType = oci_analytics.LicenseTypeEnum(licenseType.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

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

			scaleRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")
			scaleResponse, err := s.Client.ScaleAnalyticsInstance(context.Background(), scaleRequest)

			if err != nil {
				return err
			}

			workId := scaleResponse.OpcWorkRequestId
			return s.getAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultScaled, s.D.Timeout(schema.TimeoutUpdate))
		}
	}

	return nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) Delete() error {
	request := oci_analytics.DeleteAnalyticsInstanceRequest{}

	tmp := s.D.Id()
	request.AnalyticsInstanceId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.DeleteAnalyticsInstance(context.Background(), request)
	time.Sleep(2 * time.Minute) //We add this to prevent 412-PreconditionFailed, NetworkSecurityGroup cannot be deleted since it still has vnics attached to it
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
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DomainId != nil {
		s.D.Set("domain_id", *s.Res.DomainId)
	}

	if s.Res.EmailNotification != nil {
		s.D.Set("email_notification", *s.Res.EmailNotification)
	}

	s.D.Set("feature_bundle", s.Res.FeatureBundle)

	s.D.Set("feature_set", s.Res.FeatureSet)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KmsKeyId != nil {
		s.D.Set("kms_key_id", *s.Res.KmsKeyId)
	}

	s.D.Set("license_type", s.Res.LicenseType)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.NetworkEndpointDetails != nil {
		networkEndpointDetailsArray := []interface{}{}
		if networkEndpointDetailsMap := NetworkEndpointDetailsToMap(&s.Res.NetworkEndpointDetails, false); networkEndpointDetailsMap != nil {
			networkEndpointDetailsArray = append(networkEndpointDetailsArray, networkEndpointDetailsMap)
		}
		s.D.Set("network_endpoint_details", networkEndpointDetailsArray)
	} else {
		s.D.Set("network_endpoint_details", nil)
	}

	if s.Res.ServiceUrl != nil {
		s.D.Set("service_url", *s.Res.ServiceUrl)
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

func (s *AnalyticsAnalyticsInstanceResourceCrud) StartAnalyticsInstance() error {
	request := oci_analytics.StartAnalyticsInstanceRequest{}

	idTmp := s.D.Id()
	request.AnalyticsInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	_, err := s.Client.StartAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_analytics.AnalyticsInstanceLifecycleStateActive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) StopAnalyticsInstance() error {
	request := oci_analytics.StopAnalyticsInstanceRequest{}

	idTmp := s.D.Id()
	request.AnalyticsInstanceId = &idTmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	_, err := s.Client.StopAnalyticsInstance(context.Background(), request)
	if err != nil {
		return err
	}

	retentionPolicyFunc := func() bool { return s.Res.LifecycleState == oci_analytics.AnalyticsInstanceLifecycleStateInactive }
	return tfresource.WaitForResourceCondition(s, retentionPolicyFunc, s.D.Timeout(schema.TimeoutUpdate))
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

// check if the new network endpoint details and old are different
func (s *AnalyticsAnalyticsInstanceResourceCrud) hasNEPChanged(fieldKeyFormat string) (bool, error) {
	var hasChange bool
	networkEndpointTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_endpoint_type"))
	var networkEndpointType string
	if ok {
		networkEndpointType = networkEndpointTypeRaw.(string)
	} else {
		networkEndpointType = "" // default value
	}
	switch strings.ToLower(networkEndpointType) {
	case strings.ToLower("PRIVATE"):
		if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "network_security_group_ids")) {
			hasChange = true
		}
		if !hasChange && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "subnet_id")) {
			hasChange = true
		}
		if !hasChange && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "vcn_id")) {
			hasChange = true
		}
	case strings.ToLower("PUBLIC"):
		if s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_ips")) {
			hasChange = true
		}
		if !hasChange && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_services")) {
			hasChange = true
		}
		if !hasChange && s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_vcns")) {
			hasChange = true
		}
	default:
		return false, fmt.Errorf("unknown network_endpoint_type '%v' was specified", networkEndpointType)
	}
	return hasChange, nil
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) mapToNetworkEndpointDetails(fieldKeyFormat string) (oci_analytics.NetworkEndpointDetails, error) {
	var baseObject oci_analytics.NetworkEndpointDetails
	//discriminator
	networkEndpointTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_endpoint_type"))
	var networkEndpointType string
	if ok {
		networkEndpointType = networkEndpointTypeRaw.(string)
	} else {
		networkEndpointType = "" // default value
	}
	switch strings.ToLower(networkEndpointType) {
	case strings.ToLower("PRIVATE"):
		details := oci_analytics.PrivateEndpointDetails{}
		if networkSecurityGroupIds, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "network_security_group_ids")); ok {
			set := networkSecurityGroupIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "network_security_group_ids")) {
				details.NetworkSecurityGroupIds = tmp
			}
		}
		if subnetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "subnet_id")); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vcnId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vcn_id")); ok {
			tmp := vcnId.(string)
			details.VcnId = &tmp
		}
		baseObject = details
	case strings.ToLower("PUBLIC"):
		details := oci_analytics.PublicEndpointDetails{}
		if whitelistedIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelisted_ips")); ok {
			interfaces := whitelistedIps.([]interface{})
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_ips")) {
				details.WhitelistedIps = tmp
			}
		}
		if whitelistedServices, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelisted_services")); ok {
			interfaces := whitelistedServices.([]interface{})
			tmp := make([]oci_analytics.AccessControlServiceTypeEnum, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(oci_analytics.AccessControlServiceTypeEnum)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_services")) {
				details.WhitelistedServices = tmp
			}
		}
		if whitelistedVcns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelisted_vcns")); ok {
			interfaces := whitelistedVcns.([]interface{})
			tmp := make([]oci_analytics.VirtualCloudNetwork, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "whitelisted_vcns"), stateDataIndex)
				converted, err := s.mapToVirtualCloudNetwork(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_vcns")) {
				details.WhitelistedVcns = tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown network_endpoint_type '%v' was specified", networkEndpointType)
	}
	return baseObject, nil
}

func NetworkEndpointDetailsToMap(obj *oci_analytics.NetworkEndpointDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_analytics.PrivateEndpointDetails:
		result["network_endpoint_type"] = "PRIVATE"

		networkSecurityGroupIds := []interface{}{}
		for _, item := range v.NetworkSecurityGroupIds {
			networkSecurityGroupIds = append(networkSecurityGroupIds, item)
		}
		if datasource {
			result["network_security_group_ids"] = networkSecurityGroupIds
		} else {
			result["network_security_group_ids"] = schema.NewSet(tfresource.LiteralTypeHashCodeForSets, networkSecurityGroupIds)
		}

		if v.SubnetId != nil {
			result["subnet_id"] = string(*v.SubnetId)
		}

		if v.VcnId != nil {
			result["vcn_id"] = string(*v.VcnId)
		}
	case oci_analytics.PublicEndpointDetails:
		result["network_endpoint_type"] = "PUBLIC"

		result["whitelisted_ips"] = v.WhitelistedIps

		result["whitelisted_services"] = v.WhitelistedServices

		whitelistedVcns := []interface{}{}
		for _, item := range v.WhitelistedVcns {
			whitelistedVcns = append(whitelistedVcns, VirtualCloudNetworkToMap(item))
		}
		result["whitelisted_vcns"] = whitelistedVcns
	default:
		log.Printf("[WARN] Received 'network_endpoint_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) mapToVirtualCloudNetwork(fieldKeyFormat string) (oci_analytics.VirtualCloudNetwork, error) {
	result := oci_analytics.VirtualCloudNetwork{}

	if id, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id")); ok {
		tmp := id.(string)
		result.Id = &tmp
	}

	if whitelistedIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "whitelisted_ips")); ok {
		interfaces := whitelistedIps.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "whitelisted_ips")) {
			result.WhitelistedIps = tmp
		}
	}

	return result, nil
}

func VirtualCloudNetworkToMap(obj oci_analytics.VirtualCloudNetwork) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	result["whitelisted_ips"] = obj.WhitelistedIps

	return result
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_analytics.ChangeAnalyticsInstanceCompartmentRequest{}

	idTmp := s.D.Id()
	changeCompartmentRequest.AnalyticsInstanceId = &idTmp

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")

	response, err := s.Client.ChangeAnalyticsInstanceCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultCompartmentChanged, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *AnalyticsAnalyticsInstanceResourceCrud) updateNetworkEndpoint(networkEndpointDetails oci_analytics.NetworkEndpointDetails) error {

	idTmp := s.D.Id()
	changeEndPointRequest := oci_analytics.ChangeAnalyticsInstanceNetworkEndpointRequest{}
	changeEndPointRequest.NetworkEndpointDetails = networkEndpointDetails
	changeEndPointRequest.AnalyticsInstanceId = &idTmp
	changeEndPointRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics")
	response, err := s.Client.ChangeAnalyticsInstanceNetworkEndpoint(context.Background(), changeEndPointRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getAnalyticsInstanceFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "analytics"), oci_analytics.WorkRequestActionResultNetworkEndpointChanged, s.D.Timeout(schema.TimeoutUpdate))
}
