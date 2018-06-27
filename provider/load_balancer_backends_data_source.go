// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_load_balancer "github.com/oracle/oci-go-sdk/loadbalancer"

	"github.com/oracle/terraform-provider-oci/crud"
)

func BackendsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readBackends,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
				Elem:     BackendResource(),
			},
		},
	}
}

func readBackends(d *schema.ResourceData, m interface{}) error {
	sync := &BackendsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).loadBalancerClient

	return crud.ReadResource(sync)
}

type BackendsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_load_balancer.LoadBalancerClient
	Res    *oci_load_balancer.ListBackendsResponse
}

func (s *BackendsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BackendsDataSourceCrud) Get() error {
	request := oci_load_balancer.ListBackendsRequest{}

	if backendsetName, ok := s.D.GetOkExists("backendset_name"); ok {
		tmp := backendsetName.(string)
		request.BackendSetName = &tmp
	}

	if loadBalancerId, ok := s.D.GetOkExists("load_balancer_id"); ok {
		tmp := loadBalancerId.(string)
		request.LoadBalancerId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "load_balancer")

	response, err := s.Client.ListBackends(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BackendsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
		resources = ApplyFilters(f.(*schema.Set), resources, BackendsDataSource().Schema["backends"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("backends", resources); err != nil {
		panic(err)
	}

	return
}
