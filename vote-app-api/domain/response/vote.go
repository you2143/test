package response

type VoteResult struct {
	Title     string `json:"title"`
	VoteCount uint64 `json:"vote_count"`
}

type VoteResultSummary struct {
	VoteResults []VoteResult `json:"vote_results"`
}

type GetVoteIsEnd struct {
	IsEnd bool `json:"is_end"`
}
