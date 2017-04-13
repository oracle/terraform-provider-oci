// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"time"

	"github.com/MustWin/baremetal-sdk-go"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

type SwiftPasswordDatasourceCrud struct {
	crud.BaseCrud
	Res *baremetal.ListSwiftPasswords
}

func (s *SwiftPasswordDatasourceCrud) Get() (e error) {
	userID := s.D.Get("user_id").(string)

	s.Res, e = s.Client.ListSwiftPasswords(userID)
	return
}

func (s *SwiftPasswordDatasourceCrud) SetData() {
	if s.Res != nil {
		s.D.SetId(time.Now().UTC().String())
		resources := []map[string]interface{}{}
		for _, v := range s.Res.SwiftPasswords {
			res := map[string]interface{}{
				"id":             v.ID,
				"user_id":        v.UserID,
				"description":    v.Description,
				"state":          v.State,
				"inactive_state": v.InactiveStatus,
				"time_created":   v.TimeCreated.String(),
				"expires_on":     v.ExpiresOn.String(),
			}
			resources = append(resources, res)
		}
		if err := s.D.Set("passwords", resources); err != nil {
			panic(err)
		}
	}
	return
}
