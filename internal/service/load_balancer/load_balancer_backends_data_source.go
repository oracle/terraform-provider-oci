// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package load_balancer

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/v56/loadbalancer"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func LoadBalancerBackendsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readLoadBalancerBackends,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"backendset_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backends": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     LoadBalancerBackendResource(),
			},
		},
	}
}

func readLoadBalancerBackends(d *schema.ResourceData, m interface{}) error {
	sync := &LoadBalancerBackendsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).LoadBalancerClient()

	return tfresource.ReadResource(sync)
}

type LoadBalancerBackendsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListBackendsResponse
}

func (s *LoadBalancerBackendsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *LoadBalancerBackendsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListBackendsRequest{}

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListBackends(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *LoadBalancerBackendsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("LoadBalancerBackendsDataSource-", LoadBalancerBackendsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		backend := map[string]interface{}{}

		if r.Backup != nil {
			backend["backup"] = *r.Backup
		}

		if r.Drain != nil {
			backend["drain"] = *r.Drain
		}

		if r.IpAddress != nil {
			backend["ip_address"] = *r.IpAddress
		}

		if r.Name != nil {
			backend["name"] = *r.Name
		}

		if r.Offline != nil {
			backend["offline"] = *r.Offline
		}

		if r.Port != nil {
			backend["port"] = *r.Port
		}

		if r.Weight != nil {
			backend["weight"] = *r.Weight
		}

		resources = append(resources, backend)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, LoadBalancerBackendsDataSource().Schema["backends"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backends", resources); err != nil {
		return err
	}

	return nil
}
