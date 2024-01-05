// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"
)

func DatabaseVmClusterNetworkDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["exadata_infrastructure_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["vm_cluster_network_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DatabaseVmClusterNetworkResource(), fieldMap, readSingularDatabaseVmClusterNetwork)
}

func readSingularDatabaseVmClusterNetwork(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworkDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterNetworkDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetVmClusterNetworkResponse
}

func (s *DatabaseVmClusterNetworkDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterNetworkDataSourceCrud) Get() error {
	request := oci_database.GetVmClusterNetworkRequest{}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if vmClusterNetworkId, ok := s.D.GetOkExists("vm_cluster_network_id"); ok {
		tmp := vmClusterNetworkId.(string)
		request.VmClusterNetworkId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetVmClusterNetwork(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseVmClusterNetworkDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("dns", s.Res.Dns)

	drScans := []interface{}{}
	for _, item := range s.Res.DrScans {
		drScans = append(drScans, DrScanDetailsToMap(item))
	}
	s.D.Set("dr_scans", drScans)

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("ntp", s.Res.Ntp)

	scans := []interface{}{}
	for _, item := range s.Res.Scans {
		scans = append(scans, ScanDetailsToMap(item))
	}
	s.D.Set("scans", scans)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VmClusterId != nil {
		s.D.Set("vm_cluster_id", *s.Res.VmClusterId)
	}

	vmNetworks := []interface{}{}
	for _, item := range s.Res.VmNetworks {
		vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item, true))
	}
	s.D.Set("vm_networks", vmNetworks)

	return nil
}
