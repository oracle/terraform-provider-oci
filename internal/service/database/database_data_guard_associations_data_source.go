// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseDataGuardAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDataGuardAssociations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"data_guard_associations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseDataGuardAssociationResource()),
			},
		},
	}
}

func readDatabaseDataGuardAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDataGuardAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDataGuardAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDataGuardAssociationsResponse
}

func (s *DatabaseDataGuardAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDataGuardAssociationsDataSourceCrud) Get() error {
	request := oci_database.ListDataGuardAssociationsRequest{}

	if databaseId, ok := s.D.GetOkExists("database_id"); ok {
		tmp := databaseId.(string)
		request.DatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDataGuardAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDataGuardAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDataGuardAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDataGuardAssociationsDataSource-", DatabaseDataGuardAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dataGuardAssociation := map[string]interface{}{
			"database_id": *r.DatabaseId,
		}

		if r.ApplyLag != nil {
			dataGuardAssociation["apply_lag"] = *r.ApplyLag
		}

		if r.ApplyRate != nil {
			dataGuardAssociation["apply_rate"] = *r.ApplyRate
		}

		if r.Id != nil {
			dataGuardAssociation["id"] = *r.Id
		}

		if r.IsActiveDataGuardEnabled != nil {
			dataGuardAssociation["is_active_data_guard_enabled"] = *r.IsActiveDataGuardEnabled
		}

		if r.LifecycleDetails != nil {
			dataGuardAssociation["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PeerDataGuardAssociationId != nil {
			dataGuardAssociation["peer_data_guard_association_id"] = *r.PeerDataGuardAssociationId
		}

		if r.PeerDatabaseId != nil {
			dataGuardAssociation["peer_database_id"] = *r.PeerDatabaseId
		}

		if r.PeerDbHomeId != nil {
			dataGuardAssociation["peer_db_home_id"] = *r.PeerDbHomeId
		}

		if r.PeerDbSystemId != nil {
			dataGuardAssociation["peer_db_system_id"] = *r.PeerDbSystemId
		}

		dataGuardAssociation["peer_role"] = r.PeerRole

		dataGuardAssociation["protection_mode"] = r.ProtectionMode

		dataGuardAssociation["role"] = r.Role

		dataGuardAssociation["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dataGuardAssociation["time_created"] = r.TimeCreated.String()
		}

		dataGuardAssociation["transport_type"] = r.TransportType

		resources = append(resources, dataGuardAssociation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDataGuardAssociationsDataSource().Schema["data_guard_associations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("data_guard_associations", resources); err != nil {
		return err
	}

	return nil
}
