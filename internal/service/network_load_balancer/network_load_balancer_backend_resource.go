// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancerBackendResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createNetworkLoadBalancerBackend,
		Read:     readNetworkLoadBalancerBackend,
		Update:   updateNetworkLoadBalancerBackend,
		Delete:   deleteNetworkLoadBalancerBackend,
		Schema: map[string]*schema.Schema{
			// Required
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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
				ForceNew: true,
			},
			"target_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
		},
	}
}

func createNetworkLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readNetworkLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateNetworkLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteNetworkLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type NetworkLoadBalancerBackendResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_network_load_balancer.NetworkLoadBalancerClient
	Res                    *oci_network_load_balancer.Backend
	DisableNotFoundRetries bool
}

// The Create, Update, and delete operations may implicitly modify the associated backend set resource. This
// may happen concurrently with an Update to oci_network_load_balancer_backend_set. Use a per-backend set
// mutex to synchronize accesses to the backend set.
// This replicates the LBaaS (oci_loadbalancer_backend) behavior.
func (s *NetworkLoadBalancerBackendResourceCrud) GetMutex() *sync.Mutex {
	return nlbBackendSetMutexes.GetOrCreateNlbBackendSetMutex(s.D.Get("network_load_balancer_id").(string), s.D.Get("backend_set_name").(string))
}

func (s *NetworkLoadBalancerBackendResourceCrud) ID() string {
	tmp := s.determineBackendName()
	return GetNlbBackendCompositeId(tmp, s.D.Get("backend_set_name").(string), s.D.Get("network_load_balancer_id").(string))
}

func (s *NetworkLoadBalancerBackendResourceCrud) determineBackendName() string {
	if name, ok := s.D.GetOkExists("name"); ok {
		return name.(string)
	} else if targetId, ok := s.D.GetOkExists("target_id"); ok {
		return targetId.(string) + "." + strconv.Itoa(s.D.Get("port").(int))
	} else {
		return s.D.Get("ip_address").(string) + ":" + strconv.Itoa(s.D.Get("port").(int))
	}
}

func (s *NetworkLoadBalancerBackendResourceCrud) Create() error {
	request := oci_network_load_balancer.CreateBackendRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if isBackup, ok := s.D.GetOkExists("is_backup"); ok {
		tmp := isBackup.(bool)
		request.IsBackup = &tmp
	}

	if isDrain, ok := s.D.GetOkExists("is_drain"); ok {
		tmp := isDrain.(bool)
		request.IsDrain = &tmp
	}

	if isOffline, ok := s.D.GetOkExists("is_offline"); ok {
		tmp := isOffline.(bool)
		request.IsOffline = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if targetId, ok := s.D.GetOkExists("target_id"); ok {
		tmp := targetId.(string)
		request.TargetId = &tmp
	}

	if weight, ok := s.D.GetOkExists("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.CreateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	s.D.SetId(s.ID())
	return s.getBackendFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *NetworkLoadBalancerBackendResourceCrud) getBackendFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_network_load_balancer.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	nlbId, err := nlbBackendWaitForWorkRequest(workId,
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.Set("network_load_balancer_id", nlbId)

	return s.Get()
}

func backendWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
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

func nlbBackendWaitForWorkRequest(wId *string, action oci_network_load_balancer.ActionTypeEnum,
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
		errorMessage := getErrorFromNlbBackendWorkRequest(response.WorkRequest, client, retryPolicy)
		workRequestErr = fmt.Errorf("work request did not succeed, workId: %s, action: %s. Message: %s", *wId, action, errorMessage)
	}

	return identifier, workRequestErr
}

func getErrorFromNlbBackendWorkRequest(wr oci_network_load_balancer.WorkRequest,
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

func (s *NetworkLoadBalancerBackendResourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	tmp := s.determineBackendName()
	request.BackendName = &tmp

	backendName, backendSetName, networkLoadBalancerId, err := parseNlbBackendCompositeId(s.D.Id())
	if err == nil {
		request.BackendName = &backendName
		request.BackendSetName = &backendSetName
		request.NetworkLoadBalancerId = &networkLoadBalancerId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.GetBackend(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backend
	return nil
}

func (s *NetworkLoadBalancerBackendResourceCrud) Update() error {
	request := oci_network_load_balancer.UpdateBackendRequest{}

	if backendName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if isBackup, ok := s.D.GetOkExists("is_backup"); ok {
		tmp := isBackup.(bool)
		request.IsBackup = &tmp
	}

	if isDrain, ok := s.D.GetOkExists("is_drain"); ok {
		tmp := isDrain.(bool)
		request.IsDrain = &tmp
	}

	if isOffline, ok := s.D.GetOkExists("is_offline"); ok {
		tmp := isOffline.(bool)
		request.IsOffline = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	if weight, ok := s.D.GetOkExists("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.UpdateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getBackendFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer"), oci_network_load_balancer.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *NetworkLoadBalancerBackendResourceCrud) Delete() error {
	request := oci_network_load_balancer.DeleteBackendRequest{}

	if backendName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendsetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "network_load_balancer")

	response, err := s.Client.DeleteBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := networkLoadBalancerWaitForWorkRequest(workId,
		oci_network_load_balancer.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *NetworkLoadBalancerBackendResourceCrud) SetData() error {

	backendName, backendSetName, networkLoadBalancerId, err := parseNlbBackendCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &backendName)
		s.D.Set("backend_set_name", &backendSetName)
		s.D.Set("network_load_balancer_id", &networkLoadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.IsBackup != nil {
		s.D.Set("is_backup", *s.Res.IsBackup)
	}

	if s.Res.IsDrain != nil {
		s.D.Set("is_drain", *s.Res.IsDrain)
	}

	if s.Res.IsOffline != nil {
		s.D.Set("is_offline", *s.Res.IsOffline)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.Weight != nil {
		s.D.Set("weight", *s.Res.Weight)
	}

	if s.Res.TargetId != nil {
		s.D.Set("target_id", *s.Res.TargetId)
	}

	return nil
}

func GetNlbBackendCompositeId(backendName string, backendSetName string, networkLoadBalancerId string) string {
	backendName = url.PathEscape(backendName)
	backendSetName = url.PathEscape(backendSetName)
	networkLoadBalancerId = url.PathEscape(networkLoadBalancerId)
	compositeId := "networkLoadBalancers/" + networkLoadBalancerId + "/backendSets/" + backendSetName + "/backends/" + backendName
	return compositeId
}

func parseNlbBackendCompositeId(compositeId string) (backendName string, backendSetName string, networkLoadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("networkLoadBalancers/.*/backendSets/.*/backends/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	networkLoadBalancerId, _ = url.PathUnescape(parts[1])
	backendSetName, _ = url.PathUnescape(parts[3])
	backendName, _ = url.PathUnescape(parts[5])

	return
}

func NlbBackendSummaryToMap(obj oci_network_load_balancer.BackendSummary) map[string]interface{} {
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

	if obj.Weight != nil {
		result["weight"] = int(*obj.Weight)
	}

	if obj.TargetId != nil {
		result["target_id"] = string(*obj.TargetId)
	}

	return result
}
