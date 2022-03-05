// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "API文档仅用于Hello Peer研发使用。",
        "contact": {
            "name": "Frank Sheng",
            "email": "726569998@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user/getRandomUsers/{user_id}": {
            "get": {
                "description": "可通过用户ID获取用户详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取广场用户详情",
                "operationId": "get_random_users_details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserDetails"
                        }
                    }
                }
            },
            "post": {
                "description": "可通过用户ID以及用户广场设置项获取随机用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户广场随机用户列表",
                "operationId": "get_random_users_list_in_square",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "分页",
                        "name": "pagination",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Pagination"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.RandomUser"
                            }
                        }
                    }
                }
            }
        },
        "/user/getRecommendSetting/{user_id}": {
            "get": {
                "description": "可通过用户ID获取用户推荐设置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Settings"
                ],
                "summary": "获取用户推荐设置",
                "operationId": "get_recommend_settings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RecommendSetting"
                        }
                    }
                }
            }
        },
        "/user/getRecommendedUserList/{user_id}": {
            "get": {
                "description": "可通过用户ID获取推荐用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取用户推荐用户列表",
                "operationId": "get_recommended_user_list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
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
                                "$ref": "#/definitions/controllers.RecommendedUser"
                            }
                        }
                    }
                }
            }
        },
        "/user/getSquareSetting/{user_id}": {
            "get": {
                "description": "可通过用户ID获取用户广场设置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Settings"
                ],
                "summary": "获取用户广场设置",
                "operationId": "get_square_settings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SquareSettings"
                        }
                    }
                }
            }
        },
        "/user/setRecommendSetting/{user_id}": {
            "post": {
                "description": "可设置用户广场设置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Settings"
                ],
                "summary": "设置用户广场设置",
                "operationId": "set_recommend_settings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "设置",
                        "name": "settings",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RecommendSetting"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/setSquareSetting/{user_id}": {
            "post": {
                "description": "可设置用户广场设置",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Settings"
                ],
                "summary": "设置用户广场设置",
                "operationId": "set_square_settings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户ID",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "设置",
                        "name": "settings",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.SquareSettings"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.RandomUser": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 25
                },
                "coverImageUrl": {
                    "type": "string",
                    "example": "www.coverUrl.com"
                },
                "lat": {
                    "type": "number",
                    "example": 22.51
                },
                "lng": {
                    "type": "number",
                    "example": 113.95
                },
                "location": {
                    "type": "string",
                    "example": "南山区"
                },
                "occupation": {
                    "type": "string",
                    "example": "平面设计师"
                },
                "petName": {
                    "type": "string",
                    "example": "Amy"
                },
                "userName": {
                    "type": "string",
                    "example": "豆豆"
                }
            }
        },
        "controllers.RecommendedUser": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 25
                },
                "coverImageUrl": {
                    "type": "string",
                    "example": "www.coverUrl.com"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "www.imgUrl1.com",
                        " www.imgUrl2.com"
                    ]
                },
                "lat": {
                    "type": "number",
                    "example": 22.51
                },
                "lng": {
                    "type": "number",
                    "example": 113.95
                },
                "location": {
                    "type": "string",
                    "example": "南山区"
                },
                "occupation": {
                    "type": "string",
                    "example": "平面设计师"
                },
                "petName": {
                    "type": "string",
                    "example": "Amy"
                },
                "tags": {
                    "type": "string",
                    "example": "猫控 读书达人 电影爱好者"
                },
                "userName": {
                    "type": "string",
                    "example": "豆豆"
                },
                "verified": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "controllers.UserDetails": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 25
                },
                "constellation": {
                    "type": "string",
                    "example": "处女座"
                },
                "declaration": {
                    "type": "string",
                    "example": "交友宣言"
                },
                "education": {
                    "type": "string",
                    "example": "本科"
                },
                "height": {
                    "type": "string",
                    "example": "165cm"
                },
                "hobbies": {
                    "type": "string",
                    "example": "兴趣爱好"
                },
                "hometown": {
                    "type": "string",
                    "example": "湖南长沙"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "www.imgUrl1.com",
                        " www.imgUrl2.com"
                    ]
                },
                "occupation": {
                    "type": "string",
                    "example": "平面设计师"
                },
                "resident": {
                    "type": "string",
                    "example": "深圳"
                },
                "selfDesc": {
                    "type": "string",
                    "example": "自我描述"
                },
                "tags": {
                    "type": "string",
                    "example": "猫控 读书达人 电影爱好者 旅行者"
                },
                "theOne": {
                    "type": "string",
                    "example": "希望另一半的样子"
                },
                "userName": {
                    "type": "string",
                    "example": "豆豆"
                },
                "weight": {
                    "type": "string",
                    "example": "43kg"
                }
            }
        },
        "models.Pagination": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 4
                },
                "offset": {
                    "type": "integer",
                    "example": 0
                }
            }
        },
        "models.RecommendSetting": {
            "type": "object",
            "properties": {
                "ageMax": {
                    "type": "integer",
                    "example": 40
                },
                "ageMin": {
                    "type": "integer",
                    "example": 22
                },
                "gender": {
                    "type": "string",
                    "example": "女生"
                },
                "hometown": {
                    "type": "string",
                    "example": "同省优先"
                },
                "location": {
                    "type": "string",
                    "example": "只要同城"
                },
                "petLover": {
                    "type": "string",
                    "example": "喜欢就行"
                },
                "tags": {
                    "type": "string",
                    "example": "不限"
                }
            }
        },
        "models.SquareSettings": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string",
                    "example": "女生"
                },
                "location": {
                    "type": "string",
                    "example": "不限"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "1.12.243.73:8686",
	BasePath:         "/api/v0.1",
	Schemes:          []string{},
	Title:            "Hello Peer API",
	Description:      "Hello Peer是一款基于兴趣的社交应用。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
