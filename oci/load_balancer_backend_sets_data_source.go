// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v47/loadbalancer"
)

func init() {
	RegisterDatasource("oci_load_balancer_backend_sets", LoadBalancerBackendSetsDataSource())
}

func LoadBalancerBackendSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerBackendSets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backendsets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(LoadBalancerBackendSetResource()),
			},
		},
	}
}

func readLoadBalancerBackendSets(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient()

	return ReadResource(sync)
}

type LoadBalancerBackendSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListBackendSetsResponse
}

func (s *LoadBalancerBackendSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerBackendSetsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListBackendSetsRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListBackendSets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerBackendSetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("LoadBalancerBackendSetsDataSource-", LoadBalancerBackendSetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backendSet := map[string]interface{}{}

		backend := []interface{}{}
		for _, item := range r.Backends {
			backend = append(backend, BackendToMap(item))
		}
		backendSet["backend"] = backend

		if r.HealthChecker != nil {
			backendSet["health_checker"] = []interface{}{HealthCheckerToMap(r.HealthChecker)}
		} else {
			backendSet["health_checker"] = nil
		}

		if r.LbCookieSessionPersistenceConfiguration != nil {
			backendSet["lb_cookie_session_persistence_configuration"] = []interface{}{LBCookieSessionPersistenceConfigurationDetailsToMap(r.LbCookieSessionPersistenceConfiguration)}
		} else {
			backendSet["lb_cookie_session_persistence_configuration"] = nil
		}

		if r.Name != nil {
			backendSet["name"] = *r.Name
		}

		if r.Policy != nil {
			backendSet["policy"] = *r.Policy
		}

		if r.SessionPersistenceConfiguration != nil {
			backendSet["session_persistence_configuration"] = []interface{}{SessionPersistenceConfigurationDetailsToMap(r.SessionPersistenceConfiguration)}
		} else {
			backendSet["session_persistence_configuration"] = nil
		}

		if r.SslConfiguration != nil {
			backendSet["ssl_configuration"] = []interface{}{SSLConfigurationToMap(r.SslConfiguration)}
		} else {
			backendSet["ssl_configuration"] = nil
		}

		resources = append(resources, backendSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, LoadBalancerBackendSetsDataSource().Schema["backendsets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backendsets", resources); err != nil {
		return err
	}

	return nil
}
