// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/items": {
            "get": {
                "description": "Read all items",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Read all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/item.Entity"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create an item data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Create an item data",
                "parameters": [
                    {
                        "description": "Create an item",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/item.CreateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/item.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "Read an item data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Read an item data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/item.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an item",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Delete an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an item",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "item"
                ],
                "summary": "Update an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update all items",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/item.UpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/item.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "item.CreateDto": {
            "type": "object",
            "required": [
                "name",
                "point"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                },
                "product": {
                    "type": "string"
                }
            }
        },
        "item.Entity": {
            "type": "object",
            "required": [
                "code",
                "company",
                "name",
                "point",
                "product"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                },
                "product": {
                    "type": "string"
                }
            }
        },
        "item.UpdateDto": {
            "type": "object",
            "required": [
                "name",
                "point"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                }
            }
        },
        "util.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Item API",
	Description: "This is the item API of LoyaltyDLT project",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
