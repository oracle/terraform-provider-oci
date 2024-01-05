// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity

import (
	"context"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_identity "github.com/oracle/oci-go-sdk/v65/identity"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityIdentityProvidersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readIdentityIdentityProviders,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"identity_providers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(IdentityIdentityProviderResource()),
			},
		},
	}
}

func readIdentityIdentityProviders(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityIdentityProvidersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).IdentityClient()

	return tfresource.ReadResource(sync)
}

type IdentityIdentityProvidersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_identity.IdentityClient
	Res    *oci_identity.ListIdentityProvidersResponse
}

func (s *IdentityIdentityProvidersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *IdentityIdentityProvidersDataSourceCrud) Get() error {
	request := oci_identity.ListIdentityProvidersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if protocol, ok := s.D.GetOkExists("protocol"); ok {
		request.Protocol = oci_identity.ListIdentityProvidersProtocolEnum(protocol.(string))
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_identity.IdentityProviderLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "identity")

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

func (s *IdentityIdentityProvidersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("IdentityIdentityProvidersDataSource-", IdentityIdentityProvidersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		result := map[string]interface{}{}
		switch v := (r).(type) {
		case oci_identity.Saml2IdentityProvider:
			result["protocol"] = "SAML2"

			result["freeform_attributes"] = v.FreeformAttributes

			if v.Metadata != nil {
				result["metadata"] = string(*v.Metadata)
			}

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
				result["defined_tags"] = tfresource.DefinedTagsToMap(v.DefinedTags)
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

			if v.ProductType != nil {
				result["product_type"] = string(*v.ProductType)
			}

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, IdentityIdentityProvidersDataSource().Schema["identity_providers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("identity_providers", resources); err != nil {
		return err
	}

	return nil
}
