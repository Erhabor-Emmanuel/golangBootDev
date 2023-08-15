package interfaces

//Type Assertions in GO
//When working with interfaces in Go, every once in a while you will need access to an underlying
//type of an interface value. You can cast an interface to its underlying type using a type insertion

type food interface {
	getFoods()
}

func getExpenseReport(e expense) (string, float64) {

	//which is an instance of a shape.
	//"ok" is bool that is true if e was of type email
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()
	}
	s, ok := e.(sms)
	if ok {
		return s.toPhoneNumber, s.cost()
	}
	return "", 0
}

func (e email) cost() float64 {
	if !e.subscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.subscribed {
		return float64(len(s.body)) * .05
	}
	return float64(len(s.body)) * .01
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	subscribed bool
	body       string
	toAddress  string
}

type sms struct {
	subscribed    bool
	body          string
	toPhoneNumber string
}

type invalid struct{}
