// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package identity_domains

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	oci_identity_domains "github.com/oracle/oci-go-sdk/v65/identitydomains"

	"github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
)

func IdentityDomainsAppResource() *schema.Resource {
	return &schema.Resource{
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: tfresource.DefaultTimeout,
		Create:   createIdentityDomainsApp,
		Read:     readIdentityDomainsApp,
		Update:   updateIdentityDomainsApp,
		Delete:   deleteIdentityDomainsApp,
		Schema: map[string]*schema.Schema{
			// Required
			"based_on_template": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						// Optional
						"well_known_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"display_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"idcs_endpoint": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"schemas": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			// Optional
			"access_token_expiry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"alias_apps": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"all_url_schemes_allowed": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_access_control": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allow_offline": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"allowed_grants": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_operations": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"allowed_scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"fqs": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"id_of_defining_app": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"allowed_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"app_icon": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_signon_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"app_thumbnail": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"apps_network_perimeters": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"as_opc_service": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"attr_rendering_metadata": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"datatype": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"helptext": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"label": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_length": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_length": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"min_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"order": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"read_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"regexp": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"section": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"visible": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"widget": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"attribute_sets": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"attributes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"audience": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authorization": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"bypass_consent": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"certificates": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"cert_alias": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"kid": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sha1thumbprint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"x509base64certificate": {
							// binary object similar to tagSlug
							Type:     schema.TypeString,
							Computed: true,
						},
						"x5t": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"client_ip_checking": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"contact_email_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delegated_service_names": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_kmsi_token_authentication": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"error_page_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"home_page_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icon": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"id_token_enc_algo": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"identity_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"display": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idp_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"is_alias_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_enterprise_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_form_fill": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_kerberos_realm": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_login_target": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_mobile_target": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_multicloud_service_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_oauth_client": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_oauth_resource": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_obligation_capable": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_radius_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_saml_service_provider": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"is_unmanaged_app": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"is_web_tier_policy": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"landing_page_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linking_callback_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_mechanism": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_page_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logout_page_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logout_uri": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"post_logout_redirect_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"privacy_policy_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product_logo_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"product_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protectable_secondary_audiences": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"radius_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redirect_uris": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"refresh_token_expiry": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"resource_type_schema_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml_service_provider": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"scopes": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"requires_consent": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
						"fqs": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"read_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"secondary_audiences": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_params": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"value": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"service_type_urn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_type_version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"show_in_my_apps": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"signon_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"key": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"terms_of_service_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"terms_of_use": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"trust_policies": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"trust_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"urnietfparamsscimschemasoracleidcsextension_oci_tags": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"defined_tags": {
							Type:             schema.TypeList,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: tfresource.DefinedTagsDiffSuppressFunction,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},
						"freeform_tags": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"key": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
								},
							},
						},

						// Computed
						"tag_slug": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensiondbcs_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"domain_app": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"domain_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionenterprise_app_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"allow_authz_decision_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"allow_authz_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"app_resources": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"deny_authz_decision_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"deny_authz_policy": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionform_fill_app_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"configuration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_cred_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_credential_sharing_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_fill_url_match": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"form_url": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"form_url_match_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"form_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reveal_password_on_form": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"user_name_form_expression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_name_form_template": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"sync_from_template": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"configuration": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_cred_method": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_credential_sharing_group_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_fill_url_match": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"form_url": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"form_url_match_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"form_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"reveal_password_on_form": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"sync_from_template": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"user_name_form_expression": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_name_form_template": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"default_encryption_salt_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"master_key": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_renewable_age": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"max_ticket_life": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"realm_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"supported_encryption_salt_types": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ticket_flags": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionmanagedapp_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"admin_consent_granted": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"bundle_configuration_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"icf_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"required": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"confidential": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"help_message": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"bundle_pool_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"max_idle": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_objects": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"max_wait": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_evictable_idle_time_millis": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"min_idle": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"connected": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"enable_auth_sync_new_user_notification": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"enable_sync": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"enable_sync_summary_report_notification": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"flat_file_bundle_configuration_properties": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"icf_type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"required": {
										Type:     schema.TypeBool,
										Required: true,
									},

									// Optional
									"confidential": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"display_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"help_message": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"order": {
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									// Computed
								},
							},
						},
						"flat_file_connector_bundle": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"well_known_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_authoritative": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"three_legged_oauth_credential": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required

									// Optional
									"access_token": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"access_token_expiry": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"refresh_token": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"account_form_visible": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"can_be_authoritative": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"connector_bundle": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"well_known_id": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"identity_bridges": {
							Type:     schema.TypeList,
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
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"is_directory": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_on_premise_app": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_schema_customization_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_schema_discovery_supported": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_three_legged_oauth_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"is_two_legged_oauth_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"object_classes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"type": {
										Type:     schema.TypeString,
										Required: true,
									},
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"display": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"is_account_object_class": {
										Type:     schema.TypeBool,
										Optional: true,
										Computed: true,
									},
									"resource_type": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sync_config_last_modified": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"three_legged_oauth_provider_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"multicloud_service_type": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"multicloud_platform_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionopc_service_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"service_instance_identifier": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"current_federation_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"current_synchronization_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabling_next_fed_sync_modes": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"next_federation_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"next_synchronization_mode": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionradius_app_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"client_ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"include_group_in_response": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"port": {
							Type:     schema.TypeString,
							Required: true,
						},
						"secret_key": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"capture_client_ip": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"country_code_response_attribute_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_user_ip_attribute": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_membership_radius_attribute": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_membership_to_return": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional

									// Computed
									"display": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"group_name_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"password_and_otp_together": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"radius_vendor_specific_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"response_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"response_format_delimiter": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type_of_radius_app": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionrequestable_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"requestable": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"assertion_consumer_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encrypt_assertion": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"encryption_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"encryption_certificate": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"federation_protocol": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group_assertion_attributes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"condition": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"format": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group_name": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},
						"hok_acs_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hok_required": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"include_signing_cert_in_signature": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"key_encryption_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"logout_binding": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"logout_enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"logout_request_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"logout_response_url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metadata": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name_id_format": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name_id_userstore_attribute": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"partner_provider_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"partner_provider_pattern": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sign_response_or_assertion": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"signature_hash_algorithm": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"signing_certificate": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"succinct_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"user_assertion_attributes": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"name": {
										Type:     schema.TypeString,
										Required: true,
									},
									"user_store_attribute_name": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"format": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
								},
							},
						},

						// Computed
						"last_notification_sent_time": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"outbound_assertion_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// Required
									"value": {
										Type:     schema.TypeString,
										Required: true,
									},

									// Optional
									"direction": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},

									// Computed
									"ref": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"tenant_provider_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"resource_ref": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"web_tier_policy_az_control": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"web_tier_policy_json": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},

			// Computed
			"accounts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"active": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"owner_id": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"admin_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"callback_service_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_control_properties": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						// Optional

						// Computed
					},
				},
			},
			"compartment_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_in_progress": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"domain_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"editable_attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional

						// Computed
					},
				},
			},
			"granted_app_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"admin_role": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"app_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"app_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"legacy_group_name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"read_only": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"grants": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"grant_mechanism": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"grantee_id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"grantee_type": {
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
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"hashed_client_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_created_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_modified_by": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ocid": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"idcs_last_upgraded_in_release": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"idcs_prevented_operations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"infrastructure": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_database_service": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_managed_app": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_opc_service": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"meta": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required

						// Optional
						"created": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_modified": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"location": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"resource_type": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
					},
				},
			},
			"meter_as_opc_service": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"migrated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ready_to_upgrade": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenancy_ocid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_roles": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						// Required
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},

						// Optional
						"description": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"display": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},

						// Computed
						"ref": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createIdentityDomainsApp(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAppResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.CreateResource(d, sync)
}

func readIdentityDomainsApp(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAppResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpointForRead(d, "apps")
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.ReadResource(sync)
}

func updateIdentityDomainsApp(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAppResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client

	return tfresource.UpdateResource(d, sync)
}

func deleteIdentityDomainsApp(d *schema.ResourceData, m interface{}) error {
	sync := &IdentityDomainsAppResourceCrud{}
	sync.D = d
	idcsEndpoint, err := getIdcsEndpoint(d)
	if err != nil {
		return err
	}
	client, err := m.(*client.OracleClients).IdentityDomainsClientWithEndpoint(idcsEndpoint)
	if err != nil {
		return err
	}
	sync.Client = client
	sync.DisableNotFoundRetries = true

	return tfresource.DeleteResource(d, sync)
}

type IdentityDomainsAppResourceCrud struct {
	tfresource.BaseCrud
	Client                 *oci_identity_domains.IdentityDomainsClient
	Res                    *oci_identity_domains.App
	DisableNotFoundRetries bool
}

func (s *IdentityDomainsAppResourceCrud) ID() string {
	return *s.Res.Id
}

func (s *IdentityDomainsAppResourceCrud) Create() error {
	request := oci_identity_domains.CreateAppRequest{}

	if accessTokenExpiry, ok := s.D.GetOkExists("access_token_expiry"); ok {
		tmp := accessTokenExpiry.(int)
		request.AccessTokenExpiry = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if aliasApps, ok := s.D.GetOkExists("alias_apps"); ok {
		interfaces := aliasApps.([]interface{})
		tmp := make([]oci_identity_domains.AppAliasApps, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alias_apps", stateDataIndex)
			converted, err := s.mapToAppAliasApps(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alias_apps") {
			request.AliasApps = tmp
		}
	}

	if allUrlSchemesAllowed, ok := s.D.GetOkExists("all_url_schemes_allowed"); ok {
		tmp := allUrlSchemesAllowed.(bool)
		request.AllUrlSchemesAllowed = &tmp
	}

	if allowAccessControl, ok := s.D.GetOkExists("allow_access_control"); ok {
		tmp := allowAccessControl.(bool)
		request.AllowAccessControl = &tmp
	}

	if allowOffline, ok := s.D.GetOkExists("allow_offline"); ok {
		tmp := allowOffline.(bool)
		request.AllowOffline = &tmp
	}

	if allowedGrants, ok := s.D.GetOkExists("allowed_grants"); ok {
		interfaces := allowedGrants.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_grants") {
			request.AllowedGrants = tmp
		}
	}

	if allowedOperations, ok := s.D.GetOkExists("allowed_operations"); ok {
		interfaces := allowedOperations.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedOperationsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AppAllowedOperationsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_operations") {
			request.AllowedOperations = tmp
		}
	}

	if allowedScopes, ok := s.D.GetOkExists("allowed_scopes"); ok {
		interfaces := allowedScopes.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedScopes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_scopes", stateDataIndex)
			converted, err := s.mapToAppAllowedScopes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_scopes") {
			request.AllowedScopes = tmp
		}
	}

	if allowedTags, ok := s.D.GetOkExists("allowed_tags"); ok {
		interfaces := allowedTags.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_tags", stateDataIndex)
			converted, err := s.mapToAppAllowedTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_tags") {
			request.AllowedTags = tmp
		}
	}

	if appIcon, ok := s.D.GetOkExists("app_icon"); ok {
		tmp := appIcon.(string)
		request.AppIcon = &tmp
	}

	if appSignonPolicy, ok := s.D.GetOkExists("app_signon_policy"); ok {
		if tmpList := appSignonPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "app_signon_policy", 0)
			tmp, err := s.mapToAppAppSignonPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AppSignonPolicy = &tmp
		}
	}

	if appThumbnail, ok := s.D.GetOkExists("app_thumbnail"); ok {
		tmp := appThumbnail.(string)
		request.AppThumbnail = &tmp
	}

	if appsNetworkPerimeters, ok := s.D.GetOkExists("apps_network_perimeters"); ok {
		interfaces := appsNetworkPerimeters.([]interface{})
		tmp := make([]oci_identity_domains.AppAppsNetworkPerimeters, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "apps_network_perimeters", stateDataIndex)
			converted, err := s.mapToAppAppsNetworkPerimeters(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("apps_network_perimeters") {
			request.AppsNetworkPerimeters = tmp
		}
	}

	if asOPCService, ok := s.D.GetOkExists("as_opc_service"); ok {
		if tmpList := asOPCService.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "as_opc_service", 0)
			tmp, err := s.mapToAppAsOPCService(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AsOPCService = &tmp
		}
	}

	if attrRenderingMetadata, ok := s.D.GetOkExists("attr_rendering_metadata"); ok {
		interfaces := attrRenderingMetadata.([]interface{})
		tmp := make([]oci_identity_domains.AppAttrRenderingMetadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attr_rendering_metadata", stateDataIndex)
			converted, err := s.mapToAppAttrRenderingMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attr_rendering_metadata") {
			request.AttrRenderingMetadata = tmp
		}
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if audience, ok := s.D.GetOkExists("audience"); ok {
		tmp := audience.(string)
		request.Audience = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if basedOnTemplate, ok := s.D.GetOkExists("based_on_template"); ok {
		if tmpList := basedOnTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "based_on_template", 0)
			tmp, err := s.mapToAppBasedOnTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BasedOnTemplate = &tmp
		}
	}

	if bypassConsent, ok := s.D.GetOkExists("bypass_consent"); ok {
		tmp := bypassConsent.(bool)
		request.BypassConsent = &tmp
	}

	if certificates, ok := s.D.GetOkExists("certificates"); ok {
		interfaces := certificates.([]interface{})
		tmp := make([]oci_identity_domains.AppCertificates, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificates", stateDataIndex)
			converted, err := s.mapToAppCertificates(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificates") {
			request.Certificates = tmp
		}
	}

	if clientIPChecking, ok := s.D.GetOkExists("client_ip_checking"); ok {
		request.ClientIPChecking = oci_identity_domains.AppClientIPCheckingEnum(clientIPChecking.(string))
	}

	if clientType, ok := s.D.GetOkExists("client_type"); ok {
		request.ClientType = oci_identity_domains.AppClientTypeEnum(clientType.(string))
	}

	if contactEmailAddress, ok := s.D.GetOkExists("contact_email_address"); ok {
		tmp := contactEmailAddress.(string)
		request.ContactEmailAddress = &tmp
	}

	if delegatedServiceNames, ok := s.D.GetOkExists("delegated_service_names"); ok {
		interfaces := delegatedServiceNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("delegated_service_names") {
			request.DelegatedServiceNames = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if disableKmsiTokenAuthentication, ok := s.D.GetOkExists("disable_kmsi_token_authentication"); ok {
		tmp := disableKmsiTokenAuthentication.(bool)
		request.DisableKmsiTokenAuthentication = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if errorPageUrl, ok := s.D.GetOkExists("error_page_url"); ok {
		tmp := errorPageUrl.(string)
		request.ErrorPageUrl = &tmp
	}

	if homePageUrl, ok := s.D.GetOkExists("home_page_url"); ok {
		tmp := homePageUrl.(string)
		request.HomePageUrl = &tmp
	}

	if icon, ok := s.D.GetOkExists("icon"); ok {
		tmp := icon.(string)
		request.Icon = &tmp
	}

	if id, ok := s.D.GetOkExists("id"); ok {
		tmp := id.(string)
		request.Id = &tmp
	}

	if idTokenEncAlgo, ok := s.D.GetOkExists("id_token_enc_algo"); ok {
		tmp := idTokenEncAlgo.(string)
		request.IdTokenEncAlgo = &tmp
	}

	if identityProviders, ok := s.D.GetOkExists("identity_providers"); ok {
		interfaces := identityProviders.([]interface{})
		tmp := make([]oci_identity_domains.AppIdentityProviders, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "identity_providers", stateDataIndex)
			converted, err := s.mapToAppIdentityProviders(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("identity_providers") {
			request.IdentityProviders = tmp
		}
	}

	if idpPolicy, ok := s.D.GetOkExists("idp_policy"); ok {
		if tmpList := idpPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "idp_policy", 0)
			tmp, err := s.mapToAppIdpPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IdpPolicy = &tmp
		}
	}

	if isAliasApp, ok := s.D.GetOkExists("is_alias_app"); ok {
		tmp := isAliasApp.(bool)
		request.IsAliasApp = &tmp
	}

	if isEnterpriseApp, ok := s.D.GetOkExists("is_enterprise_app"); ok {
		tmp := isEnterpriseApp.(bool)
		request.IsEnterpriseApp = &tmp
	}

	if isFormFill, ok := s.D.GetOkExists("is_form_fill"); ok {
		tmp := isFormFill.(bool)
		request.IsFormFill = &tmp
	}

	if isKerberosRealm, ok := s.D.GetOkExists("is_kerberos_realm"); ok {
		tmp := isKerberosRealm.(bool)
		request.IsKerberosRealm = &tmp
	}

	if isLoginTarget, ok := s.D.GetOkExists("is_login_target"); ok {
		tmp := isLoginTarget.(bool)
		request.IsLoginTarget = &tmp
	}

	if isMobileTarget, ok := s.D.GetOkExists("is_mobile_target"); ok {
		tmp := isMobileTarget.(bool)
		request.IsMobileTarget = &tmp
	}

	if isMulticloudServiceApp, ok := s.D.GetOkExists("is_multicloud_service_app"); ok {
		tmp := isMulticloudServiceApp.(bool)
		request.IsMulticloudServiceApp = &tmp
	}

	if isOAuthClient, ok := s.D.GetOkExists("is_oauth_client"); ok {
		tmp := isOAuthClient.(bool)
		request.IsOAuthClient = &tmp
	}

	if isOAuthResource, ok := s.D.GetOkExists("is_oauth_resource"); ok {
		tmp := isOAuthResource.(bool)
		request.IsOAuthResource = &tmp
	}

	if isObligationCapable, ok := s.D.GetOkExists("is_obligation_capable"); ok {
		tmp := isObligationCapable.(bool)
		request.IsObligationCapable = &tmp
	}

	if isRadiusApp, ok := s.D.GetOkExists("is_radius_app"); ok {
		tmp := isRadiusApp.(bool)
		request.IsRadiusApp = &tmp
	}

	if isSamlServiceProvider, ok := s.D.GetOkExists("is_saml_service_provider"); ok {
		tmp := isSamlServiceProvider.(bool)
		request.IsSamlServiceProvider = &tmp
	}

	if isUnmanagedApp, ok := s.D.GetOkExists("is_unmanaged_app"); ok {
		tmp := isUnmanagedApp.(bool)
		request.IsUnmanagedApp = &tmp
	}

	if isWebTierPolicy, ok := s.D.GetOkExists("is_web_tier_policy"); ok {
		tmp := isWebTierPolicy.(bool)
		request.IsWebTierPolicy = &tmp
	}

	if landingPageUrl, ok := s.D.GetOkExists("landing_page_url"); ok {
		tmp := landingPageUrl.(string)
		request.LandingPageUrl = &tmp
	}

	if linkingCallbackUrl, ok := s.D.GetOkExists("linking_callback_url"); ok {
		tmp := linkingCallbackUrl.(string)
		request.LinkingCallbackUrl = &tmp
	}

	if loginMechanism, ok := s.D.GetOkExists("login_mechanism"); ok {
		request.LoginMechanism = oci_identity_domains.AppLoginMechanismEnum(loginMechanism.(string))
	}

	if loginPageUrl, ok := s.D.GetOkExists("login_page_url"); ok {
		tmp := loginPageUrl.(string)
		request.LoginPageUrl = &tmp
	}

	if logoutPageUrl, ok := s.D.GetOkExists("logout_page_url"); ok {
		tmp := logoutPageUrl.(string)
		request.LogoutPageUrl = &tmp
	}

	if logoutUri, ok := s.D.GetOkExists("logout_uri"); ok {
		tmp := logoutUri.(string)
		request.LogoutUri = &tmp
	}

	if name, ok := s.D.GetOkExists("name"); ok {
		tmp := name.(string)
		request.Name = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if postLogoutRedirectUris, ok := s.D.GetOkExists("post_logout_redirect_uris"); ok {
		interfaces := postLogoutRedirectUris.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("post_logout_redirect_uris") {
			request.PostLogoutRedirectUris = tmp
		}
	}

	if privacyPolicyUrl, ok := s.D.GetOkExists("privacy_policy_url"); ok {
		tmp := privacyPolicyUrl.(string)
		request.PrivacyPolicyUrl = &tmp
	}

	if productLogoUrl, ok := s.D.GetOkExists("product_logo_url"); ok {
		tmp := productLogoUrl.(string)
		request.ProductLogoUrl = &tmp
	}

	if productName, ok := s.D.GetOkExists("product_name"); ok {
		tmp := productName.(string)
		request.ProductName = &tmp
	}

	if protectableSecondaryAudiences, ok := s.D.GetOkExists("protectable_secondary_audiences"); ok {
		interfaces := protectableSecondaryAudiences.([]interface{})
		tmp := make([]oci_identity_domains.AppProtectableSecondaryAudiences, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "protectable_secondary_audiences", stateDataIndex)
			converted, err := s.mapToAppProtectableSecondaryAudiences(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("protectable_secondary_audiences") {
			request.ProtectableSecondaryAudiences = tmp
		}
	}

	if radiusPolicy, ok := s.D.GetOkExists("radius_policy"); ok {
		if tmpList := radiusPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "radius_policy", 0)
			tmp, err := s.mapToAppRadiusPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RadiusPolicy = &tmp
		}
	}

	if redirectUris, ok := s.D.GetOkExists("redirect_uris"); ok {
		interfaces := redirectUris.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("redirect_uris") {
			request.RedirectUris = tmp
		}
	}

	if refreshTokenExpiry, ok := s.D.GetOkExists("refresh_token_expiry"); ok {
		tmp := refreshTokenExpiry.(int)
		request.RefreshTokenExpiry = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if samlServiceProvider, ok := s.D.GetOkExists("saml_service_provider"); ok {
		if tmpList := samlServiceProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "saml_service_provider", 0)
			tmp, err := s.mapToAppSamlServiceProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SamlServiceProvider = &tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if scopes, ok := s.D.GetOkExists("scopes"); ok {
		interfaces := scopes.([]interface{})
		tmp := make([]oci_identity_domains.AppScopes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scopes", stateDataIndex)
			converted, err := s.mapToAppScopes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scopes") {
			request.Scopes = tmp
		}
	}

	if secondaryAudiences, ok := s.D.GetOkExists("secondary_audiences"); ok {
		interfaces := secondaryAudiences.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("secondary_audiences") {
			request.SecondaryAudiences = tmp
		}
	}

	if serviceParams, ok := s.D.GetOkExists("service_params"); ok {
		interfaces := serviceParams.([]interface{})
		tmp := make([]oci_identity_domains.AppServiceParams, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_params", stateDataIndex)
			converted, err := s.mapToAppServiceParams(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("service_params") {
			request.ServiceParams = tmp
		}
	}

	if serviceTypeURN, ok := s.D.GetOkExists("service_type_urn"); ok {
		tmp := serviceTypeURN.(string)
		request.ServiceTypeURN = &tmp
	}

	if serviceTypeVersion, ok := s.D.GetOkExists("service_type_version"); ok {
		tmp := serviceTypeVersion.(string)
		request.ServiceTypeVersion = &tmp
	}

	if showInMyApps, ok := s.D.GetOkExists("show_in_my_apps"); ok {
		tmp := showInMyApps.(bool)
		request.ShowInMyApps = &tmp
	}

	if signonPolicy, ok := s.D.GetOkExists("signon_policy"); ok {
		if tmpList := signonPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "signon_policy", 0)
			tmp, err := s.mapToAppSignonPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SignonPolicy = &tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if termsOfServiceUrl, ok := s.D.GetOkExists("terms_of_service_url"); ok {
		tmp := termsOfServiceUrl.(string)
		request.TermsOfServiceUrl = &tmp
	}

	if termsOfUse, ok := s.D.GetOkExists("terms_of_use"); ok {
		if tmpList := termsOfUse.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "terms_of_use", 0)
			tmp, err := s.mapToAppTermsOfUse(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TermsOfUse = &tmp
		}
	}

	if trustPolicies, ok := s.D.GetOkExists("trust_policies"); ok {
		interfaces := trustPolicies.([]interface{})
		tmp := make([]oci_identity_domains.AppTrustPolicies, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trust_policies", stateDataIndex)
			converted, err := s.mapToAppTrustPolicies(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("trust_policies") {
			request.TrustPolicies = tmp
		}
	}

	if trustScope, ok := s.D.GetOkExists("trust_scope"); ok {
		request.TrustScope = oci_identity_domains.AppTrustScopeEnum(trustScope.(string))
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondbcsApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondbcs_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondbcsApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondbcs_app", 0)
			tmp, err := s.mapToAppExtensionDbcsApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionenterpriseAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionenterpriseAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", 0)
			tmp, err := s.mapToAppExtensionEnterpriseAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionformFillAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionformFillAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", 0)
			tmp, err := s.mapToAppExtensionFormFillAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplate, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", 0)
			tmp, err := s.mapToAppExtensionFormFillAppTemplateAppTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionkerberosRealmApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionkerberosRealmApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", 0)
			tmp, err := s.mapToAppExtensionKerberosRealmApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmanagedappApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmanagedappApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", 0)
			tmp, err := s.mapToAppExtensionManagedappApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", 0)
			tmp, err := s.mapToAppExtensionMulticloudServiceAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionopcServiceApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionopc_service_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionopcServiceApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionopc_service_app", 0)
			tmp, err := s.mapToAppExtensionOpcServiceApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionradiusAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionradius_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionradiusAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionradius_app_app", 0)
			tmp, err := s.mapToAppExtensionRadiusAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionrequestableApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionrequestable_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionrequestableApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionrequestable_app", 0)
			tmp, err := s.mapToAppExtensionRequestableApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsamlServiceProviderApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsamlServiceProviderApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", 0)
			tmp, err := s.mapToAppExtensionSamlServiceProviderApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionwebTierPolicyApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionwebTierPolicyApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", 0)
			tmp, err := s.mapToAppExtensionWebTierPolicyApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.CreateApp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.App
	return nil
}

func (s *IdentityDomainsAppResourceCrud) Get() error {
	request := oci_identity_domains.GetAppRequest{}

	tmp := s.D.Id()
	request.AppId = &tmp

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	appId, err := parseAppCompositeId(s.D.Id())
	if err == nil {
		request.AppId = &appId
	} else {
		log.Printf("[WARN] Get() unable to parse current ID: %s", s.D.Id())
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.GetApp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.App
	return nil
}

func (s *IdentityDomainsAppResourceCrud) Update() error {
	request := oci_identity_domains.PutAppRequest{}

	if accessTokenExpiry, ok := s.D.GetOkExists("access_token_expiry"); ok {
		tmp := accessTokenExpiry.(int)
		request.AccessTokenExpiry = &tmp
	}

	if active, ok := s.D.GetOkExists("active"); ok {
		tmp := active.(bool)
		request.Active = &tmp
	}

	if aliasApps, ok := s.D.GetOkExists("alias_apps"); ok {
		interfaces := aliasApps.([]interface{})
		tmp := make([]oci_identity_domains.AppAliasApps, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "alias_apps", stateDataIndex)
			converted, err := s.mapToAppAliasApps(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("alias_apps") {
			request.AliasApps = tmp
		}
	}

	if allUrlSchemesAllowed, ok := s.D.GetOkExists("all_url_schemes_allowed"); ok {
		tmp := allUrlSchemesAllowed.(bool)
		request.AllUrlSchemesAllowed = &tmp
	}

	if allowAccessControl, ok := s.D.GetOkExists("allow_access_control"); ok {
		tmp := allowAccessControl.(bool)
		request.AllowAccessControl = &tmp
	}

	if allowOffline, ok := s.D.GetOkExists("allow_offline"); ok {
		tmp := allowOffline.(bool)
		request.AllowOffline = &tmp
	}

	if allowedGrants, ok := s.D.GetOkExists("allowed_grants"); ok {
		interfaces := allowedGrants.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_grants") {
			request.AllowedGrants = tmp
		}
	}

	if allowedOperations, ok := s.D.GetOkExists("allowed_operations"); ok {
		interfaces := allowedOperations.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedOperationsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AppAllowedOperationsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_operations") {
			request.AllowedOperations = tmp
		}
	}

	if allowedScopes, ok := s.D.GetOkExists("allowed_scopes"); ok {
		interfaces := allowedScopes.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedScopes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_scopes", stateDataIndex)
			converted, err := s.mapToAppAllowedScopes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_scopes") {
			request.AllowedScopes = tmp
		}
	}

	if allowedTags, ok := s.D.GetOkExists("allowed_tags"); ok {
		interfaces := allowedTags.([]interface{})
		tmp := make([]oci_identity_domains.AppAllowedTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "allowed_tags", stateDataIndex)
			converted, err := s.mapToAppAllowedTags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("allowed_tags") {
			request.AllowedTags = tmp
		}
	}

	if appIcon, ok := s.D.GetOkExists("app_icon"); ok {
		tmp := appIcon.(string)
		request.AppIcon = &tmp
	}

	tmp := s.D.Id()
	request.AppId = &tmp

	if appSignonPolicy, ok := s.D.GetOkExists("app_signon_policy"); ok {
		if tmpList := appSignonPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "app_signon_policy", 0)
			tmp, err := s.mapToAppAppSignonPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AppSignonPolicy = &tmp
		}
	}

	if appThumbnail, ok := s.D.GetOkExists("app_thumbnail"); ok {
		tmp := appThumbnail.(string)
		request.AppThumbnail = &tmp
	}

	if appsNetworkPerimeters, ok := s.D.GetOkExists("apps_network_perimeters"); ok {
		interfaces := appsNetworkPerimeters.([]interface{})
		tmp := make([]oci_identity_domains.AppAppsNetworkPerimeters, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "apps_network_perimeters", stateDataIndex)
			converted, err := s.mapToAppAppsNetworkPerimeters(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("apps_network_perimeters") {
			request.AppsNetworkPerimeters = tmp
		}
	}

	if asOPCService, ok := s.D.GetOkExists("as_opc_service"); ok {
		if tmpList := asOPCService.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "as_opc_service", 0)
			tmp, err := s.mapToAppAsOPCService(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.AsOPCService = &tmp
		}
	}

	if attrRenderingMetadata, ok := s.D.GetOkExists("attr_rendering_metadata"); ok {
		interfaces := attrRenderingMetadata.([]interface{})
		tmp := make([]oci_identity_domains.AppAttrRenderingMetadata, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "attr_rendering_metadata", stateDataIndex)
			converted, err := s.mapToAppAttrRenderingMetadata(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("attr_rendering_metadata") {
			request.AttrRenderingMetadata = tmp
		}
	}

	if attributeSets, ok := s.D.GetOkExists("attribute_sets"); ok {
		interfaces := attributeSets.([]interface{})
		tmp := make([]oci_identity_domains.AttributeSetsEnum, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = oci_identity_domains.AttributeSetsEnum(interfaces[i].(string))
			}
		}
		if len(tmp) != 0 || s.D.HasChange("attribute_sets") {
			request.AttributeSets = tmp
		}
	}

	if attributes, ok := s.D.GetOkExists("attributes"); ok {
		tmp := attributes.(string)
		request.Attributes = &tmp
	}

	if audience, ok := s.D.GetOkExists("audience"); ok {
		tmp := audience.(string)
		request.Audience = &tmp
	}

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if basedOnTemplate, ok := s.D.GetOkExists("based_on_template"); ok {
		if tmpList := basedOnTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "based_on_template", 0)
			tmp, err := s.mapToAppBasedOnTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.BasedOnTemplate = &tmp
		}
	}

	if bypassConsent, ok := s.D.GetOkExists("bypass_consent"); ok {
		tmp := bypassConsent.(bool)
		request.BypassConsent = &tmp
	}

	if certificates, ok := s.D.GetOkExists("certificates"); ok {
		interfaces := certificates.([]interface{})
		tmp := make([]oci_identity_domains.AppCertificates, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "certificates", stateDataIndex)
			converted, err := s.mapToAppCertificates(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("certificates") {
			request.Certificates = tmp
		}
	}

	if clientIPChecking, ok := s.D.GetOkExists("client_ip_checking"); ok {
		request.ClientIPChecking = oci_identity_domains.AppClientIPCheckingEnum(clientIPChecking.(string))
	}

	if clientType, ok := s.D.GetOkExists("client_type"); ok {
		request.ClientType = oci_identity_domains.AppClientTypeEnum(clientType.(string))
	}

	if contactEmailAddress, ok := s.D.GetOkExists("contact_email_address"); ok {
		tmp := contactEmailAddress.(string)
		request.ContactEmailAddress = &tmp
	}

	if delegatedServiceNames, ok := s.D.GetOkExists("delegated_service_names"); ok {
		interfaces := delegatedServiceNames.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("delegated_service_names") {
			request.DelegatedServiceNames = tmp
		}
	}

	if description, ok := s.D.GetOkExists("description"); ok {
		tmp := description.(string)
		request.Description = &tmp
	}

	if disableKmsiTokenAuthentication, ok := s.D.GetOkExists("disable_kmsi_token_authentication"); ok {
		tmp := disableKmsiTokenAuthentication.(bool)
		request.DisableKmsiTokenAuthentication = &tmp
	}

	if displayName, ok := s.D.GetOkExists("display_name"); ok {
		tmp := displayName.(string)
		request.DisplayName = &tmp
	}

	if errorPageUrl, ok := s.D.GetOkExists("error_page_url"); ok {
		tmp := errorPageUrl.(string)
		request.ErrorPageUrl = &tmp
	}

	if homePageUrl, ok := s.D.GetOkExists("home_page_url"); ok {
		tmp := homePageUrl.(string)
		request.HomePageUrl = &tmp
	}

	if icon, ok := s.D.GetOkExists("icon"); ok {
		tmp := icon.(string)
		request.Icon = &tmp
	}

	tmp = s.D.Id()
	request.Id = &tmp

	if idTokenEncAlgo, ok := s.D.GetOkExists("id_token_enc_algo"); ok {
		tmp := idTokenEncAlgo.(string)
		request.IdTokenEncAlgo = &tmp
	}

	if identityProviders, ok := s.D.GetOkExists("identity_providers"); ok {
		interfaces := identityProviders.([]interface{})
		tmp := make([]oci_identity_domains.AppIdentityProviders, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "identity_providers", stateDataIndex)
			converted, err := s.mapToAppIdentityProviders(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("identity_providers") {
			request.IdentityProviders = tmp
		}
	}

	if idpPolicy, ok := s.D.GetOkExists("idp_policy"); ok {
		if tmpList := idpPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "idp_policy", 0)
			tmp, err := s.mapToAppIdpPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.IdpPolicy = &tmp
		}
	}

	if isEnterpriseApp, ok := s.D.GetOkExists("is_enterprise_app"); ok {
		tmp := isEnterpriseApp.(bool)
		request.IsEnterpriseApp = &tmp
	}

	if isFormFill, ok := s.D.GetOkExists("is_form_fill"); ok {
		tmp := isFormFill.(bool)
		request.IsFormFill = &tmp
	}

	if isKerberosRealm, ok := s.D.GetOkExists("is_kerberos_realm"); ok {
		tmp := isKerberosRealm.(bool)
		request.IsKerberosRealm = &tmp
	}

	if isLoginTarget, ok := s.D.GetOkExists("is_login_target"); ok {
		tmp := isLoginTarget.(bool)
		request.IsLoginTarget = &tmp
	}

	if isMobileTarget, ok := s.D.GetOkExists("is_mobile_target"); ok {
		tmp := isMobileTarget.(bool)
		request.IsMobileTarget = &tmp
	}

	if isMulticloudServiceApp, ok := s.D.GetOkExists("is_multicloud_service_app"); ok {
		tmp := isMulticloudServiceApp.(bool)
		request.IsMulticloudServiceApp = &tmp
	}

	if isOAuthClient, ok := s.D.GetOkExists("is_oauth_client"); ok {
		tmp := isOAuthClient.(bool)
		request.IsOAuthClient = &tmp
	}

	if isOAuthResource, ok := s.D.GetOkExists("is_oauth_resource"); ok {
		tmp := isOAuthResource.(bool)
		request.IsOAuthResource = &tmp
	}

	if isObligationCapable, ok := s.D.GetOkExists("is_obligation_capable"); ok {
		tmp := isObligationCapable.(bool)
		request.IsObligationCapable = &tmp
	}

	if isRadiusApp, ok := s.D.GetOkExists("is_radius_app"); ok {
		tmp := isRadiusApp.(bool)
		request.IsRadiusApp = &tmp
	}

	if isSamlServiceProvider, ok := s.D.GetOkExists("is_saml_service_provider"); ok {
		tmp := isSamlServiceProvider.(bool)
		request.IsSamlServiceProvider = &tmp
	}

	if isWebTierPolicy, ok := s.D.GetOkExists("is_web_tier_policy"); ok {
		tmp := isWebTierPolicy.(bool)
		request.IsWebTierPolicy = &tmp
	}

	if landingPageUrl, ok := s.D.GetOkExists("landing_page_url"); ok {
		tmp := landingPageUrl.(string)
		request.LandingPageUrl = &tmp
	}

	if linkingCallbackUrl, ok := s.D.GetOkExists("linking_callback_url"); ok {
		tmp := linkingCallbackUrl.(string)
		request.LinkingCallbackUrl = &tmp
	}

	if loginMechanism, ok := s.D.GetOkExists("login_mechanism"); ok {
		request.LoginMechanism = oci_identity_domains.AppLoginMechanismEnum(loginMechanism.(string))
	}

	if loginPageUrl, ok := s.D.GetOkExists("login_page_url"); ok {
		tmp := loginPageUrl.(string)
		request.LoginPageUrl = &tmp
	}

	if logoutPageUrl, ok := s.D.GetOkExists("logout_page_url"); ok {
		tmp := logoutPageUrl.(string)
		request.LogoutPageUrl = &tmp
	}

	if logoutUri, ok := s.D.GetOkExists("logout_uri"); ok {
		tmp := logoutUri.(string)
		request.LogoutUri = &tmp
	}

	if ocid, ok := s.D.GetOkExists("ocid"); ok {
		tmp := ocid.(string)
		request.Ocid = &tmp
	}

	if postLogoutRedirectUris, ok := s.D.GetOkExists("post_logout_redirect_uris"); ok {
		interfaces := postLogoutRedirectUris.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("post_logout_redirect_uris") {
			request.PostLogoutRedirectUris = tmp
		}
	}

	if privacyPolicyUrl, ok := s.D.GetOkExists("privacy_policy_url"); ok {
		tmp := privacyPolicyUrl.(string)
		request.PrivacyPolicyUrl = &tmp
	}

	if productLogoUrl, ok := s.D.GetOkExists("product_logo_url"); ok {
		tmp := productLogoUrl.(string)
		request.ProductLogoUrl = &tmp
	}

	if productName, ok := s.D.GetOkExists("product_name"); ok {
		tmp := productName.(string)
		request.ProductName = &tmp
	}

	if protectableSecondaryAudiences, ok := s.D.GetOkExists("protectable_secondary_audiences"); ok {
		interfaces := protectableSecondaryAudiences.([]interface{})
		tmp := make([]oci_identity_domains.AppProtectableSecondaryAudiences, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "protectable_secondary_audiences", stateDataIndex)
			converted, err := s.mapToAppProtectableSecondaryAudiences(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("protectable_secondary_audiences") {
			request.ProtectableSecondaryAudiences = tmp
		}
	}

	if radiusPolicy, ok := s.D.GetOkExists("radius_policy"); ok {
		if tmpList := radiusPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "radius_policy", 0)
			tmp, err := s.mapToAppRadiusPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.RadiusPolicy = &tmp
		}
	}

	if redirectUris, ok := s.D.GetOkExists("redirect_uris"); ok {
		interfaces := redirectUris.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("redirect_uris") {
			request.RedirectUris = tmp
		}
	}

	if refreshTokenExpiry, ok := s.D.GetOkExists("refresh_token_expiry"); ok {
		tmp := refreshTokenExpiry.(int)
		request.RefreshTokenExpiry = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	if samlServiceProvider, ok := s.D.GetOkExists("saml_service_provider"); ok {
		if tmpList := samlServiceProvider.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "saml_service_provider", 0)
			tmp, err := s.mapToAppSamlServiceProvider(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SamlServiceProvider = &tmp
		}
	}

	if schemas, ok := s.D.GetOkExists("schemas"); ok {
		interfaces := schemas.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("schemas") {
			request.Schemas = tmp
		}
	}

	if scopes, ok := s.D.GetOkExists("scopes"); ok {
		interfaces := scopes.([]interface{})
		tmp := make([]oci_identity_domains.AppScopes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "scopes", stateDataIndex)
			converted, err := s.mapToAppScopes(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("scopes") {
			request.Scopes = tmp
		}
	}

	if secondaryAudiences, ok := s.D.GetOkExists("secondary_audiences"); ok {
		interfaces := secondaryAudiences.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange("secondary_audiences") {
			request.SecondaryAudiences = tmp
		}
	}

	if serviceParams, ok := s.D.GetOkExists("service_params"); ok {
		interfaces := serviceParams.([]interface{})
		tmp := make([]oci_identity_domains.AppServiceParams, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "service_params", stateDataIndex)
			converted, err := s.mapToAppServiceParams(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("service_params") {
			request.ServiceParams = tmp
		}
	}

	if serviceTypeURN, ok := s.D.GetOkExists("service_type_urn"); ok {
		tmp := serviceTypeURN.(string)
		request.ServiceTypeURN = &tmp
	}

	if serviceTypeVersion, ok := s.D.GetOkExists("service_type_version"); ok {
		tmp := serviceTypeVersion.(string)
		request.ServiceTypeVersion = &tmp
	}

	if showInMyApps, ok := s.D.GetOkExists("show_in_my_apps"); ok {
		tmp := showInMyApps.(bool)
		request.ShowInMyApps = &tmp
	}

	if signonPolicy, ok := s.D.GetOkExists("signon_policy"); ok {
		if tmpList := signonPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "signon_policy", 0)
			tmp, err := s.mapToAppSignonPolicy(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.SignonPolicy = &tmp
		}
	}

	if tags, ok := s.D.GetOkExists("tags"); ok {
		interfaces := tags.([]interface{})
		tmp := make([]oci_identity_domains.Tags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "tags", stateDataIndex)
			converted, err := s.mapTotags(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("tags") {
			request.Tags = tmp
		}
	}

	if termsOfServiceUrl, ok := s.D.GetOkExists("terms_of_service_url"); ok {
		tmp := termsOfServiceUrl.(string)
		request.TermsOfServiceUrl = &tmp
	}

	if termsOfUse, ok := s.D.GetOkExists("terms_of_use"); ok {
		if tmpList := termsOfUse.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "terms_of_use", 0)
			tmp, err := s.mapToAppTermsOfUse(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.TermsOfUse = &tmp
		}
	}

	if trustPolicies, ok := s.D.GetOkExists("trust_policies"); ok {
		interfaces := trustPolicies.([]interface{})
		tmp := make([]oci_identity_domains.AppTrustPolicies, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "trust_policies", stateDataIndex)
			converted, err := s.mapToAppTrustPolicies(fieldKeyFormat)
			if err != nil {
				return err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange("trust_policies") {
			request.TrustPolicies = tmp
		}
	}

	if trustScope, ok := s.D.GetOkExists("trust_scope"); ok {
		request.TrustScope = oci_identity_domains.AppTrustScopeEnum(trustScope.(string))
	}

	if urnietfparamsscimschemasoracleidcsextensionOCITags, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextension_oci_tags"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionOCITags.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextension_oci_tags", 0)
			tmp, err := s.mapToExtensionOCITags(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensiondbcsApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensiondbcs_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensiondbcsApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensiondbcs_app", 0)
			tmp, err := s.mapToAppExtensionDbcsApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionenterpriseAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionenterpriseAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", 0)
			tmp, err := s.mapToAppExtensionEnterpriseAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionformFillAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionformFillAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", 0)
			tmp, err := s.mapToAppExtensionFormFillAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplate, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionformFillAppTemplateAppTemplate.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", 0)
			tmp, err := s.mapToAppExtensionFormFillAppTemplateAppTemplate(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionkerberosRealmApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionkerberosRealmApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", 0)
			tmp, err := s.mapToAppExtensionKerberosRealmApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmanagedappApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmanagedappApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", 0)
			tmp, err := s.mapToAppExtensionManagedappApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionmulticloudServiceAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", 0)
			tmp, err := s.mapToAppExtensionMulticloudServiceAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionopcServiceApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionopc_service_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionopcServiceApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionopc_service_app", 0)
			tmp, err := s.mapToAppExtensionOpcServiceApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionradiusAppApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionradius_app_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionradiusAppApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionradius_app_app", 0)
			tmp, err := s.mapToAppExtensionRadiusAppApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionrequestableApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionrequestable_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionrequestableApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionrequestable_app", 0)
			tmp, err := s.mapToAppExtensionRequestableApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionsamlServiceProviderApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionsamlServiceProviderApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", 0)
			tmp, err := s.mapToAppExtensionSamlServiceProviderApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp = &tmp
		}
	}

	if urnietfparamsscimschemasoracleidcsextensionwebTierPolicyApp, ok := s.D.GetOkExists("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app"); ok {
		if tmpList := urnietfparamsscimschemasoracleidcsextensionwebTierPolicyApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormat := fmt.Sprintf("%s.%d.%%s", "urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", 0)
			tmp, err := s.mapToAppExtensionWebTierPolicyApp(fieldKeyFormat)
			if err != nil {
				return err
			}
			request.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp = &tmp
		}
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	response, err := s.Client.PutApp(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response.App
	return nil
}

func (s *IdentityDomainsAppResourceCrud) Delete() error {
	request := oci_identity_domains.DeleteAppRequest{}

	tmp := s.D.Id()
	request.AppId = &tmp

	if authorization, ok := s.D.GetOkExists("authorization"); ok {
		tmp := authorization.(string)
		request.Authorization = &tmp
	}

	if forceDelete, ok := s.D.GetOkExists("force_delete"); ok {
		tmp := forceDelete.(bool)
		request.ForceDelete = &tmp
	}

	if resourceTypeSchemaVersion, ok := s.D.GetOkExists("resource_type_schema_version"); ok {
		tmp := resourceTypeSchemaVersion.(string)
		request.ResourceTypeSchemaVersion = &tmp
	}

	request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(s.DisableNotFoundRetries, "identity_domains")

	_, err := s.Client.DeleteApp(context.Background(), request)
	return err
}

func (s *IdentityDomainsAppResourceCrud) SetData() error {

	appId, err := parseAppCompositeId(s.D.Id())
	if err == nil {
		s.D.SetId(appId)
	} else {
		log.Printf("[WARN] SetData() unable to parse current ID: %s", s.D.Id())
	}

	if s.Res.AccessTokenExpiry != nil {
		s.D.Set("access_token_expiry", *s.Res.AccessTokenExpiry)
	}

	accounts := []interface{}{}
	for _, item := range s.Res.Accounts {
		accounts = append(accounts, AppAccountsToMap(item))
	}
	s.D.Set("accounts", accounts)

	if s.Res.Active != nil {
		s.D.Set("active", *s.Res.Active)
	}

	adminRoles := []interface{}{}
	for _, item := range s.Res.AdminRoles {
		adminRoles = append(adminRoles, AppAdminRolesToMap(item))
	}
	s.D.Set("admin_roles", adminRoles)

	aliasApps := []interface{}{}
	for _, item := range s.Res.AliasApps {
		aliasApps = append(aliasApps, AppAliasAppsToMap(item))
	}
	s.D.Set("alias_apps", aliasApps)

	if s.Res.AllUrlSchemesAllowed != nil {
		s.D.Set("all_url_schemes_allowed", *s.Res.AllUrlSchemesAllowed)
	}

	if s.Res.AllowAccessControl != nil {
		s.D.Set("allow_access_control", *s.Res.AllowAccessControl)
	}

	if s.Res.AllowOffline != nil {
		s.D.Set("allow_offline", *s.Res.AllowOffline)
	}

	s.D.Set("allowed_grants", s.Res.AllowedGrants)

	s.D.Set("allowed_operations", s.Res.AllowedOperations)

	allowedScopes := []interface{}{}
	for _, item := range s.Res.AllowedScopes {
		allowedScopes = append(allowedScopes, AppAllowedScopesToMap(item))
	}
	s.D.Set("allowed_scopes", allowedScopes)

	allowedTags := []interface{}{}
	for _, item := range s.Res.AllowedTags {
		allowedTags = append(allowedTags, AppAllowedTagsToMap(item))
	}
	s.D.Set("allowed_tags", allowedTags)

	if s.Res.AppIcon != nil {
		s.D.Set("app_icon", *s.Res.AppIcon)
	}

	if s.Res.AppSignonPolicy != nil {
		s.D.Set("app_signon_policy", []interface{}{AppAppSignonPolicyToMap(s.Res.AppSignonPolicy)})
	} else {
		s.D.Set("app_signon_policy", nil)
	}

	if s.Res.AppThumbnail != nil {
		s.D.Set("app_thumbnail", *s.Res.AppThumbnail)
	}

	appsNetworkPerimeters := []interface{}{}
	for _, item := range s.Res.AppsNetworkPerimeters {
		appsNetworkPerimeters = append(appsNetworkPerimeters, AppAppsNetworkPerimetersToMap(item))
	}
	s.D.Set("apps_network_perimeters", appsNetworkPerimeters)

	if s.Res.AsOPCService != nil {
		s.D.Set("as_opc_service", []interface{}{AppAsOPCServiceToMap(s.Res.AsOPCService)})
	} else {
		s.D.Set("as_opc_service", nil)
	}

	attrRenderingMetadata := []interface{}{}
	for _, item := range s.Res.AttrRenderingMetadata {
		attrRenderingMetadata = append(attrRenderingMetadata, AppAttrRenderingMetadataToMap(item))
	}
	s.D.Set("attr_rendering_metadata", attrRenderingMetadata)

	if s.Res.Audience != nil {
		s.D.Set("audience", *s.Res.Audience)
	}

	if s.Res.BasedOnTemplate != nil {
		s.D.Set("based_on_template", []interface{}{AppBasedOnTemplateToMap(s.Res.BasedOnTemplate)})
	} else {
		s.D.Set("based_on_template", nil)
	}

	if s.Res.BypassConsent != nil {
		s.D.Set("bypass_consent", *s.Res.BypassConsent)
	}

	if s.Res.CallbackServiceUrl != nil {
		s.D.Set("callback_service_url", *s.Res.CallbackServiceUrl)
	}

	certificates := []interface{}{}
	for _, item := range s.Res.Certificates {
		certificates = append(certificates, AppCertificatesToMap(item))
	}
	s.D.Set("certificates", certificates)

	s.D.Set("client_ip_checking", s.Res.ClientIPChecking)

	if s.Res.ClientSecret != nil {
		s.D.Set("client_secret", *s.Res.ClientSecret)
	}

	s.D.Set("client_type", s.Res.ClientType)

	cloudControlProperties := []interface{}{}
	for _, item := range s.Res.CloudControlProperties {
		cloudControlProperties = append(cloudControlProperties, AppCloudControlPropertiesToMap(item))
	}
	s.D.Set("cloud_control_properties", cloudControlProperties)

	if s.Res.CompartmentOcid != nil {
		s.D.Set("compartment_ocid", *s.Res.CompartmentOcid)
	}

	if s.Res.ContactEmailAddress != nil {
		s.D.Set("contact_email_address", *s.Res.ContactEmailAddress)
	}

	s.D.Set("delegated_service_names", s.Res.DelegatedServiceNames)

	if s.Res.DeleteInProgress != nil {
		s.D.Set("delete_in_progress", *s.Res.DeleteInProgress)
	}

	if s.Res.Description != nil {
		s.D.Set("description", *s.Res.Description)
	}

	if s.Res.DisableKmsiTokenAuthentication != nil {
		s.D.Set("disable_kmsi_token_authentication", *s.Res.DisableKmsiTokenAuthentication)
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	if s.Res.DomainOcid != nil {
		s.D.Set("domain_ocid", *s.Res.DomainOcid)
	}

	editableAttributes := []interface{}{}
	for _, item := range s.Res.EditableAttributes {
		editableAttributes = append(editableAttributes, AppEditableAttributesToMap(item))
	}
	s.D.Set("editable_attributes", editableAttributes)

	if s.Res.ErrorPageUrl != nil {
		s.D.Set("error_page_url", *s.Res.ErrorPageUrl)
	}

	grantedAppRoles := []interface{}{}
	for _, item := range s.Res.GrantedAppRoles {
		grantedAppRoles = append(grantedAppRoles, AppGrantedAppRolesToMap(item))
	}
	s.D.Set("granted_app_roles", grantedAppRoles)

	grants := []interface{}{}
	for _, item := range s.Res.Grants {
		grants = append(grants, AppGrantsToMap(item))
	}
	s.D.Set("grants", grants)

	if s.Res.HashedClientSecret != nil {
		s.D.Set("hashed_client_secret", *s.Res.HashedClientSecret)
	}

	if s.Res.HomePageUrl != nil {
		s.D.Set("home_page_url", *s.Res.HomePageUrl)
	}

	if s.Res.Icon != nil {
		s.D.Set("icon", *s.Res.Icon)
	}

	if s.Res.IdTokenEncAlgo != nil {
		s.D.Set("id_token_enc_algo", *s.Res.IdTokenEncAlgo)
	}

	if s.Res.IdcsCreatedBy != nil {
		s.D.Set("idcs_created_by", []interface{}{idcsCreatedByToMap(s.Res.IdcsCreatedBy)})
	} else {
		s.D.Set("idcs_created_by", nil)
	}

	if s.Res.IdcsLastModifiedBy != nil {
		s.D.Set("idcs_last_modified_by", []interface{}{idcsLastModifiedByToMap(s.Res.IdcsLastModifiedBy)})
	} else {
		s.D.Set("idcs_last_modified_by", nil)
	}

	if s.Res.IdcsLastUpgradedInRelease != nil {
		s.D.Set("idcs_last_upgraded_in_release", *s.Res.IdcsLastUpgradedInRelease)
	}

	s.D.Set("idcs_prevented_operations", s.Res.IdcsPreventedOperations)

	identityProviders := []interface{}{}
	for _, item := range s.Res.IdentityProviders {
		identityProviders = append(identityProviders, AppIdentityProvidersToMap(item))
	}
	s.D.Set("identity_providers", identityProviders)

	if s.Res.IdpPolicy != nil {
		s.D.Set("idp_policy", []interface{}{AppIdpPolicyToMap(s.Res.IdpPolicy)})
	} else {
		s.D.Set("idp_policy", nil)
	}

	if s.Res.Infrastructure != nil {
		s.D.Set("infrastructure", *s.Res.Infrastructure)
	}

	if s.Res.IsAliasApp != nil {
		s.D.Set("is_alias_app", *s.Res.IsAliasApp)
	}

	if s.Res.IsDatabaseService != nil {
		s.D.Set("is_database_service", *s.Res.IsDatabaseService)
	}

	if s.Res.IsEnterpriseApp != nil {
		s.D.Set("is_enterprise_app", *s.Res.IsEnterpriseApp)
	}

	if s.Res.IsFormFill != nil {
		s.D.Set("is_form_fill", *s.Res.IsFormFill)
	}

	if s.Res.IsKerberosRealm != nil {
		s.D.Set("is_kerberos_realm", *s.Res.IsKerberosRealm)
	}

	if s.Res.IsLoginTarget != nil {
		s.D.Set("is_login_target", *s.Res.IsLoginTarget)
	}

	if s.Res.IsManagedApp != nil {
		s.D.Set("is_managed_app", *s.Res.IsManagedApp)
	}

	if s.Res.IsMobileTarget != nil {
		s.D.Set("is_mobile_target", *s.Res.IsMobileTarget)
	}

	if s.Res.IsMulticloudServiceApp != nil {
		s.D.Set("is_multicloud_service_app", *s.Res.IsMulticloudServiceApp)
	}

	if s.Res.IsOAuthClient != nil {
		s.D.Set("is_oauth_client", *s.Res.IsOAuthClient)
	}

	if s.Res.IsOAuthResource != nil {
		s.D.Set("is_oauth_resource", *s.Res.IsOAuthResource)
	}

	if s.Res.IsOPCService != nil {
		s.D.Set("is_opc_service", *s.Res.IsOPCService)
	}

	if s.Res.IsObligationCapable != nil {
		s.D.Set("is_obligation_capable", *s.Res.IsObligationCapable)
	}

	if s.Res.IsRadiusApp != nil {
		s.D.Set("is_radius_app", *s.Res.IsRadiusApp)
	}

	if s.Res.IsSamlServiceProvider != nil {
		s.D.Set("is_saml_service_provider", *s.Res.IsSamlServiceProvider)
	}

	if s.Res.IsUnmanagedApp != nil {
		s.D.Set("is_unmanaged_app", *s.Res.IsUnmanagedApp)
	}

	if s.Res.IsWebTierPolicy != nil {
		s.D.Set("is_web_tier_policy", *s.Res.IsWebTierPolicy)
	}

	if s.Res.LandingPageUrl != nil {
		s.D.Set("landing_page_url", *s.Res.LandingPageUrl)
	}

	if s.Res.LinkingCallbackUrl != nil {
		s.D.Set("linking_callback_url", *s.Res.LinkingCallbackUrl)
	}

	s.D.Set("login_mechanism", s.Res.LoginMechanism)

	if s.Res.LoginPageUrl != nil {
		s.D.Set("login_page_url", *s.Res.LoginPageUrl)
	}

	if s.Res.LogoutPageUrl != nil {
		s.D.Set("logout_page_url", *s.Res.LogoutPageUrl)
	}

	if s.Res.LogoutUri != nil {
		s.D.Set("logout_uri", *s.Res.LogoutUri)
	}

	if s.Res.Meta != nil {
		s.D.Set("meta", []interface{}{metaToMap(s.Res.Meta)})
	} else {
		s.D.Set("meta", nil)
	}

	if s.Res.MeterAsOPCService != nil {
		s.D.Set("meter_as_opc_service", *s.Res.MeterAsOPCService)
	}

	if s.Res.Migrated != nil {
		s.D.Set("migrated", *s.Res.Migrated)
	}

	if s.Res.Name != nil {
		s.D.Set("name", *s.Res.Name)
	}

	if s.Res.Ocid != nil {
		s.D.Set("ocid", *s.Res.Ocid)
	}

	s.D.Set("post_logout_redirect_uris", s.Res.PostLogoutRedirectUris)

	if s.Res.PrivacyPolicyUrl != nil {
		s.D.Set("privacy_policy_url", *s.Res.PrivacyPolicyUrl)
	}

	if s.Res.ProductLogoUrl != nil {
		s.D.Set("product_logo_url", *s.Res.ProductLogoUrl)
	}

	if s.Res.ProductName != nil {
		s.D.Set("product_name", *s.Res.ProductName)
	}

	protectableSecondaryAudiences := []interface{}{}
	for _, item := range s.Res.ProtectableSecondaryAudiences {
		protectableSecondaryAudiences = append(protectableSecondaryAudiences, AppProtectableSecondaryAudiencesToMap(item))
	}
	s.D.Set("protectable_secondary_audiences", protectableSecondaryAudiences)

	if s.Res.RadiusPolicy != nil {
		s.D.Set("radius_policy", []interface{}{AppRadiusPolicyToMap(s.Res.RadiusPolicy)})
	} else {
		s.D.Set("radius_policy", nil)
	}

	if s.Res.ReadyToUpgrade != nil {
		s.D.Set("ready_to_upgrade", *s.Res.ReadyToUpgrade)
	}

	s.D.Set("redirect_uris", s.Res.RedirectUris)

	if s.Res.RefreshTokenExpiry != nil {
		s.D.Set("refresh_token_expiry", *s.Res.RefreshTokenExpiry)
	}

	if s.Res.SamlServiceProvider != nil {
		s.D.Set("saml_service_provider", []interface{}{AppSamlServiceProviderToMap(s.Res.SamlServiceProvider)})
	} else {
		s.D.Set("saml_service_provider", nil)
	}

	s.D.Set("schemas", s.Res.Schemas)

	scopes := []interface{}{}
	for _, item := range s.Res.Scopes {
		scopes = append(scopes, AppScopesToMap(item))
	}
	s.D.Set("scopes", scopes)

	s.D.Set("secondary_audiences", s.Res.SecondaryAudiences)

	serviceParams := []interface{}{}
	for _, item := range s.Res.ServiceParams {
		serviceParams = append(serviceParams, AppServiceParamsToMap(item))
	}
	s.D.Set("service_params", serviceParams)

	if s.Res.ServiceTypeURN != nil {
		s.D.Set("service_type_urn", *s.Res.ServiceTypeURN)
	}

	if s.Res.ServiceTypeVersion != nil {
		s.D.Set("service_type_version", *s.Res.ServiceTypeVersion)
	}

	if s.Res.ShowInMyApps != nil {
		s.D.Set("show_in_my_apps", *s.Res.ShowInMyApps)
	}

	if s.Res.SignonPolicy != nil {
		s.D.Set("signon_policy", []interface{}{AppSignonPolicyToMap(s.Res.SignonPolicy)})
	} else {
		s.D.Set("signon_policy", nil)
	}

	tags := []interface{}{}
	for _, item := range s.Res.Tags {
		tags = append(tags, tagsToMap(item))
	}
	s.D.Set("tags", tags)

	if s.Res.TenancyOcid != nil {
		s.D.Set("tenancy_ocid", *s.Res.TenancyOcid)
	}

	if s.Res.TermsOfServiceUrl != nil {
		s.D.Set("terms_of_service_url", *s.Res.TermsOfServiceUrl)
	}

	if s.Res.TermsOfUse != nil {
		s.D.Set("terms_of_use", []interface{}{AppTermsOfUseToMap(s.Res.TermsOfUse)})
	} else {
		s.D.Set("terms_of_use", nil)
	}

	trustPolicies := []interface{}{}
	for _, item := range s.Res.TrustPolicies {
		trustPolicies = append(trustPolicies, AppTrustPoliciesToMap(item))
	}
	s.D.Set("trust_policies", trustPolicies)

	s.D.Set("trust_scope", s.Res.TrustScope)

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", []interface{}{ExtensionOCITagsToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextension_oci_tags", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_app", []interface{}{AppExtensionDbcsAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensiondbcs_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", []interface{}{AppExtensionEnterpriseAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionenterprise_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", []interface{}{AppExtensionFormFillAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", []interface{}{AppExtensionFormFillAppTemplateAppTemplateToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", []interface{}{AppExtensionKerberosRealmAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", []interface{}{AppExtensionManagedappAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmanagedapp_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", []interface{}{AppExtensionMulticloudServiceAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionopc_service_app", []interface{}{AppExtensionOpcServiceAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionopc_service_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionradius_app_app", []interface{}{AppExtensionRadiusAppAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionradius_app_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_app", []interface{}{AppExtensionRequestableAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionrequestable_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", []interface{}{AppExtensionSamlServiceProviderAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app", nil)
	}

	if s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp != nil {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", []interface{}{AppExtensionWebTierPolicyAppToMap(s.Res.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp)})
	} else {
		s.D.Set("urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app", nil)
	}

	userRoles := []interface{}{}
	for _, item := range s.Res.UserRoles {
		userRoles = append(userRoles, AppUserRolesToMap(item))
	}
	s.D.Set("user_roles", userRoles)

	return nil
}

func GetAppCompositeId(appId string, idcsEndpoint string) string {
	appId = url.PathEscape(appId)
	//id = url.PathEscape(id)
	idcsEndpoint = url.PathEscape(idcsEndpoint)
	compositeId := "idcsEndpoint/" + idcsEndpoint + "/apps/" + appId
	return compositeId
}

func parseAppCompositeId(compositeId string) (appId string, err error) {
	parts := strings.Split(compositeId, "/")
	match, _ := regexp.MatchString("idcsEndpoint/.*/apps/.*", compositeId)
	if !match || len(parts) != 4 {
		err = fmt.Errorf("illegal compositeId %s encountered", compositeId)
		return
	}
	//idcsEndpoint, _ := url.PathUnescape(parts[1])
	appId, _ = url.PathUnescape(parts[3])

	return
}

func AppToMap(obj oci_identity_domains.App) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessTokenExpiry != nil {
		result["access_token_expiry"] = int(*obj.AccessTokenExpiry)
	}

	accounts := []interface{}{}
	for _, item := range obj.Accounts {
		accounts = append(accounts, AppAccountsToMap(item))
	}
	result["accounts"] = accounts

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	adminRoles := []interface{}{}
	for _, item := range obj.AdminRoles {
		adminRoles = append(adminRoles, AppAdminRolesToMap(item))
	}
	result["admin_roles"] = adminRoles

	aliasApps := []interface{}{}
	for _, item := range obj.AliasApps {
		aliasApps = append(aliasApps, AppAliasAppsToMap(item))
	}
	result["alias_apps"] = aliasApps

	if obj.AllUrlSchemesAllowed != nil {
		result["all_url_schemes_allowed"] = bool(*obj.AllUrlSchemesAllowed)
	}

	if obj.AllowAccessControl != nil {
		result["allow_access_control"] = bool(*obj.AllowAccessControl)
	}

	if obj.AllowOffline != nil {
		result["allow_offline"] = bool(*obj.AllowOffline)
	}

	result["allowed_grants"] = obj.AllowedGrants

	result["allowed_operations"] = obj.AllowedOperations

	allowedScopes := []interface{}{}
	for _, item := range obj.AllowedScopes {
		allowedScopes = append(allowedScopes, AppAllowedScopesToMap(item))
	}
	result["allowed_scopes"] = allowedScopes

	allowedTags := []interface{}{}
	for _, item := range obj.AllowedTags {
		allowedTags = append(allowedTags, AppAllowedTagsToMap(item))
	}
	result["allowed_tags"] = allowedTags

	if obj.AppIcon != nil {
		result["app_icon"] = string(*obj.AppIcon)
	}

	if obj.AppSignonPolicy != nil {
		result["app_signon_policy"] = []interface{}{AppAppSignonPolicyToMap(obj.AppSignonPolicy)}
	}

	if obj.AppThumbnail != nil {
		result["app_thumbnail"] = string(*obj.AppThumbnail)
	}

	appsNetworkPerimeters := []interface{}{}
	for _, item := range obj.AppsNetworkPerimeters {
		appsNetworkPerimeters = append(appsNetworkPerimeters, AppAppsNetworkPerimetersToMap(item))
	}
	result["apps_network_perimeters"] = appsNetworkPerimeters

	if obj.AsOPCService != nil {
		result["as_opc_service"] = []interface{}{AppAsOPCServiceToMap(obj.AsOPCService)}
	}

	attrRenderingMetadata := []interface{}{}
	for _, item := range obj.AttrRenderingMetadata {
		attrRenderingMetadata = append(attrRenderingMetadata, AppAttrRenderingMetadataToMap(item))
	}
	result["attr_rendering_metadata"] = attrRenderingMetadata

	if obj.Audience != nil {
		result["audience"] = string(*obj.Audience)
	}

	if obj.BasedOnTemplate != nil {
		result["based_on_template"] = []interface{}{AppBasedOnTemplateToMap(obj.BasedOnTemplate)}
	}

	if obj.BypassConsent != nil {
		result["bypass_consent"] = bool(*obj.BypassConsent)
	}

	if obj.CallbackServiceUrl != nil {
		result["callback_service_url"] = string(*obj.CallbackServiceUrl)
	}

	certificates := []interface{}{}
	for _, item := range obj.Certificates {
		certificates = append(certificates, AppCertificatesToMap(item))
	}
	result["certificates"] = certificates

	result["client_ip_checking"] = string(obj.ClientIPChecking)

	if obj.ClientSecret != nil {
		result["client_secret"] = string(*obj.ClientSecret)
	}

	result["client_type"] = string(obj.ClientType)

	cloudControlProperties := []interface{}{}
	for _, item := range obj.CloudControlProperties {
		cloudControlProperties = append(cloudControlProperties, AppCloudControlPropertiesToMap(item))
	}
	result["cloud_control_properties"] = cloudControlProperties

	if obj.CompartmentOcid != nil {
		result["compartment_ocid"] = string(*obj.CompartmentOcid)
	}

	if obj.ContactEmailAddress != nil {
		result["contact_email_address"] = string(*obj.ContactEmailAddress)
	}

	result["delegated_service_names"] = obj.DelegatedServiceNames

	if obj.DeleteInProgress != nil {
		result["delete_in_progress"] = bool(*obj.DeleteInProgress)
	}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisableKmsiTokenAuthentication != nil {
		result["disable_kmsi_token_authentication"] = bool(*obj.DisableKmsiTokenAuthentication)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.DomainOcid != nil {
		result["domain_ocid"] = string(*obj.DomainOcid)
	}

	editableAttributes := []interface{}{}
	for _, item := range obj.EditableAttributes {
		editableAttributes = append(editableAttributes, AppEditableAttributesToMap(item))
	}
	result["editable_attributes"] = editableAttributes

	if obj.ErrorPageUrl != nil {
		result["error_page_url"] = string(*obj.ErrorPageUrl)
	}

	grantedAppRoles := []interface{}{}
	for _, item := range obj.GrantedAppRoles {
		grantedAppRoles = append(grantedAppRoles, AppGrantedAppRolesToMap(item))
	}
	result["granted_app_roles"] = grantedAppRoles

	grants := []interface{}{}
	for _, item := range obj.Grants {
		grants = append(grants, AppGrantsToMap(item))
	}
	result["grants"] = grants

	if obj.HashedClientSecret != nil {
		result["hashed_client_secret"] = string(*obj.HashedClientSecret)
	}

	if obj.HomePageUrl != nil {
		result["home_page_url"] = string(*obj.HomePageUrl)
	}

	if obj.Icon != nil {
		result["icon"] = string(*obj.Icon)
	}

	if obj.Id != nil {
		result["id"] = string(*obj.Id)
	}

	if obj.IdTokenEncAlgo != nil {
		result["id_token_enc_algo"] = string(*obj.IdTokenEncAlgo)
	}

	if obj.IdcsCreatedBy != nil {
		result["idcs_created_by"] = []interface{}{idcsCreatedByToMap(obj.IdcsCreatedBy)}
	}

	if obj.IdcsLastModifiedBy != nil {
		result["idcs_last_modified_by"] = []interface{}{idcsLastModifiedByToMap(obj.IdcsLastModifiedBy)}
	}

	if obj.IdcsLastUpgradedInRelease != nil {
		result["idcs_last_upgraded_in_release"] = string(*obj.IdcsLastUpgradedInRelease)
	}

	result["idcs_prevented_operations"] = obj.IdcsPreventedOperations

	identityProviders := []interface{}{}
	for _, item := range obj.IdentityProviders {
		identityProviders = append(identityProviders, AppIdentityProvidersToMap(item))
	}
	result["identity_providers"] = identityProviders

	if obj.IdpPolicy != nil {
		result["idp_policy"] = []interface{}{AppIdpPolicyToMap(obj.IdpPolicy)}
	}

	if obj.Infrastructure != nil {
		result["infrastructure"] = bool(*obj.Infrastructure)
	}

	if obj.IsAliasApp != nil {
		result["is_alias_app"] = bool(*obj.IsAliasApp)
	}

	if obj.IsDatabaseService != nil {
		result["is_database_service"] = bool(*obj.IsDatabaseService)
	}

	if obj.IsEnterpriseApp != nil {
		result["is_enterprise_app"] = bool(*obj.IsEnterpriseApp)
	}

	if obj.IsFormFill != nil {
		result["is_form_fill"] = bool(*obj.IsFormFill)
	}

	if obj.IsKerberosRealm != nil {
		result["is_kerberos_realm"] = bool(*obj.IsKerberosRealm)
	}

	if obj.IsLoginTarget != nil {
		result["is_login_target"] = bool(*obj.IsLoginTarget)
	}

	if obj.IsManagedApp != nil {
		result["is_managed_app"] = bool(*obj.IsManagedApp)
	}

	if obj.IsMobileTarget != nil {
		result["is_mobile_target"] = bool(*obj.IsMobileTarget)
	}

	if obj.IsMulticloudServiceApp != nil {
		result["is_multicloud_service_app"] = bool(*obj.IsMulticloudServiceApp)
	}

	if obj.IsOAuthClient != nil {
		result["is_oauth_client"] = bool(*obj.IsOAuthClient)
	}

	if obj.IsOAuthResource != nil {
		result["is_oauth_resource"] = bool(*obj.IsOAuthResource)
	}

	if obj.IsOPCService != nil {
		result["is_opc_service"] = bool(*obj.IsOPCService)
	}

	if obj.IsObligationCapable != nil {
		result["is_obligation_capable"] = bool(*obj.IsObligationCapable)
	}

	if obj.IsRadiusApp != nil {
		result["is_radius_app"] = bool(*obj.IsRadiusApp)
	}

	if obj.IsSamlServiceProvider != nil {
		result["is_saml_service_provider"] = bool(*obj.IsSamlServiceProvider)
	}

	if obj.IsUnmanagedApp != nil {
		result["is_unmanaged_app"] = bool(*obj.IsUnmanagedApp)
	}

	if obj.IsWebTierPolicy != nil {
		result["is_web_tier_policy"] = bool(*obj.IsWebTierPolicy)
	}

	if obj.LandingPageUrl != nil {
		result["landing_page_url"] = string(*obj.LandingPageUrl)
	}

	if obj.LinkingCallbackUrl != nil {
		result["linking_callback_url"] = string(*obj.LinkingCallbackUrl)
	}

	result["login_mechanism"] = string(obj.LoginMechanism)

	if obj.LoginPageUrl != nil {
		result["login_page_url"] = string(*obj.LoginPageUrl)
	}

	if obj.LogoutPageUrl != nil {
		result["logout_page_url"] = string(*obj.LogoutPageUrl)
	}

	if obj.LogoutUri != nil {
		result["logout_uri"] = string(*obj.LogoutUri)
	}

	if obj.Meta != nil {
		result["meta"] = []interface{}{metaToMap(obj.Meta)}
	}

	if obj.MeterAsOPCService != nil {
		result["meter_as_opc_service"] = bool(*obj.MeterAsOPCService)
	}

	if obj.Migrated != nil {
		result["migrated"] = bool(*obj.Migrated)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ocid != nil {
		result["ocid"] = string(*obj.Ocid)
	}

	result["post_logout_redirect_uris"] = obj.PostLogoutRedirectUris

	if obj.PrivacyPolicyUrl != nil {
		result["privacy_policy_url"] = string(*obj.PrivacyPolicyUrl)
	}

	if obj.ProductLogoUrl != nil {
		result["product_logo_url"] = string(*obj.ProductLogoUrl)
	}

	if obj.ProductName != nil {
		result["product_name"] = string(*obj.ProductName)
	}

	protectableSecondaryAudiences := []interface{}{}
	for _, item := range obj.ProtectableSecondaryAudiences {
		protectableSecondaryAudiences = append(protectableSecondaryAudiences, AppProtectableSecondaryAudiencesToMap(item))
	}
	result["protectable_secondary_audiences"] = protectableSecondaryAudiences

	if obj.RadiusPolicy != nil {
		result["radius_policy"] = []interface{}{AppRadiusPolicyToMap(obj.RadiusPolicy)}
	}

	if obj.ReadyToUpgrade != nil {
		result["ready_to_upgrade"] = bool(*obj.ReadyToUpgrade)
	}

	result["redirect_uris"] = obj.RedirectUris

	if obj.RefreshTokenExpiry != nil {
		result["refresh_token_expiry"] = int(*obj.RefreshTokenExpiry)
	}

	if obj.SamlServiceProvider != nil {
		result["saml_service_provider"] = []interface{}{AppSamlServiceProviderToMap(obj.SamlServiceProvider)}
	}

	result["schemas"] = obj.Schemas

	scopes := []interface{}{}
	for _, item := range obj.Scopes {
		scopes = append(scopes, AppScopesToMap(item))
	}
	result["scopes"] = scopes

	result["secondary_audiences"] = obj.SecondaryAudiences

	serviceParams := []interface{}{}
	for _, item := range obj.ServiceParams {
		serviceParams = append(serviceParams, AppServiceParamsToMap(item))
	}
	result["service_params"] = serviceParams

	if obj.ServiceTypeURN != nil {
		result["service_type_urn"] = string(*obj.ServiceTypeURN)
	}

	if obj.ServiceTypeVersion != nil {
		result["service_type_version"] = string(*obj.ServiceTypeVersion)
	}

	if obj.ShowInMyApps != nil {
		result["show_in_my_apps"] = bool(*obj.ShowInMyApps)
	}

	if obj.SignonPolicy != nil {
		result["signon_policy"] = []interface{}{AppSignonPolicyToMap(obj.SignonPolicy)}
	}

	tags := []interface{}{}
	for _, item := range obj.Tags {
		tags = append(tags, tagsToMap(item))
	}
	result["tags"] = tags

	if obj.TenancyOcid != nil {
		result["tenancy_ocid"] = string(*obj.TenancyOcid)
	}

	if obj.TermsOfServiceUrl != nil {
		result["terms_of_service_url"] = string(*obj.TermsOfServiceUrl)
	}

	if obj.TermsOfUse != nil {
		result["terms_of_use"] = []interface{}{AppTermsOfUseToMap(obj.TermsOfUse)}
	}

	trustPolicies := []interface{}{}
	for _, item := range obj.TrustPolicies {
		trustPolicies = append(trustPolicies, AppTrustPoliciesToMap(item))
	}
	result["trust_policies"] = trustPolicies

	result["trust_scope"] = string(obj.TrustScope)

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags != nil {
		result["urnietfparamsscimschemasoracleidcsextension_oci_tags"] = []interface{}{ExtensionOCITagsToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOciTags)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensiondbcs_app"] = []interface{}{AppExtensionDbcsAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionDbcsApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionenterprise_app_app"] = []interface{}{AppExtensionEnterpriseAppAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionEnterpriseAppApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionform_fill_app_app"] = []interface{}{AppExtensionFormFillAppAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate != nil {
		result["urnietfparamsscimschemasoracleidcsextensionform_fill_app_template_app_template"] = []interface{}{AppExtensionFormFillAppTemplateAppTemplateToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionFormFillAppTemplateAppTemplate)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionkerberos_realm_app"] = []interface{}{AppExtensionKerberosRealmAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionKerberosRealmApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionmanagedapp_app"] = []interface{}{AppExtensionManagedappAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionManagedappApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionmulticloud_service_app_app"] = []interface{}{AppExtensionMulticloudServiceAppAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionMulticloudServiceAppApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionopc_service_app"] = []interface{}{AppExtensionOpcServiceAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionOpcServiceApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionradius_app_app"] = []interface{}{AppExtensionRadiusAppAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRadiusAppApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionrequestable_app"] = []interface{}{AppExtensionRequestableAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionRequestableApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionsaml_service_provider_app"] = []interface{}{AppExtensionSamlServiceProviderAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionSamlServiceProviderApp)}
	}

	if obj.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp != nil {
		result["urnietfparamsscimschemasoracleidcsextensionweb_tier_policy_app"] = []interface{}{AppExtensionWebTierPolicyAppToMap(obj.UrnIetfParamsScimSchemasOracleIdcsExtensionWebTierPolicyApp)}
	}

	userRoles := []interface{}{}
	for _, item := range obj.UserRoles {
		userRoles = append(userRoles, AppUserRolesToMap(item))
	}
	result["user_roles"] = userRoles

	return result
}

func AppAccountsToMap(obj oci_identity_domains.AppAccounts) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Active != nil {
		result["active"] = bool(*obj.Active)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.OwnerId != nil {
		result["owner_id"] = string(*obj.OwnerId)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AppAdminRolesToMap(obj oci_identity_domains.AppAdminRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAliasApps(fieldKeyFormat string) (oci_identity_domains.AppAliasApps, error) {
	result := oci_identity_domains.AppAliasApps{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAliasAppsToMap(obj oci_identity_domains.AppAliasApps) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAllowAuthzPolicy(fieldKeyFormat string) (oci_identity_domains.AppAllowAuthzPolicy, error) {
	result := oci_identity_domains.AppAllowAuthzPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAllowAuthzPolicyToMap(obj *oci_identity_domains.AppAllowAuthzPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAllowedScopes(fieldKeyFormat string) (oci_identity_domains.AppAllowedScopes, error) {
	result := oci_identity_domains.AppAllowedScopes{}

	if fqs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fqs")); ok {
		tmp := fqs.(string)
		result.Fqs = &tmp
	}

	if idOfDefiningApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "id_of_defining_app")); ok {
		tmp := idOfDefiningApp.(string)
		result.IdOfDefiningApp = &tmp
	}

	if readOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_only")); ok {
		tmp := readOnly.(bool)
		result.ReadOnly = &tmp
	}

	return result, nil
}

func AppAllowedScopesToMap(obj oci_identity_domains.AppAllowedScopes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Fqs != nil {
		result["fqs"] = string(*obj.Fqs)
	}

	if obj.IdOfDefiningApp != nil {
		result["id_of_defining_app"] = string(*obj.IdOfDefiningApp)
	}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAllowedTags(fieldKeyFormat string) (oci_identity_domains.AppAllowedTags, error) {
	result := oci_identity_domains.AppAllowedTags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if readOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_only")); ok {
		tmp := readOnly.(bool)
		result.ReadOnly = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAllowedTagsToMap(obj oci_identity_domains.AppAllowedTags) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Key != nil {
		result["key"] = string(*obj.Key)
	}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAppResources(fieldKeyFormat string) (oci_identity_domains.AppAppResources, error) {
	result := oci_identity_domains.AppAppResources{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAppResourcesToMap(obj oci_identity_domains.AppAppResources) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAppSignonPolicy(fieldKeyFormat string) (oci_identity_domains.AppAppSignonPolicy, error) {
	result := oci_identity_domains.AppAppSignonPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAppSignonPolicyToMap(obj *oci_identity_domains.AppAppSignonPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAppsNetworkPerimeters(fieldKeyFormat string) (oci_identity_domains.AppAppsNetworkPerimeters, error) {
	result := oci_identity_domains.AppAppsNetworkPerimeters{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAppsNetworkPerimetersToMap(obj oci_identity_domains.AppAppsNetworkPerimeters) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAsOPCService(fieldKeyFormat string) (oci_identity_domains.AppAsOpcService, error) {
	result := oci_identity_domains.AppAsOpcService{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppAsOPCServiceToMap(obj *oci_identity_domains.AppAsOpcService) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppAttrRenderingMetadata(fieldKeyFormat string) (oci_identity_domains.AppAttrRenderingMetadata, error) {
	result := oci_identity_domains.AppAttrRenderingMetadata{}

	if datatype, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "datatype")); ok {
		tmp := datatype.(string)
		result.Datatype = &tmp
	}

	if helptext, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "helptext")); ok {
		tmp := helptext.(string)
		result.Helptext = &tmp
	}

	if label, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "label")); ok {
		tmp := label.(string)
		result.Label = &tmp
	}

	if maxLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_length")); ok {
		tmp := maxLength.(int)
		result.MaxLength = &tmp
	}

	if maxSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_size")); ok {
		tmp := maxSize.(int)
		result.MaxSize = &tmp
	}

	if minLength, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_length")); ok {
		tmp := minLength.(int)
		result.MinLength = &tmp
	}

	if minSize, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_size")); ok {
		tmp := minSize.(int)
		result.MinSize = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if order, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "order")); ok {
		tmp := order.(int)
		result.Order = &tmp
	}

	if readOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_only")); ok {
		tmp := readOnly.(bool)
		result.ReadOnly = &tmp
	}

	if regexp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "regexp")); ok {
		tmp := regexp.(string)
		result.Regexp = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	if section, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "section")); ok {
		result.Section = oci_identity_domains.AppAttrRenderingMetadataSectionEnum(section.(string))
	}

	if visible, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "visible")); ok {
		tmp := visible.(bool)
		result.Visible = &tmp
	}

	if widget, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "widget")); ok {
		result.Widget = oci_identity_domains.AppAttrRenderingMetadataWidgetEnum(widget.(string))
	}

	return result, nil
}

func AppAttrRenderingMetadataToMap(obj oci_identity_domains.AppAttrRenderingMetadata) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Datatype != nil {
		result["datatype"] = string(*obj.Datatype)
	}

	if obj.Helptext != nil {
		result["helptext"] = string(*obj.Helptext)
	}

	if obj.Label != nil {
		result["label"] = string(*obj.Label)
	}

	if obj.MaxLength != nil {
		result["max_length"] = int(*obj.MaxLength)
	}

	if obj.MaxSize != nil {
		result["max_size"] = int(*obj.MaxSize)
	}

	if obj.MinLength != nil {
		result["min_length"] = int(*obj.MinLength)
	}

	if obj.MinSize != nil {
		result["min_size"] = int(*obj.MinSize)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Order != nil {
		result["order"] = int(*obj.Order)
	}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	if obj.Regexp != nil {
		result["regexp"] = string(*obj.Regexp)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	result["section"] = string(obj.Section)

	if obj.Visible != nil {
		result["visible"] = bool(*obj.Visible)
	}

	result["widget"] = string(obj.Widget)

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppBasedOnTemplate(fieldKeyFormat string) (oci_identity_domains.AppBasedOnTemplate, error) {
	result := oci_identity_domains.AppBasedOnTemplate{}

	if lastModified, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "last_modified")); ok {
		tmp := lastModified.(string)
		result.LastModified = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if wellKnownId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "well_known_id")); ok {
		tmp := wellKnownId.(string)
		result.WellKnownId = &tmp
	}

	return result, nil
}

func AppBasedOnTemplateToMap(obj *oci_identity_domains.AppBasedOnTemplate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.LastModified != nil {
		result["last_modified"] = string(*obj.LastModified)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.WellKnownId != nil {
		result["well_known_id"] = string(*obj.WellKnownId)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppBundleConfigurationProperties(fieldKeyFormat string) (oci_identity_domains.AppBundleConfigurationProperties, error) {
	result := oci_identity_domains.AppBundleConfigurationProperties{}

	if confidential, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidential")); ok {
		tmp := confidential.(bool)
		result.Confidential = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if helpMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "help_message")); ok {
		tmp := helpMessage.(string)
		result.HelpMessage = &tmp
	}

	if icfType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icf_type")); ok {
		result.IcfType = oci_identity_domains.AppBundleConfigurationPropertiesIcfTypeEnum(icfType.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if order, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "order")); ok {
		tmp := order.(int)
		result.Order = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		interfaces := value.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func AppBundleConfigurationPropertiesToMap(obj oci_identity_domains.AppBundleConfigurationProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Confidential != nil {
		result["confidential"] = bool(*obj.Confidential)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HelpMessage != nil {
		result["help_message"] = string(*obj.HelpMessage)
	}

	result["icf_type"] = string(obj.IcfType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Order != nil {
		result["order"] = int(*obj.Order)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	result["value"] = obj.Value

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppBundlePoolConfiguration(fieldKeyFormat string) (oci_identity_domains.AppBundlePoolConfiguration, error) {
	result := oci_identity_domains.AppBundlePoolConfiguration{}

	if maxIdle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_idle")); ok {
		tmp := maxIdle.(int)
		result.MaxIdle = &tmp
	}

	if maxObjects, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_objects")); ok {
		tmp := maxObjects.(int)
		result.MaxObjects = &tmp
	}

	if maxWait, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_wait")); ok {
		tmp := maxWait.(int)
		result.MaxWait = &tmp
	}

	if minEvictableIdleTimeMillis, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_evictable_idle_time_millis")); ok {
		tmp := minEvictableIdleTimeMillis.(int)
		result.MinEvictableIdleTimeMillis = &tmp
	}

	if minIdle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "min_idle")); ok {
		tmp := minIdle.(int)
		result.MinIdle = &tmp
	}

	return result, nil
}

func AppBundlePoolConfigurationToMap(obj *oci_identity_domains.AppBundlePoolConfiguration) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MaxIdle != nil {
		result["max_idle"] = int(*obj.MaxIdle)
	}

	if obj.MaxObjects != nil {
		result["max_objects"] = int(*obj.MaxObjects)
	}

	if obj.MaxWait != nil {
		result["max_wait"] = int(*obj.MaxWait)
	}

	if obj.MinEvictableIdleTimeMillis != nil {
		result["min_evictable_idle_time_millis"] = int(*obj.MinEvictableIdleTimeMillis)
	}

	if obj.MinIdle != nil {
		result["min_idle"] = int(*obj.MinIdle)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppCertificates(fieldKeyFormat string) (oci_identity_domains.AppCertificates, error) {
	result := oci_identity_domains.AppCertificates{}

	if certAlias, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "cert_alias")); ok {
		tmp := certAlias.(string)
		result.CertAlias = &tmp
	}

	return result, nil
}

func AppCertificatesToMap(obj oci_identity_domains.AppCertificates) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CertAlias != nil {
		result["cert_alias"] = string(*obj.CertAlias)
	}

	if obj.Kid != nil {
		result["kid"] = string(*obj.Kid)
	}

	if obj.Sha1Thumbprint != nil {
		result["sha1thumbprint"] = string(*obj.Sha1Thumbprint)
	}

	if obj.X509Base64Certificate != nil {
		result["x509base64certificate"] = fmt.Sprintf("%v", *obj.X509Base64Certificate)
	}

	if obj.X5t != nil {
		result["x5t"] = string(*obj.X5t)
	}

	return result
}

func AppCloudControlPropertiesToMap(obj oci_identity_domains.AppCloudControlProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	result["values"] = obj.Values

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppConnectorBundle(fieldKeyFormat string) (oci_identity_domains.AppConnectorBundle, error) {
	result := oci_identity_domains.AppConnectorBundle{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.AppConnectorBundleTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if wellKnownId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "well_known_id")); ok {
		tmp := wellKnownId.(string)
		result.WellKnownId = &tmp
	}

	return result, nil
}

func AppConnectorBundleToMap(obj *oci_identity_domains.AppConnectorBundle) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.WellKnownId != nil {
		result["well_known_id"] = string(*obj.WellKnownId)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppDenyAuthzPolicy(fieldKeyFormat string) (oci_identity_domains.AppDenyAuthzPolicy, error) {
	result := oci_identity_domains.AppDenyAuthzPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppDenyAuthzPolicyToMap(obj *oci_identity_domains.AppDenyAuthzPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppDomainApp(fieldKeyFormat string) (oci_identity_domains.AppDomainApp, error) {
	result := oci_identity_domains.AppDomainApp{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppDomainAppToMap(obj *oci_identity_domains.AppDomainApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AppEditableAttributesToMap(obj oci_identity_domains.AppEditableAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionDbcsApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionDbcsApp, error) {
	result := oci_identity_domains.AppExtensionDbcsApp{}

	if domainApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_app")); ok {
		if tmpList := domainApp.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "domain_app"), 0)
			tmp, err := s.mapToAppDomainApp(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert domain_app, encountered error: %v", err)
			}
			result.DomainApp = &tmp
		}
	}

	if domainName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "domain_name")); ok {
		tmp := domainName.(string)
		result.DomainName = &tmp
	}

	return result, nil
}

func AppExtensionDbcsAppToMap(obj *oci_identity_domains.AppExtensionDbcsApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DomainApp != nil {
		result["domain_app"] = []interface{}{AppDomainAppToMap(obj.DomainApp)}
	}

	if obj.DomainName != nil {
		result["domain_name"] = string(*obj.DomainName)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionEnterpriseAppApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionEnterpriseAppApp, error) {
	result := oci_identity_domains.AppExtensionEnterpriseAppApp{}

	if allowAuthzDecisionTTL, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_authz_decision_ttl")); ok {
		tmp := allowAuthzDecisionTTL.(int)
		result.AllowAuthzDecisionTTL = &tmp
	}

	if allowAuthzPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "allow_authz_policy")); ok {
		if tmpList := allowAuthzPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "allow_authz_policy"), 0)
			tmp, err := s.mapToAppAllowAuthzPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert allow_authz_policy, encountered error: %v", err)
			}
			result.AllowAuthzPolicy = &tmp
		}
	}

	if appResources, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "app_resources")); ok {
		interfaces := appResources.([]interface{})
		tmp := make([]oci_identity_domains.AppAppResources, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "app_resources"), stateDataIndex)
			converted, err := s.mapToAppAppResources(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "app_resources")) {
			result.AppResources = tmp
		}
	}

	if denyAuthzDecisionTTL, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deny_authz_decision_ttl")); ok {
		tmp := denyAuthzDecisionTTL.(int)
		result.DenyAuthzDecisionTTL = &tmp
	}

	if denyAuthzPolicy, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "deny_authz_policy")); ok {
		if tmpList := denyAuthzPolicy.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "deny_authz_policy"), 0)
			tmp, err := s.mapToAppDenyAuthzPolicy(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert deny_authz_policy, encountered error: %v", err)
			}
			result.DenyAuthzPolicy = &tmp
		}
	}

	return result, nil
}

func AppExtensionEnterpriseAppAppToMap(obj *oci_identity_domains.AppExtensionEnterpriseAppApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AllowAuthzDecisionTTL != nil {
		result["allow_authz_decision_ttl"] = int(*obj.AllowAuthzDecisionTTL)
	}

	if obj.AllowAuthzPolicy != nil {
		result["allow_authz_policy"] = []interface{}{AppAllowAuthzPolicyToMap(obj.AllowAuthzPolicy)}
	}

	appResources := []interface{}{}
	for _, item := range obj.AppResources {
		appResources = append(appResources, AppAppResourcesToMap(item))
	}
	result["app_resources"] = appResources

	if obj.DenyAuthzDecisionTTL != nil {
		result["deny_authz_decision_ttl"] = int(*obj.DenyAuthzDecisionTTL)
	}

	if obj.DenyAuthzPolicy != nil {
		result["deny_authz_policy"] = []interface{}{AppDenyAuthzPolicyToMap(obj.DenyAuthzPolicy)}
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionFormFillAppApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionFormFillAppApp, error) {
	result := oci_identity_domains.AppExtensionFormFillAppApp{}

	if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "configuration")) {
			tmp := configuration.(string)
			result.Configuration = &tmp
		}
	}

	if formCredMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_cred_method")); ok {
		result.FormCredMethod = oci_identity_domains.AppExtensionFormFillAppAppFormCredMethodEnum(formCredMethod.(string))
	}

	if formCredentialSharingGroupID, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_credential_sharing_group_id")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "form_credential_sharing_group_id")) {
			tmp := formCredentialSharingGroupID.(string)
			result.FormCredentialSharingGroupID = &tmp
		}
	}

	if formFillUrlMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_fill_url_match")); ok {
		interfaces := formFillUrlMatch.([]interface{})
		tmp := make([]oci_identity_domains.AppFormFillUrlMatch, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "form_fill_url_match"), stateDataIndex)
			converted, err := s.mapToAppFormFillUrlMatch(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "form_fill_url_match")) {
			result.FormFillUrlMatch = tmp
		}
	}

	if formType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_type")); ok {
		result.FormType = oci_identity_domains.AppExtensionFormFillAppAppFormTypeEnum(formType.(string))
	}

	if revealPasswordOnForm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reveal_password_on_form")); ok {
		tmp := revealPasswordOnForm.(bool)
		result.RevealPasswordOnForm = &tmp
	}

	if syncFromTemplate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sync_from_template")); ok {
		tmp := syncFromTemplate.(bool)
		result.SyncFromTemplate = &tmp
	}

	if userNameFormExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name_form_expression")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "user_name_form_expression")) {
			tmp := userNameFormExpression.(string)
			result.UserNameFormExpression = &tmp
		}
	}

	if userNameFormTemplate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name_form_template")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "user_name_form_template")) {
			tmp := userNameFormTemplate.(string)
			result.UserNameFormTemplate = &tmp
		}
	}

	return result, nil
}

func AppExtensionFormFillAppAppToMap(obj *oci_identity_domains.AppExtensionFormFillAppApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Configuration != nil {
		result["configuration"] = string(*obj.Configuration)
	}

	result["form_cred_method"] = string(obj.FormCredMethod)

	if obj.FormCredentialSharingGroupID != nil {
		result["form_credential_sharing_group_id"] = string(*obj.FormCredentialSharingGroupID)
	}

	formFillUrlMatch := []interface{}{}
	for _, item := range obj.FormFillUrlMatch {
		formFillUrlMatch = append(formFillUrlMatch, AppFormFillUrlMatchToMap(item))
	}
	result["form_fill_url_match"] = formFillUrlMatch

	result["form_type"] = string(obj.FormType)

	if obj.RevealPasswordOnForm != nil {
		result["reveal_password_on_form"] = bool(*obj.RevealPasswordOnForm)
	}

	if obj.SyncFromTemplate != nil {
		result["sync_from_template"] = bool(*obj.SyncFromTemplate)
	}

	if obj.UserNameFormExpression != nil {
		result["user_name_form_expression"] = string(*obj.UserNameFormExpression)
	}

	if obj.UserNameFormTemplate != nil {
		result["user_name_form_template"] = string(*obj.UserNameFormTemplate)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionFormFillAppTemplateAppTemplate(fieldKeyFormat string) (oci_identity_domains.AppExtensionFormFillAppTemplateAppTemplate, error) {
	result := oci_identity_domains.AppExtensionFormFillAppTemplateAppTemplate{}

	if configuration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "configuration")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "configuration")) {
			tmp := configuration.(string)
			result.Configuration = &tmp
		}
	}

	if formCredMethod, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_cred_method")); ok {
		result.FormCredMethod = oci_identity_domains.AppExtensionFormFillAppTemplateAppTemplateFormCredMethodEnum(formCredMethod.(string))
	}

	if formCredentialSharingGroupID, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_credential_sharing_group_id")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "form_credential_sharing_group_id")) {
			tmp := formCredentialSharingGroupID.(string)
			result.FormCredentialSharingGroupID = &tmp
		}
	}

	if formFillUrlMatch, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_fill_url_match")); ok {
		interfaces := formFillUrlMatch.([]interface{})
		tmp := make([]oci_identity_domains.AppFormFillUrlMatch, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "form_fill_url_match"), stateDataIndex)
			converted, err := s.mapToAppFormFillUrlMatch(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "form_fill_url_match")) {
			result.FormFillUrlMatch = tmp
		}
	}

	if formType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_type")); ok {
		result.FormType = oci_identity_domains.AppExtensionFormFillAppTemplateAppTemplateFormTypeEnum(formType.(string))
	}

	if revealPasswordOnForm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "reveal_password_on_form")); ok {
		tmp := revealPasswordOnForm.(bool)
		result.RevealPasswordOnForm = &tmp
	}

	if syncFromTemplate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sync_from_template")); ok {
		tmp := syncFromTemplate.(bool)
		result.SyncFromTemplate = &tmp
	}

	if userNameFormExpression, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name_form_expression")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "user_name_form_expression")) {
			tmp := userNameFormExpression.(string)
			result.UserNameFormExpression = &tmp
		}
	}

	if userNameFormTemplate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_name_form_template")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "user_name_form_template")) {
			tmp := userNameFormTemplate.(string)
			result.UserNameFormTemplate = &tmp
		}
	}

	return result, nil
}

func AppExtensionFormFillAppTemplateAppTemplateToMap(obj *oci_identity_domains.AppExtensionFormFillAppTemplateAppTemplate) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Configuration != nil {
		result["configuration"] = string(*obj.Configuration)
	}

	result["form_cred_method"] = string(obj.FormCredMethod)

	if obj.FormCredentialSharingGroupID != nil {
		result["form_credential_sharing_group_id"] = string(*obj.FormCredentialSharingGroupID)
	}

	formFillUrlMatch := []interface{}{}
	for _, item := range obj.FormFillUrlMatch {
		formFillUrlMatch = append(formFillUrlMatch, AppFormFillUrlMatchToMap(item))
	}
	result["form_fill_url_match"] = formFillUrlMatch

	result["form_type"] = string(obj.FormType)

	if obj.RevealPasswordOnForm != nil {
		result["reveal_password_on_form"] = bool(*obj.RevealPasswordOnForm)
	}

	if obj.SyncFromTemplate != nil {
		result["sync_from_template"] = bool(*obj.SyncFromTemplate)
	}

	if obj.UserNameFormExpression != nil {
		result["user_name_form_expression"] = string(*obj.UserNameFormExpression)
	}

	if obj.UserNameFormTemplate != nil {
		result["user_name_form_template"] = string(*obj.UserNameFormTemplate)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionKerberosRealmApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionKerberosRealmApp, error) {
	result := oci_identity_domains.AppExtensionKerberosRealmApp{}

	if defaultEncryptionSaltType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "default_encryption_salt_type")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "default_encryption_salt_type")) {
			tmp := defaultEncryptionSaltType.(string)
			result.DefaultEncryptionSaltType = &tmp
		}
	}

	if masterKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "master_key")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "master_key")) {
			tmp := masterKey.(string)
			result.MasterKey = &tmp
		}
	}

	if maxRenewableAge, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_renewable_age")); ok {
		tmp := maxRenewableAge.(int)
		result.MaxRenewableAge = &tmp
	}

	if maxTicketLife, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "max_ticket_life")); ok {
		tmp := maxTicketLife.(int)
		result.MaxTicketLife = &tmp
	}

	if realmName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "realm_name")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "realm_name")) {
			tmp := realmName.(string)
			result.RealmName = &tmp
		}
	}

	if supportedEncryptionSaltTypes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "supported_encryption_salt_types")); ok {
		interfaces := supportedEncryptionSaltTypes.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "supported_encryption_salt_types")) {
			result.SupportedEncryptionSaltTypes = tmp
		}
	}

	if ticketFlags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ticket_flags")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "ticket_flags")) {
			tmp := ticketFlags.(int)
			result.TicketFlags = &tmp
		}
	}

	return result, nil
}

func AppExtensionKerberosRealmAppToMap(obj *oci_identity_domains.AppExtensionKerberosRealmApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.DefaultEncryptionSaltType != nil {
		result["default_encryption_salt_type"] = string(*obj.DefaultEncryptionSaltType)
	}

	if obj.MasterKey != nil {
		result["master_key"] = string(*obj.MasterKey)
	}

	if obj.MaxRenewableAge != nil {
		result["max_renewable_age"] = int(*obj.MaxRenewableAge)
	}

	if obj.MaxTicketLife != nil {
		result["max_ticket_life"] = int(*obj.MaxTicketLife)
	}

	if obj.RealmName != nil {
		result["realm_name"] = string(*obj.RealmName)
	}

	result["supported_encryption_salt_types"] = obj.SupportedEncryptionSaltTypes

	if obj.TicketFlags != nil {
		result["ticket_flags"] = int(*obj.TicketFlags)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionManagedappApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionManagedappApp, error) {
	result := oci_identity_domains.AppExtensionManagedappApp{}

	if accountFormVisible, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "account_form_visible")); ok {
		tmp := accountFormVisible.(bool)
		result.AccountFormVisible = &tmp
	}

	if adminConsentGranted, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "admin_consent_granted")); ok {
		tmp := adminConsentGranted.(bool)
		result.AdminConsentGranted = &tmp
	}

	if bundleConfigurationProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bundle_configuration_properties")); ok {
		interfaces := bundleConfigurationProperties.([]interface{})
		tmp := make([]oci_identity_domains.AppBundleConfigurationProperties, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "bundle_configuration_properties"), stateDataIndex)
			converted, err := s.mapToAppBundleConfigurationProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "bundle_configuration_properties")) {
			result.BundleConfigurationProperties = tmp
		}
	}

	if bundlePoolConfiguration, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "bundle_pool_configuration")); ok {
		if tmpList := bundlePoolConfiguration.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "bundle_pool_configuration"), 0)
			tmp, err := s.mapToAppBundlePoolConfiguration(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert bundle_pool_configuration, encountered error: %v", err)
			}
			result.BundlePoolConfiguration = &tmp
		}
	}

	if canBeAuthoritative, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "can_be_authoritative")); ok {
		tmp := canBeAuthoritative.(bool)
		result.CanBeAuthoritative = &tmp
	}

	if connected, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connected")); ok {
		tmp := connected.(bool)
		result.Connected = &tmp
	}

	if connectorBundle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "connector_bundle")); ok {
		if tmpList := connectorBundle.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "connector_bundle"), 0)
			tmp, err := s.mapToAppConnectorBundle(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert connector_bundle, encountered error: %v", err)
			}
			result.ConnectorBundle = &tmp
		}
	}

	if enableAuthSyncNewUserNotification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_auth_sync_new_user_notification")); ok {
		tmp := enableAuthSyncNewUserNotification.(bool)
		result.EnableAuthSyncNewUserNotification = &tmp
	}

	if enableSync, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_sync")); ok {
		tmp := enableSync.(bool)
		result.EnableSync = &tmp
	}

	if enableSyncSummaryReportNotification, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enable_sync_summary_report_notification")); ok {
		tmp := enableSyncSummaryReportNotification.(bool)
		result.EnableSyncSummaryReportNotification = &tmp
	}

	if flatFileBundleConfigurationProperties, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "flat_file_bundle_configuration_properties")); ok {
		interfaces := flatFileBundleConfigurationProperties.([]interface{})
		tmp := make([]oci_identity_domains.AppFlatFileBundleConfigurationProperties, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "flat_file_bundle_configuration_properties"), stateDataIndex)
			converted, err := s.mapToAppFlatFileBundleConfigurationProperties(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "flat_file_bundle_configuration_properties")) {
			result.FlatFileBundleConfigurationProperties = tmp
		}
	}

	if flatFileConnectorBundle, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "flat_file_connector_bundle")); ok {
		if tmpList := flatFileConnectorBundle.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "flat_file_connector_bundle"), 0)
			tmp, err := s.mapToAppFlatFileConnectorBundle(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert flat_file_connector_bundle, encountered error: %v", err)
			}
			result.FlatFileConnectorBundle = &tmp
		}
	}

	if identityBridges, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "identity_bridges")); ok {
		interfaces := identityBridges.([]interface{})
		tmp := make([]oci_identity_domains.AppIdentityBridges, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "identity_bridges"), stateDataIndex)
			converted, err := s.mapToAppIdentityBridges(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "identity_bridges")) {
			result.IdentityBridges = tmp
		}
	}

	if isAuthoritative, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_authoritative")); ok {
		tmp := isAuthoritative.(bool)
		result.IsAuthoritative = &tmp
	}

	if isDirectory, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_directory")); ok {
		tmp := isDirectory.(bool)
		result.IsDirectory = &tmp
	}

	if isOnPremiseApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_on_premise_app")); ok {
		tmp := isOnPremiseApp.(bool)
		result.IsOnPremiseApp = &tmp
	}

	if isSchemaCustomizationSupported, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_schema_customization_supported")); ok {
		tmp := isSchemaCustomizationSupported.(bool)
		result.IsSchemaCustomizationSupported = &tmp
	}

	if isSchemaDiscoverySupported, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_schema_discovery_supported")); ok {
		tmp := isSchemaDiscoverySupported.(bool)
		result.IsSchemaDiscoverySupported = &tmp
	}

	if isThreeLeggedOAuthEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_three_legged_oauth_enabled")); ok {
		tmp := isThreeLeggedOAuthEnabled.(bool)
		result.IsThreeLeggedOAuthEnabled = &tmp
	}

	if isTwoLeggedOAuthEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_two_legged_oauth_enabled")); ok {
		tmp := isTwoLeggedOAuthEnabled.(bool)
		result.IsTwoLeggedOAuthEnabled = &tmp
	}

	if objectClasses, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "object_classes")); ok {
		interfaces := objectClasses.([]interface{})
		tmp := make([]oci_identity_domains.AppObjectClasses, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "object_classes"), stateDataIndex)
			converted, err := s.mapToAppObjectClasses(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "object_classes")) {
			result.ObjectClasses = tmp
		}
	}

	if threeLeggedOAuthCredential, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "three_legged_oauth_credential")); ok {
		if tmpList := threeLeggedOAuthCredential.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "three_legged_oauth_credential"), 0)
			tmp, err := s.mapToAppThreeLeggedOAuthCredential(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert three_legged_oauth_credential, encountered error: %v", err)
			}
			result.ThreeLeggedOAuthCredential = &tmp
		}
	}

	if threeLeggedOAuthProviderName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "three_legged_oauth_provider_name")); ok {
		tmp := threeLeggedOAuthProviderName.(string)
		result.ThreeLeggedOAuthProviderName = &tmp
	}

	return result, nil
}

func AppExtensionManagedappAppToMap(obj *oci_identity_domains.AppExtensionManagedappApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccountFormVisible != nil {
		result["account_form_visible"] = bool(*obj.AccountFormVisible)
	}

	if obj.AdminConsentGranted != nil {
		result["admin_consent_granted"] = bool(*obj.AdminConsentGranted)
	}

	bundleConfigurationProperties := []interface{}{}
	for _, item := range obj.BundleConfigurationProperties {
		bundleConfigurationProperties = append(bundleConfigurationProperties, AppBundleConfigurationPropertiesToMap(item))
	}
	result["bundle_configuration_properties"] = bundleConfigurationProperties

	if obj.BundlePoolConfiguration != nil {
		result["bundle_pool_configuration"] = []interface{}{AppBundlePoolConfigurationToMap(obj.BundlePoolConfiguration)}
	}

	if obj.CanBeAuthoritative != nil {
		result["can_be_authoritative"] = bool(*obj.CanBeAuthoritative)
	}

	if obj.Connected != nil {
		result["connected"] = bool(*obj.Connected)
	}

	if obj.ConnectorBundle != nil {
		result["connector_bundle"] = []interface{}{AppConnectorBundleToMap(obj.ConnectorBundle)}
	}

	if obj.EnableAuthSyncNewUserNotification != nil {
		result["enable_auth_sync_new_user_notification"] = bool(*obj.EnableAuthSyncNewUserNotification)
	}

	if obj.EnableSync != nil {
		result["enable_sync"] = bool(*obj.EnableSync)
	}

	if obj.EnableSyncSummaryReportNotification != nil {
		result["enable_sync_summary_report_notification"] = bool(*obj.EnableSyncSummaryReportNotification)
	}

	flatFileBundleConfigurationProperties := []interface{}{}
	for _, item := range obj.FlatFileBundleConfigurationProperties {
		flatFileBundleConfigurationProperties = append(flatFileBundleConfigurationProperties, AppFlatFileBundleConfigurationPropertiesToMap(item))
	}
	result["flat_file_bundle_configuration_properties"] = flatFileBundleConfigurationProperties

	if obj.FlatFileConnectorBundle != nil {
		result["flat_file_connector_bundle"] = []interface{}{AppFlatFileConnectorBundleToMap(obj.FlatFileConnectorBundle)}
	}

	identityBridges := []interface{}{}
	for _, item := range obj.IdentityBridges {
		identityBridges = append(identityBridges, AppIdentityBridgesToMap(item))
	}
	result["identity_bridges"] = identityBridges

	if obj.IsAuthoritative != nil {
		result["is_authoritative"] = bool(*obj.IsAuthoritative)
	}

	if obj.IsDirectory != nil {
		result["is_directory"] = bool(*obj.IsDirectory)
	}

	if obj.IsOnPremiseApp != nil {
		result["is_on_premise_app"] = bool(*obj.IsOnPremiseApp)
	}

	if obj.IsSchemaCustomizationSupported != nil {
		result["is_schema_customization_supported"] = bool(*obj.IsSchemaCustomizationSupported)
	}

	if obj.IsSchemaDiscoverySupported != nil {
		result["is_schema_discovery_supported"] = bool(*obj.IsSchemaDiscoverySupported)
	}

	if obj.IsThreeLeggedOAuthEnabled != nil {
		result["is_three_legged_oauth_enabled"] = bool(*obj.IsThreeLeggedOAuthEnabled)
	}

	if obj.IsTwoLeggedOAuthEnabled != nil {
		result["is_two_legged_oauth_enabled"] = bool(*obj.IsTwoLeggedOAuthEnabled)
	}

	objectClasses := []interface{}{}
	for _, item := range obj.ObjectClasses {
		objectClasses = append(objectClasses, AppObjectClassesToMap(item))
	}
	result["object_classes"] = objectClasses

	if obj.SyncConfigLastModified != nil {
		result["sync_config_last_modified"] = string(*obj.SyncConfigLastModified)
	}

	if obj.ThreeLeggedOAuthCredential != nil {
		result["three_legged_oauth_credential"] = []interface{}{AppThreeLeggedOAuthCredentialToMap(obj.ThreeLeggedOAuthCredential)}
	}

	if obj.ThreeLeggedOAuthProviderName != nil {
		result["three_legged_oauth_provider_name"] = string(*obj.ThreeLeggedOAuthProviderName)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionMulticloudServiceAppApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionMulticloudServiceAppApp, error) {
	result := oci_identity_domains.AppExtensionMulticloudServiceAppApp{}

	if multicloudPlatformUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "multicloud_platform_url")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "multicloud_platform_url")) {
			tmp := multicloudPlatformUrl.(string)
			result.MulticloudPlatformUrl = &tmp
		}
	}

	if multicloudServiceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "multicloud_service_type")); ok {
		result.MulticloudServiceType = oci_identity_domains.AppExtensionMulticloudServiceAppAppMulticloudServiceTypeEnum(multicloudServiceType.(string))
	}

	return result, nil
}

func AppExtensionMulticloudServiceAppAppToMap(obj *oci_identity_domains.AppExtensionMulticloudServiceAppApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.MulticloudPlatformUrl != nil {
		result["multicloud_platform_url"] = string(*obj.MulticloudPlatformUrl)
	}

	result["multicloud_service_type"] = string(obj.MulticloudServiceType)

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionOpcServiceApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionOpcServiceApp, error) {
	result := oci_identity_domains.AppExtensionOpcServiceApp{}

	if currentFederationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "current_federation_mode")); ok {
		result.CurrentFederationMode = oci_identity_domains.AppExtensionOpcServiceAppCurrentFederationModeEnum(currentFederationMode.(string))
	}

	if currentSynchronizationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "current_synchronization_mode")); ok {
		result.CurrentSynchronizationMode = oci_identity_domains.AppExtensionOpcServiceAppCurrentSynchronizationModeEnum(currentSynchronizationMode.(string))
	}

	if enablingNextFedSyncModes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "enabling_next_fed_sync_modes")); ok {
		tmp := enablingNextFedSyncModes.(bool)
		result.EnablingNextFedSyncModes = &tmp
	}

	if nextFederationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "next_federation_mode")); ok {
		result.NextFederationMode = oci_identity_domains.AppExtensionOpcServiceAppNextFederationModeEnum(nextFederationMode.(string))
	}

	if nextSynchronizationMode, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "next_synchronization_mode")); ok {
		result.NextSynchronizationMode = oci_identity_domains.AppExtensionOpcServiceAppNextSynchronizationModeEnum(nextSynchronizationMode.(string))
	}

	if region, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "region")); ok {
		tmp := region.(string)
		result.Region = &tmp
	}

	if serviceInstanceIdentifier, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "service_instance_identifier")); ok {
		tmp := serviceInstanceIdentifier.(string)
		result.ServiceInstanceIdentifier = &tmp
	}

	return result, nil
}

func AppExtensionOpcServiceAppToMap(obj *oci_identity_domains.AppExtensionOpcServiceApp) map[string]interface{} {
	result := map[string]interface{}{}

	result["current_federation_mode"] = string(obj.CurrentFederationMode)

	result["current_synchronization_mode"] = string(obj.CurrentSynchronizationMode)

	if obj.EnablingNextFedSyncModes != nil {
		result["enabling_next_fed_sync_modes"] = bool(*obj.EnablingNextFedSyncModes)
	}

	result["next_federation_mode"] = string(obj.NextFederationMode)

	result["next_synchronization_mode"] = string(obj.NextSynchronizationMode)

	if obj.Region != nil {
		result["region"] = string(*obj.Region)
	}

	if obj.ServiceInstanceIdentifier != nil {
		result["service_instance_identifier"] = string(*obj.ServiceInstanceIdentifier)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionRadiusAppApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionRadiusAppApp, error) {
	result := oci_identity_domains.AppExtensionRadiusAppApp{}

	if captureClientIp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "capture_client_ip")); ok {
		tmp := captureClientIp.(bool)
		result.CaptureClientIp = &tmp
	}

	if clientIP, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "client_ip")); ok {
		tmp := clientIP.(string)
		result.ClientIP = &tmp
	}

	if countryCodeResponseAttributeId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "country_code_response_attribute_id")); ok {
		tmp := countryCodeResponseAttributeId.(string)
		result.CountryCodeResponseAttributeId = &tmp
	}

	if endUserIPAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "end_user_ip_attribute")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "end_user_ip_attribute")) {
			tmp := endUserIPAttribute.(string)
			result.EndUserIPAttribute = &tmp
		}
	}

	if groupMembershipRadiusAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_membership_radius_attribute")); ok {
		tmp := groupMembershipRadiusAttribute.(string)
		result.GroupMembershipRadiusAttribute = &tmp
	}

	if groupMembershipToReturn, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_membership_to_return")); ok {
		interfaces := groupMembershipToReturn.([]interface{})
		tmp := make([]oci_identity_domains.AppGroupMembershipToReturn, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "group_membership_to_return"), stateDataIndex)
			converted, err := s.mapToAppGroupMembershipToReturn(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_membership_to_return")) {
			result.GroupMembershipToReturn = tmp
		}
	}

	if groupNameFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_name_format")); ok {
		tmp := groupNameFormat.(string)
		result.GroupNameFormat = &tmp
	}

	if includeGroupInResponse, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_group_in_response")); ok {
		tmp := includeGroupInResponse.(bool)
		result.IncludeGroupInResponse = &tmp
	}

	if passwordAndOtpTogether, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "password_and_otp_together")); ok {
		tmp := passwordAndOtpTogether.(bool)
		result.PasswordAndOtpTogether = &tmp
	}

	if port, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "port")); ok {
		tmp := port.(string)
		result.Port = &tmp
	}

	if radiusVendorSpecificId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "radius_vendor_specific_id")); ok {
		tmp := radiusVendorSpecificId.(string)
		result.RadiusVendorSpecificId = &tmp
	}

	if responseFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_format")); ok {
		tmp := responseFormat.(string)
		result.ResponseFormat = &tmp
	}

	if responseFormatDelimiter, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "response_format_delimiter")); ok {
		tmp := responseFormatDelimiter.(string)
		result.ResponseFormatDelimiter = &tmp
	}

	if secretKey, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "secret_key")); ok {
		tmp := secretKey.(string)
		result.SecretKey = &tmp
	}

	if typeOfRadiusApp, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type_of_radius_app")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "type_of_radius_app")) {
			tmp := typeOfRadiusApp.(string)
			result.TypeOfRadiusApp = &tmp
		}
	}

	return result, nil
}

func AppExtensionRadiusAppAppToMap(obj *oci_identity_domains.AppExtensionRadiusAppApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.CaptureClientIp != nil {
		result["capture_client_ip"] = bool(*obj.CaptureClientIp)
	}

	if obj.ClientIP != nil {
		result["client_ip"] = string(*obj.ClientIP)
	}

	if obj.CountryCodeResponseAttributeId != nil {
		result["country_code_response_attribute_id"] = string(*obj.CountryCodeResponseAttributeId)
	}

	if obj.EndUserIPAttribute != nil {
		result["end_user_ip_attribute"] = string(*obj.EndUserIPAttribute)
	}

	if obj.GroupMembershipRadiusAttribute != nil {
		result["group_membership_radius_attribute"] = string(*obj.GroupMembershipRadiusAttribute)
	}

	groupMembershipToReturn := []interface{}{}
	for _, item := range obj.GroupMembershipToReturn {
		groupMembershipToReturn = append(groupMembershipToReturn, AppGroupMembershipToReturnToMap(item))
	}
	result["group_membership_to_return"] = groupMembershipToReturn

	if obj.GroupNameFormat != nil {
		result["group_name_format"] = string(*obj.GroupNameFormat)
	}

	if obj.IncludeGroupInResponse != nil {
		result["include_group_in_response"] = bool(*obj.IncludeGroupInResponse)
	}

	if obj.PasswordAndOtpTogether != nil {
		result["password_and_otp_together"] = bool(*obj.PasswordAndOtpTogether)
	}

	if obj.Port != nil {
		result["port"] = string(*obj.Port)
	}

	if obj.RadiusVendorSpecificId != nil {
		result["radius_vendor_specific_id"] = string(*obj.RadiusVendorSpecificId)
	}

	if obj.ResponseFormat != nil {
		result["response_format"] = string(*obj.ResponseFormat)
	}

	if obj.ResponseFormatDelimiter != nil {
		result["response_format_delimiter"] = string(*obj.ResponseFormatDelimiter)
	}

	if obj.SecretKey != nil {
		result["secret_key"] = string(*obj.SecretKey)
	}

	if obj.TypeOfRadiusApp != nil {
		result["type_of_radius_app"] = string(*obj.TypeOfRadiusApp)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionRequestableApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionRequestableApp, error) {
	result := oci_identity_domains.AppExtensionRequestableApp{}

	if requestable, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "requestable")); ok {
		tmp := requestable.(bool)
		result.Requestable = &tmp
	}

	return result, nil
}

func AppExtensionRequestableAppToMap(obj *oci_identity_domains.AppExtensionRequestableApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Requestable != nil {
		result["requestable"] = bool(*obj.Requestable)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionSamlServiceProviderApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionSamlServiceProviderApp, error) {
	result := oci_identity_domains.AppExtensionSamlServiceProviderApp{}

	if assertionConsumerUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "assertion_consumer_url")); ok {
		tmp := assertionConsumerUrl.(string)
		result.AssertionConsumerUrl = &tmp
	}

	if encryptAssertion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encrypt_assertion")); ok {
		tmp := encryptAssertion.(bool)
		result.EncryptAssertion = &tmp
	}

	if encryptionAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_algorithm")); ok {
		result.EncryptionAlgorithm = oci_identity_domains.AppExtensionSamlServiceProviderAppEncryptionAlgorithmEnum(encryptionAlgorithm.(string))
	}

	if encryptionCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "encryption_certificate")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "encryption_certificate")) {
			tmp := encryptionCertificate.(string)
			result.EncryptionCertificate = &tmp
		}
	}

	if federationProtocol, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "federation_protocol")); ok {
		result.FederationProtocol = oci_identity_domains.AppExtensionSamlServiceProviderAppFederationProtocolEnum(federationProtocol.(string))
	}

	if groupAssertionAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_assertion_attributes")); ok {
		interfaces := groupAssertionAttributes.([]interface{})
		tmp := make([]oci_identity_domains.AppGroupAssertionAttributes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "group_assertion_attributes"), stateDataIndex)
			converted, err := s.mapToAppGroupAssertionAttributes(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "group_assertion_attributes")) {
			result.GroupAssertionAttributes = tmp
		}
	}

	if hokAcsUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hok_acs_url")); ok {
		tmp := hokAcsUrl.(string)
		result.HokAcsUrl = &tmp
	}

	if hokRequired, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "hok_required")); ok {
		tmp := hokRequired.(bool)
		result.HokRequired = &tmp
	}

	if includeSigningCertInSignature, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "include_signing_cert_in_signature")); ok {
		tmp := includeSigningCertInSignature.(bool)
		result.IncludeSigningCertInSignature = &tmp
	}

	if keyEncryptionAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key_encryption_algorithm")); ok {
		result.KeyEncryptionAlgorithm = oci_identity_domains.AppExtensionSamlServiceProviderAppKeyEncryptionAlgorithmEnum(keyEncryptionAlgorithm.(string))
	}

	if logoutBinding, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logout_binding")); ok {
		result.LogoutBinding = oci_identity_domains.AppExtensionSamlServiceProviderAppLogoutBindingEnum(logoutBinding.(string))
	}

	if logoutEnabled, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logout_enabled")); ok {
		tmp := logoutEnabled.(bool)
		result.LogoutEnabled = &tmp
	}

	if logoutRequestUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logout_request_url")); ok {
		tmp := logoutRequestUrl.(string)
		result.LogoutRequestUrl = &tmp
	}

	if logoutResponseUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "logout_response_url")); ok {
		tmp := logoutResponseUrl.(string)
		result.LogoutResponseUrl = &tmp
	}

	if metadata, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "metadata")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "metadata")) {
			tmp := metadata.(string)
			result.Metadata = &tmp
		}
	}

	if nameIdFormat, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name_id_format")); ok {
		tmp := nameIdFormat.(string)
		result.NameIdFormat = &tmp
	}

	if nameIdUserstoreAttribute, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name_id_userstore_attribute")); ok {
		tmp := nameIdUserstoreAttribute.(string)
		result.NameIdUserstoreAttribute = &tmp
	}

	if outboundAssertionAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "outbound_assertion_attributes")); ok {
		if tmpList := outboundAssertionAttributes.([]interface{}); len(tmpList) > 0 {
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "outbound_assertion_attributes"), 0)
			tmp, err := s.mapToAppOutboundAssertionAttributes(fieldKeyFormatNextLevel)
			if err != nil {
				return result, fmt.Errorf("unable to convert outbound_assertion_attributes, encountered error: %v", err)
			}
			result.OutboundAssertionAttributes = &tmp
		}
	}

	if partnerProviderId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "partner_provider_id")); ok {
		tmp := partnerProviderId.(string)
		result.PartnerProviderId = &tmp
	}

	if partnerProviderPattern, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "partner_provider_pattern")); ok {
		tmp := partnerProviderPattern.(string)
		result.PartnerProviderPattern = &tmp
	}

	if signResponseOrAssertion, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "sign_response_or_assertion")); ok {
		result.SignResponseOrAssertion = oci_identity_domains.AppExtensionSamlServiceProviderAppSignResponseOrAssertionEnum(signResponseOrAssertion.(string))
	}

	if signatureHashAlgorithm, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signature_hash_algorithm")); ok {
		result.SignatureHashAlgorithm = oci_identity_domains.AppExtensionSamlServiceProviderAppSignatureHashAlgorithmEnum(signatureHashAlgorithm.(string))
	}

	if signingCertificate, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "signing_certificate")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "signing_certificate")) {
			tmp := signingCertificate.(string)
			result.SigningCertificate = &tmp
		}
	}

	if succinctId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "succinct_id")); ok {
		tmp := succinctId.(string)
		result.SuccinctId = &tmp
	}

	if tenantProviderId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tenant_provider_id")); ok {
		tmp := tenantProviderId.(string)
		result.TenantProviderId = &tmp
	}

	if userAssertionAttributes, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_assertion_attributes")); ok {
		interfaces := userAssertionAttributes.([]interface{})
		tmp := make([]oci_identity_domains.AppUserAssertionAttributes, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "user_assertion_attributes"), stateDataIndex)
			converted, err := s.mapToAppUserAssertionAttributes(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "user_assertion_attributes")) {
			result.UserAssertionAttributes = tmp
		}
	}

	return result, nil
}

func AppExtensionSamlServiceProviderAppToMap(obj *oci_identity_domains.AppExtensionSamlServiceProviderApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AssertionConsumerUrl != nil {
		result["assertion_consumer_url"] = string(*obj.AssertionConsumerUrl)
	}

	if obj.EncryptAssertion != nil {
		result["encrypt_assertion"] = bool(*obj.EncryptAssertion)
	}

	result["encryption_algorithm"] = string(obj.EncryptionAlgorithm)

	if obj.EncryptionCertificate != nil {
		result["encryption_certificate"] = string(*obj.EncryptionCertificate)
	}

	result["federation_protocol"] = string(obj.FederationProtocol)

	groupAssertionAttributes := []interface{}{}
	for _, item := range obj.GroupAssertionAttributes {
		groupAssertionAttributes = append(groupAssertionAttributes, AppGroupAssertionAttributesToMap(item))
	}
	result["group_assertion_attributes"] = groupAssertionAttributes

	if obj.HokAcsUrl != nil {
		result["hok_acs_url"] = string(*obj.HokAcsUrl)
	}

	if obj.HokRequired != nil {
		result["hok_required"] = bool(*obj.HokRequired)
	}

	if obj.IncludeSigningCertInSignature != nil {
		result["include_signing_cert_in_signature"] = bool(*obj.IncludeSigningCertInSignature)
	}

	result["key_encryption_algorithm"] = string(obj.KeyEncryptionAlgorithm)

	if obj.LastNotificationSentTime != nil {
		result["last_notification_sent_time"] = string(*obj.LastNotificationSentTime)
	}

	result["logout_binding"] = string(obj.LogoutBinding)

	if obj.LogoutEnabled != nil {
		result["logout_enabled"] = bool(*obj.LogoutEnabled)
	}

	if obj.LogoutRequestUrl != nil {
		result["logout_request_url"] = string(*obj.LogoutRequestUrl)
	}

	if obj.LogoutResponseUrl != nil {
		result["logout_response_url"] = string(*obj.LogoutResponseUrl)
	}

	if obj.Metadata != nil {
		result["metadata"] = string(*obj.Metadata)
	}

	if obj.NameIdFormat != nil {
		result["name_id_format"] = string(*obj.NameIdFormat)
	}

	if obj.NameIdUserstoreAttribute != nil {
		result["name_id_userstore_attribute"] = string(*obj.NameIdUserstoreAttribute)
	}

	if obj.OutboundAssertionAttributes != nil {
		result["outbound_assertion_attributes"] = []interface{}{AppOutboundAssertionAttributesToMap(obj.OutboundAssertionAttributes)}
	}

	if obj.PartnerProviderId != nil {
		result["partner_provider_id"] = string(*obj.PartnerProviderId)
	}

	if obj.PartnerProviderPattern != nil {
		result["partner_provider_pattern"] = string(*obj.PartnerProviderPattern)
	}

	result["sign_response_or_assertion"] = string(obj.SignResponseOrAssertion)

	result["signature_hash_algorithm"] = string(obj.SignatureHashAlgorithm)

	if obj.SigningCertificate != nil {
		result["signing_certificate"] = string(*obj.SigningCertificate)
	}

	if obj.SuccinctId != nil {
		result["succinct_id"] = string(*obj.SuccinctId)
	}

	if obj.TenantProviderId != nil {
		result["tenant_provider_id"] = string(*obj.TenantProviderId)
	}

	userAssertionAttributes := []interface{}{}
	for _, item := range obj.UserAssertionAttributes {
		userAssertionAttributes = append(userAssertionAttributes, AppUserAssertionAttributesToMap(item))
	}
	result["user_assertion_attributes"] = userAssertionAttributes

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppExtensionWebTierPolicyApp(fieldKeyFormat string) (oci_identity_domains.AppExtensionWebTierPolicyApp, error) {
	result := oci_identity_domains.AppExtensionWebTierPolicyApp{}

	if resourceRef, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_ref")); ok {
		tmp := resourceRef.(bool)
		result.ResourceRef = &tmp
	}

	if webTierPolicyAZControl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "web_tier_policy_az_control")); ok {
		result.WebTierPolicyAZControl = oci_identity_domains.AppExtensionWebTierPolicyAppWebTierPolicyAZControlEnum(webTierPolicyAZControl.(string))
	}

	if webTierPolicyJson, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "web_tier_policy_json")); ok {
		tmp := webTierPolicyJson.(string)
		result.WebTierPolicyJson = &tmp
	}

	return result, nil
}

func AppExtensionWebTierPolicyAppToMap(obj *oci_identity_domains.AppExtensionWebTierPolicyApp) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ResourceRef != nil {
		result["resource_ref"] = bool(*obj.ResourceRef)
	}

	result["web_tier_policy_az_control"] = string(obj.WebTierPolicyAZControl)

	if obj.WebTierPolicyJson != nil {
		result["web_tier_policy_json"] = string(*obj.WebTierPolicyJson)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppFlatFileBundleConfigurationProperties(fieldKeyFormat string) (oci_identity_domains.AppFlatFileBundleConfigurationProperties, error) {
	result := oci_identity_domains.AppFlatFileBundleConfigurationProperties{}

	if confidential, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "confidential")); ok {
		tmp := confidential.(bool)
		result.Confidential = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		tmp := displayName.(string)
		result.DisplayName = &tmp
	}

	if helpMessage, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "help_message")); ok {
		tmp := helpMessage.(string)
		result.HelpMessage = &tmp
	}

	if icfType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "icf_type")); ok {
		result.IcfType = oci_identity_domains.AppFlatFileBundleConfigurationPropertiesIcfTypeEnum(icfType.(string))
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if order, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "order")); ok {
		tmp := order.(int)
		result.Order = &tmp
	}

	if required, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "required")); ok {
		tmp := required.(bool)
		result.Required = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		interfaces := value.([]interface{})
		tmp := make([]string, len(interfaces))
		for i := range interfaces {
			if interfaces[i] != nil {
				tmp[i] = interfaces[i].(string)
			}
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "value")) {
			result.Value = tmp
		}
	}

	return result, nil
}

func AppFlatFileBundleConfigurationPropertiesToMap(obj oci_identity_domains.AppFlatFileBundleConfigurationProperties) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Confidential != nil {
		result["confidential"] = bool(*obj.Confidential)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.HelpMessage != nil {
		result["help_message"] = string(*obj.HelpMessage)
	}

	result["icf_type"] = string(obj.IcfType)

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Order != nil {
		result["order"] = int(*obj.Order)
	}

	if obj.Required != nil {
		result["required"] = bool(*obj.Required)
	}

	result["value"] = obj.Value

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppFlatFileConnectorBundle(fieldKeyFormat string) (oci_identity_domains.AppFlatFileConnectorBundle, error) {
	result := oci_identity_domains.AppFlatFileConnectorBundle{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	if wellKnownId, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "well_known_id")); ok {
		tmp := wellKnownId.(string)
		result.WellKnownId = &tmp
	}

	return result, nil
}

func AppFlatFileConnectorBundleToMap(obj *oci_identity_domains.AppFlatFileConnectorBundle) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	if obj.WellKnownId != nil {
		result["well_known_id"] = string(*obj.WellKnownId)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppFormFillUrlMatch(fieldKeyFormat string) (oci_identity_domains.AppFormFillUrlMatch, error) {
	result := oci_identity_domains.AppFormFillUrlMatch{}

	if formUrl, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_url")); ok {
		tmp := formUrl.(string)
		result.FormUrl = &tmp
	}

	if formUrlMatchType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "form_url_match_type")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "form_url_match_type")) {
			tmp := formUrlMatchType.(string)
			result.FormUrlMatchType = &tmp
		}
	}

	return result, nil
}

func AppFormFillUrlMatchToMap(obj oci_identity_domains.AppFormFillUrlMatch) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.FormUrl != nil {
		result["form_url"] = string(*obj.FormUrl)
	}

	if obj.FormUrlMatchType != nil {
		result["form_url_match_type"] = string(*obj.FormUrlMatchType)
	}

	return result
}

func AppGrantedAppRolesToMap(obj oci_identity_domains.AppGrantedAppRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AdminRole != nil {
		result["admin_role"] = bool(*obj.AdminRole)
	}

	if obj.AppId != nil {
		result["app_id"] = string(*obj.AppId)
	}

	if obj.AppName != nil {
		result["app_name"] = string(*obj.AppName)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.LegacyGroupName != nil {
		result["legacy_group_name"] = string(*obj.LegacyGroupName)
	}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func AppGrantsToMap(obj oci_identity_domains.AppGrants) map[string]interface{} {
	result := map[string]interface{}{}

	result["grant_mechanism"] = string(obj.GrantMechanism)

	if obj.GranteeId != nil {
		result["grantee_id"] = string(*obj.GranteeId)
	}

	result["grantee_type"] = string(obj.GranteeType)

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppGroupAssertionAttributes(fieldKeyFormat string) (oci_identity_domains.AppGroupAssertionAttributes, error) {
	result := oci_identity_domains.AppGroupAssertionAttributes{}

	if condition, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "condition")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "condition")) {
			tmp := condition.(string)
			result.Condition = &tmp
		}
	}

	if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "format")) {
			tmp := format.(string)
			result.Format = &tmp
		}
	}

	if groupName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "group_name")); ok {
		tmp := groupName.(string)
		result.GroupName = &tmp
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	return result, nil
}

func AppGroupAssertionAttributesToMap(obj oci_identity_domains.AppGroupAssertionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Condition != nil {
		result["condition"] = string(*obj.Condition)
	}

	if obj.Format != nil {
		result["format"] = string(*obj.Format)
	}

	if obj.GroupName != nil {
		result["group_name"] = string(*obj.GroupName)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppGroupMembershipToReturn(fieldKeyFormat string) (oci_identity_domains.AppGroupMembershipToReturn, error) {
	result := oci_identity_domains.AppGroupMembershipToReturn{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppGroupMembershipToReturnToMap(obj oci_identity_domains.AppGroupMembershipToReturn) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppIdentityBridges(fieldKeyFormat string) (oci_identity_domains.AppIdentityBridges, error) {
	result := oci_identity_domains.AppIdentityBridges{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppIdentityBridgesToMap(obj oci_identity_domains.AppIdentityBridges) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppIdentityProviders(fieldKeyFormat string) (oci_identity_domains.AppIdentityProviders, error) {
	result := oci_identity_domains.AppIdentityProviders{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppIdentityProvidersToMap(obj oci_identity_domains.AppIdentityProviders) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppIdpPolicy(fieldKeyFormat string) (oci_identity_domains.AppIdpPolicy, error) {
	result := oci_identity_domains.AppIdpPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppIdpPolicyToMap(obj *oci_identity_domains.AppIdpPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppObjectClasses(fieldKeyFormat string) (oci_identity_domains.AppObjectClasses, error) {
	result := oci_identity_domains.AppObjectClasses{}

	if display, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display")); ok {
		tmp := display.(string)
		result.Display = &tmp
	}

	if isAccountObjectClass, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "is_account_object_class")); ok {
		tmp := isAccountObjectClass.(bool)
		result.IsAccountObjectClass = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if resourceType, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "resource_type")); ok {
		tmp := resourceType.(string)
		result.ResourceType = &tmp
	}

	if type_, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "type")); ok {
		result.Type = oci_identity_domains.AppObjectClassesTypeEnum(type_.(string))
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppObjectClassesToMap(obj oci_identity_domains.AppObjectClasses) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.IsAccountObjectClass != nil {
		result["is_account_object_class"] = bool(*obj.IsAccountObjectClass)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.ResourceType != nil {
		result["resource_type"] = string(*obj.ResourceType)
	}

	result["type"] = string(obj.Type)

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppOutboundAssertionAttributes(fieldKeyFormat string) (oci_identity_domains.AppOutboundAssertionAttributes, error) {
	result := oci_identity_domains.AppOutboundAssertionAttributes{}

	if direction, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "direction")); ok {
		tmp := direction.(string)
		result.Direction = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppOutboundAssertionAttributesToMap(obj *oci_identity_domains.AppOutboundAssertionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Direction != nil {
		result["direction"] = string(*obj.Direction)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppProtectableSecondaryAudiences(fieldKeyFormat string) (oci_identity_domains.AppProtectableSecondaryAudiences, error) {
	result := oci_identity_domains.AppProtectableSecondaryAudiences{}

	if readOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_only")); ok {
		tmp := readOnly.(bool)
		result.ReadOnly = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppProtectableSecondaryAudiencesToMap(obj oci_identity_domains.AppProtectableSecondaryAudiences) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppRadiusPolicy(fieldKeyFormat string) (oci_identity_domains.AppRadiusPolicy, error) {
	result := oci_identity_domains.AppRadiusPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppRadiusPolicyToMap(obj *oci_identity_domains.AppRadiusPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppSamlServiceProvider(fieldKeyFormat string) (oci_identity_domains.AppSamlServiceProvider, error) {
	result := oci_identity_domains.AppSamlServiceProvider{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppSamlServiceProviderToMap(obj *oci_identity_domains.AppSamlServiceProvider) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppScopes(fieldKeyFormat string) (oci_identity_domains.AppScopes, error) {
	result := oci_identity_domains.AppScopes{}

	if description, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "description")); ok {
		tmp := description.(string)
		result.Description = &tmp
	}

	if displayName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "display_name")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "display_name")) {
			tmp := displayName.(string)
			result.DisplayName = &tmp
		}
	}

	if fqs, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "fqs")); ok {
		tmp := fqs.(string)
		result.Fqs = &tmp
	}

	if readOnly, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "read_only")); ok {
		tmp := readOnly.(bool)
		result.ReadOnly = &tmp
	}

	if requiresConsent, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "requires_consent")); ok {
		tmp := requiresConsent.(bool)
		result.RequiresConsent = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppScopesToMap(obj oci_identity_domains.AppScopes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.DisplayName != nil {
		result["display_name"] = string(*obj.DisplayName)
	}

	if obj.Fqs != nil {
		result["fqs"] = string(*obj.Fqs)
	}

	if obj.ReadOnly != nil {
		result["read_only"] = bool(*obj.ReadOnly)
	}

	if obj.RequiresConsent != nil {
		result["requires_consent"] = bool(*obj.RequiresConsent)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppServiceParams(fieldKeyFormat string) (oci_identity_domains.AppServiceParams, error) {
	result := oci_identity_domains.AppServiceParams{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppServiceParamsToMap(obj oci_identity_domains.AppServiceParams) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppSignonPolicy(fieldKeyFormat string) (oci_identity_domains.AppSignonPolicy, error) {
	result := oci_identity_domains.AppSignonPolicy{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppSignonPolicyToMap(obj *oci_identity_domains.AppSignonPolicy) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppTermsOfUse(fieldKeyFormat string) (oci_identity_domains.AppTermsOfUse, error) {
	result := oci_identity_domains.AppTermsOfUse{}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppTermsOfUseToMap(obj *oci_identity_domains.AppTermsOfUse) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppThreeLeggedOAuthCredential(fieldKeyFormat string) (oci_identity_domains.AppThreeLeggedOAuthCredential, error) {
	result := oci_identity_domains.AppThreeLeggedOAuthCredential{}

	if accessToken, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_token")); ok {
		tmp := accessToken.(string)
		result.AccessToken = &tmp
	}

	if accessTokenExpiry, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "access_token_expiry")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "access_token_expiry")) {
			tmp := accessTokenExpiry.(string)
			result.AccessTokenExpiry = &tmp
		}
	}

	if refreshToken, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "refresh_token")); ok {
		tmp := refreshToken.(string)
		result.RefreshToken = &tmp
	}

	return result, nil
}

func AppThreeLeggedOAuthCredentialToMap(obj *oci_identity_domains.AppThreeLeggedOAuthCredential) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.AccessToken != nil {
		result["access_token"] = string(*obj.AccessToken)
	}

	if obj.AccessTokenExpiry != nil {
		result["access_token_expiry"] = string(*obj.AccessTokenExpiry)
	}

	if obj.RefreshToken != nil {
		result["refresh_token"] = string(*obj.RefreshToken)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppTrustPolicies(fieldKeyFormat string) (oci_identity_domains.AppTrustPolicies, error) {
	result := oci_identity_domains.AppTrustPolicies{}

	if ref, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "ref")); ok {
		tmp := ref.(string)
		result.Ref = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func AppTrustPoliciesToMap(obj oci_identity_domains.AppTrustPolicies) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToAppUserAssertionAttributes(fieldKeyFormat string) (oci_identity_domains.AppUserAssertionAttributes, error) {
	result := oci_identity_domains.AppUserAssertionAttributes{}

	if format, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "format")); ok {
		if !IsOptionalAndEmpty(IdentityDomainsAppResource(), s.D, fmt.Sprintf(fieldKeyFormat, "format")) {
			tmp := format.(string)
			result.Format = &tmp
		}
	}

	if name, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "name")); ok {
		tmp := name.(string)
		result.Name = &tmp
	}

	if userStoreAttributeName, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "user_store_attribute_name")); ok {
		tmp := userStoreAttributeName.(string)
		result.UserStoreAttributeName = &tmp
	}

	return result, nil
}

func AppUserAssertionAttributesToMap(obj oci_identity_domains.AppUserAssertionAttributes) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Format != nil {
		result["format"] = string(*obj.Format)
	}

	if obj.Name != nil {
		result["name"] = string(*obj.Name)
	}

	if obj.UserStoreAttributeName != nil {
		result["user_store_attribute_name"] = string(*obj.UserStoreAttributeName)
	}

	return result
}

func AppUserRolesToMap(obj oci_identity_domains.AppUserRoles) map[string]interface{} {
	result := map[string]interface{}{}

	if obj.Description != nil {
		result["description"] = string(*obj.Description)
	}

	if obj.Display != nil {
		result["display"] = string(*obj.Display)
	}

	if obj.Ref != nil {
		result["ref"] = string(*obj.Ref)
	}

	if obj.Value != nil {
		result["value"] = string(*obj.Value)
	}

	return result
}

func (s *IdentityDomainsAppResourceCrud) mapToExtensionOCITags(fieldKeyFormat string) (oci_identity_domains.ExtensionOciTags, error) {
	result := oci_identity_domains.ExtensionOciTags{}

	if definedTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "defined_tags")); ok {
		interfaces := definedTags.([]interface{})
		tmp := make([]oci_identity_domains.DefinedTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "defined_tags"), stateDataIndex)
			converted, err := s.mapTodefinedTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "defined_tags")) {
			result.DefinedTags = tmp
		}
	}

	if freeformTags, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "freeform_tags")); ok {
		interfaces := freeformTags.([]interface{})
		tmp := make([]oci_identity_domains.FreeformTags, len(interfaces))
		for i := range interfaces {
			stateDataIndex := i
			fieldKeyFormatNextLevel := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(fieldKeyFormat, "freeform_tags"), stateDataIndex)
			converted, err := s.mapTofreeformTags(fieldKeyFormatNextLevel)
			if err != nil {
				return result, err
			}
			tmp[i] = converted
		}
		if len(tmp) != 0 || s.D.HasChange(fmt.Sprintf(fieldKeyFormat, "freeform_tags")) {
			result.FreeformTags = tmp
		}
	}

	if tagSlug, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "tag_slug")); ok {
		result.TagSlug = &tagSlug
	}

	return result, nil
}

func (s *IdentityDomainsAppResourceCrud) mapTodefinedTags(fieldKeyFormat string) (oci_identity_domains.DefinedTags, error) {
	result := oci_identity_domains.DefinedTags{}

	if namespace, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "namespace")); ok {
		tmp := namespace.(string)
		result.Namespace = &tmp
	}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *IdentityDomainsAppResourceCrud) mapTofreeformTags(fieldKeyFormat string) (oci_identity_domains.FreeformTags, error) {
	result := oci_identity_domains.FreeformTags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}

func (s *IdentityDomainsAppResourceCrud) mapTotags(fieldKeyFormat string) (oci_identity_domains.Tags, error) {
	result := oci_identity_domains.Tags{}

	if key, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "key")); ok {
		tmp := key.(string)
		result.Key = &tmp
	}

	if value, ok := s.D.GetOkExists(fmt.Sprintf(fieldKeyFormat, "value")); ok {
		tmp := value.(string)
		result.Value = &tmp
	}

	return result, nil
}
