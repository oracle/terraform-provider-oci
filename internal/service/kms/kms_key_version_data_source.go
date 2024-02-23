// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func KmsKeyVersionDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["key_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["key_version_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	fieldMap["management_endpoint"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return tfresource.GetSingularDataSourceItemSchema(KmsKeyVersionResource(), fieldMap, readSingularKmsKeyVersion)
}

func readSingularKmsKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyVersionDataSourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		return fmt.Errorf("management endpoint missing")
	}
	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

type KmsKeyVersionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.GetKeyVersionResponse
}

func (s *KmsKeyVersionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsKeyVersionDataSourceCrud) Get() error {
	request := oci_kms.GetKeyVersionRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	if keyVersionId, ok := s.D.GetOkExists("key_version_id"); ok {
		tmp := keyVersionId.(string)
		request.KeyVersionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.GetKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *KmsKeyVersionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(getKeyVersionCompositeId(*s.Res.KeyId, *s.Res.Id))

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ExternalKeyReferenceDetails != nil {
		s.D.Set("external_key_reference_details", []interface{}{ExternalKeyReferenceDetailsToMap(s.Res.ExternalKeyReferenceDetails)})
	} else {
		s.D.Set("external_key_reference_details", nil)
	}

	if s.Res.IsAutoRotated != nil {
		s.D.Set("is_auto_rotated", *s.Res.IsAutoRotated)
	}

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.PublicKey != nil {
		s.D.Set("public_key", *s.Res.PublicKey)
	}

	if s.Res.ReplicaDetails != nil {
		s.D.Set("replica_details", []interface{}{KeyVersionReplicaDetailsToMap(s.Res.ReplicaDetails)})
	} else {
		s.D.Set("replica_details", nil)
	}

	if s.Res.RestoredFromKeyVersionId != nil {
		s.D.Set("restored_from_key_version_id", *s.Res.RestoredFromKeyVersionId)
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

	if s.Res.RestoredFromKeyVersionId != nil {
		s.D.Set("restored_from_key_id", *s.Res.RestoredFromKeyVersionId)
	}
	return nil
}
