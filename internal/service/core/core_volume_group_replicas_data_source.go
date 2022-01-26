// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v56/core"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func CoreVolumeGroupReplicasDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readCoreVolumeGroupReplicas,
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
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"volume_group_replicas": {
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
						"member_replicas": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"volume_replica_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"size_in_gbs": {
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
						"time_last_synced": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"volume_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readCoreVolumeGroupReplicas(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupReplicasDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreVolumeGroupReplicasDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.ListVolumeGroupReplicasResponse
}

func (s *CoreVolumeGroupReplicasDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeGroupReplicasDataSourceCrud) Get() error {
	request := oci_core.ListVolumeGroupReplicasRequest{}

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

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_core.VolumeGroupReplicaLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.ListVolumeGroupReplicas(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListVolumeGroupReplicas(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *CoreVolumeGroupReplicasDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("CoreVolumeGroupReplicasDataSource-", CoreVolumeGroupReplicasDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		volumeGroupReplica := map[string]interface{}{
			"availability_domain": *r.AvailabilityDomain,
			"compartment_id":      *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			volumeGroupReplica["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			volumeGroupReplica["display_name"] = *r.DisplayName
		}

		volumeGroupReplica["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			volumeGroupReplica["id"] = *r.Id
		}

		memberReplicas := []interface{}{}
		for _, item := range r.MemberReplicas {
			memberReplicas = append(memberReplicas, MemberReplicaToMap(item))
		}
		volumeGroupReplica["member_replicas"] = memberReplicas

		if r.SizeInGBs != nil {
			volumeGroupReplica["size_in_gbs"] = strconv.FormatInt(*r.SizeInGBs, 10)
		}

		volumeGroupReplica["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			volumeGroupReplica["time_created"] = r.TimeCreated.String()
		}

		if r.TimeLastSynced != nil {
			volumeGroupReplica["time_last_synced"] = r.TimeLastSynced.String()
		}

		if r.VolumeGroupId != nil {
			volumeGroupReplica["volume_group_id"] = *r.VolumeGroupId
		}

		resources = append(resources, volumeGroupReplica)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, CoreVolumeGroupReplicasDataSource().Schema["volume_group_replicas"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("volume_group_replicas", resources); err != nil {
		return err
	}

	return nil
}
