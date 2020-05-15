package models

/*Relacion establece la relación de un usuario con otro*/
type Relacion struct {
	UserID         string `bson:"usuarioid" json:"usuarioId"`
	UserRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
