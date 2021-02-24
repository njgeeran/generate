package model

type BodyPara struct {
	ModelId 		int					`json:"model_id"`
	ContentType 	BodyParaContentType `json:"content_type"`
	Validators 		Validator 			`json:"validators"`

	JoinModel 		Model
}

type BodyParaContentType int
const (
	_BodyParaContentType = iota
	BodyParaContentType_JSON
	BodyParaContentType_FORMDATA
)