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

func DatabaseAutonomousDatabaseDataguardAssociationsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousDatabaseDataguardAssociations,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_database_dataguard_associations": {
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
						"autonomous_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_automatic_failover_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_autonomous_database_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"peer_autonomous_database_life_cycle_state": {
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
						"time_last_synced": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"transport_lag": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseAutonomousDatabaseDataguardAssociations(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseDataguardAssociationsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseDataguardAssociationsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousDatabaseDataguardAssociationsResponse
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationsDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousDatabaseDataguardAssociationsRequest{}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousDatabaseDataguardAssociations(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousDatabaseDataguardAssociations(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousDatabaseDataguardAssociationsDataSource-", DatabaseAutonomousDatabaseDataguardAssociationsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousDatabaseDataguardAssociation := map[string]interface{}{
			"autonomous_database_id": *r.AutonomousDatabaseId,
		}

		if r.ApplyLag != nil {
			autonomousDatabaseDataguardAssociation["apply_lag"] = *r.ApplyLag
		}

		if r.ApplyRate != nil {
			autonomousDatabaseDataguardAssociation["apply_rate"] = *r.ApplyRate
		}

		if r.Id != nil {
			autonomousDatabaseDataguardAssociation["id"] = *r.Id
		}

		if r.IsAutomaticFailoverEnabled != nil {
			autonomousDatabaseDataguardAssociation["is_automatic_failover_enabled"] = *r.IsAutomaticFailoverEnabled
		}

		if r.LifecycleDetails != nil {
			autonomousDatabaseDataguardAssociation["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.PeerAutonomousDatabaseId != nil {
			autonomousDatabaseDataguardAssociation["peer_autonomous_database_id"] = *r.PeerAutonomousDatabaseId
		}

		autonomousDatabaseDataguardAssociation["peer_autonomous_database_life_cycle_state"] = r.PeerAutonomousDatabaseLifeCycleState

		autonomousDatabaseDataguardAssociation["peer_role"] = r.PeerRole

		autonomousDatabaseDataguardAssociation["protection_mode"] = r.ProtectionMode

		autonomousDatabaseDataguardAssociation["role"] = r.Role

		autonomousDatabaseDataguardAssociation["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousDatabaseDataguardAssociation["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastRoleChanged != nil {
			autonomousDatabaseDataguardAssociation["time_last_role_changed"] = r.TimeLastRoleChanged.String()
		}

		if r.TimeLastSynced != nil {
			autonomousDatabaseDataguardAssociation["time_last_synced"] = r.TimeLastSynced.String()
		}

		if r.TransportLag != nil {
			autonomousDatabaseDataguardAssociation["transport_lag"] = *r.TransportLag
		}

		resources = append(resources, autonomousDatabaseDataguardAssociation)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousDatabaseDataguardAssociationsDataSource().Schema["autonomous_database_dataguard_associations"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_database_dataguard_associations", resources); err != nil {
		return err
	}

	return nil
}
