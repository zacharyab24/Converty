/* main.go
 * Last updated 20/12/24
 * converty is a cli application used to convert between different number formats
 * The current plan is to be able to convert between decimal, hex, binary and text (utf-8, may add others later)
 * The application will use flags to specify which to convert to, or have a default all option
 */

package main

import "fmt"
import "flag"
import "os"
//import "strings"

func main() {
    
    //Flags
    formatPtr := flag.String("f", "binary", "Input format, e.g. decimal, hex, binary")
    decimalPtr := flag.Bool("d", false, "Convert to decimal")
    hexPtr := flag.Bool("h", false, "Convert to hex")
    binPtr := flag.Bool("b", false, "Convert to binary")

	flag.Parse()

    //Check if the input format is valid
    if flag.Arg(0) == "" {
        fmt.Println("No input string was given")
        os.Exit(80)
    }
    inputString := flag.Arg(0)


    if *formatPtr == "binary" {
        fmt.Println("Converting from binary")
        for i := range inputString {
            // 48 is the utf8 value for 0 and 49 is the value for 1
            // Since the input will always be a utf8 string, we need to check if its equal to this value not "0" or "1"
            if inputString[i] != 48 && inputString[i] != 49 {
                fmt.Println("An invalid input string was given")
                os.Exit(80)
            }
        }
        fmt.Println(inputString)
    } else if *formatPtr == "hex" {
       fmt.Println("Converting from hex") 
    } else if *formatPtr == "decimal" {
        fmt.Println("Converting from decimal") 
    } else {
        fmt.Println("An invalid input was specified")
        os.Exit(80)
    }

    fmt.Print(decimalPtr, hexPtr, binPtr)
}
