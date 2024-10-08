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
        "/api/address/geocode": {
            "post": {
                "description": "Converts an address into geographic coordinates.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Geocode address to coordinates",
                "parameters": [
                    {
                        "description": "Address Data",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.ResponseAddress"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coordinate: Lon: Longitude, Lat: Latitude",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/address/search": {
            "post": {
                "description": "Search address based on latitude and longitude.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Search address by coordinates",
                "parameters": [
                    {
                        "description": "Lat and Lng",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.RequestAddressGeocode"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "You are in Country: Country, Town: Town, Road: Road",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Address": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "country_code": {
                    "type": "string"
                },
                "county": {
                    "type": "string"
                },
                "postcode": {
                    "type": "string"
                },
                "road": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "town": {
                    "type": "string"
                }
            }
        },
        "main.RequestAddressGeocode": {
            "type": "object",
            "properties": {
                "lat": {
                    "type": "number"
                },
                "lng": {
                    "type": "number"
                }
            }
        },
        "main.ResponseAddress": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/main.Address"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Address Search API",
	Description:      "This is an API for searching addresses and geocoding.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
