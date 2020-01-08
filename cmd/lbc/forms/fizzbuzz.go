package forms

//FizzbuzzForm ...
type FizzbuzzForm struct {
	Int1    int    `form:"int1" json:"int1" url:"int1" binding:"required"`
	Int2    int    `form:"int2" json:"int2" url:"int2" binding:"required"`
	Limit   int    `form:"limit" json:"limit" url:"limit" binding:"required"`
	String1 string `form:"string1" json:"string1" url:"string1" binding:"required"`
	String2 string `form:"string2" json:"string2" url:"string2" binding:"required"`
}
