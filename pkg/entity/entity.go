package entity

type User struct {
	CNPJ  uint64 `json:"cnpj"`
	Senha string `json:"password"`
}

type NfeProc struct {
	NFe NFe `xml:"NFe"`
}

type NFe struct {
	InfNFe InfNFe `xml:"infNFe"`
}

type InfNFe struct {
	Emit  Emit  `xml:"emit"`
	Dest  Dest  `xml:"dest"`
	Det   Det   `xml:"det"`
	Total Total `xml:"total"`
}

type Emit struct {
	CNPJ string `xml:"CNPJ"`
	// XNome     string    `xml:"xNome"`
	// XFant     string    `xml:"xFant"`
	// EnderEmit EnderEmit `xml:"enderEmit"`
	// IE        string    `xml:"IE"`
	// CRT       string    `xml:"CRT"`
}

type EnderEmit struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	XMun    string `xml:"xMun"`
	UF      string `xml:"UF"`
	CEP     string `xml:"CEP"`
	CPais   string `xml:"cPais"`
	XPais   string `xml:"xPais"`
	Fone    string `xml:"fone"`
}

type Dest struct {
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	EnderDest EnderDest `xml:"enderDest"`
	// IndIEDest string    `xml:"indIEDest"`
	// IE        string    `xml:"IE"`
}

type EnderDest struct {
	XLgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	XCpl    string `xml:"xCpl"`
	XBairro string `xml:"xBairro"`
	CMun    string `xml:"cMun"`
	// XMun    string `xml:"xMun"`
	// UF      string `xml:"UF"`
	CEP string `xml:"CEP"`
	// CPais   string `xml:"cPais"`
	// XPais   string `xml:"xPais"`
	// Fone    string `xml:"fone"`
}

type Det struct {
	Prod []*Prod `xml:"prod"`
}

type Prod struct {
	CProd  string  `xml:"cProd"`
	CEAN   string  `xml:"cEAN"`
	XProd  string  `xml:"xProd"`
	UCom   string  `xml:"uCom"`
	QCom   float64 `xml:"qCom"`
	VUnCom float64 `xml:"vUnCom"`
	// QTrib   float64 `xml:"qTrib"`
	// VUnTrib float64 `xml:"vUnTrib"`
	VFrete float64 `xml:"vFrete"`
	VSeg   float64 `xml:"vSeg"`
	VDesc  float64 `xml:"vDesc"`
	VOutro float64 `xml:"vOutro"`
	VProd  float64 `xml:"vProd"`
}

type Total struct {
	ICMSTot ICMSTot `xml:"ICMSTot"`
}

type ICMSTot struct {
	VProd  float64 `xml:"vProd"`
	VFrete float64 `xml:"vFrete"`
	VSeg   float64 `xml:"vSeg"`
	VDesc  float64 `xml:"vDesc"`
	VOutro float64 `xml:"vOutro"`
	// VAdicional float64 `xml:"vAdicional"`
	// VMargem    float64 `xml:"vMargem"`
}
