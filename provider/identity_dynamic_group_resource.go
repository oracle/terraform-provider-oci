// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func DynamicGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createDynamicGroup,
		Read:     readDynamicGroup,
		Update:   updateDynamicGroup,
		Delete:   deleteDynamicGroup,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"matching_rule": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inactive_state": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &DynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type DynamicGroupResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.DynamicGroup
	DisableNotFoundRetries bool
}

func (s *DynamicGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DynamicGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateCreating),
	}
}

func (s *DynamicGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateActive),
	}
}

func (s *DynamicGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateDeleting),
	}
}

func (s *DynamicGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateDeleted),
	}
}

func (s *DynamicGroupResourceCrud) Create() error {
	request := oci_identity.CreateDynamicGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		tmp := matchingRule.(string)
		request.MatchingRule = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *DynamicGroupResourceCrud) Get() error {
	request := oci_identity.GetDynamicGroupRequest{}

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *DynamicGroupResourceCrud) Update() error {
	request := oci_identity.UpdateDynamicGroupRequest{}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		tmp := matchingRule.(string)
		request.MatchingRule = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *DynamicGroupResourceCrud) Delete() error {
	request := oci_identity.DeleteDynamicGroupRequest{}

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteDynamicGroup(context.Background(), request)
	return err
}

func (s *DynamicGroupResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	if s.Res.MatchingRule != nil {
		s.D.Set("matching_rule", *s.Res.MatchingRule)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
