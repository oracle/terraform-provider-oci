// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_apigateway "github.com/oracle/oci-go-sdk/v41/apigateway"
	oci_common "github.com/oracle/oci-go-sdk/v41/common"
)

func init() {
	RegisterResource("oci_apigateway_gateway", ApigatewayGatewayResource())
}

func ApigatewayGatewayResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createApigatewayGateway,
		Read:     readApigatewayGateway,
		Update:   updateApigatewayGateway,
		Delete:   deleteApigatewayGateway,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"certificate_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
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
			"response_cache_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"EXTERNAL_RESP_CACHE",
								"NONE",
							}, true),
						},

						// Optional
						"authentication_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication_secret_version_number": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							ValidateFunc:     validateInt64TypeString,
							DiffSuppressFunc: int64StringDiffSuppressFunction,
						},
						"connect_timeout_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"is_ssl_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_ssl_verify_disabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"read_timeout_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"send_timeout_in_ms": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"servers": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"host": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
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

			// Computed
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ip_address": {
							Type:     schema.TypeString,
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

func createApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return CreateResource(d, sync)
}

func readApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return ReadResource(sync)
}

func updateApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient

	return UpdateResource(d, sync)
}

func deleteApigatewayGateway(d *schema.ResourceData, m interface{}) error {
	sync := &ApigatewayGatewayResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).gatewayClient()
	sync.WorkRequestsClient = m.(*OracleClients).gatewayWorkRequestsClient
	sync.DisableNotFoundRetries = true

	return DeleteResource(d, sync)
}

type ApigatewayGatewayResourceCrud struct {
	BaseCrud
	Client                 *oci_apigateway.GatewayClient
	WorkRequestsClient     *oci_apigateway.WorkRequestsClient
	Res                    *oci_apigateway.Gateway
	DisableNotFoundRetries bool
}

func (s *ApigatewayGatewayResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *ApigatewayGatewayResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_apigateway.GatewayLifecycleStateCreating),
	}
}

func (s *ApigatewayGatewayResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_apigateway.GatewayLifecycleStateActive),
	}
}

func (s *ApigatewayGatewayResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_apigateway.GatewayLifecycleStateDeleting),
	}
}

func (s *ApigatewayGatewayResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_apigateway.GatewayLifecycleStateDeleted),
	}
}

func (s *ApigatewayGatewayResourceCrud) Create() error {
	request := oci_apigateway.CreateGatewayRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
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

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if endpointType, ok := s.D.GetOkExists("endpoint_type"); ok {
		request.EndpointType = oci_apigateway.GatewayEndpointTypeEnum(endpointType.(string))
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if responseCacheDetails, ok := s.D.GetOkExists("response_cache_details"); ok {
		if tmpList := responseCacheDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_cache_details", 0)
			tmp, err := s.mapToResponseCacheDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCacheDetails = tmp
		}
	}

	if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
		tmp := subnetId.(string)
		request.SubnetId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.CreateGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGatewayFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *ApigatewayGatewayResourceCrud) getGatewayFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_apigateway.WorkRequestResourceActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	gatewayId, err := gatewayWaitForWorkRequest(workId, "gateway",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.WorkRequestsClient)

	if err != nil {
		// Try to cancel the work request
		log.Printf("[DEBUG] creation failed, attempting to cancel the workrequest: %v for identifier: %v\n", workId, gatewayId)
		_, cancelErr := s.WorkRequestsClient.CancelWorkRequest(context.Background(),
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
	s.D.SetId(*gatewayId)

	return s.Get()
}

func gatewayWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if shouldRetry(response, false, "apigateway", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_apigateway.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func gatewayWaitForWorkRequest(wId *string, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_apigateway.WorkRequestsClient) (*string, error) {
	retryPolicy := getRetryPolicy(disableFoundRetries, "apigateway")
	retryPolicy.ShouldRetryOperation = gatewayWorkRequestShouldRetryFunc(timeout)

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
		return nil, getErrorFromApigatewayGatewayWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromApigatewayGatewayWorkRequest(client *oci_apigateway.WorkRequestsClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_apigateway.WorkRequestResourceActionTypeEnum) error {
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

func (s *ApigatewayGatewayResourceCrud) Get() error {
	request := oci_apigateway.GetGatewayRequest{}

	tmp := s.D.Id()
	request.GatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.GetGateway(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Gateway
	return nil
}

func (s *ApigatewayGatewayResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_apigateway.UpdateGatewayRequest{}

	if certificateId, ok := s.D.GetOkExists("certificate_id"); ok {
		tmp := certificateId.(string)
		request.CertificateId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.GatewayId = &tmp

	if responseCacheDetails, ok := s.D.GetOkExists("response_cache_details"); ok {
		if tmpList := responseCacheDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "response_cache_details", 0)
			tmp, err := s.mapToResponseCacheDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.ResponseCacheDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.UpdateGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGatewayFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *ApigatewayGatewayResourceCrud) Delete() error {
	request := oci_apigateway.DeleteGatewayRequest{}

	tmp := s.D.Id()
	request.GatewayId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.DeleteGateway(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := gatewayWaitForWorkRequest(workId, "gateway",
		oci_apigateway.WorkRequestResourceActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.WorkRequestsClient)
	return delWorkRequestErr
}

func (s *ApigatewayGatewayResourceCrud) SetData() error {
	if s.Res.CertificateId != nil {
		s.D.Set("certificate_id", *s.Res.CertificateId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("endpoint_type", s.Res.EndpointType)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	ipAddresses := []interface{}{}
	for _, item := range s.Res.IpAddresses {
		ipAddresses = append(ipAddresses, GatewayIpAddressToMap(item))
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ResponseCacheDetails != nil {
		responseCacheDetailsArray := []interface{}{}
		if responseCacheDetailsMap := ResponseCacheDetailsToMap(&s.Res.ResponseCacheDetails); responseCacheDetailsMap != nil {
			responseCacheDetailsArray = append(responseCacheDetailsArray, responseCacheDetailsMap)
		}
		s.D.Set("response_cache_details", responseCacheDetailsArray)
	} else {
		s.D.Set("response_cache_details", nil)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func GatewaySummaryToMap(obj oci_apigateway.GatewaySummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertificateId != nil {
		result["certificate_id"] = string(*obj.CertificateId)
	}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	result["endpoint_type"] = string(obj.EndpointType)

	if obj.Hostname != nil {
		result["hostname"] = string(*obj.Hostname)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.LifecycleDetails != nil {
		result["lifecycle_details"] = string(*obj.LifecycleDetails)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = definedTagsToMap(obj.DefinedTags)
	}

	if obj.FreeformTags != nil {
		result["freeform_tags"] = obj.FreeformTags
	}

	return result
}

func GatewayIpAddressToMap(obj oci_apigateway.IpAddress) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	return result
}

func (s *ApigatewayGatewayResourceCrud) mapToResponseCacheDetails(fieldKeyFormat string) (oci_apigateway.ResponseCacheDetails, error) {
	var baseObject oci_apigateway.ResponseCacheDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("EXTERNAL_RESP_CACHE"):
		details := oci_apigateway.ExternalRespCache{}
		if authenticationSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_secret_id")); ok {
			tmp := authenticationSecretId.(string)
			details.AuthenticationSecretId = &tmp
		}
		if authenticationSecretVersionNumber, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_secret_version_number")); ok {
			tmp := authenticationSecretVersionNumber.(string)
			tmpInt64, err := strconv.ParseInt(tmp, 10, 64)
			if err != nil {
				return details, fmt.Errorf("unable to convert authenticationSecretVersionNumber string: %s to an int64 and encountered error: %v", tmp, err)
			}
			details.AuthenticationSecretVersionNumber = &tmpInt64
		}
		if connectTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connect_timeout_in_ms")); ok {
			tmp := connectTimeoutInMs.(int)
			details.ConnectTimeoutInMs = &tmp
		}
		if isSslEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ssl_enabled")); ok {
			tmp := isSslEnabled.(bool)
			details.IsSslEnabled = &tmp
		}
		if isSslVerifyDisabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_ssl_verify_disabled")); ok {
			tmp := isSslVerifyDisabled.(bool)
			details.IsSslVerifyDisabled = &tmp
		}
		if readTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_timeout_in_ms")); ok {
			tmp := readTimeoutInMs.(int)
			details.ReadTimeoutInMs = &tmp
		}
		if sendTimeoutInMs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "send_timeout_in_ms")); ok {
			tmp := sendTimeoutInMs.(int)
			details.SendTimeoutInMs = &tmp
		}
		if servers, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "servers")); ok {
			interfaces := servers.([]interface{})
			tmp := make([]oci_apigateway.ResponseCacheRespServer, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "servers"), stateDataIndex)
				converted, err := s.mapToResponseCacheRespServer(fieldKeyFormatNextLevel)
				if err != nil {
					return details, err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "servers")) {
				details.Servers = tmp
			}
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_apigateway.NoCache{}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func ResponseCacheDetailsToMap(obj *oci_apigateway.ResponseCacheDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_apigateway.ExternalRespCache:
		result["type"] = "EXTERNAL_RESP_CACHE"

		if v.AuthenticationSecretId != nil {
			result["authentication_secret_id"] = string(*v.AuthenticationSecretId)
		}

		if v.AuthenticationSecretVersionNumber != nil {
			result["authentication_secret_version_number"] = strconv.FormatInt(*v.AuthenticationSecretVersionNumber, 10)
		}

		if v.ConnectTimeoutInMs != nil {
			result["connect_timeout_in_ms"] = int(*v.ConnectTimeoutInMs)
		}

		if v.IsSslEnabled != nil {
			result["is_ssl_enabled"] = bool(*v.IsSslEnabled)
		}

		if v.IsSslVerifyDisabled != nil {
			result["is_ssl_verify_disabled"] = bool(*v.IsSslVerifyDisabled)
		}

		if v.ReadTimeoutInMs != nil {
			result["read_timeout_in_ms"] = int(*v.ReadTimeoutInMs)
		}

		if v.SendTimeoutInMs != nil {
			result["send_timeout_in_ms"] = int(*v.SendTimeoutInMs)
		}

		servers := []interface{}{}
		for _, item := range v.Servers {
			servers = append(servers, ResponseCacheRespServerToMap(item))
		}
		result["servers"] = servers
	case oci_apigateway.NoCache:
		result["type"] = "NONE"
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *ApigatewayGatewayResourceCrud) mapToResponseCacheRespServer(fieldKeyFormat string) (oci_apigateway.ResponseCacheRespServer, error) {
	result := oci_apigateway.ResponseCacheRespServer{}

	if host, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func ResponseCacheRespServerToMap(obj oci_apigateway.ResponseCacheRespServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func (s *ApigatewayGatewayResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_apigateway.ChangeGatewayCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.GatewayId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "apigateway")

	response, err := s.Client.ChangeGatewayCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getGatewayFromWorkRequest(workId, getRetryPolicy(s.DisableNotFoundRetries, "apigateway"), oci_apigateway.WorkRequestResourceActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
