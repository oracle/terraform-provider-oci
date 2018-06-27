// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/core"

	"github.com/oracle/terraform-provider-oci/crud"
)

func VolumeBackupPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readVolumeBackupPolicies,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"volume_backup_policies": {
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
						"schedules": {
							Type:     schema.TypeList,
							Computed: true,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"backup_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"offset_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"period": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"retention_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readVolumeBackupPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &VolumeBackupPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).blockstorageClient

	return crud.ReadResource(sync)
}

type VolumeBackupPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeBackupPoliciesResponse
}

func (s *VolumeBackupPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *VolumeBackupPoliciesDataSourceCrud) Get() error {
	request := oci_core.ListVolumeBackupPoliciesRequest{}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeBackupPolicies(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeBackupPolicies(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *VolumeBackupPoliciesDataSourceCrud) SetData() {
	if s.Res == nil {
		return
	}

	s.D.SetId(crud.GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeBackupPolicy := map[string]interface{}{}

		if r.DisplayName != nil {
			volumeBackupPolicy["display_name"] = *r.DisplayName
		}

		if r.Id != nil {
			volumeBackupPolicy["id"] = *r.Id
		}

		schedules := []interface{}{}
		for _, item := range r.Schedules {
			schedules = append(schedules, VolumeBackupScheduleToMap(item))
		}
		volumeBackupPolicy["schedules"] = schedules

		if r.TimeCreated != nil {
			volumeBackupPolicy["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, volumeBackupPolicy)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, VolumeBackupPoliciesDataSource().Schema["volume_backup_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_backup_policies", resources); err != nil {
		panic(err)
	}

	return
}

func VolumeBackupScheduleToMap(obj oci_core.VolumeBackupSchedule) map[string]interface{} {
	result := map[string]interface{}{}

	result["backup_type"] = string(obj.BackupType)

	if obj.OffsetSeconds != nil {
		result["offset_seconds"] = int(*obj.OffsetSeconds)
	}

	result["period"] = string(obj.Period)

	if obj.RetentionSeconds != nil {
		result["retention_seconds"] = int(*obj.RetentionSeconds)
	}

	return result
}
