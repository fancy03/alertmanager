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
  "swagger": "2.0",
  "info": {
    "description": "API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)",
    "title": "Alertmanager API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/alerts": {
      "get": {
        "description": "Get a list of alerts",
        "tags": [
          "alert"
        ],
        "operationId": "getAlerts",
        "parameters": [
          {
            "type": "boolean",
            "default": true,
            "description": "Show active alerts",
            "name": "active",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show silenced alerts",
            "name": "silenced",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show inhibited alerts",
            "name": "inhibited",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show unprocessed alerts",
            "name": "unprocessed",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "A list of matchers to filter alerts by",
            "name": "filter",
            "in": "query"
          },
          {
            "type": "string",
            "description": "A regex matching receivers to filter alerts by",
            "name": "receiver",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Get alerts response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/alert"
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "post": {
        "description": "Create new Alerts",
        "tags": [
          "alert"
        ],
        "operationId": "postAlerts",
        "parameters": [
          {
            "description": "The alerts to create",
            "name": "alerts",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/alerts"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Create alerts response"
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      }
    },
    "/receivers": {
      "get": {
        "description": "Get list of all receivers (name of notification integrations)",
        "tags": [
          "receiver"
        ],
        "operationId": "getReceivers",
        "responses": {
          "200": {
            "description": "Get receivers response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/receiver"
              }
            }
          }
        }
      }
    },
    "/silence/{silenceID}": {
      "get": {
        "description": "Get a silence by its ID",
        "tags": [
          "silence"
        ],
        "operationId": "getSilence",
        "responses": {
          "200": {
            "description": "Get silence response",
            "schema": {
              "$ref": "#/definitions/silence"
            }
          },
          "404": {
            "description": "A silence with the specified ID was not found"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "delete": {
        "description": "Delete a silence by its ID",
        "tags": [
          "silence"
        ],
        "operationId": "deleteSilence",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the silence to get",
            "name": "silenceID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Delete silence response"
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "uuid",
          "description": "ID of the silence to get",
          "name": "silenceID",
          "in": "path",
          "required": true
        }
      ]
    },
    "/silences": {
      "get": {
        "description": "Get a list of silences",
        "tags": [
          "silence"
        ],
        "operationId": "getSilences",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "A list of matchers to filter silences by",
            "name": "filter",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Get silences response",
            "schema": {
              "$ref": "#/definitions/silences"
            }
          },
          "500": {
            "$ref": "#/responses/InternalServerError"
          }
        }
      },
      "post": {
        "description": "Post a new silence or update an existing one",
        "tags": [
          "silence"
        ],
        "operationId": "postSilences",
        "parameters": [
          {
            "description": "The silence to create",
            "name": "silence",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/silence"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Create / update silence response",
            "schema": {
              "type": "object",
              "properties": {
                "silenceID": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/BadRequest"
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Get current status of an Alertmanager instance and its cluster",
        "tags": [
          "general"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "Get status response",
            "schema": {
              "$ref": "#/definitions/alertmanagerStatus"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "alert": {
      "type": "object",
      "properties": {
        "annotations": {
          "$ref": "#/definitions/labelSet"
        },
        "endsAt": {
          "type": "string",
          "format": "date-time"
        },
        "fingerprint": {
          "type": "string"
        },
        "generatorURL": {
          "type": "string",
          "format": "uri"
        },
        "labels": {
          "$ref": "#/definitions/labelSet"
        },
        "receivers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/receiver"
          }
        },
        "startsAt": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "type": "object",
          "properties": {
            "inhibitedBy": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "silencedBy": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "state": {
              "type": "string",
              "enum": [
                "unprocessed",
                "active",
                "suppressed"
              ]
            }
          }
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "alertmanagerStatus": {
      "type": "object",
      "required": [
        "name",
        "statusInCluster",
        "peers"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "peers": {
          "type": "array",
          "minimum": 0,
          "items": {
            "$ref": "#/definitions/peerStatus"
          }
        },
        "statusInCluster": {
          "type": "string"
        }
      }
    },
    "alerts": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/alert"
      }
    },
    "labelSet": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "matcher": {
      "type": "object",
      "properties": {
        "isRegex": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "regex": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "matchers": {
      "type": "array",
      "minItems": 1,
      "items": {
        "$ref": "#/definitions/matcher"
      }
    },
    "peerStatus": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "receiver": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "silence": {
      "type": "object",
      "required": [
        "matchers"
      ],
      "properties": {
        "comment": {
          "type": "string"
        },
        "createdBy": {
          "type": "string"
        },
        "endsAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string"
        },
        "matchers": {
          "$ref": "#/definitions/matchers"
        },
        "startsAt": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "type": "object",
          "properties": {
            "state": {
              "type": "string",
              "enum": [
                "expired",
                "active",
                "pending"
              ]
            }
          }
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "silences": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/silence"
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "string"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "type": "string"
      }
    }
  },
  "tags": [
    {
      "description": "General Alertmanager operations",
      "name": "general"
    },
    {
      "description": "Everything related to Alertmanager receivers",
      "name": "receiver"
    },
    {
      "description": "Everything related to Alertmanager silences",
      "name": "silence"
    },
    {
      "description": "Everything related to Alertmanager alerts",
      "name": "alert"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)",
    "title": "Alertmanager API",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "0.0.1"
  },
  "paths": {
    "/alerts": {
      "get": {
        "description": "Get a list of alerts",
        "tags": [
          "alert"
        ],
        "operationId": "getAlerts",
        "parameters": [
          {
            "type": "boolean",
            "default": true,
            "description": "Show active alerts",
            "name": "active",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show silenced alerts",
            "name": "silenced",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show inhibited alerts",
            "name": "inhibited",
            "in": "query"
          },
          {
            "type": "boolean",
            "default": true,
            "description": "Show unprocessed alerts",
            "name": "unprocessed",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "A list of matchers to filter alerts by",
            "name": "filter",
            "in": "query"
          },
          {
            "type": "string",
            "description": "A regex matching receivers to filter alerts by",
            "name": "receiver",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Get alerts response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/alert"
              }
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "Create new Alerts",
        "tags": [
          "alert"
        ],
        "operationId": "postAlerts",
        "parameters": [
          {
            "description": "The alerts to create",
            "name": "alerts",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/alerts"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Create alerts response"
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/receivers": {
      "get": {
        "description": "Get list of all receivers (name of notification integrations)",
        "tags": [
          "receiver"
        ],
        "operationId": "getReceivers",
        "responses": {
          "200": {
            "description": "Get receivers response",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/receiver"
              }
            }
          }
        }
      }
    },
    "/silence/{silenceID}": {
      "get": {
        "description": "Get a silence by its ID",
        "tags": [
          "silence"
        ],
        "operationId": "getSilence",
        "responses": {
          "200": {
            "description": "Get silence response",
            "schema": {
              "$ref": "#/definitions/silence"
            }
          },
          "404": {
            "description": "A silence with the specified ID was not found"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "description": "Delete a silence by its ID",
        "tags": [
          "silence"
        ],
        "operationId": "deleteSilence",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the silence to get",
            "name": "silenceID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Delete silence response"
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "format": "uuid",
          "description": "ID of the silence to get",
          "name": "silenceID",
          "in": "path",
          "required": true
        }
      ]
    },
    "/silences": {
      "get": {
        "description": "Get a list of silences",
        "tags": [
          "silence"
        ],
        "operationId": "getSilences",
        "parameters": [
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "A list of matchers to filter silences by",
            "name": "filter",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Get silences response",
            "schema": {
              "$ref": "#/definitions/silences"
            }
          },
          "500": {
            "description": "Internal server error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "description": "Post a new silence or update an existing one",
        "tags": [
          "silence"
        ],
        "operationId": "postSilences",
        "parameters": [
          {
            "description": "The silence to create",
            "name": "silence",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/silence"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Create / update silence response",
            "schema": {
              "type": "object",
              "properties": {
                "silenceID": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/status": {
      "get": {
        "description": "Get current status of an Alertmanager instance and its cluster",
        "tags": [
          "general"
        ],
        "operationId": "getStatus",
        "responses": {
          "200": {
            "description": "Get status response",
            "schema": {
              "$ref": "#/definitions/alertmanagerStatus"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "alert": {
      "type": "object",
      "properties": {
        "annotations": {
          "$ref": "#/definitions/labelSet"
        },
        "endsAt": {
          "type": "string",
          "format": "date-time"
        },
        "fingerprint": {
          "type": "string"
        },
        "generatorURL": {
          "type": "string",
          "format": "uri"
        },
        "labels": {
          "$ref": "#/definitions/labelSet"
        },
        "receivers": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/receiver"
          }
        },
        "startsAt": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "type": "object",
          "properties": {
            "inhibitedBy": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "silencedBy": {
              "type": "array",
              "items": {
                "type": "string"
              }
            },
            "state": {
              "type": "string",
              "enum": [
                "unprocessed",
                "active",
                "suppressed"
              ]
            }
          }
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "alertmanagerStatus": {
      "type": "object",
      "required": [
        "name",
        "statusInCluster",
        "peers"
      ],
      "properties": {
        "name": {
          "type": "string"
        },
        "peers": {
          "type": "array",
          "minimum": 0,
          "items": {
            "$ref": "#/definitions/peerStatus"
          }
        },
        "statusInCluster": {
          "type": "string"
        }
      }
    },
    "alerts": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/alert"
      }
    },
    "labelSet": {
      "type": "object",
      "additionalProperties": {
        "type": "string"
      }
    },
    "matcher": {
      "type": "object",
      "properties": {
        "isRegex": {
          "type": "boolean"
        },
        "name": {
          "type": "string"
        },
        "regex": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "matchers": {
      "type": "array",
      "minItems": 1,
      "items": {
        "$ref": "#/definitions/matcher"
      }
    },
    "peerStatus": {
      "type": "object",
      "properties": {
        "address": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "receiver": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    },
    "silence": {
      "type": "object",
      "required": [
        "matchers"
      ],
      "properties": {
        "comment": {
          "type": "string"
        },
        "createdBy": {
          "type": "string"
        },
        "endsAt": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string"
        },
        "matchers": {
          "$ref": "#/definitions/matchers"
        },
        "startsAt": {
          "type": "string",
          "format": "date-time"
        },
        "status": {
          "type": "object",
          "properties": {
            "state": {
              "type": "string",
              "enum": [
                "expired",
                "active",
                "pending"
              ]
            }
          }
        },
        "updatedAt": {
          "type": "string",
          "format": "date-time"
        }
      }
    },
    "silences": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/silence"
      }
    }
  },
  "responses": {
    "BadRequest": {
      "description": "Bad request",
      "schema": {
        "type": "string"
      }
    },
    "InternalServerError": {
      "description": "Internal server error",
      "schema": {
        "type": "string"
      }
    }
  },
  "tags": [
    {
      "description": "General Alertmanager operations",
      "name": "general"
    },
    {
      "description": "Everything related to Alertmanager receivers",
      "name": "receiver"
    },
    {
      "description": "Everything related to Alertmanager silences",
      "name": "silence"
    },
    {
      "description": "Everything related to Alertmanager alerts",
      "name": "alert"
    }
  ]
}`))
}