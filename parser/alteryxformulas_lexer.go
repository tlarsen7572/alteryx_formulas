// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 75, 680,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48,
	3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61,
	3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3,
	64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65,
	3, 65, 3, 65, 3, 66, 3, 66, 5, 66, 534, 10, 66, 3, 67, 6, 67, 537, 10,
	67, 13, 67, 14, 67, 538, 3, 68, 7, 68, 542, 10, 68, 12, 68, 14, 68, 545,
	11, 68, 3, 68, 3, 68, 6, 68, 549, 10, 68, 13, 68, 14, 68, 550, 3, 69, 3,
	69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 6, 71, 563,
	10, 71, 13, 71, 14, 71, 564, 3, 71, 3, 71, 3, 72, 3, 72, 7, 72, 571, 10,
	72, 12, 72, 14, 72, 574, 11, 72, 3, 72, 3, 72, 3, 73, 3, 73, 7, 73, 580,
	10, 73, 12, 73, 14, 73, 583, 11, 73, 3, 73, 3, 73, 3, 74, 6, 74, 588, 10,
	74, 13, 74, 14, 74, 589, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77,
	3, 77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3,
	82, 3, 83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87,
	3, 88, 3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 3,
	93, 3, 93, 3, 94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97, 3, 98,
	3, 98, 3, 99, 3, 99, 3, 100, 3, 100, 3, 101, 3, 101, 3, 102, 3, 102, 3,
	102, 3, 102, 3, 102, 3, 102, 3, 102, 3, 102, 3, 102, 3, 102, 3, 102, 3,
	103, 3, 103, 3, 103, 3, 103, 3, 103, 3, 103, 3, 103, 3, 103, 3, 103, 3,
	103, 3, 103, 3, 104, 3, 104, 3, 104, 3, 104, 3, 104, 3, 105, 3, 105, 3,
	105, 3, 105, 3, 105, 3, 105, 2, 2, 106, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7,
	13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31,
	17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49,
	26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67,
	35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85,
	44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103,
	53, 105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119,
	61, 121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135,
	69, 137, 70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 2, 151,
	2, 153, 2, 155, 2, 157, 2, 159, 2, 161, 2, 163, 2, 165, 2, 167, 2, 169,
	2, 171, 2, 173, 2, 175, 2, 177, 2, 179, 2, 181, 2, 183, 2, 185, 2, 187,
	2, 189, 2, 191, 2, 193, 2, 195, 2, 197, 2, 199, 2, 201, 2, 203, 2, 205,
	2, 207, 2, 209, 2, 3, 2, 34, 3, 2, 98, 98, 3, 2, 95, 95, 3, 2, 41, 41,
	3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 67, 99, 99, 4, 2,
	68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2,
	71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2,
	74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2,
	77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2,
	80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2,
	83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2,
	86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2,
	89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2,
	92, 92, 124, 124, 3, 2, 50, 59, 2, 656, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3,
	2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2,
	113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2,
	2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127,
	3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2,
	2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3,
	2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 3,
	211, 3, 2, 2, 2, 5, 213, 3, 2, 2, 2, 7, 215, 3, 2, 2, 2, 9, 217, 3, 2,
	2, 2, 11, 219, 3, 2, 2, 2, 13, 221, 3, 2, 2, 2, 15, 223, 3, 2, 2, 2, 17,
	225, 3, 2, 2, 2, 19, 227, 3, 2, 2, 2, 21, 230, 3, 2, 2, 2, 23, 232, 3,
	2, 2, 2, 25, 235, 3, 2, 2, 2, 27, 238, 3, 2, 2, 2, 29, 240, 3, 2, 2, 2,
	31, 243, 3, 2, 2, 2, 33, 246, 3, 2, 2, 2, 35, 249, 3, 2, 2, 2, 37, 253,
	3, 2, 2, 2, 39, 258, 3, 2, 2, 2, 41, 263, 3, 2, 2, 2, 43, 268, 3, 2, 2,
	2, 45, 274, 3, 2, 2, 2, 47, 282, 3, 2, 2, 2, 49, 287, 3, 2, 2, 2, 51, 299,
	3, 2, 2, 2, 53, 309, 3, 2, 2, 2, 55, 318, 3, 2, 2, 2, 57, 322, 3, 2, 2,
	2, 59, 327, 3, 2, 2, 2, 61, 338, 3, 2, 2, 2, 63, 347, 3, 2, 2, 2, 65, 356,
	3, 2, 2, 2, 67, 367, 3, 2, 2, 2, 69, 371, 3, 2, 2, 2, 71, 377, 3, 2, 2,
	2, 73, 385, 3, 2, 2, 2, 75, 397, 3, 2, 2, 2, 77, 401, 3, 2, 2, 2, 79, 407,
	3, 2, 2, 2, 81, 411, 3, 2, 2, 2, 83, 418, 3, 2, 2, 2, 85, 422, 3, 2, 2,
	2, 87, 426, 3, 2, 2, 2, 89, 431, 3, 2, 2, 2, 91, 434, 3, 2, 2, 2, 93, 438,
	3, 2, 2, 2, 95, 443, 3, 2, 2, 2, 97, 451, 3, 2, 2, 2, 99, 457, 3, 2, 2,
	2, 101, 461, 3, 2, 2, 2, 103, 466, 3, 2, 2, 2, 105, 471, 3, 2, 2, 2, 107,
	478, 3, 2, 2, 2, 109, 482, 3, 2, 2, 2, 111, 487, 3, 2, 2, 2, 113, 491,
	3, 2, 2, 2, 115, 494, 3, 2, 2, 2, 117, 498, 3, 2, 2, 2, 119, 502, 3, 2,
	2, 2, 121, 505, 3, 2, 2, 2, 123, 508, 3, 2, 2, 2, 125, 513, 3, 2, 2, 2,
	127, 518, 3, 2, 2, 2, 129, 525, 3, 2, 2, 2, 131, 533, 3, 2, 2, 2, 133,
	536, 3, 2, 2, 2, 135, 543, 3, 2, 2, 2, 137, 552, 3, 2, 2, 2, 139, 556,
	3, 2, 2, 2, 141, 560, 3, 2, 2, 2, 143, 568, 3, 2, 2, 2, 145, 577, 3, 2,
	2, 2, 147, 587, 3, 2, 2, 2, 149, 593, 3, 2, 2, 2, 151, 595, 3, 2, 2, 2,
	153, 597, 3, 2, 2, 2, 155, 599, 3, 2, 2, 2, 157, 601, 3, 2, 2, 2, 159,
	603, 3, 2, 2, 2, 161, 605, 3, 2, 2, 2, 163, 607, 3, 2, 2, 2, 165, 609,
	3, 2, 2, 2, 167, 611, 3, 2, 2, 2, 169, 613, 3, 2, 2, 2, 171, 615, 3, 2,
	2, 2, 173, 617, 3, 2, 2, 2, 175, 619, 3, 2, 2, 2, 177, 621, 3, 2, 2, 2,
	179, 623, 3, 2, 2, 2, 181, 625, 3, 2, 2, 2, 183, 627, 3, 2, 2, 2, 185,
	629, 3, 2, 2, 2, 187, 631, 3, 2, 2, 2, 189, 633, 3, 2, 2, 2, 191, 635,
	3, 2, 2, 2, 193, 637, 3, 2, 2, 2, 195, 639, 3, 2, 2, 2, 197, 641, 3, 2,
	2, 2, 199, 643, 3, 2, 2, 2, 201, 645, 3, 2, 2, 2, 203, 647, 3, 2, 2, 2,
	205, 658, 3, 2, 2, 2, 207, 669, 3, 2, 2, 2, 209, 674, 3, 2, 2, 2, 211,
	212, 7, 42, 2, 2, 212, 4, 3, 2, 2, 2, 213, 214, 7, 43, 2, 2, 214, 6, 3,
	2, 2, 2, 215, 216, 7, 44, 2, 2, 216, 8, 3, 2, 2, 2, 217, 218, 7, 49, 2,
	2, 218, 10, 3, 2, 2, 2, 219, 220, 7, 45, 2, 2, 220, 12, 3, 2, 2, 2, 221,
	222, 7, 47, 2, 2, 222, 14, 3, 2, 2, 2, 223, 224, 7, 63, 2, 2, 224, 16,
	3, 2, 2, 2, 225, 226, 7, 64, 2, 2, 226, 18, 3, 2, 2, 2, 227, 228, 7, 64,
	2, 2, 228, 229, 7, 63, 2, 2, 229, 20, 3, 2, 2, 2, 230, 231, 7, 62, 2, 2,
	231, 22, 3, 2, 2, 2, 232, 233, 7, 62, 2, 2, 233, 234, 7, 63, 2, 2, 234,
	24, 3, 2, 2, 2, 235, 236, 7, 35, 2, 2, 236, 237, 7, 63, 2, 2, 237, 26,
	3, 2, 2, 2, 238, 239, 7, 46, 2, 2, 239, 28, 3, 2, 2, 2, 240, 241, 7, 40,
	2, 2, 241, 242, 7, 40, 2, 2, 242, 30, 3, 2, 2, 2, 243, 244, 7, 126, 2,
	2, 244, 245, 7, 126, 2, 2, 245, 32, 3, 2, 2, 2, 246, 247, 7, 42, 2, 2,
	247, 248, 7, 43, 2, 2, 248, 34, 3, 2, 2, 2, 249, 250, 5, 149, 75, 2, 250,
	251, 5, 151, 76, 2, 251, 252, 5, 185, 93, 2, 252, 36, 3, 2, 2, 2, 253,
	254, 5, 149, 75, 2, 254, 255, 5, 153, 77, 2, 255, 256, 5, 177, 89, 2, 256,
	257, 5, 185, 93, 2, 257, 38, 3, 2, 2, 2, 258, 259, 5, 149, 75, 2, 259,
	260, 5, 185, 93, 2, 260, 261, 5, 165, 83, 2, 261, 262, 5, 175, 88, 2, 262,
	40, 3, 2, 2, 2, 263, 264, 5, 149, 75, 2, 264, 265, 5, 187, 94, 2, 265,
	266, 5, 149, 75, 2, 266, 267, 5, 175, 88, 2, 267, 42, 3, 2, 2, 2, 268,
	269, 5, 149, 75, 2, 269, 270, 5, 187, 94, 2, 270, 271, 5, 149, 75, 2, 271,
	272, 5, 175, 88, 2, 272, 273, 7, 52, 2, 2, 273, 44, 3, 2, 2, 2, 274, 275,
	5, 149, 75, 2, 275, 276, 5, 191, 96, 2, 276, 277, 5, 157, 79, 2, 277, 278,
	5, 183, 92, 2, 278, 279, 5, 149, 75, 2, 279, 280, 5, 161, 81, 2, 280, 281,
	5, 157, 79, 2, 281, 46, 3, 2, 2, 2, 282, 283, 5, 153, 77, 2, 283, 284,
	5, 157, 79, 2, 284, 285, 5, 165, 83, 2, 285, 286, 5, 171, 86, 2, 286, 48,
	3, 2, 2, 2, 287, 288, 5, 153, 77, 2, 288, 289, 5, 163, 82, 2, 289, 290,
	5, 149, 75, 2, 290, 291, 5, 183, 92, 2, 291, 292, 5, 159, 80, 2, 292, 293,
	5, 183, 92, 2, 293, 294, 5, 177, 89, 2, 294, 295, 5, 173, 87, 2, 295, 296,
	5, 165, 83, 2, 296, 297, 5, 175, 88, 2, 297, 298, 5, 187, 94, 2, 298, 50,
	3, 2, 2, 2, 299, 300, 5, 153, 77, 2, 300, 301, 5, 163, 82, 2, 301, 302,
	5, 149, 75, 2, 302, 303, 5, 183, 92, 2, 303, 304, 5, 187, 94, 2, 304, 305,
	5, 177, 89, 2, 305, 306, 5, 165, 83, 2, 306, 307, 5, 175, 88, 2, 307, 308,
	5, 187, 94, 2, 308, 52, 3, 2, 2, 2, 309, 310, 5, 153, 77, 2, 310, 311,
	5, 177, 89, 2, 311, 312, 5, 175, 88, 2, 312, 313, 5, 187, 94, 2, 313, 314,
	5, 149, 75, 2, 314, 315, 5, 165, 83, 2, 315, 316, 5, 175, 88, 2, 316, 317,
	5, 185, 93, 2, 317, 54, 3, 2, 2, 2, 318, 319, 5, 153, 77, 2, 319, 320,
	5, 177, 89, 2, 320, 321, 5, 185, 93, 2, 321, 56, 3, 2, 2, 2, 322, 323,
	5, 153, 77, 2, 323, 324, 5, 177, 89, 2, 324, 325, 5, 185, 93, 2, 325, 326,
	5, 163, 82, 2, 326, 58, 3, 2, 2, 2, 327, 328, 5, 153, 77, 2, 328, 329,
	5, 177, 89, 2, 329, 330, 5, 189, 95, 2, 330, 331, 5, 175, 88, 2, 331, 332,
	5, 187, 94, 2, 332, 333, 5, 193, 97, 2, 333, 334, 5, 177, 89, 2, 334, 335,
	5, 183, 92, 2, 335, 336, 5, 155, 78, 2, 336, 337, 5, 185, 93, 2, 337, 60,
	3, 2, 2, 2, 338, 339, 5, 155, 78, 2, 339, 340, 5, 165, 83, 2, 340, 341,
	5, 185, 93, 2, 341, 342, 5, 187, 94, 2, 342, 343, 5, 149, 75, 2, 343, 344,
	5, 175, 88, 2, 344, 345, 5, 153, 77, 2, 345, 346, 5, 157, 79, 2, 346, 62,
	3, 2, 2, 2, 347, 348, 5, 157, 79, 2, 348, 349, 5, 175, 88, 2, 349, 350,
	5, 155, 78, 2, 350, 351, 5, 185, 93, 2, 351, 352, 5, 193, 97, 2, 352, 353,
	5, 165, 83, 2, 353, 354, 5, 187, 94, 2, 354, 355, 5, 163, 82, 2, 355, 64,
	3, 2, 2, 2, 356, 357, 5, 159, 80, 2, 357, 358, 5, 165, 83, 2, 358, 359,
	5, 175, 88, 2, 359, 360, 5, 155, 78, 2, 360, 361, 5, 185, 93, 2, 361, 362,
	5, 187, 94, 2, 362, 363, 5, 183, 92, 2, 363, 364, 5, 165, 83, 2, 364, 365,
	5, 175, 88, 2, 365, 366, 5, 161, 81, 2, 366, 66, 3, 2, 2, 2, 367, 368,
	5, 157, 79, 2, 368, 369, 5, 195, 98, 2, 369, 370, 5, 179, 90, 2, 370, 68,
	3, 2, 2, 2, 371, 372, 5, 159, 80, 2, 372, 373, 5, 171, 86, 2, 373, 374,
	5, 177, 89, 2, 374, 375, 5, 177, 89, 2, 375, 376, 5, 183, 92, 2, 376, 70,
	3, 2, 2, 2, 377, 378, 5, 161, 81, 2, 378, 379, 5, 157, 79, 2, 379, 380,
	5, 187, 94, 2, 380, 381, 5, 193, 97, 2, 381, 382, 5, 177, 89, 2, 382, 383,
	5, 183, 92, 2, 383, 384, 5, 155, 78, 2, 384, 72, 3, 2, 2, 2, 385, 386,
	5, 163, 82, 2, 386, 387, 5, 157, 79, 2, 387, 388, 5, 195, 98, 2, 388, 389,
	5, 187, 94, 2, 389, 390, 5, 177, 89, 2, 390, 391, 5, 175, 88, 2, 391, 392,
	5, 189, 95, 2, 392, 393, 5, 173, 87, 2, 393, 394, 5, 151, 76, 2, 394, 395,
	5, 157, 79, 2, 395, 396, 5, 183, 92, 2, 396, 74, 3, 2, 2, 2, 397, 398,
	5, 171, 86, 2, 398, 399, 5, 177, 89, 2, 399, 400, 5, 161, 81, 2, 400, 76,
	3, 2, 2, 2, 401, 402, 5, 171, 86, 2, 402, 403, 5, 177, 89, 2, 403, 404,
	5, 161, 81, 2, 404, 405, 7, 51, 2, 2, 405, 406, 7, 50, 2, 2, 406, 78, 3,
	2, 2, 2, 407, 408, 5, 173, 87, 2, 408, 409, 5, 149, 75, 2, 409, 410, 5,
	195, 98, 2, 410, 80, 3, 2, 2, 2, 411, 412, 5, 173, 87, 2, 412, 413, 5,
	157, 79, 2, 413, 414, 5, 155, 78, 2, 414, 415, 5, 165, 83, 2, 415, 416,
	5, 149, 75, 2, 416, 417, 5, 175, 88, 2, 417, 82, 3, 2, 2, 2, 418, 419,
	5, 173, 87, 2, 419, 420, 5, 165, 83, 2, 420, 421, 5, 175, 88, 2, 421, 84,
	3, 2, 2, 2, 422, 423, 5, 173, 87, 2, 423, 424, 5, 177, 89, 2, 424, 425,
	5, 155, 78, 2, 425, 86, 3, 2, 2, 2, 426, 427, 5, 175, 88, 2, 427, 428,
	5, 189, 95, 2, 428, 429, 5, 171, 86, 2, 429, 430, 5, 171, 86, 2, 430, 88,
	3, 2, 2, 2, 431, 432, 5, 179, 90, 2, 432, 433, 5, 165, 83, 2, 433, 90,
	3, 2, 2, 2, 434, 435, 5, 179, 90, 2, 435, 436, 5, 177, 89, 2, 436, 437,
	5, 193, 97, 2, 437, 92, 3, 2, 2, 2, 438, 439, 5, 183, 92, 2, 439, 440,
	5, 149, 75, 2, 440, 441, 5, 175, 88, 2, 441, 442, 5, 155, 78, 2, 442, 94,
	3, 2, 2, 2, 443, 444, 5, 183, 92, 2, 444, 445, 5, 149, 75, 2, 445, 446,
	5, 175, 88, 2, 446, 447, 5, 155, 78, 2, 447, 448, 5, 165, 83, 2, 448, 449,
	5, 175, 88, 2, 449, 450, 5, 187, 94, 2, 450, 96, 3, 2, 2, 2, 451, 452,
	5, 183, 92, 2, 452, 453, 5, 177, 89, 2, 453, 454, 5, 189, 95, 2, 454, 455,
	5, 175, 88, 2, 455, 456, 5, 155, 78, 2, 456, 98, 3, 2, 2, 2, 457, 458,
	5, 185, 93, 2, 458, 459, 5, 165, 83, 2, 459, 460, 5, 175, 88, 2, 460, 100,
	3, 2, 2, 2, 461, 462, 5, 185, 93, 2, 462, 463, 5, 165, 83, 2, 463, 464,
	5, 175, 88, 2, 464, 465, 5, 163, 82, 2, 465, 102, 3, 2, 2, 2, 466, 467,
	5, 185, 93, 2, 467, 468, 5, 181, 91, 2, 468, 469, 5, 183, 92, 2, 469, 470,
	5, 187, 94, 2, 470, 104, 3, 2, 2, 2, 471, 472, 5, 185, 93, 2, 472, 473,
	5, 193, 97, 2, 473, 474, 5, 165, 83, 2, 474, 475, 5, 187, 94, 2, 475, 476,
	5, 153, 77, 2, 476, 477, 5, 163, 82, 2, 477, 106, 3, 2, 2, 2, 478, 479,
	5, 187, 94, 2, 479, 480, 5, 149, 75, 2, 480, 481, 5, 175, 88, 2, 481, 108,
	3, 2, 2, 2, 482, 483, 5, 187, 94, 2, 483, 484, 5, 149, 75, 2, 484, 485,
	5, 175, 88, 2, 485, 486, 5, 163, 82, 2, 486, 110, 3, 2, 2, 2, 487, 488,
	5, 165, 83, 2, 488, 489, 5, 165, 83, 2, 489, 490, 5, 159, 80, 2, 490, 112,
	3, 2, 2, 2, 491, 492, 5, 165, 83, 2, 492, 493, 5, 175, 88, 2, 493, 114,
	3, 2, 2, 2, 494, 495, 5, 175, 88, 2, 495, 496, 5, 177, 89, 2, 496, 497,
	5, 187, 94, 2, 497, 116, 3, 2, 2, 2, 498, 499, 5, 149, 75, 2, 499, 500,
	5, 175, 88, 2, 500, 501, 5, 155, 78, 2, 501, 118, 3, 2, 2, 2, 502, 503,
	5, 177, 89, 2, 503, 504, 5, 183, 92, 2, 504, 120, 3, 2, 2, 2, 505, 506,
	5, 165, 83, 2, 506, 507, 5, 159, 80, 2, 507, 122, 3, 2, 2, 2, 508, 509,
	5, 187, 94, 2, 509, 510, 5, 163, 82, 2, 510, 511, 5, 157, 79, 2, 511, 512,
	5, 175, 88, 2, 512, 124, 3, 2, 2, 2, 513, 514, 5, 157, 79, 2, 514, 515,
	5, 171, 86, 2, 515, 516, 5, 185, 93, 2, 516, 517, 5, 157, 79, 2, 517, 126,
	3, 2, 2, 2, 518, 519, 5, 157, 79, 2, 519, 520, 5, 171, 86, 2, 520, 521,
	5, 185, 93, 2, 521, 522, 5, 157, 79, 2, 522, 523, 5, 165, 83, 2, 523, 524,
	5, 159, 80, 2, 524, 128, 3, 2, 2, 2, 525, 526, 5, 157, 79, 2, 526, 527,
	5, 175, 88, 2, 527, 528, 5, 155, 78, 2, 528, 529, 5, 165, 83, 2, 529, 530,
	5, 159, 80, 2, 530, 130, 3, 2, 2, 2, 531, 534, 5, 207, 104, 2, 532, 534,
	5, 209, 105, 2, 533, 531, 3, 2, 2, 2, 533, 532, 3, 2, 2, 2, 534, 132, 3,
	2, 2, 2, 535, 537, 5, 201, 101, 2, 536, 535, 3, 2, 2, 2, 537, 538, 3, 2,
	2, 2, 538, 536, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 134, 3, 2, 2, 2,
	540, 542, 5, 201, 101, 2, 541, 540, 3, 2, 2, 2, 542, 545, 3, 2, 2, 2, 543,
	541, 3, 2, 2, 2, 543, 544, 3, 2, 2, 2, 544, 546, 3, 2, 2, 2, 545, 543,
	3, 2, 2, 2, 546, 548, 7, 48, 2, 2, 547, 549, 5, 201, 101, 2, 548, 547,
	3, 2, 2, 2, 549, 550, 3, 2, 2, 2, 550, 548, 3, 2, 2, 2, 550, 551, 3, 2,
	2, 2, 551, 136, 3, 2, 2, 2, 552, 553, 9, 2, 2, 2, 553, 554, 5, 203, 102,
	2, 554, 555, 9, 2, 2, 2, 555, 138, 3, 2, 2, 2, 556, 557, 9, 2, 2, 2, 557,
	558, 5, 205, 103, 2, 558, 559, 9, 2, 2, 2, 559, 140, 3, 2, 2, 2, 560, 562,
	7, 93, 2, 2, 561, 563, 10, 3, 2, 2, 562, 561, 3, 2, 2, 2, 563, 564, 3,
	2, 2, 2, 564, 562, 3, 2, 2, 2, 564, 565, 3, 2, 2, 2, 565, 566, 3, 2, 2,
	2, 566, 567, 7, 95, 2, 2, 567, 142, 3, 2, 2, 2, 568, 572, 9, 4, 2, 2, 569,
	571, 10, 4, 2, 2, 570, 569, 3, 2, 2, 2, 571, 574, 3, 2, 2, 2, 572, 570,
	3, 2, 2, 2, 572, 573, 3, 2, 2, 2, 573, 575, 3, 2, 2, 2, 574, 572, 3, 2,
	2, 2, 575, 576, 9, 4, 2, 2, 576, 144, 3, 2, 2, 2, 577, 581, 7, 36, 2, 2,
	578, 580, 10, 5, 2, 2, 579, 578, 3, 2, 2, 2, 580, 583, 3, 2, 2, 2, 581,
	579, 3, 2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 584, 3, 2, 2, 2, 583, 581,
	3, 2, 2, 2, 584, 585, 7, 36, 2, 2, 585, 146, 3, 2, 2, 2, 586, 588, 9, 6,
	2, 2, 587, 586, 3, 2, 2, 2, 588, 589, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2,
	589, 590, 3, 2, 2, 2, 590, 591, 3, 2, 2, 2, 591, 592, 8, 74, 2, 2, 592,
	148, 3, 2, 2, 2, 593, 594, 9, 7, 2, 2, 594, 150, 3, 2, 2, 2, 595, 596,
	9, 8, 2, 2, 596, 152, 3, 2, 2, 2, 597, 598, 9, 9, 2, 2, 598, 154, 3, 2,
	2, 2, 599, 600, 9, 10, 2, 2, 600, 156, 3, 2, 2, 2, 601, 602, 9, 11, 2,
	2, 602, 158, 3, 2, 2, 2, 603, 604, 9, 12, 2, 2, 604, 160, 3, 2, 2, 2, 605,
	606, 9, 13, 2, 2, 606, 162, 3, 2, 2, 2, 607, 608, 9, 14, 2, 2, 608, 164,
	3, 2, 2, 2, 609, 610, 9, 15, 2, 2, 610, 166, 3, 2, 2, 2, 611, 612, 9, 16,
	2, 2, 612, 168, 3, 2, 2, 2, 613, 614, 9, 17, 2, 2, 614, 170, 3, 2, 2, 2,
	615, 616, 9, 18, 2, 2, 616, 172, 3, 2, 2, 2, 617, 618, 9, 19, 2, 2, 618,
	174, 3, 2, 2, 2, 619, 620, 9, 20, 2, 2, 620, 176, 3, 2, 2, 2, 621, 622,
	9, 21, 2, 2, 622, 178, 3, 2, 2, 2, 623, 624, 9, 22, 2, 2, 624, 180, 3,
	2, 2, 2, 625, 626, 9, 23, 2, 2, 626, 182, 3, 2, 2, 2, 627, 628, 9, 24,
	2, 2, 628, 184, 3, 2, 2, 2, 629, 630, 9, 25, 2, 2, 630, 186, 3, 2, 2, 2,
	631, 632, 9, 26, 2, 2, 632, 188, 3, 2, 2, 2, 633, 634, 9, 27, 2, 2, 634,
	190, 3, 2, 2, 2, 635, 636, 9, 28, 2, 2, 636, 192, 3, 2, 2, 2, 637, 638,
	9, 29, 2, 2, 638, 194, 3, 2, 2, 2, 639, 640, 9, 30, 2, 2, 640, 196, 3,
	2, 2, 2, 641, 642, 9, 31, 2, 2, 642, 198, 3, 2, 2, 2, 643, 644, 9, 32,
	2, 2, 644, 200, 3, 2, 2, 2, 645, 646, 9, 33, 2, 2, 646, 202, 3, 2, 2, 2,
	647, 648, 5, 201, 101, 2, 648, 649, 5, 201, 101, 2, 649, 650, 5, 201, 101,
	2, 650, 651, 5, 201, 101, 2, 651, 652, 7, 47, 2, 2, 652, 653, 5, 201, 101,
	2, 653, 654, 5, 201, 101, 2, 654, 655, 7, 47, 2, 2, 655, 656, 5, 201, 101,
	2, 656, 657, 5, 201, 101, 2, 657, 204, 3, 2, 2, 2, 658, 659, 5, 203, 102,
	2, 659, 660, 7, 34, 2, 2, 660, 661, 5, 201, 101, 2, 661, 662, 5, 201, 101,
	2, 662, 663, 7, 60, 2, 2, 663, 664, 5, 201, 101, 2, 664, 665, 5, 201, 101,
	2, 665, 666, 7, 60, 2, 2, 666, 667, 5, 201, 101, 2, 667, 668, 5, 201, 101,
	2, 668, 206, 3, 2, 2, 2, 669, 670, 5, 187, 94, 2, 670, 671, 5, 183, 92,
	2, 671, 672, 5, 189, 95, 2, 672, 673, 5, 157, 79, 2, 673, 208, 3, 2, 2,
	2, 674, 675, 5, 159, 80, 2, 675, 676, 5, 149, 75, 2, 676, 677, 5, 171,
	86, 2, 677, 678, 5, 185, 93, 2, 678, 679, 5, 157, 79, 2, 679, 210, 3, 2,
	2, 2, 11, 2, 533, 538, 543, 550, 564, 572, 581, 589, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'>='", "'<'",
	"'<='", "'!='", "','", "'&&'", "'||'", "'()'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Abs",
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Log", "Log10", "Max", "Median",
	"Min", "Mod", "Null", "Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh",
	"Sqrt", "Switch", "Tan", "Tanh", "Iif", "In", "Not", "And", "Or", "If",
	"Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal", "Date",
	"Datetime", "Field", "SingleQuoteString", "DoubleQuoteString", "Whitespace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Abs", "Acos",
	"Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Log", "Log10", "Max", "Median",
	"Min", "Mod", "Null", "Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh",
	"Sqrt", "Switch", "Tan", "Tanh", "Iif", "In", "Not", "And", "Or", "If",
	"Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal", "Date",
	"Datetime", "Field", "SingleQuoteString", "DoubleQuoteString", "Whitespace",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
	"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "Digit", "DateLiteral",
	"DateTimeLiteral", "TrueLiteral", "FalseLiteral",
}

type AlteryxFormulasLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewAlteryxFormulasLexer(input antlr.CharStream) *AlteryxFormulasLexer {

	l := new(AlteryxFormulasLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "AlteryxFormulas.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlteryxFormulasLexer tokens.
const (
	AlteryxFormulasLexerT__0              = 1
	AlteryxFormulasLexerT__1              = 2
	AlteryxFormulasLexerT__2              = 3
	AlteryxFormulasLexerT__3              = 4
	AlteryxFormulasLexerT__4              = 5
	AlteryxFormulasLexerT__5              = 6
	AlteryxFormulasLexerT__6              = 7
	AlteryxFormulasLexerT__7              = 8
	AlteryxFormulasLexerT__8              = 9
	AlteryxFormulasLexerT__9              = 10
	AlteryxFormulasLexerT__10             = 11
	AlteryxFormulasLexerT__11             = 12
	AlteryxFormulasLexerT__12             = 13
	AlteryxFormulasLexerT__13             = 14
	AlteryxFormulasLexerT__14             = 15
	AlteryxFormulasLexerT__15             = 16
	AlteryxFormulasLexerAbs               = 17
	AlteryxFormulasLexerAcos              = 18
	AlteryxFormulasLexerAsin              = 19
	AlteryxFormulasLexerAtan              = 20
	AlteryxFormulasLexerAtan2             = 21
	AlteryxFormulasLexerAverage           = 22
	AlteryxFormulasLexerCeil              = 23
	AlteryxFormulasLexerCharFromInt       = 24
	AlteryxFormulasLexerCharToInt         = 25
	AlteryxFormulasLexerContains          = 26
	AlteryxFormulasLexerCos               = 27
	AlteryxFormulasLexerCosh              = 28
	AlteryxFormulasLexerCountWords        = 29
	AlteryxFormulasLexerDistance          = 30
	AlteryxFormulasLexerEndsWith          = 31
	AlteryxFormulasLexerFindString        = 32
	AlteryxFormulasLexerExp               = 33
	AlteryxFormulasLexerFloor             = 34
	AlteryxFormulasLexerGetWord           = 35
	AlteryxFormulasLexerHexToNumber       = 36
	AlteryxFormulasLexerLog               = 37
	AlteryxFormulasLexerLog10             = 38
	AlteryxFormulasLexerMax               = 39
	AlteryxFormulasLexerMedian            = 40
	AlteryxFormulasLexerMin               = 41
	AlteryxFormulasLexerMod               = 42
	AlteryxFormulasLexerNull              = 43
	AlteryxFormulasLexerPi                = 44
	AlteryxFormulasLexerPow               = 45
	AlteryxFormulasLexerRand              = 46
	AlteryxFormulasLexerRandInt           = 47
	AlteryxFormulasLexerRound             = 48
	AlteryxFormulasLexerSin               = 49
	AlteryxFormulasLexerSinh              = 50
	AlteryxFormulasLexerSqrt              = 51
	AlteryxFormulasLexerSwitch            = 52
	AlteryxFormulasLexerTan               = 53
	AlteryxFormulasLexerTanh              = 54
	AlteryxFormulasLexerIif               = 55
	AlteryxFormulasLexerIn                = 56
	AlteryxFormulasLexerNot               = 57
	AlteryxFormulasLexerAnd               = 58
	AlteryxFormulasLexerOr                = 59
	AlteryxFormulasLexerIf                = 60
	AlteryxFormulasLexerThen              = 61
	AlteryxFormulasLexerElse              = 62
	AlteryxFormulasLexerElseif            = 63
	AlteryxFormulasLexerEndif             = 64
	AlteryxFormulasLexerBool              = 65
	AlteryxFormulasLexerInteger           = 66
	AlteryxFormulasLexerDecimal           = 67
	AlteryxFormulasLexerDate              = 68
	AlteryxFormulasLexerDatetime          = 69
	AlteryxFormulasLexerField             = 70
	AlteryxFormulasLexerSingleQuoteString = 71
	AlteryxFormulasLexerDoubleQuoteString = 72
	AlteryxFormulasLexerWhitespace        = 73
)
