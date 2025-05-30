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

func DbmulticloudOracleDbAzureVaultAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureVaultAssociations,
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
			"oracle_db_azure_vault_association_id": {
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
			"oracle_db_azure_vault_association_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAzureVaultAssociationResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAzureVaultAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureVaultAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDbAzureVaultAssociationClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureVaultAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDbAzureVaultAssociationClient
	Res    *oci_dbmulticloud.ListOracleDbAzureVaultAssociationsResponse
}

func (s *DbmulticloudOracleDbAzureVaultAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureVaultAssociationsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureVaultAssociationsRequest{}

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

	if oracleDbAzureVaultAssociationId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureVaultAssociationId.(string)
		request.OracleDbAzureVaultAssociationId = &tmp
	}

	if oracleDbAzureVaultId, ok := s.D.GetOkExists("oracle_db_azure_vault_id"); ok {
		tmp := oracleDbAzureVaultId.(string)
		request.OracleDbAzureVaultId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureVaultAssociationLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureVaultAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureVaultAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureVaultAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureVaultAssociationsDataSource-", DbmulticloudOracleDbAzureVaultAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureVaultAssociation := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureVaultAssociationSummaryToMap(item))
	}
	oracleDbAzureVaultAssociation["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureVaultAssociationsDataSource().Schema["oracle_db_azure_vault_association_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureVaultAssociation["items"] = items
	}

	resources = append(resources, oracleDbAzureVaultAssociation)
	if err := s.D.Set("oracle_db_azure_vault_association_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
