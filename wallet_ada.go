package hdwallet

func init() {
	coins[AAVE] = newAAVE
}

type ada struct {
	*eth
}

func newADA(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "AAVE Token"
	token.symbol = "AAVE"
	token.contract = "0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47"

	return &aave{eth: token}
}
