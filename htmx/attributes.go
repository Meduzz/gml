package htmx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	RequestSetting func(map[string]any)

	HeaderSetting func(map[string]string)

	FormEncoding string
)

var (
	UrlEncoded = FormEncoding("application/x-www-form-urlencoded")
	Multipart  = FormEncoding("multipart/form-data")
)

func Boost() []string {
	return []string{
		"hx-boost",
		"true",
	}
}

func Confirm(message string) []string {
	return []string{
		"hx-confirm",
		message,
	}
}

func Delete(url string) []string {
	return []string{
		"hx-delete",
		url,
	}
}

func Disable() []string {
	return []string{
		"hx-disable",
	}
}

func DisabledELT(selectors ...string) []string {
	return []string{
		"hx-disabled-lt",
		strings.Join(selectors, " "),
	}
}

func Disinherit(tags string) []string {
	// *
	// hx-target
	// hx-select
	// osv
	return []string{
		"hx-disinherit",
		tags,
	}
}

func Encoding(encoding FormEncoding) []string {
	return []string{
		"hx-encoding",
		string(encoding),
	}
}

func Extension(name string) []string {
	return []string{
		"hx-ext",
		name,
	}
}

func Get(url string) []string {
	return []string{
		"hx-get",
		url,
	}
}

func Headers(settings ...HeaderSetting) []string {
	data := make(map[string]string)

	slice.ForEach(settings, func(it HeaderSetting) {
		it(data)
	})

	return []string{
		"hx-headers",
		toJsonS(data),
	}
}

func History() []string {
	return []string{
		"hx-history",
		"false",
	}
}

func HistoryELT(selector string) []string {
	return []string{
		"hx-history-elt",
		selector,
	}
}

func Include(selector string) []string {
	return []string{
		"hx-include",
		selector,
	}
}

func Indicator(selector string) []string {
	return []string{
		"hx-indicator",
		selector,
	}
}

func Inherit(selector string) []string {
	return []string{
		"hx-inherit",
		selector,
	}
}

func Params(args string) []string {
	// *
	// none
	// not <param list>
	// <param list>
	return []string{
		"hx-params",
		args,
	}
}

func Patch(url string) []string {
	return []string{
		"hx-patch",
		url,
	}
}

func Post(url string) []string {
	return []string{
		"hx-post",
		url,
	}
}

func Preserve() []string {
	return []string{
		"hx-preserve",
		"true",
	}
}

func Prompt(message string) []string {
	return []string{
		"hx-prompts",
		message,
	}
}

func PushURL(arg string) []string {
	// true
	// false
	// url
	return []string{
		"hx-push-url",
		arg,
	}
}

func Put(url string) []string {
	return []string{
		"hx-put",
		url,
	}
}

func ReplaceURL(arg string) []string {
	// true
	// false
	// url
	return []string{
		"hx-replace-url",
		arg,
	}
}

func Request(visitors ...RequestSetting) []string {
	settings := make(map[string]any)

	slice.ForEach(visitors, func(visitor RequestSetting) {
		visitor(settings)
	})

	return []string{
		"hx-request",
		toJson(settings),
	}
}

func Sync(settings string) []string {
	return []string{
		"hx-sync",
		settings,
	}
}

func Select(selector string) []string {
	return []string{
		"hx-select",
		selector,
	}
}

func SelectOOB(selector string) []string {
	return []string{
		"hx-select-oob",
		selector,
	}
}

func Swap(swap string) []string {
	return []string{
		"hx-swap",
		swap,
	}
}

func SwapOOB(swap string) []string {
	return []string{
		"hx-swap-oob",
		swap,
	}
}

func Target(selector string) []string {
	return []string{
		"hx-target",
		selector,
	}
}

func Trigger(event string) []string {
	return []string{
		"hx-trigger",
		event,
	}
}

func Validate() []string {
	return []string{
		"hx-validate",
		"true",
	}
}

func Vals(data map[string]any) []string {
	return []string{
		"hx-vals",
		toJson(data),
	}
}

func RequestTimeout(timeout int) RequestSetting {
	return func(m map[string]any) {
		m["timeout"] = timeout
	}
}

func RequestCredentials(torf bool) RequestSetting {
	return func(m map[string]any) {
		m["credentials"] = torf
	}
}

func RequestNoHeaders(torf bool) RequestSetting {
	return func(m map[string]any) {
		m["noheaders"] = torf
	}
}

func PlainHeader(header, value string) HeaderSetting {
	return func(m map[string]string) {
		m[header] = value
	}
}

func AuthorizationHerder(auth string, bearer bool) HeaderSetting {
	return func(m map[string]string) {
		if bearer {
			m["Authorization"] = fmt.Sprintf("Bearer %s", auth)
		} else {
			m["Authorization"] = auth
		}
	}
}

func toJson(data map[string]any) string {
	props := make([]string, 0)

	for k, v := range data {
		fmt.Printf("%s=%s\n", k, reflect.TypeOf(v).String())
		switch v.(type) {
		case string:
			props = append(props, fmt.Sprintf(`"%s":"%v"`, k, v))
			break
		case bool, int, int32, int64, float32, float64:
			props = append(props, fmt.Sprintf(`"%s":%v`, k, v))
			break
		default:
			continue
		}

	}

	return fmt.Sprintf("{%s}", strings.Join(props, ", "))
}

func toJsonS(data map[string]string) string {
	props := make([]string, 0)

	for k, v := range data {
		props = append(props, fmt.Sprintf(`"%s":"%s"`, k, v))
	}

	return fmt.Sprintf("{%s}", strings.Join(props, ""))
}
