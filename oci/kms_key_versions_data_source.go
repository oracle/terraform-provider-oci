// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KeyVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKeyVersions,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"key_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     GetDataSourceItemSchema(KeyVersionResource()),
			},
		},
	}
}

func readKeyVersions(d *schema.ResourceData, m interface{}) error {
	sync := &KeyVersionsDataSourceCrud{}
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

type KeyVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.ListKeyVersionsResponse
}

func (s *KeyVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KeyVersionsDataSourceCrud) Get() error {
	request := oci_kms.ListKeyVersionsRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "kms")

	response, err := s.Client.ListKeyVersions(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	request.Page = s.Res.OpcNextPage

	for request.Page != nil {
		listResponse, err := s.Client.ListKeyVersions(context.Background(), request)
		if err != nil {
			return err
		}

		s.Res.Items = append(s.Res.Items, listResponse.Items...)
		request.Page = listResponse.OpcNextPage
	}

	return nil
}

func (s *KeyVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(GenerateDataSourceID())
	resources := []map[string]interface{}{}

	for _, r := range s.Res.Items {
		keyVersion := map[string]interface{}{
			"key_id": *r.KeyId,
		}

		if r.CompartmentId != nil {
			keyVersion["compartment_id"] = *r.CompartmentId
		}

		if r.Id != nil {
			keyVersion["key_version_id"] = *r.Id
		}

		if r.TimeCreated != nil {
			keyVersion["time_created"] = r.TimeCreated.String()
		}

		if r.VaultId != nil {
			keyVersion["vault_id"] = *r.VaultId
		}

		resources = append(resources, keyVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = ApplyFilters(f.(*schema.Set), resources, KeyVersionsDataSource().Schema["key_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("key_versions", resources); err != nil {
		return err
	}

	return nil
}
