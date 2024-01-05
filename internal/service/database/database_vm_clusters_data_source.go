// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
			"vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseVmClusterResource()),
			},
		},
	}
}

func readDatabaseVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClustersResponse
}

func (s *DatabaseVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListVmClustersRequest{}

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
		request.LifecycleState = oci_database.VmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClustersDataSource-", DatabaseVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			vmCluster["availability_domain"] = *r.AvailabilityDomain
		}

		if r.CpusEnabled != nil {
			vmCluster["cpus_enabled"] = *r.CpusEnabled
			vmCluster["cpu_core_count"] = *r.CpusEnabled
		}

		if r.DataCollectionOptions != nil {
			vmCluster["data_collection_options"] = []interface{}{DataCollectionOptionsToMap(r.DataCollectionOptions)}
		} else {
			vmCluster["data_collection_options"] = nil
		}

		if r.DataStorageSizeInGBs != nil {
			vmCluster["data_storage_size_in_gb"] = *r.DataStorageSizeInGBs
		}

		if r.DataStorageSizeInTBs != nil {
			vmCluster["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			vmCluster["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		vmCluster["db_servers"] = r.DbServers

		if r.DefinedTags != nil {
			vmCluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vmCluster["display_name"] = *r.DisplayName
		}

		if r.ExadataInfrastructureId != nil {
			vmCluster["exadata_infrastructure_id"] = *r.ExadataInfrastructureId
		}

		vmCluster["freeform_tags"] = r.FreeformTags

		if r.GiVersion != nil {
			vmCluster["gi_version"] = *r.GiVersion
		}

		if r.Id != nil {
			vmCluster["id"] = *r.Id
		}

		if r.IsLocalBackupEnabled != nil {
			vmCluster["is_local_backup_enabled"] = *r.IsLocalBackupEnabled
		}

		if r.IsSparseDiskgroupEnabled != nil {
			vmCluster["is_sparse_diskgroup_enabled"] = *r.IsSparseDiskgroupEnabled
		}

		if r.LastPatchHistoryEntryId != nil {
			vmCluster["last_patch_history_entry_id"] = *r.LastPatchHistoryEntryId
		}

		vmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			vmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MemorySizeInGBs != nil {
			vmCluster["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.OcpusEnabled != nil {
			vmCluster["ocpus_enabled"] = *r.OcpusEnabled
		}

		if r.Shape != nil {
			vmCluster["shape"] = *r.Shape
		}

		vmCluster["ssh_public_keys"] = r.SshPublicKeys

		vmCluster["state"] = r.LifecycleState

		if r.SystemVersion != nil {
			vmCluster["system_version"] = *r.SystemVersion
		}

		if r.TimeCreated != nil {
			vmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			vmCluster["time_zone"] = *r.TimeZone
		}

		if r.VmClusterNetworkId != nil {
			vmCluster["vm_cluster_network_id"] = *r.VmClusterNetworkId
		}

		resources = append(resources, vmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClustersDataSource().Schema["vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
