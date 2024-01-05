// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package core

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_core "github.com/oracle/oci-go-sdk/v65/core"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func CoreVolumeGroupReplicaDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularCoreVolumeGroupReplica,
		Schema: map[string]*schema.Schema{
			"volume_group_replica_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
			"member_replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"membership_state": {
							Type:     schema.TypeString,
							Computed: true,
						},
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
	}
}

func readSingularCoreVolumeGroupReplica(d *schema.ResourceData, m interface{}) error {
	sync := &CoreVolumeGroupReplicaDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BlockstorageClient()

	return tfresource.ReadResource(sync)
}

type CoreVolumeGroupReplicaDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_core.BlockstorageClient
	Res    *oci_core.GetVolumeGroupReplicaResponse
}

func (s *CoreVolumeGroupReplicaDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *CoreVolumeGroupReplicaDataSourceCrud) Get() error {
	request := oci_core.GetVolumeGroupReplicaRequest{}

	if volumeGroupReplicaId, ok := s.D.GetOkExists("volume_group_replica_id"); ok {
		tmp := volumeGroupReplicaId.(string)
		request.VolumeGroupReplicaId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "core")

	response, err := s.Client.GetVolumeGroupReplica(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *CoreVolumeGroupReplicaDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.AvailabilityDomain != nil {
		s.D.Set("availability_domain", *s.Res.AvailabilityDomain)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	memberReplicas := []interface{}{}
	for _, item := range s.Res.MemberReplicas {
		memberReplicas = append(memberReplicas, MemberReplicaToMap(item))
	}
	s.D.Set("member_replicas", memberReplicas)

	if s.Res.SizeInGBs != nil {
		s.D.Set("size_in_gbs", strconv.FormatInt(*s.Res.SizeInGBs, 10))
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeLastSynced != nil {
		s.D.Set("time_last_synced", s.Res.TimeLastSynced.String())
	}

	if s.Res.VolumeGroupId != nil {
		s.D.Set("volume_group_id", *s.Res.VolumeGroupId)
	}

	return nil
}
