// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	oci_database "github.com/oracle/oci-go-sdk/v58/database"
)

func DatabaseKeyStoreResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createDatabaseKeyStore,
		Read:     readDatabaseKeyStore,
		Update:   updateDatabaseKeyStore,
		Delete:   deleteDatabaseKeyStore,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type_details": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"admin_username": {
							Type:     schema.TypeString,
							Required: true,
						},
						"connection_ips": {
							Type:     schema.TypeSet,
							Required: true,
							Set:      utils.LiteralTypeHashCodeForSets,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"secret_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"ORACLE_KEY_VAULT",
							}, true),
						},
						"vault_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},

			// Optional
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},

			// Computed
			"associated_databases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"lifecycle_details": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func createDatabaseKeyStore(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.CreateResource(d, sync)
}

func readDatabaseKeyStore(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.ReadResource(sync)
}

func updateDatabaseKeyStore(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()

	return tfresource.UpdateResource(d, sync)
}

func deleteDatabaseKeyStore(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseKeyStoreResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).DatabaseClient()
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type DatabaseKeyStoreResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_database.DatabaseClient
	Res                    *oci_database.KeyStore
	DisableNotFoundRetries bool
}

func (s *DatabaseKeyStoreResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *DatabaseKeyStoreResourceCrud) CreatedPending() []string {
	return []string{}
}

func (s *DatabaseKeyStoreResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_database.KeyStoreLifecycleStateActive),
	}
}

func (s *DatabaseKeyStoreResourceCrud) DeletedPending() []string {
	return []string{}
}

func (s *DatabaseKeyStoreResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_database.KeyStoreLifecycleStateDeleted),
	}
}

func (s *DatabaseKeyStoreResourceCrud) Create() error {
	request := oci_database.CreateKeyStoreRequest{}

	if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
		tmp := compartmentId.(string)
		request.CompartmentId = &tmp
	}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	if typeDetails, ok := s.D.GetOkExists("type_details"); ok {
		if tmpList := typeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "type_details", 0)
			tmp, err := s.mapToKeyStoreTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TypeDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.CreateKeyStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyStore
	return nil
}

func (s *DatabaseKeyStoreResourceCrud) Get() error {
	request := oci_database.GetKeyStoreRequest{}

	tmp := s.D.Id()
	request.KeyStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.GetKeyStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyStore
	return nil
}

func (s *DatabaseKeyStoreResourceCrud) Update() error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_database.UpdateKeyStoreRequest{}

	if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
		convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
		if err != nil {
			return err
		}
		request.DefinedTags = convertedDefinedTags
	}

	if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
		request.FreeformTags = utils.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
	}

	tmp := s.D.Id()
	request.KeyStoreId = &tmp

	if typeDetails, ok := s.D.GetOkExists("type_details"); ok {
		if tmpList := typeDetails.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "type_details", 0)
			tmp, err := s.mapToKeyStoreTypeDetails(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TypeDetails = tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	response, err := s.Client.UpdateKeyStore(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.KeyStore
	return nil
}

func (s *DatabaseKeyStoreResourceCrud) Delete() error {
	request := oci_database.DeleteKeyStoreRequest{}

	tmp := s.D.Id()
	request.KeyStoreId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.DeleteKeyStore(context.Background(), request)
	return err
}

func (s *DatabaseKeyStoreResourceCrud) SetData() error {
	associatedDatabases := []interface{}{}
	for _, item := range s.Res.AssociatedDatabases {
		associatedDatabases = append(associatedDatabases, KeyStoreAssociatedDatabaseDetailsToMap(item))
	}
	s.D.Set("associated_databases", associatedDatabases)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", tfresource.DefinedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	if s.Res.TypeDetails != nil {
		typeDetailsArray := []interface{}{}
		if typeDetailsMap := KeyStoreTypeDetailsToMap(&s.Res.TypeDetails, false); typeDetailsMap != nil {
			typeDetailsArray = append(typeDetailsArray, typeDetailsMap)
		}
		s.D.Set("type_details", typeDetailsArray)
	} else {
		s.D.Set("type_details", nil)
	}

	return nil
}

func KeyStoreAssociatedDatabaseDetailsToMap(obj oci_database.KeyStoreAssociatedDatabaseDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DbName != nil {
		result["db_name"] = string(*obj.DbName)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	return result
}

func (s *DatabaseKeyStoreResourceCrud) mapToKeyStoreTypeDetails(fieldKeyFormat string) (oci_database.KeyStoreTypeDetails, error) {
	var baseObject oci_database.KeyStoreTypeDetails
	//discriminator
	typeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type"))
	var type_ string
	if ok {
		type_ = typeRaw.(string)
	} else {
		type_ = "" // default value
	}
	switch strings.ToLower(type_) {
	case strings.ToLower("ORACLE_KEY_VAULT"):
		details := oci_database.KeyStoreTypeFromOracleKeyVaultDetails{}
		if adminUsername, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_username")); ok {
			tmp := adminUsername.(string)
			details.AdminUsername = &tmp
		}
		if connectionIps, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connection_ips")); ok {
			set := connectionIps.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "connection_ips")) {
				details.ConnectionIps = tmp
			}
		}
		if secretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_id")); ok {
			tmp := secretId.(string)
			details.SecretId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "vault_id")); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown type '%v' was specified", type_)
	}
	return baseObject, nil
}

func KeyStoreTypeDetailsToMap(obj *oci_database.KeyStoreTypeDetails, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_database.KeyStoreTypeFromOracleKeyVaultDetails:
		result["type"] = "ORACLE_KEY_VAULT"

		if v.AdminUsername != nil {
			result["admin_username"] = string(*v.AdminUsername)
		}

		result["connection_ips"] = v.ConnectionIps

		if v.SecretId != nil {
			result["secret_id"] = string(*v.SecretId)
		}

		if v.VaultId != nil {
			result["vault_id"] = string(*v.VaultId)
		}
	default:
		log.Printf("[WARN] Received 'type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *DatabaseKeyStoreResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_database.ChangeKeyStoreCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.KeyStoreId = &idTmp

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "database")

	_, err := s.Client.ChangeKeyStoreCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	if waitErr := tfresource.WaitForUpdatedState(s.D, s); waitErr != nil {
		return waitErr
	}

	return nil
}
