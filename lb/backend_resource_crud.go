// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package lb

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

type LoadBalancerBackendResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.BackendSet
}

func (s *LoadBalancerBackendResourceCrud) ID() string {
	return s.Resource.Name
}

func (s *LoadBalancerBackendResourceCrud) CustomTimeout() time.Duration {
	return 15 * time.Minute
}

func (s *LoadBalancerBackendResourceCrud) CreatedPending() []string {
	return []string{
		baremetal.ResourceProvisioning,
		baremetal.ResourceStarting,
	}
}

func (s *LoadBalancerBackendResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceRunning}
}

func (s *LoadBalancerBackendResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *LoadBalancerBackendResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func resourceMapToMetadata(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func (s *LoadBalancerBackendResourceCrud) Create() (e error) {
	// TODO: populate
	backends := []baremetal.Backend{}
	healthChecker := baremetal.HealthChecker{}
	sslConfig := baremetal.SSLConfiguration{}

	// TODO: should LoadBalancerOptions be CreateOptions
	// opts := &baremetal.CreateOptions{}
	opts := &baremetal.LoadBalancerOptions{}
	// TODO: add LoadBalancerOptions.DisplayName
	// opts.DisplayName = s.D.Get("display_name").(string)

	s.Resource, e = s.Client.CreateBackendSet(
		s.D.Get("loadbalancer_id").(string),
		s.D.Get("name").(string),
		s.D.Get("policy").(string),
		backends,
		healthChecker,
		sslConfig,
		opts,
	)
	return
}

func NewBackend(v map[string]interface{}) baremetal.Backend {
	return baremetal.Backend{
		Name:    v["name"].(string),
		Backup:  v["backup"].(bool),
		Drain:   v["drain"].(bool),
		Offline: v["offline"].(bool),
		Port:    v["port"].(int),
		Weight:  v["weight"].(int),
	}
}

func NewBackendSet(vs []interface{}) baremetal.BackendSet {
	set := baremetal.BackendSet{}
	// TODO: validation for len(vs)
	if len(vs) == 1 {
		singleton := vs[0].(map[string]interface{})
		// policy {
		set.Policy = singleton["policy"].(string)
		// }

		// backend {
		if _backends := singleton["backend"]; _backends != nil {
			backends := _backends.([]interface{})
			set.Backends = make([]baremetal.Backend, len(backends))
			for i, b := range backends {
				set.Backends[i] = NewBackend(b.(map[string]interface{}))
			}
		}
		// }

		// health_checker {
		// TODO: set.HealthChecker
		// }

		// ssl_configuration {
		// TODO: set.SSLConfiguration
		// }
	}
	return set
}

func (s *LoadBalancerBackendResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetBackendSet(s.D.Get("loadbalancer_id").(string), s.D.Get("backendset_name").(string), s.D.Get("name").(string), nil)
	return
}

func (s *LoadBalancerBackendResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateLoadBalancerBackendSetOptions{}
	// TODO: add UpdateLoadBalancerBackendSetOptions.DisplayName field
	// opts.DisplayName = s.D.Get("display_name").(string)

	s.Resource, e = s.Client.UpdateBackendSet(s.D.Get("loadbalancer_id").(string), s.D.Id(), opts)

	return
}

func (s *LoadBalancerBackendResourceCrud) SetData() {
	s.D.Set("policy", s.Resource.Policy)
	s.D.Set("name", s.Resource.Name)
	// TODO: remaining attrs
}

func (s *LoadBalancerBackendResourceCrud) Delete() (e error) {
	return s.Client.DeleteBackendSet(s.D.Get("loadbalancer_id").(string), s.D.Get("name").(string), nil)
}
