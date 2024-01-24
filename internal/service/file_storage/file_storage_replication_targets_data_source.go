// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package file_storage

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_file_storage "github.com/oracle/oci-go-sdk/v65/filestorage"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func FileStorageReplicationTargetsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readFileStorageReplicationTargets,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"availability_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"replication_targets": {
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
						"compartment_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"defined_tags": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},
						"delta_progress": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"delta_status": {
							Type:     schema.TypeString,
							Computed: true,
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
						"last_snapshot_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"lifecycle_details": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"recovery_point_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replication_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"source_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"state": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"target_id": {
							Type:     schema.TypeString,
							Computed: true,
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

func readFileStorageReplicationTargets(d *schema.ResourceData, m interface{}) error {
	sync := &FileStorageReplicationTargetsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).FileStorageClient()

	return tfresource.ReadResource(sync)
}

type FileStorageReplicationTargetsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_file_storage.FileStorageClient
	Res    *oci_file_storage.ListReplicationTargetsResponse
}

func (s *FileStorageReplicationTargetsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *FileStorageReplicationTargetsDataSourceCrud) Get() error {
	request := oci_file_storage.ListReplicationTargetsRequest{}

	if availabilityDomain, ok := s.D.GetOkExists("availability_domain"); ok {
		tmp := availabilityDomain.(string)
		request.AvailabilityDomain = &tmp
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_file_storage.ListReplicationTargetsLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "file_storage")

	response, err := s.Client.ListReplicationTargets(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListReplicationTargets(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *FileStorageReplicationTargetsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("FileStorageReplicationTargetsDataSource-", FileStorageReplicationTargetsDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		replicationTarget := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			replicationTarget["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			replicationTarget["display_name"] = *r.DisplayName
		}

		replicationTarget["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			replicationTarget["id"] = *r.Id
		}

		if r.LifecycleDetails != nil {
			replicationTarget["lifecycle_details"] = *r.LifecycleDetails
		}

		if r.RecoveryPointTime != nil {
			replicationTarget["recovery_point_time"] = r.RecoveryPointTime.String()
		}

		replicationTarget["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			replicationTarget["time_created"] = r.TimeCreated.String()
		}

		resources = append(resources, replicationTarget)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, FileStorageReplicationTargetsDataSource().Schema["replication_targets"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("replication_targets", resources); err != nil {
		return err
	}

	return nil
}
