// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"

	"github.com/oracle/terraform-provider-oci/crud"
)

func IdentityProvidersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityProviders,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"identity_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     IdentityProviderResource(),
			},
		},
	}
}

func readIdentityProviders(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProvidersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return crud.ReadResource(sync)
}

type IdentityProvidersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIdentityProvidersResponse
}

func (s *IdentityProvidersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityProvidersDataSourceCrud) Get() error {
	request := oci_identity.ListIdentityProvidersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_identity.ListIdentityProvidersProtocolEnum(protocol.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "identity")

	response, err := s.Client.ListIdentityProviders(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListIdentityProviders(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *IdentityProvidersDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		identityProvider := map[string]interface{}{
			"compartment_id": *r.GetCompartmentId(),
			"protocol":       string(oci_identity.ListIdentityProvidersProtocolSaml2),
		}

		if r.GetDefinedTags() != nil {
			identityProvider["defined_tags"] = definedTagsToMap(r.GetDefinedTags())
		}

		if r.GetDescription() != nil {
			identityProvider["description"] = *r.GetDescription()
		}

		identityProvider["freeform_tags"] = r.GetFreeformTags()

		if r.GetId() != nil {
			identityProvider["id"] = *r.GetId()
		}

		if r.GetInactiveStatus() != nil {
			identityProvider["inactive_state"] = *r.GetInactiveStatus()
		}

		if r.GetName() != nil {
			identityProvider["name"] = *r.GetName()
		}

		if r.GetProductType() != nil {
			identityProvider["product_type"] = *r.GetProductType()
		}

		identityProvider["state"] = r.GetLifecycleState()

		if r.GetTimeCreated() != nil {
			identityProvider["time_created"] = r.GetTimeCreated().String()
		}

		resources = append(resources, identityProvider)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityProvidersDataSource().Schema["identity_providers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("identity_providers", resources); err != nil {
		panic(err)
	}

	return
}
