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

func DatabaseDbServersDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseDbServers,
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
			"db_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"exadata_infrastructure_id": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func readDatabaseDbServers(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseDbServersDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseDbServersDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListDbServersResponse
}

func (s *DatabaseDbServersDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseDbServersDataSourceCrud) Get() error {
	request := oci_database.ListDbServersRequest{}

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
		request.LifecycleState = oci_database.DbServerSummaryLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListDbServers(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDbServers(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseDbServersDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseDbServersDataSource-", DatabaseDbServersDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		dbServer := map[string]interface{}{
			"compartment_id":            *r.CompartmentId,
			"exadata_infrastructure_id": *r.ExadataInfrastructureId,
		}

		dbServer["autonomous_virtual_machine_ids"] = r.AutonomousVirtualMachineIds

		dbServer["autonomous_vm_cluster_ids"] = r.AutonomousVmClusterIds

		if r.CpuCoreCount != nil {
			dbServer["cpu_core_count"] = *r.CpuCoreCount
		}

		dbServer["db_node_ids"] = r.DbNodeIds

		if r.DbNodeStorageSizeInGBs != nil {
			dbServer["db_node_storage_size_in_gbs"] = *r.DbNodeStorageSizeInGBs
		}

		if r.DbServerPatchingDetails != nil {
			dbServer["db_server_patching_details"] = []interface{}{DbServersPatchingDetailsToMap(r.DbServerPatchingDetails)}
		} else {
			dbServer["db_server_patching_details"] = nil
		}

		if r.DefinedTags != nil {
			dbServer["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			dbServer["display_name"] = *r.DisplayName
		}

		dbServer["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			dbServer["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			dbServer["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.MaxCpuCount != nil {
			dbServer["max_cpu_count"] = *r.MaxCpuCount
		}

		if r.MaxDbNodeStorageInGBs != nil {
			dbServer["max_db_node_storage_in_gbs"] = *r.MaxDbNodeStorageInGBs
		}

		if r.MaxMemoryInGBs != nil {
			dbServer["max_memory_in_gbs"] = *r.MaxMemoryInGBs
		}

		if r.MemorySizeInGBs != nil {
			dbServer["memory_size_in_gbs"] = *r.MemorySizeInGBs
		}

		if r.Shape != nil {
			dbServer["shape"] = *r.Shape
		}

		dbServer["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			dbServer["time_created"] = r.TimeCreated.String()
		}

		dbServer["vm_cluster_ids"] = r.VmClusterIds

		resources = append(resources, dbServer)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseDbServersDataSource().Schema["db_servers"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("db_servers", resources); err != nil {
		return err
	}

	return nil
}

func DbServersPatchingDetailsToMap(obj *oci_database.DbServerPatchingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EstimatedPatchDuration != nil {
		result["estimated_patch_duration"] = int(*obj.EstimatedPatchDuration)
	}

	result["patching_status"] = string(obj.PatchingStatus)

	if obj.TimePatchingEnded != nil {
		result["time_patching_ended"] = obj.TimePatchingEnded.String()
	}

	if obj.TimePatchingStarted != nil {
		result["time_patching_started"] = obj.TimePatchingStarted.String()
	}

	return result
}

func DbServerPatchingDetailsToMap(obj *oci_database.DbServerPatchingDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.EstimatedPatchDuration != nil {
		result["estimated_patch_duration"] = int(*obj.EstimatedPatchDuration)
	}

	result["patching_status"] = string(obj.PatchingStatus)

	if obj.TimePatchingEnded != nil {
		result["time_patching_ended"] = obj.TimePatchingEnded.String()
	}

	if obj.TimePatchingStarted != nil {
		result["time_patching_started"] = obj.TimePatchingStarted.String()
	}

	return result
}
