// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
package provider

import (
	"fmt"

	"strconv"

	"encoding/json"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func literalTypeHashCodeForSets(m interface{}) int {
	return hashcode.String(fmt.Sprintf("%v", m))
}

func validateBoolInSlice(valid []bool) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(bool)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be bool", k))
			return
		}

		for _, str := range valid {
			if v == str {
				return
			}
		}

		es = append(es, fmt.Errorf("expected %s to be one of %v, got %t", k, valid, v))
		return
	}
}

func validateNotEmptyString() schema.SchemaValidateFunc {
	return func(i interface{}, k string) (s []string, es []error) {
		v, ok := i.(string)
		if !ok {
			es = append(es, fmt.Errorf("expected type of %s to be string", k))
			return
		}
		if len(v) == 0 {
			es = append(es, fmt.Errorf("%s cannot be an empty string", k))
		}
		return
	}
}

func objectMapToStringMap(rm map[string]interface{}) map[string]string {
	result := map[string]string{}
	for k, v := range rm {
		result[k] = v.(string)
	}
	return result
}

func validateInt64TypeString(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)

	_, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		errors = append(errors, fmt.Errorf("%q (%q) must be a 64-bit integer", k, v))
	}
	return
}

func int64StringDiffSuppressFunction(key string, old string, new string, d *schema.ResourceData) bool {
	//  We can ignore the error since the validate func takes care of this before it reaches here
	oldIntVal, _ := strconv.ParseInt(old, 10, 64)
	newIntVal, _ := strconv.ParseInt(new, 10, 64)
	return oldIntVal == newIntVal
}

func convertMapOfStringSlicesToMapOfStrings(rm map[string][]string) (map[string]string, error) {
	result := map[string]string{}
	for k, v := range rm {
		val, err := json.Marshal(v)
		if err == nil {
			result[k] = string(val)
		} else {
			return nil, err
		}
	}
	return result, nil
}
