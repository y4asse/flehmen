// Code generated by ent, DO NOT EDIT.

package ent

import (
	"flehmen-api/ent/mbti"
	"flehmen-api/ent/schema"
	"flehmen-api/ent/sukipi"
	"flehmen-api/ent/tweet"
	"flehmen-api/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	mbtiFields := schema.Mbti{}.Fields()
	_ = mbtiFields
	// mbtiDescCreatedAt is the schema descriptor for created_at field.
	mbtiDescCreatedAt := mbtiFields[1].Descriptor()
	// mbti.DefaultCreatedAt holds the default value on creation for the created_at field.
	mbti.DefaultCreatedAt = mbtiDescCreatedAt.Default.(func() time.Time)
	sukipiFields := schema.Sukipi{}.Fields()
	_ = sukipiFields
	// sukipiDescCreatedAt is the schema descriptor for created_at field.
	sukipiDescCreatedAt := sukipiFields[5].Descriptor()
	// sukipi.DefaultCreatedAt holds the default value on creation for the created_at field.
	sukipi.DefaultCreatedAt = sukipiDescCreatedAt.Default.(func() time.Time)
	tweetFields := schema.Tweet{}.Fields()
	_ = tweetFields
	// tweetDescCreatedAt is the schema descriptor for created_at field.
	tweetDescCreatedAt := tweetFields[3].Descriptor()
	// tweet.DefaultCreatedAt holds the default value on creation for the created_at field.
	tweet.DefaultCreatedAt = tweetDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}