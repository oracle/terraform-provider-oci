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

func DbmulticloudOracleDbAzureBlobContainersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbAzureBlobContainers,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"azure_storage_account_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"azure_storage_container_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_azure_blob_container_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAzureBlobContainerResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAzureBlobContainers(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbAzureBlobContainersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OracleDBAzureBlobContainerClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbAzureBlobContainersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.OracleDBAzureBlobContainerClient
	Res    *oci_dbmulticloud.ListOracleDbAzureBlobContainersResponse
}

func (s *DbmulticloudOracleDbAzureBlobContainersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAzureBlobContainersDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbAzureBlobContainersRequest{}

	if azureStorageAccountName, ok := s.D.GetOkExists("azure_storage_account_name"); ok {
		tmp := azureStorageAccountName.(string)
		request.AzureStorageAccountName = &tmp
	}

	if azureStorageContainerName, ok := s.D.GetOkExists("azure_storage_container_name"); ok {
		tmp := azureStorageContainerName.(string)
		request.AzureStorageContainerName = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbAzureBlobContainerId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbAzureBlobContainerId.(string)
		request.OracleDbAzureBlobContainerId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbAzureBlobContainerLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAzureBlobContainers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAzureBlobContainers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAzureBlobContainersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAzureBlobContainersDataSource-", DbmulticloudOracleDbAzureBlobContainersDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAzureBlobContainer := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAzureBlobContainerSummaryToMap(item))
	}
	oracleDbAzureBlobContainer["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAzureBlobContainersDataSource().Schema["oracle_db_azure_blob_container_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAzureBlobContainer["items"] = items
	}

	resources = append(resources, oracleDbAzureBlobContainer)
	if err := s.D.Set("oracle_db_azure_blob_container_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
