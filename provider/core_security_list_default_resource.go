// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"fmt"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func DefaultSecurityListResource() *schema.Resource {
	defaultResourceSchema := ConvertToDefaultVcnResourceSchema(SecurityListResource())

	defaultResourceSchema.Create = createDefaultSecurityList
	defaultResourceSchema.Delete = deleteDefaultSecurityList

	return defaultResourceSchema
}

type DefaultSecurityListResourceCrud struct {
	SecurityListResourceCrud
}

func createDefaultSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func deleteDefaultSecurityList(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultSecurityListResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

func (s *DefaultSecurityListResourceCrud) Create() error {
	// If we are creating a default resource, then don't have to
	// actually create it. Just set the ID and update it.
	if defaultId, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		s.D.SetId(defaultId.(string))
		return s.Update()
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// This clears out all of the security rules from the list
// This is used to clear out default security list resources that can't otherwise be deleted
func (s *DefaultSecurityListResourceCrud) reset() error {
	request := oci_core.UpdateSecurityListRequest{}

	tmp := s.D.Id()
	request.SecurityListId = &tmp

	request.IngressSecurityRules = []oci_core.IngressSecurityRule{}

	request.EgressSecurityRules = []oci_core.EgressSecurityRule{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateSecurityList(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.SecurityList
	return nil
}

func (s *DefaultSecurityListResourceCrud) Delete() error {
	if _, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		// We can't actually delete a default resource.
		// Clear out its settings and mark it as deleted.
		err := s.reset()
		s.D.Set("state", s.DeletedTarget()[0])
		return err
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// You can't actually delete a default resource, so the creation target
// states are valid states for when waiting for delete to complete
func (s *DefaultSecurityListResourceCrud) DeletedPending() []string {
	return s.CreatedTarget()
}

func (s *DefaultSecurityListResourceCrud) DeletedTarget() []string {
	return s.CreatedTarget()
}
