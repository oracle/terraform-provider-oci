// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v56/database"
)

func DatabaseVmClusterNetworksDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseVmClusterNetworks,
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
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_cluster_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseVmClusterNetworkResource()),
			},
		},
	}
}

func readDatabaseVmClusterNetworks(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseVmClusterNetworksDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseVmClusterNetworksDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListVmClusterNetworksResponse
}

func (s *DatabaseVmClusterNetworksDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseVmClusterNetworksDataSourceCrud) Get() error {
	request := oci_database.ListVmClusterNetworksRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.VmClusterNetworkSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListVmClusterNetworks(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVmClusterNetworks(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseVmClusterNetworksDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseVmClusterNetworksDataSource-", DatabaseVmClusterNetworksDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		vmClusterNetwork := map[string]interface{}{
			"compartment_id":            *r.CompartmentId,
			"exadata_infrastructure_id": *r.ExadataInfrastructureId,
		}

		if r.DefinedTags != nil {
			vmClusterNetwork["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			vmClusterNetwork["display_name"] = *r.DisplayName
		}

		vmClusterNetwork["dns"] = r.Dns

		vmClusterNetwork["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			vmClusterNetwork["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			vmClusterNetwork["lifecycle_details"] = *r.LifecycleDetails
		}

		vmClusterNetwork["ntp"] = r.Ntp

		scans := []interface{}{}
		for _, item := range r.Scans {
			scans = append(scans, ScanDetailsToMap(item))
		}
		vmClusterNetwork["scans"] = scans

		vmClusterNetwork["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			vmClusterNetwork["time_created"] = r.TimeCreated.String()
		}

		if r.VmClusterId != nil {
			vmClusterNetwork["vm_cluster_id"] = *r.VmClusterId
		}

		vmNetworks := []interface{}{}
		for _, item := range r.VmNetworks {
			vmNetworks = append(vmNetworks, VmNetworkDetailsToMap(item, true))
		}
		vmClusterNetwork["vm_networks"] = vmNetworks

		resources = append(resources, vmClusterNetwork)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseVmClusterNetworksDataSource().Schema["vm_cluster_networks"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("vm_cluster_networks", resources); err != nil {
		return err
	}

	return nil
}
