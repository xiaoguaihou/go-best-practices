package core

// interface definition of dependent service

type SendSmsReponse struct {
	Fee     int    `json:"fee"`
	Sid     string `json:"sid"`
	Result  int    `json:"result"`
	Errmsg  string `json:"errmsg"`
	Success bool   `json:"success"`
	Ext     string `json:"ext"`
}
