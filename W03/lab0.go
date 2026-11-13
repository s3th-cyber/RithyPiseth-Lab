package main

import (
	"crypto/sha3"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func Sha3_256(input1 string , input2 string)  {
	h1 := []byte(input1)
	hh1 := []byte(input2)
	// compute sha1 hashed 
	h2 := sha3.Sum256(h1)
	hh2 := sha3.Sum256(hh1)
	// convert from byted to hex string
	fmt.Println("Hash (SHA3-256) : ")
	h3 := hex.EncodeToString(h2[:])
	hh3:= hex.EncodeToString(hh2[:])
	fmt.Println("\tOutput A : ",h3)
	fmt.Println("\tOutput B : ",hh3)
	
	if(h3==hh3){
		fmt.Println("\t => Matched!")
	}else{
		fmt.Println("\t => No Match!")
	}
}

func Sha512(input1 string , input2 string)  {
	h1 := []byte(input1)
	hh1 := []byte(input2)
	// compute sha1 hashed 
	h2 := sha512.Sum512(h1)
	hh2 := sha512.Sum512(hh1)
	// convert from byted to hex string
	fmt.Println("Hash (SHA512) : ")
	h3 := hex.EncodeToString(h2[:])
	hh3:= hex.EncodeToString(hh2[:])
	fmt.Println("\tOutput A : ",h3)
	fmt.Println("\tOutput B : ",hh3)
	
	if(h3==hh3){
		fmt.Println("\t => Matched!")
	}else{
		fmt.Println("\t => No Match!")
	}
}


func Sha256(input1 string , input2 string)  {
	h1 := []byte(input1)
	hh1 := []byte(input2)
	// compute sha1 hashed 
	h2 := sha256.Sum256(h1)
	hh2 := sha256.Sum256(hh1)
	// convert from byted to hex string
	fmt.Println("Hash (SHA256) : ")
	h3 := hex.EncodeToString(h2[:])
	hh3:= hex.EncodeToString(hh2[:])
	fmt.Println("\tOutput A : ",h3)
	fmt.Println("\tOutput B : ",hh3)
	
	if(h3==hh3){
		fmt.Println("\t => Matched!")
	}else{
		fmt.Println("\t => No Match!")
	}

}

func Sha1(input1 string , input2 string)  {
	h1 := []byte(input1)
	hh1 := []byte(input2)
	// compute sha1 hashed 
	h2 := sha1.Sum(h1)
	hh2 := sha1.Sum(hh1)
	// convert from byted to hex string
	fmt.Println("Hash (SHA1) : ")
	h3 := hex.EncodeToString(h2[:])
	hh3:= hex.EncodeToString(hh2[:])
	fmt.Println("\tOutput A : ",h3)
	fmt.Println("\tOutput B : ",hh3)
	
	if(h3==hh3){
		fmt.Println("\t => Matched!")
	}else{
		fmt.Println("\t => No Match!")
	}
}

func Md5(input1 string , input2 string)  {
	// convert text to bytes
	h1 := []byte(input1)
	hh1 := []byte(input2)
	// compute md5 hashed 
	h2 := md5.Sum(h1)
	hh2 := md5.Sum(hh1)
	// convert from byted to hex string
	fmt.Println("Hash (MD5) : ")
	h3 := hex.EncodeToString(h2[:])
	hh3:= hex.EncodeToString(hh2[:])
	fmt.Println("\tOutput A : ",h3)
	fmt.Println("\tOutput B : ",hh3)
	
	if(h3==hh3){
		fmt.Println("\t => Matched!")
	}else{
		fmt.Println("\t => No Match!")
	}


}


func main(){


	var i1,i2 string


	fmt.Print("Please Input Value 1 : ")
	fmt.Scan(&i1)
	fmt.Print("Please Input value 2 : ")
	fmt.Scan(&i2)
	 
	Md5(i1,i2)
	Sha1(i1,i2)
	Sha256(i1,i2)
	Sha512(i1,i2)
	Sha3_256(i1,i2)

}
