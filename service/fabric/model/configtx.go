package model

type AnchorPeer struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

type Organization struct {
	AnchorPeers []AnchorPeer `yaml:"AnchorPeers"`
}

type ConfigTx struct {
	Organizations []Organization `yaml:"Organizations"`
}
