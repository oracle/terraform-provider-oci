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

func DatabaseDbNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_server_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vm_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"db_nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_node_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"additional_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_vnic2id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"backup_vnic_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cpu_core_count": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"db_node_storage_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"db_server_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"db_system_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"fault_domain": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"host_ip_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"maintenance_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"memory_size_in_gbs": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"software_storage_size_in_gb": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_window_end": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_maintenance_window_start": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vnic2id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vnic_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDatabaseDbNodes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbNodesResponse
}

func (s *DatabaseDbNodesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodesDataSourceCrud) Get() error {
	request := oci_database.ListDbNodesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dbServerId, ok := s.D.GetOkExists("db_server_id"); ok {
		tmp := dbServerId.(string)
		request.DbServerId = &tmp
	}

	if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
		tmp := dbSystemId.(string)
		request.DbSystemId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.DbNodeSummaryLifecycleStateEnum(state.(string))
	}

	if vmClusterId, ok := s.D.GetOkExists("vm_cluster_id"); ok {
		tmp := vmClusterId.(string)
		request.VmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbNodes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response

	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbNodes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbNodesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbNodesDataSource-", DatabaseDbNodesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNode := map[string]interface{}{
			"db_system_id": *r.DbSystemId,
		}

		if r.AdditionalDetails != nil {
			dbNode["additional_details"] = *r.AdditionalDetails
		}

		if r.BackupIpId != nil {
			dbNode["backup_ip_id"] = *r.BackupIpId
		}

		if r.BackupVnic2Id != nil {
			dbNode["backup_vnic2id"] = *r.BackupVnic2Id
		}

		if r.BackupVnicId != nil {
			dbNode["backup_vnic_id"] = *r.BackupVnicId
		}

		if r.CpuCoreCount != nil {
			dbNode["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DbNodeStorageSizeInGBs != nil {
			dbNode["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DbServerId != nil {
			dbNode["db_server_id"] = *r.DbServerId
		}

		if r.DbSystemId != nil {
			dbNode["db_system_id"] = *r.DbSystemId
		}

		if r.FaultDomain != nil {
			dbNode["fault_domain"] = *r.FaultDomain
		}

		if r.HostIpId != nil {
			dbNode["host_ip_id"] = *r.HostIpId
		}

		if r.Hostname != nil {
			dbNode["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbNode["id"] = *r.Id
			dbNode["db_node_id"] = *r.Id // maintain legacy vanity id
		}

		dbNode["maintenance_type"] = r.MaintenanceType

		if r.MemorySizeInGBs != nil {
			dbNode["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.SoftwareStorageSizeInGB != nil {
			dbNode["software_storage_size_in_gb"] = *r.SoftwareStorageSizeInGB
		}

		dbNode["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbNode["time_created"] = r.TimeCreated.String()
		}

		if r.TimeMaintenanceWindowEnd != nil {
			dbNode["time_maintenance_window_end"] = r.TimeMaintenanceWindowEnd.String()
		}

		if r.TimeMaintenanceWindowStart != nil {
			dbNode["time_maintenance_window_start"] = r.TimeMaintenanceWindowStart.String()
		}

		if r.Vnic2Id != nil {
			dbNode["vnic2id"] = *r.Vnic2Id
		}

		if r.VnicId != nil {
			dbNode["vnic_id"] = *r.VnicId
		}

		resources = append(resources, dbNode)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbNodesDataSource().Schema["db_nodes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_nodes", resources); err != nil {
		return err
	}

	return nil
}
