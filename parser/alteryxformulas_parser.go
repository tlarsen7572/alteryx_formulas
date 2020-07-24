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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 506,
	4, 2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2,
	60, 10, 2, 12, 2, 14, 2, 63, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 89, 10, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 126, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 205, 10, 2, 13, 2, 14, 2, 206, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 216, 10, 2, 12, 2, 14, 2, 219, 11,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 228, 10, 2, 13, 2, 14,
	2, 229, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 239, 10, 2, 12,
	2, 14, 2, 242, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 289, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 5, 2, 300, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 5, 2, 313, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 374, 10, 2, 13, 2, 14, 2, 375, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 395, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 404, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5,
	2, 413, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 434, 10,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 478, 10, 2, 12, 2, 14, 2,
	481, 11, 2, 5, 2, 483, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 7, 2, 493, 10, 2, 12, 2, 14, 2, 496, 11, 2, 5, 2, 498, 10, 2, 3,
	2, 7, 2, 501, 10, 2, 12, 2, 14, 2, 504, 11, 2, 3, 2, 2, 3, 2, 3, 2, 2,
	4, 4, 2, 16, 16, 76, 76, 4, 2, 17, 17, 77, 77, 2, 603, 2, 433, 3, 2, 2,
	2, 4, 5, 8, 2, 1, 2, 5, 6, 7, 3, 2, 2, 6, 7, 5, 2, 2, 2, 7, 8, 7, 4, 2,
	2, 8, 434, 3, 2, 2, 2, 9, 10, 7, 78, 2, 2, 10, 11, 5, 2, 2, 2, 11, 12,
	7, 79, 2, 2, 12, 20, 5, 2, 2, 2, 13, 14, 7, 81, 2, 2, 14, 15, 5, 2, 2,
	2, 15, 16, 7, 79, 2, 2, 16, 17, 5, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 13,
	3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2,
	21, 23, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7, 80, 2, 2, 24, 25, 5,
	2, 2, 2, 25, 26, 7, 82, 2, 2, 26, 434, 3, 2, 2, 2, 27, 28, 7, 19, 2, 2,
	28, 29, 7, 3, 2, 2, 29, 30, 5, 2, 2, 2, 30, 31, 7, 4, 2, 2, 31, 434, 3,
	2, 2, 2, 32, 33, 7, 20, 2, 2, 33, 34, 7, 3, 2, 2, 34, 35, 5, 2, 2, 2, 35,
	36, 7, 4, 2, 2, 36, 434, 3, 2, 2, 2, 37, 38, 7, 21, 2, 2, 38, 39, 7, 3,
	2, 2, 39, 40, 5, 2, 2, 2, 40, 41, 7, 4, 2, 2, 41, 434, 3, 2, 2, 2, 42,
	43, 7, 22, 2, 2, 43, 44, 7, 3, 2, 2, 44, 45, 5, 2, 2, 2, 45, 46, 7, 4,
	2, 2, 46, 434, 3, 2, 2, 2, 47, 48, 7, 23, 2, 2, 48, 49, 7, 3, 2, 2, 49,
	50, 5, 2, 2, 2, 50, 51, 7, 15, 2, 2, 51, 52, 5, 2, 2, 2, 52, 53, 7, 4,
	2, 2, 53, 434, 3, 2, 2, 2, 54, 55, 7, 24, 2, 2, 55, 56, 7, 3, 2, 2, 56,
	61, 5, 2, 2, 2, 57, 58, 7, 15, 2, 2, 58, 60, 5, 2, 2, 2, 59, 57, 3, 2,
	2, 2, 60, 63, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 64,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 64, 65, 7, 4, 2, 2, 65, 434, 3, 2, 2, 2,
	66, 67, 7, 25, 2, 2, 67, 68, 7, 3, 2, 2, 68, 69, 5, 2, 2, 2, 69, 70, 7,
	4, 2, 2, 70, 434, 3, 2, 2, 2, 71, 72, 7, 26, 2, 2, 72, 73, 7, 3, 2, 2,
	73, 74, 5, 2, 2, 2, 74, 75, 7, 4, 2, 2, 75, 434, 3, 2, 2, 2, 76, 77, 7,
	27, 2, 2, 77, 78, 7, 3, 2, 2, 78, 79, 5, 2, 2, 2, 79, 80, 7, 4, 2, 2, 80,
	434, 3, 2, 2, 2, 81, 82, 7, 28, 2, 2, 82, 83, 7, 3, 2, 2, 83, 84, 5, 2,
	2, 2, 84, 85, 7, 15, 2, 2, 85, 88, 5, 2, 2, 2, 86, 87, 7, 15, 2, 2, 87,
	89, 5, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 3, 2, 2,
	2, 90, 91, 7, 4, 2, 2, 91, 434, 3, 2, 2, 2, 92, 93, 7, 29, 2, 2, 93, 94,
	7, 3, 2, 2, 94, 95, 5, 2, 2, 2, 95, 96, 7, 4, 2, 2, 96, 434, 3, 2, 2, 2,
	97, 98, 7, 30, 2, 2, 98, 99, 7, 3, 2, 2, 99, 100, 5, 2, 2, 2, 100, 101,
	7, 4, 2, 2, 101, 434, 3, 2, 2, 2, 102, 103, 7, 31, 2, 2, 103, 104, 7, 3,
	2, 2, 104, 105, 5, 2, 2, 2, 105, 106, 7, 4, 2, 2, 106, 434, 3, 2, 2, 2,
	107, 108, 7, 32, 2, 2, 108, 109, 7, 3, 2, 2, 109, 110, 5, 2, 2, 2, 110,
	111, 7, 15, 2, 2, 111, 112, 5, 2, 2, 2, 112, 113, 7, 15, 2, 2, 113, 114,
	5, 2, 2, 2, 114, 115, 7, 15, 2, 2, 115, 116, 5, 2, 2, 2, 116, 117, 7, 4,
	2, 2, 117, 434, 3, 2, 2, 2, 118, 119, 7, 33, 2, 2, 119, 120, 7, 3, 2, 2,
	120, 121, 5, 2, 2, 2, 121, 122, 7, 15, 2, 2, 122, 125, 5, 2, 2, 2, 123,
	124, 7, 15, 2, 2, 124, 126, 5, 2, 2, 2, 125, 123, 3, 2, 2, 2, 125, 126,
	3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 128, 7, 4, 2, 2, 128, 434, 3, 2,
	2, 2, 129, 130, 7, 35, 2, 2, 130, 131, 7, 3, 2, 2, 131, 132, 5, 2, 2, 2,
	132, 133, 7, 4, 2, 2, 133, 434, 3, 2, 2, 2, 134, 135, 7, 34, 2, 2, 135,
	136, 7, 3, 2, 2, 136, 137, 5, 2, 2, 2, 137, 138, 7, 15, 2, 2, 138, 139,
	5, 2, 2, 2, 139, 140, 7, 4, 2, 2, 140, 434, 3, 2, 2, 2, 141, 142, 7, 36,
	2, 2, 142, 143, 7, 3, 2, 2, 143, 144, 5, 2, 2, 2, 144, 145, 7, 4, 2, 2,
	145, 434, 3, 2, 2, 2, 146, 147, 7, 37, 2, 2, 147, 148, 7, 3, 2, 2, 148,
	149, 5, 2, 2, 2, 149, 150, 7, 15, 2, 2, 150, 151, 5, 2, 2, 2, 151, 152,
	7, 4, 2, 2, 152, 434, 3, 2, 2, 2, 153, 154, 7, 38, 2, 2, 154, 155, 7, 3,
	2, 2, 155, 156, 5, 2, 2, 2, 156, 157, 7, 4, 2, 2, 157, 434, 3, 2, 2, 2,
	158, 159, 7, 39, 2, 2, 159, 160, 7, 3, 2, 2, 160, 161, 5, 2, 2, 2, 161,
	162, 7, 15, 2, 2, 162, 163, 5, 2, 2, 2, 163, 164, 7, 15, 2, 2, 164, 165,
	5, 2, 2, 2, 165, 166, 7, 4, 2, 2, 166, 434, 3, 2, 2, 2, 167, 168, 7, 40,
	2, 2, 168, 169, 7, 3, 2, 2, 169, 170, 5, 2, 2, 2, 170, 171, 7, 4, 2, 2,
	171, 434, 3, 2, 2, 2, 172, 173, 7, 41, 2, 2, 173, 174, 7, 3, 2, 2, 174,
	175, 5, 2, 2, 2, 175, 176, 7, 15, 2, 2, 176, 177, 5, 2, 2, 2, 177, 178,
	7, 4, 2, 2, 178, 434, 3, 2, 2, 2, 179, 180, 7, 42, 2, 2, 180, 181, 7, 3,
	2, 2, 181, 182, 5, 2, 2, 2, 182, 183, 7, 4, 2, 2, 183, 434, 3, 2, 2, 2,
	184, 185, 7, 43, 2, 2, 185, 186, 7, 3, 2, 2, 186, 187, 5, 2, 2, 2, 187,
	188, 7, 4, 2, 2, 188, 434, 3, 2, 2, 2, 189, 190, 7, 44, 2, 2, 190, 191,
	7, 3, 2, 2, 191, 192, 5, 2, 2, 2, 192, 193, 7, 4, 2, 2, 193, 434, 3, 2,
	2, 2, 194, 195, 7, 45, 2, 2, 195, 196, 7, 3, 2, 2, 196, 197, 5, 2, 2, 2,
	197, 198, 7, 4, 2, 2, 198, 434, 3, 2, 2, 2, 199, 200, 7, 46, 2, 2, 200,
	201, 7, 3, 2, 2, 201, 204, 5, 2, 2, 2, 202, 203, 7, 15, 2, 2, 203, 205,
	5, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206, 3, 2, 2, 2, 206, 204, 3, 2,
	2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 7, 4, 2, 2,
	209, 434, 3, 2, 2, 2, 210, 211, 7, 47, 2, 2, 211, 212, 7, 3, 2, 2, 212,
	217, 5, 2, 2, 2, 213, 214, 7, 15, 2, 2, 214, 216, 5, 2, 2, 2, 215, 213,
	3, 2, 2, 2, 216, 219, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2,
	2, 2, 218, 220, 3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 221, 7, 4, 2, 2,
	221, 434, 3, 2, 2, 2, 222, 223, 7, 48, 2, 2, 223, 224, 7, 3, 2, 2, 224,
	227, 5, 2, 2, 2, 225, 226, 7, 15, 2, 2, 226, 228, 5, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2,
	2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 7, 4, 2, 2, 232, 434, 3, 2, 2, 2,
	233, 234, 7, 49, 2, 2, 234, 235, 7, 3, 2, 2, 235, 240, 5, 2, 2, 2, 236,
	237, 7, 15, 2, 2, 237, 239, 5, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 242,
	3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243, 3, 2,
	2, 2, 242, 240, 3, 2, 2, 2, 243, 244, 7, 4, 2, 2, 244, 434, 3, 2, 2, 2,
	245, 246, 7, 50, 2, 2, 246, 434, 7, 18, 2, 2, 247, 248, 7, 51, 2, 2, 248,
	249, 7, 3, 2, 2, 249, 250, 5, 2, 2, 2, 250, 251, 7, 15, 2, 2, 251, 252,
	5, 2, 2, 2, 252, 253, 7, 15, 2, 2, 253, 254, 5, 2, 2, 2, 254, 255, 7, 4,
	2, 2, 255, 434, 3, 2, 2, 2, 256, 257, 7, 52, 2, 2, 257, 258, 7, 3, 2, 2,
	258, 259, 5, 2, 2, 2, 259, 260, 7, 15, 2, 2, 260, 261, 5, 2, 2, 2, 261,
	262, 7, 15, 2, 2, 262, 263, 5, 2, 2, 2, 263, 264, 7, 4, 2, 2, 264, 434,
	3, 2, 2, 2, 265, 266, 7, 53, 2, 2, 266, 434, 7, 18, 2, 2, 267, 268, 7,
	54, 2, 2, 268, 269, 7, 3, 2, 2, 269, 270, 5, 2, 2, 2, 270, 271, 7, 15,
	2, 2, 271, 272, 5, 2, 2, 2, 272, 273, 7, 4, 2, 2, 273, 434, 3, 2, 2, 2,
	274, 275, 7, 55, 2, 2, 275, 434, 7, 18, 2, 2, 276, 277, 7, 56, 2, 2, 277,
	278, 7, 3, 2, 2, 278, 279, 5, 2, 2, 2, 279, 280, 7, 4, 2, 2, 280, 434,
	3, 2, 2, 2, 281, 282, 7, 57, 2, 2, 282, 283, 7, 3, 2, 2, 283, 284, 5, 2,
	2, 2, 284, 285, 7, 15, 2, 2, 285, 288, 5, 2, 2, 2, 286, 287, 7, 15, 2,
	2, 287, 289, 5, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289,
	290, 3, 2, 2, 2, 290, 291, 7, 4, 2, 2, 291, 434, 3, 2, 2, 2, 292, 293,
	7, 58, 2, 2, 293, 294, 7, 3, 2, 2, 294, 295, 5, 2, 2, 2, 295, 296, 7, 15,
	2, 2, 296, 299, 5, 2, 2, 2, 297, 298, 7, 15, 2, 2, 298, 300, 5, 2, 2, 2,
	299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301,
	302, 7, 4, 2, 2, 302, 434, 3, 2, 2, 2, 303, 304, 7, 59, 2, 2, 304, 305,
	7, 3, 2, 2, 305, 306, 5, 2, 2, 2, 306, 307, 7, 15, 2, 2, 307, 308, 5, 2,
	2, 2, 308, 309, 7, 15, 2, 2, 309, 312, 5, 2, 2, 2, 310, 311, 7, 15, 2,
	2, 311, 313, 5, 2, 2, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313,
	314, 3, 2, 2, 2, 314, 315, 7, 4, 2, 2, 315, 434, 3, 2, 2, 2, 316, 317,
	7, 60, 2, 2, 317, 318, 7, 3, 2, 2, 318, 319, 5, 2, 2, 2, 319, 320, 7, 15,
	2, 2, 320, 321, 5, 2, 2, 2, 321, 322, 7, 15, 2, 2, 322, 323, 5, 2, 2, 2,
	323, 324, 7, 4, 2, 2, 324, 434, 3, 2, 2, 2, 325, 326, 7, 61, 2, 2, 326,
	327, 7, 3, 2, 2, 327, 328, 5, 2, 2, 2, 328, 329, 7, 15, 2, 2, 329, 330,
	5, 2, 2, 2, 330, 331, 7, 4, 2, 2, 331, 434, 3, 2, 2, 2, 332, 333, 7, 62,
	2, 2, 333, 334, 7, 3, 2, 2, 334, 335, 5, 2, 2, 2, 335, 336, 7, 15, 2, 2,
	336, 337, 5, 2, 2, 2, 337, 338, 7, 4, 2, 2, 338, 434, 3, 2, 2, 2, 339,
	340, 7, 63, 2, 2, 340, 341, 7, 3, 2, 2, 341, 342, 5, 2, 2, 2, 342, 343,
	7, 4, 2, 2, 343, 434, 3, 2, 2, 2, 344, 345, 7, 64, 2, 2, 345, 346, 7, 3,
	2, 2, 346, 347, 5, 2, 2, 2, 347, 348, 7, 4, 2, 2, 348, 434, 3, 2, 2, 2,
	349, 350, 7, 65, 2, 2, 350, 351, 7, 3, 2, 2, 351, 352, 5, 2, 2, 2, 352,
	353, 7, 4, 2, 2, 353, 434, 3, 2, 2, 2, 354, 355, 7, 66, 2, 2, 355, 356,
	7, 3, 2, 2, 356, 357, 5, 2, 2, 2, 357, 358, 7, 15, 2, 2, 358, 359, 5, 2,
	2, 2, 359, 360, 7, 15, 2, 2, 360, 361, 5, 2, 2, 2, 361, 362, 7, 4, 2, 2,
	362, 434, 3, 2, 2, 2, 363, 364, 7, 67, 2, 2, 364, 365, 7, 3, 2, 2, 365,
	366, 5, 2, 2, 2, 366, 367, 7, 15, 2, 2, 367, 373, 5, 2, 2, 2, 368, 369,
	7, 15, 2, 2, 369, 370, 5, 2, 2, 2, 370, 371, 7, 15, 2, 2, 371, 372, 5,
	2, 2, 2, 372, 374, 3, 2, 2, 2, 373, 368, 3, 2, 2, 2, 374, 375, 3, 2, 2,
	2, 375, 373, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377,
	378, 7, 4, 2, 2, 378, 434, 3, 2, 2, 2, 379, 380, 7, 68, 2, 2, 380, 381,
	7, 3, 2, 2, 381, 382, 5, 2, 2, 2, 382, 383, 7, 4, 2, 2, 383, 434, 3, 2,
	2, 2, 384, 385, 7, 69, 2, 2, 385, 386, 7, 3, 2, 2, 386, 387, 5, 2, 2, 2,
	387, 388, 7, 4, 2, 2, 388, 434, 3, 2, 2, 2, 389, 390, 7, 70, 2, 2, 390,
	391, 7, 3, 2, 2, 391, 394, 5, 2, 2, 2, 392, 393, 7, 15, 2, 2, 393, 395,
	5, 2, 2, 2, 394, 392, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 396, 3, 2,
	2, 2, 396, 397, 7, 4, 2, 2, 397, 434, 3, 2, 2, 2, 398, 399, 7, 71, 2, 2,
	399, 400, 7, 3, 2, 2, 400, 403, 5, 2, 2, 2, 401, 402, 7, 15, 2, 2, 402,
	404, 5, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405,
	3, 2, 2, 2, 405, 406, 7, 4, 2, 2, 406, 434, 3, 2, 2, 2, 407, 408, 7, 72,
	2, 2, 408, 409, 7, 3, 2, 2, 409, 412, 5, 2, 2, 2, 410, 411, 7, 15, 2, 2,
	411, 413, 5, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413,
	414, 3, 2, 2, 2, 414, 415, 7, 4, 2, 2, 415, 434, 3, 2, 2, 2, 416, 417,
	7, 73, 2, 2, 417, 418, 7, 3, 2, 2, 418, 419, 5, 2, 2, 2, 419, 420, 7, 4,
	2, 2, 420, 434, 3, 2, 2, 2, 421, 434, 7, 84, 2, 2, 422, 423, 7, 8, 2, 2,
	423, 434, 7, 84, 2, 2, 424, 434, 7, 85, 2, 2, 425, 426, 7, 8, 2, 2, 426,
	434, 7, 85, 2, 2, 427, 434, 7, 89, 2, 2, 428, 434, 7, 90, 2, 2, 429, 434,
	7, 87, 2, 2, 430, 434, 7, 86, 2, 2, 431, 434, 7, 83, 2, 2, 432, 434, 7,
	88, 2, 2, 433, 4, 3, 2, 2, 2, 433, 9, 3, 2, 2, 2, 433, 27, 3, 2, 2, 2,
	433, 32, 3, 2, 2, 2, 433, 37, 3, 2, 2, 2, 433, 42, 3, 2, 2, 2, 433, 47,
	3, 2, 2, 2, 433, 54, 3, 2, 2, 2, 433, 66, 3, 2, 2, 2, 433, 71, 3, 2, 2,
	2, 433, 76, 3, 2, 2, 2, 433, 81, 3, 2, 2, 2, 433, 92, 3, 2, 2, 2, 433,
	97, 3, 2, 2, 2, 433, 102, 3, 2, 2, 2, 433, 107, 3, 2, 2, 2, 433, 118, 3,
	2, 2, 2, 433, 129, 3, 2, 2, 2, 433, 134, 3, 2, 2, 2, 433, 141, 3, 2, 2,
	2, 433, 146, 3, 2, 2, 2, 433, 153, 3, 2, 2, 2, 433, 158, 3, 2, 2, 2, 433,
	167, 3, 2, 2, 2, 433, 172, 3, 2, 2, 2, 433, 179, 3, 2, 2, 2, 433, 184,
	3, 2, 2, 2, 433, 189, 3, 2, 2, 2, 433, 194, 3, 2, 2, 2, 433, 199, 3, 2,
	2, 2, 433, 210, 3, 2, 2, 2, 433, 222, 3, 2, 2, 2, 433, 233, 3, 2, 2, 2,
	433, 245, 3, 2, 2, 2, 433, 247, 3, 2, 2, 2, 433, 256, 3, 2, 2, 2, 433,
	265, 3, 2, 2, 2, 433, 267, 3, 2, 2, 2, 433, 274, 3, 2, 2, 2, 433, 276,
	3, 2, 2, 2, 433, 281, 3, 2, 2, 2, 433, 292, 3, 2, 2, 2, 433, 303, 3, 2,
	2, 2, 433, 316, 3, 2, 2, 2, 433, 325, 3, 2, 2, 2, 433, 332, 3, 2, 2, 2,
	433, 339, 3, 2, 2, 2, 433, 344, 3, 2, 2, 2, 433, 349, 3, 2, 2, 2, 433,
	354, 3, 2, 2, 2, 433, 363, 3, 2, 2, 2, 433, 379, 3, 2, 2, 2, 433, 384,
	3, 2, 2, 2, 433, 389, 3, 2, 2, 2, 433, 398, 3, 2, 2, 2, 433, 407, 3, 2,
	2, 2, 433, 416, 3, 2, 2, 2, 433, 421, 3, 2, 2, 2, 433, 422, 3, 2, 2, 2,
	433, 424, 3, 2, 2, 2, 433, 425, 3, 2, 2, 2, 433, 427, 3, 2, 2, 2, 433,
	428, 3, 2, 2, 2, 433, 429, 3, 2, 2, 2, 433, 430, 3, 2, 2, 2, 433, 431,
	3, 2, 2, 2, 433, 432, 3, 2, 2, 2, 434, 502, 3, 2, 2, 2, 435, 436, 12, 82,
	2, 2, 436, 437, 7, 5, 2, 2, 437, 501, 5, 2, 2, 83, 438, 439, 12, 81, 2,
	2, 439, 440, 7, 6, 2, 2, 440, 501, 5, 2, 2, 82, 441, 442, 12, 80, 2, 2,
	442, 443, 7, 7, 2, 2, 443, 501, 5, 2, 2, 81, 444, 445, 12, 79, 2, 2, 445,
	446, 7, 8, 2, 2, 446, 501, 5, 2, 2, 80, 447, 448, 12, 78, 2, 2, 448, 449,
	7, 9, 2, 2, 449, 501, 5, 2, 2, 79, 450, 451, 12, 77, 2, 2, 451, 452, 7,
	10, 2, 2, 452, 501, 5, 2, 2, 78, 453, 454, 12, 76, 2, 2, 454, 455, 7, 11,
	2, 2, 455, 501, 5, 2, 2, 77, 456, 457, 12, 75, 2, 2, 457, 458, 7, 12, 2,
	2, 458, 501, 5, 2, 2, 76, 459, 460, 12, 74, 2, 2, 460, 461, 7, 13, 2, 2,
	461, 501, 5, 2, 2, 75, 462, 463, 12, 73, 2, 2, 463, 464, 7, 14, 2, 2, 464,
	501, 5, 2, 2, 74, 465, 466, 12, 70, 2, 2, 466, 467, 9, 2, 2, 2, 467, 501,
	5, 2, 2, 71, 468, 469, 12, 69, 2, 2, 469, 470, 9, 3, 2, 2, 470, 501, 5,
	2, 2, 70, 471, 472, 12, 72, 2, 2, 472, 473, 7, 74, 2, 2, 473, 482, 7, 3,
	2, 2, 474, 479, 5, 2, 2, 2, 475, 476, 7, 15, 2, 2, 476, 478, 5, 2, 2, 2,
	477, 475, 3, 2, 2, 2, 478, 481, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479,
	480, 3, 2, 2, 2, 480, 483, 3, 2, 2, 2, 481, 479, 3, 2, 2, 2, 482, 474,
	3, 2, 2, 2, 482, 483, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 501, 7, 4,
	2, 2, 485, 486, 12, 71, 2, 2, 486, 487, 7, 75, 2, 2, 487, 488, 7, 74, 2,
	2, 488, 497, 7, 3, 2, 2, 489, 494, 5, 2, 2, 2, 490, 491, 7, 15, 2, 2, 491,
	493, 5, 2, 2, 2, 492, 490, 3, 2, 2, 2, 493, 496, 3, 2, 2, 2, 494, 492,
	3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 498, 3, 2, 2, 2, 496, 494, 3, 2,
	2, 2, 497, 489, 3, 2, 2, 2, 497, 498, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2,
	499, 501, 7, 4, 2, 2, 500, 435, 3, 2, 2, 2, 500, 438, 3, 2, 2, 2, 500,
	441, 3, 2, 2, 2, 500, 444, 3, 2, 2, 2, 500, 447, 3, 2, 2, 2, 500, 450,
	3, 2, 2, 2, 500, 453, 3, 2, 2, 2, 500, 456, 3, 2, 2, 2, 500, 459, 3, 2,
	2, 2, 500, 462, 3, 2, 2, 2, 500, 465, 3, 2, 2, 2, 500, 468, 3, 2, 2, 2,
	500, 471, 3, 2, 2, 2, 500, 485, 3, 2, 2, 2, 501, 504, 3, 2, 2, 2, 502,
	500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 3, 3, 2, 2, 2, 504, 502, 3,
	2, 2, 2, 24, 20, 61, 88, 125, 206, 217, 229, 240, 288, 299, 312, 375, 394,
	403, 412, 433, 479, 482, 494, 497, 500, 502,
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
	"Exp", "Floor", "GetWord", "HexToNumber", "Iif", "IsNull", "Left", "Length",
	"Log", "Log10", "Lowercase", "Max", "Median", "Min", "Mod", "Null", "PadLeft",
	"PadRight", "Pi", "Pow", "Rand", "RandInt", "Regex_CountMatches", "Regex_Match",
	"Regex_Replace", "Replace", "Right", "Round", "Sin", "Sinh", "Sqrt", "Substring",
	"Switch", "Tan", "Tanh", "Trim", "TrimLeft", "TrimRight", "Uppercase",
	"In", "Not", "And", "Or", "If", "Then", "Else", "Elseif", "Endif", "Bool",
	"Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace",
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
	AlteryxFormulasParserEOF                = antlr.TokenEOF
	AlteryxFormulasParserT__0               = 1
	AlteryxFormulasParserT__1               = 2
	AlteryxFormulasParserT__2               = 3
	AlteryxFormulasParserT__3               = 4
	AlteryxFormulasParserT__4               = 5
	AlteryxFormulasParserT__5               = 6
	AlteryxFormulasParserT__6               = 7
	AlteryxFormulasParserT__7               = 8
	AlteryxFormulasParserT__8               = 9
	AlteryxFormulasParserT__9               = 10
	AlteryxFormulasParserT__10              = 11
	AlteryxFormulasParserT__11              = 12
	AlteryxFormulasParserT__12              = 13
	AlteryxFormulasParserT__13              = 14
	AlteryxFormulasParserT__14              = 15
	AlteryxFormulasParserT__15              = 16
	AlteryxFormulasParserAbs                = 17
	AlteryxFormulasParserAcos               = 18
	AlteryxFormulasParserAsin               = 19
	AlteryxFormulasParserAtan               = 20
	AlteryxFormulasParserAtan2              = 21
	AlteryxFormulasParserAverage            = 22
	AlteryxFormulasParserCeil               = 23
	AlteryxFormulasParserCharFromInt        = 24
	AlteryxFormulasParserCharToInt          = 25
	AlteryxFormulasParserContains           = 26
	AlteryxFormulasParserCos                = 27
	AlteryxFormulasParserCosh               = 28
	AlteryxFormulasParserCountWords         = 29
	AlteryxFormulasParserDistance           = 30
	AlteryxFormulasParserEndsWith           = 31
	AlteryxFormulasParserFindString         = 32
	AlteryxFormulasParserExp                = 33
	AlteryxFormulasParserFloor              = 34
	AlteryxFormulasParserGetWord            = 35
	AlteryxFormulasParserHexToNumber        = 36
	AlteryxFormulasParserIif                = 37
	AlteryxFormulasParserIsNull             = 38
	AlteryxFormulasParserLeft               = 39
	AlteryxFormulasParserLength             = 40
	AlteryxFormulasParserLog                = 41
	AlteryxFormulasParserLog10              = 42
	AlteryxFormulasParserLowercase          = 43
	AlteryxFormulasParserMax                = 44
	AlteryxFormulasParserMedian             = 45
	AlteryxFormulasParserMin                = 46
	AlteryxFormulasParserMod                = 47
	AlteryxFormulasParserNull               = 48
	AlteryxFormulasParserPadLeft            = 49
	AlteryxFormulasParserPadRight           = 50
	AlteryxFormulasParserPi                 = 51
	AlteryxFormulasParserPow                = 52
	AlteryxFormulasParserRand               = 53
	AlteryxFormulasParserRandInt            = 54
	AlteryxFormulasParserRegex_CountMatches = 55
	AlteryxFormulasParserRegex_Match        = 56
	AlteryxFormulasParserRegex_Replace      = 57
	AlteryxFormulasParserReplace            = 58
	AlteryxFormulasParserRight              = 59
	AlteryxFormulasParserRound              = 60
	AlteryxFormulasParserSin                = 61
	AlteryxFormulasParserSinh               = 62
	AlteryxFormulasParserSqrt               = 63
	AlteryxFormulasParserSubstring          = 64
	AlteryxFormulasParserSwitch             = 65
	AlteryxFormulasParserTan                = 66
	AlteryxFormulasParserTanh               = 67
	AlteryxFormulasParserTrim               = 68
	AlteryxFormulasParserTrimLeft           = 69
	AlteryxFormulasParserTrimRight          = 70
	AlteryxFormulasParserUppercase          = 71
	AlteryxFormulasParserIn                 = 72
	AlteryxFormulasParserNot                = 73
	AlteryxFormulasParserAnd                = 74
	AlteryxFormulasParserOr                 = 75
	AlteryxFormulasParserIf                 = 76
	AlteryxFormulasParserThen               = 77
	AlteryxFormulasParserElse               = 78
	AlteryxFormulasParserElseif             = 79
	AlteryxFormulasParserEndif              = 80
	AlteryxFormulasParserBool               = 81
	AlteryxFormulasParserInteger            = 82
	AlteryxFormulasParserDecimal            = 83
	AlteryxFormulasParserDate               = 84
	AlteryxFormulasParserDatetime           = 85
	AlteryxFormulasParserField              = 86
	AlteryxFormulasParserSingleQuoteString  = 87
	AlteryxFormulasParserDoubleQuoteString  = 88
	AlteryxFormulasParserWhitespace         = 89
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

type ReplaceFuncContext struct {
	*ExprContext
}

func NewReplaceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReplaceFuncContext {
	var p = new(ReplaceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ReplaceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceFuncContext) Replace() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserReplace, 0)
}

func (s *ReplaceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ReplaceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReplaceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterReplaceFunc(s)
	}
}

func (s *ReplaceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitReplaceFunc(s)
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

type TrimRightFuncContext struct {
	*ExprContext
}

func NewTrimRightFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrimRightFuncContext {
	var p = new(TrimRightFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TrimRightFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrimRightFuncContext) TrimRight() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTrimRight, 0)
}

func (s *TrimRightFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TrimRightFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TrimRightFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTrimRightFunc(s)
	}
}

func (s *TrimRightFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTrimRightFunc(s)
	}
}

type LeftFuncContext struct {
	*ExprContext
}

func NewLeftFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeftFuncContext {
	var p = new(LeftFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LeftFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftFuncContext) Left() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLeft, 0)
}

func (s *LeftFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LeftFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LeftFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLeftFunc(s)
	}
}

func (s *LeftFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLeftFunc(s)
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

type RegexReplaceFuncContext struct {
	*ExprContext
}

func NewRegexReplaceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexReplaceFuncContext {
	var p = new(RegexReplaceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexReplaceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexReplaceFuncContext) Regex_Replace() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_Replace, 0)
}

func (s *RegexReplaceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexReplaceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexReplaceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexReplaceFunc(s)
	}
}

func (s *RegexReplaceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexReplaceFunc(s)
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

type IsNullFuncContext struct {
	*ExprContext
}

func NewIsNullFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullFuncContext {
	var p = new(IsNullFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IsNullFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullFuncContext) IsNull() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIsNull, 0)
}

func (s *IsNullFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IsNullFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIsNullFunc(s)
	}
}

func (s *IsNullFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIsNullFunc(s)
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

type RegexMatchFuncContext struct {
	*ExprContext
}

func NewRegexMatchFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexMatchFuncContext {
	var p = new(RegexMatchFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexMatchFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexMatchFuncContext) Regex_Match() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_Match, 0)
}

func (s *RegexMatchFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexMatchFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexMatchFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexMatchFunc(s)
	}
}

func (s *RegexMatchFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexMatchFunc(s)
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

type PadLeftFuncContext struct {
	*ExprContext
}

func NewPadLeftFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PadLeftFuncContext {
	var p = new(PadLeftFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PadLeftFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadLeftFuncContext) PadLeft() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPadLeft, 0)
}

func (s *PadLeftFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PadLeftFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PadLeftFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPadLeftFunc(s)
	}
}

func (s *PadLeftFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPadLeftFunc(s)
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

type PadRightFuncContext struct {
	*ExprContext
}

func NewPadRightFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PadRightFuncContext {
	var p = new(PadRightFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PadRightFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadRightFuncContext) PadRight() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPadRight, 0)
}

func (s *PadRightFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PadRightFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PadRightFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPadRightFunc(s)
	}
}

func (s *PadRightFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPadRightFunc(s)
	}
}

type RegexCountMatchesFuncContext struct {
	*ExprContext
}

func NewRegexCountMatchesFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexCountMatchesFuncContext {
	var p = new(RegexCountMatchesFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexCountMatchesFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexCountMatchesFuncContext) Regex_CountMatches() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_CountMatches, 0)
}

func (s *RegexCountMatchesFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexCountMatchesFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexCountMatchesFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexCountMatchesFunc(s)
	}
}

func (s *RegexCountMatchesFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexCountMatchesFunc(s)
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

type LowercaseFuncContext struct {
	*ExprContext
}

func NewLowercaseFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LowercaseFuncContext {
	var p = new(LowercaseFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LowercaseFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LowercaseFuncContext) Lowercase() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLowercase, 0)
}

func (s *LowercaseFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LowercaseFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLowercaseFunc(s)
	}
}

func (s *LowercaseFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLowercaseFunc(s)
	}
}

type TrimFuncContext struct {
	*ExprContext
}

func NewTrimFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrimFuncContext {
	var p = new(TrimFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TrimFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrimFuncContext) Trim() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTrim, 0)
}

func (s *TrimFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TrimFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TrimFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTrimFunc(s)
	}
}

func (s *TrimFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTrimFunc(s)
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

type LengthFuncContext struct {
	*ExprContext
}

func NewLengthFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthFuncContext {
	var p = new(LengthFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LengthFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFuncContext) Length() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLength, 0)
}

func (s *LengthFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LengthFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLengthFunc(s)
	}
}

func (s *LengthFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLengthFunc(s)
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

type UppercaseFuncContext struct {
	*ExprContext
}

func NewUppercaseFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UppercaseFuncContext {
	var p = new(UppercaseFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UppercaseFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UppercaseFuncContext) Uppercase() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserUppercase, 0)
}

func (s *UppercaseFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UppercaseFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterUppercaseFunc(s)
	}
}

func (s *UppercaseFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitUppercaseFunc(s)
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

type RightFuncContext struct {
	*ExprContext
}

func NewRightFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightFuncContext {
	var p = new(RightFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RightFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightFuncContext) Right() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRight, 0)
}

func (s *RightFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RightFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RightFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRightFunc(s)
	}
}

func (s *RightFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRightFunc(s)
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

type TrimLeftFuncContext struct {
	*ExprContext
}

func NewTrimLeftFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrimLeftFuncContext {
	var p = new(TrimLeftFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TrimLeftFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrimLeftFuncContext) TrimLeft() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTrimLeft, 0)
}

func (s *TrimLeftFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TrimLeftFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TrimLeftFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTrimLeftFunc(s)
	}
}

func (s *TrimLeftFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTrimLeftFunc(s)
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

type SubstringFuncContext struct {
	*ExprContext
}

func NewSubstringFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubstringFuncContext {
	var p = new(SubstringFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubstringFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubstringFuncContext) Substring() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSubstring, 0)
}

func (s *SubstringFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubstringFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubstringFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSubstringFunc(s)
	}
}

func (s *SubstringFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSubstringFunc(s)
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
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
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
		localctx = NewAbsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(AlteryxFormulasParserAbs)
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
			p.Match(AlteryxFormulasParserT__1)
		}

	case 4:
		localctx = NewAcosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(30)
			p.Match(AlteryxFormulasParserAcos)
		}
		{
			p.SetState(31)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(32)
			p.expr(0)
		}
		{
			p.SetState(33)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 5:
		localctx = NewAsinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(35)
			p.Match(AlteryxFormulasParserAsin)
		}
		{
			p.SetState(36)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(37)
			p.expr(0)
		}
		{
			p.SetState(38)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 6:
		localctx = NewAtanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(AlteryxFormulasParserAtan)
		}
		{
			p.SetState(41)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(42)
			p.expr(0)
		}
		{
			p.SetState(43)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 7:
		localctx = NewAtan2FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(AlteryxFormulasParserAtan2)
		}
		{
			p.SetState(46)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(47)
			p.expr(0)
		}
		{
			p.SetState(48)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(49)
			p.expr(0)
		}
		{
			p.SetState(50)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 8:
		localctx = NewAverageFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(AlteryxFormulasParserAverage)
		}
		{
			p.SetState(53)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(54)
			p.expr(0)
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(55)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(56)
				p.expr(0)
			}

			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(62)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 9:
		localctx = NewCeilFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(AlteryxFormulasParserCeil)
		}
		{
			p.SetState(65)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(66)
			p.expr(0)
		}
		{
			p.SetState(67)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 10:
		localctx = NewCharFromIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(AlteryxFormulasParserCharFromInt)
		}
		{
			p.SetState(70)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(71)
			p.expr(0)
		}
		{
			p.SetState(72)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 11:
		localctx = NewCharToIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(AlteryxFormulasParserCharToInt)
		}
		{
			p.SetState(75)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(76)
			p.expr(0)
		}
		{
			p.SetState(77)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 12:
		localctx = NewContainsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)
			p.Match(AlteryxFormulasParserContains)
		}
		{
			p.SetState(80)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		{
			p.SetState(82)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(83)
			p.expr(0)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(84)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(85)
				p.expr(0)
			}

		}
		{
			p.SetState(88)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 13:
		localctx = NewCosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			p.Match(AlteryxFormulasParserCos)
		}
		{
			p.SetState(91)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(92)
			p.expr(0)
		}
		{
			p.SetState(93)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 14:
		localctx = NewCoshFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(95)
			p.Match(AlteryxFormulasParserCosh)
		}
		{
			p.SetState(96)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(97)
			p.expr(0)
		}
		{
			p.SetState(98)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 15:
		localctx = NewCountWordsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(AlteryxFormulasParserCountWords)
		}
		{
			p.SetState(101)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(102)
			p.expr(0)
		}
		{
			p.SetState(103)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 16:
		localctx = NewDistanceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.Match(AlteryxFormulasParserDistance)
		}
		{
			p.SetState(106)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(107)
			p.expr(0)
		}
		{
			p.SetState(108)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(109)
			p.expr(0)
		}
		{
			p.SetState(110)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(111)
			p.expr(0)
		}
		{
			p.SetState(112)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(113)
			p.expr(0)
		}
		{
			p.SetState(114)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 17:
		localctx = NewEndsWithFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.Match(AlteryxFormulasParserEndsWith)
		}
		{
			p.SetState(117)
			p.Match(AlteryxFormulasParserT__0)
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
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(121)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(122)
				p.expr(0)
			}

		}
		{
			p.SetState(125)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 18:
		localctx = NewExpFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.Match(AlteryxFormulasParserExp)
		}
		{
			p.SetState(128)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(129)
			p.expr(0)
		}
		{
			p.SetState(130)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 19:
		localctx = NewFindStringFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(AlteryxFormulasParserFindString)
		}
		{
			p.SetState(133)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(134)
			p.expr(0)
		}
		{
			p.SetState(135)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(136)
			p.expr(0)
		}
		{
			p.SetState(137)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 20:
		localctx = NewFloorFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(139)
			p.Match(AlteryxFormulasParserFloor)
		}
		{
			p.SetState(140)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(141)
			p.expr(0)
		}
		{
			p.SetState(142)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 21:
		localctx = NewGetWordFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(144)
			p.Match(AlteryxFormulasParserGetWord)
		}
		{
			p.SetState(145)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(146)
			p.expr(0)
		}
		{
			p.SetState(147)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(148)
			p.expr(0)
		}
		{
			p.SetState(149)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 22:
		localctx = NewHexToNumberFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(151)
			p.Match(AlteryxFormulasParserHexToNumber)
		}
		{
			p.SetState(152)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(153)
			p.expr(0)
		}
		{
			p.SetState(154)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 23:
		localctx = NewIifFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(156)
			p.Match(AlteryxFormulasParserIif)
		}
		{
			p.SetState(157)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(158)
			p.expr(0)
		}
		{
			p.SetState(159)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(160)
			p.expr(0)
		}
		{
			p.SetState(161)
			p.Match(AlteryxFormulasParserT__12)
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
		localctx = NewIsNullFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(AlteryxFormulasParserIsNull)
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
		localctx = NewLeftFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(170)
			p.Match(AlteryxFormulasParserLeft)
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
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(174)
			p.expr(0)
		}
		{
			p.SetState(175)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 26:
		localctx = NewLengthFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(177)
			p.Match(AlteryxFormulasParserLength)
		}
		{
			p.SetState(178)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(179)
			p.expr(0)
		}
		{
			p.SetState(180)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 27:
		localctx = NewLogFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Match(AlteryxFormulasParserLog)
		}
		{
			p.SetState(183)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(184)
			p.expr(0)
		}
		{
			p.SetState(185)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 28:
		localctx = NewLog10FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(187)
			p.Match(AlteryxFormulasParserLog10)
		}
		{
			p.SetState(188)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(189)
			p.expr(0)
		}
		{
			p.SetState(190)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 29:
		localctx = NewLowercaseFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(192)
			p.Match(AlteryxFormulasParserLowercase)
		}
		{
			p.SetState(193)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(194)
			p.expr(0)
		}
		{
			p.SetState(195)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 30:
		localctx = NewMaxFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(197)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(198)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(199)
			p.expr(0)
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(200)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(201)
				p.expr(0)
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(206)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 31:
		localctx = NewMedianFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(208)
			p.Match(AlteryxFormulasParserMedian)
		}
		{
			p.SetState(209)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(210)
			p.expr(0)
		}
		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(211)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(212)
				p.expr(0)
			}

			p.SetState(217)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(218)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 32:
		localctx = NewMinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(220)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(221)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(222)
			p.expr(0)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(223)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(224)
				p.expr(0)
			}

			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(229)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 33:
		localctx = NewModFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(231)
			p.Match(AlteryxFormulasParserMod)
		}
		{
			p.SetState(232)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(233)
			p.expr(0)
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(234)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(235)
				p.expr(0)
			}

			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(241)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 34:
		localctx = NewNullFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(243)
			p.Match(AlteryxFormulasParserNull)
		}
		{
			p.SetState(244)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 35:
		localctx = NewPadLeftFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(245)
			p.Match(AlteryxFormulasParserPadLeft)
		}
		{
			p.SetState(246)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(247)
			p.expr(0)
		}
		{
			p.SetState(248)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(249)
			p.expr(0)
		}
		{
			p.SetState(250)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(251)
			p.expr(0)
		}
		{
			p.SetState(252)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 36:
		localctx = NewPadRightFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(254)
			p.Match(AlteryxFormulasParserPadRight)
		}
		{
			p.SetState(255)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(256)
			p.expr(0)
		}
		{
			p.SetState(257)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(258)
			p.expr(0)
		}
		{
			p.SetState(259)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(260)
			p.expr(0)
		}
		{
			p.SetState(261)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 37:
		localctx = NewPiFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(263)
			p.Match(AlteryxFormulasParserPi)
		}
		{
			p.SetState(264)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 38:
		localctx = NewPowFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(265)
			p.Match(AlteryxFormulasParserPow)
		}
		{
			p.SetState(266)
			p.Match(AlteryxFormulasParserT__0)
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
		{
			p.SetState(270)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 39:
		localctx = NewRandFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)
			p.Match(AlteryxFormulasParserRand)
		}
		{
			p.SetState(273)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 40:
		localctx = NewRandIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(274)
			p.Match(AlteryxFormulasParserRandInt)
		}
		{
			p.SetState(275)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(276)
			p.expr(0)
		}
		{
			p.SetState(277)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 41:
		localctx = NewRegexCountMatchesFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(279)
			p.Match(AlteryxFormulasParserRegex_CountMatches)
		}
		{
			p.SetState(280)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(281)
			p.expr(0)
		}
		{
			p.SetState(282)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(283)
			p.expr(0)
		}
		p.SetState(286)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(284)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(285)
				p.expr(0)
			}

		}
		{
			p.SetState(288)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 42:
		localctx = NewRegexMatchFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(290)
			p.Match(AlteryxFormulasParserRegex_Match)
		}
		{
			p.SetState(291)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(292)
			p.expr(0)
		}
		{
			p.SetState(293)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(294)
			p.expr(0)
		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(295)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(296)
				p.expr(0)
			}

		}
		{
			p.SetState(299)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 43:
		localctx = NewRegexReplaceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(301)
			p.Match(AlteryxFormulasParserRegex_Replace)
		}
		{
			p.SetState(302)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(303)
			p.expr(0)
		}
		{
			p.SetState(304)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(305)
			p.expr(0)
		}
		{
			p.SetState(306)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(307)
			p.expr(0)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(308)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(309)
				p.expr(0)
			}

		}
		{
			p.SetState(312)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 44:
		localctx = NewReplaceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Match(AlteryxFormulasParserReplace)
		}
		{
			p.SetState(315)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(316)
			p.expr(0)
		}
		{
			p.SetState(317)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(318)
			p.expr(0)
		}
		{
			p.SetState(319)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(320)
			p.expr(0)
		}
		{
			p.SetState(321)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 45:
		localctx = NewRightFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)
			p.Match(AlteryxFormulasParserRight)
		}
		{
			p.SetState(324)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(325)
			p.expr(0)
		}
		{
			p.SetState(326)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(327)
			p.expr(0)
		}
		{
			p.SetState(328)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 46:
		localctx = NewRoundFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(330)
			p.Match(AlteryxFormulasParserRound)
		}
		{
			p.SetState(331)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(332)
			p.expr(0)
		}
		{
			p.SetState(333)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(334)
			p.expr(0)
		}
		{
			p.SetState(335)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 47:
		localctx = NewSinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)
			p.Match(AlteryxFormulasParserSin)
		}
		{
			p.SetState(338)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(339)
			p.expr(0)
		}
		{
			p.SetState(340)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 48:
		localctx = NewSinhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(342)
			p.Match(AlteryxFormulasParserSinh)
		}
		{
			p.SetState(343)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(344)
			p.expr(0)
		}
		{
			p.SetState(345)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 49:
		localctx = NewSqrtFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(347)
			p.Match(AlteryxFormulasParserSqrt)
		}
		{
			p.SetState(348)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(349)
			p.expr(0)
		}
		{
			p.SetState(350)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 50:
		localctx = NewSubstringFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(352)
			p.Match(AlteryxFormulasParserSubstring)
		}
		{
			p.SetState(353)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(354)
			p.expr(0)
		}
		{
			p.SetState(355)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(356)
			p.expr(0)
		}
		{
			p.SetState(357)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(358)
			p.expr(0)
		}
		{
			p.SetState(359)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 51:
		localctx = NewSwitchFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(361)
			p.Match(AlteryxFormulasParserSwitch)
		}
		{
			p.SetState(362)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(363)
			p.expr(0)
		}
		{
			p.SetState(364)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(365)
			p.expr(0)
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(366)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(367)
				p.expr(0)
			}
			{
				p.SetState(368)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(369)
				p.expr(0)
			}

			p.SetState(373)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(375)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 52:
		localctx = NewTanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(377)
			p.Match(AlteryxFormulasParserTan)
		}
		{
			p.SetState(378)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(379)
			p.expr(0)
		}
		{
			p.SetState(380)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 53:
		localctx = NewTanhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(382)
			p.Match(AlteryxFormulasParserTanh)
		}
		{
			p.SetState(383)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(384)
			p.expr(0)
		}
		{
			p.SetState(385)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 54:
		localctx = NewTrimFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(387)
			p.Match(AlteryxFormulasParserTrim)
		}
		{
			p.SetState(388)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(389)
			p.expr(0)
		}
		p.SetState(392)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(390)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(391)
				p.expr(0)
			}

		}
		{
			p.SetState(394)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 55:
		localctx = NewTrimLeftFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(396)
			p.Match(AlteryxFormulasParserTrimLeft)
		}
		{
			p.SetState(397)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(398)
			p.expr(0)
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(399)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(400)
				p.expr(0)
			}

		}
		{
			p.SetState(403)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 56:
		localctx = NewTrimRightFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(405)
			p.Match(AlteryxFormulasParserTrimRight)
		}
		{
			p.SetState(406)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(407)
			p.expr(0)
		}
		p.SetState(410)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(408)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(409)
				p.expr(0)
			}

		}
		{
			p.SetState(412)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 57:
		localctx = NewUppercaseFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(414)
			p.Match(AlteryxFormulasParserUppercase)
		}
		{
			p.SetState(415)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(416)
			p.expr(0)
		}
		{
			p.SetState(417)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 58:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(419)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 59:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(420)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(421)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 60:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(422)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 61:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(423)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(424)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 62:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(425)
			p.Match(AlteryxFormulasParserSingleQuoteString)
		}

	case 63:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(426)
			p.Match(AlteryxFormulasParserDoubleQuoteString)
		}

	case 64:
		localctx = NewDatetimeLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(427)
			p.Match(AlteryxFormulasParserDatetime)
		}

	case 65:
		localctx = NewDateLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(428)
			p.Match(AlteryxFormulasParserDate)
		}

	case 66:
		localctx = NewBoolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(429)
			p.Match(AlteryxFormulasParserBool)
		}

	case 67:
		localctx = NewExprFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(430)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(498)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplyContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 80)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 80)", ""))
				}
				{
					p.SetState(434)
					p.Match(AlteryxFormulasParserT__2)
				}
				{
					p.SetState(435)

					var _x = p.expr(81)

					localctx.(*MultiplyContext).right = _x
				}

			case 2:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivideContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(436)

				if !(p.Precpred(p.GetParserRuleContext(), 79)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 79)", ""))
				}
				{
					p.SetState(437)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(438)

					var _x = p.expr(80)

					localctx.(*DivideContext).right = _x
				}

			case 3:
				localctx = NewAddContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 78)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 78)", ""))
				}
				{
					p.SetState(440)
					p.Match(AlteryxFormulasParserT__4)
				}
				{
					p.SetState(441)

					var _x = p.expr(79)

					localctx.(*AddContext).right = _x
				}

			case 4:
				localctx = NewSubtractContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubtractContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(442)

				if !(p.Precpred(p.GetParserRuleContext(), 77)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 77)", ""))
				}
				{
					p.SetState(443)
					p.Match(AlteryxFormulasParserT__5)
				}
				{
					p.SetState(444)

					var _x = p.expr(78)

					localctx.(*SubtractContext).right = _x
				}

			case 5:
				localctx = NewEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(445)

				if !(p.Precpred(p.GetParserRuleContext(), 76)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 76)", ""))
				}
				{
					p.SetState(446)
					p.Match(AlteryxFormulasParserT__6)
				}
				{
					p.SetState(447)

					var _x = p.expr(77)

					localctx.(*EqualContext).right = _x
				}

			case 6:
				localctx = NewGreaterThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(448)

				if !(p.Precpred(p.GetParserRuleContext(), 75)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 75)", ""))
				}
				{
					p.SetState(449)
					p.Match(AlteryxFormulasParserT__7)
				}
				{
					p.SetState(450)

					var _x = p.expr(76)

					localctx.(*GreaterThanContext).right = _x
				}

			case 7:
				localctx = NewGreaterEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(451)

				if !(p.Precpred(p.GetParserRuleContext(), 74)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 74)", ""))
				}
				{
					p.SetState(452)
					p.Match(AlteryxFormulasParserT__8)
				}
				{
					p.SetState(453)

					var _x = p.expr(75)

					localctx.(*GreaterEqualContext).right = _x
				}

			case 8:
				localctx = NewLessThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(454)

				if !(p.Precpred(p.GetParserRuleContext(), 73)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 73)", ""))
				}
				{
					p.SetState(455)
					p.Match(AlteryxFormulasParserT__9)
				}
				{
					p.SetState(456)

					var _x = p.expr(74)

					localctx.(*LessThanContext).right = _x
				}

			case 9:
				localctx = NewLessEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(457)

				if !(p.Precpred(p.GetParserRuleContext(), 72)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 72)", ""))
				}
				{
					p.SetState(458)
					p.Match(AlteryxFormulasParserT__10)
				}
				{
					p.SetState(459)

					var _x = p.expr(73)

					localctx.(*LessEqualContext).right = _x
				}

			case 10:
				localctx = NewNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(460)

				if !(p.Precpred(p.GetParserRuleContext(), 71)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 71)", ""))
				}
				{
					p.SetState(461)
					p.Match(AlteryxFormulasParserT__11)
				}
				{
					p.SetState(462)

					var _x = p.expr(72)

					localctx.(*NotEqualContext).right = _x
				}

			case 11:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(463)

				if !(p.Precpred(p.GetParserRuleContext(), 68)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 68)", ""))
				}
				{
					p.SetState(464)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__13 || _la == AlteryxFormulasParserAnd) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(465)

					var _x = p.expr(69)

					localctx.(*AndContext).right = _x
				}

			case 12:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(466)

				if !(p.Precpred(p.GetParserRuleContext(), 67)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 67)", ""))
				}
				{
					p.SetState(467)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__14 || _la == AlteryxFormulasParserOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(468)

					var _x = p.expr(68)

					localctx.(*OrContext).right = _x
				}

			case 13:
				localctx = NewInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(469)

				if !(p.Precpred(p.GetParserRuleContext(), 70)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 70)", ""))
				}
				{
					p.SetState(470)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(471)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(480)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIsNull-32))|(1<<(AlteryxFormulasParserLeft-32))|(1<<(AlteryxFormulasParserLength-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserLowercase-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPadLeft-32))|(1<<(AlteryxFormulasParserPadRight-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRegex_CountMatches-32))|(1<<(AlteryxFormulasParserRegex_Match-32))|(1<<(AlteryxFormulasParserRegex_Replace-32))|(1<<(AlteryxFormulasParserReplace-32))|(1<<(AlteryxFormulasParserRight-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(AlteryxFormulasParserSubstring-64))|(1<<(AlteryxFormulasParserSwitch-64))|(1<<(AlteryxFormulasParserTan-64))|(1<<(AlteryxFormulasParserTanh-64))|(1<<(AlteryxFormulasParserTrim-64))|(1<<(AlteryxFormulasParserTrimLeft-64))|(1<<(AlteryxFormulasParserTrimRight-64))|(1<<(AlteryxFormulasParserUppercase-64))|(1<<(AlteryxFormulasParserIf-64))|(1<<(AlteryxFormulasParserBool-64))|(1<<(AlteryxFormulasParserInteger-64))|(1<<(AlteryxFormulasParserDecimal-64))|(1<<(AlteryxFormulasParserDate-64))|(1<<(AlteryxFormulasParserDatetime-64))|(1<<(AlteryxFormulasParserField-64))|(1<<(AlteryxFormulasParserSingleQuoteString-64))|(1<<(AlteryxFormulasParserDoubleQuoteString-64)))) != 0) {
					{
						p.SetState(472)
						p.expr(0)
					}
					p.SetState(477)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(473)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(474)
							p.expr(0)
						}

						p.SetState(479)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(482)
					p.Match(AlteryxFormulasParserT__1)
				}

			case 14:
				localctx = NewNotInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(483)

				if !(p.Precpred(p.GetParserRuleContext(), 69)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 69)", ""))
				}
				{
					p.SetState(484)
					p.Match(AlteryxFormulasParserNot)
				}
				{
					p.SetState(485)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(486)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(495)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserIif-32))|(1<<(AlteryxFormulasParserIsNull-32))|(1<<(AlteryxFormulasParserLeft-32))|(1<<(AlteryxFormulasParserLength-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserLowercase-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPadLeft-32))|(1<<(AlteryxFormulasParserPadRight-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRegex_CountMatches-32))|(1<<(AlteryxFormulasParserRegex_Match-32))|(1<<(AlteryxFormulasParserRegex_Replace-32))|(1<<(AlteryxFormulasParserReplace-32))|(1<<(AlteryxFormulasParserRight-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(AlteryxFormulasParserSubstring-64))|(1<<(AlteryxFormulasParserSwitch-64))|(1<<(AlteryxFormulasParserTan-64))|(1<<(AlteryxFormulasParserTanh-64))|(1<<(AlteryxFormulasParserTrim-64))|(1<<(AlteryxFormulasParserTrimLeft-64))|(1<<(AlteryxFormulasParserTrimRight-64))|(1<<(AlteryxFormulasParserUppercase-64))|(1<<(AlteryxFormulasParserIf-64))|(1<<(AlteryxFormulasParserBool-64))|(1<<(AlteryxFormulasParserInteger-64))|(1<<(AlteryxFormulasParserDecimal-64))|(1<<(AlteryxFormulasParserDate-64))|(1<<(AlteryxFormulasParserDatetime-64))|(1<<(AlteryxFormulasParserField-64))|(1<<(AlteryxFormulasParserSingleQuoteString-64))|(1<<(AlteryxFormulasParserDoubleQuoteString-64)))) != 0) {
					{
						p.SetState(487)
						p.expr(0)
					}
					p.SetState(492)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(488)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(489)
							p.expr(0)
						}

						p.SetState(494)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(497)
					p.Match(AlteryxFormulasParserT__1)
				}

			}

		}
		p.SetState(502)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
		return p.Precpred(p.GetParserRuleContext(), 80)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 79)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 78)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 77)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 76)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 75)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 74)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 73)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 72)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 71)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 68)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 67)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 70)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 69)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
