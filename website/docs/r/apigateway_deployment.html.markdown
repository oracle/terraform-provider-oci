---
subcategory: "API Gateway"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apigateway_deployment"
sidebar_current: "docs-oci-resource-apigateway-deployment"
description: |-
  Provides the Deployment resource in Oracle Cloud Infrastructure API Gateway service
---

# oci_apigateway_deployment
This resource provides the Deployment resource in Oracle Cloud Infrastructure API Gateway service.

Creates a new deployment.


## Example Usage

```hcl
resource "oci_apigateway_deployment" "test_deployment" {
	#Required
	compartment_id = var.compartment_id
	gateway_id = oci_apigateway_gateway.test_gateway.id
	path_prefix = var.deployment_path_prefix
	specification {

		#Optional
		logging_policies {

			#Optional
			access_log {

				#Optional
				is_enabled = var.deployment_specification_logging_policies_access_log_is_enabled
			}
			execution_log {

				#Optional
				is_enabled = var.deployment_specification_logging_policies_execution_log_is_enabled
				log_level = var.deployment_specification_logging_policies_execution_log_log_level
			}
		}
		request_policies {

			#Optional
			authentication {
				#Required
				type = var.deployment_specification_request_policies_authentication_type

				#Optional
				audiences = var.deployment_specification_request_policies_authentication_audiences
				cache_key = var.deployment_specification_request_policies_authentication_cache_key
				function_id = oci_functions_function.test_function.id
				is_anonymous_access_allowed = var.deployment_specification_request_policies_authentication_is_anonymous_access_allowed
				issuers = var.deployment_specification_request_policies_authentication_issuers
				max_clock_skew_in_seconds = var.deployment_specification_request_policies_authentication_max_clock_skew_in_seconds
				parameters = var.deployment_specification_request_policies_authentication_parameters
				public_keys {
					#Required
					type = var.deployment_specification_request_policies_authentication_public_keys_type

					#Optional
					is_ssl_verify_disabled = var.deployment_specification_request_policies_authentication_public_keys_is_ssl_verify_disabled
					keys {
						#Required
						format = var.deployment_specification_request_policies_authentication_public_keys_keys_format

						#Optional
						alg = var.deployment_specification_request_policies_authentication_public_keys_keys_alg
						e = var.deployment_specification_request_policies_authentication_public_keys_keys_e
						key = var.deployment_specification_request_policies_authentication_public_keys_keys_key
						key_ops = var.deployment_specification_request_policies_authentication_public_keys_keys_key_ops
						kid = var.deployment_specification_request_policies_authentication_public_keys_keys_kid
						kty = var.deployment_specification_request_policies_authentication_public_keys_keys_kty
						n = var.deployment_specification_request_policies_authentication_public_keys_keys_n
						use = var.deployment_specification_request_policies_authentication_public_keys_keys_use
					}
					max_cache_duration_in_hours = var.deployment_specification_request_policies_authentication_public_keys_max_cache_duration_in_hours
					uri = var.deployment_specification_request_policies_authentication_public_keys_uri
				}
				token_auth_scheme = var.deployment_specification_request_policies_authentication_token_auth_scheme
				token_header = var.deployment_specification_request_policies_authentication_token_header
				token_query_param = var.deployment_specification_request_policies_authentication_token_query_param
				validation_failure_policy {
					#Required
					type = var.deployment_specification_request_policies_authentication_validation_failure_policy_type

					#Optional
					client_details {
						#Required
						type = var.deployment_specification_request_policies_authentication_validation_failure_policy_client_details_type

						#Optional
						client_id = oci_apigateway_client.test_client.id
						client_secret_id = oci_vault_secret.test_secret.id
						client_secret_version_number = var.deployment_specification_request_policies_authentication_validation_failure_policy_client_details_client_secret_version_number
					}
					fallback_redirect_path = var.deployment_specification_request_policies_authentication_validation_failure_policy_fallback_redirect_path
					logout_path = var.deployment_specification_request_policies_authentication_validation_failure_policy_logout_path
					max_expiry_duration_in_hours = var.deployment_specification_request_policies_authentication_validation_failure_policy_max_expiry_duration_in_hours
					response_code = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_code
					response_header_transformations {

						#Optional
						filter_headers {

							#Optional
							items {

								#Optional
								name = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_filter_headers_items_name
							}
							type = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_filter_headers_type
						}
						rename_headers {

							#Optional
							items {

								#Optional
								from = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_rename_headers_items_from
								to = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_rename_headers_items_to
							}
						}
						set_headers {

							#Optional
							items {

								#Optional
								if_exists = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_set_headers_items_if_exists
								name = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_set_headers_items_name
								values = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_header_transformations_set_headers_items_values
							}
						}
					}
					response_message = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_message
					response_type = var.deployment_specification_request_policies_authentication_validation_failure_policy_response_type
					scopes = var.deployment_specification_request_policies_authentication_validation_failure_policy_scopes
					source_uri_details {
						#Required
						type = var.deployment_specification_request_policies_authentication_validation_failure_policy_source_uri_details_type

						#Optional
						uri = var.deployment_specification_request_policies_authentication_validation_failure_policy_source_uri_details_uri
					}
					use_cookies_for_intermediate_steps = var.deployment_specification_request_policies_authentication_validation_failure_policy_use_cookies_for_intermediate_steps
					use_cookies_for_session = var.deployment_specification_request_policies_authentication_validation_failure_policy_use_cookies_for_session
					use_pkce = var.deployment_specification_request_policies_authentication_validation_failure_policy_use_pkce
				}
				validation_policy {
					#Required
					type = var.deployment_specification_request_policies_authentication_validation_policy_type

					#Optional
					additional_validation_policy {

						#Optional
						audiences = var.deployment_specification_request_policies_authentication_validation_policy_additional_validation_policy_audiences
						issuers = var.deployment_specification_request_policies_authentication_validation_policy_additional_validation_policy_issuers
						verify_claims {

							#Optional
							is_required = var.deployment_specification_request_policies_authentication_validation_policy_additional_validation_policy_verify_claims_is_required
							key = var.deployment_specification_request_policies_authentication_validation_policy_additional_validation_policy_verify_claims_key
							values = var.deployment_specification_request_policies_authentication_validation_policy_additional_validation_policy_verify_claims_values
						}
					}
					client_details {
						#Required
						type = var.deployment_specification_request_policies_authentication_validation_policy_client_details_type

						#Optional
						client_id = oci_apigateway_client.test_client.id
						client_secret_id = oci_vault_secret.test_secret.id
						client_secret_version_number = var.deployment_specification_request_policies_authentication_validation_policy_client_details_client_secret_version_number
					}
					is_ssl_verify_disabled = var.deployment_specification_request_policies_authentication_validation_policy_is_ssl_verify_disabled
					keys {
						#Required
						format = var.deployment_specification_request_policies_authentication_validation_policy_keys_format

						#Optional
						alg = var.deployment_specification_request_policies_authentication_validation_policy_keys_alg
						e = var.deployment_specification_request_policies_authentication_validation_policy_keys_e
						key = var.deployment_specification_request_policies_authentication_validation_policy_keys_key
						key_ops = var.deployment_specification_request_policies_authentication_validation_policy_keys_key_ops
						kid = var.deployment_specification_request_policies_authentication_validation_policy_keys_kid
						kty = var.deployment_specification_request_policies_authentication_validation_policy_keys_kty
						n = var.deployment_specification_request_policies_authentication_validation_policy_keys_n
						use = var.deployment_specification_request_policies_authentication_validation_policy_keys_use
					}
					max_cache_duration_in_hours = var.deployment_specification_request_policies_authentication_validation_policy_max_cache_duration_in_hours
					source_uri_details {
						#Required
						type = var.deployment_specification_request_policies_authentication_validation_policy_source_uri_details_type

						#Optional
						uri = var.deployment_specification_request_policies_authentication_validation_policy_source_uri_details_uri
					}
					uri = var.deployment_specification_request_policies_authentication_validation_policy_uri
				}
				verify_claims {

					#Optional
					is_required = var.deployment_specification_request_policies_authentication_verify_claims_is_required
					key = var.deployment_specification_request_policies_authentication_verify_claims_key
					values = var.deployment_specification_request_policies_authentication_verify_claims_values
				}
			}
			cors {
				#Required
				allowed_origins = var.deployment_specification_request_policies_cors_allowed_origins

				#Optional
				allowed_headers = var.deployment_specification_request_policies_cors_allowed_headers
				allowed_methods = var.deployment_specification_request_policies_cors_allowed_methods
				exposed_headers = var.deployment_specification_request_policies_cors_exposed_headers
				is_allow_credentials_enabled = var.deployment_specification_request_policies_cors_is_allow_credentials_enabled
				max_age_in_seconds = var.deployment_specification_request_policies_cors_max_age_in_seconds
			}
			dynamic_authentication {
				#Required
				authentication_servers {
					#Required
					authentication_server_detail {
						#Required
						type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_type

						#Optional
						audiences = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_audiences
						function_id = oci_functions_function.test_function.id
						is_anonymous_access_allowed = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_is_anonymous_access_allowed
						issuers = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_issuers
						max_clock_skew_in_seconds = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_max_clock_skew_in_seconds
						public_keys {
							#Required
							type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_type

							#Optional
							is_ssl_verify_disabled = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_is_ssl_verify_disabled
							keys {
								#Required
								format = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_format

								#Optional
								alg = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_alg
								e = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_e
								key = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_key
								key_ops = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_key_ops
								kid = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_kid
								kty = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_kty
								n = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_n
								use = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_keys_use
							}
							max_cache_duration_in_hours = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_max_cache_duration_in_hours
							uri = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_public_keys_uri
						}
						token_auth_scheme = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_token_auth_scheme
						token_header = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_token_header
						token_query_param = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_token_query_param
						validation_failure_policy {
							#Required
							type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_type

							#Optional
							client_details {
								#Required
								type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_client_details_type

								#Optional
								client_id = oci_apigateway_client.test_client.id
								client_secret_id = oci_vault_secret.test_secret.id
								client_secret_version_number = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_client_details_client_secret_version_number
							}
							fallback_redirect_path = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_fallback_redirect_path
							logout_path = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_logout_path
							max_expiry_duration_in_hours = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_max_expiry_duration_in_hours
							response_code = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_code
							response_header_transformations {

								#Optional
								filter_headers {

									#Optional
									items {

										#Optional
										name = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_filter_headers_items_name
									}
									type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_filter_headers_type
								}
								rename_headers {

									#Optional
									items {

										#Optional
										from = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_rename_headers_items_from
										to = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_rename_headers_items_to
									}
								}
								set_headers {

									#Optional
									items {

										#Optional
										if_exists = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_set_headers_items_if_exists
										name = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_set_headers_items_name
										values = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_header_transformations_set_headers_items_values
									}
								}
							}
							response_message = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_message
							response_type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_response_type
							scopes = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_scopes
							source_uri_details {
								#Required
								type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_source_uri_details_type

								#Optional
								uri = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_source_uri_details_uri
							}
							use_cookies_for_intermediate_steps = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_use_cookies_for_intermediate_steps
							use_cookies_for_session = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_use_cookies_for_session
							use_pkce = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_failure_policy_use_pkce
						}
						validation_policy {
							#Required
							type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_type

							#Optional
							additional_validation_policy {

								#Optional
								audiences = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_additional_validation_policy_audiences
								issuers = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_additional_validation_policy_issuers
								verify_claims {

									#Optional
									is_required = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_additional_validation_policy_verify_claims_is_required
									key = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_additional_validation_policy_verify_claims_key
									values = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_additional_validation_policy_verify_claims_values
								}
							}
							client_details {
								#Required
								type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_client_details_type

								#Optional
								client_id = oci_apigateway_client.test_client.id
								client_secret_id = oci_vault_secret.test_secret.id
								client_secret_version_number = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_client_details_client_secret_version_number
							}
							is_ssl_verify_disabled = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_is_ssl_verify_disabled
							keys {
								#Required
								format = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_format

								#Optional
								alg = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_alg
								e = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_e
								key = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_key
								key_ops = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_key_ops
								kid = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_kid
								kty = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_kty
								n = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_n
								use = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_keys_use
							}
							max_cache_duration_in_hours = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_max_cache_duration_in_hours
							source_uri_details {
								#Required
								type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_source_uri_details_type

								#Optional
								uri = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_source_uri_details_uri
							}
							uri = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_validation_policy_uri
						}
						verify_claims {

							#Optional
							is_required = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_verify_claims_is_required
							key = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_verify_claims_key
							values = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_authentication_server_detail_verify_claims_values
						}
					}
					key {
						#Required
						name = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_name

						#Optional
						expression = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_expression
						is_default = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_is_default
						type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_type
						values = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_values
					}
				}
				selection_source {
					#Required
					selector = var.deployment_specification_request_policies_dynamic_authentication_selection_source_selector
					type = var.deployment_specification_request_policies_dynamic_authentication_selection_source_type
				}
			}
			mutual_tls {

				#Optional
				allowed_sans = var.deployment_specification_request_policies_mutual_tls_allowed_sans
				is_verified_certificate_required = var.deployment_specification_request_policies_mutual_tls_is_verified_certificate_required
			}
			rate_limiting {
				#Required
				rate_in_requests_per_second = var.deployment_specification_request_policies_rate_limiting_rate_in_requests_per_second
				rate_key = var.deployment_specification_request_policies_rate_limiting_rate_key
			}
			usage_plans {
				#Required
				token_locations = var.deployment_specification_request_policies_usage_plans_token_locations
			}
		}
		routes {
			#Required
			backend {
				#Required
				type = var.deployment_specification_routes_backend_type

				#Optional
				allowed_post_logout_uris = var.deployment_specification_routes_backend_allowed_post_logout_uris
				body = var.deployment_specification_routes_backend_body
				connect_timeout_in_seconds = var.deployment_specification_routes_backend_connect_timeout_in_seconds
				function_id = oci_functions_function.test_function.id
				headers {

					#Optional
					name = var.deployment_specification_routes_backend_headers_name
					value = var.deployment_specification_routes_backend_headers_value
				}
				is_ssl_verify_disabled = var.deployment_specification_routes_backend_is_ssl_verify_disabled
				post_logout_state = var.deployment_specification_routes_backend_post_logout_state
				read_timeout_in_seconds = var.deployment_specification_routes_backend_read_timeout_in_seconds
				routing_backends {
					backend {
						#Required
						type = var.deployment_specification_routes_backend_routing_backends_backend_type

						#Optional
						body = var.deployment_specification_routes_backend_routing_backends_backend_body
						connect_timeout_in_seconds = var.deployment_specification_routes_backend_routing_backends_backend_connect_timeout_in_seconds
						function_id = oci_functions_function.test_function.id
						headers {

							#Optional
							name = var.deployment_specification_routes_backend_routing_backends_backend_headers_name
							value = var.deployment_specification_routes_backend_routing_backends_backend_headers_value
						}
						is_ssl_verify_disabled = var.deployment_specification_routes_backend_routing_backends_backend_is_ssl_verify_disabled
						read_timeout_in_seconds = var.deployment_specification_routes_backend_routing_backends_backend_read_timeout_in_seconds
						send_timeout_in_seconds = var.deployment_specification_routes_backend_routing_backends_backend_send_timeout_in_seconds
						status = var.deployment_specification_routes_backend_routing_backends_backend_status
						url = var.deployment_specification_routes_backend_routing_backends_backend_url
					}
					key {
						#Required
						name = var.deployment_specification_routes_backend_routing_backends_key_name
						type = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_type
		
						#Optional
						expression = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_expression
						is_default = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_is_default
						values = var.deployment_specification_request_policies_dynamic_authentication_authentication_servers_key_values
					}
				}
				selection_source {
					#Required
					selector = var.deployment_specification_routes_backend_selection_source_selector
					type = var.deployment_specification_routes_backend_selection_source_type
				}
				send_timeout_in_seconds = var.deployment_specification_routes_backend_send_timeout_in_seconds
				status = var.deployment_specification_routes_backend_status
				url = var.deployment_specification_routes_backend_url
			}
			path = var.deployment_specification_routes_path

			#Optional
			logging_policies {

				#Optional
				access_log {

					#Optional
					is_enabled = var.deployment_specification_routes_logging_policies_access_log_is_enabled
				}
				execution_log {

					#Optional
					is_enabled = var.deployment_specification_routes_logging_policies_execution_log_is_enabled
					log_level = var.deployment_specification_routes_logging_policies_execution_log_log_level
				}
			}
			methods = var.deployment_specification_routes_methods
			request_policies {

				#Optional
				authorization {

					#Optional
					allowed_scope = var.deployment_specification_routes_request_policies_authorization_allowed_scope
					type = var.deployment_specification_routes_request_policies_authorization_type
				}
				body_validation {

					#Optional
					content {
						#Required
						media_type = var.deployment_specification_routes_request_policies_body_validation_content_media_type
						validation_type = var.deployment_specification_routes_request_policies_body_validation_content_validation_type
					}

					required = var.deployment_specification_routes_request_policies_body_validation_required
					validation_mode = var.deployment_specification_routes_request_policies_body_validation_validation_mode
				}
				cors {
					#Required
					allowed_origins = var.deployment_specification_routes_request_policies_cors_allowed_origins

					#Optional
					allowed_headers = var.deployment_specification_routes_request_policies_cors_allowed_headers
					allowed_methods = var.deployment_specification_routes_request_policies_cors_allowed_methods
					exposed_headers = var.deployment_specification_routes_request_policies_cors_exposed_headers
					is_allow_credentials_enabled = var.deployment_specification_routes_request_policies_cors_is_allow_credentials_enabled
					max_age_in_seconds = var.deployment_specification_routes_request_policies_cors_max_age_in_seconds
				}
				header_transformations {

					#Optional
					filter_headers {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_request_policies_header_transformations_filter_headers_items_name
						}
						type = var.deployment_specification_routes_request_policies_header_transformations_filter_headers_type
					}
					rename_headers {
						#Required
						items {
							#Required
							from = var.deployment_specification_routes_request_policies_header_transformations_rename_headers_items_from
							to = var.deployment_specification_routes_request_policies_header_transformations_rename_headers_items_to
						}
					}
					set_headers {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_request_policies_header_transformations_set_headers_items_name
							values = var.deployment_specification_routes_request_policies_header_transformations_set_headers_items_values

							#Optional
							if_exists = var.deployment_specification_routes_request_policies_header_transformations_set_headers_items_if_exists
						}
					}
				}
				header_validations {

					#Optional
					headers {
						#Required
						name = var.deployment_specification_routes_request_policies_header_validations_headers_name

						#Optional
						required = var.deployment_specification_routes_request_policies_header_validations_headers_required
					}
					validation_mode = var.deployment_specification_routes_request_policies_header_validations_validation_mode
				}
				query_parameter_transformations {

					#Optional
					filter_query_parameters {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_items_name
						}
						type = var.deployment_specification_routes_request_policies_query_parameter_transformations_filter_query_parameters_type
					}
					rename_query_parameters {
						#Required
						items {
							#Required
							from = var.deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_from
							to = var.deployment_specification_routes_request_policies_query_parameter_transformations_rename_query_parameters_items_to
						}
					}
					set_query_parameters {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_name
							values = var.deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_values

							#Optional
							if_exists = var.deployment_specification_routes_request_policies_query_parameter_transformations_set_query_parameters_items_if_exists
						}
					}
				}
				query_parameter_validations {

					#Optional
					parameters {
						#Required
						name = var.deployment_specification_routes_request_policies_query_parameter_validations_parameters_name

						#Optional
						required = var.deployment_specification_routes_request_policies_query_parameter_validations_parameters_required
					}
					validation_mode = var.deployment_specification_routes_request_policies_query_parameter_validations_validation_mode
				}
				response_cache_lookup {
					#Required
					type = var.deployment_specification_routes_request_policies_response_cache_lookup_type

					#Optional
					cache_key_additions = var.deployment_specification_routes_request_policies_response_cache_lookup_cache_key_additions
					is_enabled = var.deployment_specification_routes_request_policies_response_cache_lookup_is_enabled
					is_private_caching_enabled = var.deployment_specification_routes_request_policies_response_cache_lookup_is_private_caching_enabled
				}
			}
			response_policies {

				#Optional
				header_transformations {

					#Optional
					filter_headers {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_response_policies_header_transformations_filter_headers_items_name
						}
						type = var.deployment_specification_routes_response_policies_header_transformations_filter_headers_type
					}
					rename_headers {
						#Required
						items {
							#Required
							from = var.deployment_specification_routes_response_policies_header_transformations_rename_headers_items_from
							to = var.deployment_specification_routes_response_policies_header_transformations_rename_headers_items_to
						}
					}
					set_headers {
						#Required
						items {
							#Required
							name = var.deployment_specification_routes_response_policies_header_transformations_set_headers_items_name
							values = var.deployment_specification_routes_response_policies_header_transformations_set_headers_items_values

							#Optional
							if_exists = var.deployment_specification_routes_response_policies_header_transformations_set_headers_items_if_exists
						}
					}
				}
				response_cache_store {
					#Required
					time_to_live_in_seconds = var.deployment_specification_routes_response_policies_response_cache_store_time_to_live_in_seconds
					type = var.deployment_specification_routes_response_policies_response_cache_store_type
				}
			}
		}
	}

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.deployment_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `path_prefix` - (Required) A path on which to deploy all routes contained in the API deployment specification. For more information, see [Deploying an API on an API Gateway by Creating an API Deployment](https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm). 
* `specification` - (Required) (Updatable) The logical configuration of the API exposed by a deployment.
	* `logging_policies` - (Optional) (Updatable) Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
		* `access_log` - (Optional) (Updatable) Configures the logging policies for the access logs of an API Deployment. 
			* `is_enabled` - (Optional) (Updatable) Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
		* `execution_log` - (Optional) (Updatable) Configures the logging policies for the execution logs of an API Deployment. 
			* `is_enabled` - (Optional) (Updatable) Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `log_level` - (Optional) (Updatable) Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
	* `request_policies` - (Optional) (Updatable) Global behavior applied to all requests received by the API.
		* `authentication` - (Optional) (Updatable) Information on how to authenticate incoming requests.
			* `audiences` - (Required when type=JWT_AUTHENTICATION) (Updatable) The list of intended recipients for the token.
			* `cache_key` - (Applicable when type=CUSTOM_AUTHENTICATION) (Updatable) A list of keys from "parameters" attribute value whose values will be added to the cache key. 
			* `function_id` - (Required when type=CUSTOM_AUTHENTICATION) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `is_anonymous_access_allowed` - (Optional) (Updatable) Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
			* `issuers` - (Required when type=JWT_AUTHENTICATION) (Updatable) A list of parties that could have issued the token.
			* `max_clock_skew_in_seconds` - (Applicable when type=JWT_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
			* `parameters` - (Applicable when type=CUSTOM_AUTHENTICATION) (Updatable) A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression. Example: `{"foo": "request.header[abc]"}` 
			* `public_keys` - (Required when type=JWT_AUTHENTICATION) (Updatable) A set of Public Keys that will be used to verify the JWT signature.
				* `is_ssl_verify_disabled` - (Applicable when type=REMOTE_JWKS) (Updatable) Defines whether or not to uphold SSL verification. 
				* `keys` - (Applicable when type=STATIC_KEYS) (Updatable) The set of static public keys.
					* `alg` - (Required when format=JSON_WEB_KEY) (Updatable) The algorithm intended for use with this key.
					* `e` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - (Required) (Updatable) The format of the public key.
					* `key` - (Required when format=PEM) (Updatable) The content of the PEM-encoded public key.
					* `key_ops` - (Applicable when format=JSON_WEB_KEY) (Updatable) The operations for which this key is to be used.
					* `kid` - (Required when type=STATIC_KEYS) (Updatable) A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - (Required when format=JSON_WEB_KEY) (Updatable) The key type.
					* `n` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - (Applicable when format=JSON_WEB_KEY) (Updatable) The intended use of the public key.
				* `max_cache_duration_in_hours` - (Applicable when type=REMOTE_JWKS) (Updatable) The duration for which the JWKS should be cached before it is fetched again. 
				* `type` - (Required) (Updatable) Type of the public key set.
				* `uri` - (Required when type=REMOTE_JWKS) (Updatable) The uri from which to retrieve the key. It must be accessible without authentication. 
			* `token_auth_scheme` - (Applicable when type=JWT_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
			* `token_header` - (Optional) (Updatable) The name of the header containing the authentication token.
			* `token_query_param` - (Optional) (Updatable) The name of the query parameter containing the authentication token.
			* `type` - (Required) (Updatable) Type of the authentication policy to use.
			* `validation_failure_policy` - (Applicable when type=CUSTOM_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) Policy for defining behaviour on validation failure.
				* `client_details` - (Required when type=OAUTH2) (Updatable) Client App Credential details.
					* `client_id` - (Required when type=CUSTOM) (Updatable) Client ID for the OAuth2/OIDC app.
					* `client_secret_id` - (Required when type=CUSTOM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
					* `client_secret_version_number` - (Required when type=CUSTOM) (Updatable) The version number of the client secret to use.
					* `type` - (Required) (Updatable) To specify where the Client App details should be taken from.
				* `fallback_redirect_path` - (Applicable when type=OAUTH2) (Updatable) The path to be used as fallback after OAuth2.
				* `logout_path` - (Applicable when type=OAUTH2) (Updatable) The path to be used as logout.
				* `max_expiry_duration_in_hours` - (Applicable when type=OAUTH2) (Updatable) The duration for which the OAuth2 success token should be cached before it is fetched again. 
				* `response_code` - (Applicable when type=MODIFY_RESPONSE) (Updatable) HTTP response code, can include context variables.
				* `response_header_transformations` - (Applicable when type=MODIFY_RESPONSE) (Updatable) A set of transformations to apply to HTTP headers that pass through the gateway. 
					* `filter_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
						* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers. 
							* `name` - (Required when type=MODIFY_RESPONSE) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `type` - (Required when type=MODIFY_RESPONSE) (Updatable) BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
					* `rename_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Rename HTTP headers as they pass through the gateway. 
						* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers.
							* `from` - (Required when type=MODIFY_RESPONSE) (Updatable) The original case-insensitive name of the header.  This name must be unique across transformation policies. 
							* `to` - (Required when type=MODIFY_RESPONSE) (Updatable) The new name of the header.  This name must be unique across transformation policies. 
					* `set_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Set HTTP headers as they pass through the gateway. 
						* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers.
							* `if_exists` - (Applicable when type=MODIFY_RESPONSE) (Updatable) If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
							* `name` - (Required when type=MODIFY_RESPONSE) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
							* `values` - (Required when type=MODIFY_RESPONSE) (Updatable) A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
				* `response_message` - (Applicable when type=MODIFY_RESPONSE) (Updatable) HTTP response message.
				* `response_type` - (Required when type=OAUTH2) (Updatable) Response Type.
				* `scopes` - (Required when type=OAUTH2) (Updatable) List of scopes.
				* `source_uri_details` - (Required when type=OAUTH2) (Updatable) Auth endpoint details.
					* `type` - (Required) (Updatable) Type of the Uri detail.
					* `uri` - (Required when type=DISCOVERY_URI) (Updatable) The discovery URI for the auth server.
				* `type` - (Required) (Updatable) Type of the Validation failure Policy.
				* `use_cookies_for_intermediate_steps` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to use cookies for OAuth2 intermediate steps. 
				* `use_cookies_for_session` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to use cookies for session maintenance. 
				* `use_pkce` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to support PKCE. 
			* `validation_policy` - (Required when type=TOKEN_AUTHENTICATION) (Updatable) Authentication Policies for the Token Authentication types.
				* `additional_validation_policy` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) Additional JWT validation checks.
					* `audiences` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) The list of intended recipients for the token.
					* `issuers` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) A list of parties that could have issued the token.
					* `verify_claims` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) A list of claims which should be validated to consider the token valid.
						* `is_required` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
						* `key` - (Required when type=TOKEN_AUTHENTICATION) (Updatable) Name of the claim.
						* `values` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
				* `client_details` - (Required when type=REMOTE_DISCOVERY) (Updatable) Client App Credential details.
					* `client_id` - (Required when type=CUSTOM) (Updatable) Client ID for the OAuth2/OIDC app.
					* `client_secret_id` - (Required when type=CUSTOM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
					* `client_secret_version_number` - (Required when type=CUSTOM) (Updatable) The version number of the client secret to use.
					* `type` - (Required) (Updatable) To specify where the Client App details should be taken from.
				* `is_ssl_verify_disabled` - (Applicable when type=REMOTE_DISCOVERY | REMOTE_JWKS) (Updatable) Defines whether or not to uphold SSL verification. 
				* `keys` - (Applicable when type=STATIC_KEYS) (Updatable) The set of static public keys.
					* `alg` - (Required when format=JSON_WEB_KEY) (Updatable) The algorithm intended for use with this key.
					* `e` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - (Required) (Updatable) The format of the public key.
					* `key` - (Required when format=PEM) (Updatable) The content of the PEM-encoded public key.
					* `key_ops` - (Applicable when format=JSON_WEB_KEY) (Updatable) The operations for which this key is to be used.
					* `kid` - (Required when type=STATIC_KEYS) (Updatable) A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - (Required when format=JSON_WEB_KEY) (Updatable) The key type.
					* `n` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - (Applicable when format=JSON_WEB_KEY) (Updatable) The intended use of the public key.
				* `max_cache_duration_in_hours` - (Applicable when type=REMOTE_DISCOVERY | REMOTE_JWKS) (Updatable) The duration for which the introspect URL response should be cached before it is fetched again. 
				* `source_uri_details` - (Required when type=REMOTE_DISCOVERY) (Updatable) Auth endpoint details.
					* `type` - (Required) (Updatable) Type of the Uri detail.
					* `uri` - (Required when type=DISCOVERY_URI) (Updatable) The discovery URI for the auth server.
				* `type` - (Required) (Updatable) Type of the token validation policy.
				* `uri` - (Required when type=REMOTE_JWKS) (Updatable) The uri from which to retrieve the key. It must be accessible without authentication. 
			* `verify_claims` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) A list of claims which should be validated to consider the token valid.
				* `is_required` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
				* `key` - (Required when type=JWT_AUTHENTICATION) (Updatable) Name of the claim.
				* `values` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
		* `cors` - (Optional) (Updatable) Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
			* `allowed_headers` - (Optional) (Updatable) The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
			* `allowed_methods` - (Optional) (Updatable) The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
			* `allowed_origins` - (Required) (Updatable) The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
			* `exposed_headers` - (Optional) (Updatable) The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
			* `is_allow_credentials_enabled` - (Optional) (Updatable) Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
			* `max_age_in_seconds` - (Optional) (Updatable) The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
		* `dynamic_authentication` - (Optional) (Updatable) Policy on how to authenticate requests when multiple authentication options are configured for a deployment. For an incoming request, the value of selector specified under selectionSource will be matched against the keys specified for each authentication server. The authentication server whose key matches the value of selector will be used for authentication.
			* `authentication_servers` - (Required) (Updatable) List of authentication servers to choose from during dynamic authentication.
				* `authentication_server_detail` - (Required) (Updatable) Information on how to authenticate incoming requests.
					* `audiences` - (Required when type=JWT_AUTHENTICATION) (Updatable) The list of intended recipients for the token.
					* `function_id` - (Required when type=CUSTOM_AUTHENTICATION) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
					* `is_anonymous_access_allowed` - (Optional) (Updatable) Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
					* `issuers` - (Required when type=JWT_AUTHENTICATION) (Updatable) A list of parties that could have issued the token.
					* `max_clock_skew_in_seconds` - (Applicable when type=JWT_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
					* `parameters` - (Applicable when type=CUSTOM_AUTHENTICATION) (Updatable) A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression. Example: `{"foo": "request.header[abc]"}` 
					* `public_keys` - (Required when type=JWT_AUTHENTICATION) (Updatable) A set of Public Keys that will be used to verify the JWT signature.
						* `is_ssl_verify_disabled` - (Applicable when type=REMOTE_JWKS) (Updatable) Defines whether or not to uphold SSL verification. 
						* `keys` - (Applicable when type=STATIC_KEYS) (Updatable) The set of static public keys.
							* `alg` - (Required when format=JSON_WEB_KEY) (Updatable) The algorithm intended for use with this key.
							* `e` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded exponent of the RSA public key represented by this key. 
							* `format` - (Required) (Updatable) The format of the public key.
							* `key` - (Required when format=PEM) (Updatable) The content of the PEM-encoded public key.
							* `key_ops` - (Applicable when format=JSON_WEB_KEY) (Updatable) The operations for which this key is to be used.
							* `kid` - (Required when type=STATIC_KEYS) (Updatable) A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
							* `kty` - (Required when format=JSON_WEB_KEY) (Updatable) The key type.
							* `n` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded modulus of the RSA public key represented by this key. 
							* `use` - (Applicable when format=JSON_WEB_KEY) (Updatable) The intended use of the public key.
						* `max_cache_duration_in_hours` - (Applicable when type=REMOTE_JWKS) (Updatable) The duration for which the JWKS should be cached before it is fetched again. 
						* `type` - (Required) (Updatable) Type of the public key set.
						* `uri` - (Required when type=REMOTE_JWKS) (Updatable) The uri from which to retrieve the key. It must be accessible without authentication. 
					* `token_auth_scheme` - (Applicable when type=JWT_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
					* `token_header` - (Optional) (Updatable) The name of the header containing the authentication token.
					* `token_query_param` - (Optional) (Updatable) The name of the query parameter containing the authentication token.
					* `type` - (Required) (Updatable) Type of the authentication policy to use.
					* `validation_failure_policy` - (Applicable when type=CUSTOM_AUTHENTICATION | TOKEN_AUTHENTICATION) (Updatable) Policy for defining behaviour on validation failure.
						* `client_details` - (Required when type=OAUTH2) (Updatable) Client App Credential details.
							* `client_id` - (Required when type=CUSTOM) (Updatable) Client ID for the OAuth2/OIDC app.
							* `client_secret_id` - (Required when type=CUSTOM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
							* `client_secret_version_number` - (Required when type=CUSTOM) (Updatable) The version number of the client secret to use.
							* `type` - (Required) (Updatable) To specify where the Client App details should be taken from.
						* `fallback_redirect_path` - (Applicable when type=OAUTH2) (Updatable) The path to be used as fallback after OAuth2.
						* `logout_path` - (Applicable when type=OAUTH2) (Updatable) The path to be used as logout.
						* `max_expiry_duration_in_hours` - (Applicable when type=OAUTH2) (Updatable) The duration for which the OAuth2 success token should be cached before it is fetched again. 
						* `response_code` - (Applicable when type=MODIFY_RESPONSE) (Updatable) HTTP response code, can include context variables.
						* `response_header_transformations` - (Applicable when type=MODIFY_RESPONSE) (Updatable) A set of transformations to apply to HTTP headers that pass through the gateway. 
							* `filter_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
								* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers. 
									* `name` - (Required when type=MODIFY_RESPONSE) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
								* `type` - (Required when type=MODIFY_RESPONSE) (Updatable) BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
							* `rename_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Rename HTTP headers as they pass through the gateway. 
								* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers.
									* `from` - (Required when type=MODIFY_RESPONSE) (Updatable) The original case-insensitive name of the header.  This name must be unique across transformation policies. 
									* `to` - (Required when type=MODIFY_RESPONSE) (Updatable) The new name of the header.  This name must be unique across transformation policies. 
							* `set_headers` - (Applicable when type=MODIFY_RESPONSE) (Updatable) Set HTTP headers as they pass through the gateway. 
								* `items` - (Required when type=MODIFY_RESPONSE) (Updatable) The list of headers.
									* `if_exists` - (Applicable when type=MODIFY_RESPONSE) (Updatable) If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
									* `name` - (Required when type=MODIFY_RESPONSE) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
									* `values` - (Required when type=MODIFY_RESPONSE) (Updatable) A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
						* `response_message` - (Applicable when type=MODIFY_RESPONSE) (Updatable) HTTP response message.
						* `response_type` - (Required when type=OAUTH2) (Updatable) Response Type.
						* `scopes` - (Required when type=OAUTH2) (Updatable) List of scopes.
						* `source_uri_details` - (Required when type=OAUTH2) (Updatable) Auth endpoint details.
							* `type` - (Required) (Updatable) Type of the Uri detail.
							* `uri` - (Required when type=DISCOVERY_URI) (Updatable) The discovery URI for the auth server.
						* `type` - (Required) (Updatable) Type of the Validation failure Policy.
						* `use_cookies_for_intermediate_steps` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to use cookies for OAuth2 intermediate steps. 
						* `use_cookies_for_session` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to use cookies for session maintenance. 
						* `use_pkce` - (Applicable when type=OAUTH2) (Updatable) Defines whether or not to support PKCE. 
					* `validation_policy` - (Required when type=TOKEN_AUTHENTICATION) (Updatable) Authentication Policies for the Token Authentication types.
						* `additional_validation_policy` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) Additional JWT validation checks.
							* `audiences` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) The list of intended recipients for the token.
							* `issuers` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) A list of parties that could have issued the token.
							* `verify_claims` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) A list of claims which should be validated to consider the token valid.
								* `is_required` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
								* `key` - (Required when type=TOKEN_AUTHENTICATION) (Updatable) Name of the claim.
								* `values` - (Applicable when type=TOKEN_AUTHENTICATION) (Updatable) The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
						* `client_details` - (Required when type=REMOTE_DISCOVERY) (Updatable) Client App Credential details.
							* `client_id` - (Required when type=CUSTOM) (Updatable) Client ID for the OAuth2/OIDC app.
							* `client_secret_id` - (Required when type=CUSTOM) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
							* `client_secret_version_number` - (Required when type=CUSTOM) (Updatable) The version number of the client secret to use.
							* `type` - (Required) (Updatable) To specify where the Client App details should be taken from.
						* `is_ssl_verify_disabled` - (Applicable when type=REMOTE_DISCOVERY | REMOTE_JWKS) (Updatable) Defines whether or not to uphold SSL verification. 
						* `keys` - (Applicable when type=STATIC_KEYS) (Updatable) The set of static public keys.
							* `alg` - (Required when format=JSON_WEB_KEY) (Updatable) The algorithm intended for use with this key.
							* `e` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded exponent of the RSA public key represented by this key. 
							* `format` - (Required) (Updatable) The format of the public key.
							* `key` - (Required when format=PEM) (Updatable) The content of the PEM-encoded public key.
							* `key_ops` - (Applicable when format=JSON_WEB_KEY) (Updatable) The operations for which this key is to be used.
							* `kid` - (Required when type=STATIC_KEYS) (Updatable) A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
							* `kty` - (Required when format=JSON_WEB_KEY) (Updatable) The key type.
							* `n` - (Required when format=JSON_WEB_KEY) (Updatable) The base64 url encoded modulus of the RSA public key represented by this key. 
							* `use` - (Applicable when format=JSON_WEB_KEY) (Updatable) The intended use of the public key.
						* `max_cache_duration_in_hours` - (Applicable when type=REMOTE_DISCOVERY | REMOTE_JWKS) (Updatable) The duration for which the introspect URL response should be cached before it is fetched again. 
						* `source_uri_details` - (Required when type=REMOTE_DISCOVERY) (Updatable) Auth endpoint details.
							* `type` - (Required) (Updatable) Type of the Uri detail.
							* `uri` - (Required when type=DISCOVERY_URI) (Updatable) The discovery URI for the auth server.
						* `type` - (Required) (Updatable) Type of the token validation policy.
						* `uri` - (Required when type=REMOTE_JWKS) (Updatable) The uri from which to retrieve the key. It must be accessible without authentication. 
					* `verify_claims` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) A list of claims which should be validated to consider the token valid.
						* `is_required` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
						* `key` - (Required when type=JWT_AUTHENTICATION) (Updatable) Name of the claim.
						* `values` - (Applicable when type=JWT_AUTHENTICATION) (Updatable) The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
				* `key` - (Required) (Updatable) Base policy for defining how to match the context variable in an incoming request with selection keys when dynamically routing and dynamically authenticating requests.
					* `expression` - (Required when type=WILDCARD) (Updatable) A selection key string containing a wildcard to match with the context variable in an incoming request. If the context variable matches the string, the request is sent to the route or authentication server associated with the selection key. Valid wildcards are '*' (zero or more characters) and '+' (one or more characters). The string can only contain one wildcard, and the wildcard must be at the start or the end of the string.
					* `is_default` - (Optional) (Updatable) Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
					* `name` - (Required) (Updatable) Name assigned to the branch.
					* `type` - (Optional) (Updatable) Type of the selection key.
					* `values` - (Applicable when type=ANY_OF) (Updatable) The set of selection keys to match with the context variable in an incoming request. If the context variable exactly matches one of the keys in the set, the request is sent to the route or authentication server associated with the set.
			* `selection_source` - (Required) (Updatable) The type of selector to use when dynamically routing and dynamically authenticating requests.
				* `selector` - (Required) (Updatable) String describing the context variable used as selector.
				* `type` - (Required) (Updatable) Type of the Selection source to use.
		* `mutual_tls` - (Optional) (Updatable) Properties used to configure client mTLS verification when API Consumer makes connection to the gateway. 
			* `allowed_sans` - (Optional) (Updatable) Allowed list of CN or SAN which will be used for verification of certificate.
			* `is_verified_certificate_required` - (Optional) (Updatable) Determines whether to enable client verification when API Consumer makes connection to the gateway.
		* `rate_limiting` - (Optional) (Updatable) Limit the number of requests that should be handled for the specified window using a specfic key.
			* `rate_in_requests_per_second` - (Required) (Updatable) The maximum number of requests per second to allow.
			* `rate_key` - (Required) (Updatable) The key used to group requests together.
		* `usage_plans` - (Optional) (Updatable) Usage plan policies for this deployment
			* `token_locations` - (Required) (Updatable) A list of context variables specifying where API tokens may be located in a request. Example locations:
				* "request.headers[token]"
				* "request.query[token]"
				* "request.auth[Token]"
				* "request.path[TOKEN]" 
	* `routes` - (Optional) (Updatable) A list of routes that this API exposes.
		* `backend` - (Required) (Updatable) The backend to forward requests to. 
			* `allowed_post_logout_uris` - (Applicable when type=OAUTH2_LOGOUT_BACKEND) (Updatable) 
			* `body` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The body of the stock response from the mock backend.
			* `connect_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for establishing a connection with a proxied server. 
			* `function_id` - (Required when type=ORACLE_FUNCTIONS_BACKEND) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `headers` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The headers of the stock response from the mock backend.
				* `name` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Name of the header.
				* `value` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Value of the header.
			* `is_ssl_verify_disabled` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines whether or not to uphold SSL verification. 
			* `post_logout_state` - (Applicable when type=OAUTH2_LOGOUT_BACKEND) (Updatable) Defines a state that should be shared on redirecting to postLogout URL. 
			* `read_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for reading a response from the proxied server. 
			* `routing_backends` - (Required when type=DYNAMIC_ROUTING_BACKEND) (Updatable) List of backends to chose from for Dynamic Routing.
			    * `backend` - (Required) (Updatable) The backend to forward requests to.
			  		* `body` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The body of the stock response from the mock backend.
					* `connect_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for establishing a connection with a proxied server.
					* `function_id` - (Required when type=ORACLE_FUNCTIONS_BACKEND) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
					* `headers` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) The headers of the stock response from the mock backend.
						* `name` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Name of the header.
						* `value` - (Applicable when type=STOCK_RESPONSE_BACKEND) (Updatable) Value of the header.
					* `is_ssl_verify_disabled` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines whether or not to uphold SSL verification.
					* `read_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for reading a response from the proxied server.
					* `send_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for transmitting a request to the proxied server.
					* `status` - (Required when type=STOCK_RESPONSE_BACKEND) (Updatable) The status code of the stock response from the mock backend.
					* `type` - (Required) (Updatable) Type of the API backend.
					* `url` - (Required when type=HTTP_BACKEND) (Updatable) The url of the proxied server.
				* `key` - (Required) (Updatable) Information around the values for selector of an authentication/ routing branch.
					* `expression` - (Required when type=WILDCARD) (Updatable) String describing the expression with wildcards.
					* `is_default` - (Optional) (Updatable) Information regarding whether this is the default branch.
					* `name` - (Required) (Updatable) Name assigned to the branch.
					* `type` - (Required) (Updatable) Information regarding type of the selection key.
					* `values` - (Applicable when type=ANY_OF) (Updatable) Information regarding the set of values of selector for which this branch should be selected.
			* `selection_source` - (Required when type=DYNAMIC_ROUTING_BACKEND) (Updatable) Information around selector used for branching among routes/ authentication servers while dynamic routing/ authentication.
				* `selector` - (Required) (Updatable) String describing the context variable used as selector.
				* `type` - (Required) (Updatable) Type of the Selection source to use.
			* `send_timeout_in_seconds` - (Applicable when type=HTTP_BACKEND) (Updatable) Defines a timeout for transmitting a request to the proxied server. 
			* `status` - (Required when type=STOCK_RESPONSE_BACKEND) (Updatable) The status code of the stock response from the mock backend.
			* `type` - (Required) (Updatable) Type of the API backend.
			* `url` - (Required when type=HTTP_BACKEND) (Updatable) 
		* `logging_policies` - (Optional) (Updatable) Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
			* `access_log` - (Optional) (Updatable) Configures the logging policies for the access logs of an API Deployment. 
				* `is_enabled` - (Optional) (Updatable) Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `execution_log` - (Optional) (Updatable) Configures the logging policies for the execution logs of an API Deployment. 
				* `is_enabled` - (Optional) (Updatable) Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
				* `log_level` - (Optional) (Updatable) Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
		* `methods` - (Optional) (Updatable) A list of allowed methods on this route. 
		* `path` - (Required) (Updatable) A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
		* `request_policies` - (Optional) (Updatable) Behavior applied to any requests received by the API on this route. 
			* `authorization` - (Optional) (Updatable) If authentication has been performed, validate whether the request scope (if any) applies to this route. If no RouteAuthorizationPolicy is defined for a route, a policy with a type of AUTHENTICATION_ONLY is applied. 
				* `allowed_scope` - (Required when type=ANY_OF) (Updatable) A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
				* `type` - (Optional) (Updatable) Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specified. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
			* `body_validation` - (Optional) (Updatable) Validate the payload body of the incoming API requests on a specific route.
				* `content` - (Optional) (Updatable) The content of the request body.
					* `media_type` - (Required) (Updatable) The media_type is a [media type range](https://tools.ietf.org/html/rfc7231#appendix-D) subset restricted to the following schema

						media_type ::= ( / (  "*" "/" "*" ) / ( type "/" "*" ) / ( type "/" subtype ) )

						For requests that match multiple media types, only the most specific media type is applicable. e.g. `text/plain` overrides `text/*` 
					* `validation_type` - (Required) (Updatable) Validation type defines the content validation method.

						Make the validation to first parse the body as the respective format. 
				* `required` - (Optional) (Updatable) Determines if the request body is required in the request.
				* `validation_mode` - (Optional) (Updatable) Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `cors` - (Optional) (Updatable) Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
				* `allowed_headers` - (Optional) (Updatable) The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
				* `allowed_methods` - (Optional) (Updatable) The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
				* `allowed_origins` - (Required) (Updatable) The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
				* `exposed_headers` - (Optional) (Updatable) The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
				* `is_allow_credentials_enabled` - (Optional) (Updatable) Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
				* `max_age_in_seconds` - (Optional) (Updatable) The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
			* `header_transformations` - (Optional) (Updatable) A set of transformations to apply to HTTP headers that pass through the gateway. 
				* `filter_headers` - (Optional) (Updatable) Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - (Required) (Updatable) The list of headers. 
						* `name` - (Required) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `type` - (Required) (Updatable) BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
				* `rename_headers` - (Optional) (Updatable) Rename HTTP headers as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of headers.
						* `from` - (Required) (Updatable) The original case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `to` - (Required) (Updatable) The new name of the header.  This name must be unique across transformation policies. 
				* `set_headers` - (Optional) (Updatable) Set HTTP headers as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of headers.
						* `if_exists` - (Optional) (Updatable) If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - (Required) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `values` - (Required) (Updatable) A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `header_validations` - (Optional) (Updatable) Validate the HTTP headers on the incoming API requests on a specific route.
				* `headers` - (Optional) (Updatable) 
					* `name` - (Required) (Updatable) Parameter name.
					* `required` - (Optional) (Updatable) Determines if the header is required in the request.
				* `validation_mode` - (Optional) (Updatable) Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `query_parameter_transformations` - (Optional) (Updatable) A set of transformations to apply to query parameters that pass through the gateway. 
				* `filter_query_parameters` - (Optional) (Updatable) Filter parameters from the query string as they pass through the gateway.  The gateway applies filters after other transformations, so any parameters set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - (Required) (Updatable) The list of query parameters. 
						* `name` - (Required) (Updatable) The case-sensitive name of the query parameter. 
					* `type` - (Required) (Updatable) BLOCK drops any query parameters that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the parameters in the list and removes all others, so it acts as an inclusion list. 
				* `rename_query_parameters` - (Optional) (Updatable) Rename parameters on the query string as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of query parameters. 
						* `from` - (Required) (Updatable) The original case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
						* `to` - (Required) (Updatable) The new name of the query parameter.  This name must be unique across transformation policies. 
				* `set_query_parameters` - (Optional) (Updatable) Set parameters on the query string as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of query parameters. 
						* `if_exists` - (Optional) (Updatable) If a query parameter with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - (Required) (Updatable) The case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
						* `values` - (Required) (Updatable) A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `query_parameter_validations` - (Optional) (Updatable) Validate the URL query parameters on the incoming API requests on a specific route.
				* `parameters` - (Optional) (Updatable) 
					* `name` - (Required) (Updatable) Parameter name.
					* `required` - (Optional) (Updatable) Determines if the parameter is required in the request.
				* `validation_mode` - (Optional) (Updatable) Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `response_cache_lookup` - (Optional) (Updatable) Base policy for Response Cache lookup. 
				* `cache_key_additions` - (Optional) (Updatable) A list of context expressions whose values will be added to the base cache key. Values should contain an expression enclosed within ${} delimiters. Only the request context is available. 
				* `is_enabled` - (Optional) (Updatable) Whether this policy is currently enabled. 
				* `is_private_caching_enabled` - (Optional) (Updatable) Set true to allow caching responses where the request has an Authorization header. Ensure you have configured your  cache key additions to get the level of isolation across authenticated requests that you require.

					When false, any request with an Authorization header will not be stored in the Response Cache.

					If using the CustomAuthenticationPolicy then the tokenHeader/tokenQueryParam are also subject to this check. 
				* `type` - (Required) (Updatable) Type of the Response Cache Store Policy.
		* `response_policies` - (Optional) (Updatable) Behavior applied to any responses sent by the API for requests on this route. 
			* `header_transformations` - (Optional) (Updatable) A set of transformations to apply to HTTP headers that pass through the gateway. 
				* `filter_headers` - (Optional) (Updatable) Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - (Required) (Updatable) The list of headers. 
						* `name` - (Required) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `type` - (Required) (Updatable) BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
				* `rename_headers` - (Optional) (Updatable) Rename HTTP headers as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of headers.
						* `from` - (Required) (Updatable) The original case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `to` - (Required) (Updatable) The new name of the header.  This name must be unique across transformation policies. 
				* `set_headers` - (Optional) (Updatable) Set HTTP headers as they pass through the gateway. 
					* `items` - (Required) (Updatable) The list of headers.
						* `if_exists` - (Optional) (Updatable) If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - (Required) (Updatable) The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `values` - (Required) (Updatable) A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `response_cache_store` - (Optional) (Updatable) Base policy for how a response from a backend is cached in the Response Cache. 
				* `time_to_live_in_seconds` - (Required) (Updatable) Sets the number of seconds for a response from a backend being stored in the Response Cache before it expires. 
				* `type` - (Required) (Updatable) Type of the Response Cache Store Policy.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the resource is created. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.  Example: `My new resource` 
* `endpoint` - The endpoint to access this deployment on the gateway.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gateway_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource. 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed state. 
* `path_prefix` - A path on which to deploy all routes contained in the API deployment specification. For more information, see [Deploying an API on an API Gateway by Creating an API Deployment](https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm). 
* `specification` - The logical configuration of the API exposed by a deployment.
	* `logging_policies` - Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
		* `access_log` - Configures the logging policies for the access logs of an API Deployment. 
			* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
		* `execution_log` - Configures the logging policies for the execution logs of an API Deployment. 
			* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

				Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
	* `request_policies` - Global behavior applied to all requests received by the API.
		* `authentication` - Information on how to authenticate incoming requests.
			* `audiences` - The list of intended recipients for the token.
			* `cache_key` - A list of keys from "parameters" attribute value whose values will be added to the cache key. 
			* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `is_anonymous_access_allowed` - Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
			* `issuers` - A list of parties that could have issued the token.
			* `max_clock_skew_in_seconds` - The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
			* `parameters` - A map where key is a user defined string and value is a context expressions whose values will be sent to the custom auth function. Values should contain an expression. Example: `{"foo": "request.header[abc]"}` 
			* `public_keys` - A set of Public Keys that will be used to verify the JWT signature.
				* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
				* `keys` - The set of static public keys.
					* `alg` - The algorithm intended for use with this key.
					* `e` - The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - The format of the public key.
					* `key` - The content of the PEM-encoded public key.
					* `key_ops` - The operations for which this key is to be used.
					* `kid` - A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - The key type.
					* `n` - The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - The intended use of the public key.
				* `max_cache_duration_in_hours` - The duration for which the JWKS should be cached before it is fetched again. 
				* `type` - Type of the public key set.
				* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
			* `token_auth_scheme` - The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
			* `token_header` - The name of the header containing the authentication token.
			* `token_query_param` - The name of the query parameter containing the authentication token.
			* `type` - Type of the authentication policy to use.
			* `validation_failure_policy` - Policy for defining behaviour on validation failure.
				* `client_details` - Client App Credential details.
					* `client_id` - Client ID for the OAuth2/OIDC app.
					* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
					* `client_secret_version_number` - The version number of the client secret to use.
					* `type` - To specify where the Client App details should be taken from.
				* `fallback_redirect_path` - The path to be used as fallback after OAuth2.
				* `logout_path` - The path to be used as logout.
				* `max_expiry_duration_in_hours` - The duration for which the OAuth2 success token should be cached before it is fetched again. 
				* `response_code` - HTTP response code, can include context variables.
				* `response_header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
					* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
						* `items` - The list of headers. 
							* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
					* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
						* `items` - The list of headers.
							* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
							* `to` - The new name of the header.  This name must be unique across transformation policies. 
					* `set_headers` - Set HTTP headers as they pass through the gateway. 
						* `items` - The list of headers.
							* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
							* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
							* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
				* `response_message` - HTTP response message.
				* `response_type` - Response Type.
				* `scopes` - List of scopes.
				* `source_uri_details` - Auth endpoint details.
					* `type` - Type of the Uri detail.
					* `uri` - The discovery URI for the auth server.
				* `type` - Type of the Validation failure Policy.
				* `use_cookies_for_intermediate_steps` - Defines whether or not to use cookies for OAuth2 intermediate steps. 
				* `use_cookies_for_session` - Defines whether or not to use cookies for session maintenance. 
				* `use_pkce` - Defines whether or not to support PKCE. 
			* `validation_policy` - Authentication Policies for the Token Authentication types.
				* `additional_validation_policy` - Additional JWT validation checks.
					* `audiences` - The list of intended recipients for the token.
					* `issuers` - A list of parties that could have issued the token.
					* `verify_claims` - A list of claims which should be validated to consider the token valid.
						* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
						* `key` - Name of the claim.
						* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
				* `client_details` - Client App Credential details.
					* `client_id` - Client ID for the OAuth2/OIDC app.
					* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
					* `client_secret_version_number` - The version number of the client secret to use.
					* `type` - To specify where the Client App details should be taken from.
				* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
				* `keys` - The set of static public keys.
					* `alg` - The algorithm intended for use with this key.
					* `e` - The base64 url encoded exponent of the RSA public key represented by this key. 
					* `format` - The format of the public key.
					* `key` - The content of the PEM-encoded public key.
					* `key_ops` - The operations for which this key is to be used.
					* `kid` - A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
					* `kty` - The key type.
					* `n` - The base64 url encoded modulus of the RSA public key represented by this key. 
					* `use` - The intended use of the public key.
				* `max_cache_duration_in_hours` - The duration for which the introspect URL response should be cached before it is fetched again. 
				* `source_uri_details` - Auth endpoint details.
					* `type` - Type of the Uri detail.
					* `uri` - The discovery URI for the auth server.
				* `type` - Type of the token validation policy.
				* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
			* `verify_claims` - A list of claims which should be validated to consider the token valid.
				* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
				* `key` - Name of the claim.
				* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
		* `cors` - Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
			* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
			* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
			* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
			* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
			* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
			* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
		* `dynamic_authentication` - Policy on how to authenticate requests when multiple authentication options are configured for a deployment. For an incoming request, the value of selector specified under selectionSource will be matched against the keys specified for each authentication server. The authentication server whose key matches the value of selector will be used for authentication.
			* `authentication_servers` - List of authentication servers to choose from during dynamic authentication.
				* `authentication_server_detail` - Information on how to authenticate incoming requests.
					* `audiences` - The list of intended recipients for the token.
					* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
					* `is_anonymous_access_allowed` - Whether an unauthenticated user may access the API. Must be "true" to enable ANONYMOUS route authorization. 
					* `issuers` - A list of parties that could have issued the token.
					* `max_clock_skew_in_seconds` - The maximum expected time difference between the system clocks of the token issuer and the API Gateway. 
					* `public_keys` - A set of Public Keys that will be used to verify the JWT signature.
						* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
						* `keys` - The set of static public keys.
							* `alg` - The algorithm intended for use with this key.
							* `e` - The base64 url encoded exponent of the RSA public key represented by this key. 
							* `format` - The format of the public key.
							* `key` - The content of the PEM-encoded public key.
							* `key_ops` - The operations for which this key is to be used.
							* `kid` - A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
							* `kty` - The key type.
							* `n` - The base64 url encoded modulus of the RSA public key represented by this key. 
							* `use` - The intended use of the public key.
						* `max_cache_duration_in_hours` - The duration for which the JWKS should be cached before it is fetched again. 
						* `type` - Type of the public key set.
						* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
					* `token_auth_scheme` - The authentication scheme that is to be used when authenticating the token. This must to be provided if "tokenHeader" is specified. 
					* `token_header` - The name of the header containing the authentication token.
					* `token_query_param` - The name of the query parameter containing the authentication token.
					* `type` - Type of the authentication policy to use.
					* `validation_failure_policy` - Policy for defining behaviour on validation failure.
						* `client_details` - Client App Credential details.
							* `client_id` - Client ID for the OAuth2/OIDC app.
							* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
							* `client_secret_version_number` - The version number of the client secret to use.
							* `type` - To specify where the Client App details should be taken from.
						* `fallback_redirect_path` - The path to be used as fallback after OAuth2.
						* `logout_path` - The path to be used as logout.
						* `max_expiry_duration_in_hours` - The duration for which the OAuth2 success token should be cached before it is fetched again. 
						* `response_code` - HTTP response code, can include context variables.
						* `response_header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
							* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
								* `items` - The list of headers. 
									* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
								* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
							* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
								* `items` - The list of headers.
									* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
									* `to` - The new name of the header.  This name must be unique across transformation policies. 
							* `set_headers` - Set HTTP headers as they pass through the gateway. 
								* `items` - The list of headers.
									* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
									* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
									* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
						* `response_message` - HTTP response message.
						* `response_type` - Response Type.
						* `scopes` - List of scopes.
						* `source_uri_details` - Auth endpoint details.
							* `type` - Type of the Uri detail.
							* `uri` - The discovery URI for the auth server.
						* `type` - Type of the Validation failure Policy.
						* `use_cookies_for_intermediate_steps` - Defines whether or not to use cookies for OAuth2 intermediate steps. 
						* `use_cookies_for_session` - Defines whether or not to use cookies for session maintenance. 
						* `use_pkce` - Defines whether or not to support PKCE. 
					* `validation_policy` - Authentication Policies for the Token Authentication types.
						* `additional_validation_policy` - Additional JWT validation checks.
							* `audiences` - The list of intended recipients for the token.
							* `issuers` - A list of parties that could have issued the token.
							* `verify_claims` - A list of claims which should be validated to consider the token valid.
								* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
								* `key` - Name of the claim.
								* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
						* `client_details` - Client App Credential details.
							* `client_id` - Client ID for the OAuth2/OIDC app.
							* `client_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Vault Service secret resource. 
							* `client_secret_version_number` - The version number of the client secret to use.
							* `type` - To specify where the Client App details should be taken from.
						* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
						* `keys` - The set of static public keys.
							* `alg` - The algorithm intended for use with this key.
							* `e` - The base64 url encoded exponent of the RSA public key represented by this key. 
							* `format` - The format of the public key.
							* `key` - The content of the PEM-encoded public key.
							* `key_ops` - The operations for which this key is to be used.
							* `kid` - A unique key ID. This key will be used to verify the signature of a JWT with matching "kid". 
							* `kty` - The key type.
							* `n` - The base64 url encoded modulus of the RSA public key represented by this key. 
							* `use` - The intended use of the public key.
						* `max_cache_duration_in_hours` - The duration for which the introspect URL response should be cached before it is fetched again. 
						* `source_uri_details` - Auth endpoint details.
							* `type` - Type of the Uri detail.
							* `uri` - The discovery URI for the auth server.
						* `type` - Type of the token validation policy.
						* `uri` - The uri from which to retrieve the key. It must be accessible without authentication. 
					* `verify_claims` - A list of claims which should be validated to consider the token valid.
						* `is_required` - Whether the claim is required to be present in the JWT or not. If set to "false", the claim values will be matched only if the claim is present in the JWT. 
						* `key` - Name of the claim.
						* `values` - The list of acceptable values for a given claim. If this value is "null" or empty and "isRequired" set to "true", then the presence of this claim in the JWT is validated. 
				* `key` - Base policy for defining how to match the context variable in an incoming request with selection keys when dynamically routing and dynamically authenticating requests.
					* `expression` - A selection key string containing a wildcard to match with the context variable in an incoming request. If the context variable matches the string, the request is sent to the route or authentication server associated with the selection key. Valid wildcards are '*' (zero or more characters) and '+' (one or more characters). The string can only contain one wildcard, and the wildcard must be at the start or the end of the string.
					* `is_default` - Specifies whether to use the route or authentication server associated with this selection key as the default. The default is used if the value of a context variable in an incoming request does not match any of the other selection key values when dynamically routing and dynamically authenticating requests.
					* `name` - Name assigned to the branch.
					* `type` - Type of the selection key.
					* `values` - The set of selection keys to match with the context variable in an incoming request. If the context variable exactly matches one of the keys in the set, the request is sent to the route or authentication server associated with the set.
			* `selection_source` - The type of selector to use when dynamically routing and dynamically authenticating requests.
				* `selector` - String describing the context variable used as selector.
				* `type` - Type of the Selection source to use.
		* `mutual_tls` - Properties used to configure client mTLS verification when API Consumer makes connection to the gateway. 
			* `allowed_sans` - Allowed list of CN or SAN which will be used for verification of certificate.
			* `is_verified_certificate_required` - Determines whether to enable client verification when API Consumer makes connection to the gateway.
		* `rate_limiting` - Limit the number of requests that should be handled for the specified window using a specfic key.
			* `rate_in_requests_per_second` - The maximum number of requests per second to allow.
			* `rate_key` - The key used to group requests together.
		* `usage_plans` - Usage plan policies for this deployment
			* `token_locations` - A list of context variables specifying where API tokens may be located in a request. Example locations:
				* "request.headers[token]"
				* "request.query[token]"
				* "request.auth[Token]"
				* "request.path[TOKEN]" 
	* `routes` - A list of routes that this API exposes.
		* `backend` - The backend to forward requests to. 
			* `allowed_post_logout_uris` - 
			* `body` - The body of the stock response from the mock backend.
			* `connect_timeout_in_seconds` - Defines a timeout for establishing a connection with a proxied server. 
			* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource. 
			* `headers` - The headers of the stock response from the mock backend.
				* `name` - Name of the header.
				* `value` - Value of the header.
			* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification. 
			* `post_logout_state` - Defines a state that should be shared on redirecting to postLogout URL. 
			* `read_timeout_in_seconds` - Defines a timeout for reading a response from the proxied server. 
			* `routing_backends` - List of backends to chose from for Dynamic Routing.
				* `backend` - The backend to forward requests to.
					* `body` - The body of the stock response from the mock backend.
					* `connect_timeout_in_seconds` - Defines a timeout for establishing a connection with a proxied server.
					* `function_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Functions function resource.
					* `headers` - The headers of the stock response from the mock backend.
						* `name` - Name of the header.
						* `value` - Value of the header.
					* `is_ssl_verify_disabled` - Defines whether or not to uphold SSL verification.
					* `read_timeout_in_seconds` - Defines a timeout for reading a response from the proxied server.
					* `send_timeout_in_seconds` - Defines a timeout for transmitting a request to the proxied server.
					* `status` - The status code of the stock response from the mock backend.
					* `type` - Type of the API backend.
					* `url` - The url of the proxied server.
				* `key` - Information around the values for selector of an authentication/ routing branch.
					* `expression` - String describing the expression with wildcards.
					* `is_default` - Information regarding whether this is the default branch.
					* `name` - Name assigned to the branch.
					* `type` - Information regarding type of the selection key.
					* `values` - Information regarding the set of values of selector for which this branch should be selected.
			* `selection_source` - Information around selector used for branching among routes/ authentication servers while dynamic routing/ authentication.
				* `selector` - String describing the context variable used as selector.
				* `type` - Type of the Selection source to use.
			* `send_timeout_in_seconds` - Defines a timeout for transmitting a request to the proxied server. 
			* `status` - The status code of the stock response from the mock backend.
			* `type` - Type of the API backend.
			* `url` - 
		* `logging_policies` - Policies controlling the pushing of logs to Oracle Cloud Infrastructure Public Logging. 
			* `access_log` - Configures the logging policies for the access logs of an API Deployment. 
				* `is_enabled` - Enables pushing of access logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query access logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'access' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
			* `execution_log` - Configures the logging policies for the execution logs of an API Deployment. 
				* `is_enabled` - Enables pushing of execution logs to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Oracle recommends using the Oracle Cloud Infrastructure Logging service to enable, retrieve, and query execution logs for an API Deployment. If there is an active log object for the API Deployment and its category is set to 'execution' in Oracle Cloud Infrastructure Logging service, the logs will not be uploaded to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket.

					Please note that the functionality to push to the legacy Oracle Cloud Infrastructure Object Storage log archival bucket has been deprecated and will be removed in the future. 
				* `log_level` - Specifies the log level used to control logging output of execution logs. Enabling logging at a given level also enables logging at all higher levels. 
		* `methods` - A list of allowed methods on this route. 
		* `path` - A URL path pattern that must be matched on this route. The path pattern may contain a subset of RFC 6570 identifiers to allow wildcard and parameterized matching. 
		* `request_policies` - Behavior applied to any requests received by the API on this route. 
			* `authorization` - If authentication has been performed, validate whether the request scope (if any) applies to this route. If no RouteAuthorizationPolicy is defined for a route, a policy with a type of AUTHENTICATION_ONLY is applied. 
				* `allowed_scope` - A user whose scope includes any of these access ranges is allowed on this route. Access ranges are case-sensitive. 
				* `type` - Indicates how authorization should be applied. For a type of ANY_OF, an "allowedScope" property must also be specified. Otherwise, only a type is required. For a type of ANONYMOUS, an authenticated API must have the "isAnonymousAccessAllowed" property set to "true" in the authentication policy. 
			* `body_validation` - Validate the payload body of the incoming API requests on a specific route.
				* `content` - The content of the request body.
					* `media_type` - (Required) (Updatable) The media_type is a [media type range](https://tools.ietf.org/html/rfc7231#appendix-D) subset restricted to the following schema

						media_type ::= ( / (  "*" "/" "*" ) / ( type "/" "*" ) / ( type "/" subtype ) )

						For requests that match multiple media types, only the most specific media type is applicable. e.g. `text/plain` overrides `text/*` 
					* `validation_type` - Validation type defines the content validation method.

						Make the validation to first parse the body as the respective format. 
				* `required` - Determines if the request body is required in the request.
				* `validation_mode` - Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `cors` - Enable CORS (Cross-Origin-Resource-Sharing) request handling. 
				* `allowed_headers` - The list of headers that will be allowed from the client via the Access-Control-Allow-Headers header. '*' will allow all headers. 
				* `allowed_methods` - The list of allowed HTTP methods that will be returned for the preflight OPTIONS request in the Access-Control-Allow-Methods header. '*' will allow all methods. 
				* `allowed_origins` - The list of allowed origins that the CORS handler will use to respond to CORS requests. The gateway will send the Access-Control-Allow-Origin header with the best origin match for the circumstances. '*' will match any origins, and 'null' will match queries from 'file:' origins. All other origins must be qualified with the scheme, full hostname, and port if necessary. 
				* `exposed_headers` - The list of headers that the client will be allowed to see from the response as indicated by the Access-Control-Expose-Headers header. '*' will expose all headers. 
				* `is_allow_credentials_enabled` - Whether to send the Access-Control-Allow-Credentials header to allow CORS requests with cookies. 
				* `max_age_in_seconds` - The time in seconds for the client to cache preflight responses. This is sent as the Access-Control-Max-Age if greater than 0. 
			* `header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
				* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - The list of headers. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
				* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `to` - The new name of the header.  This name must be unique across transformation policies. 
				* `set_headers` - Set HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `header_validations` - Validate the HTTP headers on the incoming API requests on a specific route.
				* `headers` - 
					* `name` - Parameter name.
					* `required` - Determines if the header is required in the request.
				* `validation_mode` - Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `query_parameter_transformations` - A set of transformations to apply to query parameters that pass through the gateway. 
				* `filter_query_parameters` - Filter parameters from the query string as they pass through the gateway.  The gateway applies filters after other transformations, so any parameters set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - The list of query parameters. 
						* `name` - The case-sensitive name of the query parameter. 
					* `type` - BLOCK drops any query parameters that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the parameters in the list and removes all others, so it acts as an inclusion list. 
				* `rename_query_parameters` - Rename parameters on the query string as they pass through the gateway. 
					* `items` - The list of query parameters. 
						* `from` - The original case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
						* `to` - The new name of the query parameter.  This name must be unique across transformation policies. 
				* `set_query_parameters` - Set parameters on the query string as they pass through the gateway. 
					* `items` - The list of query parameters. 
						* `if_exists` - If a query parameter with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - The case-sensitive name of the query parameter.  This name must be unique across transformation policies. 
						* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `query_parameter_validations` - Validate the URL query parameters on the incoming API requests on a specific route.
				* `parameters` - 
					* `name` - Parameter name.
					* `required` - Determines if the parameter is required in the request.
				* `validation_mode` - Validation behavior mode.

					In `ENFORCING` mode, upon a validation failure, the request will be rejected with a 4xx response and not sent to the backend.

					In `PERMISSIVE` mode, the result of the validation will be exposed as metrics while the request will follow the normal path.

					`DISABLED` type turns the validation off. 
			* `response_cache_lookup` - Base policy for Response Cache lookup. 
				* `cache_key_additions` - A list of context expressions whose values will be added to the base cache key. Values should contain an expression enclosed within ${} delimiters. Only the request context is available. 
				* `is_enabled` - Whether this policy is currently enabled. 
				* `is_private_caching_enabled` - Set true to allow caching responses where the request has an Authorization header. Ensure you have configured your  cache key additions to get the level of isolation across authenticated requests that you require.

					When false, any request with an Authorization header will not be stored in the Response Cache.

					If using the CustomAuthenticationPolicy then the tokenHeader/tokenQueryParam are also subject to this check. 
				* `type` - Type of the Response Cache Store Policy.
		* `response_policies` - Behavior applied to any responses sent by the API for requests on this route. 
			* `header_transformations` - A set of transformations to apply to HTTP headers that pass through the gateway. 
				* `filter_headers` - Filter HTTP headers as they pass through the gateway.  The gateway applies filters after other transformations, so any headers set or renamed must also be listed here when using an ALLOW type policy. 
					* `items` - The list of headers. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
					* `type` - BLOCK drops any headers that are in the list of items, so it acts as an exclusion list.  ALLOW permits only the headers in the list and removes all others, so it acts as an inclusion list. 
				* `rename_headers` - Rename HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `from` - The original case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `to` - The new name of the header.  This name must be unique across transformation policies. 
				* `set_headers` - Set HTTP headers as they pass through the gateway. 
					* `items` - The list of headers.
						* `if_exists` - If a header with the same name already exists in the request, OVERWRITE will overwrite the value, APPEND will append to the existing value, or SKIP will keep the existing value. 
						* `name` - The case-insensitive name of the header.  This name must be unique across transformation policies. 
						* `values` - A list of new values.  Each value can be a constant or may include one or more expressions enclosed within ${} delimiters. 
			* `response_cache_store` - Base policy for how a response from a backend is cached in the Response Cache. 
				* `time_to_live_in_seconds` - Sets the number of seconds for a response from a backend being stored in the Response Cache before it expires. 
				* `type` - Type of the Response Cache Store Policy.
* `state` - The current state of the deployment.
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Deployment
	* `update` - (Defaults to 20 minutes), when updating the Deployment
	* `delete` - (Defaults to 20 minutes), when destroying the Deployment


## Import

Deployments can be imported using the `id`, e.g.

```
$ terraform import oci_apigateway_deployment.test_deployment "id"
```

