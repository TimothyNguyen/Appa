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

const { ObjectID } = require('mongodb');
const mongoose = require('mongoose');
const { v4: uuidv4 } = require('uuid');

const { Schema } = mongoose;

const UserSchema = new Schema({
    _id: {
        type: ObjectID,
        required: true,
    },
    name: {
        type: String,
        required: true,
    },
    email: {
        type: String,
        required: true,
    },
    password: {
        type: String,
        required: true,
    },
    phone_number: {
        type: String,
    },
    github_username: {
        type: String,
    },
    linkedin_link: {
        type: String,
    },
    date: {
        type: Date,
        default: Date.now,
    },
    verified: {
        type: Boolean,
        default: false,
    },
    verification_url_code: {
        type: String,
        default: () => uuidv4(),
    },
    password_reset_url_code: {
        type: String,
        default: '',
    },
    user_type: {
        type: String,
    },
})

const User = mongoose.model('users', UserSchema);
module.exports = User;