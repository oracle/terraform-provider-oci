// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package tenantmanagercontrolplane

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_tenantmanagercontrolplane "github.com/oracle/oci-go-sdk/v65/tenantmanagercontrolplane"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func TenantmanagercontrolplaneDomainGovernancesDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readTenantmanagercontrolplaneDomainGovernances,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"domain_governance_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_governance_collection": {
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
									"compartment_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"domain_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"ons_subscription_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"ons_topic_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"subscription_email": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"defined_tags": {
										Type:             schema.TypeMap,
										Optional:         true,
										Computed:         true,
										DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
										Elem:             schema.TypeString,
									},
									"freeform_tags": {
										Type:     schema.TypeMap,
										Optional: true,
										Computed: true,
										Elem:     schema.TypeString,
									},

									// Computed
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"is_governance_enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"owner_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"state": {
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
							},
						},
					},
				},
			},
		},
	}
}

func readTenantmanagercontrolplaneDomainGovernances(d *schema.ResourceData, m interface{}) error {
	sync := &TenantmanagercontrolplaneDomainGovernancesDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DomainGovernanceClient()

	return tfresource.ReadResource(sync)
}

type TenantmanagercontrolplaneDomainGovernancesDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_tenantmanagercontrolplane.DomainGovernanceClient
	Res    *oci_tenantmanagercontrolplane.ListDomainGovernancesResponse
}

func (s *TenantmanagercontrolplaneDomainGovernancesDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *TenantmanagercontrolplaneDomainGovernancesDataSourceCrud) Get() error {
	request := oci_tenantmanagercontrolplane.ListDomainGovernancesRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if domainGovernanceId, ok := s.D.GetOkExists("id"); ok {
		tmp := domainGovernanceId.(string)
		request.DomainGovernanceId = &tmp
	}

	if domainId, ok := s.D.GetOkExists("domain_id"); ok {
		tmp := domainId.(string)
		request.DomainId = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if state, ok := s.D.GetOkExists("state"); ok {
		request.LifecycleState = oci_tenantmanagercontrolplane.ListDomainGovernancesLifecycleStateEnum(state.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "tenantmanagercontrolplane")

	response, err := s.Client.ListDomainGovernances(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListDomainGovernances(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *TenantmanagercontrolplaneDomainGovernancesDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("TenantmanagercontrolplaneDomainGovernancesDataSource-", TenantmanagercontrolplaneDomainGovernancesDataSource(), s.D))
	resources := []map[string]interface{}{}
	domainGovernance := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, DomainGovernanceSummaryToMap(item))
	}
	domainGovernance["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, TenantmanagercontrolplaneDomainGovernancesDataSource().Schema["domain_governance_collection"].Elem.(*schema.Resource).Schema)
		domainGovernance["items"] = items
	}

	resources = append(resources, domainGovernance)
	if err := s.D.Set("domain_governance_collection", resources); err != nil {
		return err
	}

	return nil
}

func DomainGovernanceSummaryToMap(obj oci_tenantmanagercontrolplane.DomainGovernanceSummary) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefinedTags != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.DefinedTags)
	}

	if obj.DomainId != nil {
		result["domain_id"] = string(*obj.DomainId)
	}

	result["freeform_tags"] = obj.FreeformTags

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IsGovernanceEnabled != nil {
		result["is_governance_enabled"] = bool(*obj.IsGovernanceEnabled)
	}

	if obj.OnsSubscriptionId != nil {
		result["ons_subscription_id"] = string(*obj.OnsSubscriptionId)
	}

	if obj.OnsTopicId != nil {
		result["ons_topic_id"] = string(*obj.OnsTopicId)
	}

	if obj.OwnerId != nil {
		result["owner_id"] = string(*obj.OwnerId)
	}

	result["state"] = string(obj.LifecycleState)

	if obj.SubscriptionEmail != nil {
		result["subscription_email"] = string(*obj.SubscriptionEmail)
	}

	if obj.SystemTags != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.SystemTags)
	}

	if obj.TimeCreated != nil {
		result["time_created"] = obj.TimeCreated.String()
	}

	if obj.TimeUpdated != nil {
		result["time_updated"] = obj.TimeUpdated.String()
	}

	return result
}
