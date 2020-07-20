// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 75, 374,
	4, 2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 69, 10, 2, 12, 2,
	14, 2, 72, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 98, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 135, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 183, 10, 2, 13, 2, 14, 2, 184,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 194, 10, 2, 12, 2, 14,
	2, 197, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 206, 10,
	2, 13, 2, 14, 2, 207, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 217,
	10, 2, 12, 2, 14, 2, 220, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 274, 10,
	2, 13, 2, 14, 2, 275, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 302, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 7, 2, 346, 10, 2, 12, 2, 14, 2, 349, 11, 2, 5, 2, 351, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 361, 10, 2, 12, 2, 14,
	2, 364, 11, 2, 5, 2, 366, 10, 2, 3, 2, 7, 2, 369, 10, 2, 12, 2, 14, 2,
	372, 11, 2, 3, 2, 2, 3, 2, 3, 2, 2, 4, 4, 2, 16, 16, 60, 60, 4, 2, 17,
	17, 61, 61, 2, 449, 2, 301, 3, 2, 2, 2, 4, 5, 8, 2, 1, 2, 5, 6, 7, 3, 2,
	2, 6, 7, 5, 2, 2, 2, 7, 8, 7, 4, 2, 2, 8, 302, 3, 2, 2, 2, 9, 10, 7, 62,
	2, 2, 10, 11, 5, 2, 2, 2, 11, 12, 7, 63, 2, 2, 12, 20, 5, 2, 2, 2, 13,
	14, 7, 65, 2, 2, 14, 15, 5, 2, 2, 2, 15, 16, 7, 63, 2, 2, 16, 17, 5, 2,
	2, 2, 17, 19, 3, 2, 2, 2, 18, 13, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18,
	3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	23, 24, 7, 64, 2, 2, 24, 25, 5, 2, 2, 2, 25, 26, 7, 66, 2, 2, 26, 302,
	3, 2, 2, 2, 27, 28, 7, 57, 2, 2, 28, 29, 7, 3, 2, 2, 29, 30, 5, 2, 2, 2,
	30, 31, 7, 15, 2, 2, 31, 32, 5, 2, 2, 2, 32, 33, 7, 15, 2, 2, 33, 34, 5,
	2, 2, 2, 34, 35, 7, 4, 2, 2, 35, 302, 3, 2, 2, 2, 36, 37, 7, 19, 2, 2,
	37, 38, 7, 3, 2, 2, 38, 39, 5, 2, 2, 2, 39, 40, 7, 4, 2, 2, 40, 302, 3,
	2, 2, 2, 41, 42, 7, 20, 2, 2, 42, 43, 7, 3, 2, 2, 43, 44, 5, 2, 2, 2, 44,
	45, 7, 4, 2, 2, 45, 302, 3, 2, 2, 2, 46, 47, 7, 21, 2, 2, 47, 48, 7, 3,
	2, 2, 48, 49, 5, 2, 2, 2, 49, 50, 7, 4, 2, 2, 50, 302, 3, 2, 2, 2, 51,
	52, 7, 22, 2, 2, 52, 53, 7, 3, 2, 2, 53, 54, 5, 2, 2, 2, 54, 55, 7, 4,
	2, 2, 55, 302, 3, 2, 2, 2, 56, 57, 7, 23, 2, 2, 57, 58, 7, 3, 2, 2, 58,
	59, 5, 2, 2, 2, 59, 60, 7, 15, 2, 2, 60, 61, 5, 2, 2, 2, 61, 62, 7, 4,
	2, 2, 62, 302, 3, 2, 2, 2, 63, 64, 7, 24, 2, 2, 64, 65, 7, 3, 2, 2, 65,
	70, 5, 2, 2, 2, 66, 67, 7, 15, 2, 2, 67, 69, 5, 2, 2, 2, 68, 66, 3, 2,
	2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 73,
	3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74, 7, 4, 2, 2, 74, 302, 3, 2, 2, 2,
	75, 76, 7, 25, 2, 2, 76, 77, 7, 3, 2, 2, 77, 78, 5, 2, 2, 2, 78, 79, 7,
	4, 2, 2, 79, 302, 3, 2, 2, 2, 80, 81, 7, 26, 2, 2, 81, 82, 7, 3, 2, 2,
	82, 83, 5, 2, 2, 2, 83, 84, 7, 4, 2, 2, 84, 302, 3, 2, 2, 2, 85, 86, 7,
	27, 2, 2, 86, 87, 7, 3, 2, 2, 87, 88, 5, 2, 2, 2, 88, 89, 7, 4, 2, 2, 89,
	302, 3, 2, 2, 2, 90, 91, 7, 28, 2, 2, 91, 92, 7, 3, 2, 2, 92, 93, 5, 2,
	2, 2, 93, 94, 7, 15, 2, 2, 94, 97, 5, 2, 2, 2, 95, 96, 7, 15, 2, 2, 96,
	98, 5, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 3, 2, 2,
	2, 99, 100, 7, 4, 2, 2, 100, 302, 3, 2, 2, 2, 101, 102, 7, 29, 2, 2, 102,
	103, 7, 3, 2, 2, 103, 104, 5, 2, 2, 2, 104, 105, 7, 4, 2, 2, 105, 302,
	3, 2, 2, 2, 106, 107, 7, 30, 2, 2, 107, 108, 7, 3, 2, 2, 108, 109, 5, 2,
	2, 2, 109, 110, 7, 4, 2, 2, 110, 302, 3, 2, 2, 2, 111, 112, 7, 31, 2, 2,
	112, 113, 7, 3, 2, 2, 113, 114, 5, 2, 2, 2, 114, 115, 7, 4, 2, 2, 115,
	302, 3, 2, 2, 2, 116, 117, 7, 32, 2, 2, 117, 118, 7, 3, 2, 2, 118, 119,
	5, 2, 2, 2, 119, 120, 7, 15, 2, 2, 120, 121, 5, 2, 2, 2, 121, 122, 7, 15,
	2, 2, 122, 123, 5, 2, 2, 2, 123, 124, 7, 15, 2, 2, 124, 125, 5, 2, 2, 2,
	125, 126, 7, 4, 2, 2, 126, 302, 3, 2, 2, 2, 127, 128, 7, 33, 2, 2, 128,
	129, 7, 3, 2, 2, 129, 130, 5, 2, 2, 2, 130, 131, 7, 15, 2, 2, 131, 134,
	5, 2, 2, 2, 132, 133, 7, 15, 2, 2, 133, 135, 5, 2, 2, 2, 134, 132, 3, 2,
	2, 2, 134, 135, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 137, 7, 4, 2, 2,
	137, 302, 3, 2, 2, 2, 138, 139, 7, 35, 2, 2, 139, 140, 7, 3, 2, 2, 140,
	141, 5, 2, 2, 2, 141, 142, 7, 4, 2, 2, 142, 302, 3, 2, 2, 2, 143, 144,
	7, 34, 2, 2, 144, 145, 7, 3, 2, 2, 145, 146, 5, 2, 2, 2, 146, 147, 7, 15,
	2, 2, 147, 148, 5, 2, 2, 2, 148, 149, 7, 4, 2, 2, 149, 302, 3, 2, 2, 2,
	150, 151, 7, 36, 2, 2, 151, 152, 7, 3, 2, 2, 152, 153, 5, 2, 2, 2, 153,
	154, 7, 4, 2, 2, 154, 302, 3, 2, 2, 2, 155, 156, 7, 37, 2, 2, 156, 157,
	7, 3, 2, 2, 157, 158, 5, 2, 2, 2, 158, 159, 7, 15, 2, 2, 159, 160, 5, 2,
	2, 2, 160, 161, 7, 4, 2, 2, 161, 302, 3, 2, 2, 2, 162, 163, 7, 38, 2, 2,
	163, 164, 7, 3, 2, 2, 164, 165, 5, 2, 2, 2, 165, 166, 7, 4, 2, 2, 166,
	302, 3, 2, 2, 2, 167, 168, 7, 39, 2, 2, 168, 169, 7, 3, 2, 2, 169, 170,
	5, 2, 2, 2, 170, 171, 7, 4, 2, 2, 171, 302, 3, 2, 2, 2, 172, 173, 7, 40,
	2, 2, 173, 174, 7, 3, 2, 2, 174, 175, 5, 2, 2, 2, 175, 176, 7, 4, 2, 2,
	176, 302, 3, 2, 2, 2, 177, 178, 7, 41, 2, 2, 178, 179, 7, 3, 2, 2, 179,
	182, 5, 2, 2, 2, 180, 181, 7, 15, 2, 2, 181, 183, 5, 2, 2, 2, 182, 180,
	3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2,
	2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 7, 4, 2, 2, 187, 302, 3, 2, 2, 2,
	188, 189, 7, 42, 2, 2, 189, 190, 7, 3, 2, 2, 190, 195, 5, 2, 2, 2, 191,
	192, 7, 15, 2, 2, 192, 194, 5, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 197,
	3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 198, 3, 2,
	2, 2, 197, 195, 3, 2, 2, 2, 198, 199, 7, 4, 2, 2, 199, 302, 3, 2, 2, 2,
	200, 201, 7, 43, 2, 2, 201, 202, 7, 3, 2, 2, 202, 205, 5, 2, 2, 2, 203,
	204, 7, 15, 2, 2, 204, 206, 5, 2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 207,
	3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2,
	2, 2, 209, 210, 7, 4, 2, 2, 210, 302, 3, 2, 2, 2, 211, 212, 7, 44, 2, 2,
	212, 213, 7, 3, 2, 2, 213, 218, 5, 2, 2, 2, 214, 215, 7, 15, 2, 2, 215,
	217, 5, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 220, 3, 2, 2, 2, 218, 216,
	3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 221, 3, 2, 2, 2, 220, 218, 3, 2,
	2, 2, 221, 222, 7, 4, 2, 2, 222, 302, 3, 2, 2, 2, 223, 224, 7, 45, 2, 2,
	224, 302, 7, 18, 2, 2, 225, 226, 7, 46, 2, 2, 226, 302, 7, 18, 2, 2, 227,
	228, 7, 47, 2, 2, 228, 229, 7, 3, 2, 2, 229, 230, 5, 2, 2, 2, 230, 231,
	7, 15, 2, 2, 231, 232, 5, 2, 2, 2, 232, 233, 7, 4, 2, 2, 233, 302, 3, 2,
	2, 2, 234, 235, 7, 48, 2, 2, 235, 302, 7, 18, 2, 2, 236, 237, 7, 49, 2,
	2, 237, 238, 7, 3, 2, 2, 238, 239, 5, 2, 2, 2, 239, 240, 7, 4, 2, 2, 240,
	302, 3, 2, 2, 2, 241, 242, 7, 50, 2, 2, 242, 243, 7, 3, 2, 2, 243, 244,
	5, 2, 2, 2, 244, 245, 7, 15, 2, 2, 245, 246, 5, 2, 2, 2, 246, 247, 7, 4,
	2, 2, 247, 302, 3, 2, 2, 2, 248, 249, 7, 51, 2, 2, 249, 250, 7, 3, 2, 2,
	250, 251, 5, 2, 2, 2, 251, 252, 7, 4, 2, 2, 252, 302, 3, 2, 2, 2, 253,
	254, 7, 52, 2, 2, 254, 255, 7, 3, 2, 2, 255, 256, 5, 2, 2, 2, 256, 257,
	7, 4, 2, 2, 257, 302, 3, 2, 2, 2, 258, 259, 7, 53, 2, 2, 259, 260, 7, 3,
	2, 2, 260, 261, 5, 2, 2, 2, 261, 262, 7, 4, 2, 2, 262, 302, 3, 2, 2, 2,
	263, 264, 7, 54, 2, 2, 264, 265, 7, 3, 2, 2, 265, 266, 5, 2, 2, 2, 266,
	267, 7, 15, 2, 2, 267, 273, 5, 2, 2, 2, 268, 269, 7, 15, 2, 2, 269, 270,
	5, 2, 2, 2, 270, 271, 7, 15, 2, 2, 271, 272, 5, 2, 2, 2, 272, 274, 3, 2,
	2, 2, 273, 268, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2,
	275, 276, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 7, 4, 2, 2, 278,
	302, 3, 2, 2, 2, 279, 280, 7, 55, 2, 2, 280, 281, 7, 3, 2, 2, 281, 282,
	5, 2, 2, 2, 282, 283, 7, 4, 2, 2, 283, 302, 3, 2, 2, 2, 284, 285, 7, 56,
	2, 2, 285, 286, 7, 3, 2, 2, 286, 287, 5, 2, 2, 2, 287, 288, 7, 4, 2, 2,
	288, 302, 3, 2, 2, 2, 289, 302, 7, 68, 2, 2, 290, 291, 7, 8, 2, 2, 291,
	302, 7, 68, 2, 2, 292, 302, 7, 69, 2, 2, 293, 294, 7, 8, 2, 2, 294, 302,
	7, 69, 2, 2, 295, 302, 7, 73, 2, 2, 296, 302, 7, 74, 2, 2, 297, 302, 7,
	71, 2, 2, 298, 302, 7, 70, 2, 2, 299, 302, 7, 67, 2, 2, 300, 302, 7, 72,
	2, 2, 301, 4, 3, 2, 2, 2, 301, 9, 3, 2, 2, 2, 301, 27, 3, 2, 2, 2, 301,
	36, 3, 2, 2, 2, 301, 41, 3, 2, 2, 2, 301, 46, 3, 2, 2, 2, 301, 51, 3, 2,
	2, 2, 301, 56, 3, 2, 2, 2, 301, 63, 3, 2, 2, 2, 301, 75, 3, 2, 2, 2, 301,
	80, 3, 2, 2, 2, 301, 85, 3, 2, 2, 2, 301, 90, 3, 2, 2, 2, 301, 101, 3,
	2, 2, 2, 301, 106, 3, 2, 2, 2, 301, 111, 3, 2, 2, 2, 301, 116, 3, 2, 2,
	2, 301, 127, 3, 2, 2, 2, 301, 138, 3, 2, 2, 2, 301, 143, 3, 2, 2, 2, 301,
	150, 3, 2, 2, 2, 301, 155, 3, 2, 2, 2, 301, 162, 3, 2, 2, 2, 301, 167,
	3, 2, 2, 2, 301, 172, 3, 2, 2, 2, 301, 177, 3, 2, 2, 2, 301, 188, 3, 2,
	2, 2, 301, 200, 3, 2, 2, 2, 301, 211, 3, 2, 2, 2, 301, 223, 3, 2, 2, 2,
	301, 225, 3, 2, 2, 2, 301, 227, 3, 2, 2, 2, 301, 234, 3, 2, 2, 2, 301,
	236, 3, 2, 2, 2, 301, 241, 3, 2, 2, 2, 301, 248, 3, 2, 2, 2, 301, 253,
	3, 2, 2, 2, 301, 258, 3, 2, 2, 2, 301, 263, 3, 2, 2, 2, 301, 279, 3, 2,
	2, 2, 301, 284, 3, 2, 2, 2, 301, 289, 3, 2, 2, 2, 301, 290, 3, 2, 2, 2,
	301, 292, 3, 2, 2, 2, 301, 293, 3, 2, 2, 2, 301, 295, 3, 2, 2, 2, 301,
	296, 3, 2, 2, 2, 301, 297, 3, 2, 2, 2, 301, 298, 3, 2, 2, 2, 301, 299,
	3, 2, 2, 2, 301, 300, 3, 2, 2, 2, 302, 370, 3, 2, 2, 2, 303, 304, 12, 66,
	2, 2, 304, 305, 7, 5, 2, 2, 305, 369, 5, 2, 2, 67, 306, 307, 12, 65, 2,
	2, 307, 308, 7, 6, 2, 2, 308, 369, 5, 2, 2, 66, 309, 310, 12, 64, 2, 2,
	310, 311, 7, 7, 2, 2, 311, 369, 5, 2, 2, 65, 312, 313, 12, 63, 2, 2, 313,
	314, 7, 8, 2, 2, 314, 369, 5, 2, 2, 64, 315, 316, 12, 62, 2, 2, 316, 317,
	7, 9, 2, 2, 317, 369, 5, 2, 2, 63, 318, 319, 12, 61, 2, 2, 319, 320, 7,
	10, 2, 2, 320, 369, 5, 2, 2, 62, 321, 322, 12, 60, 2, 2, 322, 323, 7, 11,
	2, 2, 323, 369, 5, 2, 2, 61, 324, 325, 12, 59, 2, 2, 325, 326, 7, 12, 2,
	2, 326, 369, 5, 2, 2, 60, 327, 328, 12, 58, 2, 2, 328, 329, 7, 13, 2, 2,
	329, 369, 5, 2, 2, 59, 330, 331, 12, 57, 2, 2, 331, 332, 7, 14, 2, 2, 332,
	369, 5, 2, 2, 58, 333, 334, 12, 54, 2, 2, 334, 335, 9, 2, 2, 2, 335, 369,
	5, 2, 2, 55, 336, 337, 12, 53, 2, 2, 337, 338, 9, 3, 2, 2, 338, 369, 5,
	2, 2, 54, 339, 340, 12, 56, 2, 2, 340, 341, 7, 58, 2, 2, 341, 350, 7, 3,
	2, 2, 342, 347, 5, 2, 2, 2, 343, 344, 7, 15, 2, 2, 344, 346, 5, 2, 2, 2,
	345, 343, 3, 2, 2, 2, 346, 349, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 347,
	348, 3, 2, 2, 2, 348, 351, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2, 350, 342,
	3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 369, 7, 4,
	2, 2, 353, 354, 12, 55, 2, 2, 354, 355, 7, 59, 2, 2, 355, 356, 7, 58, 2,
	2, 356, 365, 7, 3, 2, 2, 357, 362, 5, 2, 2, 2, 358, 359, 7, 15, 2, 2, 359,
	361, 5, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 364, 3, 2, 2, 2, 362, 360,
	3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364, 362, 3, 2,
	2, 2, 365, 357, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2,
	367, 369, 7, 4, 2, 2, 368, 303, 3, 2, 2, 2, 368, 306, 3, 2, 2, 2, 368,
	309, 3, 2, 2, 2, 368, 312, 3, 2, 2, 2, 368, 315, 3, 2, 2, 2, 368, 318,
	3, 2, 2, 2, 368, 321, 3, 2, 2, 2, 368, 324, 3, 2, 2, 2, 368, 327, 3, 2,
	2, 2, 368, 330, 3, 2, 2, 2, 368, 333, 3, 2, 2, 2, 368, 336, 3, 2, 2, 2,
	368, 339, 3, 2, 2, 2, 368, 353, 3, 2, 2, 2, 369, 372, 3, 2, 2, 2, 370,
	368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 3, 3, 2, 2, 2, 372, 370, 3,
	2, 2, 2, 18, 20, 70, 97, 134, 184, 195, 207, 218, 275, 301, 347, 350, 362,
	365, 368, 370,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'>='", "'<'",
	"'<='", "'!='", "','", "'&&'", "'||'", "'()'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Abs",
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Log", "Log10", "Max", "Median",
	"Min", "Mod", "Null", "Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh",
	"Sqrt", "Switch", "Tan", "Tanh", "Iif", "In", "Not", "And", "Or", "If",
	"Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal", "Date",
	"Datetime", "Field", "SingleQuoteString", "DoubleQuoteString", "Whitespace",
}

var ruleNames = []string{
	"expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AlteryxFormulasParser struct {
	*antlr.BaseParser
}

func NewAlteryxFormulasParser(input antlr.TokenStream) *AlteryxFormulasParser {
	this := new(AlteryxFormulasParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AlteryxFormulas.g4"

	return this
}

// AlteryxFormulasParser tokens.
const (
	AlteryxFormulasParserEOF               = antlr.TokenEOF
	AlteryxFormulasParserT__0              = 1
	AlteryxFormulasParserT__1              = 2
	AlteryxFormulasParserT__2              = 3
	AlteryxFormulasParserT__3              = 4
	AlteryxFormulasParserT__4              = 5
	AlteryxFormulasParserT__5              = 6
	AlteryxFormulasParserT__6              = 7
	AlteryxFormulasParserT__7              = 8
	AlteryxFormulasParserT__8              = 9
	AlteryxFormulasParserT__9              = 10
	AlteryxFormulasParserT__10             = 11
	AlteryxFormulasParserT__11             = 12
	AlteryxFormulasParserT__12             = 13
	AlteryxFormulasParserT__13             = 14
	AlteryxFormulasParserT__14             = 15
	AlteryxFormulasParserT__15             = 16
	AlteryxFormulasParserAbs               = 17
	AlteryxFormulasParserAcos              = 18
	AlteryxFormulasParserAsin              = 19
	AlteryxFormulasParserAtan              = 20
	AlteryxFormulasParserAtan2             = 21
	AlteryxFormulasParserAverage           = 22
	AlteryxFormulasParserCeil              = 23
	AlteryxFormulasParserCharFromInt       = 24
	AlteryxFormulasParserCharToInt         = 25
	AlteryxFormulasParserContains          = 26
	AlteryxFormulasParserCos               = 27
	AlteryxFormulasParserCosh              = 28
	AlteryxFormulasParserCountWords        = 29
	AlteryxFormulasParserDistance          = 30
	AlteryxFormulasParserEndsWith          = 31
	AlteryxFormulasParserFindString        = 32
	AlteryxFormulasParserExp               = 33
	AlteryxFormulasParserFloor             = 34
	AlteryxFormulasParserGetWord           = 35
	AlteryxFormulasParserHexToNumber       = 36
	AlteryxFormulasParserLog               = 37
	AlteryxFormulasParserLog10             = 38
	AlteryxFormulasParserMax               = 39
	AlteryxFormulasParserMedian            = 40
	AlteryxFormulasParserMin               = 41
	AlteryxFormulasParserMod               = 42
	AlteryxFormulasParserNull              = 43
	AlteryxFormulasParserPi                = 44
	AlteryxFormulasParserPow               = 45
	AlteryxFormulasParserRand              = 46
	AlteryxFormulasParserRandInt           = 47
	AlteryxFormulasParserRound             = 48
	AlteryxFormulasParserSin               = 49
	AlteryxFormulasParserSinh              = 50
	AlteryxFormulasParserSqrt              = 51
	AlteryxFormulasParserSwitch            = 52
	AlteryxFormulasParserTan               = 53
	AlteryxFormulasParserTanh              = 54
	AlteryxFormulasParserIif               = 55
	AlteryxFormulasParserIn                = 56
	AlteryxFormulasParserNot               = 57
	AlteryxFormulasParserAnd               = 58
	AlteryxFormulasParserOr                = 59
	AlteryxFormulasParserIf                = 60
	AlteryxFormulasParserThen              = 61
	AlteryxFormulasParserElse              = 62
	AlteryxFormulasParserElseif            = 63
	AlteryxFormulasParserEndif             = 64
	AlteryxFormulasParserBool              = 65
	AlteryxFormulasParserInteger           = 66
	AlteryxFormulasParserDecimal           = 67
	AlteryxFormulasParserDate              = 68
	AlteryxFormulasParserDatetime          = 69
	AlteryxFormulasParserField             = 70
	AlteryxFormulasParserSingleQuoteString = 71
	AlteryxFormulasParserDoubleQuoteString = 72
	AlteryxFormulasParserWhitespace        = 73
)

// AlteryxFormulasParserRULE_expr is the AlteryxFormulasParser rule.
const AlteryxFormulasParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MaxFuncContext struct {
	*ExprContext
}

func NewMaxFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxFuncContext {
	var p = new(MaxFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MaxFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxFuncContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *MaxFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MaxFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MaxFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMaxFunc(s)
	}
}

func (s *MaxFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMaxFunc(s)
	}
}

type CosFuncContext struct {
	*ExprContext
}

func NewCosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CosFuncContext {
	var p = new(CosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosFuncContext) Cos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCos, 0)
}

func (s *CosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCosFunc(s)
	}
}

func (s *CosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCosFunc(s)
	}
}

type RandIntFuncContext struct {
	*ExprContext
}

func NewRandIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RandIntFuncContext {
	var p = new(RandIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RandIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RandIntFuncContext) RandInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRandInt, 0)
}

func (s *RandIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RandIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRandIntFunc(s)
	}
}

func (s *RandIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRandIntFunc(s)
	}
}

type CharToIntFuncContext struct {
	*ExprContext
}

func NewCharToIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharToIntFuncContext {
	var p = new(CharToIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CharToIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharToIntFuncContext) CharToInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCharToInt, 0)
}

func (s *CharToIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CharToIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCharToIntFunc(s)
	}
}

func (s *CharToIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCharToIntFunc(s)
	}
}

type TanFuncContext struct {
	*ExprContext
}

func NewTanFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanFuncContext {
	var p = new(TanFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TanFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanFuncContext) Tan() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTan, 0)
}

func (s *TanFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TanFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTanFunc(s)
	}
}

func (s *TanFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTanFunc(s)
	}
}

type FloorFuncContext struct {
	*ExprContext
}

func NewFloorFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloorFuncContext {
	var p = new(FloorFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloorFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorFuncContext) Floor() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserFloor, 0)
}

func (s *FloorFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloorFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFloorFunc(s)
	}
}

func (s *FloorFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFloorFunc(s)
	}
}

type NotInContext struct {
	*ExprContext
}

func NewNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInContext {
	var p = new(NotInContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotInContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *NotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *NotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotIn(s)
	}
}

func (s *NotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotIn(s)
	}
}

type TanhFuncContext struct {
	*ExprContext
}

func NewTanhFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanhFuncContext {
	var p = new(TanhFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TanhFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanhFuncContext) Tanh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTanh, 0)
}

func (s *TanhFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TanhFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTanhFunc(s)
	}
}

func (s *TanhFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTanhFunc(s)
	}
}

type GreaterThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterThanContext) GetLeft() IExprContext { return s.left }

func (s *GreaterThanContext) GetRight() IExprContext { return s.right }

func (s *GreaterThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterThanContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

type AsinFuncContext struct {
	*ExprContext
}

func NewAsinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsinFuncContext {
	var p = new(AsinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AsinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsinFuncContext) Asin() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAsin, 0)
}

func (s *AsinFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAsinFunc(s)
	}
}

func (s *AsinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAsinFunc(s)
	}
}

type AddContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddContext) GetLeft() IExprContext { return s.left }

func (s *AddContext) GetRight() IExprContext { return s.right }

func (s *AddContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddContext) SetRight(v IExprContext) { s.right = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAdd(s)
	}
}

type CoshFuncContext struct {
	*ExprContext
}

func NewCoshFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoshFuncContext {
	var p = new(CoshFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CoshFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoshFuncContext) Cosh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCosh, 0)
}

func (s *CoshFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CoshFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCoshFunc(s)
	}
}

func (s *CoshFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCoshFunc(s)
	}
}

type InContext struct {
	*ExprContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIn(s)
	}
}

type SubtractContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewSubtractContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractContext {
	var p = new(SubtractContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubtractContext) GetLeft() IExprContext { return s.left }

func (s *SubtractContext) GetRight() IExprContext { return s.right }

func (s *SubtractContext) SetLeft(v IExprContext) { s.left = v }

func (s *SubtractContext) SetRight(v IExprContext) { s.right = v }

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubtractContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubtractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSubtract(s)
	}
}

func (s *SubtractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSubtract(s)
	}
}

type NotEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqualContext) GetLeft() IExprContext { return s.left }

func (s *NotEqualContext) GetRight() IExprContext { return s.right }

func (s *NotEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *NotEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotEqual(s)
	}
}

func (s *NotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotEqual(s)
	}
}

type ParenthesisContext struct {
	*ExprContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type HexToNumberFuncContext struct {
	*ExprContext
}

func NewHexToNumberFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexToNumberFuncContext {
	var p = new(HexToNumberFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *HexToNumberFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexToNumberFuncContext) HexToNumber() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserHexToNumber, 0)
}

func (s *HexToNumberFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *HexToNumberFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterHexToNumberFunc(s)
	}
}

func (s *HexToNumberFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitHexToNumberFunc(s)
	}
}

type MedianFuncContext struct {
	*ExprContext
}

func NewMedianFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MedianFuncContext {
	var p = new(MedianFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MedianFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MedianFuncContext) Median() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMedian, 0)
}

func (s *MedianFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MedianFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MedianFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMedianFunc(s)
	}
}

func (s *MedianFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMedianFunc(s)
	}
}

type NullFuncContext struct {
	*ExprContext
}

func NewNullFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullFuncContext {
	var p = new(NullFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NullFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullFuncContext) Null() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNull, 0)
}

func (s *NullFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNullFunc(s)
	}
}

func (s *NullFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNullFunc(s)
	}
}

type DatetimeLiteralContext struct {
	*ExprContext
}

func NewDatetimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatetimeLiteralContext {
	var p = new(DatetimeLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DatetimeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeLiteralContext) Datetime() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDatetime, 0)
}

func (s *DatetimeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDatetimeLiteral(s)
	}
}

func (s *DatetimeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDatetimeLiteral(s)
	}
}

type StringLiteralContext struct {
	*ExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) SingleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSingleQuoteString, 0)
}

func (s *StringLiteralContext) DoubleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDoubleQuoteString, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type DateLiteralContext struct {
	*ExprContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DateLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLiteralContext) Date() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDate, 0)
}

func (s *DateLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateLiteral(s)
	}
}

func (s *DateLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateLiteral(s)
	}
}

type AbsFuncContext struct {
	*ExprContext
}

func NewAbsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbsFuncContext {
	var p = new(AbsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AbsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsFuncContext) Abs() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAbs, 0)
}

func (s *AbsFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAbsFunc(s)
	}
}

func (s *AbsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAbsFunc(s)
	}
}

type Atan2FuncContext struct {
	*ExprContext
}

func NewAtan2FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atan2FuncContext {
	var p = new(Atan2FuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Atan2FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atan2FuncContext) Atan2() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan2, 0)
}

func (s *Atan2FuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Atan2FuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Atan2FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtan2Func(s)
	}
}

func (s *Atan2FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtan2Func(s)
	}
}

type FindStringFuncContext struct {
	*ExprContext
}

func NewFindStringFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FindStringFuncContext {
	var p = new(FindStringFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FindStringFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindStringFuncContext) FindString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserFindString, 0)
}

func (s *FindStringFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FindStringFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FindStringFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFindStringFunc(s)
	}
}

func (s *FindStringFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFindStringFunc(s)
	}
}

type Log10FuncContext struct {
	*ExprContext
}

func NewLog10FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Log10FuncContext {
	var p = new(Log10FuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Log10FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log10FuncContext) Log10() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLog10, 0)
}

func (s *Log10FuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log10FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLog10Func(s)
	}
}

func (s *Log10FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLog10Func(s)
	}
}

type CeilFuncContext struct {
	*ExprContext
}

func NewCeilFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilFuncContext {
	var p = new(CeilFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CeilFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilFuncContext) Ceil() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCeil, 0)
}

func (s *CeilFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CeilFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCeilFunc(s)
	}
}

func (s *CeilFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCeilFunc(s)
	}
}

type EndsWithFuncContext struct {
	*ExprContext
}

func NewEndsWithFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EndsWithFuncContext {
	var p = new(EndsWithFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EndsWithFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndsWithFuncContext) EndsWith() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndsWith, 0)
}

func (s *EndsWithFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EndsWithFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EndsWithFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterEndsWithFunc(s)
	}
}

func (s *EndsWithFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitEndsWithFunc(s)
	}
}

type DistanceFuncContext struct {
	*ExprContext
}

func NewDistanceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistanceFuncContext {
	var p = new(DistanceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DistanceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceFuncContext) Distance() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDistance, 0)
}

func (s *DistanceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DistanceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DistanceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDistanceFunc(s)
	}
}

func (s *DistanceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDistanceFunc(s)
	}
}

type CharFromIntFuncContext struct {
	*ExprContext
}

func NewCharFromIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharFromIntFuncContext {
	var p = new(CharFromIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CharFromIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharFromIntFuncContext) CharFromInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCharFromInt, 0)
}

func (s *CharFromIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CharFromIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCharFromIntFunc(s)
	}
}

func (s *CharFromIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCharFromIntFunc(s)
	}
}

type BoolLiteralContext struct {
	*ExprContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) Bool() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserBool, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

type AcosFuncContext struct {
	*ExprContext
}

func NewAcosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AcosFuncContext {
	var p = new(AcosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AcosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcosFuncContext) Acos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAcos, 0)
}

func (s *AcosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AcosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAcosFunc(s)
	}
}

func (s *AcosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAcosFunc(s)
	}
}

type SqrtFuncContext struct {
	*ExprContext
}

func NewSqrtFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrtFuncContext {
	var p = new(SqrtFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SqrtFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtFuncContext) Sqrt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSqrt, 0)
}

func (s *SqrtFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SqrtFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSqrtFunc(s)
	}
}

func (s *SqrtFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSqrtFunc(s)
	}
}

type SinhFuncContext struct {
	*ExprContext
}

func NewSinhFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinhFuncContext {
	var p = new(SinhFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SinhFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinhFuncContext) Sinh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSinh, 0)
}

func (s *SinhFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SinhFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSinhFunc(s)
	}
}

func (s *SinhFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSinhFunc(s)
	}
}

type RandFuncContext struct {
	*ExprContext
}

func NewRandFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RandFuncContext {
	var p = new(RandFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RandFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RandFuncContext) Rand() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRand, 0)
}

func (s *RandFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRandFunc(s)
	}
}

func (s *RandFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRandFunc(s)
	}
}

type MinFuncContext struct {
	*ExprContext
}

func NewMinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinFuncContext {
	var p = new(MinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinFuncContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *MinFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MinFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMinFunc(s)
	}
}

func (s *MinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMinFunc(s)
	}
}

type PiFuncContext struct {
	*ExprContext
}

func NewPiFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PiFuncContext {
	var p = new(PiFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PiFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PiFuncContext) Pi() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPi, 0)
}

func (s *PiFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPiFunc(s)
	}
}

func (s *PiFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPiFunc(s)
	}
}

type AndContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetLeft() IExprContext { return s.left }

func (s *AndContext) GetRight() IExprContext { return s.right }

func (s *AndContext) SetLeft(v IExprContext) { s.left = v }

func (s *AndContext) SetRight(v IExprContext) { s.right = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) And() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAnd, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAnd(s)
	}
}

type LessThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessThanContext) GetLeft() IExprContext { return s.left }

func (s *LessThanContext) GetRight() IExprContext { return s.right }

func (s *LessThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessThanContext) SetRight(v IExprContext) { s.right = v }

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessThan(s)
	}
}

type DivideContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivideContext) GetLeft() IExprContext { return s.left }

func (s *DivideContext) GetRight() IExprContext { return s.right }

func (s *DivideContext) SetLeft(v IExprContext) { s.left = v }

func (s *DivideContext) SetRight(v IExprContext) { s.right = v }

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivideContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDivide(s)
	}
}

type ExprIfContext struct {
	*ExprContext
	ifStmnt   IExprContext
	thenStmnt IExprContext
	elseStmnt IExprContext
}

func NewExprIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIfContext {
	var p = new(ExprIfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprIfContext) GetIfStmnt() IExprContext { return s.ifStmnt }

func (s *ExprIfContext) GetThenStmnt() IExprContext { return s.thenStmnt }

func (s *ExprIfContext) GetElseStmnt() IExprContext { return s.elseStmnt }

func (s *ExprIfContext) SetIfStmnt(v IExprContext) { s.ifStmnt = v }

func (s *ExprIfContext) SetThenStmnt(v IExprContext) { s.thenStmnt = v }

func (s *ExprIfContext) SetElseStmnt(v IExprContext) { s.elseStmnt = v }

func (s *ExprIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *ExprIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *ExprIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *ExprIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *ExprIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *ExprIfContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprIfContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *ExprIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *ExprIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprIf(s)
	}
}

func (s *ExprIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprIf(s)
	}
}

type ContainsFuncContext struct {
	*ExprContext
}

func NewContainsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsFuncContext {
	var p = new(ContainsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ContainsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsFuncContext) Contains() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserContains, 0)
}

func (s *ContainsFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ContainsFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ContainsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterContainsFunc(s)
	}
}

func (s *ContainsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitContainsFunc(s)
	}
}

type MultiplyContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultiplyContext) GetLeft() IExprContext { return s.left }

func (s *MultiplyContext) GetRight() IExprContext { return s.right }

func (s *MultiplyContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultiplyContext) SetRight(v IExprContext) { s.right = v }

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultiplyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMultiply(s)
	}
}

func (s *MultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMultiply(s)
	}
}

type AtanFuncContext struct {
	*ExprContext
}

func NewAtanFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtanFuncContext {
	var p = new(AtanFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtanFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtanFuncContext) Atan() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan, 0)
}

func (s *AtanFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtanFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtanFunc(s)
	}
}

func (s *AtanFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtanFunc(s)
	}
}

type RoundFuncContext struct {
	*ExprContext
}

func NewRoundFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundFuncContext {
	var p = new(RoundFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RoundFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundFuncContext) Round() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRound, 0)
}

func (s *RoundFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RoundFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RoundFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRoundFunc(s)
	}
}

func (s *RoundFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRoundFunc(s)
	}
}

type SwitchFuncContext struct {
	*ExprContext
}

func NewSwitchFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchFuncContext {
	var p = new(SwitchFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SwitchFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchFuncContext) Switch() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSwitch, 0)
}

func (s *SwitchFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SwitchFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSwitchFunc(s)
	}
}

func (s *SwitchFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSwitchFunc(s)
	}
}

type ExprFieldContext struct {
	*ExprContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprField(s)
	}
}

type GetWordFuncContext struct {
	*ExprContext
}

func NewGetWordFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetWordFuncContext {
	var p = new(GetWordFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetWordFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetWordFuncContext) GetWord() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserGetWord, 0)
}

func (s *GetWordFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GetWordFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GetWordFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGetWordFunc(s)
	}
}

func (s *GetWordFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGetWordFunc(s)
	}
}

type OrContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetLeft() IExprContext { return s.left }

func (s *OrContext) GetRight() IExprContext { return s.right }

func (s *OrContext) SetLeft(v IExprContext) { s.left = v }

func (s *OrContext) SetRight(v IExprContext) { s.right = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) Or() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserOr, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitOr(s)
	}
}

type ExpFuncContext struct {
	*ExprContext
}

func NewExpFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFuncContext {
	var p = new(ExpFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFuncContext) Exp() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserExp, 0)
}

func (s *ExpFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExpFunc(s)
	}
}

func (s *ExpFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExpFunc(s)
	}
}

type CountWordsFuncContext struct {
	*ExprContext
}

func NewCountWordsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountWordsFuncContext {
	var p = new(CountWordsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CountWordsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountWordsFuncContext) CountWords() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCountWords, 0)
}

func (s *CountWordsFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CountWordsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCountWordsFunc(s)
	}
}

func (s *CountWordsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCountWordsFunc(s)
	}
}

type PowFuncContext struct {
	*ExprContext
}

func NewPowFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowFuncContext {
	var p = new(PowFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowFuncContext) Pow() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPow, 0)
}

func (s *PowFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPowFunc(s)
	}
}

func (s *PowFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPowFunc(s)
	}
}

type IifFuncContext struct {
	*ExprContext
}

func NewIifFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IifFuncContext {
	var p = new(IifFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IifFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IifFuncContext) Iif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIif, 0)
}

func (s *IifFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IifFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IifFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIifFunc(s)
	}
}

func (s *IifFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIifFunc(s)
	}
}

type EqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualContext) GetLeft() IExprContext { return s.left }

func (s *EqualContext) GetRight() IExprContext { return s.right }

func (s *EqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitEqual(s)
	}
}

type ModFuncContext struct {
	*ExprContext
}

func NewModFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModFuncContext {
	var p = new(ModFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ModFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModFuncContext) Mod() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMod, 0)
}

func (s *ModFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ModFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterModFunc(s)
	}
}

func (s *ModFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitModFunc(s)
	}
}

type SinFuncContext struct {
	*ExprContext
}

func NewSinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinFuncContext {
	var p = new(SinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinFuncContext) Sin() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSin, 0)
}

func (s *SinFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSinFunc(s)
	}
}

func (s *SinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSinFunc(s)
	}
}

type AverageFuncContext struct {
	*ExprContext
}

func NewAverageFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AverageFuncContext {
	var p = new(AverageFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AverageFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AverageFuncContext) Average() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAverage, 0)
}

func (s *AverageFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AverageFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AverageFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAverageFunc(s)
	}
}

func (s *AverageFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAverageFunc(s)
	}
}

type GreaterEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterEqualContext {
	var p = new(GreaterEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterEqualContext) GetLeft() IExprContext { return s.left }

func (s *GreaterEqualContext) GetRight() IExprContext { return s.right }

func (s *GreaterEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterEqual(s)
	}
}

func (s *GreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterEqual(s)
	}
}

type LessEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessEqualContext {
	var p = new(LessEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessEqualContext) GetLeft() IExprContext { return s.left }

func (s *LessEqualContext) GetRight() IExprContext { return s.right }

func (s *LessEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *LessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessEqual(s)
	}
}

func (s *LessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessEqual(s)
	}
}

type LogFuncContext struct {
	*ExprContext
}

func NewLogFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogFuncContext {
	var p = new(LogFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogFuncContext) Log() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLog, 0)
}

func (s *LogFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLogFunc(s)
	}
}

func (s *LogFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLogFunc(s)
	}
}

type NumberLiteralContext struct {
	*ExprContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) Integer() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserInteger, 0)
}

func (s *NumberLiteralContext) Decimal() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDecimal, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *AlteryxFormulasParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *AlteryxFormulasParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, AlteryxFormulasParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(4)
			p.expr(0)
		}
		{
			p.SetState(5)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewExprIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(7)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(8)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).ifStmnt = _x
		}
		{
			p.SetState(9)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(10)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).thenStmnt = _x
		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserElseif {
			{
				p.SetState(11)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(12)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).ifStmnt = _x
			}
			{
				p.SetState(13)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(14)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).thenStmnt = _x
			}

			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(21)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(22)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).elseStmnt = _x
		}
		{
			p.SetState(23)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 3:
		localctx = NewIifFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(AlteryxFormulasParserIif)
		}
		{
			p.SetState(26)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(27)
			p.expr(0)
		}
		{
			p.SetState(28)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(29)
			p.expr(0)
		}
		{
			p.SetState(30)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(31)
			p.expr(0)
		}
		{
			p.SetState(32)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 4:
		localctx = NewAbsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(AlteryxFormulasParserAbs)
		}
		{
			p.SetState(35)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(36)
			p.expr(0)
		}
		{
			p.SetState(37)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 5:
		localctx = NewAcosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			p.Match(AlteryxFormulasParserAcos)
		}
		{
			p.SetState(40)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(41)
			p.expr(0)
		}
		{
			p.SetState(42)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 6:
		localctx = NewAsinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(AlteryxFormulasParserAsin)
		}
		{
			p.SetState(45)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(46)
			p.expr(0)
		}
		{
			p.SetState(47)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 7:
		localctx = NewAtanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(AlteryxFormulasParserAtan)
		}
		{
			p.SetState(50)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(51)
			p.expr(0)
		}
		{
			p.SetState(52)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 8:
		localctx = NewAtan2FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(AlteryxFormulasParserAtan2)
		}
		{
			p.SetState(55)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(56)
			p.expr(0)
		}
		{
			p.SetState(57)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		{
			p.SetState(59)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 9:
		localctx = NewAverageFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(AlteryxFormulasParserAverage)
		}
		{
			p.SetState(62)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(63)
			p.expr(0)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(64)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(65)
				p.expr(0)
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 10:
		localctx = NewCeilFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(AlteryxFormulasParserCeil)
		}
		{
			p.SetState(74)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(75)
			p.expr(0)
		}
		{
			p.SetState(76)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 11:
		localctx = NewCharFromIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(AlteryxFormulasParserCharFromInt)
		}
		{
			p.SetState(79)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(80)
			p.expr(0)
		}
		{
			p.SetState(81)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 12:
		localctx = NewCharToIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(AlteryxFormulasParserCharToInt)
		}
		{
			p.SetState(84)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(85)
			p.expr(0)
		}
		{
			p.SetState(86)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 13:
		localctx = NewContainsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(AlteryxFormulasParserContains)
		}
		{
			p.SetState(89)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(90)
			p.expr(0)
		}
		{
			p.SetState(91)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(92)
			p.expr(0)
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(93)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(94)
				p.expr(0)
			}

		}
		{
			p.SetState(97)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 14:
		localctx = NewCosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(AlteryxFormulasParserCos)
		}
		{
			p.SetState(100)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(101)
			p.expr(0)
		}
		{
			p.SetState(102)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 15:
		localctx = NewCoshFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.Match(AlteryxFormulasParserCosh)
		}
		{
			p.SetState(105)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(106)
			p.expr(0)
		}
		{
			p.SetState(107)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 16:
		localctx = NewCountWordsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.Match(AlteryxFormulasParserCountWords)
		}
		{
			p.SetState(110)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(111)
			p.expr(0)
		}
		{
			p.SetState(112)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 17:
		localctx = NewDistanceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(AlteryxFormulasParserDistance)
		}
		{
			p.SetState(115)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(116)
			p.expr(0)
		}
		{
			p.SetState(117)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(118)
			p.expr(0)
		}
		{
			p.SetState(119)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(120)
			p.expr(0)
		}
		{
			p.SetState(121)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 18:
		localctx = NewEndsWithFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(125)
			p.Match(AlteryxFormulasParserEndsWith)
		}
		{
			p.SetState(126)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(127)
			p.expr(0)
		}
		{
			p.SetState(128)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(129)
			p.expr(0)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(130)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(131)
				p.expr(0)
			}

		}
		{
			p.SetState(134)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 19:
		localctx = NewExpFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Match(AlteryxFormulasParserExp)
		}
		{
			p.SetState(137)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 20:
		localctx = NewFindStringFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(AlteryxFormulasParserFindString)
		}
		{
			p.SetState(142)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(143)
			p.expr(0)
		}
		{
			p.SetState(144)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(145)
			p.expr(0)
		}
		{
			p.SetState(146)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 21:
		localctx = NewFloorFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.Match(AlteryxFormulasParserFloor)
		}
		{
			p.SetState(149)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(150)
			p.expr(0)
		}
		{
			p.SetState(151)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 22:
		localctx = NewGetWordFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(AlteryxFormulasParserGetWord)
		}
		{
			p.SetState(154)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(155)
			p.expr(0)
		}
		{
			p.SetState(156)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(157)
			p.expr(0)
		}
		{
			p.SetState(158)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 23:
		localctx = NewHexToNumberFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(AlteryxFormulasParserHexToNumber)
		}
		{
			p.SetState(161)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(162)
			p.expr(0)
		}
		{
			p.SetState(163)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 24:
		localctx = NewLogFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(AlteryxFormulasParserLog)
		}
		{
			p.SetState(166)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(167)
			p.expr(0)
		}
		{
			p.SetState(168)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 25:
		localctx = NewLog10FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.Match(AlteryxFormulasParserLog10)
		}
		{
			p.SetState(171)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(172)
			p.expr(0)
		}
		{
			p.SetState(173)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 26:
		localctx = NewMaxFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(175)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(176)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(177)
			p.expr(0)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(178)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(179)
				p.expr(0)
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 27:
		localctx = NewMedianFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(186)
			p.Match(AlteryxFormulasParserMedian)
		}
		{
			p.SetState(187)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(188)
			p.expr(0)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(189)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(190)
				p.expr(0)
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(196)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 28:
		localctx = NewMinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(198)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(199)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(200)
			p.expr(0)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(201)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(202)
				p.expr(0)
			}

			p.SetState(205)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(207)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 29:
		localctx = NewModFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(209)
			p.Match(AlteryxFormulasParserMod)
		}
		{
			p.SetState(210)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(211)
			p.expr(0)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(212)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(213)
				p.expr(0)
			}

			p.SetState(218)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(219)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 30:
		localctx = NewNullFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(221)
			p.Match(AlteryxFormulasParserNull)
		}
		{
			p.SetState(222)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 31:
		localctx = NewPiFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(223)
			p.Match(AlteryxFormulasParserPi)
		}
		{
			p.SetState(224)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 32:
		localctx = NewPowFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(225)
			p.Match(AlteryxFormulasParserPow)
		}
		{
			p.SetState(226)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(227)
			p.expr(0)
		}
		{
			p.SetState(228)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(229)
			p.expr(0)
		}
		{
			p.SetState(230)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 33:
		localctx = NewRandFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(232)
			p.Match(AlteryxFormulasParserRand)
		}
		{
			p.SetState(233)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 34:
		localctx = NewRandIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(234)
			p.Match(AlteryxFormulasParserRandInt)
		}
		{
			p.SetState(235)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(236)
			p.expr(0)
		}
		{
			p.SetState(237)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 35:
		localctx = NewRoundFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(239)
			p.Match(AlteryxFormulasParserRound)
		}
		{
			p.SetState(240)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(241)
			p.expr(0)
		}
		{
			p.SetState(242)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(243)
			p.expr(0)
		}
		{
			p.SetState(244)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 36:
		localctx = NewSinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(246)
			p.Match(AlteryxFormulasParserSin)
		}
		{
			p.SetState(247)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(248)
			p.expr(0)
		}
		{
			p.SetState(249)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 37:
		localctx = NewSinhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(251)
			p.Match(AlteryxFormulasParserSinh)
		}
		{
			p.SetState(252)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(253)
			p.expr(0)
		}
		{
			p.SetState(254)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 38:
		localctx = NewSqrtFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)
			p.Match(AlteryxFormulasParserSqrt)
		}
		{
			p.SetState(257)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(258)
			p.expr(0)
		}
		{
			p.SetState(259)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 39:
		localctx = NewSwitchFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(261)
			p.Match(AlteryxFormulasParserSwitch)
		}
		{
			p.SetState(262)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(263)
			p.expr(0)
		}
		{
			p.SetState(264)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(265)
			p.expr(0)
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(266)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(267)
				p.expr(0)
			}
			{
				p.SetState(268)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(269)
				p.expr(0)
			}

			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(275)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 40:
		localctx = NewTanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.Match(AlteryxFormulasParserTan)
		}
		{
			p.SetState(278)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(279)
			p.expr(0)
		}
		{
			p.SetState(280)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 41:
		localctx = NewTanhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(282)
			p.Match(AlteryxFormulasParserTanh)
		}
		{
			p.SetState(283)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(284)
			p.expr(0)
		}
		{
			p.SetState(285)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 42:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(287)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 43:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(288)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(289)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 44:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(290)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 45:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(291)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(292)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 46:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(293)
			p.Match(AlteryxFormulasParserSingleQuoteString)
		}

	case 47:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(294)
			p.Match(AlteryxFormulasParserDoubleQuoteString)
		}

	case 48:
		localctx = NewDatetimeLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(295)
			p.Match(AlteryxFormulasParserDatetime)
		}

	case 49:
		localctx = NewDateLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(296)
			p.Match(AlteryxFormulasParserDate)
		}

	case 50:
		localctx = NewBoolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(297)
			p.Match(AlteryxFormulasParserBool)
		}

	case 51:
		localctx = NewExprFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(298)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(368)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplyContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(301)

				if !(p.Precpred(p.GetParserRuleContext(), 64)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 64)", ""))
				}
				{
					p.SetState(302)
					p.Match(AlteryxFormulasParserT__2)
				}
				{
					p.SetState(303)

					var _x = p.expr(65)

					localctx.(*MultiplyContext).right = _x
				}

			case 2:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivideContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(304)

				if !(p.Precpred(p.GetParserRuleContext(), 63)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 63)", ""))
				}
				{
					p.SetState(305)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(306)

					var _x = p.expr(64)

					localctx.(*DivideContext).right = _x
				}

			case 3:
				localctx = NewAddContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 62)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 62)", ""))
				}
				{
					p.SetState(308)
					p.Match(AlteryxFormulasParserT__4)
				}
				{
					p.SetState(309)

					var _x = p.expr(63)

					localctx.(*AddContext).right = _x
				}

			case 4:
				localctx = NewSubtractContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubtractContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(310)

				if !(p.Precpred(p.GetParserRuleContext(), 61)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 61)", ""))
				}
				{
					p.SetState(311)
					p.Match(AlteryxFormulasParserT__5)
				}
				{
					p.SetState(312)

					var _x = p.expr(62)

					localctx.(*SubtractContext).right = _x
				}

			case 5:
				localctx = NewEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(313)

				if !(p.Precpred(p.GetParserRuleContext(), 60)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 60)", ""))
				}
				{
					p.SetState(314)
					p.Match(AlteryxFormulasParserT__6)
				}
				{
					p.SetState(315)

					var _x = p.expr(61)

					localctx.(*EqualContext).right = _x
				}

			case 6:
				localctx = NewGreaterThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(316)

				if !(p.Precpred(p.GetParserRuleContext(), 59)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 59)", ""))
				}
				{
					p.SetState(317)
					p.Match(AlteryxFormulasParserT__7)
				}
				{
					p.SetState(318)

					var _x = p.expr(60)

					localctx.(*GreaterThanContext).right = _x
				}

			case 7:
				localctx = NewGreaterEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 58)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 58)", ""))
				}
				{
					p.SetState(320)
					p.Match(AlteryxFormulasParserT__8)
				}
				{
					p.SetState(321)

					var _x = p.expr(59)

					localctx.(*GreaterEqualContext).right = _x
				}

			case 8:
				localctx = NewLessThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 57)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 57)", ""))
				}
				{
					p.SetState(323)
					p.Match(AlteryxFormulasParserT__9)
				}
				{
					p.SetState(324)

					var _x = p.expr(58)

					localctx.(*LessThanContext).right = _x
				}

			case 9:
				localctx = NewLessEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(325)

				if !(p.Precpred(p.GetParserRuleContext(), 56)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 56)", ""))
				}
				{
					p.SetState(326)
					p.Match(AlteryxFormulasParserT__10)
				}
				{
					p.SetState(327)

					var _x = p.expr(57)

					localctx.(*LessEqualContext).right = _x
				}

			case 10:
				localctx = NewNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 55)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 55)", ""))
				}
				{
					p.SetState(329)
					p.Match(AlteryxFormulasParserT__11)
				}
				{
					p.SetState(330)

					var _x = p.expr(56)

					localctx.(*NotEqualContext).right = _x
				}

			case 11:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 52)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 52)", ""))
				}
				{
					p.SetState(332)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__13 || _la == AlteryxFormulasParserAnd) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(333)

					var _x = p.expr(53)

					localctx.(*AndContext).right = _x
				}

			case 12:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 51)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 51)", ""))
				}
				{
					p.SetState(335)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__14 || _la == AlteryxFormulasParserOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(336)

					var _x = p.expr(52)

					localctx.(*OrContext).right = _x
				}

			case 13:
				localctx = NewInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(337)

				if !(p.Precpred(p.GetParserRuleContext(), 54)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 54)", ""))
				}
				{
					p.SetState(338)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(339)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(348)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32))|(1<<(AlteryxFormulasParserSwitch-32))|(1<<(AlteryxFormulasParserTan-32))|(1<<(AlteryxFormulasParserTanh-32))|(1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIf-32)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(AlteryxFormulasParserBool-65))|(1<<(AlteryxFormulasParserInteger-65))|(1<<(AlteryxFormulasParserDecimal-65))|(1<<(AlteryxFormulasParserDate-65))|(1<<(AlteryxFormulasParserDatetime-65))|(1<<(AlteryxFormulasParserField-65))|(1<<(AlteryxFormulasParserSingleQuoteString-65))|(1<<(AlteryxFormulasParserDoubleQuoteString-65)))) != 0) {
					{
						p.SetState(340)
						p.expr(0)
					}
					p.SetState(345)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(341)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(342)
							p.expr(0)
						}

						p.SetState(347)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(350)
					p.Match(AlteryxFormulasParserT__1)
				}

			case 14:
				localctx = NewNotInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 53)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 53)", ""))
				}
				{
					p.SetState(352)
					p.Match(AlteryxFormulasParserNot)
				}
				{
					p.SetState(353)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(354)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(363)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32))|(1<<(AlteryxFormulasParserSwitch-32))|(1<<(AlteryxFormulasParserTan-32))|(1<<(AlteryxFormulasParserTanh-32))|(1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIf-32)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(AlteryxFormulasParserBool-65))|(1<<(AlteryxFormulasParserInteger-65))|(1<<(AlteryxFormulasParserDecimal-65))|(1<<(AlteryxFormulasParserDate-65))|(1<<(AlteryxFormulasParserDatetime-65))|(1<<(AlteryxFormulasParserField-65))|(1<<(AlteryxFormulasParserSingleQuoteString-65))|(1<<(AlteryxFormulasParserDoubleQuoteString-65)))) != 0) {
					{
						p.SetState(355)
						p.expr(0)
					}
					p.SetState(360)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(356)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(357)
							p.expr(0)
						}

						p.SetState(362)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(365)
					p.Match(AlteryxFormulasParserT__1)
				}

			}

		}
		p.SetState(370)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

func (p *AlteryxFormulasParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlteryxFormulasParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 64)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 63)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 62)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 61)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 60)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 59)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 58)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 57)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 56)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 55)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 52)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 51)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 54)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 53)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
