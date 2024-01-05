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

func DatabaseAutonomousVmClusterResourceUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousVmClusterResourceUsage,
		Schema: map[string]*schema.Schema{
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"autonomous_vm_resource_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"autonomous_container_database_usage": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"available_cpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"provisioned_cpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"reclaimable_cpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"reserved_cpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"used_cpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
						"available_cpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"provisioned_cpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"reclaimable_cpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"reserved_cpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"used_cpus": {
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
			"available_autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"available_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"db_node_storage_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exadata_storage_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"is_local_backup_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory_per_oracle_compute_unit_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_size_in_gbs": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"non_provisionable_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisionable_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisioned_autonomous_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"provisioned_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reclaimable_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"reserved_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"total_container_databases": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"total_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"used_autonomous_data_storage_size_in_tbs": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"used_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func readSingularDatabaseAutonomousVmClusterResourceUsage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterResourceUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVmClusterResourceUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousVmClusterResourceUsageResponse
}

func (s *DatabaseAutonomousVmClusterResourceUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVmClusterResourceUsageDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousVmClusterResourceUsageRequest{}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousVmClusterResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousVmClusterResourceUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AutonomousDataStorageSizeInTBs != nil {
		s.D.Set("autonomous_data_storage_size_in_tbs", *s.Res.AutonomousDataStorageSizeInTBs)
	}

	autonomousVmResourceUsage := []interface{}{}
	for _, item := range s.Res.AutonomousVmResourceUsage {
		autonomousVmResourceUsage = append(autonomousVmResourceUsage, AutonomousVmResourceUsageToMap(item))
	}
	s.D.Set("autonomous_vm_resource_usage", autonomousVmResourceUsage)

	if s.Res.AvailableAutonomousDataStorageSizeInTBs != nil {
		s.D.Set("available_autonomous_data_storage_size_in_tbs", *s.Res.AvailableAutonomousDataStorageSizeInTBs)
	}

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.DbNodeStorageSizeInGBs != nil {
		s.D.Set("db_node_storage_size_in_gbs", *s.Res.DbNodeStorageSizeInGBs)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.ExadataStorageInTBs != nil {
		s.D.Set("exadata_storage_in_tbs", *s.Res.ExadataStorageInTBs)
	}

	if s.Res.IsLocalBackupEnabled != nil {
		s.D.Set("is_local_backup_enabled", *s.Res.IsLocalBackupEnabled)
	}

	if s.Res.MemoryPerOracleComputeUnitInGBs != nil {
		s.D.Set("memory_per_oracle_compute_unit_in_gbs", *s.Res.MemoryPerOracleComputeUnitInGBs)
	}

	if s.Res.MemorySizeInGBs != nil {
		s.D.Set("memory_size_in_gbs", *s.Res.MemorySizeInGBs)
	}

	if s.Res.NonProvisionableAutonomousContainerDatabases != nil {
		s.D.Set("non_provisionable_autonomous_container_databases", *s.Res.NonProvisionableAutonomousContainerDatabases)
	}

	if s.Res.ProvisionableAutonomousContainerDatabases != nil {
		s.D.Set("provisionable_autonomous_container_databases", *s.Res.ProvisionableAutonomousContainerDatabases)
	}

	if s.Res.ProvisionedAutonomousContainerDatabases != nil {
		s.D.Set("provisioned_autonomous_container_databases", *s.Res.ProvisionedAutonomousContainerDatabases)
	}

	if s.Res.ProvisionedCpus != nil {
		s.D.Set("provisioned_cpus", *s.Res.ProvisionedCpus)
	}

	if s.Res.ReclaimableCpus != nil {
		s.D.Set("reclaimable_cpus", *s.Res.ReclaimableCpus)
	}

	if s.Res.ReservedCpus != nil {
		s.D.Set("reserved_cpus", *s.Res.ReservedCpus)
	}

	if s.Res.TotalContainerDatabases != nil {
		s.D.Set("total_container_databases", *s.Res.TotalContainerDatabases)
	}

	if s.Res.TotalCpus != nil {
		s.D.Set("total_cpus", *s.Res.TotalCpus)
	}

	if s.Res.UsedAutonomousDataStorageSizeInTBs != nil {
		s.D.Set("used_autonomous_data_storage_size_in_tbs", *s.Res.UsedAutonomousDataStorageSizeInTBs)
	}

	if s.Res.UsedCpus != nil {
		s.D.Set("used_cpus", *s.Res.UsedCpus)
	}

	return nil
}
