// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package network_load_balancer

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_network_load_balancer "github.com/oracle/oci-go-sdk/v65/networkloadbalancer"
)

func NetworkLoadBalancerBackendHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularNetworkLoadBalancerBackendHealth,
		Schema: map[string]*schema.Schema{
			"backend_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_load_balancer_id": {
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

func readSingularNetworkLoadBalancerBackendHealth(d *schema.ResourceData, m interface{}) error {
	sync := &NetworkLoadBalancerBackendHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).NetworkLoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type NetworkLoadBalancerBackendHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_network_load_balancer.NetworkLoadBalancerClient
	Res    *oci_network_load_balancer.GetBackendHealthResponse
}

func (s *NetworkLoadBalancerBackendHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *NetworkLoadBalancerBackendHealthDataSourceCrud) Get() error {
	request := oci_network_load_balancer.GetBackendHealthRequest{}

	if backendName, ok := s.D.GetOkExists("backend_name"); ok {
		tmp := backendName.(string)
		request.BackendName = &tmp
	}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if networkLoadBalancerId, ok := s.D.GetOkExists("network_load_balancer_id"); ok {
		tmp := networkLoadBalancerId.(string)
		request.NetworkLoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "network_load_balancer")

	response, err := s.Client.GetBackendHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *NetworkLoadBalancerBackendHealthDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("NetworkLoadBalancerBackendHealthDataSource-", NetworkLoadBalancerBackendHealthDataSource(), s.D))

	healthCheckResults := []interface{}{}
	for _, item := range s.Res.HealthCheckResults {
		healthCheckResults = append(healthCheckResults, NlbHealthCheckResultToMap(item))
	}
	s.D.Set("health_check_results", healthCheckResults)

	s.D.Set("status", s.Res.Status)

	return nil
}

func NlbHealthCheckResultToMap(obj oci_network_load_balancer.HealthCheckResult) map[string]interface{} {
	result := map[string]interface{}{}

	result["health_check_status"] = string(obj.HealthCheckStatus)

	if obj.Timestamp != nil {
		result["timestamp"] = obj.Timestamp.String()
	}

	return result
}
