package main

//Importowanie pakietów.
import (
    "flag"
    "fmt"
    "golang.org/x/crypto/sha3"
    "os"

)

//Logo SHA-ng.
var (
 logoSHA1024 = `
 SSSSSS\  HH\   HH\  AAAAAA\              Version 1.0b7              
SS  __SS\ HH |  HH |AA  __AA\                                   
SS /  \__|HH |  HH |AA /  AA |        nnnnnnn\   gggggg\  
\SSSSSS\  HHHHHHHH |AAAAAAAA |======\ nn  __nn\ gg  __gg\ 
 \____SS\ HH  __HH |AA  __AA |\______|nn |  nn |gg /  gg |
SS\   SS |HH |  HH |AA |  AA |        nn |  nn |gg |  gg |
\SSSSSS  |HH |  HH |AA |  AA |        nn |  nn |\ggggggg |
 \______/ \__|  \__|\__|  \__|        \__|  \__| \____gg |
                                                gg\   gg |
╒────> Site:                                    \gggggg  |
╘─>https://github.com/lukaszwojcikdev/sha-ng.git \______/ `

    //Deklarowanie zmiennych
    fileName   string
    sha1024    bool
    sha2048    bool
    sha4096    bool
    sha8192    bool
    sha16384   bool
    sha32768   bool
    sha65536   bool
)
//Funkcja główna wyświetlająca informację o flagach/opcjach programu.
func main() {
    flag.StringVar(&fileName, "f", "", "[nazwa pliku]")
    flag.BoolVar(&sha1024, "sha1024", false, "[generuj skrót SHA-1024]")
    flag.BoolVar(&sha2048, "sha2048", false, "[generuj skrót SHA-2048]")
    flag.BoolVar(&sha4096, "sha4096", false, "[generuj skrót SHA-4096]")
    flag.BoolVar(&sha8192, "sha8192", false, "[generuj skrót SHA-8192]")
    flag.BoolVar(&sha16384, "sha16384", false, "[generuj skrót SHA-16384]")
    flag.BoolVar(&sha32768, "sha32768", false, "[generuj skrót SHA-32768]")
    flag.BoolVar(&sha65536, "sha65536", false, "[generuj skrót SHA-65536]")

	//Definiuje funkcję 'Usage' dla flagi: '-h'.
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "%s\n", logoSHA1024)
        fmt.Fprintf(os.Stderr, "\n")
        fmt.Fprintf(os.Stderr, "╒────> Generates SHA hashes 1024-65536 bit  \n├────> for the file using the SHA-3 Shake256 algorithm.\n")
		fmt.Fprintf(os.Stderr, "│\n")
        fmt.Fprintf(os.Stderr, "├──> Use: %s -f [file_name] [option]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "│\n")
        fmt.Fprintf(os.Stderr, "╘──> Opcje: \n")
        printFlags()
    }

	//Przetwarza argumenty wiersza poleceń i przypisuje do nich wartości flag.
    flag.Parse()

	//Wywołanie funkcji 'Usage' wyświetlającej informację o flagach.
    if fileName == "" {
        flag.Usage()
        os.Exit(1)
    }

    buf := []byte(fileName)
	
	//1024-bitowy skrót (128 bajtów).
    if sha1024 {
        h := make([]byte, 128)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-1024: %x\n", h)
    }

	//2048-bitowy skrót (256 bajtów).
    if sha2048 {
        h := make([]byte, 256)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-2048: %x\n", h)
    }

	//4096-bitowy skrót (512 bajtów).
    if sha4096 {
        h := make([]byte, 512)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-4096: %x\n", h)
    }

	//8192-bitowy skrót (1024 bajtów).
    if sha8192 {
        h := make([]byte, 1024)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-8192: %x\n", h)
    }

	//16384-bitowy skrót (2048 bajtów).
    if sha16384 {
        h := make([]byte, 2048)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-16384: %x\n", h)
    }

	//32768-bitowy skrót (4096 bajtów).
    if sha32768 {
        h := make([]byte, 4096)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-32768: %x\n", h)
    }

	//65536-bitowy skrót (8192 bajtów).
    if sha65536 {
        h := make([]byte, 8192)
        sha3.ShakeSum256(h, buf)
        fmt.Printf("SHA-65536: %x\n", h)
    }
}
//Wyświetlanie wszystkich zdefiniowanych flag i wywołanie 
//funkcji zwrotnej dla każdej z nich.
func printFlags() {
    flag.VisitAll(func(f *flag.Flag) {
       fmt.Fprintf(os.Stderr, "  -%s %s\n", f.Name, f.Usage)
    })
}
