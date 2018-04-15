package main

import "fmt"

// func main() {
// 	html, err := template.ParseFiles("index.gohtml")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	err = html.Execute(os.Stdout, nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func main() {
// 	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
//
// 	scanner := bufio.NewScanner(res.Body)
// 	defer res.Body.Close()
// 	scanner.Split(bufio.ScanWords)
//
// 	buckets := make([][]string, 12)
//
// 	for i := 0; i < 12; i++ {
// 		buckets = append(buckets, []string{})
// 	}
//
// 	for scanner.Scan() {
// 		word := scanner.Text()
// 		n := HashBucket(word, 12)
// 		buckets[n] = append(buckets[n], word)
// 	}
//
// 	for i := 0; i < 12; i++ {
// 		fmt.Println(i, " - ", len(buckets[i]))
// 	}
//
// 	fmt.Println(buckets[0])
// }
//
// func HashBucket(word string, buckets int) int {
// 	var sum int
// 	for _, v := range word {
// 		sum += int(v)
// 	}
// 	return sum % buckets
// }

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }
//
// func main() {
// 	bond := Person{"James", "Bond", 20}
// 	json.NewDecoder(os.Stdout).Decode(bond)
// }

// func main() {
// 	n := 1
// 	channel := make(chan int)
// 	done := make(chan bool)
//
// 	go func() {
// 		for i := 0; i < 100000; i++ {
// 			channel <- i
// 		}
// 		close(channel)
// 	}()
//
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			for n := range channel {
// 				fmt.Println(n)
// 			}
// 			done <- true
// 		}()
// 	}
//
// 	for i := 0; i < n; i++ {
// 		<-done
// 	}
// }

func main() {
	c := incrementor()
	fmt.Println(<-puller(c))
}

func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
