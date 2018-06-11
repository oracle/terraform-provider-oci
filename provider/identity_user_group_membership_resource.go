// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func UserGroupMembershipResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUserGroupMembership,
		Read:     readUserGroupMembership,
		Delete:   deleteUserGroupMembership,
		Schema: map[string]*schema.Schema{
			// Required
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			// The legacy provider required this but never sent it to the API (api does not accept it).
			// Make it optional for legacy users.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true, // this property is ignored, keep it optional for legacy configurations
			},
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

func createUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func deleteUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &UserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type UserGroupMembershipResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.UserGroupMembership
	DisableNotFoundRetries bool
}

func (s *UserGroupMembershipResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *UserGroupMembershipResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateCreating),
	}
}

func (s *UserGroupMembershipResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateActive),
	}
}

func (s *UserGroupMembershipResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateDeleting),
	}
}

func (s *UserGroupMembershipResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateDeleted),
	}
}

func (s *UserGroupMembershipResourceCrud) Create() error {
	request := oci_identity.AddUserToGroupRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.AddUserToGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserGroupMembership
	return nil
}

func (s *UserGroupMembershipResourceCrud) Get() error {
	request := oci_identity.GetUserGroupMembershipRequest{}

	tmp := s.D.Id()
	request.UserGroupMembershipId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUserGroupMembership(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserGroupMembership
	return nil
}

func (s *UserGroupMembershipResourceCrud) Delete() error {
	request := oci_identity.RemoveUserFromGroupRequest{}

	tmp := s.D.Id()
	request.UserGroupMembershipId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.RemoveUserFromGroup(context.Background(), request)
	return err
}

func (s *UserGroupMembershipResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.GroupId != nil {
		s.D.Set("group_id", *s.Res.GroupId)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

}
