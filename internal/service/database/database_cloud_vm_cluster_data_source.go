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

func DatabaseCloudVmClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cloud_vm_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseCloudVmClusterResource(), fieldMap, readSingularDatabaseCloudVmCluster)
}

func readSingularDatabaseCloudVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudVmClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudVmClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetCloudVmClusterResponse
}

func (s *DatabaseCloudVmClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudVmClusterDataSourceCrud) Get() error {
	request := oci_database.GetCloudVmClusterRequest{}

	if cloudVmClusterId, ok := s.D.GetOkExists("cloud_vm_cluster_id"); ok {
		tmp := cloudVmClusterId.(string)
		request.CloudVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetCloudVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseCloudVmClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	s.D.Set("backup_network_nsg_ids", s.Res.BackupNetworkNsgIds)

	if s.Res.BackupSubnetId != nil {
		s.D.Set("backup_subnet_id", *s.Res.BackupSubnetId)
	}

	if s.Res.CloudExadataInfrastructureId != nil {
		s.D.Set("cloud_exadata_infrastructure_id", *s.Res.CloudExadataInfrastructureId)
	}

	if s.Res.ClusterName != nil {
		s.D.Set("cluster_name", *s.Res.ClusterName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DataStoragePercentage != nil {
		s.D.Set("data_storage_percentage", *s.Res.DataStoragePercentage)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("disk_redundancy", s.Res.DiskRedundancy)

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Domain != nil {
		s.D.Set("domain", *s.Res.Domain)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GiVersion != nil {
		s.D.Set("gi_version", *s.Res.GiVersion)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{ExadataIormConfigToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", nil)
	}

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	if s.Res.IsSparseDiskgroupEnabled != nil {
		s.D.Set("is_sparse_diskgroup_enabled", *s.Res.IsSparseDiskgroupEnabled)
	}

	if s.Res.LastUpdateHistoryEntryId != nil {
		s.D.Set("last_update_history_entry_id", *s.Res.LastUpdateHistoryEntryId)
	}

	s.D.Set("license_model", s.Res.LicenseModel)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", strconv.FormatInt(*s.Res.ListenerPort, 10))
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.OcpuCount != nil {
		s.D.Set("ocpu_count", *s.Res.OcpuCount)
	}

	if s.Res.ScanDnsName != nil {
		s.D.Set("scan_dns_name", *s.Res.ScanDnsName)
	}

	if s.Res.ScanDnsRecordId != nil {
		s.D.Set("scan_dns_record_id", *s.Res.ScanDnsRecordId)
	}

	s.D.Set("scan_ip_ids", s.Res.ScanIpIds)

	if s.Res.ScanListenerPortTcp != nil {
		s.D.Set("scan_listener_port_tcp", *s.Res.ScanListenerPortTcp)
	}

	if s.Res.ScanListenerPortTcpSsl != nil {
		s.D.Set("scan_listener_port_tcp_ssl", *s.Res.ScanListenerPortTcpSsl)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.StorageSizeInGBs != nil {
		s.D.Set("storage_size_in_gbs", *s.Res.StorageSizeInGBs)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemVersion != nil {
		s.D.Set("system_version", *s.Res.SystemVersion)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeZone != nil {
		s.D.Set("time_zone", *s.Res.TimeZone)
	}

	s.D.Set("vip_ids", s.Res.VipIds)

	if s.Res.ZoneId != nil {
		s.D.Set("zone_id", *s.Res.ZoneId)
	}

	return nil
}
