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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/control/crash": {
            "put": {
                "description": "Make Dobby kill itself",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Suicide"
            }
        },
        "/control/goturbo/cpu": {
            "put": {
                "description": "Make Dobby create a CPU spike",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "CPU Spike",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/control/goturbo/memory": {
            "put": {
                "description": "Make Dobby create a memory spike",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Memory Spike",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/control/health/perfect": {
            "put": {
                "description": "Make Dobby healthy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Make Healthy",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/control/health/sick": {
            "put": {
                "description": "Make Dobby sick or unhealthy",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Make Unhealthy",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recover health after sometime (seconds) - E.g. 2",
                        "name": "resetInSeconds",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/control/ready/perfect": {
            "put": {
                "description": "Make Dobby ready",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Make Ready",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/control/ready/sick": {
            "put": {
                "description": "Make Dobby unready",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Control"
                ],
                "summary": "Make Unready",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recover readiness after sometime (seconds) - E.g. 2",
                        "name": "resetInSeconds",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ControlSuccess"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Get Dobby's health status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Dobby Health",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Health"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Health"
                        }
                    }
                }
            }
        },
        "/meta": {
            "get": {
                "description": "Get Dobby's metadata",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Dobby Metadata",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Metadata"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/ready": {
            "get": {
                "description": "Get Dobby's readiness",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Dobby Ready",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Ready"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/model.Ready"
                        }
                    }
                }
            }
        },
        "/return/{statusCode}": {
            "get": {
                "description": "Ask Dobby to return the status code sent by the client",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Repeat Status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Status Code - E.g. 200",
                        "name": "statusCode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Dela(milliseconds) - E.g. 1000",
                        "name": "delay",
                        "in": "query"
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "Get Dobby's version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Dobby Version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Version"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "503": {
                        "description": "Service Unavailable",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ControlSuccess": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "model.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "something went wrong"
                }
            }
        },
        "model.Health": {
            "type": "object",
            "properties": {
                "healthy": {
                    "type": "boolean"
                }
            }
        },
        "model.Metadata": {
            "type": "object",
            "properties": {
                "hostname": {
                    "type": "string",
                    "example": "dobby"
                },
                "ip": {
                    "type": "string",
                    "example": "192.168.1.100"
                }
            }
        },
        "model.Ready": {
            "type": "object",
            "properties": {
                "ready": {
                    "type": "boolean"
                }
            }
        },
        "model.Version": {
            "type": "object",
            "properties": {
                "version": {
                    "type": "string",
                    "example": "1.0.0"
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
