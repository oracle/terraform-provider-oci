// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackendSetHealthDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularBackendSetHealth,
		Schema: map[string]*schema.Schema{
			"backend_set_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"critical_state_backend_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"total_backend_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unknown_state_backend_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"warning_state_backend_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularBackendSetHealth(d *schema.ResourceData, m interface{}) error {
	sync := &BackendSetHealthDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type BackendSetHealthDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.GetBackendSetHealthResponse
}

func (s *BackendSetHealthDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BackendSetHealthDataSourceCrud) Get() error {
	request := oci_load_balancer.GetBackendSetHealthRequest{}

	if backendSetName, ok := s.D.GetOkExists("backend_set_name"); ok {
		tmp := backendSetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.GetBackendSetHealth(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BackendSetHealthDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())

	s.D.Set("critical_state_backend_names", s.Res.CriticalStateBackendNames)

	s.D.Set("status", s.Res.Status)

	if s.Res.TotalBackendCount != nil {
		s.D.Set("total_backend_count", *s.Res.TotalBackendCount)
	}

	s.D.Set("unknown_state_backend_names", s.Res.UnknownStateBackendNames)

	s.D.Set("warning_state_backend_names", s.Res.WarningStateBackendNames)

	return
}
