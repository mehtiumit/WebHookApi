// Package hookApi Code generated by swaggo/swag. DO NOT EDIT
package hookApi

import "github.com/swaggo/swag"

const docTemplatehookApi = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/content": {
            "post": {
                "description": "Create content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Create content",
                "parameters": [
                    {
                        "description": "Content",
                        "name": "content",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/content.CreateContentRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/content.CreateContentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        },
        "/v1/content/{id}": {
            "get": {
                "description": "Get content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Content"
                ],
                "summary": "Get content",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Content ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/content.ContentDto"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.CustomError"
                        }
                    }
                }
            }
        },
        "/v1/hook": {
            "post": {
                "description": "Create hook",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Hook"
                ],
                "summary": "Create hook",
                "parameters": [
                    {
                        "description": "Create new hook",
                        "name": "createHookRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/hook.CreateHookRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "type": "int"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "content.ContentDto": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "content.CreateContentRequestDto": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "content.CreateContentResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "hook.CreateHookRequest": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "contentId": {
                    "type": "string"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "models.CustomError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "errorDetail": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfohookApi holds exported Swagger Info so clients can modify it
var SwaggerInfohookApi = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/webhook/api",
	Schemes:          []string{"http"},
	Title:            "Mehti Umit - WebHook Api",
	Description:      "Mehti Umit - WebHook Api",
	InfoInstanceName: "hookApi",
	SwaggerTemplate:  docTemplatehookApi,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfohookApi.InstanceName(), SwaggerInfohookApi)
}
