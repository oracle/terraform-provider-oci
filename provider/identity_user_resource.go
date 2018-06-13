// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	oci_identity "github.com/oracle/oci-go-sdk/identity"
)

func UserResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: crud.DefaultTimeout,
		Create:   createUser,
		Read:     readUser,
		Update:   updateUser,
		Delete:   deleteUser,
		Schema: map[string]*schema.Schema{
			// The legacy provider exposed this as read-only/computed. The API requires this param. For legacy users who are
			// not supplying a value, make it optional, behind the scenes it will use the tenancy ocid if not supplied.
			// If a user supplies the value, then changes it, it requires forcing new.
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
			},
			// Required
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: definedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

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
			// @Deprecated: time_modified (removed)
			"time_modified": {
				Type:       schema.TypeString,
				Deprecated: crud.FieldDeprecated("time_modified"),
				Computed:   true,
			},
		},
	}
}

func createUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.CreateResource(d, sync)
}

func readUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

func updateUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.UpdateResource(d, sync)
}

func deleteUser(d *schema.ResourceData, m interface{}) error {
	sync := &UserResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

type UserResourceCrud struct {
	crud.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.User
	DisableNotFoundRetries bool
}

func (s *UserResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *UserResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateCreating),
	}
}

func (s *UserResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateActive),
	}
}

func (s *UserResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleting),
	}
}

func (s *UserResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.UserLifecycleStateDeleted),
	}
}

func (s *UserResourceCrud) Create() error {
	request := oci_identity.CreateUserRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else {
		c := *s.Client.ConfigurationProvider()
		if c == nil {
			return fmt.Errorf("cannot access tenancyOCID")
		}
		tenancy, err := c.TenancyOCID()
		if err != nil {
			return err
		}
		request.CompartmentId = &tenancy
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Get() error {
	request := oci_identity.GetUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Update() error {
	request := oci_identity.UpdateUserRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := mapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = objectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateUser(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.User
	return nil
}

func (s *UserResourceCrud) Delete() error {
	request := oci_identity.DeleteUserRequest{}

	tmp := s.D.Id()
	request.UserId = &tmp

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteUser(context.Background(), request)
	return err
}

func (s *UserResourceCrud) SetData() {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", *s.Res.InactiveStatus)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

}
