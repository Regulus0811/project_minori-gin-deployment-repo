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
        "/example/helloworld": {
            "get": {
                "description": "do ping",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
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
        "/gb": {
            "get": {
                "description": "全てのグループ掲示板のリストを取得します。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "全てのグループ掲示板を取得",
                "responses": {
                    "200": {
                        "description": "全てのグループ掲示板のリスト",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.GroupBoard"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "サーバーエラーが発生しました",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new group board with the provided information, including image upload.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "Create a new group board",
                "parameters": [
                    {
                        "description": "Create group board",
                        "name": "group_board_create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupBoardCreateDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Upload image file",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Group board created successfully",
                        "schema": {
                            "$ref": "#/definitions/models.GroupBoard"
                        }
                    },
                    "400": {
                        "description": "Invalid request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gb/announced": {
            "get": {
                "description": "公告されたグループ掲示板のリストを取得します。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "公告されたグループ掲示板を取得",
                "responses": {
                    "200": {
                        "description": "公告されたグループ掲示板のリスト",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/models.GroupBoard"
                                }
                            }
                        }
                    },
                    "500": {
                        "description": "サーバーエラーが発生しました",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gb/{id}": {
            "get": {
                "description": "指定されたIDのグループ掲示板の詳細を取得します。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "IDでグループ掲示板を取得",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "グループ掲示板ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "グループ掲示板が取得されました",
                        "schema": {
                            "$ref": "#/definitions/models.GroupBoard"
                        }
                    },
                    "400": {
                        "description": "無効なリクエストです",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "コードが見つかりません",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "サーバーエラーが発生しました",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "指定されたIDのグループ掲示板を削除します。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "グループ掲示板を削除",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "グループ掲示板ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "グループ掲示板が正常に削除されました",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "コードが見つかりません",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "サーバーエラーが発生しました",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "指定されたIDのグループ掲示板の詳細を更新します。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_board"
                ],
                "summary": "グループ掲示板を更新",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "グループ掲示板ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "グループ掲示板の更新",
                        "name": "group_board_update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GroupBoardUpdateDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "グループ掲示板が正常に更新されました",
                        "schema": {
                            "$ref": "#/definitions/models.GroupBoard"
                        }
                    },
                    "400": {
                        "description": "リクエストが不正です",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "コードが見つかりません",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "サーバーエラーが発生しました",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gc/checkSecretExists": {
            "get": {
                "description": "指定されたグループコードにシークレットがあるかどうかをチェックする。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_code"
                ],
                "summary": "グループコードにシークレットが存在するかチェック",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code to check",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "secretExists\" \"シークレットが存在します",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "無効なリクエストです",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "コードが見つかりません",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gc/verifyGroupCode": {
            "get": {
                "description": "グループコードと、該当する場合はそのシークレットを確認する。",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "group_code"
                ],
                "summary": "グループコードとシークレットを検証",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Code to verify",
                        "name": "code",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Secret for the code",
                        "name": "secret",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "グループコードが検証されました",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "無効なリクエストです",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "シークレットが一致しません",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "コードが見つかりません",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GroupBoardCreateDTO": {
            "type": "object",
            "required": [
                "cid",
                "content",
                "title",
                "uid"
            ],
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "content": {
                    "type": "string",
                    "example": "Sample Content"
                },
                "image": {
                    "type": "string"
                },
                "is_announced": {
                    "type": "boolean",
                    "default": false
                },
                "title": {
                    "type": "string",
                    "example": "Sample Title"
                },
                "uid": {
                    "type": "integer"
                }
            }
        },
        "dto.GroupBoardUpdateDTO": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "is_announced": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.Class": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "limitation": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.GroupBoard": {
            "type": "object",
            "properties": {
                "cid": {
                    "description": "Class ID",
                    "type": "integer"
                },
                "class": {
                    "$ref": "#/definitions/models.Class"
                },
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "isAnnounced": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                },
                "uid": {
                    "description": "User ID",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "name": {
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
