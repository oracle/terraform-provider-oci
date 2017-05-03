// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/client"
	"github.com/oracle/terraform-provider-baremetal/crud"
)

func BackendDatasource() *schema.Resource {
	return &schema.Resource{
		Read: readBackends,
		Schema: map[string]*schema.Schema{
			"load_balancer_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backendset_name": {
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

func readBackends(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(client.BareMetalClient)
	sync := &BackendDatasourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

type BackendDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListBackends
}

func (s *BackendDatasourceCrud) Get() (e error) {
	lbID := s.D.Get("load_balancer_id").(string)
	backendSetName := s.D.Get("backendset_name").(string)
	s.Res, e = s.Client.ListBackends(lbID, backendSetName)
	return
}

func (s *BackendDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.Backends {
			res := map[string]interface{}{
				"name":       v.Name,
				"ip_address": v.IPAddress,
				"port":       v.Port,
				"backup":     v.Backup,
				"drain":      v.Drain,
				"offline":    v.Offline,
				"weight":     v.Weight,
			}
			resources = append(resources, res)
		}
		s.D.Set("backends", resources)
	}
	return
}
