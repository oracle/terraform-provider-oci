---
subcategory: "DNS"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dns_steering_policy"
sidebar_current: "docs-oci-datasource-dns-steering_policy"
description: |-
  Provides details about a specific Steering Policy in Oracle Cloud Infrastructure DNS service
---

# Data Source: oci_dns_steering_policy
This data source provides details about a specific Steering Policy resource in Oracle Cloud Infrastructure DNS service.

Gets information about the specified steering policy.


## Example Usage

```hcl
data "oci_dns_steering_policy" "test_steering_policy" {
	#Required
	steering_policy_id = oci_dns_steering_policy.test_steering_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `steering_policy_id` - (Required) The OCID of the target steering policy.


## Attributes Reference

The following attributes are exported:

* `answers` - The set of all answers that can potentially issue from the steering policy. 
	* `is_disabled` - Set this property to `true` to indicate that the answer is administratively disabled, such as when the corresponding server is down for maintenance. An answer's `isDisabled` property can be referenced in `answerCondition` properties in rules using `answer.isDisabled`.
	* `name` - A user-friendly name for the answer, unique within the steering policy. An answer's `name` property can be referenced in `answerCondition` properties of rules using `answer.name`.
	* `pool` - The freeform name of a group of one or more records in which this record is included, such as "LAX data center". An answer's `pool` property can be referenced in `answerCondition` properties of rules using `answer.pool`.
	* `rdata` - The record's data, as whitespace-delimited tokens in type-specific presentation format. All RDATA is normalized and the returned presentation of your RDATA may differ from its initial input. For more information about RDATA, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm). 
	* `rtype` - The type of DNS record, such as A or CNAME. Only A, AAAA, and CNAME are supported. For more information, see [Supported DNS Resource Record Types](https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm). 
* `compartment_id` - The OCID of the compartment containing the steering policy.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `display_name` - A user-friendly name for the steering policy. Does not have to be unique and can be changed. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
* `health_check_monitor_id` - The OCID of the health check monitor providing health data about the answers of the steering policy. A steering policy answer with `rdata` matching a monitored endpoint will use the health data of that endpoint. A steering policy answer with `rdata` not matching any monitored endpoint will be assumed healthy.

	 **Note:** To use the Health Check monitoring feature in a steering policy, a monitor must be created using the Health Checks service first. For more information on how to create a monitor, please see [Managing Health Checks](https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Tasks/managinghealthchecks.htm). 
* `id` - The OCID of the resource.
* `rules` - The series of rules that will be processed in sequence to reduce the pool of answers to a response for any given request.

	 The first rule receives a shuffled list of all answers, and every other rule receives the list of answers emitted by the one preceding it. The last rule populates the response. 
	* `cases` - An array of `caseConditions`. A rule may optionally include a sequence of cases defining alternate configurations for how it should behave during processing for any given DNS query. When a rule has no sequence of `cases`, it is always evaluated with the same configuration during processing. When a rule has an empty sequence of `cases`, it is always ignored during processing. When a rule has a non-empty sequence of `cases`, its behavior during processing is configured by the first matching `case` in the sequence. When a rule has no matching cases the rule is ignored. A rule case with no `caseCondition` always matches. A rule case with a `caseCondition` matches only when that expression evaluates to true for the given query. 
		* `answer_data` - An array of `SteeringPolicyPriorityAnswerData` objects.
			* `answer_condition` - An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties. 
			* `should_keep` - Keeps the answer only if the value is `true`.
			* `value` - The rank assigned to the set of answers that match the expression in `answerCondition`. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between `0` and `255`. 
		* `case_condition` - An expression that uses conditions at the time of a DNS query to indicate whether a case matches. Conditions may include the geographical location, IP subnet, or ASN the DNS query originated. **Example:** If you have an office that uses the subnet `192.0.2.0/24` you could use a `caseCondition` expression `query.client.address in ('192.0.2.0/24')` to define a case that matches queries from that office. 
		* `count` - The number of answers allowed to remain after the limit rule has been processed, keeping only the first of the remaining answers in the list. Example: If the `count` property is set to `2` and four answers remain before the limit rule is processed, only the first two answers in the list will remain after the limit rule has been processed. 
	* `default_answer_data` - Defines a default set of answer conditions and values that are applied to an answer when `cases` is not defined for the rule, or a matching case does not have any matching `answerCondition`s in its `answerData`. `defaultAnswerData` is not applied if `cases` is defined and there are no matching cases. In this scenario, the next rule will be processed. 
		* `answer_condition` - An expression that is used to select a set of answers that match a condition. For example, answers with matching pool properties. 
		* `should_keep` - Keeps the answer only if the value is `true`.
		* `value` - The rank assigned to the set of answers that match the expression in `answerCondition`. Answers with the lowest values move to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between `0` and `255`. 
	* `default_count` - Defines a default count if `cases` is not defined for the rule or a matching case does not define `count`. `defaultCount` is **not** applied if `cases` is defined and there are no matching cases. In this scenario, the next rule will be processed. If no rules remain to be processed, the answer will be chosen from the remaining list of answers. 
	* `description` - A user-defined description of the rule's purpose or behavior.
	* `rule_type` - The type of a rule determines its sorting/filtering behavior.
		* `FILTER` - Filters the list of answers based on their defined boolean data. Answers remain only if their `shouldKeep` value is `true`.
		* `HEALTH` - Removes answers from the list if their `rdata` matches a target in the health check monitor referenced by the steering policy and the target is reported down.
		* `WEIGHTED` - Uses a number between 0 and 255 to determine how often an answer will be served in relation to other answers. Anwers with a higher weight will be served more frequently.
		* `PRIORITY` - Uses a defined rank value of answers to determine which answer to serve, moving those with the lowest values to the beginning of the list without changing the relative order of those with the same value. Answers can be given a value between `0` and `255`.
		* `LIMIT` - Filters answers that are too far down the list. Parameter `defaultCount` specifies how many answers to keep. **Example:** If `defaultCount` has a value of `2` and there are five answers left, when the `LIMIT` rule is processed, only the first two answers will remain in the list. 
* `self` - The canonical absolute URL of the resource.
* `state` - The current state of the resource.
* `template` - A set of predefined rules based on the desired purpose of the steering policy. Each template utilizes Traffic Management's rules in a different order to produce the desired results when answering DNS queries.

	 **Example:** The `FAILOVER` template determines answers by filtering the policy's answers using the `FILTER` rule first, then the following rules in succession: `HEALTH`, `PRIORITY`, and `LIMIT`. This gives the domain dynamic failover capability.

	 It is **strongly recommended** to use a template other than `CUSTOM` when creating a steering policy.

	 All templates require the rule order to begin with an unconditional `FILTER` rule that keeps answers contingent upon `answer.isDisabled != true`, except for `CUSTOM`. A defined `HEALTH` rule must follow the `FILTER` rule if the policy references a `healthCheckMonitorId`. The last rule of a template must must be a `LIMIT` rule. For more information about templates and code examples, see [Traffic Management API Guide](https://docs.cloud.oracle.com/iaas/Content/TrafficManagement/Concepts/trafficmanagementapi.htm).

	**Template Types**
	* `FAILOVER` - Uses health check information on your endpoints to determine which DNS answers to serve. If an endpoint fails a health check, the answer for that endpoint will be removed from the list of available answers until the endpoint is detected as healthy.
	* `LOAD_BALANCE` - Distributes web traffic to specified endpoints based on defined weights.
	* `ROUTE_BY_GEO` - Answers DNS queries based on the query's geographic location. For a list of geographic locations to route by, see [Traffic Management Geographic Locations](https://docs.cloud.oracle.com/iaas/Content/TrafficManagement/Reference/trafficmanagementgeo.htm).
	* `ROUTE_BY_ASN` - Answers DNS queries based on the query's originating ASN.
	* `ROUTE_BY_IP` - Answers DNS queries based on the query's IP address.
	* `CUSTOM` - Allows a customized configuration of rules. 
* `time_created` - The date and time the resource was created, expressed in RFC 3339 timestamp format.

	**Example:** `2016-07-22T17:23:59:60Z` 
* `ttl` - The Time To Live (TTL) for responses from the steering policy, in seconds. If not specified during creation, a value of 30 seconds will be used. 

