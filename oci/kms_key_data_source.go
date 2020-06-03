// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func init() {
	RegisterDatasource("oci_kms_key", KmsKeyDataSource())
}

func KmsKeyDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["management_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(KmsKeyResource(), fieldMap, readSingularKmsKey)
}

func readSingularKmsKey(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}

	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

type KmsKeyDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.GetKeyResponse
}

func (s *KmsKeyDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsKeyDataSourceCrud) Get() error {
	request := oci_kms.GetKeyRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.GetKey(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsKeyDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.CurrentKeyVersion != nil {
		s.D.Set("current_key_version", *s.Res.CurrentKeyVersion)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.KeyShape != nil {
		s.D.Set("key_shape", []interface{}{KeyShapeToMap(s.Res.KeyShape)})
	} else {
		s.D.Set("key_shape", nil)
	}

	if s.Res.RestoredFromKeyId != nil {
		s.D.Set("restored_from_key_id", *s.Res.RestoredFromKeyId)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TimeOfDeletion != nil {
		s.D.Set("time_of_deletion", s.Res.TimeOfDeletion.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}
