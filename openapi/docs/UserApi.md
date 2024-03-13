# \UserApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_user**](UserApi.md#create_user) | **POST** /users | Create a new User
[**delete_user**](UserApi.md#delete_user) | **DELETE** /users/{id} | Deletes a User by ID
[**list_user**](UserApi.md#list_user) | **GET** /users | List Users
[**list_user_card**](UserApi.md#list_user_card) | **GET** /users/{id}/card | List attached Cards
[**list_user_ue**](UserApi.md#list_user_ue) | **GET** /users/{id}/ue | List attached Ues
[**read_user**](UserApi.md#read_user) | **GET** /users/{id} | Find a User by ID
[**update_user**](UserApi.md#update_user) | **PATCH** /users/{id} | Updates a User



## create_user

> models::UserCreate create_user(create_user_request)
Create a new User

Creates a new User and persists it to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_user_request** | [**CreateUserRequest**](CreateUserRequest.md) | User to create | [required] |

### Return type

[**models::UserCreate**](UserCreate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_user

> delete_user(id)
Deletes a User by ID

Deletes the User with the requested ID.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the User | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_user

> Vec<models::UserList> list_user(page, items_per_page)
List Users

List Users.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::UserList>**](UserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_user_card

> Vec<models::UserCardList> list_user_card(id, page, items_per_page)
List attached Cards

List attached Cards.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the User | [required] |
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::UserCardList>**](User_CardList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_user_ue

> Vec<models::UserUeList> list_user_ue(id, page, items_per_page)
List attached Ues

List attached Ues.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the User | [required] |
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::UserUeList>**](User_UeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_user

> models::UserRead read_user(id)
Find a User by ID

Finds the User with the requested ID and returns it.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the User | [required] |

### Return type

[**models::UserRead**](UserRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_user

> models::UserUpdate update_user(id, update_user_request)
Updates a User

Updates a User and persists changes to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the User | [required] |
**update_user_request** | [**UpdateUserRequest**](UpdateUserRequest.md) | User properties to update | [required] |

### Return type

[**models::UserUpdate**](UserUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

