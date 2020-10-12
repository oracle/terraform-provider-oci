// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v27/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_container_database_dataguard_associations", DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource())
}

func DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousContainerDatabaseDataguardAssociations,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_container_database_dataguard_associations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"apply_lag": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"apply_rate": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_autonomous_container_database_dataguard_association_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_autonomous_container_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_lifecycle_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"protection_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"role": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_last_role_changed": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousContainerDatabaseDataguardAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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

	s.D.SetId(GenerateDataSourceHashID("DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource-", DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource(), s.D))
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

		resources = append(resources, autonomousContainerDatabaseDataguardAssociation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousContainerDatabaseDataguardAssociationsDataSource().Schema["autonomous_container_database_dataguard_associations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_database_dataguard_associations", resources); err != nil {
		return err
	}

	return nil
}
