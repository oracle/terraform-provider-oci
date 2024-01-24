// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package kms

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"log"
	"net/url"
	"regexp"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_kms "github.com/oracle/oci-go-sdk/v65/keymanagement"
)

func KmsKeyVersionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createKmsKeyVersion,
		Read:     readKmsKeyVersion,
		Delete:   deleteKmsKeyVersion,
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
			"external_key_version_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"time_of_deletion": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_key_reference_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"external_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_key_version_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replica_details": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"replication_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"restored_from_key_version_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
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
			"restored_from_key_id": {
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

func createKmsKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyVersionResourceCrud{}
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

	return tfresource.CreateResource(d, sync)
}

func readKmsKeyVersion(d *schema.ResourceData, m interface{}) error {
	sync := &KmsKeyVersionResourceCrud{}
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
	client, err := m.(*client.OracleClients).KmsManagementClientWithEndpoint(endpoint.(string))
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

func deleteKmsKeyVersion(d *schema.ResourceData, m interface{}) error {
	// prevent kms version deletion as part of testing as version deletion is only applicable when the version is not the current version of the key
	disableKmsVersionDeletion, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("disable_kms_version_delete", "false"))
	if disableKmsVersionDeletion {
		return nil
	}

	sync := &KmsKeyVersionResourceCrud{}
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

	return tfresource.DeleteResource(d, sync)
}

// existing ID() returned by getKeyVersionCompositeId method is of format keys/(.*)/keyVersions/(.*) and
// that is not the format that readKmsKeyVersion expects managementEndpoint/{managementEndpoint}/keys/{keyId}/keyVersions/{keyVersionId}
// GetCompositeKeyVersionId is only used for resource discovery and it returns the Id as expected by readKmsKeyVersion
// terraform import oci_kms_key_version.test_key_version "managementEndpoint/{managementEndpoint}/keys/{keyId}/keyVersions/{keyVersionId}"
func GetCompositeKeyVersionId(managementEndpoint string, keyId string, keyVersionId string) string {
	compositeId := "managementEndpoint/" + managementEndpoint + "/keys/" + keyId + "/keyVersions/" + keyVersionId
	return compositeId
}

type KmsKeyVersionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_kms.KmsManagementClient
	Res                    *oci_kms.KeyVersion
	DisableNotFoundRetries bool
}

func (s *KmsKeyVersionResourceCrud) ID() string {
	return getKeyVersionCompositeId(*s.Res.KeyId, *s.Res.Id)
}

func (s *KmsKeyVersionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_kms.KeyVersionLifecycleStateCreating),
		string(oci_kms.KeyVersionLifecycleStateEnabling),
	}
}

func (s *KmsKeyVersionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_kms.KeyVersionLifecycleStateEnabled),
	}
}

func (s *KmsKeyVersionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_kms.KeyVersionLifecycleStateDisabled),
		string(oci_kms.KeyVersionLifecycleStateDeleting),
		string(oci_kms.KeyVersionLifecycleStateSchedulingDeletion),
	}
}

func (s *KmsKeyVersionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_kms.KeyVersionLifecycleStateDeleted),
		string(oci_kms.KeyVersionLifecycleStatePendingDeletion),
	}
}

func (s *KmsKeyVersionResourceCrud) Create() error {
	request := oci_kms.CreateKeyVersionRequest{}

	if externalKeyVersionId, ok := s.D.GetOkExists("external_key_version_id"); ok {
		tmp := externalKeyVersionId.(string)
		request.CreateKeyMetadataDetails.ExternalKeyVersionId = &tmp
	}

	if keyId, ok := s.D.GetOkExists("key_id"); ok {
		tmp := keyId.(string)
		request.KeyId = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.CreateKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}
	//has to wait some time, otherwise subsequent querying will fail
	time.Sleep(time.Second * 30)
	s.Res = &response.KeyVersion
	return nil
}

func (s *KmsKeyVersionResourceCrud) Get() error {
	request := oci_kms.GetKeyVersionRequest{}

	keyId, keyVersionId, err := parseKeyVersionCompositeId(s.D.Id())
	if err == nil {
		request.KeyId = &keyId
		request.KeyVersionId = &keyVersionId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	response, err := s.Client.GetKeyVersion(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyVersion
	return nil
}

func (s *KmsKeyVersionResourceCrud) Delete() error {
	request := oci_kms.ScheduleKeyVersionDeletionRequest{}

	keyId, keyVersionId, err := parseKeyVersionCompositeId(s.D.Id())
	if err == nil {
		request.KeyId = &keyId
		request.KeyVersionId = &keyVersionId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
		return err
	}

	if timeOfDeletion, ok := s.D.GetOkExists("time_of_deletion"); ok {
		tmpTime, err := time.Parse(time.RFC3339Nano, timeOfDeletion.(string))
		if err != nil {
			return err
		}
		request.TimeOfDeletion = &oci_common.SDKTime{Time: tmpTime}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "kms")

	_, error := s.Client.ScheduleKeyVersionDeletion(context.Background(), request)
	return error
}

func (s *KmsKeyVersionResourceCrud) SetData() error {

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

	if s.Res.ExternalKeyReferenceDetails != nil {
		s.D.Set("external_key_reference_details", []interface{}{ExternalKeyReferenceDetailsToMap(s.Res.ExternalKeyReferenceDetails)})
	} else {
		s.D.Set("external_key_reference_details", nil)
	}

	if s.Res.IsPrimary != nil {
		s.D.Set("is_primary", *s.Res.IsPrimary)
	}

	if s.Res.KeyId != nil {
		s.D.Set("key_id", *s.Res.KeyId)
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

func KeyVersionReplicaDetailsToMap(obj *oci_kms.KeyVersionReplicaDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReplicationId != nil {
		result["replication_id"] = string(*obj.ReplicationId)
	}

	return result
}
