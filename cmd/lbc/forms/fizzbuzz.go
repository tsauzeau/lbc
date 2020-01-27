package forms

// FizzbuzzForm for the fizzbuzz controller
// multiples of Int1 are replaced by Str1
// multiples of Int2 are replaced by Str2
// multiples of Int1 and Int2 are replaced by Str1Str2
// in the Limit number
type FizzbuzzForm struct {
	Int1    int    `form:"int1" json:"int1" url:"int1" binding:"required"`
	Int2    int    `form:"int2" json:"int2" url:"int2" binding:"required"`
	Limit   int    `form:"limit" json:"limit" url:"limit" binding:"required"`
	String1 string `form:"string1" json:"string1" url:"string1" binding:"required"`
	String2 string `form:"string2" json:"string2" url:"string2" binding:"required"`
}
