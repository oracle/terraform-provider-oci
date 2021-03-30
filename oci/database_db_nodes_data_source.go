// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v38/database"
)

func init() {
	RegisterDatasource("oci_database_db_nodes", DatabaseDbNodesDataSource())
}

func DatabaseDbNodesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbNodes,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
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
				Elem:     GetDataSourceItemSchema(DatabaseDbNodeDataSource()),
			},
		},
	}
}

func readDatabaseDbNodes(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient()

	return ReadResource(sync)
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

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

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

	s.D.SetId(GenerateDataSourceHashID("DatabaseDbNodesDataSource-", DatabaseDbNodesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbNode := map[string]interface{}{
			"db_system_id": *r.DbSystemId,
		}

		if r.AdditionalDetails != nil {
			dbNode["additional_details"] = *r.AdditionalDetails
		}

		if r.BackupVnicId != nil {
			dbNode["backup_vnic_id"] = *r.BackupVnicId
		}

		if r.FaultDomain != nil {
			dbNode["fault_domain"] = *r.FaultDomain
		}

		if r.Hostname != nil {
			dbNode["hostname"] = *r.Hostname
		}

		if r.Id != nil {
			dbNode["id"] = *r.Id
			dbNode["db_node_id"] = *r.Id // maintain legacy vanity id
		}

		dbNode["maintenance_type"] = r.MaintenanceType

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

		if r.VnicId != nil {
			dbNode["vnic_id"] = *r.VnicId
		}

		resources = append(resources, dbNode)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, DatabaseDbNodesDataSource().Schema["db_nodes"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_nodes", resources); err != nil {
		return err
	}

	return nil
}
