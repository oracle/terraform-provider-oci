// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsVaultReplicasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsVaultReplicas,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"vault_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vault_replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"crypto_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"management_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readKmsVaultReplicas(d *schema.ResourceData, m interface{}) error {
	sync := &KmsVaultReplicasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).KmsVaultClient()

	return tfresource.ReadResource(sync)
}

type KmsVaultReplicasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsVaultClient
	Res    *oci_kms.ListVaultReplicasResponse
}

func (s *KmsVaultReplicasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsVaultReplicasDataSourceCrud) Get() error {
	request := oci_kms.ListVaultReplicasRequest{}

	if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
		tmp := vaultId.(string)
		request.VaultId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.ListVaultReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	// workaround for redundant pagination token returned in response.
	// We found pagination token is returned even when result is only one page,
	// and querying with this token in next request causes error.
	// As the replica shouldn't take more than one page, remove token handling for now.
	/*
		request.Page = s.Res.OpcNextPage

		for request.Page != nil {
			listResponse, err := s.Client.ListVaultReplicas(context.Background(), request)
			if err != nil {
				return err
			}

			s.Res.Items = append(s.Res.Items, listResponse.Items...)
			request.Page = listResponse.OpcNextPage
		}*/

	return nil
}

func (s *KmsVaultReplicasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsVaultReplicasDataSource-", KmsVaultReplicasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vaultReplica := map[string]interface{}{}

		if r.CryptoEndpoint != nil {
			vaultReplica["crypto_endpoint"] = *r.CryptoEndpoint
		}

		if r.ManagementEndpoint != nil {
			vaultReplica["management_endpoint"] = *r.ManagementEndpoint
		}

		if r.Region != nil {
			vaultReplica["region"] = *r.Region
		}

		vaultReplica["status"] = r.Status

		resources = append(resources, vaultReplica)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, KmsVaultReplicasDataSource().Schema["vault_replicas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vault_replicas", resources); err != nil {
		return err
	}

	return nil
}
