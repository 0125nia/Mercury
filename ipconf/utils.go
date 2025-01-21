package ipconf

func ipConfResp() Response {
	return Response{
		Message: "ok",
		Code:    0,
		Data:    struct{}{}, // todo
	}
}
