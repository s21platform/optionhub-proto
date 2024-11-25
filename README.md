# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [optionhub.proto](#optionhub-proto)
    - [AddIn](#-AddIn)
    - [AddOut](#-AddOut)
    - [DeleteByIdIn](#-DeleteByIdIn)
    - [DeleteByIdOut](#-DeleteByIdOut)
    - [Empty](#-Empty)
    - [GetAllOut](#-GetAllOut)
    - [GetByIdIn](#-GetByIdIn)
    - [GetByIdOut](#-GetByIdOut)
    - [GetByNameIn](#-GetByNameIn)
    - [GetByNameOut](#-GetByNameOut)
    - [GetPreviewOut](#-GetPreviewOut)
    - [Record](#-Record)
    - [SetByIdIn](#-SetByIdIn)
    - [SetByIdOut](#-SetByIdOut)
  
    - [OptionhubService](#-OptionhubService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="optionhub-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## optionhub.proto



<a name="-AddIn"></a>

### AddIn
message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | value that will be added |






<a name="-AddOut"></a>

### AddOut
message response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | id of added value |
| value | [string](#string) |  |  |






<a name="-DeleteByIdIn"></a>

### DeleteByIdIn
message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | value with such id will be marked as inactive |






<a name="-DeleteByIdOut"></a>

### DeleteByIdOut
message response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ok | [bool](#bool) |  | flag that says if deletion was successful |






<a name="-Empty"></a>

### Empty







<a name="-GetAllOut"></a>

### GetAllOut
message response for all active values


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Record](#Record) | repeated |  |






<a name="-GetByIdIn"></a>

### GetByIdIn
message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | id of the row in the db |






<a name="-GetByIdOut"></a>

### GetByIdOut
message response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="-GetByNameIn"></a>

### GetByNameIn



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="-GetByNameOut"></a>

### GetByNameOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Record](#Record) | repeated |  |






<a name="-GetPreviewOut"></a>

### GetPreviewOut



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [Record](#Record) | repeated |  |






<a name="-Record"></a>

### Record



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |






<a name="-SetByIdIn"></a>

### SetByIdIn
message request


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="-SetByIdOut"></a>

### SetByIdOut
message response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | id of value that will be changed |
| value | [string](#string) |  |  |





 

 

 


<a name="-OptionhubService"></a>

### OptionhubService


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetOsBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetOsPreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetOsByID | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| GetAllOs | [.Empty](#Empty) | [.GetAllOut](#GetAllOut) |  |
| AddOs | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |
| SetOsByID | [.SetByIdIn](#SetByIdIn) | [.SetByIdOut](#SetByIdOut) |  |
| DeleteOsByID | [.DeleteByIdIn](#DeleteByIdIn) | [.DeleteByIdOut](#DeleteByIdOut) |  |
| GetWorkPlaceBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetWorkPlacePreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetWorkPlaceById | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| AddWorkPlace | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |
| GetStudyPlaceBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetStudyPlacePreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetStudyPlaceById | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| AddStudyPlace | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |
| GetHobbyBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetHobbyPreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetHobbyById | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| AddHobby | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |
| GetSkillBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetSkillPreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetSkillById | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| AddSkill | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |
| GetCityBySearchName | [.GetByNameIn](#GetByNameIn) | [.GetByNameOut](#GetByNameOut) |  |
| GetCityPreview | [.Empty](#Empty) | [.GetPreviewOut](#GetPreviewOut) |  |
| GetCityById | [.GetByIdIn](#GetByIdIn) | [.GetByIdOut](#GetByIdOut) |  |
| AddCity | [.AddIn](#AddIn) | [.AddOut](#AddOut) |  |

 



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

