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

func DbmulticloudOracleDbGcpKeyRingsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDbmulticloudOracleDbGcpKeyRings,
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
			"oracle_db_gcp_connector_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_gcp_key_ring_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"oracle_db_gcp_key_ring_summary_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(DbmulticloudOracleDbGcpKeyRingResource()),
						},
					},
				},
			},
		},
	}
}

func readDbmulticloudOracleDbGcpKeyRings(d *schema.ResourceData, m interface{}) error {
	sync := &DbmulticloudOracleDbGcpKeyRingsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DbMulticloudGCPProviderClient()

	return tfresource.ReadResource(sync)
}

type DbmulticloudOracleDbGcpKeyRingsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_dbmulticloud.DbMulticloudGCPProviderClient
	Res    *oci_dbmulticloud.ListOracleDbGcpKeyRingsResponse
}

func (s *DbmulticloudOracleDbGcpKeyRingsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DbmulticloudOracleDbGcpKeyRingsDataSourceCrud) Get() error {
	request := oci_dbmulticloud.ListOracleDbGcpKeyRingsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if oracleDbGcpConnectorId, ok := s.D.GetOkExists("oracle_db_gcp_connector_id"); ok {
		tmp := oracleDbGcpConnectorId.(string)
		request.OracleDbGcpConnectorId = &tmp
	}

	if oracleDbGcpKeyRingId, ok := s.D.GetOkExists("id"); ok {
		tmp := oracleDbGcpKeyRingId.(string)
		request.OracleDbGcpKeyRingId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_dbmulticloud.OracleDbGcpKeyRingLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "dbmulticloud")

	response, err := s.Client.ListOracleDbGcpKeyRings(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListOracleDbGcpKeyRings(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DbmulticloudOracleDbGcpKeyRingsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DbmulticloudOracleDbGcpKeyRingsDataSource-", DbmulticloudOracleDbGcpKeyRingsDataSource(), s.D))
	resources := []map[string]interface{}{}
	oracleDbGcpKeyRing := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, OracleDbGcpKeyRingSummaryToMap(item))
	}
	oracleDbGcpKeyRing["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DbmulticloudOracleDbGcpKeyRingsDataSource().Schema["oracle_db_gcp_key_ring_summary_collection"].Elem.(*schema.Resource).Schema)
		oracleDbGcpKeyRing["items"] = items
	}

	resources = append(resources, oracleDbGcpKeyRing)
	if err := s.D.Set("oracle_db_gcp_key_ring_summary_collection", resources); err != nil {
		return err
	}

	return nil
}
