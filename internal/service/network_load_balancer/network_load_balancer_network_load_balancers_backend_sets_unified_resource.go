// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

func NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified,
		Read:     readNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified,
		Update:   updateNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified,
		Delete:   deleteNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified,
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
						"dns": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"domain_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"query_class": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"query_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rcodes": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"transport_protocol": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
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
			"backends": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      backendsHashCodeForSets,
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
			"ip_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_fail_open": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_preserve_source": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnified(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_load_balancer.NetworkLoadBalancerClient
	Res                    *oci_network_load_balancer.BackendSet
	DisableNotFoundRetries bool
}

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) ID() string {
	return GetNetworkLoadBalancersBackendSetsUnifiedCompositeId(s.D.Get("name").(string), s.D.Get("network_load_balancer_id").(string))
}

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) Create() error {
	request := oci_network_load_balancer.CreateBackendSetRequest{}

	if backends, ok := s.D.GetOkExists("backends"); ok {
		set := backends.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_network_load_balancer.BackendDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := backendsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backends", stateDataIndex)
			converted, err := s.mapToBackendDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("backends") {
			request.Backends = tmp
		}
	}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HealthChecker = &tmp
		}
	}

	if ipVersion, ok := s.D.GetOkExists("ip_version"); ok {
		request.IpVersion = oci_network_load_balancer.IpVersionEnum(ipVersion.(string))
	}

	if isFailOpen, ok := s.D.GetOkExists("is_fail_open"); ok {
		tmp := isFailOpen.(bool)
		request.IsFailOpen = &tmp
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
	s.D.SetId(s.ID())
	return s.getNetworkLoadBalancersBackendSetsUnifiedFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) getNetworkLoadBalancersBackendSetsUnifiedFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_load_balancer.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	networkLoadBalancersBackendSetsUnifiedId, err := networkLoadBalancersBackendSetsUnifiedWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.Set("network_load_balancer_id", networkLoadBalancersBackendSetsUnifiedId)

	return s.Get()
}

func networkLoadBalancersBackendSetsUnifiedWaitForWorkRequest(wId *string, action oci_network_load_balancer.ActionTypeEnum,
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
		errorMessage := getErrorFromNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedWorkRequest(response.WorkRequest, client, retryPolicy)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, action: %s. Message: %s", *wId, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromNetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedWorkRequest(wr oci_network_load_balancer.WorkRequest,
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

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	backendSetName, networkLoadBalancerId, err := parseNetworkLoadBalancersBackendSetsUnifiedCompositeId(s.D.Id())
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

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) Update() error {
	request := oci_network_load_balancer.UpdateBackendSetRequest{}

	if backendSetName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if backends, ok := s.D.GetOkExists("backends"); ok {
		set := backends.(*schema.Set)
		interfaces := set.List()
		tmp := make([]oci_network_load_balancer.BackendDetails, len(interfaces))
		for i := range interfaces {
			stateDataIndex := backendsHashCodeForSets(interfaces[i])
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "backends", stateDataIndex)
			converted, err := s.mapToBackendDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("backends") {
			request.Backends = tmp
		}
	}

	if healthChecker, ok := s.D.GetOkExists("health_checker"); ok {
		if tmpList := healthChecker.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "health_checker", 0)
			tmp, err := s.mapToHealthCheckerDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.HealthChecker = &tmp
		}
	}

	if ipVersion, ok := s.D.GetOkExists("ip_version"); ok {
		request.IpVersion = oci_network_load_balancer.IpVersionEnum(ipVersion.(string))
	}

	if isFailOpen, ok := s.D.GetOkExists("is_fail_open"); ok {
		tmp := isFailOpen.(bool)
		request.IsFailOpen = &tmp
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
	return s.getNetworkLoadBalancersBackendSetsUnifiedFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) Delete() error {
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

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) SetData() error {

	backendSetName, networkLoadBalancerId, err := parseNetworkLoadBalancersBackendSetsUnifiedCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &backendSetName)
		s.D.Set("network_load_balancer_id", &networkLoadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	backends := []interface{}{}
	for _, item := range s.Res.Backends {
		backends = append(backends, BackendToMap(item))
	}
	s.D.Set("backends", schema.NewSet(backendsHashCodeForSets, backends))

	if s.Res.HealthChecker != nil {
		s.D.Set("health_checker", []interface{}{HealthCheckerToMap(s.Res.HealthChecker)})
	} else {
		s.D.Set("health_checker", nil)
	}

	s.D.Set("ip_version", s.Res.IpVersion)

	if s.Res.IsFailOpen != nil {
		s.D.Set("is_fail_open", *s.Res.IsFailOpen)
	}

	if s.Res.IsPreserveSource != nil {
		s.D.Set("is_preserve_source", *s.Res.IsPreserveSource)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("policy", s.Res.Policy)

	return nil
}

func GetNetworkLoadBalancersBackendSetsUnifiedCompositeId(backendSetName string, networkLoadBalancerId string) string {
	backendSetName = url.PathEscape(backendSetName)
	networkLoadBalancerId = url.PathEscape(networkLoadBalancerId)
	compositeId := "networkLoadBalancers/" + networkLoadBalancerId + "/backendSets/" + backendSetName
	return compositeId
}

func parseNetworkLoadBalancersBackendSetsUnifiedCompositeId(compositeId string) (backendSetName string, networkLoadBalancerId string, err error) {
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

func BackendToMap(obj oci_network_load_balancer.Backend) map[string]interface{} {
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

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) mapToBackendDetails(fieldKeyFormat string) (oci_network_load_balancer.BackendDetails, error) {
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

func BackendSetSummaryToMap(obj oci_network_load_balancer.BackendSetSummary) map[string]interface{} {
	result := map[string]interface{}{}

	// TODO: How does this work?
	backends := []interface{}{}
	for _, item := range obj.Backends {
		backends = append(backends, BackendToMap(item))
	}
	result["backends"] = backends

	if obj.HealthChecker != nil {
		result["health_checker"] = []interface{}{HealthCheckerToMap(obj.HealthChecker)}
	}

	result["ip_version"] = string(obj.IpVersion)

	if obj.IsFailOpen != nil {
		result["is_fail_open"] = bool(*obj.IsFailOpen)
	}

	if obj.IsPreserveSource != nil {
		result["is_preserve_source"] = bool(*obj.IsPreserveSource)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["policy"] = string(obj.Policy)

	return result
}

func DnsHealthCheckRCodesToMap(obj oci_network_load_balancer.DnsHealthCheckRCodesEnum) map[string]interface{} {
	result := map[string]interface{}{}

	return result
}

func DnsHealthCheckerDetailsToMap(obj *oci_network_load_balancer.DnsHealthCheckerDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	result["query_class"] = string(obj.QueryClass)

	result["query_type"] = string(obj.QueryType)

	rcodes := make([]string, 0, 4)
	for _, item := range obj.Rcodes {
		rcodes = append(rcodes, string(item))
	}

	result["rcodes"] = rcodes

	result["transport_protocol"] = string(obj.TransportProtocol)

	return result
}

func HealthCheckerToMap(obj *oci_network_load_balancer.HealthChecker) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Dns != nil {
		result["dns"] = []interface{}{DnsHealthCheckerDetailsToMap(obj.Dns)}
	}

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

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) mapToHealthCheckerDetails(fieldKeyFormat string) (oci_network_load_balancer.HealthCheckerDetails, error) {
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
	} else if protocol == "DNS" {
		if dns, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "dns")); ok {
			if tmpList := dns.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "dns"), 0)
				tmp, err := s.mapToDnsHealthCheckerDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return result, fmt.Errorf("unable to convert dns, encountered error: %v", err)
				}
				result.Dns = &tmp
			}
		}
	}

	return result, nil
}

func (s *NetworkLoadBalancerNetworkLoadBalancersBackendSetsUnifiedResourceCrud) mapToDnsHealthCheckerDetails(fieldKeyFormat string) (oci_network_load_balancer.DnsHealthCheckerDetails, error) {
	result := oci_network_load_balancer.DnsHealthCheckerDetails{}

	if domainName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_name")); ok {
		tmp := domainName.(string)
		result.DomainName = &tmp
	}

	transportProtocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "transport_protocol"))
	if ok {
		result.TransportProtocol = oci_network_load_balancer.DnsHealthCheckTransportProtocolsEnum(transportProtocol.(string))
	}

	queryClass, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_class"))
	if ok {
		result.QueryClass = oci_network_load_balancer.DnsHealthCheckQueryClassesEnum(queryClass.(string))
	}

	queryType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "query_type"))
	if ok {
		result.QueryType = oci_network_load_balancer.DnsHealthCheckQueryTypesEnum(queryType.(string))
	}
	rCodes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "rcodes"))
	if ok {
		tmp := rCodes.([]interface{})
		result.Rcodes = make([]oci_network_load_balancer.DnsHealthCheckRCodesEnum, 0, 4)
		for _, v := range tmp {
			rCode, ok := oci_network_load_balancer.GetMappingDnsHealthCheckRCodesEnum(v.(string))
			if ok {
				result.Rcodes = append(result.Rcodes, rCode)
			}
		}
	}

	return result, nil
}

func backendsHashCodeForSets(v interface{}) int {
	var buf bytes.Buffer
	m := v.(map[string]interface{})
	if ipAddress, ok := m["ip_address"]; ok && ipAddress != "" {
		buf.WriteString(fmt.Sprintf("%v-", ipAddress))
	}
	if isBackup, ok := m["is_backup"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isBackup))
	}
	if isDrain, ok := m["is_drain"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isDrain))
	}
	if isOffline, ok := m["is_offline"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", isOffline))
	}
	if port, ok := m["port"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", port))
	}
	if targetId, ok := m["target_id"]; ok && targetId != "" {
		buf.WriteString(fmt.Sprintf("%v-", targetId))
	}
	if weight, ok := m["weight"]; ok {
		buf.WriteString(fmt.Sprintf("%v-", weight))
	}
	return utils.GetStringHashcode(buf.String())
}
