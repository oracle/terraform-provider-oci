// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabaseDataguardAssociations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_container_database_dataguard_associations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousContainerDatabaseDataguardAssociationResource()),
			},
		},
	}
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousContainerDatabaseDataguardAssociationsResponse
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousContainerDatabaseDataguardAssociationsRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousContainerDatabaseDataguardAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousContainerDatabaseDataguardAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource-", DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousContainerDatabaseDataguardAssociation := map[string]interface{}{
			"autonomous_container_database_id": *r.AutonomousContainerDatabaseId,
		}

		if r.ApplyLag != nil {
			autonomousContainerDatabaseDataguardAssociation["apply_lag"] = *r.ApplyLag
		}

		if r.ApplyRate != nil {
			autonomousContainerDatabaseDataguardAssociation["apply_rate"] = *r.ApplyRate
		}

		if r.Id != nil {
			autonomousContainerDatabaseDataguardAssociation["id"] = *r.Id
		}

		if r.IsAutomaticFailoverEnabled != nil {
			autonomousContainerDatabaseDataguardAssociation["is_automatic_failover_enabled"] = *r.IsAutomaticFailoverEnabled
		}

		if r.LifecycleDetails != nil {
			autonomousContainerDatabaseDataguardAssociation["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PeerAutonomousContainerDatabaseDataguardAssociationId != nil {
			autonomousContainerDatabaseDataguardAssociation["peer_autonomous_container_database_dataguard_association_id"] = *r.PeerAutonomousContainerDatabaseDataguardAssociationId
		}

		if r.PeerAutonomousContainerDatabaseId != nil {
			autonomousContainerDatabaseDataguardAssociation["peer_autonomous_container_database_id"] = *r.PeerAutonomousContainerDatabaseId
		}

		autonomousContainerDatabaseDataguardAssociation["peer_lifecycle_state"] = r.PeerLifecycleState

		autonomousContainerDatabaseDataguardAssociation["peer_role"] = r.PeerRole

		autonomousContainerDatabaseDataguardAssociation["protection_mode"] = r.ProtectionMode

		autonomousContainerDatabaseDataguardAssociation["role"] = r.Role

		autonomousContainerDatabaseDataguardAssociation["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousContainerDatabaseDataguardAssociation["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastRoleChanged != nil {
			autonomousContainerDatabaseDataguardAssociation["time_last_role_changed"] = r.TimeLastRoleChanged.String()
		}

		if r.TimeLastSynced != nil {
			autonomousContainerDatabaseDataguardAssociation["time_last_synced"] = r.TimeLastSynced.String()
		}

		if r.TransportLag != nil {
			autonomousContainerDatabaseDataguardAssociation["transport_lag"] = *r.TransportLag
		}

		resources = append(resources, autonomousContainerDatabaseDataguardAssociation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource().Schema["autonomous_container_database_dataguard_associations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_database_dataguard_associations", resources); err != nil {
		return err
	}

	return nil
}
