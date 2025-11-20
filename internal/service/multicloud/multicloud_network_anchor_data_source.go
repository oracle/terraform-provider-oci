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

func MulticloudNetworkAnchorDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularMulticloudNetworkAnchor,
		Schema: map[string]*schema.Schema{
			// Required
			"network_anchor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_service_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Optional
			"external_location": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// Computed
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_anchor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_anchor_lifecycle_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lifecycle_details": {
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
			"setup_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_placement_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_metadata_item": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"network_anchor_connection_status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"vcn": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"vcn_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cidr_blocks": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"backup_cidr_blocks": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"dns_label": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dns": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"custom_domain_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"subnets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional

									// Computed
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"subnet_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"label": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"dns_listening_endpoint_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_forwarding_endpoint_ip_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_forwarding_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
			"cloud_service_provider_metadata_item": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"odb_network_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"cidr_blocks": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_anchor_uri": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"dns_forwarding_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func readSingularMulticloudNetworkAnchor(d *schema.ResourceData, m interface{}) error {
	sync := &MulticloudNetworkAnchorDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).OmhubNetworkAnchorClient()

	return tfresource.ReadResource(sync)
}

type MulticloudNetworkAnchorDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_multicloud.OmhubNetworkAnchorClient
	Res    *oci_multicloud.GetNetworkAnchorResponse
}

func (s *MulticloudNetworkAnchorDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MulticloudNetworkAnchorDataSourceCrud) Get() error {
	request := oci_multicloud.GetNetworkAnchorRequest{}

	// Required

	if networkAnchorId, ok := s.D.GetOkExists("network_anchor_id"); ok {
		tmp := networkAnchorId.(string)
		request.NetworkAnchorId = &tmp
	}

	if subscriptionServiceName, ok := s.D.GetOkExists("subscription_service_name"); ok {
		request.SubscriptionServiceName = oci_multicloud.GetNetworkAnchorSubscriptionServiceNameEnum(subscriptionServiceName.(string))
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	// Optional

	if externalLocation, ok := s.D.GetOkExists("external_location"); ok {
		tmp := externalLocation.(string)
		request.ExternalLocation = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "multicloud")

	response, err := s.Client.GetNetworkAnchor(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MulticloudNetworkAnchorDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.Id != nil {
		s.D.Set("id", *s.Res.Id)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ResourceAnchorId != nil {
		s.D.Set("resource_anchor_id", *s.Res.ResourceAnchorId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	// NOTE: Latest SDK (v65.105.0) has breaking changes where the LifecycleState property has been renamed to NetworkAnchorLifecycleState
	// s.D.Set("network_anchor_lifecycle_state", string(s.Res.LifecycleState))

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	s.D.Set("setup_mode", string(s.Res.SetupMode))

	if s.Res.ClusterPlacementGroupId != nil {
		s.D.Set("cluster_placement_group_id", *s.Res.ClusterPlacementGroupId)
	}

	if s.Res.OciMetadataItem != nil {
		s.D.Set("oci_metadata_item", []interface{}{OciNetworkMetadataToMap(s.Res.OciMetadataItem)})
	} else {
		s.D.Set("oci_metadata_item", nil)
	}

	if s.Res.CloudServiceProviderMetadataItem != nil {
		s.D.Set("cloud_service_provider_metadata_item", []interface{}{CloudServiceProviderNetworkMetadataItemToMap(s.Res.CloudServiceProviderMetadataItem)})
	} else {
		s.D.Set("cloud_service_provider_metadata_item", nil)
	}

	return nil
}

func CloudServiceProviderNetworkMetadataItemToMap(obj *oci_multicloud.CloudServiceProviderNetworkMetadataItem) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.OdbNetworkId != nil {
		result["odb_network_id"] = string(*obj.OdbNetworkId)
	}

	result["cidr_blocks"] = obj.CidrBlocks

	if obj.NetworkAnchorUri != nil {
		result["network_anchor_uri"] = string(*obj.NetworkAnchorUri)
	}

	dnsForwardingConfig := []interface{}{}
	for _, item := range obj.DnsForwardingConfig {
		dnsForwardingConfig = append(dnsForwardingConfig, item)
	}
	result["dns_forwarding_config"] = dnsForwardingConfig

	return result
}

func OciDnsToMap(obj *oci_multicloud.OciDns) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CustomDomainName != nil {
		result["custom_domain_name"] = string(*obj.CustomDomainName)
	}

	return result
}

func OciNetworkMetadataToMap(obj *oci_multicloud.OciNetworkMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	result["network_anchor_connection_status"] = string(obj.NetworkAnchorConnectionStatus)

	if obj.Vcn != nil {
		result["vcn"] = []interface{}{OciVcnToMap(obj.Vcn)}
	}

	if obj.Dns != nil {
		result["dns"] = []interface{}{OciDnsToMap(obj.Dns)}
	}

	subnets := []interface{}{}
	for _, item := range obj.Subnets {
		subnets = append(subnets, OciNetworkSubnetToMap(item))
	}
	result["subnets"] = subnets

	if obj.DnsListeningEndpointIpAddress != nil {
		result["dns_listening_endpoint_ip_address"] = string(*obj.DnsListeningEndpointIpAddress)
	}

	if obj.DnsForwardingEndpointIpAddress != nil {
		result["dns_forwarding_endpoint_ip_address"] = string(*obj.DnsForwardingEndpointIpAddress)
	}

	dnsForwardingConfig := []interface{}{}
	for _, item := range obj.DnsForwardingConfig {
		dnsForwardingConfig = append(dnsForwardingConfig, item)
	}
	result["dns_forwarding_config"] = dnsForwardingConfig

	return result
}

func OciNetworkSubnetToMap(obj oci_multicloud.OciNetworkSubnet) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.SubnetId != nil {
		result["subnet_id"] = string(*obj.SubnetId)
	}

	result["type"] = string(obj.Type)

	return result
}

func OciVcnToMap(obj *oci_multicloud.OciVcn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.VcnId != nil {
		result["vcn_id"] = string(*obj.VcnId)
	}

	result["cidr_blocks"] = obj.CidrBlocks

	result["backup_cidr_blocks"] = obj.BackupCidrBlocks

	if obj.DnsLabel != nil {
		result["dns_label"] = string(*obj.DnsLabel)
	}

	return result
}
