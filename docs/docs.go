// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://uzsk.iamtakagi.net",
        "contact": {
            "name": "yude",
            "email": "i@yude.jp"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/license/mit"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/discord/{uuid}": {
            "get": {
                "description": "fetch UUID from provided Discord ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "discord"
                ],
                "summary": "fetch UUID from provided Discord ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "who to retrieve",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/avatar/{part}/bedrock/{xuid}": {
            "get": {
                "description": "Get the specified part of player's skin image",
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "avatar"
                ],
                "summary": "Get player's skin image",
                "parameters": [
                    {
                        "type": "string",
                        "description": "which part to retrieve",
                        "name": "part",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "XUID of target Bedrock player",
                        "name": "xuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/login": {
            "get": {
                "description": "Login with Discord",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Login with Discord",
                "responses": {}
            }
        },
        "/login/callback": {
            "get": {
                "description": "callback endpoint for Discord login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "callback endpoint for Discord login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer token",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Random state for validating request",
                        "name": "state",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.JWTResponse"
                            }
                        },
                        "headers": {
                            "X-Auth-Token": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "description": "retrieve information of authenticated user",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "retrieve information of authenticated user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JSON Web Token",
                        "name": "X-Auth-Token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Me"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/profiles": {
            "get": {
                "description": "Get a list of profiles with optional filtering and sorting, etc.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Get profiles",
                "parameters": [
                    {
                        "type": "string",
                        "default": "",
                        "example": "550e8400-e29b-41d4-a716-446655440000",
                        "description": "Filter criteria",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "asc",
                        "example": "desc",
                        "description": "Sort order",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "example": 0,
                        "description": "Offset for pagination",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 50,
                        "example": 10,
                        "description": "Limit for pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "play_time",
                        "description": "Order by field",
                        "name": "order_by",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Profile"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/profiles/{uuid}": {
            "get": {
                "description": "Get a profile by UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profiles"
                ],
                "summary": "Get profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "UUID of target profile",
                        "name": "uuid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Profile"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        },
        "/servers": {
            "get": {
                "description": "Get servers registered to uzsk-api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Get servers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/structs.Server"
                            }
                        }
                    }
                }
            }
        },
        "/servers/{name}": {
            "get": {
                "description": "Get servers registered to uzsk-api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Get server",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of target server",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.Server"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/structs.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "structs.Avatar": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "face": {
                    "type": "string"
                },
                "head": {
                    "type": "string"
                }
            }
        },
        "structs.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "structs.JWTResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "structs.Me": {
            "type": "object",
            "properties": {
                "session_expire_at": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "structs.Profile": {
            "type": "object",
            "properties": {
                "avatar": {
                    "$ref": "#/definitions/structs.Avatar"
                },
                "currency": {
                    "type": "integer"
                },
                "experience": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "initial_login_date": {
                    "type": "string"
                },
                "is_bedrock": {
                    "type": "boolean"
                },
                "last_login_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "total_build_blocks": {
                    "type": "integer"
                },
                "total_destroy_blocks": {
                    "type": "integer"
                },
                "total_mob_kills": {
                    "type": "integer"
                },
                "total_play_time": {
                    "type": "integer"
                },
                "uuid": {
                    "type": "string"
                },
                "xuid": {
                    "type": "string"
                }
            }
        },
        "structs.Server": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "uzsk-api.iamtakagi.net",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "uzsk-api",
	Description:      "Public Web API for uzsk.iamtakagi.net",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
