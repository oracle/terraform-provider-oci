// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package opensearch

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_opensearch "github.com/oracle/oci-go-sdk/v65/opensearch"

	"github.com/oracle/terraform-provider-oci/internal/client"

	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func OpensearchOpensearchClusterDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["opensearch_cluster_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(OpensearchOpensearchClusterResource(), fieldMap, readSingularOpensearchOpensearchCluster)
}

func readSingularOpensearchOpensearchCluster(d *schema.ResourceData, m interface{}) error {
	sync := &OpensearchOpensearchClusterDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OpensearchClusterClient()

	return tfresource.ReadResource(sync)
}

type OpensearchOpensearchClusterDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_opensearch.OpensearchClusterClient
	Res    *oci_opensearch.GetOpensearchClusterResponse
}

func (s *OpensearchOpensearchClusterDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *OpensearchOpensearchClusterDataSourceCrud) Get() error {
	request := oci_opensearch.GetOpensearchClusterRequest{}

	if opensearchClusterId, ok := s.D.GetOkExists("opensearch_cluster_id"); ok {
		tmp := opensearchClusterId.(string)
		request.OpensearchClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "opensearch")

	response, err := s.Client.GetOpensearchCluster(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *OpensearchOpensearchClusterDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("availability_domains", s.Res.AvailabilityDomains)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DataNodeCount != nil {
		s.D.Set("data_node_count", *s.Res.DataNodeCount)
	}

	if s.Res.DataNodeHostBareMetalShape != nil {
		s.D.Set("data_node_host_bare_metal_shape", *s.Res.DataNodeHostBareMetalShape)
	}

	if s.Res.DataNodeHostMemoryGB != nil {
		s.D.Set("data_node_host_memory_gb", *s.Res.DataNodeHostMemoryGB)
	}

	if s.Res.DataNodeHostOcpuCount != nil {
		s.D.Set("data_node_host_ocpu_count", *s.Res.DataNodeHostOcpuCount)
	}

	s.D.Set("data_node_host_type", s.Res.DataNodeHostType)

	if s.Res.DataNodeStorageGB != nil {
		s.D.Set("data_node_storage_gb", *s.Res.DataNodeStorageGB)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Fqdn != nil {
		s.D.Set("fqdn", *s.Res.Fqdn)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MasterNodeCount != nil {
		s.D.Set("master_node_count", *s.Res.MasterNodeCount)
	}

	if s.Res.MasterNodeHostBareMetalShape != nil {
		s.D.Set("master_node_host_bare_metal_shape", *s.Res.MasterNodeHostBareMetalShape)
	}

	if s.Res.MasterNodeHostMemoryGB != nil {
		s.D.Set("master_node_host_memory_gb", *s.Res.MasterNodeHostMemoryGB)
	}

	if s.Res.MasterNodeHostOcpuCount != nil {
		s.D.Set("master_node_host_ocpu_count", *s.Res.MasterNodeHostOcpuCount)
	}

	s.D.Set("master_node_host_type", s.Res.MasterNodeHostType)

	if s.Res.OpendashboardFqdn != nil {
		s.D.Set("opendashboard_fqdn", *s.Res.OpendashboardFqdn)
	}

	if s.Res.OpendashboardNodeCount != nil {
		s.D.Set("opendashboard_node_count", *s.Res.OpendashboardNodeCount)
	}

	if s.Res.OpendashboardNodeHostMemoryGB != nil {
		s.D.Set("opendashboard_node_host_memory_gb", *s.Res.OpendashboardNodeHostMemoryGB)
	}

	if s.Res.OpendashboardNodeHostOcpuCount != nil {
		s.D.Set("opendashboard_node_host_ocpu_count", *s.Res.OpendashboardNodeHostOcpuCount)
	}

	if s.Res.OpendashboardPrivateIp != nil {
		s.D.Set("opendashboard_private_ip", *s.Res.OpendashboardPrivateIp)
	}

	if s.Res.OpensearchFqdn != nil {
		s.D.Set("opensearch_fqdn", *s.Res.OpensearchFqdn)
	}

	if s.Res.OpensearchPrivateIp != nil {
		s.D.Set("opensearch_private_ip", *s.Res.OpensearchPrivateIp)
	}

	if s.Res.SecurityMasterUserName != nil {
		s.D.Set("security_master_user_name", *s.Res.SecurityMasterUserName)
	}

	if s.Res.SecurityMasterUserPasswordHash != nil {
		s.D.Set("security_master_user_password_hash", *s.Res.SecurityMasterUserPasswordHash)
	}

	s.D.Set("security_mode", s.Res.SecurityMode)

	if s.Res.SoftwareVersion != nil {
		s.D.Set("software_version", *s.Res.SoftwareVersion)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SubnetCompartmentId != nil {
		s.D.Set("subnet_compartment_id", *s.Res.SubnetCompartmentId)
	}

	if s.Res.SubnetId != nil {
		s.D.Set("subnet_id", *s.Res.SubnetId)
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.TotalStorageGB != nil {
		s.D.Set("total_storage_gb", *s.Res.TotalStorageGB)
	}

	if s.Res.VcnCompartmentId != nil {
		s.D.Set("vcn_compartment_id", *s.Res.VcnCompartmentId)
	}

	if s.Res.VcnId != nil {
		s.D.Set("vcn_id", *s.Res.VcnId)
	}

	return nil
}
