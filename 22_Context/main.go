/*Context paketi; istek tabanlı verileri,
iptal sinyallerini, deadline bilgilerini bir API kapsamı içindeki tüm fonksiyonlara, goroutine’lere geçirebilmek için tasarlandı.*/

/*
	type Context interface {
		// Done metodu, context tamamlandığında veya zaman aşımına uğradığında
	  // kapanan bir channel döndürür.

Done() <-chan struct{}

	// Err metodu, context'in neden kapandığını anlatan bir error döndürür.

// Genellikle Done kanalı kapandıktan sonra bunu kullanırız.
Err() error

// Deadline metodu, eğer context herhangi bir deadline'a sahipse deadline zamanını döndürür.
// `ok` değeri context bir deadline'a sahipse true, değilse false olur.
Deadline() (deadline time.Time, ok bool)

// Value metodu, context'in içerisinde saklanan bilgilere ulaşmak için kullanılır.
Value(key interface{}) interface{}
}
*/
/* package main

import (
	"context"
	"fmt"
	"time"
)

func a(ctx context.Context) {
	fmt.Println("a çalışmaya başladı")

	select {
	case <-time.After(1 * time.Second): // işlemler 1 saniye sürüyor
		fmt.Println("a çalışmasını bitirdi")
	case <-ctx.Done():
		fmt.Println("context iptal edildi")
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go a(ctx) // a()'yı ayrı bi goroutine içinde başlattık.

	time.Sleep(500 * time.Millisecond) // context'i iptal etmeden önce 500ms uyuyoruz.
	cancel()                           // context'i iptal ediyoruz.

	time.Sleep(1 * time.Second) // a()'nın çıktısını görebilmek için biraz bekliyoruz.
	fmt.Println("program sonlandı")
}
*/
/* package main

import (
	"context"
	"fmt"
	"time"
)

func a(ctx context.Context) {
	fmt.Println("a çalışmaya başladı")

	select {
	case <-time.After(2 * time.Second): // işlemler 2 saniye sürüyor.
		fmt.Println("a çalışmasını bitirdi")
	case <-ctx.Done():
		fmt.Println("context iptal edildi")
	}
}

func main() {
	// 1 saniye sonra otomatik olarak iptal olacak bir context oluşturduk.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go a(ctx) // a()'yı ayrı bi goroutine içinde başlattık.

	time.Sleep(5 * time.Second) // Sonuçları görebilmek için biraz uyuyoruz.
	fmt.Println("program sonlandı")
} */

/* package main

import (
	"context"
	"fmt"
	"time"
)

func a(ctx context.Context) {
	fmt.Println("a çalışmaya başladı")

	select {
	case <-ctx.Done():
		fmt.Println("context iptal edildi")
	default:
		if yeniDeger := ctx.Value("YeniDeger"); yeniDeger != nil {
			fmt.Println("context içinden okunan değer:", yeniDeger)
		}
	}
}

func main() {
	// Context'in içine bilgi gömdük, yeni bir context oluşturduk.
	ctx := context.WithValue(context.Background(), "YeniDeger", "Merhaba Dünya")

	go a(ctx) // a()'yı ayrı bi goroutine içinde başlattık.

	time.Sleep(1 * time.Second) // Sonuçları görebilmek için biraz uyuyoruz.
	fmt.Println("program sonlandı")
} */

/*
	package main

import (

	"context"
	"fmt"
	"time"

)

	func a(ctx context.Context, idx int) {
		fmt.Println(idx, "çalışmaya başladı")

		select {
		case <-time.After(2 * time.Second):
			fmt.Println(idx, "çalışma sonlandı")
		case <-ctx.Done():
			fmt.Println(idx, "context iptal edildi")

		}
	}

	func main() {
		// Context'in içine bilgi gömdük, yeni bir context oluşturduk.
		baseCtx, cancel := context.WithCancel(context.Background())

		for i := 1; i <= 5; i++ {
			newCtx := context.WithValue(baseCtx, "YeniDeger", i) // baseCtx'den yeni bir context üretilir.
			go a(newCtx, i)                                      // a() yeni context ile çağırılır.
		}

		cancel() // baseCtx iptal edilir. Böylece baseCtx'den türetilen bütün context'ler iptal edilmiş olur.

		time.Sleep(5 * time.Second) // Sonuçları görebilmek için biraz uyuyoruz.
		fmt.Println("program sonlandı")
	}
*/
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func sendReq(ctx context.Context) error {
	fmt.Println("istek atılıyor")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://context.free.beeceptor.com/", nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("cevap alındı")
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := sendReq(ctx); err != nil {
		fmt.Println("sendReq err:", err)
	}

	fmt.Println("program sonlandı")
}
