package request

type GetMasterPerson struct {
	ID uint64 `uri:"id" binding:"required"`
}

type CreateMasterPerson struct {
	Name string `form:"name"`
}

type UpdateMasterPerson struct {
	ID   uint64 `uri:"id" binding:"required"`
	Name string `form:"name"`
}

type DeleteMasterPerson struct {
	ID uint64 `uri:"id" binding:"required"`
}
