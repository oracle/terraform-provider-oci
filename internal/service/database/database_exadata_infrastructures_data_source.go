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

func DatabaseExadataInfrastructuresDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseExadataInfrastructures,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"exadata_infrastructures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(DatabaseExadataInfrastructureResource()),
			},
		},
	}
}

func readDatabaseExadataInfrastructures(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseExadataInfrastructuresDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseExadataInfrastructuresDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListExadataInfrastructuresResponse
}

func (s *DatabaseExadataInfrastructuresDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseExadataInfrastructuresDataSourceCrud) Get() error {
	request := oci_database.ListExadataInfrastructuresRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.ExadataInfrastructureSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListExadataInfrastructures(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListExadataInfrastructures(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseExadataInfrastructuresDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseExadataInfrastructuresDataSource-", DatabaseExadataInfrastructuresDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		exadataInfrastructure := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.ActivatedStorageCount != nil {
			exadataInfrastructure["activated_storage_count"] = *r.ActivatedStorageCount
		}

		if r.AdditionalStorageCount != nil {
			exadataInfrastructure["additional_storage_count"] = *r.AdditionalStorageCount
		}

		if r.AdminNetworkCIDR != nil {
			exadataInfrastructure["admin_network_cidr"] = *r.AdminNetworkCIDR
		}

		if r.CloudControlPlaneServer1 != nil {
			exadataInfrastructure["cloud_control_plane_server1"] = *r.CloudControlPlaneServer1
		}

		if r.CloudControlPlaneServer2 != nil {
			exadataInfrastructure["cloud_control_plane_server2"] = *r.CloudControlPlaneServer2
		}

		if r.ComputeCount != nil {
			exadataInfrastructure["compute_count"] = *r.ComputeCount
		}

		contacts := []interface{}{}
		for _, item := range r.Contacts {
			contacts = append(contacts, ExadataInfrastructureContactToMap(item))
		}
		exadataInfrastructure["contacts"] = contacts

		if r.CorporateProxy != nil {
			exadataInfrastructure["corporate_proxy"] = *r.CorporateProxy
		}

		if r.CpusEnabled != nil {
			exadataInfrastructure["cpus_enabled"] = *r.CpusEnabled
		}

		if r.CsiNumber != nil {
			exadataInfrastructure["csi_number"] = *r.CsiNumber
		}

		if r.DataStorageSizeInTBs != nil {
			exadataInfrastructure["data_storage_size_in_tbs"] = *r.DataStorageSizeInTBs
		}

		if r.DbNodeStorageSizeInGBs != nil {
			exadataInfrastructure["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DefinedTags != nil {
			exadataInfrastructure["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			exadataInfrastructure["display_name"] = *r.DisplayName
		}

		exadataInfrastructure["dns_server"] = r.DnsServer

		exadataInfrastructure["freeform_tags"] = r.FreeformTags

		if r.Gateway != nil {
			exadataInfrastructure["gateway"] = *r.Gateway
		}

		if r.Id != nil {
			exadataInfrastructure["id"] = *r.Id
		}

		if r.InfiniBandNetworkCIDR != nil {
			exadataInfrastructure["infini_band_network_cidr"] = *r.InfiniBandNetworkCIDR
		}

		if r.LifecycleDetails != nil {
			exadataInfrastructure["lifecycle_details"] = *r.LifecycleDetails
		}

		exadataInfrastructure["maintenance_slo_status"] = r.MaintenanceSLOStatus

		if r.MaintenanceWindow != nil {
			exadataInfrastructure["maintenance_window"] = []interface{}{ExadataInfrastructureMaintenanceWindowToMap(r.MaintenanceWindow)}
		} else {
			exadataInfrastructure["maintenance_window"] = nil
		}

		if r.MaxCpuCount != nil {
			exadataInfrastructure["max_cpu_count"] = *r.MaxCpuCount
		}

		if r.MaxDataStorageInTBs != nil {
			exadataInfrastructure["max_data_storage_in_tbs"] = *r.MaxDataStorageInTBs
		}

		if r.MaxDbNodeStorageInGBs != nil {
			exadataInfrastructure["max_db_node_storage_in_gbs"] = *r.MaxDbNodeStorageInGBs
		}

		if r.MaxMemoryInGBs != nil {
			exadataInfrastructure["max_memory_in_gbs"] = *r.MaxMemoryInGBs
		}

		if r.MemorySizeInGBs != nil {
			exadataInfrastructure["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.Netmask != nil {
			exadataInfrastructure["netmask"] = *r.Netmask
		}

		exadataInfrastructure["ntp_server"] = r.NtpServer

		if r.Shape != nil {
			exadataInfrastructure["shape"] = *r.Shape
		}

		exadataInfrastructure["state"] = r.LifecycleState

		if r.StorageCount != nil {
			exadataInfrastructure["storage_count"] = *r.StorageCount
		}

		if r.TimeCreated != nil {
			exadataInfrastructure["time_created"] = r.TimeCreated.String()
		}

		if r.TimeZone != nil {
			exadataInfrastructure["time_zone"] = *r.TimeZone
		}

		resources = append(resources, exadataInfrastructure)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseExadataInfrastructuresDataSource().Schema["exadata_infrastructures"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("exadata_infrastructures", resources); err != nil {
		return err
	}

	return nil
}
