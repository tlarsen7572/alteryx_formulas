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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 65, 566,
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
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41,
	3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 55, 3, 55, 3, 56, 3, 56, 5, 56, 420, 10, 56, 3, 57, 6, 57, 423, 10,
	57, 13, 57, 14, 57, 424, 3, 58, 7, 58, 428, 10, 58, 12, 58, 14, 58, 431,
	11, 58, 3, 58, 3, 58, 6, 58, 435, 10, 58, 13, 58, 14, 58, 436, 3, 59, 3,
	59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 6, 61, 449,
	10, 61, 13, 61, 14, 61, 450, 3, 61, 3, 61, 3, 62, 3, 62, 7, 62, 457, 10,
	62, 12, 62, 14, 62, 460, 11, 62, 3, 62, 3, 62, 3, 63, 3, 63, 7, 63, 466,
	10, 63, 12, 63, 14, 63, 469, 11, 63, 3, 63, 3, 63, 3, 64, 6, 64, 474, 10,
	64, 13, 64, 14, 64, 475, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67,
	3, 67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3,
	72, 3, 73, 3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77,
	3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3,
	83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87, 3, 88,
	3, 88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 3, 92, 3,
	92, 3, 92, 3, 92, 3, 92, 3, 92, 3, 92, 3, 92, 3, 92, 3, 93, 3, 93, 3, 93,
	3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 94, 3, 94, 3,
	94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 2, 2, 96, 3,
	3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40,
	79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49,
	97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147,
	2, 149, 2, 151, 2, 153, 2, 155, 2, 157, 2, 159, 2, 161, 2, 163, 2, 165,
	2, 167, 2, 169, 2, 171, 2, 173, 2, 175, 2, 177, 2, 179, 2, 181, 2, 183,
	2, 185, 2, 187, 2, 189, 2, 3, 2, 34, 3, 2, 98, 98, 3, 2, 95, 95, 3, 2,
	41, 41, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 67, 99, 99,
	4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102,
	4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105,
	4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108,
	4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111,
	4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114,
	4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117,
	4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120,
	4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123,
	4, 2, 92, 92, 124, 124, 3, 2, 50, 59, 2, 542, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 3, 191, 3, 2, 2, 2, 5, 193, 3, 2, 2, 2, 7, 195, 3,
	2, 2, 2, 9, 197, 3, 2, 2, 2, 11, 199, 3, 2, 2, 2, 13, 201, 3, 2, 2, 2,
	15, 203, 3, 2, 2, 2, 17, 205, 3, 2, 2, 2, 19, 207, 3, 2, 2, 2, 21, 210,
	3, 2, 2, 2, 23, 212, 3, 2, 2, 2, 25, 215, 3, 2, 2, 2, 27, 218, 3, 2, 2,
	2, 29, 220, 3, 2, 2, 2, 31, 223, 3, 2, 2, 2, 33, 226, 3, 2, 2, 2, 35, 229,
	3, 2, 2, 2, 37, 233, 3, 2, 2, 2, 39, 238, 3, 2, 2, 2, 41, 243, 3, 2, 2,
	2, 43, 248, 3, 2, 2, 2, 45, 254, 3, 2, 2, 2, 47, 262, 3, 2, 2, 2, 49, 267,
	3, 2, 2, 2, 51, 271, 3, 2, 2, 2, 53, 276, 3, 2, 2, 2, 55, 285, 3, 2, 2,
	2, 57, 289, 3, 2, 2, 2, 59, 295, 3, 2, 2, 2, 61, 299, 3, 2, 2, 2, 63, 305,
	3, 2, 2, 2, 65, 309, 3, 2, 2, 2, 67, 316, 3, 2, 2, 2, 69, 320, 3, 2, 2,
	2, 71, 324, 3, 2, 2, 2, 73, 329, 3, 2, 2, 2, 75, 332, 3, 2, 2, 2, 77, 336,
	3, 2, 2, 2, 79, 341, 3, 2, 2, 2, 81, 349, 3, 2, 2, 2, 83, 355, 3, 2, 2,
	2, 85, 359, 3, 2, 2, 2, 87, 364, 3, 2, 2, 2, 89, 369, 3, 2, 2, 2, 91, 373,
	3, 2, 2, 2, 93, 377, 3, 2, 2, 2, 95, 380, 3, 2, 2, 2, 97, 384, 3, 2, 2,
	2, 99, 388, 3, 2, 2, 2, 101, 391, 3, 2, 2, 2, 103, 394, 3, 2, 2, 2, 105,
	399, 3, 2, 2, 2, 107, 404, 3, 2, 2, 2, 109, 411, 3, 2, 2, 2, 111, 419,
	3, 2, 2, 2, 113, 422, 3, 2, 2, 2, 115, 429, 3, 2, 2, 2, 117, 438, 3, 2,
	2, 2, 119, 442, 3, 2, 2, 2, 121, 446, 3, 2, 2, 2, 123, 454, 3, 2, 2, 2,
	125, 463, 3, 2, 2, 2, 127, 473, 3, 2, 2, 2, 129, 479, 3, 2, 2, 2, 131,
	481, 3, 2, 2, 2, 133, 483, 3, 2, 2, 2, 135, 485, 3, 2, 2, 2, 137, 487,
	3, 2, 2, 2, 139, 489, 3, 2, 2, 2, 141, 491, 3, 2, 2, 2, 143, 493, 3, 2,
	2, 2, 145, 495, 3, 2, 2, 2, 147, 497, 3, 2, 2, 2, 149, 499, 3, 2, 2, 2,
	151, 501, 3, 2, 2, 2, 153, 503, 3, 2, 2, 2, 155, 505, 3, 2, 2, 2, 157,
	507, 3, 2, 2, 2, 159, 509, 3, 2, 2, 2, 161, 511, 3, 2, 2, 2, 163, 513,
	3, 2, 2, 2, 165, 515, 3, 2, 2, 2, 167, 517, 3, 2, 2, 2, 169, 519, 3, 2,
	2, 2, 171, 521, 3, 2, 2, 2, 173, 523, 3, 2, 2, 2, 175, 525, 3, 2, 2, 2,
	177, 527, 3, 2, 2, 2, 179, 529, 3, 2, 2, 2, 181, 531, 3, 2, 2, 2, 183,
	533, 3, 2, 2, 2, 185, 544, 3, 2, 2, 2, 187, 555, 3, 2, 2, 2, 189, 560,
	3, 2, 2, 2, 191, 192, 7, 42, 2, 2, 192, 4, 3, 2, 2, 2, 193, 194, 7, 43,
	2, 2, 194, 6, 3, 2, 2, 2, 195, 196, 7, 44, 2, 2, 196, 8, 3, 2, 2, 2, 197,
	198, 7, 49, 2, 2, 198, 10, 3, 2, 2, 2, 199, 200, 7, 45, 2, 2, 200, 12,
	3, 2, 2, 2, 201, 202, 7, 47, 2, 2, 202, 14, 3, 2, 2, 2, 203, 204, 7, 63,
	2, 2, 204, 16, 3, 2, 2, 2, 205, 206, 7, 64, 2, 2, 206, 18, 3, 2, 2, 2,
	207, 208, 7, 64, 2, 2, 208, 209, 7, 63, 2, 2, 209, 20, 3, 2, 2, 2, 210,
	211, 7, 62, 2, 2, 211, 22, 3, 2, 2, 2, 212, 213, 7, 62, 2, 2, 213, 214,
	7, 63, 2, 2, 214, 24, 3, 2, 2, 2, 215, 216, 7, 35, 2, 2, 216, 217, 7, 63,
	2, 2, 217, 26, 3, 2, 2, 2, 218, 219, 7, 46, 2, 2, 219, 28, 3, 2, 2, 2,
	220, 221, 7, 40, 2, 2, 221, 222, 7, 40, 2, 2, 222, 30, 3, 2, 2, 2, 223,
	224, 7, 126, 2, 2, 224, 225, 7, 126, 2, 2, 225, 32, 3, 2, 2, 2, 226, 227,
	7, 42, 2, 2, 227, 228, 7, 43, 2, 2, 228, 34, 3, 2, 2, 2, 229, 230, 5, 129,
	65, 2, 230, 231, 5, 131, 66, 2, 231, 232, 5, 165, 83, 2, 232, 36, 3, 2,
	2, 2, 233, 234, 5, 129, 65, 2, 234, 235, 5, 133, 67, 2, 235, 236, 5, 157,
	79, 2, 236, 237, 5, 165, 83, 2, 237, 38, 3, 2, 2, 2, 238, 239, 5, 129,
	65, 2, 239, 240, 5, 165, 83, 2, 240, 241, 5, 145, 73, 2, 241, 242, 5, 155,
	78, 2, 242, 40, 3, 2, 2, 2, 243, 244, 5, 129, 65, 2, 244, 245, 5, 167,
	84, 2, 245, 246, 5, 129, 65, 2, 246, 247, 5, 155, 78, 2, 247, 42, 3, 2,
	2, 2, 248, 249, 5, 129, 65, 2, 249, 250, 5, 167, 84, 2, 250, 251, 5, 129,
	65, 2, 251, 252, 5, 155, 78, 2, 252, 253, 7, 52, 2, 2, 253, 44, 3, 2, 2,
	2, 254, 255, 5, 129, 65, 2, 255, 256, 5, 171, 86, 2, 256, 257, 5, 137,
	69, 2, 257, 258, 5, 163, 82, 2, 258, 259, 5, 129, 65, 2, 259, 260, 5, 141,
	71, 2, 260, 261, 5, 137, 69, 2, 261, 46, 3, 2, 2, 2, 262, 263, 5, 133,
	67, 2, 263, 264, 5, 137, 69, 2, 264, 265, 5, 145, 73, 2, 265, 266, 5, 151,
	76, 2, 266, 48, 3, 2, 2, 2, 267, 268, 5, 133, 67, 2, 268, 269, 5, 157,
	79, 2, 269, 270, 5, 165, 83, 2, 270, 50, 3, 2, 2, 2, 271, 272, 5, 133,
	67, 2, 272, 273, 5, 157, 79, 2, 273, 274, 5, 165, 83, 2, 274, 275, 5, 143,
	72, 2, 275, 52, 3, 2, 2, 2, 276, 277, 5, 135, 68, 2, 277, 278, 5, 145,
	73, 2, 278, 279, 5, 165, 83, 2, 279, 280, 5, 167, 84, 2, 280, 281, 5, 129,
	65, 2, 281, 282, 5, 155, 78, 2, 282, 283, 5, 133, 67, 2, 283, 284, 5, 137,
	69, 2, 284, 54, 3, 2, 2, 2, 285, 286, 5, 137, 69, 2, 286, 287, 5, 175,
	88, 2, 287, 288, 5, 159, 80, 2, 288, 56, 3, 2, 2, 2, 289, 290, 5, 139,
	70, 2, 290, 291, 5, 151, 76, 2, 291, 292, 5, 157, 79, 2, 292, 293, 5, 157,
	79, 2, 293, 294, 5, 163, 82, 2, 294, 58, 3, 2, 2, 2, 295, 296, 5, 151,
	76, 2, 296, 297, 5, 157, 79, 2, 297, 298, 5, 141, 71, 2, 298, 60, 3, 2,
	2, 2, 299, 300, 5, 151, 76, 2, 300, 301, 5, 157, 79, 2, 301, 302, 5, 141,
	71, 2, 302, 303, 7, 51, 2, 2, 303, 304, 7, 50, 2, 2, 304, 62, 3, 2, 2,
	2, 305, 306, 5, 153, 77, 2, 306, 307, 5, 129, 65, 2, 307, 308, 5, 175,
	88, 2, 308, 64, 3, 2, 2, 2, 309, 310, 5, 153, 77, 2, 310, 311, 5, 137,
	69, 2, 311, 312, 5, 135, 68, 2, 312, 313, 5, 145, 73, 2, 313, 314, 5, 129,
	65, 2, 314, 315, 5, 155, 78, 2, 315, 66, 3, 2, 2, 2, 316, 317, 5, 153,
	77, 2, 317, 318, 5, 145, 73, 2, 318, 319, 5, 155, 78, 2, 319, 68, 3, 2,
	2, 2, 320, 321, 5, 153, 77, 2, 321, 322, 5, 157, 79, 2, 322, 323, 5, 135,
	68, 2, 323, 70, 3, 2, 2, 2, 324, 325, 5, 155, 78, 2, 325, 326, 5, 169,
	85, 2, 326, 327, 5, 151, 76, 2, 327, 328, 5, 151, 76, 2, 328, 72, 3, 2,
	2, 2, 329, 330, 5, 159, 80, 2, 330, 331, 5, 145, 73, 2, 331, 74, 3, 2,
	2, 2, 332, 333, 5, 159, 80, 2, 333, 334, 5, 157, 79, 2, 334, 335, 5, 173,
	87, 2, 335, 76, 3, 2, 2, 2, 336, 337, 5, 163, 82, 2, 337, 338, 5, 129,
	65, 2, 338, 339, 5, 155, 78, 2, 339, 340, 5, 135, 68, 2, 340, 78, 3, 2,
	2, 2, 341, 342, 5, 163, 82, 2, 342, 343, 5, 129, 65, 2, 343, 344, 5, 155,
	78, 2, 344, 345, 5, 135, 68, 2, 345, 346, 5, 145, 73, 2, 346, 347, 5, 155,
	78, 2, 347, 348, 5, 167, 84, 2, 348, 80, 3, 2, 2, 2, 349, 350, 5, 163,
	82, 2, 350, 351, 5, 157, 79, 2, 351, 352, 5, 169, 85, 2, 352, 353, 5, 155,
	78, 2, 353, 354, 5, 135, 68, 2, 354, 82, 3, 2, 2, 2, 355, 356, 5, 165,
	83, 2, 356, 357, 5, 145, 73, 2, 357, 358, 5, 155, 78, 2, 358, 84, 3, 2,
	2, 2, 359, 360, 5, 165, 83, 2, 360, 361, 5, 145, 73, 2, 361, 362, 5, 155,
	78, 2, 362, 363, 5, 143, 72, 2, 363, 86, 3, 2, 2, 2, 364, 365, 5, 165,
	83, 2, 365, 366, 5, 161, 81, 2, 366, 367, 5, 163, 82, 2, 367, 368, 5, 167,
	84, 2, 368, 88, 3, 2, 2, 2, 369, 370, 5, 167, 84, 2, 370, 371, 5, 129,
	65, 2, 371, 372, 5, 155, 78, 2, 372, 90, 3, 2, 2, 2, 373, 374, 5, 145,
	73, 2, 374, 375, 5, 145, 73, 2, 375, 376, 5, 139, 70, 2, 376, 92, 3, 2,
	2, 2, 377, 378, 5, 145, 73, 2, 378, 379, 5, 155, 78, 2, 379, 94, 3, 2,
	2, 2, 380, 381, 5, 155, 78, 2, 381, 382, 5, 157, 79, 2, 382, 383, 5, 167,
	84, 2, 383, 96, 3, 2, 2, 2, 384, 385, 5, 129, 65, 2, 385, 386, 5, 155,
	78, 2, 386, 387, 5, 135, 68, 2, 387, 98, 3, 2, 2, 2, 388, 389, 5, 157,
	79, 2, 389, 390, 5, 163, 82, 2, 390, 100, 3, 2, 2, 2, 391, 392, 5, 145,
	73, 2, 392, 393, 5, 139, 70, 2, 393, 102, 3, 2, 2, 2, 394, 395, 5, 167,
	84, 2, 395, 396, 5, 143, 72, 2, 396, 397, 5, 137, 69, 2, 397, 398, 5, 155,
	78, 2, 398, 104, 3, 2, 2, 2, 399, 400, 5, 137, 69, 2, 400, 401, 5, 151,
	76, 2, 401, 402, 5, 165, 83, 2, 402, 403, 5, 137, 69, 2, 403, 106, 3, 2,
	2, 2, 404, 405, 5, 137, 69, 2, 405, 406, 5, 151, 76, 2, 406, 407, 5, 165,
	83, 2, 407, 408, 5, 137, 69, 2, 408, 409, 5, 145, 73, 2, 409, 410, 5, 139,
	70, 2, 410, 108, 3, 2, 2, 2, 411, 412, 5, 137, 69, 2, 412, 413, 5, 155,
	78, 2, 413, 414, 5, 135, 68, 2, 414, 415, 5, 145, 73, 2, 415, 416, 5, 139,
	70, 2, 416, 110, 3, 2, 2, 2, 417, 420, 5, 187, 94, 2, 418, 420, 5, 189,
	95, 2, 419, 417, 3, 2, 2, 2, 419, 418, 3, 2, 2, 2, 420, 112, 3, 2, 2, 2,
	421, 423, 5, 181, 91, 2, 422, 421, 3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424,
	422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 114, 3, 2, 2, 2, 426, 428,
	5, 181, 91, 2, 427, 426, 3, 2, 2, 2, 428, 431, 3, 2, 2, 2, 429, 427, 3,
	2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 432, 3, 2, 2, 2, 431, 429, 3, 2, 2,
	2, 432, 434, 7, 48, 2, 2, 433, 435, 5, 181, 91, 2, 434, 433, 3, 2, 2, 2,
	435, 436, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437,
	116, 3, 2, 2, 2, 438, 439, 9, 2, 2, 2, 439, 440, 5, 183, 92, 2, 440, 441,
	9, 2, 2, 2, 441, 118, 3, 2, 2, 2, 442, 443, 9, 2, 2, 2, 443, 444, 5, 185,
	93, 2, 444, 445, 9, 2, 2, 2, 445, 120, 3, 2, 2, 2, 446, 448, 7, 93, 2,
	2, 447, 449, 10, 3, 2, 2, 448, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450,
	448, 3, 2, 2, 2, 450, 451, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 453,
	7, 95, 2, 2, 453, 122, 3, 2, 2, 2, 454, 458, 9, 4, 2, 2, 455, 457, 10,
	4, 2, 2, 456, 455, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2, 458, 456, 3, 2, 2,
	2, 458, 459, 3, 2, 2, 2, 459, 461, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 461,
	462, 9, 4, 2, 2, 462, 124, 3, 2, 2, 2, 463, 467, 7, 36, 2, 2, 464, 466,
	10, 5, 2, 2, 465, 464, 3, 2, 2, 2, 466, 469, 3, 2, 2, 2, 467, 465, 3, 2,
	2, 2, 467, 468, 3, 2, 2, 2, 468, 470, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2,
	470, 471, 7, 36, 2, 2, 471, 126, 3, 2, 2, 2, 472, 474, 9, 6, 2, 2, 473,
	472, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 473, 3, 2, 2, 2, 475, 476,
	3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 478, 8, 64, 2, 2, 478, 128, 3, 2,
	2, 2, 479, 480, 9, 7, 2, 2, 480, 130, 3, 2, 2, 2, 481, 482, 9, 8, 2, 2,
	482, 132, 3, 2, 2, 2, 483, 484, 9, 9, 2, 2, 484, 134, 3, 2, 2, 2, 485,
	486, 9, 10, 2, 2, 486, 136, 3, 2, 2, 2, 487, 488, 9, 11, 2, 2, 488, 138,
	3, 2, 2, 2, 489, 490, 9, 12, 2, 2, 490, 140, 3, 2, 2, 2, 491, 492, 9, 13,
	2, 2, 492, 142, 3, 2, 2, 2, 493, 494, 9, 14, 2, 2, 494, 144, 3, 2, 2, 2,
	495, 496, 9, 15, 2, 2, 496, 146, 3, 2, 2, 2, 497, 498, 9, 16, 2, 2, 498,
	148, 3, 2, 2, 2, 499, 500, 9, 17, 2, 2, 500, 150, 3, 2, 2, 2, 501, 502,
	9, 18, 2, 2, 502, 152, 3, 2, 2, 2, 503, 504, 9, 19, 2, 2, 504, 154, 3,
	2, 2, 2, 505, 506, 9, 20, 2, 2, 506, 156, 3, 2, 2, 2, 507, 508, 9, 21,
	2, 2, 508, 158, 3, 2, 2, 2, 509, 510, 9, 22, 2, 2, 510, 160, 3, 2, 2, 2,
	511, 512, 9, 23, 2, 2, 512, 162, 3, 2, 2, 2, 513, 514, 9, 24, 2, 2, 514,
	164, 3, 2, 2, 2, 515, 516, 9, 25, 2, 2, 516, 166, 3, 2, 2, 2, 517, 518,
	9, 26, 2, 2, 518, 168, 3, 2, 2, 2, 519, 520, 9, 27, 2, 2, 520, 170, 3,
	2, 2, 2, 521, 522, 9, 28, 2, 2, 522, 172, 3, 2, 2, 2, 523, 524, 9, 29,
	2, 2, 524, 174, 3, 2, 2, 2, 525, 526, 9, 30, 2, 2, 526, 176, 3, 2, 2, 2,
	527, 528, 9, 31, 2, 2, 528, 178, 3, 2, 2, 2, 529, 530, 9, 32, 2, 2, 530,
	180, 3, 2, 2, 2, 531, 532, 9, 33, 2, 2, 532, 182, 3, 2, 2, 2, 533, 534,
	5, 181, 91, 2, 534, 535, 5, 181, 91, 2, 535, 536, 5, 181, 91, 2, 536, 537,
	5, 181, 91, 2, 537, 538, 7, 47, 2, 2, 538, 539, 5, 181, 91, 2, 539, 540,
	5, 181, 91, 2, 540, 541, 7, 47, 2, 2, 541, 542, 5, 181, 91, 2, 542, 543,
	5, 181, 91, 2, 543, 184, 3, 2, 2, 2, 544, 545, 5, 183, 92, 2, 545, 546,
	7, 34, 2, 2, 546, 547, 5, 181, 91, 2, 547, 548, 5, 181, 91, 2, 548, 549,
	7, 60, 2, 2, 549, 550, 5, 181, 91, 2, 550, 551, 5, 181, 91, 2, 551, 552,
	7, 60, 2, 2, 552, 553, 5, 181, 91, 2, 553, 554, 5, 181, 91, 2, 554, 186,
	3, 2, 2, 2, 555, 556, 5, 167, 84, 2, 556, 557, 5, 163, 82, 2, 557, 558,
	5, 169, 85, 2, 558, 559, 5, 137, 69, 2, 559, 188, 3, 2, 2, 2, 560, 561,
	5, 139, 70, 2, 561, 562, 5, 129, 65, 2, 562, 563, 5, 151, 76, 2, 563, 564,
	5, 165, 83, 2, 564, 565, 5, 137, 69, 2, 565, 190, 3, 2, 2, 2, 11, 2, 419,
	424, 429, 436, 450, 458, 467, 475, 3, 8, 2, 2,
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
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "Cos", "Cosh", "Distance",
	"Exp", "Floor", "Log", "Log10", "Max", "Median", "Min", "Mod", "Null",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh", "Sqrt", "Tan",
	"Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif", "Endif",
	"Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Abs", "Acos",
	"Asin", "Atan", "Atan2", "Average", "Ceil", "Cos", "Cosh", "Distance",
	"Exp", "Floor", "Log", "Log10", "Max", "Median", "Min", "Mod", "Null",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh", "Sqrt", "Tan",
	"Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif", "Endif",
	"Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace", "A", "B", "C", "D", "E", "F", "G", "H",
	"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
	"X", "Y", "Z", "Digit", "DateLiteral", "DateTimeLiteral", "TrueLiteral",
	"FalseLiteral",
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
	AlteryxFormulasLexerCos               = 24
	AlteryxFormulasLexerCosh              = 25
	AlteryxFormulasLexerDistance          = 26
	AlteryxFormulasLexerExp               = 27
	AlteryxFormulasLexerFloor             = 28
	AlteryxFormulasLexerLog               = 29
	AlteryxFormulasLexerLog10             = 30
	AlteryxFormulasLexerMax               = 31
	AlteryxFormulasLexerMedian            = 32
	AlteryxFormulasLexerMin               = 33
	AlteryxFormulasLexerMod               = 34
	AlteryxFormulasLexerNull              = 35
	AlteryxFormulasLexerPi                = 36
	AlteryxFormulasLexerPow               = 37
	AlteryxFormulasLexerRand              = 38
	AlteryxFormulasLexerRandInt           = 39
	AlteryxFormulasLexerRound             = 40
	AlteryxFormulasLexerSin               = 41
	AlteryxFormulasLexerSinh              = 42
	AlteryxFormulasLexerSqrt              = 43
	AlteryxFormulasLexerTan               = 44
	AlteryxFormulasLexerIif               = 45
	AlteryxFormulasLexerIn                = 46
	AlteryxFormulasLexerNot               = 47
	AlteryxFormulasLexerAnd               = 48
	AlteryxFormulasLexerOr                = 49
	AlteryxFormulasLexerIf                = 50
	AlteryxFormulasLexerThen              = 51
	AlteryxFormulasLexerElse              = 52
	AlteryxFormulasLexerElseif            = 53
	AlteryxFormulasLexerEndif             = 54
	AlteryxFormulasLexerBool              = 55
	AlteryxFormulasLexerInteger           = 56
	AlteryxFormulasLexerDecimal           = 57
	AlteryxFormulasLexerDate              = 58
	AlteryxFormulasLexerDatetime          = 59
	AlteryxFormulasLexerField             = 60
	AlteryxFormulasLexerSingleQuoteString = 61
	AlteryxFormulasLexerDoubleQuoteString = 62
	AlteryxFormulasLexerWhitespace        = 63
)
