// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func VaultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVaults,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vaults": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(VaultResource()),
			},
		},
	}
}

func readVaults(d *schema.ResourceData, m interface{}) error {
	sync := &VaultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).kmsVaultClient

	return ReadResource(sync)
}

type VaultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.ListVaultsResponse
}

func (s *VaultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VaultsDataSourceCrud) Get() error {
	request := oci_kms.ListVaultsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.ListVaults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVaults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VaultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vault := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CryptoEndpoint != nil {
			vault["crypto_endpoint"] = *r.CryptoEndpoint
		}

		if r.DisplayName != nil {
			vault["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			vault["id"] = *r.Id
		}

		if r.ManagementEndpoint != nil {
			vault["management_endpoint"] = *r.ManagementEndpoint
		}

		vault["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			vault["time_created"] = r.TimeCreated.String()
		}

		vault["vault_type"] = r.VaultType

		resources = append(resources, vault)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VaultsDataSource().Schema["vaults"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vaults", resources); err != nil {
		return err
	}

	return nil
}
