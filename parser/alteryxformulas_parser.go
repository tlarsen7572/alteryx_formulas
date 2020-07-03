// Code generated from C:/repositories/alteryx_formulas/grammar\AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 38, 437,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 5, 2, 28, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 34, 10, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 6, 4, 58, 10, 4,
	13, 4, 14, 4, 59, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 69, 10,
	4, 3, 4, 3, 4, 3, 4, 7, 4, 74, 10, 4, 12, 4, 14, 4, 77, 11, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 6, 5, 84, 10, 5, 13, 5, 14, 5, 85, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 95, 10, 5, 13, 5, 14, 5, 96, 3, 5, 3,
	5, 5, 5, 101, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 6, 6, 126, 10, 6, 13, 6, 14, 6, 127, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 141, 10, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 155,
	10, 6, 12, 6, 14, 6, 158, 11, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 172, 10, 7, 13, 7, 14, 7, 173, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6, 7, 183, 10, 7, 13, 7, 14, 7,
	184, 3, 7, 3, 7, 5, 7, 189, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 6, 8, 212, 10, 8, 13, 8, 14, 8, 213, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 224, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 6, 9, 231, 10, 9, 13, 9, 14, 9, 232, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 6, 9, 242, 10, 9, 13, 9, 14, 9, 243, 3, 9, 3, 9, 5, 9, 248,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 285, 10, 10, 12, 10, 14, 10, 288,
	11, 10, 5, 10, 290, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 7, 10, 301, 10, 10, 12, 10, 14, 10, 304, 11, 10, 5, 10,
	306, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 7, 10, 340, 10, 10, 12, 10, 14, 10, 343, 11, 10, 5, 10,
	345, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 356, 10, 10, 12, 10, 14, 10, 359, 11, 10, 5, 10, 361, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 395, 10, 10, 12, 10, 14, 10, 398, 11, 10, 5, 10, 400, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 411,
	10, 10, 12, 10, 14, 10, 414, 11, 10, 5, 10, 416, 10, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 422, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 7, 10, 430, 10, 10, 12, 10, 14, 10, 433, 11, 10, 3, 11, 3, 11, 3, 11,
	2, 5, 6, 10, 18, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 5, 4, 2, 16,
	16, 23, 23, 4, 2, 17, 17, 24, 24, 3, 2, 36, 37, 2, 508, 2, 27, 3, 2, 2,
	2, 4, 33, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 100, 3, 2, 2, 2, 10, 140, 3,
	2, 2, 2, 12, 188, 3, 2, 2, 2, 14, 223, 3, 2, 2, 2, 16, 247, 3, 2, 2, 2,
	18, 421, 3, 2, 2, 2, 20, 434, 3, 2, 2, 2, 22, 28, 5, 4, 3, 2, 23, 28, 5,
	10, 6, 2, 24, 28, 5, 14, 8, 2, 25, 28, 5, 6, 4, 2, 26, 28, 5, 18, 10, 2,
	27, 22, 3, 2, 2, 2, 27, 23, 3, 2, 2, 2, 27, 24, 3, 2, 2, 2, 27, 25, 3,
	2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 3, 3, 2, 2, 2, 29, 30, 7, 3, 2, 2, 30,
	31, 7, 35, 2, 2, 31, 34, 7, 4, 2, 2, 32, 34, 7, 35, 2, 2, 33, 29, 3, 2,
	2, 2, 33, 32, 3, 2, 2, 2, 34, 5, 3, 2, 2, 2, 35, 36, 8, 4, 1, 2, 36, 37,
	7, 3, 2, 2, 37, 38, 5, 6, 4, 2, 38, 39, 7, 4, 2, 2, 39, 69, 3, 2, 2, 2,
	40, 41, 7, 25, 2, 2, 41, 42, 5, 18, 10, 2, 42, 43, 7, 26, 2, 2, 43, 44,
	5, 6, 4, 2, 44, 45, 7, 27, 2, 2, 45, 46, 5, 6, 4, 2, 46, 47, 7, 29, 2,
	2, 47, 69, 3, 2, 2, 2, 48, 49, 7, 25, 2, 2, 49, 50, 5, 18, 10, 2, 50, 51,
	7, 26, 2, 2, 51, 57, 5, 6, 4, 2, 52, 53, 7, 28, 2, 2, 53, 54, 5, 18, 10,
	2, 54, 55, 7, 26, 2, 2, 55, 56, 5, 6, 4, 2, 56, 58, 3, 2, 2, 2, 57, 52,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 61, 3, 2, 2, 2, 61, 62, 7, 27, 2, 2, 62, 63, 5, 6, 4, 2, 63, 64, 7,
	29, 2, 2, 64, 69, 3, 2, 2, 2, 65, 69, 5, 8, 5, 2, 66, 69, 5, 20, 11, 2,
	67, 69, 7, 35, 2, 2, 68, 35, 3, 2, 2, 2, 68, 40, 3, 2, 2, 2, 68, 48, 3,
	2, 2, 2, 68, 65, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	75, 3, 2, 2, 2, 70, 71, 12, 8, 2, 2, 71, 72, 7, 5, 2, 2, 72, 74, 5, 6,
	4, 9, 73, 70, 3, 2, 2, 2, 74, 77, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 7, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 78, 79, 7, 19, 2, 2,
	79, 80, 7, 3, 2, 2, 80, 83, 5, 6, 4, 2, 81, 82, 7, 6, 2, 2, 82, 84, 5,
	6, 4, 2, 83, 81, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85,
	86, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 88, 7, 4, 2, 2, 88, 101, 3, 2,
	2, 2, 89, 90, 7, 20, 2, 2, 90, 91, 7, 3, 2, 2, 91, 94, 5, 6, 4, 2, 92,
	93, 7, 6, 2, 2, 93, 95, 5, 6, 4, 2, 94, 92, 3, 2, 2, 2, 95, 96, 3, 2, 2,
	2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99,
	7, 4, 2, 2, 99, 101, 3, 2, 2, 2, 100, 78, 3, 2, 2, 2, 100, 89, 3, 2, 2,
	2, 101, 9, 3, 2, 2, 2, 102, 103, 8, 6, 1, 2, 103, 104, 7, 3, 2, 2, 104,
	105, 5, 10, 6, 2, 105, 106, 7, 4, 2, 2, 106, 141, 3, 2, 2, 2, 107, 141,
	5, 12, 7, 2, 108, 109, 7, 25, 2, 2, 109, 110, 5, 18, 10, 2, 110, 111, 7,
	26, 2, 2, 111, 112, 5, 10, 6, 2, 112, 113, 7, 27, 2, 2, 113, 114, 5, 10,
	6, 2, 114, 115, 7, 29, 2, 2, 115, 141, 3, 2, 2, 2, 116, 117, 7, 25, 2,
	2, 117, 118, 5, 18, 10, 2, 118, 119, 7, 26, 2, 2, 119, 125, 5, 10, 6, 2,
	120, 121, 7, 28, 2, 2, 121, 122, 5, 18, 10, 2, 122, 123, 7, 26, 2, 2, 123,
	124, 5, 10, 6, 2, 124, 126, 3, 2, 2, 2, 125, 120, 3, 2, 2, 2, 126, 127,
	3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129, 3, 2,
	2, 2, 129, 130, 7, 27, 2, 2, 130, 131, 5, 10, 6, 2, 131, 132, 7, 29, 2,
	2, 132, 141, 3, 2, 2, 2, 133, 141, 7, 31, 2, 2, 134, 135, 7, 9, 2, 2, 135,
	141, 7, 31, 2, 2, 136, 141, 7, 32, 2, 2, 137, 138, 7, 9, 2, 2, 138, 141,
	7, 32, 2, 2, 139, 141, 7, 35, 2, 2, 140, 102, 3, 2, 2, 2, 140, 107, 3,
	2, 2, 2, 140, 108, 3, 2, 2, 2, 140, 116, 3, 2, 2, 2, 140, 133, 3, 2, 2,
	2, 140, 134, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2, 140,
	139, 3, 2, 2, 2, 141, 156, 3, 2, 2, 2, 142, 143, 12, 14, 2, 2, 143, 144,
	7, 7, 2, 2, 144, 155, 5, 10, 6, 15, 145, 146, 12, 13, 2, 2, 146, 147, 7,
	8, 2, 2, 147, 155, 5, 10, 6, 14, 148, 149, 12, 12, 2, 2, 149, 150, 7, 5,
	2, 2, 150, 155, 5, 10, 6, 13, 151, 152, 12, 11, 2, 2, 152, 153, 7, 9, 2,
	2, 153, 155, 5, 10, 6, 12, 154, 142, 3, 2, 2, 2, 154, 145, 3, 2, 2, 2,
	154, 148, 3, 2, 2, 2, 154, 151, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156,
	154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 11, 3, 2, 2, 2, 158, 156, 3,
	2, 2, 2, 159, 160, 7, 18, 2, 2, 160, 161, 7, 3, 2, 2, 161, 162, 5, 10,
	6, 2, 162, 163, 7, 6, 2, 2, 163, 164, 5, 10, 6, 2, 164, 165, 7, 4, 2, 2,
	165, 189, 3, 2, 2, 2, 166, 167, 7, 19, 2, 2, 167, 168, 7, 3, 2, 2, 168,
	171, 5, 10, 6, 2, 169, 170, 7, 6, 2, 2, 170, 172, 5, 10, 6, 2, 171, 169,
	3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 176, 7, 4, 2, 2, 176, 189, 3, 2, 2, 2,
	177, 178, 7, 20, 2, 2, 178, 179, 7, 3, 2, 2, 179, 182, 5, 10, 6, 2, 180,
	181, 7, 6, 2, 2, 181, 183, 5, 10, 6, 2, 182, 180, 3, 2, 2, 2, 183, 184,
	3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 3, 2,
	2, 2, 186, 187, 7, 4, 2, 2, 187, 189, 3, 2, 2, 2, 188, 159, 3, 2, 2, 2,
	188, 166, 3, 2, 2, 2, 188, 177, 3, 2, 2, 2, 189, 13, 3, 2, 2, 2, 190, 191,
	7, 3, 2, 2, 191, 192, 5, 14, 8, 2, 192, 193, 7, 4, 2, 2, 193, 224, 3, 2,
	2, 2, 194, 195, 7, 25, 2, 2, 195, 196, 5, 18, 10, 2, 196, 197, 7, 26, 2,
	2, 197, 198, 5, 14, 8, 2, 198, 199, 7, 27, 2, 2, 199, 200, 5, 14, 8, 2,
	200, 201, 7, 29, 2, 2, 201, 224, 3, 2, 2, 2, 202, 203, 7, 25, 2, 2, 203,
	204, 5, 18, 10, 2, 204, 205, 7, 26, 2, 2, 205, 211, 5, 14, 8, 2, 206, 207,
	7, 28, 2, 2, 207, 208, 5, 18, 10, 2, 208, 209, 7, 26, 2, 2, 209, 210, 5,
	14, 8, 2, 210, 212, 3, 2, 2, 2, 211, 206, 3, 2, 2, 2, 212, 213, 3, 2, 2,
	2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	216, 7, 27, 2, 2, 216, 217, 5, 14, 8, 2, 217, 218, 7, 29, 2, 2, 218, 224,
	3, 2, 2, 2, 219, 224, 5, 16, 9, 2, 220, 224, 7, 34, 2, 2, 221, 224, 7,
	33, 2, 2, 222, 224, 7, 35, 2, 2, 223, 190, 3, 2, 2, 2, 223, 194, 3, 2,
	2, 2, 223, 202, 3, 2, 2, 2, 223, 219, 3, 2, 2, 2, 223, 220, 3, 2, 2, 2,
	223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 15, 3, 2, 2, 2, 225, 226,
	7, 19, 2, 2, 226, 227, 7, 3, 2, 2, 227, 230, 5, 14, 8, 2, 228, 229, 7,
	6, 2, 2, 229, 231, 5, 14, 8, 2, 230, 228, 3, 2, 2, 2, 231, 232, 3, 2, 2,
	2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234,
	235, 7, 4, 2, 2, 235, 248, 3, 2, 2, 2, 236, 237, 7, 20, 2, 2, 237, 238,
	7, 3, 2, 2, 238, 241, 5, 14, 8, 2, 239, 240, 7, 6, 2, 2, 240, 242, 5, 14,
	8, 2, 241, 239, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2,
	243, 244, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 7, 4, 2, 2, 246,
	248, 3, 2, 2, 2, 247, 225, 3, 2, 2, 2, 247, 236, 3, 2, 2, 2, 248, 17, 3,
	2, 2, 2, 249, 250, 8, 10, 1, 2, 250, 251, 7, 3, 2, 2, 251, 252, 5, 18,
	10, 2, 252, 253, 7, 4, 2, 2, 253, 422, 3, 2, 2, 2, 254, 255, 5, 6, 4, 2,
	255, 256, 7, 10, 2, 2, 256, 257, 5, 6, 4, 2, 257, 422, 3, 2, 2, 2, 258,
	259, 5, 6, 4, 2, 259, 260, 7, 11, 2, 2, 260, 261, 5, 6, 4, 2, 261, 422,
	3, 2, 2, 2, 262, 263, 5, 6, 4, 2, 263, 264, 7, 12, 2, 2, 264, 265, 5, 6,
	4, 2, 265, 422, 3, 2, 2, 2, 266, 267, 5, 6, 4, 2, 267, 268, 7, 13, 2, 2,
	268, 269, 5, 6, 4, 2, 269, 422, 3, 2, 2, 2, 270, 271, 5, 6, 4, 2, 271,
	272, 7, 14, 2, 2, 272, 273, 5, 6, 4, 2, 273, 422, 3, 2, 2, 2, 274, 275,
	5, 6, 4, 2, 275, 276, 7, 15, 2, 2, 276, 277, 5, 6, 4, 2, 277, 422, 3, 2,
	2, 2, 278, 279, 5, 6, 4, 2, 279, 280, 7, 21, 2, 2, 280, 289, 7, 3, 2, 2,
	281, 286, 5, 6, 4, 2, 282, 283, 7, 6, 2, 2, 283, 285, 5, 6, 4, 2, 284,
	282, 3, 2, 2, 2, 285, 288, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 286, 287,
	3, 2, 2, 2, 287, 290, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 289, 281, 3, 2,
	2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 3, 2, 2, 2, 291, 292, 7, 4, 2, 2,
	292, 422, 3, 2, 2, 2, 293, 294, 5, 6, 4, 2, 294, 295, 7, 22, 2, 2, 295,
	296, 7, 21, 2, 2, 296, 305, 7, 3, 2, 2, 297, 302, 5, 6, 4, 2, 298, 299,
	7, 6, 2, 2, 299, 301, 5, 6, 4, 2, 300, 298, 3, 2, 2, 2, 301, 304, 3, 2,
	2, 2, 302, 300, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2,
	304, 302, 3, 2, 2, 2, 305, 297, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306,
	307, 3, 2, 2, 2, 307, 308, 7, 4, 2, 2, 308, 422, 3, 2, 2, 2, 309, 310,
	5, 10, 6, 2, 310, 311, 7, 10, 2, 2, 311, 312, 5, 10, 6, 2, 312, 422, 3,
	2, 2, 2, 313, 314, 5, 10, 6, 2, 314, 315, 7, 11, 2, 2, 315, 316, 5, 10,
	6, 2, 316, 422, 3, 2, 2, 2, 317, 318, 5, 10, 6, 2, 318, 319, 7, 12, 2,
	2, 319, 320, 5, 10, 6, 2, 320, 422, 3, 2, 2, 2, 321, 322, 5, 10, 6, 2,
	322, 323, 7, 13, 2, 2, 323, 324, 5, 10, 6, 2, 324, 422, 3, 2, 2, 2, 325,
	326, 5, 10, 6, 2, 326, 327, 7, 14, 2, 2, 327, 328, 5, 10, 6, 2, 328, 422,
	3, 2, 2, 2, 329, 330, 5, 10, 6, 2, 330, 331, 7, 15, 2, 2, 331, 332, 5,
	10, 6, 2, 332, 422, 3, 2, 2, 2, 333, 334, 5, 10, 6, 2, 334, 335, 7, 21,
	2, 2, 335, 344, 7, 3, 2, 2, 336, 341, 5, 10, 6, 2, 337, 338, 7, 6, 2, 2,
	338, 340, 5, 10, 6, 2, 339, 337, 3, 2, 2, 2, 340, 343, 3, 2, 2, 2, 341,
	339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342, 345, 3, 2, 2, 2, 343, 341,
	3, 2, 2, 2, 344, 336, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 3, 2,
	2, 2, 346, 347, 7, 4, 2, 2, 347, 422, 3, 2, 2, 2, 348, 349, 5, 10, 6, 2,
	349, 350, 7, 22, 2, 2, 350, 351, 7, 21, 2, 2, 351, 360, 7, 3, 2, 2, 352,
	357, 5, 10, 6, 2, 353, 354, 7, 6, 2, 2, 354, 356, 5, 10, 6, 2, 355, 353,
	3, 2, 2, 2, 356, 359, 3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 357, 358, 3, 2,
	2, 2, 358, 361, 3, 2, 2, 2, 359, 357, 3, 2, 2, 2, 360, 352, 3, 2, 2, 2,
	360, 361, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 7, 4, 2, 2, 363,
	422, 3, 2, 2, 2, 364, 365, 5, 14, 8, 2, 365, 366, 7, 10, 2, 2, 366, 367,
	5, 14, 8, 2, 367, 422, 3, 2, 2, 2, 368, 369, 5, 14, 8, 2, 369, 370, 7,
	11, 2, 2, 370, 371, 5, 14, 8, 2, 371, 422, 3, 2, 2, 2, 372, 373, 5, 14,
	8, 2, 373, 374, 7, 12, 2, 2, 374, 375, 5, 14, 8, 2, 375, 422, 3, 2, 2,
	2, 376, 377, 5, 14, 8, 2, 377, 378, 7, 13, 2, 2, 378, 379, 5, 14, 8, 2,
	379, 422, 3, 2, 2, 2, 380, 381, 5, 14, 8, 2, 381, 382, 7, 14, 2, 2, 382,
	383, 5, 14, 8, 2, 383, 422, 3, 2, 2, 2, 384, 385, 5, 14, 8, 2, 385, 386,
	7, 15, 2, 2, 386, 387, 5, 14, 8, 2, 387, 422, 3, 2, 2, 2, 388, 389, 5,
	14, 8, 2, 389, 390, 7, 21, 2, 2, 390, 399, 7, 3, 2, 2, 391, 396, 5, 14,
	8, 2, 392, 393, 7, 6, 2, 2, 393, 395, 5, 14, 8, 2, 394, 392, 3, 2, 2, 2,
	395, 398, 3, 2, 2, 2, 396, 394, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397,
	400, 3, 2, 2, 2, 398, 396, 3, 2, 2, 2, 399, 391, 3, 2, 2, 2, 399, 400,
	3, 2, 2, 2, 400, 401, 3, 2, 2, 2, 401, 402, 7, 4, 2, 2, 402, 422, 3, 2,
	2, 2, 403, 404, 5, 14, 8, 2, 404, 405, 7, 22, 2, 2, 405, 406, 7, 21, 2,
	2, 406, 415, 7, 3, 2, 2, 407, 412, 5, 14, 8, 2, 408, 409, 7, 6, 2, 2, 409,
	411, 5, 14, 8, 2, 410, 408, 3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 410,
	3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 416, 3, 2, 2, 2, 414, 412, 3, 2,
	2, 2, 415, 407, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2,
	417, 418, 7, 4, 2, 2, 418, 422, 3, 2, 2, 2, 419, 422, 7, 30, 2, 2, 420,
	422, 7, 35, 2, 2, 421, 249, 3, 2, 2, 2, 421, 254, 3, 2, 2, 2, 421, 258,
	3, 2, 2, 2, 421, 262, 3, 2, 2, 2, 421, 266, 3, 2, 2, 2, 421, 270, 3, 2,
	2, 2, 421, 274, 3, 2, 2, 2, 421, 278, 3, 2, 2, 2, 421, 293, 3, 2, 2, 2,
	421, 309, 3, 2, 2, 2, 421, 313, 3, 2, 2, 2, 421, 317, 3, 2, 2, 2, 421,
	321, 3, 2, 2, 2, 421, 325, 3, 2, 2, 2, 421, 329, 3, 2, 2, 2, 421, 333,
	3, 2, 2, 2, 421, 348, 3, 2, 2, 2, 421, 364, 3, 2, 2, 2, 421, 368, 3, 2,
	2, 2, 421, 372, 3, 2, 2, 2, 421, 376, 3, 2, 2, 2, 421, 380, 3, 2, 2, 2,
	421, 384, 3, 2, 2, 2, 421, 388, 3, 2, 2, 2, 421, 403, 3, 2, 2, 2, 421,
	419, 3, 2, 2, 2, 421, 420, 3, 2, 2, 2, 422, 431, 3, 2, 2, 2, 423, 424,
	12, 6, 2, 2, 424, 425, 9, 2, 2, 2, 425, 430, 5, 18, 10, 7, 426, 427, 12,
	5, 2, 2, 427, 428, 9, 3, 2, 2, 428, 430, 5, 18, 10, 6, 429, 423, 3, 2,
	2, 2, 429, 426, 3, 2, 2, 2, 430, 433, 3, 2, 2, 2, 431, 429, 3, 2, 2, 2,
	431, 432, 3, 2, 2, 2, 432, 19, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 434, 435,
	9, 4, 2, 2, 435, 21, 3, 2, 2, 2, 37, 27, 33, 59, 68, 75, 85, 96, 100, 127,
	140, 154, 156, 173, 184, 188, 213, 223, 232, 243, 247, 286, 289, 302, 305,
	341, 344, 357, 360, 396, 399, 412, 415, 421, 429, 431,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'+'", "','", "'*'", "'/'", "'-'", "'='", "'>'", "'>='",
	"'<'", "'<='", "'!='", "'&&'", "'||'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Pow",
	"Min", "Max", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif",
	"Endif", "Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace",
}

var ruleNames = []string{
	"formula", "fieldExpr", "stringExpr", "stringFunction", "numberExpr", "numberFunction",
	"dateExpr", "dateFunction", "boolExpr", "str",
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
	AlteryxFormulasParserPow               = 16
	AlteryxFormulasParserMin               = 17
	AlteryxFormulasParserMax               = 18
	AlteryxFormulasParserIn                = 19
	AlteryxFormulasParserNot               = 20
	AlteryxFormulasParserAnd               = 21
	AlteryxFormulasParserOr                = 22
	AlteryxFormulasParserIf                = 23
	AlteryxFormulasParserThen              = 24
	AlteryxFormulasParserElse              = 25
	AlteryxFormulasParserElseif            = 26
	AlteryxFormulasParserEndif             = 27
	AlteryxFormulasParserBool              = 28
	AlteryxFormulasParserInteger           = 29
	AlteryxFormulasParserDecimal           = 30
	AlteryxFormulasParserDate              = 31
	AlteryxFormulasParserDatetime          = 32
	AlteryxFormulasParserField             = 33
	AlteryxFormulasParserSingleQuoteString = 34
	AlteryxFormulasParserDoubleQuoteString = 35
	AlteryxFormulasParserWhitespace        = 36
)

// AlteryxFormulasParser rules.
const (
	AlteryxFormulasParserRULE_formula        = 0
	AlteryxFormulasParserRULE_fieldExpr      = 1
	AlteryxFormulasParserRULE_stringExpr     = 2
	AlteryxFormulasParserRULE_stringFunction = 3
	AlteryxFormulasParserRULE_numberExpr     = 4
	AlteryxFormulasParserRULE_numberFunction = 5
	AlteryxFormulasParserRULE_dateExpr       = 6
	AlteryxFormulasParserRULE_dateFunction   = 7
	AlteryxFormulasParserRULE_boolExpr       = 8
	AlteryxFormulasParserRULE_str            = 9
)

// IFormulaContext is an interface to support dynamic dispatch.
type IFormulaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormulaContext differentiates from other interfaces.
	IsFormulaContext()
}

type FormulaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormulaContext() *FormulaContext {
	var p = new(FormulaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_formula
	return p
}

func (*FormulaContext) IsFormulaContext() {}

func NewFormulaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormulaContext {
	var p = new(FormulaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_formula

	return p
}

func (s *FormulaContext) GetParser() antlr.Parser { return s.parser }

func (s *FormulaContext) FieldExpr() IFieldExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldExprContext)
}

func (s *FormulaContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *FormulaContext) DateExpr() IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *FormulaContext) StringExpr() IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *FormulaContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *FormulaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormulaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormulaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFormula(s)
	}
}

func (s *FormulaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFormula(s)
	}
}

func (p *AlteryxFormulasParser) Formula() (localctx IFormulaContext) {
	localctx = NewFormulaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AlteryxFormulasParserRULE_formula)

	defer func() {
		p.ExitRule()
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

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.FieldExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.numberExpr(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.DateExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(23)
			p.stringExpr(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(24)
			p.boolExpr(0)
		}

	}

	return localctx
}

// IFieldExprContext is an interface to support dynamic dispatch.
type IFieldExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldExprContext differentiates from other interfaces.
	IsFieldExprContext()
}

type FieldExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExprContext() *FieldExprContext {
	var p = new(FieldExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_fieldExpr
	return p
}

func (*FieldExprContext) IsFieldExprContext() {}

func NewFieldExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExprContext {
	var p = new(FieldExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_fieldExpr

	return p
}

func (s *FieldExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExprContext) CopyFrom(ctx *FieldExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FieldParenthesisContext struct {
	*FieldExprContext
}

func NewFieldParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldParenthesisContext {
	var p = new(FieldParenthesisContext)

	p.FieldExprContext = NewEmptyFieldExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldExprContext))

	return p
}

func (s *FieldParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldParenthesisContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *FieldParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFieldParenthesis(s)
	}
}

func (s *FieldParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFieldParenthesis(s)
	}
}

type AnyFieldContext struct {
	*FieldExprContext
}

func NewAnyFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AnyFieldContext {
	var p = new(AnyFieldContext)

	p.FieldExprContext = NewEmptyFieldExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FieldExprContext))

	return p
}

func (s *AnyFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *AnyFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAnyField(s)
	}
}

func (s *AnyFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAnyField(s)
	}
}

func (p *AlteryxFormulasParser) FieldExpr() (localctx IFieldExprContext) {
	localctx = NewFieldExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AlteryxFormulasParserRULE_fieldExpr)

	defer func() {
		p.ExitRule()
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

	p.SetState(31)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlteryxFormulasParserT__0:
		localctx = NewFieldParenthesisContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(28)
			p.Match(AlteryxFormulasParserField)
		}
		{
			p.SetState(29)
			p.Match(AlteryxFormulasParserT__1)
		}

	case AlteryxFormulasParserField:
		localctx = NewAnyFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(AlteryxFormulasParserField)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringExprContext is an interface to support dynamic dispatch.
type IStringExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringExprContext differentiates from other interfaces.
	IsStringExprContext()
}

type StringExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringExprContext() *StringExprContext {
	var p = new(StringExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_stringExpr
	return p
}

func (*StringExprContext) IsStringExprContext() {}

func NewStringExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringExprContext {
	var p = new(StringExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_stringExpr

	return p
}

func (s *StringExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StringExprContext) CopyFrom(ctx *StringExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StringExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringIfContext struct {
	*StringExprContext
}

func NewStringIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringIfContext {
	var p = new(StringIfContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *StringIfContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *StringIfContext) Then() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, 0)
}

func (s *StringIfContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringIfContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *StringIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *StringIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringIf(s)
	}
}

func (s *StringIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringIf(s)
	}
}

type StringFuncContext struct {
	*StringExprContext
}

func NewStringFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringFuncContext {
	var p = new(StringFuncContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringFuncContext) StringFunction() IStringFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringFunctionContext)
}

func (s *StringFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringFunc(s)
	}
}

func (s *StringFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringFunc(s)
	}
}

type ConcatenateContext struct {
	*StringExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewConcatenateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConcatenateContext {
	var p = new(ConcatenateContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *ConcatenateContext) GetLeft() IStringExprContext { return s.left }

func (s *ConcatenateContext) GetRight() IStringExprContext { return s.right }

func (s *ConcatenateContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *ConcatenateContext) SetRight(v IStringExprContext) { s.right = v }

func (s *ConcatenateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcatenateContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *ConcatenateContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *ConcatenateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterConcatenate(s)
	}
}

func (s *ConcatenateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitConcatenate(s)
	}
}

type StringLiteralContext struct {
	*StringExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) Str() IStrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStrContext)
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

type StringParenthesisContext struct {
	*StringExprContext
}

func NewStringParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringParenthesisContext {
	var p = new(StringParenthesisContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringParenthesisContext) StringExpr() IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringParenthesis(s)
	}
}

func (s *StringParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringParenthesis(s)
	}
}

type StringElseIfContext struct {
	*StringExprContext
}

func NewStringElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringElseIfContext {
	var p = new(StringElseIfContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringElseIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *StringElseIfContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *StringElseIfContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *StringElseIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *StringElseIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *StringElseIfContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringElseIfContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringElseIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *StringElseIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *StringElseIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *StringElseIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *StringElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringElseIf(s)
	}
}

func (s *StringElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringElseIf(s)
	}
}

type StringFieldContext struct {
	*StringExprContext
}

func NewStringFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringFieldContext {
	var p = new(StringFieldContext)

	p.StringExprContext = NewEmptyStringExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringExprContext))

	return p
}

func (s *StringFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *StringFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringField(s)
	}
}

func (s *StringFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringField(s)
	}
}

func (p *AlteryxFormulasParser) StringExpr() (localctx IStringExprContext) {
	return p.stringExpr(0)
}

func (p *AlteryxFormulasParser) stringExpr(_p int) (localctx IStringExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStringExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStringExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, AlteryxFormulasParserRULE_stringExpr, _p)
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStringParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(34)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(35)
			p.stringExpr(0)
		}
		{
			p.SetState(36)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewStringIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(38)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(39)
			p.boolExpr(0)
		}
		{
			p.SetState(40)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(41)
			p.stringExpr(0)
		}
		{
			p.SetState(42)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(43)
			p.stringExpr(0)
		}
		{
			p.SetState(44)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 3:
		localctx = NewStringElseIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(46)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(47)
			p.boolExpr(0)
		}
		{
			p.SetState(48)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(49)
			p.stringExpr(0)
		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserElseif {
			{
				p.SetState(50)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(51)
				p.boolExpr(0)
			}
			{
				p.SetState(52)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(53)
				p.stringExpr(0)
			}

			p.SetState(57)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(59)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(60)
			p.stringExpr(0)
		}
		{
			p.SetState(61)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 4:
		localctx = NewStringFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)
			p.StringFunction()
		}

	case 5:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Str()
		}

	case 6:
		localctx = NewStringFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConcatenateContext(p, NewStringExprContext(p, _parentctx, _parentState))
			localctx.(*ConcatenateContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_stringExpr)
			p.SetState(68)

			if !(p.Precpred(p.GetParserRuleContext(), 6)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
			}
			{
				p.SetState(69)
				p.Match(AlteryxFormulasParserT__2)
			}
			{
				p.SetState(70)

				var _x = p.stringExpr(7)

				localctx.(*ConcatenateContext).right = _x
			}

		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IStringFunctionContext is an interface to support dynamic dispatch.
type IStringFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringFunctionContext differentiates from other interfaces.
	IsStringFunctionContext()
}

type StringFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringFunctionContext() *StringFunctionContext {
	var p = new(StringFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_stringFunction
	return p
}

func (*StringFunctionContext) IsStringFunctionContext() {}

func NewStringFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringFunctionContext {
	var p = new(StringFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_stringFunction

	return p
}

func (s *StringFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *StringFunctionContext) CopyFrom(ctx *StringFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StringFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringMaxContext struct {
	*StringFunctionContext
}

func NewStringMaxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringMaxContext {
	var p = new(StringMaxContext)

	p.StringFunctionContext = NewEmptyStringFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringFunctionContext))

	return p
}

func (s *StringMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringMaxContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *StringMaxContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringMaxContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringMax(s)
	}
}

func (s *StringMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringMax(s)
	}
}

type StringMinContext struct {
	*StringFunctionContext
}

func NewStringMinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringMinContext {
	var p = new(StringMinContext)

	p.StringFunctionContext = NewEmptyStringFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StringFunctionContext))

	return p
}

func (s *StringMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringMinContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *StringMinContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringMinContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringMin(s)
	}
}

func (s *StringMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringMin(s)
	}
}

func (p *AlteryxFormulasParser) StringFunction() (localctx IStringFunctionContext) {
	localctx = NewStringFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AlteryxFormulasParserRULE_stringFunction)
	var _la int

	defer func() {
		p.ExitRule()
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlteryxFormulasParserMin:
		localctx = NewStringMinContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(77)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(78)
			p.stringExpr(0)
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(79)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(80)
				p.stringExpr(0)
			}

			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(85)
			p.Match(AlteryxFormulasParserT__1)
		}

	case AlteryxFormulasParserMax:
		localctx = NewStringMaxContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(88)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(89)
			p.stringExpr(0)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(90)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(91)
				p.stringExpr(0)
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(96)
			p.Match(AlteryxFormulasParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberExprContext is an interface to support dynamic dispatch.
type INumberExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberExprContext differentiates from other interfaces.
	IsNumberExprContext()
}

type NumberExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberExprContext() *NumberExprContext {
	var p = new(NumberExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_numberExpr
	return p
}

func (*NumberExprContext) IsNumberExprContext() {}

func NewNumberExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberExprContext {
	var p = new(NumberExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_numberExpr

	return p
}

func (s *NumberExprContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberExprContext) CopyFrom(ctx *NumberExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumberExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberParenthesisContext struct {
	*NumberExprContext
}

func NewNumberParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberParenthesisContext {
	var p = new(NumberParenthesisContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *NumberParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberParenthesisContext) NumberExpr() INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberParenthesis(s)
	}
}

func (s *NumberParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberParenthesis(s)
	}
}

type AddContext struct {
	*NumberExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *AddContext) GetLeft() INumberExprContext { return s.left }

func (s *AddContext) GetRight() INumberExprContext { return s.right }

func (s *AddContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *AddContext) SetRight(v INumberExprContext) { s.right = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *AddContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
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

type NumberFuncContext struct {
	*NumberExprContext
}

func NewNumberFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberFuncContext {
	var p = new(NumberFuncContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *NumberFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberFuncContext) NumberFunction() INumberFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberFunctionContext)
}

func (s *NumberFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberFunc(s)
	}
}

func (s *NumberFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberFunc(s)
	}
}

type NumberFieldContext struct {
	*NumberExprContext
}

func NewNumberFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberFieldContext {
	var p = new(NumberFieldContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *NumberFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *NumberFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberField(s)
	}
}

func (s *NumberFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberField(s)
	}
}

type SubtractContext struct {
	*NumberExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewSubtractContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractContext {
	var p = new(SubtractContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *SubtractContext) GetLeft() INumberExprContext { return s.left }

func (s *SubtractContext) GetRight() INumberExprContext { return s.right }

func (s *SubtractContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *SubtractContext) SetRight(v INumberExprContext) { s.right = v }

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *SubtractContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
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

type NumberIfContext struct {
	*NumberExprContext
}

func NewNumberIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberIfContext {
	var p = new(NumberIfContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *NumberIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *NumberIfContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *NumberIfContext) Then() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, 0)
}

func (s *NumberIfContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberIfContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *NumberIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *NumberIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberIf(s)
	}
}

func (s *NumberIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberIf(s)
	}
}

type NumberElseIfContext struct {
	*NumberExprContext
}

func NewNumberElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberElseIfContext {
	var p = new(NumberElseIfContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *NumberElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberElseIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *NumberElseIfContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *NumberElseIfContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *NumberElseIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *NumberElseIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *NumberElseIfContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberElseIfContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberElseIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *NumberElseIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *NumberElseIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *NumberElseIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *NumberElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberElseIf(s)
	}
}

func (s *NumberElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberElseIf(s)
	}
}

type DivideContext struct {
	*NumberExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *DivideContext) GetLeft() INumberExprContext { return s.left }

func (s *DivideContext) GetRight() INumberExprContext { return s.right }

func (s *DivideContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *DivideContext) SetRight(v INumberExprContext) { s.right = v }

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *DivideContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
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

type IntegerContext struct {
	*NumberExprContext
}

func NewIntegerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerContext {
	var p = new(IntegerContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) Integer() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserInteger, 0)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitInteger(s)
	}
}

type DecimalContext struct {
	*NumberExprContext
}

func NewDecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalContext {
	var p = new(DecimalContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalContext) Decimal() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDecimal, 0)
}

func (s *DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDecimal(s)
	}
}

func (s *DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDecimal(s)
	}
}

type MultiplyContext struct {
	*NumberExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.NumberExprContext = NewEmptyNumberExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberExprContext))

	return p
}

func (s *MultiplyContext) GetLeft() INumberExprContext { return s.left }

func (s *MultiplyContext) GetRight() INumberExprContext { return s.right }

func (s *MultiplyContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *MultiplyContext) SetRight(v INumberExprContext) { s.right = v }

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *MultiplyContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
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

func (p *AlteryxFormulasParser) NumberExpr() (localctx INumberExprContext) {
	return p.numberExpr(0)
}

func (p *AlteryxFormulasParser) numberExpr(_p int) (localctx INumberExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewNumberExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx INumberExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, AlteryxFormulasParserRULE_numberExpr, _p)
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
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(101)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(102)
			p.numberExpr(0)
		}
		{
			p.SetState(103)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewNumberFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			p.NumberFunction()
		}

	case 3:
		localctx = NewNumberIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(106)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(107)
			p.boolExpr(0)
		}
		{
			p.SetState(108)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(109)
			p.numberExpr(0)
		}
		{
			p.SetState(110)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(111)
			p.numberExpr(0)
		}
		{
			p.SetState(112)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 4:
		localctx = NewNumberElseIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(115)
			p.boolExpr(0)
		}
		{
			p.SetState(116)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(117)
			p.numberExpr(0)
		}
		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserElseif {
			{
				p.SetState(118)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(119)
				p.boolExpr(0)
			}
			{
				p.SetState(120)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(121)
				p.numberExpr(0)
			}

			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(127)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(128)
			p.numberExpr(0)
		}
		{
			p.SetState(129)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 5:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 6:
		localctx = NewIntegerContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			p.Match(AlteryxFormulasParserT__6)
		}
		{
			p.SetState(133)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 7:
		localctx = NewDecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(134)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 8:
		localctx = NewDecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(135)
			p.Match(AlteryxFormulasParserT__6)
		}
		{
			p.SetState(136)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 9:
		localctx = NewNumberFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(137)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplyContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_numberExpr)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(141)
					p.Match(AlteryxFormulasParserT__4)
				}
				{
					p.SetState(142)

					var _x = p.numberExpr(13)

					localctx.(*MultiplyContext).right = _x
				}

			case 2:
				localctx = NewDivideContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				localctx.(*DivideContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_numberExpr)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(144)
					p.Match(AlteryxFormulasParserT__5)
				}
				{
					p.SetState(145)

					var _x = p.numberExpr(12)

					localctx.(*DivideContext).right = _x
				}

			case 3:
				localctx = NewAddContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				localctx.(*AddContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_numberExpr)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(147)
					p.Match(AlteryxFormulasParserT__2)
				}
				{
					p.SetState(148)

					var _x = p.numberExpr(11)

					localctx.(*AddContext).right = _x
				}

			case 4:
				localctx = NewSubtractContext(p, NewNumberExprContext(p, _parentctx, _parentState))
				localctx.(*SubtractContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_numberExpr)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(150)
					p.Match(AlteryxFormulasParserT__6)
				}
				{
					p.SetState(151)

					var _x = p.numberExpr(10)

					localctx.(*SubtractContext).right = _x
				}

			}

		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// INumberFunctionContext is an interface to support dynamic dispatch.
type INumberFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberFunctionContext differentiates from other interfaces.
	IsNumberFunctionContext()
}

type NumberFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberFunctionContext() *NumberFunctionContext {
	var p = new(NumberFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_numberFunction
	return p
}

func (*NumberFunctionContext) IsNumberFunctionContext() {}

func NewNumberFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberFunctionContext {
	var p = new(NumberFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_numberFunction

	return p
}

func (s *NumberFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberFunctionContext) CopyFrom(ctx *NumberFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumberFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PowContext struct {
	*NumberFunctionContext
}

func NewPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowContext {
	var p = new(PowContext)

	p.NumberFunctionContext = NewEmptyNumberFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberFunctionContext))

	return p
}

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) Pow() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPow, 0)
}

func (s *PowContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *PowContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *PowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPow(s)
	}
}

func (s *PowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPow(s)
	}
}

type NumberMaxContext struct {
	*NumberFunctionContext
}

func NewNumberMaxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberMaxContext {
	var p = new(NumberMaxContext)

	p.NumberFunctionContext = NewEmptyNumberFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberFunctionContext))

	return p
}

func (s *NumberMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberMaxContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *NumberMaxContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberMaxContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberMax(s)
	}
}

func (s *NumberMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberMax(s)
	}
}

type NumberMinContext struct {
	*NumberFunctionContext
}

func NewNumberMinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberMinContext {
	var p = new(NumberMinContext)

	p.NumberFunctionContext = NewEmptyNumberFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberFunctionContext))

	return p
}

func (s *NumberMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberMinContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *NumberMinContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberMinContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberMin(s)
	}
}

func (s *NumberMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberMin(s)
	}
}

func (p *AlteryxFormulasParser) NumberFunction() (localctx INumberFunctionContext) {
	localctx = NewNumberFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AlteryxFormulasParserRULE_numberFunction)
	var _la int

	defer func() {
		p.ExitRule()
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

	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlteryxFormulasParserPow:
		localctx = NewPowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.Match(AlteryxFormulasParserPow)
		}
		{
			p.SetState(158)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(159)
			p.numberExpr(0)
		}
		{
			p.SetState(160)
			p.Match(AlteryxFormulasParserT__3)
		}
		{
			p.SetState(161)
			p.numberExpr(0)
		}
		{
			p.SetState(162)
			p.Match(AlteryxFormulasParserT__1)
		}

	case AlteryxFormulasParserMin:
		localctx = NewNumberMinContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(165)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(166)
			p.numberExpr(0)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(167)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(168)
				p.numberExpr(0)
			}

			p.SetState(171)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(173)
			p.Match(AlteryxFormulasParserT__1)
		}

	case AlteryxFormulasParserMax:
		localctx = NewNumberMaxContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
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
			p.numberExpr(0)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(178)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(179)
				p.numberExpr(0)
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(184)
			p.Match(AlteryxFormulasParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDateExprContext is an interface to support dynamic dispatch.
type IDateExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateExprContext differentiates from other interfaces.
	IsDateExprContext()
}

type DateExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateExprContext() *DateExprContext {
	var p = new(DateExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_dateExpr
	return p
}

func (*DateExprContext) IsDateExprContext() {}

func NewDateExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateExprContext {
	var p = new(DateExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_dateExpr

	return p
}

func (s *DateExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DateExprContext) CopyFrom(ctx *DateExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DateParenthesisContext struct {
	*DateExprContext
}

func NewDateParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateParenthesisContext {
	var p = new(DateParenthesisContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

	return p
}

func (s *DateParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateParenthesisContext) DateExpr() IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateParenthesis(s)
	}
}

func (s *DateParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateParenthesis(s)
	}
}

type DatetimeLiteralContext struct {
	*DateExprContext
}

func NewDatetimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatetimeLiteralContext {
	var p = new(DatetimeLiteralContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

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

type DateElseIfContext struct {
	*DateExprContext
}

func NewDateElseIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateElseIfContext {
	var p = new(DateElseIfContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

	return p
}

func (s *DateElseIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateElseIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *DateElseIfContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *DateElseIfContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *DateElseIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *DateElseIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *DateElseIfContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateElseIfContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateElseIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *DateElseIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *DateElseIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *DateElseIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *DateElseIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateElseIf(s)
	}
}

func (s *DateElseIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateElseIf(s)
	}
}

type DateFuncContext struct {
	*DateExprContext
}

func NewDateFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateFuncContext {
	var p = new(DateFuncContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

	return p
}

func (s *DateFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateFuncContext) DateFunction() IDateFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateFunctionContext)
}

func (s *DateFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateFunc(s)
	}
}

func (s *DateFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateFunc(s)
	}
}

type DateFieldContext struct {
	*DateExprContext
}

func NewDateFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateFieldContext {
	var p = new(DateFieldContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

	return p
}

func (s *DateFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *DateFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateField(s)
	}
}

func (s *DateFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateField(s)
	}
}

type DateIfContext struct {
	*DateExprContext
}

func NewDateIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateIfContext {
	var p = new(DateIfContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

	return p
}

func (s *DateIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *DateIfContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *DateIfContext) Then() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, 0)
}

func (s *DateIfContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateIfContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *DateIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *DateIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateIf(s)
	}
}

func (s *DateIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateIf(s)
	}
}

type DateLiteralContext struct {
	*DateExprContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.DateExprContext = NewEmptyDateExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateExprContext))

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

func (p *AlteryxFormulasParser) DateExpr() (localctx IDateExprContext) {
	localctx = NewDateExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AlteryxFormulasParserRULE_dateExpr)
	var _la int

	defer func() {
		p.ExitRule()
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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDateParenthesisContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(189)
			p.DateExpr()
		}
		{
			p.SetState(190)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewDateIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(193)
			p.boolExpr(0)
		}
		{
			p.SetState(194)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(195)
			p.DateExpr()
		}
		{
			p.SetState(196)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(197)
			p.DateExpr()
		}
		{
			p.SetState(198)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 3:
		localctx = NewDateElseIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(201)
			p.boolExpr(0)
		}
		{
			p.SetState(202)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(203)
			p.DateExpr()
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserElseif {
			{
				p.SetState(204)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(205)
				p.boolExpr(0)
			}
			{
				p.SetState(206)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(207)
				p.DateExpr()
			}

			p.SetState(211)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(213)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(214)
			p.DateExpr()
		}
		{
			p.SetState(215)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 4:
		localctx = NewDateFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.DateFunction()
		}

	case 5:
		localctx = NewDatetimeLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(218)
			p.Match(AlteryxFormulasParserDatetime)
		}

	case 6:
		localctx = NewDateLiteralContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(219)
			p.Match(AlteryxFormulasParserDate)
		}

	case 7:
		localctx = NewDateFieldContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(220)
			p.Match(AlteryxFormulasParserField)
		}

	}

	return localctx
}

// IDateFunctionContext is an interface to support dynamic dispatch.
type IDateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateFunctionContext differentiates from other interfaces.
	IsDateFunctionContext()
}

type DateFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateFunctionContext() *DateFunctionContext {
	var p = new(DateFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_dateFunction
	return p
}

func (*DateFunctionContext) IsDateFunctionContext() {}

func NewDateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateFunctionContext {
	var p = new(DateFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_dateFunction

	return p
}

func (s *DateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *DateFunctionContext) CopyFrom(ctx *DateFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DateMaxContext struct {
	*DateFunctionContext
}

func NewDateMaxContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateMaxContext {
	var p = new(DateMaxContext)

	p.DateFunctionContext = NewEmptyDateFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateFunctionContext))

	return p
}

func (s *DateMaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateMaxContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *DateMaxContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateMaxContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateMaxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateMax(s)
	}
}

func (s *DateMaxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateMax(s)
	}
}

type DateMinContext struct {
	*DateFunctionContext
}

func NewDateMinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateMinContext {
	var p = new(DateMinContext)

	p.DateFunctionContext = NewEmptyDateFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DateFunctionContext))

	return p
}

func (s *DateMinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateMinContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *DateMinContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateMinContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateMinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateMin(s)
	}
}

func (s *DateMinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateMin(s)
	}
}

func (p *AlteryxFormulasParser) DateFunction() (localctx IDateFunctionContext) {
	localctx = NewDateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AlteryxFormulasParserRULE_dateFunction)
	var _la int

	defer func() {
		p.ExitRule()
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

	p.SetState(245)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AlteryxFormulasParserMin:
		localctx = NewDateMinContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(224)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(225)
			p.DateExpr()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(226)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(227)
				p.DateExpr()
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(232)
			p.Match(AlteryxFormulasParserT__1)
		}

	case AlteryxFormulasParserMax:
		localctx = NewDateMaxContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(234)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(235)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(236)
			p.DateExpr()
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__3 {
			{
				p.SetState(237)
				p.Match(AlteryxFormulasParserT__3)
			}
			{
				p.SetState(238)
				p.DateExpr()
			}

			p.SetState(241)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(243)
			p.Match(AlteryxFormulasParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprContext differentiates from other interfaces.
	IsBoolExprContext()
}

type BoolExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprContext() *BoolExprContext {
	var p = new(BoolExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_boolExpr
	return p
}

func (*BoolExprContext) IsBoolExprContext() {}

func NewBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprContext {
	var p = new(BoolExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_boolExpr

	return p
}

func (s *BoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprContext) CopyFrom(ctx *BoolExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberInContext struct {
	*BoolExprContext
}

func NewNumberInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberInContext {
	var p = new(NumberInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberInContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberInContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *NumberInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberIn(s)
	}
}

func (s *NumberInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberIn(s)
	}
}

type StringGreaterThanContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringGreaterThanContext {
	var p = new(StringGreaterThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringGreaterThanContext) GetLeft() IStringExprContext { return s.left }

func (s *StringGreaterThanContext) GetRight() IStringExprContext { return s.right }

func (s *StringGreaterThanContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringGreaterThanContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringGreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringGreaterThanContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringGreaterThanContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringGreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringGreaterThan(s)
	}
}

func (s *StringGreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringGreaterThan(s)
	}
}

type DateInContext struct {
	*BoolExprContext
}

func NewDateInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateInContext {
	var p = new(DateInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateInContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateInContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *DateInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateIn(s)
	}
}

func (s *DateInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateIn(s)
	}
}

type BoolParenthesisContext struct {
	*BoolExprContext
}

func NewBoolParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolParenthesisContext {
	var p = new(BoolParenthesisContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *BoolParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolParenthesisContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *BoolParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterBoolParenthesis(s)
	}
}

func (s *BoolParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitBoolParenthesis(s)
	}
}

type NumberNotInContext struct {
	*BoolExprContext
}

func NewNumberNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberNotInContext {
	var p = new(NumberNotInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberNotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberNotInContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberNotInContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberNotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *NumberNotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *NumberNotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberNotIn(s)
	}
}

func (s *NumberNotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberNotIn(s)
	}
}

type BoolLiteralContext struct {
	*BoolExprContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

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

type StringNotInContext struct {
	*BoolExprContext
}

func NewStringNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringNotInContext {
	var p = new(StringNotInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringNotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringNotInContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringNotInContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringNotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *StringNotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *StringNotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringNotIn(s)
	}
}

func (s *StringNotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringNotIn(s)
	}
}

type BoolFieldContext struct {
	*BoolExprContext
}

func NewBoolFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolFieldContext {
	var p = new(BoolFieldContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *BoolFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *BoolFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterBoolField(s)
	}
}

func (s *BoolFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitBoolField(s)
	}
}

type DateGreaterEqualContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateGreaterEqualContext {
	var p = new(DateGreaterEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateGreaterEqualContext) GetLeft() IDateExprContext { return s.left }

func (s *DateGreaterEqualContext) GetRight() IDateExprContext { return s.right }

func (s *DateGreaterEqualContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateGreaterEqualContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateGreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateGreaterEqualContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateGreaterEqualContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateGreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateGreaterEqual(s)
	}
}

func (s *DateGreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateGreaterEqual(s)
	}
}

type DateLessThanContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLessThanContext {
	var p = new(DateLessThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateLessThanContext) GetLeft() IDateExprContext { return s.left }

func (s *DateLessThanContext) GetRight() IDateExprContext { return s.right }

func (s *DateLessThanContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateLessThanContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateLessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLessThanContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateLessThanContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateLessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateLessThan(s)
	}
}

func (s *DateLessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateLessThan(s)
	}
}

type AndContext struct {
	*BoolExprContext
	left  IBoolExprContext
	right IBoolExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *AndContext) GetLeft() IBoolExprContext { return s.left }

func (s *AndContext) GetRight() IBoolExprContext { return s.right }

func (s *AndContext) SetLeft(v IBoolExprContext) { s.left = v }

func (s *AndContext) SetRight(v IBoolExprContext) { s.right = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *AndContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
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

type DateNotEqualContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateNotEqualContext {
	var p = new(DateNotEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateNotEqualContext) GetLeft() IDateExprContext { return s.left }

func (s *DateNotEqualContext) GetRight() IDateExprContext { return s.right }

func (s *DateNotEqualContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateNotEqualContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateNotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateNotEqualContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateNotEqualContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateNotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateNotEqual(s)
	}
}

func (s *DateNotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateNotEqual(s)
	}
}

type DateEqualContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateEqualContext {
	var p = new(DateEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateEqualContext) GetLeft() IDateExprContext { return s.left }

func (s *DateEqualContext) GetRight() IDateExprContext { return s.right }

func (s *DateEqualContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateEqualContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateEqualContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateEqualContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateEqual(s)
	}
}

func (s *DateEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateEqual(s)
	}
}

type StringLessEqualContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLessEqualContext {
	var p = new(StringLessEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringLessEqualContext) GetLeft() IStringExprContext { return s.left }

func (s *StringLessEqualContext) GetRight() IStringExprContext { return s.right }

func (s *StringLessEqualContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringLessEqualContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringLessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLessEqualContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringLessEqualContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringLessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringLessEqual(s)
	}
}

func (s *StringLessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringLessEqual(s)
	}
}

type NumberEqualContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberEqualContext {
	var p = new(NumberEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberEqualContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberEqualContext) GetRight() INumberExprContext { return s.right }

func (s *NumberEqualContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberEqualContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberEqualContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberEqualContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberEqual(s)
	}
}

func (s *NumberEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberEqual(s)
	}
}

type StringLessThanContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLessThanContext {
	var p = new(StringLessThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringLessThanContext) GetLeft() IStringExprContext { return s.left }

func (s *StringLessThanContext) GetRight() IStringExprContext { return s.right }

func (s *StringLessThanContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringLessThanContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringLessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLessThanContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringLessThanContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringLessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringLessThan(s)
	}
}

func (s *StringLessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringLessThan(s)
	}
}

type NumberLessEqualContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLessEqualContext {
	var p = new(NumberLessEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberLessEqualContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberLessEqualContext) GetRight() INumberExprContext { return s.right }

func (s *NumberLessEqualContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberLessEqualContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberLessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLessEqualContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberLessEqualContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberLessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberLessEqual(s)
	}
}

func (s *NumberLessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberLessEqual(s)
	}
}

type OrContext struct {
	*BoolExprContext
	left  IBoolExprContext
	right IBoolExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *OrContext) GetLeft() IBoolExprContext { return s.left }

func (s *OrContext) GetRight() IBoolExprContext { return s.right }

func (s *OrContext) SetLeft(v IBoolExprContext) { s.left = v }

func (s *OrContext) SetRight(v IBoolExprContext) { s.right = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *OrContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
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

type NumberGreaterThanContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberGreaterThanContext {
	var p = new(NumberGreaterThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberGreaterThanContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberGreaterThanContext) GetRight() INumberExprContext { return s.right }

func (s *NumberGreaterThanContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberGreaterThanContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberGreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberGreaterThanContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberGreaterThanContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberGreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberGreaterThan(s)
	}
}

func (s *NumberGreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberGreaterThan(s)
	}
}

type DateLessEqualContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLessEqualContext {
	var p = new(DateLessEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateLessEqualContext) GetLeft() IDateExprContext { return s.left }

func (s *DateLessEqualContext) GetRight() IDateExprContext { return s.right }

func (s *DateLessEqualContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateLessEqualContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateLessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLessEqualContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateLessEqualContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateLessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateLessEqual(s)
	}
}

func (s *DateLessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateLessEqual(s)
	}
}

type StringEqualContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringEqualContext {
	var p = new(StringEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringEqualContext) GetLeft() IStringExprContext { return s.left }

func (s *StringEqualContext) GetRight() IStringExprContext { return s.right }

func (s *StringEqualContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringEqualContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringEqualContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringEqualContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringEqual(s)
	}
}

func (s *StringEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringEqual(s)
	}
}

type StringInContext struct {
	*BoolExprContext
}

func NewStringInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringInContext {
	var p = new(StringInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringInContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringInContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *StringInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringIn(s)
	}
}

func (s *StringInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringIn(s)
	}
}

type StringGreaterEqualContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringGreaterEqualContext {
	var p = new(StringGreaterEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringGreaterEqualContext) GetLeft() IStringExprContext { return s.left }

func (s *StringGreaterEqualContext) GetRight() IStringExprContext { return s.right }

func (s *StringGreaterEqualContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringGreaterEqualContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringGreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringGreaterEqualContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringGreaterEqualContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringGreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringGreaterEqual(s)
	}
}

func (s *StringGreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringGreaterEqual(s)
	}
}

type DateGreaterThanContext struct {
	*BoolExprContext
	left  IDateExprContext
	right IDateExprContext
}

func NewDateGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateGreaterThanContext {
	var p = new(DateGreaterThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateGreaterThanContext) GetLeft() IDateExprContext { return s.left }

func (s *DateGreaterThanContext) GetRight() IDateExprContext { return s.right }

func (s *DateGreaterThanContext) SetLeft(v IDateExprContext) { s.left = v }

func (s *DateGreaterThanContext) SetRight(v IDateExprContext) { s.right = v }

func (s *DateGreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateGreaterThanContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateGreaterThanContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateGreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateGreaterThan(s)
	}
}

func (s *DateGreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateGreaterThan(s)
	}
}

type StringNotEqualContext struct {
	*BoolExprContext
	left  IStringExprContext
	right IStringExprContext
}

func NewStringNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringNotEqualContext {
	var p = new(StringNotEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *StringNotEqualContext) GetLeft() IStringExprContext { return s.left }

func (s *StringNotEqualContext) GetRight() IStringExprContext { return s.right }

func (s *StringNotEqualContext) SetLeft(v IStringExprContext) { s.left = v }

func (s *StringNotEqualContext) SetRight(v IStringExprContext) { s.right = v }

func (s *StringNotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringNotEqualContext) AllStringExpr() []IStringExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStringExprContext)(nil)).Elem())
	var tst = make([]IStringExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStringExprContext)
		}
	}

	return tst
}

func (s *StringNotEqualContext) StringExpr(i int) IStringExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStringExprContext)
}

func (s *StringNotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringNotEqual(s)
	}
}

func (s *StringNotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringNotEqual(s)
	}
}

type NumberGreaterEqualContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberGreaterEqualContext {
	var p = new(NumberGreaterEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberGreaterEqualContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberGreaterEqualContext) GetRight() INumberExprContext { return s.right }

func (s *NumberGreaterEqualContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberGreaterEqualContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberGreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberGreaterEqualContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberGreaterEqualContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberGreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberGreaterEqual(s)
	}
}

func (s *NumberGreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberGreaterEqual(s)
	}
}

type DateNotInContext struct {
	*BoolExprContext
}

func NewDateNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateNotInContext {
	var p = new(DateNotInContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *DateNotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateNotInContext) AllDateExpr() []IDateExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDateExprContext)(nil)).Elem())
	var tst = make([]IDateExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDateExprContext)
		}
	}

	return tst
}

func (s *DateNotInContext) DateExpr(i int) IDateExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDateExprContext)
}

func (s *DateNotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *DateNotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *DateNotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateNotIn(s)
	}
}

func (s *DateNotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateNotIn(s)
	}
}

type NumberLessThanContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLessThanContext {
	var p = new(NumberLessThanContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberLessThanContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberLessThanContext) GetRight() INumberExprContext { return s.right }

func (s *NumberLessThanContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberLessThanContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberLessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLessThanContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberLessThanContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberLessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberLessThan(s)
	}
}

func (s *NumberLessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberLessThan(s)
	}
}

type NumberNotEqualContext struct {
	*BoolExprContext
	left  INumberExprContext
	right INumberExprContext
}

func NewNumberNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberNotEqualContext {
	var p = new(NumberNotEqualContext)

	p.BoolExprContext = NewEmptyBoolExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolExprContext))

	return p
}

func (s *NumberNotEqualContext) GetLeft() INumberExprContext { return s.left }

func (s *NumberNotEqualContext) GetRight() INumberExprContext { return s.right }

func (s *NumberNotEqualContext) SetLeft(v INumberExprContext) { s.left = v }

func (s *NumberNotEqualContext) SetRight(v INumberExprContext) { s.right = v }

func (s *NumberNotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberNotEqualContext) AllNumberExpr() []INumberExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberExprContext)(nil)).Elem())
	var tst = make([]INumberExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberExprContext)
		}
	}

	return tst
}

func (s *NumberNotEqualContext) NumberExpr(i int) INumberExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberExprContext)
}

func (s *NumberNotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberNotEqual(s)
	}
}

func (s *NumberNotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberNotEqual(s)
	}
}

func (p *AlteryxFormulasParser) BoolExpr() (localctx IBoolExprContext) {
	return p.boolExpr(0)
}

func (p *AlteryxFormulasParser) boolExpr(_p int) (localctx IBoolExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, AlteryxFormulasParserRULE_boolExpr, _p)
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
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBoolParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(248)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(249)
			p.boolExpr(0)
		}
		{
			p.SetState(250)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewStringEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(252)

			var _x = p.stringExpr(0)

			localctx.(*StringEqualContext).left = _x
		}
		{
			p.SetState(253)
			p.Match(AlteryxFormulasParserT__7)
		}
		{
			p.SetState(254)

			var _x = p.stringExpr(0)

			localctx.(*StringEqualContext).right = _x
		}

	case 3:
		localctx = NewStringGreaterThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(256)

			var _x = p.stringExpr(0)

			localctx.(*StringGreaterThanContext).left = _x
		}
		{
			p.SetState(257)
			p.Match(AlteryxFormulasParserT__8)
		}
		{
			p.SetState(258)

			var _x = p.stringExpr(0)

			localctx.(*StringGreaterThanContext).right = _x
		}

	case 4:
		localctx = NewStringGreaterEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)

			var _x = p.stringExpr(0)

			localctx.(*StringGreaterEqualContext).left = _x
		}
		{
			p.SetState(261)
			p.Match(AlteryxFormulasParserT__9)
		}
		{
			p.SetState(262)

			var _x = p.stringExpr(0)

			localctx.(*StringGreaterEqualContext).right = _x
		}

	case 5:
		localctx = NewStringLessThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(264)

			var _x = p.stringExpr(0)

			localctx.(*StringLessThanContext).left = _x
		}
		{
			p.SetState(265)
			p.Match(AlteryxFormulasParserT__10)
		}
		{
			p.SetState(266)

			var _x = p.stringExpr(0)

			localctx.(*StringLessThanContext).right = _x
		}

	case 6:
		localctx = NewStringLessEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(268)

			var _x = p.stringExpr(0)

			localctx.(*StringLessEqualContext).left = _x
		}
		{
			p.SetState(269)
			p.Match(AlteryxFormulasParserT__11)
		}
		{
			p.SetState(270)

			var _x = p.stringExpr(0)

			localctx.(*StringLessEqualContext).right = _x
		}

	case 7:
		localctx = NewStringNotEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(272)

			var _x = p.stringExpr(0)

			localctx.(*StringNotEqualContext).left = _x
		}
		{
			p.SetState(273)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(274)

			var _x = p.stringExpr(0)

			localctx.(*StringNotEqualContext).right = _x
		}

	case 8:
		localctx = NewStringInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.stringExpr(0)
		}
		{
			p.SetState(277)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(278)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AlteryxFormulasParserField-33))|(1<<(AlteryxFormulasParserSingleQuoteString-33))|(1<<(AlteryxFormulasParserDoubleQuoteString-33)))) != 0) {
			{
				p.SetState(279)
				p.stringExpr(0)
			}
			p.SetState(284)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(280)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(281)
					p.stringExpr(0)
				}

				p.SetState(286)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(289)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 9:
		localctx = NewStringNotInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(291)
			p.stringExpr(0)
		}
		{
			p.SetState(292)
			p.Match(AlteryxFormulasParserNot)
		}
		{
			p.SetState(293)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(294)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(AlteryxFormulasParserField-33))|(1<<(AlteryxFormulasParserSingleQuoteString-33))|(1<<(AlteryxFormulasParserDoubleQuoteString-33)))) != 0) {
			{
				p.SetState(295)
				p.stringExpr(0)
			}
			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(296)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(297)
					p.stringExpr(0)
				}

				p.SetState(302)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(305)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 10:
		localctx = NewNumberEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(307)

			var _x = p.numberExpr(0)

			localctx.(*NumberEqualContext).left = _x
		}
		{
			p.SetState(308)
			p.Match(AlteryxFormulasParserT__7)
		}
		{
			p.SetState(309)

			var _x = p.numberExpr(0)

			localctx.(*NumberEqualContext).right = _x
		}

	case 11:
		localctx = NewNumberGreaterThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(311)

			var _x = p.numberExpr(0)

			localctx.(*NumberGreaterThanContext).left = _x
		}
		{
			p.SetState(312)
			p.Match(AlteryxFormulasParserT__8)
		}
		{
			p.SetState(313)

			var _x = p.numberExpr(0)

			localctx.(*NumberGreaterThanContext).right = _x
		}

	case 12:
		localctx = NewNumberGreaterEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)

			var _x = p.numberExpr(0)

			localctx.(*NumberGreaterEqualContext).left = _x
		}
		{
			p.SetState(316)
			p.Match(AlteryxFormulasParserT__9)
		}
		{
			p.SetState(317)

			var _x = p.numberExpr(0)

			localctx.(*NumberGreaterEqualContext).right = _x
		}

	case 13:
		localctx = NewNumberLessThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(319)

			var _x = p.numberExpr(0)

			localctx.(*NumberLessThanContext).left = _x
		}
		{
			p.SetState(320)
			p.Match(AlteryxFormulasParserT__10)
		}
		{
			p.SetState(321)

			var _x = p.numberExpr(0)

			localctx.(*NumberLessThanContext).right = _x
		}

	case 14:
		localctx = NewNumberLessEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(323)

			var _x = p.numberExpr(0)

			localctx.(*NumberLessEqualContext).left = _x
		}
		{
			p.SetState(324)
			p.Match(AlteryxFormulasParserT__11)
		}
		{
			p.SetState(325)

			var _x = p.numberExpr(0)

			localctx.(*NumberLessEqualContext).right = _x
		}

	case 15:
		localctx = NewNumberNotEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(327)

			var _x = p.numberExpr(0)

			localctx.(*NumberNotEqualContext).left = _x
		}
		{
			p.SetState(328)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(329)

			var _x = p.numberExpr(0)

			localctx.(*NumberNotEqualContext).right = _x
		}

	case 16:
		localctx = NewNumberInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(331)
			p.numberExpr(0)
		}
		{
			p.SetState(332)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(333)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__6)|(1<<AlteryxFormulasParserPow)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf)|(1<<AlteryxFormulasParserInteger)|(1<<AlteryxFormulasParserDecimal))) != 0) || _la == AlteryxFormulasParserField {
			{
				p.SetState(334)
				p.numberExpr(0)
			}
			p.SetState(339)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(335)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(336)
					p.numberExpr(0)
				}

				p.SetState(341)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(344)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 17:
		localctx = NewNumberNotInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(346)
			p.numberExpr(0)
		}
		{
			p.SetState(347)
			p.Match(AlteryxFormulasParserNot)
		}
		{
			p.SetState(348)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(349)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__6)|(1<<AlteryxFormulasParserPow)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf)|(1<<AlteryxFormulasParserInteger)|(1<<AlteryxFormulasParserDecimal))) != 0) || _la == AlteryxFormulasParserField {
			{
				p.SetState(350)
				p.numberExpr(0)
			}
			p.SetState(355)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(351)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(352)
					p.numberExpr(0)
				}

				p.SetState(357)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(360)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 18:
		localctx = NewDateEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(362)

			var _x = p.DateExpr()

			localctx.(*DateEqualContext).left = _x
		}
		{
			p.SetState(363)
			p.Match(AlteryxFormulasParserT__7)
		}
		{
			p.SetState(364)

			var _x = p.DateExpr()

			localctx.(*DateEqualContext).right = _x
		}

	case 19:
		localctx = NewDateGreaterThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(366)

			var _x = p.DateExpr()

			localctx.(*DateGreaterThanContext).left = _x
		}
		{
			p.SetState(367)
			p.Match(AlteryxFormulasParserT__8)
		}
		{
			p.SetState(368)

			var _x = p.DateExpr()

			localctx.(*DateGreaterThanContext).right = _x
		}

	case 20:
		localctx = NewDateGreaterEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(370)

			var _x = p.DateExpr()

			localctx.(*DateGreaterEqualContext).left = _x
		}
		{
			p.SetState(371)
			p.Match(AlteryxFormulasParserT__9)
		}
		{
			p.SetState(372)

			var _x = p.DateExpr()

			localctx.(*DateGreaterEqualContext).right = _x
		}

	case 21:
		localctx = NewDateLessThanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(374)

			var _x = p.DateExpr()

			localctx.(*DateLessThanContext).left = _x
		}
		{
			p.SetState(375)
			p.Match(AlteryxFormulasParserT__10)
		}
		{
			p.SetState(376)

			var _x = p.DateExpr()

			localctx.(*DateLessThanContext).right = _x
		}

	case 22:
		localctx = NewDateLessEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(378)

			var _x = p.DateExpr()

			localctx.(*DateLessEqualContext).left = _x
		}
		{
			p.SetState(379)
			p.Match(AlteryxFormulasParserT__11)
		}
		{
			p.SetState(380)

			var _x = p.DateExpr()

			localctx.(*DateLessEqualContext).right = _x
		}

	case 23:
		localctx = NewDateNotEqualContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(382)

			var _x = p.DateExpr()

			localctx.(*DateNotEqualContext).left = _x
		}
		{
			p.SetState(383)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(384)

			var _x = p.DateExpr()

			localctx.(*DateNotEqualContext).right = _x
		}

	case 24:
		localctx = NewDateInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(386)
			p.DateExpr()
		}
		{
			p.SetState(387)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(388)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf)|(1<<AlteryxFormulasParserDate))) != 0) || _la == AlteryxFormulasParserDatetime || _la == AlteryxFormulasParserField {
			{
				p.SetState(389)
				p.DateExpr()
			}
			p.SetState(394)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(390)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(391)
					p.DateExpr()
				}

				p.SetState(396)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(399)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 25:
		localctx = NewDateNotInContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(401)
			p.DateExpr()
		}
		{
			p.SetState(402)
			p.Match(AlteryxFormulasParserNot)
		}
		{
			p.SetState(403)
			p.Match(AlteryxFormulasParserIn)
		}
		{
			p.SetState(404)
			p.Match(AlteryxFormulasParserT__0)
		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserMin)|(1<<AlteryxFormulasParserMax)|(1<<AlteryxFormulasParserIf)|(1<<AlteryxFormulasParserDate))) != 0) || _la == AlteryxFormulasParserDatetime || _la == AlteryxFormulasParserField {
			{
				p.SetState(405)
				p.DateExpr()
			}
			p.SetState(410)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AlteryxFormulasParserT__3 {
				{
					p.SetState(406)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(407)
					p.DateExpr()
				}

				p.SetState(412)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(415)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 26:
		localctx = NewBoolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(417)
			p.Match(AlteryxFormulasParserBool)
		}

	case 27:
		localctx = NewBoolFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(418)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(429)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(427)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndContext(p, NewBoolExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_boolExpr)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(422)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__13 || _la == AlteryxFormulasParserAnd) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(423)

					var _x = p.boolExpr(5)

					localctx.(*AndContext).right = _x
				}

			case 2:
				localctx = NewOrContext(p, NewBoolExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_boolExpr)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(425)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__14 || _la == AlteryxFormulasParserOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)

					var _x = p.boolExpr(4)

					localctx.(*OrContext).right = _x
				}

			}

		}
		p.SetState(431)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IStrContext is an interface to support dynamic dispatch.
type IStrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrContext differentiates from other interfaces.
	IsStrContext()
}

type StrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrContext() *StrContext {
	var p = new(StrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_str
	return p
}

func (*StrContext) IsStrContext() {}

func NewStrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrContext {
	var p = new(StrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_str

	return p
}

func (s *StrContext) GetParser() antlr.Parser { return s.parser }

func (s *StrContext) SingleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSingleQuoteString, 0)
}

func (s *StrContext) DoubleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDoubleQuoteString, 0)
}

func (s *StrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStr(s)
	}
}

func (s *StrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStr(s)
	}
}

func (p *AlteryxFormulasParser) Str() (localctx IStrContext) {
	localctx = NewStrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AlteryxFormulasParserRULE_str)
	var _la int

	defer func() {
		p.ExitRule()
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		_la = p.GetTokenStream().LA(1)

		if !(_la == AlteryxFormulasParserSingleQuoteString || _la == AlteryxFormulasParserDoubleQuoteString) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *AlteryxFormulasParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *StringExprContext = nil
		if localctx != nil {
			t = localctx.(*StringExprContext)
		}
		return p.StringExpr_Sempred(t, predIndex)

	case 4:
		var t *NumberExprContext = nil
		if localctx != nil {
			t = localctx.(*NumberExprContext)
		}
		return p.NumberExpr_Sempred(t, predIndex)

	case 8:
		var t *BoolExprContext = nil
		if localctx != nil {
			t = localctx.(*BoolExprContext)
		}
		return p.BoolExpr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlteryxFormulasParser) StringExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AlteryxFormulasParser) NumberExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *AlteryxFormulasParser) BoolExpr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
