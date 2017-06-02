// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"log"
	"strconv"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func LoadBalancerBackendResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancerBackend,
		Read:   readLoadBalancerBackend,
		Update: updateLoadBalancerBackend,
		Delete: deleteLoadBalancerBackend,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"backendset_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"backup": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"drain": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"offline": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"weight": {
				Type:     schema.TypeInt,
				Optional: true,
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

func createLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.CreateResource(d, sync)
}

func readLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.ReadResource(sync)
}

func updateLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(client.BareMetalClient)
	return crud.DeleteResource(d, sync)
}

type LoadBalancerBackendResourceCrud struct {
	crud.BaseCrud
	WorkRequest *baremetal.WorkRequest
	Resource    *baremetal.Backend
}

func (s *LoadBalancerBackendResourceCrud) buildID() string {
	return s.D.Get("ip_address").(string) + ":" + strconv.Itoa(s.D.Get("port").(int))
}

func (s *LoadBalancerBackendResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Resource, s.WorkRequest)
	log.Printf("ID in load balancer backend ID(): %v", id)
	if id != nil {
		return *id
	}
	if workSuccess {
		// Always inferred this way
		return s.buildID()
	}
	return ""
}

func (s *LoadBalancerBackendResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
		baremetal.ResourceFailed,
	}
}

func (s *LoadBalancerBackendResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
	}
}

func makeBackendOptions(data *schema.ResourceData) *baremetal.CreateLoadBalancerBackendOptions {
	opts := &baremetal.CreateLoadBalancerBackendOptions{}
	if v, ok := data.GetOk("backup"); ok {
		opts.Backup = v.(bool)
	}
	if v, ok := data.GetOk("drain"); ok {
		opts.Drain = v.(bool)
	}
	if v, ok := data.GetOk("offline"); ok {
		opts.Offline = v.(bool)
	}
	if v, ok := data.GetOk("weight"); ok {
		opts.Weight = v.(int)
	}
	return opts
}

func (s *LoadBalancerBackendResourceCrud) Create() (e error) {

	opts := makeBackendOptions(s.D)

	var workReqID string
	workReqID, e = s.Client.CreateBackend(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("backendset_name").(string),
		s.D.Get("ip_address").(string),
		s.D.Get("port").(int),
		opts,
	)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendResourceCrud) Get() (e error) {
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.BaseCrud, s.WorkRequest)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}
	s.Resource, e = s.Client.GetBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.buildID(), nil)
	return
}

func (s *LoadBalancerBackendResourceCrud) Update() (e error) {
	opts := makeBackendOptions(s.D)

	var workReqID string
	workReqID, e = s.Client.UpdateBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.D.Id(), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendResourceCrud) SetData() {
	if s.Resource == nil {
		return
	}
	s.D.Set("backup", s.Resource.Backup)
	s.D.Set("drain", s.Resource.Drain)
	s.D.Set("offline", s.Resource.Offline)
	s.D.Set("weight", s.Resource.Weight)
}

func (s *LoadBalancerBackendResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.D.Id(), nil)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
