package response

import "awesomeProject111/model/serializer"

type IpResponse struct {
	SuccessData *serializer.IpSuccessData
	FailData    *serializer.IpFailData
}
