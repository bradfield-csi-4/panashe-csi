package metrics

import (
	"encoding/csv"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type UserId int
type UserMap map[UserId]*User

type Address struct {
	fullAddress string
	zip         int
}

type DollarAmount struct {
	cents uint64
}

type Payment struct {
	amount DollarAmount
	time   time.Time
}

type User struct {
	id       UserId
	name     string
	age      int
	address  Address
	payments []Payment
}

func AverageAge(ages []int) float64 {
	total1, total2, total3, total4, total5 := 0, 0, 0, 0, 0
	agesLen := len(ages)
	for i := 0; i < agesLen; i+=5 {
		total1 += ages[i]
		total2 += ages[i+1]
		total3 += ages[i+2]
		total4 += ages[i+3]
		total5 += ages[i+4]
	}
	if agesLen % 5 != 0 {
		for i := agesLen-(agesLen % 5); i < agesLen; i++ {
			total1 += ages[i]
		}
	}
	return float64(total1+total2+total3+total4+total5) / float64(len(ages))
}

func AveragePaymentAmount(payments []int) float64 {
	total1, total2, total3, total4, total5 := 0, 0, 0, 0, 0
	paymentsLen := len(payments)
	for i := 0; i < paymentsLen; i+=5 {
		total1 += payments[i]
		total2 += payments[i+1]
		total3 += payments[i+2]
		total4 += payments[i+3]
		total5 += payments[i+4]
	}
	if paymentsLen % 5 != 0 {
		for i := paymentsLen-(paymentsLen % 5); i < paymentsLen; i++ {
			total1 += payments[i]
		}
	}
	return float64(total1+total2+total3+total4+total5) / float64(len(payments)) / 100
}


func StdDevPaymentAmount(payments []int) float64 {
	count, mean, M2 := 0.0, 0.0, 0.0
	for _, p := range payments {
		count += 1
		amount := float64(p)
		delta := amount - mean
		mean += delta / count
		delta2 := amount - mean
		M2 += delta * delta2
	}

    return math.Sqrt(M2 / count) / 100
}

func LoadData() ([]int, []int, UserMap) {
	f, err := os.Open("users.csv")
	if err != nil {
		log.Fatalln("Unable to read users.csv", err)
	}
	reader := csv.NewReader(f)
	userLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse users.csv as csv", err)
	}

	users := make(UserMap, len(userLines))
	ages := make([]int, len(userLines))
	loopCount := 0
	for i, line := range userLines {
		id, _ := strconv.Atoi(line[0])
		name := line[1]
		age, _ := strconv.Atoi(line[2])
		ages[i] = age
		address := line[3]
		zip, _ := strconv.Atoi(line[3])
		users[UserId(id)] = &User{UserId(id), name, age, Address{address, zip}, []Payment{}}
		loopCount += 1
	}

	f, err = os.Open("payments.csv")
	if err != nil {
		log.Fatalln("Unable to read payments.csv", err)
	}
	reader = csv.NewReader(f)
	paymentLines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Unable to parse payments.csv as csv", err)
	}

	payments := make([]int, len(paymentLines))
	for i, line := range paymentLines {
		userId, _ := strconv.Atoi(line[2])
		paymentCents, _ := strconv.Atoi(line[0])
		datetime, _ := time.Parse(time.RFC3339, line[1])
		payments[i] = paymentCents
		payment := Payment{
			DollarAmount{uint64(paymentCents)},
			datetime,
		}
		users[UserId(userId)].payments = append(users[UserId(userId)].payments, payment)
	}

	return ages, payments, users
}
