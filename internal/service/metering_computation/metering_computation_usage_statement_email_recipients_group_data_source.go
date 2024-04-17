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

func MeteringComputationUsageStatementEmailRecipientsGroupDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["compartment_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["email_recipients_group_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["subscription_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(MeteringComputationUsageStatementEmailRecipientsGroupResource(), fieldMap, readSingularMeteringComputationUsageStatementEmailRecipientsGroup)
}

func readSingularMeteringComputationUsageStatementEmailRecipientsGroup(d *schema.ResourceData, m interface{}) error {
	sync := &MeteringComputationUsageStatementEmailRecipientsGroupDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).UsageapiClient()

	return tfresource.ReadResource(sync)
}

type MeteringComputationUsageStatementEmailRecipientsGroupDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_metering_computation.UsageapiClient
	Res    *oci_metering_computation.GetEmailRecipientsGroupResponse
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupDataSourceCrud) Get() error {
	request := oci_metering_computation.GetEmailRecipientsGroupRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if emailRecipientsGroupId, ok := s.D.GetOkExists("email_recipients_group_id"); ok {
		tmp := emailRecipientsGroupId.(string)
		request.EmailRecipientsGroupId = &tmp
	}

	if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
		tmp := subscriptionId.(string)
		request.SubscriptionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "metering_computation")

	response, err := s.Client.GetEmailRecipientsGroup(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *MeteringComputationUsageStatementEmailRecipientsGroupDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	recipientsList := []interface{}{}
	for _, item := range s.Res.RecipientsList {
		recipientsList = append(recipientsList, EmailRecipientToMap(item))
	}
	s.D.Set("recipients_list", recipientsList)

	s.D.Set("state", s.Res.LifecycleState)

	return nil
}
