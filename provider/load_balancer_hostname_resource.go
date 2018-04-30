// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func HostnameResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createHostname,
		Read:     readHostname,
		Update:   updateHostname,
		Delete:   deleteHostname,
		Schema: map[string]*schema.Schema{
			// Required
			"hostname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updateHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteHostname(d *schema.ResourceData, m interface{}) error {
	sync := &HostnameResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type HostnameResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Hostname
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *HostnameResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *HostnameResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *HostnameResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *HostnameResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *HostnameResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *HostnameResourceCrud) Create() error {
	request := oci_load_balancer.CreateHostnameRequest{}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateHostname(context.Background(), request)
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

func (s *HostnameResourceCrud) Get() error {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	request := oci_load_balancer.GetHostnameRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetHostname(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Hostname
	return nil
}

func (s *HostnameResourceCrud) Update() error {
	request := oci_load_balancer.UpdateHostnameRequest{}

	if hostname, ok := s.D.GetOkExists("hostname"); ok {
		tmp := hostname.(string)
		request.Hostname = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateHostname(context.Background(), request)
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

func (s *HostnameResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteHostnameRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteHostname(context.Background(), request)

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

func (s *HostnameResourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

}
