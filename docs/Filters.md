### Data Sources Filtering

Data sources that return lists of resources support filtering semantics. 
To employ a filter include this block in your data source definition:

```hcl
filter {
	name = ""
	values = [""]
}
```

The `name` value corresponds to the qualified property name to filter with
and the `values` lists can contain one or more values filter with.  

Nested Properties and map elements can be addressed by qualifying the property name with parent property name
Example r1 will give all the instances which have `source_type` image
Example r2 will give all the instances which contain a defined tag with value "42" for key `CostCenter` in the namespace `Operations`
```hcl
data "oci_core_instances" "r1" {
  ...
  filter {
    name = "source_details.source_type"
    value = ["image"]
  }
}

data "oci_core_instances" "r2" {
  ...
  filter {
    name = "defined_tags.Operations.CostCenter"
    value = ["42"]
  }
}

```

Multiple `values` work as an **OR** type filter. In the shape 
example below, the resulting data source would contain both VM 
shapes _Standard 1.1_ and _Standard 1.2_:
```hcl
data "oci_core_shape" "t" {
  ...
  filter {
    name = "name"
    values = ["VM.Standard1.1", "VM.Standard1.2"]
  }
}
```

Multiple filters blocks can be composed to form **AND** type comparisons. The example below will return a data source containing 
_running instances_ in the _first AD_ of a region:
```hcl
data "oci_core_instances" "s" {
	...
  filter {
    name = "availability_domain"
    values = ["\\w*-AD-1"]
    regex = true
  }

  filter {
    name = "state"
    values = ["RUNNING"]
  }
}
        
```

As shown above, filters can also employ regular expressions. By setting
`regex = true`, each item in the `values` list will be treated as a 
regular expression. Note that backslashes in strings for regular
expression special characters need to be escaped with another slash,
shown above as the first `\` before `\w` in `"\\w*-AD-1"`.

### Limitations
Drilling into lists of structured objects is not currently supported. If these properties are targeted no results will be returned from the datasource.
