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

func DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readDatabaseCloudAutonomousVmClusterAcdResourceUsages,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"cloud_autonomous_vm_cluster_id": {
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

func readDatabaseCloudAutonomousVmClusterAcdResourceUsages(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

type DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.ListCloudAutonomousVmClusterAcdResourceUsageResponse
}

func (s *DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSourceCrud) Get() error {
	request := oci_database.ListCloudAutonomousVmClusterAcdResourceUsageRequest{}

	if cloudAutonomousVmClusterId, ok := s.D.GetOkExists("cloud_autonomous_vm_cluster_id"); ok {
		tmp := cloudAutonomousVmClusterId.(string)
		request.CloudAutonomousVmClusterId = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "database")

	response, err := s.Client.ListCloudAutonomousVmClusterAcdResourceUsage(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListCloudAutonomousVmClusterAcdResourceUsage(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSource-", DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		cloudAutonomousVmClusterAcdResourceUsage := map[string]interface{}{}

		autonomousContainerDatabaseVmUsage := []interface{}{}
		for _, item := range r.AutonomousContainerDatabaseVmUsage {
			autonomousContainerDatabaseVmUsage = append(autonomousContainerDatabaseVmUsage, AcdAvmResourceStatsToMap(item))
		}
		cloudAutonomousVmClusterAcdResourceUsage["autonomous_container_database_vm_usage"] = autonomousContainerDatabaseVmUsage

		if r.AvailableCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["available_cpus"] = *r.AvailableCpus
		}

		if r.DefinedTags != nil {
			cloudAutonomousVmClusterAcdResourceUsage["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			cloudAutonomousVmClusterAcdResourceUsage["display_name"] = *r.DisplayName
		}

		cloudAutonomousVmClusterAcdResourceUsage["freeform_tags"] = r.FreeformTags
		cloudAutonomousVmClusterAcdResourceUsage["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			cloudAutonomousVmClusterAcdResourceUsage["id"] = *r.Id
		}

		if r.LargestProvisionableAutonomousDatabaseInCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["largest_provisionable_autonomous_database_in_cpus"] = *r.LargestProvisionableAutonomousDatabaseInCpus
		}

		cloudAutonomousVmClusterAcdResourceUsage["provisionable_cpus"] = r.ProvisionableCpus
		cloudAutonomousVmClusterAcdResourceUsage["provisionable_cpus"] = r.ProvisionableCpus

		if r.ProvisionedCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["provisioned_cpus"] = *r.ProvisionedCpus
		}

		if r.ReclaimableCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["reclaimable_cpus"] = *r.ReclaimableCpus
		}

		if r.ReservedCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["reserved_cpus"] = *r.ReservedCpus
		}

		if r.UsedCpus != nil {
			cloudAutonomousVmClusterAcdResourceUsage["used_cpus"] = *r.UsedCpus
		}

		resources = append(resources, cloudAutonomousVmClusterAcdResourceUsage)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, DatabaseCloudAutonomousVmClusterAcdResourceUsagesDataSource().Schema["autonomous_container_database_resource_usages"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("autonomous_container_database_resource_usages", resources); err != nil {
		return err
	}

	return nil
}

func AcdAvmResourceStatsMap(obj oci_database.AcdAvmResourceStats) map[string]interface{} {
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
