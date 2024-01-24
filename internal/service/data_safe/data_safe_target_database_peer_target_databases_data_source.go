// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package data_safe

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataSafeTargetDatabasePeerTargetDatabasesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDataSafeTargetDatabasePeerTargetDatabases,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"target_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"peer_target_database_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     DataSafeTargetDatabasePeerTargetDatabaseResource(),
						},
					},
				},
			},
		},
	}
}

func readDataSafeTargetDatabasePeerTargetDatabases(d *schema.ResourceData, m interface{}) error {
	sync := &DataSafeTargetDatabasePeerTargetDatabasesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DataSafeClient()

	return tfresource.ReadResource(sync)
}

type DataSafeTargetDatabasePeerTargetDatabasesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_data_safe.DataSafeClient
	Res    *oci_data_safe.ListPeerTargetDatabasesResponse
}

func (s *DataSafeTargetDatabasePeerTargetDatabasesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataSafeTargetDatabasePeerTargetDatabasesDataSourceCrud) Get() error {
	request := oci_data_safe.ListPeerTargetDatabasesRequest{}

	if targetDatabaseId, ok := s.D.GetOkExists("target_database_id"); ok {
		tmp := targetDatabaseId.(string)
		request.TargetDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "data_safe")

	response, err := s.Client.ListPeerTargetDatabases(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListPeerTargetDatabases(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DataSafeTargetDatabasePeerTargetDatabasesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DataSafeTargetDatabasePeerTargetDatabasesDataSource-", DataSafeTargetDatabasePeerTargetDatabasesDataSource(), s.D))
	resources := []map[string]interface{}{}
	targetDatabasePeerTargetDatabase := map[string]interface{}{}

	if s.Res.CompartmentId != nil {
		targetDatabasePeerTargetDatabase["compartment_id"] = *s.Res.CompartmentId
	}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, PeerTargetDatabaseSummaryToMap(item))
	}
	targetDatabasePeerTargetDatabase["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, DataSafeTargetDatabasePeerTargetDatabasesDataSource().Schema["peer_target_database_collection"].Elem.(*schema.Resource).Schema)
		targetDatabasePeerTargetDatabase["items"] = items
	}

	resources = append(resources, targetDatabasePeerTargetDatabase)
	if err := s.D.Set("peer_target_database_collection", resources); err != nil {
		return err
	}

	return nil
}
