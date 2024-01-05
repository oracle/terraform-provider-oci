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

func CloudMigrationsMigrationPlanAvailableShapeDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCloudMigrationsMigrationPlanAvailableShape,
		Schema: map[string]*schema.Schema{
			"availability_domain": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
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
			// Computed
			"items": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
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
		DeprecationMessage: tfresource.DatasourceDeprecatedForAnother("oci_cloud_migrations_migration_plan_available_shape", "oci_cloud_migrations_migration_plan_available_shapes"),
	}
}

func readSingularCloudMigrationsMigrationPlanAvailableShape(d *schema.ResourceData, m interface{}) error {
	sync := &CloudMigrationsMigrationPlanAvailableShapeDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MigrationClient()

	return tfresource.ReadResource(sync)
}

type CloudMigrationsMigrationPlanAvailableShapeDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_cloud_migrations.MigrationClient
	Res    *oci_cloud_migrations.ListAvailableShapesResponse
}

func (s *CloudMigrationsMigrationPlanAvailableShapeDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CloudMigrationsMigrationPlanAvailableShapeDataSourceCrud) Get() error {
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
	return nil
}

func (s *CloudMigrationsMigrationPlanAvailableShapeDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CloudMigrationsMigrationPlanAvailableShapeDataSource-", CloudMigrationsMigrationPlanAvailableShapeDataSource(), s.D))

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, AvailableShapeSummaryToMap(item))
	}
	s.D.Set("items", items)

	return nil
}
