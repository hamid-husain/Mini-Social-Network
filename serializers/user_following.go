package serializers

type UserFollowRequest struct {
	UserIDs []uint `json:"user_ids"`
}
