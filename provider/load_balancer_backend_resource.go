// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strconv"
	"strings"
	"sync"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func BackendResource() *schema.Resource {
	return &schema.Resource{
		Create: createBackend,
		Read:   readBackend,
		Update: updateBackend,
		Delete: deleteBackend,
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

func createBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updateBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteBackend(d *schema.ResourceData, m interface{}) error {
	sync := &BackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type BackendResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Backend
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

// The create, update, and delete operations may implicitly modify the associated backend set resource. This
// may happen concurrently with an update to oci_loadbalancer_backend_set. Use a per-backend set
// mutex to synchronize accesses to the backend set.
func (s *BackendResourceCrud) GetMutex() *sync.Mutex {
	return lbBackendSetMutexes.GetOrCreateBackendSetMutex(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string))
}

func (s *BackendResourceCrud) buildID() string {
	return s.D.Get("ip_address").(string) + ":" + strconv.Itoa(s.D.Get("port").(int))
}

func (s *BackendResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		// Always inferred this way
		return s.buildID()
	}
	return ""
}

func (s *BackendResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *BackendResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *BackendResourceCrud) Create() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *BackendResourceCrud) Get() error {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetBackend(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Backend
	return nil
}

func (s *BackendResourceCrud) Update() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateBackend(context.Background(), request)
	if err != nil {
		return err
	}

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *BackendResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteBackend(context.Background(), request)

	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	getWorkRequestRequest.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	s.D.SetId(*workReqID)
	return nil
}

func (s *BackendResourceCrud) SetData() {
	if s.Res == nil {
		return
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

}
