package models

import "gorm.io/gorm"

type Crunchy_prereq struct {
    gorm.Model
    kubectl string `json:"kubectl" gorm:"text;not null;default:null"`
    git string `json:"git" gorm:"text;not null;default:null"'
}