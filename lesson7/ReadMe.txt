
- fmt.Printf, которая записывает результаты в стандартный поток вывода(файл)
- fmt.Sprintf, которая возвращает результат в виде строки string. 
Благодаря интерфейсам - обе функции являются, по сути, "обертками" для третьей функции, fmt.Fprintf, 
которая ничего не знает о том, что будет с результатом, который она вычисляет:

package fmt

func Fprintf (w io.Writer, format string, a rgs ... interfa ce{ } ) (int, error)

func P rintf(format string, a rgs ... interface{}) (int, error ) {
return Fprintf (os.Stdout , format , a rgs ... )
}

func Sprintf (format string, args ... interface{}) string {
var buf bytes.Buffer
Fprintf (&buf, format, args ... )
return buf . String()
} 