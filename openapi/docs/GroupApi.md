# \GroupApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_group**](GroupApi.md#create_group) | **POST** /groups | Create a new Group
[**delete_group**](GroupApi.md#delete_group) | **DELETE** /groups/{id} | Deletes a Group by ID
[**list_group**](GroupApi.md#list_group) | **GET** /groups | List Groups
[**list_group_users**](GroupApi.md#list_group_users) | **GET** /groups/{id}/users | List attached Users
[**read_group**](GroupApi.md#read_group) | **GET** /groups/{id} | Find a Group by ID
[**update_group**](GroupApi.md#update_group) | **PATCH** /groups/{id} | Updates a Group



## create_group

> models::GroupCreate create_group(create_group_request)
Create a new Group

Creates a new Group and persists it to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_group_request** | [**CreateGroupRequest**](CreateGroupRequest.md) | Group to create | [required] |

### Return type

[**models::GroupCreate**](GroupCreate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_group

> delete_group(id)
Deletes a Group by ID

Deletes the Group with the requested ID.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Group | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_group

> Vec<models::GroupList> list_group(page, items_per_page)
List Groups

List Groups.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::GroupList>**](GroupList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_group_users

> Vec<models::GroupUsersList> list_group_users(id, page, items_per_page)
List attached Users

List attached Users.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Group | [required] |
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::GroupUsersList>**](Group_UsersList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_group

> models::GroupRead read_group(id)
Find a Group by ID

Finds the Group with the requested ID and returns it.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Group | [required] |

### Return type

[**models::GroupRead**](GroupRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_group

> models::GroupUpdate update_group(id, update_group_request)
Updates a Group

Updates a Group and persists changes to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Group | [required] |
**update_group_request** | [**UpdateGroupRequest**](UpdateGroupRequest.md) | Group properties to update | [required] |

### Return type

[**models::GroupUpdate**](GroupUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

