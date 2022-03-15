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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "ACCWeb project",
            "url": "https://github.com/assetto-corsa-web/accweb"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/assetto-corsa-web/accweb/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/instance": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Create new acc instance information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Create new acc instance information",
                "parameters": [
                    {
                        "description": "Instance data",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.SaveInstancePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.InstancePayload"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get acc instance information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Get acc instance information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.InstancePayload"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Saves acc instance information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Saves acc instance information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Instance data",
                        "name": "instance",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/app.SaveInstancePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.InstancePayload"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Delete acc instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Delete acc instance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}/clone": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Clones acc instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Clones acc instance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}/export": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get acc instance configuration files",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Get acc instance configuration files",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Zip file with all cofiguration files",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}/live": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get acc instance live information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Get acc instance live information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.LiveServerInstancePayload"
                        }
                    },
                    "404": {
                        "description": ""
                    }
                }
            }
        },
        "/instance/{id}/logs": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Get acc instance logs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Get acc instance logs",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.accWebInstanceLogs"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}/start": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Starts acc instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Starts acc instance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/instance/{id}/stop": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Stops acc instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "instances"
                ],
                "summary": "Stops acc instance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Instance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.AccWError"
                        }
                    }
                }
            }
        },
        "/servers": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "List all ACC dedicated servers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "List all ACC dedicated servers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/app.ListServerItem"
                            }
                        }
                    }
                }
            }
        },
        "/servers/stop-all": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Stops all running ACC dedicated servers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "servers"
                ],
                "summary": "Stops all running ACC dedicated servers",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "app.AccWError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "app.InstancePayload": {
            "type": "object",
            "properties": {
                "acc": {
                    "$ref": "#/definitions/instance.AccConfigFiles"
                },
                "accWeb": {
                    "$ref": "#/definitions/instance.AccWebConfigJson"
                },
                "id": {
                    "type": "string"
                },
                "is_running": {
                    "type": "boolean"
                },
                "path": {
                    "type": "string"
                },
                "pid": {
                    "type": "integer"
                }
            }
        },
        "app.ListServerItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "isRunning": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "nrClients": {
                    "type": "integer"
                },
                "pid": {
                    "type": "integer"
                },
                "serverState": {
                    "type": "string"
                },
                "sessionPhase": {
                    "type": "string"
                },
                "sessionType": {
                    "type": "string"
                },
                "tcpPort": {
                    "type": "integer"
                },
                "track": {
                    "type": "string"
                },
                "udpPort": {
                    "type": "integer"
                }
            }
        },
        "app.LiveServerInstancePayload": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "isRunning": {
                    "type": "boolean"
                },
                "live": {
                    "$ref": "#/definitions/instance.LiveState"
                },
                "name": {
                    "type": "string"
                },
                "nrClients": {
                    "type": "integer"
                },
                "pid": {
                    "type": "integer"
                },
                "serverState": {
                    "type": "string"
                },
                "sessionPhase": {
                    "type": "string"
                },
                "sessionType": {
                    "type": "string"
                },
                "tcpPort": {
                    "type": "integer"
                },
                "track": {
                    "type": "string"
                },
                "udpPort": {
                    "type": "integer"
                }
            }
        },
        "app.SaveInstancePayload": {
            "type": "object",
            "properties": {
                "acc": {
                    "$ref": "#/definitions/instance.AccConfigFiles"
                },
                "accWeb": {
                    "$ref": "#/definitions/instance.AccWebConfigJson"
                }
            }
        },
        "app.accWebInstanceLogs": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "logs": {
                    "type": "string"
                }
            }
        },
        "instance.AccConfigFiles": {
            "type": "object",
            "properties": {
                "assistRules": {
                    "$ref": "#/definitions/instance.AssistRulesJson"
                },
                "bop": {
                    "$ref": "#/definitions/instance.BopJson"
                },
                "configuration": {
                    "$ref": "#/definitions/instance.ConfigurationJson"
                },
                "entrylist": {
                    "$ref": "#/definitions/instance.EntrylistJson"
                },
                "event": {
                    "$ref": "#/definitions/instance.EventJson"
                },
                "eventRules": {
                    "$ref": "#/definitions/instance.EventRulesJson"
                },
                "settings": {
                    "$ref": "#/definitions/instance.SettingsJson"
                }
            }
        },
        "instance.AccWebConfigJson": {
            "type": "object",
            "properties": {
                "autoStart": {
                    "type": "boolean"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "md5Sum": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "instance.AssistRulesJson": {
            "type": "object",
            "properties": {
                "configVersion": {
                    "type": "integer"
                },
                "disableAutoClutch": {
                    "type": "integer"
                },
                "disableAutoEngineStart": {
                    "type": "integer"
                },
                "disableAutoGear": {
                    "type": "integer"
                },
                "disableAutoLights": {
                    "type": "integer"
                },
                "disableAutoPitLimiter": {
                    "type": "integer"
                },
                "disableAutoWiper": {
                    "type": "integer"
                },
                "disableAutosteer": {
                    "type": "integer"
                },
                "disableIdealLine": {
                    "type": "integer"
                },
                "stabilityControlLevelMax": {
                    "type": "integer"
                }
            }
        },
        "instance.BopJson": {
            "type": "object",
            "properties": {
                "configVersion": {
                    "type": "integer"
                },
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.BopSettings"
                    }
                }
            }
        },
        "instance.BopSettings": {
            "type": "object",
            "properties": {
                "ballastKg": {
                    "type": "integer"
                },
                "carModel": {
                    "type": "integer"
                },
                "restrictor": {
                    "type": "integer"
                },
                "track": {
                    "type": "string"
                }
            }
        },
        "instance.CarState": {
            "type": "object",
            "properties": {
                "bestLapMS": {
                    "type": "integer"
                },
                "carID": {
                    "type": "integer"
                },
                "carModel": {
                    "type": "integer"
                },
                "currentDriver": {
                    "$ref": "#/definitions/instance.DriverState"
                },
                "drivers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.DriverState"
                    }
                },
                "fuel": {
                    "type": "integer"
                },
                "laps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.LapState"
                    }
                },
                "lastLapMS": {
                    "type": "integer"
                },
                "lastLapTimestampMS": {
                    "type": "integer"
                },
                "nrLaps": {
                    "type": "integer"
                },
                "position": {
                    "type": "integer"
                },
                "raceNumber": {
                    "type": "integer"
                }
            }
        },
        "instance.ConfigurationJson": {
            "type": "object",
            "properties": {
                "configVersion": {
                    "type": "integer"
                },
                "lanDiscovery": {
                    "type": "integer"
                },
                "maxConnections": {
                    "type": "integer"
                },
                "registerToLobby": {
                    "type": "integer"
                },
                "tcpPort": {
                    "type": "integer"
                },
                "udpPort": {
                    "type": "integer"
                }
            }
        },
        "instance.DriverSettings": {
            "type": "object",
            "properties": {
                "driverCategory": {
                    "type": "integer"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "playerID": {
                    "type": "string"
                },
                "shortName": {
                    "type": "string"
                }
            }
        },
        "instance.DriverState": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "playerID": {
                    "type": "string"
                }
            }
        },
        "instance.EntrySettings": {
            "type": "object",
            "properties": {
                "ballastKg": {
                    "type": "integer"
                },
                "customCar": {
                    "type": "string"
                },
                "defaultGridPosition": {
                    "type": "integer"
                },
                "drivers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.DriverSettings"
                    }
                },
                "forcedCarModel": {
                    "type": "integer"
                },
                "isServerAdmin": {
                    "type": "integer"
                },
                "overrideCarModelForCustomCar": {
                    "type": "integer"
                },
                "overrideDriverInfo": {
                    "type": "integer"
                },
                "raceNumber": {
                    "type": "integer"
                },
                "restrictor": {
                    "type": "integer"
                }
            }
        },
        "instance.EntrylistJson": {
            "type": "object",
            "properties": {
                "configVersion": {
                    "type": "integer"
                },
                "entries": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.EntrySettings"
                    }
                },
                "forceEntryList": {
                    "type": "integer"
                }
            }
        },
        "instance.EventJson": {
            "type": "object",
            "properties": {
                "ambientTemp": {
                    "type": "integer"
                },
                "cloudLevel": {
                    "type": "number"
                },
                "configVersion": {
                    "type": "integer"
                },
                "isFixedConditionQualification": {
                    "type": "integer"
                },
                "postQualySeconds": {
                    "type": "integer"
                },
                "postRaceSeconds": {
                    "type": "integer"
                },
                "preRaceWaitingTimeSeconds": {
                    "type": "integer"
                },
                "rain": {
                    "type": "number"
                },
                "sessionOverTimeSeconds": {
                    "type": "integer"
                },
                "sessions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/instance.SessionSettings"
                    }
                },
                "simracerWeatherConditions": {
                    "type": "integer"
                },
                "track": {
                    "type": "string"
                },
                "trackTemp": {
                    "type": "integer"
                },
                "weatherRandomness": {
                    "type": "integer"
                }
            }
        },
        "instance.EventRulesJson": {
            "type": "object",
            "properties": {
                "configVersion": {
                    "type": "integer"
                },
                "driverStintTimeSec": {
                    "type": "integer"
                },
                "isMandatoryPitstopRefuellingRequired": {
                    "type": "boolean"
                },
                "isMandatoryPitstopSwapDriverRequired": {
                    "type": "boolean"
                },
                "isMandatoryPitstopTyreChangeRequired": {
                    "type": "boolean"
                },
                "isRefuellingAllowedInRace": {
                    "type": "boolean"
                },
                "isRefuellingTimeFixed": {
                    "type": "boolean"
                },
                "mandatoryPitstopCount": {
                    "type": "integer"
                },
                "maxDriversCount": {
                    "type": "integer"
                },
                "maxTotalDrivingTime": {
                    "type": "integer"
                },
                "pitWindowLengthSec": {
                    "type": "integer"
                },
                "qualifyStandingType": {
                    "type": "integer"
                },
                "tyreSetCount": {
                    "type": "integer"
                }
            }
        },
        "instance.LapState": {
            "type": "object",
            "properties": {
                "carID": {
                    "type": "integer"
                },
                "driverIndex": {
                    "type": "integer"
                },
                "flags": {
                    "type": "integer"
                },
                "fuel": {
                    "type": "integer"
                },
                "hasCut": {
                    "type": "boolean"
                },
                "inLap": {
                    "type": "boolean"
                },
                "lapTimeMS": {
                    "type": "integer"
                },
                "outLap": {
                    "type": "boolean"
                },
                "s1": {
                    "type": "string"
                },
                "s2": {
                    "type": "string"
                },
                "s3": {
                    "type": "string"
                },
                "sessionOver": {
                    "type": "boolean"
                },
                "timestampMS": {
                    "type": "integer"
                }
            }
        },
        "instance.LiveState": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/instance.CarState"
                    }
                },
                "nrClients": {
                    "type": "integer"
                },
                "serverState": {
                    "type": "string"
                },
                "sessionPhase": {
                    "type": "string"
                },
                "sessionRemaining": {
                    "type": "integer"
                },
                "sessionType": {
                    "type": "string"
                },
                "track": {
                    "type": "string"
                }
            }
        },
        "instance.SessionSettings": {
            "type": "object",
            "properties": {
                "dayOfWeekend": {
                    "type": "integer"
                },
                "hourOfDay": {
                    "type": "integer"
                },
                "sessionDurationMinutes": {
                    "type": "integer"
                },
                "sessionType": {
                    "type": "string"
                },
                "timeMultiplier": {
                    "type": "integer"
                }
            }
        },
        "instance.SettingsJson": {
            "type": "object",
            "properties": {
                "adminPassword": {
                    "type": "string"
                },
                "allowAutoDQ": {
                    "type": "integer"
                },
                "carGroup": {
                    "type": "string"
                },
                "centralEntryListPath": {
                    "type": "string"
                },
                "configVersion": {
                    "type": "integer"
                },
                "dumpEntryList": {
                    "type": "integer"
                },
                "dumpLeaderboards": {
                    "type": "integer"
                },
                "formationLapType": {
                    "type": "integer"
                },
                "ignorePrematureDisconnects": {
                    "type": "integer"
                },
                "isRaceLocked": {
                    "type": "integer"
                },
                "maxCarSlots": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "racecraftRatingRequirement": {
                    "type": "integer"
                },
                "randomizeTrackWhenEmpty": {
                    "type": "integer"
                },
                "safetyRatingRequirement": {
                    "type": "integer"
                },
                "serverName": {
                    "type": "string"
                },
                "shortFormationLap": {
                    "type": "integer"
                },
                "spectatorPassword": {
                    "type": "string"
                },
                "trackMedalsRequirement": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "ACCWeb API documentation",
	Description:      "Accweb api documentation",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}