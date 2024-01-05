// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpSddcDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["sddc_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OcvpSddcResource(), fieldMap, readSingularOcvpSddc)
}

func readSingularOcvpSddc(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpSddcDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).SddcClient()
	sync.ClusterClient = m.(*client.OracleClients).ClusterClient()
	return tfresource.ReadResource(sync)
}

type OcvpSddcDataSourceCrud struct {
	D             *schema.ResourceData
	Client        *oci_ocvp.SddcClient
	ClusterClient *oci_ocvp.ClusterClient
	Res           *oci_ocvp.GetSddcResponse
}

func (s *OcvpSddcDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpSddcDataSourceCrud) Get() error {
	request := oci_ocvp.GetSddcRequest{}

	if sddcId, ok := s.D.GetOkExists("sddc_id"); ok {
		tmp := sddcId.(string)
		request.SddcId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetSddc(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpSddcDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.ClustersCount != nil {
		s.D.Set("clusters_count", *s.Res.ClustersCount)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	actualEsxiHostCount, err := CalculateActualEsxiHostCount(s.Res.Id, s.Res.CompartmentId, s.ClusterClient)
	if err != nil {
		return nil
	}
	s.D.Set("esxi_hosts_count", actualEsxiHostCount)
	s.D.Set("actual_esxi_hosts_count", actualEsxiHostCount)

	if s.Res.InitialConfiguration != nil {
		s.D.Set("initial_configuration", []interface{}{InitialConfigurationToMap(s.Res.InitialConfiguration,
			s.D.GetOk, s.D.HasChange, s.D.GetChange, true)})
	} else {
		s.D.Set("initial_configuration", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EsxiSoftwareVersion != nil {
		s.D.Set("esxi_software_version", *s.Res.EsxiSoftwareVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.HcxFqdn != nil {
		s.D.Set("hcx_fqdn", *s.Res.HcxFqdn)
	}

	s.D.Set("hcx_mode", s.Res.HcxMode)
	switch s.Res.HcxMode {
	case oci_ocvp.HcxModesDisabled:
		s.D.Set("is_hcx_enabled", false)
		s.D.Set("is_hcx_enterprise_enabled", false)

	case oci_ocvp.HcxModesAdvanced:
		s.D.Set("is_hcx_enabled", true)
		s.D.Set("is_hcx_enterprise_enabled", false)

	case oci_ocvp.HcxModesEnterprise:
		s.D.Set("is_hcx_enabled", true)
		s.D.Set("is_hcx_enterprise_enabled", true)
	}

	hcxOnPremLicenses := []interface{}{}
	for _, item := range s.Res.HcxOnPremLicenses {
		hcxOnPremLicenses = append(hcxOnPremLicenses, HcxLicenseSummaryToMap(item))
	}
	s.D.Set("hcx_on_prem_licenses", hcxOnPremLicenses)

	if len(s.Res.HcxOnPremLicenses) > 0 {
		s.D.Set("hcx_on_prem_key", s.Res.HcxOnPremLicenses[0].ActivationKey)
	}

	s.D.Set("nsx_overlay_segment_name", "WORKLOAD")

	if s.Res.HcxPrivateIpId != nil {
		s.D.Set("hcx_private_ip_id", *s.Res.HcxPrivateIpId)
	}

	if s.Res.IsHcxPendingDowngrade != nil {
		s.D.Set("is_hcx_pending_downgrade", *s.Res.IsHcxPendingDowngrade)
	}

	if s.Res.IsSingleHostSddc != nil {
		s.D.Set("is_single_host_sddc", *s.Res.IsSingleHostSddc)
	}

	if s.Res.NsxEdgeUplinkIpId != nil {
		s.D.Set("nsx_edge_uplink_ip_id", *s.Res.NsxEdgeUplinkIpId)
	}

	if s.Res.NsxManagerFqdn != nil {
		s.D.Set("nsx_manager_fqdn", *s.Res.NsxManagerFqdn)
	}

	if s.Res.NsxManagerPrivateIpId != nil {
		s.D.Set("nsx_manager_private_ip_id", *s.Res.NsxManagerPrivateIpId)
	}

	if s.Res.NsxManagerUsername != nil {
		s.D.Set("nsx_manager_username", *s.Res.NsxManagerUsername)
	}

	if s.Res.SshAuthorizedKeys != nil {
		s.D.Set("ssh_authorized_keys", *s.Res.SshAuthorizedKeys)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeHcxBillingCycleEnd != nil {
		s.D.Set("time_hcx_billing_cycle_end", s.Res.TimeHcxBillingCycleEnd.String())
	}

	if s.Res.TimeHcxLicenseStatusUpdated != nil {
		s.D.Set("time_hcx_license_status_updated", s.Res.TimeHcxLicenseStatusUpdated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VcenterFqdn != nil {
		s.D.Set("vcenter_fqdn", *s.Res.VcenterFqdn)
	}

	if s.Res.VcenterPrivateIpId != nil {
		s.D.Set("vcenter_private_ip_id", *s.Res.VcenterPrivateIpId)
	}

	if s.Res.VcenterUsername != nil {
		s.D.Set("vcenter_username", *s.Res.VcenterUsername)
	}

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	if s.Res.HcxMode != oci_ocvp.HcxModesDisabled {
		hcxPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeHcx)
		if err != nil {
			return err
		}
		if hcxPassword != nil {
			s.D.Set("hcx_initial_password", *hcxPassword)
		}
	}

	nsxPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeNsx)
	if err != nil {
		return err
	}
	if nsxPassword != nil {
		s.D.Set("nsx_manager_initial_password", *nsxPassword)
	}

	vCenterPassword, err := GetSddcPassword(s.Client, s.D.Id(), oci_ocvp.RetrievePasswordTypeVcenter)
	if err != nil {
		return err
	}
	if vCenterPassword != nil {
		s.D.Set("vcenter_initial_password", *vCenterPassword)
	}

	err = s.SetDataClusterValues(s.Res.Id, s.Res.CompartmentId, s.ClusterClient)

	if err != nil {
		return err
	}

	return nil
}

func (s *OcvpSddcDataSourceCrud) SetDataClusterValues(sddcId *string, compartmentId *string, clusterClient *oci_ocvp.ClusterClient) error {
	clusterSummary, err := GetManagementClusterSummary(sddcId, compartmentId, clusterClient)
	if err != nil {
		return err
	}
	clusterId := clusterSummary.Id
	log.Printf("[DEBUG] setting values from cluster %s", *clusterId)

	req := oci_ocvp.GetClusterRequest{}
	req.ClusterId = clusterId
	clusterResponse, err := clusterClient.GetCluster(context.Background(), req)

	if err != nil {
		log.Printf("[ERROR] failed to get cluster id : '%s'", *clusterId)
		return err
	}
	log.Printf("[DEBUG] setting vshere upgrade objects")
	vsphereUpgradeObjects := []interface{}{}
	for _, item := range clusterResponse.VsphereUpgradeObjects {
		vsphereUpgradeObjects = append(vsphereUpgradeObjects, VsphereUpgradeObjectToMap(item))
	}
	err = s.D.Set("vsphere_upgrade_objects", vsphereUpgradeObjects)
	if err != nil {
		return err
	}

	if len(vsphereUpgradeObjects) > 0 {
		s.D.Set("vsphere_upgrade_guide", "vsphereUpgradeGuide_place_holder")
	}

	log.Printf("[DEBUG] setting upgrade licenses")
	upgradeLicenses := []interface{}{}
	for _, item := range clusterResponse.UpgradeLicenses {
		upgradeLicenses = append(upgradeLicenses, VsphereLicenseToMap(item))
	}
	err = s.D.Set("upgrade_licenses", upgradeLicenses)
	if err != nil {
		return err
	}

	if clusterResponse.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", clusterResponse.CapacityReservationId)
	}

	if clusterResponse.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", clusterResponse.ComputeAvailabilityDomain)
	}

	datastores := []interface{}{}
	for _, item := range clusterResponse.Datastores {
		datastores = append(datastores, DatastoreDetailsToMap(item))
	}
	s.D.Set("datastores", datastores)

	if clusterResponse.InitialHostOcpuCount != nil {
		s.D.Set("initial_host_ocpu_count", clusterResponse.InitialHostOcpuCount)
	}

	if clusterResponse.InitialHostShapeName != nil {
		s.D.Set("initial_host_shape_name", clusterResponse.InitialHostShapeName)
	}

	s.D.Set("initial_sku", clusterResponse.InitialCommitment)

	if clusterResponse.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", clusterResponse.InstanceDisplayNamePrefix)
	}

	if clusterResponse.IsShieldedInstanceEnabled != nil {
		s.D.Set("is_shielded_instance_enabled", clusterResponse.IsShieldedInstanceEnabled)
	}

	if clusterResponse.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", clusterResponse.WorkloadNetworkCidr)
	}

	networkConfiguration := clusterResponse.NetworkConfiguration
	if networkConfiguration.NsxEdgeUplink1VlanId != nil {
		s.D.Set("nsx_edge_uplink1vlan_id", networkConfiguration.NsxEdgeUplink1VlanId)
	}
	if networkConfiguration.NsxEdgeUplink2VlanId != nil {
		s.D.Set("nsx_edge_uplink2vlan_id", networkConfiguration.NsxEdgeUplink2VlanId)
	}
	if networkConfiguration.NsxEdgeVTepVlanId != nil {
		s.D.Set("nsx_edge_vtep_vlan_id", networkConfiguration.NsxEdgeVTepVlanId)
	}
	if networkConfiguration.NsxVTepVlanId != nil {
		s.D.Set("nsx_vtep_vlan_id", networkConfiguration.NsxVTepVlanId)
	}
	if networkConfiguration.ProvisioningSubnetId != nil {
		s.D.Set("provisioning_subnet_id", networkConfiguration.ProvisioningSubnetId)
	}
	if networkConfiguration.ProvisioningVlanId != nil {
		s.D.Set("provisioning_vlan_id", networkConfiguration.ProvisioningVlanId)
	}
	if networkConfiguration.ReplicationVlanId != nil {
		s.D.Set("replication_vlan_id", networkConfiguration.ReplicationVlanId)
	}
	if networkConfiguration.VmotionVlanId != nil {
		s.D.Set("vmotion_vlan_id", networkConfiguration.VmotionVlanId)
	}
	if networkConfiguration.VsanVlanId != nil {
		s.D.Set("vsan_vlan_id", networkConfiguration.VsanVlanId)
	}
	if networkConfiguration.VsphereVlanId != nil {
		s.D.Set("vsphere_vlan_id", networkConfiguration.VsphereVlanId)
	}

	if networkConfiguration.HcxVlanId != nil {
		s.D.Set("hcx_vlan_id", networkConfiguration.HcxVlanId)
	}
	return nil
}
