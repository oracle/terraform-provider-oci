// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v39/database"
)

func init() {
	RegisterDatasource("oci_database_autonomous_vm_clusters", DatabaseAutonomousVmClustersDataSource())
}

func DatabaseAutonomousVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousVmClusters,
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
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(DatabaseAutonomousVmClusterResource()),
			},
		},
	}
}

func readDatabaseAutonomousVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
}

type DatabaseAutonomousVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousVmClustersResponse
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousVmClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousVmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceHashID("DatabaseAutonomousVmClustersDataSource-", DatabaseAutonomousVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousVmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailableCpus != nil {
			autonomousVmCluster["available_cpus"] = *r.AvailableCpus
		}

		if r.AvailableDataStorageSizeInTBs != nil {
			autonomousVmCluster["available_data_storage_size_in_tbs"] = *r.AvailableDataStorageSizeInTBs
		}

		if r.CpusEnabled != nil {
			autonomousVmCluster["cpus_enabled"] = *r.CpusEnabled
		}

		if r.DataStorageSizeInTBs != nil {
			autonomousVmCluster["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			autonomousVmCluster["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DefinedTags != nil {
			autonomousVmCluster["defined_tags"] = definedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousVmCluster["display_name"] = *r.DisplayName
		}

		if r.ExadataInfrastructureId != nil {
			autonomousVmCluster["exadata_infrastructure_id"] = *r.ExadataInfrastructureId
		}

		autonomousVmCluster["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousVmCluster["id"] = *r.Id
		}

		if r.IsLocalBackupEnabled != nil {
			autonomousVmCluster["is_local_backup_enabled"] = *r.IsLocalBackupEnabled
		}

		autonomousVmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			autonomousVmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MemorySizeInGBs != nil {
			autonomousVmCluster["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		autonomousVmCluster["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			autonomousVmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			autonomousVmCluster["time_zone"] = *r.TimeZone
		}

		if r.VmClusterNetworkId != nil {
			autonomousVmCluster["vm_cluster_network_id"] = *r.VmClusterNetworkId
		}

		resources = append(resources, autonomousVmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousVmClustersDataSource().Schema["autonomous_vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
