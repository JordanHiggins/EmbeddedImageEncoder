Embedded Image Encoder
======================

This application converts bitmap images in standard formats (BMP, GIF, JPG, PNG) into the 12-bit format used by the Nokia 6100 LCD screen. This application was created as a complement to embedded code created as part of the EECS 3100 Microsystems program of Fall 2014.

This application can convert images to 3 formats:

 - "Bitmap" - Padded 12-bit color. This format is comparable to the "Native" format below, but it is somewhat simpler to understand. Obsoleted by "Native" format.
 - "Native" - Packed 12-bit color. This format is native to the Nokia 6100 LCD screen, and as such is the most CPU and memory efficient to use for color images. Intended for color images.
 - "Template" - 1-bit color. This format allows for the foreground and background colors to be specified programmatically, and is highly compact. Intended for fonts or icons.

Output is formatted to embed the image in program code, intended for use with the IAR ARM assembler.
