// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousVirtualMachineDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousVirtualMachine,
		Schema: map[string]*schema.Schema{
			"autonomous_virtual_machine_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
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
			"db_server_display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"db_server_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseAutonomousVirtualMachine(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVirtualMachineDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVirtualMachineDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousVirtualMachineResponse
}

func (s *DatabaseAutonomousVirtualMachineDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVirtualMachineDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousVirtualMachineRequest{}

	if autonomousVirtualMachineId, ok := s.D.GetOkExists("autonomous_virtual_machine_id"); ok {
		tmp := autonomousVirtualMachineId.(string)
		request.AutonomousVirtualMachineId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousVirtualMachine(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousVirtualMachineDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousVmClusterId != nil {
		s.D.Set("autonomous_vm_cluster_id", *s.Res.AutonomousVmClusterId)
	}

	if s.Res.ClientIpAddress != nil {
		s.D.Set("client_ip_address", *s.Res.ClientIpAddress)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerDisplayName != nil {
		s.D.Set("db_server_display_name", *s.Res.DbServerDisplayName)
	}

	if s.Res.DbServerId != nil {
		s.D.Set("db_server_id", *s.Res.DbServerId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.VmName != nil {
		s.D.Set("vm_name", *s.Res.VmName)
	}

	return nil
}
