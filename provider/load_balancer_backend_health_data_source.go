// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackendHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBackendHealth,
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

func readSingularBackendHealth(d *schema.ResourceData, m interface{}) error {
	sync := &BackendHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type BackendHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetBackendHealthResponse
}

func (s *BackendHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BackendHealthDataSourceCrud) Get() error {
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetBackendHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BackendHealthDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	healthCheckResults := []interface{}{}
	for _, item := range s.Res.HealthCheckResults {
		healthCheckResults = append(healthCheckResults, HealthCheckResultToMap(item))
	}
	s.D.Set("health_check_results", healthCheckResults)

	s.D.Set("status", s.Res.Status)

	return
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
