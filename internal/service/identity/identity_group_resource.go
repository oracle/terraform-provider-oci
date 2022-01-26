// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/globalvar"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	oci_identity "github.com/oracle/oci-go-sdk/v56/identity"
)

func IdentityGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityGroup,
		Read:     readIdentityGroup,
		Update:   updateIdentityGroup,
		Delete:   deleteIdentityGroup,
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
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
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

func createIdentityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.Configuration = m.(*client.OracleClients).Configuration

	return tfresource.CreateResource(d, sync)
}

func readIdentityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Configuration          map[string]string
	Res                    *oci_identity.Group
	DisableNotFoundRetries bool
}

func (s *IdentityGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.GroupLifecycleStateCreating),
	}
}

func (s *IdentityGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.GroupLifecycleStateActive),
	}
}

func (s *IdentityGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.GroupLifecycleStateDeleting),
	}
}

func (s *IdentityGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.GroupLifecycleStateDeleted),
	}
}

func (s *IdentityGroupResourceCrud) Create() error {
	request := oci_identity.CreateGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	} else { // @next-break: remove
		// Prevent potentially inferring wrong TenancyOCID from InstancePrincipal
		if auth := s.Configuration["auth"]; strings.ToLower(auth) == strings.ToLower(globalvar.AuthInstancePrincipalSetting) {
			return fmt.Errorf("compartment_id must be specified for this resource")
		}
		// Maintain legacy contract of compartment_id defaulting to tenancy ocid if not specified
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
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityGroupResourceCrud) Get() error {
	request := oci_identity.GetGroupRequest{}

	tmp := s.D.Id()
	request.GroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityGroupResourceCrud) Update() error {
	request := oci_identity.UpdateGroupRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
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
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.GroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.Group
	return nil
}

func (s *IdentityGroupResourceCrud) Delete() error {
	request := oci_identity.DeleteGroupRequest{}

	tmp := s.D.Id()
	request.GroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteGroup(context.Background(), request)
	return err
}

func (s *IdentityGroupResourceCrud) SetData() error {
	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func (s *IdentityGroupResourceCrud) ExtraWaitPostCreateDelete() time.Duration {
	return time.Duration(2 * time.Second)
}
