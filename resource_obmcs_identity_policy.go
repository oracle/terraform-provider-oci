// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package main

import (
	"github.com/MustWin/baremetal-sdk-go"
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"
)

func PolicyResource() *schema.Resource {
	policySchema := make(map[string]*schema.Schema)

	for key, value := range identitySchemaWithID {
		policySchema[key] = value
	}

	policySchema["statements"] = &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
	policySchema["inactive_state"] = &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	}
	policySchema["version_date"] = &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	}

	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createPolicy,
		Read:     readPolicy,
		Update:   updatePolicy,
		Delete:   deletePolicy,
		Schema:   policySchema,
	}
}

func createPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.CreateResource(d, sync)
}

func readPolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.ReadResource(sync)
}

func updatePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client
	return crud.UpdateResource(d, sync)
}

func deletePolicy(d *schema.ResourceData, m interface{}) (e error) {
	client := m.(*baremetal.Client)
	sync := &PolicyResourceCrud{}
	sync.D = d
	sync.Client = client
	return sync.Delete()
}

type PolicyResourceCrud struct {
	*crud.IdentitySync
	crud.BaseCrud
	Res *baremetal.Policy
}

func (s *PolicyResourceCrud) ID() string {
	return s.Res.ID
}

func (s *PolicyResourceCrud) State() string {
	return s.Res.State
}

func (s *PolicyResourceCrud) CreatedPending() []string {
	return []string{baremetal.ResourceCreating}
}

func (s *PolicyResourceCrud) CreatedTarget() []string {
	return []string{baremetal.ResourceActive}
}

func (s *PolicyResourceCrud) DeletedPending() []string {
	return []string{baremetal.ResourceDeleting}
}

func (s *PolicyResourceCrud) DeletedTarget() []string {
	return []string{baremetal.ResourceDeleted}
}

func (s *PolicyResourceCrud) toStringArray(vals interface{}) []string {
	arr := vals.([]interface{})
	result := []string{}
	for _, val := range arr {
		result = append(result, val.(string))
	}
	return result
}

func (s *PolicyResourceCrud) Create() (e error) {
	name := s.D.Get("name").(string)
	description := s.D.Get("description").(string)
	compartmentID := s.D.Get("compartment_id").(string)
	statements := s.toStringArray(s.D.Get("statements"))

	s.Res, e = s.Client.CreatePolicy(name, description, compartmentID, statements, nil)
	return
}

func (s *PolicyResourceCrud) Get() (e error) {
	res, e := s.Client.GetPolicy(s.D.Id())
	if e == nil {
		s.Res = res
	}
	return
}

func (s *PolicyResourceCrud) Update() (e error) {
	opts := &baremetal.UpdatePolicyOptions{}
	if description, ok := s.D.GetOk("description"); ok {
		opts.Description = description.(string)
	}

	if rawStatements, ok := s.D.GetOk("statements"); ok {
		statements := s.toStringArray(rawStatements)
		opts.Statements = statements
	}

	s.Res, e = s.Client.UpdatePolicy(s.D.Id(), opts)
	return
}

func (s *PolicyResourceCrud) SetData() {
	s.D.Set("statements", s.Res.Statements)
	s.D.Set("name", s.Res.Name)
	s.D.Set("description", s.Res.Description)
	s.D.Set("compartment_id", s.Res.CompartmentID)
	s.D.Set("state", s.Res.State)
	s.D.Set("time_created", s.Res.TimeCreated.String())
}

func (s *PolicyResourceCrud) Delete() (e error) {
	return s.Client.DeletePolicy(s.D.Id(), nil)
}
