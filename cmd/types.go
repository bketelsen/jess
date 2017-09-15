package cmd

type Conf struct {
	SubscriptionID string
	StorageName    string
	ResourceGroup  string
	Location       string
	ShareName      string
	StorageKey     string
	VaultName      string
}

const Config = `storagename: {{.StorageName}}
subscriptionid: {{.SubscriptionID}}
resourcegroup: {{.ResourceGroup}}
location: {{.Location}}
sharename: {{.ShareName}}
storagekey: {{.StorageKey}}
vaultname: {{.VaultName}}
`
