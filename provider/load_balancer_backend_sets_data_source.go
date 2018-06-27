// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackendSetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBackendSets,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backendsets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     BackendSetResource(),
			},
		},
	}
}

func readBackendSets(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type BackendSetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListBackendSetsResponse
}

func (s *BackendSetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BackendSetsDataSourceCrud) Get() error {
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

func (s *BackendSetsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
			backendSet["session_persistence_configuration"] = []interface{}{}
		}

		if r.SslConfiguration != nil {
			backendSet["ssl_configuration"] = []interface{}{SSLConfigurationToMap(r.SslConfiguration)}
		} else {
			backendSet["session_persistence_configuration"] = []interface{}{}
		}

		resources = append(resources, backendSet)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, BackendSetsDataSource().Schema["backendsets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backendsets", resources); err != nil {
		panic(err)
	}

	return
}
