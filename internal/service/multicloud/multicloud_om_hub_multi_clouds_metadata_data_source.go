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

func MulticloudOmHubMultiCloudsMetadataDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMulticloudOmHubMultiCloudsMetadata,
		Schema: map[string]*schema.Schema{
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// Computed
			"multi_cloud_metadata_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"compartment_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"time_created": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subscription_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
									},
									"defined_tags": {
										Type:     schema.TypeMap,
										Computed: true,
										Elem:     schema.TypeString,
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

func readSingularMulticloudOmHubMultiCloudsMetadata(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudOmHubMultiCloudsMetadataDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).MultiCloudsMetadataClient()

	return tfresource.ReadResource(sync)
}

type MulticloudOmHubMultiCloudsMetadataDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.MultiCloudsMetadataClient
	Res    *oci_multicloud.ListMultiCloudMetadataResponse
}

func (s *MulticloudOmHubMultiCloudsMetadataDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudOmHubMultiCloudsMetadataDataSourceCrud) Get() error {
	request := oci_multicloud.ListMultiCloudMetadataRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if limit, ok := s.D.GetOkExists("limit"); ok {
		tmp := limit.(int)
		request.Limit = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.ListMultiCloudMetadata(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MulticloudOmHubMultiCloudsMetadataDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MulticloudOmHubMultiCloudsMetadataDataSource-", MulticloudOmHubMultiCloudsMetadataDataSource(), s.D))
	collection := []map[string]interface{}{}
	MetadataCollection := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, MultiCloudMetadataSummaryToMap(item))
	}
	MetadataCollection["items"] = items

	collection = append(collection, MetadataCollection)
	if err := s.D.Set("multi_cloud_metadata_collection", collection); err != nil {
		return err
	}

	return nil
}

func MultiCloudMetadataSummaryToMap(obj oci_multicloud.MultiCloudMetadataSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CompartmentId != nil {
		result["compartment_id"] = string(*obj.CompartmentId)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.SubscriptionId != nil {
		result["subscription_id"] = string(*obj.SubscriptionId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	return result
}
