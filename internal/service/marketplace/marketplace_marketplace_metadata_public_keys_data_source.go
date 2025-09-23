// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package marketplace

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_marketplace "github.com/oracle/oci-go-sdk/v65/marketplace"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MarketplaceMarketplaceMetadataPublicKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMarketplaceMarketplaceMetadataPublicKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"marketplace_metadata_public_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"certificate_chain": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"certificate_thumbprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"exponent": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_algorithm": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_use": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"modulus": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readMarketplaceMarketplaceMetadataPublicKeys(d *schema.ResourceData, m interface{}) error {
	sync := &MarketplaceMarketplaceMetadataPublicKeysDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MarketplaceClient()

	return tfresource.ReadResource(sync)
}

type MarketplaceMarketplaceMetadataPublicKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_marketplace.MarketplaceClient
	Res    *oci_marketplace.ListMarketplaceMetadataPublicKeysResponse
}

func (s *MarketplaceMarketplaceMetadataPublicKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MarketplaceMarketplaceMetadataPublicKeysDataSourceCrud) Get() error {
	request := oci_marketplace.ListMarketplaceMetadataPublicKeysRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "marketplace")

	response, err := s.Client.ListMarketplaceMetadataPublicKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListMarketplaceMetadataPublicKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MarketplaceMarketplaceMetadataPublicKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MarketplaceMarketplaceMetadataPublicKeysDataSource-", MarketplaceMarketplaceMetadataPublicKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		marketplaceMetadataPublicKey := map[string]interface{}{}

		marketplaceMetadataPublicKey["certificate_chain"] = r.CertificateChain

		if r.CertificateThumbprint != nil {
			marketplaceMetadataPublicKey["certificate_thumbprint"] = *r.CertificateThumbprint
		}

		if r.Exponent != nil {
			marketplaceMetadataPublicKey["exponent"] = *r.Exponent
		}

		if r.KeyAlgorithm != nil {
			marketplaceMetadataPublicKey["key_algorithm"] = *r.KeyAlgorithm
		}

		if r.KeyId != nil {
			marketplaceMetadataPublicKey["key_id"] = *r.KeyId
		}

		if r.KeyType != nil {
			marketplaceMetadataPublicKey["key_type"] = *r.KeyType
		}

		if r.KeyUse != nil {
			marketplaceMetadataPublicKey["key_use"] = *r.KeyUse
		}

		if r.Modulus != nil {
			marketplaceMetadataPublicKey["modulus"] = *r.Modulus
		}

		resources = append(resources, marketplaceMetadataPublicKey)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, MarketplaceMarketplaceMetadataPublicKeysDataSource().Schema["marketplace_metadata_public_keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("marketplace_metadata_public_keys", resources); err != nil {
		return err
	}

	return nil
}
