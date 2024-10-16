// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseExadbVmClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["exadb_vm_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseExadbVmClusterResource(), fieldMap, readSingularDatabaseExadbVmCluster)
}

func readSingularDatabaseExadbVmCluster(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadbVmClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetExadbVmClusterResponse
}

func (s *DatabaseExadbVmClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadbVmClusterDataSourceCrud) Get() error {
	request := oci_database.GetExadbVmClusterRequest{}

	if exadbVmClusterId, ok := s.D.GetOkExists("exadb_vm_cluster_id"); ok {
		tmp := exadbVmClusterId.(string)
		request.ExadbVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetExadbVmCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseExadbVmClusterDataSourceCrud) SetData() error {
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

	if s.Res.ClusterName != nil {
		s.D.Set("cluster_name", *s.Res.ClusterName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataCollectionOptions != nil {
		s.D.Set("data_collection_options", []interface{}{DataCollectionOptionsToMap(s.Res.DataCollectionOptions)})
	} else {
		s.D.Set("data_collection_options", nil)
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

	if s.Res.ExascaleDbStorageVaultId != nil {
		s.D.Set("exascale_db_storage_vault_id", *s.Res.ExascaleDbStorageVaultId)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.GiVersion != nil {
		s.D.Set("gi_version", *s.Res.GiVersion)
	}

	if s.Res.GridImageId != nil {
		s.D.Set("grid_image_id", *s.Res.GridImageId)
	}

	s.D.Set("grid_image_type", s.Res.GridImageType)

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	if s.Res.IormConfigCache != nil {
		s.D.Set("iorm_config_cache", []interface{}{ExadataIormConfigToMap(s.Res.IormConfigCache)})
	} else {
		s.D.Set("iorm_config_cache", nil)
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

	nodeConfig, nodeResourceList := getNodeConfigAndNodeListInResponse(s.Res.Id, s.Res.CompartmentId, s.Res.EnabledECpuCount, s.Res.TotalECpuCount, s.Res.VmFileSystemStorage, s.Res.MemorySizeInGBs, s.Res.SnapshotFileSystemStorage, s.Res.TotalFileSystemStorage, s.Client)
	s.D.Set("node_config", []interface{}{nodeConfig})
	s.D.Set("node_resource", nodeResourceList)

	s.D.Set("nsg_ids", s.Res.NsgIds)

	if s.Res.PrivateZoneId != nil {
		s.D.Set("private_zone_id", *s.Res.PrivateZoneId)
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

	s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(s.Res.SecurityAttributes))

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("ssh_public_keys", s.Res.SshPublicKeys)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
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
