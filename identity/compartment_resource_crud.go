// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package identity

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/client"
	"github.com/MustWin/terraform-Oracle-BareMetal-Provider/crud"
	"github.com/hashicorp/terraform/helper/schema"
)

type CompartmentResourceCrud struct {
	*crud.IdentitySync
	D      *schema.ResourceData
	Client client.BareMetalClient
	Res    *baremetal.Compartment
}

func (s *CompartmentResourceCrud) ID() string {
	return s.Res.ID
}

func (s *CompartmentResourceCrud) State() string {
	return s.Res.State
}

func (s *CompartmentResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *CompartmentResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *CompartmentResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *CompartmentResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *CompartmentResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateCompartment(name, description, nil)
	return
}

func (s *CompartmentResourceCrud) Get() (e error) {
	s.Res, e = s.Client.GetCompartment(s.D.Id())
	return
}

func (s *CompartmentResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateIdentityOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}
	s.Res, e = s.Client.UpdateCompartment(s.D.Id(), opts)
	return
}

func (s *CompartmentResourceCrud) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
