// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v58/loadbalancer"
)

func LoadBalancerBackendHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularLoadBalancerBackendHealth,
		Schema: map[string]*schema.Schema{
			"backend_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"health_check_results": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"health_check_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subnet_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"timestamp": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularLoadBalancerBackendHealth(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerBackendHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetBackendHealthResponse
}

func (s *LoadBalancerBackendHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerBackendHealthDataSourceCrud) Get() error {
	request := oci_load_balancer.GetBackendHealthRequest{}

	if backendName, ok := s.D.GetOkExists("backend_name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetBackendHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerBackendHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerBackendHealthDataSource-", LoadBalancerBackendHealthDataSource(), s.D))

	healthCheckResults := []interface{}{}
	for _, item := range s.Res.HealthCheckResults {
		healthCheckResults = append(healthCheckResults, HealthCheckResultToMap(item))
	}
	s.D.Set("health_check_results", healthCheckResults)

	s.D.Set("status", s.Res.Status)

	return nil
}

func HealthCheckResultToMap(obj oci_load_balancer.HealthCheckResult) map[string]interface{} {
	result := map[string]interface{}{}

	result["health_check_status"] = string(obj.HealthCheckStatus)

	if obj.SourceIpAddress != nil {
		result["source_ip_address"] = string(*obj.SourceIpAddress)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	return result
}
