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

func SecurityAttributeSecurityAttributesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSecurityAttributeSecurityAttributes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"security_attribute_namespace_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(SecurityAttributeSecurityAttributeResource()),
			},
		},
	}
}

func readSecurityAttributeSecurityAttributes(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

type SecurityAttributeSecurityAttributesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_security_attribute.SecurityAttributeClient
	Res    *oci_security_attribute.ListSecurityAttributesResponse
}

func (s *SecurityAttributeSecurityAttributesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecurityAttributeSecurityAttributesDataSourceCrud) Get() error {
	request := oci_security_attribute.ListSecurityAttributesRequest{}

	if securityAttributeNamespaceId, ok := s.D.GetOkExists("security_attribute_namespace_id"); ok {
		tmp := securityAttributeNamespaceId.(string)
		request.SecurityAttributeNamespaceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_security_attribute.SecurityAttributeLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "security_attribute")

	response, err := s.Client.ListSecurityAttributes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityAttributes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SecurityAttributeSecurityAttributesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SecurityAttributeSecurityAttributesDataSource-", SecurityAttributeSecurityAttributesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityAttribute := map[string]interface{}{
			"security_attribute_namespace_id": *r.SecurityAttributeNamespaceId,
		}

		if r.CompartmentId != nil {
			securityAttribute["compartment_id"] = *r.CompartmentId
		}

		if r.Description != nil {
			securityAttribute["description"] = *r.Description
		}

		if r.Id != nil {
			securityAttribute["id"] = *r.Id
		}

		if r.IsRetired != nil {
			securityAttribute["is_retired"] = *r.IsRetired
		}

		if r.Name != nil {
			securityAttribute["name"] = *r.Name
		}

		if r.SecurityAttributeNamespaceName != nil {
			securityAttribute["security_attribute_namespace_name"] = *r.SecurityAttributeNamespaceName
		}

		securityAttribute["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			securityAttribute["time_created"] = r.TimeCreated.String()
		}

		if r.Type != nil {
			securityAttribute["type"] = *r.Type
		}

		resources = append(resources, securityAttribute)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, SecurityAttributeSecurityAttributesDataSource().Schema["security_attributes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("security_attributes", resources); err != nil {
		return err
	}

	return nil
}
