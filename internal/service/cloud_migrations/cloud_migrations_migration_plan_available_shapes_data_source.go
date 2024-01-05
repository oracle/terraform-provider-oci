// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package cloud_migrations

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CloudMigrationsMigrationPlanAvailableShapesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCloudMigrationsMigrationPlanAvailableShapes,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dvh_host_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"migration_plan_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"reserved_capacity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_shapes_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"availability_domain": {
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
									"gpu_description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"gpus": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"local_disk_description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"local_disks": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"local_disks_total_size_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"max_vnic_attachments": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"memory_in_gbs": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"min_total_baseline_ocpus_required": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"networking_bandwidth_in_gbps": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"ocpus": {
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"pagination_token": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"processor_description": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"shape": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"system_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func readCloudMigrationsMigrationPlanAvailableShapes(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanAvailableShapesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsMigrationPlanAvailableShapesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.ListAvailableShapesResponse
}

func (s *CloudMigrationsMigrationPlanAvailableShapesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsMigrationPlanAvailableShapesDataSourceCrud) Get() error {
	request := oci_cloud_migrations.ListAvailableShapesRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if dvhHostId, ok := s.D.GetOkExists("dvh_host_id"); ok {
		tmp := dvhHostId.(string)
		request.DvhHostId = &tmp
	}

	if migrationPlanId, ok := s.D.GetOkExists("migration_plan_id"); ok {
		tmp := migrationPlanId.(string)
		request.MigrationPlanId = &tmp
	}

	if reservedCapacityId, ok := s.D.GetOkExists("reserved_capacity_id"); ok {
		tmp := reservedCapacityId.(string)
		request.ReservedCapacityId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "cloud_migrations")

	response, err := s.Client.ListAvailableShapes(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListAvailableShapes(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CloudMigrationsMigrationPlanAvailableShapesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudMigrationsMigrationPlanAvailableShapesDataSource-", CloudMigrationsMigrationPlanAvailableShapesDataSource(), s.D))
	resources := []map[string]interface{}{}
	migrationPlanAvailableShape := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableShapeSummaryToMap(item))
	}
	migrationPlanAvailableShape["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, CloudMigrationsMigrationPlanAvailableShapesDataSource().Schema["available_shapes_collection"].Elem.(*schema.Resource).Schema)
		migrationPlanAvailableShape["items"] = items
	}

	resources = append(resources, migrationPlanAvailableShape)
	if err := s.D.Set("available_shapes_collection", resources); err != nil {
		return err
	}

	return nil
}

func AvailableShapeSummaryToMap(obj oci_cloud_migrations.AvailableShapeSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AvailabilityDomain != nil {
		result["availability_domain"] = string(*obj.AvailabilityDomain)
	}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	result["freeform_tags"] = obj.FreeformTags
	result["freeform_tags"] = obj.FreeformTags

	if obj.GpuDescription != nil {
		result["gpu_description"] = string(*obj.GpuDescription)
	}

	if obj.Gpus != nil {
		result["gpus"] = int(*obj.Gpus)
	}

	if obj.LocalDiskDescription != nil {
		result["local_disk_description"] = string(*obj.LocalDiskDescription)
	}

	if obj.LocalDisks != nil {
		result["local_disks"] = int(*obj.LocalDisks)
	}

	if obj.LocalDisksTotalSizeInGBs != nil {
		result["local_disks_total_size_in_gbs"] = float32(*obj.LocalDisksTotalSizeInGBs)
	}

	if obj.MaxVnicAttachments != nil {
		result["max_vnic_attachments"] = int(*obj.MaxVnicAttachments)
	}

	if obj.MemoryInGBs != nil {
		result["memory_in_gbs"] = float32(*obj.MemoryInGBs)
	}

	if obj.MinTotalBaselineOcpusRequired != nil {
		result["min_total_baseline_ocpus_required"] = float32(*obj.MinTotalBaselineOcpusRequired)
	}

	if obj.NetworkingBandwidthInGbps != nil {
		result["networking_bandwidth_in_gbps"] = float32(*obj.NetworkingBandwidthInGbps)
	}

	if obj.Ocpus != nil {
		result["ocpus"] = float32(*obj.Ocpus)
	}

	if obj.PaginationToken != nil {
		result["pagination_token"] = string(*obj.PaginationToken)
	}

	if obj.ProcessorDescription != nil {
		result["processor_description"] = string(*obj.ProcessorDescription)
	}

	if obj.Shape != nil {
		result["shape"] = string(*obj.Shape)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
