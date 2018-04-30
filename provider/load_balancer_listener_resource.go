// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"
)

func ListenerResource() *schema.Resource {
	return &schema.Resource{
		Create: createListener,
		Read:   readListener,
		Update: updateListener,
		Delete: deleteListener,
		Schema: map[string]*schema.Schema{
			// Required
			"default_backend_set_name": {
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
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional
			"connection_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"idle_timeout_in_seconds": {
							Type:     schema.TypeInt,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"hostname_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"path_route_set_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_configuration": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"certificate_name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
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

						// Computed
					},
				},
			},

			// Computed
			// internal for work request access
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.CreateResource(d, sync)
}

func readListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	return crud.ReadResource(sync)
}

func updateListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.UpdateResource(d, sync)
}

func deleteListener(d *schema.ResourceData, m interface{}) error {
	sync := &ListenerResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type ListenerResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_load_balancer.LoadBalancerClient
	Res                    *oci_load_balancer.Listener
	DisableNotFoundRetries bool
	WorkRequest            *oci_load_balancer.WorkRequest
}

func (s *ListenerResourceCrud) ID() string {
	id, workSuccess := crud.LoadBalancerResourceID(s.Res, s.WorkRequest)
	if id != nil {
		return *id
	}
	if workSuccess {
		return s.D.Get("name").(string)
	}
	return ""
}

func (s *ListenerResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *ListenerResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *ListenerResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateInProgress),
		string(oci_load_balancer.WorkRequestLifecycleStateAccepted),
	}
}

func (s *ListenerResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_load_balancer.WorkRequestLifecycleStateSucceeded),
		string(oci_load_balancer.WorkRequestLifecycleStateFailed),
	}
}

func (s *ListenerResourceCrud) Create() error {
	request := oci_load_balancer.CreateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToConnectionConfiguration(tmpList[0].(map[string]interface{}))
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	request.HostnameNames = []string{}
	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.HostnameNames = tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSSLConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.CreateListener(context.Background(), request)
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

func (s *ListenerResourceCrud) Get() (e error) {
	// key: {workRequestID} || {loadBalancerID,name}
	_, stillWorking, err := crud.LoadBalancerResourceGet(s.Client, s.D, s.WorkRequest, getRetryPolicy(s.DisableNotFoundRetries, "load_balancer"))
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

func (s *ListenerResourceCrud) GetListener(loadBalancerID, name string) (*oci_load_balancer.Listener, error) {
	request := oci_load_balancer.GetLoadBalancerRequest{}
	request.LoadBalancerId = &loadBalancerID
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.GetLoadBalancer(context.Background(), request)
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

func (s *ListenerResourceCrud) Update() error {
	request := oci_load_balancer.UpdateListenerRequest{}

	if connectionConfiguration, ok := s.D.GetOkExists("connection_configuration"); ok {
		if tmpList := connectionConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToConnectionConfiguration(tmpList[0].(map[string]interface{}))
			request.ConnectionConfiguration = &tmp
		}
	}

	if defaultBackendSetName, ok := s.D.GetOkExists("default_backend_set_name"); ok {
		tmp := defaultBackendSetName.(string)
		request.DefaultBackendSetName = &tmp
	}

	request.HostnameNames = []string{}
	if hostnameNames, ok := s.D.GetOkExists("hostname_names"); ok {
		interfaces := hostnameNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i, toBeConverted := range interfaces {
			tmp[i] = toBeConverted.(string)
		}
		request.HostnameNames = tmp
	}
	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}

	if pathRouteSetName, ok := s.D.GetOkExists("path_route_set_name"); ok {
		tmp := pathRouteSetName.(string)
		request.PathRouteSetName = &tmp
	}

	if port, ok := s.D.GetOkExists("port"); ok {
		tmp := port.(int)
		request.Port = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		tmp := protocol.(string)
		request.Protocol = &tmp
	}

	if sslConfiguration, ok := s.D.GetOkExists("ssl_configuration"); ok {
		if tmpList := sslConfiguration.([]interface{}); len(tmpList) > 0 {
			tmp := mapToSSLConfigurationDetails(tmpList[0].(map[string]interface{}))
			request.SslConfiguration = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.UpdateListener(context.Background(), request)
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

func (s *ListenerResourceCrud) Delete() error {
	if strings.Contains(s.D.Id(), "ocid1.loadbalancerworkrequest") {
		return nil
	}
	request := oci_load_balancer.DeleteListenerRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.ListenerName = &tmp
	}
	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "load_balancer")

	response, err := s.Client.DeleteListener(context.Background(), request)

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

func (s *ListenerResourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	if s.Res.ConnectionConfiguration != nil {
		s.D.Set("connection_configuration", []interface{}{ConnectionConfigurationToMap(s.Res.ConnectionConfiguration)})
	} else {
		s.D.Set("connection_configuration", []interface{}{})
	}
	if s.Res.DefaultBackendSetName != nil {
		s.D.Set("default_backend_set_name", *s.Res.DefaultBackendSetName)
	}
	if s.Res.HostnameNames != nil {
		s.D.Set("hostname_names", s.Res.HostnameNames)
	}
	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}
	if s.Res.PathRouteSetName != nil {
		s.D.Set("path_route_set_name", *s.Res.PathRouteSetName)
	}
	if s.Res.Port != nil {
		s.D.Set("port", *s.Res.Port)
	}
	if s.Res.Protocol != nil {
		s.D.Set("protocol", *s.Res.Protocol)
	}
	if s.Res.SslConfiguration != nil {
		s.D.Set("ssl_configuration", []interface{}{SSLConfigurationToMap(s.Res.SslConfiguration)})
	} else {
		s.D.Set("ssl_configuration", []interface{}{})
	}
}

func mapToConnectionConfiguration(raw map[string]interface{}) oci_load_balancer.ConnectionConfiguration {
	result := oci_load_balancer.ConnectionConfiguration{}

	if idleTimeoutInSeconds, ok := raw["idle_timeout_in_seconds"]; ok {
		tmp := idleTimeoutInSeconds.(int)
		result.IdleTimeout = &tmp
	}

	return result
}

func ConnectionConfigurationToMap(obj *oci_load_balancer.ConnectionConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IdleTimeout != nil {
		result["idle_timeout_in_seconds"] = int(*obj.IdleTimeout)
	}

	return result
}
