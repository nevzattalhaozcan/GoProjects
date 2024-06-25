package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const usernameFile = "usernames.txt"
const passwordFile = "pass.txt"
const balanceFile = "balance.txt"
const logFile = "log.txt"

var balance float64
var amount float64

var username string
var password string

var appRunningCondition bool = true
var profileRunningCondition bool = true

//-----------------MAIN FUNCTION-------------------------------

func main() {

	for appRunningCondition {
		bankUI()
	}

}

//-----------------UI FUNCTIONS-------------------------------

func bankUI() {

	fmt.Println("---GO BANK---")
	fmt.Println("1. Hesap Olustur")
	fmt.Println("2. Giris Yap")
	fmt.Println("3. Uygulamayi kapat")

	var choice int
	fmt.Print("Tercihiniz: ")
	fmt.Scan(&choice)

	if choice == 1 {

		username, password := createUser()

		saveInfo(username, usernameFile)
		saveInfo(password, passwordFile)

		profileRunningCondition = true
		profileUI()

	} else if choice == 2 {

		login()

	} else if choice == 3 {

		displayGoodbyeMessage()
		appRunningCondition = false

	} else {

		fmt.Println("---------------------")
		fmt.Println("hatali secim")
		fmt.Println("---------------------")

	}
}

func profileUI() {

	for profileRunningCondition {

		fmt.Println("---------------------")
		fmt.Printf(">>>Kullanici: %v<<<\n", username)
		fmt.Println("---------------------")
		fmt.Println("Isleminizi Seciniz:")
		fmt.Println("1. Hesabi kontrol et")
		fmt.Println("2. Para yatir")
		fmt.Println("3. Para cek")
		fmt.Println("4. Onceki islemleri goruntule")
		fmt.Println("5. Cikis yap")

		fmt.Print("Tercihiniz: ")
		var option int
		fmt.Scan(&option)

		switch option {
		case 1:
			displayBalance(username)
		case 2:
			processedAmount := depositMoney()
			log("para yatirma", processedAmount)
		case 3:
			processedAmount := withdrawMoney()
			log("para cekme", processedAmount)
		case 4:
			displayPreviousTransactions()
		case 5:
			fmt.Println("---------------------")
			fmt.Println("Basariyla cikis yapildi")
			fmt.Println("---------------------")
			profileRunningCondition = false
			return
		default:
			fmt.Println("---------------------")
			fmt.Println("hatali secim")
			fmt.Println("---------------------")
		}

	}

}

//-------------BANK TRANSACTIONS-----------------

func usernameCheck(username string) bool {
	var checkResult bool
	username = strings.Trim(username, " ")
	if username == "" {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("kullanici adi bos birakilamaz")
		fmt.Println("---------------------")
	} else if len(username) > 10 {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("kullanici adi 10 karakterden az olmalidir")
		fmt.Println("---------------------")
	} else {
		checkResult = true
	}

	return checkResult
}

func passwordCheck(password string) bool {
	var checkResult bool
	var digitCheck bool
	var letterCheck bool

	for _, char := range password {
		if unicode.IsDigit(char) {
			digitCheck = true
		} else if unicode.IsLetter(char) {
			letterCheck = true
		} else {
			letterCheck = false
			digitCheck = false
		}
	}
	if password == "" {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("sifre bos birakilamaz")
		fmt.Println("---------------------")
	} else if len(password) != 6 {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("sifre 6 karakter olmak zorundadir")
		fmt.Println("---------------------")
	} else if !digitCheck {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("sifre en az bir tane rakam icermelidir")
		fmt.Println("---------------------")
	} else if !letterCheck {
		checkResult = false
		fmt.Println("---------------------")
		fmt.Println("sifre en az bir tane harf icermelidir")
		fmt.Println("---------------------")
	} else {
		checkResult = true
	}

	return checkResult
}

func checkUserExist(username string) bool {
	for _, user := range strings.Split(convertUserListToString(), ",") {
		if username == user {
			return true
		}
	}
	return false
}

func usernamePasswordMatch(username, password string) bool {

	for _, user := range strings.Split(convertUserListToString(), ",") {

		for _, pass := range strings.Split(convertPasswordListToString(), ",") {

			if username == user && password == pass {
				return true
			}

		}

	}

	return false
}

func createUser() (string, string) {
	fmt.Print("Kullanici adinizi belirleyiniz: ")
	fmt.Scan(&username)

	for checkUserExist(username) {
		fmt.Println("---------------------")
		fmt.Println("bu kullanici adi kullanimdadir")
		fmt.Println("---------------------")
		fmt.Print("Kullanici adinizi belirleyiniz: ")
		fmt.Scan(&username)
	}

	for !usernameCheck(username) {
		fmt.Print("Kullanici adinizi belirleyiniz: ")
		fmt.Scan(&username)
	}

	fmt.Print("Alti haneli bir sifre belirleyiniz: ")
	fmt.Scan(&password)

	for !passwordCheck(password) {
		fmt.Print("Alti haneli bir sifre belirleyiniz: ")
		fmt.Scan(&password)
	}
	balance = findUserBalanceAndUpdate(username, 0)
	return username, password
}

func login() {

	fmt.Print("Kullanici adinizi giriniz: ")
	fmt.Scan(&username)

	if !checkUserExist(username) {

		fmt.Println("---------------------")
		fmt.Println("bu kullanici adi mevcut degildir")
		fmt.Println("---------------------")
		profileRunningCondition = false
		return
	}

	fmt.Print("Sifrenizi giriniz: ")
	fmt.Scan(&password)

	if usernamePasswordMatch(username, password) {

		profileRunningCondition = true
		profileUI()

	} else {

		for !usernamePasswordMatch(username, password) {

			fmt.Println("---------------------")
			fmt.Println("yanlis sifre/kullanici adi")
			fmt.Println("---------------------")

			fmt.Print("Kullanici adinizi giriniz: ")
			fmt.Scan(&username)
			fmt.Print("Sifrenizi giriniz: ")
			fmt.Scan(&password)

		}

		profileRunningCondition = true
		profileUI()

	}
}

func findUserBalanceAndUpdate(username string, balance float64) float64 {

userLoop:
	for _, user := range strings.Split(convertUserListToString(), ",") {

		for _, balanceStr := range strings.Split(convertBalanceListToString(), ",") {

			if user == "" {
				break userLoop
			} else if user == username {
				updateBalance(balance)
				break
			} else {
				balanceRecord, _ := strconv.ParseFloat(balanceStr, 64)
				updateBalance(balanceRecord)
				break
			}

		}

	}
	return balance
}

//-------------ACCOUNT TRANSACTIONS-----------------

func depositMoney() float64 {

	fmt.Println("---------------------")
	fmt.Print("Miktari girin: ")
	fmt.Scan(&amount)

	balance += amount

	findUserBalanceAndUpdate(username, balance)

	fmt.Println("---------------------")
	fmt.Printf("Basarili! Yeni hesap miktari: %v\n", balance)
	fmt.Println("---------------------")

	return amount

}

func withdrawMoney() float64 {

	fmt.Println("---------------------")
	fmt.Print("Miktari girin: ")
	fmt.Scan(&amount)

	balance -= amount

	findUserBalanceAndUpdate(username, balance)

	fmt.Println("---------------------")
	fmt.Printf("Basarili! Yeni hesap miktari: %v\n", balance)
	fmt.Println("---------------------")

	return amount

}

func displayBalance(username string) {

	fmt.Println("---------------------")
	fmt.Printf("Hesabiniz: %v\n", findUserBalance(username))
	fmt.Println("---------------------")

}

func displayPreviousTransactions() {

	var logByte, _ = os.ReadFile(logFile)
	var logs string = string(logByte)

	fmt.Println("---------------------")
	fmt.Println(logs)
	fmt.Println("---------------------")

}

//----------------HELPER FUNCTIONS--------------------

func log(process string, amount float64) {

	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	logMessage := fmt.Sprintf("--->KAYIT: %s - %.2f\n", process, amount)

	file.WriteString(logMessage)

}

func saveInfo(info, fileName string) {

	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	record := fmt.Sprintf("%v,", info)

	file.WriteString(record)

}

func displayGoodbyeMessage() {

	fmt.Println("---------------------")
	fmt.Println("Gule gule!")
	fmt.Println("Go Bank'i tercih ettiginiz icin tesekkurler!")
	fmt.Println("---------------------")

}

func convertUserListToString() string {

	userList, _ := os.ReadFile(usernameFile)
	users := string(userList)

	return users

}

func convertPasswordListToString() string {

	passwordList, _ := os.ReadFile(passwordFile)
	passwords := string(passwordList)

	return passwords

}

func convertBalanceListToString() string {

	balanceList, _ := os.ReadFile(balanceFile)
	balances := string(balanceList)

	return balances

}

func updateBalance(balance float64) {

	file, _ := os.OpenFile(balanceFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	balanceRecord := fmt.Sprintf("%v,", balance)

	file.WriteString(balanceRecord)
}

func findUserBalance(username string) float64 {
	var userBalance float64
userLoop:
	for _, user := range strings.Split(convertUserListToString(), ",") {

		for _, balanceStr := range strings.Split(convertBalanceListToString(), ",") {

			if user == "" {
				break
			} else if user == username {
				userBalance, _ = strconv.ParseFloat(balanceStr, 64)
				break userLoop
			}

		}
	}
	return userBalance
}
