// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package dbmulticloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_dbmulticloud "github.com/oracle/oci-go-sdk/v65/dbmulticloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DbmulticloudOracleDbAwsIdentityConnectorsDataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDbmulticloudOracleDbAwsIdentityConnectorsWithContext,
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
			"oracle_db_aws_identity_connector_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbAwsIdentityConnectorResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbAwsIdentityConnectorsWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DbmulticloudOracleDbAwsIdentityConnectorsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudAwsProviderClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DbmulticloudOracleDbAwsIdentityConnectorsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudAwsProviderClient
	Res    *oci_dbmulticloud.ListOracleDbAwsIdentityConnectorsResponse
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorsDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_dbmulticloud.ListOracleDbAwsIdentityConnectorsRequest{}

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
		request.LifecycleState = oci_dbmulticloud.OracleDbAwsIdentityConnectorLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbAwsIdentityConnectors(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbAwsIdentityConnectors(ctx, request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbAwsIdentityConnectorsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbAwsIdentityConnectorsDataSource-", DbmulticloudOracleDbAwsIdentityConnectorsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbAwsIdentityConnector := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbAwsIdentityConnectorSummaryToMap(item))
	}
	oracleDbAwsIdentityConnector["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbAwsIdentityConnectorsDataSource().Schema["oracle_db_aws_identity_connector_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbAwsIdentityConnector["items"] = items
	}

	resources = append(resources, oracleDbAwsIdentityConnector)
	if err := s.D.Set("oracle_db_aws_identity_connector_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
