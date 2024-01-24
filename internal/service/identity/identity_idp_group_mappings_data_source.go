// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityIdpGroupMappingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityIdpGroupMappings,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idp_group_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityIdpGroupMappingResource()),
			},
		},
	}
}

func readIdentityIdpGroupMappings(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdpGroupMappingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityIdpGroupMappingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIdpGroupMappingsResponse
}

func (s *IdentityIdpGroupMappingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIdpGroupMappingsDataSourceCrud) Get() error {
	request := oci_identity.ListIdpGroupMappingsRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

	response, err := s.Client.ListIdpGroupMappings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIdpGroupMappings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityIdpGroupMappingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityIdpGroupMappingsDataSource-", IdentityIdpGroupMappingsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		idpGroupMapping := map[string]interface{}{
			"identity_provider_id": *r.IdpId,
		}

		if r.CompartmentId != nil {
			idpGroupMapping["compartment_id"] = *r.CompartmentId
		}

		if r.GroupId != nil {
			idpGroupMapping["group_id"] = *r.GroupId
		}

		if r.Id != nil {
			idpGroupMapping["id"] = *r.Id
		}

		if r.IdpGroupName != nil {
			idpGroupMapping["idp_group_name"] = *r.IdpGroupName
		}

		if r.InactiveStatus != nil {
			idpGroupMapping["inactive_state"] = strconv.FormatInt(*r.InactiveStatus, 10)
		}

		idpGroupMapping["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			idpGroupMapping["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, idpGroupMapping)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityIdpGroupMappingsDataSource().Schema["idp_group_mappings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("idp_group_mappings", resources); err != nil {
		return err
	}

	return nil
}
