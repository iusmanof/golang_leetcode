// govulncheck ./...
//
// Result
// PS C:\Users\H\GolandProjectsmulti_module_workspaces\vuln-tutorial> govulncheck ./
// === Symbol Results ===
//
// Vulnerability #1: GO-2026-5018
// Invoking pathological RSA/DSA parameters may cause DoS in
// golang.org/x/crypto/ssh
// More info: https://pkg.go.dev/vuln/GO-2026-5018
// Module: golang.org/x/crypto
// Found in: golang.org/x/crypto@v0.48.0
// Fixed in: golang.org/x/crypto@v0.52.0
// Example traces found:
// #1: main.go:12:34: vuln.main calls ssh.ParseRawPrivateKey
//
// Your code is affected by 1 vulnerability from 1 module.
// This scan also found 11 vulnerabilities in packages you import and 5
// vulnerabilities in modules you require, but your code doesn't appear to call
// these vulnerabilities.
// Use '-show verbose' for more details.

// module vuln.tutorial
//
// go 1.27.1
//
// //require golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
package main

func main() {
	// Пытаемся распарсить некорректный приватный ключ.
	// В этой версии метод ParseRawPrivateKey содержит критическую уязвимость!
	//_, err := ssh.ParseRawPrivateKey([]byte("invalid_key"))
	//if err != nil {
	//	fmt.Println("Error:", err)
	//}
}
