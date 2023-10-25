package entity

type User struct {
	CNPJ  uint64 `json:"cnpj"`
	Senha string `json:"password"`
}

type NFe struct {
	Emit       Emit   `xml:"NFe>infNFe>emit"`
	Dest       Dest   `xml:"NFe>infNFe>dest"`
	Prod       []Prod `xml:"NFe>infNFe>det"`
	VAdicional float64
	VMargem    float64
}

type Emit struct {
	CNPJ  string `xml:"CNPJ"`
	XNome string `xml:"xNome"`
}

type Dest struct {
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	Email     string    `xml:"email"`
	EnderDest EnderDest `xml:"enderDest"`
}

type EnderDest struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	CEP     string `xml:"CEP"`
	Fone    string `xml:"fone"`
}

type Prod struct {
	CProd  string  `xml:"cProd"`
	CEAN   string  `xml:"cEAN"`
	XProd  string  `xml:"xProd"`
	UCom   string  `xml:"uCom"`
	QCom   float64 `xml:"qCom"`
	VUnCom float64 `xml:"vUnCom"`
	VCusto float64
	VPreco float64
}
