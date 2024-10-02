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

func SecurityAttributeSecurityAttributeNamespacesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSecurityAttributeSecurityAttributeNamespaces,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id_in_subtree": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_attribute_namespaces": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(SecurityAttributeSecurityAttributeNamespaceResource()),
			},
		},
	}
}

func readSecurityAttributeSecurityAttributeNamespaces(d *schema.ResourceData, m interface{}) error {
	sync := &SecurityAttributeSecurityAttributeNamespacesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SecurityAttributeClient()

	return tfresource.ReadResource(sync)
}

type SecurityAttributeSecurityAttributeNamespacesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_security_attribute.SecurityAttributeClient
	Res    *oci_security_attribute.ListSecurityAttributeNamespacesResponse
}

func (s *SecurityAttributeSecurityAttributeNamespacesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *SecurityAttributeSecurityAttributeNamespacesDataSourceCrud) Get() error {
	request := oci_security_attribute.ListSecurityAttributeNamespacesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if compartmentIdInSubtree, ok := s.D.GetOkExists("compartment_id_in_subtree"); ok {
		tmp := compartmentIdInSubtree.(bool)
		request.CompartmentIdInSubtree = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_security_attribute.SecurityAttributeNamespaceLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "security_attribute")

	response, err := s.Client.ListSecurityAttributeNamespaces(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListSecurityAttributeNamespaces(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *SecurityAttributeSecurityAttributeNamespacesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("SecurityAttributeSecurityAttributeNamespacesDataSource-", SecurityAttributeSecurityAttributeNamespacesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		securityAttributeNamespace := map[string]interface{}{}

		if r.CompartmentId != nil {
			securityAttributeNamespace["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			securityAttributeNamespace["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.Description != nil {
			securityAttributeNamespace["description"] = *r.Description
		}

		securityAttributeNamespace["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			securityAttributeNamespace["id"] = *r.Id
		}

		if r.IsRetired != nil {
			securityAttributeNamespace["is_retired"] = *r.IsRetired
		}

		securityAttributeNamespace["mode"] = r.Mode

		if r.Name != nil {
			securityAttributeNamespace["name"] = *r.Name
		}

		securityAttributeNamespace["state"] = r.LifecycleState

		if r.SystemTags != nil {
			securityAttributeNamespace["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.TimeCreated != nil {
			securityAttributeNamespace["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, securityAttributeNamespace)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, SecurityAttributeSecurityAttributeNamespacesDataSource().Schema["security_attribute_namespaces"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("security_attribute_namespaces", resources); err != nil {
		return err
	}

	return nil
}
