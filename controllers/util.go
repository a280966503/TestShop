package controllers

func JsonResponse(flag bool,message string) map[string]interface{}  {
	resp := make(map[string]interface{})

	resp["flag"]=flag
	resp["message"]=message

	return resp
}

