package main

import "gorm.io/gorm"

//globally accepted here.
type Employee struct {
	gorm.Model
	EmpName   string
	EmpSalary int `json:"salary"`
	Email     string
}
