// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/oracle/bmcs-go-sdk"

	"github.com/oracle/terraform-provider-oci/crud"
)

func LoadBalancerListenerResource() *schema.Resource {
	return &schema.Resource{
		Create: createLoadBalancerListener,
		Read:   readLoadBalancerListener,
		Update: updateLoadBalancerListener,
		Delete: deleteLoadBalancerListener,
		Schema: map[string]*schema.Schema{
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
			"default_backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_configuration": SSLConfigSchema,
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.CreateResource(d, sync)
}

func readLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.ReadResource(sync)
}

func updateLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).client
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).clientWithoutNotFoundRetries
	return crud.DeleteResource(d, sync)
}

type LoadBalancerListenerResourceCrud struct {
	crud.BaseCrud
	WorkRequest *baremetal.WorkRequest
	Resource    *baremetal.Listener
}

// ID uniquely identifies the listener and its parent load balancer
func (s *LoadBalancerListenerResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Resource, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *LoadBalancerListenerResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerListenerResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
		baremetal.WorkRequestFailed,
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
		baremetal.WorkRequestFailed,
	}
}

func (s *LoadBalancerListenerResourceCrud) sslConfig() (sslConfig *baremetal.SSLConfiguration) {
	vs := s.D.Get("ssl_configuration").([]interface{})
	if len(vs) == 1 {
		sslConfig = new(baremetal.SSLConfiguration)
		v := vs[0].(map[string]interface{})
		sslConfig.CertificateName = v["certificate_name"].(string)
		sslConfig.VerifyDepth = v["verify_depth"].(int)
		sslConfig.VerifyPeerCertificate = v["verify_peer_certificate"].(bool)
		return sslConfig
	}

	return nil
}

func (s *LoadBalancerListenerResourceCrud) Create() (e error) {
	var workReqID string
	workReqID, e = s.Client.CreateListener(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("name").(string),
		s.D.Get("default_backend_set_name").(string),
		s.D.Get("protocol").(string),
		s.D.Get("port").(int),
		s.sslConfig(),
		nil, // neither OPCClientRequestID nor RetryToken is needed
	)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerListenerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID,name}
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.BaseCrud, s.WorkRequest)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}

	res, e := s.GetListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	if e == nil {
		s.Resource = res
	}
	return
}

// TODO: move this into the SDK, onto the client
func (s *LoadBalancerListenerResourceCrud) GetListener(loadBalancerID, name string) (*baremetal.Listener, error) {
	lb, err := s.Client.GetLoadBalancer(loadBalancerID, nil)
	if err != nil {
		return nil, err
	}
	l := lb.Listeners[name]
	if l.Name == name {
		return &l, nil
	}
	return nil, fmt.Errorf("Listener %s on load balancer %s does not exist", name, loadBalancerID)
}

func (s *LoadBalancerListenerResourceCrud) Update() (e error) {

	opts := &baremetal.UpdateLoadBalancerListenerOptions{
		DefaultBackendSetName: s.D.Get("default_backend_set_name").(string),
		Port:     s.D.Get("port").(int),
		Protocol: s.D.Get("protocol").(string),
	}
	opts.SSLConfig = s.sslConfig()
	log.Printf("SSL CONFIGURATION: %v", opts.SSLConfig)

	var workReqID string
	workReqID, e = s.Client.UpdateListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), opts)
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

func (s *LoadBalancerListenerResourceCrud) SetData() {
	if s.Resource == nil {
		return
	}
	s.D.Set("name", s.Resource.Name)
	s.D.Set("default_backend_set_name", s.Resource.DefaultBackendSetName)
	s.D.Set("port", s.Resource.Port)
	s.D.Set("protocol", s.Resource.Protocol)
	s.D.Set("ssl_configuration", s.Resource.SSLConfig)
}

func (s *LoadBalancerListenerResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), nil)
	if e != nil {
		return
	}
	s.D.SetId(workReqID)
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
