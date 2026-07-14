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

func DataccVmClusterNetworkDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["vm_cluster_network_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchemaWithContext(DataccVmClusterNetworkResource(), fieldMap, readSingularDataccVmClusterNetworkWithContext)
}

func readSingularDataccVmClusterNetworkWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &DataccVmClusterNetworkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BaseinfraClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

type DataccVmClusterNetworkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_datacc.BaseinfraClient
	Res    *oci_datacc.GetVmClusterNetworkResponse
}

func (s *DataccVmClusterNetworkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DataccVmClusterNetworkDataSourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_datacc.GetVmClusterNetworkRequest{}

	if vmClusterNetworkId, ok := s.D.GetOkExists("vm_cluster_network_id"); ok {
		tmp := vmClusterNetworkId.(string)
		request.VmClusterNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "datacc")

	response, err := s.Client.GetVmClusterNetwork(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DataccVmClusterNetworkDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AssociatedResourceId != nil {
		s.D.Set("associated_resource_id", *s.Res.AssociatedResourceId)
	}

	if s.Res.BaseVmClusterId != nil {
		s.D.Set("base_vm_cluster_id", *s.Res.BaseVmClusterId)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	s.D.Set("consumer_type", s.Res.ConsumerType)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns_servers", s.Res.DnsServers)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.InfrastructureId != nil {
		s.D.Set("infrastructure_id", *s.Res.InfrastructureId)
	}

	if s.Res.IsScanEnabled != nil {
		s.D.Set("is_scan_enabled", *s.Res.IsScanEnabled)
	}

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.ListenerPort != nil {
		s.D.Set("listener_port", *s.Res.ListenerPort)
	}

	if s.Res.ListenerPortSsl != nil {
		s.D.Set("listener_port_ssl", *s.Res.ListenerPortSsl)
	}

	if s.Res.NodeCount != nil {
		s.D.Set("node_count", *s.Res.NodeCount)
	}

	s.D.Set("ntp_servers", s.Res.NtpServers)

	scans := []interface{}{}
	for _, item := range s.Res.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	s.D.Set("scans", scans)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	vmNetworks := []interface{}{}
	for _, item := range s.Res.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item))
	}
	s.D.Set("vm_networks", vmNetworks)

	return nil
}
