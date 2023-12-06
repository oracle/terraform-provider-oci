// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/v65/database"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DatabaseAutonomousVmClusterAcdResourceUsagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseAutonomousVmClusterAcdResourceUsages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"autonomous_vm_cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"autonomous_container_database_resource_usages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

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
						"id": {
							Type:     schema.TypeString,
							Computed: true,
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
				},
			},
		},
	}
}

func readDatabaseAutonomousVmClusterAcdResourceUsages(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseAutonomousVmClusterAcdResourceUsagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseAutonomousVmClusterAcdResourceUsagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListAutonomousVmClusterAcdResourceUsageResponse
}

func (s *DatabaseAutonomousVmClusterAcdResourceUsagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseAutonomousVmClusterAcdResourceUsagesDataSourceCrud) Get() error {
	request := oci_database.ListAutonomousVmClusterAcdResourceUsageRequest{}

	if autonomousVmClusterId, ok := s.D.GetOkExists("autonomous_vm_cluster_id"); ok {
		tmp := autonomousVmClusterId.(string)
		request.AutonomousVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListAutonomousVmClusterAcdResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAutonomousVmClusterAcdResourceUsage(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseAutonomousVmClusterAcdResourceUsagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseAutonomousVmClusterAcdResourceUsagesDataSource-", DatabaseAutonomousVmClusterAcdResourceUsagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		autonomousVmClusterAcdResourceUsage := map[string]interface{}{}

		autonomousContainerDatabaseVmUsage := []interface{}{}
		for _, item := range r.AutonomousContainerDatabaseVmUsage {
			autonomousContainerDatabaseVmUsage = append(autonomousContainerDatabaseVmUsage, AcdAvmResourceStatsToMap(item))
		}
		autonomousVmClusterAcdResourceUsage["autonomous_container_database_vm_usage"] = autonomousContainerDatabaseVmUsage

		if r.AvailableCpus != nil {
			autonomousVmClusterAcdResourceUsage["available_cpus"] = *r.AvailableCpus
		}

		if r.DefinedTags != nil {
			autonomousVmClusterAcdResourceUsage["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			autonomousVmClusterAcdResourceUsage["display_name"] = *r.DisplayName
		}

		autonomousVmClusterAcdResourceUsage["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			autonomousVmClusterAcdResourceUsage["id"] = *r.Id
		}

		if r.LargestProvisionableAutonomousDatabaseInCpus != nil {
			autonomousVmClusterAcdResourceUsage["largest_provisionable_autonomous_database_in_cpus"] = *r.LargestProvisionableAutonomousDatabaseInCpus
		}

		autonomousVmClusterAcdResourceUsage["provisionable_cpus"] = r.ProvisionableCpus

		if r.ProvisionedCpus != nil {
			autonomousVmClusterAcdResourceUsage["provisioned_cpus"] = *r.ProvisionedCpus
		}

		if r.ReclaimableCpus != nil {
			autonomousVmClusterAcdResourceUsage["reclaimable_cpus"] = *r.ReclaimableCpus
		}

		if r.ReservedCpus != nil {
			autonomousVmClusterAcdResourceUsage["reserved_cpus"] = *r.ReservedCpus
		}

		if r.UsedCpus != nil {
			autonomousVmClusterAcdResourceUsage["used_cpus"] = *r.UsedCpus
		}

		resources = append(resources, autonomousVmClusterAcdResourceUsage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseAutonomousVmClusterAcdResourceUsagesDataSource().Schema["autonomous_container_database_resource_usages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_database_resource_usages", resources); err != nil {
		return err
	}

	return nil
}
