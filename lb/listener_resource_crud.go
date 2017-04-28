// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"fmt"
	"log"
	"strings"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type LoadBalancerListenerResourceCrud struct {
	crud.BaseCrud
	WorkRequest *baremetal.WorkRequest
	Resource    *baremetal.Listener
}

// ID uniquely identifies the listener and its parent load balancer
func (s *LoadBalancerListenerResourceCrud) ID() string {
	log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.ID: WorkRequest: %#v", s.WorkRequest)
	if s.WorkRequest != nil && s.WorkRequest.State != baremetal.WorkRequestSucceeded {
		log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.ID: WorkRequest.ID: %s", s.WorkRequest.ID)
		return s.WorkRequest.ID
	}

	id := fmt.Sprintf("%s/listener/%s", s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.ID: %#v", id)
	return id
}

// RefreshWorkRequest returns the last updated workRequest
func (s *LoadBalancerListenerResourceCrud) RefreshWorkRequest() (*baremetal.WorkRequest, error) {
	if s.WorkRequest == nil {
		return nil, nil
	}
	wr, err := s.Client.GetWorkRequest(s.WorkRequest.ID, nil)
	if err != nil {
		return nil, err
	}
	s.WorkRequest = wr
	return wr, nil
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
	}
}

func (s *LoadBalancerListenerResourceCrud) sslConfig() (sslConfig *baremetal.SSLConfiguration) {
	sslConfig = &baremetal.SSLConfiguration{}
	vs := s.D.Get("ssl_configuration").([]interface{})
	if len(vs) == 1 {
		v := vs[0].(map[string]interface{})
		sslConfig.CertificateName = v["certificate_name"].(string)
		sslConfig.VerifyDepth = v["verify_depth"].(int)
		sslConfig.VerifyPeerCertificate = v["verify_peer_certificate"].(bool)
	}

	return
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
	id := s.D.Id()
	log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.Get: ID: %#v", id)

	// NOTE: if the id is for a work request, refresh its state. then refresh the listener.
	if strings.HasPrefix(id, "ocid1.loadbalancerworkrequest.") {
		log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.Get: ID is for WorkRequest, refreshing")
		s.WorkRequest, e = s.Client.GetWorkRequest(id, nil)
		log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.Get: WorkRequest: %#v", s.WorkRequest)
		e = s.D.Set("state", s.WorkRequest.State)
		if s.WorkRequest.State == baremetal.WorkRequestSucceeded {
			// set state for the next phase
			// unset work request on success
			s.WorkRequest = nil
		} else {
			log.Printf("[DEBUG] lb.LoadBalancerListenerResourceCrud.Get: Work Request.State: %#v != Succeeded", s.WorkRequest.State)
			s.D.Set("state", s.WorkRequest.State)
			// We do not have a completed work request, so we short-circuit out
			return
		}
	}

	l, e := s.GetListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	if e != nil {
		return e
	}
	s.Resource = l
	return nil
}

// TODO: move this into the SDK, onto the client
func (s *LoadBalancerListenerResourceCrud) GetListener(loadBalancerID, name string) (*baremetal.Listener, error) {
	log.Printf("[DEBUG] lb.GetListener: loadBalancerID: %#v, name: %#v", loadBalancerID, name)

	// API does not have GetListener(loadBalancerID, name), query all and filter
	lb, err := s.Client.GetLoadBalancer(loadBalancerID, nil)
	if err != nil {
		return nil, err
	}
	l := lb.Listeners[name]
	log.Printf("[DEBUG] lb.GetListener: LoadBalancer: %#v", lb)
	if l.Name == name {
		return &l, nil
	}
	return nil, fmt.Errorf("No listener found with load_balancer: %v, name: %v; got %#v", loadBalancerID, name, l)
}

func (s *LoadBalancerListenerResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateLoadBalancerListenerOptions{
		DefaultBackendSetName: s.D.Get("default_backend_set_name").(string),
		Port:      s.D.Get("port").(int),
		Protocol:  s.D.Get("protocol").(string),
		SSLConfig: *s.sslConfig(),
	}

	var workReqID string
	workReqID, e = s.Client.UpdateListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerListenerResourceCrud) SetData() {
	// load_balancer_id is not returned, but we should be able to trust it
	// s.D.Set("load_balancer_id", s.Resource.LoadBalancerID)

	if s.Resource == nil {
		panic("Listener Resource is nil, cannot SetData")
	}
	s.D.Set("name", s.Resource.Name)
	s.D.Set("default_backend_set_name", s.Resource.DefaultBackendSetName)
	s.D.Set("port", s.Resource.Port)
	s.D.Set("protocol", s.Resource.Protocol)
	// TODO: verify testing
	s.D.Set("ssl_configuration", s.Resource.SSLConfig)
}

func (s *LoadBalancerListenerResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), nil)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}
