package firebase

type firebaseSetting struct {
	Type                        string `json:"type"  env:"TYPE"`
	Project_id                  string `json:"project_id"env:"PROJECT_ID"`
	Private_key_id              string `json:"private_key_id" env:"PRIVATE_KEY_ID"`
	Private_key                 string `json:"private_key" env:"PRIVATE_KEY2"`
	Client_email                string `json:"client_email" env:"CLIENT_EMAIL"`
	Client_id                   string `json:"client_id" env:"CLIENT_ID"`
	Auth_uri                    string `json:"auth_uri" env:"AUTH_URI"`
	Token_uri                   string `json:"token_uri" env:"TOKEN_URI"`
	Auth_provider_x509_cert_url string `json:"auth_provider_x509_cert_url" env:"AUTH_PROVIDER_X509_CERT_URL"`
	Client_x509_cert_url        string `json:"client_x509_cert_url" env:"CLIENT_X509_CERT_URL"`
}

var setting firebaseSetting

func init() {
	setting = firebaseSetting{
		Type:                        "service_account",
		Project_id:                  "pupgo-e03ef",
		Private_key_id:              "b8e7eff967143fd022f8722eabf4521227a3a2fb",
		Private_key:                 "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCphvRKHFPRezpR\n+cs86Er3CehJh4Qw+XinsOK477MvHiP8hAcU4/fJI/158/ZM6FBmIWFQmLsH9ZKL\n5zzylV6S+IJM1ySPBx0ZY9XcBoU/qMZ22gqLflJryfc+KSM+GVzT7sODz+g3+lGs\no8CY6XUZZYi1e8M12+TRFKMLNoGRUhIWQmb548R/VJ/IjeebLlj0mTfNSG9b60BQ\n5fLSFZgMidCh7LZmOceuJkDKxomqIgVPDySinzApBmSEhtIuwVfkkLow5fwBw2RB\ndNPUoWZRfWQOTZ6vR9A8IM0/CGWqMvSsrjxrTJTg75z5e2/8ePpq+NPr1cVnridB\nVocofqPrAgMBAAECggEABpeCkLEo1b53DB8PduiDDnHI8ytaIPAb2IbdwgMTUnPA\nK0oOaszaS2ERvlrnLdEF0GOOho9agPfcr8eB6cc3MORDrSetAEnXS95O+anrygfa\nbeo1uqq/vieusiU1/Y+qooMz6UwPDzaHJU6ZpyzaM0vVjA/qO983WaVepRUzLEqQ\ntnBrWcZRlpNgTZW0wTPKWVh2gbZl2lrP3ytYD4PWmEWYzQ9EfdZ9X7qBGvNpTbU5\n4DLeP7kkmuOV+H/X3b+bGi4qPhtUwpzGEwS4LUo6NvkvF4g6qm5xNpL4FzSBtWZV\nFwE86tmPgJ/Q/qnBae+9zVp9EKp1wf8cvpmCaopP7QKBgQDlTtooVEE6UsKneDpR\ncdtgmgl3VMQPHwqw5C4k1WJdMUIzEXJIXp2FoecU++Av+/m34E9+wqOSBuz8yKpT\nlkhjv+9PUYF423B7y+DO/VpzTlvwY5af9mXkZIG+X3IhThOzyOCnVM9dWcpky6hk\nV74fFFcgvR2VtjZLr2sbmC4LrwKBgQC9QrFlHeYJ6ubAM9mml0x3lsTKW4Daco4+\naVJjFOExOph0UKX22RrF8jHRVhHKkeOcyw315i5aTo5FTMzmgreG4Setwn+1c/0O\nhHLQ6FUtZ3B/NP2wT91+5eZq0QGUOFF4U8krhk5eLmdth2+AcZE5imSwUnrrPWDz\nqKJcy8EOhQKBgBEd3bL89xF/dABEPTYvJY0Ecohxz2DBlG632anuM7V5I9PkDX/R\ncDi1aSZr1sQ+LtnG6KgxpzwTQEVuRSiQIz8u6JInJ176Il9bTKCm0MWip8I97NRr\n9BckWXr4bPCHf3kAGaTj88aoGS+E9EDpO8veHNLYywiN3Wew5HHe3jEJAoGBAIGN\nLps17SqqxguZMoqLlMdjyA9wtXJS6jWkqMW0HDYzkvpD3mdr05zeHeXa0dDlLnY8\ntLC3QNGGZnLfkH7VycUlSKDzu7G81ONkHKgN8Yhj/yjEYEeZU9gyjRhfO4J3TqVc\nTR9jDy/++cOK6vN0SeGmbygXPZbbbAs+Su3Ud7F1AoGBAL/FFB8WCHekIBxHuyab\nz9LrpbpeEer1LcP2CR+skCmouU+KdGGv1GjYPtaZsJU9gwLOokIq+pzLSPIiKWl1\nzHeKVmj5BNKRWM4aiDt3stdVro0GtQxeYqn1xp/ViPAcRaxLbUDRta+/hARtZSRe\nD3ZBxaaRlpAw5qP91aI6h4AQ\n-----END PRIVATE KEY-----\n",
		Client_email:                "firebase-adminsdk-sy8tp@pupgo-e03ef.iam.gserviceaccount.com",
		Client_id:                   "103353080089807288854",
		Auth_uri:                    "https://accounts.google.com/o/oauth2/auth",
		Token_uri:                   "https://oauth2.googleapis.com/token",
		Auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
		Client_x509_cert_url:        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-sy8tp%40pupgo-e03ef.iam.gserviceaccount.com",
	}
}
