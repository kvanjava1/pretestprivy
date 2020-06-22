package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// init variable untuk hasil
	var resProcess bool
	var resValidation bool
	var e error

	// get input dari console menggunakan for supaya bisa ber ulang2
	fmt.Printf("input:\n")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// validasi karakter
		resValidation, e = validationChars(scanner.Text())

		// apakah sesuai dengan karakter yg di ijinkan ?
		// ya >> process || tidak >> pesan nil
		if resValidation == false {
			fmt.Printf("", e)
		} else {
			resProcess = process(scanner.Text())
			fmt.Printf("hasil input '" + scanner.Text() + "' adalah " + strconv.FormatBool(resProcess) + "\n")
		}
	}

}

func process(command string) bool {
	// comparasi string sesuai yg di input untuk mendapatkan hasil
	if command == "[" {
		return false
	} else if command == "[]" {
		return true
	} else if command == "{" {
		return false
	} else if command == "{}" {
		return true
	}

	return false

}

func validationChars(command string) (resValidation bool, e error) {
	// kalau inputan kosong false
	if len(command) == 0 {
		return false, nil
	}

	// validasi menggunakan regex hanya kakrakter >> {}[]
	re := regexp.MustCompile("^[/{/[/}\\] ]*$")
	if re.MatchString(command) {
		return true, nil
	}

	return false, nil
}
