// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"strconv"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseCloudVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseCloudVmClusters,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_exadata_infrastructure_id": {
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
			"cloud_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseCloudVmClusterResource()),
			},
		},
	}
}

func readDatabaseCloudVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListCloudVmClustersResponse
}

func (s *DatabaseCloudVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListCloudVmClustersRequest{}

	if cloudExadataInfrastructureId, ok := s.D.GetOkExists("cloud_exadata_infrastructure_id"); ok {
		tmp := cloudExadataInfrastructureId.(string)
		request.CloudExadataInfrastructureId = &tmp
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
		request.LifecycleState = oci_database.CloudVmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListCloudVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseCloudVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudVmClustersDataSource-", DatabaseCloudVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cloudVmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			cloudVmCluster["availability_domain"] = *r.AvailabilityDomain
		}

		cloudVmCluster["backup_network_nsg_ids"] = r.BackupNetworkNsgIds

		if r.BackupSubnetId != nil {
			cloudVmCluster["backup_subnet_id"] = *r.BackupSubnetId
		}

		if r.CloudExadataInfrastructureId != nil {
			cloudVmCluster["cloud_exadata_infrastructure_id"] = *r.CloudExadataInfrastructureId
		}

		if r.ClusterName != nil {
			cloudVmCluster["cluster_name"] = *r.ClusterName
		}

		if r.CpuCoreCount != nil {
			cloudVmCluster["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DataStoragePercentage != nil {
			cloudVmCluster["data_storage_percentage"] = *r.DataStoragePercentage
		}

		if r.DefinedTags != nil {
			cloudVmCluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		cloudVmCluster["disk_redundancy"] = r.DiskRedundancy

		if r.DisplayName != nil {
			cloudVmCluster["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			cloudVmCluster["domain"] = *r.Domain
		}

		cloudVmCluster["freeform_tags"] = r.FreeformTags

		if r.GiVersion != nil {
			cloudVmCluster["gi_version"] = *r.GiVersion
		}

		if r.Hostname != nil {
			cloudVmCluster["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			cloudVmCluster["id"] = *r.Id
		}

		if r.IsLocalBackupEnabled != nil {
			cloudVmCluster["is_local_backup_enabled"] = *r.IsLocalBackupEnabled
		}

		if r.IsSparseDiskgroupEnabled != nil {
			cloudVmCluster["is_sparse_diskgroup_enabled"] = *r.IsSparseDiskgroupEnabled
		}

		if r.LastUpdateHistoryEntryId != nil {
			cloudVmCluster["last_update_history_entry_id"] = *r.LastUpdateHistoryEntryId
		}

		cloudVmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			cloudVmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ListenerPort != nil {
			cloudVmCluster["listener_port"] = strconv.FormatInt(*r.ListenerPort, 10)
		}

		if r.NodeCount != nil {
			cloudVmCluster["node_count"] = *r.NodeCount
		}

		cloudVmCluster["nsg_ids"] = r.NsgIds

		if r.OcpuCount != nil {
			cloudVmCluster["ocpu_count"] = *r.OcpuCount
		}

		if r.ScanDnsName != nil {
			cloudVmCluster["scan_dns_name"] = *r.ScanDnsName
		}

		if r.ScanDnsRecordId != nil {
			cloudVmCluster["scan_dns_record_id"] = *r.ScanDnsRecordId
		}

		cloudVmCluster["scan_ip_ids"] = r.ScanIpIds

		if r.ScanListenerPortTcp != nil {
			cloudVmCluster["scan_listener_port_tcp"] = *r.ScanListenerPortTcp
		}

		if r.ScanListenerPortTcpSsl != nil {
			cloudVmCluster["scan_listener_port_tcp_ssl"] = *r.ScanListenerPortTcpSsl
		}

		if r.Shape != nil {
			cloudVmCluster["shape"] = *r.Shape
		}

		cloudVmCluster["ssh_public_keys"] = r.SshPublicKeys

		cloudVmCluster["state"] = r.LifecycleState

		if r.StorageSizeInGBs != nil {
			cloudVmCluster["storage_size_in_gbs"] = *r.StorageSizeInGBs
		}

		if r.SubnetId != nil {
			cloudVmCluster["subnet_id"] = *r.SubnetId
		}

		if r.SystemVersion != nil {
			cloudVmCluster["system_version"] = *r.SystemVersion
		}

		if r.TimeCreated != nil {
			cloudVmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			cloudVmCluster["time_zone"] = *r.TimeZone
		}

		cloudVmCluster["vip_ids"] = r.VipIds

		if r.ZoneId != nil {
			cloudVmCluster["zone_id"] = *r.ZoneId
		}

		resources = append(resources, cloudVmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseCloudVmClustersDataSource().Schema["cloud_vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("cloud_vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
