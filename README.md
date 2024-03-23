# Jsons

This repository provides a comprehensive toolkit for traversing, accessing, updating, and deleting elements within JSON data structures using customizable paths. Whether you're working with complex JSON files or APIs, this repository equips you with the necessary tools to efficiently manipulate JSON objects with ease.

Key Features:

1. JSON Path Traversal: Easily navigate through JSON structures using intuitive path expressions to locate specific elements or nested objects.
2. Get Operations: Retrieve values from JSON objects based on specified paths, allowing for seamless extraction of data.
3. Update Functionality: Modify JSON values effortlessly by updating elements identified by their corresponding paths.
4. Delete Capability: Remove unwanted elements or entire branches from JSON structures by specifying the paths to be deleted.
5. Flexible Path Syntax: Supports various path syntaxes, including dot notation, bracket notation, and wildcard characters, enabling versatile traversal options.
6. Error Handling: Robust error handling mechanisms to gracefully manage edge cases and invalid paths.
7. Customization: Tailor the repository to your specific needs with customizable options for path parsing and operation execution.
8. Documentation and Examples: Comprehensive documentation and usage examples to facilitate quick integration and adoption.
9. Compatibility: Compatible with popular programming languages and environments, ensuring seamless integration into your existing projects.
10. Active Maintenance: Regular updates and maintenance to address issues, incorporate feedback, and enhance functionality based on community contributions.

Function: Get

Description:
- Get retrieves the value corresponding to the specified key path from the provided JSON data. It supports accessing nested values using a JSON path format, allowing for flexible extraction of data from complex JSON structures.

Parameters:

- jsonData (map[string]interface{}): The JSON data from which the value is to be retrieved.
- key (string): The key path indicating the location of the desired value within the JSON data. This key path can be a simple key or a JSON path containing nested keys separated by dots (.) and enclosed within ${}.

Returns:

- interface{}: The value corresponding to the specified key path within the JSON data.
- error: An error value, if any occurred during the retrieval process. If the retrieval is successful, this will be nil; otherwise, it will contain information about the encountered error.

Function: Set

Description:
- Set updates the value at the specified key path within the provided JSON data. It supports modifying nested values using a JSON path format, allowing for seamless updating of data within complex JSON structures.

Parameters:

- jsonMap (map[string]interface{}): The JSON data in which the value is to be updated.
- key (string): The key path indicating the location where the value should be updated within the JSON data. This key path can be a simple key or a JSON path containing nested keys separated by dots (.) and enclosed within ${}.
- value (interface{}): The new value to be set at the specified key path within the JSON data.
  
Function: Delete

Description:
- Delete removes the value at the specified key path from the provided JSON data. It supports deleting nested values and, if necessary, removing empty parent keys along the specified path.

Parameters:

- jsonMap (map[string]interface{}): The JSON data from which the value is to be deleted.
- key (string): The key path indicating the location of the value to be deleted within the JSON data. This key path can be a simple key or a JSON path containing nested 
keys separated by dots (.) and enclosed within ${}.

