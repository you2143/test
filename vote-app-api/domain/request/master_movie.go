package request

type GetMasterMovie struct {
	ID uint64 `uri:"id" binding:"required"`
}

type CreateMasterMovie struct {
	Title string `form:"title"`
}

type UpdateMasterMovie struct {
	ID    uint64 `uri:"id" binding:"required"`
	Title string `form:"title"`
}

type DeleteMasterMovie struct {
	ID uint64 `uri:"id" binding:"required"`
}
