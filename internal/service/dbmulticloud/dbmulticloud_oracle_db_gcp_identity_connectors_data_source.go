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

func DbmulticloudOracleDbGcpIdentityConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbGcpIdentityConnectors,
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
			"resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_gcp_identity_connector_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbGcpIdentityConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbGcpIdentityConnectors(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpIdentityConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbGcpIdentityConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudGCPProviderClient
	Res    *oci_dbmulticloud.ListOracleDbGcpIdentityConnectorsResponse
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbGcpIdentityConnectorsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if resourceId, ok := s.D.GetOkExists("resource_id"); ok {
		tmp := resourceId.(string)
		request.ResourceId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbGcpIdentityConnectorLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbGcpIdentityConnectors(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbGcpIdentityConnectors(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbGcpIdentityConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbGcpIdentityConnectorsDataSource-", DbmulticloudOracleDbGcpIdentityConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbGcpIdentityConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbGcpIdentityConnectorSummaryToMap(item))
	}
	oracleDbGcpIdentityConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbGcpIdentityConnectorsDataSource().Schema["oracle_db_gcp_identity_connector_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbGcpIdentityConnector["items"] = items
	}

	resources = append(resources, oracleDbGcpIdentityConnector)
	if err := s.D.Set("oracle_db_gcp_identity_connector_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
