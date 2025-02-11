package constants

type dataEndpoints struct {
	GET_IDENTITY     string
	GET_ACCOUNTS     string
	GET_BALANCE      string
	GET_TRANSACTIONS string
}

type metadataEndpoints struct {
	GET_ACCOUNTS string
}

type paymentEndpoints struct {
	GET_BENEFICIARIES  string
	CREATE_BENEFICIARY string
	CREATE_TRANSFER    string
	TRANSFER_AUTOFLOW  string
}

type authEndpoints struct {
	DELINK_USER    string
	EXCHANGE_TOKEN string
}

type operationEndpoints struct {
	OPERATION_STATUS string
}

type dapiEndpoints struct {
	BASE_URL       string
	DATA_URLS      dataEndpoints
	METADATA_URLS  metadataEndpoints
	PAYMENT_URLS   paymentEndpoints
	AUTH_URLS      authEndpoints
	OPERATION_URLS operationEndpoints
}

const DD_URL = "https://dd.dapi.co"

// DAPI_URL is the base var that holds all supported API endpoints
var DAPI_URL = &dapiEndpoints{
	BASE_URL: "https://api.dapi.co/v2",
	DATA_URLS: dataEndpoints{
		GET_IDENTITY:     "/data/identity/get",
		GET_ACCOUNTS:     "/data/accounts/get",
		GET_BALANCE:      "/data/balance/get",
		GET_TRANSACTIONS: "/data/transactions/get",
	},
	METADATA_URLS: metadataEndpoints{
		GET_ACCOUNTS: "/metadata/accounts/get",
	},
	PAYMENT_URLS: paymentEndpoints{
		GET_BENEFICIARIES:  "/payment/beneficiaries/get",
		CREATE_BENEFICIARY: "/payment/beneficiaries/create",
		CREATE_TRANSFER:    "/payment/transfer/create",
		TRANSFER_AUTOFLOW:  "/payment/transfer/autoflow",
	},
	AUTH_URLS: authEndpoints{
		DELINK_USER:    "/users/delinkuser",
		EXCHANGE_TOKEN: "/auth/ExchangeToken",
	},

	OPERATION_URLS: operationEndpoints{
		OPERATION_STATUS: "/operation/status",
	},
}
