// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"
)

func LoadBalancerBackendResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createLoadBalancerBackend,
		Read:     readLoadBalancerBackend,
		Update:   updateLoadBalancerBackend,
		Delete:   deleteLoadBalancerBackend,
		Schema: map[string]*schema.Schema{
			// Required
			"backendset_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancer_id": {
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
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"drain": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"offline": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},

			// Computed
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.CreateResource(d, sync)
}

func readLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

func updateLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteLoadBalancerBackend(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type LoadBalancerBackendResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Backend
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

// The Create, Update, and delete operations may implicitly modify the associated backend set resource. This
// may happen concurrently with an Update to oci_loadbalancer_backend_set. Use a per-backend set
// mutex to synchronize accesses to the backend set.
func (s *LoadBalancerBackendResourceCrud) GetMutex() *sync.Mutex {
	return lbBackendSetMutexes.GetOrCreateBackendSetMutex(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string))
}

func (s *LoadBalancerBackendResourceCrud) buildID() string {
	return s.D.Get("ip_address").(string) + ":" + strconv.Itoa(s.D.Get("port").(int))
}

func (s *LoadBalancerBackendResourceCrud) ID() string {
	if s.WorkRequest != nil {
		if s.WorkRequest.LifecycleState == oci_load_balancer.WorkRequestLifecycleStateSucceeded {
			//API expects backendName to be in ip_address:port format
			return GetBackendCompositeId(s.buildID(), s.D.Get("backendset_name").(string), s.D.Get("load_balancer_id").(string))
		} else {
			return *s.WorkRequest.Id
		}
	}
	return ""
}

func (s *LoadBalancerBackendResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerBackendResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerBackendResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerBackendResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerBackendResourceCrud) Create() error {
	request := oci_load_balancer.CreateBackendRequest{}

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if backup, ok := s.D.GetOkExists("backup"); ok {
		tmp := backup.(bool)
		request.Backup = &tmp
	}

	if drain, ok := s.D.GetOkExists("drain"); ok {
		tmp := drain.(bool)
		request.Drain = &tmp
	}

	if ipAddress, ok := s.D.GetOkExists("ip_address"); ok {
		tmp := ipAddress.(string)
		request.IpAddress = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if offline, ok := s.D.GetOkExists("offline"); ok {
		tmp := offline.(bool)
		request.Offline = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if weight, ok := s.D.GetOkExists("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerBackendResourceCrud) Get() error {
	_, stillWorking, err := tfresource.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetBackendRequest{}

	tmp := s.buildID()
	request.BackendName = &tmp

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if !strings.HasPrefix(s.D.Id(), "ocid1.loadbalancerworkrequest.") {
		backendName, backendsetName, loadBalancerId, err := parseBackendCompositeId(s.D.Id())
		if err == nil {
			request.BackendName = &backendName
			request.BackendSetName = &backendsetName
			request.LoadBalancerId = &loadBalancerId
		} else {
			log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetBackend(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backend
	return nil
}

func (s *LoadBalancerBackendResourceCrud) Update() error {
	request := oci_load_balancer.UpdateBackendRequest{}

	if backendName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if backup, ok := s.D.GetOkExists("backup"); ok {
		tmp := backup.(bool)
		request.Backup = &tmp
	}

	if drain, ok := s.D.GetOkExists("drain"); ok {
		tmp := drain.(bool)
		request.Drain = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if offline, ok := s.D.GetOkExists("offline"); ok {
		tmp := offline.(bool)
		request.Offline = &tmp
	}

	if weight, ok := s.D.GetOkExists("weight"); ok {
		tmp := weight.(int)
		request.Weight = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerBackendResourceCrud) Delete() error {
	request := oci_load_balancer.DeleteBackendRequest{}

	if backendName, ok := s.D.GetOkExists("name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = tfresource.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	return nil
}

func (s *LoadBalancerBackendResourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	backendName, backendsetName, loadBalancerId, err := parseBackendCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("name", &backendName)
		s.D.Set("backendset_name", &backendsetName)
		s.D.Set("load_balancer_id", &loadBalancerId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.Backup != nil {
		s.D.Set("backup", *s.Res.Backup)
	}

	if s.Res.Drain != nil {
		s.D.Set("drain", *s.Res.Drain)
	}

	if s.Res.IpAddress != nil {
		s.D.Set("ip_address", *s.Res.IpAddress)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Offline != nil {
		s.D.Set("offline", *s.Res.Offline)
	}

	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}

	if s.Res.Weight != nil {
		s.D.Set("weight", *s.Res.Weight)
	}

	return nil
}

func GetBackendCompositeId(backendName string, backendsetName string, loadBalancerId string) string {
	backendName = url.PathEscape(backendName)
	backendsetName = url.PathEscape(backendsetName)
	loadBalancerId = url.PathEscape(loadBalancerId)
	compositeId := "loadBalancers/" + loadBalancerId + "/backendSets/" + backendsetName + "/backends/" + backendName
	return compositeId
}

func parseBackendCompositeId(compositeId string) (backendName string, backendsetName string, loadBalancerId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("loadBalancers/.*/backendSets/.*/backends/.*", compositeId)
	if !match || len(parts) != 6 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	loadBalancerId, _ = url.PathUnescape(parts[1])
	backendsetName, _ = url.PathUnescape(parts[3])
	backendName, _ = url.PathUnescape(parts[5])

	return
}
