// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package metering_computation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_metering_computation "github.com/oracle/oci-go-sdk/v65/usageapi"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func MeteringComputationUsageStatementEmailRecipientsGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readMeteringComputationUsageStatementEmailRecipientsGroups,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"email_recipients_group_collection": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"items": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     tfresource.GetDataSourceItemSchema(MeteringComputationUsageStatementEmailRecipientsGroupResource()),
						},
					},
				},
			},
		},
	}
}

func readMeteringComputationUsageStatementEmailRecipientsGroups(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupsDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationUsageStatementEmailRecipientsGroupsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.ListEmailRecipientsGroupsResponse
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupsDataSourceCrud) Get() error {
	request := oci_metering_computation.ListEmailRecipientsGroupsRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.ListEmailRecipientsGroups(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListEmailRecipientsGroups(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("MeteringComputationUsageStatementEmailRecipientsGroupsDataSource-", MeteringComputationUsageStatementEmailRecipientsGroupsDataSource(), s.D))
	resources := []map[string]interface{}{}
	usageStatementEmailRecipientsGroup := map[string]interface{}{}

	items := []interface{}{}
	for _, item := range s.Res.Items {
		items = append(items, EmailRecipientsGroupSummaryToMap(item, s.D.Get("subscription_id").(string)))
	}
	usageStatementEmailRecipientsGroup["items"] = items

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		items = tfresource.ApplyFiltersInCollection(f.(*schema.Set), items, MeteringComputationUsageStatementEmailRecipientsGroupsDataSource().Schema["email_recipients_group_collection"].Elem.(*schema.Resource).Schema)
		usageStatementEmailRecipientsGroup["items"] = items
	}

	resources = append(resources, usageStatementEmailRecipientsGroup)
	if err := s.D.Set("email_recipients_group_collection", resources); err != nil {
		return err
	}

	return nil
}
