// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v56/networkloadbalancer"
)

func NetworkLoadBalancerBackendSetResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkLoadBalancerBackendSet,
		Read:     readNetworkLoadBalancerBackendSet,
		Update:   updateNetworkLoadBalancerBackendSet,
		Delete:   deleteNetworkLoadBalancerBackendSet,
		Schema: map[string]*schema.Schema{
			// Required
			"health_checker": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"protocol": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"interval_in_millis": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"request_data": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"response_body_regex": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"response_data": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"retries": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"return_code": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"timeout_in_millis": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"url_path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			// Optional
			"is_preserve_source": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
			"backends": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_backup": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_drain": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"is_offline": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"weight": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
		},
	}
}

func createNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkLoadBalancerBackendSet(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendSetResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkLoadBalancerBackendSetResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_load_balancer.NetworkLoadBalancerClient
	Res                    *oci_network_load_balancer.BackendSet
	DisableNotFoundRetries bool
}

// The oci_network_load_balancer_backend resource may implicitly modify this backend set and this could happen concurrently.
// Use a per-backend set mutex to synchronize accesses to the backend set.
// This replicates the LBaaS (oci_loadbalancer_backend) behavior.
func (s *NetworkLoadBalancerBackendSetResourceCrud) GetMutex() *sync.Mutex {
	return nlbBackendSetMutexes.GetOrCreateNlbBackendSetMutex(s.D.Get("network_load_balancer_id").(string), s.D.Get("name").(string))
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) ID() string {
	return GetNlbBackendSetCompositeId(s.D.Get("name").(string), s.D.Get("network_load_balancer_id").(string))
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) Create() error {
	request := oci_network_load_balancer.CreateBackendSetRequest{}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToNetworkLoadBalancerHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HealthChecker = &tmp
		}
	}

	if ipVersion, ok := s.D.GetOkExists("ip_version"); ok {
		request.IpVersion = oci_network_load_balancer.IpVersionEnum(ipVersion.(string))
	}

	if isPreserveSource, ok := s.D.GetOkExists("is_preserve_source"); ok {
		tmp := isPreserveSource.(bool)
		request.IsPreserveSource = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		request.Policy = oci_network_load_balancer.NetworkLoadBalancingPolicyEnum(policy.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.CreateBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBackendSetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) getBackendSetFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_load_balancer.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	nlbId, err := nlbBackendSetWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.Set("network_load_balancer_id", nlbId)

	return s.Get()
}

func backendSetWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "network_load_balancer", startTime) {
			return true
		}

		// Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_network_load_balancer.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func nlbBackendSetWaitForWorkRequest(wId *string, action oci_network_load_balancer.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_network_load_balancer.NetworkLoadBalancerClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "network_load_balancer")
	retryPolicy.ShouldRetryOperation = networkLoadBalancerWorkRequestShouldRetryFunc(timeout)

	response := oci_network_load_balancer.GetWorkRequestResponse{}
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			string(oci_network_load_balancer.OperationStatusInProgress),
			string(oci_network_load_balancer.OperationStatusAccepted),
			string(oci_network_load_balancer.OperationStatusCanceling),
		},
		Target: []string{
			string(oci_network_load_balancer.OperationStatusSucceeded),
			string(oci_network_load_balancer.OperationStatusFailed),
			string(oci_network_load_balancer.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_network_load_balancer.GetWorkRequestRequest{
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
		if strings.Contains(strings.ToLower(*res.EntityType), "networkloadbalancer") {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	var workRequestErr error
	if response.WorkRequest.Status == oci_network_load_balancer.OperationStatusFailed {
		errorMessage := getErrorFromNlbBackendSetWorkRequest(response.WorkRequest, client, retryPolicy)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, action: %s. Message: %s", *wId, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromNlbBackendSetWorkRequest(wr oci_network_load_balancer.WorkRequest,
	client *oci_network_load_balancer.NetworkLoadBalancerClient, retryPolicy *oci_common.RetryPolicy) string {
	// Fetch the list of work request errors
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_network_load_balancer.ListWorkRequestErrorsRequest{
			WorkRequestId: wr.Id,
			CompartmentId: wr.CompartmentId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})

	if err != nil {
		return "Unknown failure reason"
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.WorkRequestErrorCollection.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")
	return errorMessage
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	backendSetName, networkLoadBalancerId, err := parseNlbBackendSetCompositeId(s.D.Id())
	if err == nil {
		request.BackendSetName = &backendSetName
		request.NetworkLoadBalancerId = &networkLoadBalancerId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.GetBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.BackendSet
	return nil
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) Update() error {
	request := oci_network_load_balancer.UpdateBackendSetRequest{}

	// @CODEGEN: Backends are marked computed in this resource, so will do a GET and include the results in the UPDATE, although they are not a required parameter.
	//           This behavior is intentionally set identical to LBaaS
	// Side-note: There is a potential for a race condition if the backend are added at the same time outside Terraform
	err := s.Get()
	if err != nil {
		return err
	}

	backends := []interface{}{}
	for _, item := range s.Res.Backends {
		backends = append(backends, NlbBackendToMap(item))
	}

	set := schema.NewSet(nlbBackendHashCodeForSets, backends)
	interfaces := set.List()
	tmp := make([]oci_network_load_balancer.BackendDetails, len(interfaces))
	for i := range interfaces {
		stateDataIndex := i
		fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backends", stateDataIndex)
		converted, err := s.mapToNlbBackendDetails(fieldKeyFormat)
		if err != nil {
			return err
		}
		tmp[i] = converted
	}
	request.Backends = tmp

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToNetworkLoadBalancerHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HealthChecker = &tmp
		}
	}
	if ipVersion, ok := s.D.GetOkExists("ip_version"); ok {
		request.IpVersion = oci_network_load_balancer.IpVersionEnum(ipVersion.(string))
	}
	if isPreserveSource, ok := s.D.GetOkExists("is_preserve_source"); ok {
		tmp := isPreserveSource.(bool)
		request.IsPreserveSource = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	if policy, ok := s.D.GetOkExists("policy"); ok {
		tmp := policy.(string)
		request.Policy = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBackendSetFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) Delete() error {
	request := oci_network_load_balancer.DeleteBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.DeleteBackendSet(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkLoadBalancerWaitForWorkRequest(workId,
		oci_network_load_balancer.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) SetData() error {

	backendSetName, networkLoadBalancerId, err := parseNlbBackendSetCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &backendSetName)
		s.D.Set("network_load_balancer_id", &networkLoadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	backends := []interface{}{}
	for _, item := range s.Res.Backends {
		backends = append(backends, NlbBackendToMap(item))
	}
	s.D.Set("backends", backends)

	if s.Res.HealthChecker != nil {
		s.D.Set("health_checker", []interface{}{NlbHealthCheckerToMap(s.Res.HealthChecker)})
	} else {
		s.D.Set("health_checker", nil)
	}
	s.D.Set("ip_version", s.Res.IpVersion)

	if s.Res.IsPreserveSource != nil {
		s.D.Set("is_preserve_source", *s.Res.IsPreserveSource)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("policy", s.Res.Policy)

	return nil
}

func GetNlbBackendSetCompositeId(backendSetName string, networkLoadBalancerId string) string {
	backendSetName = url.PathEscape(backendSetName)
	networkLoadBalancerId = url.PathEscape(networkLoadBalancerId)
	compositeId := "networkLoadBalancers/" + networkLoadBalancerId + "/backendSets/" + backendSetName
	return compositeId
}

func parseNlbBackendSetCompositeId(compositeId string) (backendSetName string, networkLoadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkLoadBalancers/.*/backendSets/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkLoadBalancerId, _ = url.PathUnescape(parts[1])
	backendSetName, _ = url.PathUnescape(parts[3])

	return
}

func NlbBackendToMap(obj oci_network_load_balancer.Backend) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IpAddress != nil {
		result["ip_address"] = string(*obj.IpAddress)
	}

	if obj.IsBackup != nil {
		result["is_backup"] = bool(*obj.IsBackup)
	}

	if obj.IsDrain != nil {
		result["is_drain"] = bool(*obj.IsDrain)
	}

	if obj.IsOffline != nil {
		result["is_offline"] = bool(*obj.IsOffline)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	return result
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) mapToNlbBackendDetails(fieldKeyFormat string) (oci_network_load_balancer.BackendDetails, error) {
	result := oci_network_load_balancer.BackendDetails{}

	if ipAddress, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ip_address")); ok {
		tmp := ipAddress.(string)
		result.IpAddress = &tmp
	}

	if isBackup, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_backup")); ok {
		tmp := isBackup.(bool)
		result.IsBackup = &tmp
	}

	if isDrain, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_drain")); ok {
		tmp := isDrain.(bool)
		result.IsDrain = &tmp
	}

	if isOffline, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_offline")); ok {
		tmp := isOffline.(bool)
		result.IsOffline = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if targetId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "target_id")); ok {
		tmp := targetId.(string)
		result.TargetId = &tmp
	}

	if weight, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "weight")); ok {
		tmp := weight.(int)
		result.Weight = &tmp
	}

	return result, nil
}

func NlbBackendSetSummaryToMap(obj oci_network_load_balancer.BackendSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	backends := []interface{}{}
	for _, item := range obj.Backends {
		backends = append(backends, NlbBackendToMap(item))
	}
	result["backends"] = backends

	if obj.HealthChecker != nil {
		result["health_checker"] = []interface{}{NlbHealthCheckerToMap(obj.HealthChecker)}
	}

	result["ip_version"] = string(obj.IpVersion)

	if obj.IsPreserveSource != nil {
		result["is_preserve_source"] = bool(*obj.IsPreserveSource)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["policy"] = string(obj.Policy)

	return result
}

func NlbHealthCheckerToMap(obj *oci_network_load_balancer.HealthChecker) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IntervalInMillis != nil {
		result["interval_in_millis"] = int(*obj.IntervalInMillis)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	result["protocol"] = string(obj.Protocol)

	if obj.RequestData != nil {
		result["request_data"] = base64.StdEncoding.EncodeToString(obj.RequestData)
	}

	if obj.ResponseBodyRegex != nil {
		result["response_body_regex"] = string(*obj.ResponseBodyRegex)
	}

	if obj.ResponseData != nil {
		result["response_data"] = base64.StdEncoding.EncodeToString(obj.ResponseData)
	}

	if obj.Retries != nil {
		result["retries"] = int(*obj.Retries)
	}

	if obj.ReturnCode != nil {
		result["return_code"] = int(*obj.ReturnCode)
	}

	if obj.TimeoutInMillis != nil {
		result["timeout_in_millis"] = int(*obj.TimeoutInMillis)
	}

	if obj.UrlPath != nil {
		result["url_path"] = string(*obj.UrlPath)
	}

	return result
}

func (s *NetworkLoadBalancerBackendSetResourceCrud) mapToNetworkLoadBalancerHealthCheckerDetails(fieldKeyFormat string) (oci_network_load_balancer.HealthCheckerDetails, error) {
	result := oci_network_load_balancer.HealthCheckerDetails{}

	if intervalInMillis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "interval_in_millis")); ok {
		tmp := intervalInMillis.(int)
		result.IntervalInMillis = &tmp
	}

	if timeoutInMillis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "timeout_in_millis")); ok {
		tmp := timeoutInMillis.(int)
		result.TimeoutInMillis = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	if retries, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "retries")); ok {
		tmp := retries.(int)
		result.Retries = &tmp
	}

	protocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "protocol"))
	if ok {
		result.Protocol = oci_network_load_balancer.HealthCheckProtocolsEnum(protocol.(string))
	}

	if protocol == "TCP" || protocol == "UDP" {
		if requestData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "request_data")); ok {
			tmp := requestData.(string)
			decoded, err := base64.StdEncoding.DecodeString(tmp)
			if err != nil {
				return result, err
			}
			result.RequestData = decoded
		}

		if responseData, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_data")); ok {
			tmp := responseData.(string)
			decoded, err := base64.StdEncoding.DecodeString(tmp)
			if err != nil {
				return result, err
			}
			result.ResponseData = decoded
		}
	} else if protocol == "HTTP" || protocol == "HTTPS" {
		if responseBodyRegex, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_body_regex")); ok {
			tmp := responseBodyRegex.(string)
			result.ResponseBodyRegex = &tmp
		}

		if returnCode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "return_code")); ok {
			tmp := returnCode.(int)
			result.ReturnCode = &tmp
		}

		if urlPath, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "url_path")); ok {
			tmp := urlPath.(string)
			result.UrlPath = &tmp
		}
	}

	return result, nil
}

func nlbBackendHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if backup, ok := m["backup"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", backup))
	}
	if drain, ok := m["drain"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", drain))
	}
	if ipAddress, ok := m["ip_address"]; ok && ipAddress != "" {
		buf.WriteString(fmt.Sprintf("%v-", ipAddress))
	}
	if offline, ok := m["offline"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", offline))
	}
	if port, ok := m["port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", port))
	}
	if targetId, ok := m["target_id"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", targetId))
	}
	if weight, ok := m["weight"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", weight))
	}
	return hashcode.String(buf.String())
}
