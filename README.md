## Printing double-sided A6 booklets using A4 printer

This tool calculates the page number re-ordering for double-sided booklet printing.

Compiling:

You need the [Go programming language](https://golang.org/) installed. Set it up proper, esp. ```GOPATH```.

```
go get github.com/ERnsTL/a6-booklet-on-a4
go install github.com/ERnsTL/a6-booklet-on-a4
```

You should now have a ```bin/a6-booklet-on-a4``` executable.

Input: A PDF file with A6-sized pages (your booklet pages) with a page number with a multiple of 8 (fill up with blank pages at the end if necessary). Example PDFs included for testing and getting used to the program and the printing process.

Example:

```
bin/a6-booklet-on-a4 -in test16.pdf -out printthis.pdf -pages 16
```

Output:

```
[...]
print order: [2 15 6 11 16 1 12 5 4 13 8 9 14 3 10 7]
pdftk A=test16.pdf cat A2 A15 A6 A11 A16 A1 A12 A5 A4 A13 A8 A9 A14 A3 A10 A7 output printthis.pdf
```

Now copy-paste or otherwise execute the last line. You need ```pdftk``` installed for re-ordering the pages. On Ubuntu, Debian, Mint etc. this is done using ```sudo apt install pdftk```.

The result will be a PDF with re-ordered pages, which you can print.

If you have a duplex printer, this is easy, otherwise if you have a simple one-sided printer then follow these steps:

1. Run ```evince printthis.pdf``` to open it in the PDF viewer.
1. Press Ctrl+P to go to the print dialog.
1. Select printing 4 pages on one - in this case 4 A6-sized pages for each A4 page -, select ordering left to right and top to bottom, limit pages (the A4 ones) to the odd pages, set scaling to enlarge to the print area. Print.
1. Take the paper stack out, turn it 180 degrees horizontally and re-insert it onto the top of your paper tray.
1. Go to the print dialog again. Set all the settings as before, but limit pages to even pages and where you select your printer and number of copies, enable the option to reverse ordering (of the A4 pages) since first printed page is now at the bottom. Print.
1. Take the paper stack out and cut it horizontally to horizontal A5 sheets.
1. Put the stack which has the first page on it (should be the top one) onto the other one.
1. Fold horizontally to get a A6 booklet, nicely printed on both sides.
1. Staple or bind as desired. Long-arm staplers are nice.

Feedback and improvements welcome.

License: GPLv3+
