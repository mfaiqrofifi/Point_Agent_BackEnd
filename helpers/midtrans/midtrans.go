package midtrans

import (
	"strconv"

	"github.com/veritrans/go-midtrans"
)

func GetPembayaranUrl(toko string, email string, total int, id int) (string, error) {
	midclient := midtrans.NewClient()
	midclient.ServerKey = "SB-Mid-server-tEG8D8pdK86RV-O09gTajtwl"
	midclient.ClientKey = "SB-Mid-client-LPnXdCTu57kZymCS"
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: email,
			FName: toko,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(id),
			GrossAmt: int64(total),
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
}
