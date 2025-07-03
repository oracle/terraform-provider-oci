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

func DbmulticloudOracleDbAzureBlobMountsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureBlobMounts,
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
			"oracle_db_azure_blob_container_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_blob_mount_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_blob_mount_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAzureBlobMountResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAzureBlobMounts(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobMountsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobMountClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureBlobMountsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureBlobMountClient
	Res    *oci_dbmulticloud.ListOracleDbAzureBlobMountsResponse
}

func (s *DbmulticloudOracleDbAzureBlobMountsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureBlobMountsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureBlobMountsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbAzureBlobContainerId, ok := s.D.GetOkExists("oracle_db_azure_blob_container_id"); ok {
		tmp := oracleDbAzureBlobContainerId.(string)
		request.OracleDbAzureBlobContainerId = &tmp
	}

	if oracleDbAzureBlobMountId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureBlobMountId.(string)
		request.OracleDbAzureBlobMountId = &tmp
	}

	if oracleDbAzureConnectorId, ok := s.D.GetOkExists("oracle_db_azure_connector_id"); ok {
		tmp := oracleDbAzureConnectorId.(string)
		request.OracleDbAzureConnectorId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureBlobMountLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureBlobMounts(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureBlobMounts(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureBlobMountsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureBlobMountsDataSource-", DbmulticloudOracleDbAzureBlobMountsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureBlobMount := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureBlobMountSummaryToMap(item))
	}
	oracleDbAzureBlobMount["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureBlobMountsDataSource().Schema["oracle_db_azure_blob_mount_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureBlobMount["items"] = items
	}

	resources = append(resources, oracleDbAzureBlobMount)
	if err := s.D.Set("oracle_db_azure_blob_mount_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
