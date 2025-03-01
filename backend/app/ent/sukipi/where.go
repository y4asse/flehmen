// Code generated by ent, DO NOT EDIT.

package sukipi

import (
	"flehmen-api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldName, v))
}

// LikedAt applies equality check predicate on the "liked_at" field. It's identical to LikedAtEQ.
func LikedAt(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldLikedAt, v))
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldWeight, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldHeight, v))
}

// XID applies equality check predicate on the "x_id" field. It's identical to XIDEQ.
func XID(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldXID, v))
}

// Hobby applies equality check predicate on the "hobby" field. It's identical to HobbyEQ.
func Hobby(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldHobby, v))
}

// Birthday applies equality check predicate on the "birthday" field. It's identical to BirthdayEQ.
func Birthday(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldBirthday, v))
}

// ShoesSize applies equality check predicate on the "shoesSize" field. It's identical to ShoesSizeEQ.
func ShoesSize(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldShoesSize, v))
}

// Family applies equality check predicate on the "family" field. It's identical to FamilyEQ.
func Family(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldFamily, v))
}

// NearlyStation applies equality check predicate on the "nearly_station" field. It's identical to NearlyStationEQ.
func NearlyStation(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldNearlyStation, v))
}

// Mbti applies equality check predicate on the "mbti" field. It's identical to MbtiEQ.
func Mbti(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldMbti, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldName, v))
}

// LikedAtEQ applies the EQ predicate on the "liked_at" field.
func LikedAtEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldLikedAt, v))
}

// LikedAtNEQ applies the NEQ predicate on the "liked_at" field.
func LikedAtNEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldLikedAt, v))
}

// LikedAtIn applies the In predicate on the "liked_at" field.
func LikedAtIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldLikedAt, vs...))
}

// LikedAtNotIn applies the NotIn predicate on the "liked_at" field.
func LikedAtNotIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldLikedAt, vs...))
}

// LikedAtGT applies the GT predicate on the "liked_at" field.
func LikedAtGT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldLikedAt, v))
}

// LikedAtGTE applies the GTE predicate on the "liked_at" field.
func LikedAtGTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldLikedAt, v))
}

// LikedAtLT applies the LT predicate on the "liked_at" field.
func LikedAtLT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldLikedAt, v))
}

// LikedAtLTE applies the LTE predicate on the "liked_at" field.
func LikedAtLTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldLikedAt, v))
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldWeight, v))
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldWeight, v))
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldWeight, vs...))
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldWeight, vs...))
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldWeight, v))
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldWeight, v))
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldWeight, v))
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldWeight, v))
}

// WeightIsNil applies the IsNil predicate on the "weight" field.
func WeightIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldWeight))
}

// WeightNotNil applies the NotNil predicate on the "weight" field.
func WeightNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldWeight))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldHeight, v))
}

// HeightIsNil applies the IsNil predicate on the "height" field.
func HeightIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldHeight))
}

// HeightNotNil applies the NotNil predicate on the "height" field.
func HeightNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldHeight))
}

// XIDEQ applies the EQ predicate on the "x_id" field.
func XIDEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldXID, v))
}

// XIDNEQ applies the NEQ predicate on the "x_id" field.
func XIDNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldXID, v))
}

// XIDIn applies the In predicate on the "x_id" field.
func XIDIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldXID, vs...))
}

// XIDNotIn applies the NotIn predicate on the "x_id" field.
func XIDNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldXID, vs...))
}

// XIDGT applies the GT predicate on the "x_id" field.
func XIDGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldXID, v))
}

// XIDGTE applies the GTE predicate on the "x_id" field.
func XIDGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldXID, v))
}

// XIDLT applies the LT predicate on the "x_id" field.
func XIDLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldXID, v))
}

// XIDLTE applies the LTE predicate on the "x_id" field.
func XIDLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldXID, v))
}

// XIDContains applies the Contains predicate on the "x_id" field.
func XIDContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldXID, v))
}

// XIDHasPrefix applies the HasPrefix predicate on the "x_id" field.
func XIDHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldXID, v))
}

// XIDHasSuffix applies the HasSuffix predicate on the "x_id" field.
func XIDHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldXID, v))
}

// XIDIsNil applies the IsNil predicate on the "x_id" field.
func XIDIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldXID))
}

// XIDNotNil applies the NotNil predicate on the "x_id" field.
func XIDNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldXID))
}

// XIDEqualFold applies the EqualFold predicate on the "x_id" field.
func XIDEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldXID, v))
}

// XIDContainsFold applies the ContainsFold predicate on the "x_id" field.
func XIDContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldXID, v))
}

// HobbyEQ applies the EQ predicate on the "hobby" field.
func HobbyEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldHobby, v))
}

// HobbyNEQ applies the NEQ predicate on the "hobby" field.
func HobbyNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldHobby, v))
}

// HobbyIn applies the In predicate on the "hobby" field.
func HobbyIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldHobby, vs...))
}

// HobbyNotIn applies the NotIn predicate on the "hobby" field.
func HobbyNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldHobby, vs...))
}

// HobbyGT applies the GT predicate on the "hobby" field.
func HobbyGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldHobby, v))
}

// HobbyGTE applies the GTE predicate on the "hobby" field.
func HobbyGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldHobby, v))
}

// HobbyLT applies the LT predicate on the "hobby" field.
func HobbyLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldHobby, v))
}

// HobbyLTE applies the LTE predicate on the "hobby" field.
func HobbyLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldHobby, v))
}

// HobbyContains applies the Contains predicate on the "hobby" field.
func HobbyContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldHobby, v))
}

// HobbyHasPrefix applies the HasPrefix predicate on the "hobby" field.
func HobbyHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldHobby, v))
}

// HobbyHasSuffix applies the HasSuffix predicate on the "hobby" field.
func HobbyHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldHobby, v))
}

// HobbyIsNil applies the IsNil predicate on the "hobby" field.
func HobbyIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldHobby))
}

// HobbyNotNil applies the NotNil predicate on the "hobby" field.
func HobbyNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldHobby))
}

// HobbyEqualFold applies the EqualFold predicate on the "hobby" field.
func HobbyEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldHobby, v))
}

// HobbyContainsFold applies the ContainsFold predicate on the "hobby" field.
func HobbyContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldHobby, v))
}

// BirthdayEQ applies the EQ predicate on the "birthday" field.
func BirthdayEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldBirthday, v))
}

// BirthdayNEQ applies the NEQ predicate on the "birthday" field.
func BirthdayNEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldBirthday, v))
}

// BirthdayIn applies the In predicate on the "birthday" field.
func BirthdayIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldBirthday, vs...))
}

// BirthdayNotIn applies the NotIn predicate on the "birthday" field.
func BirthdayNotIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldBirthday, vs...))
}

// BirthdayGT applies the GT predicate on the "birthday" field.
func BirthdayGT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldBirthday, v))
}

// BirthdayGTE applies the GTE predicate on the "birthday" field.
func BirthdayGTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldBirthday, v))
}

// BirthdayLT applies the LT predicate on the "birthday" field.
func BirthdayLT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldBirthday, v))
}

// BirthdayLTE applies the LTE predicate on the "birthday" field.
func BirthdayLTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldBirthday, v))
}

// BirthdayIsNil applies the IsNil predicate on the "birthday" field.
func BirthdayIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldBirthday))
}

// BirthdayNotNil applies the NotNil predicate on the "birthday" field.
func BirthdayNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldBirthday))
}

// ShoesSizeEQ applies the EQ predicate on the "shoesSize" field.
func ShoesSizeEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldShoesSize, v))
}

// ShoesSizeNEQ applies the NEQ predicate on the "shoesSize" field.
func ShoesSizeNEQ(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldShoesSize, v))
}

// ShoesSizeIn applies the In predicate on the "shoesSize" field.
func ShoesSizeIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldShoesSize, vs...))
}

// ShoesSizeNotIn applies the NotIn predicate on the "shoesSize" field.
func ShoesSizeNotIn(vs ...float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldShoesSize, vs...))
}

// ShoesSizeGT applies the GT predicate on the "shoesSize" field.
func ShoesSizeGT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldShoesSize, v))
}

// ShoesSizeGTE applies the GTE predicate on the "shoesSize" field.
func ShoesSizeGTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldShoesSize, v))
}

// ShoesSizeLT applies the LT predicate on the "shoesSize" field.
func ShoesSizeLT(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldShoesSize, v))
}

// ShoesSizeLTE applies the LTE predicate on the "shoesSize" field.
func ShoesSizeLTE(v float64) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldShoesSize, v))
}

// ShoesSizeIsNil applies the IsNil predicate on the "shoesSize" field.
func ShoesSizeIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldShoesSize))
}

// ShoesSizeNotNil applies the NotNil predicate on the "shoesSize" field.
func ShoesSizeNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldShoesSize))
}

// FamilyEQ applies the EQ predicate on the "family" field.
func FamilyEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldFamily, v))
}

// FamilyNEQ applies the NEQ predicate on the "family" field.
func FamilyNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldFamily, v))
}

// FamilyIn applies the In predicate on the "family" field.
func FamilyIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldFamily, vs...))
}

// FamilyNotIn applies the NotIn predicate on the "family" field.
func FamilyNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldFamily, vs...))
}

// FamilyGT applies the GT predicate on the "family" field.
func FamilyGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldFamily, v))
}

// FamilyGTE applies the GTE predicate on the "family" field.
func FamilyGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldFamily, v))
}

// FamilyLT applies the LT predicate on the "family" field.
func FamilyLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldFamily, v))
}

// FamilyLTE applies the LTE predicate on the "family" field.
func FamilyLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldFamily, v))
}

// FamilyContains applies the Contains predicate on the "family" field.
func FamilyContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldFamily, v))
}

// FamilyHasPrefix applies the HasPrefix predicate on the "family" field.
func FamilyHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldFamily, v))
}

// FamilyHasSuffix applies the HasSuffix predicate on the "family" field.
func FamilyHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldFamily, v))
}

// FamilyIsNil applies the IsNil predicate on the "family" field.
func FamilyIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldFamily))
}

// FamilyNotNil applies the NotNil predicate on the "family" field.
func FamilyNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldFamily))
}

// FamilyEqualFold applies the EqualFold predicate on the "family" field.
func FamilyEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldFamily, v))
}

// FamilyContainsFold applies the ContainsFold predicate on the "family" field.
func FamilyContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldFamily, v))
}

// NearlyStationEQ applies the EQ predicate on the "nearly_station" field.
func NearlyStationEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldNearlyStation, v))
}

// NearlyStationNEQ applies the NEQ predicate on the "nearly_station" field.
func NearlyStationNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldNearlyStation, v))
}

// NearlyStationIn applies the In predicate on the "nearly_station" field.
func NearlyStationIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldNearlyStation, vs...))
}

// NearlyStationNotIn applies the NotIn predicate on the "nearly_station" field.
func NearlyStationNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldNearlyStation, vs...))
}

// NearlyStationGT applies the GT predicate on the "nearly_station" field.
func NearlyStationGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldNearlyStation, v))
}

// NearlyStationGTE applies the GTE predicate on the "nearly_station" field.
func NearlyStationGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldNearlyStation, v))
}

// NearlyStationLT applies the LT predicate on the "nearly_station" field.
func NearlyStationLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldNearlyStation, v))
}

// NearlyStationLTE applies the LTE predicate on the "nearly_station" field.
func NearlyStationLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldNearlyStation, v))
}

// NearlyStationContains applies the Contains predicate on the "nearly_station" field.
func NearlyStationContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldNearlyStation, v))
}

// NearlyStationHasPrefix applies the HasPrefix predicate on the "nearly_station" field.
func NearlyStationHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldNearlyStation, v))
}

// NearlyStationHasSuffix applies the HasSuffix predicate on the "nearly_station" field.
func NearlyStationHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldNearlyStation, v))
}

// NearlyStationIsNil applies the IsNil predicate on the "nearly_station" field.
func NearlyStationIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldNearlyStation))
}

// NearlyStationNotNil applies the NotNil predicate on the "nearly_station" field.
func NearlyStationNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldNearlyStation))
}

// NearlyStationEqualFold applies the EqualFold predicate on the "nearly_station" field.
func NearlyStationEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldNearlyStation, v))
}

// NearlyStationContainsFold applies the ContainsFold predicate on the "nearly_station" field.
func NearlyStationContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldNearlyStation, v))
}

// MbtiEQ applies the EQ predicate on the "mbti" field.
func MbtiEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldMbti, v))
}

// MbtiNEQ applies the NEQ predicate on the "mbti" field.
func MbtiNEQ(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldMbti, v))
}

// MbtiIn applies the In predicate on the "mbti" field.
func MbtiIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldMbti, vs...))
}

// MbtiNotIn applies the NotIn predicate on the "mbti" field.
func MbtiNotIn(vs ...string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldMbti, vs...))
}

// MbtiGT applies the GT predicate on the "mbti" field.
func MbtiGT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldMbti, v))
}

// MbtiGTE applies the GTE predicate on the "mbti" field.
func MbtiGTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldMbti, v))
}

// MbtiLT applies the LT predicate on the "mbti" field.
func MbtiLT(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldMbti, v))
}

// MbtiLTE applies the LTE predicate on the "mbti" field.
func MbtiLTE(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldMbti, v))
}

// MbtiContains applies the Contains predicate on the "mbti" field.
func MbtiContains(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContains(FieldMbti, v))
}

// MbtiHasPrefix applies the HasPrefix predicate on the "mbti" field.
func MbtiHasPrefix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasPrefix(FieldMbti, v))
}

// MbtiHasSuffix applies the HasSuffix predicate on the "mbti" field.
func MbtiHasSuffix(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldHasSuffix(FieldMbti, v))
}

// MbtiIsNil applies the IsNil predicate on the "mbti" field.
func MbtiIsNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIsNull(FieldMbti))
}

// MbtiNotNil applies the NotNil predicate on the "mbti" field.
func MbtiNotNil() predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotNull(FieldMbti))
}

// MbtiEqualFold applies the EqualFold predicate on the "mbti" field.
func MbtiEqualFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEqualFold(FieldMbti, v))
}

// MbtiContainsFold applies the ContainsFold predicate on the "mbti" field.
func MbtiContainsFold(v string) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldContainsFold(FieldMbti, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Sukipi {
	return predicate.Sukipi(sql.FieldLTE(FieldCreatedAt, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Sukipi {
	return predicate.Sukipi(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Sukipi {
	return predicate.Sukipi(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sukipi) predicate.Sukipi {
	return predicate.Sukipi(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sukipi) predicate.Sukipi {
	return predicate.Sukipi(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sukipi) predicate.Sukipi {
	return predicate.Sukipi(sql.NotPredicates(p))
}
