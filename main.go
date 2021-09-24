/*
 * Author(s):
 * - Mattia Cabrini <dev@mattiacabrini.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */
package main

import (
	"log"
	"net/http"
	"os"
)

func pu(s int) {
	print("Usage: \ndir-server <port> <directory path>\n\nExample:\ndir-server 3000 /home/nick/awesomeProject")
	os.Exit(s)
}

// Arg 1 = Port
// Arg 2 = Directory
func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		pu(0)
	}

	fs := http.FileServer(http.Dir(os.Args[2]))
	http.Handle("/", fs)

	log.Println("Listening on :" + os.Args[1] + "...")
	err := http.ListenAndServe("0.0.0.0:"+os.Args[1], nil)

	if err != nil {
		log.Fatal(err)
	}
}
