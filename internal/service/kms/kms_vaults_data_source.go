// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v58/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func KmsVaultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsVaults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vaults": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(KmsVaultResource()),
			},
		},
	}
}

func readKmsVaults(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

type KmsVaultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.ListVaultsResponse
}

func (s *KmsVaultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsVaultsDataSourceCrud) Get() error {
	request := oci_kms.ListVaultsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

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

func (s *KmsVaultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsVaultsDataSource-", KmsVaultsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vault := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.CryptoEndpoint != nil {
			vault["crypto_endpoint"] = *r.CryptoEndpoint
		}

		if r.DefinedTags != nil {
			vault["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vault["display_name"] = *r.DisplayName
		}

		vault["freeform_tags"] = r.FreeformTags

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, KmsVaultsDataSource().Schema["vaults"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vaults", resources); err != nil {
		return err
	}

	return nil
}
