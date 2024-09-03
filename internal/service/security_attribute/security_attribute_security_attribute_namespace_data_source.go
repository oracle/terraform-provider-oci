// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package security_attribute

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_security_attribute "github.com/oracle/oci-go-sdk/v65/securityattribute"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func SecurityAttributeSecurityAttributeNamespaceDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["security_attribute_namespace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(SecurityAttributeSecurityAttributeNamespaceResource(), fieldMap, readSingularSecurityAttributeSecurityAttributeNamespace)
}

func readSingularSecurityAttributeSecurityAttributeNamespace(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespaceDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

type SecurityAttributeSecurityAttributeNamespaceDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_security_attribute.SecurityAttributeClient
	Res    *oci_security_attribute.GetSecurityAttributeNamespaceResponse
}

func (s *SecurityAttributeSecurityAttributeNamespaceDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecurityAttributeSecurityAttributeNamespaceDataSourceCrud) Get() error {
	request := oci_security_attribute.GetSecurityAttributeNamespaceRequest{}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "security_attribute")

	response, err := s.Client.GetSecurityAttributeNamespace(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespaceDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

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

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	s.D.Set("mode", s.Res.Mode)

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
