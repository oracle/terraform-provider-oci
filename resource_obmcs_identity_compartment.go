// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"strings"

	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-baremetal/crud"
)

// ResourceIdentityCompartment exposes an IdentityCompartment Resource
func CompartmentResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createCompartment,
		Read:     readCompartment,
		Update:   updateCompartment,
		Delete:   deleteCompartment,
		Schema:   baseIdentitySchemaWithID,
	}
}

func createCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updateCompartment(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &CompartmentResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deleteCompartment(d *schema.ResourceData, m interface{}) (e error) {
	sync := &CompartmentResourceCrud{}
	sync.D = d
	return crud.DeleteResource(d, sync)
}

type CompartmentResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.Compartment
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

func (s *CompartmentResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateCompartment(name, description, nil)
	// Compartments can't be destroyed, so we shouldn't complain about them being created.
	if e != nil && strings.Contains(e.Error(), "already exists") {
		e = nil
		list, err := s.Client.ListCompartments(nil) // TODO: This won't paginate...
		if err != nil {
			e = err
			return
		}
		for _, compartment := range list.Compartments {
			if compartment.Name == name {
				s.Res = &compartment
				break
			}
		}
	}
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

func (s *CompartmentResourceCrud) Delete() (e error) {
	// Compartments cannot be deleted. Just pretend it worked.
	e = nil
	return
}

func (s *CompartmentResourceCrud) SetData() {
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}
