// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAzureVaultsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureVaults,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_resource_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_vault_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_vault_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAzureVaultResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAzureVaults(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureVaultsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureVaultClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureVaultsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureVaultClient
	Res    *oci_dbmulticloud.ListOracleDbAzureVaultsResponse
}

func (s *DbmulticloudOracleDbAzureVaultsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureVaultsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureVaultsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbAzureConnectorId, ok := s.D.GetOkExists("oracle_db_azure_connector_id"); ok {
		tmp := oracleDbAzureConnectorId.(string)
		request.OracleDbAzureConnectorId = &tmp
	}

	if oracleDbAzureResourceGroup, ok := s.D.GetOkExists("oracle_db_azure_resource_group"); ok {
		tmp := oracleDbAzureResourceGroup.(string)
		request.OracleDbAzureResourceGroup = &tmp
	}

	if oracleDbAzureVaultId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureVaultId.(string)
		request.OracleDbAzureVaultId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureVaultLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureVaults(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureVaults(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureVaultsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureVaultsDataSource-", DbmulticloudOracleDbAzureVaultsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureVault := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureVaultSummaryToMap(item))
	}
	oracleDbAzureVault["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureVaultsDataSource().Schema["oracle_db_azure_vault_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureVault["items"] = items
	}

	resources = append(resources, oracleDbAzureVault)
	if err := s.D.Set("oracle_db_azure_vault_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
