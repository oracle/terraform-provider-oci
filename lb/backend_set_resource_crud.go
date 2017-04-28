// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type LoadBalancerBackendSetResourceCrud struct {
	crud.BaseCrud
	WorkRequest  *baremetal.WorkRequest
	Resource     *baremetal.BackendSet
	ResourceName string
}

func (s *LoadBalancerBackendSetResourceCrud) ID() string {
	if s.Resource != nil && s.Resource.Name != "" {
		return s.Resource.Name
	}
	if s.WorkRequest != nil && s.WorkRequest.State == baremetal.WorkRequestSucceeded {
		return s.ResourceName
	}
	return s.D.Get("name").(string)
}

// RefreshWorkRequest returns the last updated workRequest
func (s *LoadBalancerBackendSetResourceCrud) RefreshWorkRequest() (*baremetal.WorkRequest, error) {
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

func (s *LoadBalancerBackendSetResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) CreatedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedPending() []string {
	return []string{
		baremetal.ResourceWaitingForWorkRequest,
		baremetal.WorkRequestInProgress,
		baremetal.WorkRequestAccepted,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) DeletedTarget() []string {
	return []string{
		baremetal.ResourceSucceededWorkRequest,
		baremetal.WorkRequestSucceeded,
	}
}

func (s *LoadBalancerBackendSetResourceCrud) Create() (e error) {
	workReqID, e := s.Client.CreateBackendSet(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("name").(string),
		s.D.Get("policy").(string),
		s.backends(),
		s.healthChecker(),
		s.sslConfig(),
		nil,
	)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetBackendSet(
		s.D.Get("load_balancer_id").(string),
		s.D.Get("name").(string),
		nil,
	)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) Update() (e error) {
	healthChecker := baremetal.HealthChecker{
		IntervalInMS:      s.D.Get("health_checker.interval_ms").(int),
		Port:              s.D.Get("health_checker.port").(int),
		Protocol:          s.D.Get("health_checker.protocol").(string),
		ResponseBodyRegex: s.D.Get("health_checker.response_body_regex").(string),
	}
	sslConfig := baremetal.SSLConfiguration{
		CertificateName:       s.D.Get("ssl_configuration.certificate_name").(string),
		VerifyDepth:           s.D.Get("ssl_configuration.verify_depth").(int),
		VerifyPeerCertificate: s.D.Get("ssl_configuration.verify_peer_certificate").(bool),
	}
	opts := &baremetal.UpdateLoadBalancerBackendSetOptions{
		Policy:        s.D.Get("policy").(string),
		SSLConfig:     sslConfig,
		HealthChecker: healthChecker,
	}

	var workReqID string
	workReqID, e = s.Client.UpdateBackendSet(s.D.Get("load_balancer_id").(string), s.D.Id(), opts)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) SetData() {
	if s.Resource == nil {
		panic("BackendSet Resource is nil, cannot SetData")
	}
	s.D.Set("policy", s.Resource.Policy)
	s.D.Set("name", s.Resource.Name)
	s.D.Set("health_checker", map[string]interface{}{
		"interval_ms":         s.Resource.HealthChecker.IntervalInMS,
		"port":                s.Resource.HealthChecker.Port,
		"protocol":            s.Resource.HealthChecker.Protocol,
		"response_body_regex": s.Resource.HealthChecker.ResponseBodyRegex,
	})
	s.D.Set("ssl_configuration", map[string]interface{}{
		"certificate_name":        s.Resource.SSLConfig.CertificateName,
		"verify_depth":            s.Resource.SSLConfig.VerifyDepth,
		"verify_peer_certificate": s.Resource.SSLConfig.VerifyPeerCertificate,
	})
	backends := make([]map[string]interface{}, len(s.Resource.Backends))
	for i, v := range s.Resource.Backends {
		backends[i] = map[string]interface{}{
			"backup":     v.Backup,
			"drain":      v.Drain,
			"ip_address": v.IPAddress,
			"name":       v.Name,
			"offline":    v.Offline,
			"port":       v.Port,
			"weight":     v.Weight,
		}
	}
	s.D.Set("backend", backends)
}

func (s *LoadBalancerBackendSetResourceCrud) Delete() (e error) {
	var workReqID string
	workReqID, e = s.Client.DeleteBackendSet(s.D.Get("load_balancer_id").(string), s.D.Get("name").(string), nil)
	if e != nil {
		return
	}
	s.WorkRequest, e = s.Client.GetWorkRequest(workReqID, nil)
	return
}

func (s *LoadBalancerBackendSetResourceCrud) sslConfig() (sslConfig *baremetal.SSLConfiguration) {
	sslConfig = new(baremetal.SSLConfiguration)
	vs := s.D.Get("ssl_configuration").([]interface{})
	if len(vs) == 1 {
		v := vs[0].(map[string]interface{})
		sslConfig.CertificateName = v["certificate_name"].(string)
		sslConfig.VerifyDepth = v["verify_depth"].(int)
		sslConfig.VerifyPeerCertificate = v["verify_peer_certificate"].(bool)
	}

	return
}

func (s *LoadBalancerBackendSetResourceCrud) healthChecker() *baremetal.HealthChecker {
	healthChecker := new(baremetal.HealthChecker)
	vs := s.D.Get("health_checker").([]interface{})
	if len(vs) == 1 {
		v := vs[0].(map[string]interface{})
		healthChecker.IntervalInMS = v["interval_ms"].(int)
		healthChecker.Port = v["port"].(int)
		healthChecker.Protocol = v["protocol"].(string)
		healthChecker.ResponseBodyRegex = v["response_body_regex"].(string)
	}
	return healthChecker
}
func (s *LoadBalancerBackendSetResourceCrud) backends() []baremetal.Backend {
	vs := s.D.Get("backend").([]interface{})
	backends := make([]baremetal.Backend, len(vs))
	for i := range vs {
		v := vs[i].(map[string]interface{})
		backends[i] = baremetal.Backend{
			Backup:    v["backup"].(bool),
			Drain:     v["drain"].(bool),
			IPAddress: v["ip_address"].(string),
			Name:      v["name"].(string),
			Offline:   v["offline"].(bool),
			Port:      v["port"].(int),
			Weight:    v["weight"].(int),
		}
	}
	return backends
}
