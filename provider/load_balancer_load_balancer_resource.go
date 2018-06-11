// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func LoadBalancerResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Create: createLoadBalancer,
		Read:   readLoadBalancer,
		Update: updateLoadBalancer,
		Delete: deleteLoadBalancer,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shape": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_ids": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_addresses": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

func updateLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancer(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type LoadBalancerResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.LoadBalancer
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *LoadBalancerResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return *s.WorkRequest.LoadBalancerId
	}
	return ""
}

func (s *LoadBalancerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateCreating),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateActive),
		string(oci_load_balancer.LoadBalancerLifecycleStateFailed),
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleting),
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.LoadBalancerLifecycleStateDeleted),
	}
}

func (s *LoadBalancerResourceCrud) Create() error {
	request := oci_load_balancer.CreateLoadBalancerRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if isPrivate, ok := s.D.GetOkExists("is_private"); ok {
		tmp := isPrivate.(bool)
		request.IsPrivate = &tmp
	}

	if shape, ok := s.D.GetOkExists("shape"); ok {
		tmp := shape.(string)
		request.ShapeName = &tmp
	}

	request.SubnetIds = []string{}
	if subnetIds, ok := s.D.GetOkExists("subnet_ids"); ok {
		interfaces := subnetIds.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.SubnetIds = tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateLoadBalancer(context.Background(), request)
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
	s.D.Set("state", s.WorkRequest.LifecycleState)
	return nil
}

func (s *LoadBalancerResourceCrud) Get() error {
	id, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	if id == "" && s.WorkRequest != nil {
		id = *s.WorkRequest.LoadBalancerId
		s.D.SetId(id)
	}

	request := oci_load_balancer.GetLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.LoadBalancer
	return nil
}

func (s *LoadBalancerResourceCrud) Update() error {
	request := oci_load_balancer.UpdateLoadBalancerRequest{}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateLoadBalancer(context.Background(), request)
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

func (s *LoadBalancerResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteLoadBalancerRequest{}

	tmp := s.D.Id()
	request.LoadBalancerId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteLoadBalancer(context.Background(), request)

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

func (s *LoadBalancerResourceCrud) SetData() {
	// The first time this is called, we haven't actually fetched the resource yet, we just got a work request
	if s.Res == nil || s.Res.Id == nil {
		return
	}
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	/*
		if s.Res.Id != nil {
			s.D.Set("id", *s.Res.Id)
		}
	*/

	ipAddresses := []string{}
	for _, ad := range s.Res.IpAddresses {
		if ad.IpAddress != nil {
			ipAddresses = append(ipAddresses, *ad.IpAddress)
		}
	}
	s.D.Set("ip_addresses", ipAddresses)

	if s.Res.IsPrivate != nil {
		s.D.Set("is_private", *s.Res.IsPrivate)
	}

	if s.Res.ShapeName != nil {
		s.D.Set("shape", *s.Res.ShapeName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	s.D.Set("subnet_ids", s.Res.SubnetIds)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
