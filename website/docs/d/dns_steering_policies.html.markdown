---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policies"
sidebar_current: "docs-oci-datasource-dns-steering_policies"
description: |-
  Provides the list of Steering Policies in Oracle Cloud Infrastructure Dns service
---

# Data Source: oci_dns_steering_policies
This data source provides the list of Steering Policies in Oracle Cloud Infrastructure Dns service.

Gets a list of all steering policies in the specified compartment.


## Example Usage

```hcl
data "oci_dns_steering_policies" "test_steering_policies" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.steering_policy_display_name}"
	display_name_contains = "${var.steering_policy_display_name_contains}"
	health_check_monitor_id = "${oci_dns_health_check_monitor.test_health_check_monitor.id}"
	id = "${var.steering_policy_id}"
	state = "${var.steering_policy_state}"
	template = "${var.steering_policy_template}"
	time_created_greater_than_or_equal_to = "${var.steering_policy_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.steering_policy_time_created_less_than}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment the resource belongs to.
* `display_name` - (Optional) The displayName of a resource.
* `display_name_contains` - (Optional) The partial displayName of a resource. Will match any resource whose name (case-insensitive) contains the provided value. 
* `health_check_monitor_id` - (Optional) Search by health check monitor OCID. Will match any resource whose health check monitor id matches the provided value. 
* `id` - (Optional) The OCID of a resource.
* `state` - (Optional) The state of a resource.
* `template` - (Optional) Search by template type. Will match any resource whose template type matches the provided value. 
* `time_created_greater_than_or_equal_to` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created on or after the indicated time. 
* `time_created_less_than` - (Optional) An [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp that states all returned resources were created before the indicated time. 


## Attributes Reference

The following attributes are exported:

* `steering_policies` - The list of steering_policies.

### SteeringPolicy Reference

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
	* `description` - Your description of the rule's purpose and/or behavior.
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

