// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/serlip06",
            "email": "714220023@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chartitem": {
            "get": {
                "description": "Mengambil semua data cartitem.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chartitem"
                ],
                "summary": "Get All Data cartitem.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartItem"
                        }
                    }
                }
            }
        },
        "/chartitem/{id}": {
            "get": {
                "description": "Ambil per ID data cartitem.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chartitem"
                ],
                "summary": "Get By ID Data Cartitem.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/customer": {
            "get": {
                "description": "Mengambil semua data customer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get All Data Customer.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Customer"
                        }
                    }
                }
            }
        },
        "/customer/{id}": {
            "get": {
                "description": "Ambil per ID data customer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Get By ID Data Customer.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Customer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data customer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Delete data customer.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/deletechartitem/{id}": {
            "delete": {
                "description": "Hapus data cartitem.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chartitem"
                ],
                "summary": "Delete data cartitem.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/deleteproduk/{id}": {
            "delete": {
                "description": "Hapus data produk.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produk"
                ],
                "summary": "Delete data produk.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insert": {
            "post": {
                "description": "Input data customer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Insert data customer.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Customer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insertchartitem": {
            "post": {
                "description": "Input data cartitem.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chartitem"
                ],
                "summary": "Insert data cartitem.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqCartItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insertproduk": {
            "post": {
                "description": "Input data produk.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produk"
                ],
                "summary": "Insert data produk.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqProduk"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Produk"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/produk": {
            "get": {
                "description": "Mengambil semua data Produk.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produk"
                ],
                "summary": "Get All Data Produk.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Produk"
                        }
                    }
                }
            }
        },
        "/produk/{id}": {
            "get": {
                "description": "Ambil per ID data produk.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produk"
                ],
                "summary": "Get By ID Data Produk.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Produk"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data customer.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Update data customer.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Customer"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/updatechartitem/{id}": {
            "put": {
                "description": "Ubah data cartitem.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chartitem"
                ],
                "summary": "Update data cartitem.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqCartItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.CartItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/updateproduk/{id}": {
            "put": {
                "description": "Ubah data produk.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produk"
                ],
                "summary": "Update data produk.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqProduk"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Produk"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CartItem": {
            "type": "object",
            "properties": {
                "_id": {
                    "description": "ID unik untuk item keranjang",
                    "type": "string",
                    "example": "1234567"
                },
                "gambar": {
                    "description": "Gambar produk",
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "description": "Harga produk pada saat dimasukkan ke keranjang",
                    "type": "integer",
                    "example": 5000
                },
                "id_produk": {
                    "description": "Referensi ke ID Produk",
                    "type": "string",
                    "example": "1234567"
                },
                "id_user": {
                    "type": "string",
                    "example": "1234567"
                },
                "is_selected": {
                    "description": "Tambahkan flag ini",
                    "type": "boolean",
                    "example": true
                },
                "nama_produk": {
                    "description": "nama untuk produknya",
                    "type": "string",
                    "example": "ikan bakar"
                },
                "quantity": {
                    "description": "Jumlah produk dalam keranjang",
                    "type": "integer",
                    "example": 1
                },
                "sub_total": {
                    "description": "Total harga (Harga * Quantity)",
                    "type": "integer",
                    "example": 2000
                }
            }
        },
        "controller.Customer": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "1234567"
                },
                "alamat": {
                    "type": "string",
                    "example": "jl.sarijadi"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "nasi goreng dengan telor dan daging"
                },
                "email": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Xaviera_89@gmail.com",
                        "Putri_90@gmail.com"
                    ]
                },
                "gambar": {
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "type": "integer",
                    "example": 15000
                },
                "nama": {
                    "type": "string",
                    "example": "xavieraa putri"
                },
                "nama_produk": {
                    "type": "string",
                    "example": "Nasi Goreng"
                },
                "phone_number": {
                    "type": "string",
                    "example": "085798654096"
                },
                "stok": {
                    "type": "string",
                    "example": "10"
                }
            }
        },
        "controller.Produk": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "1234567"
                },
                "created_at": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "minuman teh manis yang menyegarkan"
                },
                "gambar": {
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "type": "integer",
                    "example": 10000
                },
                "kategori": {
                    "type": "string",
                    "example": "makanan"
                },
                "nama_produk": {
                    "type": "string",
                    "example": "nama makanan/minuman"
                },
                "stok": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "controller.ReqCartItem": {
            "type": "object",
            "properties": {
                "gambar": {
                    "description": "Gambar produk",
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "description": "Harga produk pada saat dimasukkan ke keranjang",
                    "type": "integer",
                    "example": 5000
                },
                "id_produk": {
                    "description": "Referensi ke ID Produk",
                    "type": "string",
                    "example": "1234567"
                },
                "id_user": {
                    "type": "string",
                    "example": "1234567"
                },
                "is_selected": {
                    "description": "Tambahkan flag ini",
                    "type": "boolean",
                    "example": true
                },
                "nama_produk": {
                    "description": "nama untuk produknya",
                    "type": "string",
                    "example": "ikan bakar"
                },
                "quantity": {
                    "description": "Jumlah produk dalam keranjang",
                    "type": "integer",
                    "example": 1
                },
                "sub_total": {
                    "description": "Total harga (Harga * Quantity)",
                    "type": "integer",
                    "example": 2000
                }
            }
        },
        "controller.ReqCustomer": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string",
                    "example": "jl.sarijadi"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "nasi goreng dengan telor dan daging"
                },
                "email": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Xaviera_89@gmail.com",
                        "Putri_90@gmail.com"
                    ]
                },
                "gambar": {
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "type": "integer",
                    "example": 15000
                },
                "nama": {
                    "type": "string",
                    "example": "Tes swager"
                },
                "nama_produk": {
                    "type": "string",
                    "example": "Nasi Goreng"
                },
                "phone_number": {
                    "type": "string",
                    "example": "085798654096"
                },
                "stok": {
                    "type": "string",
                    "example": "10"
                }
            }
        },
        "controller.ReqProduk": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deskripsi": {
                    "type": "string",
                    "example": "minuman teh manis yang menyegarkan"
                },
                "gambar": {
                    "type": "string",
                    "example": "https://i.pinimg.com/564x/94/82/ab/9482ab2e248d249e7daa7fd6924c8d3b.jpg"
                },
                "harga": {
                    "type": "integer",
                    "example": 10000
                },
                "kategori": {
                    "type": "string",
                    "example": "makanan"
                },
                "nama_produk": {
                    "type": "string",
                    "example": "test swagger"
                },
                "stok": {
                    "type": "integer",
                    "example": 5
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "ats-714220023-serlipariela-38bba14820aa.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAGGER ULBI",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
