// Code generated by "i18n-stringer -type Code,Test,Single"; DO NOT EDIT.

package test

import (
	"context"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[CodeOK-1]
	_ = x[CodeErr-2]
	_ = x[CodeFail-3]
	_ = x[CodeRange1-20001]
	_ = x[CodeRange2-20002]
	_ = x[CodeRange3-20003]
	_ = x[CodeRange4-20004]
	_ = x[CodeRange5-20204]
	_ = x[CodeRange6-20205]
	_ = x[CodeRange7-20206]
	_ = x[CodeRange8-20206]
	_ = x[CodeRange9-20301]
	_ = x[CodeRange10-20302]
	_ = x[CodeTe1-20310]
	_ = x[CodeTe2-20311]
	_ = x[CodeSe1-20400]
	_ = x[CodeSe2-20401]
	_ = x[CodeSe3-20410]
	_ = x[CodeSe4-20411]
	_ = x[CodeAe1-20420]
	_ = x[CodeAe2-20421]
	_ = x[CodeBe1-20430]
	_ = x[CodeBe2-20431]
	_ = x[CodeCe1-20440]
	_ = x[CodeCe2-20441]
	_ = x[CodeDe1-20450]
	_ = x[CodeDe2-20451]
	_ = x[CodeEe1-20460]
	_ = x[CodeEe2-20461]
	_ = x[CodeFe1-20470]
	_ = x[CodeFe2-20471]
	_ = x[CodeFe3-20472]
	_ = x[CodeGe1-20480]
	_ = x[CodeGe2-20481]
	_ = x[CodeXe1-20490]
	_ = x[CodeXe2-20491]
	_ = x[CodeXe3-20491]
}

const _Code_name = "CodeOKCodeErrCodeFailCodeRange1CodeRange2CodeRange3CodeRange4CodeRange5CodeRange6CodeRange7CodeRange9CodeRange10CodeTe1CodeTe2CodeSe1CodeSe2CodeSe3CodeSe4CodeAe1CodeAe2CodeBe1CodeBe2CodeCe1CodeCe2CodeDe1CodeDe2CodeEe1CodeEe2CodeFe1CodeFe2CodeFe3CodeGe1CodeGe2CodeXe1CodeXe2"

var _Code_map = map[Code]string{
	1:     _Code_name[0:6],
	2:     _Code_name[6:13],
	3:     _Code_name[13:21],
	20001: _Code_name[21:31],
	20002: _Code_name[31:41],
	20003: _Code_name[41:51],
	20004: _Code_name[51:61],
	20204: _Code_name[61:71],
	20205: _Code_name[71:81],
	20206: _Code_name[81:91],
	20301: _Code_name[91:101],
	20302: _Code_name[101:112],
	20310: _Code_name[112:119],
	20311: _Code_name[119:126],
	20400: _Code_name[126:133],
	20401: _Code_name[133:140],
	20410: _Code_name[140:147],
	20411: _Code_name[147:154],
	20420: _Code_name[154:161],
	20421: _Code_name[161:168],
	20430: _Code_name[168:175],
	20431: _Code_name[175:182],
	20440: _Code_name[182:189],
	20441: _Code_name[189:196],
	20450: _Code_name[196:203],
	20451: _Code_name[203:210],
	20460: _Code_name[210:217],
	20461: _Code_name[217:224],
	20470: _Code_name[224:231],
	20471: _Code_name[231:238],
	20472: _Code_name[238:245],
	20480: _Code_name[245:252],
	20481: _Code_name[252:259],
	20490: _Code_name[259:266],
	20491: _Code_name[266:273],
}

func (i Code) String() string {
	if str, ok := _Code_map[i]; ok {
		return str
	}
	return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
}

// _Code_defaultLocale default locale generated pass by flag -defaultlocale, Don't assign directly
var _Code_defaultLocale = "en"

// _Code_supported All supported locales and text offset information
var _Code_supported = map[string]int{"zh-hk": 0}

// _Code_ctxKey Key of set used locale from context.Context Value
var _Code_ctxKey = "i18nLocale"

func _Code_isLocaleSupport(locale string) bool {
	_, ok := _Code_supported[locale]
	return ok
}

// _Code_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Code_isLocaleSupport is false
func _Code_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Code_defaultLocale
	}
	v := ctx.Value(_Code_ctxKey)
	if v != nil && _Code_isLocaleSupport(v.(string)) {
		return v.(string)
	}
	return _Code_defaultLocale
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[TestCase01-10]
	_ = x[TestCase02-11]
	_ = x[TestCase03-12]
	_ = x[TestCase04-1001]
	_ = x[TestCase05-1002]
	_ = x[TestCase06-1003]
}

const (
	_Test_name_0 = "TestCase01TestCase02TestCase03"
	_Test_name_1 = "TestCase04TestCase05TestCase06"
)

var (
	_Test_index_0 = [...]uint8{0, 10, 20, 30}
	_Test_index_1 = [...]uint8{0, 10, 20, 30}
)

// String type to fmt.Stringer interface use default locale
func (i Test) String() string {
	switch {
	case 10 <= i && i <= 12:
		i -= 10
		return _Test_name_0[_Test_index_0[i]:_Test_index_0[i+1]]
	case 1001 <= i && i <= 1003:
		i -= 1001
		return _Test_name_1[_Test_index_1[i]:_Test_index_1[i+1]]
	default:
		return "Test(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}

// _Test_defaultLocale default locale generated pass by flag -defaultlocale, Don't assign directly
var _Test_defaultLocale = "en"

// _Test_supported All supported locales and text offset information
var _Test_supported = map[string]int{"zh-hk": 0}

// _Test_ctxKey Key of set used locale from context.Context Value
var _Test_ctxKey = "i18nLocale"

func _Test_isLocaleSupport(locale string) bool {
	_, ok := _Test_supported[locale]
	return ok
}

// _Test_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Test_isLocaleSupport is false
func _Test_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Test_defaultLocale
	}
	v := ctx.Value(_Test_ctxKey)
	if v != nil && _Test_isLocaleSupport(v.(string)) {
		return v.(string)
	}
	return _Test_defaultLocale
}

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the i18n-stringer command to generate them again.
	var x [1]struct{}
	_ = x[Sig01-300]
	_ = x[Sig02-301]
	_ = x[Sig03-302]
	_ = x[Sig04-303]
	_ = x[Sig05-304]
	_ = x[Sig06-305]
	_ = x[Sig07-306]
	_ = x[Sig08-307]
	_ = x[Sig09-308]
	_ = x[Sig10-309]
	_ = x[Sig11-310]
	_ = x[Sig12-311]
	_ = x[Sig13-312]
	_ = x[Sig14-313]
	_ = x[Sig15-314]
	_ = x[Sig16-315]
	_ = x[Sig17-316]
	_ = x[Sig18-317]
	_ = x[Sig19-318]
	_ = x[Sig20-319]
	_ = x[Sig21-320]
	_ = x[Sig22-321]
	_ = x[Sig23-322]
	_ = x[Sig24-323]
}

const _Single_name = "Sig01Sig02Sig03Sig04Sig05Sig06Sig07Sig08Sig09Sig10Sig11Sig12Sig13Sig14Sig15Sig16Sig17Sig18Sig19Sig20Sig21Sig22Sig23Sig24"

var _Single_index = [...]uint8{0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100, 105, 110, 115, 120}

// String type to fmt.Stringer interface use default locale
func (i Single) String() string {
	i -= 300
	if i < 0 || i >= Single(len(_Single_index)-1) {
		return "Single(" + strconv.FormatInt(int64(i+300), 10) + ")"
	}
	return _Single_name[_Single_index[i]:_Single_index[i+1]]
}

// _Single_defaultLocale default locale generated pass by flag -defaultlocale, Don't assign directly
var _Single_defaultLocale = "en"

// _Single_supported All supported locales and text offset information
var _Single_supported = map[string]int{"zh-hk": 0}

// _Single_ctxKey Key of set used locale from context.Context Value
var _Single_ctxKey = "i18nLocale"

func _Single_isLocaleSupport(locale string) bool {
	_, ok := _Single_supported[locale]
	return ok
}

// _Single_localeFromCtxWithFallback retrieves and returns language locale name from context.
// It returns default locale when _Single_isLocaleSupport is false
func _Single_localeFromCtxWithFallback(ctx context.Context) string {
	if ctx == nil {
		return _Single_defaultLocale
	}
	v := ctx.Value(_Single_ctxKey)
	if v != nil && _Single_isLocaleSupport(v.(string)) {
		return v.(string)
	}
	return _Single_defaultLocale
}