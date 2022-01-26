// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreVolumeBackupPoliciesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeBackupPolicies,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_backup_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(CoreVolumeBackupPolicyResource()),
			},
		},
	}
}

func readCoreVolumeBackupPolicies(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeBackupPoliciesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreVolumeBackupPoliciesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeBackupPoliciesResponse
}

func (s *CoreVolumeBackupPoliciesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeBackupPoliciesDataSourceCrud) Get() error {
	request := oci_core.ListVolumeBackupPoliciesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

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

func (s *CoreVolumeBackupPoliciesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVolumeBackupPoliciesDataSource-", CoreVolumeBackupPoliciesDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeBackupPolicy := map[string]interface{}{}

		if r.CompartmentId != nil {
			volumeBackupPolicy["compartment_id"] = *r.CompartmentId
		}

		if r.DefinedTags != nil {
			volumeBackupPolicy["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DestinationRegion != nil {
			volumeBackupPolicy["destination_region"] = *r.DestinationRegion
		}

		if r.DisplayName != nil {
			volumeBackupPolicy["display_name"] = *r.DisplayName
		}

		volumeBackupPolicy["freeform_tags"] = r.FreeformTags

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
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVolumeBackupPoliciesDataSource().Schema["volume_backup_policies"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_backup_policies", resources); err != nil {
		return err
	}

	return nil
}
