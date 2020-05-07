package prime

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"sync"
)

var primeList = []uint64{}

// ReadPrimes read the prime numbers from the file ./primes
func ReadPrimes() {
	f, err := os.Open("/Users/xianlinfeng/Documents/Go/src/Mathematics/prime/primes")
	Handle(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		k, _ := strconv.ParseUint(scanner.Text(), 10, 64)
		primeList = append(primeList, k)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// WritePrimes write the content of primeList into file ./primes
func WritePrimes() {
	f, err := os.Create("/Users/xianlinfeng/Documents/Go/src/Mathematics/prime/primes")
	Handle(err)
	defer f.Close()

	for _, k := range primeList {
		_, err := f.WriteString(strconv.FormatUint(k, 10) + "\n")
		Handle(err)
	}

	f.Sync()
}

// IsPrime check if the number is a prime number
func IsPrime(n uint64) bool {
	if len(primeList) == 0 {
		ReadPrimes()
	}
	factors := []uint64{}

	for _, p := range primeList {
		if p*p <= n && n%p == 0 {
			factors = append(factors, p)
			break
		}
	}

	if len(factors) == 0 {
		return true
	}
	return false
}

// GetNextPrimes find the next n prime numbers
func GetNextPrimes(n int) {
	if len(primeList) == 0 {
		ReadPrimes()
	}
	l := len(primeList)
	numbers := make(chan uint64, 1000)
	cancelChan := make(chan bool, 0)
	var wg sync.WaitGroup

	go func() {
		i := primeList[l-1]
		for {
			i++
			numbers <- i
		}
	}()

	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func() {
			for {
				if isCancelled(cancelChan) {
					wg.Done()
					return
				} else if len(primeList) == l+n {
					cancel(cancelChan)
					wg.Done()
					return
				}
				num := <-numbers
				if IsPrime(num) {
					primeList = append(primeList, num)
				}
			}
		}()
	}

	wg.Wait()

	// fmt.Println(primeList)

	WritePrimes()
}

// GetFactors get the factors of a integer number n
func GetFactors(n uint64) []uint64 {
	if len(primeList) == 0 {
		ReadPrimes()
	}
	factors := []uint64{}

	for _, p := range primeList {
	Again:
		if n%p == 0 {
			factors = append(factors, p)
			n /= p
			goto Again
		}
	}

	return factors
}

// Handle Handle an error
func Handle(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func cancel(cancelChan chan bool) {
	close(cancelChan)
}

func isCancelled(cancelChan chan bool) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

// IsPrime check is the number is a prime number or not
// func IsPrime(num uint64) {
// 	var i uint64
// 	var wg sync.WaitGroup
// 	pool := make(chan int, 200)

// 	for i = 3; i < num; i++ {
// 		factors := make(chan uint64, 100)

// 		for _, p := range primeList {
// 			if float64(p) <= math.Sqrt(float64(i)) {
// 				wg.Add(1)
// 				go func(i, p uint64, pool chan int) {
// 					pool <- 1
// 					if i%p == 0 {
// 						factors <- p
// 					}
// 					wg.Done()
// 					<-pool
// 				}(i, p, pool)
// 			}
// 		}

// 		wg.Wait()
// 		if len(factors) == 0 {
// 			primeList = append(primeList, i)

// 		}
// 	}

// 	fmt.Println(primeList)
// }
