package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BranchType string
const (
	CSE BranchType = "CSE"
	DCS BranchType = "DCS"
	ECE BranchType = "ECE"
	DEC BranchType = "DEC"
	MNC BranchType = "MNC"
	EE BranchType = "EE"
	PH BranchType = "PHY"
	CV BranchType = "CVL"
	MS BranchType = "MS"
)

type GenderType string
const (
	Male GenderType = "Male"
	Female GenderType = "Female"
	Others GenderType = "Others"
)

type Student struct {
	ID				primitive.ObjectID			`bson:"_id,omitempty" json:"_id"`
	Name			string						`bson:"name" json:"name" binding:"required"`
	Email			string						`bson:"email" json:"email" binding:"required"`
	Branch			BranchType					`bson:"branch" json:"branch" binding:"required,oneof=CSE DCS ECE DEC MNC EE PHY CVL MS"`
	RollNumber		string						`bson:"roll_number" json:"roll_number" binding:"required"`
	Gender			GenderType					`bson:"gender" json:"gender" binding:"required,oneof=Male Female Others"`
	CourseCode		string						`bson:"course_code" json:"course_code" binding:"required"`
}
