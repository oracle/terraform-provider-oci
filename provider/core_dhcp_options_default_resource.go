// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"

	"github.com/oracle/terraform-provider-oci/crud"

	"fmt"

	oci_core "github.com/oracle/oci-go-sdk/core"
)

func DefaultDhcpOptionsResource() *schema.Resource {
	defaultResourceSchema := ConvertToDefaultVcnResourceSchema(DhcpOptionsResource())

	defaultResourceSchema.Create = createDefaultDhcpOptions
	defaultResourceSchema.Delete = deleteDefaultDhcpOptions

	return defaultResourceSchema
}

type DefaultDhcpOptionsResourceCrud struct {
	DhcpOptionsResourceCrud
}

func createDefaultDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient

	return crud.CreateResource(d, sync)
}

func deleteDefaultDhcpOptions(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultDhcpOptionsResourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).virtualNetworkClient
	sync.DisableNotFoundRetries = true

	return crud.DeleteResource(d, sync)
}

func (s *DefaultDhcpOptionsResourceCrud) Create() error {
	// If we are creating a default resource, then don't have to
	// actually create it. Just set the ID and update it.
	if defaultId, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		s.D.SetId(defaultId.(string))
		return s.Update()
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// This creates a DHCP option with no dns servers
// This is used to clear out default DHCP options resources that can't otherwise be deleted
func (s *DefaultDhcpOptionsResourceCrud) reset() error {
	request := oci_core.UpdateDhcpOptionsRequest{}

	tmp := s.D.Id()
	request.DhcpId = &tmp

	request.Options = []oci_core.DhcpOption{
		oci_core.DhcpDnsOption{
			CustomDnsServers: []string{},
			ServerType:       "VcnLocalPlusInternet",
		},
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDhcpOptions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DhcpOptions
	return nil
}

func (s *DefaultDhcpOptionsResourceCrud) Delete() error {
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
func (s *DefaultDhcpOptionsResourceCrud) DeletedPending() []string {
	return s.CreatedTarget()
}

func (s *DefaultDhcpOptionsResourceCrud) DeletedTarget() []string {
	return s.CreatedTarget()
}
