package resources

type FetchByIdUri struct {
	Id uint `uri:"id" binding:"required"`
}
