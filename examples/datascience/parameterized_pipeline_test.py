#!/usr/bin/env python
import json

try:
    # Define the data to write
    data = {
        "ocpus": "2",
        "memory": "32",
        "flexShape": "VM.Standard.E5.Flex",
        "nonFlexShape": "VM.Standard2.1",
        "blockStorage": "100"
    }

    # Write to the JSON file
    with open("/home/datascience/output.json", 'w') as f:
        json.dump(data, f, indent=2)
    print(f"Successfully wrote data to /home/datascience/output.json")

except Exception as e:
    print(f"An error occurred: {str(e)}")