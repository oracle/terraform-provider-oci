// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package multicloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_multicloud "github.com/oracle/oci-go-sdk/v65/multicloud"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MulticloudResourceAnchorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMulticloudResourceAnchor,
		Schema: map[string]*schema.Schema{
			"resource_anchor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"should_fetch_compartment_name": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			// Computed
			"cloud_service_provider_metadata_item": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"account_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"csp_additional_properties": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem:     schema.TypeString,
						},

						"csp_resource_anchor_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"csp_resource_anchor_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"project_number": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_anchor_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_anchor_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_group": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_name": {
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
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"linked_compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"linked_compartment_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"setup_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_anchor_subscription_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_type": {
				Type:     schema.TypeString,
				Computed: true,
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
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularMulticloudResourceAnchor(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudResourceAnchorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OmhubResourceAnchorClient()

	return tfresource.ReadResource(sync)
}

type MulticloudResourceAnchorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.OmhubResourceAnchorClient
	Res    *oci_multicloud.GetResourceAnchorResponse
}

func (s *MulticloudResourceAnchorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudResourceAnchorDataSourceCrud) Get() error {
	request := oci_multicloud.GetResourceAnchorRequest{}

	if resourceAnchorId, ok := s.D.GetOkExists("resource_anchor_id"); ok {
		tmp := resourceAnchorId.(string)
		request.ResourceAnchorId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.GetResourceAnchorSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if shouldFetchCompartmentName, ok := s.D.GetOkExists("should_fetch_compartment_name"); ok {
		tmp := shouldFetchCompartmentName.(bool)
		request.ShouldFetchCompartmentName = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.GetResourceAnchor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MulticloudResourceAnchorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	if s.Res.Id != nil {
		s.D.SetId(*s.Res.Id)
	}

	if s.Res.CloudServiceProviderMetadataItem != nil {
		cloudServiceProviderMetadataItemArray := []interface{}{}
		if cloudServiceProviderMetadataItemMap := CloudServiceProviderMetadataItemToMap(&s.Res.CloudServiceProviderMetadataItem); cloudServiceProviderMetadataItemMap != nil {
			cloudServiceProviderMetadataItemArray = append(cloudServiceProviderMetadataItemArray, cloudServiceProviderMetadataItemMap)
		}
		s.D.Set("cloud_service_provider_metadata_item", cloudServiceProviderMetadataItemArray)
	} else {
		s.D.Set("cloud_service_provider_metadata_item", nil)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CompartmentName != nil {
		s.D.Set("compartment_name", *s.Res.CompartmentName)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LinkedCompartmentId != nil {
		s.D.Set("linked_compartment_id", *s.Res.LinkedCompartmentId)
	}

	if s.Res.LinkedCompartmentName != nil {
		s.D.Set("linked_compartment_name", *s.Res.LinkedCompartmentName)
	}

	if s.Res.Region != nil {
		s.D.Set("region", *s.Res.Region)
	}

	s.D.Set("setup_mode", s.Res.SetupMode)

	s.D.Set("lifecycle_state", s.Res.LifecycleState)

	if s.Res.SubscriptionId != nil {
		s.D.Set("resource_anchor_subscription_id", *s.Res.SubscriptionId)
	}

	s.D.Set("subscription_type", s.Res.SubscriptionType)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}

func CloudServiceProviderMetadataItemToMap(obj *oci_multicloud.CloudServiceProviderMetadataItem) map[string]interface{} {
	result := map[string]interface{}{}

	if (*obj).GetRegion() != nil {
		result["region"] = string(*(*obj).GetRegion())
	}

	if (*obj).GetResourceAnchorName() != nil {
		result["resource_anchor_name"] = (*obj).GetResourceAnchorName()
	}

	if (*obj).GetCspResourceAnchorId() != nil {
		result["csp_resource_anchor_id"] = string(*(*obj).GetCspResourceAnchorId())
	}

	if (*obj).GetCspResourceAnchorName() != nil {
		result["csp_resource_anchor_name"] = string(*(*obj).GetCspResourceAnchorName())
	}

	if (*obj).GetResourceAnchorUri() != nil {
		result["resource_anchor_uri"] = string(*(*obj).GetResourceAnchorUri())
	}

	if (*obj).GetCspAdditionalProperties() != nil {
		result["csp_additional_properties"] = (*obj).GetCspAdditionalProperties()
	}

	switch v := (*obj).(type) {
	case oci_multicloud.AwsCloudServiceProviderMetadataItem:
		result["subscription_type"] = "ORACLEDBATAWS"

		if v.AccountId != nil {
			result["account_id"] = string(*v.AccountId)
		}

	case oci_multicloud.AzureCloudServiceProviderMetadataItem:
		result["subscription_type"] = "ORACLEDBATAZURE"

		if v.ResourceGroup != nil {
			result["resource_group"] = string(*v.ResourceGroup)
		}

		if v.Subscription != nil {
			result["subscription"] = string(*v.Subscription)
		}
	case oci_multicloud.GcpCloudServiceProviderMetadataItem:
		result["subscription_type"] = "ORACLEDBATGOOGLE"

		if v.ProjectNumber != nil {
			result["project_number"] = string(*v.ProjectNumber)
		}
	default:
		log.Printf("[WARN] Received 'subscription_type' of unknown type %v", *obj)
		return nil
	}

	return result
}
