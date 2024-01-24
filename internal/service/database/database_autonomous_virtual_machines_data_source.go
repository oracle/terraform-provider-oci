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

func DatabaseAutonomousVirtualMachinesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousVirtualMachines,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_virtual_machines": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func readDatabaseAutonomousVirtualMachines(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVirtualMachinesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVirtualMachinesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousVirtualMachinesResponse
}

func (s *DatabaseAutonomousVirtualMachinesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVirtualMachinesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousVirtualMachinesRequest{}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_database.AutonomousVirtualMachineSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousVirtualMachines(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousVirtualMachines(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousVirtualMachinesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousVirtualMachinesDataSource-", DatabaseAutonomousVirtualMachinesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousVirtualMachine := map[string]interface{}{
			"autonomous_vm_cluster_id": *r.AutonomousVmClusterId,
			"compartment_id":           *r.CompartmentId,
		}

		if r.ClientIpAddress != nil {
			autonomousVirtualMachine["client_ip_address"] = *r.ClientIpAddress
		}

		if r.CpuCoreCount != nil {
			autonomousVirtualMachine["cpu_core_count"] = *r.CpuCoreCount
		}

		if r.DbNodeStorageSizeInGBs != nil {
			autonomousVirtualMachine["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DbServerDisplayName != nil {
			autonomousVirtualMachine["db_server_display_name"] = *r.DbServerDisplayName
		}

		if r.DbServerId != nil {
			autonomousVirtualMachine["db_server_id"] = *r.DbServerId
		}

		if r.DefinedTags != nil {
			autonomousVirtualMachine["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		autonomousVirtualMachine["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousVirtualMachine["id"] = *r.Id
		}

		if r.MemorySizeInGBs != nil {
			autonomousVirtualMachine["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		autonomousVirtualMachine["state"] = r.LifecycleState

		if r.VmName != nil {
			autonomousVirtualMachine["vm_name"] = *r.VmName
		}

		resources = append(resources, autonomousVirtualMachine)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousVirtualMachinesDataSource().Schema["autonomous_virtual_machines"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_virtual_machines", resources); err != nil {
		return err
	}

	return nil
}
