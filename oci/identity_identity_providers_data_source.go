// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/identity"
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
				Elem:     GetDataSourceItemSchema(IdentityProviderResource()),
			},
		},
	}
}

func readIdentityProviders(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityProvidersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).identityClient

	return ReadResource(sync)
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

func (s *IdentityProvidersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_identity.Saml2IdentityProvider:
			result["protocol"] = "SAML2"

			if v.MetadataUrl != nil {
				result["metadata_url"] = string(*v.MetadataUrl)
			}

			if v.RedirectUrl != nil {
				result["redirect_url"] = string(*v.RedirectUrl)
			}

			if v.SigningCertificate != nil {
				result["signing_certificate"] = string(*v.SigningCertificate)
			}

			if v.CompartmentId != nil {
				result["compartment_id"] = string(*v.CompartmentId)
			}

			if v.DefinedTags != nil {
				result["defined_tags"] = definedTagsToMap(v.DefinedTags)
			}

			if v.Description != nil {
				result["description"] = string(*v.Description)
			}

			result["freeform_tags"] = v.FreeformTags

			if v.Id != nil {
				result["id"] = string(*v.Id)
			}

			if v.InactiveStatus != nil {
				result["inactive_state"] = strconv.FormatInt(*v.InactiveStatus, 10)
			}

			if v.Name != nil {
				result["name"] = string(*v.Name)
			}

			result["product_type"] = string(*v.ProductType)

			result["state"] = string(v.LifecycleState)

			if v.TimeCreated != nil {
				result["time_created"] = v.TimeCreated.String()
			}
		default:
			log.Printf("[WARN] Received 'protocol' of unknown type %v", r)
			return nil
		}

		resources = append(resources, result)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, IdentityProvidersDataSource().Schema["identity_providers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("identity_providers", resources); err != nil {
		return err
	}

	return nil
}
