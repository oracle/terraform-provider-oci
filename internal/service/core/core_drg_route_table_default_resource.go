// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	oci_core "github.com/oracle/oci-go-sdk/v65/core"
)

func DefaultCoreDrgRouteTableResource() *schema.Resource {
	defaultResourceSchema := ConvertToDefaultDrgRouteTableSchema(CoreDrgRouteTableResource())

	defaultResourceSchema.Create = createDefaultDrgRouteTable
	defaultResourceSchema.Delete = deleteDefaultDrgRouteTable

	return defaultResourceSchema
}

type DefaultDrgRouteTableResourceCrud struct {
	CoreDrgRouteTableResourceCrud
}

func createDefaultDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	return tfresource.CreateResource(d, sync)
}

func deleteDefaultDrgRouteTable(d *schema.ResourceData, m interface{}) error {
	sync := &DefaultDrgRouteTableResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).VirtualNetworkClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

func (s *DefaultDrgRouteTableResourceCrud) Create() error {
	// If we are creating a default resource, then don't have to actually Create
	// it. Just set the ID and Update it.
	if defaultId, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		s.D.SetId(defaultId.(string))
		return s.Update()
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

func (s *DefaultDrgRouteTableResourceCrud) reset() error {
	request := oci_core.UpdateDrgRouteTableRequest{}

	tmp := s.D.Id()
	request.DrgRouteTableId = &tmp

	// DRG route tables do not expose a service-side reset API. Clear the
	// mutable fields that the update API supports so the management resource can
	// behave like the other default-resource wrappers.
	request.DefinedTags = map[string]map[string]interface{}{}
	request.FreeformTags = map[string]string{}
	resetEcmp := false
	request.IsEcmpEnabled = &resetEcmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "core")

	response, err := s.Client.UpdateDrgRouteTable(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.DrgRouteTable

	if s.Res.ImportDrgRouteDistributionId != nil {
		return s.RemoveImportRouteDistribution()
	}

	return nil
}

func (s *DefaultDrgRouteTableResourceCrud) Delete() error {
	if _, ok := s.D.GetOkExists("manage_default_resource_id"); ok {
		// We can't actually delete a default resource.
		// Clear out its supported mutable settings and mark it as deleted.
		err := s.reset()
		s.D.Set("state", s.DeletedTarget()[0])
		return err
	}

	return fmt.Errorf("Default resource does not have a manage_default_resource_id set")
}

// You can't actually delete a default resource, so the creation target states
// are valid states for when waiting for delete to complete.
func (s *DefaultDrgRouteTableResourceCrud) DeletedPending() []string {
	return s.CreatedTarget()
}

func (s *DefaultDrgRouteTableResourceCrud) DeletedTarget() []string {
	return s.CreatedTarget()
}
