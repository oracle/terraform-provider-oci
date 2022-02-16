// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseAutonomousDatabaseDataguardAssociationDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousDatabaseDataguardAssociation,
		Schema: map[string]*schema.Schema{
			"autonomous_database_dataguard_association_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"autonomous_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"apply_lag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apply_rate": {
				Type:     schema.TypeString,
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
		},
	}
}

func readSingularDatabaseAutonomousDatabaseDataguardAssociation(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousDatabaseDataguardAssociationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousDatabaseDataguardAssociationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousDatabaseDataguardAssociationResponse
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousDatabaseDataguardAssociationRequest{}

	if autonomousDatabaseDataguardAssociationId, ok := s.D.GetOkExists("autonomous_database_dataguard_association_id"); ok {
		tmp := autonomousDatabaseDataguardAssociationId.(string)
		request.AutonomousDatabaseDataguardAssociationId = &tmp
	}

	if autonomousDatabaseId, ok := s.D.GetOkExists("autonomous_database_id"); ok {
		tmp := autonomousDatabaseId.(string)
		request.AutonomousDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousDatabaseDataguardAssociation(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousDatabaseDataguardAssociationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ApplyLag != nil {
		s.D.Set("apply_lag", *s.Res.ApplyLag)
	}

	if s.Res.ApplyRate != nil {
		s.D.Set("apply_rate", *s.Res.ApplyRate)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.PeerAutonomousDatabaseId != nil {
		s.D.Set("peer_autonomous_database_id", *s.Res.PeerAutonomousDatabaseId)
	}

	s.D.Set("peer_autonomous_database_life_cycle_state", s.Res.PeerAutonomousDatabaseLifeCycleState)

	s.D.Set("peer_role", s.Res.PeerRole)

	s.D.Set("protection_mode", s.Res.ProtectionMode)

	s.D.Set("role", s.Res.Role)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastRoleChanged != nil {
		s.D.Set("time_last_role_changed", s.Res.TimeLastRoleChanged.String())
	}

	return nil
}
