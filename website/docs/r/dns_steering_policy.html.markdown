---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policy"
sidebar_current: "docs-oci-resource-dns-steering_policy"
description: |-
  Provides the Steering Policy resource in Oracle Cloud Infrastructure Dns service
---

# oci_dns_steering_policy
This resource provides the Steering Policy resource in Oracle Cloud Infrastructure Dns service.

Creates a new steering policy in the specified compartment.


## Example Usage

```hcl
resource "oci_dns_steering_policy" "test_steering_policy" {
	#Required
	compartment_id = "${var.compartment_id}"
	display_name = "${var.steering_policy_display_name}"
	template = "${var.steering_policy_template}"

	#Optional
	answers {
		#Required
		name = "${var.steering_policy_answers_name}"
		rdata = "${var.steering_policy_answers_rdata}"
		rtype = "${var.steering_policy_answers_rtype}"

		#Optional
		is_disabled = "${var.steering_policy_answers_is_disabled}"
		pool = "${var.steering_policy_answers_pool}"
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	health_check_monitor_id = "${oci_dns_health_check_monitor.test_health_check_monitor.id}"
	rules {
		#Required
		rule_type = "${var.steering_policy_rules_rule_type}"

		#Optional
		cases {

			#Optional
			answer_data {

				#Optional
				answer_condition = "${var.steering_policy_rules_cases_answer_data_answer_condition}"
				should_keep = "${var.steering_policy_rules_cases_answer_data_should_keep}"
				value = "${var.steering_policy_rules_cases_answer_data_value}"
			}
			case_condition = "${var.steering_policy_rules_cases_case_condition}"
			count = "${var.steering_policy_rules_cases_count}"
		}
		default_answer_data {

			#Optional
			answer_condition = "${var.steering_policy_rules_default_answer_data_answer_condition}"
			should_keep = "${var.steering_policy_rules_default_answer_data_should_keep}"
			value = "${var.steering_policy_rules_default_answer_data_value}"
		}
		default_count = "${var.steering_policy_rules_default_count}"
		description = "${var.steering_policy_rules_description}"
	}
	ttl = "${var.steering_policy_ttl}"
}
```

## Argument Reference

If a change to the Steering Policy will result in the destruction and recreation of the resource, the Steering Policy will be temporarily removed from all attached domains while it is being updated. Since this could cause a temporary outage we recommend that you create DNS records at the affected domains with default values. Those records will be used to resolve DNS queries for the affected domains while the Steering Policy is offline.

The following arguments are supported:

* `answers` - (Optional) The set of all answers that can potentially issue from the steering policy. 
	* `is_disabled` - (Optional) Whether or not an answer should be excluded from responses, e.g. because the corresponding server is down for maintenance. Note, however, that such filtering is not automatic and will only take place if a rule implements it. 
	* `name` - (Required) A user-friendly name for the answer, unique within the steering policy.
	* `pool` - (Optional) The freeform name of a group of one or more records (e.g., a data center or a geographic region) in which this one is included. 
	* `rdata` - (Required) The record's data, as whitespace-delimited tokens in type-specific presentation format. 
	* `rtype` - (Required) The canonical name for the record's type. Only A, AAAA, and CNAME are supported. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `compartment_id` - (Required) The OCID of the compartment containing the steering policy.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name for the steering policy. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}` 
* `health_check_monitor_id` - (Optional) (Updatable) The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with `rdata` matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with `rdata` not matching any monitored endpoint will be assumed healthy. 
* `rules` - (Optional) The pipeline of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.

	The first rule receives a shuffled list of all answers, and every other rule receives the list of answers emitted by the one preceding it. The last rule populates the response. 
	* `cases` - (Optional) 
		* `answer_data` - (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED)
			* `answer_condition` - (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) 
			* `should_keep` - (Applicable when rule_type=FILTER) Keep the answer if the value is `true`.
			* `value` - (Required when rule_type=PRIORITY | WEIGHTED) 
		* `case_condition` - (Applicable when rule_type=FILTER | HEALTH | LIMIT | PRIORITY | WEIGHTED) 
		* `count` - (Required when rule_type=LIMIT) 
	* `default_answer_data` - (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) Defines a default set of answer conditions and values that are applied to an answer when `cases` is not defined for the rule or a matching case does not have any matching `answerCondition`s in its `answerData`. `defaultAnswerData` is **not** applied if `cases` is defined and there are no matching cases. 
		* `answer_condition` - (Applicable when rule_type=FILTER | PRIORITY | WEIGHTED) 
		* `should_keep` - (Applicable when rule_type=FILTER) Keep the answer if the value is `true`.
		* `value` - (Required when rule_type=PRIORITY | WEIGHTED) 
	* `default_count` - (Applicable when rule_type=LIMIT) Defines a default count if `cases` is not defined for the rule or a matching case does not define `count`. `defaultCount` is **not** applied if `cases` is defined and there are no matching cases. 
	* `rule_type` - (Required) The type of a rule determines its sorting/filtering behavior.
		* FILTER rules filter the list of answers (e.g., to remove those with hosts that are down for maintenance). Answers remain if and only if their associated data is `true`.
		* HEALTH rules remove answers from the list if their `rdata` matches a target in the health check monitor referenced by the steering policy and the target is reported down.
		* WEIGHTED rules probabilistically move answers with greater associated integer data to the beginning of the list.
		* PRIORITY rules sort answers by associated integer data, moving those with the lowest values to the beginning of the list without changing the relative order of those with the same value.
		* LIMIT rules filter away answers that are too far down the list. Parameter "count" specifies how many answers to keep. 
* `template` - (Required) (Updatable) The common pattern (or lack thereof) to which the steering policy adheres. This value restricts the possible configurations of rules, but thereby supports specifically tailored interfaces. Values other than "CUSTOM" require the rules to begin with an unconditional FILTER that keeps answers contingent upon `answer.isDisabled != true`, followed _if and only if the policy references a health check monitor_ by an unconditional HEALTH rule, and require the last rule to be an unconditional LIMIT. What must precede the LIMIT rule is determined by the template value:
	* FAILOVER requires exactly an unconditional PRIORITY rule that ranks answers by pool. Each answer pool must have a unique priority value assigned to it. Answer data must be defined in the `defaultAnswerData` property for the rule and the `cases` property must not be defined.
	* LOAD_BALANCE requires exactly an unconditional WEIGHTED rule that shuffles answers by name. Answer data must be defined in the `defaultAnswerData` property for the rule and the `cases` property must not be defined.
	* ROUTE_BY_GEO requires exactly one PRIORITY rule that ranks answers by pool using the geographical location of the client as a condition. Within that rule you may only use `query.client.geoKey` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* ROUTE_BY_ASN requires exactly one PRIORITY rule that ranks answers by pool using the ASN of the client as a condition. Within that rule you may only use `query.client.asn` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* ROUTE_BY_IP requires exactly one PRIORITY rule that ranks answers by pool using the IP subnet of the client as a condition. Within that rule you may only use `query.client.address` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* CUSTOM allows an arbitrary configuration of rules.

	For an existing steering policy, the template value may be changed to any of the supported options but the resulting policy must conform to the requirements for the new template type or else a Bad Request error will be returned. 
* `ttl` - (Optional) (Updatable) The Time To Live for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `answers` - The set of all answers that can potentially issue from the steering policy. 
	* `is_disabled` - Whether or not an answer should be excluded from responses, e.g. because the corresponding server is down for maintenance. Note, however, that such filtering is not automatic and will only take place if a rule implements it. 
	* `name` - A user-friendly name for the answer, unique within the steering policy.
	* `pool` - The freeform name of a group of one or more records (e.g., a data center or a geographic region) in which this one is included. 
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. 
	* `rtype` - The canonical name for the record's type. Only A, AAAA, and CNAME are supported. For more information, see [Resource Record (RR) TYPEs](https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4). 
* `compartment_id` - The OCID of the compartment containing the steering policy.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name for the steering policy. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}` 
* `health_check_monitor_id` - The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with `rdata` matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with `rdata` not matching any monitored endpoint will be assumed healthy. 
* `id` - The OCID of the resource.
* `rules` - The pipeline of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.

	The first rule receives a shuffled list of all answers, and every other rule receives the list of answers emitted by the one preceding it. The last rule populates the response. 
	* `cases` - 
		* `answer_data` - 
			* `answer_condition` - 
			* `should_keep` - Keep the answer if the value is `true`.
			* `value` - 
		* `case_condition` - 
		* `count` - 
	* `default_answer_data` - Defines a default set of answer conditions and values that are applied to an answer when `cases` is not defined for the rule or a matching case does not have any matching `answerCondition`s in its `answerData`. `defaultAnswerData` is **not** applied if `cases` is defined and there are no matching cases. 
		* `answer_condition` - 
		* `should_keep` - Keep the answer if the value is `true`.
		* `value` - 
	* `default_count` - Defines a default count if `cases` is not defined for the rule or a matching case does not define `count`. `defaultCount` is **not** applied if `cases` is defined and there are no matching cases. 
	* `rule_type` - The type of a rule determines its sorting/filtering behavior.
		* FILTER rules filter the list of answers (e.g., to remove those with hosts that are down for maintenance). Answers remain if and only if their associated data is `true`.
		* HEALTH rules remove answers from the list if their `rdata` matches a target in the health check monitor referenced by the steering policy and the target is reported down.
		* WEIGHTED rules probabilistically move answers with greater associated integer data to the beginning of the list.
		* PRIORITY rules sort answers by associated integer data, moving those with the lowest values to the beginning of the list without changing the relative order of those with the same value.
		* LIMIT rules filter away answers that are too far down the list. Parameter "count" specifies how many answers to keep. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `template` - The common pattern (or lack thereof) to which the steering policy adheres. This value restricts the possible configurations of rules, but thereby supports specifically tailored interfaces. Values other than "CUSTOM" require the rules to begin with an unconditional FILTER that keeps answers contingent upon `answer.isDisabled != true`, followed _if and only if the policy references a health check monitor_ by an unconditional HEALTH rule, and require the last rule to be an unconditional LIMIT. What must precede the LIMIT rule is determined by the template value:
	* FAILOVER requires exactly an unconditional PRIORITY rule that ranks answers by pool. Each answer pool must have a unique priority value assigned to it. Answer data must be defined in the `defaultAnswerData` property for the rule and the `cases` property must not be defined.
	* LOAD_BALANCE requires exactly an unconditional WEIGHTED rule that shuffles answers by name. Answer data must be defined in the `defaultAnswerData` property for the rule and the `cases` property must not be defined.
	* ROUTE_BY_GEO requires exactly one PRIORITY rule that ranks answers by pool using the geographical location of the client as a condition. Within that rule you may only use `query.client.geoKey` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* ROUTE_BY_ASN requires exactly one PRIORITY rule that ranks answers by pool using the ASN of the client as a condition. Within that rule you may only use `query.client.asn` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* ROUTE_BY_IP requires exactly one PRIORITY rule that ranks answers by pool using the IP subnet of the client as a condition. Within that rule you may only use `query.client.address` in the `caseCondition` expressions for defining the cases. For each case in the PRIORITY rule each answer pool must have a unique priority value assigned to it. Answer data can only be defined within cases and `defaultAnswerData` cannot be used in the PRIORITY rule.
	* CUSTOM allows an arbitrary configuration of rules.

	For an existing steering policy, the template value may be changed to any of the supported options but the resulting policy must conform to the requirements for the new template type or else a Bad Request error will be returned. 
* `time_created` - The date and time the resource was created in "YYYY-MM-ddThh:mmZ" format with a Z offset, as defined by RFC 3339.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `ttl` - The Time To Live for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used. 

## Import

SteeringPolicies can be imported using the `id`, e.g.

```
$ terraform import oci_dns_steering_policy.test_steering_policy "id"
```

