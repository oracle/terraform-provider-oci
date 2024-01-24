// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"
)

func IdentityUserGroupMembershipResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityUserGroupMembership,
		Read:     readIdentityUserGroupMembership,
		Delete:   deleteIdentityUserGroupMembership,
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
				Optional: true, // this property is ignored, keep it optional for legacy.Configurations
			},
			"inactive_state": {
				Type:     schema.TypeString,
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

func createIdentityUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func deleteIdentityUserGroupMembership(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityUserGroupMembershipResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityUserGroupMembershipResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.UserGroupMembership
	DisableNotFoundRetries bool
}

func (s *IdentityUserGroupMembershipResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityUserGroupMembershipResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateCreating),
	}
}

func (s *IdentityUserGroupMembershipResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateActive),
	}
}

func (s *IdentityUserGroupMembershipResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateDeleting),
	}
}

func (s *IdentityUserGroupMembershipResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UserGroupMembershipLifecycleStateDeleted),
	}
}

func (s *IdentityUserGroupMembershipResourceCrud) Create() error {
	request := oci_identity.AddUserToGroupRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if userId, ok := s.D.GetOkExists("user_id"); ok {
		tmp := userId.(string)
		request.UserId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.AddUserToGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserGroupMembership
	return nil
}

func (s *IdentityUserGroupMembershipResourceCrud) Get() error {
	request := oci_identity.GetUserGroupMembershipRequest{}

	tmp := s.D.Id()
	request.UserGroupMembershipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUserGroupMembership(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.UserGroupMembership
	return nil
}

func (s *IdentityUserGroupMembershipResourceCrud) Delete() error {
	request := oci_identity.RemoveUserFromGroupRequest{}

	tmp := s.D.Id()
	request.UserGroupMembershipId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.RemoveUserFromGroup(context.Background(), request)
	return err
}

func (s *IdentityUserGroupMembershipResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.GroupId != nil {
		s.D.Set("group_id", *s.Res.GroupId)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.UserId != nil {
		s.D.Set("user_id", *s.Res.UserId)
	}

	return nil
}
