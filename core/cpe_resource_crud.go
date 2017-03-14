// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
)

type CpeResourceCrud struct {
	crud.BaseCrud
	Resource *baremetal.Cpe
}

func (s *CpeResourceCrud) ID() string {
	return s.Resource.ID
}

func (s *CpeResourceCrud) Create() (e error) {
	compartmentID := s.D.Get("compartment_id").(string)
	ipAddress := s.D.Get("ip_address").(string)

	opts := &baremetal.CreateOptions{}
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.CreateCpe(compartmentID, ipAddress, opts)
	return
}

func (s *CpeResourceCrud) Get() (e error) {
	s.Resource, e = s.Client.GetCpe(s.D.Id())
	return
}

func (s *CpeResourceCrud) Update() (e error) {
	opts := &baremetal.IfMatchDisplayNameOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Resource, e = s.Client.UpdateCpe(compartmentID, opts)
	return
}

func (s *CpeResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Resource.CompartmentID)
	s.D.Set("display_name", s.Resource.DisplayName)
	s.D.Set("ip_address", s.Resource.IPAddress)
	s.D.Set("time_created", s.Resource.TimeCreated.String())
}

func (s *CpeResourceCrud) Delete() (e error) {
	return s.Client.DeleteCpe(s.D.Id(), nil)
}
