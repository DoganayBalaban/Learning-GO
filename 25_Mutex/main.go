package main

import (
	"fmt"
	"sync"
)

//Bu nesne aynı anda sadece bir işlemi gerçekleştiriyor. Bu yüzden sıra işlemi sağlıyor.
//Önce başlayan asenkron işlem  ilk sırada oluyor. Tamamlanınca diğerine sıra geçiyor.

var mt sync.Mutex //global mutex nesnesi

func paraÇek(bakiye *float64, çekilecekMiktar float64, wg *sync.WaitGroup) { /*
	* mt isimli mutex'i bu işlem yapılırken kilitliyoruz.
	* bu sayede mt mutex'ini başka işlemler kullanamıyor.
	 */
	mt.Lock()

	/*
		bu kısımda asenkron olmasını istemediğimiz işlemi yapalım.
	*/
	*bakiye -= 25
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)

	/*
	* diğer işlemlerinde kullanabilmesi için mutex'i tekrardan açalım.
	* mt mutex açılınca diğer asenkron işlemdeki mt mutex'i çalışmaya başlar.
	 */
	mt.Unlock()
	fmt.Println("Çekme işlemi tamamlandı.")

	/*
	* waitgroup ile işlemin tamamlandığını belirttik.
	* böylece wg havuzu 2'den 1'e düştü
	 */
	wg.Done()
}
func paraYatır(bakiye *float64, yatırılacakMiktar float64, wg *sync.WaitGroup) {
	mt.Lock()
	*bakiye += 65
	fmt.Printf("Yeni Bakiye: %.2f\n", *bakiye)
	mt.Unlock()
	fmt.Println("Yatırma işlemi tamamlandı.")
	wg.Done()
}
func main() {
	/*
	* asenkron işlemlerimizin, ana iş parçacığında tamamlanmasını
	* beklemek için waitgroup nesnesi oluşturalım
	 */
	var wg sync.WaitGroup

	//2 fonksiyonu da bekleyeceğimiz için Add'e 2 yazalım
	wg.Add(2)

	//fonksiyonlarımızın kullancağı bakiye değişkenimiz
	var bakiye float64 = 100
	fmt.Printf("İlk Bakiye: %.2f\n", bakiye)

	/*
	* paraÇek ve paraYatır fonksiyonlarımızı aynı anda başlatıyoruz.
	* hangisi daha önce başlarsa mutex sırasına ilk o girer. bu esnada diğer
	* fonksiyon mutex'in açılmasını bekler.
	 */
	go paraÇek(&bakiye, 25, &wg)
	go paraYatır(&bakiye, 65, &wg)

	/*
	* ana iş parçacığı tamamlandığında asenkron çalışan fonksiyonları beklemez.
	* beklemediğinde de asenkron fonksiyonlar çalışmadan program sonlanır.
	* ana iş parçacığının asenkron işlemleri beklemesi için waitgroup sonucunun 0 olmasını bekleriz.
	* wg.Add(2) yazarak 2 adet wg.Done() fonksiyonu çalıştığında wg.Add(0) olur ve
	* wg.Wait() tamamlanır ve program başka işlemler yapılmıyor ise sonlanır.
	 */
	wg.Wait()
}
