// User - info required for login
/*
type User struct {
	ID                   primitive.ObjectID  `bson:"_id" json:"_id,omitempty" `
	Name                 string              `bson:"name" json:"name,omitempty"`
	Email                string              `bson:"email" json:"email,omitempty" binding:"required"`
	Password             string              `bson:"password" json:"password,omitempty" binding:"required"`
	Phone                string              `bson:"phone_number" json:"phone_number,omitempty"`
	GithubUsername       string              `bson:"github_username" json:"github_username,omitempty"`
	Linkedin             string              `bson:"linkedin" json:"linkedin,omitempty"`
	IsVerified           bool                `bson:"is_verified" json:"is_verified,false"`
	Date                 primitive.Timestamp `bson:"date" json:"date,omitempty"`
	PasswordResetURLCode string              `bson:"password_reset_url_code" json:"password_reset_url_code,omitempty"`
	VerificationURLCode  string              `bson:"verification_url_code" json: "verification_url_code,omitempty"`
	UserType             string              `bson:"user_type" json: "user_type,omitempty"`
}
*/

const mongoose = require('mongoose');
const { v4: uuidv4 } = require('uuid');

const { Schema } = mongoose;

