# \UeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_ue**](UeApi.md#create_ue) | **POST** /ues | Create a new Ue
[**delete_ue**](UeApi.md#delete_ue) | **DELETE** /ues/{id} | Deletes a Ue by ID
[**list_ue**](UeApi.md#list_ue) | **GET** /ues | List Ues
[**read_ue**](UeApi.md#read_ue) | **GET** /ues/{id} | Find a Ue by ID
[**read_ue_owner**](UeApi.md#read_ue_owner) | **GET** /ues/{id}/owner | Find the attached User
[**update_ue**](UeApi.md#update_ue) | **PATCH** /ues/{id} | Updates a Ue



## create_ue

> models::UeCreate create_ue(create_ue_request)
Create a new Ue

Creates a new Ue and persists it to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_ue_request** | [**CreateUeRequest**](CreateUeRequest.md) | Ue to create | [required] |

### Return type

[**models::UeCreate**](UeCreate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_ue

> delete_ue(id)
Deletes a Ue by ID

Deletes the Ue with the requested ID.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Ue | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_ue

> Vec<models::UeList> list_ue(page, items_per_page)
List Ues

List Ues.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::UeList>**](UeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_ue

> models::UeRead read_ue(id)
Find a Ue by ID

Finds the Ue with the requested ID and returns it.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Ue | [required] |

### Return type

[**models::UeRead**](UeRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_ue_owner

> models::UeOwnerRead read_ue_owner(id)
Find the attached User

Find the attached User of the Ue with the given ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Ue | [required] |

### Return type

[**models::UeOwnerRead**](Ue_OwnerRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_ue

> models::UeUpdate update_ue(id, update_ue_request)
Updates a Ue

Updates a Ue and persists changes to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Ue | [required] |
**update_ue_request** | [**UpdateUeRequest**](UpdateUeRequest.md) | Ue properties to update | [required] |

### Return type

[**models::UeUpdate**](UeUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

