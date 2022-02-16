// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	oci_identity "github.com/oracle/oci-go-sdk/v58/identity"
)

func IdentityIdpGroupMappingResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityIdpGroupMapping,
		Read:     readIdentityIdpGroupMapping,
		Update:   updateIdentityIdpGroupMapping,
		Delete:   deleteIdentityIdpGroupMapping,
		Schema: map[string]*schema.Schema{
			// Required
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"idp_group_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
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

func createIdentityIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.CreateResource(d, sync)
}

func readIdentityIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

func updateIdentityIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityIdpGroupMapping(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdpGroupMappingResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityIdpGroupMappingResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity.IdentityClient
	Res                    *oci_identity.IdpGroupMapping
	DisableNotFoundRetries bool
}

func (s *IdentityIdpGroupMappingResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityIdpGroupMappingResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateCreating),
	}
}

func (s *IdentityIdpGroupMappingResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateActive),
	}
}

func (s *IdentityIdpGroupMappingResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateDeleting),
	}
}

func (s *IdentityIdpGroupMappingResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_identity.IdpGroupMappingLifecycleStateDeleted),
	}
}

func (s *IdentityIdpGroupMappingResourceCrud) Create() error {
	request := oci_identity.CreateIdpGroupMappingRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if idpGroupName, ok := s.D.GetOkExists("idp_group_name"); ok {
		tmp := idpGroupName.(string)
		request.IdpGroupName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.CreateIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdentityIdpGroupMappingResourceCrud) Get() error {
	request := oci_identity.GetIdpGroupMappingRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	identityProviderId, mappingId, err := parseIdpGroupMappingCompositeId(s.D.Id())
	if err == nil {
		request.IdentityProviderId = &identityProviderId
		request.MappingId = &mappingId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.GetIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdentityIdpGroupMappingResourceCrud) Update() error {
	request := oci_identity.UpdateIdpGroupMappingRequest{}

	if groupId, ok := s.D.GetOkExists("group_id"); ok {
		tmp := groupId.(string)
		request.GroupId = &tmp
	}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	if idpGroupName, ok := s.D.GetOkExists("idp_group_name"); ok {
		tmp := idpGroupName.(string)
		request.IdpGroupName = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	response, err := s.Client.UpdateIdpGroupMapping(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.IdpGroupMapping
	return nil
}

func (s *IdentityIdpGroupMappingResourceCrud) Delete() error {
	request := oci_identity.DeleteIdpGroupMappingRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	tmp := s.D.Id()
	request.MappingId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity")

	_, err := s.Client.DeleteIdpGroupMapping(context.Background(), request)
	return err
}

func (s *IdentityIdpGroupMappingResourceCrud) SetData() error {

	identityProviderId, mappingId, err := parseIdpGroupMappingCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("identity_provider_id", identityProviderId)
		s.D.SetId(mappingId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.GroupId != nil {
		s.D.Set("group_id", *s.Res.GroupId)
	}

	if s.Res.IdpId != nil {
		s.D.Set("identity_provider_id", *s.Res.IdpId)
	}

	if s.Res.IdpGroupName != nil {
		s.D.Set("idp_group_name", *s.Res.IdpGroupName)
	}

	if s.Res.InactiveStatus != nil {
		s.D.Set("inactive_state", strconv.FormatInt(*s.Res.InactiveStatus, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}

func GetIdpGroupMappingCompositeId(identityProviderId string, mappingId string) string {
	identityProviderId = url.PathEscape(identityProviderId)
	mappingId = url.PathEscape(mappingId)
	compositeId := "identityProviders/" + identityProviderId + "/groupMappings/" + mappingId
	return compositeId
}

func parseIdpGroupMappingCompositeId(compositeId string) (identityProviderId string, mappingId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("identityProviders/.*/groupMappings/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	identityProviderId, _ = url.PathUnescape(parts[1])
	mappingId, _ = url.PathUnescape(parts[3])

	return
}
