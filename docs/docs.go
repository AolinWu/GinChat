// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/index": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/CreateUser": {
            "post": {
                "tags": [
                    "首页"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "userpasswd",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/DeleteUser": {
            "post": {
                "tags": [
                    "首页"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "userpasswd",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/GetUserList": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/QueryUsers": {
            "get": {
                "tags": [
                    "首页"
                ],
                "summary": "通过用户名查询用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/QueryUsersWithPasswd": {
            "get": {
                "tags": [
                    "首页"
                ],
                "summary": "通过用户名和密码查询用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "userpasswd",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
