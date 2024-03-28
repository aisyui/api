/*
 * Ent Schema API
 *
 * This is an auto generated API description made out of an Ent schema definition
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 * Generated by: https://openapi-generator.tech
 */

use crate::models;

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct ListCard400Response {
    #[serde(rename = "code")]
    pub code: i32,
    #[serde(rename = "status")]
    pub status: String,
    #[serde(rename = "errors", default, with = "::serde_with::rust::double_option", skip_serializing_if = "Option::is_none")]
    pub errors: Option<Option<serde_json::Value>>,
}

impl ListCard400Response {
    pub fn new(code: i32, status: String) -> ListCard400Response {
        ListCard400Response {
            code,
            status,
            errors: None,
        }
    }
}
