module example

go 1.25.8

replace github.com/sikalabsx/sikalabs-encrypted-go => ../..

require (
	github.com/sikalabs/slu v0.101.0
	github.com/sikalabsx/sikalabs-encrypted-go v0.0.0-00010101000000-000000000000
)

require (
	github.com/sikalabs/sikalabs-crypt-go v0.0.0-20251105174759-2ff0fb3fb552 // indirect
	golang.org/x/crypto v0.51.0 // indirect
)
