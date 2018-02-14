package zendeskModel

import (
	"time"
)

//APIResponse struct for the API Response of the GetTicketList API
type APIResponse struct {
	Tickets []TicketList `json:"tickets"`
}

//TicketList struct for the object of API Response List
type TicketList struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	AssigneeID  int       `json:"assignee_id"`
	BrandID     int       `json:"brand_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Tags        []string  `json:"tags"`
	SubmitterID int       `json:"submitter_id"`
	Status      string    `json:"status"`
	URL         string    `json:"url"`
	RequesterID int       `json:"requester_id"`
}

//FindCurrencyMapping is a function to find all currency mappings and if revrese mapping exist for the query pair then this function
//will set the value of flag to false to notify the reverse mapping case in EvaluateQuotes Function
// func FindCurrencyMapping(baseCurrency string, quoteCurrency string, amount string) (string,bool,ProductMap,error) {

// 	//Denotes currency pair provided in the request
// 	key := baseCurrency+"-"+quoteCurrency

// 	//Reverse of the above currency pair to find out if given pair doesn't exist does its reverse exists
// 	reverseKey := quoteCurrency+"-"+baseCurrency

// 	URL := zendeskUtil.BaseURL

// 	//Http Client object for the request
// 	client := &http.Client{}
// 	req, _ := http.NewRequest("GET", URL,nil)

// 	//Make request
// 	resp, error := client.Do(req)
// 	//If API returns success code then return the currency mapping
// 	if( resp.StatusCode >= 200 && resp.StatusCode < 300 ) {
// 		var data []ProductMap

// 		//Store the currency pair as key and ProductMap struct as value
// 		dataMap := make(map[string]ProductMap)
// 		bodyBytes, _ := ioutil.ReadAll(resp.Body)
// 		error = json.Unmarshal(bodyBytes, &data)
// 		for _,val := range data {
// 			dataMap[val.ID] = val
// 		}
// 		//If Currency Pair Found
// 		if v, ok := dataMap[key]; ok {
// 			maxSize,_ := strconv.ParseFloat(v.BaseMaxSize, 64)
// 			minSize,_ := strconv.ParseFloat(v.BaseMinSize, 64)
// 			newAmount,_ := strconv.ParseFloat(amount,64)

// 			//
// 			if (maxSize >= newAmount && minSize <= newAmount) {
// 				return key,true,v,nil
// 			}else {
// 				return zendeskUtil.CurrencyPairErrorMessage,false,ProductMap{},fmt.Errorf("Error")
// 			}
// 		}else if v, ok := dataMap[reverseKey]; ok{			//If Reverse Currency Pair found
// 			return reverseKey,false,v,nil
// 		}else {
// 			return zendeskUtil.CurrencyPairNotFound,false,ProductMap{},fmt.Errorf("Not Found")
// 		}
// 	}

// 	//Return Error if both currency pair and its reverse not found
// 	return "",false,ProductMap{},error
// }

// //EvaluateQuotes is a function to evaluate the most appropriate quote for the given Request
// func EvaluateQuotes(quoteResponse QuoteResponse, params SampleGdaxQuoteModel,flag bool,pMap ProductMap) (Result,error) {
// 	result := Result{}
// 	AskQuotes := ParseAsks(quoteResponse.Asks)				//Store the list of Asks
// 	BidsQuotes := ParseBids(quoteResponse.Bids)				//Store the list of Bids
// 	BidType := params.Action
// 	amount,_ := strconv.ParseFloat(params.Amount,64)
// 	finalPrice := 0.0
// 	finalSize := 0.0

// 	if (strings.ToLower(BidType) == zendeskUtil.Buy) {			//If Action type is Buy
// 		tempAmount := amount
// 		n := len(AskQuotes)
// 		i := 0

// 		//Find first crossover point means when Asks cross bids value
// 		for BidsQuotes[i].Price > AskQuotes[i].Price {
// 			if i < n {
// 				i++
// 			}else {
// 				break
// 			}
// 		}

// 		//Aggregated weighted sum to reach the exact amount given in API Request
// 		for i < n {
// 			size,_ := strconv.ParseFloat(AskQuotes[i].Size,64)
// 			price,_ := strconv.ParseFloat(AskQuotes[i].Price,64)

// 			//Flag is true whene exact currency pair matches otherwise will use the reverse of it.
// 			if flag == true {
// 				if tempAmount > size {
// 					finalPrice += size*price
// 					tempAmount -= size
// 				} else {
// 					finalPrice += tempAmount*price
// 					tempAmount = 0.0
// 					break;
// 				}
// 			}else {
// 				if tempAmount > price*size {
// 					finalSize += size
// 					tempAmount -= price*size
// 				}else {
// 					finalSize += tempAmount/price
// 					tempAmount = 0.0
// 					break
// 				}
// 			}
// 			i++;
// 		}
// 		if flag == true {
// 			if tempAmount > 0 {
// 				finalPrice = 0
// 			}else {
// 				finalPrice = finalPrice/amount
// 			}

// 			if finalPrice > 0 {
// 				result.Currency = params.QuoteCurrency
// 				result.Price = strconv.FormatFloat(finalPrice,'f',5,64)
// 				result.Total = strconv.FormatFloat(finalPrice*amount,'f',5,64)
// 				return result,nil
// 			}
// 		}else {
// 			maxSize,_ := strconv.ParseFloat(pMap.BaseMaxSize,64)
// 			minSize,_ := strconv.ParseFloat(pMap.BaseMinSize,64)
// 			if tempAmount > 0 {
// 				return Result{}, fmt.Errorf(zendeskUtil.InsufficientAmountMessage)
// 			}else if finalSize > maxSize {
// 				finalSize = 0
// 				return Result{}, fmt.Errorf(zendeskUtil.AmountLimitExceeded)
// 			}else if finalSize < minSize {
// 				return Result{}, fmt.Errorf("Quote Price should be at least "+ strconv.FormatFloat(minSize,'f',5,64))
// 			}
// 			result.Currency = params.QuoteCurrency
// 			result.Price = strconv.FormatFloat(finalSize/amount,'f',10,64)
// 			result.Total = strconv.FormatFloat(finalSize,'f',5,64)
// 			return result,nil

// 		}
// 		return Result{}, fmt.Errorf(zendeskUtil.InsufficientAmountMessage)

// 	}else if (strings.ToLower(BidType) == zendeskUtil.Sell) {				//If Action type is Sell
// 		tempAmount := amount
// 		n := len(BidsQuotes)
// 		i := 0

// 		//Find the first Ask value which is greater than Bids value
// 		for BidsQuotes[i].Price > AskQuotes[i].Price {
// 			if i < n {
// 				i++
// 			}else {
// 				break
// 			}
// 		}

// 		//Flag is true whene exact currency pair matches otherwise will use the reverse of it.
// 		for i < n {
// 			size,_ := strconv.ParseFloat(BidsQuotes[i].Size,64)
// 			price,_ := strconv.ParseFloat(BidsQuotes[i].Price,64)
// 			if flag == true {
// 				if tempAmount > size {
// 					finalPrice += size*price
// 					tempAmount -= size
// 				} else {
// 					finalPrice += tempAmount*price
// 					tempAmount = 0.0
// 					break;
// 				}
// 			}else {
// 				if tempAmount > price*size {
// 					finalSize += size
// 					tempAmount -= price*size
// 				}else {
// 					finalSize += tempAmount/price
// 					tempAmount = 0.0
// 					break
// 				}
// 			}
// 			i++;
// 		}
// 		if flag == true {
// 			if tempAmount > 0 {
// 				finalPrice = 0
// 			}else {
// 				finalPrice = finalPrice/amount
// 			}

// 			if finalPrice > 0 {
// 				result.Currency = params.QuoteCurrency
// 				result.Price = strconv.FormatFloat(finalPrice,'f',5,64)
// 				result.Total = strconv.FormatFloat(finalPrice*amount,'f',5,64)
// 				return result,nil
// 			}
// 		}else {
// 			maxSize,_ := strconv.ParseFloat(pMap.BaseMaxSize,64)
// 			minSize,_ := strconv.ParseFloat(pMap.BaseMinSize,64)
// 			if tempAmount > 0 {
// 				return Result{}, fmt.Errorf(zendeskUtil.InsufficientAmountMessage)
// 			}else if finalSize > maxSize {
// 				finalSize = 0
// 				return Result{}, fmt.Errorf(zendeskUtil.AmountLimitExceeded)
// 			}else if finalSize < minSize {
// 				return Result{}, fmt.Errorf("Amount recieved should be at least "+ strconv.FormatFloat(minSize,'f',5,64))
// 			}
// 			result.Currency = params.QuoteCurrency
// 			result.Price = strconv.FormatFloat(finalSize/amount,'f',10,64)
// 			result.Total = strconv.FormatFloat(finalSize,'f',5,64)
// 			return result,nil
// 		}
// 	}
// 	return Result{}, fmt.Errorf(zendeskUtil.InsufficientAmountMessage)
// }

// //ParseBids is a function to parse the bids from the API Response
// func ParseBids(bids []interface{}) []Quote {
// 	BidsQuotes := []Quote{}
// 	for _,val := range bids {
// 		temp := Quote{}
// 		switch typedValue := val.(type) {
// 		case []interface{}:
// 			for i,x := range typedValue {
// 				switch typed := x.(type) {
// 				case string:
// 					if i == 0 {
// 						temp.Price = typed
// 					}else {
// 						temp.Size = typed
// 					}
// 				case float64:
// 					temp.Num = typed
// 				default:
// 					fmt.Println(zendeskUtil.IncompatibleDatatypeMessage,typed)
// 				}
// 			}
// 			break
// 		default:
// 			fmt.Println(zendeskUtil.IncompatibleDatatypeMessage,typedValue)
// 		}
// 		BidsQuotes = append(BidsQuotes,temp)
// 	}
// 	return BidsQuotes
// }

// //ParseAsks is a function to parse the Asks slice from the API Response
// func ParseAsks(Asks []interface{}) []Quote {
// 	AskQuotes := []Quote{}
// 	for _,val := range Asks {
// 		temp := Quote{}
// 		switch typedValue := val.(type) {
// 		case []interface{}:
// 			for i,x := range typedValue {
// 				switch typed := x.(type) {
// 				case string:
// 					if i == 0 {
// 						temp.Price = typed
// 					}else {
// 						temp.Size = typed
// 					}
// 				case float64:
// 					temp.Num = typed
// 				default:
// 					fmt.Println(zendeskUtil.IncompatibleDatatypeMessage,typed)
// 				}
// 			}
// 			break
// 		default:
// 			fmt.Println(zendeskUtil.IncompatibleDatatypeMessage,typedValue)
// 		}
// 		AskQuotes = append(AskQuotes,temp)
// 	}
// 	return AskQuotes
// }
