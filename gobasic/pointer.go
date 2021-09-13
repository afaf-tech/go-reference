package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pass by reference
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address1 Address = Address{City: "Subang", Province: "Jawa Barat", Country: "Indonesia"}
	var address2 *Address = &address1 // becomes pointer to address1

	address2.City = "Bandung"

	// address1 will not be changed due to not using * in reassignment process
	address2 = &Address{City: "Gresik", Province: "Jawa Timur", Country: "Indonesia"}
	fmt.Println("address1", address1)
	fmt.Println("address2", address2)

	var address3 Address = Address{City: "Malang", Province: "Jawa Timur", Country: "Indonesia"}
	var address4 *Address = &address3 // becomes pointer to address3
	// address3 will be changed due to using * in reassignment process
	*address4 = Address{City: "Bojonegoro", Province: "Jawa Timur", Country: "Indonesia"}
	fmt.Println("address3", address3)
	fmt.Println("address4", address4)

	// keyword new to make empty pointer
	var address5 *Address = new(Address)
	address5.City = "Palembang"
	fmt.Println("address5", address5)

	var alamat = Address{City: "Denpasar", Province: "Bali", Country: ""}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println("alamat", alamat)
}
