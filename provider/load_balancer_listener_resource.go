// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func ListenerResource() *schema.Resource {
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
			"ssl_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"verify_depth": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  5,
						},
						"verify_peer_certificate": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
					},
				},
			},
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
	sync.Client = m.(*OracleClients).loadBalancerClient
	return crud.CreateResource(d, sync)
}

func readLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	return crud.ReadResource(sync)
}

func updateLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	return crud.UpdateResource(d, sync)
}

func deleteLoadBalancerListener(d *schema.ResourceData, m interface{}) (e error) {
	sync := &LoadBalancerListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true
	return crud.DeleteResource(d, sync)
}

type LoadBalancerListenerResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	WorkRequest            *oci_load_balancer.WorkRequest
	Res                    *oci_load_balancer.Listener
	DisableNotFoundRetries bool
}

// ID uniquely identifies the listener and its parent load balancer
func (s *LoadBalancerListenerResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
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
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerListenerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *LoadBalancerListenerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *LoadBalancerListenerResourceCrud) sslConfig() (sslConfig *oci_load_balancer.SslConfigurationDetails) {
	vs := s.D.Get("ssl_configuration").([]interface{})
	if len(vs) == 1 {
		sslConfig = new(oci_load_balancer.SslConfigurationDetails)
		v := vs[0].(map[string]interface{})
		certificateNameStr := v["certificate_name"].(string)
		sslConfig.CertificateName = &certificateNameStr
		verifyDepthInt := v["verify_depth"].(int)
		sslConfig.VerifyDepth = &verifyDepthInt
		verifyPeerCertificateBool := v["verify_peer_certificate"].(bool)
		sslConfig.VerifyPeerCertificate = &verifyPeerCertificateBool
		return sslConfig
	}

	return nil
}

func (s *LoadBalancerListenerResourceCrud) Create() (e error) {
	request := oci_load_balancer.CreateListenerRequest{}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	request.SslConfiguration = s.sslConfig()

	response, err := s.Client.CreateListener(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *LoadBalancerListenerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID,name}
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	if stillWorking {
		return nil
	}

	res, e := s.GetListener(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string))
	if e == nil {
		s.Res = res
	}
	return
}

func (s *LoadBalancerListenerResourceCrud) GetListener(loadBalancerID, name string) (*oci_load_balancer.Listener, error) {
	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerID
	response, err := s.Client.GetLoadBalancer(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return nil, err
	}
	lb := &response.LoadBalancer
	if lb != nil && lb.Listeners != nil {
		if l, ok := lb.Listeners[name]; ok {
			if l.Name != nil && *l.Name == name {
				return &l, nil
			}
		}
	}
	return nil, fmt.Errorf("Listener %s on load balancer %s does not exist", name, loadBalancerID)
}

func (s *LoadBalancerListenerResourceCrud) Update() (e error) {
	request := oci_load_balancer.UpdateListenerRequest{}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	request.SslConfiguration = s.sslConfig()

	response, err := s.Client.UpdateListener(context.Background(), request, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return
	}
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	err = crud.LoadBalancerWaitForWorkRequest(s.Client, s.D, s.WorkRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}

	return s.Get()
}

func (s *LoadBalancerListenerResourceCrud) Delete() (e error) {
	request := oci_load_balancer.DeleteListenerRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}
	response, e := s.Client.DeleteListener(context.Background(), request)
	workReqID := response.OpcWorkRequestId
	getWorkRequestRequest := oci_load_balancer.GetWorkRequestRequest{}
	getWorkRequestRequest.WorkRequestId = workReqID
	workRequestResponse, err := s.Client.GetWorkRequest(context.Background(), getWorkRequestRequest, getRetryOptions(s.DisableNotFoundRetries, "load_balancer")...)
	if err != nil {
		return err
	}
	s.WorkRequest = &workRequestResponse.WorkRequest
	return nil
}

func (s *LoadBalancerListenerResourceCrud) SetData() {
	if s.Res == nil {
		return
	}
	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
	}
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}
	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}
	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}
	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", *s.Res.SslConfiguration)
	}
}
