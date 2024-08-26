# Learning Go

![Go Logo](https://go.dev/images/go-logo-white.svg)

| References                                                                                           |
| ---------------------------------------------------------------------------------------------------- |
| 👉🏻 [Go by Example](https://gobyexample.com/)                                                         |
| 👉🏻 [Validator](https://github.com/go-playground/validator)                                           |
| 👉🏻 [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin) |

## 📦 `go_by_example` Package

This is a collection of Go scripts each corresponding to a topic, most of which is copied from the **Go by Example** website, with a small modification from me.

## 📦 `validator` Package

This package contains example scripts that demonstrate how to use the `validator` package by `go-playground`. `validator` implements value validations for structs
and individual fields based on tags. The scripts are copied from the examples from the validator's GitHub repository, with a small modification from me.

## 📦 `records_api` Package

> **⚠️ Still Work in Progress**

Based on the **Tutorial: Developing a RESTful API with Go and Gin**, this is a simple RESTful API that uses the Gin framework to provide access to store selling vintage recordings on vinyl. Extra features are added to the original tutorial, such as validation, using a database, and more.

## 🚀 Running the scripts

To run any of the scripts, uncomment the corresponding function call in the `main.go`, and run like this:

```bash
go run .
```
