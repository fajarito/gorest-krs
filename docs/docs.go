// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://bkkbn.go.id",
        "contact": {
            "name": "Direktorat Teknologi Informasi dan Data",
            "url": "https://bkkbn.go.id",
            "email": "dittifdok@bkkbn.go.id"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/krsbykab": {
            "get": {
                "description": "Get Aggregate Keluarga Beresiko Stunting By Kecamatan",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Aggregate Keluarga Beresiko Stunting By Kecamatan"
                ],
                "summary": "Get Aggregate Keluarga Beresiko Stunting By Kecamatan",
                "operationId": "get-krs-by-kec",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province Code",
                        "name": "kdprov",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kabupaten Code",
                        "name": "kdkab",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/krk_kab.Krk"
                        }
                    }
                }
            }
        },
        "/krsbykec": {
            "get": {
                "description": "Get Aggregate Keluarga Beresiko Stunting By Kelurahan",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Aggregate Keluarga Beresiko Stunting By Kelurahan"
                ],
                "summary": "Get Aggregate Keluarga Beresiko Stunting By Kelurahan",
                "operationId": "get-krs-by-kelurahan",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province Code",
                        "name": "kdprov",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kabupaten Code",
                        "name": "kdkab",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Kecamatan Code",
                        "name": "kdkec",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/krs_kec.Krs"
                        }
                    }
                }
            }
        },
        "/krsbynas": {
            "get": {
                "description": "Get Aggregate Keluarga Beresiko Stunting By Province",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Aggregate Keluarga Beresiko Stunting By Province"
                ],
                "summary": "Get Aggregate Keluarga Beresiko Stunting By Province",
                "operationId": "get-krs-by-prov",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/krs_nas.KrsNas"
                        }
                    }
                }
            }
        },
        "/krsbyprov": {
            "get": {
                "description": "Get Aggregate Keluarga Beresiko Stunting By Kabupaten",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Get Aggregate Keluarga Beresiko Stunting By Kabupaten"
                ],
                "summary": "Get Aggregate Keluarga Beresiko Stunting By Kabupaten",
                "operationId": "get-krs-by-kab",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province Code",
                        "name": "kdprov",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/krs_prov.KrsProv"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "krk_kab.Krk": {
            "type": "object",
            "properties": {
                "idKabupaten": {
                    "type": "integer"
                },
                "idKecamatan": {
                    "type": "integer"
                },
                "idProvinsi": {
                    "type": "integer"
                },
                "jumlahAirTidakLayak": {
                    "type": "integer"
                },
                "jumlahBaduta": {
                    "type": "integer"
                },
                "jumlahBalita": {
                    "type": "integer"
                },
                "jumlahBukanPesertaKbModern": {
                    "type": "integer"
                },
                "jumlahJambanTidakLayak": {
                    "type": "integer"
                },
                "jumlahKeluarga": {
                    "type": "integer"
                },
                "jumlahKeluargaBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahKeluargaSasaran": {
                    "type": "integer"
                },
                "jumlahKeluargaTidakBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahPus": {
                    "type": "integer"
                },
                "jumlahPusHamil": {
                    "type": "integer"
                },
                "kodeKabupaten": {
                    "type": "string"
                },
                "kodeKecamatan": {
                    "type": "string"
                },
                "kodeProvinsi": {
                    "type": "string"
                },
                "namaKabupaten": {
                    "type": "string"
                },
                "namaKecamatan": {
                    "type": "string"
                },
                "namaProvinsi": {
                    "type": "string"
                },
                "peringkatKesejahteraanDiatas4": {
                    "type": "integer"
                },
                "prioritas1": {
                    "type": "integer"
                },
                "prioritas2": {
                    "type": "integer"
                },
                "prioritas3": {
                    "type": "integer"
                },
                "prioritas4": {
                    "type": "integer"
                }
            }
        },
        "krs_kec.Krs": {
            "type": "object",
            "properties": {
                "idKabupaten": {
                    "type": "integer"
                },
                "idKecamatan": {
                    "type": "integer"
                },
                "idKelurahan": {
                    "type": "integer"
                },
                "idProvinsi": {
                    "type": "integer"
                },
                "jumlahAirTidakLayak": {
                    "type": "integer"
                },
                "jumlahBaduta": {
                    "type": "integer"
                },
                "jumlahBalita": {
                    "type": "integer"
                },
                "jumlahBukanPesertaKbModern": {
                    "type": "integer"
                },
                "jumlahJambanTidakLayak": {
                    "type": "integer"
                },
                "jumlahKeluarga": {
                    "type": "integer"
                },
                "jumlahKeluargaBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahKeluargaSasaran": {
                    "type": "integer"
                },
                "jumlahKeluargaTidakBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahPus": {
                    "type": "integer"
                },
                "jumlahPusHamil": {
                    "type": "integer"
                },
                "kodeKabupaten": {
                    "type": "string"
                },
                "kodeKecamatan": {
                    "type": "string"
                },
                "kodeKelurahan": {
                    "type": "string"
                },
                "kodeProvinsi": {
                    "type": "string"
                },
                "namaKabupaten": {
                    "type": "string"
                },
                "namaKecamatan": {
                    "type": "string"
                },
                "namaKelurahan": {
                    "type": "string"
                },
                "namaProvinsi": {
                    "type": "string"
                },
                "peringkatKesejahteraanDiatas4": {
                    "type": "integer"
                },
                "prioritas1": {
                    "type": "integer"
                },
                "prioritas2": {
                    "type": "integer"
                },
                "prioritas3": {
                    "type": "integer"
                },
                "prioritas4": {
                    "type": "integer"
                }
            }
        },
        "krs_nas.KrsNas": {
            "type": "object",
            "properties": {
                "idProvinsi": {
                    "type": "integer"
                },
                "jumlahAirTidakLayak": {
                    "type": "integer"
                },
                "jumlahBaduta": {
                    "type": "integer"
                },
                "jumlahBalita": {
                    "type": "integer"
                },
                "jumlahBukanPesertaKbModern": {
                    "type": "integer"
                },
                "jumlahJambanTidakLayak": {
                    "type": "integer"
                },
                "jumlahKeluarga": {
                    "type": "integer"
                },
                "jumlahKeluargaBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahKeluargaSasaran": {
                    "type": "integer"
                },
                "jumlahKeluargaTidakBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahPus": {
                    "type": "integer"
                },
                "jumlahPusHamil": {
                    "type": "integer"
                },
                "kodeProvinsi": {
                    "type": "string"
                },
                "namaProvinsi": {
                    "type": "string"
                },
                "peringkatKesejahteraanDiatas4": {
                    "type": "integer"
                },
                "prioritas1": {
                    "type": "integer"
                },
                "prioritas2": {
                    "type": "integer"
                },
                "prioritas3": {
                    "type": "integer"
                },
                "prioritas4": {
                    "type": "integer"
                }
            }
        },
        "krs_prov.KrsProv": {
            "type": "object",
            "properties": {
                "idKabupaten": {
                    "type": "integer"
                },
                "idProvinsi": {
                    "type": "integer"
                },
                "jumlahAirTidakLayak": {
                    "type": "integer"
                },
                "jumlahBaduta": {
                    "type": "integer"
                },
                "jumlahBalita": {
                    "type": "integer"
                },
                "jumlahBukanPesertaKbModern": {
                    "type": "integer"
                },
                "jumlahJambanTidakLayak": {
                    "type": "integer"
                },
                "jumlahKeluarga": {
                    "type": "integer"
                },
                "jumlahKeluargaBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahKeluargaSasaran": {
                    "type": "integer"
                },
                "jumlahKeluargaTidakBeresikoStunting": {
                    "type": "integer"
                },
                "jumlahPus": {
                    "type": "integer"
                },
                "jumlahPusHamil": {
                    "type": "integer"
                },
                "kodeKabupaten": {
                    "type": "string"
                },
                "kodeProvinsi": {
                    "type": "string"
                },
                "namaKabupaten": {
                    "type": "string"
                },
                "namaProvinsi": {
                    "type": "string"
                },
                "peringkatKesejahteraanDiatas4": {
                    "type": "integer"
                },
                "prioritas1": {
                    "type": "integer"
                },
                "prioritas2": {
                    "type": "integer"
                },
                "prioritas3": {
                    "type": "integer"
                },
                "prioritas4": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.1",
	Host:             "gorest-2022-krs-ds-bkkbn-2022-gorest-krs.apps.ocp-dev.bkkbn.go.id",
	BasePath:         "/v1/api",
	Schemes:          []string{},
	Title:            "BKKBN Digital Service - Rekapitulasi Keluarga Beresiko Stunting",
	Description:      "Digital Service BKKBN for Integration",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
