/*
 * Copyright (c) 2018
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

// Command line argument printer.
package main
import(
  "fmt"
  "os"
  "strconv"
)

func main(){
  // Version 1 :
  // var s, sep string
  // sep = " "
  // for i := 0 ; i < len(os.Args) ; i++{
  //   s += sep  + os.Args[i]
  // }
  // fmt.Println(s)

  // Version 2:
  // s,sep := ""," "
  // for _, arg := range  os.Args[1:]{
  //   s+= sep + arg;
  // }
  // fmt.Println(s)

  //Version 3:
  //fmt.Println(strings.Join(os.Args[1:]," "))

  //Version 4 : Most efficient for slices. Adds extra parentheses.
  //fmt.Println(os.Args[1:])

  // Version 5 : Args and index printer
  s,sep,term := "", " -> ","\n"
  for index,arg := range os.Args[0:]{
    s+= strconv.Itoa(index) + sep + arg + term;
  }
  fmt.Println(s)
}
