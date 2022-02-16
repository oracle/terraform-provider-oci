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

func DatabaseAutonomousExadataInfrastructuresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousExadataInfrastructures,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_exadata_infrastructures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseAutonomousExadataInfrastructureResource()),
			},
		},
	}
}

func readDatabaseAutonomousExadataInfrastructures(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousExadataInfrastructuresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousExadataInfrastructuresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousExadataInfrastructuresResponse
}

func (s *DatabaseAutonomousExadataInfrastructuresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousExadataInfrastructuresDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousExadataInfrastructuresRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousExadataInfrastructureSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousExadataInfrastructures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousExadataInfrastructures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousExadataInfrastructuresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousExadataInfrastructuresDataSource-", DatabaseAutonomousExadataInfrastructuresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousExadataInfrastructure := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			autonomousExadataInfrastructure["availability_domain"] = *r.AvailabilityDomain
		}

		if r.DefinedTags != nil {
			autonomousExadataInfrastructure["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousExadataInfrastructure["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			autonomousExadataInfrastructure["domain"] = *r.Domain
		}

		autonomousExadataInfrastructure["freeform_tags"] = r.FreeformTags

		if r.Hostname != nil {
			autonomousExadataInfrastructure["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			autonomousExadataInfrastructure["id"] = *r.Id
		}

		if r.LastMaintenanceRunId != nil {
			autonomousExadataInfrastructure["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		autonomousExadataInfrastructure["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousExadataInfrastructure["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			autonomousExadataInfrastructure["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			autonomousExadataInfrastructure["maintenance_window"] = nil
		}

		if r.NextMaintenanceRunId != nil {
			autonomousExadataInfrastructure["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		autonomousExadataInfrastructure["nsg_ids"] = r.NsgIds

		if r.ScanDnsName != nil {
			autonomousExadataInfrastructure["scan_dns_name"] = *r.ScanDnsName
		}

		if r.Shape != nil {
			autonomousExadataInfrastructure["shape"] = *r.Shape
		}

		autonomousExadataInfrastructure["state"] = r.LifecycleState

		if r.SubnetId != nil {
			autonomousExadataInfrastructure["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			autonomousExadataInfrastructure["time_created"] = r.TimeCreated.String()
		}

		if r.ZoneId != nil {
			autonomousExadataInfrastructure["zone_id"] = *r.ZoneId
		}

		resources = append(resources, autonomousExadataInfrastructure)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousExadataInfrastructuresDataSource().Schema["autonomous_exadata_infrastructures"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_exadata_infrastructures", resources); err != nil {
		return err
	}

	return nil
}
