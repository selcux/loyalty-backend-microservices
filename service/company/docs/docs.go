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
        "/companies": {
            "get": {
                "description": "Get all companies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Read all company data",
                "parameters": [
                    {
                        "description": "Read Companies",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/company.Company"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/company.Company"
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
                "description": "Create a new company",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Create a company",
                "parameters": [
                    {
                        "description": "New Company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/company.CreateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/company.Company"
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
        "/companies/{id}": {
            "get": {
                "description": "Get a company data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Read a company data",
                "parameters": [
                    {
                        "description": "Read Company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/company.Company"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/company.Company"
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
            "put": {
                "description": "Update a company data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Update a company data",
                "parameters": [
                    {
                        "description": "Update Company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/company.UpdateDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/company.Company"
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
                "description": "Delete a company data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "company"
                ],
                "summary": "Delete a company data",
                "parameters": [
                    {
                        "description": "Delete Company",
                        "name": "company",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/company.Company"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/company.Company"
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
        "company.Company": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "company.CreateDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "company.UpdateDto": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "localhost:80",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Company API",
	Description: "This is the company API of LoyaltyDLT project",
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