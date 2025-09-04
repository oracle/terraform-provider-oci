// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MulticloudOmHubMultiCloudMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMulticloudOmHubMultiCloudMetadata,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"base_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"base_subscription_id": {
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
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMulticloudOmHubMultiCloudMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudOmHubMultiCloudMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudsMetadataClient()

	return tfresource.ReadResource(sync)
}

type MulticloudOmHubMultiCloudMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MultiCloudsMetadataClient
	Res    *oci_multicloud.GetMultiCloudMetadataResponse
}

func (s *MulticloudOmHubMultiCloudMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudOmHubMultiCloudMetadataDataSourceCrud) Get() error {
	request := oci_multicloud.GetMultiCloudMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.GetMultiCloudMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MulticloudOmHubMultiCloudMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudOmHubMultiCloudMetadataDataSource-", MulticloudOmHubMultiCloudMetadataDataSource(), s.D))

	if s.Res.CompartmentId != nil {
		s.D.Set("base_compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.SubscriptionId != nil {
		s.D.Set("base_subscription_id", *s.Res.SubscriptionId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
