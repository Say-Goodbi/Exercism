package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen" : 6,
        "dozen" : 12,
        "small_gross": 120,
        "gross": 144,
        "great_gross": 1728,
    }
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	score, exists := units[unit]
    if !exists {
    	return false 
    }
    _, exists = bill[item]
    if exists {
        bill[item] += score
    } else {
    	bill[item] = score
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	bill_score, exists := bill[item]
    score, exists2 := units[unit]
    if !exists || !exists2 || bill_score - score < 0 {
        return false
    }
    if (bill_score - score) == 0 {
        delete(bill, item)
    } else { 
    	bill[item] -= score
    }
    return true
    
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	score, exists := bill[item]
    return score, exists
    
}
