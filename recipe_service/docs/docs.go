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
        "/recipe": {
            "post": {
                "description": "Returns a list of recipes based on the provided ingredients",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get Recipes by Ingredients",
                "parameters": [
                    {
                        "description": "Array of ingredients, e.g., {\\",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.InputIngredient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with a list of recipes",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request: Invalid input or empty array",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/recipe/random": {
            "get": {
                "description": "Returns a list of random recipes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get Random Recipes",
                "responses": {
                    "200": {
                        "description": "List of random recipes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/recipe.Entity"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/recipe/{id}": {
            "get": {
                "description": "Returns a recipe with the specified ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipes"
                ],
                "summary": "Get Recipe By ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Details of the recipe",
                        "schema": {
                            "$ref": "#/definitions/recipe.Entity"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.InputIngredient": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "recipe.Entity": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageUrl": {
                    "type": "string"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/recipe.Ingredient"
                    }
                },
                "instructions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/recipe.Step"
                    }
                },
                "matching": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "usedIngredientCount": {
                    "type": "integer"
                }
            }
        },
        "recipe.Ingredient": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "unit": {
                    "type": "string"
                }
            }
        },
        "recipe.Step": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/recipe.Ingredient"
                    }
                },
                "number": {
                    "type": "integer"
                },
                "step": {
                    "type": "string"
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
