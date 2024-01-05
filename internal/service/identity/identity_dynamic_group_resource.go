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

func IdentityDynamicGroupResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDynamicGroup,
		Read:     readIdentityDynamicGroup,
		Update:   updateIdentityDynamicGroup,
		Delete:   deleteIdentityDynamicGroup,
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

func createIdentityDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDynamicGroup(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDynamicGroupResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDynamicGroupResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.DynamicGroup
	DisableNotFoundRetries bool
}

func (s *IdentityDynamicGroupResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDynamicGroupResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateCreating),
	}
}

func (s *IdentityDynamicGroupResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateActive),
	}
}

func (s *IdentityDynamicGroupResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateDeleting),
	}
}

func (s *IdentityDynamicGroupResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.DynamicGroupLifecycleStateDeleted),
	}
}

func (s *IdentityDynamicGroupResourceCrud) Create() error {
	request := oci_identity.CreateDynamicGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
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
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		tmp := matchingRule.(string)
		request.MatchingRule = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *IdentityDynamicGroupResourceCrud) Get() error {
	request := oci_identity.GetDynamicGroupRequest{}

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *IdentityDynamicGroupResourceCrud) Update() error {
	request := oci_identity.UpdateDynamicGroupRequest{}

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

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if matchingRule, ok := s.D.GetOkExists("matching_rule"); ok {
		tmp := matchingRule.(string)
		request.MatchingRule = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateDynamicGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DynamicGroup
	return nil
}

func (s *IdentityDynamicGroupResourceCrud) Delete() error {
	request := oci_identity.DeleteDynamicGroupRequest{}

	tmp := s.D.Id()
	request.DynamicGroupId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteDynamicGroup(context.Background(), request)
	return err
}

func (s *IdentityDynamicGroupResourceCrud) SetData() error {
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

	return nil
}
