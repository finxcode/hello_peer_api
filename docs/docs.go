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
        "/auth/authlogin": {
            "post": {
                "description": "通过传递微信登录code以及getUserProfile返回字段换取用户token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "授权用户登录",
                "operationId": "authlogin",
                "parameters": [
                    {
                        "description": "微信登录code以及getUserProfile返回字段",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserProfileForm"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    }
                }
            }
        },
        "/auth/autologin": {
            "post": {
                "description": "通过传递微信登录code换取用户token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "未授权用户登录",
                "operationId": "autologin",
                "parameters": [
                    {
                        "description": "微信登录code",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AutoLogin"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    }
                }
            }
        },
        "/settings/getRecommendSetting": {
            "get": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        "/settings/getSquareSetting": {
            "get": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        "/settings/setRecommendSetting": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        "/settings/setSquareSetting": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        },
        "/user/getRandomUserDetails/{user_id}": {
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
            }
        },
        "/user/getRandomUsers": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        "/user/getRecommendedUserList": {
            "get": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
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
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
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
        "/user/setUserAvatar": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "用户可设置头像",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "设置用户头像",
                "operationId": "set_user_avatar",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "头像文件",
                        "name": "content",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/user/setUserBasicInfo": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "可通过用户token设置用户基础信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "设置用户基础信息",
                "operationId": "set_user_basic_info",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户基础信息",
                        "name": "basicInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.BasicInfo"
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
        "/user/setUserGender": {
            "post": {
                "security": [
                    {
                        "x-token": []
                    }
                ],
                "description": "可通过用户token设置用户性别",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "设置用户性别",
                "operationId": "set_user_gender",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "用户性别",
                        "name": "gender",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserGender"
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
                "location": {
                    "type": "string",
                    "example": "深圳"
                },
                "occupation": {
                    "type": "string",
                    "example": "平面设计师"
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
        "models.AutoLogin": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "xtdad-fdfdsf"
                }
            }
        },
        "models.BasicInfo": {
            "type": "object",
            "properties": {
                "birth_day": {
                    "type": "string",
                    "example": "2010-10-09"
                },
                "constellation": {
                    "type": "string",
                    "example": "白羊座"
                },
                "education": {
                    "type": "string",
                    "example": "本科"
                },
                "gender": {
                    "type": "integer",
                    "example": 1
                },
                "height": {
                    "type": "number",
                    "example": 178
                },
                "hometown": {
                    "type": "string",
                    "example": "重庆"
                },
                "location": {
                    "type": "string",
                    "example": "深圳"
                },
                "marriage": {
                    "type": "string",
                    "example": "未婚"
                },
                "occupation": {
                    "type": "string",
                    "example": "设计师"
                },
                "weight": {
                    "type": "number",
                    "example": 56.5
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
                "age_Max": {
                    "type": "integer",
                    "example": 40
                },
                "age_Min": {
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
                "pet_Lover": {
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
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "cdaefj93sds-eqedsdsdsa-3sadasdsss"
                },
                "expiresIn": {
                    "type": "integer",
                    "example": 186624000
                },
                "tokenType": {
                    "type": "string",
                    "example": "bearer"
                }
            }
        },
        "models.UserGender": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "models.UserProfileForm": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "xtdad-fdfdsf"
                },
                "encryptedData": {
                    "type": "string",
                    "example": "7SfFtStsHqKZYhbIkke3BH2bCRzGD15T0jEiUtuksrl9lDeHm9LsPmswJymBXuinPCiXkZhd/uq7s7pACTvbWuvvoKEwz5fAJ6Vr9bTx79XVxiIN4r+Fwm6QHO9DjPkFrxTGAZvMYLyH6IOyOV/nmmlMoBM3G4peSnBi1qCYukwlyCMNp67lb93wSiPAoI7eRhYYw8ayPTsZ/MAJ9CBBUiCwM5aFOUWrMKNTikeq7YVjNCv7KCz0LJTrMKda0YMS0J/034L8x9vJ1OnIkxlWVMQEy/f55IfWVHI1I1fSKd5azzyVKXCbWDpU0PLJnU8XM/l4L7ZUlDOcRMR5KQVGhB9rIjVkykdXUPQK87v8lpnitslK06XceOJqDjK6mRkhJWOYpFUozZa6idFV6xmLZX8bkBsLxczzp1h/satEH7rIz3nKbxd3O1c+3dI2soSt8qFtaumcGdwhenTm+at0gxccAp8JD8PZiB5ZDLTofZIQ4RmI004SIExYUDZUje9mZO+3aC8McVwzrEyK7NKD/NZ5/dYPgDRwzBl1Vm99niY="
                },
                "iv": {
                    "type": "string",
                    "example": "z3tGYrgMcbLzd0qXqZuduQ=="
                },
                "rawData": {
                    "type": "string",
                    "example": ""
                },
                "signature": {
                    "type": "string",
                    "example": ""
                },
                "userInfo": {
                    "type": "string",
                    "example": ""
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
