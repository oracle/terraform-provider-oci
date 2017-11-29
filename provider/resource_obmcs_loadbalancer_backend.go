// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
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
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerBackend(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerBackendResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
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

func (s *LoadBalancerBackendResourceCrud) Create() (e error) {

	opts := &baremetal.CreateLoadBalancerBackendOptions{}
	if v, ok := s.D.GetOk("backup"); ok {
		opts.Backup = v.(bool)
	}
	if v, ok := s.D.GetOk("drain"); ok {
		opts.Drain = v.(bool)
	}
	if v, ok := s.D.GetOk("offline"); ok {
		opts.Offline = v.(bool)
	}
	if v, ok := s.D.GetOk("weight"); ok {
		opts.Weight = v.(int)
	}

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
	res, e := s.Client.GetBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.buildID(), nil)
	if e == nil {
		s.Resource = res
	}
	return
}

func (s *LoadBalancerBackendResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateLoadBalancerBackendOptions{}
	if v, ok := s.D.GetOk("backup"); ok {
		opts.Backup = v.(bool)
	}
	if v, ok := s.D.GetOk("drain"); ok {
		opts.Drain = v.(bool)
	}
	if v, ok := s.D.GetOk("offline"); ok {
		opts.Offline = v.(bool)
	}
	if v, ok := s.D.GetOk("weight"); ok {
		opts.Weight = v.(int)
	}

	var workReqID string
	workReqID, e = s.Client.UpdateBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.D.Id(), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	if e != nil {
		return
	}
	e = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest)
	if e != nil {
		return
	}
	return s.Get()
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
	// TODO: make sure this actually works
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return
	}
	var workReqID string
	workReqID, e = s.Client.DeleteBackend(s.D.Get("load_balancer_id").(string), s.D.Get("backendset_name").(string), s.D.Id(), nil)
	if e != nil {
		return
	}
	s.D.SetId(workReqID)
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
