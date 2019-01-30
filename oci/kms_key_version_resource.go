// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"

	"log"
	"net/url"
	"regexp"

	oci_kms "github.com/oracle/oci-go-sdk/keymanagement"
)

func KeyVersionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: DefaultTimeout,
		Create:   createKeyVersion,
		Read:     readKeyVersion,
		Delete:   deleteKeyVersion,
		Schema: map[string]*schema.Schema{
			// Required
			"key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KeyVersionResourceCrud{}
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

	return CreateResource(d, sync)
}

func readKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KeyVersionResourceCrud{}
	sync.D = d
	endpoint, ok := d.GetOkExists("management_endpoint")
	if !ok {
		//Import use case:
		id := d.Id()
		regex, _ := regexp.Compile("^managementEndpoint/(.*)/keys/(.*)/keyVersions/(.*)$")
		tokens := regex.FindStringSubmatch(id)
		if len(tokens) == 4 {
			endpoint = tokens[1]
			d.Set("management_endpoint", endpoint)
			d.Set("key_id", tokens[2])
			d.Set("key_version_id", tokens[3])
			d.SetId(getKeyVersionCompositeId(tokens[2], tokens[3]))
		} else {
			return fmt.Errorf("id %s should be of format: managementEndpoint/{managementEndpoint}/keys/{keyId}/keyVersions/{keyVersionId}", id)
		}
	}
	client, err := m.(*OracleClients).KmsManagementClient(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return ReadResource(sync)
}

func deleteKeyVersion(d *schema.ResourceData, m interface{}) error {
	return nil
}

type KeyVersionResourceCrud struct {
	BaseCrud
	Client                 *oci_kms.KmsManagementClient
	Res                    *oci_kms.KeyVersion
	DisableNotFoundRetries bool
}

func (s *KeyVersionResourceCrud) ID() string {
	return getKeyVersionCompositeId(*s.Res.KeyId, *s.Res.Id)
}

func (s *KeyVersionResourceCrud) Create() error {
	request := oci_kms.CreateKeyVersionRequest{}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyVersion
	return nil
}

func (s *KeyVersionResourceCrud) Get() error {
	request := oci_kms.GetKeyVersionRequest{}

	keyId, keyVersionId, err := parseKeyVersionCompositeId(s.D.Id())
	if err == nil {
		request.KeyId = &keyId
		request.KeyVersionId = &keyVersionId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyVersion
	return nil
}

func (s *KeyVersionResourceCrud) SetData() error {

	keyId, keyVersionId, err := parseKeyVersionCompositeId(s.D.Id())
	if err == nil {
		s.D.Set("key_id", &keyId)
		s.D.Set("key_version_id", &keyVersionId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
	}

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.VaultId != nil {
		s.D.Set("vault_id", *s.Res.VaultId)
	}

	return nil
}

func getKeyVersionCompositeId(keyId string, keyVersionId string) string {
	keyId = url.PathEscape(keyId)
	keyVersionId = url.PathEscape(keyVersionId)
	compositeId := "keys/" + keyId + "/keyVersions/" + keyVersionId
	return compositeId
}

func parseKeyVersionCompositeId(compositeId string) (keyId string, keyVersionId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("keys/.*/keyVersions/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	keyId, _ = url.PathUnescape(parts[1])
	keyVersionId, _ = url.PathUnescape(parts[3])

	return
}
