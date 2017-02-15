// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package core

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
)

type DrgResourceCrud struct {
	crud.BaseCrud
	Res *baremetal.Drg
}

func (s *DrgResourceCrud) ID() string {
	return s.Res.ID
}

func (s *DrgResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceProvisioning}
}

func (s *DrgResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceAvailable}
}

func (s *DrgResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceTerminating}
}

func (s *DrgResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceTerminated}
}

func (s *DrgResourceCrud) State() string {
	return s.Res.State
}

func (s *DrgResourceCrud) Create() (e error) {
	opts := &baremetal.CreateOptions{}
	compartmentID := s.D.Get("compartment_id").(string)
	displayName, ok := s.D.GetOk("display_name")
	if ok {
		opts.DisplayName = displayName.(string)
	}

	s.Res, e = s.Client.CreateDrg(compartmentID, opts)

	return
}

func (s *DrgResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetDrg(s.D.Id())
	return
}

func (s *DrgResourceCrud) SetData() {
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("display_name", s.Res.DisplayName)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *DrgResourceCrud) Delete() (e error) {
	return s.Client.DeleteDrg(s.D.Id(), nil)
}
