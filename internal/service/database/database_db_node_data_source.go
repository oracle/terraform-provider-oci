// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v58/database"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func DatabaseDbNodeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDbNode,
		Schema: map[string]*schema.Schema{
			"db_node_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
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
	}
}

func readSingularDatabaseDbNode(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbNodeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbNodeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDbNodeResponse
}

func (s *DatabaseDbNodeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbNodeDataSourceCrud) Get() error {
	request := oci_database.GetDbNodeRequest{}

	if dbNodeId, ok := s.D.GetOkExists("db_node_id"); ok {
		tmp := dbNodeId.(string)
		request.DbNodeId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDbNode(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbNodeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AdditionalDetails != nil {
		s.D.Set("additional_details", *s.Res.AdditionalDetails)
	}

	if s.Res.BackupIpId != nil {
		s.D.Set("backup_ip_id", *s.Res.BackupIpId)
	}

	if s.Res.BackupVnic2Id != nil {
		s.D.Set("backup_vnic2id", *s.Res.BackupVnic2Id)
	}

	if s.Res.BackupVnicId != nil {
		s.D.Set("backup_vnic_id", *s.Res.BackupVnicId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerId != nil {
		s.D.Set("db_server_id", *s.Res.DbServerId)
	}

	if s.Res.DbSystemId != nil {
		s.D.Set("db_system_id", *s.Res.DbSystemId)
	}

	if s.Res.FaultDomain != nil {
		s.D.Set("fault_domain", *s.Res.FaultDomain)
	}

	if s.Res.HostIpId != nil {
		s.D.Set("host_ip_id", *s.Res.HostIpId)
	}

	if s.Res.Hostname != nil {
		s.D.Set("hostname", *s.Res.Hostname)
	}

	s.D.Set("maintenance_type", s.Res.MaintenanceType)

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.SoftwareStorageSizeInGB != nil {
		s.D.Set("software_storage_size_in_gb", *s.Res.SoftwareStorageSizeInGB)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeMaintenanceWindowEnd != nil {
		s.D.Set("time_maintenance_window_end", s.Res.TimeMaintenanceWindowEnd.String())
	}

	if s.Res.TimeMaintenanceWindowStart != nil {
		s.D.Set("time_maintenance_window_start", s.Res.TimeMaintenanceWindowStart.String())
	}

	if s.Res.Vnic2Id != nil {
		s.D.Set("vnic2id", *s.Res.Vnic2Id)
	}

	if s.Res.VnicId != nil {
		s.D.Set("vnic_id", *s.Res.VnicId)
	}

	return nil
}
