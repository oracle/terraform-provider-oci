// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

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
