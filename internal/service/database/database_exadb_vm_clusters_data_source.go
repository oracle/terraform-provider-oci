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

func DatabaseExadbVmClustersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExadbVmClusters,
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
			"exascale_db_storage_vault_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadb_vm_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExadbVmClusterResource()),
			},
		},
	}
}

func readDatabaseExadbVmClusters(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadbVmClustersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadbVmClustersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExadbVmClustersResponse
}

func (s *DatabaseExadbVmClustersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadbVmClustersDataSourceCrud) Get() error {
	request := oci_database.ListExadbVmClustersRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if exascaleDbStorageVaultId, ok := s.D.GetOkExists("exascale_db_storage_vault_id"); ok {
		tmp := exascaleDbStorageVaultId.(string)
		request.ExascaleDbStorageVaultId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExadbVmClusterSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExadbVmClusters(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExadbVmClusters(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExadbVmClustersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadbVmClustersDataSource-", DatabaseExadbVmClustersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exadbVmCluster := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.AvailabilityDomain != nil {
			exadbVmCluster["availability_domain"] = *r.AvailabilityDomain
		}

		exadbVmCluster["backup_network_nsg_ids"] = r.BackupNetworkNsgIds

		if r.BackupSubnetId != nil {
			exadbVmCluster["backup_subnet_id"] = *r.BackupSubnetId
		}

		if r.ClusterName != nil {
			exadbVmCluster["cluster_name"] = *r.ClusterName
		}

		if r.DataCollectionOptions != nil {
			exadbVmCluster["data_collection_options"] = []interface{}{DataCollectionOptionsToMap(r.DataCollectionOptions)}
		} else {
			exadbVmCluster["data_collection_options"] = nil
		}

		if r.DefinedTags != nil {
			exadbVmCluster["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			exadbVmCluster["display_name"] = *r.DisplayName
		}

		if r.Domain != nil {
			exadbVmCluster["domain"] = *r.Domain
		}

		if r.ExascaleDbStorageVaultId != nil {
			exadbVmCluster["exascale_db_storage_vault_id"] = *r.ExascaleDbStorageVaultId
		}

		exadbVmCluster["freeform_tags"] = r.FreeformTags

		if r.GiVersion != nil {
			exadbVmCluster["gi_version"] = *r.GiVersion
		}

		if r.GridImageId != nil {
			exadbVmCluster["grid_image_id"] = *r.GridImageId
		}

		exadbVmCluster["grid_image_type"] = r.GridImageType

		if r.Hostname != nil {
			exadbVmCluster["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			exadbVmCluster["id"] = *r.Id
		}

		if r.LastUpdateHistoryEntryId != nil {
			exadbVmCluster["last_update_history_entry_id"] = *r.LastUpdateHistoryEntryId
		}

		exadbVmCluster["license_model"] = r.LicenseModel

		if r.LifecycleDetails != nil {
			exadbVmCluster["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.ListenerPort != nil {
			exadbVmCluster["listener_port"] = strconv.FormatInt(*r.ListenerPort, 10)
		}

		nodeConfg, nodeResourceList := getNodeConfigAndNodeListInResponse(r.Id, r.CompartmentId, r.EnabledECpuCount, r.TotalECpuCount, r.VmFileSystemStorage, r.MemorySizeInGBs, r.SnapshotFileSystemStorage, r.TotalFileSystemStorage, s.Client)
		exadbVmCluster["node_config"] = []interface{}{nodeConfg}
		exadbVmCluster["node_resource"] = nodeResourceList

		exadbVmCluster["nsg_ids"] = r.NsgIds

		if r.PrivateZoneId != nil {
			exadbVmCluster["private_zone_id"] = *r.PrivateZoneId
		}

		if r.ScanDnsName != nil {
			exadbVmCluster["scan_dns_name"] = *r.ScanDnsName
		}

		if r.ScanDnsRecordId != nil {
			exadbVmCluster["scan_dns_record_id"] = *r.ScanDnsRecordId
		}

		exadbVmCluster["scan_ip_ids"] = r.ScanIpIds

		if r.ScanListenerPortTcp != nil {
			exadbVmCluster["scan_listener_port_tcp"] = *r.ScanListenerPortTcp
		}

		if r.ScanListenerPortTcpSsl != nil {
			exadbVmCluster["scan_listener_port_tcp_ssl"] = *r.ScanListenerPortTcpSsl
		}

		exadbVmCluster["security_attributes"] = tfresource.SecurityAttributesToMap(r.SecurityAttributes)

		if r.Shape != nil {
			exadbVmCluster["shape"] = *r.Shape
		}

		exadbVmCluster["ssh_public_keys"] = r.SshPublicKeys

		exadbVmCluster["state"] = r.LifecycleState

		if r.SubnetId != nil {
			exadbVmCluster["subnet_id"] = *r.SubnetId
		}

		if r.SystemTags != nil {
			exadbVmCluster["system_tags"] = tfresource.SystemTagsToMap(r.SystemTags)
		}

		if r.SystemVersion != nil {
			exadbVmCluster["system_version"] = *r.SystemVersion
		}

		if r.TimeCreated != nil {
			exadbVmCluster["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			exadbVmCluster["time_zone"] = *r.TimeZone
		}

		exadbVmCluster["vip_ids"] = r.VipIds

		if r.ZoneId != nil {
			exadbVmCluster["zone_id"] = *r.ZoneId
		}

		resources = append(resources, exadbVmCluster)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExadbVmClustersDataSource().Schema["exadb_vm_clusters"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exadb_vm_clusters", resources); err != nil {
		return err
	}

	return nil
}
