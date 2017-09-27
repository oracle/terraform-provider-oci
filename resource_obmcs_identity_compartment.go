// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"strings"

	"github.com/oracle/bmcs-go-sdk"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
	"github.com/oracle/terraform-provider-oci/options"
)

// ResourceIdentityCompartment exposes an IdentityCompartment Resource
func CompartmentResource() *schema.Resource {
	compartmentSchema := map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
			ForceNew: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"compartment_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"inactive_state": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"time_created": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"time_modified": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createCompartment,
		Read:     readCompartment,
		Update:   updateCompartment,
		Delete:   deleteCompartment,
		Schema:   compartmentSchema,
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

func listAllCompartments(s *CompartmentResourceCrud) (result *baremetal.ListCompartments, e error) {
	opts := &baremetal.ListOptions{}
	options.SetListOptions(s.D, opts)

	result = &baremetal.ListCompartments{Compartments: []baremetal.Compartment{}}

	for {
		var page *baremetal.ListCompartments
		if page, e = s.Client.ListCompartments(opts); e != nil {
			break
		}

		result.Compartments = append(result.Compartments, page.Compartments...)

		if hasNexPage := options.SetNextPageOption(page.NextPage, &opts.PageListOptions); !hasNexPage {
			break
		}
	}
	return
}

func (s *CompartmentResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	s.Res, e = s.Client.CreateCompartment(name, description, nil)
	// Compartments can't be destroyed, so we shouldn't complain about them being created.
	if e != nil && strings.Contains(e.Error(), "already exists") {
		e = nil
		list, err := listAllCompartments(s)
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
	res, e := s.Client.GetCompartment(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *CompartmentResourceCrud) Update() (e error) {
	opts := &baremetal.UpdateCompartmentOptions{}
	if name, ok := s.D.GetOk("name"); ok {
		opts.Name = name.(string)
	}
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
