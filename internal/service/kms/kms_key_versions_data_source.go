// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"regexp"

	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	oci_kms "github.com/oracle/oci-go-sdk/v56/keymanagement"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
)

func KmsKeyVersionsDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readKmsKeyVersions,
		Schema: map[string]*schema.Schema{
			"filter": tfresource.DataSourceFiltersSchema(),
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
				Elem:     tfresource.GetDataSourceItemSchema(KmsKeyVersionResource()),
			},
		},
	}
}

func readKmsKeyVersions(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyVersionsDataSourceCrud{}
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

type KmsKeyVersionsDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_kms.KmsManagementClient
	Res    *oci_kms.ListKeyVersionsResponse
}

func (s *KmsKeyVersionsDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *KmsKeyVersionsDataSourceCrud) Get() error {
	request := oci_kms.ListKeyVersionsRequest{}

	request.KeyId = getKeyID(s)

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(false, "kms")

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

// with resource discovery s.D.GetOkExists("key_id") can return one of the two things
// 1) keyId (key ocid) (or)
// 2) managementEndpoint/{managementEndpoint}/keys/{keyId}
// getKeyID method handles both and will return the key OCID
func getKeyID(s *KmsKeyVersionsDataSourceCrud) *string {
	var finalKeyId string
	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		regex, _ := regexp.Compile("^managementEndpoint/(.*)/keys/(.*)$")
		tokens := regex.FindStringSubmatch(keyId.(string))
		if len(tokens) == 3 {
			finalKeyId = tokens[2]
		} else {
			finalKeyId = keyId.(string)
		}
	}
	return &finalKeyId
}

func (s *KmsKeyVersionsDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(tfresource.GenerateDataSourceHashID("KmsKeyVersionsDataSource-", KmsKeyVersionsDataSource(), s.D))
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

		keyVersion["state"] = r.LifecycleState

		if r.TimeCreated != nil {
			keyVersion["time_created"] = r.TimeCreated.String()
		}

		if r.TimeOfDeletion != nil {
			keyVersion["time_of_deletion"] = r.TimeOfDeletion.String()
		}

		if r.VaultId != nil {
			keyVersion["vault_id"] = *r.VaultId
		}

		resources = append(resources, keyVersion)
	}

	if f, fOk := s.D.GetOkExists("filter"); fOk {
		resources = tfresource.ApplyFilters(f.(*schema.Set), resources, KmsKeyVersionsDataSource().Schema["key_versions"].Elem.(*schema.Resource).Schema)
	}

	if err := s.D.Set("key_versions", resources); err != nil {
		return err
	}

	return nil
}
