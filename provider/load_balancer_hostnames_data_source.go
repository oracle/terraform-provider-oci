// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func HostnamesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readHostnames,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hostnames": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     HostnameResource(),
			},
		},
	}
}

func readHostnames(d *schema.ResourceData, m interface{}) error {
	sync := &HostnamesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type HostnamesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListHostnamesResponse
}

func (s *HostnamesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *HostnamesDataSourceCrud) Get() error {
	request := oci_load_balancer.ListHostnamesRequest{}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListHostnames(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *HostnamesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		hostname := map[string]interface{}{}

		if r.Hostname != nil {
			hostname["hostname"] = *r.Hostname
		}

		if r.Name != nil {
			hostname["name"] = *r.Name
		}

		resources = append(resources, hostname)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, HostnamesDataSource().Schema["hostnames"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("hostnames", resources); err != nil {
		panic(err)
	}

	return
}
