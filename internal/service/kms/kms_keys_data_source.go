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

func KmsKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsKeys,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
			"algorithm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"curve_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"length": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"protection_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     tfresource.GetDataSourceItemSchema(KmsKeyResource()),
			},
		},
	}
}

func readKmsKeys(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeysDataSourceCrud{}
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

type KmsKeysDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.ListKeysResponse
}

func (s *KmsKeysDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsKeysDataSourceCrud) Get() error {
	request := oci_kms.ListKeysRequest{}

	if algorithm, ok := s.D.GetOkExists("algorithm"); ok {
		request.Algorithm = oci_kms.ListKeysAlgorithmEnum(algorithm.(string))
	}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if curveId, ok := s.D.GetOkExists("curve_id"); ok {
		request.CurveId = oci_kms.ListKeysCurveIdEnum(curveId.(string))
	}

	if length, ok := s.D.GetOkExists("length"); ok {
		tmp := length.(int)
		request.Length = &tmp
	}

	if protectionMode, ok := s.D.GetOkExists("protection_mode"); ok {
		request.ProtectionMode = oci_kms.ListKeysProtectionModeEnum(protectionMode.(string))
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

	response, err := s.Client.ListKeys(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKeys(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *KmsKeysDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsKeysDataSource-", KmsKeysDataSource(), s.D))
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		key := map[string]interface{}{
			"compartment_id": *r.CompartmentId,
		}

		if r.DefinedTags != nil {
			key["defined_tags"] = tfresource.DefinedTagsToMap(r.DefinedTags)
		}

		if r.DisplayName != nil {
			key["display_name"] = *r.DisplayName
		}

		if r.ExternalKeyReferenceDetails != nil {
			key["external_key_reference_details"] = []interface{}{ExternalKeyReferenceDetailsToMap(r.ExternalKeyReferenceDetails)}
		} else {
			key["external_key_reference_details"] = nil
		}

		key["freeform_tags"] = r.FreeformTags

		if r.Id != nil {
			key["id"] = *r.Id
		}

		if r.IsAutoRotationEnabled != nil {
			key["is_auto_rotation_enabled"] = *r.IsAutoRotationEnabled
		}

		key["protection_mode"] = r.ProtectionMode

		key["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			key["time_created"] = r.TimeCreated.String()
		}

		if r.VaultId != nil {
			key["vault_id"] = *r.VaultId
		}

		resources = append(resources, key)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, KmsKeysDataSource().Schema["keys"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("keys", resources); err != nil {
		return err
	}

	return nil
}
