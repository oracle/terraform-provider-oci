// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package golden_gate

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_golden_gate "github.com/oracle/oci-go-sdk/v65/goldengate"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func GoldenGateConnectionResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts:      tfresource.DefaultTimeout,
		CreateContext: createGoldenGateConnectionWithContext,
		ReadContext:   readGoldenGateConnectionWithContext,
		UpdateContext: updateGoldenGateConnectionWithContext,
		DeleteContext: deleteGoldenGateConnectionWithContext,
		Schema: map[string]*schema.Schema{
			// Required
			"compartment_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"connection_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
				ValidateFunc: validation.StringInSlice([]string{
					"AI_MODEL",
					"AI_PROVIDER",
					"AMAZON_KINESIS",
					"AMAZON_REDSHIFT",
					"AMAZON_S3",
					"AZURE_DATA_LAKE_STORAGE",
					"AZURE_SYNAPSE_ANALYTICS",
					"DATABRICKS",
					"DATABRICKS_ENGINE",
					"DB2",
					"ELASTICSEARCH",
					"GENERIC",
					"GOLDENGATE",
					"GOOGLE_BIGQUERY",
					"GOOGLE_CLOUD_STORAGE",
					"GOOGLE_PUBSUB",
					"HDFS",
					"ICEBERG",
					"JAVA_MESSAGE_SERVICE",
					"KAFKA",
					"KAFKA_SCHEMA_REGISTRY",
					"MICROSOFT_FABRIC",
					"MICROSOFT_SQLSERVER",
					"MONGODB",
					"MYSQL",
					"OCI_OBJECT_STORAGE",
					"ORACLE",
					"ORACLE_AI_DATA_PLATFORM",
					"ORACLE_AI_DATA_PLATFORM_ENGINE",
					"ORACLE_NOSQL",
					"POSTGRESQL",
					"REDIS",
					"SNOWFLAKE",
					"SPARK",
				}, true),
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"technology_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			// Optional
			"access_key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_key": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("account_key", "account_key_secret_id"),
			},
			"account_key_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"additional_attributes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"artifact_location": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"authentication_type": {
										Type:             schema.TypeString,
										Required:         true,
										DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
										ValidateFunc: validation.StringInSlice([]string{
											"CUSTOMER_SECRET_KEY",
											"PRE_AUTHENTICATED_REQUEST",
											"RESOURCE_PRINCIPAL",
											"USER_PRINCIPAL",
										}, true),
									},

									// Optional
									"access_key_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bucket": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"par_url": {
										Type:       schema.TypeString,
										Optional:   true,
										Sensitive:  true,
										Deprecated: tfresource.FieldDeprecatedForAnother("par_url", "par_url_secret_id"),
									},
									"par_url_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"path": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"private_key_file": {
										Type:       schema.TypeString,
										Optional:   true,
										Sensitive:  true,
										Deprecated: tfresource.FieldDeprecatedForAnother("private_key_file", "private_key_file_secret_id"),
									},
									"private_key_file_secret_id": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										Sensitive: true,
									},
									"private_key_passphrase": {
										Type:       schema.TypeString,
										Optional:   true,
										Sensitive:  true,
										Deprecated: tfresource.FieldDeprecatedForAnother("private_key_passphrase", "private_key_passphrase_secret_id"),
									},
									"private_key_passphrase_secret_id": {
										Type:      schema.TypeString,
										Optional:  true,
										Computed:  true,
										Sensitive: true,
									},
									"public_key_fingerprint": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"region": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"secret_access_key": {
										Type:       schema.TypeString,
										Optional:   true,
										Sensitive:  true,
										Deprecated: tfresource.FieldDeprecatedForAnother("secret_access_key", "secret_access_key_secret_id"),
									},
									"secret_access_key_secret_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"tenancy_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"user_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"location_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AZURE_DATE_LAKE_STORAGE",
								"OCI_OBJECT_STORAGE",
								"UC_VOLUME",
							}, true),
						},

						// Optional
						"account_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_key_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoint_suffix": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"filesystem": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"path": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"should_overwrite_existing": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"storage_account": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"auth_details": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"auth_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"API_KEY",
								"OCI_GEN_AI",
							}, true),
						},

						// Optional
						"api_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"api_key_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"base_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_fingerprint": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenancy_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"auth_secret": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("auth_secret", "auth_secret_secret_id"),
			},
			"auth_secret_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_token": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("auth_token", "auth_token_secret_id"),
			},
			"auth_token_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_authority_host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bootstrap_servers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
						"private_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"catalog": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"catalog_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"GLUE",
								"HADOOP",
								"NESSIE",
								"POLARIS",
								"REST",
							}, true),
						},

						// Optional
						"branch": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_secret": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("client_secret", "client_secret_secret_id"),
						},
						"client_secret_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"glue_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"principal_role": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"properties": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("properties", "properties_secret_id"),
						},
						"properties_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uri": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"client_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_secret": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("client_secret", "client_secret_secret_id"),
			},
			"client_secret_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cluster_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_placement_group_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"connection_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_string": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"connection_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"consumer_properties": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"core_site_xml": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"db_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"defined_tags": {
				Type:             schema.TypeMap,
				Optional:         true,
				Computed:         true,
				DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
				Elem:             schema.TypeString,
			},
			"deployment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"does_use_secret_ids": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"freeform_tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_connection_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_initial_context_factory": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_provider_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"jndi_security_credentials": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("jndi_security_credentials", "jndi_security_credentials_secret_id"),
			},
			"jndi_security_credentials_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"jndi_security_principal": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_store": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("key_store", "key_store_secret_id"),
			},
			"key_store_password": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("key_store_password", "key_store_password_secret_id"),
			},
			"key_store_password_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_store_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_store_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"locks": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"message": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						// Computed
						"related_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"time_created": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"max_input_chars": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"model_key": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_lock_override": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"nsg_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Set:      tfresource.LiteralTypeHashCodeForSets,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"password": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("password", "password_secret_id"),
			},
			"password_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"private_key_file": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("private_key_file", "private_key_file_secret_id"),
			},
			"private_key_file_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key_passphrase": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("private_key_passphrase", "private_key_passphrase_secret_id"),
			},
			"private_key_passphrase_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"producer_properties": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"provider_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"provider_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"public_key_fingerprint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redis_cluster_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_auth_secret_key": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("rest_auth_secret_key", "rest_auth_secret_key_secret_id"),
			},
			"rest_auth_secret_key_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rest_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_method": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rpc_cipher": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sas_token": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("sas_token", "sas_token_secret_id"),
			},
			"sas_token_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secret_access_key": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("secret_access_key", "secret_access_key_secret_id"),
			},
			"secret_access_key_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"security_attributes": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"security_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_account_key_file": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("service_account_key_file", "service_account_key_file_secret_id"),
			},
			"service_account_key_file_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"session_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"should_use_jndi": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_use_resource_principal": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"should_validate_server_certificate": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ssl_ca": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_cert": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_client_keystash": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("ssl_client_keystash", "ssl_client_keystash_secret_id"),
			},
			"ssl_client_keystash_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_client_keystoredb": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("ssl_client_keystoredb", "ssl_client_keystoredb_secret_id"),
			},
			"ssl_client_keystoredb_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_crl": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_key": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("ssl_key", "ssl_key_secret_id"),
			},
			"ssl_key_password": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("ssl_key_password", "ssl_key_password_secret_id"),
			},
			"ssl_key_password_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_key_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_server_certificate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storage": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"storage_type": {
							Type:             schema.TypeString,
							Required:         true,
							DiffSuppressFunc: tfresource.EqualIgnoreCaseSuppressDiff,
							ValidateFunc: validation.StringInSlice([]string{
								"AMAZON_S3",
								"AZURE_DATA_LAKE_STORAGE",
								"GOOGLE_CLOUD_STORAGE",
								"NONE",
								"OCI_OBJECT_STORAGE_S3_API",
							}, true),
						},

						// Optional
						"access_key_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_key": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("account_key", "account_key_secret_id"),
						},
						"account_key_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"account_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bucket": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"container": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"project_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scheme_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secret_access_key": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("secret_access_key", "secret_access_key_secret_id"),
						},
						"secret_access_key_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_account_key_file": {
							Type:       schema.TypeString,
							Optional:   true,
							Computed:   true,
							Deprecated: tfresource.FieldDeprecatedForAnother("service_account_key_file", "service_account_key_file_secret_id"),
						},
						"service_account_key_file_secret_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"storage_credential_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stream_pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"submit_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"subscription_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tenancy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_ca_file": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_certificate_key_file": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Deprecated: tfresource.FieldDeprecatedForAnother("tls_certificate_key_file", "tls_certificate_key_file_secret_id"),
			},
			"tls_certificate_key_file_password": {
				Type:       schema.TypeString,
				Optional:   true,
				Computed:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("tls_certificate_key_file_password", "tls_certificate_key_file_password_secret_id"),
			},
			"tls_certificate_key_file_password_secret_id": {
				Type:      schema.TypeString,
				Optional:  true,
				Computed:  true,
				Sensitive: true,
			},
			"tls_certificate_key_file_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tls_protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trust_store": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("trust_store", "trust_store_secret_id"),
			},
			"trust_store_password": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("trust_store_password", "trust_store_password_secret_id"),
			},
			"trust_store_password_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_store_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"trust_store_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vault_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wallet": {
				Type:       schema.TypeString,
				Optional:   true,
				Sensitive:  true,
				Deprecated: tfresource.FieldDeprecatedForAnother("wallet", "wallet_secret_id"),
			},
			"wallet_secret_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"workspace_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			// Computed
			"ingress_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional

						// Computed
						"ingress_ip": {
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
			"private_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_tags": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_refresh": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createGoldenGateConnectionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	err := tfresource.HandleDiagError(m, tfresource.CreateResourceWithContext(ctx, d, sync))
	// always rewrite to false
	sync.D.Set("trigger_refresh", false)
	return err
}

func readGoldenGateConnectionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	return tfresource.HandleDiagError(m, tfresource.ReadResourceWithContext(ctx, sync))
}

func updateGoldenGateConnectionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()

	var refreshOnly = false
	if value, ok := sync.D.GetOkExists("trigger_refresh"); ok && value.(bool) {
		// if there is any other change than trigger_refresh, let's call update, otherwise refresh is enough
		refreshOnly = !sync.D.HasChangeExcept("trigger_refresh")
	}

	var err error
	if refreshOnly {
		err = sync.refreshConnection()
	} else {
		err = tfresource.UpdateResourceWithContext(ctx, d, sync)
	}

	// always rewrite to false
	sync.D.Set("trigger_refresh", false)
	return tfresource.HandleDiagError(m, err)
}

func (s *GoldenGateConnectionResourceCrud) refreshConnection() error {
	refreshConnectionRequest := oci_golden_gate.RefreshConnectionRequest{}

	idTmp := s.D.Id()
	refreshConnectionRequest.ConnectionId = &idTmp
	refreshConnectionRequest.RefreshConnectionDetails = oci_golden_gate.DefaultRefreshConnectionDetails{}
	refreshConnectionRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		refreshConnectionRequest.IsLockOverride = &tmp
	}

	response, err := s.Client.RefreshConnection(context.Background(), refreshConnectionRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, refreshWorkRequestErr := connectionWaitForWorkRequest(workId, "goldengateconnection",
		oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate), s.DisableNotFoundRetries, s.Client)
	if refreshWorkRequestErr != nil {
		return refreshWorkRequestErr
	}

	if e := s.GetWithContext(context.Background()); e != nil {
		return e
	}

	if e := s.SetData(); e != nil {
		return e
	}

	return nil
}

func deleteGoldenGateConnectionWithContext(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	sync := &GoldenGateConnectionResourceCrud{}
	sync.D = d
	sync.Client = m.(*client.OracleClients).GoldenGateClient()
	sync.DisableNotFoundRetries = true

	return tfresource.HandleDiagError(m, tfresource.DeleteResourceWithContext(ctx, d, sync))
}

type GoldenGateConnectionResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_golden_gate.GoldenGateClient
	Res                    *oci_golden_gate.Connection
	DisableNotFoundRetries bool
}

func (s *GoldenGateConnectionResourceCrud) ID() string {
	connection := *s.Res
	return *connection.GetId()
}

func (s *GoldenGateConnectionResourceCrud) CreatedPending() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateCreating),
	}
}

func (s *GoldenGateConnectionResourceCrud) CreatedTarget() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateActive),
		string(oci_golden_gate.ConnectionLifecycleStateNeedsAttention),
	}
}

func (s *GoldenGateConnectionResourceCrud) DeletedPending() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateDeleting),
	}
}

func (s *GoldenGateConnectionResourceCrud) DeletedTarget() []string {
	return []string{
		string(oci_golden_gate.ConnectionLifecycleStateDeleted),
	}
}

func (s *GoldenGateConnectionResourceCrud) CreateWithContext(ctx context.Context) error {
	request := oci_golden_gate.CreateConnectionRequest{}
	err := s.populateTopLevelPolymorphicCreateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.CreateConnection(context.Background(), request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	var identifier *string
	identifier = response.GetId()
	if identifier != nil {
		s.D.SetId(*identifier)
	}
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeCreated, s.D.Timeout(schema.TimeoutCreate))
}

func (s *GoldenGateConnectionResourceCrud) getConnectionFromWorkRequest(workId *string, retryPolicy *oci_common.RetryPolicy,
	actionTypeEnum oci_golden_gate.ActionTypeEnum, timeout time.Duration) error {

	// Wait until it finishes
	connectionId, err := connectionWaitForWorkRequest(workId, "goldengateconnection",
		actionTypeEnum, timeout, s.DisableNotFoundRetries, s.Client)

	if err != nil {
		return err
	}
	s.D.SetId(*connectionId)

	return s.GetWithContext(context.Background())
}

func connectionWorkRequestShouldRetryFunc(timeout time.Duration) func(response oci_common.OCIOperationResponse) bool {
	startTime := time.Now()
	stopTime := startTime.Add(timeout)
	return func(response oci_common.OCIOperationResponse) bool {

		// Stop after timeout has elapsed
		if time.Now().After(stopTime) {
			return false
		}

		// Make sure we stop on default rules
		if tfresource.ShouldRetry(response, false, "golden_gate", startTime) {
			return true
		}

		//  Only stop if the time Finished is set
		if workRequestResponse, ok := response.Response.(oci_golden_gate.GetWorkRequestResponse); ok {
			return workRequestResponse.TimeFinished == nil
		}
		return false
	}
}

func connectionWaitForWorkRequest(wId *string, entityType string, action oci_golden_gate.ActionTypeEnum,
	timeout time.Duration, disableFoundRetries bool, client *oci_golden_gate.GoldenGateClient) (*string, error) {
	retryPolicy := tfresource.GetRetryPolicy(disableFoundRetries, "golden_gate")
	retryPolicy.ShouldRetryOperation = connectionWorkRequestShouldRetryFunc(timeout)

	response := oci_golden_gate.GetWorkRequestResponse{}
	stateConf := &retry.StateChangeConf{
		Pending: []string{
			string(oci_golden_gate.OperationStatusInProgress),
			string(oci_golden_gate.OperationStatusAccepted),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Target: []string{
			string(oci_golden_gate.OperationStatusSucceeded),
			string(oci_golden_gate.OperationStatusFailed),
			string(oci_golden_gate.OperationStatusCanceled),
		},
		Refresh: func() (interface{}, string, error) {
			var err error
			response, err = client.GetWorkRequest(context.Background(),
				oci_golden_gate.GetWorkRequestRequest{
					WorkRequestId: wId,
					RequestMetadata: oci_common.RequestMetadata{
						RetryPolicy: retryPolicy,
					},
				})
			wr := &response.WorkRequest
			return wr, string(wr.Status), err
		},
		Timeout: timeout,
	}
	if _, e := stateConf.WaitForState(); e != nil {
		return nil, e
	}

	var identifier *string
	// The work request response contains an array of objects that finished the operation
	for _, res := range response.Resources {
		if strings.Contains(strings.ToLower(*res.EntityType), entityType) {
			if res.ActionType == action {
				identifier = res.Identifier
				break
			}
		}
	}

	// The workrequest may have failed, check for errors if identifier is not found or work failed or got cancelled
	if identifier == nil || response.Status == oci_golden_gate.OperationStatusFailed || response.Status == oci_golden_gate.OperationStatusCanceled {
		return nil, getErrorFromGoldenGateConnectionWorkRequest(client, wId, retryPolicy, entityType, action)
	}

	return identifier, nil
}

func getErrorFromGoldenGateConnectionWorkRequest(client *oci_golden_gate.GoldenGateClient, workId *string, retryPolicy *oci_common.RetryPolicy, entityType string, action oci_golden_gate.ActionTypeEnum) error {
	response, err := client.ListWorkRequestErrors(context.Background(),
		oci_golden_gate.ListWorkRequestErrorsRequest{
			WorkRequestId: workId,
			RequestMetadata: oci_common.RequestMetadata{
				RetryPolicy: retryPolicy,
			},
		})
	if err != nil {
		return err
	}

	allErrs := make([]string, 0)
	for _, wrkErr := range response.Items {
		allErrs = append(allErrs, *wrkErr.Message)
	}
	errorMessage := strings.Join(allErrs, "\n")

	workRequestErr := fmt.Errorf("work request did not succeed, workId: %s, entity: %s, action: %s. Message: %s", *workId, entityType, action, errorMessage)

	return workRequestErr
}

func (s *GoldenGateConnectionResourceCrud) GetWithContext(ctx context.Context) error {
	request := oci_golden_gate.GetConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.GetConnection(ctx, request)
	if err != nil {
		return err
	}

	s.Res = &response.Connection
	return nil
}

func (s *GoldenGateConnectionResourceCrud) UpdateWithContext(ctx context.Context) error {
	if compartment, ok := s.D.GetOkExists("compartment_id"); ok && s.D.HasChange("compartment_id") {
		oldRaw, newRaw := s.D.GetChange("compartment_id")
		if newRaw != "" && oldRaw != "" {
			err := s.updateCompartment(compartment)
			if err != nil {
				return err
			}
		}
	}
	request := oci_golden_gate.UpdateConnectionRequest{}
	err := s.populateTopLevelPolymorphicUpdateConnectionRequest(&request)
	if err != nil {
		return err
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.UpdateConnection(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}

func (s *GoldenGateConnectionResourceCrud) DeleteWithContext(ctx context.Context) error {
	request := oci_golden_gate.DeleteConnectionRequest{}

	tmp := s.D.Id()
	request.ConnectionId = &tmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		request.IsLockOverride = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.DeleteConnection(ctx, request)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	// Wait until it finishes
	_, delWorkRequestErr := connectionWaitForWorkRequest(workId, "goldengateconnection",
		oci_golden_gate.ActionTypeDeleted, s.D.Timeout(schema.TimeoutDelete), s.DisableNotFoundRetries, s.Client)
	return delWorkRequestErr
}

func (s *GoldenGateConnectionResourceCrud) SetData() error {
	switch v := (*s.Res).(type) {
	case oci_golden_gate.AiModelConnection:
		s.D.Set("connection_type", "AI_MODEL")

		if v.AuthDetails != nil {
			authDetailsArray := []interface{}{}
			if authDetailsMap := AiModelAuthDetailsToMap(&v.AuthDetails); authDetailsMap != nil {
				authDetailsMap = preserveAiModelAuthDetailsWriteOnlyFields(s.D, authDetailsMap)
				authDetailsArray = append(authDetailsArray, authDetailsMap)
			}
			s.D.Set("auth_details", authDetailsArray)
		} else {
			s.D.Set("auth_details", nil)
		}

		if v.MaxInputChars != nil {
			s.D.Set("max_input_chars", *v.MaxInputChars)
		}

		if v.ModelKey != nil {
			s.D.Set("model_key", *v.ModelKey)
		}

		if v.ProviderId != nil {
			s.D.Set("provider_id", *v.ProviderId)
		}

		s.D.Set("provider_type", v.ProviderType)

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AiProviderConnection:
		s.D.Set("connection_type", "AI_PROVIDER")

		if v.AuthDetails != nil {
			authDetailsArray := []interface{}{}
			if authDetailsMap := AiModelAuthDetailsToMap(&v.AuthDetails); authDetailsMap != nil {
				authDetailsMap = preserveAiModelAuthDetailsWriteOnlyFields(s.D, authDetailsMap)
				authDetailsArray = append(authDetailsArray, authDetailsMap)
			}
			s.D.Set("auth_details", authDetailsArray)
		} else {
			s.D.Set("auth_details", nil)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		} else {
			s.D.Set("security_attributes", map[string]interface{}{})
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AmazonKinesisConnection:
		s.D.Set("connection_type", "AMAZON_KINESIS")

		if v.AccessKeyId != nil {
			s.D.Set("access_key_id", *v.AccessKeyId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			s.D.Set("secret_access_key_secret_id", *v.SecretAccessKeySecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AmazonRedshiftConnection:
		s.D.Set("connection_type", "AMAZON_REDSHIFT")

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AmazonS3Connection:
		s.D.Set("connection_type", "AMAZON_S3")

		if v.AccessKeyId != nil {
			s.D.Set("access_key_id", *v.AccessKeyId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			s.D.Set("secret_access_key_secret_id", *v.SecretAccessKeySecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AzureDataLakeStorageConnection:
		s.D.Set("connection_type", "AZURE_DATA_LAKE_STORAGE")

		if v.AccountKeySecretId != nil {
			s.D.Set("account_key_secret_id", *v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			s.D.Set("account_name", *v.AccountName)
		}

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.AzureAuthorityHost != nil {
			s.D.Set("azure_authority_host", *v.AzureAuthorityHost)
		}

		if v.AzureTenantId != nil {
			s.D.Set("azure_tenant_id", *v.AzureTenantId)
		}

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.SasTokenSecretId != nil {
			s.D.Set("sas_token_secret_id", *v.SasTokenSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.AzureSynapseConnection:
		s.D.Set("connection_type", "AZURE_SYNAPSE_ANALYTICS")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.DatabricksConnection:
		s.D.Set("connection_type", "DATABRICKS")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.StorageCredentialName != nil {
			s.D.Set("storage_credential_name", *v.StorageCredentialName)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.DatabricksEngineConnection:
		s.D.Set("connection_type", "DATABRICKS_ENGINE")

		if v.ArtifactLocation != nil {
			artifactLocationArray := []interface{}{}
			if artifactLocationMap := DatabricksEngineArtifactLocationToMap(&v.ArtifactLocation); artifactLocationMap != nil {
				artifactLocationMap = preserveArtifactLocationWriteOnlyFields(s.D, artifactLocationMap)
				artifactLocationArray = append(artifactLocationArray, artifactLocationMap)
			}
			s.D.Set("artifact_location", artifactLocationArray)
		} else {
			s.D.Set("artifact_location", nil)
		}

		if v.AuthTokenSecretId != nil {
			s.D.Set("auth_token_secret_id", *v.AuthTokenSecretId)
		}

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.WorkspaceUrl != nil {
			s.D.Set("workspace_url", *v.WorkspaceUrl)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		} else {
			s.D.Set("security_attributes", map[string]interface{}{})
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.Db2Connection:
		s.D.Set("connection_type", "DB2")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslClientKeystashSecretId != nil {
			s.D.Set("ssl_client_keystash_secret_id", *v.SslClientKeystashSecretId)
		}

		if v.SslClientKeystoredbSecretId != nil {
			s.D.Set("ssl_client_keystoredb_secret_id", *v.SslClientKeystoredbSecretId)
		}

		if v.SslServerCertificate != nil {
			s.D.Set("ssl_server_certificate", *v.SslServerCertificate)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.ElasticsearchConnection:
		s.D.Set("connection_type", "ELASTICSEARCH")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.Fingerprint != nil {
			s.D.Set("fingerprint", *v.Fingerprint)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.Servers != nil {
			s.D.Set("servers", *v.Servers)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GenericConnection:
		s.D.Set("connection_type", "GENERIC")

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoldenGateConnection:
		s.D.Set("connection_type", "GOLDENGATE")

		if v.DeploymentId != nil {
			s.D.Set("deployment_id", *v.DeploymentId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoogleBigQueryConnection:
		s.D.Set("connection_type", "GOOGLE_BIGQUERY")

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GoogleCloudStorageConnection:
		s.D.Set("connection_type", "GOOGLE_CLOUD_STORAGE")

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.GooglePubSubConnection:
		s.D.Set("connection_type", "GOOGLE_PUBSUB")

		if v.ServiceAccountKeyFileSecretId != nil {
			s.D.Set("service_account_key_file_secret_id", *v.ServiceAccountKeyFileSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.HdfsConnection:
		s.D.Set("connection_type", "HDFS")

		if v.CoreSiteXml != nil {
			s.D.Set("core_site_xml", *v.CoreSiteXml)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.IcebergConnection:
		s.D.Set("connection_type", "ICEBERG")

		if v.Catalog != nil {
			catalogArray := []interface{}{}
			if catalogMap := IcebergCatalogToMap(&v.Catalog); catalogMap != nil {
				catalogMap = preserveIcebergCatalogWriteOnlyFields(s.D, catalogMap)
				catalogArray = append(catalogArray, catalogMap)
			}
			s.D.Set("catalog", catalogArray)
		} else {
			s.D.Set("catalog", nil)
		}

		if v.Storage != nil {
			storageArray := []interface{}{}
			if storageMap := IcebergStorageToMap(&v.Storage); storageMap != nil {
				storageMap = preserveIcebergStorageWriteOnlyFields(s.D, storageMap)
				storageArray = append(storageArray, storageMap)
			}
			s.D.Set("storage", storageArray)
		} else {
			s.D.Set("storage", nil)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.JavaMessageServiceConnection:
		s.D.Set("connection_type", "JAVA_MESSAGE_SERVICE")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ConnectionFactory != nil {
			s.D.Set("connection_factory", *v.ConnectionFactory)
		}

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.JndiConnectionFactory != nil {
			s.D.Set("jndi_connection_factory", *v.JndiConnectionFactory)
		}

		if v.JndiInitialContextFactory != nil {
			s.D.Set("jndi_initial_context_factory", *v.JndiInitialContextFactory)
		}

		if v.JndiProviderUrl != nil {
			s.D.Set("jndi_provider_url", *v.JndiProviderUrl)
		}

		if v.JndiSecurityCredentialsSecretId != nil {
			s.D.Set("jndi_security_credentials_secret_id", *v.JndiSecurityCredentialsSecretId)
		}

		if v.JndiSecurityPrincipal != nil {
			s.D.Set("jndi_security_principal", *v.JndiSecurityPrincipal)
		}

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldUseJndi != nil {
			s.D.Set("should_use_jndi", *v.ShouldUseJndi)
		}

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.KafkaConnection:
		s.D.Set("connection_type", "KAFKA")

		bootstrapServers := []interface{}{}
		for _, item := range v.BootstrapServers {
			bootstrapServers = append(bootstrapServers, KafkaBootstrapServerToMap(item))
		}
		s.D.Set("bootstrap_servers", bootstrapServers)

		if v.ClusterId != nil {
			s.D.Set("cluster_id", *v.ClusterId)
		}

		if v.ConsumerProperties != nil {
			s.D.Set("consumer_properties", *v.ConsumerProperties)
		}

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.ProducerProperties != nil {
			s.D.Set("producer_properties", *v.ProducerProperties)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
		}

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		if v.StreamPoolId != nil {
			s.D.Set("stream_pool_id", *v.StreamPoolId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.KafkaSchemaRegistryConnection:
		s.D.Set("connection_type", "KAFKA_SCHEMA_REGISTRY")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

		if v.Url != nil {
			s.D.Set("url", *v.Url)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MicrosoftFabricConnection:
		s.D.Set("connection_type", "MICROSOFT_FABRIC")

		if v.ClientId != nil {
			s.D.Set("client_id", *v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			s.D.Set("client_secret_secret_id", *v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			s.D.Set("endpoint", *v.Endpoint)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenantId != nil {
			s.D.Set("tenant_id", *v.TenantId)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MicrosoftSqlserverConnection:
		s.D.Set("connection_type", "MICROSOFT_SQLSERVER")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.ShouldValidateServerCertificate != nil {
			s.D.Set("should_validate_server_certificate", *v.ShouldValidateServerCertificate)
		}

		if v.SslCa != nil {
			s.D.Set("ssl_ca", *v.SslCa)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MongoDbConnection:
		s.D.Set("connection_type", "MONGODB")

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		s.D.Set("technology_type", v.TechnologyType)

		if v.TlsCaFile != nil {
			s.D.Set("tls_ca_file", *v.TlsCaFile)
		}

		if v.TlsCertificateKeyFilePasswordSecretId != nil {
			s.D.Set("tls_certificate_key_file_password_secret_id", *v.TlsCertificateKeyFilePasswordSecretId)
		}

		if v.TlsCertificateKeyFileSecretId != nil {
			s.D.Set("tls_certificate_key_file_secret_id", *v.TlsCertificateKeyFileSecretId)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.MysqlConnection:
		s.D.Set("connection_type", "MYSQL")

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DbSystemId != nil {
			s.D.Set("db_system_id", *v.DbSystemId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslCa != nil {
			s.D.Set("ssl_ca", *v.SslCa)
		}

		if v.SslCert != nil {
			s.D.Set("ssl_cert", *v.SslCert)
		}

		if v.SslCrl != nil {
			s.D.Set("ssl_crl", *v.SslCrl)
		}

		if v.SslKeySecretId != nil {
			s.D.Set("ssl_key_secret_id", *v.SslKeySecretId)
		}

		s.D.Set("ssl_mode", v.SslMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OciObjectStorageConnection:
		s.D.Set("connection_type", "OCI_OBJECT_STORAGE")

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			s.D.Set("public_key_fingerprint", *v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenancyId != nil {
			s.D.Set("tenancy_id", *v.TenancyId)
		}

		if v.UserId != nil {
			s.D.Set("user_id", *v.UserId)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleConnection:
		s.D.Set("connection_type", "ORACLE")

		s.D.Set("authentication_mode", v.AuthenticationMode)

		if v.ConnectionString != nil {
			s.D.Set("connection_string", *v.ConnectionString)
		}

		if v.DatabaseId != nil {
			s.D.Set("database_id", *v.DatabaseId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("session_mode", v.SessionMode)

		s.D.Set("technology_type", v.TechnologyType)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.WalletSecretId != nil {
			s.D.Set("wallet_secret_id", *v.WalletSecretId)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleAiDataPlatformConnection:
		s.D.Set("connection_type", "ORACLE_AI_DATA_PLATFORM")

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			s.D.Set("public_key_fingerprint", *v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenancyId != nil {
			s.D.Set("tenancy_id", *v.TenancyId)
		}

		if v.UserId != nil {
			s.D.Set("user_id", *v.UserId)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleAiDataPlatformEngineConnection:
		s.D.Set("connection_type", "ORACLE_AI_DATA_PLATFORM_ENGINE")

		s.D.Set("technology_type", v.TechnologyType)

		if v.WorkspaceUrl != nil {
			s.D.Set("workspace_url", *v.WorkspaceUrl)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		} else {
			s.D.Set("security_attributes", map[string]interface{}{})
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.OracleNosqlConnection:
		s.D.Set("connection_type", "ORACLE_NOSQL")

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			s.D.Set("public_key_fingerprint", *v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			s.D.Set("region", *v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			s.D.Set("should_use_resource_principal", *v.ShouldUseResourcePrincipal)
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TenancyId != nil {
			s.D.Set("tenancy_id", *v.TenancyId)
		}

		if v.UserId != nil {
			s.D.Set("user_id", *v.UserId)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.PostgresqlConnection:
		s.D.Set("connection_type", "POSTGRESQL")
		s.D.Set("technology_type", v.TechnologyType)

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		s.D.Set("additional_attributes", additionalAttributes)

		if v.DatabaseName != nil {
			s.D.Set("database_name", *v.DatabaseName)
		}

		if v.DbSystemId != nil {
			s.D.Set("db_system_id", *v.DbSystemId)
		}

		if v.Host != nil {
			s.D.Set("host", *v.Host)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.Port != nil {
			s.D.Set("port", *v.Port)
		}

		if v.PrivateIp != nil {
			s.D.Set("private_ip", *v.PrivateIp)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslCa != nil {
			s.D.Set("ssl_ca", *v.SslCa)
		}

		if v.SslCert != nil {
			s.D.Set("ssl_cert", *v.SslCert)
		}

		if v.SslCrl != nil {
			s.D.Set("ssl_crl", *v.SslCrl)
		}

		if v.SslKeySecretId != nil {
			s.D.Set("ssl_key_secret_id", *v.SslKeySecretId)
		}

		s.D.Set("ssl_mode", v.SslMode)

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.RedisConnection:
		s.D.Set("connection_type", "REDIS")
		s.D.Set("technology_type", v.TechnologyType)

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.RedisClusterId != nil {
			s.D.Set("redis_cluster_id", *v.RedisClusterId)
		}

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.Servers != nil {
			s.D.Set("servers", *v.Servers)
		}

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.SnowflakeConnection:
		s.D.Set("connection_type", "SNOWFLAKE")

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ConnectionUrl != nil {
			s.D.Set("connection_url", *v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			s.D.Set("password_secret_id", *v.PasswordSecretId)
		}

		if v.PrivateKeyFileSecretId != nil {
			s.D.Set("private_key_file_secret_id", *v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			s.D.Set("private_key_passphrase_secret_id", *v.PrivateKeyPassphraseSecretId)
		}

		if v.Username != nil {
			s.D.Set("username", *v.Username)
		}

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", nsgIds)

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		s.D.Set("technology_type", v.TechnologyType)

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	case oci_golden_gate.SparkConnection:
		s.D.Set("connection_type", "SPARK")

		if v.ArtifactLocation != nil {
			artifactLocationArray := []interface{}{}
			if artifactLocationMap := SparkArtifactLocationToMap(&v.ArtifactLocation); artifactLocationMap != nil {
				artifactLocationMap = preserveArtifactLocationWriteOnlyFields(s.D, artifactLocationMap)
				artifactLocationArray = append(artifactLocationArray, artifactLocationMap)
			}
			s.D.Set("artifact_location", artifactLocationArray)
		} else {
			s.D.Set("artifact_location", nil)
		}

		if v.AuthSecretSecretId != nil {
			s.D.Set("auth_secret_secret_id", *v.AuthSecretSecretId)
		}

		s.D.Set("authentication_type", v.AuthenticationType)

		if v.ClusterEndpoint != nil {
			s.D.Set("cluster_endpoint", *v.ClusterEndpoint)
		}

		if v.KeyStorePasswordSecretId != nil {
			s.D.Set("key_store_password_secret_id", *v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			s.D.Set("key_store_secret_id", *v.KeyStoreSecretId)
		}

		s.D.Set("key_store_type", v.KeyStoreType)

		if v.RestAuthSecretKeySecretId != nil {
			s.D.Set("rest_auth_secret_key_secret_id", *v.RestAuthSecretKeySecretId)
		}

		if v.RestEndpoint != nil {
			s.D.Set("rest_endpoint", *v.RestEndpoint)
		}

		s.D.Set("rpc_cipher", v.RpcCipher)

		s.D.Set("security_protocol", v.SecurityProtocol)

		if v.SslKeyPasswordSecretId != nil {
			s.D.Set("ssl_key_password_secret_id", *v.SslKeyPasswordSecretId)
		}

		s.D.Set("submit_mode", v.SubmitMode)

		s.D.Set("technology_type", v.TechnologyType)

		s.D.Set("tls_protocol", v.TlsProtocol)

		if v.TrustStorePasswordSecretId != nil {
			s.D.Set("trust_store_password_secret_id", *v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			s.D.Set("trust_store_secret_id", *v.TrustStoreSecretId)
		}

		s.D.Set("trust_store_type", v.TrustStoreType)

		if v.ClusterPlacementGroupId != nil {
			s.D.Set("cluster_placement_group_id", *v.ClusterPlacementGroupId)
		}

		if v.CompartmentId != nil {
			s.D.Set("compartment_id", *v.CompartmentId)
		}

		if v.DefinedTags != nil {
			s.D.Set("defined_tags", tfresource.DefinedTagsToMap(v.DefinedTags))
		}

		if v.Description != nil {
			s.D.Set("description", *v.Description)
		}

		if v.DisplayName != nil {
			s.D.Set("display_name", *v.DisplayName)
		}

		if v.DoesUseSecretIds != nil {
			s.D.Set("does_use_secret_ids", *v.DoesUseSecretIds)
		}

		s.D.Set("freeform_tags", v.FreeformTags)

		if v.Id != nil {
			s.D.SetId(*v.Id)
		}

		ingressIps := []interface{}{}
		for _, item := range v.IngressIps {
			ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
		}
		s.D.Set("ingress_ips", ingressIps)

		if v.KeyId != nil {
			s.D.Set("key_id", *v.KeyId)
		}

		if v.LifecycleDetails != nil {
			s.D.Set("lifecycle_details", *v.LifecycleDetails)
		}

		locks := []interface{}{}
		for _, item := range v.Locks {
			locks = append(locks, ResourceLockToMap(item))
		}
		s.D.Set("locks", locks)

		nsgIds := []interface{}{}
		for _, item := range v.NsgIds {
			nsgIds = append(nsgIds, item)
		}
		s.D.Set("nsg_ids", schema.NewSet(tfresource.LiteralTypeHashCodeForSets, nsgIds))

		s.D.Set("routing_method", v.RoutingMethod)

		if v.SecurityAttributes != nil {
			s.D.Set("security_attributes", tfresource.SecurityAttributesToMap(v.SecurityAttributes))
		} else {
			s.D.Set("security_attributes", map[string]interface{}{})
		}

		s.D.Set("state", v.LifecycleState)

		if v.SubnetId != nil {
			s.D.Set("subnet_id", *v.SubnetId)
		}

		if v.SubscriptionId != nil {
			s.D.Set("subscription_id", *v.SubscriptionId)
		}

		if v.SystemTags != nil {
			s.D.Set("system_tags", tfresource.SystemTagsToMap(v.SystemTags))
		}

		if v.TimeCreated != nil {
			s.D.Set("time_created", v.TimeCreated.String())
		}

		if v.TimeUpdated != nil {
			s.D.Set("time_updated", v.TimeUpdated.String())
		}

		if v.VaultId != nil {
			s.D.Set("vault_id", *v.VaultId)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", *s.Res)
		return nil
	}
	return nil
}

func preserveAiModelAuthDetailsWriteOnlyFields(d *schema.ResourceData, authDetailsMap map[string]interface{}) map[string]interface{} {
	existingAuthDetails, ok := d.GetOkExists("auth_details")
	if !ok {
		return authDetailsMap
	}

	existingList, ok := existingAuthDetails.([]interface{})
	if !ok || len(existingList) == 0 || existingList[0] == nil {
		return authDetailsMap
	}

	existingMap, ok := existingList[0].(map[string]interface{})
	if !ok {
		return authDetailsMap
	}

	for _, key := range []string{"api_key", "api_key_secret_id"} {
		if _, exists := authDetailsMap[key]; exists {
			continue
		}
		if value, ok := existingMap[key]; ok && value != nil && value != "" {
			authDetailsMap[key] = value
		}
	}

	return authDetailsMap
}

func preserveIcebergCatalogWriteOnlyFields(d *schema.ResourceData, catalogMap map[string]interface{}) map[string]interface{} {
	existingCatalog, ok := d.GetOkExists("catalog")
	if !ok {
		return catalogMap
	}

	existingList, ok := existingCatalog.([]interface{})
	if !ok || len(existingList) == 0 || existingList[0] == nil {
		return catalogMap
	}

	existingMap, ok := existingList[0].(map[string]interface{})
	if !ok {
		return catalogMap
	}

	for _, key := range []string{"client_secret", "properties"} {
		if _, exists := catalogMap[key]; exists {
			continue
		}
		if value, ok := existingMap[key]; ok && value != nil && value != "" {
			catalogMap[key] = value
		}
	}

	return catalogMap
}

func preserveIcebergStorageWriteOnlyFields(d *schema.ResourceData, storageMap map[string]interface{}) map[string]interface{} {
	existingStorage, ok := d.GetOkExists("storage")
	if !ok {
		return storageMap
	}

	existingList, ok := existingStorage.([]interface{})
	if !ok || len(existingList) == 0 || existingList[0] == nil {
		return storageMap
	}

	existingMap, ok := existingList[0].(map[string]interface{})
	if !ok {
		return storageMap
	}

	for _, key := range []string{"secret_access_key", "account_key", "service_account_key_file"} {
		if _, exists := storageMap[key]; exists {
			continue
		}
		if value, ok := existingMap[key]; ok && value != nil && value != "" {
			storageMap[key] = value
		}
	}

	return storageMap
}

func preserveArtifactLocationWriteOnlyFields(d *schema.ResourceData, artifactLocationMap map[string]interface{}) map[string]interface{} {
	existingArtifactLocation, ok := d.GetOkExists("artifact_location")
	if !ok {
		return artifactLocationMap
	}

	existingList, ok := existingArtifactLocation.([]interface{})
	if !ok || len(existingList) == 0 || existingList[0] == nil {
		return artifactLocationMap
	}

	existingMap, ok := existingList[0].(map[string]interface{})
	if !ok || existingMap["location_type"] != artifactLocationMap["location_type"] {
		return artifactLocationMap
	}

	preserveNonEmptyMapValues(artifactLocationMap, existingMap, "account_key")

	configuration, ok := firstNestedMap(artifactLocationMap, "configuration")
	if !ok {
		return artifactLocationMap
	}
	existingConfiguration, ok := firstNestedMap(existingMap, "configuration")
	if !ok || existingConfiguration["authentication_type"] != configuration["authentication_type"] {
		return artifactLocationMap
	}

	preserveNonEmptyMapValues(configuration, existingConfiguration,
		"par_url", "private_key_file", "private_key_passphrase", "secret_access_key")
	return artifactLocationMap
}

func firstNestedMap(parent map[string]interface{}, key string) (map[string]interface{}, bool) {
	nested, ok := parent[key].([]interface{})
	if !ok || len(nested) == 0 || nested[0] == nil {
		return nil, false
	}
	value, ok := nested[0].(map[string]interface{})
	return value, ok
}

func preserveNonEmptyMapValues(target, source map[string]interface{}, keys ...string) {
	for _, key := range keys {
		if _, exists := target[key]; exists {
			continue
		}
		if value, ok := source[key]; ok && value != nil && value != "" {
			target[key] = value
		}
	}
}

func (s *GoldenGateConnectionResourceCrud) mapToAddResourceLockDetails(fieldKeyFormat string) (oci_golden_gate.AddResourceLockDetails, error) {
	result := oci_golden_gate.AddResourceLockDetails{}

	if message, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "message")); ok {
		tmp := message.(string)
		result.Message = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_golden_gate.AddResourceLockDetailsTypeEnum(type_.(string))
	}

	return result, nil
}

func (s *GoldenGateConnectionResourceCrud) mapToAiModelAuthDetails(fieldKeyFormat string) (oci_golden_gate.AiModelAuthDetails, error) {
	var baseObject oci_golden_gate.AiModelAuthDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_type"))
	var authType string
	if ok {
		authType = authTypeRaw.(string)
	} else {
		authType = "" // default value
	}
	switch strings.ToLower(authType) {
	case strings.ToLower("API_KEY"):
		details := oci_golden_gate.ApiKeyAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if baseUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "base_url")); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_GEN_AI"):
		details := oci_golden_gate.OciGenAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if keyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "key_fingerprint")); ok {
			tmp := keyFingerprint.(string)
			details.KeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown auth_type '%v' was specified", authType)
	}
	return baseObject, nil
}

func AiModelAuthDetailsToMap(obj *oci_golden_gate.AiModelAuthDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.ApiKeyAiModelAuthDetails:
		result["auth_type"] = "API_KEY"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}
	case oci_golden_gate.OciGenAiModelAuthDetails:
		result["auth_type"] = "OCI_GEN_AI"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.KeyFingerprint != nil {
			result["key_fingerprint"] = string(*v.KeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'auth_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToAiModelAuthDetailsSummary(fieldKeyFormat string) (oci_golden_gate.AiModelAuthDetailsSummary, error) {
	var baseObject oci_golden_gate.AiModelAuthDetailsSummary
	//discriminator
	authTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_type"))
	var authType string
	if ok {
		authType = authTypeRaw.(string)
	} else {
		authType = "" // default value
	}
	switch strings.ToLower(authType) {
	case strings.ToLower("API_KEY"):
		details := oci_golden_gate.ApiKeyAiModelAuthDetailsSummary{}
		if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
			tmp := apiKeySecretId.(string)
			details.ApiKeySecretId = &tmp
		}
		if baseUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "base_url")); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_GEN_AI"):
		details := oci_golden_gate.OciGenAiModelAuthDetailsSummary{}
		if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
			tmp := apiKeySecretId.(string)
			details.ApiKeySecretId = &tmp
		}
		if keyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "key_fingerprint")); ok {
			tmp := keyFingerprint.(string)
			details.KeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown auth_type '%v' was specified", authType)
	}
	return baseObject, nil
}

func AiModelAuthDetailsSummaryToMap(obj *oci_golden_gate.AiModelAuthDetailsSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.ApiKeyAiModelAuthDetailsSummary:
		result["auth_type"] = "API_KEY"

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}
	case oci_golden_gate.OciGenAiModelAuthDetailsSummary:
		result["auth_type"] = "OCI_GEN_AI"

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.KeyFingerprint != nil {
			result["key_fingerprint"] = string(*v.KeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'auth_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func ConnectionSummaryToMap(obj oci_golden_gate.ConnectionSummary, datasource bool) map[string]interface{} {
	result := map[string]interface{}{}

	//set common fields
	result["id"] = *obj.GetId()

	if obj.GetDisplayName() != nil {
		result["display_name"] = *obj.GetDisplayName()
	}
	result["compartment_id"] = *obj.GetCompartmentId()
	result["state"] = string(obj.GetLifecycleState())

	if obj.GetTimeCreated() != nil {
		result["time_created"] = obj.GetTimeCreated().String()
	}
	if obj.GetTimeUpdated() != nil {
		result["time_updated"] = obj.GetTimeUpdated().String()
	}
	if obj.GetDescription() != nil {
		result["description"] = *obj.GetDescription()
	}
	result["freeform_tags"] = obj.GetFreeformTags()
	if obj.GetDefinedTags() != nil {
		result["defined_tags"] = tfresource.DefinedTagsToMap(obj.GetDefinedTags())
	}
	if obj.GetSystemTags() != nil {
		result["system_tags"] = tfresource.SystemTagsToMap(obj.GetSystemTags())
	}
	if obj.GetLifecycleDetails() != nil {
		result["lifecycle_details"] = *obj.GetLifecycleDetails()
	}

	if obj.GetVaultId() != nil {
		result["vault_id"] = *obj.GetVaultId()
	}

	if obj.GetSecurityAttributes() != nil {
		result["security_attributes"] = tfresource.SecurityAttributesToMap(obj.GetSecurityAttributes())
	}

	if obj.GetSubnetId() != nil {
		result["subnet_id"] = *obj.GetSubnetId()
	}

	ingressIps := []interface{}{}
	for _, item := range obj.GetIngressIps() {
		ingressIps = append(ingressIps, IngressIpDetailsToMap(item))
	}
	result["ingress_ips"] = ingressIps

	if obj.GetNsgIds() != nil {
		result["nsg_ids"] = obj.GetNsgIds()
	}
	result["routing_method"] = string(obj.GetRoutingMethod())

	// set type specific fields
	switch v := (obj).(type) {
	case oci_golden_gate.AiModelConnectionSummary:
		result["connection_type"] = "AI_MODEL"

		if v.AuthDetails != nil {
			authDetailsArray := []interface{}{}
			if authDetailsMap := AiModelAuthDetailsSummaryToMap(&v.AuthDetails); authDetailsMap != nil {
				authDetailsArray = append(authDetailsArray, authDetailsMap)
			}
			result["auth_details"] = authDetailsArray
		}

		if v.MaxInputChars != nil {
			result["max_input_chars"] = int(*v.MaxInputChars)
		}

		if v.ModelKey != nil {
			result["model_key"] = string(*v.ModelKey)
		}

		if v.ProviderId != nil {
			result["provider_id"] = string(*v.ProviderId)
		}

		result["provider_type"] = string(v.ProviderType)

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AiProviderConnectionSummary:
		result["connection_type"] = "AI_PROVIDER"

		if v.AuthDetails != nil {
			authDetailsArray := []interface{}{}
			if authDetailsMap := AiModelAuthDetailsSummaryToMap(&v.AuthDetails); authDetailsMap != nil {
				authDetailsArray = append(authDetailsArray, authDetailsMap)
			}
			result["auth_details"] = authDetailsArray
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AmazonKinesisConnectionSummary:
		result["connection_type"] = "AMAZON_KINESIS"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AmazonRedshiftConnectionSummary:
		result["connection_type"] = "AMAZON_REDSHIFT"

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.AmazonS3ConnectionSummary:
		result["connection_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AzureDataLakeStorageConnectionSummary:
		result["connection_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		result["authentication_type"] = string(v.AuthenticationType)

		if v.AzureAuthorityHost != nil {
			result["azure_authority_host"] = string(*v.AzureAuthorityHost)
		}

		if v.AzureTenantId != nil {
			result["azure_tenant_id"] = string(*v.AzureTenantId)
		}

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.SasTokenSecretId != nil {
			result["sas_token_secret_id"] = string(*v.SasTokenSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.AzureSynapseConnectionSummary:
		result["connection_type"] = "AZURE_SYNAPSE_ANALYTICS"

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.DatabricksConnectionSummary:
		result["connection_type"] = "DATABRICKS"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.StorageCredentialName != nil {
			result["storage_credential_name"] = string(*v.StorageCredentialName)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.DatabricksEngineConnectionSummary:
		result["connection_type"] = "DATABRICKS_ENGINE"

		if v.ArtifactLocation != nil {
			artifactLocationArray := []interface{}{}
			if artifactLocationMap := DatabricksEngineArtifactLocationSummaryToMap(&v.ArtifactLocation); artifactLocationMap != nil {
				artifactLocationArray = append(artifactLocationArray, artifactLocationMap)
			}
			result["artifact_location"] = artifactLocationArray
		}

		if v.AuthTokenSecretId != nil {
			result["auth_token_secret_id"] = string(*v.AuthTokenSecretId)
		}

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.WorkspaceUrl != nil {
			result["workspace_url"] = string(*v.WorkspaceUrl)
		}
	case oci_golden_gate.Db2ConnectionSummary:
		result["connection_type"] = "DB2"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.SslClientKeystashSecretId != nil {
			result["ssl_client_keystash_secret_id"] = string(*v.SslClientKeystashSecretId)
		}

		if v.SslClientKeystoredbSecretId != nil {
			result["ssl_client_keystoredb_secret_id"] = string(*v.SslClientKeystoredbSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.ElasticsearchConnectionSummary:
		result["connection_type"] = "ELASTICSEARCH"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.Servers != nil {
			result["servers"] = string(*v.Servers)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.GenericConnectionSummary:
		result["connection_type"] = "GENERIC"

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.GoldenGateConnectionSummary:
		result["connection_type"] = "GOLDENGATE"

		if v.DeploymentId != nil {
			result["deployment_id"] = string(*v.DeploymentId)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.GoogleBigQueryConnectionSummary:
		result["connection_type"] = "GOOGLE_BIGQUERY"

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.GoogleCloudStorageConnectionSummary:
		result["connection_type"] = "GOOGLE_CLOUD_STORAGE"

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.GooglePubSubConnectionSummary:
		result["connection_type"] = "GOOGLE_PUBSUB"

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.HdfsConnectionSummary:
		result["connection_type"] = "HDFS"

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.IcebergConnectionSummary:
		result["connection_type"] = "ICEBERG"

		if v.Catalog != nil {
			catalogArray := []interface{}{}
			if catalogMap := IcebergCatalogSummaryToMap(&v.Catalog); catalogMap != nil {
				catalogArray = append(catalogArray, catalogMap)
			}
			result["catalog"] = catalogArray
		}

		if v.Storage != nil {
			storageArray := []interface{}{}
			if storageMap := IcebergStorageSummaryToMap(&v.Storage); storageMap != nil {
				storageArray = append(storageArray, storageMap)
			}
			result["storage"] = storageArray
		}

		result["technology_type"] = string(v.TechnologyType)
	case oci_golden_gate.JavaMessageServiceConnectionSummary:
		result["connection_type"] = "JAVA_MESSAGE_SERVICE"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ConnectionFactory != nil {
			result["connection_factory"] = string(*v.ConnectionFactory)
		}

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.JndiConnectionFactory != nil {
			result["jndi_connection_factory"] = string(*v.JndiConnectionFactory)
		}

		if v.JndiInitialContextFactory != nil {
			result["jndi_initial_context_factory"] = string(*v.JndiInitialContextFactory)
		}

		if v.JndiProviderUrl != nil {
			result["jndi_provider_url"] = string(*v.JndiProviderUrl)
		}

		if v.JndiSecurityCredentialsSecretId != nil {
			result["jndi_security_credentials_secret_id"] = string(*v.JndiSecurityCredentialsSecretId)
		}

		if v.JndiSecurityPrincipal != nil {
			result["jndi_security_principal"] = string(*v.JndiSecurityPrincipal)
		}

		if v.KeyStorePasswordSecretId != nil {
			result["key_store_password_secret_id"] = string(*v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			result["key_store_secret_id"] = string(*v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.ShouldUseJndi != nil {
			result["should_use_jndi"] = bool(*v.ShouldUseJndi)
		}

		if v.SslKeyPasswordSecretId != nil {
			result["ssl_key_password_secret_id"] = string(*v.SslKeyPasswordSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			result["trust_store_password_secret_id"] = string(*v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			result["trust_store_secret_id"] = string(*v.TrustStoreSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.KafkaConnectionSummary:
		result["connection_type"] = "KAFKA"

		bootstrapServers := []interface{}{}
		for _, item := range v.BootstrapServers {
			bootstrapServers = append(bootstrapServers, KafkaBootstrapServerToMap(item))
		}
		result["bootstrap_servers"] = bootstrapServers

		if v.ClusterId != nil {
			result["cluster_id"] = string(*v.ClusterId)
		}

		if v.KeyStorePasswordSecretId != nil {
			result["key_store_password_secret_id"] = string(*v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			result["key_store_secret_id"] = string(*v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.ShouldUseResourcePrincipal != nil {
			result["should_use_resource_principal"] = bool(*v.ShouldUseResourcePrincipal)
		}

		if v.SslKeyPasswordSecretId != nil {
			result["ssl_key_password_secret_id"] = string(*v.SslKeyPasswordSecretId)
		}

		if v.StreamPoolId != nil {
			result["stream_pool_id"] = string(*v.StreamPoolId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			result["trust_store_password_secret_id"] = string(*v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			result["trust_store_secret_id"] = string(*v.TrustStoreSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.KafkaSchemaRegistryConnectionSummary:
		result["connection_type"] = "KAFKA_SCHEMA_REGISTRY"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.KeyStorePasswordSecretId != nil {
			result["key_store_password_secret_id"] = string(*v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			result["key_store_secret_id"] = string(*v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		if v.SslKeyPasswordSecretId != nil {
			result["ssl_key_password_secret_id"] = string(*v.SslKeyPasswordSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			result["trust_store_password_secret_id"] = string(*v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			result["trust_store_secret_id"] = string(*v.TrustStoreSecretId)
		}

		if v.Url != nil {
			result["url"] = string(*v.Url)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MicrosoftFabricConnectionSummary:
		result["connection_type"] = "MICROSOFT_FABRIC"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenantId != nil {
			result["tenant_id"] = string(*v.TenantId)
		}
	case oci_golden_gate.MicrosoftSqlserverConnectionSummary:
		result["connection_type"] = "MICROSOFT_SQLSERVER"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.ShouldValidateServerCertificate != nil {
			result["should_validate_server_certificate"] = bool(*v.ShouldValidateServerCertificate)
		}

		if v.SslCa != nil {
			result["ssl_ca"] = string(*v.SslCa)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MongoDbConnectionSummary:
		result["connection_type"] = "MONGODB"

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		result["technology_type"] = string(v.TechnologyType)

		if v.TlsCertificateKeyFilePasswordSecretId != nil {
			result["tls_certificate_key_file_password_secret_id"] = string(*v.TlsCertificateKeyFilePasswordSecretId)
		}

		if v.TlsCertificateKeyFileSecretId != nil {
			result["tls_certificate_key_file_secret_id"] = string(*v.TlsCertificateKeyFileSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.MysqlConnectionSummary:
		result["connection_type"] = "MYSQL"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.SslKeySecretId != nil {
			result["ssl_key_secret_id"] = string(*v.SslKeySecretId)
		}

		result["ssl_mode"] = string(v.SslMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.OciObjectStorageConnectionSummary:
		result["connection_type"] = "OCI_OBJECT_STORAGE"

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			result["should_use_resource_principal"] = bool(*v.ShouldUseResourcePrincipal)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	case oci_golden_gate.OracleConnectionSummary:
		result["connection_type"] = "ORACLE"

		result["authentication_mode"] = string(v.AuthenticationMode)

		if v.ConnectionString != nil {
			result["connection_string"] = string(*v.ConnectionString)
		}

		if v.DatabaseId != nil {
			result["database_id"] = string(*v.DatabaseId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["session_mode"] = string(v.SessionMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}

		if v.WalletSecretId != nil {
			result["wallet_secret_id"] = string(*v.WalletSecretId)
		}
	case oci_golden_gate.OracleAiDataPlatformConnectionSummary:
		result["connection_type"] = "ORACLE_AI_DATA_PLATFORM"

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			result["should_use_resource_principal"] = bool(*v.ShouldUseResourcePrincipal)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	case oci_golden_gate.OracleAiDataPlatformEngineConnectionSummary:
		result["connection_type"] = "ORACLE_AI_DATA_PLATFORM_ENGINE"

		result["technology_type"] = string(v.TechnologyType)

		if v.WorkspaceUrl != nil {
			result["workspace_url"] = string(*v.WorkspaceUrl)
		}
	case oci_golden_gate.OracleNosqlConnectionSummary:
		result["connection_type"] = "ORACLE_NOSQL"

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.ShouldUseResourcePrincipal != nil {
			result["should_use_resource_principal"] = bool(*v.ShouldUseResourcePrincipal)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	case oci_golden_gate.PostgresqlConnectionSummary:
		result["connection_type"] = "POSTGRESQL"

		additionalAttributes := []interface{}{}
		for _, item := range v.AdditionalAttributes {
			additionalAttributes = append(additionalAttributes, NameValuePairToMap(item))
		}
		result["additional_attributes"] = additionalAttributes

		if v.DatabaseName != nil {
			result["database_name"] = string(*v.DatabaseName)
		}

		if v.DbSystemId != nil {
			result["db_system_id"] = string(*v.DbSystemId)
		}

		if v.Host != nil {
			result["host"] = string(*v.Host)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.Port != nil {
			result["port"] = int(*v.Port)
		}

		if v.PrivateIp != nil {
			result["private_ip"] = string(*v.PrivateIp)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.SslKeySecretId != nil {
			result["ssl_key_secret_id"] = string(*v.SslKeySecretId)
		}

		result["ssl_mode"] = string(v.SslMode)

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.RedisConnectionSummary:
		result["connection_type"] = "REDIS"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.KeyStorePasswordSecretId != nil {
			result["key_store_password_secret_id"] = string(*v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			result["key_store_secret_id"] = string(*v.KeyStoreSecretId)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.RedisClusterId != nil {
			result["redis_cluster_id"] = string(*v.RedisClusterId)
		}

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.Servers != nil {
			result["servers"] = string(*v.Servers)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.TrustStorePasswordSecretId != nil {
			result["trust_store_password_secret_id"] = string(*v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			result["trust_store_secret_id"] = string(*v.TrustStoreSecretId)
		}

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	case oci_golden_gate.SparkConnectionSummary:
		result["connection_type"] = "SPARK"

		if v.ArtifactLocation != nil {
			artifactLocationArray := []interface{}{}
			if artifactLocationMap := SparkArtifactLocationSummaryToMap(&v.ArtifactLocation); artifactLocationMap != nil {
				artifactLocationArray = append(artifactLocationArray, artifactLocationMap)
			}
			result["artifact_location"] = artifactLocationArray
		}

		if v.AuthSecretSecretId != nil {
			result["auth_secret_secret_id"] = string(*v.AuthSecretSecretId)
		}

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ClusterEndpoint != nil {
			result["cluster_endpoint"] = string(*v.ClusterEndpoint)
		}

		if v.KeyStorePasswordSecretId != nil {
			result["key_store_password_secret_id"] = string(*v.KeyStorePasswordSecretId)
		}

		if v.KeyStoreSecretId != nil {
			result["key_store_secret_id"] = string(*v.KeyStoreSecretId)
		}

		result["key_store_type"] = string(v.KeyStoreType)

		if v.RestAuthSecretKeySecretId != nil {
			result["rest_auth_secret_key_secret_id"] = string(*v.RestAuthSecretKeySecretId)
		}

		if v.RestEndpoint != nil {
			result["rest_endpoint"] = string(*v.RestEndpoint)
		}

		result["rpc_cipher"] = string(v.RpcCipher)

		result["security_protocol"] = string(v.SecurityProtocol)

		if v.SslKeyPasswordSecretId != nil {
			result["ssl_key_password_secret_id"] = string(*v.SslKeyPasswordSecretId)
		}

		result["submit_mode"] = string(v.SubmitMode)

		result["technology_type"] = string(v.TechnologyType)

		result["tls_protocol"] = string(v.TlsProtocol)

		if v.TrustStorePasswordSecretId != nil {
			result["trust_store_password_secret_id"] = string(*v.TrustStorePasswordSecretId)
		}

		if v.TrustStoreSecretId != nil {
			result["trust_store_secret_id"] = string(*v.TrustStoreSecretId)
		}

		result["trust_store_type"] = string(v.TrustStoreType)
	case oci_golden_gate.SnowflakeConnectionSummary:
		result["connection_type"] = "SNOWFLAKE"

		result["authentication_type"] = string(v.AuthenticationType)

		if v.ConnectionUrl != nil {
			result["connection_url"] = string(*v.ConnectionUrl)
		}

		if v.PasswordSecretId != nil {
			result["password_secret_id"] = string(*v.PasswordSecretId)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		result["technology_type"] = string(v.TechnologyType)

		if v.Username != nil {
			result["username"] = string(*v.Username)
		}
	default:
		log.Printf("[WARN] Received 'connection_type' of unknown type %v", obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateAiModelAuthDetails(fieldKeyFormat string) (oci_golden_gate.CreateAiModelAuthDetails, error) {
	var baseObject oci_golden_gate.CreateAiModelAuthDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_type"))
	var authType string
	if ok {
		authType = authTypeRaw.(string)
	} else {
		authType = "" // default value
	}
	switch strings.ToLower(authType) {
	case strings.ToLower("API_KEY"):
		details := oci_golden_gate.CreateApiKeyAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if baseUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "base_url")); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_GEN_AI"):
		details := oci_golden_gate.CreateOciGenAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if keyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "key_fingerprint")); ok {
			tmp := keyFingerprint.(string)
			details.KeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown auth_type '%v' was specified", authType)
	}
	return baseObject, nil
}

func CreateAiModelAuthDetailsToMap(obj *oci_golden_gate.CreateAiModelAuthDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateApiKeyAiModelAuthDetails:
		result["auth_type"] = "API_KEY"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}
	case oci_golden_gate.CreateOciGenAiModelAuthDetails:
		result["auth_type"] = "OCI_GEN_AI"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.KeyFingerprint != nil {
			result["key_fingerprint"] = string(*v.KeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'auth_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateIcebergCatalogDetails(fieldKeyFormat string) (oci_golden_gate.CreateIcebergCatalogDetails, error) {
	var baseObject oci_golden_gate.CreateIcebergCatalogDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	catalogTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_type"))
	var catalogType string
	if ok {
		catalogType = catalogTypeRaw.(string)
	} else {
		catalogType = "" // default value
	}
	switch strings.ToLower(catalogType) {
	case strings.ToLower("GLUE"):
		details := oci_golden_gate.CreateGlueIcebergCatalogDetails{}
		if glueId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "glue_id")); ok {
			tmp := glueId.(string)
			details.GlueId = &tmp
		}
		baseObject = details
	case strings.ToLower("HADOOP"):
		details := oci_golden_gate.CreateHadoopIcebergCatalogDetails{}
		baseObject = details
	case strings.ToLower("NESSIE"):
		details := oci_golden_gate.CreateNessieIcebergCatalogDetails{}
		if branch, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("POLARIS"):
		details := oci_golden_gate.CreatePolarisIcebergCatalogDetails{}
		if clientId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if !useSecretIds {
			if clientSecret, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret")); ok {
				tmp := clientSecret.(string)
				details.ClientSecret = &tmp
			}
		}
		if useSecretIds {
			if clientSecretSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret_secret_id")); ok {
				tmp := clientSecretSecretId.(string)
				details.ClientSecretSecretId = &tmp
			}
		}
		if name, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if principalRole, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "principal_role")); ok {
			tmp := principalRole.(string)
			details.PrincipalRole = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("REST"):
		details := oci_golden_gate.CreateRestIcebergCatalogDetails{}
		if !useSecretIds {
			if properties, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
				tmp := properties.(string)
				details.Properties = &tmp
			}
		}
		if useSecretIds {
			if propertiesSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties_secret_id")); ok {
				tmp := propertiesSecretId.(string)
				details.PropertiesSecretId = &tmp
			}
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown catalog_type '%v' was specified", catalogType)
	}
	return baseObject, nil
}

func CreateIcebergCatalogDetailsToMap(obj *oci_golden_gate.CreateIcebergCatalogDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateGlueIcebergCatalogDetails:
		result["catalog_type"] = "GLUE"

		if v.GlueId != nil {
			result["glue_id"] = string(*v.GlueId)
		}
	case oci_golden_gate.CreateHadoopIcebergCatalogDetails:
		result["catalog_type"] = "HADOOP"
	case oci_golden_gate.CreateNessieIcebergCatalogDetails:
		result["catalog_type"] = "NESSIE"

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.CreatePolarisIcebergCatalogDetails:
		result["catalog_type"] = "POLARIS"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecret != nil {
			result["client_secret"] = string(*v.ClientSecret)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.PrincipalRole != nil {
			result["principal_role"] = string(*v.PrincipalRole)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.CreateRestIcebergCatalogDetails:
		result["catalog_type"] = "REST"

		if v.Properties != nil {
			result["properties"] = string(*v.Properties)
		}

		if v.PropertiesSecretId != nil {
			result["properties_secret_id"] = string(*v.PropertiesSecretId)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'catalog_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateIcebergStorageDetails(fieldKeyFormat string) (oci_golden_gate.CreateIcebergStorageDetails, error) {
	var baseObject oci_golden_gate.CreateIcebergStorageDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.CreateAmazonS3IcebergStorageDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if schemeType, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "scheme_type")); ok {
			details.SchemeType = oci_golden_gate.AmazonS3IcebergStorageSchemeTypeEnum(schemeType.(string))
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.CreateAzureDataLakeStorageIcebergStorageDetails{}
		if !useSecretIds {
			if accountKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key")); ok {
				tmp := accountKey.(string)
				details.AccountKey = &tmp
			}
		}
		if useSecretIds {
			if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
				tmp := accountKeySecretId.(string)
				details.AccountKeySecretId = &tmp
			}
		}
		if accountName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_name")); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if container, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "container")); ok {
			tmp := container.(string)
			details.Container = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.CreateGoogleCloudStorageIcebergStorageDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if projectId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		if !useSecretIds {
			if serviceAccountKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file")); ok {
				tmp := serviceAccountKeyFile.(string)
				details.ServiceAccountKeyFile = &tmp
			}
		}
		if useSecretIds {
			if serviceAccountKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file_secret_id")); ok {
				tmp := serviceAccountKeyFileSecretId.(string)
				details.ServiceAccountKeyFileSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_golden_gate.CreateNoIcebergStorageDetails{}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE_S3_API"):
		details := oci_golden_gate.CreateOciObjectStorageS3ApiIcebergStorageDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func CreateIcebergStorageDetailsToMap(obj *oci_golden_gate.CreateIcebergStorageDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateAmazonS3IcebergStorageDetails:
		result["storage_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["scheme_type"] = string(v.SchemeType)

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.CreateAzureDataLakeStorageIcebergStorageDetails:
		result["storage_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountKey != nil {
			result["account_key"] = string(*v.AccountKey)
		}

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		if v.Container != nil {
			result["container"] = string(*v.Container)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_golden_gate.CreateGoogleCloudStorageIcebergStorageDetails:
		result["storage_type"] = "GOOGLE_CLOUD_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ProjectId != nil {
			result["project_id"] = string(*v.ProjectId)
		}

		if v.ServiceAccountKeyFile != nil {
			result["service_account_key_file"] = string(*v.ServiceAccountKeyFile)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}
	case oci_golden_gate.CreateNoIcebergStorageDetails:
		result["storage_type"] = "NONE"
	case oci_golden_gate.CreateOciObjectStorageS3ApiIcebergStorageDetails:
		result["storage_type"] = "OCI_OBJECT_STORAGE_S3_API"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	default:
		log.Printf("[WARN] Received 'storage_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateKafkaBootstrapServer(fieldKeyFormat string) (oci_golden_gate.CreateKafkaBootstrapServer, error) {
	result := oci_golden_gate.CreateKafkaBootstrapServer{}

	if host, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func CreateKafkaBootstrapServerToMap(obj oci_golden_gate.CreateKafkaBootstrapServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToIcebergCatalog(fieldKeyFormat string) (oci_golden_gate.IcebergCatalog, error) {
	var baseObject oci_golden_gate.IcebergCatalog
	//discriminator
	catalogTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_type"))
	var catalogType string
	if ok {
		catalogType = catalogTypeRaw.(string)
	} else {
		catalogType = "" // default value
	}
	switch strings.ToLower(catalogType) {
	case strings.ToLower("GLUE"):
		details := oci_golden_gate.GlueIcebergCatalog{}
		if glueId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "glue_id")); ok {
			tmp := glueId.(string)
			details.GlueId = &tmp
		}
		baseObject = details
	case strings.ToLower("HADOOP"):
		details := oci_golden_gate.HadoopIcebergCatalog{}
		baseObject = details
	case strings.ToLower("NESSIE"):
		details := oci_golden_gate.NessieIcebergCatalog{}
		if branch, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("POLARIS"):
		details := oci_golden_gate.PolarisIcebergCatalog{}
		if clientId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret_secret_id")); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if name, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if principalRole, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "principal_role")); ok {
			tmp := principalRole.(string)
			details.PrincipalRole = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("REST"):
		details := oci_golden_gate.RestIcebergCatalog{}
		if propertiesSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties_secret_id")); ok {
			tmp := propertiesSecretId.(string)
			details.PropertiesSecretId = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown catalog_type '%v' was specified", catalogType)
	}
	return baseObject, nil
}

func IcebergCatalogToMap(obj *oci_golden_gate.IcebergCatalog) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.GlueIcebergCatalog:
		result["catalog_type"] = "GLUE"

		if v.GlueId != nil {
			result["glue_id"] = string(*v.GlueId)
		}
	case oci_golden_gate.HadoopIcebergCatalog:
		result["catalog_type"] = "HADOOP"
	case oci_golden_gate.NessieIcebergCatalog:
		result["catalog_type"] = "NESSIE"

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.PolarisIcebergCatalog:
		result["catalog_type"] = "POLARIS"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.PrincipalRole != nil {
			result["principal_role"] = string(*v.PrincipalRole)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.RestIcebergCatalog:
		result["catalog_type"] = "REST"

		if v.PropertiesSecretId != nil {
			result["properties_secret_id"] = string(*v.PropertiesSecretId)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'catalog_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToIcebergCatalogSummary(fieldKeyFormat string) (oci_golden_gate.IcebergCatalogSummary, error) {
	var baseObject oci_golden_gate.IcebergCatalogSummary
	//discriminator
	catalogTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_type"))
	var catalogType string
	if ok {
		catalogType = catalogTypeRaw.(string)
	} else {
		catalogType = "" // default value
	}
	switch strings.ToLower(catalogType) {
	case strings.ToLower("GLUE"):
		details := oci_golden_gate.GlueIcebergCatalogSummary{}
		if glueId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "glue_id")); ok {
			tmp := glueId.(string)
			details.GlueId = &tmp
		}
		baseObject = details
	case strings.ToLower("HADOOP"):
		details := oci_golden_gate.HadoopIcebergCatalogSummary{}
		baseObject = details
	case strings.ToLower("NESSIE"):
		details := oci_golden_gate.NessieIcebergCatalogSummary{}
		if branch, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("POLARIS"):
		details := oci_golden_gate.PolarisIcebergCatalogSummary{}
		if clientId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret_secret_id")); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if name, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if principalRole, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "principal_role")); ok {
			tmp := principalRole.(string)
			details.PrincipalRole = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("REST"):
		details := oci_golden_gate.RestIcebergCatalogSummary{}
		if propertiesSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties_secret_id")); ok {
			tmp := propertiesSecretId.(string)
			details.PropertiesSecretId = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown catalog_type '%v' was specified", catalogType)
	}
	return baseObject, nil
}

func IcebergCatalogSummaryToMap(obj *oci_golden_gate.IcebergCatalogSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.GlueIcebergCatalogSummary:
		result["catalog_type"] = "GLUE"

		if v.GlueId != nil {
			result["glue_id"] = string(*v.GlueId)
		}
	case oci_golden_gate.HadoopIcebergCatalogSummary:
		result["catalog_type"] = "HADOOP"
	case oci_golden_gate.NessieIcebergCatalogSummary:
		result["catalog_type"] = "NESSIE"

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.PolarisIcebergCatalogSummary:
		result["catalog_type"] = "POLARIS"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.PrincipalRole != nil {
			result["principal_role"] = string(*v.PrincipalRole)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.RestIcebergCatalogSummary:
		result["catalog_type"] = "REST"

		if v.PropertiesSecretId != nil {
			result["properties_secret_id"] = string(*v.PropertiesSecretId)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'catalog_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToIcebergStorage(fieldKeyFormat string) (oci_golden_gate.IcebergStorage, error) {
	var baseObject oci_golden_gate.IcebergStorage
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.AmazonS3IcebergStorage{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if schemeType, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "scheme_type")); ok {
			details.SchemeType = oci_golden_gate.AmazonS3IcebergStorageSchemeTypeEnum(schemeType.(string))
		}
		if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.AzureDataLakeStorageIcebergStorage{}
		if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if accountName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_name")); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if container, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "container")); ok {
			tmp := container.(string)
			details.Container = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.GoogleCloudStorageIcebergStorage{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if projectId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file_secret_id")); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_golden_gate.NoIcebergStorage{}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE_S3_API"):
		details := oci_golden_gate.OciObjectStorageS3ApiIcebergStorage{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func IcebergStorageToMap(obj *oci_golden_gate.IcebergStorage) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.AmazonS3IcebergStorage:
		result["storage_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["scheme_type"] = string(v.SchemeType)

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.AzureDataLakeStorageIcebergStorage:
		result["storage_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		if v.Container != nil {
			result["container"] = string(*v.Container)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_golden_gate.GoogleCloudStorageIcebergStorage:
		result["storage_type"] = "GOOGLE_CLOUD_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ProjectId != nil {
			result["project_id"] = string(*v.ProjectId)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}
	case oci_golden_gate.NoIcebergStorage:
		result["storage_type"] = "NONE"
	case oci_golden_gate.OciObjectStorageS3ApiIcebergStorage:
		result["storage_type"] = "OCI_OBJECT_STORAGE_S3_API"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	default:
		log.Printf("[WARN] Received 'storage_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToIcebergStorageSummary(fieldKeyFormat string) (oci_golden_gate.IcebergStorageSummary, error) {
	var baseObject oci_golden_gate.IcebergStorageSummary
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.AmazonS3IcebergStorageSummary{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if schemeType, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "scheme_type")); ok {
			details.SchemeType = oci_golden_gate.AmazonS3IcebergStorageSchemeTypeEnum(schemeType.(string))
		}
		if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.AzureDataLakeStorageIcebergStorageSummary{}
		if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if accountName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_name")); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if container, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "container")); ok {
			tmp := container.(string)
			details.Container = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.GoogleCloudStorageIcebergStorageSummary{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if projectId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file_secret_id")); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_golden_gate.NoIcebergStorageSummary{}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE_S3_API"):
		details := oci_golden_gate.OciObjectStorageS3ApiIcebergStorageSummary{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func IcebergStorageSummaryToMap(obj *oci_golden_gate.IcebergStorageSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.AmazonS3IcebergStorageSummary:
		result["storage_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["scheme_type"] = string(v.SchemeType)

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.AzureDataLakeStorageIcebergStorageSummary:
		result["storage_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		if v.Container != nil {
			result["container"] = string(*v.Container)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_golden_gate.GoogleCloudStorageIcebergStorageSummary:
		result["storage_type"] = "GOOGLE_CLOUD_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ProjectId != nil {
			result["project_id"] = string(*v.ProjectId)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}
	case oci_golden_gate.NoIcebergStorageSummary:
		result["storage_type"] = "NONE"
	case oci_golden_gate.OciObjectStorageS3ApiIcebergStorageSummary:
		result["storage_type"] = "OCI_OBJECT_STORAGE_S3_API"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	default:
		log.Printf("[WARN] Received 'storage_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func IngressIpDetailsToMap(obj oci_golden_gate.IngressIpDetails) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.IngressIp != nil {
		result["ingress_ip"] = string(*obj.IngressIp)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToKafkaBootstrapServer(fieldKeyFormat string) (oci_golden_gate.KafkaBootstrapServer, error) {
	result := oci_golden_gate.KafkaBootstrapServer{}

	if host, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "host")); ok {
		tmp := host.(string)
		result.Host = &tmp
	}

	if port, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(int)
		result.Port = &tmp
	}

	return result, nil
}

func KafkaBootstrapServerToMap(obj oci_golden_gate.KafkaBootstrapServer) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Host != nil {
		result["host"] = string(*obj.Host)
	}

	if obj.Port != nil {
		result["port"] = int(*obj.Port)
	}

	if obj.PrivateIp != nil {
		result["private_ip"] = string(*obj.PrivateIp)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToNameValuePair(fieldKeyFormat string) (oci_golden_gate.NameValuePair, error) {
	result := oci_golden_gate.NameValuePair{}

	if name, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func NameValuePairToMap(obj oci_golden_gate.NameValuePair) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateAiModelAuthDetails(fieldKeyFormat string) (oci_golden_gate.UpdateAiModelAuthDetails, error) {
	var baseObject oci_golden_gate.UpdateAiModelAuthDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "auth_type"))
	var authType string
	if ok {
		authType = authTypeRaw.(string)
	} else {
		authType = "" // default value
	}
	switch strings.ToLower(authType) {
	case strings.ToLower("API_KEY"):
		details := oci_golden_gate.UpdateApiKeyAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if baseUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "base_url")); ok {
			tmp := baseUrl.(string)
			details.BaseUrl = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_GEN_AI"):
		details := oci_golden_gate.UpdateOciGenAiModelAuthDetails{}
		if !useSecretIds {
			if apiKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key")); ok {
				tmp := apiKey.(string)
				details.ApiKey = &tmp
			}
		}
		if useSecretIds {
			if apiKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "api_key_secret_id")); ok {
				tmp := apiKeySecretId.(string)
				details.ApiKeySecretId = &tmp
			}
		}
		if keyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "key_fingerprint")); ok {
			tmp := keyFingerprint.(string)
			details.KeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown auth_type '%v' was specified", authType)
	}
	return baseObject, nil
}

func UpdateAiModelAuthDetailsToMap(obj *oci_golden_gate.UpdateAiModelAuthDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateApiKeyAiModelAuthDetails:
		result["auth_type"] = "API_KEY"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.BaseUrl != nil {
			result["base_url"] = string(*v.BaseUrl)
		}
	case oci_golden_gate.UpdateOciGenAiModelAuthDetails:
		result["auth_type"] = "OCI_GEN_AI"

		if v.ApiKey != nil {
			result["api_key"] = string(*v.ApiKey)
		}

		if v.ApiKeySecretId != nil {
			result["api_key_secret_id"] = string(*v.ApiKeySecretId)
		}

		if v.KeyFingerprint != nil {
			result["key_fingerprint"] = string(*v.KeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'auth_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateIcebergCatalogDetails(fieldKeyFormat string) (oci_golden_gate.UpdateIcebergCatalogDetails, error) {
	var baseObject oci_golden_gate.UpdateIcebergCatalogDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	catalogTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "catalog_type"))
	var catalogType string
	if ok {
		catalogType = catalogTypeRaw.(string)
	} else {
		catalogType = "" // default value
	}
	switch strings.ToLower(catalogType) {
	case strings.ToLower("GLUE"):
		details := oci_golden_gate.UpdateGlueIcebergCatalogDetails{}
		if glueId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "glue_id")); ok {
			tmp := glueId.(string)
			details.GlueId = &tmp
		}
		baseObject = details
	case strings.ToLower("HADOOP"):
		details := oci_golden_gate.UpdateHadoopIcebergCatalogDetails{}
		baseObject = details
	case strings.ToLower("NESSIE"):
		details := oci_golden_gate.UpdateNessieIcebergCatalogDetails{}
		if branch, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "branch")); ok {
			tmp := branch.(string)
			details.Branch = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("POLARIS"):
		details := oci_golden_gate.UpdatePolarisIcebergCatalogDetails{}
		if clientId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_id")); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if !useSecretIds {
			if clientSecret, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret")); ok {
				tmp := clientSecret.(string)
				details.ClientSecret = &tmp
			}
		}
		if useSecretIds {
			if clientSecretSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "client_secret_secret_id")); ok {
				tmp := clientSecretSecretId.(string)
				details.ClientSecretSecretId = &tmp
			}
		}
		if name, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "name")); ok {
			tmp := name.(string)
			details.Name = &tmp
		}
		if principalRole, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "principal_role")); ok {
			tmp := principalRole.(string)
			details.PrincipalRole = &tmp
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	case strings.ToLower("REST"):
		details := oci_golden_gate.UpdateRestIcebergCatalogDetails{}
		if !useSecretIds {
			if properties, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties")); ok {
				tmp := properties.(string)
				details.Properties = &tmp
			}
		}
		if useSecretIds {
			if propertiesSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "properties_secret_id")); ok {
				tmp := propertiesSecretId.(string)
				details.PropertiesSecretId = &tmp
			}
		}
		if uri, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "uri")); ok {
			tmp := uri.(string)
			details.Uri = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown catalog_type '%v' was specified", catalogType)
	}
	return baseObject, nil
}

func UpdateIcebergCatalogDetailsToMap(obj *oci_golden_gate.UpdateIcebergCatalogDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateGlueIcebergCatalogDetails:
		result["catalog_type"] = "GLUE"

		if v.GlueId != nil {
			result["glue_id"] = string(*v.GlueId)
		}
	case oci_golden_gate.UpdateHadoopIcebergCatalogDetails:
		result["catalog_type"] = "HADOOP"
	case oci_golden_gate.UpdateNessieIcebergCatalogDetails:
		result["catalog_type"] = "NESSIE"

		if v.Branch != nil {
			result["branch"] = string(*v.Branch)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.UpdatePolarisIcebergCatalogDetails:
		result["catalog_type"] = "POLARIS"

		if v.ClientId != nil {
			result["client_id"] = string(*v.ClientId)
		}

		if v.ClientSecret != nil {
			result["client_secret"] = string(*v.ClientSecret)
		}

		if v.ClientSecretSecretId != nil {
			result["client_secret_secret_id"] = string(*v.ClientSecretSecretId)
		}

		if v.Name != nil {
			result["name"] = string(*v.Name)
		}

		if v.PrincipalRole != nil {
			result["principal_role"] = string(*v.PrincipalRole)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	case oci_golden_gate.UpdateRestIcebergCatalogDetails:
		result["catalog_type"] = "REST"

		if v.Properties != nil {
			result["properties"] = string(*v.Properties)
		}

		if v.PropertiesSecretId != nil {
			result["properties_secret_id"] = string(*v.PropertiesSecretId)
		}

		if v.Uri != nil {
			result["uri"] = string(*v.Uri)
		}
	default:
		log.Printf("[WARN] Received 'catalog_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateIcebergStorageDetails(fieldKeyFormat string) (oci_golden_gate.UpdateIcebergStorageDetails, error) {
	var baseObject oci_golden_gate.UpdateIcebergStorageDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	storageTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_type"))
	var storageType string
	if ok {
		storageType = storageTypeRaw.(string)
	} else {
		storageType = "" // default value
	}
	switch strings.ToLower(storageType) {
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.UpdateAmazonS3IcebergStorageDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if schemeType, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "scheme_type")); ok {
			details.SchemeType = oci_golden_gate.AmazonS3IcebergStorageSchemeTypeEnum(schemeType.(string))
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.UpdateAzureDataLakeStorageIcebergStorageDetails{}
		if !useSecretIds {
			if accountKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key")); ok {
				tmp := accountKey.(string)
				details.AccountKey = &tmp
			}
		}
		if useSecretIds {
			if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
				tmp := accountKeySecretId.(string)
				details.AccountKeySecretId = &tmp
			}
		}
		if accountName, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_name")); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if container, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "container")); ok {
			tmp := container.(string)
			details.Container = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		baseObject = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.UpdateGoogleCloudStorageIcebergStorageDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if projectId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "project_id")); ok {
			tmp := projectId.(string)
			details.ProjectId = &tmp
		}
		if !useSecretIds {
			if serviceAccountKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file")); ok {
				tmp := serviceAccountKeyFile.(string)
				details.ServiceAccountKeyFile = &tmp
			}
		}
		if useSecretIds {
			if serviceAccountKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "service_account_key_file_secret_id")); ok {
				tmp := serviceAccountKeyFileSecretId.(string)
				details.ServiceAccountKeyFileSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("NONE"):
		details := oci_golden_gate.UpdateNoIcebergStorageDetails{}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE_S3_API"):
		details := oci_golden_gate.UpdateOciObjectStorageS3ApiIcebergStorageDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if endpoint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint")); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown storage_type '%v' was specified", storageType)
	}
	return baseObject, nil
}

func UpdateIcebergStorageDetailsToMap(obj *oci_golden_gate.UpdateIcebergStorageDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateAmazonS3IcebergStorageDetails:
		result["storage_type"] = "AMAZON_S3"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		result["scheme_type"] = string(v.SchemeType)

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.UpdateAzureDataLakeStorageIcebergStorageDetails:
		result["storage_type"] = "AZURE_DATA_LAKE_STORAGE"

		if v.AccountKey != nil {
			result["account_key"] = string(*v.AccountKey)
		}

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.AccountName != nil {
			result["account_name"] = string(*v.AccountName)
		}

		if v.Container != nil {
			result["container"] = string(*v.Container)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}
	case oci_golden_gate.UpdateGoogleCloudStorageIcebergStorageDetails:
		result["storage_type"] = "GOOGLE_CLOUD_STORAGE"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.ProjectId != nil {
			result["project_id"] = string(*v.ProjectId)
		}

		if v.ServiceAccountKeyFile != nil {
			result["service_account_key_file"] = string(*v.ServiceAccountKeyFile)
		}

		if v.ServiceAccountKeyFileSecretId != nil {
			result["service_account_key_file_secret_id"] = string(*v.ServiceAccountKeyFileSecretId)
		}
	case oci_golden_gate.UpdateNoIcebergStorageDetails:
		result["storage_type"] = "NONE"
	case oci_golden_gate.UpdateOciObjectStorageS3ApiIcebergStorageDetails:
		result["storage_type"] = "OCI_OBJECT_STORAGE_S3_API"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Endpoint != nil {
			result["endpoint"] = string(*v.Endpoint)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	default:
		log.Printf("[WARN] Received 'storage_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateDatabricksEngineArtifactLocationDetails(fieldKeyFormat string) (oci_golden_gate.CreateDatabricksEngineArtifactLocationDetails, error) {
	var baseObject oci_golden_gate.CreateDatabricksEngineArtifactLocationDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("AZURE_DATE_LAKE_STORAGE"):
		details := oci_golden_gate.CreateAzureDataLakeStorageDatabricksEngineArtifactLocationDetails{}
		if !useSecretIds {
			if accountKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key")); ok {
				tmp := accountKey.(string)
				details.AccountKey = &tmp
			}
		}
		if useSecretIds {
			if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
				tmp := accountKeySecretId.(string)
				details.AccountKeySecretId = &tmp
			}
		}
		if endpointSuffix, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint_suffix")); ok {
			tmp := endpointSuffix.(string)
			details.EndpointSuffix = &tmp
		}
		if filesystem, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "filesystem")); ok {
			tmp := filesystem.(string)
			details.Filesystem = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if storageAccount, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "storage_account")); ok {
			tmp := storageAccount.(string)
			details.StorageAccount = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.CreateOciObjectStorageDatabricksEngineArtifactLocationDetails{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToCreateDatabricksEngineOciObjectStorageConfigDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("UC_VOLUME"):
		details := oci_golden_gate.CreateUcVolumeDatabricksEngineArtifactLocationDetails{}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func CreateDatabricksEngineArtifactLocationDetailsToMap(obj *oci_golden_gate.CreateDatabricksEngineArtifactLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateAzureDataLakeStorageDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "AZURE_DATE_LAKE_STORAGE"

		if v.AccountKey != nil {
			result["account_key"] = string(*v.AccountKey)
		}

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.EndpointSuffix != nil {
			result["endpoint_suffix"] = string(*v.EndpointSuffix)
		}

		if v.Filesystem != nil {
			result["filesystem"] = string(*v.Filesystem)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.StorageAccount != nil {
			result["storage_account"] = string(*v.StorageAccount)
		}
	case oci_golden_gate.CreateOciObjectStorageDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := CreateDatabricksEngineOciObjectStorageConfigDetailsToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	case oci_golden_gate.CreateUcVolumeDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "UC_VOLUME"

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateDatabricksEngineOciObjectStorageConfigDetails(fieldKeyFormat string) (oci_golden_gate.CreateDatabricksEngineOciObjectStorageConfigDetails, error) {
	var baseObject oci_golden_gate.CreateDatabricksEngineOciObjectStorageConfigDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.CreateDatabricksEngineSecretKeyOciObjectStorageConfigDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.CreateDatabricksEngineParOciObjectStorageConfigDetails{}
		if !useSecretIds {
			if parUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url")); ok {
				tmp := parUrl.(string)
				details.ParUrl = &tmp
			}
		}
		if useSecretIds {
			if parUrlSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
				tmp := parUrlSecretId.(string)
				details.ParUrlSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.CreateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.CreateDatabricksEngineUserPrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if !useSecretIds {
			if privateKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file")); ok {
				tmp := privateKeyFile.(string)
				details.PrivateKeyFile = &tmp
			}
		}
		if useSecretIds {
			if privateKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
				tmp := privateKeyFileSecretId.(string)
				details.PrivateKeyFileSecretId = &tmp
			}
		}
		if !useSecretIds {
			if privateKeyPassphrase, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase")); ok {
				tmp := privateKeyPassphrase.(string)
				details.PrivateKeyPassphrase = &tmp
			}
		}
		if useSecretIds {
			if privateKeyPassphraseSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
				tmp := privateKeyPassphraseSecretId.(string)
				details.PrivateKeyPassphraseSecretId = &tmp
			}
		}
		if publicKeyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func CreateDatabricksEngineOciObjectStorageConfigDetailsToMap(obj *oci_golden_gate.CreateDatabricksEngineOciObjectStorageConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateDatabricksEngineSecretKeyOciObjectStorageConfigDetails:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.CreateDatabricksEngineParOciObjectStorageConfigDetails:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrl != nil {
			result["par_url"] = string(*v.ParUrl)
		}

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.CreateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.CreateDatabricksEngineUserPrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFile != nil {
			result["private_key_file"] = string(*v.PrivateKeyFile)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphrase != nil {
			result["private_key_passphrase"] = string(*v.PrivateKeyPassphrase)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateSparkArtifactLocationDetails(fieldKeyFormat string) (oci_golden_gate.CreateSparkArtifactLocationDetails, error) {
	var baseObject oci_golden_gate.CreateSparkArtifactLocationDetails
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.CreateOciObjectStorageSparkArtifactLocationDetails{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToCreateSparkOciObjectStorageConfigDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func CreateSparkArtifactLocationDetailsToMap(obj *oci_golden_gate.CreateSparkArtifactLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateOciObjectStorageSparkArtifactLocationDetails:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := CreateSparkOciObjectStorageConfigDetailsToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToCreateSparkOciObjectStorageConfigDetails(fieldKeyFormat string) (oci_golden_gate.CreateSparkOciObjectStorageConfigDetails, error) {
	var baseObject oci_golden_gate.CreateSparkOciObjectStorageConfigDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.CreateSparkSecretKeyOciObjectStorageConfigDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.CreateSparkParOciObjectStorageConfigDetails{}
		if !useSecretIds {
			if parUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url")); ok {
				tmp := parUrl.(string)
				details.ParUrl = &tmp
			}
		}
		if useSecretIds {
			if parUrlSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
				tmp := parUrlSecretId.(string)
				details.ParUrlSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.CreateSparkResourcePrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.CreateSparkUserPrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if !useSecretIds {
			if privateKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file")); ok {
				tmp := privateKeyFile.(string)
				details.PrivateKeyFile = &tmp
			}
		}
		if useSecretIds {
			if privateKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
				tmp := privateKeyFileSecretId.(string)
				details.PrivateKeyFileSecretId = &tmp
			}
		}
		if !useSecretIds {
			if privateKeyPassphrase, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase")); ok {
				tmp := privateKeyPassphrase.(string)
				details.PrivateKeyPassphrase = &tmp
			}
		}
		if useSecretIds {
			if privateKeyPassphraseSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
				tmp := privateKeyPassphraseSecretId.(string)
				details.PrivateKeyPassphraseSecretId = &tmp
			}
		}
		if publicKeyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func CreateSparkOciObjectStorageConfigDetailsToMap(obj *oci_golden_gate.CreateSparkOciObjectStorageConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.CreateSparkSecretKeyOciObjectStorageConfigDetails:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.CreateSparkParOciObjectStorageConfigDetails:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrl != nil {
			result["par_url"] = string(*v.ParUrl)
		}

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.CreateSparkResourcePrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.CreateSparkUserPrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFile != nil {
			result["private_key_file"] = string(*v.PrivateKeyFile)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphrase != nil {
			result["private_key_passphrase"] = string(*v.PrivateKeyPassphrase)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToDatabricksEngineArtifactLocation(fieldKeyFormat string) (oci_golden_gate.DatabricksEngineArtifactLocation, error) {
	var baseObject oci_golden_gate.DatabricksEngineArtifactLocation
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("AZURE_DATE_LAKE_STORAGE"):
		details := oci_golden_gate.AzureDataLakeStorageDatabricksEngineArtifactLocation{}
		if accountKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if endpointSuffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_suffix")); ok {
			tmp := endpointSuffix.(string)
			details.EndpointSuffix = &tmp
		}
		if filesystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filesystem")); ok {
			tmp := filesystem.(string)
			details.Filesystem = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if storageAccount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_account")); ok {
			tmp := storageAccount.(string)
			details.StorageAccount = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.OciObjectStorageDatabricksEngineArtifactLocation{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToDatabricksEngineOciObjectStorageConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("UC_VOLUME"):
		details := oci_golden_gate.UcVolumeDatabricksEngineArtifactLocation{}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func DatabricksEngineArtifactLocationToMap(obj *oci_golden_gate.DatabricksEngineArtifactLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.AzureDataLakeStorageDatabricksEngineArtifactLocation:
		result["location_type"] = "AZURE_DATE_LAKE_STORAGE"

		if v.ShouldOverwriteExisting != nil {
			result["should_overwrite_existing"] = bool(*v.ShouldOverwriteExisting)
		}

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.EndpointSuffix != nil {
			result["endpoint_suffix"] = string(*v.EndpointSuffix)
		}

		if v.Filesystem != nil {
			result["filesystem"] = string(*v.Filesystem)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.StorageAccount != nil {
			result["storage_account"] = string(*v.StorageAccount)
		}
	case oci_golden_gate.OciObjectStorageDatabricksEngineArtifactLocation:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.ShouldOverwriteExisting != nil {
			result["should_overwrite_existing"] = bool(*v.ShouldOverwriteExisting)
		}

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := DatabricksEngineOciObjectStorageConfigToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	case oci_golden_gate.UcVolumeDatabricksEngineArtifactLocation:
		result["location_type"] = "UC_VOLUME"

		if v.ShouldOverwriteExisting != nil {
			result["should_overwrite_existing"] = bool(*v.ShouldOverwriteExisting)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToDatabricksEngineArtifactLocationSummary(fieldKeyFormat string) (oci_golden_gate.DatabricksEngineArtifactLocationSummary, error) {
	var baseObject oci_golden_gate.DatabricksEngineArtifactLocationSummary
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("AZURE_DATE_LAKE_STORAGE"):
		details := oci_golden_gate.AzureDataLakeStorageDatabricksEngineArtifactLocationSummary{}
		if accountKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if endpointSuffix, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "endpoint_suffix")); ok {
			tmp := endpointSuffix.(string)
			details.EndpointSuffix = &tmp
		}
		if filesystem, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "filesystem")); ok {
			tmp := filesystem.(string)
			details.Filesystem = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if storageAccount, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "storage_account")); ok {
			tmp := storageAccount.(string)
			details.StorageAccount = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.OciObjectStorageDatabricksEngineArtifactLocationSummary{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToDatabricksEngineOciObjectStorageConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("UC_VOLUME"):
		details := oci_golden_gate.UcVolumeDatabricksEngineArtifactLocationSummary{}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func DatabricksEngineArtifactLocationSummaryToMap(obj *oci_golden_gate.DatabricksEngineArtifactLocationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.AzureDataLakeStorageDatabricksEngineArtifactLocationSummary:
		result["location_type"] = "AZURE_DATE_LAKE_STORAGE"

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.EndpointSuffix != nil {
			result["endpoint_suffix"] = string(*v.EndpointSuffix)
		}

		if v.Filesystem != nil {
			result["filesystem"] = string(*v.Filesystem)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.StorageAccount != nil {
			result["storage_account"] = string(*v.StorageAccount)
		}
	case oci_golden_gate.OciObjectStorageDatabricksEngineArtifactLocationSummary:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := DatabricksEngineOciObjectStorageConfigToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	case oci_golden_gate.UcVolumeDatabricksEngineArtifactLocationSummary:
		result["location_type"] = "UC_VOLUME"

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToDatabricksEngineOciObjectStorageConfig(fieldKeyFormat string) (oci_golden_gate.DatabricksEngineOciObjectStorageConfig, error) {
	var baseObject oci_golden_gate.DatabricksEngineOciObjectStorageConfig
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.DatabricksEngineSecretKeyOciObjectStorageConfig{}
		if accessKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.DatabricksEngineParOciObjectStorageConfig{}
		if parUrlSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
			tmp := parUrlSecretId.(string)
			details.ParUrlSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.DatabricksEngineResourcePrincipalOciObjectStorageConfig{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.DatabricksEngineUserPrincipalOciObjectStorageConfig{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func DatabricksEngineOciObjectStorageConfigToMap(obj *oci_golden_gate.DatabricksEngineOciObjectStorageConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.DatabricksEngineSecretKeyOciObjectStorageConfig:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.DatabricksEngineParOciObjectStorageConfig:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.DatabricksEngineResourcePrincipalOciObjectStorageConfig:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.DatabricksEngineUserPrincipalOciObjectStorageConfig:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToSparkArtifactLocation(fieldKeyFormat string) (oci_golden_gate.SparkArtifactLocation, error) {
	var baseObject oci_golden_gate.SparkArtifactLocation
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.OciObjectStorageSparkArtifactLocation{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToSparkOciObjectStorageConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func SparkArtifactLocationToMap(obj *oci_golden_gate.SparkArtifactLocation) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.OciObjectStorageSparkArtifactLocation:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.ShouldOverwriteExisting != nil {
			result["should_overwrite_existing"] = bool(*v.ShouldOverwriteExisting)
		}

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := SparkOciObjectStorageConfigToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToSparkArtifactLocationSummary(fieldKeyFormat string) (oci_golden_gate.SparkArtifactLocationSummary, error) {
	var baseObject oci_golden_gate.SparkArtifactLocationSummary
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.OciObjectStorageSparkArtifactLocationSummary{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToSparkOciObjectStorageConfig(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func SparkArtifactLocationSummaryToMap(obj *oci_golden_gate.SparkArtifactLocationSummary) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.OciObjectStorageSparkArtifactLocationSummary:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := SparkOciObjectStorageConfigToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToSparkOciObjectStorageConfig(fieldKeyFormat string) (oci_golden_gate.SparkOciObjectStorageConfig, error) {
	var baseObject oci_golden_gate.SparkOciObjectStorageConfig
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.SparkSecretKeyOciObjectStorageConfig{}
		if accessKeyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.SparkParOciObjectStorageConfig{}
		if parUrlSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
			tmp := parUrlSecretId.(string)
			details.ParUrlSecretId = &tmp
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.SparkResourcePrincipalOciObjectStorageConfig{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.SparkUserPrincipalOciObjectStorageConfig{}
		if bucket, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func SparkOciObjectStorageConfigToMap(obj *oci_golden_gate.SparkOciObjectStorageConfig) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.SparkSecretKeyOciObjectStorageConfig:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.SparkParOciObjectStorageConfig:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.SparkResourcePrincipalOciObjectStorageConfig:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.SparkUserPrincipalOciObjectStorageConfig:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateDatabricksEngineArtifactLocationDetails(fieldKeyFormat string) (oci_golden_gate.UpdateDatabricksEngineArtifactLocationDetails, error) {
	var baseObject oci_golden_gate.UpdateDatabricksEngineArtifactLocationDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("AZURE_DATE_LAKE_STORAGE"):
		details := oci_golden_gate.UpdateAzureDataLakeStorageDatabricksEngineArtifactLocationDetails{}
		if !useSecretIds {
			if accountKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key")); ok {
				tmp := accountKey.(string)
				details.AccountKey = &tmp
			}
		}
		if useSecretIds {
			if accountKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "account_key_secret_id")); ok {
				tmp := accountKeySecretId.(string)
				details.AccountKeySecretId = &tmp
			}
		}
		if endpointSuffix, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "endpoint_suffix")); ok {
			tmp := endpointSuffix.(string)
			details.EndpointSuffix = &tmp
		}
		if filesystem, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "filesystem")); ok {
			tmp := filesystem.(string)
			details.Filesystem = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if storageAccount, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "storage_account")); ok {
			tmp := storageAccount.(string)
			details.StorageAccount = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.UpdateOciObjectStorageDatabricksEngineArtifactLocationDetails{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToUpdateDatabricksEngineOciObjectStorageConfigDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	case strings.ToLower("UC_VOLUME"):
		details := oci_golden_gate.UpdateUcVolumeDatabricksEngineArtifactLocationDetails{}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func UpdateDatabricksEngineArtifactLocationDetailsToMap(obj *oci_golden_gate.UpdateDatabricksEngineArtifactLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateAzureDataLakeStorageDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "AZURE_DATE_LAKE_STORAGE"

		if v.AccountKey != nil {
			result["account_key"] = string(*v.AccountKey)
		}

		if v.AccountKeySecretId != nil {
			result["account_key_secret_id"] = string(*v.AccountKeySecretId)
		}

		if v.EndpointSuffix != nil {
			result["endpoint_suffix"] = string(*v.EndpointSuffix)
		}

		if v.Filesystem != nil {
			result["filesystem"] = string(*v.Filesystem)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.StorageAccount != nil {
			result["storage_account"] = string(*v.StorageAccount)
		}
	case oci_golden_gate.UpdateOciObjectStorageDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := UpdateDatabricksEngineOciObjectStorageConfigDetailsToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	case oci_golden_gate.UpdateUcVolumeDatabricksEngineArtifactLocationDetails:
		result["location_type"] = "UC_VOLUME"

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateDatabricksEngineOciObjectStorageConfigDetails(fieldKeyFormat string) (oci_golden_gate.UpdateDatabricksEngineOciObjectStorageConfigDetails, error) {
	var baseObject oci_golden_gate.UpdateDatabricksEngineOciObjectStorageConfigDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.UpdateDatabricksEngineSecretKeyOciObjectStorageConfigDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.UpdateDatabricksEngineParOciObjectStorageConfigDetails{}
		if !useSecretIds {
			if parUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url")); ok {
				tmp := parUrl.(string)
				details.ParUrl = &tmp
			}
		}
		if useSecretIds {
			if parUrlSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
				tmp := parUrlSecretId.(string)
				details.ParUrlSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.UpdateDatabricksEngineUserPrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if !useSecretIds {
			if privateKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file")); ok {
				tmp := privateKeyFile.(string)
				details.PrivateKeyFile = &tmp
			}
		}
		if useSecretIds {
			if privateKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
				tmp := privateKeyFileSecretId.(string)
				details.PrivateKeyFileSecretId = &tmp
			}
		}
		if !useSecretIds {
			if privateKeyPassphrase, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase")); ok {
				tmp := privateKeyPassphrase.(string)
				details.PrivateKeyPassphrase = &tmp
			}
		}
		if useSecretIds {
			if privateKeyPassphraseSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
				tmp := privateKeyPassphraseSecretId.(string)
				details.PrivateKeyPassphraseSecretId = &tmp
			}
		}
		if publicKeyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func UpdateDatabricksEngineOciObjectStorageConfigDetailsToMap(obj *oci_golden_gate.UpdateDatabricksEngineOciObjectStorageConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateDatabricksEngineSecretKeyOciObjectStorageConfigDetails:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.UpdateDatabricksEngineParOciObjectStorageConfigDetails:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrl != nil {
			result["par_url"] = string(*v.ParUrl)
		}

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.UpdateDatabricksEngineResourcePrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.UpdateDatabricksEngineUserPrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFile != nil {
			result["private_key_file"] = string(*v.PrivateKeyFile)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphrase != nil {
			result["private_key_passphrase"] = string(*v.PrivateKeyPassphrase)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateSparkArtifactLocationDetails(fieldKeyFormat string) (oci_golden_gate.UpdateSparkArtifactLocationDetails, error) {
	var baseObject oci_golden_gate.UpdateSparkArtifactLocationDetails
	//discriminator
	locationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "location_type"))
	var locationType string
	if ok {
		locationType = locationTypeRaw.(string)
	} else {
		locationType = "" // default value
	}
	switch strings.ToLower(locationType) {
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.UpdateOciObjectStorageSparkArtifactLocationDetails{}
		if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
			if tmpList := configuration.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "configuration"), 0)
				tmp, err := s.mapToUpdateSparkOciObjectStorageConfigDetails(fieldKeyFormatNextLevel)
				if err != nil {
					return details, fmt.Errorf("unable to convert configuration, encountered error: %v", err)
				}
				details.Configuration = tmp
			}
		}
		if shouldOverwriteExisting, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "should_overwrite_existing")); ok {
			tmp := shouldOverwriteExisting.(bool)
			details.ShouldOverwriteExisting = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown location_type '%v' was specified", locationType)
	}
	return baseObject, nil
}

func UpdateSparkArtifactLocationDetailsToMap(obj *oci_golden_gate.UpdateSparkArtifactLocationDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateOciObjectStorageSparkArtifactLocationDetails:
		result["location_type"] = "OCI_OBJECT_STORAGE"

		if v.Configuration != nil {
			configurationArray := []interface{}{}
			if configurationMap := UpdateSparkOciObjectStorageConfigDetailsToMap(&v.Configuration); configurationMap != nil {
				configurationArray = append(configurationArray, configurationMap)
			}
			result["configuration"] = configurationArray
		}
	default:
		log.Printf("[WARN] Received 'location_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) mapToUpdateSparkOciObjectStorageConfigDetails(fieldKeyFormat string) (oci_golden_gate.UpdateSparkOciObjectStorageConfigDetails, error) {
	var baseObject oci_golden_gate.UpdateSparkOciObjectStorageConfigDetails
	useSecretIds := false
	if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
		useSecretIds = doesUseSecretIds.(bool)
	}
	//discriminator
	authenticationTypeRaw, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "authentication_type"))
	var authenticationType string
	if ok {
		authenticationType = authenticationTypeRaw.(string)
	} else {
		authenticationType = "" // default value
	}
	switch strings.ToLower(authenticationType) {
	case strings.ToLower("CUSTOMER_SECRET_KEY"):
		details := oci_golden_gate.UpdateSparkSecretKeyOciObjectStorageConfigDetails{}
		if accessKeyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "access_key_id")); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if !useSecretIds {
			if secretAccessKey, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key")); ok {
				tmp := secretAccessKey.(string)
				details.SecretAccessKey = &tmp
			}
		}
		if useSecretIds {
			if secretAccessKeySecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "secret_access_key_secret_id")); ok {
				tmp := secretAccessKeySecretId.(string)
				details.SecretAccessKeySecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("PRE_AUTHENTICATED_REQUEST"):
		details := oci_golden_gate.UpdateSparkParOciObjectStorageConfigDetails{}
		if !useSecretIds {
			if parUrl, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url")); ok {
				tmp := parUrl.(string)
				details.ParUrl = &tmp
			}
		}
		if useSecretIds {
			if parUrlSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "par_url_secret_id")); ok {
				tmp := parUrlSecretId.(string)
				details.ParUrlSecretId = &tmp
			}
		}
		baseObject = details
	case strings.ToLower("RESOURCE_PRINCIPAL"):
		details := oci_golden_gate.UpdateSparkResourcePrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		baseObject = details
	case strings.ToLower("USER_PRINCIPAL"):
		details := oci_golden_gate.UpdateSparkUserPrincipalOciObjectStorageConfigDetails{}
		if bucket, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "bucket")); ok {
			tmp := bucket.(string)
			details.Bucket = &tmp
		}
		if namespace, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
			tmp := namespace.(string)
			details.Namespace = &tmp
		}
		if path, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "path")); ok {
			tmp := path.(string)
			details.Path = &tmp
		}
		if !useSecretIds {
			if privateKeyFile, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file")); ok {
				tmp := privateKeyFile.(string)
				details.PrivateKeyFile = &tmp
			}
		}
		if useSecretIds {
			if privateKeyFileSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_file_secret_id")); ok {
				tmp := privateKeyFileSecretId.(string)
				details.PrivateKeyFileSecretId = &tmp
			}
		}
		if !useSecretIds {
			if privateKeyPassphrase, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase")); ok {
				tmp := privateKeyPassphrase.(string)
				details.PrivateKeyPassphrase = &tmp
			}
		}
		if useSecretIds {
			if privateKeyPassphraseSecretId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "private_key_passphrase_secret_id")); ok {
				tmp := privateKeyPassphraseSecretId.(string)
				details.PrivateKeyPassphraseSecretId = &tmp
			}
		}
		if publicKeyFingerprint, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "public_key_fingerprint")); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "region")); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if tenancyId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "tenancy_id")); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOk(fmt.Sprintf(fieldKeyFormat, "user_id")); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		baseObject = details
	default:
		return nil, fmt.Errorf("unknown authentication_type '%v' was specified", authenticationType)
	}
	return baseObject, nil
}

func UpdateSparkOciObjectStorageConfigDetailsToMap(obj *oci_golden_gate.UpdateSparkOciObjectStorageConfigDetails) map[string]interface{} {
	result := map[string]interface{}{}
	switch v := (*obj).(type) {
	case oci_golden_gate.UpdateSparkSecretKeyOciObjectStorageConfigDetails:
		result["authentication_type"] = "CUSTOMER_SECRET_KEY"

		if v.AccessKeyId != nil {
			result["access_key_id"] = string(*v.AccessKeyId)
		}

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.SecretAccessKey != nil {
			result["secret_access_key"] = string(*v.SecretAccessKey)
		}

		if v.SecretAccessKeySecretId != nil {
			result["secret_access_key_secret_id"] = string(*v.SecretAccessKeySecretId)
		}
	case oci_golden_gate.UpdateSparkParOciObjectStorageConfigDetails:
		result["authentication_type"] = "PRE_AUTHENTICATED_REQUEST"

		if v.ParUrl != nil {
			result["par_url"] = string(*v.ParUrl)
		}

		if v.ParUrlSecretId != nil {
			result["par_url_secret_id"] = string(*v.ParUrlSecretId)
		}
	case oci_golden_gate.UpdateSparkResourcePrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "RESOURCE_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}
	case oci_golden_gate.UpdateSparkUserPrincipalOciObjectStorageConfigDetails:
		result["authentication_type"] = "USER_PRINCIPAL"

		if v.Bucket != nil {
			result["bucket"] = string(*v.Bucket)
		}

		if v.Namespace != nil {
			result["namespace"] = string(*v.Namespace)
		}

		if v.Path != nil {
			result["path"] = string(*v.Path)
		}

		if v.PrivateKeyFile != nil {
			result["private_key_file"] = string(*v.PrivateKeyFile)
		}

		if v.PrivateKeyFileSecretId != nil {
			result["private_key_file_secret_id"] = string(*v.PrivateKeyFileSecretId)
		}

		if v.PrivateKeyPassphrase != nil {
			result["private_key_passphrase"] = string(*v.PrivateKeyPassphrase)
		}

		if v.PrivateKeyPassphraseSecretId != nil {
			result["private_key_passphrase_secret_id"] = string(*v.PrivateKeyPassphraseSecretId)
		}

		if v.PublicKeyFingerprint != nil {
			result["public_key_fingerprint"] = string(*v.PublicKeyFingerprint)
		}

		if v.Region != nil {
			result["region"] = string(*v.Region)
		}

		if v.TenancyId != nil {
			result["tenancy_id"] = string(*v.TenancyId)
		}

		if v.UserId != nil {
			result["user_id"] = string(*v.UserId)
		}
	default:
		log.Printf("[WARN] Received 'authentication_type' of unknown type %v", *obj)
		return nil
	}

	return result
}

func (s *GoldenGateConnectionResourceCrud) populateTopLevelPolymorphicCreateConnectionRequest(request *oci_golden_gate.CreateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("AI_MODEL"):
		details := oci_golden_gate.CreateAiModelConnectionDetails{}
		if authDetails, ok := s.D.GetOkExists("auth_details"); ok {
			if tmpList := authDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_details", 0)
				tmp, err := s.mapToCreateAiModelAuthDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthDetails = tmp
			}
		}
		if maxInputChars, ok := s.D.GetOkExists("max_input_chars"); ok {
			tmp := maxInputChars.(int)
			details.MaxInputChars = &tmp
		}
		if modelKey, ok := s.D.GetOkExists("model_key"); ok {
			tmp := modelKey.(string)
			details.ModelKey = &tmp
		}
		if providerId, ok := s.D.GetOk("provider_id"); ok {
			tmp := providerId.(string)
			details.ProviderId = &tmp
		}
		if providerType, ok := s.D.GetOkExists("provider_type"); ok {
			details.ProviderType = oci_golden_gate.CreateAiModelConnectionDetailsProviderTypeEnum(providerType.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AiModelConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AI_PROVIDER"):
		details := oci_golden_gate.CreateAiProviderConnectionDetails{}
		if authDetails, ok := s.D.GetOkExists("auth_details"); ok {
			if tmpList := authDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_details", 0)
				tmp, err := s.mapToCreateAiModelAuthDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthDetails = tmp
			}
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AiProviderConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AMAZON_KINESIS"):
		details := oci_golden_gate.CreateAmazonKinesisConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists("secret_access_key_secret_id"); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonKinesisConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AMAZON_REDSHIFT"):
		details := oci_golden_gate.CreateAmazonRedshiftConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonRedshiftConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.CreateAmazonS3ConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists("secret_access_key_secret_id"); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AmazonS3ConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.CreateAzureDataLakeStorageConnectionDetails{}
		if accountKey, ok := s.D.GetOkExists("account_key"); ok {
			tmp := accountKey.(string)
			details.AccountKey = &tmp
		}
		if accountKeySecretId, ok := s.D.GetOkExists("account_key_secret_id"); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if accountName, ok := s.D.GetOkExists("account_name"); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.AzureDataLakeStorageConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if azureAuthorityHost, ok := s.D.GetOkExists("azure_authority_host"); ok {
			tmp := azureAuthorityHost.(string)
			details.AzureAuthorityHost = &tmp
		}
		if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
			tmp := azureTenantId.(string)
			details.AzureTenantId = &tmp
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if sasToken, ok := s.D.GetOkExists("sas_token"); ok {
			tmp := sasToken.(string)
			details.SasToken = &tmp
		}
		if sasTokenSecretId, ok := s.D.GetOkExists("sas_token_secret_id"); ok {
			tmp := sasTokenSecretId.(string)
			details.SasTokenSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AzureDataLakeStorageConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("AZURE_SYNAPSE_ANALYTICS"):
		details := oci_golden_gate.CreateAzureSynapseConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.AzureSynapseConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("DATABRICKS"):
		details := oci_golden_gate.CreateDatabricksConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.DatabricksConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if storageCredentialName, ok := s.D.GetOkExists("storage_credential_name"); ok {
			tmp := storageCredentialName.(string)
			details.StorageCredentialName = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.DatabricksConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("DATABRICKS_ENGINE"):
		details := oci_golden_gate.CreateDatabricksEngineConnectionDetails{}
		useSecretIds := false
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			useSecretIds = doesUseSecretIds.(bool)
		}
		if artifactLocation, ok := s.D.GetOkExists("artifact_location"); ok {
			if tmpList := artifactLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_location", 0)
				tmp, err := s.mapToCreateDatabricksEngineArtifactLocationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ArtifactLocation = tmp
			}
		}
		if !useSecretIds {
			if authToken, ok := s.D.GetOk("auth_token"); ok {
				tmp := authToken.(string)
				details.AuthToken = &tmp
			}
		}
		if useSecretIds {
			if authTokenSecretId, ok := s.D.GetOk("auth_token_secret_id"); ok {
				tmp := authTokenSecretId.(string)
				details.AuthTokenSecretId = &tmp
			}
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.DatabricksEngineConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if !useSecretIds {
			if clientSecret, ok := s.D.GetOk("client_secret"); ok {
				tmp := clientSecret.(string)
				details.ClientSecret = &tmp
			}
		}
		if useSecretIds {
			if clientSecretSecretId, ok := s.D.GetOk("client_secret_secret_id"); ok {
				tmp := clientSecretSecretId.(string)
				details.ClientSecretSecretId = &tmp
			}
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.DatabricksEngineConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if workspaceUrl, ok := s.D.GetOkExists("workspace_url"); ok {
			tmp := workspaceUrl.(string)
			details.WorkspaceUrl = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("DB2"):
		details := oci_golden_gate.CreateDb2ConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.Db2ConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslClientKeystash, ok := s.D.GetOkExists("ssl_client_keystash"); ok {
			tmp := sslClientKeystash.(string)
			details.SslClientKeystash = &tmp
		}
		if sslClientKeystashSecretId, ok := s.D.GetOkExists("ssl_client_keystash_secret_id"); ok {
			tmp := sslClientKeystashSecretId.(string)
			details.SslClientKeystashSecretId = &tmp
		}
		if sslClientKeystoredb, ok := s.D.GetOkExists("ssl_client_keystoredb"); ok {
			tmp := sslClientKeystoredb.(string)
			details.SslClientKeystoredb = &tmp
		}
		if sslClientKeystoredbSecretId, ok := s.D.GetOkExists("ssl_client_keystoredb_secret_id"); ok {
			tmp := sslClientKeystoredbSecretId.(string)
			details.SslClientKeystoredbSecretId = &tmp
		}
		if sslServerCertificate, ok := s.D.GetOkExists("ssl_server_certificate"); ok {
			tmp := sslServerCertificate.(string)
			details.SslServerCertificate = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.Db2ConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ELASTICSEARCH"):
		details := oci_golden_gate.CreateElasticsearchConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.ElasticsearchConnectionAuthenticationTypeEnum(authenticationType.(string))

		}
		if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
			tmp := fingerprint.(string)
			details.Fingerprint = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.ElasticsearchConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.ElasticsearchConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GENERIC"):
		details := oci_golden_gate.CreateGenericConnectionDetails{}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GenericConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOLDENGATE"):
		details := oci_golden_gate.CreateGoldenGateConnectionDetails{}
		if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
			tmp := deploymentId.(string)
			details.DeploymentId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoldenGateConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOOGLE_BIGQUERY"):
		details := oci_golden_gate.CreateGoogleBigQueryConnectionDetails{}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoogleBigQueryConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.CreateGoogleCloudStorageConnectionDetails{}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GoogleCloudStorageConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("GOOGLE_PUBSUB"):
		details := oci_golden_gate.CreateGooglePubSubConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.GooglePubSubConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("HDFS"):
		details := oci_golden_gate.CreateHdfsConnectionDetails{}
		if coreSiteXml, ok := s.D.GetOkExists("core_site_xml"); ok {
			tmp := coreSiteXml.(string)
			details.CoreSiteXml = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.HdfsConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ICEBERG"):
		details := oci_golden_gate.CreateIcebergConnectionDetails{}
		if catalog, ok := s.D.GetOkExists("catalog"); ok {
			if tmpList := catalog.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog", 0)
				tmp, err := s.mapToCreateIcebergCatalogDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Catalog = tmp
			}
		}
		if storage, ok := s.D.GetOkExists("storage"); ok {
			if tmpList := storage.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage", 0)
				tmp, err := s.mapToCreateIcebergStorageDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Storage = tmp
			}
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.IcebergConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("JAVA_MESSAGE_SERVICE"):
		details := oci_golden_gate.CreateJavaMessageServiceConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.JavaMessageServiceConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionFactory, ok := s.D.GetOkExists("connection_factory"); ok {
			tmp := connectionFactory.(string)
			details.ConnectionFactory = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if jndiConnectionFactory, ok := s.D.GetOkExists("jndi_connection_factory"); ok {
			tmp := jndiConnectionFactory.(string)
			details.JndiConnectionFactory = &tmp
		}
		if jndiInitialContextFactory, ok := s.D.GetOkExists("jndi_initial_context_factory"); ok {
			tmp := jndiInitialContextFactory.(string)
			details.JndiInitialContextFactory = &tmp
		}
		if jndiProviderUrl, ok := s.D.GetOkExists("jndi_provider_url"); ok {
			tmp := jndiProviderUrl.(string)
			details.JndiProviderUrl = &tmp
		}
		if jndiSecurityCredentials, ok := s.D.GetOkExists("jndi_security_credentials"); ok {
			tmp := jndiSecurityCredentials.(string)
			details.JndiSecurityCredentials = &tmp
		}
		if jndiSecurityCredentialsSecretId, ok := s.D.GetOkExists("jndi_security_credentials_secret_id"); ok {
			tmp := jndiSecurityCredentialsSecretId.(string)
			details.JndiSecurityCredentialsSecretId = &tmp
		}
		if jndiSecurityPrincipal, ok := s.D.GetOkExists("jndi_security_principal"); ok {
			tmp := jndiSecurityPrincipal.(string)
			details.JndiSecurityPrincipal = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.JavaMessageServiceConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if shouldUseJndi, ok := s.D.GetOkExists("should_use_jndi"); ok {
			tmp := shouldUseJndi.(bool)
			details.ShouldUseJndi = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.JavaMessageServiceConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("KAFKA"):
		details := oci_golden_gate.CreateKafkaConnectionDetails{}
		if bootstrapServers, ok := s.D.GetOk("bootstrap_servers"); ok {
			interfaces := bootstrapServers.([]interface{})
			tmp := make([]oci_golden_gate.CreateKafkaBootstrapServer, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bootstrap_servers", stateDataIndex)
				converted, err := s.mapToCreateKafkaBootstrapServer(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("bootstrap_servers") {
				details.BootstrapServers = tmp
			}
		}
		if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
			tmp := clusterId.(string)
			details.ClusterId = &tmp
		}
		if consumerProperties, ok := s.D.GetOkExists("consumer_properties"); ok {
			tmp := consumerProperties.(string)
			details.ConsumerProperties = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if producerProperties, ok := s.D.GetOkExists("producer_properties"); ok {
			tmp := producerProperties.(string)
			details.ProducerProperties = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.KafkaConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
			tmp := streamPoolId.(string)
			details.StreamPoolId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.KafkaConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("KAFKA_SCHEMA_REGISTRY"):
		details := oci_golden_gate.CreateKafkaSchemaRegistryConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.KafkaSchemaRegistryConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.KafkaSchemaRegistryConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MICROSOFT_FABRIC"):
		details := oci_golden_gate.CreateMicrosoftFabricConnectionDetails{}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MicrosoftFabricConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
			tmp := tenantId.(string)
			details.TenantId = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MICROSOFT_SQLSERVER"):
		details := oci_golden_gate.CreateMicrosoftSqlserverConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MicrosoftSqlserverConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldValidateServerCertificate, ok := s.D.GetOkExists("should_validate_server_certificate"); ok {
			tmp := shouldValidateServerCertificate.(bool)
			details.ShouldValidateServerCertificate = &tmp
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MicrosoftSqlserverConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MONGODB"):
		details := oci_golden_gate.CreateMongoDbConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MongoDbConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MongoDbConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tlsCaFile, ok := s.D.GetOkExists("tls_ca_file"); ok {
			tmp := tlsCaFile.(string)
			details.TlsCaFile = &tmp
		}
		if tlsCertificateKeyFile, ok := s.D.GetOkExists("tls_certificate_key_file"); ok {
			tmp := tlsCertificateKeyFile.(string)
			details.TlsCertificateKeyFile = &tmp
		}
		if tlsCertificateKeyFilePassword, ok := s.D.GetOkExists("tls_certificate_key_file_password"); ok {
			tmp := tlsCertificateKeyFilePassword.(string)
			details.TlsCertificateKeyFilePassword = &tmp
		}
		if tlsCertificateKeyFilePasswordSecretId, ok := s.D.GetOkExists("tls_certificate_key_file_password_secret_id"); ok {
			tmp := tlsCertificateKeyFilePasswordSecretId.(string)
			details.TlsCertificateKeyFilePasswordSecretId = &tmp
		}
		if tlsCertificateKeyFileSecretId, ok := s.D.GetOkExists("tls_certificate_key_file_secret_id"); ok {
			tmp := tlsCertificateKeyFileSecretId.(string)
			details.TlsCertificateKeyFileSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_golden_gate.CreateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslKeySecretId, ok := s.D.GetOkExists("ssl_key_secret_id"); ok {
			tmp := sslKeySecretId.(string)
			details.SslKeySecretId = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_golden_gate.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.MysqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.CreateOciObjectStorageConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OciObjectStorageConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_golden_gate.CreateOracleConnectionDetails{}
		if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
			details.AuthenticationMode = oci_golden_gate.OracleConnectionAuthenticationModeEnum(authenticationMode.(string))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
			details.SessionMode = oci_golden_gate.OracleConnectionSessionModeEnum(sessionMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if wallet, ok := s.D.GetOkExists("wallet"); ok {
			tmp := wallet.(string)
			details.Wallet = &tmp
		}
		if walletSecretId, ok := s.D.GetOkExists("wallet_secret_id"); ok {
			tmp := walletSecretId.(string)
			details.WalletSecretId = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE_AI_DATA_PLATFORM"):
		details := oci_golden_gate.CreateOracleAiDataPlatformConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleAiDataPlatformConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE_AI_DATA_PLATFORM_ENGINE"):
		details := oci_golden_gate.CreateOracleAiDataPlatformEngineConnectionDetails{}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleAiDataPlatformEngineConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if workspaceUrl, ok := s.D.GetOkExists("workspace_url"); ok {
			tmp := workspaceUrl.(string)
			details.WorkspaceUrl = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("ORACLE_NOSQL"):
		details := oci_golden_gate.CreateOracleNosqlConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.OracleNosqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_golden_gate.CreatePostgresqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.PostgresqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslKeySecretId, ok := s.D.GetOkExists("ssl_key_secret_id"); ok {
			tmp := sslKeySecretId.(string)
			details.SslKeySecretId = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_golden_gate.PostgresqlConnectionSslModeEnum(sslMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.PostgresqlConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("REDIS"):
		details := oci_golden_gate.CreateRedisConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.RedisConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
			tmp := redisClusterId.(string)
			details.RedisClusterId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.RedisConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.RedisConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("SNOWFLAKE"):
		details := oci_golden_gate.CreateSnowflakeConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SnowflakeConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.SnowflakeConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	case strings.ToLower("SPARK"):
		details := oci_golden_gate.CreateSparkConnectionDetails{}
		useSecretIds := false
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			useSecretIds = doesUseSecretIds.(bool)
		}
		if artifactLocation, ok := s.D.GetOkExists("artifact_location"); ok {
			if tmpList := artifactLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_location", 0)
				tmp, err := s.mapToCreateSparkArtifactLocationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ArtifactLocation = tmp
			}
		}
		if !useSecretIds {
			if authSecret, ok := s.D.GetOk("auth_secret"); ok {
				tmp := authSecret.(string)
				details.AuthSecret = &tmp
			}
		}
		if useSecretIds {
			if authSecretSecretId, ok := s.D.GetOk("auth_secret_secret_id"); ok {
				tmp := authSecretSecretId.(string)
				details.AuthSecretSecretId = &tmp
			}
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SparkConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clusterEndpoint, ok := s.D.GetOkExists("cluster_endpoint"); ok {
			tmp := clusterEndpoint.(string)
			details.ClusterEndpoint = &tmp
		}
		if !useSecretIds {
			if keyStore, ok := s.D.GetOk("key_store"); ok {
				tmp := keyStore.(string)
				details.KeyStore = &tmp
			}
			if keyStorePassword, ok := s.D.GetOk("key_store_password"); ok {
				tmp := keyStorePassword.(string)
				details.KeyStorePassword = &tmp
			}
		}
		if useSecretIds {
			if keyStorePasswordSecretId, ok := s.D.GetOk("key_store_password_secret_id"); ok {
				tmp := keyStorePasswordSecretId.(string)
				details.KeyStorePasswordSecretId = &tmp
			}
			if keyStoreSecretId, ok := s.D.GetOk("key_store_secret_id"); ok {
				tmp := keyStoreSecretId.(string)
				details.KeyStoreSecretId = &tmp
			}
		}
		if keyStoreType, ok := s.D.GetOkExists("key_store_type"); ok {
			details.KeyStoreType = oci_golden_gate.SparkConnectionKeyStoreTypeEnum(keyStoreType.(string))
		}
		if !useSecretIds {
			if restAuthSecretKey, ok := s.D.GetOk("rest_auth_secret_key"); ok {
				tmp := restAuthSecretKey.(string)
				details.RestAuthSecretKey = &tmp
			}
		}
		if useSecretIds {
			if restAuthSecretKeySecretId, ok := s.D.GetOk("rest_auth_secret_key_secret_id"); ok {
				tmp := restAuthSecretKeySecretId.(string)
				details.RestAuthSecretKeySecretId = &tmp
			}
		}
		if restEndpoint, ok := s.D.GetOkExists("rest_endpoint"); ok {
			tmp := restEndpoint.(string)
			details.RestEndpoint = &tmp
		}
		if rpcCipher, ok := s.D.GetOkExists("rpc_cipher"); ok {
			details.RpcCipher = oci_golden_gate.SparkConnectionRpcCipherEnum(rpcCipher.(string))
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.SparkConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if !useSecretIds {
			if sslKeyPassword, ok := s.D.GetOk("ssl_key_password"); ok {
				tmp := sslKeyPassword.(string)
				details.SslKeyPassword = &tmp
			}
		}
		if useSecretIds {
			if sslKeyPasswordSecretId, ok := s.D.GetOk("ssl_key_password_secret_id"); ok {
				tmp := sslKeyPasswordSecretId.(string)
				details.SslKeyPasswordSecretId = &tmp
			}
		}
		if submitMode, ok := s.D.GetOkExists("submit_mode"); ok {
			details.SubmitMode = oci_golden_gate.SparkConnectionSubmitModeEnum(submitMode.(string))
		}
		if technologyType, ok := s.D.GetOkExists("technology_type"); ok {
			details.TechnologyType = oci_golden_gate.SparkConnectionTechnologyTypeEnum(technologyType.(string))
		}
		if tlsProtocol, ok := s.D.GetOkExists("tls_protocol"); ok {
			details.TlsProtocol = oci_golden_gate.SparkConnectionTlsProtocolEnum(tlsProtocol.(string))
		}
		if !useSecretIds {
			if trustStore, ok := s.D.GetOk("trust_store"); ok {
				tmp := trustStore.(string)
				details.TrustStore = &tmp
			}
			if trustStorePassword, ok := s.D.GetOk("trust_store_password"); ok {
				tmp := trustStorePassword.(string)
				details.TrustStorePassword = &tmp
			}
		}
		if useSecretIds {
			if trustStorePasswordSecretId, ok := s.D.GetOk("trust_store_password_secret_id"); ok {
				tmp := trustStorePasswordSecretId.(string)
				details.TrustStorePasswordSecretId = &tmp
			}
			if trustStoreSecretId, ok := s.D.GetOk("trust_store_secret_id"); ok {
				tmp := trustStoreSecretId.(string)
				details.TrustStoreSecretId = &tmp
			}
		}
		if trustStoreType, ok := s.D.GetOkExists("trust_store_type"); ok {
			details.TrustStoreType = oci_golden_gate.SparkConnectionTrustStoreTypeEnum(trustStoreType.(string))
		}
		if clusterPlacementGroupId, ok := s.D.GetOkExists("cluster_placement_group_id"); ok {
			tmp := clusterPlacementGroupId.(string)
			details.ClusterPlacementGroupId = &tmp
		}
		if compartmentId, ok := s.D.GetOkExists("compartment_id"); ok {
			tmp := compartmentId.(string)
			details.CompartmentId = &tmp
		}
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if locks, ok := s.D.GetOkExists("locks"); ok {
			interfaces := locks.([]interface{})
			tmp := make([]oci_golden_gate.AddResourceLockDetails, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "locks", stateDataIndex)
				converted, err := s.mapToAddResourceLockDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("locks") {
				details.Locks = tmp
			}
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if subscriptionId, ok := s.D.GetOkExists("subscription_id"); ok {
			tmp := subscriptionId.(string)
			details.SubscriptionId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.CreateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
}

func (s *GoldenGateConnectionResourceCrud) populateTopLevelPolymorphicUpdateConnectionRequest(request *oci_golden_gate.UpdateConnectionRequest) error {
	//discriminator
	connectionTypeRaw, ok := s.D.GetOkExists("connection_type")
	var connectionType string
	if ok {
		connectionType = connectionTypeRaw.(string)
	} else {
		connectionType = "" // default value
	}
	switch strings.ToLower(connectionType) {
	case strings.ToLower("AI_MODEL"):
		details := oci_golden_gate.UpdateAiModelConnectionDetails{}
		if authDetails, ok := s.D.GetOkExists("auth_details"); ok {
			if tmpList := authDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_details", 0)
				tmp, err := s.mapToUpdateAiModelAuthDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthDetails = tmp
			}
		}
		if maxInputChars, ok := s.D.GetOkExists("max_input_chars"); ok {
			tmp := maxInputChars.(int)
			details.MaxInputChars = &tmp
		}
		if modelKey, ok := s.D.GetOkExists("model_key"); ok {
			tmp := modelKey.(string)
			details.ModelKey = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AI_PROVIDER"):
		details := oci_golden_gate.UpdateAiProviderConnectionDetails{}
		if authDetails, ok := s.D.GetOkExists("auth_details"); ok {
			if tmpList := authDetails.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "auth_details", 0)
				tmp, err := s.mapToUpdateAiModelAuthDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.AuthDetails = tmp
			}
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AMAZON_KINESIS"):
		details := oci_golden_gate.UpdateAmazonKinesisConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists("secret_access_key_secret_id"); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AMAZON_REDSHIFT"):
		details := oci_golden_gate.UpdateAmazonRedshiftConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AMAZON_S3"):
		details := oci_golden_gate.UpdateAmazonS3ConnectionDetails{}
		if accessKeyId, ok := s.D.GetOkExists("access_key_id"); ok {
			tmp := accessKeyId.(string)
			details.AccessKeyId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if secretAccessKey, ok := s.D.GetOkExists("secret_access_key"); ok {
			tmp := secretAccessKey.(string)
			details.SecretAccessKey = &tmp
		}
		if secretAccessKeySecretId, ok := s.D.GetOkExists("secret_access_key_secret_id"); ok {
			tmp := secretAccessKeySecretId.(string)
			details.SecretAccessKeySecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AZURE_DATA_LAKE_STORAGE"):
		details := oci_golden_gate.UpdateAzureDataLakeStorageConnectionDetails{}
		if accountKey, ok := s.D.GetOkExists("account_key"); ok {
			tmp := accountKey.(string)
			details.AccountKey = &tmp
		}
		if accountKeySecretId, ok := s.D.GetOkExists("account_key_secret_id"); ok {
			tmp := accountKeySecretId.(string)
			details.AccountKeySecretId = &tmp
		}
		if accountName, ok := s.D.GetOkExists("account_name"); ok {
			tmp := accountName.(string)
			details.AccountName = &tmp
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.AzureDataLakeStorageConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if azureAuthorityHost, ok := s.D.GetOkExists("azure_authority_host"); ok {
			tmp := azureAuthorityHost.(string)
			details.AzureAuthorityHost = &tmp
		}
		if azureTenantId, ok := s.D.GetOkExists("azure_tenant_id"); ok {
			tmp := azureTenantId.(string)
			details.AzureTenantId = &tmp
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if sasToken, ok := s.D.GetOkExists("sas_token"); ok {
			tmp := sasToken.(string)
			details.SasToken = &tmp
		}
		if sasTokenSecretId, ok := s.D.GetOkExists("sas_token_secret_id"); ok {
			tmp := sasTokenSecretId.(string)
			details.SasTokenSecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("AZURE_SYNAPSE_ANALYTICS"):
		details := oci_golden_gate.UpdateAzureSynapseConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("DATABRICKS"):
		details := oci_golden_gate.UpdateDatabricksConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.DatabricksConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if storageCredentialName, ok := s.D.GetOkExists("storage_credential_name"); ok {
			tmp := storageCredentialName.(string)
			details.StorageCredentialName = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("DATABRICKS_ENGINE"):
		details := oci_golden_gate.UpdateDatabricksEngineConnectionDetails{}
		useSecretIds := false
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			useSecretIds = doesUseSecretIds.(bool)
		}
		if artifactLocation, ok := s.D.GetOkExists("artifact_location"); ok {
			if tmpList := artifactLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_location", 0)
				tmp, err := s.mapToUpdateDatabricksEngineArtifactLocationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ArtifactLocation = tmp
			}
		}
		if !useSecretIds {
			if authToken, ok := s.D.GetOk("auth_token"); ok {
				tmp := authToken.(string)
				details.AuthToken = &tmp
			}
		}
		if useSecretIds {
			if authTokenSecretId, ok := s.D.GetOk("auth_token_secret_id"); ok {
				tmp := authTokenSecretId.(string)
				details.AuthTokenSecretId = &tmp
			}
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.DatabricksEngineConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if !useSecretIds {
			if clientSecret, ok := s.D.GetOk("client_secret"); ok {
				tmp := clientSecret.(string)
				details.ClientSecret = &tmp
			}
		}
		if useSecretIds {
			if clientSecretSecretId, ok := s.D.GetOk("client_secret_secret_id"); ok {
				tmp := clientSecretSecretId.(string)
				details.ClientSecretSecretId = &tmp
			}
		}
		if workspaceUrl, ok := s.D.GetOkExists("workspace_url"); ok {
			tmp := workspaceUrl.(string)
			details.WorkspaceUrl = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("DB2"):
		details := oci_golden_gate.UpdateDb2ConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.Db2ConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslClientKeystash, ok := s.D.GetOkExists("ssl_client_keystash"); ok {
			tmp := sslClientKeystash.(string)
			details.SslClientKeystash = &tmp
		}
		if sslClientKeystashSecretId, ok := s.D.GetOkExists("ssl_client_keystash_secret_id"); ok {
			tmp := sslClientKeystashSecretId.(string)
			details.SslClientKeystashSecretId = &tmp
		}
		if sslClientKeystoredb, ok := s.D.GetOkExists("ssl_client_keystoredb"); ok {
			tmp := sslClientKeystoredb.(string)
			details.SslClientKeystoredb = &tmp
		}
		if sslClientKeystoredbSecretId, ok := s.D.GetOkExists("ssl_client_keystoredb_secret_id"); ok {
			tmp := sslClientKeystoredbSecretId.(string)
			details.SslClientKeystoredbSecretId = &tmp
		}
		if sslServerCertificate, ok := s.D.GetOkExists("ssl_server_certificate"); ok {
			tmp := sslServerCertificate.(string)
			details.SslServerCertificate = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ELASTICSEARCH"):
		details := oci_golden_gate.UpdateElasticsearchConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.ElasticsearchConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if fingerprint, ok := s.D.GetOkExists("fingerprint"); ok {
			tmp := fingerprint.(string)
			details.Fingerprint = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.ElasticsearchConnectionSecurityProtocolEnum(securityProtocol.(string))

		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GENERIC"):
		details := oci_golden_gate.UpdateGenericConnectionDetails{}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOLDENGATE"):
		details := oci_golden_gate.UpdateGoldenGateConnectionDetails{}
		if deploymentId, ok := s.D.GetOkExists("deployment_id"); ok {
			tmp := deploymentId.(string)
			details.DeploymentId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOOGLE_BIGQUERY"):
		details := oci_golden_gate.UpdateGoogleBigQueryConnectionDetails{}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOOGLE_CLOUD_STORAGE"):
		details := oci_golden_gate.UpdateGoogleCloudStorageConnectionDetails{}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("GOOGLE_PUBSUB"):
		details := oci_golden_gate.UpdateGooglePubSubConnectionDetails{}
		if serviceAccountKeyFile, ok := s.D.GetOkExists("service_account_key_file"); ok {
			tmp := serviceAccountKeyFile.(string)
			details.ServiceAccountKeyFile = &tmp
		}
		if serviceAccountKeyFileSecretId, ok := s.D.GetOkExists("service_account_key_file_secret_id"); ok {
			tmp := serviceAccountKeyFileSecretId.(string)
			details.ServiceAccountKeyFileSecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("HDFS"):
		details := oci_golden_gate.UpdateHdfsConnectionDetails{}
		if coreSiteXml, ok := s.D.GetOkExists("core_site_xml"); ok {
			tmp := coreSiteXml.(string)
			details.CoreSiteXml = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ICEBERG"):
		details := oci_golden_gate.UpdateIcebergConnectionDetails{}
		if catalog, ok := s.D.GetOkExists("catalog"); ok {
			if tmpList := catalog.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "catalog", 0)
				tmp, err := s.mapToUpdateIcebergCatalogDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Catalog = tmp
			}
		}
		if storage, ok := s.D.GetOkExists("storage"); ok {
			if tmpList := storage.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "storage", 0)
				tmp, err := s.mapToUpdateIcebergStorageDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.Storage = tmp
			}
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("JAVA_MESSAGE_SERVICE"):
		details := oci_golden_gate.UpdateJavaMessageServiceConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.JavaMessageServiceConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionFactory, ok := s.D.GetOkExists("connection_factory"); ok {
			tmp := connectionFactory.(string)
			details.ConnectionFactory = &tmp
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if jndiConnectionFactory, ok := s.D.GetOkExists("jndi_connection_factory"); ok {
			tmp := jndiConnectionFactory.(string)
			details.JndiConnectionFactory = &tmp
		}
		if jndiInitialContextFactory, ok := s.D.GetOkExists("jndi_initial_context_factory"); ok {
			tmp := jndiInitialContextFactory.(string)
			details.JndiInitialContextFactory = &tmp
		}
		if jndiProviderUrl, ok := s.D.GetOkExists("jndi_provider_url"); ok {
			tmp := jndiProviderUrl.(string)
			details.JndiProviderUrl = &tmp
		}
		if jndiSecurityCredentials, ok := s.D.GetOkExists("jndi_security_credentials"); ok {
			tmp := jndiSecurityCredentials.(string)
			details.JndiSecurityCredentials = &tmp
		}
		if jndiSecurityCredentialsSecretId, ok := s.D.GetOkExists("jndi_security_credentials_secret_id"); ok {
			tmp := jndiSecurityCredentialsSecretId.(string)
			details.JndiSecurityCredentialsSecretId = &tmp
		}
		if jndiSecurityPrincipal, ok := s.D.GetOkExists("jndi_security_principal"); ok {
			tmp := jndiSecurityPrincipal.(string)
			details.JndiSecurityPrincipal = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.JavaMessageServiceConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldUseJndi, ok := s.D.GetOkExists("should_use_jndi"); ok {
			tmp := shouldUseJndi.(bool)
			details.ShouldUseJndi = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("KAFKA"):
		details := oci_golden_gate.UpdateKafkaConnectionDetails{}
		if bootstrapServers, ok := s.D.GetOk("bootstrap_servers"); ok {
			interfaces := bootstrapServers.([]interface{})
			tmp := make([]oci_golden_gate.KafkaBootstrapServer, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "bootstrap_servers", stateDataIndex)
				converted, err := s.mapToKafkaBootstrapServer(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("bootstrap_servers") {
				details.BootstrapServers = tmp
			}
		}
		if clusterId, ok := s.D.GetOkExists("cluster_id"); ok {
			tmp := clusterId.(string)
			details.ClusterId = &tmp
		}
		if consumerProperties, ok := s.D.GetOkExists("consumer_properties"); ok {
			tmp := consumerProperties.(string)
			details.ConsumerProperties = &tmp
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if producerProperties, ok := s.D.GetOkExists("producer_properties"); ok {
			tmp := producerProperties.(string)
			details.ProducerProperties = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.KafkaConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if streamPoolId, ok := s.D.GetOkExists("stream_pool_id"); ok {
			tmp := streamPoolId.(string)
			details.StreamPoolId = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("KAFKA_SCHEMA_REGISTRY"):
		details := oci_golden_gate.UpdateKafkaSchemaRegistryConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.KafkaSchemaRegistryConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if sslKeyPassword, ok := s.D.GetOkExists("ssl_key_password"); ok {
			tmp := sslKeyPassword.(string)
			details.SslKeyPassword = &tmp
		}
		if sslKeyPasswordSecretId, ok := s.D.GetOkExists("ssl_key_password_secret_id"); ok {
			tmp := sslKeyPasswordSecretId.(string)
			details.SslKeyPasswordSecretId = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if url, ok := s.D.GetOkExists("url"); ok {
			tmp := url.(string)
			details.Url = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MICROSOFT_FABRIC"):
		details := oci_golden_gate.UpdateMicrosoftFabricConnectionDetails{}
		if clientId, ok := s.D.GetOkExists("client_id"); ok {
			tmp := clientId.(string)
			details.ClientId = &tmp
		}
		if clientSecret, ok := s.D.GetOkExists("client_secret"); ok {
			tmp := clientSecret.(string)
			details.ClientSecret = &tmp
		}
		if clientSecretSecretId, ok := s.D.GetOkExists("client_secret_secret_id"); ok {
			tmp := clientSecretSecretId.(string)
			details.ClientSecretSecretId = &tmp
		}
		if endpoint, ok := s.D.GetOkExists("endpoint"); ok {
			tmp := endpoint.(string)
			details.Endpoint = &tmp
		}
		if tenantId, ok := s.D.GetOkExists("tenant_id"); ok {
			tmp := tenantId.(string)
			details.TenantId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MICROSOFT_SQLSERVER"):
		details := oci_golden_gate.UpdateMicrosoftSqlserverConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MicrosoftSqlserverConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if shouldValidateServerCertificate, ok := s.D.GetOkExists("should_validate_server_certificate"); ok {
			tmp := shouldValidateServerCertificate.(bool)
			details.ShouldValidateServerCertificate = &tmp
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MONGODB"):
		details := oci_golden_gate.UpdateMongoDbConnectionDetails{}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MongoDbConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if tlsCaFile, ok := s.D.GetOkExists("tls_ca_file"); ok {
			tmp := tlsCaFile.(string)
			details.TlsCaFile = &tmp
		}
		if tlsCertificateKeyFile, ok := s.D.GetOkExists("tls_certificate_key_file"); ok {
			tmp := tlsCertificateKeyFile.(string)
			details.TlsCertificateKeyFile = &tmp
		}
		if tlsCertificateKeyFilePassword, ok := s.D.GetOkExists("tls_certificate_key_file_password"); ok {
			tmp := tlsCertificateKeyFilePassword.(string)
			details.TlsCertificateKeyFilePassword = &tmp
		}
		if tlsCertificateKeyFilePasswordSecretId, ok := s.D.GetOkExists("tls_certificate_key_file_password_secret_id"); ok {
			tmp := tlsCertificateKeyFilePasswordSecretId.(string)
			details.TlsCertificateKeyFilePasswordSecretId = &tmp
		}
		if tlsCertificateKeyFileSecretId, ok := s.D.GetOkExists("tls_certificate_key_file_secret_id"); ok {
			tmp := tlsCertificateKeyFileSecretId.(string)
			details.TlsCertificateKeyFileSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("MYSQL"):
		details := oci_golden_gate.UpdateMysqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.MysqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslKeySecretId, ok := s.D.GetOkExists("ssl_key_secret_id"); ok {
			tmp := sslKeySecretId.(string)
			details.SslKeySecretId = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_golden_gate.MysqlConnectionSslModeEnum(sslMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("OCI_OBJECT_STORAGE"):
		details := oci_golden_gate.UpdateOciObjectStorageConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE"):
		details := oci_golden_gate.UpdateOracleConnectionDetails{}
		if authenticationMode, ok := s.D.GetOkExists("authentication_mode"); ok {
			details.AuthenticationMode = oci_golden_gate.OracleConnectionAuthenticationModeEnum(authenticationMode.(string))
		}
		if connectionString, ok := s.D.GetOkExists("connection_string"); ok {
			tmp := connectionString.(string)
			details.ConnectionString = &tmp
		}
		if databaseId, ok := s.D.GetOkExists("database_id"); ok {
			tmp := databaseId.(string)
			details.DatabaseId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if sessionMode, ok := s.D.GetOkExists("session_mode"); ok {
			details.SessionMode = oci_golden_gate.OracleConnectionSessionModeEnum(sessionMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		if wallet, ok := s.D.GetOkExists("wallet"); ok {
			tmp := wallet.(string)
			details.Wallet = &tmp
		}
		if walletSecretId, ok := s.D.GetOkExists("wallet_secret_id"); ok {
			tmp := walletSecretId.(string)
			details.WalletSecretId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE_AI_DATA_PLATFORM"):
		details := oci_golden_gate.UpdateOracleAiDataPlatformConnectionDetails{}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE_AI_DATA_PLATFORM_ENGINE"):
		details := oci_golden_gate.UpdateOracleAiDataPlatformEngineConnectionDetails{}
		if workspaceUrl, ok := s.D.GetOkExists("workspace_url"); ok {
			tmp := workspaceUrl.(string)
			details.WorkspaceUrl = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("ORACLE_NOSQL"):
		details := oci_golden_gate.UpdateOracleNosqlConnectionDetails{}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if publicKeyFingerprint, ok := s.D.GetOkExists("public_key_fingerprint"); ok {
			tmp := publicKeyFingerprint.(string)
			details.PublicKeyFingerprint = &tmp
		}
		if region, ok := s.D.GetOkExists("region"); ok {
			tmp := region.(string)
			details.Region = &tmp
		}
		if shouldUseResourcePrincipal, ok := s.D.GetOkExists("should_use_resource_principal"); ok {
			tmp := shouldUseResourcePrincipal.(bool)
			details.ShouldUseResourcePrincipal = &tmp
		}
		if tenancyId, ok := s.D.GetOkExists("tenancy_id"); ok {
			tmp := tenancyId.(string)
			details.TenancyId = &tmp
		}
		if userId, ok := s.D.GetOkExists("user_id"); ok {
			tmp := userId.(string)
			details.UserId = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("POSTGRESQL"):
		details := oci_golden_gate.UpdatePostgresqlConnectionDetails{}
		if additionalAttributes, ok := s.D.GetOk("additional_attributes"); ok {
			interfaces := additionalAttributes.([]interface{})
			tmp := make([]oci_golden_gate.NameValuePair, len(interfaces))
			for i := range interfaces {
				stateDataIndex := i
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "additional_attributes", stateDataIndex)
				converted, err := s.mapToNameValuePair(fieldKeyFormat)
				if err != nil {
					return err
				}
				tmp[i] = converted
			}
			if len(tmp) != 0 || s.D.HasChange("additional_attributes") {
				details.AdditionalAttributes = tmp
			}
		}
		if databaseName, ok := s.D.GetOkExists("database_name"); ok {
			tmp := databaseName.(string)
			details.DatabaseName = &tmp
		}
		if dbSystemId, ok := s.D.GetOkExists("db_system_id"); ok {
			tmp := dbSystemId.(string)
			details.DbSystemId = &tmp
		}
		if host, ok := s.D.GetOkExists("host"); ok {
			tmp := host.(string)
			details.Host = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if port, ok := s.D.GetOkExists("port"); ok {
			tmp := port.(int)
			details.Port = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.PostgresqlConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if sslCa, ok := s.D.GetOkExists("ssl_ca"); ok {
			tmp := sslCa.(string)
			details.SslCa = &tmp
		}
		if sslCert, ok := s.D.GetOkExists("ssl_cert"); ok {
			tmp := sslCert.(string)
			details.SslCert = &tmp
		}
		if sslCrl, ok := s.D.GetOkExists("ssl_crl"); ok {
			tmp := sslCrl.(string)
			details.SslCrl = &tmp
		}
		if sslKey, ok := s.D.GetOkExists("ssl_key"); ok {
			tmp := sslKey.(string)
			details.SslKey = &tmp
		}
		if sslKeySecretId, ok := s.D.GetOkExists("ssl_key_secret_id"); ok {
			tmp := sslKeySecretId.(string)
			details.SslKeySecretId = &tmp
		}
		if sslMode, ok := s.D.GetOkExists("ssl_mode"); ok {
			details.SslMode = oci_golden_gate.PostgresqlConnectionSslModeEnum(sslMode.(string))
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("REDIS"):
		details := oci_golden_gate.UpdateRedisConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.RedisConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if keyStore, ok := s.D.GetOkExists("key_store"); ok {
			tmp := keyStore.(string)
			details.KeyStore = &tmp
		}
		if keyStorePassword, ok := s.D.GetOkExists("key_store_password"); ok {
			tmp := keyStorePassword.(string)
			details.KeyStorePassword = &tmp
		}
		if keyStorePasswordSecretId, ok := s.D.GetOkExists("key_store_password_secret_id"); ok {
			tmp := keyStorePasswordSecretId.(string)
			details.KeyStorePasswordSecretId = &tmp
		}
		if keyStoreSecretId, ok := s.D.GetOkExists("key_store_secret_id"); ok {
			tmp := keyStoreSecretId.(string)
			details.KeyStoreSecretId = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if redisClusterId, ok := s.D.GetOkExists("redis_cluster_id"); ok {
			tmp := redisClusterId.(string)
			details.RedisClusterId = &tmp
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.RedisConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if servers, ok := s.D.GetOkExists("servers"); ok {
			tmp := servers.(string)
			details.Servers = &tmp
		}
		if trustStore, ok := s.D.GetOkExists("trust_store"); ok {
			tmp := trustStore.(string)
			details.TrustStore = &tmp
		}
		if trustStorePassword, ok := s.D.GetOkExists("trust_store_password"); ok {
			tmp := trustStorePassword.(string)
			details.TrustStorePassword = &tmp
		}
		if trustStorePasswordSecretId, ok := s.D.GetOkExists("trust_store_password_secret_id"); ok {
			tmp := trustStorePasswordSecretId.(string)
			details.TrustStorePasswordSecretId = &tmp
		}
		if trustStoreSecretId, ok := s.D.GetOkExists("trust_store_secret_id"); ok {
			tmp := trustStoreSecretId.(string)
			details.TrustStoreSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("SNOWFLAKE"):
		details := oci_golden_gate.UpdateSnowflakeConnectionDetails{}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SnowflakeConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if connectionUrl, ok := s.D.GetOkExists("connection_url"); ok {
			tmp := connectionUrl.(string)
			details.ConnectionUrl = &tmp
		}
		if password, ok := s.D.GetOkExists("password"); ok {
			tmp := password.(string)
			details.Password = &tmp
		}
		if passwordSecretId, ok := s.D.GetOkExists("password_secret_id"); ok {
			tmp := passwordSecretId.(string)
			details.PasswordSecretId = &tmp
		}
		if privateKeyFile, ok := s.D.GetOkExists("private_key_file"); ok {
			tmp := privateKeyFile.(string)
			details.PrivateKeyFile = &tmp
		}
		if privateKeyFileSecretId, ok := s.D.GetOkExists("private_key_file_secret_id"); ok {
			tmp := privateKeyFileSecretId.(string)
			details.PrivateKeyFileSecretId = &tmp
		}
		if privateKeyPassphrase, ok := s.D.GetOkExists("private_key_passphrase"); ok {
			tmp := privateKeyPassphrase.(string)
			details.PrivateKeyPassphrase = &tmp
		}
		if privateKeyPassphraseSecretId, ok := s.D.GetOkExists("private_key_passphrase_secret_id"); ok {
			tmp := privateKeyPassphraseSecretId.(string)
			details.PrivateKeyPassphraseSecretId = &tmp
		}
		if username, ok := s.D.GetOkExists("username"); ok {
			tmp := username.(string)
			details.Username = &tmp
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	case strings.ToLower("SPARK"):
		details := oci_golden_gate.UpdateSparkConnectionDetails{}
		useSecretIds := false
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			useSecretIds = doesUseSecretIds.(bool)
		}
		if artifactLocation, ok := s.D.GetOkExists("artifact_location"); ok {
			if tmpList := artifactLocation.([]interface{}); len(tmpList) > 0 {
				fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "artifact_location", 0)
				tmp, err := s.mapToUpdateSparkArtifactLocationDetails(fieldKeyFormat)
				if err != nil {
					return err
				}
				details.ArtifactLocation = tmp
			}
		}
		if !useSecretIds {
			if authSecret, ok := s.D.GetOk("auth_secret"); ok {
				tmp := authSecret.(string)
				details.AuthSecret = &tmp
			}
		}
		if useSecretIds {
			if authSecretSecretId, ok := s.D.GetOk("auth_secret_secret_id"); ok {
				tmp := authSecretSecretId.(string)
				details.AuthSecretSecretId = &tmp
			}
		}
		if authenticationType, ok := s.D.GetOkExists("authentication_type"); ok {
			details.AuthenticationType = oci_golden_gate.SparkConnectionAuthenticationTypeEnum(authenticationType.(string))
		}
		if clusterEndpoint, ok := s.D.GetOkExists("cluster_endpoint"); ok {
			tmp := clusterEndpoint.(string)
			details.ClusterEndpoint = &tmp
		}
		if !useSecretIds {
			if keyStore, ok := s.D.GetOk("key_store"); ok {
				tmp := keyStore.(string)
				details.KeyStore = &tmp
			}
			if keyStorePassword, ok := s.D.GetOk("key_store_password"); ok {
				tmp := keyStorePassword.(string)
				details.KeyStorePassword = &tmp
			}
		}
		if useSecretIds {
			if keyStorePasswordSecretId, ok := s.D.GetOk("key_store_password_secret_id"); ok {
				tmp := keyStorePasswordSecretId.(string)
				details.KeyStorePasswordSecretId = &tmp
			}
			if keyStoreSecretId, ok := s.D.GetOk("key_store_secret_id"); ok {
				tmp := keyStoreSecretId.(string)
				details.KeyStoreSecretId = &tmp
			}
		}
		if keyStoreType, ok := s.D.GetOkExists("key_store_type"); ok {
			details.KeyStoreType = oci_golden_gate.SparkConnectionKeyStoreTypeEnum(keyStoreType.(string))
		}
		if !useSecretIds {
			if restAuthSecretKey, ok := s.D.GetOk("rest_auth_secret_key"); ok {
				tmp := restAuthSecretKey.(string)
				details.RestAuthSecretKey = &tmp
			}
		}
		if useSecretIds {
			if restAuthSecretKeySecretId, ok := s.D.GetOk("rest_auth_secret_key_secret_id"); ok {
				tmp := restAuthSecretKeySecretId.(string)
				details.RestAuthSecretKeySecretId = &tmp
			}
		}
		if restEndpoint, ok := s.D.GetOkExists("rest_endpoint"); ok {
			tmp := restEndpoint.(string)
			details.RestEndpoint = &tmp
		}
		if rpcCipher, ok := s.D.GetOkExists("rpc_cipher"); ok {
			details.RpcCipher = oci_golden_gate.SparkConnectionRpcCipherEnum(rpcCipher.(string))
		}
		if securityProtocol, ok := s.D.GetOkExists("security_protocol"); ok {
			details.SecurityProtocol = oci_golden_gate.SparkConnectionSecurityProtocolEnum(securityProtocol.(string))
		}
		if !useSecretIds {
			if sslKeyPassword, ok := s.D.GetOk("ssl_key_password"); ok {
				tmp := sslKeyPassword.(string)
				details.SslKeyPassword = &tmp
			}
		}
		if useSecretIds {
			if sslKeyPasswordSecretId, ok := s.D.GetOk("ssl_key_password_secret_id"); ok {
				tmp := sslKeyPasswordSecretId.(string)
				details.SslKeyPasswordSecretId = &tmp
			}
		}
		if submitMode, ok := s.D.GetOkExists("submit_mode"); ok {
			details.SubmitMode = oci_golden_gate.SparkConnectionSubmitModeEnum(submitMode.(string))
		}
		if tlsProtocol, ok := s.D.GetOkExists("tls_protocol"); ok {
			details.TlsProtocol = oci_golden_gate.SparkConnectionTlsProtocolEnum(tlsProtocol.(string))
		}
		if !useSecretIds {
			if trustStore, ok := s.D.GetOk("trust_store"); ok {
				tmp := trustStore.(string)
				details.TrustStore = &tmp
			}
			if trustStorePassword, ok := s.D.GetOk("trust_store_password"); ok {
				tmp := trustStorePassword.(string)
				details.TrustStorePassword = &tmp
			}
		}
		if useSecretIds {
			if trustStorePasswordSecretId, ok := s.D.GetOk("trust_store_password_secret_id"); ok {
				tmp := trustStorePasswordSecretId.(string)
				details.TrustStorePasswordSecretId = &tmp
			}
			if trustStoreSecretId, ok := s.D.GetOk("trust_store_secret_id"); ok {
				tmp := trustStoreSecretId.(string)
				details.TrustStoreSecretId = &tmp
			}
		}
		if trustStoreType, ok := s.D.GetOkExists("trust_store_type"); ok {
			details.TrustStoreType = oci_golden_gate.SparkConnectionTrustStoreTypeEnum(trustStoreType.(string))
		}
		tmp := s.D.Id()
		request.ConnectionId = &tmp
		if definedTags, ok := s.D.GetOkExists("defined_tags"); ok {
			convertedDefinedTags, err := tfresource.MapToDefinedTags(definedTags.(map[string]interface{}))
			if err != nil {
				return err
			}
			details.DefinedTags = convertedDefinedTags
		}
		if description, ok := s.D.GetOkExists("description"); ok {
			tmp := description.(string)
			details.Description = &tmp
		}
		if displayName, ok := s.D.GetOkExists("display_name"); ok {
			tmp := displayName.(string)
			details.DisplayName = &tmp
		}
		if doesUseSecretIds, ok := s.D.GetOkExists("does_use_secret_ids"); ok {
			tmp := doesUseSecretIds.(bool)
			details.DoesUseSecretIds = &tmp
		}
		if freeformTags, ok := s.D.GetOkExists("freeform_tags"); ok {
			details.FreeformTags = tfresource.ObjectMapToStringMap(freeformTags.(map[string]interface{}))
		}
		if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
			tmp := isLockOverride.(bool)
			request.IsLockOverride = &tmp
		}
		if keyId, ok := s.D.GetOkExists("key_id"); ok {
			tmp := keyId.(string)
			details.KeyId = &tmp
		}
		if nsgIds, ok := s.D.GetOkExists("nsg_ids"); ok {
			set := nsgIds.(*schema.Set)
			interfaces := set.List()
			tmp := make([]string, len(interfaces))
			for i := range interfaces {
				if interfaces[i] != nil {
					tmp[i] = interfaces[i].(string)
				}
			}
			if len(tmp) != 0 || s.D.HasChange("nsg_ids") {
				details.NsgIds = tmp
			}
		}
		if routingMethod, ok := s.D.GetOkExists("routing_method"); ok {
			details.RoutingMethod = oci_golden_gate.RoutingMethodEnum(routingMethod.(string))
		}
		if securityAttributes, ok := s.D.GetOkExists("security_attributes"); ok {
			details.SecurityAttributes = tfresource.MapToSecurityAttributes(securityAttributes.(map[string]interface{}))
		}
		if subnetId, ok := s.D.GetOkExists("subnet_id"); ok {
			tmp := subnetId.(string)
			details.SubnetId = &tmp
		}
		if vaultId, ok := s.D.GetOkExists("vault_id"); ok {
			tmp := vaultId.(string)
			details.VaultId = &tmp
		}
		request.UpdateConnectionDetails = details
	default:
		return fmt.Errorf("unknown connection_type '%v' was specified", connectionType)
	}
	return nil
}

func (s *GoldenGateConnectionResourceCrud) updateCompartment(compartment interface{}) error {
	changeCompartmentRequest := oci_golden_gate.ChangeConnectionCompartmentRequest{}

	compartmentTmp := compartment.(string)
	changeCompartmentRequest.CompartmentId = &compartmentTmp

	idTmp := s.D.Id()
	changeCompartmentRequest.ConnectionId = &idTmp

	if isLockOverride, ok := s.D.GetOkExists("is_lock_override"); ok {
		tmp := isLockOverride.(bool)
		changeCompartmentRequest.IsLockOverride = &tmp
	}

	changeCompartmentRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate")

	response, err := s.Client.ChangeConnectionCompartment(context.Background(), changeCompartmentRequest)
	if err != nil {
		return err
	}

	workId := response.OpcWorkRequestId
	return s.getConnectionFromWorkRequest(workId, tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "golden_gate"), oci_golden_gate.ActionTypeUpdated, s.D.Timeout(schema.TimeoutUpdate))
}
