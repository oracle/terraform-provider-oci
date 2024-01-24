// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package bastion

import (
	"context"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_bastion "github.com/oracle/oci-go-sdk/v65/bastion"
)

func BastionSessionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["session_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(BastionSessionResource(), fieldMap, readSingularBastionSession)
}

func readSingularBastionSession(d *schema.ResourceData, m interface{}) error {
	sync := &BastionSessionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).BastionClient()

	return tfresource.ReadResource(sync)
}

type BastionSessionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_bastion.BastionClient
	Res    *oci_bastion.GetSessionResponse
}

func (s *BastionSessionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *BastionSessionDataSourceCrud) Get() error {
	request := oci_bastion.GetSessionRequest{}

	if sessionId, ok := s.D.GetOkExists("session_id"); ok {
		tmp := sessionId.(string)
		request.SessionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "bastion")

	response, err := s.Client.GetSession(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *BastionSessionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.BastionId != nil {
		s.D.Set("bastion_id", *s.Res.BastionId)
	}

	if s.Res.BastionName != nil {
		s.D.Set("bastion_name", *s.Res.BastionName)
	}

	if s.Res.BastionPublicHostKeyInfo != nil {
		s.D.Set("bastion_public_host_key_info", *s.Res.BastionPublicHostKeyInfo)
	}

	if s.Res.BastionUserName != nil {
		s.D.Set("bastion_user_name", *s.Res.BastionUserName)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.KeyDetails != nil {
		s.D.Set("key_details", []interface{}{PublicKeyDetailsToMap(s.Res.KeyDetails)})
	} else {
		s.D.Set("key_details", nil)
	}

	s.D.Set("key_type", s.Res.KeyType)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.SessionTtlInSeconds != nil {
		s.D.Set("session_ttl_in_seconds", *s.Res.SessionTtlInSeconds)
	}

	s.D.Set("ssh_metadata", s.Res.SshMetadata)

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TargetResourceDetails != nil {
		targetResourceDetailsArray := []interface{}{}
		if targetResourceDetailsMap := TargetResourceDetailsToMap(&s.Res.TargetResourceDetails); targetResourceDetailsMap != nil {
			targetResourceDetailsArray = append(targetResourceDetailsArray, targetResourceDetailsMap)
		}
		s.D.Set("target_resource_details", targetResourceDetailsArray)
	} else {
		s.D.Set("target_resource_details", nil)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeUpdated != nil {
		s.D.Set("time_updated", s.Res.TimeUpdated.String())
	}

	return nil
}
