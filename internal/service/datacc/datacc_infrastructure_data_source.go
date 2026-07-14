// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package datacc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DataccInfrastructureDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataccInfrastructureResource(), fieldMap, readSingularDataccInfrastructureWithContext)
}

func readSingularDataccInfrastructureWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccInfrastructureDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccInfrastructureDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.GetInfrastructureResponse
}

func (s *DataccInfrastructureDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccInfrastructureDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetInfrastructureRequest{}

	if infrastructureId, ok := s.D.GetOkExists("infrastructure_id"); ok {
		tmp := infrastructureId.(string)
		request.InfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.GetInfrastructure(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataccInfrastructureDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AcfsFileSystemStorageInGbs != nil {
		s.D.Set("acfs_file_system_storage_in_gbs", *s.Res.AcfsFileSystemStorageInGbs)
	}

	if s.Res.AcfsFileSystemUsedStorageInGbs != nil {
		s.D.Set("acfs_file_system_used_storage_in_gbs", *s.Res.AcfsFileSystemUsedStorageInGbs)
	}

	if s.Res.AdminNetworkcidr != nil {
		s.D.Set("admin_networkcidr", *s.Res.AdminNetworkcidr)
	}

	s.D.Set("backup_network_bonding_interface", s.Res.BackupNetworkBondingInterface)

	s.D.Set("backup_network_bonding_mode", s.Res.BackupNetworkBondingMode)

	s.D.Set("client_network_bonding_interface", s.Res.ClientNetworkBondingInterface)

	s.D.Set("client_network_bonding_mode", s.Res.ClientNetworkBondingMode)

	if s.Res.CloudControlPlaneServer1 != nil {
		s.D.Set("cloud_control_plane_server1", *s.Res.CloudControlPlaneServer1)
	}

	if s.Res.CloudControlPlaneServer2 != nil {
		s.D.Set("cloud_control_plane_server2", *s.Res.CloudControlPlaneServer2)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeCapacity != nil {
		s.D.Set("compute_capacity", []interface{}{ComputeCapacityDetailsToMap(s.Res.ComputeCapacity)})
	} else {
		s.D.Set("compute_capacity", nil)
	}

	contacts := []interface{}{}
	for _, item := range s.Res.Contacts {
		contacts = append(contacts, InfrastructureContactToMap(item))
	}
	s.D.Set("contacts", contacts)

	if s.Res.CorporateProxy != nil {
		s.D.Set("corporate_proxy", *s.Res.CorporateProxy)
	}

	s.D.Set("cps_network_bonding_interface", s.Res.CpsNetworkBondingInterface)

	s.D.Set("cps_network_bonding_mode", s.Res.CpsNetworkBondingMode)

	if s.Res.DataDiskPercentage != nil {
		s.D.Set("data_disk_percentage", *s.Res.DataDiskPercentage)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_servers", s.Res.DnsServers)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.Gateway != nil {
		s.D.Set("gateway", *s.Res.Gateway)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	if s.Res.MaintenanceWindow != nil {
		s.D.Set("maintenance_window", []interface{}{MaintenanceWindowToMap(s.Res.MaintenanceWindow)})
	} else {
		s.D.Set("maintenance_window", nil)
	}

	if s.Res.Netmask != nil {
		s.D.Set("netmask", *s.Res.Netmask)
	}

	if s.Res.NetworkAdapterConfiguration != nil {
		s.D.Set("network_adapter_configuration", *s.Res.NetworkAdapterConfiguration)
	}

	s.D.Set("ntp_servers", s.Res.NtpServers)

	if s.Res.RackSerialNumber != nil {
		s.D.Set("rack_serial_number", *s.Res.RackSerialNumber)
	}

	if s.Res.RecoDiskPercentage != nil {
		s.D.Set("reco_disk_percentage", *s.Res.RecoDiskPercentage)
	}

	servers := []interface{}{}
	for _, item := range s.Res.Servers {
		servers = append(servers, InfrastructureServerToMap(item))
	}
	s.D.Set("servers", servers)

	s.D.Set("shape", s.Res.Shape)

	s.D.Set("ssd_configuration_requested", s.Res.SsdConfigurationRequested)

	s.D.Set("state", s.Res.LifecycleState)

	storageCapacity := []interface{}{}
	for _, item := range s.Res.StorageCapacity {
		storageCapacity = append(storageCapacity, StorageCapacityDetailsToMap(item))
	}
	s.D.Set("storage_capacity", storageCapacity)

	if s.Res.SubscriptionPlanNumber != nil {
		s.D.Set("subscription_plan_number", *s.Res.SubscriptionPlanNumber)
	}

	s.D.Set("system_model", s.Res.SystemModel)

	if s.Res.SystemStorageCapacity != nil {
		s.D.Set("system_storage_capacity", []interface{}{SystemStorageCapacityDetailsToMap(s.Res.SystemStorageCapacity)})
	} else {
		s.D.Set("system_storage_capacity", nil)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeActivated != nil {
		s.D.Set("time_activated", s.Res.TimeActivated.String())
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastStateUpdated != nil {
		s.D.Set("time_last_state_updated", s.Res.TimeLastStateUpdated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TimeValidated != nil {
		s.D.Set("time_validated", s.Res.TimeValidated.String())
	}

	if s.Res.Version != nil {
		s.D.Set("version", *s.Res.Version)
	}

	if s.Res.VlanId != nil {
		s.D.Set("vlan_id", *s.Res.VlanId)
	}

	return nil
}
