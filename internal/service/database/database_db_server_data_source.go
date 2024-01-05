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

func DatabaseDbServerDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseDbServer,
		Schema: map[string]*schema.Schema{
			"db_server_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"exadata_infrastructure_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"autonomous_virtual_machine_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"autonomous_vm_cluster_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_node_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"db_server_patching_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"estimated_patch_duration": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"patching_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_patching_ended": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_patching_started": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"defined_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_db_node_storage_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_memory_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"shape": {
				Type:     schema.TypeString,
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
			"vm_cluster_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readSingularDatabaseDbServer(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbServerDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbServerDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetDbServerResponse
}

func (s *DatabaseDbServerDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbServerDataSourceCrud) Get() error {
	request := oci_database.GetDbServerRequest{}

	if dbServerId, ok := s.D.GetOkExists("db_server_id"); ok {
		tmp := dbServerId.(string)
		request.DbServerId = &tmp
	}

	if exadataInfrastructureId, ok := s.D.GetOkExists("exadata_infrastructure_id"); ok {
		tmp := exadataInfrastructureId.(string)
		request.ExadataInfrastructureId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetDbServer(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseDbServerDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	s.D.Set("autonomous_virtual_machine_ids", s.Res.AutonomousVirtualMachineIds)

	s.D.Set("autonomous_vm_cluster_ids", s.Res.AutonomousVmClusterIds)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CpuCoreCount != nil {
		s.D.Set("cpu_core_count", *s.Res.CpuCoreCount)
	}

	s.D.Set("db_node_ids", s.Res.DbNodeIds)

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DbServerPatchingDetails != nil {
		s.D.Set("db_server_patching_details", []interface{}{DbServerPatchingDetailsToMap(s.Res.DbServerPatchingDetails)})
	} else {
		s.D.Set("db_server_patching_details", nil)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.MaxCpuCount != nil {
		s.D.Set("max_cpu_count", *s.Res.MaxCpuCount)
	}

	if s.Res.MaxDbNodeStorageInGBs != nil {
		s.D.Set("max_db_node_storage_in_gbs", *s.Res.MaxDbNodeStorageInGBs)
	}

	if s.Res.MaxMemoryInGBs != nil {
		s.D.Set("max_memory_in_gbs", *s.Res.MaxMemoryInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.Shape != nil {
		s.D.Set("shape", *s.Res.Shape)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("vm_cluster_ids", s.Res.VmClusterIds)

	return nil
}
