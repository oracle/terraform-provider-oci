// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package delegate_access_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_delegate_access_control "github.com/oracle/oci-go-sdk/v65/delegateaccesscontrol"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func DelegateAccessControlDelegationControlDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["delegation_control_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(DelegateAccessControlDelegationControlResource(), fieldMap, readSingularDelegateAccessControlDelegationControl)
}

func readSingularDelegateAccessControlDelegationControl(d *schema.ResourceData, m interface{}) error {
	sync := &DelegateAccessControlDelegationControlDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DelegateAccessControlClient()

	return tfresource.ReadResource(sync)
}

type DelegateAccessControlDelegationControlDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_delegate_access_control.DelegateAccessControlClient
	Res    *oci_delegate_access_control.GetDelegationControlResponse
}

func (s *DelegateAccessControlDelegationControlDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DelegateAccessControlDelegationControlDataSourceCrud) Get() error {
	request := oci_delegate_access_control.GetDelegationControlRequest{}

	if delegationControlId, ok := s.D.GetOkExists("delegation_control_id"); ok {
		tmp := delegationControlId.(string)
		request.DelegationControlId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "delegate_access_control")

	response, err := s.Client.GetDelegationControl(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DelegateAccessControlDelegationControlDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	s.D.Set("delegation_subscription_ids", s.Res.DelegationSubscriptionIds)

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.IsAutoApproveDuringMaintenance != nil {
		s.D.Set("is_auto_approve_during_maintenance", *s.Res.IsAutoApproveDuringMaintenance)
	}

	if s.Res.LifecycleStateDetails != nil {
		s.D.Set("lifecycle_state_details", *s.Res.LifecycleStateDetails)
	}

	s.D.Set("notification_message_format", s.Res.NotificationMessageFormat)

	if s.Res.NotificationTopicId != nil {
		s.D.Set("notification_topic_id", *s.Res.NotificationTopicId)
	}

	if s.Res.NumApprovalsRequired != nil {
		s.D.Set("num_approvals_required", *s.Res.NumApprovalsRequired)
	}

	s.D.Set("pre_approved_service_provider_action_names", s.Res.PreApprovedServiceProviderActionNames)

	s.D.Set("resource_ids", s.Res.ResourceIds)

	s.D.Set("resource_type", s.Res.ResourceType)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.SystemTags != nil {
		s.D.Set("system_tags", tfresource.SystemTagsToMap(s.Res.SystemTags))
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeDeleted != nil {
		s.D.Set("time_deleted", s.Res.TimeDeleted.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	if s.Res.VaultKeyId != nil {
		s.D.Set("vault_key_id", *s.Res.VaultKeyId)
	}

	return nil
}
