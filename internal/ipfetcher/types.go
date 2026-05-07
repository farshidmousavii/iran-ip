package ipfetcher

type ASNResponse struct {
	Data struct {
		Resources struct {
			ASN  []string `json:"asn"`
			IPV4 []string `json:"ipv4"`
			IPV6 []string `json:"ipv6"`
		} `json:"resources"`
	} `json:"data"`
}

type ASNJob struct {
	ASN string
}

type PrefixResponse struct {
	Data struct {
		Prefixes []Prefix `json:"prefixes"`
	} `json:"data"`
}

type Prefix struct {
	Prefix string `json:"prefix"`
}

type PrefixResult struct {
	ASN    string
	Prefix Prefix
	Err    error
}
