### Data Sources Filtering

Data sources that return lists of resources support filtering semantics. 
To employ a filter include this block in your data source definition:

```
filter {
	name = ""
	values = [""]
}
```

The `name` value corresponds to the property name to filter with 
and the `values` lists can contain one or more values filter with. 

Multiple `values` work as an **OR** type filter. In the shape 
example below, the resulting data source would contain both VM 
shapes _Standard 1.1_ and _Standard 1.2_:
```
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
```
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
Currently filters can only target top level attributes of a 
resource (or top level arrays of strings). 

Drilling into properties of structured objects or lists of structured objects is not currently 
supported. If these properties are targeted no results will be returned from the datasource.
