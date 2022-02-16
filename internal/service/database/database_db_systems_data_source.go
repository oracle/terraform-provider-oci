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

func DatabaseDbSystemsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbSystems,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"backup_id": {
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
			"db_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseDbSystemResource()),
			},
		},
	}
}

func readDatabaseDbSystems(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbSystemsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbSystemsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbSystemsResponse
}

func (s *DatabaseDbSystemsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbSystemsDataSourceCrud) Get() error {
	request := oci_database.ListDbSystemsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if backupId, ok := s.D.GetOkExists("backup_id"); ok {
		tmp := backupId.(string)
		request.BackupId = &tmp
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
		request.LifecycleState = oci_database.DbSystemSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbSystems(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbSystems(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbSystemsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbSystemsDataSource-", DatabaseDbSystemsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbSystem := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			dbSystem["availability_domain"] = *r.AvailabilityDomain
		}

		dbSystem["backup_network_nsg_ids"] = r.BackupNetworkNsgIds

		if r.BackupSubnetId != nil {
			dbSystem["backup_subnet_id"] = *r.BackupSubnetId
		}

		if r.ClusterName != nil {
			dbSystem["cluster_name"] = *r.ClusterName
		}

		if r.CpuCoreCount != nil {
			dbSystem["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DataStoragePercentage != nil {
			dbSystem["data_storage_percentage"] = *r.DataStoragePercentage
		}

		if r.DataStorageSizeInGBs != nil {
			dbSystem["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		dbSystem["database_edition"] = r.DatabaseEdition

		if r.DbSystemOptions != nil {
			dbSystem["db_system_options"] = []interface{}{DbSystemOptionsToMap(r.DbSystemOptions)}
		} else {
			dbSystem["db_system_options"] = nil
		}

		if r.DefinedTags != nil {
			dbSystem["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		dbSystem["disk_redundancy"] = r.DiskRedundancy

		if r.DisplayName != nil {
			dbSystem["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			dbSystem["domain"] = *r.Domain
		}

		dbSystem["fault_domains"] = r.FaultDomains

		dbSystem["freeform_tags"] = r.FreeformTags

		if r.Hostname != nil {
			dbSystem["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbSystem["id"] = *r.Id
		}

		if r.KmsKeyId != nil {
			dbSystem["kms_key_id"] = *r.KmsKeyId
		}

		if r.LastMaintenanceRunId != nil {
			dbSystem["last_maintenance_run_id"] = *r.LastMaintenanceRunId
		}

		if r.LastPatchHistoryEntryId != nil {
			dbSystem["last_patch_history_entry_id"] = *r.LastPatchHistoryEntryId
		}

		dbSystem["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			dbSystem["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ListenerPort != nil {
			dbSystem["listener_port"] = *r.ListenerPort
		}

		if r.MaintenanceWindow != nil {
			dbSystem["maintenance_window"] = []interface{}{MaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			dbSystem["maintenance_window"] = nil
		}

		if r.NextMaintenanceRunId != nil {
			dbSystem["next_maintenance_run_id"] = *r.NextMaintenanceRunId
		}

		if r.NodeCount != nil {
			dbSystem["node_count"] = *r.NodeCount
		}

		dbSystem["nsg_ids"] = r.NsgIds

		if r.PointInTimeDataDiskCloneTimestamp != nil {
			dbSystem["point_in_time_data_disk_clone_timestamp"] = r.PointInTimeDataDiskCloneTimestamp.String()
		}

		if r.RecoStorageSizeInGB != nil {
			dbSystem["reco_storage_size_in_gb"] = *r.RecoStorageSizeInGB
		}

		if r.ScanDnsName != nil {
			dbSystem["scan_dns_name"] = *r.ScanDnsName
		}

		if r.ScanDnsRecordId != nil {
			dbSystem["scan_dns_record_id"] = *r.ScanDnsRecordId
		}

		dbSystem["scan_ip_ids"] = r.ScanIpIds

		if r.Shape != nil {
			dbSystem["shape"] = *r.Shape
		}

		if r.SourceDbSystemId != nil {
			dbSystem["source_db_system_id"] = *r.SourceDbSystemId
		}

		if r.SparseDiskgroup != nil {
			dbSystem["sparse_diskgroup"] = *r.SparseDiskgroup
		}

		dbSystem["ssh_public_keys"] = r.SshPublicKeys

		dbSystem["state"] = r.LifecycleState

		if r.SubnetId != nil {
			dbSystem["subnet_id"] = *r.SubnetId
		}

		if r.TimeCreated != nil {
			dbSystem["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			dbSystem["time_zone"] = *r.TimeZone
		}

		if r.Version != nil {
			dbSystem["version"] = *r.Version
		}

		dbSystem["vip_ids"] = r.VipIds

		if r.ZoneId != nil {
			dbSystem["zone_id"] = *r.ZoneId
		}

		resources = append(resources, dbSystem)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbSystemsDataSource().Schema["db_systems"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_systems", resources); err != nil {
		return err
	}

	return nil
}
