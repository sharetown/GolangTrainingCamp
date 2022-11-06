package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("%d %b %x %q\n", i, i, i, i)
	}
	//%q      单引号围绕的字符字面值，由Go语法安全地转义

	/*输出：
	60 111100 3c '<'
	61 111101 3d '='
	62 111110 3e '>'
	63 111111 3f '?'
	64 1000000 40 '@'
	65 1000001 41 'A'
	66 1000010 42 'B'
	67 1000011 43 'C'
	68 1000100 44 'D'
	69 1000101 45 'E'
	70 1000110 46 'F'
	71 1000111 47 'G'
	72 1001000 48 'H'
	73 1001001 49 'I'
	74 1001010 4a 'J'
	75 1001011 4b 'K'
	76 1001100 4c 'L'
	77 1001101 4d 'M'
	78 1001110 4e 'N'
	79 1001111 4f 'O'
	80 1010000 50 'P'
	81 1010001 51 'Q'
	82 1010010 52 'R'
	83 1010011 53 'S'
	84 1010100 54 'T'
	85 1010101 55 'U'
	86 1010110 56 'V'
	87 1010111 57 'W'
	88 1011000 58 'X'
	89 1011001 59 'Y'
	90 1011010 5a 'Z'
	91 1011011 5b '['
	92 1011100 5c '\\'
	93 1011101 5d ']'
	94 1011110 5e '^'
	95 1011111 5f '_'
	96 1100000 60 '`'
	97 1100001 61 'a'
	98 1100010 62 'b'
	99 1100011 63 'c'
	100 1100100 64 'd'
	101 1100101 65 'e'
	102 1100110 66 'f'
	103 1100111 67 'g'
	104 1101000 68 'h'
	105 1101001 69 'i'
	106 1101010 6a 'j'
	107 1101011 6b 'k'
	108 1101100 6c 'l'
	109 1101101 6d 'm'
	110 1101110 6e 'n'
	111 1101111 6f 'o'
	112 1110000 70 'p'
	113 1110001 71 'q'
	114 1110010 72 'r'
	115 1110011 73 's'
	116 1110100 74 't'
	117 1110101 75 'u'
	118 1110110 76 'v'
	119 1110111 77 'w'
	120 1111000 78 'x'
	121 1111001 79 'y'
	*/
}
