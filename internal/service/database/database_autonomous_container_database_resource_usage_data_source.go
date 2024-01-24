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

func DatabaseAutonomousContainerDatabaseResourceUsageDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularDatabaseAutonomousContainerDatabaseResourceUsage,
		Schema: map[string]*schema.Schema{
			"autonomous_container_database_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"autonomous_container_database_vm_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
			"largest_provisionable_autonomous_database_in_cpus": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"provisionable_cpus": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
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
	}
}

func readSingularDatabaseAutonomousContainerDatabaseResourceUsage(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousContainerDatabaseResourceUsageDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousContainerDatabaseResourceUsageDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetAutonomousContainerDatabaseResourceUsageResponse
}

func (s *DatabaseAutonomousContainerDatabaseResourceUsageDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousContainerDatabaseResourceUsageDataSourceCrud) Get() error {
	request := oci_database.GetAutonomousContainerDatabaseResourceUsageRequest{}

	if autonomousContainerDatabaseId, ok := s.D.GetOkExists("autonomous_container_database_id"); ok {
		tmp := autonomousContainerDatabaseId.(string)
		request.AutonomousContainerDatabaseId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.GetAutonomousContainerDatabaseResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseAutonomousContainerDatabaseResourceUsageDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	autonomousContainerDatabaseVmUsage := []interface{}{}
	for _, item := range s.Res.AutonomousContainerDatabaseVmUsage {
		autonomousContainerDatabaseVmUsage = append(autonomousContainerDatabaseVmUsage, AcdAvmResourceStatsToMap(item))
	}
	s.D.Set("autonomous_container_database_vm_usage", autonomousContainerDatabaseVmUsage)

	if s.Res.AvailableCpus != nil {
		s.D.Set("available_cpus", *s.Res.AvailableCpus)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LargestProvisionableAutonomousDatabaseInCpus != nil {
		s.D.Set("largest_provisionable_autonomous_database_in_cpus", *s.Res.LargestProvisionableAutonomousDatabaseInCpus)
	}

	s.D.Set("provisionable_cpus", s.Res.ProvisionableCpus)

	if s.Res.ProvisionedCpus != nil {
		s.D.Set("provisioned_cpus", *s.Res.ProvisionedCpus)
	}

	if s.Res.ReclaimableCpus != nil {
		s.D.Set("reclaimable_cpus", *s.Res.ReclaimableCpus)
	}

	if s.Res.ReservedCpus != nil {
		s.D.Set("reserved_cpus", *s.Res.ReservedCpus)
	}

	if s.Res.UsedCpus != nil {
		s.D.Set("used_cpus", *s.Res.UsedCpus)
	}

	return nil
}

func AcdAvmResourceStatsToMap(obj oci_database.AcdAvmResourceStats) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.ProvisionedCpus != nil {
		result["provisioned_cpus"] = float32(*obj.ProvisionedCpus)
	}

	if obj.ReclaimableCpus != nil {
		result["reclaimable_cpus"] = float32(*obj.ReclaimableCpus)
	}

	if obj.ReservedCpus != nil {
		result["reserved_cpus"] = float32(*obj.ReservedCpus)
	}

	if obj.UsedCpus != nil {
		result["used_cpus"] = float32(*obj.UsedCpus)
	}

	return result
}
