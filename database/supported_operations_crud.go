// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package database

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type SupportedOperationDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListSupportedOperations
}

func (s *SupportedOperationDatasourceCrud) Get() (e error) {
	var list *baremetal.ListSupportedOperations
	if list, e = s.Client.ListSupportedOperations(); e != nil {
		return
	}

	s.Res = list
	return
}

func (s *SupportedOperationDatasourceCrud) SetData() {
	if s.Res != nil {
		resources := []map[string]interface{}{}
		s.D.SetId(time.Now().UTC().String())
		for _, v := range s.Res.SupportedOperations {
			res := map[string]interface{}{
				"id": v.ID,
			}
			resources = append(resources, res)
		}
		s.D.Set("supported_operations", resources)
	}
	return
}
