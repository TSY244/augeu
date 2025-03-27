// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "BrightPath api",
    "title": "BrightPath",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/version": {
      "get": {
        "security": [],
        "responses": {
          "200": {
            "description": "返回 BrightPath Api 版本号",
            "schema": {
              "$ref": "#/definitions/Version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionFailure": {
      "type": "object",
      "required": [
        "success",
        "from",
        "reason"
      ],
      "properties": {
        "from": {
          "type": "string",
          "default": "unknown"
        },
        "reason": {
          "type": "string",
          "default": "unexpected stack overflow"
        },
        "success": {
          "type": "boolean",
          "default": false
        }
      }
    },
    "PageMeta": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "default": 1
        },
        "size": {
          "type": "integer",
          "default": 10
        },
        "total": {
          "type": "integer"
        }
      }
    },
    "UnauthorizedError": {
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "code": {
          "description": "403",
          "type": "integer"
        },
        "message": {
          "description": "没有权限",
          "type": "string"
        }
      }
    },
    "Version": {
      "type": "object",
      "properties": {
        "version": {
          "description": "版本号",
          "type": "string",
          "default": "0.0.1"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "BrightPath api",
    "title": "BrightPath",
    "version": "0.0.1"
  },
  "basePath": "/api/v1",
  "paths": {
    "/version": {
      "get": {
        "security": [],
        "responses": {
          "200": {
            "description": "返回 BrightPath Api 版本号",
            "schema": {
              "$ref": "#/definitions/Version"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ActionFailure": {
      "type": "object",
      "required": [
        "success",
        "from",
        "reason"
      ],
      "properties": {
        "from": {
          "type": "string",
          "default": "unknown"
        },
        "reason": {
          "type": "string",
          "default": "unexpected stack overflow"
        },
        "success": {
          "type": "boolean",
          "default": false
        }
      }
    },
    "PageMeta": {
      "type": "object",
      "properties": {
        "page": {
          "type": "integer",
          "default": 1
        },
        "size": {
          "type": "integer",
          "default": 10
        },
        "total": {
          "type": "integer"
        }
      }
    },
    "UnauthorizedError": {
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "code": {
          "description": "403",
          "type": "integer"
        },
        "message": {
          "description": "没有权限",
          "type": "string"
        }
      }
    },
    "Version": {
      "type": "object",
      "properties": {
        "version": {
          "description": "版本号",
          "type": "string",
          "default": "0.0.1"
        }
      }
    }
  }
}`))
}
