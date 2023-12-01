// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package ocvp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_ocvp "github.com/oracle/oci-go-sdk/v65/ocvp"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OcvpClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OcvpClusterResource(), fieldMap, readSingularOcvpCluster)
}

func readSingularOcvpCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OcvpClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).ClusterClient()

	return tfresource.ReadResource(sync)
}

type OcvpClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_ocvp.ClusterClient
	Res    *oci_ocvp.GetClusterResponse
}

func (s *OcvpClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OcvpClusterDataSourceCrud) Get() error {
	request := oci_ocvp.GetClusterRequest{}

	if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
		tmp := clusterId.(string)
		request.ClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "ocvp")

	response, err := s.Client.GetCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OcvpClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CapacityReservationId != nil {
		s.D.Set("capacity_reservation_id", *s.Res.CapacityReservationId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ComputeAvailabilityDomain != nil {
		s.D.Set("compute_availability_domain", *s.Res.ComputeAvailabilityDomain)
	}

	datastores := []interface{}{}
	for _, item := range s.Res.Datastores {
		datastores = append(datastores, DatastoreDetailsToMap(item))
	}
	s.D.Set("datastores", datastores)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.EsxiHostsCount != nil {
		s.D.Set("esxi_hosts_count", *s.Res.EsxiHostsCount)
		s.D.Set("actual_esxi_hosts_count", *s.Res.EsxiHostsCount)
	}

	if s.Res.EsxiSoftwareVersion != nil {
		s.D.Set("esxi_software_version", *s.Res.EsxiSoftwareVersion)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)
	s.D.Set("freeform_tags", s.Res.FreeformTags)

	s.D.Set("initial_commitment", s.Res.InitialCommitment)

	if s.Res.InitialHostOcpuCount != nil {
		s.D.Set("initial_host_ocpu_count", *s.Res.InitialHostOcpuCount)
	}

	if s.Res.InitialHostShapeName != nil {
		s.D.Set("initial_host_shape_name", *s.Res.InitialHostShapeName)
	}

	if s.Res.InstanceDisplayNamePrefix != nil {
		s.D.Set("instance_display_name_prefix", *s.Res.InstanceDisplayNamePrefix)
	}

	if s.Res.IsShieldedInstanceEnabled != nil {
		s.D.Set("is_shielded_instance_enabled", *s.Res.IsShieldedInstanceEnabled)
	}

	if s.Res.NetworkConfiguration != nil {
		s.D.Set("network_configuration", []interface{}{NetworkConfigurationToMap(s.Res.NetworkConfiguration)})
	} else {
		s.D.Set("network_configuration", nil)
	}

	if s.Res.SddcId != nil {
		s.D.Set("sddc_id", *s.Res.SddcId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	upgradeLicenses := []interface{}{}
	for _, item := range s.Res.UpgradeLicenses {
		upgradeLicenses = append(upgradeLicenses, VsphereLicenseToMap(item))
	}
	s.D.Set("upgrade_licenses", upgradeLicenses)

	if s.Res.VmwareSoftwareVersion != nil {
		s.D.Set("vmware_software_version", *s.Res.VmwareSoftwareVersion)
	}

	s.D.Set("vsphere_type", s.Res.VsphereType)

	vsphereUpgradeObjects := []interface{}{}
	for _, item := range s.Res.VsphereUpgradeObjects {
		vsphereUpgradeObjects = append(vsphereUpgradeObjects, VsphereUpgradeObjectToMap(item))
	}
	s.D.Set("vsphere_upgrade_objects", vsphereUpgradeObjects)

	if s.Res.WorkloadNetworkCidr != nil {
		s.D.Set("workload_network_cidr", *s.Res.WorkloadNetworkCidr)
	}

	return nil
}
