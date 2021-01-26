// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v34/database"
)

func init() {
	RegisterDatasource("oci_database_cloud_exadata_infrastructures", DatabaseCloudExadataInfrastructuresDataSource())
}

func DatabaseCloudExadataInfrastructuresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseCloudExadataInfrastructures,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
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
			"cloud_exadata_infrastructures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseCloudExadataInfrastructureResource()),
			},
		},
	}
}

func readDatabaseCloudExadataInfrastructures(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudExadataInfrastructuresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseCloudExadataInfrastructuresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListCloudExadataInfrastructuresResponse
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) Get() error {
	request := oci_database.ListCloudExadataInfrastructuresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.CloudExadataInfrastructureSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListCloudExadataInfrastructures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudExadataInfrastructures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseCloudExadataInfrastructuresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseCloudExadataInfrastructuresDataSource-", DatabaseCloudExadataInfrastructuresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cloudExadataInfrastructure := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			cloudExadataInfrastructure["availability_domain"] = *r.AvailabilityDomain
		}

		if r.AvailableStorageSizeInGBs != nil {
			cloudExadataInfrastructure["available_storage_size_in_gbs"] = *r.AvailableStorageSizeInGBs
		}

		if r.ComputeCount != nil {
			cloudExadataInfrastructure["compute_count"] = *r.ComputeCount
		}

		if r.DefinedTags != nil {
			cloudExadataInfrastructure["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			cloudExadataInfrastructure["display_name"] = *r.DisplayName
		}

		cloudExadataInfrastructure["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			cloudExadataInfrastructure["id"] = *r.Id
		}

		if r.LastMaintenanceRunId != nil {
			cloudExadataInfrastructure["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LifecycleDetails != nil {
			cloudExadataInfrastructure["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaintenanceWindow != nil {
			cloudExadataInfrastructure["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			cloudExadataInfrastructure["maintenance_window"] = nil
		}

		if r.NextMaintenanceRunId != nil {
			cloudExadataInfrastructure["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.Shape != nil {
			cloudExadataInfrastructure["shape"] = *r.Shape
		}

		cloudExadataInfrastructure["state"] = r.LifecycleState

		if r.StorageCount != nil {
			cloudExadataInfrastructure["storage_count"] = *r.StorageCount
		}

		if r.TimeCreated != nil {
			cloudExadataInfrastructure["time_created"] = r.TimeCreated.String()
		}

		if r.TotalStorageSizeInGBs != nil {
			cloudExadataInfrastructure["total_storage_size_in_gbs"] = *r.TotalStorageSizeInGBs
		}

		resources = append(resources, cloudExadataInfrastructure)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseCloudExadataInfrastructuresDataSource().Schema["cloud_exadata_infrastructures"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cloud_exadata_infrastructures", resources); err != nil {
		return err
	}

	return nil
}
