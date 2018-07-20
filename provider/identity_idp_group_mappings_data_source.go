// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IdpGroupMappingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdpGroupMappings,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"identity_provider_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idp_group_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     crud.GetDataSourceItemSchema(IdpGroupMappingResource()),
			},
		},
	}
}

func readIdpGroupMappings(d *schema.ResourceData, m interface{}) error {
	sync := &IdpGroupMappingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type IdpGroupMappingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIdpGroupMappingsResponse
}

func (s *IdpGroupMappingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdpGroupMappingsDataSourceCrud) Get() error {
	request := oci_identity.ListIdpGroupMappingsRequest{}

	if identityProviderId, ok := s.D.GetOkExists("identity_provider_id"); ok {
		tmp := identityProviderId.(string)
		request.IdentityProviderId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

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

func (s *IdpGroupMappingsDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
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
			idpGroupMapping["inactive_state"] = *r.InactiveStatus
		}

		idpGroupMapping["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			idpGroupMapping["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, idpGroupMapping)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdpGroupMappingsDataSource().Schema["idp_group_mappings"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("idp_group_mappings", resources); err != nil {
		panic(err)
	}

	return
}
