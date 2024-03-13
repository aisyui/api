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
pub struct CreateGroupRequest {
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "password")]
    pub password: String,
    #[serde(rename = "users", skip_serializing_if = "Option::is_none")]
    pub users: Option<Vec<i32>>,
}

impl CreateGroupRequest {
    pub fn new(name: String, password: String) -> CreateGroupRequest {
        CreateGroupRequest {
            name,
            password,
            users: None,
        }
    }
}

