package request

type CreateVote struct {
	MovieID  uint `form:"movie_id"`
	PersonID uint `form:"person_id"`
}
