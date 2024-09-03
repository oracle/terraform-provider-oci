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

func SecurityAttributeSecurityAttributeDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["security_attribute_name"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["security_attribute_namespace_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(SecurityAttributeSecurityAttributeResource(), fieldMap, readSingularSecurityAttributeSecurityAttribute)
}

func readSingularSecurityAttributeSecurityAttribute(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

type SecurityAttributeSecurityAttributeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_security_attribute.SecurityAttributeClient
	Res    *oci_security_attribute.GetSecurityAttributeResponse
}

func (s *SecurityAttributeSecurityAttributeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecurityAttributeSecurityAttributeDataSourceCrud) Get() error {
	request := oci_security_attribute.GetSecurityAttributeRequest{}

	if securityAttributeName, ok := s.D.GetOkExists("security_attribute_name"); ok {
		tmp := securityAttributeName.(string)
		request.SecurityAttributeName = &tmp
	}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "security_attribute")

	response, err := s.Client.GetSecurityAttribute(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *SecurityAttributeSecurityAttributeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.IsRetired != nil {
		s.D.Set("is_retired", *s.Res.IsRetired)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.SecurityAttributeNamespaceName != nil {
		s.D.Set("security_attribute_namespace_name", *s.Res.SecurityAttributeNamespaceName)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.Type != nil {
		s.D.Set("type", *s.Res.Type)
	}

	if s.Res.Validator != nil {
		validatorArray := []interface{}{}
		if validatorMap := BaseSecurityAttributeValidatorToMap(&s.Res.Validator); validatorMap != nil {
			validatorArray = append(validatorArray, validatorMap)
		}
		s.D.Set("validator", validatorArray)
	} else {
		s.D.Set("validator", nil)
	}

	return nil
}
