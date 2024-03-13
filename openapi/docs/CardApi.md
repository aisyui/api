# \CardApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_card**](CardApi.md#create_card) | **POST** /cards | Create a new Card
[**delete_card**](CardApi.md#delete_card) | **DELETE** /cards/{id} | Deletes a Card by ID
[**draw_done**](CardApi.md#draw_done) | **PUT** /cards/{id}/d | Draws a card item as done.
[**draw_start**](CardApi.md#draw_start) | **PATCH** /users/{id}/card/start | Draws a card item as done.
[**list_card**](CardApi.md#list_card) | **GET** /cards | List Cards
[**read_card**](CardApi.md#read_card) | **GET** /cards/{id} | Find a Card by ID
[**read_card_owner**](CardApi.md#read_card_owner) | **GET** /cards/{id}/owner | Find the attached User
[**update_card**](CardApi.md#update_card) | **PATCH** /cards/{id} | Updates a Card



## create_card

> models::CardCreate create_card(create_card_request)
Create a new Card

Creates a new Card and persists it to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**create_card_request** | [**CreateCardRequest**](CreateCardRequest.md) | Card to create | [required] |

### Return type

[**models::CardCreate**](CardCreate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## delete_card

> delete_card(id)
Deletes a Card by ID

Deletes the Card with the requested ID.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Card | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## draw_done

> draw_done(id)
Draws a card item as done.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## draw_start

> draw_start(id)
Draws a card item as done.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** |  | [required] |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## list_card

> Vec<models::CardList> list_card(page, items_per_page)
List Cards

List Cards.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**page** | Option<**i32**> | what page to render |  |
**items_per_page** | Option<**i32**> | item count to render per page |  |

### Return type

[**Vec<models::CardList>**](CardList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_card

> models::CardRead read_card(id)
Find a Card by ID

Finds the Card with the requested ID and returns it.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Card | [required] |

### Return type

[**models::CardRead**](CardRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## read_card_owner

> models::CardOwnerRead read_card_owner(id)
Find the attached User

Find the attached User of the Card with the given ID

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Card | [required] |

### Return type

[**models::CardOwnerRead**](Card_OwnerRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## update_card

> models::CardUpdate update_card(id, update_card_request)
Updates a Card

Updates a Card and persists changes to storage.

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**id** | **i32** | ID of the Card | [required] |
**update_card_request** | [**UpdateCardRequest**](UpdateCardRequest.md) | Card properties to update | [required] |

### Return type

[**models::CardUpdate**](CardUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

