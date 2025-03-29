# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [optionhub.v1.proto](#optionhub-v1-proto)
    - [GetOptionRequestsOut](#-GetOptionRequestsOut)
    - [OptionRequestItem](#-OptionRequestItem)
    - [SetAttributeByIdIn](#-SetAttributeByIdIn)
  
    - [OptionhubService](#-OptionhubService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="optionhub-v1-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## optionhub.v1.proto



<a name="-GetOptionRequestsOut"></a>

### GetOptionRequestsOut
message response with requested options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| optionRequestItem | [OptionRequestItem](#OptionRequestItem) | repeated | array of items |






<a name="-OptionRequestItem"></a>

### OptionRequestItem
Describe


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| option_request_id | [int64](#int64) |  | id of requested note in db |
| option_request_value | [string](#string) |  | value of requested option |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | time of create note |
| attribute_value | [string](#string) |  | value of attribute where option requested in |
| attribute_id | [string](#string) |  | id of requested attribute |






<a name="-SetAttributeByIdIn"></a>

### SetAttributeByIdIn
message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| attribute_id | [int64](#int64) |  | id of the row in the db |
| value | [string](#string) |  |  |
| parent_id | [int64](#int64) |  |  |





 

 

 


<a name="-OptionhubService"></a>

### OptionhubService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| SetAttributeByID | [.SetAttributeByIdIn](#SetAttributeByIdIn) | [.google.protobuf.Empty](#google-protobuf-Empty) |  |
| GetOptionRequests | [.google.protobuf.Empty](#google-protobuf-Empty) | [.GetOptionRequestsOut](#GetOptionRequestsOut) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

