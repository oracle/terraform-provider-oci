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

func DatabaseAutonomousExadataInfrastructureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["autonomous_exadata_infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseAutonomousExadataInfrastructureResource(), fieldMap, readSingularDatabaseAutonomousExadataInfrastructure)
}

func readSingularDatabaseAutonomousExadataInfrastructure(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousExadataInfrastructureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousExadataInfrastructureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousExadataInfrastructureResponse
}

func (s *DatabaseAutonomousExadataInfrastructureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousExadataInfrastructureDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousExadataInfrastructureRequest{}

	if autonomousExadataInfrastructureId, ok := s.D.GetOkExists("autonomous_exadata_infrastructure_id"); ok {
		tmp := autonomousExadataInfrastructureId.(string)
		request.AutonomousExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousExadataInfrastructure(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousExadataInfrastructureDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.LastMaintenanceRunId != nil {
		s.D.Set("last_maintenance_run_id", *s.Res.LastMaintenanceRunId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.NextMaintenanceRunId != nil {
		s.D.Set("next_maintenance_run_id", *s.Res.NextMaintenanceRunId)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.ScanDnsName != nil {
		s.D.Set("scan_dns_name", *s.Res.ScanDnsName)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.ZoneId != nil {
		s.D.Set("zone_id", *s.Res.ZoneId)
	}

	return nil
}
