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
pub struct UpdateGroupRequest {
    #[serde(rename = "name", skip_serializing_if = "Option::is_none")]
    pub name: Option<String>,
    #[serde(rename = "users", skip_serializing_if = "Option::is_none")]
    pub users: Option<Vec<i32>>,
}

impl UpdateGroupRequest {
    pub fn new() -> UpdateGroupRequest {
        UpdateGroupRequest {
            name: None,
            users: None,
        }
    }
}

