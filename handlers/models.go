package handlers

type UrlCreationRequest struct {
  LongUrl string `json:"long_url" binding:"required"`
  UserId  string `json:"user_id"  binding:"required"`
}